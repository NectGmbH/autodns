// Code generated by go-swagger; DO NOT EDIT.

package domain_tasks

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/NectGmbH/autodns/models"
)

// DomainServicesUpdateReader is a Reader for the DomainServicesUpdate structure.
type DomainServicesUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DomainServicesUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewDomainServicesUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewDomainServicesUpdateOK creates a DomainServicesUpdateOK with default headers values
func NewDomainServicesUpdateOK() *DomainServicesUpdateOK {
	return &DomainServicesUpdateOK{}
}

/*DomainServicesUpdateOK handles this case with default header values.

successful operation
*/
type DomainServicesUpdateOK struct {
	Payload *models.JSONResponseDataJSONNoData
}

func (o *DomainServicesUpdateOK) Error() string {
	return fmt.Sprintf("[PUT /domain/_services][%d] domainServicesUpdateOK  %+v", 200, o.Payload)
}

func (o *DomainServicesUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.JSONResponseDataJSONNoData)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
