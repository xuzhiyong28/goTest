package errors

import (
	"bytes"
	"fmt"
	"strings"
)

type ErrorType string

// 错误类型常量
const (
	ERROR_TYPE_DOWNLOADER ErrorType = "downloader error"
	// ERROR_TYPE_ANALYZER 代表分析器错误。
	ERROR_TYPE_ANALYZER ErrorType = "analyzer error"
	// ERROR_TYPE_PIPELINE 代表条目处理管道错误。
	ERROR_TYPE_PIPELINE ErrorType = "pipeline error"
	// ERROR_TYPE_SCHEDULER 代表调度器错误。
	ERROR_TYPE_SCHEDULER ErrorType = "scheduler error"
)

type CrawlerError interface {
	// Type 用于获得错误的类型。
	Type() ErrorType
	// Error 用于获得错误提示信息。
	Error() string
}

type myCrawlerError struct {
	// errType 代表错误的类型
	errType ErrorType
	// errMsg 代表错误的提示信息。
	errMsg string
	// fullErrMsg 代表完整的错误提示信息。
	fullErrMsg string
}

func (ce *myCrawlerError) Type() ErrorType {
	return ce.errType
}

func (ce *myCrawlerError) Error() string {
	if ce.fullErrMsg == "" {
		ce.genFullErrMsg()
	}
	return ce.fullErrMsg
}

// genFullErrMsg 用于生成错误提示信息，并给相应的字段赋值
func (ce *myCrawlerError) genFullErrMsg() {
	var buffer bytes.Buffer
	buffer.WriteString("crawler error: ")
	if ce.errType != "" {
		buffer.WriteString(string(ce.errType))
		buffer.WriteString(": ")
	}
	buffer.WriteString(ce.errMsg)
	ce.fullErrMsg = fmt.Sprintf("%s", buffer.String())
	return
}

// IllegalParameterError 代表非法的参数的错误类型。
type IllegalParameterError struct {
	msg string
}

// NewIllegalParameterError 会创建一个IllegalParameterError类型的实例。
func NewIllegalParameterError(errMsg string) IllegalParameterError {
	return IllegalParameterError{
		msg: fmt.Sprintf("illegal parameter: %s", strings.TrimSpace(errMsg)),
	}
}

func (ipe IllegalParameterError) Error() string {
	return ipe.msg
}
