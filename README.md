# genurl

`genurl` is a powerful and efficient tool for generating URL combinations with customizable options. Designed for use in brute-force testing and domain discovery, `genurl` allows you to generate combinations of alphabetic and alphanumeric strings to replace placeholders in specified domains.

## Features

- **Customizable Combinations**: Generate combinations of alphabetic or alphanumeric strings.
- **Flexible Length**: Specify the length of the combinations to be generated.
- **Multiple Domains**: Accepts a single domain or a list of domains from a file.
- **Output Options**: Save results to a specified output file or display them in the terminal.
- **Silent Mode**: Run the tool in silent mode to suppress output in the terminal.

## Usage

### Command Line Options

- `-a` : Generate alphabetic combinations.
- `-ad` : Generate alphabetic and numeric combinations.
- `-l` : Length of combinations (default: 6).
- `-d` : Single domain with placeholder.
- `-f` : File containing list of domains with placeholders.
- `-o` : Output file name (default: output.txt).
- `-silent` : Run the program in silent mode.

### Example Commands

```bash
# Generate alphabetic combinations of length 6 with a single domain
./genurl -a -l 6 -d "production-[here].test.com" -o myoutput.txt

# Generate alphabetic and numeric combinations of length 8 with domains from a file
./genurl -ad -l 8 -f domains.txt -o myoutput.txt

# Run in silent mode
./genurl -a -l 6 -d "production-[here].test.com" -silent
```

## Installation

1. Clone the repository:
   ```bash
   git clone https://github.com/yourusername/genurl.git
   ```

2. Build the tool:
   ```bash
   cd genurl
   go build generate_urls.go
   ```

3. Run the tool:
   ```bash
   ./genurl [options]
   ```
