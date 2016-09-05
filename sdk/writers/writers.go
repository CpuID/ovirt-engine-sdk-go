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

package ovirtsdk4

import (
	"time"
)

// TODO: do we want to mandate some functions on any type we know implements XmlWriter?
type XmlWriter interface{}

//
// This is the base functions for all the XML writers used by the SDK. It contains the utility methods used by all of them.
//

// Writes an element with the given name and string value.
func writerWriteString(writer *XmlWriter, name string, value string) {
	// TODO: implement
}

// Converts the given boolean value to an string.
func writerRenderBoolean(value bool) string {
	// TODO: implement
}

// Writes an element with the given name and boolean value.
func writerWriteBoolean(writer *XmlWriter, name string, value string) {
	// TODO: implement
}

// Converts the given integer value to an string.
func writerRenderInteger(value int) string {
	// TODO: implement
}

// Writes an element with the given name and integer value.
func writerWriteInteger(writer *XmlWriter, name string, value int) {
	// TODO: implement
}

// Converts the given decimal value to an string.
// TODO: int or float?
func writerRenderDecimal(value int) string {
	// TODO: implement
}

// Writes an element with the given name and decimal value.
// TODO: int or float?
func writerWriteDecimal(writer *XmlWriter, name string, value int) {
	// TODO: implement
}

// Converts the given date value to an string.
func writerRenderDate(value *time.Time) string {
	// TODO: implement
}

// Writes an element with the given name and date value.
func writerWriteDate(writer *XmlWriter, name string, value time.Time) {
	// TODO: implement
}
