set GOOS=linux
set GOARCH=amd64
go mod tidy
go build -o main main.go
build-lambda-zip.exe -output main.zip main
del main
set GOOS=windows