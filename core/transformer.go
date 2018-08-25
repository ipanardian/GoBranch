/*
 * The MIT License (MIT)
 *
 * Copyright (c) 2015 Ian Coleman
 * Copyright (c) 2018 Ma_124, <github.com/Ma124>
 * Copyright (c) 2018 ipanardian <https://github.com/ipanardian>
 *
 * Permission is hereby granted, free of charge, to any person obtaining a copy
 * of this software and associated documentation files (the "Software"), to deal
 * in the Software without restriction, including without limitation the rights
 * to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
 * copies of the Software, and to permit persons to whom the Software is
 * furnished to do so, Subject to the following conditions:
 *
 * The above copyright notice and this permission notice shall be included in all
 * copies or Substantial portions of the Software.
 *
 * THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
 * IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
 * FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
 * AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
 * LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
 * OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
 * SOFTWARE.
 */

package core

import (
	"log"
	"reflect"
	"regexp"
	"strings"
)

// isZero returns true if the passed value is the zero object
func isZero(v reflect.Value) bool {
	switch v.Kind() {
	case reflect.Slice, reflect.Map:
		return v.Len() == 0
	}

	return reflect.DeepEqual(v.Interface(), reflect.Zero(v.Type()).Interface())
}

//TransformInput Trim and ToLower input
func TransformInput(ans interface{}) interface{} {
	if isZero(reflect.ValueOf(ans)) {
		return nil
	}
	ans = strings.Trim(ans.(string), " ")
	ans = strings.ToLower(ans.(string))

	return ans
}

//ToValidBranchName Transform string to valid git branch name
func ToValidBranchName(s string, convention string, isCustom bool) string {
	var pattern string
	if isCustom {
		pattern = `[^a-zA-Z0-9\/\-\_\.\s]+`
	} else {
		pattern = `[^a-zA-Z0-9\s]+`
	}

	reg, err := regexp.Compile(pattern)
	if err != nil {
		log.Fatal(err)
	}
	processedStr := reg.ReplaceAllString(s, " ")
	processedStr = strings.Join(strings.Fields(processedStr), " ")

	var validBranchName string
	switch convention {
	case namingConvention.ToKebab:
		validBranchName = ToKebab(processedStr)
	default:
		validBranchName = ToSnake(processedStr)
	}

	return validBranchName
}

//ToSnake Converts a string to snake_case
func ToSnake(s string) string {
	return ToDelimited(s, '_')
}

//ToScreamingSnake Converts a string to SCREAMING_SNAKE_CASE
func ToScreamingSnake(s string) string {
	return ToScreamingDelimited(s, '_', true)
}

//ToKebab Converts a string to kebab-case
func ToKebab(s string) string {
	return ToDelimited(s, '-')
}

//ToScreamingKebab Converts a string to SCREAMING-KEBAB-CASE
func ToScreamingKebab(s string) string {
	return ToScreamingDelimited(s, '-', true)
}

//ToDelimited Converts a string to delimited.snake.case (in this case `del = '.'`)
func ToDelimited(s string, del uint8) string {
	return ToScreamingDelimited(s, del, false)
}

//ToScreamingDelimited Converts a string to SCREAMING.DELIMITED.SNAKE.CASE (in this case `del = '.'; screaming = true`) or delimited.snake.case (in this case `del = '.'; screaming = false`)
func ToScreamingDelimited(s string, del uint8, screaming bool) string {
	s = strings.Trim(s, " ")
	n := ""
	for i, v := range s {
		// treat acronyms as words, eg for JSONData -> JSON is a whole word
		nextCaseIsChanged := false
		if i+1 < len(s) {
			next := s[i+1]
			if (v >= 'A' && v <= 'Z' && next >= 'a' && next <= 'z') || (v >= 'a' && v <= 'z' && next >= 'A' && next <= 'Z') {
				nextCaseIsChanged = true
			}
		}

		if i > 0 && n[len(n)-1] != del && nextCaseIsChanged {
			// add underscore if next letter case type is changed
			if v >= 'A' && v <= 'Z' {
				n += string(del) + string(v)
			} else if v >= 'a' && v <= 'z' {
				n += string(v) + string(del)
			}
		} else if v == ' ' || v == '_' || v == '-' {
			// replace spaces/underscores with delimiters
			n += string(del)
		} else {
			n = n + string(v)
		}
	}

	if screaming {
		n = strings.ToUpper(n)
	} else {
		n = strings.ToLower(n)
	}
	return n
}
