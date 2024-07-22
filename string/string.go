// Copyright 2024 Willard Lu
//
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE file or at
// https://opensource.org/licenses/MIT.
package youling_string

import (
	"strings"
)

// Reads a string between two specific strings
//
//		ori: source string
//		left: left string
//		right: right string
//		returns: returns a string between two specific strings
//		Note: The left string is searched from the far left and the right string
//	 is searched from the far right.
func ReadBetween(ori string, left string, right string) string {
	s := strings.Index(ori, left)
	if s == -1 {
		return ""
	}
	s += len(left)
	e := strings.LastIndex(ori[s:], right)
	if e == -1 {
		return ""
	}
	return ori[s:][:e]
}
