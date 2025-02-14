//                           _       _
// __      _____  __ ___   ___  __ _| |_ ___
// \ \ /\ / / _ \/ _` \ \ / / |/ _` | __/ _ \
//  \ V  V /  __/ (_| |\ V /| | (_| | ||  __/
//   \_/\_/ \___|\__,_| \_/ |_|\__,_|\__\___|
//
//  Copyright © 2016 - 2022 SeMI Technologies B.V. All rights reserved.
//
//  CONTACT: hello@semi.technology
//

// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/semi-technologies/weaviate/genesis/models"
)

// GenesisPeersRegisterReader is a Reader for the GenesisPeersRegister structure.
type GenesisPeersRegisterReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GenesisPeersRegisterReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGenesisPeersRegisterOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGenesisPeersRegisterBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGenesisPeersRegisterForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGenesisPeersRegisterOK creates a GenesisPeersRegisterOK with default headers values
func NewGenesisPeersRegisterOK() *GenesisPeersRegisterOK {
	return &GenesisPeersRegisterOK{}
}

/*GenesisPeersRegisterOK handles this case with default header values.

Successfully registred the peer to the network.
*/
type GenesisPeersRegisterOK struct {
	Payload *models.PeerRegistrationResponse
}

func (o *GenesisPeersRegisterOK) Error() string {
	return fmt.Sprintf("[POST /peers/register][%d] genesisPeersRegisterOK  %+v", 200, o.Payload)
}

func (o *GenesisPeersRegisterOK) GetPayload() *models.PeerRegistrationResponse {
	return o.Payload
}

func (o *GenesisPeersRegisterOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {
	o.Payload = new(models.PeerRegistrationResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGenesisPeersRegisterBadRequest creates a GenesisPeersRegisterBadRequest with default headers values
func NewGenesisPeersRegisterBadRequest() *GenesisPeersRegisterBadRequest {
	return &GenesisPeersRegisterBadRequest{}
}

/*GenesisPeersRegisterBadRequest handles this case with default header values.

The weaviate peer is not reachable from the Gensis service.
*/
type GenesisPeersRegisterBadRequest struct {
}

func (o *GenesisPeersRegisterBadRequest) Error() string {
	return fmt.Sprintf("[POST /peers/register][%d] genesisPeersRegisterBadRequest ", 400)
}

func (o *GenesisPeersRegisterBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {
	return nil
}

// NewGenesisPeersRegisterForbidden creates a GenesisPeersRegisterForbidden with default headers values
func NewGenesisPeersRegisterForbidden() *GenesisPeersRegisterForbidden {
	return &GenesisPeersRegisterForbidden{}
}

/*GenesisPeersRegisterForbidden handles this case with default header values.

You are not allowed on the network.
*/
type GenesisPeersRegisterForbidden struct {
}

func (o *GenesisPeersRegisterForbidden) Error() string {
	return fmt.Sprintf("[POST /peers/register][%d] genesisPeersRegisterForbidden ", 403)
}

func (o *GenesisPeersRegisterForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {
	return nil
}
