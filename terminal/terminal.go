package terminal

import (
	"bufio"
	"os"
	"strings"
)

func Input() string {
	reader := bufio.NewReader(os.Stdin)
	entered, err := reader.ReadString('\n')
	if err != nil {
		panic(err)
	}

	return strings.TrimSpace(entered)
}
