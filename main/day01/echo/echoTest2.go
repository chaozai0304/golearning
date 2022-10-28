package main
import (
	"fmt"
	"os"
)

// func main(){
// 	str, sep := "",""
// 	for i, arg := range os.Args[1:] {
// 		fmt.Println(i)
// 		str += sep + arg
// 		sep=" "
// 	}
// 	fmt.Println(str)
// }
func main(){
	str, sep := "",""
	for _, arg := range os.Args[1:] {
		str += sep + arg
		sep=" "
	}
	fmt.Println(str)
}