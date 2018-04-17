# go-overlap

Overlap timezone utility

## Install

```sh
go get github.com/sirodoht/go-overlap
```

Also, there are compiled binaries at the [releases page](https://github.com/sirodoht/go-overlap/releases).

After downloading you need to make them executable:
```sh
chmod +x go-overlap-darwin-amd64
mv go-overlap-darwin-amd64 overlap
./overlap utc utc-3
```

## Use

```
$ overlap utc-4 utc+3

+-----------+-----------+
|   UTC-4   |   UTC+3   |
+-----------+-----------+
|   20:00   |   03:00   |
|   21:00   |   04:00   |
|   22:00   |   05:00   |
|   23:00   |   06:00   |
|   00:00   |   07:00   |
|   01:00   |   08:00   |
|   02:00   |   09:00   |
|   03:00   |   10:00   |
|   04:00   |   11:00   |
|   05:00   |   12:00   |
|   06:00   |   13:00   |
|   07:00   |   14:00   |
|   08:00   |   15:00   |
|   09:00   |   16:00   |
|   10:00   |   17:00   |
|   11:00   |   18:00   |
|   12:00   |   19:00   |
|   13:00   |   20:00   |
|   14:00   |   21:00   |
|   15:00   |   22:00   |
|   16:00   |   23:00   |
|   17:00   |   00:00   |
|   18:00   |   01:00   |
|   19:00   |   02:00   |
+-----------+-----------+
```

## License

MIT
