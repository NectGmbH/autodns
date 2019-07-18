// Code generated by go-swagger; DO NOT EDIT.

package certificate_tasks

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/NectGmbH/autodns/models"
)

// CertificateCreateReader is a Reader for the CertificateCreate structure.
type CertificateCreateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CertificateCreateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewCertificateCreateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewCertificateCreateOK creates a CertificateCreateOK with default headers values
func NewCertificateCreateOK() *CertificateCreateOK {
	return &CertificateCreateOK{}
}

/*CertificateCreateOK handles this case with default header values.

successful operation
*/
type CertificateCreateOK struct {
	Payload *models.JSONResponseDataObjectJob
}

func (o *CertificateCreateOK) Error() string {
	return fmt.Sprintf("[POST /certificate][%d] certificateCreateOK  %+v", 200, o.Payload)
}

func (o *CertificateCreateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.JSONResponseDataObjectJob)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}