# gerd

## install
```
$ git clone https://github.com/ritarock/gerd.git
$ cd gerd
$ make install
```

## Usage
```
$ gerd -h
make entity relationship diagram

Usage:
  gerd [flags]
  gerd [command]

Available Commands:
  completion  Generate the autocompletion script for the specified shell
  help        Help about any command
  server      show created mermaid file

Flags:
  -a, --address string    connection address (required)
  -d, --db string         connection db name (required)
  -h, --help              help for gerd
  -p, --password string   database password (required)
  -u, --user string       database user name (required)

$ gerd server -h
show mermaid file.
access to http://localhost:8080

Usage:
  gerd server [flags]

Flags:
  -h, --help   help for server
```

## Sample
```
$ docker-compose up database
$ gerd -u user -p pass -a localhost:3306 -d app

$ gerd server

server started
http://localhost:8080
```
![](https://raw.githubusercontent.com/ritarock/gerd/main/etc/mermaid.png)


## development
Run `cobra-cli` on docker

```bash
$ docker-compose run cobra-cli <command>
```
