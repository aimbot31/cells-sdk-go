# Cells API client

**Warning, this is a Work In Progress**.

Rest API Client for Pydio Cells.

Current SDK has been update on **Aug, the 30th 2018** with git commit **[406a82d](https://github.com/pydio/cells/commit/406a82dbe316cbb8c55f740b2b109c5c6a902fb6)**.

## Overview

This API client was generated by the [go-swagger](https://github.com/go-swagger/go-swagger) project.

For more information, please visit [https://pydio.com](https://pydio.com)

## Installation

Put the package under your project folder and add the following in import:

```go
    "github.com/pydio/cells-sdk-go"
```

## Update SDK

_This is for the maintainers of cells-sdk-go project only._

To update the Cells go sdk, you must follow the steps below:

_First time only_:

```sh
# go to the roots of this directory, typically:
cd $GOPATH/src/github.com/pydio/cells-sdk-go

# If necessary, retrieve swagger binary, rename it and give execution permission

## for linux
wget https://github.com/go-swagger/go-swagger/releases/download/0.16.0/swagger_linux_amd64

## for Mac OS
wget https://github.com/go-swagger/go-swagger/releases/download/0.16.0/swagger_darwin_amd64

mv swagger_linux_amd64 swagger
chmod u+x swagger
```

_After each API Spec modification_:

```sh
# retrieve latest spec file
wget https://raw.githubusercontent.com/pydio/cells/master/common/proto/rest/rest.swagger.json
# as the time of writing, we rather use the corresponding branch in Cells
wget https://raw.githubusercontent.com/pydio/cells/gokilledphpstars/common/proto/rest/rest.swagger.json


# You might also delete folder models and client
rm -r client
rm -r models

# simply generate updated code
./swagger generate client --skip-validation -f rest.swagger.json
```

> _WARNING_: we currently have some glitches during import process and must do some extra twick for the go code to compile. See below.

```sh
# Apply the tweak to workaround int64 serialisation issue between protobuf and swagger
go run build/main.go tweak-model
```

You should also update version information at the top of this page.

_**NOTE**: we use the --skip-validation flag to avoid circular issues with object that make reference to same type of objects, typically activities and jobs. See [issue #957 in go-swagger repository](https://github.com/go-swagger/go-swagger/issues/957) for more details._
