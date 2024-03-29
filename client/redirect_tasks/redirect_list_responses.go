// Code generated by go-swagger; DO NOT EDIT.

package redirect_tasks

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/NectGmbH/autodns/models"
)

// RedirectListReader is a Reader for the RedirectList structure.
type RedirectListReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *RedirectListReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewRedirectListOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewRedirectListOK creates a RedirectListOK with default headers values
func NewRedirectListOK() *RedirectListOK {
	return &RedirectListOK{}
}

/*RedirectListOK handles this case with default header values.

successful operation
*/
type RedirectListOK struct {
	Payload *models.JSONResponseDataRedirect
}

func (o *RedirectListOK) Error() string {
	return fmt.Sprintf("[POST /redirect/_search][%d] redirectListOK  %+v", 200, o.Payload)
}

func (o *RedirectListOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.JSONResponseDataRedirect)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
