// +build glfw

package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/freneticmonkey/imgui-go"

	"github.com/freneticmonkey/imgui-go-examples/internal/demo"
	"github.com/freneticmonkey/imgui-go-examples/internal/platforms"
	"github.com/freneticmonkey/imgui-go-examples/internal/renderers"
)

func main() {

	viewports := false
	docking := false

	flag.BoolVar(&viewports, "viewports", false, "enable viewports mode" )
	flag.BoolVar(&docking, "docking", true, "enable docking mode")
	flag.Parse()

	context := imgui.CreateContext(nil)
	defer context.Destroy()
	io := imgui.CurrentIO()
	flags := imgui.ConfigFlagEnableDocking | imgui.ConfigFlagEnableViewports //| imgui.ConfigFlagNavEnableKeyboard |
	io.SetConfigFlags(flags)

	platform, err := platforms.NewGLFW(io, platforms.GLFWClientAPIOpenGL3, viewports, docking)
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
