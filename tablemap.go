// Copyright 2019 JustAnotherOrganization (justanother.org)

// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use these files except in compliance with the License.
// You may obtain a copy of the License at

// 	http://www.apache.org/licenses/LICENSE-2.0

// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package tester

import (
	"reflect"
	"testing"
)

// These are two reflect.Type we will need to reference later
var (
	testingPTRType = reflect.TypeOf(&testing.T{})
	stringType     = reflect.TypeOf("")
)

// TableMap wants the testing object, a map[string]X, and a
// func(t *testing.T, name string, test X)
func TableMap(t *testing.T, tableMap interface{}, tester interface{}) {
	// Get the type of our tableMap, ensure it's a map
	tableMapType := reflect.TypeOf(tableMap)
	if tableMapType.Kind() != reflect.Map {
		t.Fatalf("Improper use of TableMap: Provided tableMap is not a map")
	}

	// XType is the map's value, and the tester's test object
	// TODO:
	// - Support arrays, so we can table our table test.
	XType := tableMapType.Elem()

	// Get the type of our tester function, ensure it's a function
	testerType := reflect.TypeOf(tester)
	if testerType.Kind() != reflect.Func {
		t.Fatalf("Improper use of TableMap: Provided tester is not a func")
	}

	// The tester function needs to have at least 3 parameters.
	if testerType.NumIn() != 3 {
		t.Fatalf("Improper use of TableMap: Provided tester must look like "+
			"func(*testing.T, string, %v", XType)
	}

	// These are the parameters that our tester function should have.
	// Used an array so it's easier to do matching
	expectedParameters := []reflect.Type{
		testingPTRType,
		stringType,
		XType,
	}

	// Check the function parameters against what we're expecting
	for i := 0; i < 3; i++ {
		if testerType.In(i) != expectedParameters[i] {
			// TODO Deuglify this
			t.Fatalf("Improper use of TableMap: \nProvided tester must look "+
				"like func(*testing.T, string, %v)\nGot func(%v, %v, %v) "+
				"instead.",
				XType,
				testerType.In(0),
				testerType.In(1),
				testerType.In(2),
			)
		}
	}

	// We will need to get the values of our objects to interact with them
	tableMapValue := reflect.ValueOf(tableMap)
	testerValue := reflect.ValueOf(tester)
	for _, key := range tableMapValue.MapKeys() {
		// t.Run ensures that we have mapping to the specific test
		t.Run(
			key.String(),
			func(test *testing.T) {
				// Use reflection to call our function
				testerValue.Call([]reflect.Value{
					reflect.ValueOf(test),
					key,
					tableMapValue.MapIndex(key),
				})
			},
		)
	}
}
