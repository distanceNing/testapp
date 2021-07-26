package comm

type Status struct {
	code int
	msg  string
}

func NewStatus() Status {
	return Status{0, "ok"}
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