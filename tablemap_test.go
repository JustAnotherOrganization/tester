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

package tester_test

import (
	"fmt"
	"testing"

	"justanother.org/tester"
)

var ErrEmpty = fmt.Errorf("Empty Path")

// TODO:
// - Make a test table to test our test table.
func TestTableMapTester(t *testing.T) {
	tester.TableMap(t,
		map[string]TestA{
			"Basic Happy Path": {
				"test",
				"test:test",
				nil,
			},
			"Empty Path": {
				"",
				"",
				ErrEmpty,
			},
			"Weird Characters Path": {
				":*''",
				":*''::*''",
				nil,
			},
		},
		func(t *testing.T, name string, test TestA) {
			got, err := TableDriven(test.In)
			if test.Expected != got {
				t.Fatalf("Expected test to be %#v, got %#v", got, test.Expected)
			}
			if test.Err != err {
				t.Fatalf("Expected err to be %#v, got %#v", test.Err, err)
			}
		},
	)
}

func TableDriven(str string) (string, error) {
	if str == "" {
		return "", ErrEmpty
	}
	return fmt.Sprintf("%s:%s", str, str), nil
}

// TestA ...
type TestA struct {
	In       string
	Expected string
	Err      error
}
