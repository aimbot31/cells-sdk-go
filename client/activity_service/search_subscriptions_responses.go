// Code generated by go-swagger; DO NOT EDIT.

package activity_service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/pydio/cells-sdk-go/models"
)

// SearchSubscriptionsReader is a Reader for the SearchSubscriptions structure.
type SearchSubscriptionsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SearchSubscriptionsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewSearchSubscriptionsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewSearchSubscriptionsOK creates a SearchSubscriptionsOK with default headers values
func NewSearchSubscriptionsOK() *SearchSubscriptionsOK {
	return &SearchSubscriptionsOK{}
}

/*SearchSubscriptionsOK handles this case with default header values.

SearchSubscriptionsOK search subscriptions o k
*/
type SearchSubscriptionsOK struct {
	Payload *models.RestSubscriptionsCollection
}

func (o *SearchSubscriptionsOK) Error() string {
	return fmt.Sprintf("[POST /activity/subscriptions][%d] searchSubscriptionsOK  %+v", 200, o.Payload)
}

func (o *SearchSubscriptionsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RestSubscriptionsCollection)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
