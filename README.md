# RIPE Databse Go Client
This is a Go client for the RIPE Database RESTful Web Service API. It is a simple wrapper around the API, providing a more convenient way to interact with the RIPE Database.

## Usage
```bash
ripedb <resource> <key>
```

## Example
```bash
$ ripedb organisation ORG-CEOf1-RIPE
# organisation:       ORG-CEOf1-RIPE
# org-name:           CERN - European Organization for Nuclear Research
# country:            CH
# org-type:           LIR
# address:            CERN
# address:            CH-1211
# address:            Geneva 23
# address:            SWITZERLAND
# phone:              +41 22 76 72613
# admin-c:            EM1969
# admin-c:            DGR6-RIPE
# tech-c:             CERN513
# abuse-c:            CERN513
# mnt-ref:            RIPE-NCC-HM-MNT
# mnt-ref:            CERN-MNT
# mnt-by:             RIPE-NCC-HM-MNT
# mnt-by:             CERN-MNT
# created:            2004-04-17T11:02:00Z
# last-modified:      2021-06-29T07:00:23Z
# source:             RIPE
```
