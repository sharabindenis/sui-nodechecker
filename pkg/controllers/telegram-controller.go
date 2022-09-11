package controllers

import (
	"fmt"
	tgbotapi "github.com/Syfaro/telegram-bot-api"
	"reflect"
	"strconv"
)

func TelegramBot() {
	bot, err := tgbotapi.NewBotAPI("5436979593:AAEEoh_5A6-OeUEdXEdnTBeRcMbdCMk1hXk")
	if err != nil {
		panic(err)
	}
	//fmt.Println(bot)
	bot.Debug = false

	//Устанавливаем время обновления
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	//Получаем обновления от бота
	updates, err := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message == nil {
			continue
		}
		if reflect.TypeOf(update.Message.Text).Kind() == reflect.String && update.Message.Text != "" {
			switch update.Message.Text {
			case "/start":
				//Отправлем сообщение
				msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Hi, i'm sui checker. Configure me for alert or info messages. Type /")
				bot.Send(msg)
				fmt.Println(msg)
			case "/ip":
				//fmt.Println(SuiTaskHolder)
				ser := SuiTaskHolder[Sch.Ip]
				//fmt.Println(ser.Ip)
				msg := tgbotapi.NewMessage(update.Message.Chat.ID, ser.Ip)
				bot.Send(msg)
				//fmt.Println(msg)
			default:
				ms := update.Message.Text
				//fmt.Println(ms)
				CreateScheduleByBot(ms, update.Message.Chat.ID)

				msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Checker started on "+ms)
				bot.Send(msg)
				////fmt.Println(msg)
			}

		}
	}
}

func TelegramBotAlert(chtid int64, tt int) {
	bot, err := tgbotapi.NewBotAPI("5436979593:AAEEoh_5A6-OeUEdXEdnTBeRcMbdCMk1hXk")
	if err != nil {
		panic(err)
	}
	fmt.Println(bot)
	bot.Debug = false

	t := strconv.Itoa(tt)
	msg := tgbotapi.NewMessage(chtid, t)
	bot.Send(msg)

}
