// Copyright ©2021 The go-fonts Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package melt_test

import (
	"bytes"
	"crypto/md5"
	"fmt"
	"testing"

	"github.com/go-fonts/liberation/liberationmonobold"
	"github.com/go-fonts/liberation/liberationmonobolditalic"
	"github.com/go-fonts/liberation/liberationmonoitalic"
	"github.com/go-fonts/liberation/liberationmonoregular"
	"golang.org/x/image/font/sfnt"

	"github.com/go-fonts/melt"
)

func TestWriteCollection(t *testing.T) {
	parse := func(raw []byte) *sfnt.Font {
		fnt, err := sfnt.Parse(raw)
		if err != nil {
			t.Fatalf("could not parse font: %+v", err)
		}
		return fnt
	}

	ttfs := []*sfnt.Font{
		parse(liberationmonobold.TTF),
		parse(liberationmonobolditalic.TTF),
		parse(liberationmonoitalic.TTF),
		parse(liberationmonoregular.TTF),
	}

	buf := new(bytes.Buffer)
	err := melt.WriteCollection(buf, ttfs...)
	if err != nil {
		t.Fatalf("could not create OTC: %+v", err)
	}

	sum := fmt.Sprintf("%x", md5.Sum(buf.Bytes()))
	if got, want := sum, "1c12db5f8e582ee29828d7face845aaa"; got != want {
		t.Fatalf("invalid md5 checksum:\ngot= %s\nwant=%s\n", got, want)
	}
}
