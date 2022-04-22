package main

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
	"os"
	"strconv"
	"tg_pranje_bot/queue"
)

func main() {
	bot, apiErr := tgbotapi.NewBotAPI(os.Getenv("TELEGRAM_APITOKEN"))
	if apiErr != nil {
		log.Panic(apiErr)
	}
	tgGroupId, convErr := strconv.Atoi(os.Getenv("TELEGRAM_GROUP_ID"))
	if convErr != nil {
		log.Panic(convErr)
	}

	bot.Debug = false
	log.Printf("Authorized on account %s", bot.Self.UserName)
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60
	updates := bot.GetUpdatesChan(u)

	for update := range updates {
		func() {
			defer func() {
				if r := recover(); r != nil {
					fmt.Printf("panic occured: %v\n", r)
				}
			}()

			if update.Message != nil {
				if update.Message.NewChatMembers != nil {
					for _, user := range update.Message.NewChatMembers {
						if user.UserName != "" && !user.IsBot {
							// Great new user added in a chat
							msg := tgbotapi.NewMessage(update.Message.Chat.ID, "")
							msg.Text = fmt.Sprintf("Bok @%s!\n Ja ovdje održavam red. Lijepo se ponašaj.", user.UserName)
						}
					}
				} else if update.Message.LeftChatMember != nil {
					msg := tgbotapi.NewMessage(update.Message.Chat.ID, "")
					msg.Text = fmt.Sprintf("Chao @%s!\n Vidimo se!", update.Message.LeftChatMember.UserName)
				}

				// ignore any non-Message updates
				if update.Message == nil {
					return
				}

				// ignore any non-command Messages
				if !update.Message.IsCommand() {
					return
				}

				// Create a new MessageConfig. We don't have text yet, so we leave it empty.
				msg := tgbotapi.NewMessage(update.Message.Chat.ID, "")
				// Extract the command from the Message.
				switch update.Message.Command() {
				case "start":
					msg.Text = "Hey!\nI'm a laundry queue keeper!\nUse the folowing commands:\nhelp - Request bot usage instructions\nqueue - Get the laundry queue\npush - Get in line for laundry\npop - Get out of the laundry queue"
				case "help":
					//msg.ReplyToMessageID = update.Message.MessageID
					msg.Text = "help - Request bot usage instructions\nqueue - Get the laundry queue\npush - Get in line for laundry\npop - Get out of the laundry queue"
				case "queue":
					//msg.ReplyToMessageID = update.Message.MessageID
					msg.Text = queue.PrintQueue()
				case "push":
					if queue.Push(update.Message.Chat.UserName) {
						//msg.ReplyToMessageID = update.Message.MessageID
						msg.Text = "You are get in line for laundry"
					} else {
						//msg.ReplyToMessageID = update.Message.MessageID
						msg.Text = "You are already in line for laundry"
					}
				case "pop":
					pop := queue.Pop(update.Message.Chat.UserName)
					if pop == "last" {
						//msg.ReplyToMessageID = update.Message.MessageID
						msg.Text = "You are got out from the laundry queue"
					} else if pop == "empty" {
						//msg.ReplyToMessageID = update.Message.MessageID
						msg.Text = "Laundry queue is already empty"
					} else if pop == "denied" {
						//msg.ReplyToMessageID = update.Message.MessageID
						msg.Text = "You can't get someone out of the laundry queue"
					} else {
						//msg.ReplyToMessageID = update.Message.MessageID
						msg = tgbotapi.NewMessage(int64(tgGroupId), "")
						msg.Text = fmt.Sprintf("Hey @%s!\n Are you ready for some loundry staf?\n You are next in the queue!\n So wake up and hurry! :D", pop)
					}
				default:
					//msg.ReplyToMessageID = update.Message.MessageID
					msg.Text = "Unknown command"
				}

				if _, err := bot.Send(msg); err != nil {
					log.Println(err)
				}
			}
		}()
	}
}
