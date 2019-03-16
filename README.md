[![Build Status](https://travis-ci.org/mchirico/go_slicestore.svg?branch=master)](https://travis-ci.org/mchirico/go_slicestore)
[![Maintainability](https://api.codeclimate.com/v1/badges/fc67d1c49ea5ee570bb0/maintainability)](https://codeclimate.com/github/mchirico/go_slicestore/maintainability)
[![codecov](https://codecov.io/gh/mchirico/go_slicestore/branch/master/graph/badge.svg)](https://codecov.io/gh/mchirico/go_slicestore)
# go_slicestore
Pull Data from Slice Store

[IBM API Reference](https://www.ibm.com/support/knowledgecenter/STXNRM_3.14.3/coss.doc/managerapi834a.html#managerapi-gen833)





## Build with Vendor

```bash

git clone https://github.com/mchirico/go_slicestore.git
cd go_slicestore

export GO111MODULE=on
go mod init
# Below will put all packages in a vendor folder
go mod vendor



go test -v -mod=vendor ./...

# Don't forget the "." in "./cmd/getdata" below
go build -v -mod=vendor ./cmd/getdata

# You can now run

./getdata

```