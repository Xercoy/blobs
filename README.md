![https://github.com/xercoy/blobs](blobs-logo.png)

# Blobs - Generate files of various sizes.

[![Build Status](https://travis-ci.org/Xercoy/blobs.svg?branch=master)](https://travis-ci.org/Xercoy/blobs)
[![GoDoc](https://godoc.org/github.com/xercoy/blobs?status.png)](http://godoc.org/github.com/xercoy/blobs)
[![Coverage Status](https://coveralls.io/repos/github/Xercoy/blobs/badge.svg?branch=master)](https://coveralls.io/github/Xercoy/blobs?branch=master)
[![Gitter](https://badges.gitter.im/Xercoy/blobs.svg)](https://gitter.im/Xercoy/blobs?utm_source=badge&utm_medium=badge&utm_campaign=pr-badge)

===

Use Blobs to create files of various sizes in bytes, megabytes, gigabytes, and terabytes.

# Usage
```
Blobs 1.0

 Usage: blobs <options> <options>...
  -amount int
    	Number of files to be created. (default 1)
  -dest string
    	Destination of created globs. (default "./")
  -help
    	Displays flag attributes & usage information.
  -o string
    	Format specifier for blob file name. %d is for the number sequence of the file. (default "%d.dat")
  -random
    	Random number of blobs ranging from 1 to the value of the amount flag.
  -unit string
    	Unit of space for the glob. (default "MB")
```

# Sample Usage

### TODO

- Tests, tests, and more tests
- Error checking and handling for flags and other values
- Improved configuration parsing