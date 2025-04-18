package main
import (
	"fmt"
	"os"
)

func main(){
	var str,sep string
	for i :=1;i<len(os.Args);i++ {
		str += sep + os.Args[i]
	}
	fmt.Println(str)
}