# melt

[![GitHub release](https://img.shields.io/github/release/go-fonts/melt.svg)](https://github.com/go-fonts/melt/releases)
[![GoDoc](https://godoc.org/github.com/go-fonts/melt?status.svg)](https://godoc.org/github.com/go-fonts/melt)
[![License](https://img.shields.io/badge/License-BSD--3-blue.svg)](https://github.com/go-fonts/melt/blob/main/LICENSE)

`melt` provides a set of tools related with fonts, in Go.

## Example

```
$> go get github.com/go-fonts/melt/cmd/otf2otc
$> otf2otc -h
Usage: otf2otc [options] file1.ttf file2.ttf [file3.ttf [...]]

ex:
 $> otf2otc -o foo.otc file1.ttf file2.ttf
 $> otf2otc -o foo.otc file1.ttf file2.otf file3.otf file4.ttf

options:
  -o string
    	path to output OpenType Collection file (default "out.otc")


$> otf2otc -o out.otc LiberationMono-*.ttf
$> ll out.otc
-rw-r--r-- 1 binet binet 1.2M Jun 17 11:27 out.otc
```
