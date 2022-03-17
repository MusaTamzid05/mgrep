package operations

import (
	"path/filepath"
	"os"
	"fmt"
	"strings"
)

type Replacer struct {
	fileExt string
}

func MakeReplacer(fileExt string) Replacer {
	return Replacer{fileExt : fileExt}
}

func (r* Replacer) Run(dirPath , target string) {
	err := filepath.Walk(dirPath, func(path string, info os.FileInfo, err error) error {

		if strings.HasSuffix(path, r.fileExt) == false {
			return nil
		}

		fmt.Println(path)

		return nil

	})

	if err != nil {
		fmt.Println(err)
	}
}

