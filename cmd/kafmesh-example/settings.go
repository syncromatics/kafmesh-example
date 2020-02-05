package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/syncromatics/go-kit/database"
)

type settings struct {
	Registry         string
	Brokers          []string
	DatabaseSettings *database.PostgresDatabaseSettings
}

func getSettingsFromEnv() (*settings, error) {
	errors := []string{}

	ok := false

	registry, ok := os.LookupEnv("REGISTRY")
	if !ok {
		errors = append(errors, "REGISTRY")
	}

	brokers, ok := os.LookupEnv("BROKERS")
	if !ok {
		errors = append(errors, "BROKERS")
	}

	ds := database.PostgresDatabaseSettings{}

	ds.Host, ok = os.LookupEnv("DATABASE_HOST")
	if !ok {
		errors = append(errors, "DATABASE_HOST")
	}

	ds.Name, ok = os.LookupEnv("DATABASE_NAME")
	if !ok {
		errors = append(errors, "DATABASE_NAME")
	}

	ds.User, ok = os.LookupEnv("DATABASE_USER")
	if !ok {
		errors = append(errors, "DATABASE_USER")
	}

	ds.Password, ok = os.LookupEnv("DATABASE_PASSWORD")
	if !ok {
		errors = append(errors, "DATABASE_PASSWORD")
	}

	if len(errors) > 0 {
		return nil, fmt.Errorf("Missing required environment variables: %s", strings.Join(errors, ", "))
	}

	return &settings{
		Registry:         registry,
		Brokers:          strings.Split(brokers, ","),
		DatabaseSettings: &ds,
	}, nil
}
