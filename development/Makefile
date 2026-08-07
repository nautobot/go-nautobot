.PHONY: all ci get-api update-version create-bindings

all: get-api update-version create-bindings

# Validation target for CI. Deliberately omits update-version: bumping the
# client version is a release action, and it rewrites api/nautobot_version,
# which would make a "regenerated output matches the committed tree" check fail
# on every run. This regenerates against the committed version instead, so the
# result is comparable to what is checked in.
ci: get-api create-bindings

get-api:
	cd get-api && ./get-api.sh

update-version:
	cd update-version && ./update-version.sh

create-bindings:
	cd create-bindings && ./create-bindings.sh
