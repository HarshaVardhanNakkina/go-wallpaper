package setwallpaper

import (
	"fmt"
	"syscall"
	"unsafe"
)

const (
	SPI_SETDESKWALLPAPER = 0x0014

	uiParam = 0x0000

	SPIF_UPDATEINIFILE = 0x01
	SPIF_SENDCHANGE    = 0x02
)

var (
	user32DLL           = syscall.NewLazyDLL("user32.dll")
	procSystemParamInfo = user32DLL.NewProc("SystemParametersInfoW")
)

func SetWallpaper(filepath string) {
	imagePath, _ := syscall.UTF16PtrFromString(filepath)
	fmt.Println("Setting wallpaper ...")
	procSystemParamInfo.Call(
		SPI_SETDESKWALLPAPER, uiParam, uintptr(unsafe.Pointer(imagePath)), uintptr(SPIF_UPDATEINIFILE|SPIF_SENDCHANGE))
}
