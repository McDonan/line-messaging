package main

import (
	"errors"
	"github.com/line/line-bot-sdk-go/v8/linebot"
	"github.com/line/line-bot-sdk-go/v8/linebot/messaging_api"
	"github.com/line/line-bot-sdk-go/v8/linebot/webhook"
	"log"
	"net/http"
)

func callback(w http.ResponseWriter, req *http.Request) {
	log.Println("/callback called...")

	channelSecret, bot := initBot()

	cb, err := webhook.ParseRequest(channelSecret, req)
	if err != nil {
		log.Printf("Cannot parse request: %+v\n", err)
		if errors.Is(err, linebot.ErrInvalidSignature) {
			w.WriteHeader(400)
		} else {
			w.WriteHeader(500)
		}
		return
	}

	log.Println("Handling events...")
	for _, event := range cb.Events {
		log.Printf("/callback called%+v...\n", event)

		switch e := event.(type) {
		case webhook.MessageEvent:
			switch message := e.Message.(type) {
			case webhook.TextMessageContent:
				if _, err = bot.ReplyMessage(
					&messaging_api.ReplyMessageRequest{
						ReplyToken: e.ReplyToken,
						Messages: []messaging_api.MessageInterface{
							messaging_api.TextMessage{
								Text: message.Text,
							},
						},
					},
				); err != nil {
					log.Print(err)
				} else {
					log.Println("Sent text reply.")
				}
			case webhook.ImageMessageContent:

				//case webhook.StickerMessageContent:
				//	replyMessage := fmt.Sprintf(
				//		"sticker id is %s, stickerResourceType is %s", message.StickerId, message.StickerResourceType)
				//	if _, err = bot.ReplyMessage(
				//		&messaging_api.ReplyMessageRequest{
				//			ReplyToken: e.ReplyToken,
				//			Messages: []messaging_api.MessageInterface{
				//				messaging_api.TextMessage{
				//					Text: replyMessage,
				//				},
				//			},
				//		}); err != nil {
				//		log.Print(err)
				//	} else {
				//		log.Println("Sent sticker reply.")
				//	}
				//default:
				//	log.Printf("Unsupported message content: %T\n", e.Message)
			}
		default:
			log.Printf("Unsupported message: %T\n", event)
		}
	}
}
