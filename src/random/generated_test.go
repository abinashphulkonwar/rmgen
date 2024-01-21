package random_test

import (
	"os"
	"testing"

	"github.com/abinashphulkonwar/rmgen/src/file"
	"github.com/abinashphulkonwar/rmgen/src/random"
)

func TestRandomGeneration(t *testing.T) {
	os.Setenv("mode", "test")
	config := file.NewConfig("config.json")
	output := file.NewOutput()
	config.Init()
	output.Init()
	defer output.Close()
	defer config.Close()
	random.Generated(10)
}
