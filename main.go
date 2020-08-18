package main

import (
	"math/rand"
	"strings"
	"time"

	ubot "github.com/UBotPlatform/UBot.Common.Go"
)

var api *ubot.AppApi

func onReceiveChatMessage(bot string, msgType ubot.MsgType, source string, sender string, message string, info ubot.MsgInfo) (ubot.EventResultType, error) {
	trimmedMessage := strings.TrimSpace(message)
	if trimmedMessage == "诗歌" || trimmedMessage == "作诗" {
		rand.Seed(time.Now().UnixNano())
		_ = api.SendChatMessage(bot, msgType, source, sender, makePoem())
		return ubot.CompleteEvent, nil
	}
	return ubot.IgnoreEvent, nil
}

func makePoem() string {
	var builder strings.Builder
	n := 6 + rand.Intn(21)
	for i := 0; i < n; i++ {
		if i != 0 {
			builder.WriteString("\n")
		}
		sentence := sentences[rand.Intn(len(sentences))]
		sentence = fillWithFragments(sentence, "DD", fragmentDD)
		sentence = fillWithFragments(sentence, "DJ", fragmentDJ)
		sentence = fillWithFragments(sentence, "MM", fragmentMM)
		sentence = fillWithFragments(sentence, "TT", fragmentTT)
		sentence = fillWithFragments(sentence, "XX", fragmentXX)
		builder.WriteString(sentence)
	}
	return builder.String()
}

func fillWithFragments(template string, key string, fragments []string) string {
	result := template
	for index := strings.Index(result, key); index != -1; index = strings.Index(result, key) {
		fragment := fragments[rand.Intn(len(fragments))]
		result = result[:index] + fragment + result[index+len(key):]
	}
	return result
}

func main() {
	err := ubot.HostApp("Poet", func(e *ubot.AppApi) *ubot.App {
		api = e
		return &ubot.App{
			OnReceiveChatMessage: onReceiveChatMessage,
		}
	})
	ubot.AssertNoError(err)
}
