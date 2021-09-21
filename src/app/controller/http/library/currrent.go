package library_controller

import (
	controller "funnel/app/controller/http"
	"funnel/app/errors"
	"funnel/app/service/libraryService"
	"funnel/app/utils"

	"github.com/gin-gonic/gin"
)

// @Summary 图书馆当前借书记录
// @Description 图书馆当前借书记录
// @Tags 图书馆
// @Produce  json
// @Param username body string true "用户名"
// @Param password body string true "密码"
// @Success 200 json  {"code":200,"data":[{...}],"msg":"OK"}
// @Failure 400 json  {"code":400,"data":null,"msg":""}
// @Router /student/libraryService/current [post]
func LibraryCurrentBorrow(context *gin.Context) {
	user, err := controller.LoginHandle(context, libraryService.GetUser)
	if err != nil {
		return
	}
	books, err := libraryService.GetCurrentBorrow(user)

	if err == errors.ERR_Session_Expired {
		user, err = controller.LoginHandle(context, libraryService.GetUser)
		if err != nil {
			controller.ErrorHandle(context, err)
			return
		}
		books, err = libraryService.GetCurrentBorrow(user)
	}

	if err != nil {
		controller.ErrorHandle(context, err)
		return
	}
	utils.ContextDataResponseJson(context, utils.SuccessResponseJson(books))
}
