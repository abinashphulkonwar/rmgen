package commandshandler

import (
	"github.com/abinashphulkonwar/rmgen/src/file"
	"github.com/abinashphulkonwar/rmgen/src/random"
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
