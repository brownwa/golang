# Creating a package
- $ cd /path/to/toolbox
- $ go mod init some.domain/toolbox
- Write code that imports the module, as shown in toolbox.go
- $ go install some.domain/toolbox
- Add Go binary install directory to PATH, then execute binary
  - $ export PATH=$PATH:$(dirname $(go list -f '{{.Target}}' .))
  - $ toolbox
- [optional] Build and run in the toolbox directory
  - $ go build toolbox.go
  - $ ./toolbox

# References
- [Your first program](https://golang.org/doc/code.html#Command)