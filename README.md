# Glob - Generate files of various sizes.

Use Glob to create files of various sizes in bytes, megabytes, gigabytes, and terabytes.

# Usage

Specify the path, data unit, number of files, and content source:

```
glob --unit="B" --dest="./" --amount=3 --mode="./"
```

### Flags

Existing:

- Unit of space (B, MB, GB, TB)
- Number of files
- current directory

Future:

- Content of file (urandom, uzero, custom)
- Increment