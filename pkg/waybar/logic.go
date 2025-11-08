package waybar

import (
	"encoding/json"
	"fmt"
	"log/slog"
)

func (m *ModuleOutput) Print(logger *slog.Logger) {
	output, err := json.Marshal(&m)
	if err != nil {
		logger.Error("failed to marshal output", slog.Any("error", err))
		return
	}

	fmt.Println(string(output))
}
