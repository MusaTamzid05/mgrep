package main

import (
	"mgrep/operations"
	"flag"
	"fmt"
	"os"
)


func Replace(searchDir, srcStr, dstStr, fileExt string) {
	replacer := operations.MakeReplacer(fileExt)
	replacer.Run(searchDir, srcStr, dstStr)
}

func Match(searchDir, srcStr, fileExt string) {
	matcher := operations.MakeMatcher(fileExt)
	matcher.Run(searchDir, srcStr)
}



func main() {
	searchDirStrPtr := flag.String("search", "", " Directory to search")
	patternStrPtr := flag.String("pattern", "", "Search pattern")
	fileExtStrPtr := flag.String("ext", "", "file exts from which to search")
	replaceStrPtr := flag.String("replace", "", "The new replcae str")

	flag.Parse()


	if *searchDirStrPtr == "" {
		fmt.Println("Must specify a search diirectory")
		fmt.Println("usage: -search")
		os.Exit(1)
	}


	if *patternStrPtr == "" {
		fmt.Println("Must specify a pattern")
		fmt.Println("usage: -pattern")
		os.Exit(1)
	}


	if *fileExtStrPtr == "" {
		fmt.Println("Must specify a file ext")
		fmt.Println("usage: -ext ")
		os.Exit(1)
	}

	if *replaceStrPtr != "" {
		Replace(*searchDirStrPtr, *patternStrPtr, *replaceStrPtr, *fileExtStrPtr)
		return
	}

	Match(*searchDirStrPtr, *patternStrPtr, *fileExtStrPtr)

}
