// Code generated by go-swagger; DO NOT EDIT.

package poll_tasks

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/NectGmbH/autodns/models"
)

// PollConfirmReader is a Reader for the PollConfirm structure.
type PollConfirmReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PollConfirmReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewPollConfirmOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPollConfirmOK creates a PollConfirmOK with default headers values
func NewPollConfirmOK() *PollConfirmOK {
	return &PollConfirmOK{}
}

/*PollConfirmOK handles this case with default header values.

successful operation
*/
type PollConfirmOK struct {
	Payload *models.JSONResponseDataJSONNoData
}

func (o *PollConfirmOK) Error() string {
	return fmt.Sprintf("[PUT /poll/{id}][%d] pollConfirmOK  %+v", 200, o.Payload)
}

func (o *PollConfirmOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.JSONResponseDataJSONNoData)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}