
build:
ifeq ($(OS),Windows_NT)
	# build asset-downloader
	go build -o ../bin/downloader.exe ../cmd/downloader/main.go
	# download assets
	../bin/downloader.exe
	# build main
	go build -o ../bin/mega8.exe ../cmd/mega8/main.go
else
	# build asset-downloader
	go build -o ../bin/downloader ../cmd/downloader/main.go
	# download assets
	../bin/downloader
	# build main
	go build -o ../bin/mega8 ../cmd/mega8/main.go
endif

