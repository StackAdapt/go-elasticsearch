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

// Privileges type alias.
//
// https://github.com/elastic/elasticsearch-specification/blob/135ae054e304239743b5777ad8d41cb2c9091d35/specification/security/has_privileges/types.ts#L48-L48
type Privileges map[bool]string

// PrivilegesBuilder holds Privileges struct and provides a builder API.
type PrivilegesBuilder struct {
	v Privileges
}

// NewPrivileges provides a builder for the Privileges struct.
func NewPrivilegesBuilder() *PrivilegesBuilder {
	return &PrivilegesBuilder{}
}

// Build finalize the chain and returns the Privileges struct
func (b *PrivilegesBuilder) Build() Privileges {
	return b.v
}

func (b *PrivilegesBuilder) Privileges(value Privileges) *PrivilegesBuilder {
	b.v = value
	return b
}