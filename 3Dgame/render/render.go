package render

import (
	"fmt"
	"time"

	"github.com/ZayaSansanch/golang/3Dgame/data"
)

func motionRender() {
	for i := 0; i < len(data.GameMap); i++ {
		for j := 0; j < len(data.GameMap[i]); j++ {
			// fmt.Print(" ", data.GameMap[i][j])
			fmt.Print(" ", data.GameMap[i][j])

			if j == 9 {
				fmt.Println("")
			}
		}
	}

	fmt.Println("")

	for i := 0; i < len(data.Motion); i++ {
		for j := 0; j < len(data.Motion[i]); j++ {
			// fmt.Print(" ", data.Motion[i][j])
			fmt.Print(" ", data.Motion[i][j])

			if j == 9 {
				fmt.Print(" data: ", i)
				fmt.Println("")
			}
		}
	}
}

func Render() {
	fmt.Println("Render - start")

	for {
		tm1 := time.Now().Second()
		motionRender()
		tm2 := time.Now().Second()
		tm3 := 40 - (tm2 - tm1)
		time.Sleep(tm3 * time.Second)
	}
}

// for i := 0; i < len(data.GameMap); i++ {
// 	for j := 0; j < len(data.GameMap[i]); j++ {
// 		fmt.Print(" ", data.GameMap[i][j])

// 		if j == 9 {
// 			fmt.Println("")
// 		}
// 	}
// }

// func addGameMap() {
// 	for mI := 0; mI < len(data.Motion); mI++ {
// 		for mJ := 0; mJ < len(data.Motion[mI]); mJ++ {
// 			for gI := 0; gI < len(data.GameMap); gI++ {
// 				for gJ := 0; gJ < len(data.GameMap[gI]); gJ++ {
// 					// data.Motion[mJ][mI] = data.GameMap[gI][gJ]
// 					data.Motion[mJ][mI] = 1
// 				}
// 			}
// 		}
// 	}
// }
