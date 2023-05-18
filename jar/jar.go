package jar

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

func PrintJARFiles() {
	labymodDir := filepath.Join(os.Getenv("APPDATA"), ".minecraft", "LabyMod")
	prefix := "addons-"

	// Проверяем, существует ли указанная директория
	if _, err := os.Stat(labymodDir); os.IsNotExist(err) {
		fmt.Println("Директория не существует:", labymodDir)
		return
	}

	// Проверяем, является ли путь директорией
	info, err := os.Stat(labymodDir)
	if err != nil {
		fmt.Println("Произошла ошибка при проверке директории:", err)
		return
	}
	if !info.IsDir() {
		fmt.Println("Указанный путь не является директорией:", labymodDir)
		return
	}

	err = filepath.Walk(labymodDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if !info.IsDir() {
			return nil // пропускаем файлы, проверяем только директории
		}

		dirName := filepath.Base(path)
		if strings.HasPrefix(dirName, prefix) {
			files, err := ioutil.ReadDir(path)
			if err != nil {
				return err
			}

			for _, file := range files {
				if !file.IsDir() && filepath.Ext(file.Name()) == ".jar" {
					fmt.Println("Найден JAR-файл:", filepath.Join(path, file.Name()))
				}
			}
		}

		return nil
	})

	if err != nil {
		fmt.Println("Произошла ошибка:", err)
	}
}
