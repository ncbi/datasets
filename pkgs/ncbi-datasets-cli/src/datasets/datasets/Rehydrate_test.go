package datasets

import (
	"net/http"
	"net/url"
	"path/filepath"
	"testing"

	"github.com/spf13/afero"
	"github.com/spf13/afero/mem"
	"github.com/stretchr/testify/require"

	"fmt"
	"io/fs"
	"os/exec"
	"reflect"

	"bou.ke/monkey"
)

func TestMin(t *testing.T) {
	val := min(2, 1)
	require.Equal(t, val, 1)

	val = min(10, 10)
	require.Equal(t, val, 10)

	val = min(-2, -1)
	require.Equal(t, val, -2)
}

func TestPatchUrlNoProxy(t *testing.T) {
	val := patchURL("urlstring", nil)
	require.Equal(t, val, "urlstring")

	val = patchURL("", nil)
	require.Equal(t, val, "")

	val = patchURL("foo", nil)
	require.Equal(t, val, "foo")

	val = patchURL("api.ncbi.nlm.nih.gov", nil)
	require.Equal(t, val, "api.ncbi.nlm.nih.gov")
}

func TestPatchUrlProxy(t *testing.T) {
	proxy, err := url.Parse("http://myproxy/x")
	require.NoError(t, err)

	val := patchURL("", proxy)
	require.Equal(t, val, "")

	val = patchURL("foo", proxy)
	require.Equal(t, val, "foo")

	val = patchURL("http://api.ncbi.nlm.nih.gov/y", proxy)
	require.Equal(t, val, "http://myproxy/y")

	val = patchURL("https://api.ncbi.nlm.nih.gov/datasets/foo", proxy)
	require.Equal(t, val, "http://myproxy/foo")
}

func TestPatchUrlDatasetsProxy(t *testing.T) {
	proxy, err := url.Parse("http://myproxy/datasets/x")
	require.NoError(t, err)

	val := patchURL("", proxy)
	require.Equal(t, val, "")

	val = patchURL("foo", proxy)
	require.Equal(t, val, "foo")

	val = patchURL("http://api.ncbi.nlm.nih.gov/y", proxy)
	require.Equal(t, val, "http://myproxy/y")

	val = patchURL("https://api.ncbi.nlm.nih.gov/datasets/foo", proxy)
	require.Equal(t, val, "http://myproxy/datasets/foo")
}

func TestFileRetrieval(t *testing.T) {
	initTestVars()

	argNoProgress = true
	err := downloadMultipleFiles(test_fl)
	require.NoError(t, err)

	for _, test_file := range test_fl {
		verifyFile(t, test_file.path)
	}
}

func TestMissingFileRetrieval(t *testing.T) {
	initTestVars()

	setHttpError(http.StatusNotFound, 0, 1000)
	err := downloadMultipleFiles(test_fl)
	require.Error(t, err)

	exists, err := afero.Exists(afs, test_fl[0].path)
	require.NoError(t, err)
	require.False(t, exists)
}

func TestFileRetrievalReadOnlyError(t *testing.T) {
	initTestVars()
	afs = afero.NewReadOnlyFs(afs)

	err := downloadMultipleFiles(test_fl)
	require.Error(t, err)
	// ReadOnlyFs returns 'operation not permitted'
}

func TestOutOfDiskSpace(t *testing.T) {
	err := doCrasher(t, func() error {
		initTestVars()

		var d *mem.File
		monkey.PatchInstanceMethod(reflect.TypeOf(d), "Write",
			func(c *mem.File, p []byte) (n int, err error) {
				var perr fs.PathError
				perr.Op = "write"
				perr.Path = test_fl[0].path
				perr.Err = fmt.Errorf("out of disk space")
				return 0, &perr
			})
		defer monkey.UnpatchInstanceMethod(reflect.TypeOf(d), "Write")

		err := downloadMultipleFiles(test_fl)
		require.Error(t, err)
		return err
	})

	if e, ok := err.(*exec.ExitError); ok && !e.Success() {
		return
	}
	t.Fatalf("process ran with err %v", err)
}

func TestRehydrateRetryTooManyRequests(t *testing.T) {
	initTestVars()

	const expectedFailures = 2
	setHttpError(http.StatusTooManyRequests, 0, expectedFailures)
	err := downloadMultipleFiles(test_fl)
	require.NoError(t, err)
	require.Equal(t, len(test_fl)+expectedFailures, test_attempted)
	for _, test_file := range test_fl {
		verifyFile(t, test_file.path)
	}
	verifyRequestHeaders(t)
}

func TestRehydrateRetryInternalServerError(t *testing.T) {
	initTestVars()

	const expectedFailures = 2
	setHttpError(http.StatusInternalServerError, 0, expectedFailures)
	err := downloadMultipleFiles(test_fl)
	require.NoError(t, err)
	require.Equal(t, len(test_fl)+expectedFailures, test_attempted)
	for _, test_file := range test_fl {
		verifyFile(t, test_file.path)
	}
	verifyRequestHeaders(t)
}

func TestRehydrateFileSkipDownloaded(t *testing.T) {
	initTestVars()
	// set up one file on disk
	var fileDir = filepath.Dir(test_fl[0].path)
	afs.MkdirAll(fileDir, 0755)
	out, _ := afs.Create(test_fl[0].path)
	out.Close()
	exists, _ := afero.Exists(afs, test_fl[0].path)
	require.True(t, exists)

	err := downloadMultipleFiles(test_fl[0:1])
	require.NoError(t, err)
	require.Equal(t, 0, test_attempted)

	err = downloadMultipleFiles(test_fl)
	require.NoError(t, err)
	require.Equal(t, len(test_fl)-1, test_attempted)
}

func TestRehydrateFileDownloadErrorDelete(t *testing.T) {
	initTestVars()
	setHttpError(http.StatusInternalServerError, 0, 20)
	test_file := test_fl[0:1]
	downloadMultipleFiles(test_file)
	exists, err := afero.Exists(afs, test_file[0].path)
	require.NoError(t, err)
	require.False(t, exists)
}
