package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Printf("%s path_to_save_file\n\n", os.Args[0])
		fmt.Println("Example:")
		fmt.Printf("\t%s /Users/yourMacUser/.config/StardewValley/Saves/YourStardewCharacter_116989742/YourStardewCharacter_116989742\n", os.Args[0])
		os.Exit(1)
		return
	}

	saveFilePath := os.Args[1]
	fmt.Println(saveFilePath)
}
