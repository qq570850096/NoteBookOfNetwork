package NextDay

import (
	"strconv"
	"testing"
)

func TestBigMonth(t *testing.T) {
	f := Unit{}
	year,month,day := 2008,12,0

	s1 := strconv.Itoa(year)
	s2 := strconv.Itoa(month)
	s3 := strconv.Itoa(day)
	f.NextDay(s1,s2,s3)
}
