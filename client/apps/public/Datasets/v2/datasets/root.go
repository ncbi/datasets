package datasets

import (
	"context"
	"crypto/rand"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"runtime"
	"strconv"
	"strings"
	"sync"
	"time"

	cmdflags "datasets_cli/v2/datasets/flags"
	datasets_util "datasets/util"
	datasets_command "datasets/util/command"

	"github.com/gosuri/uiprogress"
	cleanhttp "github.com/hashicorp/go-cleanhttp"
	retry_http "github.com/hashicorp/go-retryablehttp"
	"github.com/spf13/cobra"

	openapi "datasets/openapi/v2"
)

var (
	// AppVersion is the application version string whose value is set at build time
	AppVersion = "undefined"

	versionMessage string
	userMessage    string
	clientHeaders  = make(map[string]string)

	// High level command line arguments
	argSynMon     bool
	argApiKey     string
	argGatewayURL string
	argNoProgress bool
	// argsVersion        bool

	// Default retry configuration
	defaultRetryWaitMin = 1 * time.Second
	defaultRetryWaitMax = 30 * time.Second

	maxNumRetries = uint8(10)

	// defaultLogger is the logger provided with defaultClient
	defaultLogger = log.New(io.Discard, "", log.LstdFlags)

	request_count    uint64
	request_count_mu sync.Mutex
	progress         *uiprogress.Progress
)

const dateFormat = "MM/DD/YYYY"

// TODO: give these loose global flag vars extensible structuring
var globalDebugFlag = cmdflags.NewDebugFlag()

func versionRunE(cmd *cobra.Command, args []string) error {
	fmt.Println("datasets version:", AppVersion)
	cli, err := createOAClient()
	if err != nil {
		return err
	}
	_, resp, err := cli.VersionAPI.Version(context.TODO()).Execute()
	if err == nil {
		return checkResponseHeaders(resp)
	}
	if resp != nil && resp.StatusCode >= 300 {
		msg := getGatewayRuntimeError(err)
		if msg != "" {
			err = errors.New("[gateway] " + msg)
		} else {
			err = errors.New(resp.Status)
		}
	}
	return err
}

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "datasets",
	Short: "NCBI Datasets",
	Long: `datasets is a command-line tool that is used to query and download biological sequence data
across all domains of life from NCBI databases.

Refer to NCBI's [download and install](https://www.ncbi.nlm.nih.gov/datasets/docs/v2/download-and-install/) documentation for information about getting started with the command-line tools.`,
	Args: cobra.MaximumNArgs(0),
}

// ReturnTopLevelCobraCommand returns the top-level cobra command, used for documenting the CLI
func ReturnTopLevelCobraCommand() *cobra.Command {
	return rootCmd
}

func minOf(vars ...int) int {
	min := vars[0]

	for _, i := range vars {
		if min > i {
			min = i
		}
	}

	return min
}

// helper for converting interface{} obj to json strings
func objToJSON(obj interface{}) (string, error) {
	jsonBuf, err := json.Marshal(obj)
	if err != nil {
		return "", err
	}
	return string(jsonBuf), err
}

func printResults(Payload interface{}) {
	json, err := objToJSON(Payload)
	if err != nil {
		rootCmd.PrintErr("\nError converting payload to json\n")
		return
	}
	fmt.Printf("%s\n", json)
}

func printResultsNoNewline(Payload interface{}) {
	json, err := objToJSON(Payload)
	if err != nil {
		rootCmd.PrintErr("\nError converting payload to json\n")
		return
	}
	fmt.Printf("%s", json)
}

// versionCmd represents the version command
var versionCmd = &cobra.Command{
	Use:    "version [flags]",
	Short:  "Print the version of this client and exit",
	Long:   "Print the version of this client and exit.",
	Hidden: true,
	Args:   cobra.MaximumNArgs(0),
	RunE:   versionRunE,
}

