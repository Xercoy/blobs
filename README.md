# Blob - Generate files of various sizes.

Use Blob to create files of various sizes in bytes, megabytes, gigabytes, and terabytes.

# Usage

Specify the content of the blob via stdin.

Provide the data unit (unit), number of files (amount), destination (dest), and the file name format (o):
```
echo "foobar" | go run main.go --unit="MB" --amount=2 --dest="./tmp" --o="FILE_%d.dat"
```

### Flags

`unit` -  string flag, the unit of data to be created. Valid values are MB, GB, TB, KB, and B. Case sensitive for now. Default = "MB"

`amount` - An integer flag of the number of files to be created. Default = 1

`dest` - A string flag, the directory path of the output. Default = "./"

`o` - A string flag to specify the file name format. The %d format specifier will denote the number sequence of the file. Default = "%d.dat"

### TODO

- Tests, tests, and more tests
- Error checking and handling for flags and other values
- Improved configuration parsing