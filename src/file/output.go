package file

import (
	"crypto/md5"
	"encoding/hex"
	"os"
	"time"

	applicationerror "github.com/abinashphulkonwar/go-random-data-generation/src/application_error"
	"github.com/google/uuid"
)

type Output struct {
	f *os.File
}

var OutputPointer *Output = &Output{}

func NewOutput() *Output {
	pointer := &Output{}
	OutputPointer = pointer
	return pointer
}

func (o *Output) Open() {

	connection, err := os.Create(o.file_name())
	applicationerror.ErrorChecker(&err)
	o.f = connection
}

func (o *Output) Write(data *[]byte) int {
	i, err := o.f.Write(*data)
	applicationerror.ErrorChecker(&err)
	return i
}

func (o *Output) file_name() string {
	id := uuid.New().String()
	if id == "" {
		hasher := md5.New()
		time_now := time.Now().String()
		hasher.Write([]byte(time_now))
		id = hex.EncodeToString(hasher.Sum(nil))
	}
	return id[0:8] + ".txt"
}

func (o *Output) Close() {
	o.f.Close()
}
