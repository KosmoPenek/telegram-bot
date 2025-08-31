package main

import (
	"log"

	tg "github.com/KosmoPenek/telegram-bot/internal/clients"
	"github.com/KosmoPenek/telegram-bot/internal/config"
	"github.com/KosmoPenek/telegram-bot/internal/model/messages"
)

func main() {
	config, err := config.New()
	if err != nil {
		log.Fatal("config init failed: ", err)
	}

	tgClient, err := tg.New(config)
	if err != nil {
		log.Fatal("tgClient init failed: ", err)
	}

	msgModel := messages.New(tgClient)

	tgClient.ListenUpdates(msgModel)
}
