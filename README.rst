Contractor Client Libraries
===========================

Object/Service definitaions for go and python

Golang
------

import via the standard golang module method, adding to you imports::

  import (
  	contractor "github.com/t3kton/contractor_client/go"
  )


then tell go to download the library::

  go get github.com/t3kton/contractor_client/go



To Generate
-----------

../../cinp/utils/cinp-codegen -l go -d go/autogen -s contractor http://127.0.0.1:8888/api/v1/
cd go
golint ./...
go vet ./...
