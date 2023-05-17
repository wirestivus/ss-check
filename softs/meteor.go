package softs

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
)

func Meteor() error {
	// URL для скачивания файла
	url := "https://github.com/rehellsing/ss-check/blob/main/prgs/meteor.jar?raw=true"

	// Определите путь к папке appdata/.minecraft/mods
	appDataDir, err := getAppDataDir()
	if err != nil {
		return err
	}
	modsDir := filepath.Join(appDataDir, ".minecraft", "mods")

	// Создайте папку mods, если она не существует
	err = os.MkdirAll(modsDir, 0755)
	if err != nil {
		return err
	}

	// Создайте HTTP-запрос для скачивания файла
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Создайте файл для записи загруженных данных
	filePath := filepath.Join(modsDir, "meteor.jar")
	file, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

	// Скопируйте данные из ответа HTTP в файл
	_, err = io.Copy(file, resp.Body)
	if err != nil {
		return err
	}

	fmt.Println("Файл успешно загружен и перемещен в папку mods.")
	return nil
}

func getAppDataDir() (string, error) {
	// Получите путь к папке appdata
	appDataDir, err := os.UserCacheDir()
	if err != nil {
		return "", err
	}

	return appDataDir, nil
}
