# go
[go lang](https://go.dev/doc/install)

## Install Go on Linux

1. Download it

```
get https://go.dev/dl/go1.19.linux-amd64.tar.gz
```

2. Remove previous version and install

```
rm -rf /usr/local/go && tar -C /usr/local -xzf go1.19.linux-amd64.tar.gz
```

3. Add /usr/local/go/bin to the PATH environment variable

```
export PATH=$PATH:/usr/local/go/bin
```

## check go version

Check go version by following command -

```
>> go version
go version go1.19 linux/amd64
```
