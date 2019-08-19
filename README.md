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

## Update Divisions
```bash
# example
cd cmd
go run generate.go
cp ../data.go ~/projects/mygoproject/data.go
# Modify the package name specified for you.
```
```go
//  example data.go
package mygoproject

import (
    gb2260 "github.com/cn/GB2260.go"
)

func init() {
	gb2260.Divisions = map[string]map[string]string{
		"example_revision": map[string]string{"Code": "Name",
			"110000": "北京市",
			"110101": "东城区",
			"110102": "西城区",
			"110105": "朝阳区",
			"110106": "丰台区",
			"110107": "石景山区",
			"110108": "海淀区",
			"110109": "门头沟区",
		},
	}
}
// example main.go
func main() {
    gb := gb2260.NewGB2260("example_revision")
    division := gb.Get("360426")
}
```

## Spec:

https://github.com/cn/GB2260/blob/develop/spec.md
