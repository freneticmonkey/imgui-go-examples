// +build glfw

package platforms

import (
	"fmt"
	"github.com/freneticmonkey/imgui-go"
	"github.com/go-gl/glfw/v3.2/glfw"
	"unsafe"

	// "github.com/go-gl/glfw/v3.2/glfw"
)

func (platform *GLFW) SetupViewportHandling() {

	pIO := imgui.CurrentPlatformIO()
	
	pIO.SetCallbackCreateWindow(platform.CreateWindow)

	pIO.SetCallbackDestroyWindow(platform.DestroyWindow)

	pIO.SetCallbackShowWindow(platform.ShowWindow)

	pIO.SetCallbackSetWindowPos(platform.SetWindowPos)
	pIO.SetCallbackGetWindowPos(platform.GetWindowPos)

	pIO.SetCallbackSetWindowSize(platform.SetWindowSize)
	pIO.SetCallbackGetWindowSize(platform.GetWindowPos)

	pIO.SetCallbackSetWindowFocus(platform.SetWindowFocus)
	pIO.SetCallbackGetWindowFocus(platform.GetWindowFocus)

	pIO.SetCallbackGetWindowMinimized(platform.GetWindowMinimised)

	pIO.SetCallbackSetWindowTitle(platform.SetWindowTitle)
	pIO.SetCallbackRenderWindow(platform.RenderWindow)
	pIO.SetCallbackSwapBuffers(platform.SwapBuffers)

	// FIXME: Enable Platform Viewports support
	flags := 0
	flags |= 1 << 10

	// FIXME: Enable Renderer Viewports support
	flags |= 1 << 12

	platform.imguiIO.SetBackendFlags(flags)

	platform.platformIO = pIO
}

func (platform *GLFW) CreateWindow(v *imgui.Viewport) {
	fmt.Println("Create Window Called")

	glfw.WindowHint(glfw.Visible, 0)
	glfw.WindowHint(glfw.Focused, 0)

	//flags := v.GetFlags()
	//decorated := 0
	//if flags & imgui.ViewportFlags_NoDecoration == 0 {
	//	decorated = false
	//}
	glfw.WindowHint(glfw.Decorated, 0)

	w, h := v.GetSize()
	x, y := v.GetPosition()


	// FIXME: Only set share (window) if in OpenGL mode
	win, err := glfw.CreateWindow(w, h, "No Title Yet", nil, platform.window)
	if err != nil {
		glfw.Terminate()
		fmt.Errorf("failed to create window: %v", err)
		return
	}
	v.SetPlatformHandle(unsafe.Pointer(win))
	win.SetPos(x, y)

	platform.installCallbacks(win)

	// FIXME: Again only set if in OpenGL mode
	win.MakeContextCurrent()
	glfw.SwapInterval(0)
}

func (platform *GLFW) DestroyWindow(v *imgui.Viewport) {
	fmt.Println("DestroyWindow Called")

	window := (*glfw.Window)(v.GetData())
	if window != nil {
		// FIXME: Impl WindowOwned
		window.Destroy()
		v.SetData(nil)
		v.SetPlatformHandle(nil)
	}
}

func (platform *GLFW) ShowWindow(v *imgui.Viewport) {
	fmt.Println("ShowWindow Called")

	window := (*glfw.Window)(v.GetData())
	if window != nil {
		window.Show()
	}
}

func (platform *GLFW) SetWindowPos(v *imgui.Viewport, x, y float64) {
	fmt.Println("SetWindowPos Called")
}

func (platform *GLFW) GetWindowPos(v *imgui.Viewport) (float64, float64) {
	fmt.Println("GetWindowPos Called")
	return 0, 0
}

func (platform *GLFW) SetWindowSize(v *imgui.Viewport, x, y float64) {
	fmt.Println("SetWindowSize Called")
}

func (platform *GLFW) GetWindowSize(v *imgui.Viewport) (float64, float64) {
	fmt.Println("GetWindowSize Called")
	return 0, 0
}

func (platform *GLFW) SetWindowFocus(v *imgui.Viewport) {
	fmt.Println("SetWindowFocus Called")
}

func (platform *GLFW) GetWindowFocus(v *imgui.Viewport) bool {
	fmt.Println("GetWindowFocus Called")
	return false
}

func (platform *GLFW) GetWindowMinimised(v *imgui.Viewport) bool {
	fmt.Println("GetWindowMinimised Called")
	return false
}

func (platform *GLFW) SetWindowTitle(v *imgui.Viewport, title string) {
	fmt.Println("SetWindowFocus Called")
}

func (platform *GLFW) RenderWindow(v *imgui.Viewport) {
	fmt.Println("RenderWindow Called")
	window := (*glfw.Window)(v.GetData())
	if window != nil {
		window.MakeContextCurrent()
	}
}

func (platform *GLFW) SwapBuffers(v *imgui.Viewport) {
	fmt.Println("SwapBuffers Called")
	window := (*glfw.Window)(v.GetData())
	if window != nil {
		window.MakeContextCurrent()
		window.SwapBuffers()
	}
}

func (platform *GLFW) updateMonitors() {
	monitors := glfw.GetMonitors()
	platform.platformIO.ClearMonitors()
	for _, monitor := range monitors {
		x, y := monitor.GetPos()
		vidMode := monitor.GetVideoMode()
		platform.platformIO.AddMonitor(x,y,vidMode.Width,vidMode.Height)
	}
}
