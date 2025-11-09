package main

import (
	"fmt"
	"log/slog"
	"strings"
	"time"

	"github.com/subeenregmi/music-waybar-module/internal/handlers"
	"github.com/subeenregmi/music-waybar-module/pkg/playerctl"
)

const (
	MAX_ATTEMPTS = 5
	MAX_LENGTH   = 25
)

func main() {
	logger := slog.Default()

	var lastPlayer string

	for {
		time.Sleep(time.Second)

		players, err := playerctl.Players()
		if err != nil {
			continue
		}

		logger.Info("players", slog.Any("players", players))

		for _, player := range players {
			status, err := playerctl.Status(player)
			if err != nil {
				logger.Error(
					"failed to get player status",
					slog.Any("player", player),
					slog.Any("error", err),
				)
				continue
			}

			if status == playerctl.STATUS_PAUSED {
				fmt.Println("last player", lastPlayer)
			}

			if status == playerctl.STATUS_PLAYING &&
				player == playerctl.SPOTIFY {
				handlers.SpotifyHandler(logger)
				lastPlayer = playerctl.SPOTIFY
				continue
			}

			if status == playerctl.STATUS_PLAYING &&
				strings.Contains(player, "firefox") {
				fmt.Println("firefox")
				handlers.FirefoxHandler(logger)
				lastPlayer = player
				continue
			}
		}
	}
}
