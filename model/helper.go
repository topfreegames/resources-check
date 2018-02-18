package model

import (
	"fmt"
	"time"
)

func currentTime() string {
	t := time.Now().UTC()
	return fmt.Sprint(t.Format("2006-01-02 15:04:05"))
}
