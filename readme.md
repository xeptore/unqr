# unqr

Minimal CLI to decode QR codes from image files and print the text to stdout.

## Install

```bash
go install github.com/xeptore/unqr@latest
```

Or build from source:

```bash
go build -o unqr .
```

Pre-built binaries are available on the [Releases](https://github.com/xeptore/unqr/releases) page.

## Usage

```bash
unqr path/to/qrcode.png
```

Supports JPEG and PNG images. On success, the decoded text is printed to stdout. Errors are written to stderr.
