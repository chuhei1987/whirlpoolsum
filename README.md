# WHIRLPOOLSUM(2)
## whirlpoolsum Implementation written in Go

<PRE>
Whirlpool Hashsum Tool - ALBANESE Lab (c) 2020-2021

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
