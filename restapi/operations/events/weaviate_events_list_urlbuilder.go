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


// Editing this file might prove futile when you re-run the generate command

import (
	"errors"
	"net/url"
	golangswaggerpaths "path"

	"github.com/go-openapi/swag"
)

// WeaviateEventsListURL generates an URL for the weaviate events list operation
type WeaviateEventsListURL struct {
	Alt         *string
	CommandID   []string
	DeviceID    []string
	EndTimeMs   *string
	Fields      *string
	Hl          *string
	Key         *string
	MaxResults  *int64
	OauthToken  *string
	PrettyPrint *bool
	QuotaUser   *string
	StartIndex  *int64
	StartTimeMs *string
	Token       *string
	Type        *string
	UserIP      *string

	_basePath string
	// avoid unkeyed usage
	_ struct{}
}

// WithBasePath sets the base path for this url builder, only required when it's different from the
// base path specified in the swagger spec.
// When the value of the base path is an empty string
func (o *WeaviateEventsListURL) WithBasePath(bp string) *WeaviateEventsListURL {
	o.SetBasePath(bp)
	return o
}

// SetBasePath sets the base path for this url builder, only required when it's different from the
// base path specified in the swagger spec.
// When the value of the base path is an empty string
func (o *WeaviateEventsListURL) SetBasePath(bp string) {
	o._basePath = bp
}

// Build a url path and query string
func (o *WeaviateEventsListURL) Build() (*url.URL, error) {
	var result url.URL

	var _path = "/events"

	_basePath := o._basePath
	if _basePath == "" {
		_basePath = "/weaviate/v1-alpha"
	}
	result.Path = golangswaggerpaths.Join(_basePath, _path)

	qs := make(url.Values)

	var alt string
	if o.Alt != nil {
		alt = *o.Alt
	}
	if alt != "" {
		qs.Set("alt", alt)
	}

	var commandIDIR []string
	for _, commandIDI := range o.CommandID {
		commandIDIS := commandIDI
		if commandIDIS != "" {
			commandIDIR = append(commandIDIR, commandIDIS)
		}
	}

	commandID := swag.JoinByFormat(commandIDIR, "multi")

	for _, qsv := range commandID {
		qs.Add("commandId", qsv)
	}

	var deviceIDIR []string
	for _, deviceIDI := range o.DeviceID {
		deviceIDIS := deviceIDI
		if deviceIDIS != "" {
			deviceIDIR = append(deviceIDIR, deviceIDIS)
		}
	}

	deviceID := swag.JoinByFormat(deviceIDIR, "multi")

	for _, qsv := range deviceID {
		qs.Add("deviceId", qsv)
	}

	var endTimeMs string
	if o.EndTimeMs != nil {
		endTimeMs = *o.EndTimeMs
	}
	if endTimeMs != "" {
		qs.Set("endTimeMs", endTimeMs)
	}

	var fields string
	if o.Fields != nil {
		fields = *o.Fields
	}
	if fields != "" {
		qs.Set("fields", fields)
	}

	var hl string
	if o.Hl != nil {
		hl = *o.Hl
	}
	if hl != "" {
		qs.Set("hl", hl)
	}

	var key string
	if o.Key != nil {
		key = *o.Key
	}
	if key != "" {
		qs.Set("key", key)
	}

	var maxResults string
	if o.MaxResults != nil {
		maxResults = swag.FormatInt64(*o.MaxResults)
	}
	if maxResults != "" {
		qs.Set("maxResults", maxResults)
	}

	var oauthToken string
	if o.OauthToken != nil {
		oauthToken = *o.OauthToken
	}
	if oauthToken != "" {
		qs.Set("oauth_token", oauthToken)
	}

	var prettyPrint string
	if o.PrettyPrint != nil {
		prettyPrint = swag.FormatBool(*o.PrettyPrint)
	}
	if prettyPrint != "" {
		qs.Set("prettyPrint", prettyPrint)
	}

	var quotaUser string
	if o.QuotaUser != nil {
		quotaUser = *o.QuotaUser
	}
	if quotaUser != "" {
		qs.Set("quotaUser", quotaUser)
	}

	var startIndex string
	if o.StartIndex != nil {
		startIndex = swag.FormatInt64(*o.StartIndex)
	}
	if startIndex != "" {
		qs.Set("startIndex", startIndex)
	}

	var startTimeMs string
	if o.StartTimeMs != nil {
		startTimeMs = *o.StartTimeMs
	}
	if startTimeMs != "" {
		qs.Set("startTimeMs", startTimeMs)
	}

	var token string
	if o.Token != nil {
		token = *o.Token
	}
	if token != "" {
		qs.Set("token", token)
	}

	var typeVar string
	if o.Type != nil {
		typeVar = *o.Type
	}
	if typeVar != "" {
		qs.Set("type", typeVar)
	}

	var userIP string
	if o.UserIP != nil {
		userIP = *o.UserIP
	}
	if userIP != "" {
		qs.Set("userIp", userIP)
	}

	result.RawQuery = qs.Encode()

	return &result, nil
}

// Must is a helper function to panic when the url builder returns an error
func (o *WeaviateEventsListURL) Must(u *url.URL, err error) *url.URL {
	if err != nil {
		panic(err)
	}
	if u == nil {
		panic("url can't be nil")
	}
	return u
}

// String returns the string representation of the path with query string
func (o *WeaviateEventsListURL) String() string {
	return o.Must(o.Build()).String()
}

// BuildFull builds a full url with scheme, host, path and query string
func (o *WeaviateEventsListURL) BuildFull(scheme, host string) (*url.URL, error) {
	if scheme == "" {
		return nil, errors.New("scheme is required for a full url on WeaviateEventsListURL")
	}
	if host == "" {
		return nil, errors.New("host is required for a full url on WeaviateEventsListURL")
	}

	base, err := o.Build()
	if err != nil {
		return nil, err
	}

	base.Scheme = scheme
	base.Host = host
	return base, nil
}

// StringFull returns the string representation of a complete url
func (o *WeaviateEventsListURL) StringFull(scheme, host string) string {
	return o.Must(o.BuildFull(scheme, host)).String()
}
