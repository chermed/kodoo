# Installation

## Linux

```
curl -Lo kodoo https://github.com/chermed/kodoo/releases/download/{{version}}/kodoo-linux-amd64-{{version}} && chmod +x kodoo && sudo mv kodoo /usr/local/bin
```

## macOS

```
curl -Lo kodoo https://github.com/chermed/kodoo/releases/download/{{version}}/kodoo-darwin-amd64-{{version}} && chmod +x kodoo && sudo mv kodoo /usr/local/bin
```

## Windows

```
https://github.com/chermed/kodoo/releases/download/{{version}}/kodoo-windows-amd64-{{version}}.exe
```

## Docker image

`chermed/kodoo:{{version}}`
```
docker run -it --rm -v $(pwd):/.kodoo --net host chermed/kodoo:{{version}} init-config

docker run -it --rm -v $(pwd):/.kodoo --net host chermed/kodoo:{{version}}
```


