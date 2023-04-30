# Jakonda - CLI Tools for managing collections of photos

Jakonda is a Golang package that provides a set of command-line tools for managing photos.  
With Jakonda, you can quickly and easily perform various operations on your photo collections.

## Features

- Print a tree of your photo collections
- Remove **RAW** files for which there is **JPEG** file
- Stay tuned, more to be added!

## Installation

To use Jakonda, you'll need to have [Golang installed](https://go.dev/doc/install) on your system.  
Once you have Golang installed, you can install Jakonda using the following command:

```sh
go install github.com/astappiev/jakonda
```

## Usage

To use Jakonda, simply run the `jakonda` command followed by the operation you want to perform and any necessary options.  
For example, to sort photos in a directory by date, you would run the following command:

```sh
jakonda tree /path/to/photos
```

For a full list of available operations and options, run `jakonda help`.

## Contributing

Contributions are always welcome! If you'd like to contribute to Jakonda, please follow these steps:

1. Fork the Jakonda repository
2. Create a new branch for your changes
3. Make your changes and test them thoroughly
4. Submit a pull request to the main Jakonda repository

## License

Jakonda is licensed under the MIT License. See the [LICENSE](LICENSE) file for more information.
