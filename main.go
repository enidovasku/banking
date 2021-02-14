package main

import (
	"github.com/enidovasku/banking/app"
	"github.com/enidovasku/banking/logger"
)

func main() {
	logger.Info("Starting the application!")
	app.Start()
}
