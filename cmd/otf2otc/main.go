// Copyright ©2021 The go-fonts Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Command otf2otc creates an OpenType Collection font file from a set of
// OpenType/TTF font files.
//
//   Usage: otf2otc [options] file1.ttf file2.ttf [file3.ttf [...]]
//
//   ex:
//    $> otf2otc -o foo.otc file1.ttf file2.ttf
//    $> otf2otc -o foo.otc file1.ttf file2.otf file3.otf file4.ttf
//
//   options:
//     -o string
//       	path to output OpenType Collection file (default "out.otc")
package main // import "github.com/go-fonts/melt/cmd/otf2otc"

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/go-fonts/melt"
	"golang.org/x/image/font/sfnt"
)

func main() {
	log.SetFlags(0)
	log.SetPrefix("otf2otc: ")

	var (
		oname = flag.String("o", "out.otc", "path to output OpenType Collection file")
	)

	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, `Usage: otf2otc [options] file1.ttf file2.ttf [file3.ttf [...]]

ex:
 $> otf2otc -o foo.otc file1.ttf file2.ttf
 $> otf2otc -o foo.otc file1.ttf file2.otf file3.otf file4.ttf

options:
`,
		)
		flag.PrintDefaults()
	}

	flag.Parse()

	if flag.NArg() <= 1 {
		log.Fatalf("missing input OTF/TTF files")
	}

	err := xmain(*oname, flag.Args())
	if err != nil {
		log.Fatalf("could not create collection file: %+v", err)
	}
}

func xmain(oname string, fnames []string) error {
	var (
		fnts = make([]*sfnt.Font, 0, len(fnames))
	)

	for _, fname := range fnames {
		raw, err := os.ReadFile(fname)
		if err != nil {
			return fmt.Errorf("could not read %q: %w", fname, err)
		}
		fnt, err := sfnt.Parse(raw)
		if err != nil {
			return fmt.Errorf("could not parse OTF/TTF file %q: %w", fname, err)
		}
		fnts = append(fnts, fnt)
	}

	o, err := os.Create(oname)
	if err != nil {
		return fmt.Errorf(
			"could not create output OpenType collection file %q: %w",
			oname, err,
		)
	}
	defer o.Close()

	err = melt.WriteCollection(o, fnts...)
	if err != nil {
		return fmt.Errorf("could not write output OpenType collection file: %w", err)
	}

	err = o.Close()
	if err != nil {
		return fmt.Errorf("could not close output OpenType collection file: %w", err)
	}

	return nil
}
