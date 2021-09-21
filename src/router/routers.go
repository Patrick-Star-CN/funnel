package router

import (
	canteen_controller "funnel/app/controller/http/canteen"
	library_controller "funnel/app/controller/http/library"
	schoolCardController "funnel/app/controller/http/schoolcard"
	"funnel/app/controller/http/zfController"
	"funnel/app/middleware"
	"github.com/gin-gonic/gin"
)

func SetupRouter(r *gin.Engine) *gin.Engine {

	student := r.Group("/student")
	{
		zf := student.Group("/zf", middleware.CheckUsernamePassword)
		{
			term := zf.Group("", middleware.CheckTermInfoForm)
			{
				term.POST("/score/info", zfController.GetScoreDetail)
				term.POST("/score", zfController.GetScore)
				term.POST("/table", zfController.GetClassTable)
				term.POST("/exam", zfController.GetExamInfo)
			}
			zf.POST("/room", zfController.GetRoomInfo)
			zf.POST("/program", zfController.GetProgInfo)
		}
		library := student.Group("/library", middleware.CheckUsernamePassword)
		{
			library.POST("/borrow/history", library_controller.LibraryBorrowHistory)
			library.POST("/borrow/current", library_controller.LibraryCurrentBorrow)
			library.POST("/borrow/reborrow", library_controller.LibraryReBorrow)
		}
		card := student.Group("/card", middleware.CheckUsernamePassword)
		{
			card.POST("/balance", schoolCardController.CardBalance)
			card.POST("/today", schoolCardController.CardToday)
			card.POST("/history", schoolCardController.CardHistory)
		}
	}
	canteen := r.Group("/canteen")
	{
		canteen.GET("/flow", canteen_controller.Flow) // 关于餐厅客流量的路由
	}

	return r
}
