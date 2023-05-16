package mods

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func PrintDFMods() {
	modsDir := filepath.Join(os.Getenv("APPDATA"), ".minecraft")
	fmt.Println("")
	fmt.Printf("Список модов в: \n")
	fmt.Printf("%s\n", modsDir)
	fmt.Println("")
	fmt.Printf("Моды в каталогах:")
	fmt.Println("")
	err := filepath.Walk(modsDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() && strings.HasPrefix(info.Name(), "mods") {
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
		fmt.Println("Ошибка при поиске модов в каталогах:", err)
		os.Exit(1)
	}
}
