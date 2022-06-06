package main

import (
	"fmt"

	"github.com/AllenDang/giu"
)

func onClickMe() {
	fmt.Println("Hello world!")
}

func onImSoCute() {
	fmt.Println("Im sooooooo cute!!")
}

func loop() {
	giu.SingleWindow().Layout(
		giu.Label("Hello world from giu"),
		giu.Row(
			giu.Button("Click Me").OnClick(onClickMe),
			giu.Button("I'm so cute").OnClick(onImSoCute),
		),
	)
}

func main() {
	wnd := giu.NewMasterWindow("Hello world", 400, 200, giu.MasterWindowFlagsNotResizable)
	wnd.Run(loop)
}
