package main

import (
	"fmt"
	"github.com/line/line-bot-sdk-go/v8/linebot/messaging_api"
	"log"
	"net/http"
	"time"
)

func check(w http.ResponseWriter, req *http.Request) {
	log.Println("/check called...")

	_, bot := initBot()

	userId := "Ud6be7b58a758c9b5ce1518f7c73224a2"
	now := time.Now().UnixNano()
	b := now%2 == 0
	isPremiumUser := b
	if !isPremiumUser {
		if _, err := bot.PushMessage(&messaging_api.PushMessageRequest{
			To: userId,
			Messages: []messaging_api.MessageInterface{messaging_api.TextMessage{
				Text: "สมัครบริการแบบพรีเมียม เพื่อให้น้องมานะ เอไอช่วยสอนการบ้านได้ง่ายๆ \" แตะที่นี่เพื่อสมัคร \"",
			}},
		}, ""); err != nil {
			w.WriteHeader(500)
		}
	}

	user, err := bot.GetProfile(userId)
	if err != nil {
		w.WriteHeader(500)
	}

	primaryStudent := b
	if _, err := bot.PushMessage(&messaging_api.PushMessageRequest{
		To: userId,
		Messages: []messaging_api.MessageInterface{
			messaging_api.TextMessage{
				Text: fmt.Sprintf("น้องมานะ เอไอพร้อมแล้วที่จะช่วยแนะนำการบ้าน ให้กับน้อง%sครับ", user.DisplayName),
			},
			messaging_api.TextMessage{
				Text:       "เลือกวิชาที่ต้องการให้แนะนำตามด้านล่างได้เลยครับ",
				QuickReply: &messaging_api.QuickReply{Items: getQuickReplies(primaryStudent)},
			},
		},
	}, ""); err != nil {
		w.WriteHeader(500)
	}
}
