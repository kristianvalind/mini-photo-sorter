package mps

import (
	"testing"
	"time"
)

func TestGetDate(t *testing.T) {
	jackdawTime, err := GetDate("../../testdata/jpeg/jackdaw.jpeg")
	if err != nil {
		t.Fatalf("jackdaw test failed: %v", err)
	}

	jackdawRefTime, _ := time.Parse(exifTimeFormat, "2020:06:13 19:03:26")

	if !jackdawTime.Equal(jackdawRefTime) {
		t.Fatalf("jackdaw test time failure: expected %v, but got %v", jackdawRefTime, jackdawTime)
	}

	screenshotTime, err := GetDate("../../testdata/png/iphone screenshot.png")
	if err != nil {
		t.Fatalf("screenshot test failed: %v", err)
	}

	screenshotRefTime, _ := time.Parse(exifTimeFormat, "2011:11:09 00:47:56")
	if !screenshotTime.Equal(screenshotRefTime) {
		t.Fatalf("screenshot test time failure: expected %v, but got %v", screenshotRefTime, screenshotTime)
	}

}
