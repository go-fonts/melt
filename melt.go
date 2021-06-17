// Copyright ©2021 The go-fonts Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package melt provides tools for fonts.
package melt // import "github.com/go-fonts/melt"

import (
	"bytes"
	"fmt"

	"golang.org/x/image/font/sfnt"
)

const (
	ttcHdrSize       = 4 + 2*4
	sfntDirSize      = 4 + 4*2
	sfntDirEntrySize = 4 + 3*4
)

type font struct {
	typ     uint32
	ntbls   uint16
	search  uint16
	eselect uint16
	rshift  uint16

	tbls []*table
}

func newFont(f *sfnt.Font) (fnt font, err error) {
	buf := new(bytes.Buffer)
	_, err = f.WriteSourceTo(nil, buf)
	if err != nil {
		return fnt, fmt.Errorf("melt: could not retrieve raw font data: %w", err)
	}

	raw := buf.Bytes()
	r := &reader{p: raw}
	fnt.typ = r.u32()
	fnt.ntbls = r.u16()
	fnt.search = r.u16()
	fnt.eselect = r.u16()
	fnt.rshift = r.u16()

	for i := 0; i < int(fnt.ntbls); i++ {
		fnt.tbls = append(fnt.tbls, r.table(raw))
	}

	return fnt, nil
}

func (fnt *font) table(name uint32) *table {
	for i := range fnt.tbls {
		tbl := fnt.tbls[i]
		if tbl.name == name {
			return tbl
		}
	}
	return nil
}

type table struct {
	name   uint32
	chksum uint32
	len    uint32
	offset uint32
	data   []byte
}

func (tbl *table) equal(rhs *table) bool {
	return rhs.chksum == tbl.chksum && rhs.len == tbl.len && bytes.Equal(rhs.data, tbl.data)
}
