dst=dist

build:
	go build -ldflags "-s -X main.version=`git rev-parse --short HEAD`" -o $(dst)/lotus main.go
	-cp -rv locales themes db $(dst)/

clean:
	-rm -rv $(dst)
