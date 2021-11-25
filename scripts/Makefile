
build:
ifeq ($(OS),Windows_NT)
	go build -o ../bin/downloader.exe ../cmd/downloader/main.go
	../bin/downloader.exe
	go build -o ../bin/mega8.exe ../cmd/mega8/main.go
else
	go build -o ../bin/downloader ../cmd/downloader/main.go
	../bin/downloader
	go build -o ../bin/mega8 ../cmd/mega8/main.go
endif

