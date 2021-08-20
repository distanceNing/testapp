package common

const (
	ErrSystem = 1
	ErrTokenTimeOut
	ErrUserNotExist
	ErrUserAlreadyExist
	ErrPasswordNotMatch
	ErrTokenNotMatch
	ErrJsonDecodeFail
	ErrRequest

	ErrDbDupKey = 100
	ErrNoAffected
	ErrRecordNotExist

	ErrNeedRetry = 1000
)

type ErrorCode struct {
	code int
	msg  string
}

func (s *ErrorCode) Error() string {
	return s.msg
}

func NewSuccCode() ErrorCode {
	return ErrorCode{0, "ok"}
}

func NewErrorCode(code int, msg string) ErrorCode {
	return ErrorCode{code, msg}
}

func (s *ErrorCode) Ok() bool {
	return s.code == 0
}

func (s *ErrorCode) Code() int {
	return s.code
}

func (s *ErrorCode) Msg() string {
	return s.msg
}

func (s *ErrorCode) Set(code int, msg string) {
	s.code = code
	s.msg = msg
}

//func GetCode(err error) int {
//	errcode := err.(ErrorCode)
//}
