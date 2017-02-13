# gb2260

[![Build Status](https://img.shields.io/travis/damonchen/gb2260.svg?style=flat)](https://travis-ci.org/damonchen/gb2260)


go implement https://github.com/cn/GB2260.go


## Installation

Get the code:

```go
go get github.com/cn/GB2260.go
```
or
```go
go get github.com/damonchen/gb2260
```

## Usage

```go
    
import (
    gb2260 "github.com/cn/GB2260.go"
    // or
    // gb2260 "github.com/damonchen/gb2660"
)

gb := gb2260.NewGB2260("")
division := gb.Get("360426")

```




## Spec:

https://github.com/cn/GB2260/blob/develop/spec.md
