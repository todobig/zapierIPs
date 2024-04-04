package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
)

type IPInfo struct {
	Prefixes []struct {
		IP_Prefix string `json:"ip_prefix"`
	} `json:"prefixes"`
}

func main() {
	// Read the entire JSON file
	fileContent, err := ioutil.ReadFile("D:\\aws\\ip-ranges.json")
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	// Parse JSON data
	var ipInfo IPInfo
	err = json.Unmarshal(fileContent, &ipInfo)
	if err != nil {
		fmt.Println("Error parsing JSON:", err)
		return
	}

	// Create a regular expression to match IP addresses
	ipRegex := regexp.MustCompile(`\b(?:\d{1,3}\.){3}\d{1,3}/\d{1,2}\b`)

	// Open a new text file to write IP addresses
	outputFile, err := os.Create("D:\\aws\\extracted_ips.txt")
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer outputFile.Close()

	// Create a writer to write to the text file
	writer := bufio.NewWriter(outputFile)

	// Extract IP addresses from the "ip_prefix" field and write to the text file
	for _, prefix := range ipInfo.Prefixes {
		ips := ipRegex.FindAllString(prefix.IP_Prefix, -1)

		// Write each found IP address to the text file
		for _, ip := range ips {
			_, err := writer.WriteString(ip + "\n")
			if err != nil {
				fmt.Println("Error writing to file:", err)
				return
			}
		}
	}

	// Flush the writer to ensure all data is written to the file
	writer.Flush()

	fmt.Println("IP addresses extracted and written to D:\\aws\\extracted_ips.txt")
}
