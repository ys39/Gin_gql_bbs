/*
* コントローラー層
* リクエストを受け取り、モデル層に処理を依頼し、レスポンスを返す
 */

package controllers

import (
	"bbs-project/models"
	"bbs-project/views"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// 投稿一覧取得のコントローラ
func GetPosts(c *gin.Context) {
	// クエリパラメータの取得
	pageStr := c.Query("page")
	perPageStr := c.Query("per_page")

	// ページ数の確認
	page, err := strconv.Atoi(pageStr)
	if err != nil || page < 1 {
		views.RenderError(c, http.StatusBadRequest, "Invalid page parameter", "The 'page' parameter must be a positive integer.")
		return
	}

	// 1ページあたりの表示数の確認
	perPage, err := strconv.Atoi(perPageStr)
	if err != nil || perPage < 1 {
		views.RenderError(c, http.StatusBadRequest, "Invalid per_page parameter", "The 'per_page' parameter must be a positive integer.")
		return
	}

	start := (page - 1) * perPage
	end := start + perPage

	if start >= len(models.Posts) {
		views.RenderPosts(c, []models.Post{})
		return
	}

	if end > len(models.Posts) {
		end = len(models.Posts)
	}

	views.RenderPosts(c, models.Posts[start:end])
}

// 投稿の詳細取得のコントローラ
func GetPostDetail(c *gin.Context) {
	// パスパラメータの取得
	idStr := c.Param("id")

	// IDの確認
	id, err := strconv.Atoi(idStr)
	if err != nil || id < 1 {
		views.RenderError(c, http.StatusBadRequest, "Invalid ID", "The 'id' parameter must be a positive integer.")
		return
	}

	// models.Postsの中からIDが一致する投稿を探す
	for _, post := range models.Posts {
		if post.ID == id {
			views.RenderPost(c, post)
			return
		}
	}

	views.RenderError(c, http.StatusNotFound, "Resource not found", "The post with ID "+idStr+" was not found.")
}

// 新規投稿作成のコントローラ
func CreatePost(c *gin.Context) {

	var newPost models.Post // 新規投稿データを格納するための構造体
	// リクエストボディのJSONデータを構造体にバインドし、失敗した場合はエラーレスポンスを返す
	if err := c.ShouldBindJSON(&newPost); err != nil {
		views.RenderError(c, http.StatusBadRequest, "Invalid request parameters", err.Error())
		return
	}

	// jsonデータへの追加
	// データベースを利用している場合は、INSERT文を実行する
	newPost.ID = len(models.Posts) + 1
	models.Posts = append(models.Posts, newPost)

	views.RenderPost(c, newPost)
}

// 投稿の削除のコントローラ
func DeletePost(c *gin.Context) {
	// パスパラメータの取得
	idStr := c.Param("id")

	// IDの確認
	id, err := strconv.Atoi(idStr)
	if err != nil || id < 1 {
		views.RenderError(c, http.StatusBadRequest, "Invalid ID", "The 'id' parameter must be a positive integer.")
		return
	}

	// jsonデータの中からIDが一致する投稿を探し、削除する
	// データベースを利用している場合は、DELETE文を実行する
	for i, post := range models.Posts {
		if post.ID == id {
			models.Posts = append(models.Posts[:i], models.Posts[i+1:]...)
			views.RenderSuccess(c, "Post deleted")
			return
		}
	}

	views.RenderError(c, http.StatusNotFound, "Resource not found", "The post with ID "+idStr+" was not found.")
}

// 投稿の更新のコントローラ
func UpdatePost(c *gin.Context) {
	// パスパラメータの取得
	idStr := c.Param("id")

	// IDの確認
	id, err := strconv.Atoi(idStr)
	if err != nil || id < 1 {
		views.RenderError(c, http.StatusBadRequest, "Invalid ID", "The 'id' parameter must be a positive integer.")
		return
	}

	var updatedPost models.Post // 更新後の投稿データを格納するための構造体
	// リクエストボディのJSONデータを構造体にバインドし、失敗した場合はエラーレスポンスを返す
	if err := c.ShouldBindJSON(&updatedPost); err != nil {
		views.RenderError(c, http.StatusBadRequest, "Invalid request parameters", err.Error())
		return
	}

	// jsonデータの中からIDが一致する投稿を探し、更新する
	// データベースを利用している場合は、UPDATE文を実行する
	for i, post := range models.Posts {
		if post.ID == id {
			models.Posts[i].Title = updatedPost.Title
			models.Posts[i].Content = updatedPost.Content
			views.RenderPost(c, models.Posts[i])
			return
		}
	}

	views.RenderError(c, http.StatusNotFound, "Resource not found", "The post with ID "+idStr+" was not found.")
}
