# Go-Prime

This is a test program for using Go for Command Line programs.
Go makes it simple to create and insatll programs in the command line.
The purpose of this program is to give the user some 'prime' functionalities in the command line.

### Install
1. Download the repository using `git clone` or download the zip file
1. `cd goprime`
1. Use either `go build` or `go install`, if you have a binary folder setup for go programs, to generate an executable binary

A program all abut prime numbers

### Functions
- Prime - Tells if a number is prime
- List - List all the Prime numbers from 0 to the number specified
- Factors - Lists all the prime factors of a number
- FactorOf - Determines if a number is a prime factor of another number (Ex: -fn=factorof (number) (prime number in question))

### Usage
Go Build method
```bash
./goprime -fn=[function] [number]
```

or 

Go Install Method
```bash
goprime -fn=[function] [number]
```
