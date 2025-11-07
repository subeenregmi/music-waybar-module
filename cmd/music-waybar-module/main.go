package main

import (
	"encoding/json"
	"fmt"
	"log/slog"
	"time"

	"github.com/subeenregmi/music-waybar-module/pkg/playerctl"
	"github.com/subeenregmi/music-waybar-module/pkg/waybar"
)

const (
	MAX_ATTEMPTS = 5
	MAX_LENGTH   = 35
)

func main() {
	logger := slog.Default()

	failedAttempts := 0
	for failedAttempts <= MAX_ATTEMPTS {

		_, err := playerctl.Status()
		if err != nil {
			logger.Error("failed to get playertctl status", slog.Any("error", err))
			failedAttempts++
			continue
		}

		metadata, err := playerctl.GetMetadata()
		if err != nil {
			logger.Error("failed to execute metadata fetch", slog.Any("error", err))
			failedAttempts++
			continue
		}

		var text string
		if metadata.Title == "Advertisement" {
			text = "advert"
		} else {
			text = fmt.Sprintf("%v - %v", metadata.Artist, metadata.Title)
		}

		if len(text) >= MAX_LENGTH {
			text = metadata.Title
		}

		output := waybar.ModuleOutput{
			Text: text,
		}

		bOutput, err := json.Marshal(output)
		if err != nil {
			logger.Error("failed to marshal output", slog.Any("error", err))
			failedAttempts++
			continue
		}

		fmt.Println(string(bOutput))

		time.Sleep(time.Second)
	}
}
