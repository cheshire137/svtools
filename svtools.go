package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/cheshire137/svtools/pkg/models"
)

func main() {
	var inPath string
	flag.StringVar(&inPath, "in", "",
		"Path to a Stardew Valley save file, e.g, /Users/yourMacUser/.config/StardewValley/Saves/YourStardewCharacter_116989742/YourStardewCharacter_116989742")

	var outPath string
	flag.StringVar(&outPath, "out", "",
		"Path to where a duplicate of the specified save file should be made")

	flag.Parse()
	if len(inPath) < 1 {
		flag.PrintDefaults()
		os.Exit(0)
		return
	}

	saveFile := models.NewSaveFile(inPath)
	err := saveFile.Load()
	if err != nil {
		fmt.Println("failed to open save file: " + err.Error())
		os.Exit(1)
		return
	}

	fmt.Println(saveFile.Data)

	if len(outPath) > 0 {
		err := saveFile.Save(outPath, true)
		if err != nil {
			fmt.Println("failed to copy save file: " + err.Error())
			os.Exit(1)
			return
		}
		fmt.Println("\n\nsaved copy of save file to " + outPath)
	}
}
