package main

import (
	"fmt"
	"log"
	"time"

	"github.com/ilyakaznacheev/cleanenv"
)

type config struct {
	AdminID     int64  `json:"admin_id" env:"ADMIN_ID" env-default:""`
	UserID      string `json:"user_id" env:"USER_ID" env-default:""`
	GroupChatID int64  `json:"group_chat_id" env:"GROUPCHAT_ID" env-default:""`
	BotToken    string `json:"bot_token" env:"BOT_TOKEN" env-default:""`
	Interval    int    `json:"interval" env:"INTERVAL" env-default:"5"`
	Debug       bool   `json:"debug" env:"DEBUG" env-default:"false"`
}

func (c *config) Load() error {
	if err := cleanenv.ReadEnv(c); err != nil {
		return fmt.Errorf("config load error: %s", err)
	}

	if c.Debug {
		lastItemTime = time.Time{}
		log.Printf("Config: %+v", c)
	}

	if c.UserID == "" {
		return fmt.Errorf("missing required fields: USER_ID")
	}

	if c.BotToken == "" {
		return fmt.Errorf("missing required fields: BOT_TOKEN")
	}

	if c.GroupChatID == 0 {
		return fmt.Errorf("missing required fields: GROUPCHAT_ID")
	}

	if c.AdminID == 0 {
		return fmt.Errorf("missing required fields: ADMIN_ID")
	}

	return nil
}

type Transaction struct {
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

	transactionTime time.Time
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
						Date         string        `json:"date"`
						Count        int           `json:"count"`
						Transactions []Transaction `json:"transactions"`
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
