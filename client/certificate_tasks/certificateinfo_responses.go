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

// CertificateinfoReader is a Reader for the Certificateinfo structure.
type CertificateinfoReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CertificateinfoReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewCertificateinfoOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewCertificateinfoOK creates a CertificateinfoOK with default headers values
func NewCertificateinfoOK() *CertificateinfoOK {
	return &CertificateinfoOK{}
}

/*CertificateinfoOK handles this case with default header values.

successful operation
*/
type CertificateinfoOK struct {
	Payload *models.JSONResponseDataCertificate
}

func (o *CertificateinfoOK) Error() string {
	return fmt.Sprintf("[GET /certificate/{id}][%d] certificateinfoOK  %+v", 200, o.Payload)
}

func (o *CertificateinfoOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.JSONResponseDataCertificate)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
