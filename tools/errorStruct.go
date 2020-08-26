package tools

import (
	"fmt"
	"time"
)

type errorStruct struct {
	Time time.Time
	Code int
	info string
}

func (e *errorStruct) Error(code int, info string) string {
	e.Time = time.Now()
	e.Code = code
	e.info = info
	return fmt.Sprintf("time:%s code:%d info:%s", e.Time, e.Code, e.info)
}