// DefaultRetryPolicy provides a default callback for Client.CheckRetry, which
// will retry on connection errors and server errors.
func DefaultRetryPolicy(ctx context.Context, resp *http.Response, err error) (bool, error) {
	// do not retry on context.Canceled or context.DeadlineExceeded
	if ctx.Err() != nil {
		return false, ctx.Err()
	}

	// don't propagate other errors
	shouldRetry, throw_away_error := baseRetryPolicy(resp, err)
	if throw_away_error != nil {
		defaultLogger.Println("Error while retrying HTTP request: ", throw_away_error)
	}
	return shouldRetry, nil
}

func baseRetryPolicy(resp *http.Response, err error) (bool, error) {
	if resp == nil {
		return false, err
	}
	// if err != nil {
	// 	if v, ok := err.(*url.Error); ok {
	// 		// Don't retry if the error was due to too many redirects.
	// 		if redirectsErrorRe.MatchString(v.Error()) {
	// 			return false, v
	// 		}

	// 		// Don't retry if the error was due to an invalid protocol scheme.
	// 		if schemeErrorRe.MatchString(v.Error()) {
	// 			return false, v
	// 		}

	// 		// Don't retry if the error was due to TLS cert verification failure.
	// 		if _, ok := v.Err.(x509.UnknownAuthorityError); ok {
	// 			return false, v
	// 		}
	// 	}

	// 	// The error is likely recoverable so retry.
	// 	return true, nil
	// }

	// 429 Too Many Requests is recoverable. Sometimes the server puts
	// a Retry-After response header to indicate when the server is
	// available to start processing request from client.
	if resp.StatusCode == http.StatusTooManyRequests {
		return true, nil
	}

	// Check the response code. We retry on 500-range responses to allow
	// the server time to recover, as 500's are typically not permanent
	// errors and may relate to outages on the server side. This will catch
	// invalid response codes as well, like 0 and 999.
	if resp.StatusCode == 0 || (resp.StatusCode >= 500 && resp.StatusCode != 501) {
		return true, fmt.Errorf("unexpected HTTP status %s", resp.Status)
	}

	return false, nil
}

func newRetryHttpClient(numRetries uint8) *http.Client {
	retryClient := retry_http.Client{
		HTTPClient:   cleanhttp.DefaultPooledClient(),
		Logger:       defaultLogger,
		RetryWaitMin: defaultRetryWaitMin,
		RetryWaitMax: defaultRetryWaitMax,
		RetryMax:     int(numRetries),
		CheckRetry:   DefaultRetryPolicy,
		Backoff:      retry_http.DefaultBackoff,
	}
	return retryClient.StandardClient()
}

func initRetryableClient() *http.Client {
	c := newRetryHttpClient(maxNumRetries)
	t := c.Transport
	c.Transport = LoggingRoundTripper{Proxied: t}

	return c
}

func createOAClient() (cli *openapi.APIClient, err error) {
	cfg := openapi.NewConfiguration()

	cfg.HTTPClient = initRetryableClient()

	for k, v := range clientHeaders {
		cfg.AddDefaultHeader(k, v)
	}

	if argSynMon {
		cfg.UserAgent = "datasets/monitoring/datasets_cli"
	}

	if err = updateOATransportConfig(cfg, argGatewayURL); err != nil {
		return
	}
	if globalDebugFlag.Debug() {
		cfg.Debug = true
	}
	cli = openapi.NewAPIClient(cfg)
	return
}

func updateOATransportConfig(cfg *openapi.Configuration, gatewayURL string) (err error) {
	if gatewayURL == "" {
		return
	}

	configs := openapi.ServerConfigurations{
		{
			URL: gatewayURL,
		},
	}
	cfg.Servers = configs
	return
}

