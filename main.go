package main

import (
	"fmt"
	"os"

	"github.com/rehellsing/ss-check/dll"
	"github.com/rehellsing/ss-check/jar"
)

func main() {
	dll.PrintDLLFiles()
	jar.PrintJARFiles()
	process.CheckInjectedDLL()
	fmt.Println("Нажмите Enter для выхода.")
	fmt.Scanln() // Ожидание нажатия клавиши Enter
	os.Exit(0)
}
