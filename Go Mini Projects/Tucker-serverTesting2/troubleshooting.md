## All install packages

- `goconvey`
- `testify/assert`
  - `go install github.com/stretchr/testify/assert@latest`
  
## Cannot install `goconvey`

For some reason, the command (`go get github.com/smartystreets/goconvey`) suggested in the official guide doesn't work. Instead, I tried using `go install` command 'and it worked:

`go install github.com/smartystreets/goconvey@latest`
