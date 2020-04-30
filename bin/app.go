package main

import (
	"github.com/just1689/ttt/io"
	"github.com/sirupsen/logrus"
	"os"
)

func main() {
	logrus.Println("Starting server")
	io.StartServer(os.Getenv("listen"))
	select {}
}