func checkResponseHeaders(resp *http.Response) (err error) {
	if resp == nil {
		err = errors.New("Bad response")
		return
	}

	// fmt.Println("Num headers: " + strconv.Itoa(len(resp.Header)))
	// Validation (checkmarx)
	if len(resp.Header) > 256 {
		err = errors.New("Number of headers exceeded maximum count: " + strconv.Itoa(len(resp.Header)))
		return
	}

	for key, value := range resp.Header {
		switch key {
		case "X-Datasets-Version-Message":
			if versionMessage == "" {
				fmt.Fprintln(os.Stderr, value[0])
				versionMessage = value[0]
			}

		case "X-Datasets-User-Message":
			if userMessage == "" {
				fmt.Fprintln(os.Stderr, value[0])
				userMessage = value[0]
			}
		}
	}
	return
}

func strToInt32List(strs []string) (geneInts []int32) {
	var err error = nil
	if geneInts, err = strToInt32ListErr(strs); err != nil {
		defaultLogger.Println("Failure to coerce input to integer: ", err)
	}
	return
}

func strToInt32ListErr(strs []string) (geneInts []int32, err error) {
	hasError := false
	for _, idFullStr := range strs {
		for _, idStr := range strings.Split(idFullStr, ",") {
			geneInt64, e := strconv.ParseInt(idStr, 10, 32)
			geneInt32 := int32(geneInt64)
			if e != nil {
				hasError = true
			} else {
				geneInts = append(geneInts, geneInt32)
			}
		}
	}
	if hasError {
		err = errors.New("unable to parse integer value")
	}
	return
}

func strToInt64List(strs []string) (geneInts []int64) {
	var err error = nil
	if geneInts, err = strToInt64ListErr(strs); err != nil {
		defaultLogger.Println("Failure to coerce input to integer: ", err)
	}
	return
}

func strToInt64ListErr(strs []string) (geneInts []int64, err error) {
	hasError := false
	for _, idFullStr := range strs {
		for _, idStr := range strings.Split(idFullStr, ",") {
			geneInt, e := strconv.ParseInt(idStr, 10, 64)
			if e != nil {
				hasError = true
			} else {
				geneInts = append(geneInts, geneInt)
			}
		}
	}
	if hasError {
		err = errors.New("unable to parse integer value")
	}
	return
}

func ParentCommandRunE(*cobra.Command, []string) error {
	return fmt.Errorf("Continue with one of the sub-commands")
}

func getGatewayRuntimeError(err error) string {
	if openAPIErr, ok := err.(openapi.GenericOpenAPIError); ok {
		model := openAPIErr.Model()
		status := model.(openapi.RpcStatus)
		return status.GetMessage()
	}
	return err.Error()
}

func handleHTTPResponseError(resp *http.Response, inError error) (err error) {
	if inError == nil {
		err = checkResponseHeaders(resp)
		return
	}
	rootCmd.SilenceUsage = true
	msg := getGatewayRuntimeError(inError)
	if msg != "" {
		err = errors.New("[gateway] " + msg)
		return
	}

	if resp != nil && resp.StatusCode >= 300 {
		err = errors.New(resp.Status)
		return
	}

	err = inError
	return
}

func handleHTTPResponseWithCustomErr(resp *http.Response, inError error, printfTemplate string) (err error) {
	e := CreateErrorMessageFromMessageOrError(resp, inError, printfTemplate)
	err = handleHTTPResponseError(resp, e)
	return
}

func handleHTTPResponse(resp *http.Response, inError error) (err error) {
	e := CreateErrorMessageFromMessageOrError(resp, inError, "Gateway Error (%s)")
	err = handleHTTPResponseError(resp, e)
	return
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	exitval := 0
	err := rootCmd.Execute()
	if err != nil {
		exitval = 1
	}
	if userMessage != "" {
		exitval = 1
	}
	os.Exit(exitval)
}

func GeneratePHID() string {
	phid := make([]byte, 12)
	if _, err := rand.Read(phid); err != nil {
		defaultLogger.Fatalln("Failure to create phid")
	}
	var list []string
	for _, x := range phid {
		list = append(list, fmt.Sprintf("%02X", x))
	}
	result := strings.Join(list, "")

	return result
}

