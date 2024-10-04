package resolver_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"bbs-gql-project/routers"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

// TestInit はテストの初期化を行う関数
func setupTestRouter() (*gin.Engine, *httptest.ResponseRecorder) {
	// Ginのテスト用モードを設定
	gin.SetMode(gin.TestMode)

	// ルーターを初期化
	r := routers.SetupRouter()

	// テスト用のレスポンスレコーダーを作成
	w := httptest.NewRecorder()

	return r, w
}

// 投稿作成のテスト
func TestCreatePost(t *testing.T) {
	r, w := setupTestRouter()

	// GraphQLのミューテーションクエリで投稿を作成
	newPost := map[string]interface{}{
		"query": `
			mutation {
				createPost(input: {title: "Test Post", content: "This is a test post"}) {
					id
					title
					content
				}
			}
		`,
	}

	jsonValue, _ := json.Marshal(newPost)
	req, _ := http.NewRequest("POST", "/v1/gql/query", bytes.NewBuffer(jsonValue))
	req.Header.Set("Content-Type", "application/json")

	r.ServeHTTP(w, req)

	// ステータスコードの確認
	assert.Equal(t, http.StatusOK, w.Code)

	// レスポンスの内容を確認
	var response map[string]interface{}
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.Nil(t, err)

	data := response["data"].(map[string]interface{})
	createdPost := data["createPost"].(map[string]interface{})
	assert.Equal(t, "Test Post", createdPost["title"])
	assert.Equal(t, "This is a test post", createdPost["content"])
}

// 投稿取得のテスト
func TestGetPost(t *testing.T) {
	r, w := setupTestRouter()

	// GraphQLのクエリで投稿を取得
	query := map[string]interface{}{
		"query": `
			query {
				getPost(id: "5") {
					id
					title
					content
				}
			}
		`,
	}

	jsonValue, _ := json.Marshal(query)
	req, _ := http.NewRequest("POST", "/v1/gql/query", bytes.NewBuffer(jsonValue))
	req.Header.Set("Content-Type", "application/json")

	r.ServeHTTP(w, req)

	// ステータスコードの確認
	assert.Equal(t, http.StatusOK, w.Code)

	// レスポンスの内容を確認
	var response map[string]interface{}
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.Nil(t, err)

	data := response["data"].(map[string]interface{})
	post := data["getPost"].(map[string]interface{})
	assert.Equal(t, "5", post["id"])
	assert.Equal(t, "投稿5", post["title"])
	assert.Equal(t, "サンプル投稿5", post["content"])
}

// 投稿一覧取得のテスト
func TestGetAllPosts(t *testing.T) {
	r, w := setupTestRouter()

	// GraphQLのクエリで投稿一覧を取得
	query := map[string]interface{}{
		"query": `
			query {
				getAllPosts(page: 1, per_page: 10) {
					id
					title
					content
				}
			}
		`,
	}

	jsonValue, _ := json.Marshal(query)
	req, _ := http.NewRequest("POST", "/v1/gql/query", bytes.NewBuffer(jsonValue))
	req.Header.Set("Content-Type", "application/json")

	r.ServeHTTP(w, req)

	// ステータスコードの確認
	assert.Equal(t, http.StatusOK, w.Code)

	// レスポンスの内容を確認
	var response map[string]interface{}
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.Nil(t, err)

	data := response["data"].(map[string]interface{})
	posts := data["getAllPosts"].([]interface{})

	// 期待する数の投稿があることを確認
	assert.NotEmpty(t, posts)
}

// 投稿の更新テスト
func TestUpdatePost(t *testing.T) {
	r, w := setupTestRouter()

	// GraphQLのミューテーションで投稿を更新
	updatePost := map[string]interface{}{
		"query": `
			mutation {
				updatePost(id: "1", input: {title: "Updated Title", content: "Updated Content"}) {
					id
					title
					content
				}
			}
		`,
	}

	jsonValue, _ := json.Marshal(updatePost)
	req, _ := http.NewRequest("POST", "/v1/gql/query", bytes.NewBuffer(jsonValue))
	req.Header.Set("Content-Type", "application/json")

	r.ServeHTTP(w, req)

	// ステータスコードの確認
	assert.Equal(t, http.StatusOK, w.Code)

	// レスポンスの内容を確認
	var response map[string]interface{}
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.Nil(t, err)

	data := response["data"].(map[string]interface{})
	updatedPost := data["updatePost"].(map[string]interface{})
	assert.Equal(t, "Updated Title", updatedPost["title"])
	assert.Equal(t, "Updated Content", updatedPost["content"])
}

// 投稿削除のテスト
func TestDeletePost(t *testing.T) {
	r, w := setupTestRouter()

	// GraphQLのミューテーションで投稿を削除
	deletePost := map[string]interface{}{
		"query": `
			mutation {
				deletePost(id: "1")
			}
		`,
	}

	jsonValue, _ := json.Marshal(deletePost)
	req, _ := http.NewRequest("POST", "/v1/gql/query", bytes.NewBuffer(jsonValue))
	req.Header.Set("Content-Type", "application/json")

	r.ServeHTTP(w, req)

	// ステータスコードの確認
	assert.Equal(t, http.StatusOK, w.Code)

	// レスポンスの内容を確認
	var response map[string]interface{}
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.Nil(t, err)

	data := response["data"].(map[string]interface{})
	assert.True(t, data["deletePost"].(bool))
}
