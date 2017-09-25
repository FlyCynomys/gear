package error

import "github.com/FlyCynomys/tools/err"

var (
	ErrorUserNotExist  = err.New(40001, "user not exist")
	ErrorUserHasExist  = err.New(40002, "user not exist")
	ErrorPasswordError = err.New(40003, "password error")
)
