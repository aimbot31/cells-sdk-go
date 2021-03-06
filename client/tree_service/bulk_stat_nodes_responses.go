// Code generated by go-swagger; DO NOT EDIT.

package tree_service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/pydio/cells-sdk-go/models"
)

// BulkStatNodesReader is a Reader for the BulkStatNodes structure.
type BulkStatNodesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *BulkStatNodesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewBulkStatNodesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewBulkStatNodesOK creates a BulkStatNodesOK with default headers values
func NewBulkStatNodesOK() *BulkStatNodesOK {
	return &BulkStatNodesOK{}
}

/*BulkStatNodesOK handles this case with default header values.

BulkStatNodesOK bulk stat nodes o k
*/
type BulkStatNodesOK struct {
	Payload *models.RestBulkMetaResponse
}

func (o *BulkStatNodesOK) Error() string {
	return fmt.Sprintf("[POST /tree/stats][%d] bulkStatNodesOK  %+v", 200, o.Payload)
}

func (o *BulkStatNodesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RestBulkMetaResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
