// +build glfw

package platforms

import (
	"fmt"
	"github.com/freneticmonkey/imgui-go"
	"github.com/go-gl/glfw/v3.2/glfw"
	"unsafe"

	// "github.com/go-gl/glfw/v3.2/glfw"
)

func (platform *GLFW) SetupViewportHandling(viewports, docking bool) {

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
	if viewports {
		flags := 0
		flags |= imgui.BackendFlags_PlatformHasViewports

		// FIXME: Enable Renderer Viewports support
		flags |= imgui.BackendFlags_RendererHasViewports

		platform.imguiIO.SetBackendFlags(flags)
	}

	glfw.SetMonitorCallback(platform.MonitorChangedCallback)
	
	platform.platformIO = pIO
}

func (platform *GLFW) MonitorChangedCallback(m *glfw.Monitor, _ glfw.MonitorEvent) {
	platform.updateMonitors = true
	fmt.Printf("Monitor Callback triggered\n")
}

func (platform *GLFW) CreateWindow(v *imgui.Viewport) {

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

	v.Title = "No Title Yet"
	// FIXME: Only set share (window) if in OpenGL mode
	win, err := glfw.CreateWindow(w, h, v.Title, nil, platform.window)
	if err != nil {
		glfw.Terminate()
		fmt.Errorf("failed to create window: %v", err)
		return
	}
	v.Owned = true
	v.SetPlatformHandle(unsafe.Pointer(win))
	v.SetData(unsafe.Pointer(win))
	win.SetPos(x, y)

	fmt.Printf("Created [%s] : Window pos: x: %d y: %d\n", v.Title, x, y)

	platform.installCallbacks(win)


	win.SetCloseCallback(func(w *glfw.Window){
		v.PlatformRequestClose()
	})

	win.SetPosCallback(func(w *glfw.Window, _, _ int){
		v.PlatformRequestMove()
	})

	win.SetSizeCallback(func(w *glfw.Window, _, _ int){
		// TODO: Size frame handling for glfw
		v.PlatformRequestResize()
	})

	// FIXME: Again only set if in OpenGL mode
	win.MakeContextCurrent()
	glfw.SwapInterval(0)
}

func (platform *GLFW) DestroyWindow(v *imgui.Viewport) {

	window := (*glfw.Window)(v.GetData())
	if window != nil {
		if v.Owned {
			//x, y := window.GetPos()
			//fmt.Printf("Destroying [%s] : Window pos: x: %d y: %d\n", "No Title Yet", x, y)
			window.Destroy()
		}
		imgui.DeleteViewport(v)
	}
}

func (platform *GLFW) ShowWindow(v *imgui.Viewport) {

	window := (*glfw.Window)(v.GetData())
	if window != nil {
		window.Show()
	}
}

func (platform *GLFW) SetWindowPos(v *imgui.Viewport, x, y float64) {
	window := (*glfw.Window)(v.GetData())
	if window != nil {
		window.SetPos(int(x), int(y))
		fmt.Printf(" [%s] : Set Window pos: x: %f y: %f\n", v.Title, x, y)
	}
}

func (platform *GLFW) GetWindowPos(v *imgui.Viewport) (float64, float64) {
	window := (*glfw.Window)(v.GetData())
	if window != nil {
		x, y := window.GetPos()
		return float64(x), float64(y)
	}
	return 0, 0
}

func (platform *GLFW) SetWindowSize(v *imgui.Viewport, width, height float64) {
	window := (*glfw.Window)(v.GetData())
	if window != nil {
		fmt.Printf("[%s] : Set Window size: x: %f y: %f\n", v.Title, width, height)
		window.SetSize(int(width), int(height))
	}
}

func (platform *GLFW) GetWindowSize(v *imgui.Viewport) (float64, float64) {
	window := (*glfw.Window)(v.GetData())
	if window != nil {
		width, height := window.GetSize()
		fmt.Printf("[%s] : Get Window size: x: %f y: %f\n", v.Title, width, height)
		return float64(width), float64(height)
	}
	return 0, 0
}

func (platform *GLFW) SetWindowFocus(v *imgui.Viewport) {
	window := (*glfw.Window)(v.GetData())
	if window != nil {
		window.Focus()
	}
}

func (platform *GLFW) GetWindowFocus(v *imgui.Viewport) bool {
	window := (*glfw.Window)(v.GetData())
	if window != nil {
		return window.GetAttrib(glfw.Focused) != 0
	}
	return false
}

func (platform *GLFW) GetWindowMinimised(v *imgui.Viewport) bool {
	window := (*glfw.Window)(v.GetData())
	if window != nil {
		return window.GetAttrib(glfw.Iconified) != 0
	}
	return false
}

func (platform *GLFW) SetWindowTitle(v *imgui.Viewport, title string) {
	window := (*glfw.Window)(v.GetData())
	if window != nil {
		window.SetTitle(title)
		v.Title = title
	}
}

func (platform *GLFW) RenderWindow(v *imgui.Viewport) {
	window := (*glfw.Window)(v.GetData())
	if window != nil {
		window.MakeContextCurrent()
	}
}

func (platform *GLFW) SwapBuffers(v *imgui.Viewport) {
	window := (*glfw.Window)(v.GetData())
	if window != nil {
		window.MakeContextCurrent()
		window.SwapBuffers()
	}
}

func (platform *GLFW) doUpdateMonitors() {
	monitors := glfw.GetMonitors()
	platform.platformIO.ClearMonitors()
	for i, monitor := range monitors {
		x, y := monitor.GetPos()
		fmt.Printf("Monitor [%d] : Window Pos: x: %d y: %d\n", i, x, y)
		pw, ph := monitor.GetPhysicalSize()
		fmt.Printf("Monitor [%d] : Window Physical Size: w: %d h: %d\n", i, pw, ph)
		vidMode := monitor.GetVideoMode()
		fmt.Printf("Monitor [%d] : Video Mode: w: %d h: %d\n", i, vidMode.Width, vidMode.Height)

		platform.platformIO.AddMonitor(x,y,vidMode.Width,vidMode.Height)
	}
	platform.updateMonitors = false
}
