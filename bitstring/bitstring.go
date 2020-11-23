package bitstring

import (
    "math"
    "math/bits"
)

type Bitstring struct {
    data []uint8
    w, h uint
}

func NewBitstring(w, h uint) *Bitstring {
    size := math.Round((float64(w * h) / 8.0) + 0.5)
    data := make([]uint8, int(size))
    b := Bitstring{
        data: data,
        w: w,
        h: h,
    }
    return &b
}

func (b *Bitstring) Pop() (sum int) {
    for i := 0; i < len(b.data); i++ {
        sum += bits.OnesCount8(b.data[i])
    }
    return
}

func (b *Bitstring) ToString() (str string) {
    var x, y uint
    for y = 0; y < b.h; y++ {
        for x = 0; x < b.w; x++ {
            if b.Get(x, y) {
                str += "X"
            } else {
                str += "-"
            }
        }
        str += "\n"
    }
    return
}

func (b *Bitstring) Get(x, y uint) bool {
    x += y * b.w
    if x > (b.w * b.h) {
        return false
    }
    i1 := x / 8
    i2 := x % 8
    return (1 & (b.data[i1] >> i2)) > 0

}

func (b *Bitstring) Set(x, y uint, val bool) {
    x += y * b.w
    if x > (b.w * b.h) {
        return
    }
    i1 := x / 8
    i2 := x % 8
    if val {
        b.data[i1] |= (1 << i2)
    } else {
        b.data[i1] = b.data[i1] &^ (1 << i2)
    }
}

func (op1 *Bitstring) Xor (op2 *Bitstring) *Bitstring {
    res := NewBitstring(op1.w, op1.h)
    for i := 0; i < len(res.data); i ++ {
        res.data[i] = op1.data[i] ^ op2.data[i]
    }
    return res
}

