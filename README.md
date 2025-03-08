# go-dns-speedtest/go-dns-speedtest/README.md

# Go DNS Speedtest

Go DNS Speedtest is a simple and efficient tool for measuring the response time of various DNS servers. This project is implemented in Go and provides functionalities to load DNS server lists, perform speed tests, and output the results in a user-friendly format.

## Features

- Load DNS server lists from local files or online sources.
- Measure DNS response times with configurable test repetitions.
- Display results with statistical analysis including average, standard deviation, maximum, and minimum response times.
- Export results to CSV format.

## Installation

To get started with Go DNS Speedtest, ensure you have Go installed on your machine. You can download it from [the official Go website](https://golang.org/dl/).

Clone the repository:

```bash
git clone https://github.com/Kakune55/DNSspeedtest.git
cd DNSspeedtest
```

Navigate to the project directory and run:

```bash
go mod tidy
```

## Usage

To run the DNS speed test, execute the following command:

```bash
go run cmd/dnsspeedtest/main.go
```

You will be prompted to choose between a quick test or an average value test. Follow the on-screen instructions to proceed.

## Configuration

The application can be configured by modifying the configuration file located in the `internal/config` directory. Ensure that the DNS server list is correctly specified.

## Contributing

Contributions are welcome! If you have suggestions for improvements or new features, feel free to open an issue or submit a pull request.

## License

This project is licensed under the MIT License. See the LICENSE file for more details.