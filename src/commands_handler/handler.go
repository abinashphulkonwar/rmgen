package commandshandler

import (
	"github.com/abinashphulkonwar/go-random-data-generation/src/file"
	"github.com/abinashphulkonwar/go-random-data-generation/src/random"
)

func Handler(file_path string, count int) {
	config := file.NewConfig(file_path)
	output := file.NewOutput()
	config.Init()
	output.Init()
	defer output.Close()
	defer config.Close()
	random.Generated(count)
}
