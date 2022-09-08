package format

import (
	"fmt"
	"time"
)

type Marshaler interface {
	MarshalJSON() ([]byte, error)
}

type JSONTime time.Time

func (t JSONTime) MarshalJSON() ([]byte, error) {
	timeStr := fmt.Sprintf("\"%s\"", time.Time(t).Format("2006-01-02 15:04:05"))
	return []byte(timeStr), nil
}
