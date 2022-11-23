package main

import (
	"fmt"
	"log"
	"os"
	"github.com/joho/godotenv"
	"go-es-testcode/src/infrastructure"
)

func main() {
		// 環境変数ファイルの読込
		err := godotenv.Load(fmt.Sprintf("config/%s.env", os.Getenv("GO_ENV")))
		if err != nil {
			log.Fatal(err)
		}
		r := infrastructure.NewRouting()
		infrastructure.Run(r)
}
