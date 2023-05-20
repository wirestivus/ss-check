package dll

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func PrintDLLFiles() {
	downloadsDir, err := filepath.Abs(filepath.Join(os.Getenv("USERPROFILE"), "Downloads"))
	if err != nil {
		fmt.Println("")
		fmt.Println("Ошибка при получении каталога загрузок:", err)
		os.Exit(1)
	}

	fmt.Println("")
	fmt.Println("DLL-файлы в каталоге Downloads:")
	fmt.Println("")

	foundDLL := false // Флаг для отслеживания наличия DLL-файлов

	err = filepath.Walk(downloadsDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() && strings.HasSuffix(info.Name(), ".dll") {
			fmt.Println(path)
			foundDLL = true // Устанавливаем флаг, если найден хотя бы один DLL-файл
		}
		return nil
	})
	if err != nil {
		fmt.Println("Ошибка при поиске файлов DLL в каталоге Downloads:", err)
		os.Exit(1)
	}

	if !foundDLL {
		fmt.Println("DLL-файлы не найдены в каталоге Downloads.")
	}
}
