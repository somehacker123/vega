package orders

import (
	"testing"
	"github.com/magiconair/properties/assert"
	"time"
)

func TestUnixTimestamp(t *testing.T) {

	cases := []struct {
		expected uint64
		datetime string
	}{
		{1415792726371,"2014-11-12T11:45:26.371Z"},
		{1514768400000,"2018-01-01T01:00:00.000Z"},
		{406054800111,"1982-11-13T17:00:00.111Z"},
		{1591935315123,"2020-06-12T04:15:15.123Z"},
	}

	for _, c := range cases {
		layout := "2006-01-02T15:04:05.000Z"
		parsed, _ := time.Parse(layout , c.datetime)
		res := unixTimestamp(parsed)

		assert.Equal(t, res, c.expected)
	}
}