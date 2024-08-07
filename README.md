# HeaCat-API - The System Monitoring API

## Introduction
HeaCat-API is a system monitoring API that provides information about the system's health and performance. It is designed to be used by system administrators and developers to monitor the system's health and performance. The API provides information about the system's CPU usage, memory usage, disk usage, network usage and more. It also provides information about the system's uptime, load average, and other system metrics.

## Features
- CPU monitoring
- Memory monitoring
- Disk monitoring (W.I.P)
- System Information API

## Installation
1. Clone the repository
```bash
git clone https://github.com/heacat/heacat-api.git
```
2. Install the dependencies
```bash
cd heacat-api
go mod tidy
```
3. Build the project
```bash
go build src/main.go
```
4. Run the project
```bash
./main
```

## Usage
1. Create a config file
```bash
cp config-example.yaml config.yaml
nano config.yaml                  # or use your favorite text editor
```
2. Run the program
```bash
./main
```