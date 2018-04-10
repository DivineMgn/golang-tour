// https://go-tour-ru-ru.appspot.com/methods/23

package main

import (
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	reader io.Reader
}

func main() {
	str := strings.NewReader("Lbh penpxrq gur pbqr!")
	reader := rot13Reader{str}
	io.Copy(os.Stdout, &reader)
}

func (rot13R rot13Reader) Read(data []byte) (int, error) {
	cnt, err := rot13R.reader.Read(data)

	if err != nil {
		return 0, err
	}

	for idx := 0; idx < cnt; idx++ {
		data[idx] = decryptRot13(data[idx])
	}

	return cnt, nil
}

func decryptRot13(char byte) byte {
	switch {
	case (char >= 'A' && char < 'N') ||
		(char >= 'a' && char < 'n'):
		return char + 13
	case (char > 'M' && char <= 'Z') ||
		(char > 'm' && char <= 'z'):
		return char - 13
	default:
		return char
	}
}
