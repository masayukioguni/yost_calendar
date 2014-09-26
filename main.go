package main

import (
	"github.com/codegangsta/cli"
	"os"
	"strings"
)

func yost_to_westean(text string) string {
	calendarText := ""
	yost_table := map[string]string{
		"a": "10",
		"b": "11",
		"c": "12",
		"d": "13",
		"e": "14",
		"f": "15",
		"g": "16",
		"h": "17",
		"i": "18",
		"j": "19",
		"k": "20",
		"l": "21",
		"m": "22",
		"n": "23",
		"o": "24",
		"p": "25",
		"q": "26",
		"r": "27",
		"s": "28",
		"t": "29",
		"u": "30",
		"v": "31",
		"w": "32",
		"x": "33",
		"y": "34",
		"z": "35"}
	for _, r := range text {
		calendarText += yost_table[string(r)]
	}
	return calendarText
}

func main() {
	app := cli.NewApp()
	app.Name = "yost_carendar"
	app.Usage = "yost calendar conversion tool"

	app.Action = func(c *cli.Context) {
		text := c.Args()[0]
		println(yost_to_westean(strings.ToLower(text)))
	}

	app.Run(os.Args)
}
