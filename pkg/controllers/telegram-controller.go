package controllers

import (
	"fmt"
	tgbotapi "github.com/Syfaro/telegram-bot-api"
	"log"
	regexp2 "regexp"
	"strconv"
	"strings"
)

const TOKEN = "5436979593:AAEEoh_5A6-OeUEdXEdnTBeRcMbdCMk1hXk"

func TelegramBot() {
	bot, err := tgbotapi.NewBotAPI(TOKEN)
	if err != nil {
		panic(err)
	}
	//fmt.Println(bot)
	bot.Debug = false
	log.Printf("Authorized on account %s", bot.Self.UserName)

	//Устанавливаем время обновления
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	//Получаем обновления от бота
	updates, err := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message == nil { // ignore any non-Message updates
			continue
		}

		if !update.Message.IsCommand() { // ignore any non-command Messages
			continue
		}

		// Create a new MessageConfig. We don't have text yet,
		// so we leave it empty.
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, "")

		// Extract the command from the Message.
		switch update.Message.Command() {
		case "ip":
			beforeplace := update.Message.Text
			afterreplace := strings.ReplaceAll(beforeplace, "/ip ", "")
			regexp := regexp2.MustCompile(`(25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)(\.(25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)){3}`)
			fmt.Println(regexp.MatchString(afterreplace))
			if regexp.MatchString(afterreplace) {
				ip := "http://" + afterreplace + ":9000"
				CreateScheduleByBot(ip, update.Message.Chat.ID)
				msg.Text = "Ok! Schedule started for " + ip
			} else {
				msg.Text = "Not valid ip, send me fullnode address like 1.1.1.1"
			}
		case "start":
			msg.Text = "Hi :) Send me fullnode ip like /ip 1.1.1.1 and I will start checker every 30 seconds"
		case "check":
			beforeplace := update.Message.Text
			afterreplace := strings.ReplaceAll(beforeplace, "/check ", "")
			regexp := regexp2.MustCompile(`(25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)(\.(25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)){3}`)
			fmt.Println(regexp.MatchString(afterreplace))
			if regexp.MatchString(afterreplace) {
				ip := "http://" + afterreplace + ":9000"
				ok := CheckIp(ip)
				fmt.Println(ok)
				if ok != nil {
					msg.Text = ok.Error()
				} else {
					msg.Text = "`" + afterreplace + "`" + " Ok!"
				}
			} else {
				msg.Text = "Not valid ip, send me fullnode address like 1.1.1.1"
			}
		default:
			msg.Text = "I don't know that command"
		}

		if _, err := bot.Send(msg); err != nil {
			log.Println(err)
		}
	}

	//for update := range updates {
	//	if update.Message == nil {
	//		continue
	//	}
	//	if reflect.TypeOf(update.Message.Text).Kind() == reflect.String && update.Message.Text != "" {
	//		switch update.Message.Text {
	//		case "/start":
	//			//Отправлем сообщение
	//			msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Hi, i'm sui checker. Configure me for alert or info messages. Type /")
	//			bot.Send(msg)
	//			fmt.Println(msg)
	//		case "/ip":
	//			//fmt.Println(SuiTaskHolder)
	//			ser := SuiTaskHolder[Sch.Ip]
	//			//fmt.Println(ser.Ip)
	//			msg := tgbotapi.NewMessage(update.Message.Chat.ID, ser.Ip)
	//			bot.Send(msg)
	//			//fmt.Println(msg)
	//		default:
	//			ms := update.Message.Text
	//			//fmt.Println(ms)
	//			CreateScheduleByBot(ms, update.Message.Chat.ID)
	//
	//			msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Checker started on "+ms)
	//			bot.Send(msg)
	//			////fmt.Println(msg)
	//		}
	//
	//	}
	//}
}

func TelegramBotAlert(chtid int64, tt int, ip string) {
	bot, err := tgbotapi.NewBotAPI(TOKEN)
	if err != nil {
		panic(err)
	}
	//fmt.Println(bot)
	//bot.Debug = false
	t := strconv.Itoa(tt)
	mssg := "Total transaction " + t + " for " + ip
	msg := tgbotapi.NewMessage(chtid, mssg)
	bot.Send(msg)
}
