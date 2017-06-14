package main

import (
	"os"

	"github.com/elastic/beats/libbeat/beat"

	"github.com/stanxii/cockroachbeat/beater"
)

func main() {
	err := beat.Run("cockroachbeat", "", beater.New)
	if err != nil {
		os.Exit(1)
	}
}
