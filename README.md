## Pre-requisite

This project requires Golang to be installed - https://go.dev/doc/install

## Description

`HIPAA-Verifier` is responsible for the following - 

* Act as a registry for _services_ that exchange data, described by - 
  * a `source` or the `data owner` - who is sending the data ?
  * a `destination` or the `data requester` - who is requesting the data ?
  * a `type` - the record type that is being exchanged
  * a list of `protect-rule`s (defined ahead) 
  * a list of `verify-rule`s (defined ahead)

* Define a `identify` path and `protect-rule`/`verify-rule` corresponding to this field -
  * a `protect-rule` defines what logic/function is to be used to _protect_ a certain field's value
  * a `verify-rule` defines what logic/function is to be used to _verify_ that a certain field's value is protected

## Current Implementation

* registration is part of the `hipaa/systems.go` file as a hard-coded value.
* rule definitions are defined in the `hipaa/rule.go` file as a map of specific PHI field and corresponding rule
* actual implementations of protect and verify of data are in `hipaa/protect.go`

## Running 

1. Download dependencies -
```shell
make deps
```

2. Run the server -

```shell
make run
```

This will bring up a webserver at `http://localhost:4000/` having the following 2 endpoints - 

`/protect` - returns protected data after applying all `protect-rules` based on incoming request data
`/verify` - returns a true/false (boolean) based on whether the incoming request data passes verification rules or not.

The other services will call these endpoints to protect and check their compliance.
