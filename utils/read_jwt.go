package utils

import (
	"bufio"
	"fmt"
	"go-spotify-cli/constants"
	"os"
	"strings"
	"time"
)

func isTokenExpired(expiryTime time.Time) bool {
	return time.Now().After(expiryTime)
}

func ReadJWTToken() string {
	file, err := os.OpenFile(constants.TempFileName, os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		return ""
	}

	defer func() {
		closeErr := file.Close()
		if closeErr != nil {
			fmt.Println("Error closing file:", closeErr)
		}
	}()

	scanner := bufio.NewScanner(file)
	var token string
	var expiresIn string

	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.SplitN(line, "=", 2)
		if len(parts) != 2 {
			continue
		}

		switch parts[0] {
		case "jwtToken":
			token = parts[1]
		case "expiresIn":
			expiresIn = parts[1]
		}
	}

	if token == "" || expiresIn == "" {
		return ""
	}

	storedExpiryTime, err := time.Parse(time.RFC3339, expiresIn)
	if err != nil {
		fmt.Println("error converting expiresIn to the time.Time format", err)
	}

	tokenExpired := isTokenExpired(storedExpiryTime)

	if tokenExpired {
		fmt.Println("Token expired, getting a new one")
		return ""
	}

	if err := scanner.Err(); err != nil {
		return ""
	}

	fmt.Println("Token cache hit")

	return token
}