// Copyright ©2021 The go-fonts Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package melt

import (
	"encoding/binary"
)

type reader struct {
	p []byte
	c int
}

func (r *reader) u8() uint8 {
	v := r.p[r.c]
	r.c++
	return v
}

func (r *reader) u16() uint16 {
	c := r.c
	v := binary.BigEndian.Uint16(r.p[c:])
	r.c += 2
	return v
}

func (r *reader) u32() uint32 {
	c := r.c
	v := binary.BigEndian.Uint32(r.p[c:])
	r.c += 4
	return v
}

func (r *reader) bytes(n int) []byte {
	c := r.c
	v := r.p[c : c+n]
	r.c += n
	return v
}

func (r *reader) table(raw []byte) *table {
	tbl := &table{
		name:   r.u32(),
		chksum: r.u32(),
		offset: r.u32(),
		len:    r.u32(),
	}
	tbl.data = raw[tbl.offset : tbl.offset+tbl.len]
	return tbl
}
