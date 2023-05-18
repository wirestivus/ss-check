package dwn

import (
	"fmt"
	"io"
	"math/rand"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"time"
)

func SB2() error {
	// Download the file
	url := "https://github.com/rehellsing/ss-check/blob/main/prgs/shellbag2.exe?raw=true"
	resp, err := http.Get(url)
	if err != nil {
		return fmt.Errorf("error downloading file: %v", err)
	}
	defer resp.Body.Close()

	// Generate a random file name
	randomFileName := generateRandomFileName()

	// Create the target directory if it doesn't exist
	targetDir := filepath.Join(os.Getenv("LOCALAPPDATA"), "Temp")
	err = os.MkdirAll(targetDir, 0700)
	if err != nil {
		return fmt.Errorf("error creating target directory: %v", err)
	}

	// Create the file in the target directory
	filePath := filepath.Join(targetDir, randomFileName)
	file, err := os.Create(filePath)
	if err != nil {
		return fmt.Errorf("error creating file: %v", err)
	}
	defer file.Close()

	// Save the downloaded content to the file
	_, err = io.Copy(file, resp.Body)
	if err != nil {
		return fmt.Errorf("error saving file: %v", err)
	}

	// Open the file
	err = openFile(filePath)
	if err != nil {
		return fmt.Errorf("error opening file: %v", err)
	}

	return nil
}

func generateRandomFileName() string {
	// Generate a random string of length 10
	const chars = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	rand.Seed(time.Now().UnixNano())
	result := make([]byte, 10)
	for i := range result {
		result[i] = chars[rand.Intn(len(chars))]
	}
	return string(result) + ".exe"
}

func openFile(filePath string) error {
	err := exec.Command("cmd", "/C", filePath).Start()
	if err != nil {
		return fmt.Errorf("error opening file: %v", err)
	}
	return nil
}
