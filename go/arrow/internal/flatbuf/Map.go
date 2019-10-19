// Licensed to the Apache Software Foundation (ASF) under one
// or more contributor license agreements.  See the NOTICE file
// distributed with this work for additional information
// regarding copyright ownership.  The ASF licenses this file
// to you under the Apache License, Version 2.0 (the
// "License"); you may not use this file except in compliance
// with the License.  You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package flatbuf

import (
	flatbuffers "github.com/google/flatbuffers/go"
)

/// A Map is a logical nested type that is represented as
///
/// List<entry: Struct<key: K, value: V>>
///
/// In this layout, the keys and values are each respectively contiguous. We do
/// not constrain the key and value types, so the application is responsible
/// for ensuring that the keys are hashable and unique. Whether the keys are sorted
/// may be set in the metadata for this field
///
/// In a Field with Map type, the Field has a child Struct field, which then
/// has two children: key type and the second the value type. The names of the
/// child fields may be respectively "entry", "key", and "value", but this is
/// not enforced
///
/// Map
///   - child[0] entry: Struct
///     - child[0] key: K
///     - child[1] value: V
///
/// Neither the "entry" field nor the "key" field may be nullable.
///
/// The metadata is structured so that Arrow systems without special handling
/// for Map can make Map an alias for List. The "layout" attribute for the Map
/// field must have the same contents as a List.
type Map struct {
	_tab flatbuffers.Table
}

func GetRootAsMap(buf []byte, offset flatbuffers.UOffsetT) *Map {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &Map{}
	x.Init(buf, n+offset)
	return x
}

func (rcv *Map) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *Map) Table() flatbuffers.Table {
	return rcv._tab
}

/// Set to true if the keys within each value are sorted
func (rcv *Map) KeysSorted() bool {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.GetBool(o + rcv._tab.Pos)
	}
	return false
}

/// Set to true if the keys within each value are sorted
func (rcv *Map) MutateKeysSorted(n bool) bool {
	return rcv._tab.MutateBoolSlot(4, n)
}

func MapStart(builder *flatbuffers.Builder) {
	builder.StartObject(1)
}
func MapAddKeysSorted(builder *flatbuffers.Builder, keysSorted bool) {
	builder.PrependBoolSlot(0, keysSorted, false)
}
func MapEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
