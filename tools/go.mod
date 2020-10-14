module github.com/kamilsk/egg/tools

go 1.13

require (
	github.com/golang/mock v1.4.4
	github.com/golangci/golangci-lint v1.31.0
	github.com/kyoh86/looppointer v0.1.6
	golang.org/x/exp v0.0.0-20200513190911-00229845015e
	golang.org/x/tools v0.5.0
)

replace golang.org/x/tools => github.com/kamilsk/go-tools v0.0.5
