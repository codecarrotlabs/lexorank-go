package lexorank

import (
    "errors"
    "lexorank-go/numeralsystem"
    "math/big"
    "strings"
)

type LexoDecimal struct {
    Mag   *LexoInteger
    Sig   int
}

func Half(system numeralsystem.LexoNumeralSystem) *LexoDecimal {
    mid := system.GetBase() / 2
    return MakeLexoDecimal(NewLexoInteger(big.NewInt(int64(mid)), 1, system.GetBase()), 1)
}

func ParseLexoDecimal(str string, system numeralsystem.LexoNumeralSystem) (*LexoDecimal, error) {
    radix := system.GetRadixPointChar()
    partialIndex := strings.Index(str, radix)
    if strings.LastIndex(str, radix) != partialIndex {
        return nil, errors.New("more than one radix point")
    }
    if partialIndex < 0 {
        integer, err := ParseLexoInteger(str, system)
        if err != nil {
            return nil, err
        }
        return MakeLexoDecimal(integer, 0), nil
    }
    intStr := str[:partialIndex] + str[partialIndex+1:]
    integer, err := ParseLexoInteger(intStr, system)
    if err != nil {
        return nil, err
    }
    return MakeLexoDecimal(integer, len(str)-1-partialIndex), nil
}

func FromLexoInteger(integer *LexoInteger) *LexoDecimal {
    return MakeLexoDecimal(integer, 0)
}

func MakeLexoDecimal(integer *LexoInteger, sig int) *LexoDecimal {
    if integer.IsZero() {
        return &LexoDecimal{Mag: integer, Sig: 0}
    }
    zeroCount := 0
    for i := 0; i < sig && integer.GetMag(i) == 0; i++ {
        zeroCount++
    }
    newInteger := integer.ShiftRight(zeroCount)
    newSig := sig - zeroCount
    return &LexoDecimal{Mag: newInteger, Sig: newSig}
}

func (d *LexoDecimal) GetSystem() numeralsystem.LexoNumeralSystem {
    return d.Mag.System
}

func (d *LexoDecimal) Add(other *LexoDecimal) *LexoDecimal {
    tmag := d.Mag
    tsig := d.Sig
    omag := other.Mag
    osig := other.Sig
    for tsig < osig {
        tmag = tmag.ShiftLeft(1)
        tsig++
    }
    for tsig > osig {
        omag = omag.ShiftLeft(1)
        osig++
    }
    return MakeLexoDecimal(tmag.Add(omag), tsig)
}

func (d *LexoDecimal) Subtract(other *LexoDecimal) *LexoDecimal {
    thisMag := d.Mag
    thisSig := d.Sig
    otherMag := other.Mag
    otherSig := other.Sig
    for thisSig < otherSig {
        thisMag = thisMag.ShiftLeft(1)
        thisSig++
    }
    for thisSig > otherSig {
        otherMag = otherMag.ShiftLeft(1)
        otherSig++
    }
    return MakeLexoDecimal(thisMag.Subtract(otherMag), thisSig)
}

func (d *LexoDecimal) Multiply(other *LexoDecimal) *LexoDecimal {
    return MakeLexoDecimal(d.Mag.Multiply(other.Mag), d.Sig+other.Sig)
}

func (d *LexoDecimal) Floor() *LexoInteger {
    return d.Mag.ShiftRight(d.Sig)
}

func (d *LexoDecimal) Ceil() *LexoInteger {
    if d.IsExact() {
        return d.Mag
    }
    floor := d.Floor()
    return floor.Add(OneLexoInteger(floor.System))
}

func (d *LexoDecimal) IsExact() bool {
    if d.Sig == 0 {
        return true
    }
    for i := 0; i < d.Sig; i++ {
        if d.Mag.GetMag(i) != 0 {
            return false
        }
    }
    return true
}

func (d *LexoDecimal) GetScale() int {
    return d.Sig
}

func (d *LexoDecimal) SetScale(nsig int, ceiling bool) *LexoDecimal {
    if nsig >= d.Sig {
        return d
    }
    if nsig < 0 {
        nsig = 0
    }
    diff := d.Sig - nsig
    nmag := d.Mag.ShiftRight(diff)
    if ceiling {
        nmag = nmag.Add(OneLexoInteger(nmag.System))
    }
    return MakeLexoDecimal(nmag, nsig)
}

func (d *LexoDecimal) CompareTo(other *LexoDecimal) int {
    tMag := d.Mag
    oMag := other.Mag
    if d.Sig > other.Sig {
        oMag = oMag.ShiftLeft(d.Sig - other.Sig)
    } else if d.Sig < other.Sig {
        tMag = tMag.ShiftLeft(other.Sig - d.Sig)
    }
    return tMag.CompareTo(oMag)
}

func (d *LexoDecimal) Format() string {
    intStr := d.Mag.Format()
    if d.Sig == 0 {
        return intStr
    }
    sb := intStr
    head := ""
    if len(sb) > 0 {
        head = string(sb[0])
    }
    specialHead := head == d.Mag.System.GetPositiveChar() || head == d.Mag.System.GetNegativeChar()
    if specialHead {
        sb = sb[1:]
    }
    for len(sb) < d.Sig+1 {
        sb = d.Mag.System.ToChar(0) + sb
    }
    radix := d.Mag.System.GetRadixPointChar()
    sb = sb[:len(sb)-d.Sig] + radix + sb[len(sb)-d.Sig:]
    if len(sb)-d.Sig == 0 {
        sb = d.Mag.System.ToChar(0) + sb
    }
    if specialHead {
        sb = head + sb
    }
    return sb
}

func (d *LexoDecimal) Equals(other *LexoDecimal) bool {
    if d == other {
        return true
    }
    if other == nil {
        return false
    }
    return d.Mag.Equals(other.Mag) && d.Sig == other.Sig
}

func (d *LexoDecimal) String() string {
    return d.Format()
}