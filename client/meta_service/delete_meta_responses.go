// Code generated by go-swagger; DO NOT EDIT.

package meta_service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/pydio/cells-sdk-go/models"
)

// DeleteMetaReader is a Reader for the DeleteMeta structure.
type DeleteMetaReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteMetaReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewDeleteMetaOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewDeleteMetaOK creates a DeleteMetaOK with default headers values
func NewDeleteMetaOK() *DeleteMetaOK {
	return &DeleteMetaOK{}
}

/*DeleteMetaOK handles this case with default header values.

DeleteMetaOK delete meta o k
*/
type DeleteMetaOK struct {
	Payload *models.TreeNode
}

func (o *DeleteMetaOK) Error() string {
	return fmt.Sprintf("[POST /meta/delete/{NodePath}][%d] deleteMetaOK  %+v", 200, o.Payload)
}

func (o *DeleteMetaOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.TreeNode)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
