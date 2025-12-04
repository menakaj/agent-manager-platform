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

package openchoreosvc

import "time"

type AgentComponent struct {
	Name             string     `json:"name"`
	DisplayName      string     `json:"displayName,omitempty"`
	Description      string     `json:"description,omitempty"`
	ProjectName      string     `json:"projectName"`
	CreatedAt        time.Time  `json:"createdAt"`
	Status           string     `json:"status,omitempty"`
	Repository       Repository `json:"buildConfig,omitempty"`
	BuildTemplateRef string     `json:"buildTemplateRef,omitempty"`
}

type Repository struct {
	RepoURL string `json:"repoURL"`
	Branch  string `json:"branch,omitempty"`
	AppPath string `json:"appPath,omitempty"`
}
