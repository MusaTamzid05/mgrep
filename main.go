package main

import (
	"mgrep/operations"
)


func Replace(searchDir, srcStr, dstStr, fileExt string) {
	replacer := operations.MakeReplacer(fileExt)
	replacer.Run(searchDir, srcStr, dstStr)
}

func Match(searchDir, srcStr string) {
	matcher := operations.MakeMatcher("cpp")
	matcher.Run(".", "zet")
}



func main() {
	//Replace(".", ",main", "zet", "cpp")

	Match(".", "zet")

}
