package cc

import (
	"testing"
)

func TestIndexAndFP(t *testing.T) {
	data := []byte("seif")
	numBuckets := uint(1024)
	i1, fp := getIndexAndFingerprint(data, numBuckets)
	i2 := getAltIndex(fp, i1, numBuckets)
	i11 := getAltIndex(fp, i2, numBuckets)
	i22 := getAltIndex(fp, i1, numBuckets)
	if i1 != i11 {
		t.Errorf("Expected i1 == i11, instead %d != %d", i1, i11)
	}
	if i2 != i22 {
		t.Errorf("Expected i2 == i22, instead %d != %d", i2, i22)
	}
}
