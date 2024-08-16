package main

/*
#cgo CFLAGS: -I.
#cgo LDFLAGS: -L. -lraylib -lm
#include "raylib.h"
*/
import "C"

type Object struct {
	position C.Vector2
	velocity C.Vector2
	color    C.Color
}

var objects []*Object
var playerSize C.Vector2
var collisionDamp C.float
var gravity C.Vector2

func init() {
	playerSize = C.Vector2{100, 100}
	gravity = C.Vector2{0, 100}
	collisionDamp = 1

}

func initObjects() {
	objects = make([]*Object, 20)
	for i := 0; i < 20; i++ {
		objects[i] = &Object{
			position: C.Vector2{
				C.float(C.GetRandomValue(0, C.GetScreenWidth()-C.int(playerSize.x))),
				C.float(C.GetRandomValue(0, C.GetScreenHeight()-C.int(playerSize.y))),
			},
			velocity: C.Vector2{200, 200},
			color:    C.ColorFromHSV(360*(C.float(C.GetRandomValue(0, 100))/100.0), 1, 1),
		}
	}
}

func gameFrame() {
	dt := C.GetFrameTime()

	C.BeginDrawing()
	// C.ClearBackground(C.RAYWHITE)
	C.ClearBackground(C.Color{
		0x18, 0x18, 0x18, 0xFF,
	})

	for _, object := range objects {
		object.velocity.x += gravity.x * dt
		object.velocity.y += gravity.y * dt

		nx := object.position.x + object.velocity.x*dt
		if nx < 0 || nx+playerSize.x > C.float(C.GetScreenWidth()) {
			object.velocity.x *= -collisionDamp
			object.color = C.ColorFromHSV(360*(C.float(C.GetRandomValue(0, 100))/100.0), 1, 1)
		} else {
			object.position.x = nx
		}

		ny := object.position.y + object.velocity.y*dt
		if ny < 0 || ny+playerSize.y > C.float(C.GetScreenHeight()) {
			object.velocity.y *= -collisionDamp
			object.color = C.ColorFromHSV(360*(C.float(C.GetRandomValue(0, 100))/100.0), 1, 1)
		} else {
			object.position.y = ny
		}

		C.DrawRectangleV(object.position, playerSize, object.color)
	}

	C.EndDrawing()
}

func main() {
	C.InitWindow(800, 600, C.CString("Raylib Window"))
	defer C.CloseWindow()

	initObjects()

	C.SetTargetFPS(60)

	for !C.WindowShouldClose() {
		gameFrame()
	}
}
