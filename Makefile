dst=dist

build:
	go build -ldflags "-s -X main.version=`git rev-parse --short HEAD`" -o $(dst)/lotus main.go
	-cp -rv locales db $(dst)/
	cd front-vue && npm run build
	-cp -rv front-vue/dist $(dst)/public

clean:
	-rm -rv $(dst)
