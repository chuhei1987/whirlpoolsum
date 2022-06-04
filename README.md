# WHIRLPOOLSUM(2)
[![ISC License](http://img.shields.io/badge/license-ISC-blue.svg)](https://github.com/pedroalbanese/whirlpoolsum/blob/master/LICENSE.md) 
[![GitHub downloads](https://img.shields.io/github/downloads/pedroalbanese/whirlpoolsum/total.svg?logo=github&logoColor=white)](https://github.com/pedroalbanese/whirlpoolsum/releases)
[![GoDoc](https://godoc.org/github.com/pedroalbanese/whirlpoolsum?status.png)](http://godoc.org/github.com/pedroalbanese/whirlpoolsum)
[![Go Report Card](https://goreportcard.com/badge/github.com/pedroalbanese/whirlpoolsum)](https://goreportcard.com/report/github.com/pedroalbanese/whirlpoolsum)
[![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/pedroalbanese/whirlpoolsum)](https://golang.org)
[![GitHub release (latest by date)](https://img.shields.io/github/v/release/pedroalbanese/whirlpoolsum)](https://github.com/pedroalbanese/whirlpoolsum/releases)
### Whirlpool Recursive Hashsum Implementation written in Go

<PRE>
Usage of whirlpoolsum:
whirlpoolsum [-c &lt;hash.ext&gt;] [-r] &lt;file.ext&gt;
  -c string
        Check hashsum file.
  -r    Process directories recursively.
</PRE>
  
### Examples:

#### Generate hashsum list:
```sh
$ ./whirlpoolsum [-r] "*.*" > hash.txt
```

#### Check hashsum file:
```sh
$ ./whirlpoolsum -c hash.txt
$ echo $?
```

## License

This project is licensed under the ISC License.
##### Copyright (c) 2020-2022 Pedro F. Albanese - ALBANESE Research Lab.
