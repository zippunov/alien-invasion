package infrastructure

import (
	"errors"
	"flag"
)

// Config is holder for the settings given with application parpams and provides default values
type Config struct {
	mapFilePath string
	aliensCount int
	outFilePath string
	log         func(format string, a ...any)
	Help        bool
}

// InitConfig validates application params and creates new Config instance
func InitConfig(log func(format string, a ...any)) (Config, error) {
	var (
		mapFile     string
		aliensCount uint
		outFile     string
		help        bool
	)
	flag.StringVar(&mapFile, "f", "", "")
	flag.UintVar(&aliensCount, "n", 0, "")
	flag.StringVar(&outFile, "o", "", "")
	flag.BoolVar(&help, "h", false, "")
	flag.Parse()

	config := Config{
		mapFilePath: mapFile,
		aliensCount: int(aliensCount),
		outFilePath: outFile,
		log:         log,
		Help:        help,
	}

	if !config.Help {
		if len(mapFile) == 0 {
			return Config{}, errors.New("missing map file path")
		}
		if aliensCount == 0 {
			return Config{}, errors.New("aliens number must be greater than 0")
		}
	}

	return config, nil
}
