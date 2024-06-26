#!/usr/bin/env bash
set -euo pipefail

echo "================================================================================"
echo "Clone repo: ${CLONE_REPO} @ ${PKG_VER} to /tmp/__BUILD__"
echo "================================================================================"

# shellcheck disable=2154
git clone "${CLONE_REPO}" --branch "${PKG_VER}" --single-branch /tmp/__BUILD__
cd /tmp/__BUILD__

# linux/amd64 (Intel64)
GOARCH=amd64 go build -ldflags "-w -s" -trimpath -o /usr/local/bin/iam-policy-json-to-terraform_amd64 .

# linux/arm64 (ARM64)
GOARCH=arm64 go build -ldflags "-w -s" -trimpath -o /usr/local/bin/iam-policy-json-to-terraform_arm64 .
