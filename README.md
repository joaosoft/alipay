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

## Setup
Copy default configs
```
make setup
```

## Dependency Management
>### Dep

Project dependencies are managed using Dep. Read more about [Dep](https://github.com/golang/dep).
* Install dependencies: `dep ensure`
* Update dependencies: `dep ensure -update`


>### Go
```
go get github.com/joaosoft/alipay
```

### Code
```go
import (
	"github.com/joaosoft/alipay"
)

func main() {
    ap, err := alipay.NewAliPay()
    if err != nil {
        panic(err)
    }
    
    if err = ap.Start(); err != nil {
        panic(err)
    }
}

```

## Known issues

## Follow me at
Facebook: https://www.facebook.com/joaosoft

LinkedIn: https://www.linkedin.com/in/jo%C3%A3o-ribeiro-b2775438/

##### If you have something to add, please let me know joaosoft@gmail.com