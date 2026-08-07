#!/bin/bash
set -euxo pipefail

echo "Updating go-nautobot version"

VERSION_FILE="../../api/nautobot_version"

# Sourcing local_dev.env would overwrite a NAUTOBOT_VER supplied by the
# environment (CI, or the release workflow's `tag` input), so remember any
# externally provided value and restore it afterwards. The file remains the
# default when nothing is set externally. This keeps this step consistent with
# get-api and create-bindings, where the shell environment already takes
# precedence over --env-file during docker-compose interpolation.
EXTERNAL_NAUTOBOT_VER="${NAUTOBOT_VER:-}"
. "../local_dev.env"
if [ -n "$EXTERNAL_NAUTOBOT_VER" ]; then
    NAUTOBOT_VER="$EXTERNAL_NAUTOBOT_VER"
fi

CURRENT_VERSION=$(head -n 1 $VERSION_FILE)
CURRENT_MAJOR_MINOR_VER=${CURRENT_VERSION%.*}

# Using only major and minor version to get the api_version
# 1.3.3 -> 1.3
# 1.3.7 -> 1.3
# 1.4.0 -> 1.4
MAJOR_MINOR_VER=${NAUTOBOT_VER%.*}


if [ "$CURRENT_MAJOR_MINOR_VER" = "$MAJOR_MINOR_VER" ]; then
    # Get the Patch version string
    NEW_PATCH_VERSION=$(echo $CURRENT_VERSION | awk -F '.' '{ print $3;}')
    # Remove suffixes in version
    NEW_PATCH_VERSION=${NEW_PATCH_VERSION%-*}
    # Increment Patch version string with 1
    NEW_PATCH_VERSION=$((${NEW_PATCH_VERSION} + 1))
    NEW_TAG=${CURRENT_MAJOR_MINOR_VER}.$NEW_PATCH_VERSION
else
    NEW_TAG=${MAJOR_MINOR_VER}.0
fi

# TODO: remove beta when it's in production
FINAL_NEW_TAG=${NEW_TAG}-beta

echo $FINAL_NEW_TAG > "$VERSION_FILE"

echo "go-nautobot client version updated to $FINAL_NEW_TAG"
