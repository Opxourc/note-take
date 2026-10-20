# Note-Take

Note-Take is a lightweight, interactive command-line app for organizing notes in plain-text files. Create a file for a topic, add notes to it, and list or view your saved notes without leaving the terminal.

## Run

Install Go, then run the app from the project directory:

```sh
go run .
```

Notes are stored as files in the project's `notes` directory.

## Commands

```text
new -f <file>                 Create a note file
new -n <file> <note>          Add a note to an existing file
list -a                       List all note files
list -f <file>                Show all notes in a file
help                          Show command help
exit                          Quit
```

Example: `new -n journal.txt "Remember to take a break"`

## License

This project is licensed under the [MIT License](LICENSE).
