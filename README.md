# WHIRLPOOLSUM(2)
[![ISC License](http://img.shields.io/badge/license-ISC-blue.svg)](https://github.com/pedroalbanese/whirlpoolsum/blob/master/LICENSE.md) 
[![GoDoc](https://godoc.org/github.com/pedroalbanese/whirlpoolsum?status.png)](http://godoc.org/github.com/pedroalbanese/whirlpoolsum)
[![Go Report Card](https://goreportcard.com/badge/github.com/pedroalbanese/whirlpoolsum)](https://goreportcard.com/report/github.com/pedroalbanese/whirlpoolsum)
### Whirlpool Recursive Hashsum Implementation written in Go

<PRE>
Usage of whirlpoolsum:
whirlpoolsum [-v] [-c &lt;hash.ext&gt;] [-r] -t &lt;file.ext&gt;

  -c string
        Check hashsum file.
  -r    Process directories recursively.
  -t string
        Target file/wildcard to generate hashsum list.
  -v    Verbose mode. (The exit code is always 0 in this mode)
  </PRE>
  
### Examples:

#### Generate hashsum list:
<pre>
$ ./whirlpoolsum [-r] -t "*.*" > hash.txt
</pre>
##### Always works in binary mode. 

#### Check hashsum file:
<pre>
$ ./whirlpoolsum [-v] -c hash.txt
</pre>
##### Exit code is always 0 in vebose mode. 

## License

This project is licensed under the ISC License.
##### Copyright (c) 2020-2021 Pedro Albanese - ALBANESE Lab.
