// Code generated by go-swagger; DO NOT EDIT.

package backup_mx_tasks

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/NectGmbH/autodns/models"
)

// BackupMxDeleteReader is a Reader for the BackupMxDelete structure.
type BackupMxDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *BackupMxDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewBackupMxDeleteOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewBackupMxDeleteOK creates a BackupMxDeleteOK with default headers values
func NewBackupMxDeleteOK() *BackupMxDeleteOK {
	return &BackupMxDeleteOK{}
}

/*BackupMxDeleteOK handles this case with default header values.

successful operation
*/
type BackupMxDeleteOK struct {
	Payload *models.JSONResponseDataJSONNoData
}

func (o *BackupMxDeleteOK) Error() string {
	return fmt.Sprintf("[DELETE /backupMx/{domain}][%d] backupMxDeleteOK  %+v", 200, o.Payload)
}

func (o *BackupMxDeleteOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.JSONResponseDataJSONNoData)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
