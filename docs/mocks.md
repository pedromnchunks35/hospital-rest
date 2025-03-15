# Mock
- We are using mockgen from google
## Steps to mock
```
## INSTALL MOCKGEN
go install github.com/golang/mock/mockgen@latest

## ADD THE GO BINARIES INTO OUR PATH
export PATH=$PATH:$(go env GOPATH)/bin
## OR FOR MACOS (Figure out in your OS because idc)
echo "export PATH=$PATH:$(go env GOPATH)/bin" >> $HOME/.zshrc && source 
$HOME/.zshrc

## GENERATE THE MOCKS
mockgen -source=<FILE HERE THE INTERFACE IS>/*.go 
-destination=<FILE WHERE THE MOCK WILL BE>/*_mock.go -package=<MOCKS PACKAGE>
## IN OUR PROJECT
mockgen -source=service/testMock.go 
-destination=mocks/radar_mock.go -package=mocks
```
