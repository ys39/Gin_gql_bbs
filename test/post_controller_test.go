package controllers_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"bbs-project/models"
	"bbs-project/routers"

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

func TestCreatePost(t *testing.T) {
	r, w := setupTestRouter()

	// モックデータを初期化
	models.Posts = []models.Post{
		{ID: 1, Title: "投稿1", Content: "これはサンプル投稿1です"},
	}

	// テスト用の投稿データを作成
	testPost := models.Post{
		Title:   "新しい投稿",
		Content: "これは新しい投稿の内容です",
	}
	testPostJson, _ := json.Marshal(testPost)

	// テスト用のHTTPリクエストを作成
	req, err := http.NewRequest(http.MethodPost, "/v1/api/create", strings.NewReader(string(testPostJson)))
	if err != nil {
		t.Fatalf("リクエストの作成に失敗しました: %v", err)
	}
	req.Header.Set("Content-Type", "application/json")

	// テスト用のHTTPリクエストを実行
	r.ServeHTTP(w, req)

	// レスポンスのステータスコードを確認
	assert.Equal(t, http.StatusOK, w.Code)

	// 期待されるレスポンスを定義
	expectedPost := models.Post{
		ID:      2,
		Title:   "新しい投稿",
		Content: "これは新しい投稿の内容です",
	}
	expectedPostJson, _ := json.Marshal(expectedPost)

	// 期待されるレスポンスと実際のレスポンスが一致することを確認
	assert.JSONEq(t, string(expectedPostJson), w.Body.String())

	// モックデータに新しい投稿が追加されていることを確認
	assert.Equal(t, 2, models.Posts[1].ID)
	assert.Equal(t, "新しい投稿", models.Posts[1].Title)
	assert.Equal(t, "これは新しい投稿の内容です", models.Posts[1].Content)
}

func TestGetPosts(t *testing.T) {
	r, w := setupTestRouter()

	// モックデータを初期化
	models.Posts = []models.Post{
		{ID: 1, Title: "投稿1", Content: "これはサンプル投稿1です"},
		{ID: 2, Title: "投稿2", Content: "これはサンプル投稿2です"},
		{ID: 3, Title: "投稿3", Content: "これはサンプル投稿3です"},
	}

	// テスト用のHTTPリクエストを作成
	req, err := http.NewRequest(http.MethodGet, "/v1/api/list?page=1&per_page=2", nil)
	if err != nil {
		t.Fatalf("リクエストの作成に失敗しました: %v", err)
	}

	// テスト用のHTTPリクエストを実行
	r.ServeHTTP(w, req)

	// レスポンスのステータスコードを確認
	assert.Equal(t, http.StatusOK, w.Code)

	// 期待されるレスポンスを定義
	expectedPosts := []models.Post{
		{ID: 1, Title: "投稿1", Content: "これはサンプル投稿1です"},
		{ID: 2, Title: "投稿2", Content: "これはサンプル投稿2です"},
	}
	expectedPostsJson, _ := json.Marshal(expectedPosts)

	// 期待されるレスポンスと実際のレスポンスが一致することを確認
	assert.JSONEq(t, string(expectedPostsJson), w.Body.String())
}

func TestGetPostDetail(t *testing.T) {
	r, w := setupTestRouter()

	// モックデータを初期化
	models.Posts = []models.Post{
		{ID: 1, Title: "投稿1", Content: "これはサンプル投稿1です"},
		{ID: 2, Title: "投稿2", Content: "これはサンプル投稿2です"},
		{ID: 3, Title: "投稿3", Content: "これはサンプル投稿3です"},
	}

	// テスト用のHTTPリクエストを作成
	req, err := http.NewRequest(http.MethodGet, "/v1/api/detail/2", nil)
	if err != nil {
		t.Fatalf("リクエストの作成に失敗しました: %v", err)
	}

	// テスト用のHTTPリクエストを実行
	r.ServeHTTP(w, req)

	// レスポンスのステータスコードを確認
	assert.Equal(t, http.StatusOK, w.Code)

	// 期待されるレスポンスを定義
	expectedPost := models.Post{
		ID:      2,
		Title:   "投稿2",
		Content: "これはサンプル投稿2です",
	}
	expectedPostJson, _ := json.Marshal(expectedPost)

	// 期待されるレスポンスと実際のレスポンスが一致することを確認
	assert.JSONEq(t, string(expectedPostJson), w.Body.String())
}

func TestDeletePost(t *testing.T) {
	r, w := setupTestRouter()

	// モックデータを初期化
	models.Posts = []models.Post{
		{ID: 1, Title: "投稿1", Content: "これはサンプル投稿1です"},
		{ID: 2, Title: "投稿2", Content: "これはサンプル投稿2です"},
		{ID: 3, Title: "投稿3", Content: "これはサンプル投稿3です"},
	}

	// テスト用のHTTPリクエストを作成
	req, err := http.NewRequest(http.MethodDelete, "/v1/api/delete/2", nil)
	if err != nil {
		t.Fatalf("リクエストの作成に失敗しました: %v", err)
	}

	// テスト用のHTTPリクエストを実行
	r.ServeHTTP(w, req)

	// レスポンスのステータスコードを確認
	assert.Equal(t, http.StatusOK, w.Code)

	// モックデータから削除されたことを確認
	assert.Equal(t, 2, len(models.Posts))
	assert.Equal(t, 1, models.Posts[0].ID)
	assert.Equal(t, 3, models.Posts[1].ID)
}

func TestUpdatePost(t *testing.T) {
	r, w := setupTestRouter()

	// モックデータを初期化
	models.Posts = []models.Post{
		{ID: 1, Title: "投稿1", Content: "これはサンプル投稿1です"},
		{ID: 2, Title: "投稿2", Content: "これはサンプル投稿2です"},
		{ID: 3, Title: "投稿3", Content: "これはサンプル投稿3です"},
	}

	// テスト用の投稿データを作成
	testPost := models.Post{
		Title:   "更新された投稿",
		Content: "これは更新された投稿の内容です",
	}
	testPostJson, _ := json.Marshal(testPost)

	// テスト用のHTTPリクエストを作成
	req, err := http.NewRequest(http.MethodPut, "/v1/api/update/2", strings.NewReader(string(testPostJson)))
	if err != nil {
		t.Fatalf("リクエストの作成に失敗しました: %v", err)
	}
	req.Header.Set("Content-Type", "application/json")

	// テスト用のHTTPリクエストを実行
	r.ServeHTTP(w, req)

	// レスポンスのステータスコードを確認
	assert.Equal(t, http.StatusOK, w.Code)

	// 期待されるレスポンスを定義
	expectedPost := models.Post{
		ID:      2,
		Title:   "更新された投稿",
		Content: "これは更新された投稿の内容です",
	}
	expectedPostJson, _ := json.Marshal(expectedPost)

	// 期待されるレスポンスと実際のレスポンスが一致することを確認
	assert.JSONEq(t, string(expectedPostJson), w.Body.String())

	// モックデータが更新されたことを確認
	assert.Equal(t, 2, models.Posts[1].ID)
	assert.Equal(t, "更新された投稿", models.Posts[1].Title)
	assert.Equal(t, "これは更新された投稿の内容です", models.Posts[1].Content)
}
