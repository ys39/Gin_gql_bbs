/*
* エントリーポイント
 */

package main

import (
	"bbs-project/routers"
)

func main() {
	r := routers.SetupRouter()
	r.Run(":8080")
}
