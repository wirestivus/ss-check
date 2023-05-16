package jar

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func PrintJARFiles() {
	labymodDir := filepath.Join(os.Getenv("APPDATA"), ".minecraft", "LabyMod")
	fmt.Println("")
	fmt.Printf("Поиск JAR файлов в каталогах аддонов LabyMod в: \n")
	fmt.Printf("%s\n", labymodDir)
	fmt.Println("")
	fmt.Printf("JAR-файлы в каталогах аддонов LabyMod:")
	fmt.Println("")
	err := filepath.Walk(labymodDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() && strings.HasPrefix(info.Name(), "addons-") {
			err = filepath.Walk(path, func(path string, info os.FileInfo, err error) error {
				if err != nil {
					return err
				}
				if !info.IsDir() && strings.HasSuffix(info.Name(), ".jar") {
					fmt.Println(path)
				}
				return nil
			})
			if err != nil {
				return err
			}
		}
		return nil
	})
	if err != nil {
		fmt.Println("Ошибка при поиске файлов JAR в каталогах аддонов LabyMod:", err)
		os.Exit(1)
	}
}
