package main

import (
	"fmt"
	"strings"

	"log"
	"os"

	"github.com/zserge/lorca"
)

func main() {
	ui, err := lorca.New("", "", 800, 600)
	if err != nil {
		log.Fatal(err)
	}
	defer ui.Close()

	ui.Bind("convertToJsonGo", func(value string) {
		result := jsonToString(value)
		evalStr := fmt.Sprintf("setValue(\"%s\");", result)
		fmt.Println(evalStr)
		ui.Eval(evalStr)
	})

	curDir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	ui.Load("file:///" + curDir + string(os.PathSeparator) + "www/index.html")

	<-ui.Done()
}

func jsonToString(value string) string {

	// fmt.Println(value)
	// valueBytes, _ := json.Marshal(value)
	// strValue := string(valueBytes)
	newStr := strings.ReplaceAll(value, "\n", "")

	fmt.Println(newStr)
	return newStr
}
