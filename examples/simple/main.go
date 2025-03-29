package main

import "github.com/jupiterrider/purego-sdl3/sdl"

func main() {
	if err := sdl.SetHint(sdl.HintRenderVSync, "1"); err != nil {
		panic(err)
	}

	defer sdl.Quit()
	if err := sdl.Init(sdl.InitVideo); err != nil {
		panic(err)
	}

	window, renderer, err := sdl.CreateWindowAndRenderer("Hello, World!", 1280, 720, sdl.WindowResizable)
	if err != nil {
		panic(err)
	}
	defer sdl.DestroyRenderer(renderer)
	defer sdl.DestroyWindow(window)

	sdl.SetRenderDrawColor(renderer, 100, 150, 200, 255)

Outer:
	for {
		var event sdl.Event
		for sdl.PollEvent(&event) {
			switch event.Type() {
			case sdl.EventQuit:
				break Outer
			case sdl.EventKeyDown:
				if event.Key().Scancode == sdl.ScancodeEscape {
					break Outer
				}
			}
		}
		sdl.RenderClear(renderer)
		sdl.RenderPresent(renderer)
	}
}
