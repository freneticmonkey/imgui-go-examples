// +build glfw

package main

import (
	"fmt"
	"os"

	"github.com/freneticmonkey/imgui-go"

	"github.com/frenticmonkey/imgui-go-examples/internal/demo"
	"github.com/frenticmonkey/imgui-go-examples/internal/platforms"
	"github.com/frenticmonkey/imgui-go-examples/internal/renderers"
)

func main() {
	context := imgui.CreateContext(nil)
	defer context.Destroy()
	io := imgui.CurrentIO()
	flags := imgui.ConfigFlagEnableDocking | imgui.ConfigFlagEnableViewports //| imgui.ConfigFlagNavEnableKeyboard |
	io.SetConfigFlags(flags)



	platform, err := platforms.NewGLFW(io, platforms.GLFWClientAPIOpenGL3)
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(-1)
	}
	defer platform.Dispose()

	renderer, err := renderers.NewOpenGL3(io)
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(-1)
	}
	defer renderer.Dispose()

	demo.Run(platform, renderer)
}
