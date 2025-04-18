package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"os/exec"
	"os"
)

func mkdirDir() {
	//查询数据
	rows, err := db.Query("")
	checkErr(err)
	for rows.Next() {
		var shell string
		err = rows.Scan(&shell)
		cmd := exec.Command("mkdir","-p",shell)
		out, err := cmd.CombinedOutput()
		if err != nil {
			fmt.Println(err)
		}else{
			fmt.Println(out)

		}
	}
}

var db *sql.DB

var create_time string

func main() {
	var create_date = os.Args[1]
	var time = os.Args[2]
	create_time = create_date+" "+time
	var err error
	db, err = sql.Open("postgres", "host= port= user= password= dbname=meta sslmode=disable")
	checkErr(err)
	fmt.Println(err)
	mkdirDir()
	rows, err := db.Query("")
	checkErr(err)
	for rows.Next() {
		var cpsource string 
		var copydesc string
		err = rows.Scan(&cpsource,&copydesc)
		fmt.Println(cpsource)
		fmt.Println(copydesc)
		cmd := exec.Command("cp",cpsource,copydesc)
		out, err := cmd.CombinedOutput()
		if err != nil {
			fmt.Println(err)
		}else{
			fmt.Println(out)

		}
	}

}
func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
