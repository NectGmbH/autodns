// Code generated by go-swagger; DO NOT EDIT.

package mail_proxy_tasks

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/NectGmbH/autodns/models"
)

// MailProxyInfoReader is a Reader for the MailProxyInfo structure.
type MailProxyInfoReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *MailProxyInfoReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewMailProxyInfoOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewMailProxyInfoOK creates a MailProxyInfoOK with default headers values
func NewMailProxyInfoOK() *MailProxyInfoOK {
	return &MailProxyInfoOK{}
}

/*MailProxyInfoOK handles this case with default header values.

successful operation
*/
type MailProxyInfoOK struct {
	Payload *models.JSONResponseDataMailProxy
}

func (o *MailProxyInfoOK) Error() string {
	return fmt.Sprintf("[GET /mailProxy/{domain}][%d] mailProxyInfoOK  %+v", 200, o.Payload)
}

func (o *MailProxyInfoOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.JSONResponseDataMailProxy)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
