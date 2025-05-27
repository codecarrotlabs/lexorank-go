package lexorankInteger

import (
	"fmt"
	"lexorank-go/utils"
)

// LexoInteger represents an integer in a custom numeral system with sign and magnitude.
type LexoInteger struct {
    System ILexoNumeralSystem
    Sign   int
    Mag    []int
}

// ParseLexoInteger parses a string into a LexoInteger using the given numeral system.
func ParseLexoInteger(strFull string, system ILexoNumeralSystem) (*LexoInteger, error) {
    str := strFull
    sign := 1
    if len(strFull) > 0 && string(strFull[0]) == system.GetPositiveChar() {
        str = strFull[1:]
    } else if len(strFull) > 0 && string(strFull[0]) == system.GetNegativeChar() {
        str = strFull[1:]
        sign = -1
    }
    mag := make([]int, len(str))
    strIndex := len(mag) - 1
    for magIndex := 0; strIndex >= 0; magIndex++ {
        d, err := system.ToDigit(string(str[strIndex]))
        if err != nil {
            return nil, err
        }
        mag[magIndex] = d
        strIndex--
    }
    return MakeLexoInteger(system, sign, mag), nil
}

func ZeroLexoInteger(sys ILexoNumeralSystem) *LexoInteger {
    return &LexoInteger{System: sys, Sign: 0, Mag: []int{0}}
}

func OneLexoInteger(sys ILexoNumeralSystem) *LexoInteger {
    return MakeLexoInteger(sys, 1, []int{1})
}

func MakeLexoInteger(sys ILexoNumeralSystem, sign int, mag []int) *LexoInteger {
    actualLength := len(mag)
    for actualLength > 0 && mag[actualLength-1] == 0 {
        actualLength--
    }
    if actualLength == 0 {
        return ZeroLexoInteger(sys)
    }
    if actualLength == len(mag) {
        return &LexoInteger{System: sys, Sign: sign, Mag: mag}
    }
    nmag := make([]int, actualLength)
    copy(nmag, mag[:actualLength])
    return &LexoInteger{System: sys, Sign: sign, Mag: nmag}
}

func addInts(sys ILexoNumeralSystem, l, r []int) []int {
    estimatedSize := max(len(l), len(r))
    result := make([]int, estimatedSize)
    carry := 0
    for i := 0; i < estimatedSize; i++ {
        lnum := 0
        if i < len(l) {
            lnum = l[i]
        }
        rnum := 0
        if i < len(r) {
            rnum = r[i]
        }
        sum := lnum + rnum + carry
        carry = 0
        for sum >= sys.GetBase() {
            sum -= sys.GetBase()
            carry++
        }
        result[i] = sum
    }
    return extendWithCarry(result, carry)
}

func extendWithCarry(mag []int, carry int) []int {
    if carry > 0 {
        extendedMag := make([]int, len(mag)+1)
        copy(extendedMag, mag)
        extendedMag[len(extendedMag)-1] = carry
        return extendedMag
    }
    return mag
}

func subtractInts(sys ILexoNumeralSystem, l, r []int) []int {
    rComplement := complementInts(sys, r, len(l))
    rSum := addInts(sys, l, rComplement)
    rSum[len(rSum)-1] = 0
    return addInts(sys, rSum, []int{1})
}

func multiplyInts(sys ILexoNumeralSystem, l, r []int) []int {
    result := make([]int, len(l)+len(r))
    for li := 0; li < len(l); li++ {
        for ri := 0; ri < len(r); ri++ {
            resultIndex := li + ri
            result[resultIndex] += l[li] * r[ri]
            for result[resultIndex] >= sys.GetBase() {
                result[resultIndex] -= sys.GetBase()
                result[resultIndex+1]++
            }
        }
    }
    return result
}

func complementInts(sys ILexoNumeralSystem, mag []int, digits int) []int {
    if digits <= 0 {
        panic("Expected at least 1 digit")
    }
    nmag := make([]int, digits)
    for i := range nmag {
        nmag[i] = sys.GetBase() - 1
    }
    for i := 0; i < len(mag); i++ {
        nmag[i] = sys.GetBase() - 1 - mag[i]
    }
    return nmag
}

func compareInts(l, r []int) int {
    if len(l) < len(r) {
        return -1
    }
    if len(l) > len(r) {
        return 1
    }
    for i := len(l) - 1; i >= 0; i-- {
        if l[i] < r[i] {
            return -1
        }
        if l[i] > r[i] {
            return 1
        }
    }
    return 0
}

func (li *LexoInteger) Add(other *LexoInteger) *LexoInteger {
    li.checkSystem(other)
    if li.IsZero() {
        return other
    }
    if other.IsZero() {
        return li
    }
    if li.Sign != other.Sign {
        if li.Sign == -1 {
            return li.Negate().Subtract(other).Negate()
        }
        return li.Subtract(other.Negate())
    }
    result := addInts(li.System, li.Mag, other.Mag)
    return MakeLexoInteger(li.System, li.Sign, result)
}

