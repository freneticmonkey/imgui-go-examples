// +build glfw

package platforms

import (
	"fmt"
	"github.com/freneticmonkey/imgui-go"
	// "github.com/go-gl/glfw/v3.2/glfw"
)

func SetupViewportHandling(io imgui.IO) imgui.PlatformIO {

	pIO := imgui.CurrentPlatformIO()
	
	pIO.SetCallbackCreateWindow(CreateWindow)
	pIO.SetCallbackDestroyWindow(DestroyWindow)

	pIO.SetCallbackShowWindow(ShowWindow)

	pIO.SetCallbackSetWindowPos(SetWindowPos)
	pIO.SetCallbackGetWindowPos(GetWindowPos)

	pIO.SetCallbackSetWindowSize(SetWindowSize)
	pIO.SetCallbackGetWindowSize(GetWindowPos)

	pIO.SetCallbackSetWindowFocus(SetWindowFocus)
	pIO.SetCallbackGetWindowFocus(GetWindowFocus)

	pIO.SetCallbackGetWindowMinimized(GetWindowMinimised)

	pIO.SetCallbackSetWindowTitle(SetWindowTitle)
	pIO.SetCallbackRenderWindow(RenderWindow)
	pIO.SetCallbackSwapBuffers(SwapBuffers)

	// FIXME: Enable Platform Viewports support
	flags := 0
	flags |= 1 << 10

	// FIXME: Enable Renderer Viewports support
	flags |= 1 << 12

	io.SetBackendFlags(flags)

	return pIO
}

func CreateWindow(v *imgui.Viewport) {
	fmt.Println("Create Window Called")
}

func DestroyWindow(v *imgui.Viewport) {
	fmt.Println("DestroyWindow Called")
}

func ShowWindow(v *imgui.Viewport) {
	fmt.Println("ShowWindow Called")
}

func SetWindowPos(v *imgui.Viewport, x, y float64) {
	fmt.Println("SetWindowPos Called")
}

func GetWindowPos(v *imgui.Viewport) (float64, float64) {
	fmt.Println("GetWindowPos Called")
	return 0, 0
}

func SetWindowSize(v *imgui.Viewport, x, y float64) {
	fmt.Println("SetWindowSize Called")
}

func GetWindowSize(v *imgui.Viewport) (float64, float64) {
	fmt.Println("GetWindowSize Called")
	return 0, 0
}

func SetWindowFocus(v *imgui.Viewport) {
	fmt.Println("SetWindowFocus Called")
}

func GetWindowFocus(v *imgui.Viewport) bool {
	fmt.Println("GetWindowFocus Called")
	return false
}

func GetWindowMinimised(v *imgui.Viewport) bool {
	fmt.Println("GetWindowMinimised Called")
	return false
}

func SetWindowTitle(v *imgui.Viewport, title string) {
	fmt.Println("SetWindowFocus Called")
}

func RenderWindow(v *imgui.Viewport) {
	fmt.Println("RenderWindow Called")
}

func SwapBuffers(v *imgui.Viewport) {
	fmt.Println("SwapBuffers Called")
}


