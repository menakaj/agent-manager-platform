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

type EndpointType string

const (
	EndpointTypeDefault EndpointType = "DEFAULT"
	EndpointTypeCustom  EndpointType = "CUSTOM"
)

type AgentType string

const (
	InternalAgent AgentType = "internal"
	ExternalAgent AgentType = "external"
)

type ResourceType string

const (
	ResourceTypeAgent   ResourceType = "agent"
	ResourceTypeProject ResourceType = "project"
)

// Name generation constants
const (
	MaxResourceNameLength     = 25
	RandomSuffixLength        = 2
	ValidCandidateLength      = MaxResourceNameLength - RandomSuffixLength - 1 // 1 for hyphen
	MaxNameGenerationAttempts = 10                                             // Prevent infinite loop
	NameGenerationAlphabet    = "abcdefghijklmnopqrstuvwxyz"
)

type SupportedLanguages string

const (
	LanguageJava      SupportedLanguages = "java"
	LanguagePython    SupportedLanguages = "python"
	LanguageNodeJS    SupportedLanguages = "nodejs"
	LanguageGo        SupportedLanguages = "go"
	LanguagePHP       SupportedLanguages = "php"
	LanguageRuby      SupportedLanguages = "ruby"
	LanguageDotNet    SupportedLanguages = "dotnet"
	LanguageBallerina SupportedLanguages = "ballerina"
)

// Buildpack represents the configuration for a buildpack
type Buildpack struct {
	SupportedVersions  string `json:"supportedVersions"`
	DisplayName        string `json:"displayName"`
	Language           string `json:"language"`
	VersionEnvVariable string `json:"versionEnvVariable"`
	Provider           string `json:"provider"`
}

// Buildpacks contains all supported buildpack configurations
var Buildpacks = []Buildpack{
	{
		SupportedVersions:  "8,11,17,21",
		DisplayName:        "Java",
		Language:           "java",
		VersionEnvVariable: "GOOGLE_RUNTIME_VERSION",
		Provider:           "Google",
	},
	{
		SupportedVersions:  "3.10.x,3.11.x,3.12.x,3.13.x",
		DisplayName:        "Python",
		Language:           "python",
		VersionEnvVariable: "GOOGLE_PYTHON_VERSION",
		Provider:           "Google",
	},
	{
		SupportedVersions:  "12.x.x,14.x.x,16.x.x,18.x.x,20.x.x,22.x.x,24.x.x",
		DisplayName:        "NodeJS",
		Language:           "nodejs",
		VersionEnvVariable: "GOOGLE_NODEJS_VERSION",
		Provider:           "Google",
	},
	{
		SupportedVersions:  "1.x",
		DisplayName:        "Go",
		Language:           "go",
		VersionEnvVariable: "GOOGLE_GO_VERSION",
		Provider:           "Google",
	},
	{
		SupportedVersions:  "8.1.x,8.2.x,8.3.x,8.4.x",
		DisplayName:        "PHP",
		Language:           "php",
		VersionEnvVariable: "GOOGLE_RUNTIME_VERSION",
		Provider:           "Google",
	},
	{
		SupportedVersions:  "3.1.x,3.2.x,3.3.x,3.4.x",
		DisplayName:        "Ruby",
		Language:           "ruby",
		VersionEnvVariable: "GOOGLE_RUNTIME_VERSION",
		Provider:           "Google",
	},
	{
		SupportedVersions:  "6.x,7.x,8.x",
		DisplayName:        ".NET",
		Language:           "dotnet",
		VersionEnvVariable: "GOOGLE_RUNTIME_VERSION",
		Provider:           "Google",
	},
	{
		SupportedVersions:  "",
		DisplayName:        "Ballerina",
		Language:           "ballerina",
		VersionEnvVariable: "",
		Provider:           "AMP-Ballerina",
	},
}
