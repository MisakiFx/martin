package constant

const (
	StatusCodeSuccess         = 0
	StatusCodeTokenCheckError = -1
	StatusCodeServiceError    = -2
	StatusCodeInputError      = -3
)

const (
	CodeMessageSuccess         = "成功"
	CodeMessageTokenCheckError = "token校验失败"
	CodeMessageServiceError    = "服务器内部错误"
	CodeMessageInputError      = "输入不合法"
)

var StatusCodeMessageMap = map[int]string{
	StatusCodeSuccess:         CodeMessageSuccess,
	StatusCodeTokenCheckError: CodeMessageTokenCheckError,
	StatusCodeServiceError:    CodeMessageServiceError,
	StatusCodeInputError:      CodeMessageInputError,
}
