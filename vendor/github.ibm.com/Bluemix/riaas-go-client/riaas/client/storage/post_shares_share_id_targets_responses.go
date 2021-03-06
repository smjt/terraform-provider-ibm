// Code generated by go-swagger; DO NOT EDIT.

package storage

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.ibm.com/Bluemix/riaas-go-client/riaas/models"
)

// PostSharesShareIDTargetsReader is a Reader for the PostSharesShareIDTargets structure.
type PostSharesShareIDTargetsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostSharesShareIDTargetsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 201:
		result := NewPostSharesShareIDTargetsCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewPostSharesShareIDTargetsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewPostSharesShareIDTargetsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPostSharesShareIDTargetsCreated creates a PostSharesShareIDTargetsCreated with default headers values
func NewPostSharesShareIDTargetsCreated() *PostSharesShareIDTargetsCreated {
	return &PostSharesShareIDTargetsCreated{}
}

/*PostSharesShareIDTargetsCreated handles this case with default header values.

dummy
*/
type PostSharesShareIDTargetsCreated struct {
	Payload *models.Sharetarget
}

func (o *PostSharesShareIDTargetsCreated) Error() string {
	return fmt.Sprintf("[POST /shares/{share_id}/targets][%d] postSharesShareIdTargetsCreated  %+v", 201, o.Payload)
}

func (o *PostSharesShareIDTargetsCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Sharetarget)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostSharesShareIDTargetsBadRequest creates a PostSharesShareIDTargetsBadRequest with default headers values
func NewPostSharesShareIDTargetsBadRequest() *PostSharesShareIDTargetsBadRequest {
	return &PostSharesShareIDTargetsBadRequest{}
}

/*PostSharesShareIDTargetsBadRequest handles this case with default header values.

error
*/
type PostSharesShareIDTargetsBadRequest struct {
	Payload *models.Riaaserror
}

func (o *PostSharesShareIDTargetsBadRequest) Error() string {
	return fmt.Sprintf("[POST /shares/{share_id}/targets][%d] postSharesShareIdTargetsBadRequest  %+v", 400, o.Payload)
}

func (o *PostSharesShareIDTargetsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Riaaserror)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostSharesShareIDTargetsNotFound creates a PostSharesShareIDTargetsNotFound with default headers values
func NewPostSharesShareIDTargetsNotFound() *PostSharesShareIDTargetsNotFound {
	return &PostSharesShareIDTargetsNotFound{}
}

/*PostSharesShareIDTargetsNotFound handles this case with default header values.

error
*/
type PostSharesShareIDTargetsNotFound struct {
	Payload *models.Riaaserror
}

func (o *PostSharesShareIDTargetsNotFound) Error() string {
	return fmt.Sprintf("[POST /shares/{share_id}/targets][%d] postSharesShareIdTargetsNotFound  %+v", 404, o.Payload)
}

func (o *PostSharesShareIDTargetsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Riaaserror)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
