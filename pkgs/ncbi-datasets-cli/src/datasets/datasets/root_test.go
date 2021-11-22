package datasets

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"sync"
	"testing"

	openapi "datasets_cli/v1/openapi"

	"github.com/spf13/afero"
	"github.com/stretchr/testify/require"
)

var (
	test_http_server *httptest.Server
	test_http_req    *http.Request
	test_attempted   int
	test_server_lock sync.Mutex
	test_fl          []fetchLine
)

func TestMain(m *testing.M) {
	// Write code here to run before tests
	initTestVars()
	// Start a local HTTP server
	test_http_server = httptest.NewServer(http.HandlerFunc(
		func(rw http.ResponseWriter, req *http.Request) {
			test_min_attempts := 0
			test_max_attempts := 3

			test_server_lock.Lock()
			defer test_server_lock.Unlock()
			test_attempted++
			test_http_req = req
			if test_error, err := strconv.Atoi(req.Header.Get("Error")); err == nil {
				// Don't return error until this many calls, default to 0
				if tme, err := strconv.Atoi(req.Header.Get("Min-Error")); err == nil {
					test_min_attempts = tme
				}
				// Don't return error after this many calls, default to 3
				if tma, err := strconv.Atoi(req.Header.Get("Max-Error")); err == nil {
					test_max_attempts = tma
				}

				if test_attempted > test_min_attempts {
					if test_attempted <= test_max_attempts {
						rw.WriteHeader(test_error)
					}
				}
			}
			rw.Write([]byte("{}"))
		}))

	test_fl = []fetchLine{
		{url: fmt.Sprintf("%s/test", test_http_server.URL), path: "random/dir/test"},
		{url: fmt.Sprintf("%s/test2", test_http_server.URL), path: "random/dir/test2"},
		{url: fmt.Sprintf("%s/test3", test_http_server.URL), path: "random/dir/test3"},
	}
	argProxyURL = test_http_server.URL
	argApiKey = "ABC123"
	argNoProgress = true
	//argDebug = true

	// Run tests
	exitVal := m.Run()

	// Write code here to run after tests

	// Exit with exit value from tests
	os.Exit(exitVal)
}

func initTestVars() {
	afs = afero.NewMemMapFs()
	test_attempted = 0
	test_http_req = nil
	delete(clientHeaders, "Error")
	delete(clientHeaders, "Max-Error")
	delete(clientHeaders, "Min-Error")
}

func setHttpError(http_status int, min_count int, max_count int) {
	clientHeaders["Error"] = strconv.Itoa(http_status)
	clientHeaders["Max-Error"] = strconv.Itoa(max_count)
	clientHeaders["Min-Error"] = strconv.Itoa(min_count)
}

func verifyFile(t *testing.T, path string) {
	exists, err := afero.Exists(afs, path)
	require.NoError(t, err)
	require.True(t, exists)

	empty, err := afero.IsEmpty(afs, path)
	require.NoError(t, err)
	require.False(t, empty)
}

func verifyRequestHeaders(t *testing.T) {
	found_phid := false
	found_apikey := false

	for k, v := range test_http_req.Header {
		if k == "Ncbi-Phid" {
			found_phid = true
			require.Equal(t, len(v), 1)
		}
		if k == "Api-Key" {
			found_apikey = true
			require.Equal(t, len(v), 1)
		}
	}
	require.True(t, found_phid)
	require.True(t, found_apikey)
}

func doCrasher(t *testing.T, c func() error) (err error) {
	if os.Getenv("BE_CRASHER") == "1" {
		c()
		return
	}
	cmd := exec.Command(os.Args[0], fmt.Sprintf("-test.run=%s", t.Name()))
	cmd.Env = append(os.Environ(), "BE_CRASHER=1")
	err = cmd.Run()
	return
}

func TestOAConfig(t *testing.T) {
	cfg := openapi.NewConfiguration()
	err := updateOATransportConfig(cfg, "")
	require.NoError(t, err)

	err = updateOATransportConfig(cfg, "abc123")
	require.NoError(t, err)
}

func TestOAClient(t *testing.T) {
	_, err := createOAClient()
	require.NoError(t, err)

	clientHeaders = make(map[string]string)
	_, err = createOAClient()
	require.NoError(t, err)
}

func TestStrToInt64List(t *testing.T) {
	val := strToInt64List([]string{})
	require.ElementsMatch(t, val, []int64{})

	val = strToInt64List([]string{""})
	require.ElementsMatch(t, val, []int64{})

	val = strToInt64List([]string{"a"})
	require.ElementsMatch(t, val, []int64{})

	val = strToInt64List([]string{"1"})
	require.Exactly(t, val, []int64{1})

	val = strToInt64List([]string{"1", "4", "2"})
	require.Exactly(t, val, []int64{1, 4, 2})

	val = strToInt64List(strings.Split("1e36", ""))
	require.Exactly(t, val, []int64{1, 3, 6})
}

func TestStrToInt64ListErr(t *testing.T) {
	val, err := strToInt64ListErr([]string{})
	require.ElementsMatch(t, val, []int64{})

	val, err = strToInt64ListErr([]string{"1"})
	require.Exactly(t, val, []int64{1})

	val, err = strToInt64ListErr(strings.Split("1e36", ""))
	require.Error(t, err)
	require.Exactly(t, val, []int64{1, 3, 6})
}

func TestArgsFromList(t *testing.T) {
	inputFile := ""

	args, err := getArgsFromListOrFile([]string{}, inputFile)
	require.NoError(t, err)
	require.ElementsMatch(t, args, []string{})

	args, err = getArgsFromListOrFile([]string{"1,2,3"}, inputFile)
	require.NoError(t, err)
	require.Exactly(t, args, []string{"1", "2", "3"})

	args, err = getArgsFromListOrFile(strings.Split("123", ""), inputFile)
	require.NoError(t, err)
	require.Exactly(t, args, []string{"1", "2", "3"})

	args, err = getArgsFromListOrFile([]string{"a1,a2,a3"}, inputFile)
	require.NoError(t, err)
	require.Exactly(t, args, []string{"a1", "a2", "a3"})

	// non-existent file
	inputFile = "foo"
	args, err = getArgsFromListOrFile([]string{}, inputFile)
	require.Error(t, err)
	require.Exactly(t, args, []string(nil))

	// can't supply both args and input file
	inputFile = "foo"
	args, err = getArgsFromListOrFile([]string{"a1,a2,a3"}, inputFile)
	require.Error(t, err)
	require.ElementsMatch(t, args, []string{})
}
