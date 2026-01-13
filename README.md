# Go Reloaded

<div align="center">

[![Go](https://img.shields.io/badge/Go-00ADD8?style=for-the-badge&logo=go&logoColor=white)](https://golang.org/)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg?style=for-the-badge)](https://opensource.org/licenses/MIT)

![Demo](assets/demo.gif)

**Intelligent text processing and formatting tool built in Go**

</div>

---

## 📋 Table of Contents

- [🎯 About](#-about)
- [✨ Features](#-features)
- [🚀 Installation](#-installation)
- [💻 Usage](#-usage)
- [📚 Examples](#-examples)
- [🧪 Testing](#-testing)
- [📁 Project Structure](#-project-structure)
- [🔧 How It Works](#-how-it-works)
- [📄 License](#-license)
- [👤 Author](#-author)
- [🙏 Acknowledgements](#-acknowledgements)

---

## 🎯 About

**Go Reloaded** is a command-line text processing utility that automatically applies formatting and text modification rules to transform messy text into clean, properly formatted content. It handles punctuation spacing, quote formatting, case modifications, numeric conversions, and grammatical corrections.

Built as part of the 01 Founders curriculum, this project demonstrates proficiency in:

- Text parsing and manipulation
- Regular expressions and pattern matching
- File I/O operations
- Modular Go architecture
- Comprehensive unit testing

---

## ✨ Features

### 🔤 **Text Case Modification**

Transform text case with inline tags:

- `(up)` - Convert to UPPERCASE
- `(low)` - Convert to lowercase
- `(cap)` - Capitalize First Letter
- `(up, N)` - Convert last N words to UPPERCASE
- `(low, N)` - Convert last N words to lowercase
- `(cap, N)` - Capitalize last N words

### 📐 **Punctuation Formatting**

Automatically fix spacing around punctuation marks:

- Attach punctuation to preceding word
- Add proper spacing after punctuation
- Handle `.` `,` `!` `?` `:` `;`

### 💬 **Quote Formatting**

Clean up quote placement:

- Remove extra spaces inside quotes
- Properly attach quotes to enclosed text
- Handle single quotes (`'`)

### 🔢 **Numeric Conversion**

Convert between number systems:

- `(bin)` - Binary to decimal (e.g., `1010 (bin)` → `10`)
- `(hex)` - Hexadecimal to decimal (e.g., `1A (hex)` → `26`)

### 📝 **Article Correction**

Smart grammar fixes:

- Automatically change "a" to "an" before vowel sounds
- Handles both uppercase and lowercase

---

## 🚀 Installation

### Prerequisites

- Go 1.16 or later
- Git

### Setup

```bash
# Clone the repository
git clone https://learn.01founders.co/git/iyoussef/Go-Reloaded.git
cd Go-Reloaded

# Verify installation
go version
```

---

## 💻 Usage

### Basic Command

```bash
go run ./cmd <input-filename> <output-filename>
```

**Important:**

- Input files must be placed in `text-files/` directory
- Output will be written to `outputs/` directory
- Specify only the filename, not the full path

### Quick Start

```bash
# Try with sample files
go run ./cmd sample-1.txt result-1.txt
go run ./cmd sample-2.txt result-2.txt
go run ./cmd sample-5.txt result-5.txt

# Check the results
cat outputs/result-1.txt
```

### Using Your Own Files

1. **Place your input file** in `text-files/` directory
2. **Run the command:**

```bash
   go run ./cmd myfile.txt myoutput.txt
```

3. **Check the result** in `outputs/myoutput.txt`

---

## 📚 Examples

### Example 1: Case Modification

**Input:** `text-files/sample-1.txt`

```
it (cap) was the best of times, it was the worst of times (up) ,
it was the age of wisdom, it was the age of foolishness (cap, 6)
```

**Output:** `outputs/result-1.txt`

```
It was the best of times, it was the worst of TIMES,
it was the age of wisdom, It Was The Age Of Foolishness
```

---

### Example 2: Multiple Transformations

**Input:** `text-files/sample-2.txt`

```
If I make you BREAKFAST IN BED (low, 3) just say thank you instead of: how (cap) did you get in my house (up, 2) ?
```

**Output:** `outputs/result-2.txt`

```
If I make you breakfast in bed just say thank you instead of: How did you get in MY HOUSE?
```

---

### Example 3: Numeric Conversion

**Input:** `text-files/sample-3.txt`

```
I have to pack 101 (bin) outfits. Packed 1a (hex) just to be sure
```

**Output:** `outputs/result-3.txt`

```
I have to pack 5 outfits. Packed 26 just to be sure
```

---

### Example 4: Punctuation Formatting

**Input:** `text-files/sample-4.txt`

```
Don not be sad ,because sad backwards is das . And das not good
```

**Output:** `outputs/result-4.txt`

```
Don not be sad, because sad backwards is das. And das not good
```

---

### Example 5: Article Correction & Quotes

**Input:** `text-files/sample-5.txt`

```
harold wilson (cap, 2) : ' I am a optimist ,but a optimist who carries a raincoat . '
```

**Output:** `outputs/result-5.txt`

```
Harold Wilson: 'I am an optimist, but an optimist who carries a raincoat.'
```

---

## 🧪 Testing

### Run All Tests

```bash
# Run complete test suite
go test -v ./unit_tests/

# Run with coverage
go test -v -cover ./unit_tests/
```

### Run Specific Tests

```bash
# Test case modifications
go test -v ./unit_tests/ -run TestCase

# Test hexadecimal conversion
go test -v ./unit_tests/ -run TestHex

# Test binary conversion
go test -v ./unit_tests/ -run TestBin

# Test punctuation formatting
go test -v ./unit_tests/ -run TestPunctuation

# Test article correction
go test -v ./unit_tests/ -run TestA_An

# Test file utilities
go test -v ./unit_tests/ -run TestFileUtils
```

### Test Coverage

The project includes comprehensive unit tests for:

- ✅ Case transformations (`case_test.go`)
- ✅ Hexadecimal conversion (`hex_test.go`)
- ✅ Binary conversion (`bin_test.go`)
- ✅ Punctuation formatting (`punctuation_test.go`)
- ✅ Article correction (`a_an_test.go`)
- ✅ File operations (`fileutils_test.go`)

---

## 📁 Project Structure

```
Go-Reloaded/
├── cmd/
│   └── main.go              # Entry point and CLI interface
├── text-files/              # Input files directory
│   ├── sample-1.txt         # Case modification example
│   ├── sample-2.txt         # Multiple transformations
│   ├── sample-3.txt         # Numeric conversion
│   ├── sample-4.txt         # Punctuation formatting
│   └── sample-5.txt         # Article & quote formatting
├── outputs/                 # Output files directory (generated results)
├── internal/
│   ├── textmod/             # Text modification logic
│   │   ├── textmod.go       # Main text processing
│   │   ├── case.go          # Case transformations
│   │   ├── punctuation.go   # Punctuation formatting
│   │   ├── a_an.go          # Article correction
│   │   ├── hex.go           # Hex conversion
│   │   └── bin.go           # Binary conversion
│   └── utils/
│       └── fileutils.go     # File I/O operations
├── unit_tests/              # Comprehensive test suite
│   ├── case_test.go
│   ├── punctuation_test.go
│   ├── a_an_test.go
│   ├── hex_test.go
│   ├── bin_test.go
│   └── fileutils_test.go
├── assets/
│   └── demo.gif             # Demo recording
├── go.mod                   # Go module definition
├── go.sum                   # Go dependencies
├── LICENSE.txt              # MIT License
└── README.md                # This file
```

---

## 🔧 How It Works

### Processing Pipeline

```
Input File → Read → Modify Text → Write → Output File
```

### Modification Order

1. **Numeric Conversions** - Convert hex/binary to decimal
2. **Case Modifications** - Apply (up), (low), (cap) transformations
3. **Article Correction** - Fix "a" → "an" before vowels
4. **Punctuation Formatting** - Adjust spacing around punctuation
5. **Quote Formatting** - Clean up quote placement

### Core Modules

**`textmod.go`**

- Orchestrates all text modifications
- Applies transformations in correct order
- Ensures no conflicts between operations

**`case.go`**

- Handles uppercase, lowercase, and capitalization
- Supports numbered variants (e.g., `(up, 3)`)
- Preserves text structure

**`punctuation.go`**

- Fixes spacing around `.` `,` `!` `?` `:` `;`
- Ensures proper attachment to words
- Maintains readability

**`a_an.go`**

- Detects vowel sounds
- Converts "a" to "an" automatically
- Preserves original case

**`hex.go` & `bin.go`**

- Convert number systems to decimal
- Validate input format
- Handle edge cases

---

## 📄 License

This project is licensed under the **MIT License**.

```
MIT License

Copyright (c) 2025 IbsYoussef

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.
```

See [LICENSE.txt](LICENSE.txt) for full details.

---

## 👤 Author

**[IbsYoussef](https://github.com/IbsYoussef)** - Built as part of the 01 Founders Piscine-Go curriculum.

---

## 🙏 Acknowledgements

- Built as part of the **[01 Founders](https://01-edu.org/)** Piscine-Go curriculum
- Inspired by real-world text processing challenges
- Thanks to the Go community for excellent documentation

---

<div align="center">

**[⬆ Back to Top](#go-reloaded)**

</div>
