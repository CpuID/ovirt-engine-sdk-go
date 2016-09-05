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

package readers

import (
	"time"
)

// TODO: do we want to mandate some functions on any type we know implements XmlReader?
type XmlReader interface{}

//
// This is the base functions for all the readers of the SDK. It contains the utility methods used by all of them.
//

// Reads a string value, assuming that the cursor is positioned at the start element that contains the value.
func ReadString(reader *XmlReader) string {
	// TODO: implement
}

// Reads a list of string values, assuming that the cursor is positioned at the start of the element that contains
// the first value.
func ReadStrings(reader *XmlReader) []string {
	// TODO: implement
}

// Converts the given text to a boolean value.
func ParseBoolean(text *string) bool {
	// TODO: implement
}

// Reads a boolean value, assuming that the cursor is positioned at the start element that contains the value.
func ReadBoolean(reader *XmlReader) bool {
	// TODO: implement
}

// Reads a list of boolean values, assuming that the cursor is positioned at the start element that contains
// the values.
func ReadBooleans(reader *XmlReader) []bool {
	// TODO: implement
}

// Converts the given text to an integer value.
func ParseInteger(text *string) int {
	// TODO: implement
}

// Reads an integer value, assuming that the cursor is positioned at the start element that contains the value.
func ReadInteger(reader *XmlReader) int {
	// TODO: implement
}

// Reads a list of integer values, assuming that the cursor is positioned at the start element that contains the
// values.
func ReadIntegers(reader *XmlReader) []int {
	// TODO: implement
}

// Converts the given text to a decimal value.
// TODO: int or float?
func ParseDecimal(text *string) int {
	// TODO: implement
}

// Reads a decimal value, assuming that the cursor is positioned at the start element that contains the value.
// TODO: int or float?
func ReadDecimal(reader *XmlReader) int {
	// TODO: implement
}

// Reads a list of decimal values, assuming that the cursor is positioned at the start element that contains the
// values.
// TODO: int or float?
func ReadDecimals(reader *XmlReader) []int {
	// TODO: implement
}

// Converts the given text to a date value.
func ParseDate(text *string) time.Time {
	// TODO: implement
}

// Reads a date value, assuming that the cursor is positioned at the start element that contains the value.
func ReadDate(reader *XmlReader) time.Time {
	// TODO: implement
}

// Reads a list of dates values, assuming that the cursor is positioned at the start element that contains the
// values.
func ReadDates(reader *XmlReader) []time.Time {
	// TODO: implement
}