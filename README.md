Sleipnir
========

Sleipnir is a very little proxy which can serve some local files instead of the requested ones.

Usage
-----

First, copy the `config.csv.dist` file to `config.csv`.

Then, edit the csv. This is the pattern of each row :

```text
<url to catch>,<content type of the response>,<location of the file to server instead>
```

Run it :

```sh
go build sleipnir.go
./sleipnir
```

or download a binary if exists.

Usage :

```sh
./sleipnir -h
Usage of ./sleipnir:
  -a=":8888": Bind to this address:port
  -c="config.csv": Config file
  -h=false: Print this help
  -v=false: Verbose
```
