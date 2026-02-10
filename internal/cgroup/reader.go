package cgroup

import (
	"os"
	"strconv"
	"strings"
)

func ReadFile(path string) (string, error) {
	b, err := os.ReadFile(path)
	if err != nil {
		return "", err
	}
	return strings.TrimSpace(string(b)), nil
}

func ReadInt(path string) (int64, error) {
	s, err := ReadFile(path)
	if err != nil {
		return 0, err
	}
	return strconv.ParseInt(s, 10, 64)
}
