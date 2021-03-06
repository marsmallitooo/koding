package j_provisioner

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"

	"koding/remoteapi/models"
)

// PostRemoteAPIJProvisionerSetAccessIDReader is a Reader for the PostRemoteAPIJProvisionerSetAccessID structure.
type PostRemoteAPIJProvisionerSetAccessIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostRemoteAPIJProvisionerSetAccessIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewPostRemoteAPIJProvisionerSetAccessIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPostRemoteAPIJProvisionerSetAccessIDOK creates a PostRemoteAPIJProvisionerSetAccessIDOK with default headers values
func NewPostRemoteAPIJProvisionerSetAccessIDOK() *PostRemoteAPIJProvisionerSetAccessIDOK {
	return &PostRemoteAPIJProvisionerSetAccessIDOK{}
}

/*PostRemoteAPIJProvisionerSetAccessIDOK handles this case with default header values.

OK
*/
type PostRemoteAPIJProvisionerSetAccessIDOK struct {
	Payload PostRemoteAPIJProvisionerSetAccessIDOKBody
}

func (o *PostRemoteAPIJProvisionerSetAccessIDOK) Error() string {
	return fmt.Sprintf("[POST /remote.api/JProvisioner.setAccess/{id}][%d] postRemoteApiJProvisionerSetAccessIdOK  %+v", 200, o.Payload)
}

func (o *PostRemoteAPIJProvisionerSetAccessIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*PostRemoteAPIJProvisionerSetAccessIDOKBody post remote API j provisioner set access ID o k body
swagger:model PostRemoteAPIJProvisionerSetAccessIDOKBody
*/
type PostRemoteAPIJProvisionerSetAccessIDOKBody struct {
	models.JProvisioner

	models.DefaultResponse
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (o *PostRemoteAPIJProvisionerSetAccessIDOKBody) UnmarshalJSON(raw []byte) error {

	var postRemoteAPIJProvisionerSetAccessIDOKBodyAO0 models.JProvisioner
	if err := swag.ReadJSON(raw, &postRemoteAPIJProvisionerSetAccessIDOKBodyAO0); err != nil {
		return err
	}
	o.JProvisioner = postRemoteAPIJProvisionerSetAccessIDOKBodyAO0

	var postRemoteAPIJProvisionerSetAccessIDOKBodyAO1 models.DefaultResponse
	if err := swag.ReadJSON(raw, &postRemoteAPIJProvisionerSetAccessIDOKBodyAO1); err != nil {
		return err
	}
	o.DefaultResponse = postRemoteAPIJProvisionerSetAccessIDOKBodyAO1

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (o PostRemoteAPIJProvisionerSetAccessIDOKBody) MarshalJSON() ([]byte, error) {
	var _parts [][]byte

	postRemoteAPIJProvisionerSetAccessIDOKBodyAO0, err := swag.WriteJSON(o.JProvisioner)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, postRemoteAPIJProvisionerSetAccessIDOKBodyAO0)

	postRemoteAPIJProvisionerSetAccessIDOKBodyAO1, err := swag.WriteJSON(o.DefaultResponse)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, postRemoteAPIJProvisionerSetAccessIDOKBodyAO1)

	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this post remote API j provisioner set access ID o k body
func (o *PostRemoteAPIJProvisionerSetAccessIDOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.JProvisioner.Validate(formats); err != nil {
		res = append(res, err)
	}

	if err := o.DefaultResponse.Validate(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
