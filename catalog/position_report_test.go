// Copyright 2015 Martin Hebnes Pedersen (LA5NTA). All rights reserved.
// Use of this source code is governed by the MIT-license that can be
// found in the LICENSE file.

package catalog

import (
	"os"
	"testing"
	"time"
)

func TestDecToDM(t *testing.T) {
	latTests := map[float64]string{
		-4.974: "04-58.4400S",
		-0.5:   "00-30.0000S",
		0.0:    "00-00.0000 ",
		0.5:    "00-30.0000N",
		60.132: "60-07.9200N",
	}
	lonTests := map[float64]string{
		-180.0: "180-00.0000W",
		-60.50: "060-30.0000W",
		-0.5:   "000-30.0000W",
		0.0:    "000-00.0000 ",
		0.5:    "000-30.0000E",
		003.50: "003-30.0000E",
		153.50: "153-30.0000E",
		180.0:  "180-00.0000E",
	}

	for deg, expect := range latTests {
		if got := decToMinDec(deg, true); got != expect {
			t.Errorf("On input %f, expected %s got %s", deg, expect, got)
		}
	}
	for deg, expect := range lonTests {
		if got := decToMinDec(deg, false); got != expect {
			t.Errorf("On input %f, expected %s got %s", deg, expect, got)
		}
	}
}

func omitErr(v *Course, _ error) *Course { return v }

func TestCourseStringer2(t *testing.T) {
	// omitErr := func(c *Course, err error) *Course { return c }
	tests := map[string]*Course{
		"123T": omitErr(NewCourse(123, false)),
		"123M": omitErr(NewCourse(123, true)),
	}
	for expect, in := range tests {
		t.Run(expect, func(t *testing.T) {
			if in == nil && expect != "" {
				t.Fatal("got unexpected nil")
			}
			if got := in.String(); got != expect {
				t.Errorf("Got %q, expected %q", got, expect)
			}
		})
	}
}

func TestCourseStringer(t *testing.T) {
	tests := map[string]Course{
		"123T": {Digits: [3]byte{'1', '2', '3'}, Magnetic: false},
		"123M": {Digits: [3]byte{'1', '2', '3'}, Magnetic: true},
	}
	for expect, in := range tests {
		t.Run(expect, func(t *testing.T) {
			if got := in.String(); got != expect {
				t.Errorf("Got %q, expected %q", got, expect)
			}
		})
	}
}

func ExamplePosReport_Message() {
	lat := 60.18
	lon := 5.3972

	posRe := PosReport{
		Date:    time.Now(),
		Lat:     &lat,
		Lon:     &lon,
		Comment: "Hjemme QTH",
	}
	msg := posRe.Message("N0CALL")
	msg.Write(os.Stdout)
}
