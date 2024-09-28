package lib

import (
    "strconv"
)

type Move struct {
    Row int
    Col int
}

func (m Move) GetString() string {
    return string(CHARS[m.Col]) + strconv.Itoa(8 - m.Row)
}

