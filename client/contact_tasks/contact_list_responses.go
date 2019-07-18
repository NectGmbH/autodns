// Code generated by go-swagger; DO NOT EDIT.

package contact_tasks

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/NectGmbH/autodns/models"
)

// ContactListReader is a Reader for the ContactList structure.
type ContactListReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ContactListReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewContactListOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewContactListOK creates a ContactListOK with default headers values
func NewContactListOK() *ContactListOK {
	return &ContactListOK{}
}

/*ContactListOK handles this case with default header values.

successful operation
*/
type ContactListOK struct {
	Payload *models.JSONResponseDataContact
}

func (o *ContactListOK) Error() string {
	return fmt.Sprintf("[POST /contact/_search][%d] contactListOK  %+v", 200, o.Payload)
}

func (o *ContactListOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.JSONResponseDataContact)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
