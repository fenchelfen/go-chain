// Code generated by go-swagger; DO NOT EDIT.

package peers

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"
)

// MineOKCode is the HTTP code returned for type MineOK
const MineOKCode int = 200

/*MineOK success

swagger:response mineOK
*/
type MineOK struct {
}

// NewMineOK creates MineOK with default headers values
func NewMineOK() *MineOK {

	return &MineOK{}
}

// WriteResponse to the client
func (o *MineOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(200)
}
