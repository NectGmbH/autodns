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

// CertificateReissueReader is a Reader for the CertificateReissue structure.
type CertificateReissueReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CertificateReissueReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewCertificateReissueOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewCertificateReissueOK creates a CertificateReissueOK with default headers values
func NewCertificateReissueOK() *CertificateReissueOK {
	return &CertificateReissueOK{}
}

/*CertificateReissueOK handles this case with default header values.

successful operation
*/
type CertificateReissueOK struct {
	Payload *models.JSONResponseDataObjectJob
}

func (o *CertificateReissueOK) Error() string {
	return fmt.Sprintf("[PUT /certificate/{id}][%d] certificateReissueOK  %+v", 200, o.Payload)
}

func (o *CertificateReissueOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.JSONResponseDataObjectJob)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
