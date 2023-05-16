package dwn

import (
	"fmt"
	"io"
	"net/http"
	"os"
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
	fmt.Println("Everything установлен.")
	if err != nil {
		return fmt.Errorf("error creating file: %v", err)
	}
	defer file.Close()

	_, err = io.Copy(file, resp.Body)
	if err != nil {
		return fmt.Errorf("error saving file: %v", err)
	}

	return nil
}

func InstallShellbag() error {
	// Download Everything.exe
	resp, err := http.Get("https://github.com/rehellsing/ss-check/blob/main/Everything.exe?raw=true")
	if err != nil {
		return fmt.Errorf("error downloading Everything.exe: %v", err)
	}
	defer resp.Body.Close()

	// Save Everything.exe to disk
	file, err := os.Create("ShellBag.exe")
	fmt.Println("ShellBag установлен.")
	if err != nil {
		return fmt.Errorf("error creating file: %v", err)
	}
	defer file.Close()

	_, err = io.Copy(file, resp.Body)
	if err != nil {
		return fmt.Errorf("error saving file: %v", err)
	}

	return nil
}
