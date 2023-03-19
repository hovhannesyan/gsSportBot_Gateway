package main

import (
	"fmt"
	"github.com/NicoNex/echotron/v3"
	"github.com/hovhannesyan/gsSportBot_Gateway/pkg/telegram"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	botToken := os.Getenv("TOKEN")
	userServiceURL := os.Getenv("USER_SERVICE_URL")
	eventServiceURL := os.Getenv("EVENT_SERVICE_URL")
	connectionServiceURL := os.Getenv("CONNECTION_SERVICE_URL")

	newBot := telegram.NewBot(botToken, userServiceURL, eventServiceURL, connectionServiceURL)
	dsp := echotron.NewDispatcher(botToken, newBot)

	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		if err := dsp.Poll(); err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	}()

	<-sigs

	fmt.Println("Stopping...")

	//bot.StopLongPolling()

	fmt.Println("Long polling done")
	fmt.Println("Done")
}
