package random_test

import (
	"os"
	"testing"

	"github.com/abinashphulkonwar/go-random-data-generation/src/file"
	"github.com/abinashphulkonwar/go-random-data-generation/src/random"
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
