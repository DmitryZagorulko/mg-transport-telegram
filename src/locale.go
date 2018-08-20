package main

import (
	"io/ioutil"

	"github.com/nicksnyder/go-i18n/v2/i18n"
	"golang.org/x/text/language"
	"gopkg.in/yaml.v2"
)

var (
	localizer *i18n.Localizer
	bundle    = &i18n.Bundle{DefaultLanguage: language.English}
	matcher   = language.NewMatcher([]language.Tag{
		language.English,
		language.Russian,
		language.Spanish,
	})
)

func loadTranslateFile() {
	bundle.RegisterUnmarshalFunc("yml", yaml.Unmarshal)
	files, err := ioutil.ReadDir("translate")
	if err != nil {
		panic(err)
	}
	for _, f := range files {
		if !f.IsDir() {
			bundle.MustLoadMessageFile("translate/" + f.Name())
		}
	}
}

func setLocale(al string) {
	tag, _ := language.MatchStrings(matcher, al)
	localizer = i18n.NewLocalizer(bundle, tag.String())
}

func getLocalizedMessage(messageID string) string {
	return localizer.MustLocalize(&i18n.LocalizeConfig{MessageID: messageID})
}

func getLocale() map[string]string {
	return map[string]string{
		"ButtonSave":  getLocalizedMessage("button_save"),
		"ApiKey":      getLocalizedMessage("api_key"),
		"TabSettings": getLocalizedMessage("tab_settings"),
		"TabBots":     getLocalizedMessage("tab_bots"),
		"TableName":   getLocalizedMessage("table_name"),
		"TableToken":  getLocalizedMessage("table_token"),
		"AddBot":      getLocalizedMessage("add_bot"),
		"TableDelete": getLocalizedMessage("table_delete"),
		"Title":       getLocalizedMessage("title"),
	}
}