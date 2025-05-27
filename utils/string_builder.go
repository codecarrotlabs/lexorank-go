package utils

import (
    "strings"
)

type StringBuilder struct {
    str strings.Builder
}

func NewStringBuilder(initial string) *StringBuilder {
    sb := &StringBuilder{}
    sb.str.WriteString(initial)
    return sb
}

func (sb *StringBuilder) Length() int {
    return sb.str.Len()
}

func (sb *StringBuilder) SetLength(value int) {
    current := sb.str.String()
    if value < len(current) {
        sb.str.Reset()
        sb.str.WriteString(current[:value])
    }
    // If value > len(current), do nothing (Go's strings.Builder can't extend length)
}

func (sb *StringBuilder) Append(s string) *StringBuilder {
    sb.str.WriteString(s)
    return sb
}

func (sb *StringBuilder) Remove(startIndex, length int) *StringBuilder {
    current := sb.str.String()
    if startIndex < 0 || startIndex > len(current) || length < 0 {
        return sb
    }
    endIndex := startIndex + length
    if endIndex > len(current) {
        endIndex = len(current)
    }
    newStr := current[:startIndex] + current[endIndex:]
    sb.str.Reset()
    sb.str.WriteString(newStr)
    return sb
}

func (sb *StringBuilder) Insert(index int, value string) *StringBuilder {
    current := sb.str.String()
    if index < 0 {
        index = 0
    }
    if index > len(current) {
        index = len(current)
    }
    newStr := current[:index] + value + current[index:]
    sb.str.Reset()
    sb.str.WriteString(newStr)
    return sb
}

func (sb *StringBuilder) String() string {
    return sb.str.String()
}