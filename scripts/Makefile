
build:
	# build asset-downloader
	go build -o ../bin/downloader ../cmd/downloader/main.go
	# download assets
	../bin/downloader
	# build main
	go build -o ../bin/mega8 ../cmd/mega8/main.go
