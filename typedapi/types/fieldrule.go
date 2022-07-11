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

// FieldRule type.
//
// https://github.com/elastic/elasticsearch-specification/blob/135ae054e304239743b5777ad8d41cb2c9091d35/specification/security/_types/RoleMappingRule.ts#L33-L42
type FieldRule struct {
	Dn       *Names      `json:"dn,omitempty"`
	Groups   *Names      `json:"groups,omitempty"`
	Metadata interface{} `json:"metadata,omitempty"`
	Realm    *Realm      `json:"realm,omitempty"`
	Username *Name       `json:"username,omitempty"`
}

// FieldRuleBuilder holds FieldRule struct and provides a builder API.
type FieldRuleBuilder struct {
	v *FieldRule
}

// NewFieldRule provides a builder for the FieldRule struct.
func NewFieldRuleBuilder() *FieldRuleBuilder {
	r := FieldRuleBuilder{
		&FieldRule{},
	}

	return &r
}

// Build finalize the chain and returns the FieldRule struct
func (rb *FieldRuleBuilder) Build() FieldRule {
	return *rb.v
}

func (rb *FieldRuleBuilder) Dn(dn *NamesBuilder) *FieldRuleBuilder {
	v := dn.Build()
	rb.v.Dn = &v
	return rb
}

func (rb *FieldRuleBuilder) Groups(groups *NamesBuilder) *FieldRuleBuilder {
	v := groups.Build()
	rb.v.Groups = &v
	return rb
}

func (rb *FieldRuleBuilder) Metadata(metadata interface{}) *FieldRuleBuilder {
	rb.v.Metadata = metadata
	return rb
}

func (rb *FieldRuleBuilder) Realm(realm *RealmBuilder) *FieldRuleBuilder {
	v := realm.Build()
	rb.v.Realm = &v
	return rb
}

func (rb *FieldRuleBuilder) Username(username Name) *FieldRuleBuilder {
	rb.v.Username = &username
	return rb
}