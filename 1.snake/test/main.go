package main

import (
    "github.com/faiface/pixel"
    "github.com/faiface/pixel/pixelgl"
    "golang.org/x/image/colornames"

    "snake/draw"
    "snake/snake"

    "fmt"
)

func run() {
    cfg := pixelgl.WindowConfig {
        Title: "Pixel Rocks!",
        Bounds: pixel.R(0, 0, 1024, 768),
        VSync: true,
    }
    win, err := pixelgl.NewWindow(cfg)
    if err != nil {
        panic(err)
    }

    for !win.Closed() {
        win.Clear(colornames.Aliceblue)
        //Draw.DrawBlock(colornames.Black, 10, 15).Draw(win)
        draw.DrawRectangle(colornames.Black, 10, 15, 2, 3).Draw(win)
        win.Update()
    }
}

func main() {
    fmt.Printf("%+v", snake.SNAKE_COLOR)
    pixelgl.Run(run)
}
