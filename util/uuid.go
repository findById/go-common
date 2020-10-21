package util

import (
	"fmt"
	"github.com/google/uuid"
	"os"
	"strings"
	"time"
)

var (
	pid = os.Getpid()
)

func MustUUID() string {
	return uuid.New().String()
}

func NewUUID() (string, error) {
	v, err := uuid.NewRandom()
	if err != nil {
		return "", err
	}
	return v.String(), nil
}

func NewTraceId() string {
	return fmt.Sprintf("trace-id-%d-%s",
		pid,
		strings.ReplaceAll(time.Now().Format("2006.01.02.15.04.05.999999"), ".", ""))
}
