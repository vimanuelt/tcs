# TCS (Tree and Content Snapshot)

TCS is a command-line tool for capturing and comparing directory structures and their contents. It supports creating snapshots of directory trees and generating delta reports that highlight changes between snapshots.

---

## Features
- Capture the structure and metadata of directories and files.
- Generate delta reports to track added, removed, or modified files.
- Efficiently handle large directories.
- Modular design for extensibility.

---

## Installation

### Build from Source
Ensure you have Go installed on your system. Then, run:

```bash
go build -o tcs main.go
```

This will generate an executable named `tcs` in the current directory.

### Using Makefile
To install `tcs` system-wide, use the provided `Makefile`:

#### Install:
```bash
sudo make install
```
This installs the `tcs` executable to `/usr/local/bin`.

#### Uninstall:
```bash
sudo make uninstall
```
This removes the `tcs` executable from `/usr/local/bin`.

---

## Usage

### 1. Snapshot Mode
Capture a snapshot of a directory's structure and contents.

#### Command:
```bash
./tcs -mode=snapshot -directory=<path_to_directory> -output-file=<output_snapshot_file>
```

#### Example:
```bash
./tcs -mode=snapshot -directory=/path/to/dir -output-file=snapshot.json
```

This will create a file `snapshot.json` containing the snapshot of `/path/to/dir`.

---

### 2. Delta Mode
Compare a directory's current state against a previous snapshot and generate a delta report.

#### Command:
```bash
./tcs -mode=delta -directory=<path_to_directory> -previous-snapshot=<previous_snapshot_file> -delta-file=<delta_report_file>
```

#### Example:
```bash
./tcs -mode=delta -directory=/path/to/dir -previous-snapshot=snapshot.json -delta-file=delta.json
```

This will generate a `delta.json` file containing the differences between the current state of `/path/to/dir` and the previous snapshot stored in `snapshot.json`.

---

## Command-Line Arguments

| Argument               | Description                                         |
|------------------------|-----------------------------------------------------|
| `-mode`                | Operation mode: `snapshot` or `delta`.             |
| `-directory`           | Path of the directory to process.                  |
| `-output-file`         | Path to save the snapshot (only for `snapshot`).    |
| `-previous-snapshot`   | Path to the previous snapshot (only for `delta`).   |
| `-delta-file`          | Path to save the delta report (only for `delta`).   |

---

## Testing

Run all unit and integration tests to verify functionality:

```bash
go test ./...
```

---

## Contributing

Contributions are welcome! Please follow these steps:
- Fork the repository.
- Create a new branch for your feature or bugfix.
- Commit your changes and submit a pull request.

---

## License

This project is licensed under the MIT License. See the `LICENSE` file for details.

---

## Contact
For any questions or issues, please create a GitHub issue or reach out to the maintainer.

