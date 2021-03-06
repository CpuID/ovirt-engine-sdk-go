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

package writers

import (
	"fmt"
	"time"

	"github.com/CpuID/ovirt-engine-sdk-go/sdk/ov_xml"
)

//
// This is the base functions for all the XML writers used by the SDK. It contains the utility methods used by all of them.
//

// Writes an element with the given name and string value.
func WriteString(writer *ov_xml.XmlWriter, name string, value string) {
	// TODO: implement
}

// Converts the given boolean value to an string.
func RenderBoolean(value bool) string {
	return fmt.Sprintf("%t", value)
}

// Writes an element with the given name and boolean value.
func WriteBoolean(writer *ov_xml.XmlWriter, name string, value string) {
	// TODO: implement
}

// Converts the given integer value to an string.
func RenderInteger(value int) string {
	return fmt.Sprintf("%d", value)
}

// Writes an element with the given name and integer value.
func WriteInteger(writer *ov_xml.XmlWriter, name string, value int) {
	// TODO: implement
}

// Converts the given decimal value to an string.
// TODO: int or float?
func RenderDecimal(value float64) string {
	return fmt.Sprintf("%f", value)
}

// Writes an element with the given name and decimal value.
// TODO: int or float?
func WriteDecimal(writer *ov_xml.XmlWriter, name string, value float64) {
	// TODO: implement
}

// Converts the given date value to an string.
func RenderDate(value *time.Time) string {
	// TODO: implement
}

// Writes an element with the given name and date value.
func WriteDate(writer *ov_xml.XmlWriter, name string, value time.Time) {
	// TODO: implement
}
