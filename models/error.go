package models

import (
	"fmt"
	"net/http"
)

// カスタムエラー構造体
type AppError struct {
	Code    int    `json:"code"`    // HTTPステータスコード
	Message string `json:"message"` // エラーメッセージ
	Detail  string `json:"detail"`  // エラー詳細
}

// Errorメソッドを実装して、エラーメッセージを返す
func (e *AppError) Error() string {
	return fmt.Sprintf("code: %d, message: %s, detail: %s", e.Code, e.Message, e.Detail)
}

// 新しいエラーを作成
func NewAppError(code int, message string, detail string) *AppError {
	return &AppError{
		Code:    code,
		Message: message,
		Detail:  detail,
	}
}

// 404 Not Found
func NotFoundError(message string, detail string) *AppError {
	return NewAppError(http.StatusNotFound, message, detail)
}

// 400 Bad Request
func BadRequestError(message string, detail string) *AppError {
	return NewAppError(http.StatusBadRequest, message, detail)
}

// 500 Internal Server Error
func InternalServerError(message string, detail string) *AppError {
	return NewAppError(http.StatusInternalServerError, message, detail)
}
