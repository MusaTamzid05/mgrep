package operations

import (
	"os"
	"bufio"
	"strings"
	"fmt"
)

func ReadLines(filePath string) ([]string, error ) {
	lines := []string{}

	file, err := os.Open(filePath)

	if err != nil {
		return lines, err

	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	return lines, nil
}

func ShowMatches(filePath string, target string) error  {

	lines, err := ReadLines(filePath)

	if err != nil {
		return err
	}

	fmt.Println("Path\t\tIndex\t\tLine")

	for index, line := range lines {
		if strings.Contains(line, target) {
			fmt.Println(filePath, "\t", index + 1, "\t\t", line)
		}
	}

	return nil
}
