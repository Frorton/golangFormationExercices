// NOTE: Golint is deprecated and frozen.
// There's no drop-in replacement for it, but tools such as Staticcheck and go vet should be used instead.

/*
● gofmt
	○ formats go code
● go vet
	○ reports suspicious constructs
			To use it go to your directory where all files are located :
			go vet ./...
● golint
	○ reports poor coding style
		⚠ Isn't perfect
			To use it go to your directory where all files are located :
			golint ./... will report for the directory you're in
			+ all the files down the respective directories
● staticcheck
	staticcheck . will check the current package, and
	staticcheck ./... will check all packages directly under the one you're in.
*/

package main
