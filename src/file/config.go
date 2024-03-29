package file

import (
	"encoding/json"
	"os"

	applicationerror "github.com/abinashphulkonwar/rmgen/src/application_error"
)

const (
	EMAIL_TYPE   = "email"
	NAME_TYPE    = "name"
	PHONE_TYPE   = "phone"
	ADDRESS_TYPE = "address"
	DATE_TYPE    = "date"
	NUMBER_TYPE  = "number"
	DEFUALT      = "text"
)

type Field struct {
	Name string `json:"field_name"`
	Type string `json:"type"`
}

type Config struct {
	path   string
	f      *os.File
	fields []Field
}

var ConfigPointer *Config = nil

func NewConfig(path string) *Config {
	pointer := &Config{
		path: path,
	}

	ConfigPointer = pointer
	return pointer
}

func (c *Config) Init() {
	connection, err := os.OpenFile(c.path, os.O_RDONLY, 0666)
	if os.IsExist(err) {
		println("file not existed")
	}
	applicationerror.ErrorChecker(&err)
	c.f = connection
}
func (c *Config) Parse() *[]Field {
	buf := make([]byte, 1024)
	index, err := c.f.Read(buf)
	applicationerror.ErrorChecker(&err)
	var temp []Field
	err = json.Unmarshal(buf[0:index], &temp)
	applicationerror.ErrorChecker(&err)
	println("==================")
	for _, v := range temp {
		println("field name: ", v.Name, " type: ", v.Type)
		c.fields = append(c.fields, v)
	}
	println("==================")

	return &c.fields

}

func (c *Config) Get() *[]Field {
	if len(c.fields) == 0 {
		c.Parse()
	}
	return &c.fields
}

func (c *Config) Close() {
	c.f.Close()
}
