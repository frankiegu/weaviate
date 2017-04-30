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
 package acl_entries




import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/weaviate/weaviate/models"
)

// WeaviateACLEntriesInsertCreatedCode is the HTTP code returned for type WeaviateACLEntriesInsertCreated
const WeaviateACLEntriesInsertCreatedCode int = 201

/*WeaviateACLEntriesInsertCreated Successful created.

swagger:response weaviateAclEntriesInsertCreated
*/
type WeaviateACLEntriesInsertCreated struct {

	/*
	  In: Body
	*/
	Payload *models.ACLEntry `json:"body,omitempty"`
}

// NewWeaviateACLEntriesInsertCreated creates WeaviateACLEntriesInsertCreated with default headers values
func NewWeaviateACLEntriesInsertCreated() *WeaviateACLEntriesInsertCreated {
	return &WeaviateACLEntriesInsertCreated{}
}

// WithPayload adds the payload to the weaviate Acl entries insert created response
func (o *WeaviateACLEntriesInsertCreated) WithPayload(payload *models.ACLEntry) *WeaviateACLEntriesInsertCreated {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the weaviate Acl entries insert created response
func (o *WeaviateACLEntriesInsertCreated) SetPayload(payload *models.ACLEntry) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *WeaviateACLEntriesInsertCreated) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(201)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// WeaviateACLEntriesInsertNotImplementedCode is the HTTP code returned for type WeaviateACLEntriesInsertNotImplemented
const WeaviateACLEntriesInsertNotImplementedCode int = 501

/*WeaviateACLEntriesInsertNotImplemented Not (yet) implemented.

swagger:response weaviateAclEntriesInsertNotImplemented
*/
type WeaviateACLEntriesInsertNotImplemented struct {
}

// NewWeaviateACLEntriesInsertNotImplemented creates WeaviateACLEntriesInsertNotImplemented with default headers values
func NewWeaviateACLEntriesInsertNotImplemented() *WeaviateACLEntriesInsertNotImplemented {
	return &WeaviateACLEntriesInsertNotImplemented{}
}

// WriteResponse to the client
func (o *WeaviateACLEntriesInsertNotImplemented) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(501)
}
