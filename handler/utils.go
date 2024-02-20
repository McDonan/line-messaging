package handler

import (
	"github.com/line/line-bot-sdk-go/v8/linebot/messaging_api"
	"log"
)

func initBot() (string, *messaging_api.MessagingApiAPI) {
	channelSecret := "4bb2b4063637ee24567232d80b3dc674"
	channelToken := "jraOO5bHcqPrgHT41Z+9geqP8kM5DCSGspBIOdKGQ5YiQaxU4Y8SDleFljrYuOvKVY6gA4sYb2/BwkHfajHVNM93KDXdXPw3bJA09Tndu1NSnXa9t+vv3+3ibA/rrfbXRyVas6v2/YnNq3IdGk5bpwdB04t89/1O/w1cDnyilFU="
	bot, err := messaging_api.NewMessagingApiAPI(channelToken)
	if err != nil {
		log.Fatal(err)
	}

	return channelSecret, bot
}

var commonSubjects = map[string]string{
	"soc":    "สังคมศึกษา",
	"thai":   "ภาษาไทย",
	"eng":    "ภาษาต่างประเทศ",
	"art":    "ศิลปะ",
	"health": "สุขศึกษา",
	"other":  "วิชาอื่นๆ",
}

var primarySubject = map[string]string{
	"sci": "วิทยาศาสตร์",
}

var highSchoolSubjects = map[string]string{
	"phys": "ฟิสิกส์",
	"chem": "เคมี",
	"bio":  "ชีววิทยา",
	"tech": "เทคโนโลยี",
}

func getQuickReplies(isPrimaryStudent bool) []messaging_api.QuickReplyItem {
	res := []messaging_api.QuickReplyItem{{
		Action: messaging_api.MessageAction{
			Label: "math",
			Text:  "คณิตศาสตร์",
		},
	}}
	for key, value := range commonSubjects {
		res = append(res, messaging_api.QuickReplyItem{
			Action: messaging_api.MessageAction{
				Label: key,
				Text:  value,
			},
		})
	}

	extraSubjects := highSchoolSubjects
	if isPrimaryStudent {
		extraSubjects = primarySubject
	}

	for key, value := range extraSubjects {
		commonSubjects[key] = value
	}

	for key, value := range commonSubjects {
		res = append(res, messaging_api.QuickReplyItem{
			Action: messaging_api.UriAction{
				Label: value,
				Uri:   "https://85b2-115-87-239-131.ngrok-free.app/subject/" + key,
			},
		})
	}

	return res
}
