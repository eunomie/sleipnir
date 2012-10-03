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

Run it : `go run main.go`

The proxy use the port `8888` by default.