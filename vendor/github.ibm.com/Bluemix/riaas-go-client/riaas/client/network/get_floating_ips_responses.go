// Code generated by go-swagger; DO NOT EDIT.

package network

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.ibm.com/Bluemix/riaas-go-client/riaas/models"
)

// GetFloatingIpsReader is a Reader for the GetFloatingIps structure.
type GetFloatingIpsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetFloatingIpsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetFloatingIpsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 500:
		result := NewGetFloatingIpsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetFloatingIpsOK creates a GetFloatingIpsOK with default headers values
func NewGetFloatingIpsOK() *GetFloatingIpsOK {
	return &GetFloatingIpsOK{}
}

/*GetFloatingIpsOK handles this case with default header values.

dummy
*/
type GetFloatingIpsOK struct {
	Payload *models.GetFloatingIpsOKBody
}

func (o *GetFloatingIpsOK) Error() string {
	return fmt.Sprintf("[GET /floating_ips][%d] getFloatingIpsOK  %+v", 200, o.Payload)
}

func (o *GetFloatingIpsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GetFloatingIpsOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetFloatingIpsInternalServerError creates a GetFloatingIpsInternalServerError with default headers values
func NewGetFloatingIpsInternalServerError() *GetFloatingIpsInternalServerError {
	return &GetFloatingIpsInternalServerError{}
}

/*GetFloatingIpsInternalServerError handles this case with default header values.

error
*/
type GetFloatingIpsInternalServerError struct {
	Payload *models.Riaaserror
}

func (o *GetFloatingIpsInternalServerError) Error() string {
	return fmt.Sprintf("[GET /floating_ips][%d] getFloatingIpsInternalServerError  %+v", 500, o.Payload)
}

func (o *GetFloatingIpsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Riaaserror)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
