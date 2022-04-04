package datasets

import (
	"bufio"
	"bytes"
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

	cleanhttp "github.com/hashicorp/go-cleanhttp"
	retry_http "github.com/hashicorp/go-retryablehttp"
	"github.com/spf13/cobra"
	flag "github.com/spf13/pflag"

	openapi "datasets_cli/v1/openapi"
	datasets_util "datasets_cli/v1/util"
	datasets_command "datasets_cli/v1/util/command"
)

var (
	// AppVersion is the application version string whose value is set at build time
	AppVersion = "undefined"

	serverVersion    string
	versionMessage   string
	userMessage      string
	clientHeaders    = make(map[string]string)
	jsonAcceptHeader = map[string]string{"Accept": "application/json"}
	ncbi_phid        string

	// various command-line args
	argDebug           bool
	argSynMon          bool
	argApiKey          string
	argProxyURL        string
	argRefseqOnly      bool
	argReferenceOnly   bool
	argAssmAccsOnly    bool
	argLimit           string
	argInputFile       string
	argIDArgs          []string
	argJsonLinesFormat bool
	argJsonFormat      bool
	argTaxExactMatch   bool
	argTaxon           string
	argNoProgress      bool

	// Default retry configuration
	defaultRetryWaitMin = 1 * time.Second
	defaultRetryWaitMax = 30 * time.Second

	maxNumRetries = uint8(10)

	// defaultLogger is the logger provided with defaultClient
	defaultLogger = log.New(io.Discard, "", log.LstdFlags)

	request_count    uint64
	request_count_mu sync.Mutex
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "datasets",
	Short: "NCBI Datasets",
	Long: `datasets is a command-line tool that is used to query and download biological sequence data
across all domains of life from NCBI databases.

Refer to NCBI's [command line quickstart](https://www.ncbi.nlm.nih.gov/datasets/docs/quickstarts/command-line-tools/) documentation for information about getting started with the command-line tools.`,
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
	printResultsNoNewline(Payload)
	fmt.Print("\n")
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
	Use:   "version [flags]",
	Short: "print the version of this client and exit",
	Long:  "Print the version of this client and exit.",
	Args:  cobra.MaximumNArgs(0),
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Println(AppVersion)
		cli, err := createOAClient()
		if err != nil {
			return err
		}
		_, resp, err := cli.VersionApi.Version(nil).Execute()
		return handleHTTPResponse(resp, err)
	},
}

func CustomRequestLogHook(_ retry_http.Logger, request *http.Request, _ int) {

	if len(request.URL.RawQuery) == 0 {
		request.URL.RawQuery = "api_key=" + argApiKey
	} else if strings.Index(request.URL.RawQuery, "api_key") == -1 {
		request.URL.RawQuery = request.URL.RawQuery + "&api_key=" + argApiKey
	}
}

