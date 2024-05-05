package tampering

import (
	"time"

	"github.com/liamg/grace/filter"
	"github.com/liamg/grace/tracer"
)

var tamper *filter.Filter

func Parse(input string) (*filter.Filter, error) {
	if input == "" {
		return nil, nil
	}
	f, err := filter.Parse(input)
	if err != nil {
		return nil, err
	}
	tamper = f
	return tamper, nil
}

func Delay(syscall *tracer.Syscall) {
	if tamper == nil {
		return
	}

	if tamper.Match(syscall, true) {
		time.Sleep(time.Second)
	}
}
