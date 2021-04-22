module gostarter

go 1.16

//module is collection of go packages, there can be more than one package in go module
// the go.mod file defines the modules module path
// the go command enables the use of modules when the current directory or any parent directory has a go.mod, provided the directory is outside $GOPATH/src
// Read more at https://blog.golang.org/using-go-modules

require (
	golang.org/x/crypto v0.0.0-20210322153248-0c34fe9e7dc2
	golang.org/x/text v0.3.6 // indirect
	rsc.io/quote v1.5.2
	rsc.io/sampler v1.3.1 // indirect
)
