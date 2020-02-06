This post introduced these workflows using Go modules:

go mod init creates a new module, initializing the go.mod file that describes it.
go build, go test, and other package-building commands add new dependencies to go.mod as needed.
go list -m all prints the current module’s dependencies.
go get changes the required version of a dependency (or adds a new dependency).
go mod tidy removes unused dependencies.

The indirect comment indicates a dependency is not used directly by this module, only indirectly by other module dependencies. See go help modules for details.


go help modules
go mod init example.com/hello
go get golang.org/x/text
go list -m -versions rsc.io/sampler
go list -m rsc.io/q...
go get rsc.io/sampler@v1.3.1
go doc rsc.io/quote/v3