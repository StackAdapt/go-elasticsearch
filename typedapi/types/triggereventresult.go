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

// TriggerEventResult type.
//
// https://github.com/elastic/elasticsearch-specification/blob/135ae054e304239743b5777ad8d41cb2c9091d35/specification/watcher/_types/Trigger.ts#L39-L43
type TriggerEventResult struct {
	Manual        TriggerEventContainer `json:"manual"`
	TriggeredTime DateString            `json:"triggered_time"`
	Type          string                `json:"type"`
}

// TriggerEventResultBuilder holds TriggerEventResult struct and provides a builder API.
type TriggerEventResultBuilder struct {
	v *TriggerEventResult
}

// NewTriggerEventResult provides a builder for the TriggerEventResult struct.
func NewTriggerEventResultBuilder() *TriggerEventResultBuilder {
	r := TriggerEventResultBuilder{
		&TriggerEventResult{},
	}

	return &r
}

// Build finalize the chain and returns the TriggerEventResult struct
func (rb *TriggerEventResultBuilder) Build() TriggerEventResult {
	return *rb.v
}

func (rb *TriggerEventResultBuilder) Manual(manual *TriggerEventContainerBuilder) *TriggerEventResultBuilder {
	v := manual.Build()
	rb.v.Manual = v
	return rb
}

func (rb *TriggerEventResultBuilder) TriggeredTime(triggeredtime DateString) *TriggerEventResultBuilder {
	rb.v.TriggeredTime = triggeredtime
	return rb
}

func (rb *TriggerEventResultBuilder) Type_(type_ string) *TriggerEventResultBuilder {
	rb.v.Type = type_
	return rb
}