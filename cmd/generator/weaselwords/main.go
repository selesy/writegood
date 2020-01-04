package main

import (
	"github.com/selesy/writegood/pkg/generator"
	log "github.com/sirupsen/logrus"
)

func main() {
	err := generator.GenerateWeaselWords()
	if err != nil {
		log.Fatal(err)
	}
}
