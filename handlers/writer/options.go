package writer

import (
	"io"
	"io/ioutil"
	"github.com/Tlantic/meerkats"
)

var (
	DiscardOutput = Output(ioutil.Discard)
)

func Output(w io.Writer) meerkats.HandlerOption {
	return meerkats.HandlerOption(func(handler meerkats.Handler) {
		handler.(*writerLogger).w = w
	})
}

