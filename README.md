# Project Rational

This project was created because I often have lots of ESP32 devices plugged in, and it gets hard to find them just by trial and error.

This project is a Go application that uses GoReleaser for building and releasing. It supports multiple operating systems and architectures.

## Installation

To install the dependencies, run:

```sh
go mod tidy
```

## Building

To build the project for all supported platforms, run:

```sh
goreleaser build --snapshot --rm-dist
```

### macOS Requirements

If you are building on macOS, you need to have the following installed:

```sh
brew tap SergioBenitez/osxct
brew install x86_64-unknown-linux-gnu
```

## Releasing

To create a release, follow these steps:

1. Create a new tag:

    ```sh
    git tag -a v1.0.0 -m "Release version 1.0.0"
    ```

2. Push the tag to the remote repository:

    ```sh
    git push origin v1.0.0
    ```

3. Run GoReleaser:

    ```sh
    goreleaser release --rm-dist
    ```

## Configuration

The build and release process is configured using the `.goreleaser.yaml` file. Below is a brief overview of the configuration:

- **Builds**: Configures the builds for Linux, Darwin (macOS), and Windows.
- **Archives**: Defines the archive format and naming template.
- **Changelog**: Configures the changelog generation.
- **Release**: Adds a footer to the release notes.

For more details, refer to the [GoReleaser documentation](https://goreleaser.com).

## Contributing

Feel free to open issues or submit pull requests if you find any bugs or have suggestions for improvements.

## License

This project is licensed under the MIT License.