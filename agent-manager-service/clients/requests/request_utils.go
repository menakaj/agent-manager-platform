// Copyright (c) 2025, WSO2 LLC (http://www.wso2.com). All Rights Reserved.
//
// This software is the property of WSO2 LLC and its suppliers, if any.
// Dissemination of any information or reproduction of any material contained
// herein is strictly forbidden, unless permitted by WSO2 in accordance with
// the WSO2 Commercial License available at http://wso2.com/licenses.
// For specific language governing the permissions and limitations under
// this license, please see the license as well as any agreement you've
// entered into with WSO2 governing the purchase of this software and any
// associated services.

package requests

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

// HttpRequest represents a retryable HTTP request.
type HttpRequest struct {
	// Name is a human-readable name for the request. To be used in logs.
	// e.g. "packageName.methodName"
	Name string

	URL    string
	Method string
	Query  map[string]string

	headers   http.Header
	body      []byte
	createErr error
}

func (r *HttpRequest) SetHeader(key, value string) *HttpRequest {
	if r.headers == nil {
		r.headers = http.Header{}
	}
	r.headers.Set(key, value)
	return r
}

func (r *HttpRequest) SetQuery(key, value string) *HttpRequest {
	if r.Query == nil {
		r.Query = make(map[string]string)
	}
	r.Query[key] = value
	return r
}

func (r *HttpRequest) SetJson(body any) *HttpRequest {
	v, err := json.Marshal(body)
	if err != nil {
		r.createErr = fmt.Errorf("failed to encode request body: %w", err)
		return r
	}
	r.body = v
	r.SetHeader("Content-Type", "application/json")
	return r
}

func (r *HttpRequest) buildHttpRequest(ctx context.Context) (*http.Request, error) {
	if r.createErr != nil {
		return nil, r.createErr
	}
	request, err := http.NewRequestWithContext(ctx, r.Method, r.URL, bytes.NewReader(r.body))
	if err != nil {
		return nil, err
	}
	q := request.URL.Query()
	for key, value := range r.Query {
		q.Add(key, value)
	}
	request.URL.RawQuery = q.Encode()
	request.Header = r.headers
	return request, nil
}
