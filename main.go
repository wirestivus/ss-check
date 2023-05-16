package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/rehellsing/ss-check/dll"
	"github.com/rehellsing/ss-check/evr"
	"github.com/rehellsing/ss-check/jar"
	"github.com/rehellsing/ss-check/mods"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Println("Выберите действие:")
		fmt.Println("0. Вывести DLL-файлы")
		fmt.Println("1. Вывести аддоны LabyMod")
		fmt.Println("2. Вывести обычные моды (.minecraft/mods)")
		fmt.Println("3. Вывести и DLL-файлы, и аддоны LabyMod, и моды")
		fmt.Println("4. Скачать и установить Everything")
		fmt.Println("5. Выйти")

		choice, err := readInput(reader)
		if err != nil {
			fmt.Println("Ошибка при чтении выбора:", err)
			continue
		}

		switch choice {
		case 0:
			dll.PrintDLLFiles()
		case 1:
			jar.PrintJARFiles()
		case 2:
			mods.PrintDFMods()
		case 3:
			dll.PrintDLLFiles()
			jar.PrintJARFiles()
			mods.PrintDFMods()
		case 4:
			evr.InstallEverything()
		case 5:
			fmt.Println("Выход")
			return
		default:
			fmt.Println("Неверный выбор. Попробуйте еще раз.")
		}

		fmt.Println("Нажмите Enter для продолжения.")
		fmt.Scanln()
	}
}

func readInput(reader *bufio.Reader) (int, error) {
	input, err := reader.ReadString('\n')
	if err != nil {
		return 0, err
	}

	input = strings.TrimSpace(input) // Удалить пробелы и символы новой строки

	choice, err := strconv.Atoi(input)
	if err != nil {
		return 0, err
	}

	return choice, nil
}
