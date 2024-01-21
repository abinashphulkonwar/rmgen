package file_test

import (
	"bytes"
	"encoding/gob"
	"testing"

	applicationerror "github.com/abinashphulkonwar/rmgen/src/application_error"
	"github.com/abinashphulkonwar/rmgen/src/file"
	"github.com/brianvoe/gofakeit/v6"
)

func TestOutput(t *testing.T) {
	output := file.NewOutput()
	output.Init()
	buf := new(bytes.Buffer)
	enc := gob.NewEncoder(buf)
	if err := enc.Encode(*gofakeit.Address()); err != nil {
		applicationerror.ErrorChecker(&err)
	}
	data := buf.Bytes()
	output.Write(&data)
	defer output.Close()

}
