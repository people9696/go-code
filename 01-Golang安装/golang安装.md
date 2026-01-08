# 下载

问Go官方网站 [golang.org](https://golang.org/dl/) 下载适合你操作系统的Go安装包。Go支持Windows、macOS和Linux系统。

使用 `go version`验证是否安装成功。


# 配置环境变量

Go需要配置两个重要的环境变量：GOROOT和GOPATH。

## Linux/macOS

```
# 在 ~/.bashrc 或 ~/.zshrc 中添加
export GOROOT=/usr/local/go
export GOPATH=$HOME/go
export PATH=$PATH:$GOROOT/bin:$GOPATH/bin
```

## Windows:

在系统环境变量中添加GOROOT和GOPATH
