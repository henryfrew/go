package util

import (
	"bufio"
	"fmt"
	"os"
)

func input(msg string) string {
	fmt.Println(msg)
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	line := scanner.Text()
	return string(line)
}