// This type implements the http.RoundTripper interface
type LoggingRoundTripper struct {
	Proxied http.RoundTripper
}

func (lrt LoggingRoundTripper) RoundTrip(req *http.Request) (*http.Response, error) {
	// Safely increment the request_counter
	request_count_mu.Lock()
	request_count++
	count := request_count
	request_count_mu.Unlock()

	phid := fmt.Sprintf("%s.%d", req.Header.Get("NCBI-PHID"), count)
	req.Header.Set("NCBI-PHID", phid)
	return lrt.Proxied.RoundTrip(req)
}

func useEnv(envVarName, argName string) (val string) {
	val = os.Getenv(envVarName)
	return
}

func setVersionRecursively(cmd *cobra.Command) {
	cmd.Version = AppVersion
	for _, c := range cmd.Commands() {
		setVersionRecursively(c)
	}
}

func init() {
	progress = uiprogress.New()
	progress.SetOut(os.Stderr)

	datasets_util.AddUsageSections(
		"datasets",
		&datasets_util.UsageSections{
			&datasets_util.UsageSection{
				SectionText: "Data Retrieval Commands",
				Commands:    []string{"summary", "download", "rehydrate"},
			},
			&datasets_util.UsageSection{
				SectionText: "Miscellaneous Commands",
				Commands:    []string{"completion", "version", "help"},
			},
		},
	)
	rootCmd.SetUsageTemplate(getUsageTemplate())
	rootCmd.SetHelpTemplate(getHelpTemplate())
	rootCmd.SetVersionTemplate("datasets version: {{.Version}}\n")

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.
	pFlags := rootCmd.PersistentFlags()
	cmdflags.RegisterAllFlags([]cmdflags.FlagInterface{globalDebugFlag}, pFlags)
	pFlags.StringVar(&argApiKey, "api-key", useEnv("NCBI_API_KEY", "api-key"), "Specify an NCBI API key")

	pFlags.Uint8Var(&maxNumRetries, "max-retries", 10, "Maximum number of retries on API calls, max 255")
	if err := pFlags.MarkHidden("max-retries"); err != nil {
		defaultLogger.Fatalln("Invalid attempt to create hidden flag")
	}

	pFlags.StringVar(&argGatewayURL, "gateway-url", useEnv("NCBI_GATEWAY_URL", "gateway-url"), "API endpoint proxy")
	if err := pFlags.MarkHidden("gateway-url"); err != nil {
		defaultLogger.Fatalln("Invalid attempt to create hidden flag")
	}

	pFlags.BoolVar(&argSynMon, "synmon", false, "Mark request as synthetic monitoring")
	if err := pFlags.MarkHidden("synmon"); err != nil {
		defaultLogger.Fatalln("Invalid attempt to create hidden flag")
	}

	pFlags.BoolP("help", "", false, "Print detailed help about a datasets command")
	pFlags.BoolP("version", "", false, "Print version of datasets")

	cobra.EnableCommandSorting = false

	// add top-level commands
	rootCmd.AddCommand(createSummaryCmd())
	rootCmd.AddCommand(downloadCmd)
	rootCmd.AddCommand(rehydrateCmd)

	rootCmd.AddCommand(datasets_command.NewAutocompleteCmd(rootCmd))
	rootCmd.AddCommand(versionCmd)

	rootCmd.SetFlagErrorFunc(cmdflags.ErrorFlagHandler)

	rootCmd.SetHelpCommand(&cobra.Command{
		Use:    "no-help",
		Hidden: true,
	})

	// This has to happen AFTER all subcommands are added
	setVersionRecursively(rootCmd)

	clientHeaders["X-Datasets-Client"] = "datasets-cli"
	clientHeaders["X-Datasets-Client-OS"] = runtime.GOOS
	clientHeaders["X-Datasets-Client-Arch"] = runtime.GOARCH
	clientHeaders["X-Datasets-Client-Version"] = AppVersion
	clientHeaders["X-Datasets-Client-Cmd"] = strings.Join(os.Args[1:], " ")

	ncbi_sid := os.Getenv("HTTP_NCBI_SID")
	if ncbi_sid != "" {
		clientHeaders["NCBI-SID"] = ncbi_sid
	}

	ncbi_phid := os.Getenv("HTTP_NCBI_PHID")
	if ncbi_phid != "" {
		clientHeaders["NCBI-PHID"] = ncbi_phid
	} else {
		clientHeaders["NCBI-PHID"] = GeneratePHID()
	}

	if len(argApiKey) > 0 {
		clientHeaders["api-key"] = argApiKey
	}

	l5d_dtab := os.Getenv("L5D_DTAB")
	if l5d_dtab != "" {
		clientHeaders["L5D_DTAB"] = l5d_dtab
	}

	cmdflags.RetrieveTaxIdForTaxon = RetrieveTaxIdForTaxon
}

