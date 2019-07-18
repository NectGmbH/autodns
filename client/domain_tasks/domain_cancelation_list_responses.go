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

// DomainCancelationListReader is a Reader for the DomainCancelationList structure.
type DomainCancelationListReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DomainCancelationListReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewDomainCancelationListOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewDomainCancelationListOK creates a DomainCancelationListOK with default headers values
func NewDomainCancelationListOK() *DomainCancelationListOK {
	return &DomainCancelationListOK{}
}

/*DomainCancelationListOK handles this case with default header values.

successful operation
*/
type DomainCancelationListOK struct {
	Payload *models.JSONResponseDataDomainCancelation
}

func (o *DomainCancelationListOK) Error() string {
	return fmt.Sprintf("[POST /domain/cancelation/_search][%d] domainCancelationListOK  %+v", 200, o.Payload)
}

func (o *DomainCancelationListOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.JSONResponseDataDomainCancelation)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
