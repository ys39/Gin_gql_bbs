/*
* ルーティングの設定を行う
 */

package routers

import (
	"bbs-project/controllers"

	"github.com/gin-gonic/gin"
)

// ルーティングの設定
func SetupRouter() *gin.Engine {
	r := gin.Default()

	// /v1/api に関連するエンドポイントをグループ化
	api := r.Group("/v1/api")
	{
		api.GET("/list", controllers.GetPosts)
		api.GET("/detail/:id", controllers.GetPostDetail)
		api.POST("/create", controllers.CreatePost)
		api.DELETE("/delete/:id", controllers.DeletePost)
		api.PUT("/update/:id", controllers.UpdatePost)
	}

	return r
}
