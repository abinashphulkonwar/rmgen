# rmgen - Random Data Generator CLI Tool

## Overview

**_rmgen_** is a command-line interface (CLI) tool written in Golang for generating random data based on a JSON configuration file. It's designed to make it easy for developers to quickly generate sample data for testing and development purposes.

## Features

- Simple and intuitive JSON configuration for defining data fields and types.
- CSV output for easy integration with various applications and tools.
- Customizable number of records to generate.
- Supports various data types, such as email, name, etc.

## Installation

To use **rmgen**, you need to have Golang installed. If you don't have Golang installed, you can download it here.

1. Clone the repository:

```bash
git clone https://github.com/abinashphulkonwar/rmgen
cd rmgen
```

2. Build the executable:

```bash
go build
```

3. Run the tool:

```bash
./rmgen.exe -f path/to/your/config.json -c 1000
```

## Configuration

Create a JSON configuration file to specify the fields and data types to generate. Example:

```json
[
  {
    "field_name": "user email",
    "type": "email"
  },
  {
    "field_name": "user name",
    "type": "name"
  }
]
```

## Usage

```bash
./rmgen.exe -f path/to/your/config.json -c 1000
```

- -f: Specify the path to your JSON configuration file.
- -c: Specify the number of records to generate.

## Examples

Generate 1000 records based on the provided configuration:

```bash
./rmgen.exe -f path/to/your/config.json -c 1000
```

## Contributing

Feel free to contribute by opening issues or pull requests. Any feedback or suggestions are highly appreciated.

## License

This project is licensed under the MIT License - see the [LICENSE](https://github.com/your-username/rmgen/blob/main/LICENSE) file for details.
