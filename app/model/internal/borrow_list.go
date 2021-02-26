// ==========================================================================
// This is auto-generated by gf cli tool. DO NOT EDIT THIS FILE MANUALLY.
// ==========================================================================

package internal

import (
	"github.com/gogf/gf/os/gtime"
)

// BorrowList is the golang structure for table borrow_list.
type BorrowList struct {
	Id       int         `orm:"id,primary" json:"id"`       //
	BookId   int         `orm:"book_id"    json:"bookId"`   // 用户账号
	UserId   int         `orm:"user_id"    json:"userId"`   // 用户密码
	LendDate *gtime.Time `orm:"lend_date"  json:"lendDate"` // 借出日期
	BackDate *gtime.Time `orm:"back_date"  json:"backDate"` // 归还日期
}
