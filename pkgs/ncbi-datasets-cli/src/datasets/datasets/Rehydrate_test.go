package datasets

import (
	"net/url"
	"testing"

	"github.com/spf13/afero"
	"github.com/spf13/afero/mem"
	"github.com/stretchr/testify/require"

	"bou.ke/monkey"
	"fmt"
	"io/fs"
	"os"
	"os/exec"
	"reflect"
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
	afs = afero.NewMemMapFs()

	fl := []fetchLine{
		{url: "https://www.ncbi.nlm.nih.gov/favicon.ico", path: "random/dir/test.ico"},
	}
	argNoProgress = true
	err := downloadMultipleFiles(fl)
	require.NoError(t, err)

	exists, err := afero.Exists(afs, fl[0].path)
	require.NoError(t, err)
	require.True(t, exists)

	empty, err := afero.IsEmpty(afs, fl[0].path)
	require.NoError(t, err)
	require.False(t, empty)
}

func TestMissingFileRetrieval(t *testing.T) {
	afs = afero.NewMemMapFs()

	fl := []fetchLine{
		{url: "https://www.ncbi.nlm.nih.gov/missing.test", path: "random/dir/test.ico"},
	}
	argNoProgress = true
	err := downloadMultipleFiles(fl)
	require.Error(t, err)

	exists, err := afero.Exists(afs, fl[0].path)
	require.NoError(t, err)
	require.False(t, exists)
}

func TestFileRetrievalReadOnlyError(t *testing.T) {
	afs = afero.NewReadOnlyFs(afero.NewMemMapFs())

	fl := []fetchLine{
		{url: "https://www.ncbi.nlm.nih.gov/favicon.ico", path: "random/dir/test.ico"},
	}
	argNoProgress = true
	err := downloadMultipleFiles(fl)
	require.Error(t, err)
	// ReadOnlyFs returns 'operation not permitted'
}

func TestOutOfDiskSpace(t *testing.T) {
	Crasher := func() {
		afs = afero.NewMemMapFs()

		fl := []fetchLine{
			{url: "https://www.ncbi.nlm.nih.gov/favicon.ico", path: "random/dir/test.ico"},
		}

		var d *mem.File
		monkey.PatchInstanceMethod(reflect.TypeOf(d), "Write",
			func(c *mem.File, p []byte) (n int, err error) {
				var perr fs.PathError
				fmt.Printf("%#v\n", c)
				perr.Op = "write"
				perr.Path = fl[0].path
				perr.Err = fmt.Errorf("out of disk space")
				return 0, &perr
			})
		defer monkey.UnpatchInstanceMethod(reflect.TypeOf(d), "Write")

		argNoProgress = true

		err := downloadMultipleFiles(fl)
		require.Error(t, err)
	}

	if os.Getenv("BE_CRASHER") == "1" {
		Crasher()
		return
	}
	cmd := exec.Command(os.Args[0], fmt.Sprintf("-test.run=%s", t.Name()))
	cmd.Env = append(os.Environ(), "BE_CRASHER=1")
	err := cmd.Run()
	if e, ok := err.(*exec.ExitError); ok && !e.Success() {
		return
	}
	t.Fatalf("process ran with err %v, want exit status 255", err)
}
