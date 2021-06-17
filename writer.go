// Copyright ©2021 The go-fonts Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package melt

import (
	"encoding/binary"
	"io"
)

type writer struct {
	w   io.Writer
	buf []byte
	err error

	c int
}

func newWriter(w io.Writer) *writer {
	return &writer{w: w, buf: make([]byte, 4)}
}

func (w *writer) u8(v uint8) {
	w.buf[0] = v
	w.write(w.buf[:1])
}

func (w *writer) u16(v uint16) {
	binary.BigEndian.PutUint16(w.buf, v)
	w.write(w.buf[:2])
}

func (w *writer) u32(v uint32) {
	binary.BigEndian.PutUint32(w.buf, v)
	w.write(w.buf[:4])
}

func (w *writer) write(p []byte) {
	if w.err != nil {
		return
	}
	n, err := w.w.Write(p)
	if err != nil {
		w.err = err
	}
	w.c += n
}
