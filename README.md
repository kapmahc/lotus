# lotus
A web application build by go-lang.

## Development(for archlinux)
### Install go

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

go get -u github.com/beego/bee
go get -u bitbucket.org/liamstask/goose/cmd/goose

go get -u github.com/kapmahc/lotus
```

## Build
```
cd $GOPATH/src/github.com/kapmahc/lotus
goose up
bee pack
```

## Database creation

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


## Documents
* [googse](https://bitbucket.org/liamstask/goose/)
* [beego](https://beego.me/docs/intro/)
