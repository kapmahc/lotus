# lotus
A web application build by go-lang and reactjs.

## Install go

### For ubuntu

```
add-apt-repository ppa:ubuntu-lxc/lxd-stable
sudo apt-get update
sudo apt-get install golang
```

### For archlinux

```
sudo pacman -S go go-tools
```

### Add to your .bashrc or .zshrc

```
GOPATH=$HOME/go
PATH=$GOPATH/bin:$PATH
export GOPATH PATH
```

### For development

```
go get -u github.com/nsf/gocode
go get -u github.com/derekparker/delve/cmd/dlv
go get -u github.com/alecthomas/gometalinter
go get -u github.com/golang/lint/golint

go get -u github.com/kardianos/govendor

go get -u bitbucket.org/liamstask/goose/cmd/goose
go get -u github.com/kapmahc/lotus
```

### Start
```
cd $GOPATH/src/github.com/kapmahc/lotus
govendor sync
cd demo && go run main.go   # backend server

cd front-react
npm install
npm run start # frontend server
```

## Database creation

### postgresql

```
psql -U postgres
CREATE DATABASE db-name WITH ENCODING = 'UTF8';
CREATE USER user-name WITH PASSWORD 'change-me';
GRANT ALL PRIVILEGES ON DATABASE db-name TO user-name;
```

* ExecStartPre=/usr/bin/postgresql-check-db-dir ${PGROOT}/data (code=exited, status=1/FAILURE)

```
initdb  -D '/var/lib/postgres/data'
```

## Build

```
cd $GOPATH/src/github.com/kapmahc/lotus
make
ls -lh dist
```

## Documents
- [react](https://facebook.github.io/react/docs/getting-started.html)
- [react-bootstrap](http://react-bootstrap.github.io/)
- [redux](http://redux.js.org/docs/basics/UsageWithReact.html)
- [gin](https://github.com/gin-gonic/gin)
- [gorm](http://jinzhu.me/gorm/)
- [locale](https://blog.golang.org/matchlang)
- [govendor](https://github.com/kardianos/govendor)


- [go-plus](https://atom.io/packages/go-plus)
- [atom-beautify](https://atom.io/packages/atom-beautify)
- [react](https://atom.io/packages/react)

- [react dev-tools](https://facebook.github.io/react/blog/2015/09/02/new-react-developer-tools.html)
