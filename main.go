package main

import (
	"encoding/json"
	"fmt"
	"log"
	"time"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/gocolly/colly/v2"
	"golang.org/x/text/language"
	"golang.org/x/text/message"

	_ "time/tzdata"
)

var currencyPrinter = message.NewPrinter(language.English)

func crawlData(user string, c *colly.Collector) (*ResponseAPI, error) {
	var err error
	data := ResponseAPI{}

	// Extract comment
	c.OnHTML("#__NEXT_DATA__", func(e *colly.HTMLElement) {
		// c.OnHTML("div.user-transaction-history-section table.table tr.last-item-of-day", func(e *colly.HTMLElement) {
		err = json.Unmarshal([]byte(e.Text), &data)
	})

	c.Visit("https://thiennguyen.app/user/sanchoicauvong")
	if err != nil {
		return nil, fmt.Errorf("unmarshal error: %s", err)
	}

	return &data, nil
}

func waitAndSend(bot *tgbotapi.BotAPI, chatID int64, transToSend chan *Transaction) {
	for trans := range transToSend {
		amount := currencyPrinter.Sprintf("%d", trans.TransactionAmount)

		transType := "Nhận tiền"
		if trans.Type == "DEBIT" {
			transType = "Chi tiền"
		}

		bot.Send(tgbotapi.NewMessage(chatID,
			fmt.Sprintf(`Thời gian: %s
		Loại: %s
		Số tiền: ₫ %s
		Phí: ₫ %d
		Nội dung: %s
		Người thụ hưởng: %s`, trans.transactionTime.Local().Format(time.DateTime), transType, amount, trans.FeeAmount, trans.Narrative, trans.UserName)))
	}
}

var lastItemTime time.Time = time.Now()

func main() {
	var c config

	if err := c.Load(); err != nil {
		log.Panic(fmt.Errorf("config load error: %s", err))
	}

	// bot API
	bot, err := tgbotapi.NewBotAPI(c.BotToken)
	if err != nil {
		log.Panic(err)
	}
	bot.Debug = true
	log.Printf("Authorized on account %s", bot.Self.UserName)

	transToSend := make(chan *Transaction)
	go waitAndSend(bot, c.GroupChatID, transToSend)

	bot.Send(tgbotapi.NewMessage(int64(c.AdminID), "Bot started"))
	blockingCrawlData(c, bot, transToSend)
	// crawler

}

func blockingCrawlData(c config, bot *tgbotapi.BotAPI, transToSend chan<- *Transaction) {
	// time format does not include in json data, explicitly set it (thiennguyen.app use Vietnam timezone)
	Ho_Chi_Minh_TZ, err := time.LoadLocation("Asia/Ho_Chi_Minh")
	if err != nil {
		bot.Send(tgbotapi.NewMessage(int64(c.AdminID), fmt.Sprintf("Error: %s", err)))
		log.Panic(fmt.Errorf("time load location error: %s", err))
	}
	crawler := colly.NewCollector()

	for {
		data, err := crawlData(c.UserID, crawler)
		if err != nil {
			bot.Send(tgbotapi.NewMessage(int64(c.AdminID), fmt.Sprintf("Error: %s", err)))
		} else {
			lastItemClone := lastItemTime
			lastItemTime = time.Now()
			groups := data.Props.PageProps.PageData.BankAccountTransactions.Groups
			for i := len(groups) - 1; i >= 0; i-- {
				transs := groups[i].Transactions

				for j := len(transs) - 1; j >= 0; j-- {
					trans := transs[j]
					trans.transactionTime, err = time.ParseInLocation("2006-01-02T15:04:05", trans.TransactionTime, Ho_Chi_Minh_TZ)
					if err != nil {
						bot.Send(tgbotapi.NewMessage(int64(c.AdminID), fmt.Sprintf("Error: %s", err)))
					} else {
						if trans.transactionTime.After(lastItemClone) {
							transToSend <- &trans
						}
					}

				}

			}
		}

		if c.Debug {
			log.Printf("Last item time: %s", lastItemTime)
		}
		time.Sleep(time.Duration(c.Interval) * time.Second)
	}

}
