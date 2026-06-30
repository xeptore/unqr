# Changelog

This project uses simple versioning: `v1`, `v2`, `v3`, and so on.

## v2

Changed:

- Release workflow now triggers only on simple version tags (`v1`, `v2`, …) instead of any `v*` tag
- Release archive names no longer include the version prefix (e.g. `unqr_linux_amd64.tar.gz` instead of `unqr_v1_linux_amd64.tar.gz`)

Maintenance:

- Bump `actions/download-artifact` from v7 to v8

## v1

Initial release including:

- Minimal CLI that decodes QR codes from image files and prints the text to stdout
- JPEG and PNG image support
- Error messages on stderr with exit codes for usage errors, file/decode failures, and close errors
- Cross-platform release binaries for common OS/arch combinations
- CI workflow that runs `go vet` and cross-compiles on every push and pull request
- Release workflow triggered by version tags
- README with install, `go run`, and usage instructions
- MIT license
