package raindrops

import (
	"bytes"
	"strconv"
)

func Process(number int) (raindrops string) {
	buffer := bytes.NewBufferString("")

	if number%3 == 0 {
		buffer.WriteString("Pling")
	}
	if number%5 == 0 {
		buffer.WriteString("Plang")
	}
	if number%7 == 0 {
		buffer.WriteString("Plong")
	}

	raindropsBytes := buffer.Bytes()
	raindrops = string(raindropsBytes)

	if raindrops == "" {
		raindrops = strconv.Itoa(number)
	}
	return
}