func getUsageTemplate() string {
	return `
{{- if .HasAvailableSubCommands}}
  {{- "\n  "}}{{.CommandPath}} [command]
{{- end}}
{{- if gt (len .Aliases) 0}}
{{- "\n\n"}}Aliases
  {{- "\n  "}}{{.NameAndAliases}}
{{- end}}
{{- $cmd := .}}
{{- if .HasAvailableSubCommands}}
  {{- range (usageSections $cmd) -}}
    {{- $section := .}}
    {{- "\n\n"}}{{$section.SectionText}}
    {{- range (usageCommands $cmd.Commands $section) -}}
      {{- if (or .IsAvailableCommand (eq .Name "help")) -}}
	    {{- "\n  " }}{{rpad .Name .NamePadding}} {{.Short}}
      {{- end}}
    {{- end}}
  {{- end}}
{{- end}}
Use {{.CommandPath}} <command> --help for detailed help about a command.
`
}

func getHelpTemplate() string {
	return `
{{- .Long }}

Usage
{{- if .Runnable}}
  {{- "\n  "}}{{.UseLine}}
{{- end}}
{{- if .HasAvailableSubCommands}}
  {{- "\n  "}}{{.CommandPath}} [command]
{{- end}}
{{- if gt (len .Aliases) 0}}
{{- "\n\n"}}Aliases
  {{- "\n  "}}{{.NameAndAliases}}
{{- end}}
{{- if .HasExample -}}
  {{- "\n\n"}}Sample Commands
{{.Example}}
{{- end}}
{{- $cmd := .}}
{{- if .HasAvailableSubCommands}}
  {{- range (usageSections $cmd) -}}
    {{- $section := .}}
    {{- "\n\n"}}{{$section.SectionText}}
    {{- range (usageCommands $cmd.Commands $section) -}}
      {{- if (or .IsAvailableCommand (eq .Name "help")) -}}
	    {{- "\n  " }}{{rpad .Name .NamePadding}} {{.Short}}
      {{- end}}
    {{- end}}
  {{- end}}
{{- end}}
{{- if .HasAvailableLocalFlags}}
{{- "\n\n"}}Flags
  {{- "\n"}}{{expandFields $cmd .LocalFlags.FlagUsages }}
{{- end}}
{{- if .HasAvailableInheritedFlags}}
  {{- "\n\n"}}Global Flags
  {{- "\n"}}{{.InheritedFlags.FlagUsages}}
{{- end}}

{{- if .HasHelpSubCommands}}

Additional help topics:
  {{- range .Commands -}}
    {{- if .IsAdditionalHelpTopicCommand -}}
	  {{- "\n  "}}{{rpad .CommandPath .CommandPathPadding}} {{.Short}}
	{{- end}}
  {{- end}}
{{- end}}

{{- if .HasAvailableSubCommands}}
Use {{.CommandPath}} <command> --help for detailed help about a command.
{{- end}}
`
}
