// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/speakeasy-sdks/test-go/v3/pkg/models/shared"
	"net/http"
)

type DeleteUserv1Request struct {
	// UserID
	ID string `pathParam:"style=simple,explode=false,name=id"`
}

func (o *DeleteUserv1Request) GetID() string {
	if o == nil {
		return ""
	}
	return o.ID
}

type DeleteUserv1Response struct {
	// HTTP response content type for this operation
	ContentType string
	// Default error response
	Error *shared.Error
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// OK
	Success *shared.Success
}

func (o *DeleteUserv1Response) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *DeleteUserv1Response) GetError() *shared.Error {
	if o == nil {
		return nil
	}
	return o.Error
}

func (o *DeleteUserv1Response) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *DeleteUserv1Response) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *DeleteUserv1Response) GetSuccess() *shared.Success {
	if o == nil {
		return nil
	}
	return o.Success
}
