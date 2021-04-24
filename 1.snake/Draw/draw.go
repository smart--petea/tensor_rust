package Draw

import (
    "image/color"
    "github.com/faiface/pixel/imdraw"
    "github.com/faiface/pixel"
)

const BLOCK_SIZE float64 = 25.0

func ToCoord(gameCoord int32) float64 {
    return float64(gameCoord) * BLOCK_SIZE
}

func DrawBlock(colour color.Color, x, y int32) *imdraw.IMDraw {
    gui_x := ToCoord(x)
    gui_y := ToCoord(y)

    imd := imdraw.New(nil)

    imd.Color = colour
    imd.Push(pixel.V(gui_x, gui_y))
    imd.Push(pixel.V(gui_x + BLOCK_SIZE, gui_y + BLOCK_SIZE))
    imd.Rectangle(0)

    return imd
}
