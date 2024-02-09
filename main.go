package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"time"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/gocolly/colly/v2"
	"github.com/ilyakaznacheev/cleanenv"
	"golang.org/x/text/language"
	"golang.org/x/text/message"
)

type config struct {
	AdminID     int    `json:"admin_id" env:"ADMIN_ID" env-default:""`
	GroupChatID int    `json:"group_chat_id" env:"GROUPCHAT_ID" env-default:""`
	BotToken    string `json:"bot_token" env:"BOT_TOKEN" env-default:""`
}

type ResponseAPI struct {
	Props struct {
		PageProps struct {
			PageData struct {
				User struct {
					ID           string `json:"id"`
					Fullname     string `json:"fullname"`
					Username     string `json:"username"`
					ProfilePhoto struct {
						Small struct {
							URL       string `json:"url"`
							Dimension struct {
								Width  int `json:"width"`
								Height int `json:"height"`
							} `json:"dimension"`
						} `json:"small"`
						Medium struct {
							URL       string `json:"url"`
							Dimension struct {
								Width  int `json:"width"`
								Height int `json:"height"`
							} `json:"dimension"`
						} `json:"medium"`
						Large struct {
							URL       string `json:"url"`
							Dimension struct {
								Width  int `json:"width"`
								Height int `json:"height"`
							} `json:"dimension"`
						} `json:"large"`
					} `json:"profilePhoto"`
					Verified            bool   `json:"verified"`
					Followed            bool   `json:"followed"`
					EnableDonate        bool   `json:"enableDonate"`
					Blocked             bool   `json:"blocked"`
					Gender              string `json:"gender"`
					ShowFrame           bool   `json:"showFrame"`
					FramePhoto          string `json:"framePhoto"`
					DefaultProfilePhoto bool   `json:"defaultProfilePhoto"`
					DefaultBannerPhoto  bool   `json:"defaultBannerPhoto"`
					CreatedTime         string `json:"createdTime"`
					UserType            int    `json:"userType"`
					CountDonate         int    `json:"countDonate"`
					CountLike           int    `json:"countLike"`
					Amount              any    `json:"amount"`
					DonateTargetID      any    `json:"donateTargetId"`
					Rank                int    `json:"rank"`
					TotalDonate         any    `json:"totalDonate"`
					TotalDonator        any    `json:"totalDonator"`
					Reward              any    `json:"reward"`
					TotalMoney          any    `json:"totalMoney"`
					TotalOrg            any    `json:"totalOrg"`
					BannerPhoto         struct {
						Small struct {
							URL       string `json:"url"`
							Dimension struct {
								Width  int `json:"width"`
								Height int `json:"height"`
							} `json:"dimension"`
						} `json:"small"`
						Medium struct {
							URL       string `json:"url"`
							Dimension struct {
								Width  int `json:"width"`
								Height int `json:"height"`
							} `json:"dimension"`
						} `json:"medium"`
						Large struct {
							URL       string `json:"url"`
							Dimension struct {
								Width  int `json:"width"`
								Height int `json:"height"`
							} `json:"dimension"`
						} `json:"large"`
					} `json:"bannerPhoto"`
					ReceivedMoney any    `json:"receivedMoney"`
					Biography     string `json:"biography"`
					NotedMoney    any    `json:"notedMoney"`
					NbDonator     int    `json:"nbDonator"`
					NbDonation    int    `json:"nbDonation"`
					NbFollower    int    `json:"nbFollower"`
					NbLike        int    `json:"nbLike"`
					FacebookURL   string `json:"facebookUrl"`
					YoutubeURL    any    `json:"youtubeUrl"`
					TwitterURL    any    `json:"twitterUrl"`
					InstagramURL  any    `json:"instagramUrl"`
					TiktokURL     any    `json:"tiktokUrl"`
					BankAccounts  []struct {
						ID                        string `json:"id"`
						BankCode                  string `json:"bankCode"`
						BankName                  string `json:"bankName"`
						BankPhoto                 string `json:"bankPhoto"`
						AccountNo                 string `json:"accountNo"`
						AccountName               string `json:"accountName"`
						AccountBalance            int    `json:"accountBalance"`
						SupportTransactionHistory bool   `json:"supportTransactionHistory"`
						DefaultAccount            bool   `json:"defaultAccount"`
						UpdatedTime               string `json:"updatedTime"`
					} `json:"bankAccounts"`
					DefaultBankAccount struct {
						ID                        string `json:"id"`
						BankCode                  string `json:"bankCode"`
						BankName                  string `json:"bankName"`
						BankPhoto                 string `json:"bankPhoto"`
						AccountNo                 string `json:"accountNo"`
						AccountName               string `json:"accountName"`
						AccountBalance            int    `json:"accountBalance"`
						SupportTransactionHistory bool   `json:"supportTransactionHistory"`
						DefaultAccount            bool   `json:"defaultAccount"`
						UpdatedTime               string `json:"updatedTime"`
					} `json:"defaultBankAccount"`
					DefaultBankAccountVietQrCode string `json:"defaultBankAccountVietQrCode"`
					Permalink                    string `json:"permalink"`
					RecentDonators               []struct {
						User      any  `json:"user"`
						Incognito bool `json:"incognito"`
						Amount    int  `json:"amount"`
						UserName  any  `json:"userName"`
					} `json:"recentDonators"`
					NbPost          int    `json:"nbPost"`
					ActiveStatus    int    `json:"activeStatus"`
					Email           string `json:"email"`
					Phone           any    `json:"phone"`
					Website         any    `json:"website"`
					ContactEmail    string `json:"contactEmail"`
					Representative  any    `json:"representative"`
					FieldOfActivity any    `json:"fieldOfActivity"`
					Address         any    `json:"address"`
					Incognito       bool   `json:"incognito"`
					AdminChat       bool   `json:"adminChat"`
				} `json:"user"`
				BankAccountOverview struct {
					ID                        string `json:"id"`
					BankCode                  string `json:"bankCode"`
					BankName                  string `json:"bankName"`
					BankPhoto                 string `json:"bankPhoto"`
					AccountNo                 string `json:"accountNo"`
					AccountName               string `json:"accountName"`
					AccountBalance            int    `json:"accountBalance"`
					SupportTransactionHistory bool   `json:"supportTransactionHistory"`
					DefaultAccount            bool   `json:"defaultAccount"`
					UpdatedTime               string `json:"updatedTime"`
					TotalDebit                int    `json:"totalDebit"`
					TotalCredit               int    `json:"totalCredit"`
				} `json:"bankAccountOverview"`
				BankAccountTransactions struct {
					Groups []struct {
						Date         string `json:"date"`
						Count        int    `json:"count"`
						Transactions []struct {
							ID                      string `json:"id"`
							BankAccountID           string `json:"bankAccountId"`
							TransactionTime         string `json:"transactionTime"`
							Type                    string `json:"type"`
							Method                  string `json:"method"`
							TransactionAmount       int    `json:"transactionAmount"`
							FeeAmount               int    `json:"feeAmount"`
							OtherBankCode           string `json:"otherBankCode"`
							OtherBankName           string `json:"otherBankName"`
							OtherBankPhoto          string `json:"otherBankPhoto"`
							OtherAccountDisplayName string `json:"otherAccountDisplayName"`
							OtherAccountNo          string `json:"otherAccountNo"`
							OtherAccountName        string `json:"otherAccountName"`
							Narrative               string `json:"narrative"`
							NarrativeHiddenConfig   any    `json:"narrativeHiddenConfig"`
							HiddenNarrative         bool   `json:"hiddenNarrative"`
							Donate                  any    `json:"donate"`
							ShowDonateMessage       bool   `json:"showDonateMessage"`
							Note                    any    `json:"note"`
							HiddenNote              bool   `json:"hiddenNote"`
							UserName                string `json:"userName"`
							BankAccountNo           any    `json:"bankAccountNo"`
						} `json:"transactions"`
					} `json:"groups"`
					Count      int    `json:"count"`
					NextPageID string `json:"nextPageId"`
				} `json:"bankAccountTransactions"`
				BankAccountReports []struct {
					Date        string `json:"date"`
					TotalDebit  int    `json:"totalDebit"`
					TotalCredit int    `json:"totalCredit"`
				} `json:"bankAccountReports"`
			} `json:"pageData"`
			AppData struct {
				FacebookPageURL       string `json:"facebookPageUrl"`
				SocialNetworkShareURL string `json:"socialNetworkShareUrl"`
				GoogleFormsURL        string `json:"googleFormsUrl"`
				Hotline               string `json:"hotline"`
				Email                 string `json:"email"`
				AllowRegisterKols     bool   `json:"allowRegisterKols"`
				FundraisingStepsURL   string `json:"fundraisingStepsUrl"`
				DeepLinkMiniapp       string `json:"deepLinkMiniapp"`
				AllowDeepLink         bool   `json:"allowDeepLink"`
				LinkDownLoadApp       string `json:"linkDownLoadApp"`
			} `json:"appData"`
		} `json:"pageProps"`
	} `json:"props"`
	Page  string `json:"page"`
	Query struct {
		ID string `json:"id"`
	} `json:"query"`
	BuildID       string `json:"buildId"`
	RuntimeConfig struct {
		AppStore struct {
			URL      string `json:"url"`
			AppID    string `json:"appId"`
			TeamID   string `json:"teamId"`
			BundleID string `json:"bundleId"`
		} `json:"appStore"`
		PlayStore struct {
			URL                string `json:"url"`
			PackageName        string `json:"packageName"`
			Sha256Fingerprints string `json:"sha256Fingerprints"`
		} `json:"playStore"`
		PageLink        string `json:"pageLink"`
		CanonicalDomain string `json:"canonicalDomain"`
		GtmID           string `json:"gtmId"`
		APIBaseURL      string `json:"apiBaseUrl"`
		UseHsts         bool   `json:"useHsts"`
		UseCsp          bool   `json:"useCsp"`
		APIPublicURL    string `json:"apiPublicUrl"`
		AppIDFacebook   string `json:"appIdFacebook"`
		PageIDFacebook  string `json:"pageIdFacebook"`
		Maps            struct {
			APIURL      string  `json:"apiUrl"`
			APIKey      string  `json:"apiKey"`
			AccessToken string  `json:"accessToken"`
			APIVersion  float64 `json:"apiVersion"`
		} `json:"maps"`
	} `json:"runtimeConfig"`
	IsFallback   bool  `json:"isFallback"`
	CustomServer bool  `json:"customServer"`
	Gip          bool  `json:"gip"`
	AppGip       bool  `json:"appGip"`
	ScriptLoader []any `json:"scriptLoader"`
}

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

