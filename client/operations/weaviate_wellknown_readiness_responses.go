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

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// WeaviateWellknownReadinessReader is a Reader for the WeaviateWellknownReadiness structure.
type WeaviateWellknownReadinessReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *WeaviateWellknownReadinessReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewWeaviateWellknownReadinessOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 503:
		result := NewWeaviateWellknownReadinessServiceUnavailable()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewWeaviateWellknownReadinessOK creates a WeaviateWellknownReadinessOK with default headers values
func NewWeaviateWellknownReadinessOK() *WeaviateWellknownReadinessOK {
	return &WeaviateWellknownReadinessOK{}
}

/*WeaviateWellknownReadinessOK handles this case with default header values.

The application has completed its start-up routine and is ready to accept traffic.
*/
type WeaviateWellknownReadinessOK struct {
}

func (o *WeaviateWellknownReadinessOK) Error() string {
	return fmt.Sprintf("[GET /.well-known/ready][%d] weaviateWellknownReadinessOK ", 200)
}

func (o *WeaviateWellknownReadinessOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewWeaviateWellknownReadinessServiceUnavailable creates a WeaviateWellknownReadinessServiceUnavailable with default headers values
func NewWeaviateWellknownReadinessServiceUnavailable() *WeaviateWellknownReadinessServiceUnavailable {
	return &WeaviateWellknownReadinessServiceUnavailable{}
}

/*WeaviateWellknownReadinessServiceUnavailable handles this case with default header values.

The application is currently not able to serve traffic. If other horizontal replicas of weaviate are available and they are capable of receiving traffic, all traffic should be redirected there instead.
*/
type WeaviateWellknownReadinessServiceUnavailable struct {
}

func (o *WeaviateWellknownReadinessServiceUnavailable) Error() string {
	return fmt.Sprintf("[GET /.well-known/ready][%d] weaviateWellknownReadinessServiceUnavailable ", 503)
}

func (o *WeaviateWellknownReadinessServiceUnavailable) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
