package random

import (
	"github.com/abinashphulkonwar/go-random-data-generation/src/file"
	"github.com/brianvoe/gofakeit/v6"
)

const MAX_FILE_SIZE = 1024 * 1024 * 5

func Generated() {
	if file.OutputPointer == nil {
		panic("Output pointer is nil")
	}
	if file.ConfigPointer == nil {
		panic("Config pointer is nil")
	}
	buf := *new([]byte)
	new_line := []byte("\n")
	separator := []byte(",")
	length := len((*file.ConfigPointer.Get()))
	for index, v := range *file.ConfigPointer.Get() {
		buf = append(buf, []byte(v.Type)...)
		if index < length-1 {
			buf = append(buf, separator...)
		}
	}
	buf = append(buf, new_line...)
	length_of_random_data := 10000
	for i := 0; i < length_of_random_data; i++ {
		for index, v := range *file.ConfigPointer.Get() {
			text := generated_handler(v.Type)
			buf = append(buf, []byte(text)...)
			if index < length-1 {
				buf = append(buf, separator...)
			}
		}
		if i != length_of_random_data {
			buf = append(buf, new_line...)
		}
		if len(buf) > MAX_FILE_SIZE {
			file.OutputPointer.Write(&buf)
			buf = *new([]byte)
		}
	}
	file.OutputPointer.Write(&buf)
}

func generated_handler(Type string) string {
	switch Type {
	case file.EMAIL_TYPE:
		return gofakeit.Email()
	case file.NAME_TYPE:
		return gofakeit.Name()
	case file.PHONE_TYPE:
		return gofakeit.Phone()
	case file.ADDRESS_TYPE:
		address := gofakeit.Address()
		return address.Address + " " + address.City + " " + address.State + " " + address.Zip
	}

	return gofakeit.Emoji()
}
