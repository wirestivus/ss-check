package evr

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"os/exec"
)

func InstallEverything() error {
	// Download Everything.exe
	resp, err := http.Get("https://github.com/rehellsing/ss-check/blob/main/Everything.exe?raw=true")
	if err != nil {
		return fmt.Errorf("error downloading Everything.exe: %v", err)
	}
	defer resp.Body.Close()

	// Save Everything.exe to disk
	file, err := os.Create("Everything.exe")
	if err != nil {
		return fmt.Errorf("error creating file: %v", err)
	}
	defer file.Close()

	_, err = io.Copy(file, resp.Body)
	if err != nil {
		return fmt.Errorf("error saving file: %v", err)
	}

	// Run Everything.exe installer
	cmd := exec.Command("Everything.exe", "/SILENT")
	err = cmd.Run()
	if err != nil {
		return fmt.Errorf("error running installer: %v", err)
	}

	return nil
}
