package checkTools

import (
	"os"
	"runtime"
)

func SetConfigDir() string {
	var result string
	POSIX := "/.config/BalloonUpdateTool/"
	switch runtime.GOOS {
	case "windows":
		result = os.Getenv("APPDATA") + "\\BalloonUpdateTool\\" // Windows
	case "darwin":
		result = os.Getenv("HOME") + "/Library/Application Support/BalloonUpdateTool/" // Mac OS X
	case "linux":
		result = os.Getenv("HOME") + POSIX // Linux
	case "freebsd":
		result = os.Getenv("HOME") + POSIX // FreeBSD
	case "openbsd":
		result = os.Getenv("HOME") + POSIX // OpenBSD
	case "netbsd":
		result = os.Getenv("HOME") + POSIX // NetBSD
	case "dragonfly":
		result = os.Getenv("HOME") + POSIX // DragonFly BSD
	case "plan9":
		result = os.Getenv("HOME") + POSIX // Plan 9
	case "solaris":
		result = os.Getenv("HOME") + POSIX // Solaris

	}
	return result
}
func FailExitCode() {
	if runtime.GOOS != "windows" {
		os.Exit(1)
	} else {
		os.Exit(255)
	}
}
