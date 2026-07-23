# Setup

To Create the go-bindings manually, you need to install:
- `docker-compose`
- `docker`
- `git`
- `make`

Then run:

- `make` in this directory to get the openapi spec and generate the go bindings.
- `make get-api` in this directory to only get the openapi spec.
- `make create-bindings` in this directory to only generate the go bindings using the exising openapi.yaml

You will need root permissions to create the docker containers.
