package datasets

import (
	"bytes"
	"errors"
	"fmt"
	"net/http"

	openapi "datasets/openapi/v2"
)

func CreateErrorMessageFromMessageOrError(resp *http.Response, inError error, printfTemplate string) (err error) {
	if inError != nil {
		err = inError
	} else if resp == nil {
		err = errors.New("Gateway Error")
	} else if resp.StatusCode >= 300 {
		err = fmt.Errorf(printfTemplate, resp.Status)
	}
	return
}

func MessagesToError(messages []openapi.V2reportsMessage) (err error) {
	if len(messages) > 0 {
		var errorMsg bytes.Buffer
		MessagesToErrorBuf(messages, &errorMsg)
		if errorMsg.Len() > 0 {
			err = errors.New(errorMsg.String())
		}
	}
	return
}

func MessagesToErrorBuf(messages []openapi.V2reportsMessage, errors *bytes.Buffer) {
	if len(messages) > 0 {
		for _, message := range messages {
			fmt.Fprintf(errors, "%s\n", message.Error.GetMessage())
		}
	}
}
