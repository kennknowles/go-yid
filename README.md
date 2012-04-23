Go Yid!
=======

*Y*acc *I*s *D*ead

Parsing (actually just a recognizer at the moment) for 
arbitrary context-free-grammars (i.e. recursive regexes) 
inspired by and based on the work listed below in Further Reading.

Written in Go because [April is "learn a new programming language" month](http://matt.might.net/articles/programmers-resolutions/) 

Enjoy! :-)


Run it!
-------

This comes with an OK test suite and a main program for benchmarking
a particular ambiguous grammar.

    $ git clone https://github.com/kennknowles/go-yid.git
    $ cd go-yid
    $ GOPATH=$PWD go test yid
    $ GOPATH=$PWD go run main.go | tee timings.csv


The Story and Caveats
---------------------

Matt Might [called for implementations](http://matt.might.net/articles/parsing-with-derivatives/) 
in various languages, preferably with benchmarks. I have gone halfway - this is a half-baked implementation in
an esoteric language that is not really well suited to this problem. But I
did take amateurish timings that indicate the "controversial" exponential
example is certainly recognized in less than exponential time using this technique.

Caveats:

 - This implementation is slow by a huge constant factor, for which I am
   certainly to blame. Is it the emulation of algebraic data types? Are
   there places I could be more mutatey? Should I put laziness in places
   other than where it is needed for termination? Probably all of the above.
 - Thus far, I have only implemented a recognizer, not a parser, because this is
   really toy code and I just wanted to see it go and get some benchmarks.
 - There are probably bugs. In particular my recursive examples do not seem 
   to always have the memory layout I expect, and it changes based on 
   unrelated code, even in other files. This seems to affect only
   the pretty-printing test suite.
 - I probably won't take it further because I've had my fun with Go, but I
   encourage anyone interested in parsing to fork and improve, or write
   their own.

So, you can see the timings I took (on a modern-ish MacBook Pro with an SSD) in timings.csv
and a badly-labeled graphic of them in timings.png (Y-Axis is seconds, X-Axis is O(length of input))


Further reading
---------------

In roughly reverse chronological (causal?) order.

 * [Parsing With Derivatives](http://matt.might.net/papers/might2011derivatives.pdf) by Matthew Might, David Darais, and Daniel Spiewak, International Conference on Functional Programming, 2011.
 * [Yacc is dead: An update](http://matt.might.net/articles/parsing-with-derivatives/) by Matt Might on his blog.
 * [Yacc Is Not Dead](http://research.swtch.com/yaccalive) by Russ Cox on his blog.
 * [Yacc Is Dead](http://arxiv.org/abs/1010.5023) by Matthew Might and David Darais, rejected from ESOP 2009.
 * [Regular-expression derivatives reexamined](http://www.ccs.neu.edu/home/turon/re-deriv.pdf) by Scott Owens, John Reppy, and Aaron Turon, Journal of Functional Programming, 2009.
 * [Derivatives of Regular Expressions](http://dl.acm.org/citation.cfm?id=321249) by Janusz Brzozowski, Journal of the ACMi, 1964. (paywall, sorry!)


Copyright & License
-------------------
Copyright 2012- Kenneth Knowles

Licensed under the Apache License, Version 2.0 (the "License"); you may not use
this file except in compliance with the License. You may obtain a copy of the
License at

http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software distributed
under the License is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR
CONDITIONS OF ANY KIND, either express or implied. See the License for the
specific language governing permissions and limitations under the License.



