# RIPE Databse Go Client
This is a Go client for the RIPE Database RESTful Web Service API. It is a simple wrapper around the API, providing a more convenient way to interact with the RIPE Database.

## Features

- Generic resource queries
- Create/Update/Delete operations on single resources
- Authentification schemes
  - Anonymous
  - Password
  - (*Soon*) X.509

## Library
### Installation
```bash
go get github.com/frederic-arr/ripedb-go/cmd/ripedb@latest
```

### Usage

A basic Go program that fetches the CERN organisation from the RIPE Database and prints its name.

```go
package main

import (
	"fmt"
	"os"

	"github.com/frederic-arr/ripedb-go/ripedb"
)

func main() {
    client = ripedb.NewRipeAnonymousClient()
    org, err := client.GetOrganisation("ORG-CEOf1-RIPE")
    if err != nil {
        log.Fatal(err)
        os.Exit(1)
    }

    fmt.Println(org.OrgName)
}
```


## CLI
### Installation
```bash
go install github.com/frederic-arr/ripedb-go/cmd/ripedb@latest
```

### Usage
```bash
ripedb <resource> <key>
```

### Example
```bash
$ ripedb organisation ORG-CEOf1-RIPE
```

![Terminal screenshot of the output](./.github/assets/cern-dark.png#gh-dark-mode-only)
![Terminal screenshot of the output](./.github/assets/cern-light.png#gh-light-mode-only)

<details>

<summary>Text output</summary>

```
# This is the RIPE Database search service.
# The objects are in RPSL (RFC 2622) format.
# The RIPE Database is subject to Terms and Conditions.
organisation:  ORG-CEOf1-RIPE
org-name:      CERN - European Organization for Nuclear Research
country:       CH
org-type:      LIR
address:       CERN
address:       CH-1211
address:       Geneva 23
address:       SWITZERLAND
phone:         +41 22 76 72613
admin-c:       EM1969
admin-c:       DGR6-RIPE
tech-c:        CERN513
abuse-c:       CERN513
mnt-ref:       RIPE-NCC-HM-MNT
mnt-ref:       CERN-MNT
mnt-by:        RIPE-NCC-HM-MNT
mnt-by:        CERN-MNT
created:       2004-04-17T11:02:00Z
last-modified: 2021-06-29T07:00:23Z
source:        RIPE # Filtered
```

</details>
