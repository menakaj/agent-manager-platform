// Copyright (c) 2025, WSO2 LLC (http://www.wso2.com). All Rights Reserved.
//
// This software is the property of WSO2 LLC and its suppliers, if any.
// Dissemination of any information or reproduction of any material contained
// herein is strictly forbidden, unless permitted by WSO2 in accordance with
// the WSO2 Commercial License available at http://wso2.com/licenses.
// For specific language governing the permissions and limitations under
// this license, please see the license as well as any agreement you've
// entered into with WSO2 governing the purchase of this software and any
// associated services.

package config

import (
	"fmt"
	"log/slog"
	"os"
	"strconv"
)

type configReader struct {
	errors []error
}

func (c *configReader) logAndExitIfErrorsFound() {
	if len(c.errors) > 0 {
		var errors []string
		for _, err := range c.errors {
			errors = append(errors, err.Error())
		}
		slog.Error("configReader: errors found while reading config", "errors", errors)
		os.Exit(1)
	}
}

func (c *configReader) readRequiredString(envVarName string) string {
	v := os.Getenv(envVarName)
	if v == "" {
		c.errors = append(c.errors, fmt.Errorf("required environment variable %s not found", envVarName))
	}
	return v
}

func (c *configReader) readOptionalString(envVarName string, defaultValue string) string {
	v := os.Getenv(envVarName)
	if v == "" {
		return defaultValue
	}
	return v
}

func (c *configReader) readOptionalInt64(envVarName string, defaultValue int64) int64 {
	v := os.Getenv(envVarName)
	if v == "" {
		return defaultValue
	}
	value, err := strconv.ParseInt(v, 10, 64)
	if err != nil {
		c.errors = append(c.errors, fmt.Errorf("optional environment variable %s is not a valid integer [%w]", envVarName, err))
		return 0
	}
	return value
}

func (c *configReader) readNullableInt64(envVarName string) *int64 {
	v := os.Getenv(envVarName)
	if v == "" {
		return nil
	}
	value, err := strconv.ParseInt(v, 10, 64)
	if err != nil {
		c.errors = append(c.errors, fmt.Errorf("nullable environment variable %s is not a valid integer [%w]", envVarName, err))
		return nil
	}
	return &value
}

func (c *configReader) readOptionalBool(envVarName string, defaultValue bool) bool {
	v := os.Getenv(envVarName)
	if v == "" {
		return defaultValue
	}
	value, err := strconv.ParseBool(v)
	if err != nil {
		return defaultValue
	}
	return value
}
