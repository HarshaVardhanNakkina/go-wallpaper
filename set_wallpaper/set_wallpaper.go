package setwallpaper

import (
	"fmt"
	"io/fs"
	"io/ioutil"
	"os"
	"path/filepath"
	"syscall"
	"unsafe"

	"github.com/mitchellh/go-homedir"
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
var defaultFileMode fs.FileMode = 0644

func SetWallpaper(filename string, rawImg []byte) error {
	homeDir, err := homedir.Dir()
	if err != nil {
		return err
	}
	fileLocation := filepath.Join(homeDir, "Pictures", "go-wallpaper")
	filepath := filepath.Join(fileLocation, filename)
	createDirIfNotExists(fileLocation)

	fmt.Println("Saving image@", filepath)
	err = ioutil.WriteFile(filepath, rawImg, defaultFileMode)
	if err != nil {
		return err
	}
	imagePath, _ := syscall.UTF16PtrFromString(filepath)
	fmt.Println("Setting wallpaper...")
	procSystemParamInfo.Call(
		SPI_SETDESKWALLPAPER, uiParam, uintptr(unsafe.Pointer(imagePath)), uintptr(SPIF_UPDATEINIFILE|SPIF_SENDCHANGE))

	return nil
}

func createDirIfNotExists(fileLocation string) error {
	_, err := os.Stat(fileLocation)

	if os.IsNotExist(err) {
		errDir := os.MkdirAll(fileLocation, defaultFileMode)
		if errDir != nil {
			return errDir
		}
	}
	return nil
}
