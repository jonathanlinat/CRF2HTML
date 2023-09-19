# crf2html

[![Build and Release](https://github.com/jonathanlinat/crf2html/actions/workflows/build-and-release.yml/badge.svg)](https://github.com/jonathanlinat/crf2html/actions/workflows/build-and-release.yml)

## Introduction

The `crf2html` program generates an HTML page that displays image textures from a specified directory or a CRF/ZIP file. The tool resizes the images, encodes them as base64, and organizes them into a neat HTML page.

## Features

- Read image files from both directories and CRF/ZIP files.
- Supports multiple image formats including `.pcx`, `.gif`, `.png`, `.jpg`, and `.tga`.
- Ability to resize images and encode them as base64 for inline embedding in HTML.
- Organizes images by families, based on their directory or path structure.
- Easily customizable output through command-line arguments.

## Installation

### Option 1: Precompiled Binary

You can download a precompiled binary for your platform from the [releases](https://github.com/jonathanlinat/crf2html/releases) section of this repository.

### Option 2: Compile from Source

> **Important**
>
> It is recommended to install and use at least **Go v1.21**. Here are the [corresponding instructions](https://go.dev/doc/install).

---

If you prefer to compile the program yourself, follow these steps:

1. Clone this repository to your local machine:

   ```bash
   git clone https://github.com/jonathanlinat/crf2html.git
   ```

2. Change to the project directory:

   ```bash
   cd crf2html
   ```

3. Build the program:

   ```bash
   go build -o crf2html main.go
   ```

This will generate a binary, `crf2html` or `crf2html.exe`, in the project directory.

## Usage

- `source_path`: Path to the directory containing image files or a CRF/ZIP file.
- `output_path`: Path to the HTML file to be generated.
- `-title "Page Title"` (optional): Custom title for the HTML page. If not provided, the default title is `Textures`.
- `-size 64` (optional): Custom thumbnail size for the HTML page. If not provided, the default size is `128`.

### Linux

```bash
./crf2html source_path output_path [-title "Page Title"]
```

### Windows

```bash
crf2html.exe source_path output_path [-title "Page Title"]
```

### Example

Here's an example of how to use `crf2html` to create an HTML page:

#### Linux

```bash
./crf2html ./fam.crf ./textures.html -title "My Custom Title" -size 64
```

#### Windows

```bash
crf2html.exe C:\path\to\source\fam.crf C:\path\to\output\textures.html -title "My Custom Title" -size 64
```

This command will generate an HTML page named `textures.html` in the current directory, showcasing the image textures from the `fam.crf` source, with the custom title `My Custom Title`.

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.

## Acknowledgments

`crf2html` uses the following third-party Go packages:

- [nfnt/resize](https://github.com/nfnt/resize) for image resizing.
- [samuel/go-pcx/pcx](https://github.com/samuel/go-pcx/pcx) for PCX image format support.

---

Feel free to contribute to this project, report issues, or suggest improvements!
