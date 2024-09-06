package logger

import (
	"log/slog"
	"os"
)

func Config() (*slog.Logger, error) {
	fi, err := os.OpenFile("logs.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		return nil, err
	}

	handler := slog.NewJSONHandler(fi, nil)
	return slog.New(handler), nil
}
