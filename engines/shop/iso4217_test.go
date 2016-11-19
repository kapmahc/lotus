package shop_test

import (
	"encoding/xml"
	"os"
	"testing"

	"github.com/kapmahc/lotus/engines/shop"
)

func TestISO4217(t *testing.T) {
	fd, err := os.Open("../../db/seed/list_one.xml")
	if err != nil {
		t.Fatal(err)
	}
	defer fd.Close()

	var val shop.ISO4217
	dec := xml.NewDecoder(fd)
	if err = dec.Decode(&val); err != nil {
		t.Fatal(err)
	}

	t.Logf("%+v", val)
}
