/*
* モデル層
* データベースとのやり取りを行う(簡易的にサンプルデータを使用)
* データ構造体を定義する
 */

package models

// 投稿データ構造体を定義する
// 「タグ」機能を用いることで、構造体のフィールドとJSONデータの間で変換を行う
type Post struct {
	ID      int    `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
}

// サンプルデータ
// 本来はデータベースから取得するが、ここでは簡易的にサンプルデータを使用
var Posts = []Post{
	{ID: 1, Title: "投稿1", Content: "サンプル投稿1"},
	{ID: 2, Title: "投稿2", Content: "サンプル投稿2"},
	{ID: 3, Title: "投稿3", Content: "サンプル投稿3"},
	{ID: 4, Title: "投稿4", Content: "サンプル投稿4"},
	{ID: 5, Title: "投稿5", Content: "サンプル投稿5"},
	{ID: 6, Title: "投稿6", Content: "サンプル投稿6"},
	{ID: 7, Title: "投稿7", Content: "サンプル投稿7"},
	{ID: 8, Title: "投稿8", Content: "サンプル投稿8"},
	{ID: 9, Title: "投稿9", Content: "サンプル投稿9"},
	{ID: 10, Title: "投稿10", Content: "サンプル投稿10"},
}
