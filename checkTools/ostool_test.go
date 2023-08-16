package checkTools

import (
	"os"
	"testing"
)

func TestSetConfigDir(t *testing.T) {
	tests := []struct {
		name string
		want string
	}{
		{"testingWindows", os.Getenv("APPDATA") + "\\BalloonUpdateTool\\"},
		{"testingPOSIX", os.Getenv("HOME") + "/.config/BalloonUpdateTool/"},
		{"testingMacOS", os.Getenv("HOME") + "/Library/Application Support/BalloonUpdateTool/"}, // Mac OS X
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SetConfigDir(); got != tt.want {
				t.Errorf("SetConfigDir() = %v, want %v", got, tt.want)
			}
		})
	}
}
