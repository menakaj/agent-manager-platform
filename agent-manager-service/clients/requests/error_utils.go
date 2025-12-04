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

package requests

import (
	"fmt"
)

type HttpError struct {
	StatusCode int
	Body       string
	err        error
}

func (e *HttpError) Error() string {
	if e.err != nil {
		return fmt.Sprintf("failed with status code %d and internal error %s", e.StatusCode, e.err)
	}
	if e.Body == "" {
		return fmt.Sprintf("failed with status code %d", e.StatusCode)
	}
	return fmt.Sprintf("failed with status code %d [%s]", e.StatusCode, e.Body)
}
