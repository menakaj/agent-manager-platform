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

package utils

import (
	"fmt"
	"strconv"

	"github.com/google/uuid"
)

func StrPointerAsStr(v *string, defaultValue string) string {
	if v == nil {
		return defaultValue
	}
	return *v
}

func ParseUUID(s string) (uuid.UUID, error) {
	id, err := uuid.Parse(s)
	if err != nil {
		return uuid.Nil, fmt.Errorf("invalid UUID %q: %w", s, err)
	}
	return id, nil
}

func BoolAsString(v bool) string {
	return strconv.FormatBool(v)
}
