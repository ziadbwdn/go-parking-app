# 🚗 Go Parking System

A backend application for automated parking lot management built with Go, following SOLID principles and clean architecture patterns.

## 📋 Table of Contents

- [Overview](#overview)
- [Features](#features)
- [Architecture](#architecture)
- [Quick Start](#quick-start)
- [Usage](#usage)
- [Docker Setup](#docker-setup)
- [Project Structure](#project-structure)
- [Commands](#commands)
- [Testing](#testing)
- [Development](#development)

## 🎯 Overview

This parking system manages a parking lot with `n` slots, automatically assigns the nearest available slot to incoming cars, calculates parking charges based on duration, and processes various parking operations through a command-based interface.

### Key Business Rules
- **Slot Assignment**: Cars are assigned to the nearest available slot (lowest slot number)
- **Pricing**: $10 for first 2 hours, $10 for every additional hour
- **Commands**: System processes commands from input files for automation

## ✨ Features

- **Interactive File Selection**: Choose from available test scenarios
- **Automated Slot Assignment**: Nearest-to-entry slot allocation
- **Dynamic Pricing**: Time-based parking fee calculation
- **Command Processing**: File-based command execution
- **Clean Architecture**: SOLID principles implementation
- **Docker Support**: Containerized deployment
- **Comprehensive Testing**: Unit and integration tests

## 🏗️ Architecture

This application follows Clean Architecture principles with clear separation of concerns:

```
Domain Layer     → Core business entities and interfaces
Use Cases Layer  → Application business rules
Infrastructure   → External concerns (file I/O, persistence)
Interface Layer  → Command line interface and handlers
```

### SOLID Principles Implementation

'SOLID Principles are fully implemented here'

## 🚀 Quick Start

### Prerequisites
- Go 1.21 or higher
- Docker (optional, for containerized deployment)

### Local Development

1. **Clone the repository**
   ```bash
   git clone <repository-url>
   cd go-parking-app
   ```

2. **Install dependencies**
   ```bash
   go mod tidy
   ```

3. **Build the application**
   ```bash
   go build -o parking cmd/parking/main.go
   ```

4. **Run the application**
   ```bash
   ./parking
   ```

### Using Go Run (Development)
```bash
go run cmd/parking/main.go ./input/default-input.txt
```

## 💻 Usage

### Interactive Mode
When you run the application without arguments, it presents an interactive menu:

```
🚗 Parking System - Input File Selection
=======================================

Available input files:
  1. default-input.txt
  2. v01-input.txt
  3. v02-input.txt

Select a file (1-3): 
```

### Available Test Scenarios

| File | Description |
|------|-------------|
| `default-input.txt` | Complete feature demonstration |
| `v01-input.txt` | Small parking lot scenario (limited capacity) |
| `v02-input.txt` | Full parking lot scenario (capacity stress test) |

### Sample Output
```
Allocated slot number: 1
Allocated slot number: 2
Registration number KA-01-HH-3141 with Slot Number 6 is free with Charge $30
Slot No. Registration No.
1       KA-01-HH-1234
2       KA-01-HH-9999
```

## 🐳 Docker Setup

### Using Docker Compose (Recommended)

1. **Build and run**
   ```bash
   docker-compose up --build
   ```

2. **Run in detached mode**
   ```bash
   docker-compose up -d --build
   ```

3. **View logs**
   ```bash
   docker-compose logs -f parking-app
   ```

4. **Stop the application**
   ```bash
   docker-compose down
   ```

### Using Docker directly

1. **Build the image**
   ```bash
   docker build -t go-parking-app .
   ```

2. **Run the container**
   ```bash
   docker run -it --rm go-parking-app
   ```

### Interactive Docker Usage
The Docker container runs in interactive mode, allowing you to select input files just like the local version.

## 📁 Project Structure

```
go-parking-app/
├── cmd/
│   └── parking/
│       └── main.go              # Application entry point
├── internal/
│   ├── app/
│   │   └── app.go               # Application orchestration
│   ├── domain/
│   │   ├── models/              # Core entities
│   │   │   ├── car.go
│   │   │   ├── slot.go
│   │   │   └── ticket.go
│   │   └── services/            # Business interfaces
│   │       ├── parking.go
│   │       └── pricing.go
│   ├── infrastructure/
│   │   ├── commandline/
│   │   │   └── handler.go       # Command processing
│   │   └── persistence/
│   │       └── memory/
│   │           └── parking_repository.go
│   ├── test/                    # Test files
│   │   ├── parking_repository_test.go
│   │   ├── parking_services_test.go
│   │   └── pricing_service_test.go
│   └── usecases/
│       └── usecases.go          # Business use cases
├── input/                       # Test scenarios
│   ├── default-input.txt
│   ├── v01-input.txt
│   └── v02-input.txt
├── pkg/
│   └── utils/
│       └── file_reader.go       # Utility functions
├── Dockerfile
├── docker-compose.yml
├── go.mod
└── README.md
```

## 🎮 Commands

| Command | Format | Description | Example |
|---------|--------|-------------|---------|
| Create Lot | `create_parking_lot {capacity}` | Initialize parking lot | `create_parking_lot 6` |
| Park Car | `park {registration_number}` | Park a car | `park KA-01-HH-1234` |
| Car Exit | `leave {registration_number} {hours}` | Process departure | `leave KA-01-HH-1234 4` |
| Status | `status` | Show current status | `status` |

### Pricing Rules
- **First 2 hours**: $10
- **Additional hours**: $10 per hour
- **Examples**:
  - 1 hour = $10
  - 2 hours = $10  
  - 3 hours = $20
  - 4 hours = $30

## 🧪 Testing

### Run all tests
```bash
go test ./...
```

### Run specific test package
```bash
go test ./internal/test/
```

### Run tests with coverage
```bash
go test -cover ./...
```

### Test scenarios included
- Repository operations
- Parking service business logic
- Pricing calculations
- Edge cases (full lot, invalid operations)

## 🛠️ Development

### Adding New Test Scenarios

1. Create a new `.txt` file in the `input/` directory
2. Add commands following the API format
3. The application will automatically detect and offer the new scenario

### Extending Functionality

The clean architecture makes it easy to:
- Add new command types
- Implement different pricing strategies
- Switch storage backends
- Add new output formats

### Code Style
- Follow Go idioms and conventions
- Use meaningful variable names
- Include comprehensive error handling
- Write tests for new functionality

## 📝 Example Input File (Demo)

Try to input this txt below:
``` 
create_parking_lot 6
park KA-01-HH-1234
park KA-01-HH-9999
park KA-01-BB-0001
leave KA-01-HH-1234 4
status
```

## 🤝 Contributing

1. Fork the repository
2. Create a feature branch
3. Make your changes
4. Add tests for new functionality
5. Ensure all tests pass
6. Submit a pull request

## 📄 License

This project is licensed under the MIT License - see the LICENSE file for details.

## 🎯 Future Enhancements

- [ ] REST API interface
- [ ] Database persistence
- [ ] Real-time parking status updates
- [ ] Mobile app integration
- [ ] Advanced pricing models
- [ ] Parking reservations
- [ ] Analytics and reporting