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
	"context"
	"net/http"
	"slices"
	"time"

	"github.com/hashicorp/go-retryablehttp"
)

var TransientHTTPErrorCodes = []int{
	http.StatusTooManyRequests,    // 429
	http.StatusBadGateway,         // 502
	http.StatusServiceUnavailable, // 503
	http.StatusGatewayTimeout,     // 504
}

var TransientHTTPGETErrorCodes = []int{
	http.StatusTooManyRequests,     // 429
	http.StatusInternalServerError, // 500
	http.StatusBadGateway,          // 502
	http.StatusServiceUnavailable,  // 503
	http.StatusGatewayTimeout,      // 504
}

type RequestRetryConfig struct {
	RetryWaitMin time.Duration
	RetryWaitMax time.Duration
	// RetryAttemptsMax is the maximum number of retries to attempt. 0 for no retries.
	RetryAttemptsMax int
	// AttemptTimeout is the maximum time allowed for a single request attempt.
	AttemptTimeout time.Duration
	// RetryOnStatus is a function that returns true if the request should be retried based on the status code.
	RetryOnStatus func(status int) bool
}

func (cfg RequestRetryConfig) withDefaults(req *HttpRequest) RequestRetryConfig {
	if cfg.RetryWaitMin == 0 {
		cfg.RetryWaitMin = 1 * time.Second
	}
	if cfg.RetryWaitMax == 0 {
		cfg.RetryWaitMax = 10 * time.Second
	}
	if cfg.RetryAttemptsMax == 0 {
		cfg.RetryAttemptsMax = 3
	}
	if cfg.AttemptTimeout == 0 {
		cfg.AttemptTimeout = 3 * time.Minute
	}
	if cfg.RetryOnStatus == nil {
		cfg.RetryOnStatus = func(status int) bool {
			if req.Method == http.MethodGet || req.Method == http.MethodDelete {
				return slices.Contains(TransientHTTPGETErrorCodes, status)
			}
			return slices.Contains(TransientHTTPErrorCodes, status)
		}
	}
	return cfg
}

func (cfg RequestRetryConfig) makeCheckRetry() retryablehttp.CheckRetry {
	return func(ctx context.Context, resp *http.Response, err error) (bool, error) {
		if err != nil { // not nil for network errors, context.DeadlineExceeded etc.
			return retryablehttp.DefaultRetryPolicy(ctx, resp, err)
		}
		if cfg.RetryOnStatus != nil {
			return cfg.RetryOnStatus(resp.StatusCode), nil
		}
		return false, nil
	}
}
