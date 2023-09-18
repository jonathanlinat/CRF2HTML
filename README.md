# CRF2HTML

## Introduction

CRF2HTML is a command-line utility inspired by the Thief series of video games, including Thief: The Dark Project (Thief Gold) and Thief II: The Metal Age. In these classic games, textures and images were stored in proprietary formats like CRF and PCX. CRF2HTML aims to bring a piece of that nostalgic world to modern web development.

The program is designed to generate an HTML page that beautifully showcases the textures found in Thief series CRF files and other image formats (PCX, GIF, PNG, and JPG). It seamlessly resizes and encodes these textures as base64, making it easy to embed them in an organized HTML page.

Whether you're a fan of the Thief series or simply interested in working with these classic texture formats, CRF2HTML provides a convenient way to create galleries and showcases of these vintage textures for various creative and nostalgic purposes.

> **Based and inspired by** https://github.com/vfig/thieftextures

## Features

- Converts Thief series textures and other image formats to base64 for easy embedding in an HTML page.
- Organizes textures by family and sorts them alphabetically by caption.
- Allows customization of the page title, so you can give your texture showcase a unique name.
- Supports various image formats, including PCX, GIF, PNG, and JPG.
- Provides an option to specify a custom background color for the images.
- Generates a clean and organized HTML page with embedded textures, reminiscent of the Thief games' texture menus.

## Prerequisites

Before using CRF2HTML, make sure you have the following prerequisites installed on your system:

- Go programming language (if you need to compile the program)
- ZIP file utilities (if you plan to use ZIP files as a data source)

## Installation

### Compile from Source

If you prefer to compile the program yourself, follow these steps:

1. Clone this repository to your local machine:

   ```bash
   git clone https://github.com/jonathanlinat/CRF2HTML.git
   ```

2. Change to the project directory:

   ```bash
   cd CRF2HTML
   ```

3. Build the program:

   ```bash
   go build main.go
   ```

This will generate an executable binary named `main` in the project directory.

## Usage

To use CRF2HTML, follow the usage instructions below:

```bash
./main source_path output_path [-title "Page Title"]
```

- `source_path`: Path to the directory containing image files or a CRF/ZIP file.
- `output_path`: Path to the HTML file to be generated.
- `-title "Page Title"` (optional): Custom title for the HTML page. If not provided, the default title is "Textures."

### Example

Here's an example of how to use CRF2HTML to create an HTML page:

```bash
./main ./fam.crf ./textures.html -title "My Custom Title"
```

This command will generate an HTML page named `textures.html` in the current directory, showcasing the image textures from the `./fam.crf` source, with the custom title "My Custom Title."

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.

## Acknowledgments

CRF2HTML uses the following third-party Go packages:

- [nfnt/resize](https://github.com/nfnt/resize) for image resizing.
- [samuel/go-pcx/pcx](https://github.com/samuel/go-pcx/pcx) for PCX image format support.

---

Feel free to contribute to this project, report issues, or suggest improvements!
