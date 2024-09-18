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

package types

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"strconv"

	"github.com/StackAdapt/go-elasticsearch/v8/typedapi/types/enums/shardroutingstate"
)

// ShardRouting type.
//
// https://github.com/elastic/elasticsearch-specification/blob/8e91c0692c0235474a0c21bb7e9716a8430e8533/specification/indices/stats/types.ts#L162-L167
type ShardRouting struct {
	Node           string                              `json:"node"`
	Primary        bool                                `json:"primary"`
	RelocatingNode *string                             `json:"relocating_node,omitempty"`
	State          shardroutingstate.ShardRoutingState `json:"state"`
}

func (s *ShardRouting) UnmarshalJSON(data []byte) error {

	dec := json.NewDecoder(bytes.NewReader(data))

	for {
		t, err := dec.Token()
		if err != nil {
			if errors.Is(err, io.EOF) {
				break
			}
			return err
		}

		switch t {

		case "node":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return fmt.Errorf("%s | %w", "Node", err)
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.Node = o

		case "primary":
			var tmp any
			dec.Decode(&tmp)
			switch v := tmp.(type) {
			case string:
				value, err := strconv.ParseBool(v)
				if err != nil {
					return fmt.Errorf("%s | %w", "Primary", err)
				}
				s.Primary = value
			case bool:
				s.Primary = v
			}

		case "relocating_node":
			var tmp json.RawMessage
			if err := dec.Decode(&tmp); err != nil {
				return fmt.Errorf("%s | %w", "RelocatingNode", err)
			}
			o := string(tmp[:])
			o, err = strconv.Unquote(o)
			if err != nil {
				o = string(tmp[:])
			}
			s.RelocatingNode = &o

		case "state":
			if err := dec.Decode(&s.State); err != nil {
				return fmt.Errorf("%s | %w", "State", err)
			}

		}
	}
	return nil
}

// NewShardRouting returns a ShardRouting.
func NewShardRouting() *ShardRouting {
	r := &ShardRouting{}

	return r
}
