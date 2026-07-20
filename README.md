# PW - CLI Password Tool

## About
A CLI tool to create random passwords and check password strength.

## Motivation
I needed a quick and reliable way to generate random passwords. By building my own tool, I could fully trust the integrity and encryption of the generated passwords.

## Quick Start
1. Install Go (if not already installed):
```bash
sudo apt install go
```
2. Clone the repository:
```bash
git clone https://github.com/hendrik-ro/crypto-pw.git
```
3. Navigate to the project directory:
```bash
cd crypto-pw
```
4. Build the project:
```bash
go build
```
5. Run the tool:
```bash
./pw
```

## Features
- **generate [n]**: Generate a random password of length n (default 15).
- **check**: Check password strength against common decryption methods.

## Usage
```bash
$ go run main.go
launching pw...
type 'help' for available commands
PW >
```

## Contributing
If you'd like to contribute, please fork the repository and open a pull request to the `main` branch.

## Prerequisites
- **Go** 1.21+

## License
MIT License - See [LICENSE.md](LICENSE.md) for details.
