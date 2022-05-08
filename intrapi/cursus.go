package intrapi

import "fmt"

type cursusId int

const (
	CURSUS_42CURSUS  cursusId = 22
	CURSUS_42PISCINE          = 10
)

func (c cursusId) String() string {
	if c > 0 {
		return fmt.Sprintf("&filter[cursus]=%d", c-1)
	}
	return ""
}
