# UCI Chess Engine in Go

This repository contains a Universal Chess Interface (UCI) compatible chess engine implemented in Go.

## About

This chess engine is designed to communicate with chess GUIs using the UCI protocol. It provides a foundation for developing a fully functional chess engine with customizable options.

Key features:
- UCI protocol implementation
- Configurable engine options via command-line flags
- Modular design for easy extension and improvement

## Requirements

- Go 1.16 or higher

## Building

To build the chess engine, follow these steps:

1. Clone the repository:
   ```
   git clone https://github.com/Mike17K/chess-engine.git
   cd chess-engine
   ```

2. Build the project:
   ```
   go build
   ```

This will create an executable named `chess_engine` (or `chess_engine.exe` on Windows) in the current directory.

## Running

To run the chess engine with default settings:
```
./chess_engine
```

To run the chess engine with custom options, use the following command:
```
./chess_engine -hash <hash_size> -nalimov-path <nalimov_path> -nalimov-cache <nalimov_cache> -ponder <ponder> -own-book <own_book> -multipv <multipv>
```

Replace `<hash_size>`, `<nalimov_path>`, `<nalimov_cache>`, `<ponder>`, `<own_book>`, and `<multipv>` with the desired values.
