package main

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/subeenregmi/spotify-waybar-module/pkg/waybar"
)

func main() {
	m := waybar.Module{
		Text: "Hello World!",
	}

	for count := range 100 {
		c := m
		c.Text = fmt.Sprintf("%s %d", m.Text, count)

		bytes, err := json.Marshal(c)
		if err != nil {
			panic(err)
		}

		fmt.Println(string(bytes))
		time.Sleep(time.Second * 3)
	}
}
