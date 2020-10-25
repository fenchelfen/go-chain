// Code generated by go-swagger; DO NOT EDIT.

package peers

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"gochain/models"
)

// RegisterNodeOKCode is the HTTP code returned for type RegisterNodeOK
const RegisterNodeOKCode int = 200

/*RegisterNodeOK success

swagger:response registerNodeOK
*/
type RegisterNodeOK struct {

	/*
	  In: Body
	*/
	Payload *models.Chain `json:"body,omitempty"`
}

// NewRegisterNodeOK creates RegisterNodeOK with default headers values
func NewRegisterNodeOK() *RegisterNodeOK {

	return &RegisterNodeOK{}
}

// WithPayload adds the payload to the register node o k response
func (o *RegisterNodeOK) WithPayload(payload *models.Chain) *RegisterNodeOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the register node o k response
func (o *RegisterNodeOK) SetPayload(payload *models.Chain) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *RegisterNodeOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
