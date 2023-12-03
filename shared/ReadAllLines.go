package shared

import (
	"bufio"
	"io"
	"os"
	"strings"
)

// ReadAllLines reads all lines from a file and returns them as a slice of strings.
func ReadAllLines(path string) ([]string, error) {
	var lines []string

	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}

	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			panic(err)
		}
	}(file)

	reader := bufio.NewReader(file)

	for {
		line, err := reader.ReadString('\n')
		if err != nil && err != io.EOF {
			return lines, err
		}
		line = strings.TrimRight(line, "\r\n")
		lines = append(lines, line)

		if err == io.EOF {
			break
		}
	}

	return lines, nil
}
