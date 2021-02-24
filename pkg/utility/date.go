package utility

import (
	"fmt"
	"time"
)

func TimeStampToDate(timestamp int64) time.Time {
	tm := time.Unix(timestamp, 0)
	return tm
}

func GetFormattedTime(d time.Time) string {
	hours, minutes, _ := d.Clock()
	return fmt.Sprintf("%02d:%02d", hours, minutes)
}
