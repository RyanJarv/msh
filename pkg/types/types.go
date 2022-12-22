package types

import "io"

type Process interface {
	io.ReaderFrom
	io.WriterTo
	Run() error
}

type Step interface {
	io.ReaderFrom
	io.WriterTo
	Process
	Process(Process)
}

type Deployable interface {
	Deploy() error
}
