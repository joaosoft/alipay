AliPay
================

[![Build Status](https://travis-ci.org/joaosoft/alipay.svg?branch=master)](https://travis-ci.org/joaosoft/alipay) | [![codecov](https://codecov.io/gh/joaosoft/alipay/branch/master/graph/badge.svg)](https://codecov.io/gh/joaosoft/alipay) | [![Go Report Card](https://goreportcard.com/badge/github.com/joaosoft/alipay)](https://goreportcard.com/report/github.com/joaosoft/alipay) | [![GoDoc](https://godoc.org/github.com/joaosoft/alipay?status.svg)](https://godoc.org/github.com/joaosoft/alipay)

A simple client for stripe api to receive money from alipay.

## With support:

### Customer
* Create
  
### Payment
* Create 
* Confirm

###### If i miss something or you have something interesting, please be part of this project. Let me know! My contact is at the end.

## Dependency Management
>### Dep

Project dependencies are managed using Dep. Read more about [Dep](https://github.com/golang/dep).
* Install dependencies: `dep ensure`
* Update dependencies: `dep ensure -update`


>### Go
```
go get github.com/joaosoft/alipay
```

## Usage 
This examples are available in the project at [alipay/examples](https://github.com/joaosoft/alipay/tree/master/examples)

### Code
```go
import (
	"github.com/joaosoft/alipay"
)

func main() {
    if err := globals.AliPayInstance.Start(); err != nil {
        panic(err)
    }
}
```

## Known issues

## Follow me at
Facebook: https://www.facebook.com/joaosoft

LinkedIn: https://www.linkedin.com/in/jo%C3%A3o-ribeiro-b2775438/

##### If you have something to add, please let me know joaosoft@gmail.com
