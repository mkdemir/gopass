# GOPass

`gopass` is a command-line tool designed to generate secure passwords according to user-specified criteria. This utility allows users to customize their passwords with options for length, digits, and symbols.

## Features

- Customize password length
- Option to include digits in the password
- Option to include symbols in the password

## Installation

Ensure you have Go installed on your system (Go 1.15 or higher recommended) before using this tool.

### Installing from Source

Follow these steps to install the project locally:

```bash
git clone https://github.com/yourusername/gopass.git
cd gopass
go build -o gopass
```

These commands will clone the project and build an executable file.

## Usage

To use the `gopass` application with different options, execute the following commands in your terminal:

```bash
./gopass --length 12 --digits --symbols
```

This example will generate a 12-character long password that includes digits and symbols.

### Parameters

- `--length, -l`: Specifies the length of the password (default is 8 characters).
- `--digits, -d`: Include digits in the password (default is false).
- `--symbols, -s`: Include symbols in the password (default is false).

## Contributing

Contributions are welcome! Please feel free to submit pull requests or open issues to discuss proposed changes or enhancements.

## License

Specify your license here, or if you have not yet chosen a license, it's a good idea to include a note that the project is currently unlicensed.
