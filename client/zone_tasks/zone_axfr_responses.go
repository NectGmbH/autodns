// Code generated by go-swagger; DO NOT EDIT.

package zone_tasks

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/NectGmbH/autodns/models"
)

// ZoneAxfrReader is a Reader for the ZoneAxfr structure.
type ZoneAxfrReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ZoneAxfrReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewZoneAxfrOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewZoneAxfrOK creates a ZoneAxfrOK with default headers values
func NewZoneAxfrOK() *ZoneAxfrOK {
	return &ZoneAxfrOK{}
}

/*ZoneAxfrOK handles this case with default header values.

successful operation
*/
type ZoneAxfrOK struct {
	Payload *models.JSONResponseDataString
}

func (o *ZoneAxfrOK) Error() string {
	return fmt.Sprintf("[GET /zone/{name}/{nameserver}/_axfr][%d] zoneAxfrOK  %+v", 200, o.Payload)
}

func (o *ZoneAxfrOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.JSONResponseDataString)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