func (li *LexoInteger) Subtract(other *LexoInteger) *LexoInteger {
    li.checkSystem(other)
    if li.IsZero() {
        return other.Negate()
    }
    if other.IsZero() {
        return li
    }
    if li.Sign != other.Sign {
        if li.Sign == -1 {
            return li.Negate().Add(other).Negate()
        }
        return li.Add(other.Negate())
    }
    cmp := compareInts(li.Mag, other.Mag)
    if cmp == 0 {
        return ZeroLexoInteger(li.System)
    }
    if cmp < 0 {
        return MakeLexoInteger(li.System, ifThenElse(li.Sign == -1, 1, -1), subtractInts(li.System, other.Mag, li.Mag))
    }
    return MakeLexoInteger(li.System, ifThenElse(li.Sign == -1, -1, 1), subtractInts(li.System, li.Mag, other.Mag))
}

func (li *LexoInteger) Multiply(other *LexoInteger) *LexoInteger {
    li.checkSystem(other)
    if li.IsZero() {
        return li
    }
    if other.IsZero() {
        return other
    }
    if li.isOneish() {
        return MakeLexoInteger(li.System, signProduct(li.Sign, other.Sign), other.Mag)
    }
    if other.isOneish() {
        return MakeLexoInteger(li.System, signProduct(li.Sign, other.Sign), li.Mag)
    }
    newMag := multiplyInts(li.System, li.Mag, other.Mag)
    return MakeLexoInteger(li.System, signProduct(li.Sign, other.Sign), newMag)
}

func (li *LexoInteger) Negate() *LexoInteger {
    if li.IsZero() {
        return li
    }
    return MakeLexoInteger(li.System, ifThenElse(li.Sign == 1, -1, 1), li.Mag)
}

func (li *LexoInteger) ShiftLeft(times int) *LexoInteger {
    if times == 0 {
        return li
    }
    if times < 0 {
        return li.ShiftRight(-times)
    }
    nmag := make([]int, len(li.Mag)+times)
    copy(nmag[times:], li.Mag)
    return MakeLexoInteger(li.System, li.Sign, nmag)
}

func (li *LexoInteger) ShiftRight(times int) *LexoInteger {
    if len(li.Mag)-times <= 0 {
        return ZeroLexoInteger(li.System)
    }
    nmag := make([]int, len(li.Mag)-times)
    copy(nmag, li.Mag[times:])
    return MakeLexoInteger(li.System, li.Sign, nmag)
}

func (li *LexoInteger) Complement() *LexoInteger {
    return li.ComplementDigits(len(li.Mag))
}

func (li *LexoInteger) ComplementDigits(digits int) *LexoInteger {
    return MakeLexoInteger(li.System, li.Sign, complementInts(li.System, li.Mag, digits))
}

func (li *LexoInteger) IsZero() bool {
    return li.Sign == 0 && len(li.Mag) == 1 && li.Mag[0] == 0
}

func (li *LexoInteger) IsOne() bool {
    return li.Sign == 1 && len(li.Mag) == 1 && li.Mag[0] == 1
}

func (li *LexoInteger) GetMag(index int) int {
    return li.Mag[index]
}

func (li *LexoInteger) CompareTo(other *LexoInteger) int {
    if li == other {
        return 0
    }
    if other == nil {
        return 1
    }
    if li.Sign == -1 {
        if other.Sign == -1 {
            cmp := compareInts(li.Mag, other.Mag)
            if cmp == -1 {
                return 1
            }
            if cmp == 1 {
                return -1
            }
            return 0
        }
        return -1
    }
    if li.Sign == 1 {
        if other.Sign == 1 {
            return compareInts(li.Mag, other.Mag)
        }
        return 1
    }
    if other.Sign == -1 {
        return 1
    }
    if other.Sign == 1 {
        return -1
    }
    return 0
}

func (li *LexoInteger) GetSystem() ILexoNumeralSystem {
    return li.System
}

func (li *LexoInteger) Format() string {
    if li.IsZero() {
        ch, _ := li.System.ToChar(0)
        return string(ch)
    }
    sb := utils.NewStringBuilder()
    for i := 0; i < len(li.Mag); i++ {
        ch, _ := li.System.ToChar(li.Mag[i])
        sb = string(ch) + sb
    }
    if li.Sign == -1 {
        sb = li.System.GetNegativeChar() + sb
    }
    return sb
}

func (li *LexoInteger) Equals(other *LexoInteger) bool {
    if li == other {
        return true
    }
    if other == nil {
        return false
    }
    return li.System.GetBase() == other.System.GetBase() && li.CompareTo(other) == 0
}

func (li *LexoInteger) String() string {
    return li.Format()
}

// --- helpers ---

func (li *LexoInteger) isOneish() bool {
    return len(li.Mag) == 1 && li.Mag[0] == 1
}

func (li *LexoInteger) checkSystem(other *LexoInteger) {
    if li.System.GetBase() != other.System.GetBase() {
        panic("Expected numbers of same numeral sys")
    }
}

func signProduct(a, b int) int {
    if a == b {
        return 1
    }
    return -1
}

func ifThenElse(cond bool, a, b int) int {
    if cond {
        return a
    }
    return b
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}