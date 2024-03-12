package main

import (
	"github.com/gocolly/colly/v2"
	"log/slog"
)

func GetWorkshopItemTitle(itemId string) (title string) {
	title = itemId

	c := colly.NewCollector()

	c.OnHTML("div.workshopItemTitle", func(element *colly.HTMLElement) {
		title = element.Text
	})

	c.OnError(func(response *colly.Response, err error) {
		slog.Error("sth happened trying to request", "error", err)
	})

	err := c.Visit("https://steamcommunity.com/sharedfiles/filedetails/?id=" + itemId)
	if err != nil {
		slog.Error("sth happened trying to visit", "error", err)
	}

	return title
}
