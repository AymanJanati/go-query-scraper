# go-query-scraper
# Go Web Scraper

A simple web scraper built in Go that fetches the titles of web pages concurrently.

## Table of Contents

- [Features](#features)
- [Installation](#installation)
- [Usage](#usage)

## Features

- Concurrently fetches titles of multiple URLs.
- Simple command-line interface for user input.
- Uses Go's goroutines and channels for efficient fetching.

## Installation

1. Make sure you have Go installed on your machine. You can download it from [the official Go website](https://golang.org/dl/).
2. Clone this repository to your local machine:
   ```bash
   git clone https://github.com/AymanJanati/go-query-scraper.git
   ```
3. Navigate to the project directory:
   ```bash
   cd go-query-scraper
   ```
4. Install the required dependencies:
   ```bash
   go mod tidy
   ```

## Usage

1. Run the application:
   ```bash
   go run main.go
   ```
2. The program will fetch titles from a predefined list of URLs. You can modify the urls slice in the main function to add your own URLs.
