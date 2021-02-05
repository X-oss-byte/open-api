// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/netlify/open-api/v2/go/models"
)

// ListSiteSubmissionsReader is a Reader for the ListSiteSubmissions structure.
type ListSiteSubmissionsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListSiteSubmissionsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewListSiteSubmissionsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewListSiteSubmissionsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewListSiteSubmissionsOK creates a ListSiteSubmissionsOK with default headers values
func NewListSiteSubmissionsOK() *ListSiteSubmissionsOK {
	return &ListSiteSubmissionsOK{}
}

/*ListSiteSubmissionsOK handles this case with default header values.

OK
*/
type ListSiteSubmissionsOK struct {
	Payload []*models.Submission
}

func (o *ListSiteSubmissionsOK) Error() string {
	return fmt.Sprintf("[GET /sites/{site_id}/submissions][%d] listSiteSubmissionsOK  %+v", 200, o.Payload)
}

func (o *ListSiteSubmissionsOK) GetPayload() []*models.Submission {
	return o.Payload
}

func (o *ListSiteSubmissionsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListSiteSubmissionsDefault creates a ListSiteSubmissionsDefault with default headers values
func NewListSiteSubmissionsDefault(code int) *ListSiteSubmissionsDefault {
	return &ListSiteSubmissionsDefault{
		_statusCode: code,
	}
}

/*ListSiteSubmissionsDefault handles this case with default header values.

error
*/
type ListSiteSubmissionsDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the list site submissions default response
func (o *ListSiteSubmissionsDefault) Code() int {
	return o._statusCode
}

func (o *ListSiteSubmissionsDefault) Error() string {
	return fmt.Sprintf("[GET /sites/{site_id}/submissions][%d] listSiteSubmissions default  %+v", o._statusCode, o.Payload)
}

func (o *ListSiteSubmissionsDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *ListSiteSubmissionsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