var lastItemTime time.Time = time.Time{}

func main() {
	currencyPrinter := message.NewPrinter(language.English)
	// config
	configFile := flag.String("config", "config.json", "config file")
	flag.Parse()

	var c config
	if err := cleanenv.ReadConfig(*configFile, &c); err != nil {
		log.Panic(err)
	}

	if err := cleanenv.ReadEnv(&c); err != nil {
		log.Panic(err)
	}

	fmt.Println(c.AdminID)
	// bot API
	bot, err := tgbotapi.NewBotAPI(c.BotToken)
	if err != nil {
		log.Panic(err)
	}
	bot.Debug = true
	log.Printf("Authorized on account %s", bot.Self.UserName)
	// crawler
	crawler := colly.NewCollector()
	crawler.UserAgent = "Mozilla/5.0 (X11; Linux x86_64; rv:109.0) Gecko/20100101 Firefox/116.0"
	data, err := crawlData("sanchoicauvong", crawler)
	if err != nil {
		bot.Send(tgbotapi.NewMessage(int64(c.AdminID), fmt.Sprintf("Error: %s", err)))
	} else {
		// type = CREDIT
		count := 0
		lastItemClone := lastItemTime
		lastItemTime = time.Now()
		for _, v := range data.Props.PageProps.PageData.BankAccountTransactions.Groups {
			for _, v2 := range v.Transactions {
				if v2.Type == "CREDIT" {
					fmt.Println(v2.TransactionTime)
					fmt.Println(v2.TransactionAmount)
					fmt.Println(v2.FeeAmount)
					fmt.Println(v2.Narrative)
					fmt.Println(v2.UserName)
					fmt.Println("=====================================")
					// 2024-01-28T21:11:00
					TransactionTime, err := time.Parse("2006-01-02T15:04:05", v2.TransactionTime)
					if err != nil {
						bot.Send(tgbotapi.NewMessage(int64(c.AdminID), fmt.Sprintf("Error: %s", err)))
					} else {
						if TransactionTime.After(lastItemClone) {
							amount := currencyPrinter.Sprintf("%d", v2.TransactionAmount)
							bot.Send(tgbotapi.NewMessage(int64(c.GroupChatID),
								fmt.Sprintf(`Thời gian: %s
							Số tiền: %s
							Phí: %d
							Nội dung: %s
							Người thụ hưởng: %s`, TransactionTime.Local().Format(time.DateTime), amount, v2.FeeAmount, v2.Narrative, v2.UserName)))
						}
					}
					count++
				}
			}

		}
		fmt.Println(count)

	}

}
