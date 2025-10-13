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

## Build Instructions

```bash
# Build the application
go build ./cmd/fof9editor

# Or use make
make build
```

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
