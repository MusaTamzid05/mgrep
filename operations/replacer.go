package operations

import (
	"path/filepath"
	"os"
	"fmt"
	"strings"
	"bufio"
)


type Replacer struct {
	fileExt string
}

func MakeReplacer(fileExt string) Replacer {
	return Replacer{fileExt : fileExt}
}

func (r* Replacer) Run(dirPath , target, newStr string) {
	err := filepath.Walk(dirPath, func(path string, info os.FileInfo, err error) error {

		if strings.HasSuffix(path, r.fileExt) == false {
			return nil
		}

		lines, err := ReadLines(path)

		if err != nil {
			fmt.Println(err)
			return err
		}

		indexes := r.collectMatchData(lines, target)

		if len(indexes) == 0 {
			return  nil
		}

		replacedLines := r.replace(lines, indexes, target, newStr)

		r.replaceFile(path, replacedLines)
		
		return nil

	})

	if err != nil {
		fmt.Println(err)
	}
}

func (r * Replacer) collectMatchData(lines []string, target string) []int {
	// this slow but will do for now

	indexes := []int{}

	for i, line := range lines {
		if strings.Contains(line, target) == false {
			continue
		}

		indexes = append(indexes,i )

	}

	return indexes
}

func (r *Replacer) replace(lines []string, indexes []int, target, newStr string) []string {

	for _ , index := range indexes {
		line := lines[index]
		newLine := strings.Replace(line, target, newStr, -1)
		lines[index] = newLine
	}

	return lines


}


func (r *Replacer) replaceFile(path string, lines [] string) {
	err := os.Remove(path)

	if err != nil {
		fmt.Println(err)
		return
	}

	r.saveLines(path, lines)

}

func (r* Replacer) saveLines(path string, lines [] string) {
	file, err := os.OpenFile(path, os.O_APPEND | os.O_CREATE | os.O_WRONLY, 0644)

	if err != nil {
		fmt.Println(err)
		return
	}

	defer file.Close()

	dataWriter := bufio.NewWriter(file)

	for _, line:= range lines {
		 dataWriter.WriteString(line + "\n")
	}
	dataWriter.Flush()
}
