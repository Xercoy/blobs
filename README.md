# Glob - Generate files of various sizes.

Use Glob to create files of various sizes in bytes, megabytes, gigabytes, and terabytes.

# Usage

Specify the content of the blob via stdin.

Provide the data unit (unit), number of files (amount), and destination (dest):
```
echo "foobar" | go run main.go --unit="MB" --amount=2 --dest="./tmp"
```

### Flags

`unit` -  string flag, the unit of data to be created. Valid values are MB, GB, TB, KB, and B. Case sensitive for now.

`amount` - An integer flag of the number of files to be created.

`dest` - A string flag, the directory path of the output.

### TODO

- Tests, tests, and more tests
- Error checking and handling
- Improved configuration parsing