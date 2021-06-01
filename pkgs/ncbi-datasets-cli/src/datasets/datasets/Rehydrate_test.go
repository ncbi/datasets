package datasets

import (
	"net/url"
	"testing"

	"github.com/stretchr/testify/require"
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
