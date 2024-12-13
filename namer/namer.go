package namer

import (
	"fmt"
	"time"
)

func GenerateName() string {
	const prefix = "vistar"
	timestamp := time.Now().Format("20060102_150405") // Format: YYYYMMDD_HHMMSS
	return fmt.Sprintf("%s_%s", prefix, timestamp)
}