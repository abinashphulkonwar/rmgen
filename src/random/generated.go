package random

import (
	"github.com/abinashphulkonwar/go-random-data-generation/src/file"
	"github.com/brianvoe/gofakeit/v6"
)

func Generated() {
	if file.ConfigPointer == nil {
		panic("Config pointer is nil")
	}

	for i := 0; i <= 10; i++ {
		for _, v := range *file.ConfigPointer.Get() {
			text := generated_handler(v.Type)
			println(text)
		}
	}

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
