 # aro-ado-work-items

 * This project uses an Azure Devops Golang RESTful client to simplify creating work items. The point is to be able to create work items efficiently without having to battle against a clunky user interface.

## Setup

* Create a folder called secrets

```
mkdir secrets
```

* Put your ADO PAT token in the `secrets/env` file

```
cat > secrets/env <<'EOF'
ADO_PAT=<your token here>
EOF
```

* `ADO_PAT` is your personal access token from Azure DevOps. You can use this anonymously, but stricter rate limits apply.

* Source your `./env` file

```
. ./env
```

## Build

```
make build
```

## Run

```
make run
```

## Parameters


## Example

* Make run

```
make run
```

* Go Run

```
go run . 
```