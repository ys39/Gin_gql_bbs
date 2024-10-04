/*
* エントリーポイント
 */

package main

import (
	"bbs-gql-project/routers"
)

func main() {
	r := routers.SetupRouter()
	r.Run(":8080")
}
