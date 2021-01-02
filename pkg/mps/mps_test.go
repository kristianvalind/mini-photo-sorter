package mps

import (
	"testing"
	"time"
)

func TestGetDate(t *testing.T) {
	jackdawTime, err := GetDate("../../testdata/jackdaw.jpeg")
	if err != nil {
		t.Fatalf("jackdaw test failed: %v", err)
	}

	jackdawRefTime, _ := time.Parse(exifTimeFormat, "2020:06:13 19:03:26")

	if !jackdawTime.Equal(jackdawRefTime) {
		t.Fatalf("jackdaw test time failure: expected %v, but got %v", jackdawRefTime, jackdawTime)
	}
}
