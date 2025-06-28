package tests

import (
	"bytes"
	"os"
	"reflect"
)

func ExecCliFunction(execFunc interface{}, args ...interface{}) (string, error) {
	var buf bytes.Buffer
	oldStdout := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	reflect.ValueOf(execFunc).Call(func() []reflect.Value {
		vals := make([]reflect.Value, len(args))
		for i, a := range args {
			vals[i] = reflect.ValueOf(a)
		}
		return vals
	}())

	err := w.Close()
	if err != nil {
		return "", err
	}
	os.Stdout = oldStdout
	_, err = buf.ReadFrom(r)
	if err != nil {
		return "", err
	}
	err = r.Close()
	if err != nil {
		return "", err
	}
	output := buf.String()

	return output, nil
}
