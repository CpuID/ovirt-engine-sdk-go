//
// Copyright (c) 2015-2016 Red Hat, Inc. / Nathan Sullivan
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//   http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//

package ov_xml

import (
	"encoding/xml"
)

// Port of C ov_xml implementation to native Go. Cleaner implementation for the SDK.

// TODO: implement, replicate logic in ov_xml_reader.c

type XmlReader struct{}

// TODO: review types
func NewXmlReader(self string, io string) *XmlReader {
	var reader XmlReader
	// TODO: implement
	return *reader
}

// TODO: review types
func (r *XmlReader) Read() {
	// TODO: implement
}

// TODO: review types
func (r *XmlReader) Forward() {
	// TODO: implement
}

// TODO: review types
func (r *XmlReader) NodeName() {
	// TODO: implement
}

// TODO: review types
func (r *XmlReader) EmptyElement() {
	// TODO: implement
}

// TODO: review types
func (r *XmlReader) GetAttribute() {
	// TODO: implement
}

// TODO: review types
func (r *XmlReader) ReadElement() {
	// TODO: implement
}

// TODO: review types
func (r *XmlReader) ReadElements() {
	// TODO: implement
}

// TODO: review types
func (r *XmlReader) NextElement() {
	// TODO: implement
}

// TODO: required?
// TODO: review types
func (r *XmlReader) Close() {
	// TODO: implement
}
