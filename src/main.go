package main

import (
	"os"
	"regexp"

	"github.com/jessevdk/go-flags"

	"github.com/op/go-logging"
)

// Options struct
type Options struct {
	Config string `short:"c" long:"config" default:"config.yml" description:"Path to configuration file"`
}

const Type = "telegram"
const MaxCharsCount uint16 = 4096

var (
	config   *TransportConfig
	orm      *Orm
	logger   *logging.Logger
	options  Options
	parser   = flags.NewParser(&options, flags.Default)
	rx       = regexp.MustCompile(`/+$`)
	currency = map[string]string{
		"rub": "₽",
		"uah": "₴",
		"byr": "Br",
		"kzt": "₸",
		"usd": "$",
		"eur": "€",
	}
)

func main() {
	if _, err := parser.Parse(); err != nil {
		if flagsErr, ok := err.(*flags.Error); ok && flagsErr.Type == flags.ErrHelp {
			os.Exit(0)
		} else {
			os.Exit(1)
		}
	}
}
