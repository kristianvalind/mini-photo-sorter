package mps

import "testing"

func TestGetDate(t *testing.T) {
	jackdawTime, err := GetDate("../../testdata/jackdaw.jpeg")
	if err != nil {
		t.Fatalf("jackdaw test failed: %v", err)
	}

	t.Log(jackdawTime)
}
