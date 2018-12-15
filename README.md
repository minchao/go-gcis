# go-gcis

[![Build Status](https://travis-ci.com/minchao/go-gcis.svg?branch=master)](https://travis-ci.com/minchao/go-gcis)

go-gcis is a Go client library for accessing the [GCIS API](https://data.gcis.nat.gov.tw).

## Getting started

Use `go get` to download the library into your $GOPATH.

```bash
go get -u https://github.com/minchao/go-gcis
```

### Usage

This example shows how to make a request to the GCIS API.

```go
package main

import (
	"context"
	"fmt"

	"github.com/minchao/go-gcis/gcis"
)

func main() {
	client := gcis.NewClient()

	info, _, err := client.Company.GetBasicInformation(context.Background(),
		&gcis.CompanyBasicInformationInput{BusinessAccountingNO: "20828393"})
	if err != nil {
		panic("failed to get company basic information, " + err.Error())
	}

	fmt.Println("resp", info)
}
```

## License

This library is distributed under the MIT license found in the [LICENSE](./LICENSE) file.
