package gen

import (
    "../gol"
)

type state struct {
    x, y uint
    bool on
}

type rect struct {
    states []state
}

func MakeSvg(l *Life, period uint) (str string) {
    var rects []rect
    var open []*rect
}

