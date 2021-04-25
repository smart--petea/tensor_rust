package snake

import (
    "snake/snake/Direction"
)

type Snake struct {
    direction Direction.Direction
    body *LinkedListBlock
    tail *Block
}

func NewSnake(x, y int32) *Snake {
    body := NewLinkedListBlock()
    body.PushBack(NewBlock(x + 2, y))
    body.PushBack(NewBlock(x + 1, y))
    body.PushBack(NewBlock(x    , y))

    return &Snake {
        direction: Direction.Right,
        body: body,
    }
}
