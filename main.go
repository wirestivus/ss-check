package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/shirou/gopsutil/process"
)

func main() {

	downloadsDir, err := filepath.Abs(filepath.Join(os.Getenv("USERPROFILE"), "Downloads"))
	if err != nil {
		fmt.Println("")
		fmt.Println("Ошибка при получении каталога загрузок:", err)
		os.Exit(1)
	}

	// Ищем все .dll файлы в папке Downloads
	fmt.Println("")
	fmt.Println("DLL-файлы в каталоге Downloads:")
	fmt.Println("")
	err = filepath.Walk(downloadsDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() && strings.HasSuffix(info.Name(), ".dll") {
			fmt.Println(path)
		}
		return nil
	})
	if err != nil {
		fmt.Println("Ошибка при поиске файлов DLL в каталоге Downloads:", err)
		os.Exit(1)
	}

	labymodDir := filepath.Join(os.Getenv("APPDATA"), ".minecraft", "LabyMod")
	fmt.Println("")
	fmt.Printf("Поиск JAR файлов в каталогах аддонов LabyMod в: \n")
	fmt.Printf("%s\n", labymodDir)
	fmt.Println("")
	fmt.Printf("JAR-файлы в каталогах аддонов LabyMod:")
	fmt.Println("")
	err = filepath.Walk(labymodDir, func(path string, info os.FileInfo, err error) error {
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

	// проверяем существование пути
	_, err = os.Stat(labymodDir)
	if err != nil {
		fmt.Printf("Ошибка: не удалось найти каталог LabyMod: \n")
		fmt.Printf("%s\n", labymodDir)
	}

	// Проверяем, загружены ли какие-то .dll файлы в процесс javaw.exe
	processes, err := process.Processes()
	if err != nil {
		fmt.Println("")
		fmt.Println("Ошибка при получении процессов:", err)
		os.Exit(1)
	}
	dllInjected := false
	for _, proc := range processes {
		name, err := proc.Name()
		if err == nil && strings.ToLower(name) == "javaw.exe" {
			cmdline, err := proc.Cmdline()
			if err == nil && strings.Contains(cmdline, ".dll") && strings.HasSuffix(cmdline, ".dll") {
				fmt.Println("")
				fmt.Println("DLL-файл внедрен в процесс javaw.exe с PID:", proc.Pid)
				dllInjected = true
			}
		}
	}
	if !dllInjected {
		fmt.Println("")
		fmt.Println("Ни один DLL-файл не внедряется ни в один процесс javaw.exe")
	}

	fmt.Println("Нажмите Enter для выхода.")
	fmt.Scanln() // Ожидание нажатия клавиши Enter
	os.Exit(0)
}
