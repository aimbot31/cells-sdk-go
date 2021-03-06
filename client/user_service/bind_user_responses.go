// Code generated by go-swagger; DO NOT EDIT.

package user_service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/pydio/cells-sdk-go/models"
)

// BindUserReader is a Reader for the BindUser structure.
type BindUserReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *BindUserReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewBindUserOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewBindUserOK creates a BindUserOK with default headers values
func NewBindUserOK() *BindUserOK {
	return &BindUserOK{}
}

/*BindUserOK handles this case with default header values.

BindUserOK bind user o k
*/
type BindUserOK struct {
	Payload *models.RestBindResponse
}

func (o *BindUserOK) Error() string {
	return fmt.Sprintf("[POST /user/{Login}/bind][%d] bindUserOK  %+v", 200, o.Payload)
}

func (o *BindUserOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RestBindResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
