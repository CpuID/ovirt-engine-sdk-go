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

package services

//
// This is the base functions for all the services of the SDK. It contains the utility methods used by all of them.
//

// The options permitted on any service remove operation.
type RemoveOpts struct {
	Async bool // Indicates if the remove should be performed asynchronously.
}

// Creates and raises an error containing the details of the given HTTP response and fault.
//
// This method is intended for internal use by other components of the SDK. Refrain from using it directly, as
// backwards compatibility isn't guaranteed.
func raiseError(response TODO, fault TODO) {
	// TODO: implement
}

// Reads the response body assuming that it contains a fault message, converts it to an Error and raises it.
//
// This method is intended for internal use by other components of the SDK. Refrain from using it directly, as
// backwards compatibility isn't guaranteed.
func checkFault(response TODO) {
	// TODO: implement
}

// Reads the response body assuming that it contains an action, checks if it contains a fault message, and if it
// does converts it to an Error and raises it. If it doesn't contain a fault then it just returns the action object.
//
// This method is intended for internal use by other components of the SDK. Refrain from using it directly, as
// backwards compatibility isn't guaranteed.
func checkAction(response TODO) {
	// TODO: implement
}
