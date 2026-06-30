# unqr

Minimal CLI to decode QR codes from image files and print the text to stdout.

## Install

Run without installing:

```bash
go run github.com/xeptore/unqr@latest path/to/qrcode.png
```

Or by installing it:

```bash
go install github.com/xeptore/unqr@latest
```

Pre-built binaries are available on the [Releases](https://github.com/xeptore/unqr/releases) page.

## Usage

```bash
unqr path/to/qrcode.png
```

Supports JPEG and PNG images. On success, the decoded text is printed to stdout. Errors are written to stderr.
