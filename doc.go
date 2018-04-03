/*
Package gox is a minimalistic extension to Go. It contains constants, helpers
and utilities which could have been part of Go itself (could have been
built-in).

The package is minimalistic, and introduces no dependency to any package.
Most of the functions are eligible for inlining. And don't worry if you're not
using some of the functions, the compiler will exclude those from your binary.

An easy way to use this library is to "dot-import" the package so identifiers
will be directly available, see the package example.

*/
package gox