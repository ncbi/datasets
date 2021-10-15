package datasets

import (
	"github.com/stretchr/testify/require"
	"net/http"
	"testing"
)

func TestSummaryRetryTooManyRequests(t *testing.T) {

	initTestVars()
	const expectedFailures = 2
	setHttpError(http.StatusTooManyRequests, 0, expectedFailures)

	argIDArgs = []string{"2"}

	err := cmdRunSummaryGeneID(summaryGeneIDCmd, argIDArgs)
	require.NoError(t, err)
	require.Equal(t, expectedFailures+1, test_attempted)
}
