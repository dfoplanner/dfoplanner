package npconf

import (
	"strings"
)

// Copyright 2014 Unknwon
//
// Licensed under the Apache License, Version 2.0 (the "License"): you may
// not use this file except in compliance with the License. You may obtain
// a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS, WITHOUT
// WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the
// License for the specific language governing permissions and limitations
// under the License.


// Section represents a config section.
type Section struct {
	name     string
	rawBody  string
	contentArray []string
}

func newSection(name string, rawBody string) *Section {
	return &Section{
		name:     name,
		rawBody:  rawBody,
	}
}

func (s *Section) ParseToArray() {
	// [some section]
	// `whatever`
	// `whatever2`
	// [/some section]
	s.contentArray = strings.Split(s.rawBody, "\n")
}

// Name returns name of Section.
func (s *Section) Name() string {
	return s.name
}
