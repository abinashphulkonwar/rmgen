package random_test

import (
	"testing"

	"github.com/abinashphulkonwar/go-random-data-generation/src/file"
	"github.com/abinashphulkonwar/go-random-data-generation/src/random"
)

func TestRandomGeneration(t *testing.T) {
	config := file.NewConfig("config.json")
	config.Init()
	defer config.Close()
	random.Generated()
}
