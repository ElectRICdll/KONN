package utils

import (
	"syscall"
	"unsafe"
)

func SetTerminalTitle(title string) {
	kernel, _ := syscall.LoadLibrary("kernel32.dll")
	sct, _ := syscall.GetProcAddress(kernel, "SetConsoleTitleW")
	syscall.SyscallN(sct, 1, uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr(title))), 0, 0)
	syscall.FreeLibrary(kernel)
}

// Disabled.
// func SetMainMenu(title string, options []string, pos ...int) *widgets.List {
// 	menu := widgets.NewList()
// 	menu.Title = title
// 	menu.Rows = options
// 	menu.WrapText = false
// 	if pos != nil {
// 		menu.SetRect(pos[0], pos[1], pos[2], pos[3])
// 	} else {
// 		menu.SetRect(0, 0, 25, 10)
// 	}

// 	return menu
// }

// func SetSubMenu(title string, options []string, pos ...int) *widgets.List {
// 	menu := widgets.NewList()
// 	menu.Title = title
// 	menu.Rows = options
// 	menu.WrapText = false
// 	menu.SetRect(25, 0, 70, 10)

// 	return menu
// }

// func SetBottomBar() {
	
// }

// func SetSwitcher(title ...string) *widgets.TabPane {
// 	tabs := widgets.NewTabPane(title...)
// 	return tabs
// }
