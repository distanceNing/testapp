package common

const (
	ErrSystem = 1
	ErrTokenTimeOut
	ErrUserNotExist
	ErrUserAlreadyExist
	ErrPasswordNotMatch
	ErrTokenNotMatch
	ErrJsonDecodeFail
	ErrDbDupKey = 10
	ErrNoAffected
)

type Status struct {
	code int
	msg  string
}

func NewStatus() Status {
	return Status{0, "ok"}
}

func (s *Status) Ok() bool {
	return s.code == 0
}

func (s *Status) Code() int {
	return s.code
}

func (s *Status) Msg() string {
	return s.msg
}

func (s *Status) Set(code int, msg string) {
	s.code = code
	s.msg = msg
}
