# gb2260

[![Build Status](https://img.shields.io/travis/cn/gb2260.svg?style=flat)](https://travis-ci.org/cn/gb2260)


go implement https://github.com/cn/GB2260.go


## Installation

Get the code:

```go
go get github.com/cn/GB2260.go
```

## Usage

```go
    
import (
    gb2260 "github.com/cn/GB2260.go"
)

gb := gb2260.NewGB2260("")
division := gb.Get("360426")

```




## Spec:

https://github.com/cn/GB2260/blob/api-design/spec.md
