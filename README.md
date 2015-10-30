# close2someone

## how to install and run

- install go
- set GOPATH
- install revel
- clone this repo to src/github.com/yamanetoshi
- run this project
- open localhost:9000

### install Go

for example

```
$ sudo apt-get install go
```

### set GOPATH

make directory and add it to ~/.bash_profile

```
export GOLANG=$HOME/.go
```

and source it

### install revel

below

```
$ go get github.com/revel/revel
$ go get github.com/revel/cmd/revel
```

and add PATH line

```
PATH=$PATH:$GOPATH/bin
```

### clone this repo to src/github.com/yamanetoshi

below

```
$ cd $GOPATH/src/github
$ mkdir yamanetoshi && cd yamanetoshi
$ git clone https://github.com/yamanetoshi/close2someone

### run this project

```
$ revel run github.com/yamanetoshi/close2someone
```

### open localhost:9000

below

```
$ w3m http://localhost:9000
```
