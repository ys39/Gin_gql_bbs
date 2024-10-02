/*
* ビュー層
* REST APIのレスポンスをビューとしてここに記述する
 */

package views

import (
	"bbs-project/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// 投稿一覧のレンダリング
func RenderPosts(c *gin.Context, posts []models.Post) {
	c.JSON(http.StatusOK, posts)
}

// 特定投稿のレンダリング
func RenderPost(c *gin.Context, post models.Post) {
	c.JSON(http.StatusOK, post)
}

// 成功レスポンスのレンダリング
func RenderSuccess(c *gin.Context, message string) {
	c.JSON(http.StatusOK, gin.H{
		"message": message,
	})
}

// エラーレスポンスのレンダリング
func RenderError(c *gin.Context, code int, message string, detail string) {
	c.JSON(code, gin.H{
		"code":    code,
		"message": message,
		"detail":  detail,
	})
}
