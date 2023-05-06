package config

import (
	"context"

	"github.com/joho/godotenv"
)

func Init(_ context.Context) error {
	err := godotenv.Load(".env")
	if err != nil {
		return err
	}

	return nil
}
