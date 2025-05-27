package utils

type StringBuilder struct {
	str string
}

func (sb *StringBuilder) Write(s string) {
	sb.str += s
}

func (sb *StringBuilder) String() string {
	return sb.str
}