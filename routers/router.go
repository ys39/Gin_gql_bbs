/*
* ルーティングの設定を行う
 */

package routers

import (
	"bbs-gql-project/graph"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"

	"github.com/gin-gonic/gin"
)

// GraphQLハンドラを定義
func graphqlHandler() gin.HandlerFunc {
	h := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{}}))

	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}

// Playgroundハンドラを定義
func playgroundHandler() gin.HandlerFunc {
	h := playground.Handler("GraphQL", "/v1/gql/query")

	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}

// ルーティングの設定
func SetupRouter() *gin.Engine {
	r := gin.Default()

	// /v1/gql に関連するエンドポイントをグループ化
	api := r.Group("/v1/gql")
	{
		api.POST("/query", graphqlHandler())
		api.GET("/", playgroundHandler())
	}

	return r
}
