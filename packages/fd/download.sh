#!/usr/bin/env bash
set -euo pipefail

# linux/amd64 (Intel64) glibc
download-asset get \
    --owner-repo "sharkdp/fd" \
    --pattern "fd-v{{.Ver}}-x86_64-unknown-linux-gnu.{{.Ext}}" \
    --archive-path "fd-v{{.Ver}}-x86_64-unknown-linux-gnu/fd" \
    --write-to-bin "fd_glibc_amd64" \
    ;

# linux/arm64 (ARM64) glibc
download-asset get \
    --owner-repo "sharkdp/fd" \
    --pattern "fd-v{{.Ver}}-aarch64-unknown-linux-gnu.{{.Ext}}" \
    --archive-path "fd-v{{.Ver}}-aarch64-unknown-linux-gnu/fd" \
    --write-to-bin "fd_glibc_arm64" \
    ;

# Already exists in the Alpine Linux repositories.
