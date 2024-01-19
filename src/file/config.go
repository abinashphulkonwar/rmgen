package file

import (
	"encoding/json"
	"os"

	applicationerror "github.com/abinashphulkonwar/go-random-data-generation/src/application_error"
)

type Field struct {
	Name string `json:"field_name"`
	Type string `json:"type"`
}

type Config struct {
	path   string
	f      *os.File
	Fields []Field
}

func NewConfig(path string) *Config {
	return &Config{
		path: path,
	}

}

func (c *Config) Init() {
	connection, err := os.OpenFile(c.path, os.O_RDWR|os.O_CREATE, 0666)
	applicationerror.ErrorChecker(&err)
	c.f = connection
}
func (c *Config) Parse() {
	buf := make([]byte, 1024)
	index, err := c.f.Read(buf)
	applicationerror.ErrorChecker(&err)
	println(string(buf))
	var temp []Field
	err = json.Unmarshal(buf[0:index], &temp)
	applicationerror.ErrorChecker(&err)

	for _, v := range temp {
		println(v.Name)
		println(v.Type)
		println("==================")
		c.Fields = append(c.Fields, v)
	}

}

func (c *Config) Close() {
	c.f.Close()
}
