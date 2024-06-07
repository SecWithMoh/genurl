package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

// Function to generate combinations and process them on the fly
func generateCombinations(chars []rune, length int, process func(string)) {
	var generate func(prefix string, length int)
	generate = func(prefix string, length int) {
		if length == 0 {
			process(prefix)
			return
		}
		for _, ch := range chars {
			generate(prefix+string(ch), length-1)
		}
	}
	generate("", length)
}

// Function to read domains from a file
func readDomainsFromFile(filename string) ([]string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var domains []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		domains = append(domains, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return domains, nil
}

// Function to log errors to a file
func logError(err error) {
	f, ferr := os.OpenFile("error.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if ferr != nil {
		fmt.Println("Error opening error log file:", ferr)
		return
	}
	defer f.Close()

	logger := bufio.NewWriter(f)
	timestamp := time.Now().Format("2006-01-02 15:04:05")
	logger.WriteString(fmt.Sprintf("%s: %v\n", timestamp, err))
	logger.Flush()
}

func main() {
	// Define command-line flags
	alphabetic := flag.Bool("a", false, "Generate alphabetic combinations")
	alphanumeric := flag.Bool("ad", false, "Generate alphabetic and numeric combinations")
	lengthFlag := flag.String("l", "6", "Length of combinations")
	domain := flag.String("d", "", "Single domain with placeholder")
	file := flag.String("f", "", "File containing list of domains with placeholders")
	output := flag.String("o", "output.txt", "Output file name")
	silent := flag.Bool("silent", false, "Run program in silent mode")

	// Parse command-line flags
	flag.Parse()

	// Check if silent mode is enabled
	if *silent {
		// Redirect stdout and stderr to /dev/null
		os.Stdout, _ = os.Open(os.DevNull)
		os.Stderr, _ = os.Open(os.DevNull)
	}

	// Check if no combination generation flag was provided
	if !*alphabetic && !*alphanumeric {
		fmt.Println("Usage:")
		fmt.Println("  -a   Generate alphabetic combinations")
		fmt.Println("  -ad  Generate alphabetic and numeric combinations")
		fmt.Println("  -l   Length of combinations (default: 6)")
		fmt.Println("  -d   Single domain with placeholder")
		fmt.Println("  -f   File containing list of domains with placeholders")
		fmt.Println("  -o   Output file name (default: output.txt)")
		fmt.Println("  -silent  Run program in silent mode")
		os.Exit(1)
	}

	// Check if no domain or file flag was provided
	if *domain == "" && *file == "" {
		fmt.Println("You must specify either a single domain (-d) or a file containing domains (-f)")
		os.Exit(1)
	}

	// Parse the length flag
	length, err := strconv.Atoi(*lengthFlag)
	if err != nil || length <= 0 {
		logError(fmt.Errorf("invalid length specified: %v", err))
		fmt.Println("Invalid length specified")
		os.Exit(1)
	}

	// Define the character sets based on flags
	var chars []rune
	if *alphanumeric {
		for ch := 'a'; ch <= 'z'; ch++ {
			chars = append(chars, ch)
		}
		for ch := '0'; ch <= '9'; ch++ {
			chars = append(chars, ch)
		}
	} else if *alphabetic {
		for ch := 'a'; ch <= 'z'; ch++ {
			chars = append(chars, ch)
		}
	}

	// Create a slice to hold domains
	var domains []string

	// Read domains from file if specified
	if *file != "" {
		domains, err = readDomainsFromFile(*file)
		if err != nil {
			logError(fmt.Errorf("error reading domains from file: %v", err))
			fmt.Println("Error reading domains from file:", err)
			os.Exit(1)
		}
	} else {
		// Use the single domain specified
		domains = append(domains, *domain)
	}

	// Replace the placeholder with each combination and print or save the URLs
	processCombination := func(combination string) {
		for _, domain := range domains {
			url := strings.Replace(domain, "[here]", combination, -1)
			if *output != "" {
				saveToFile(*output, url)
			} else {
				fmt.Println(url)
			}
		}
	}

	// Generate combinations and process them
	generateCombinations(chars, length, processCombination)
}

// Function to save content to a file
func saveToFile(filename, content string) {
	file, err := os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		logError(fmt.Errorf("error opening output file: %v", err))
		fmt.Println("Error opening output file:", err)
		os.Exit(1)
	}
	defer file.Close()

	if _, err := file.WriteString(content + "\n"); err != nil {
		logError(fmt.Errorf("error writing to output file: %v", err))
		fmt.Println("Error writing to output file:", err)
		os.Exit(1)
	}
}
