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

var objects []Object
var playerSize C.Vector2

func init() {
	playerSize = C.Vector2{100, 100}

	objects = make([]Object, 20)
	for i := 0; i < 20; i++ {
		objects[i] = Object{
			position: C.Vector2{
				C.float(C.GetRandomValue(0, C.GetScreenWidth()-C.int(playerSize.x))),
				C.float(C.GetRandomValue(0, C.GetScreenHeight()-C.int(playerSize.y))),
			},
			velocity: C.Vector2{200, 200},
			color:    C.ColorFromHSV(360*C.float(C.GetRandomValue(0, 100)/100.0), 1, 1),
		}
	}
}

func main() {
	C.InitWindow(C.int(800), C.int(600), C.CString("Hello, from go"))
	defer C.CloseWindow()

	C.SetTargetFPS(60)
}
