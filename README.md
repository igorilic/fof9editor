# Front Office Football Nine CSV Editor

A desktop application for creating and editing custom leagues for Front Office Football Nine.

## Description

FOF9 Editor is a comprehensive CSV editor that allows users to:
- Create custom leagues with custom players, coaches, and teams
- Edit existing league data with validation
- Import data from the default game files
- Manage league settings, salary caps, and schedules

## Requirements

- Go 1.21 or later
- Windows 10/11 (primary target platform)
- For development on Linux/WSL: X11 and OpenGL development libraries (or use Windows for building)

### Linux/WSL Dependencies (for testing UI locally)

On Ubuntu/Debian-based systems:
```bash
sudo apt-get install libgl1-mesa-dev xorg-dev
```

## Build Instructions

### On Windows
```bash
# Build the application
go build ./cmd/fof9editor

# Or use make
make build
```

### Cross-compile from Linux/WSL to Windows
```bash
# Set environment variables for Windows target
GOOS=windows GOARCH=amd64 CGO_ENABLED=1 CC=x86_64-w64-mingw32-gcc go build -o fof9editor.exe ./cmd/fof9editor
```

**Note**: Fyne requires CGO, so cross-compilation needs a C cross-compiler (like mingw-w64)

## Run Instructions

```bash
# Run directly
go run ./cmd/fof9editor

# Or run the built executable
./bin/fof9editor.exe
```

## Development

### Project Structure

```
fof9editor/
├── cmd/fof9editor/        # Application entry point
├── internal/
│   ├── app/               # Application lifecycle
│   ├── models/            # Data models
│   ├── data/              # CSV I/O
│   ├── validation/        # Validation rules
│   ├── ui/                # User interface
│   └── state/             # Application state
├── pkg/utils/             # Utility functions
└── testdata/fixtures/     # Test data
```

### Testing

```bash
go test ./...
```

## License

See LICENSE file for details.

## Documentation

- [Specification](spec.md) - Complete feature specification
- [Implementation Plan](plan.md) - Development roadmap
- [TODO](todo.md) - Progress tracking
