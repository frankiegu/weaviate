/*                          _       _
 *__      _____  __ ___   ___  __ _| |_ ___
 *\ \ /\ / / _ \/ _` \ \ / / |/ _` | __/ _ \
 * \ V  V /  __/ (_| |\ V /| | (_| | ||  __/
 *  \_/\_/ \___|\__,_| \_/ |_|\__,_|\__\___|
 *
 * Copyright © 2016 Weaviate. All rights reserved.
 * LICENSE: https://github.com/weaviate/weaviate/blob/master/LICENSE
 * AUTHOR: Bob van Luijt (bob@weaviate.com)
 * See www.weaviate.com for details
 * Contact: @weaviate_iot / yourfriends@weaviate.com
 */
 package events




import (
	"net/http"

	"github.com/go-openapi/runtime"
)

// WeaviateEventsRecordDeviceEventsCreatedCode is the HTTP code returned for type WeaviateEventsRecordDeviceEventsCreated
const WeaviateEventsRecordDeviceEventsCreatedCode int = 201

/*WeaviateEventsRecordDeviceEventsCreated Successful created.

swagger:response weaviateEventsRecordDeviceEventsCreated
*/
type WeaviateEventsRecordDeviceEventsCreated struct {
}

// NewWeaviateEventsRecordDeviceEventsCreated creates WeaviateEventsRecordDeviceEventsCreated with default headers values
func NewWeaviateEventsRecordDeviceEventsCreated() *WeaviateEventsRecordDeviceEventsCreated {
	return &WeaviateEventsRecordDeviceEventsCreated{}
}

// WriteResponse to the client
func (o *WeaviateEventsRecordDeviceEventsCreated) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(201)
}

// WeaviateEventsRecordDeviceEventsNotImplementedCode is the HTTP code returned for type WeaviateEventsRecordDeviceEventsNotImplemented
const WeaviateEventsRecordDeviceEventsNotImplementedCode int = 501

/*WeaviateEventsRecordDeviceEventsNotImplemented Not (yet) implemented.

swagger:response weaviateEventsRecordDeviceEventsNotImplemented
*/
type WeaviateEventsRecordDeviceEventsNotImplemented struct {
}

// NewWeaviateEventsRecordDeviceEventsNotImplemented creates WeaviateEventsRecordDeviceEventsNotImplemented with default headers values
func NewWeaviateEventsRecordDeviceEventsNotImplemented() *WeaviateEventsRecordDeviceEventsNotImplemented {
	return &WeaviateEventsRecordDeviceEventsNotImplemented{}
}

// WriteResponse to the client
func (o *WeaviateEventsRecordDeviceEventsNotImplemented) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(501)
}
