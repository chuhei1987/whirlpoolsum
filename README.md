# WHIRLPOOLSUM(2)
## whirlpoolsum Implementation written in Go

<PRE>
Whirlpool Hashsum Tool - ALBANESE Lab (c) 2020-2021

Usage of whirlpoolsum:
whirlpoolsum [-v] [-c &lt;hash.ext&gt;] -t &lt;file.ext&gt;

  -c string
        Check hashsum file.
  -t string
        Target file/wildcard to generate hashsum list.
  -v    Verbose mode. (for CHECK command)
  </PRE>
  
### Examples:

#### Generate hashsum list:
<pre>
$ ./whirlpoolsum -t "*.*" > hash.txt
</pre>

#### Generate recursive hashsum list:
<pre>
$ find . -type f -name "*.*" -exec ./whirlpoolsum -t '{}' \; > hash.txt 
</pre>
##### Always works in binary mode. 

#### Check hashsum file:
<pre>
$ ./whirlpoolsum [-v] -c hash.txt
</pre>
##### Exit code is always 0 in vebose mode. 
