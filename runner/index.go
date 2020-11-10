package runner

import (
	"errors"
	"io"
)

var inpError = errors.New("any_error_msg_here")
type storage interface {
	Get(string) (string, error)
	Set(string, string) error
}

type runner struct {
	database storage
}

func NewRunner (db storage) runner {
	return runner{db}
}

func (r runner) Run (output io.StringWriter, args [] string) error {
	if len(args)<3{
		return inpError
	}
	switch args[1] {
	case "set":
		if len(args) < 4 {
			return inpError
		}
		if err := r.database.Set(args[2], args[3] + "\n"); err != nil {
			return err
		}
	case "get":
		v, err := r.database.Get(args[2])
		if err != nil {
			return err
		}
		output.WriteString(v + "\n")
	default:
		return inpError
	}

	return nil
}