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

The `<url to catch>` must be a valid regexp.

By example you can create the following config file :

```text
path,contentType,file
*.png,image/jpeg,myLolCat.jpg
```

which replace all png images by a nice jpeg cat.

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

Examples
--------

See the beautiful example in the download section!
