package operations

import (
	"os"
	"bufio"
	"strings"
	"fmt"
	"path/filepath"
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


type Matcher struct {
	fileExt string
}

func MakeMatcher(fileExt string) Matcher{
	return Matcher{fileExt : fileExt}
}


func (m* Matcher) Run(dirPath , target string) {
	err := filepath.Walk(dirPath, func(path string, info os.FileInfo, err error) error {

		if strings.HasSuffix(path, m.fileExt) == false {
			return nil
		}

		err =  m.showMatches(path, target)

		if err != nil {
			return err
		}
		return nil

	})

	if err != nil {
		fmt.Println(err)
	}
}


func (m* Matcher) showMatches(filePath string, target string) error  {

	lines, err := ReadLines(filePath)

	if err != nil {
		return err
	}

	//fmt.Println("Path\t\tIndex\t\tLine")

	for index, line := range lines {
		if strings.Contains(line, target) {
			//fmt.Println(filePath, "\t", index + 1, "\t\t", line)
			m.prettyPrint(filePath, line, index)
		}
	}

	return nil
}

func (m *Matcher) prettyPrint(filePath, line string, index int) {
	filePathSpaceCount := 80
	spaceCount := filePathSpaceCount - len(filePath)
	fmt.Printf("%s ", filePath)

	m.generateSpace(spaceCount)

	fmt.Printf("%s ", line)
	spaceCount = filePathSpaceCount - len(line)

	m.generateSpace(spaceCount)

	fmt.Printf("%d ", index)
	fmt.Printf("\n")
}

func (m *Matcher) generateSpace(spaceCount int) {
	for i := 0; i < spaceCount; i += 1 {
		fmt.Printf(" ")
	}
}
