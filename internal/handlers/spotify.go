package handlers

import (
	"fmt"
	"html"
	"log/slog"

	"github.com/subeenregmi/music-waybar-module/pkg/playerctl"
	"github.com/subeenregmi/music-waybar-module/pkg/waybar"
)

func SpotifyHandler(logger *slog.Logger) {
	status, err := playerctl.Status(playerctl.SPOTIFY)
	if err != nil {
		output := waybar.ModuleOutput{
			Text:  "",
			Class: "unknown",
		}

		output.Print(logger)

		return
	}

	metadata, err := playerctl.GetMetadata(playerctl.SPOTIFY)
	if err != nil {
		logger.Error("failed to execute metadata fetch", slog.Any("error", err))
		output := waybar.ModuleOutput{
			Text:  "",
			Class: "unknown",
		}

		output.Print(logger)

		return
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

	if len(text) >= waybar.MAX_LENGTH {
		text = metadata.Title
	}

	text = html.EscapeString(text)

	output.Text = fmt.Sprintf(" %v", text)

	if status == playerctl.STATUS_PAUSED {
		output.Class = "paused"
	}

	output.Print(logger)
}
