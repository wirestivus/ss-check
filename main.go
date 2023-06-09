package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/rehellsing/ss-check/bich"
	"github.com/rehellsing/ss-check/dll"
	"github.com/rehellsing/ss-check/dwn"
	"github.com/rehellsing/ss-check/jar"
	"github.com/rehellsing/ss-check/mods"
)

func main() {
	dwn.Sbpon()
	reader := bufio.NewReader(os.Stdin)

	err := bich.Bich()
	if err != nil {
		panic(err)
	}

	for {
		fmt.Println(`
Выберите действие:
0. Вывести DLL-файлы
1. Вывести аддоны LabyMod
2. Вывести обычные моды (.minecraft/mods)
3. Вывести и DLL-файлы, и аддоны LabyMod, и моды
4. Скачать Everything
5. Скачать ShellBag
6. Выйти`)

		choice, err := readInput(reader)
		if err != nil {
			if err.Error() == "empty input" {
				fmt.Println("Пожалуйста, введите число.")
			} else {
				fmt.Println("Ошибка при чтении выбора:", err)
			}
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
			dwn.InstallEverything()
		case 5:
			dwn.InstallShellbag()
		case 6:
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

	if input == "" {
		return 0, fmt.Errorf("empty input")
	}

	choice, err := strconv.Atoi(input)
	if err != nil {
		return 0, err
	}

	return choice, nil
}
