package snake

import (
    "fmt"
)

type Direction string

const (
    Up Direction = "Up"
    Down Direction = "Down"
    Left Direction = "Left"
    Right Direction = "Right"
)

func (d Direction) Opposite() Direction {
    switch d {
    case Up: return Down
    case Down: return Up
    case Left: return Right
    case Right: return Left
    default:
        panic(fmt.Sprintf("Unrecognized direction=%+v", d))
    }
}
