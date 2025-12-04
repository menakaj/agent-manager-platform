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

import "errors"

var (
	ErrProjectNotFound           = errors.New("project not found")
	ErrAgentAlreadyExists        = errors.New("agent already exists")
	ErrAgentNotFound             = errors.New("agent not found")
	ErrOrganizationNotFound      = errors.New("organization not found")
	ErrBuildNotFound             = errors.New("build not found")
	ErrEnvironmentNotFound       = errors.New("environment not found")
	ErrOrganizationAlreadyExists = errors.New("organization already exists")
	ErrProjectAlreadyExists      = errors.New("project already exists")
)
