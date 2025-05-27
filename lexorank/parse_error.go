package lexorank

import "fmt"

type ParseError struct {
	Msg string
}

func (e *ParseError) Error() string {
	return fmt.Sprintf("Parse Error: %s", e.Msg)
}