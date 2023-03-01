package main

import (
	"time"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func seconds(time int32) int32 {
	return int32((time % 60))
}

func minutes(time int32) int32 {
	return int32((time % 3600 / 60))
}

func hours(time int32) int32 {
	return int32((time % 86400 / 3600))
}

func main() {
	rl.InitWindow(400, 400, "mewl")
	rl.SetTargetFPS(60)

	razzmatazz := rl.NewColor(227, 37, 107, 255)

	for !rl.WindowShouldClose() {
		now := int32(time.Now().Local().Unix())
		seconds := seconds(now)
		minutes := minutes(now)
		hours := hours(now)

		rl.BeginDrawing()
		rl.ClearBackground(rl.Black)

		rl.DrawRectangle(0, 0, 400, 400, razzmatazz)
		rl.DrawRectangle(0, 0, 400, (40 + hours * 15), rl.Black)
		rl.DrawRectangle(40, 40, 320, 320, rl.Black)
		rl.DrawRectangleLines(40, 40, 320, 320, razzmatazz)
		rl.DrawLine(40 + seconds * 5, 40, 40 + seconds * 5, 360, razzmatazz)
		rl.DrawLine(40, 40 + minutes * 5, 360, 40 + minutes * 5, razzmatazz)

		rl.EndDrawing()
	}

	rl.CloseWindow()
}