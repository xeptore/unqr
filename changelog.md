# Changelog

This project uses simple versioning: `v1`, `v2`, `v3`, and so on.

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
