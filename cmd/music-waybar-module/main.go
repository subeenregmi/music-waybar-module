package main

import (
	"fmt"
	"html"
	"log/slog"
	"time"

	"github.com/subeenregmi/music-waybar-module/pkg/playerctl"
	"github.com/subeenregmi/music-waybar-module/pkg/waybar"
)

const (
	MAX_ATTEMPTS = 5
	MAX_LENGTH   = 25
)

func main() {
	logger := slog.Default()

	for {
		time.Sleep(time.Second)

		status, err := playerctl.Status()
		if err != nil {
			// logger.Error("failed to get playertctl status", slog.Any("error", err))
			output := waybar.ModuleOutput{
				Text:  "",
				Class: "unknown",
			}

			output.Print(logger)

			continue
		}

		metadata, err := playerctl.GetMetadata()
		if err != nil {
			logger.Error("failed to execute metadata fetch", slog.Any("error", err))
			output := waybar.ModuleOutput{
				Text:  "",
				Class: "unknown",
			}

			output.Print(logger)

			continue
		}

		output := waybar.ModuleOutput{}

		var text string
		if metadata.Title == "Advertisement" {
			text = "advert"
			output.Class = "advert"
		} else {
			text = fmt.Sprintf("%v - %v", metadata.Artist, metadata.Title)
			output.Class = "normal"
		}

		if len(text) >= MAX_LENGTH {
			text = metadata.Title
		}

		text = html.EscapeString(text)

		output.Text = fmt.Sprintf(" %v", text)

		if status == playerctl.STATUS_PAUSED {
			output.Class = "paused"
		}

		output.Print(logger)
	}
}
