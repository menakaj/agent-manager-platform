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

package api

import (
	"context"
	"net/http"
	"time"

	"github.com/wso2-enterprise/agent-management-platform/agent-manager-service/config"
	"github.com/wso2-enterprise/agent-management-platform/agent-manager-service/db"
	"github.com/wso2-enterprise/agent-management-platform/agent-manager-service/utils"
)

func registerHealthCheck(mux *http.ServeMux) {
	mux.HandleFunc("GET /healthz", func(w http.ResponseWriter, r *http.Request) {
		ctx, cancel := context.WithTimeout(r.Context(), time.Duration(config.GetConfig().HealthCheckTimeoutSeconds)*time.Second)
		defer cancel()

		var dbRes *int
		if result := db.DB(ctx).Raw("SELECT 1").Scan(&dbRes); result.Error != nil {
			utils.WriteErrorResponse(w, http.StatusInternalServerError, "database connection error")
			return
		}
		response := map[string]interface{}{
			"message":   "agent-manager-service is healthy",
			"timestamp": time.Now(),
		}
		utils.WriteSuccessResponse(w, http.StatusOK, response)
	})
}
