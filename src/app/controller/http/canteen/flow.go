package canteen_controller

import (
	"fmt"
	"funnel/app/errors"
	"funnel/app/service/canteen_service"
	"funnel/app/utils"

	"github.com/gin-gonic/gin"
)

// 餐厅人流量
func Flow(content *gin.Context) {
	data, err := canteen_service.FetchFlow()
	if err == nil {
		utils.ContextDataResponseJson(
			content,
			utils.SuccessResponseJson(data))
	} else {
		utils.ContextDataResponseJson(
			content,
			utils.FailResponseJson(errors.RequestFailed, fmt.Sprintf("%v: %v", data, err)))
	}
}
