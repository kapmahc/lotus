# lotus
A web application build by go-lang.

## Development(for archlinux)
* install gvm
```
bash < <(curl -s -S -L https://raw.githubusercontent.com/moovweb/gvm/master/binscripts/gvm-installer)
```
Or if you are using zsh just change bash with zsh

* after re-signin, install go
```
gvm install go1.7.3 -B
gvm use go1.7.3 --default
```


* install tools

```
go get -u github.com/nsf/gocode
go get -u github.com/derekparker/delve/cmd/dlv
go get -u github.com/alecthomas/gometalinter
go get -u github.com/golang/lint/golint
go get -u golang.org/x/tools/cmd/goimports

go get -u github.com/kardianos/govendor

go get -u github.com/kapmahc/lotus
# OR
git clone https://github.com/kapmahc/lotus.git $GOPATH/src/github.com/kapmahc/lotus
```

* fix "go run main.go" very slowly
```
go get -u github.com/mattn/go-sqlite3
rm -rf vendor/github.com/mattn/go-sqlite3
```

* run backend
```
cd $GOPATH/src/github.com/kapmahc/lotus
govendor sync
sudo pacman -S inotify-tools
./run.sh
```

* run frontend
```
cd front-vue
npm install
npm run dev
```

## Build
```
cd $GOPATH/src/github.com/kapmahc/lotus
make
ls -l dist
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
* [vuejs](https://vuejs.org/v2/guide/)
* [vuex](https://vuex.vuejs.org/en/getting-started.html)
* [vue-router](https://router.vuejs.org/en/index.html)
* [vue devtools](https://github.com/vuejs/vue-devtools)
* [gin](https://github.com/gin-gonic/gin)
* [gorm](http://jinzhu.me/gorm/)
