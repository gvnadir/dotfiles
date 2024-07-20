package fileops

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

func GetFloatFromFile(fileName string) (float64, error) {
	data, err := os.ReadFile(fileName)

	// try to remove the balance.txt file in order to get the error message and the fallback value
	if err != nil {
		return 1000, errors.New("failed to find file")
	}

	valueText := string(data) // string() here is used for convert byte[] to string
	value, err := strconv.ParseFloat(valueText, 64)

	if err != nil {
		return 1000, errors.New("failed to parse stored balance value")
	}

	return value, nil
}

func WriteFloatToFile(value float64, fileName string) {
	valueText := fmt.Sprint(value) // fmt.Sprint() is used for convert float64 to string
	os.WriteFile(fileName, []byte(valueText), 0644)
}
