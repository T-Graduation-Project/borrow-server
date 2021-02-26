package router

import (
	"github.com/T-Graduation-Project/borrow-server/app/api"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
)

// 你可以将路由注册放到一个文件中管理，
// 也可以按照模块拆分到不同的文件中管理，
// 但统一都放到router目录下。
func init() {
	s := g.Server()
	s.Group("/", func(group *ghttp.RouterGroup) {
		group.ALL("/borrow/getBookList", api.Borrow.GetBookList)
		group.ALL("/borrow/borrowBook", api.Borrow.BorrowBook)
	})
}
