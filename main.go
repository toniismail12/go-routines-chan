package main

import (
	"bufio"
	"ch-export/database"
	"ch-export/handler"
	"fmt"
	"os"
	"strings"
)

func main() {

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("---------- import data create by toni ismail ------------")
	fmt.Print("nama file: ")

	text, _ := reader.ReadString('\n')

	text = strings.Replace(text, "\n", "", -1)

	// connection database
	database.Connect()

	// insert file name
	nama_file := text + ".xlsx"

	handler.Import(nama_file)

}
