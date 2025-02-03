# RIPE DB Go Client
This is a Go client for the RIPE Database RESTful Web Service API. It is a simple wrapper around the API, providing a more convenient way to interact with the RIPE Database.

## Features

- Generic resource queries
- Create/Update/Delete operations on single resources
- Authentification schemes
  - Anonymous
  - Password & [API Key](https://docs.db.ripe.net/23.Appendices/11-Appendix-K--API-Keys.html)
  - [X.509 Client Certificate](https://docs.db.ripe.net/Appendices/Appendix-I--Client-Certificate-Authentication/)

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

### Authentication

The CLI supports the following authentication schemes:

- Anonymous
- Password

#### Anonymous

The anonymous authentication scheme is the default one. It does not require any credentials.

#### Password & API Key

The password authentication scheme requires a username and a password. The username is the RIPE Database user handle.

```bash
$ ripedb --user <username> --password <password> get organisation ORG-CEOf1-RIPE
```

Alternatively, you can set the `RIPEDB_USER` and `RIPEDB_PASSWORD` environment variables.

> [!CAUTION]
> It is possible to provider the password without the username, but it is not recommended.
> This will pass the password as a query parameter in the URL (instead of the Authorization header).

#### X.509 Client Certificate

The X.509 client certificate authentication scheme requires a client certificate and a private key.

```bash
$ ripedb --cert <cert> --key <key> get organisation ORG-CEOf1-RIPE
```

Alternatively, you can set the `RIPEDB_CERTFILE` and `RIPEDB_KEYFILE` environment variables.

Please refer to the [Appendix I - Client Certificate Authentication](https://docs.db.ripe.net/Appendices/Appendix-I--Client-Certificate-Authentication/) for more information on how to generate a client certificate.

### Example
```bash
$ ripedb get organisation ORG-CEOf1-RIPE
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

## License

Copyright (c) The RIPE DB Go Client Authors. [Apache 2.0 License](./LICENSE).
