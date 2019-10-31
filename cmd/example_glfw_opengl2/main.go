// +build glfw

package main

import (
	"fmt"
	"os"

	"github.com/freneticmonkey/imgui-go"

	"github.com/freneticmonkey/imgui-go-examples/internal/demo"
	"github.com/freneticmonkey/imgui-go-examples/internal/platforms"
	"github.com/freneticmonkey/imgui-go-examples/internal/renderers"
)

func main() {
	context := imgui.CreateContext(nil)
	defer context.Destroy()
	io := imgui.CurrentIO()

	platform, err := platforms.NewGLFW(io, platforms.GLFWClientAPIOpenGL2)
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(-1)
	}
	defer platform.Dispose()

	renderer, err := renderers.NewOpenGL2(io)
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(-1)
	}
	defer renderer.Dispose()

	demo.Run(platform, renderer)
}
