package macromeasures

import (
	"strconv"
	"testing"
	"time"
)

func TestUnmarshalJSON(t *testing.T) {
	mtime := &Time{}
	json := []byte("1472083200")
	if err := mtime.UnmarshalJSON(json); err != nil {
		t.Errorf("UnmarshalJSON returned error: %v", err)
	}
	want := time.Date(2016, time.August, 25, 0, 0, 0, 0, time.UTC)
	if mtime.Time.Unix() != want.Unix() {
		t.Errorf("UnmarshalJSON set time to %v, wanted %v", mtime.Time.Unix(), want.Unix())
	}
}

func TestMarshallJSON(t *testing.T) {
	now := time.Now().UTC()
	mtime := &Time{Time: now}
	resp, err := mtime.MarshalJSON()
	if err != nil {
		t.Errorf("MarshalJSON returned error: %v", err)
	}
	str, _ := strconv.Unquote(string(resp))
	ts, err := strconv.Atoi(str)
	if err != nil {
		t.Errorf("strconv.Atoi returned error: %v", err)
	}
	if now.Unix() != int64(ts) {
		t.Errorf("MarshalJSON set time to %v, wanted %v", ts, now.Unix())
	}
}
