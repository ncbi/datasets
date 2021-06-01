package datasets

import (
	"strings"
	"testing"

	openapi "main/openapi_client"

	"github.com/stretchr/testify/require"
)

func TestOAConfig(t *testing.T) {
	cfg := openapi.NewConfiguration()
	err := updateOATransportConfig(cfg, "")
	require.NoError(t, err)

	err = updateOATransportConfig(cfg, "abc123")
	require.NoError(t, err)
}

func TestOAClient(t *testing.T) {
	cli, err := createOAClient()
	require.NoError(t, err)
	require.Equal(t, len(cli.GetConfig().DefaultHeader), 5)

	clientHeaders = make(map[string]string)
	cli, err = createOAClient()
	require.NoError(t, err)
	require.Equal(t, len(cli.GetConfig().DefaultHeader), 0)
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
	require.Error(t, err)
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
