// Copyright ©2021 The go-fonts Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package melt

import (
	"fmt"
	"io"

	"golang.org/x/image/font/sfnt"
)

// WriteCollection writes the slice of sfnt fonts to the provided
// writer as an OpenType font collection.
func WriteCollection(w io.Writer, ttfs ...*sfnt.Font) error {
	var (
		fnts []font
		tbls []*table
		tmap = make(map[uint32][]*table)
		tidx = make(map[int]uint32)
	)

	for _, ttf := range ttfs {
		fnt, err := newFont(ttf)
		if err != nil {
			return fmt.Errorf("melt: could not parse font: %w", err)
		}
		fnts = append(fnts, fnt)
	}

	for i := range fnts {
		fnt := &fnts[i]
		for i := range fnt.tbls {
			tbl := fnt.tbls[i]
			name := tbl.name
			ts, ok := tmap[name]
			switch {
			case ok:
				var matched bool
				for _, te := range ts {
					if te.equal(tbl) {
						matched = true
						fnt.tbls[i] = te
						break
					}
				}
				if !matched {
					tmap[name] = append(tmap[name], tbl)
				}
			default:
				tmap[name] = append(tmap[name], tbl)
				tidx[i] = name
			}
		}
	}

	for i := 0; i < len(tidx); i++ {
		name := tidx[i]
		tbls = append(tbls, tmap[name]...)
	}

	ww := newWriter(w)
	ww.u32(0x74746366) // "ttcf"
	ww.u32(0x00010000)
	ww.u32(uint32(len(fnts)))
	offset := uint32(ttcHdrSize) + uint32(len(fnts)*4)
	for _, fnt := range fnts {
		ww.u32(offset)
		offset += sfntDirSize + uint32(len(fnt.tbls))*sfntDirEntrySize
	}

	for i := range tbls {
		tbl := tbls[i]
		pad := tbl.len + (4-(tbl.len&3))&3
		tbl.offset = offset
		offset += pad
	}

	for _, fnt := range fnts {
		ww.u32(fnt.typ)
		ww.u16(uint16(len(fnt.tbls)))
		ww.u16(fnt.search)
		ww.u16(fnt.eselect)
		ww.u16(fnt.rshift)
		for _, tbl := range fnt.tbls {
			ww.u32(tbl.name)
			ww.u32(tbl.chksum)
			ww.u32(tbl.offset)
			ww.u32(tbl.len)
		}
	}

	for _, tbl := range tbls {
		pad := tbl.len + (4-(tbl.len&3))&3
		ww.write(tbl.data)
		n := pad - tbl.len
		ww.write(make([]byte, n))
	}

	if ww.err != nil {
		return fmt.Errorf(
			"melt: could not write OpenType font collection: %w",
			ww.err,
		)
	}

	return nil
}
