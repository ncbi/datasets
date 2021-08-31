package util

import (
	"bytes"
	"errors"
	"fmt"
	"net/http"

	openapi "main/openapi_client"
)

func HandleHttpResponseWithCustomErr(resp *http.Response, inError error, printfTemplate string) (err error) {
	if inError != nil {
		err = inError
	} else if resp == nil {
		err = errors.New("Gateway Error")
	} else if resp.StatusCode >= 300 {
		err = fmt.Errorf(printfTemplate, resp.Status)
	}
	return
}

func HandleHttpResponse(resp *http.Response, inError error) (err error) {
	err = HandleHttpResponseWithCustomErr(resp, inError, "Gateway Error (%s)")
	return
}

func MessagesToError(messages []openapi.V1Message) (err error) {
	if len(messages) > 0 {
		var errorMsg bytes.Buffer
		MessagesToErrorBuf(messages, &errorMsg)
		if errorMsg.Len() > 0 {
			err = errors.New(errorMsg.String())
		}
	}
	return
}

func MessagesToErrorBuf(messages []openapi.V1Message, errors *bytes.Buffer) {
	if len(messages) > 0 {
		for _, message := range messages {
			fmt.Fprintf(errors, "%s\n", message.Error.GetMessage())
		}
	}
}
