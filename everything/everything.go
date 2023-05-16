package everything

import (
	"os"
	"os/exec"
	"path/filepath"
)

func everything() error {
	// Получаем путь к рабочему столу пользователя
	desktopPath, err := os.UserHomeDir()
	if err != nil {
		return err
	}

	// Формируем полный путь к файлу установщика на рабочем столе
	installerPath := filepath.Join(desktopPath, "Everything-1.4.1.1023.x86-Setup.exe")

	// Запускаем установщик в тихом режиме
	cmd := exec.Command(installerPath, "/SILENT")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err = cmd.Run()
	if err != nil {
		return err
	}

	return nil
}
