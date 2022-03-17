package main

import (
	"mgrep/operations"
)


func main() {


	replacer := operations.MakeReplacer("cpp")
	replacer.Run("/Users/musakhan/c++_pro/image_operation_from_scratch", "")

}
