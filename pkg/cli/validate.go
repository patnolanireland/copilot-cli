// Copyright 2019 Amazon.com, Inc. or its affiliates. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0

package cli

import (
	"errors"
	"unicode"
)

var (
	errValueEmpty           = errors.New("value must not be empty")
	errValueTooLong         = errors.New("value must not exceed 255 characters")
	errValueNotAlphanumeric = errors.New("value must be alphanumeric: [A-Za-z0-9]")
	errValueNotAString      = errors.New("value must be a string")
)

func validateProjectName(val interface{}) error {
	// TODO(nick): add logic to determine project name uniqueness in the scope of an AWS account
	s, ok := val.(string)
	if !ok {
		return errValueNotAString
	}
	if s == "" {
		return errValueEmpty
	}
	if len(s) > 255 {
		return errValueTooLong
	}
	if !isAlphanumeric(s) {
		return errValueNotAlphanumeric
	}
	return nil
}

func validateApplicationName(val interface{}) error {
	// TODO(nick): add logic to determine application name uniqeness in the scope of a project
	s, ok := val.(string)
	if !ok {
		return errValueNotAString
	}
	if s == "" {
		return errValueEmpty
	}
	if len(s) > 255 {
		return errValueTooLong
	}
	if !isAlphanumeric(s) {
		return errValueNotAlphanumeric
	}
	return nil
}

func isAlphanumeric(s string) bool {
	for _, r := range s {
		if !(unicode.IsLetter(r) || unicode.IsNumber(r)) {
			return false
		}
	}
	return true
}