package tcpClient

import (
	"bufio"
	"fmt"
	"os"

	"../define"
)

func ReadLine(path string, lineNumber int) string {
	file, err := os.Open(path)
	if err != nil {
		return ""
	}
	fileScanner := bufio.NewScanner(file)
	lineCount := 1
	for fileScanner.Scan() {
		if lineCount == lineNumber {
			return fileScanner.Text()
		}
		lineCount++
	}
	defer file.Close()
	return ""
}

func WriteLine(file, msg string) int {
	f, err := os.OpenFile(file, os.O_CREATE|os.O_APPEND|os.O_RDWR, 0660)
	if err != nil {
		fmt.Printf("Cannot open file %s!\n", file)
		return define.Err_open
	}
	defer f.Close()
	f.WriteString(msg + "\n")
	return define.Err_null
}

func CountsLine(path string) int {
	file, err := os.Open(path)
	if err != nil {
		return 0
	}
	defer file.Close()
	fd := bufio.NewReader(file)
	count := 0
	for {
		_, err := fd.ReadString('\n')
		if err != nil {
			break
		}
		count++

	}
	return count
}

// func main() {
// 	WriteLine("d:/t.log", "ggg")
// 	msg := ReadLine("d:/t.log", 6)
// 	fmt.Println(msg)
// }
