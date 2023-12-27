package helper

import (
	"bufio"
	"encoding/csv"
	"io"
)

func ReadAllRecordsWithoutHeader(file io.ReadSeeker) ([][]string, error) {
	firstRow, err := bufio.NewReader(file).ReadSlice('\n')
	if err != nil {
		return nil, err
	}

	_, err = file.Seek(int64(len(firstRow)), io.SeekStart)
	if err != nil {
		return nil, err
	}

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}

	return records, nil
}
