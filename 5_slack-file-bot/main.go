package main

import (
	"fmt"
	"os"

	"github.com/slack-go/slack"
)

func main() {

	os.Setenv("SlACK_BOT_TOKEN", "<redacted>")
	os.Setenv("CHANNEL_ID", "<redacted>")

	api := slack.New(os.Getenv("SlACK_BOT_TOKEN"))
	ChannelArr := []string{os.Getenv("CHANNEL_ID")}
	fileArr := []string{"go.mod"}

	for i := 0; i < len(fileArr); i++ {
		params := slack.FileUploadParameters{
			Channels: ChannelArr,
			File:     fileArr[i],
		}
		file, err := api.UploadFile(params)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Printf("Name %s", file.Name)
	}
}
