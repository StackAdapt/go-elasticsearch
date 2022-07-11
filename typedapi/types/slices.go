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
// https://github.com/elastic/elasticsearch-specification/tree/135ae054e304239743b5777ad8d41cb2c9091d35


package types

import (
	"github.com/elastic/go-elasticsearch/v8/typedapi/types/enums/slicescalculation"
)

// Slices holds the union for the following types:
//     int
//     slicescalculation.SlicesCalculation
//
// https://github.com/elastic/elasticsearch-specification/blob/135ae054e304239743b5777ad8d41cb2c9091d35/specification/_types/common.ts#L319-L324
type Slices interface{}

// SlicesBuilder holds Slices struct and provides a builder API.
type SlicesBuilder struct {
	v Slices
}

// NewSlices provides a builder for the Slices struct.
func NewSlicesBuilder() *SlicesBuilder {
	return &SlicesBuilder{}
}

// Build finalize the chain and returns the Slices struct
func (u *SlicesBuilder) Build() Slices {
	return u.v
}

func (u *SlicesBuilder) Int(int int) *SlicesBuilder {
	u.v = &int
	return u
}

func (u *SlicesBuilder) SlicesCalculation(slicescalculation slicescalculation.SlicesCalculation) *SlicesBuilder {
	u.v = &slicescalculation
	return u
}