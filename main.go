package main

import (
	"mgrep/operations"
)


func Replace(searchDir, srcStr, dstStr, fileExt string) {
	replacer := operations.MakeReplacer(fileExt)
	replacer.Run(searchDir, srcStr, dstStr)
}



func main() {
	Replace(".", ",main", "zet", "cpp")

}
