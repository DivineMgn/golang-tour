// https://go-tour-ru-ru.appspot.com/methods/22

package main

import "golang.org/x/tour/reader"

func main() {
	reader.Validate(MyReader{})
}

// MyReader -my own reader
type MyReader struct{}

func (reader MyReader) Read(data []byte) (int, error) {
	len := len(data)

	for idx := range data {
		data[idx] = 'A'
	}

	return len, nil
}
