// Licensed to Elasticsearch B.V. under one or more contributor
// license agreements. See the NOTICE file distributed with
// this work for additional information regarding copyright
// ownership. Elasticsearch B.V. licenses this file to you under
// the Apache License, Version 2.0 (the "License"); you may
// not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.

// Code generated from the elasticsearch-specification DO NOT EDIT.
// https://github.com/elastic/elasticsearch-specification/tree/8e91c0692c0235474a0c21bb7e9716a8430e8533

package updateactivefiltering

import (
	"github.com/StackAdapt/go-elasticsearch/v8/typedapi/types/enums/result"
)

// Response holds the response body struct for the package updateactivefiltering
//
// https://github.com/elastic/elasticsearch-specification/blob/8e91c0692c0235474a0c21bb7e9716a8430e8533/specification/connector/update_active_filtering/ConnectorUpdateActiveFilteringResponse.ts#L22-L26
type Response struct {
	Result result.Result `json:"result"`
}

// NewResponse returns a Response
func NewResponse() *Response {
	r := &Response{}
	return r
}
