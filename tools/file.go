package tools

import (
	"bufio"
	"os"
	"path"
	"strings"
)

func ReadOne(filename string) (string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return "", err
	}
	scanner := bufio.NewScanner(file)
	if scanner.Scan() {
		return scanner.Text(), nil
	}
	return "", nil
}

func PathToName(pathStr string) string {
	return strings.TrimRight(path.Base(pathStr), path.Ext(pathStr))
}