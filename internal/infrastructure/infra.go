package infrastructure

import (
	"github.com/zippunov/alien-invasion/internal/usecases"
	"io"
	"os"
)

// Compile check to verify Interface Compliance. See
// https://github.com/uber-go/guide/blob/master/style.md#verify-interface-compliance
var _ usecases.IInfra = (*Infra)(nil)

// Infra holds references to all external sources required by the application and
// required configuration parameters
type Infra struct {
	reader      io.ReadCloser
	aliensCount int
	writer      io.WriteCloser
	log         func(format string, a ...any)
}

// Shutdown does clean up at the end of application work.
func (i *Infra) Shutdown() {
	_ = i.reader.Close()
	_ = i.writer.Close()
}

// In is a part of usecases.IInfra interface implementation
func (i *Infra) In() io.Reader {
	return i.reader
}

// Out is a part of usecases.IInfra interface implementation
func (i *Infra) Out() io.Writer {
	return i.writer
}

// AliensCount is a part of usecases.IInfra interface implementation
func (i *Infra) AliensCount() int {
	return i.aliensCount
}

// Log is a part of usecases.IInfra interface implementation
func (i *Infra) Log() func(format string, a ...any) {
	return i.log
}

// InitInfra does initialization of the all application external resources
// according to given configuration
func InitInfra(config Config) (Infra, error) {
	inFile, err := os.Open(config.mapFilePath)
	if err != nil {
		return Infra{}, err
	}
	var outFile *os.File
	if len(config.outFilePath) != 0 {
		if outFile, err = os.Create(config.outFilePath); err != nil {
			return Infra{}, err
		}
	}
	if outFile == nil {
		outFile = os.Stdout
	}
	return Infra{
		reader:      inFile,
		aliensCount: config.aliensCount,
		writer:      outFile,
		log:         config.log,
	}, nil
}