func newRetryHttpClient(numRetries uint8) *http.Client {

	retryClient := retry_http.Client{
		HTTPClient:   cleanhttp.DefaultPooledClient(),
		Logger:       defaultLogger,
		RetryWaitMin: defaultRetryWaitMin,
		RetryWaitMax: defaultRetryWaitMax,
		RetryMax:     int(numRetries),
		CheckRetry:   retry_http.DefaultRetryPolicy,
		Backoff:      retry_http.DefaultBackoff,
	}
	if len(argApiKey) > 0 {
		retryClient.RequestLogHook = CustomRequestLogHook
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

	if argSynMon {
		cfg.UserAgent = "datasets/monitoring/datasets_cli"
	}

	err = updateOATransportConfig(cfg, argProxyURL)
	if err != nil {
		return
	}
	cfg.Debug = argDebug
	cli = openapi.NewAPIClient(cfg)
	return
}

func updateOATransportConfig(cfg *openapi.Configuration, proxyURL string) (err error) {
	if proxyURL == "" {
		return
	}

	configs := openapi.ServerConfigurations{
		{
			URL: proxyURL,
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
	for key, value := range resp.Header {
		switch key {
		case "X-Datasets-Version":
			serverVersion = value[0]
		case "X-Datasets-Version-Message":
			versionMessage = value[0]
		case "X-Datasets-User-Message":
			userMessage = value[0]
		}
	}
	return
}

func strToInt32List(strs []string) (geneInts []int32) {
	geneInts, _ = strToInt32ListErr(strs)
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
	geneInts, _ = strToInt64ListErr(strs)
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

func readLines(fp io.Reader) []string {
	var lines []string
	scanner := bufio.NewScanner(fp)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines
}

func registerHiddenStringSlicePair(flags *flag.FlagSet, storage *[]string, longName, shortFlag string, defaultValue []string, description string) {
	hiddenName := "_" + longName
	flags.StringSliceVarP(storage, hiddenName, shortFlag, defaultValue, description)
	flags.MarkHidden(hiddenName)
	flags.StringSliceVar(storage, longName, defaultValue, description)
}

func registerHiddenStringPair(flags *flag.FlagSet, storage *string, longName, shortFlag string, defaultValue string, description string) {
	hiddenName := "_" + longName
	flags.StringVarP(storage, hiddenName, shortFlag, defaultValue, description)
	flags.MarkHidden(hiddenName)
	flags.StringVar(storage, longName, defaultValue, description)
}

func registerHiddenBoolPair(flags *flag.FlagSet, storage *bool, longName, shortFlag string, defaultValue bool, description string) {
	hiddenName := "_" + longName
	flags.BoolVarP(storage, hiddenName, shortFlag, defaultValue, description)
	flags.MarkHidden(hiddenName)
	flags.BoolVar(storage, longName, defaultValue, description)
}

func registerHiddenIntPair(flags *flag.FlagSet, storage *int, longName, shortFlag string, defaultValue int, description string) {
	hiddenName := "_" + longName
	flags.IntVarP(storage, hiddenName, shortFlag, defaultValue, description)
	flags.MarkHidden(hiddenName)
	flags.IntVar(storage, longName, defaultValue, description)
}

// If neither args nor argInputFile specified, no error.
func getArgsFromListOrFile(args []string, argInputFile string) (idArgs []string, err error) {
	var errorMsg bytes.Buffer

	if len(args) != 0 {
		if argInputFile != "" {
			err = fmt.Errorf("Accepts either argument or file, not both")
			return
		}
		if len(args) == 1 {
			idArgs = strings.Split(args[0], ",")
		} else {
			idArgs = args
		}
		return
	}

	if argInputFile == "-" {
		idArgs = readLines(os.Stdin)
	} else if argInputFile != "" {
		fp, fileErr := os.Open(argInputFile)
		if fileErr != nil {
			err = fmt.Errorf("'%s' opening input file: '%s'", fileErr.Error(), argInputFile)
			return
		}
		defer fp.Close()
		idArgs = readLines(fp)
		// Check if any geneIDs read
		if len(idArgs) == 0 {
			fmt.Fprintf(
				&errorMsg,
				"No identifiers read from file: '%s'\n       File should have 1 identifier per row and no spaces or quotes",
				argInputFile,
			)
		}
	}

	if errorMsg.Len() > 0 {
		err = errors.New(errorMsg.String())
	}
	return
}

/*
func GetRuntimeStreamError(err error) openapi.RuntimeStreamError {
	return err.(openapi.GenericOpenAPIError).Model().(openapi.RuntimeStreamError)
}
*/

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
	if resp != nil && resp.StatusCode >= 300 {
		msg := getGatewayRuntimeError(inError)
		err = errors.New("[gateway] " + msg)
	}
	rootCmd.SilenceUsage = true
	return
}

func handleHTTPResponseWithCustomErr(resp *http.Response, inError error, printfTemplate string) (err error) {
	e := datasets_util.HandleHttpResponseWithCustomErr(resp, inError, printfTemplate)
	err = handleHTTPResponseError(resp, e)
	return
}

func handleHTTPResponse(resp *http.Response, inError error) (err error) {
	e := datasets_util.HandleHttpResponse(resp, inError)
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
	if versionMessage != "" {
		fmt.Fprintln(os.Stderr, versionMessage)
	}
	if userMessage != "" {
		fmt.Fprintln(os.Stderr, userMessage)
		exitval = 1
	}
	os.Exit(exitval)
}

func GeneratePHID() string {
	phid := make([]byte, 12)
	rand.Read(phid)
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

	phid := fmt.Sprintf("%s.%d", ncbi_phid, count)
	req.Header.Set("NCBI-PHID", phid)

	if argApiKey != "" {
		req.Header.Set("Api-Key", argApiKey)
	}

	for k, v := range clientHeaders {
		req.Header.Set(k, v)
	}

	return lrt.Proxied.RoundTrip(req)
}

func useEnv(envVarName, argName string) (val string) {
	val = os.Getenv(envVarName)
	return
}

func useEnvBool(envVarName, argName string, defaultVal bool) bool {
	if val, err := strconv.ParseBool(useEnv(envVarName, argName)); err == nil {
		return val
	}
	return defaultVal
}

func init() {
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
	datasets_util.InitRootCommand(rootCmd)

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	rootCmd.PersistentFlags().StringVar(&argApiKey, "api-key", useEnv("NCBI_API_KEY", "api-key"), "NCBI Datasets API Key")

	rootCmd.PersistentFlags().Uint8Var(&maxNumRetries, "max-retries", 10, "Maximum number of retries on API calls, max 255")
	rootCmd.PersistentFlags().MarkHidden("max-retries")

	rootCmd.PersistentFlags().StringVar(&argProxyURL, "proxy", useEnv("http_proxy", "proxy"), "API endpoint proxy")
	rootCmd.PersistentFlags().MarkHidden("proxy")

	rootCmd.PersistentFlags().BoolVar(&argSynMon, "synmon", false, "Mark request as synthetic monitoring")
	rootCmd.PersistentFlags().MarkHidden("synmon")

	rootCmd.PersistentFlags().BoolVar(&argDebug, "debug", useEnvBool("DATASETS_DEBUG", "debug", false), "Emit debugging info")
	rootCmd.PersistentFlags().MarkHidden("debug")

	rootCmd.PersistentFlags().BoolVar(&argNoProgress, "no-progressbar", false, "hide progress bar")

	// todo: hide help flag
	// rootCmd.DebugFlags()  - this has no help flag. cannot find out how to modify its behavior

	cobra.EnableCommandSorting = false

	// add top-level commands
	rootCmd.AddCommand(summaryCmd)
	rootCmd.AddCommand(downloadCmd)
	rootCmd.AddCommand(getAssemblyDescriptorsCmd)
	rootCmd.AddCommand(getGeneDescriptorsCmd)
	rootCmd.AddCommand(rehydrateCmd)

	rootCmd.AddCommand(datasets_command.NewAutocompleteCmd(rootCmd))
	rootCmd.AddCommand(versionCmd)

	clientHeaders["X-Datasets-Client"] = "datasets-cli"
	clientHeaders["X-Datasets-Client-OS"] = runtime.GOOS
	clientHeaders["X-Datasets-Client-Arch"] = runtime.GOARCH
	clientHeaders["X-Datasets-Client-Version"] = AppVersion
	clientHeaders["X-Datasets-Client-Cmd"] = strings.Join(os.Args[1:], " ")

	ncbi_sid := os.Getenv("HTTP_NCBI_SID")
	if ncbi_sid != "" {
		clientHeaders["NCBI-SID"] = ncbi_sid
	}

	if ncbi_phid = os.Getenv("HTTP_NCBI_PHID"); ncbi_phid == "" {
		ncbi_phid = GeneratePHID()
	}

	l5d_dtab := os.Getenv("L5D_DTAB")
	if l5d_dtab != "" {
		clientHeaders["L5D_DTAB"] = l5d_dtab
	}
}
