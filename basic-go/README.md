# go
[go lang](https://go.dev/doc/install)

[go lang doc](https://go.dev/ref/spec)

[Followed](https://www.youtube.com/playlist?list=PLRAV69dS1uWQGDQoBYMZWKjzuhCaOnBpa)
## Install Go on Linux

1. Download it

```
wget https://go.dev/dl/go1.19.linux-amd64.tar.gz
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

# Updated Go on MacBook Pro
From `17-files` I have updated go version to `go version go1.19.4 darwin/arm64`.

Watch Installation [Video](https://youtu.be/3u6pZkNRCXg)

## Steps
1. search for install [go](https://go.dev/doc/install) and click on install. My Install link was `https://go.dev/dl/go1.19.4.darwin-amd64.pkg`

2. export GOROOT and PATH

``` 
export GOROOT=/usr/local/go
export PATH=/usr/local/go/bin:$PATH
```

