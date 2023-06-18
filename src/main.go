package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"go-es-testcode/src/infrastructure"
	"log"
	"os"
)

func main() {
	// 環境変数ファイルの読込
	err := godotenv.Load(fmt.Sprintf("config/%s.env", os.Getenv("GO_ENV")))
	if err != nil {
		log.Fatal(err)
	}
	r := infrastructure.NewRouter()
	infrastructure.Run(r)
}
