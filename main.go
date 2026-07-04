package main

import (
	"fmt"
	"os"

	"JacquesFernandes/gopro/dirops"
	"JacquesFernandes/gopro/gopro"
)

func main() {
	goProPath, err := gopro.GetGoProPath()
	if err != nil {
		panic(err)
	}

	fmt.Println("goProPath:", goProPath)
	goProFiles, err := os.ReadDir(goProPath)
	if err != nil {
		panic(err)
	}

	groups, err := dirops.GroupDirEntriesByDate(goProFiles)
	if err != nil {
		panic(err)
	}

	for _, group := range groups {
		fmt.Println("Group for date", group.Date)
		for index, dirEntry := range group.Entries {
			fmt.Println(index, "-", dirEntry.Name())
		}
	}
}
