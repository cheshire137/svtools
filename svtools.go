package main

import (
	"fmt"
	"os"

	"github.com/cheshire137/svtools/pkg/xml"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Printf("%s path_to_save_file\n\n", os.Args[0])
		fmt.Println("Example:")
		fmt.Printf("\t%s /Users/yourMacUser/.config/StardewValley/Saves/YourStardewCharacter_116989742/YourStardewCharacter_116989742\n", os.Args[0])
		os.Exit(1)
		return
	}

	saveFile := xml.SaveFile{Path: os.Args[1]}
	err := saveFile.Load()
	if err != nil {
		fmt.Println("failed to open save file: " + err.Error())
		os.Exit(1)
		return
	}

	fmt.Println(saveFile.Data)
}
