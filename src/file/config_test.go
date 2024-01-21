package file_test

import (
	"testing"

	"github.com/abinashphulkonwar/rmgen/src/file"
)

func TestCongig(t *testing.T) {
	config := file.NewConfig("config.json")
	config.Init()
	config.Parse()
	defer config.Close()
}
