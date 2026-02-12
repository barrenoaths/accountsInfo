package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

const (
	bold  = "\033[1m"
	reset = "\033[0m"
)

func printBold(label, value string) {
	fmt.Printf("%s%s %s %s\n", bold, label, reset, value)
}

func main() {

	file, err := os.Open("/etc/passwd")
	if err != nil {
		fmt.Println("Unable to read /etc/passwd file")
		return
	}
	defer file.Close()

	var accounstList []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		accounstList = append(accounstList, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file: ", err)
	}

	for _, singleAccount := range accounstList {
		fields := strings.Split(singleAccount, ":")
		if fields[6] != "/usr/sbin/nologin" {
			// fmt.Printf("%sUsername:%s \t%s\n%sUID:%s \t\t%s\n%sGID:%s \t\t%s\n%sShell:%s \t\t%s\n\n",
			// 	bold, reset, fields[0],
			// 	bold, reset, fields[2],
			// 	bold, reset, fields[3],
			// 	bold, reset, fields[6])
			printBold("Username:\t", fields[0])
			printBold("UID:\t\t", fields[2])
			printBold("GID:\t\t", fields[3])
			printBold("Shell:\t\t", fields[6])
			fmt.Println()
		}

	}

}
