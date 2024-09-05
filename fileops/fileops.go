package fileops

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

func PersistFloatIntoFile(value float64, fileName string) {
	_ = os.WriteFile(fileName, []byte(fmt.Sprintf("%f", value)), 0644)
}

func ReadFloatFromFile(fileName string, defaultBalance float64) (float64, error) {
	bytes, err := os.ReadFile(fileName)
	if err != nil {
		return defaultBalance, errors.New("unable to find file")
	}
	valueText := string(bytes)
	value, err := strconv.ParseFloat(valueText, 64)
	if err != nil {
		return defaultBalance, errors.New("unable to parse balance file")
	}
	return value, nil
}
