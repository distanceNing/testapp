package errcode

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
	Code int
	Msg  string
}

func Code(e error) int {
	if e == nil {
		return 0
	}
	err, ok := e.(*ErrorCode)
	if !ok {
		return -1
	}

	if err == (*ErrorCode)(nil) {
		return 0
	}
	return int(err.Code)
}

func Msg(e error) string {
	if e == nil {
		return "ok"
	}
	err, ok := e.(*ErrorCode)
	if !ok {
		return e.Error()
	}

	if err == (*ErrorCode)(nil) {
		return "ok"
	}
	return string(err.Msg)
}

func (s *ErrorCode) Error() string {
	return s.Msg
}

func NewErrorCode(code int, msg string) error {
	return &ErrorCode{code, msg}
}

func (s *ErrorCode) Set(code int, msg string) {
	s.Code = code
	s.Msg = msg
}
