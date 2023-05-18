package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"sync"

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

	var wg sync.WaitGroup

	for {
		fmt.Println("Выберите действие:")
		fmt.Println("0. Вывести DLL-файлы")
		fmt.Println("1. Вывести аддоны LabyMod")
		fmt.Println("2. Вывести обычные моды (.minecraft/mods)")
		fmt.Println("3. Вывести и DLL-файлы, и аддоны LabyMod, и моды")
		fmt.Println("4. Скачать Everything")
		fmt.Println("5. Скачать ShellBag")
		fmt.Println("6. Выйти")

		choice, err := readInput(reader)
		if err != nil {
			fmt.Println("Ошибка при чтении выбора:", err)
			continue
		}

		switch choice {
		case 0:
			wg.Add(1)
			go func() {
				defer wg.Done()
				dll.PrintDLLFiles()
			}()
		case 1:
			wg.Add(1)
			go func() {
				defer wg.Done()
				jar.PrintJARFiles()
			}()
		case 2:
			wg.Add(1)
			go func() {
				defer wg.Done()
				mods.PrintDFMods()
			}()
		case 3:
			wg.Add(3)
			go func() {
				defer wg.Done()
				dll.PrintDLLFiles()
			}()
			go func() {
				defer wg.Done()
				jar.PrintJARFiles()
			}()
			go func() {
				defer wg.Done()
				mods.PrintDFMods()
			}()
		case 4:
			wg.Add(1)
			go func() {
				defer wg.Done()
				dwn.InstallEverything()
			}()
		case 5:
			wg.Add(1)
			go func() {
				defer wg.Done()
				dwn.InstallShellbag()
			}()
		case 6:
			fmt.Println("Выход")
			return
		default:
			fmt.Println("Неверный выбор. Попробуйте еще раз.")
		}

		fmt.Println("Нажмите Enter для продолжения.")
		fmt.Scanln()
	}

	wg.Wait()
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
