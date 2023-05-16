package richpresence

import (
	"time"

	"github.com/hugolgst/rich-go/client"
)

func SetupRichPresence() error {
	err := client.Login("1080214625808822423")
	if err != nil {
		return err
	}

	start := time.Now()
	err = client.SetActivity(client.Activity{
		State:      "mixhub.fun",
		Details:    "Проходит проверку на читы",
		LargeImage: "mixhub",
		LargeText:  "Я на проверке :D",
		Timestamps: &client.Timestamps{
			Start: &start,
		},
		Buttons: []*client.Button{
			{
				Label: "shop.mixhub.fun",
				Url:   "https://shop.mixhub.fun",
			},
		},
	})

	return err
}
