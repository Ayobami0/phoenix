# Phoenix

A system configuration management tool that uses "ash" config files to setup systems.

## Overview

Phoenix allows you to define your system configuration in "ash" (a derivative of YAML, you could also use YAML files) files, which can then be applied directly to your system or converted into portable shell scripts. This makes it perfect for system setup, configuration management, and creating reproducible environments.

## Installation

```bash
# Clone the repository
git clone https://github.com/Ayobami0/phoenix.git

# Build the project
cd phoenix
go build

# Install the binary (optional)
sudo mv phoenix /usr/local/bin/
```

## Usage

Phoenix provides two main commands for working with ash configuration files:

### Rise

Apply configuration from an ash file to the current system:

```bash
phoenix rise [OPTIONS] <ashname>
```

Options:
- `--silent, -s`: Run in silent mode without printing status messages
- `--exclude`: Skip specific components (comma-separated list)

Example:
```bash
# Apply the developer workstation configuration
phoenix rise dev-workstation.ash

# Apply configuration but skip the docker component
phoenix rise --exclude docker server.ash
```

### Spawn

Generate a portable shell script from an ash file:

```bash
phoenix spawn [OPTIONS] <ashname>
```

Options:
- `--silent, -s`: Generate script without printing status messages
- `--out, -o`: Write output to specified file (default is the ash file name, use "-" for stdout)
- `--executable`: Make the output file executable (chmod +x)
- `--compress`: Generate a compressed, minimal script
- `--shell`: Specify output shell type (bash, zsh, powershell, default is bash)
- `--exclude`: Skip specific components (comma-separated list)

Example:
```bash
# Generate an executable bash script
phoenix spawn --executable web-server.ash

# Generate a compressed zsh script and output to a specific file
phoenix spawn --compress --shell=zsh --out setup.sh arch-setup.ash

# Generate a script and print to stdout
phoenix spawn --out - minimal.ash
```

## Ash Configuration Files

Ash files define the desired state of your system. Example ash files are included in the repository under the `examples/` directory for reference.

## Version

Current version: 1.0.0

## License

Phoenix CLI is open-source and available under the [MIT License](./LICENSE)

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add some amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

## Features and Improvements
- [] Storing existing files and directories before changes are made
- [] Caching state to allow for recall if an error occurs
- [] Possibility of rollback to previous state
