package main

import (
	"mgrep/operations"
)


func main() {


	replacer := operations.MakeReplacer("cpp")
	replacer.Run(".", "")

}
