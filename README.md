# Go-Reloaded

**Go-Reloaded** is a command-line text processing utility program written in Go that applies formatting and text modification rules to text files. 
It adjusts punctuation spacing, corrects quote placement, and modifies text case based on inline tags.

## Table of Contents
1. [📝 About](#-about)
2. [✨ Features](#-features)
3. [📦 Installation](#-installation)
4. [🚀 Usage](#-usage)
5. [📁 File structure](#-file-structure)
6. [⚙️ Testing](#️-testing)
7. [🤝 Contributions](#-contributions)
8. [🙏 Acknowledgements](#-acknowledgements)
9. [📄 License](#-license)

## 📝 About
**Go-Reloaded** automates the cleanup and formatting of text files by applying a variety of modifications based on inline tags. This helps ensure that your documents follow consistent styling rules and improves their overall readability.

### What it does

**go-reloaded** performs several types of text modifications:

#### 1. Punctuation Formatting
- **Purpose:**  
  Adjusts spacing around punctuation marks (such as `.`, `,`, `!`, `?`, `:`, and `;`) so that they are attached directly to the previous word and followed by exactly one space.
- **Example:**  
  - Input:  
    ```
    I was sitting over there ,and then BAMM !!
    ```
  - Output:  
    ```
    I was sitting over there, and then BAMM!!
    ```

#### 2. Quote Formatting
- **Purpose:**  
  Ensures that single quotes are positioned directly next to the text they enclose by removing any extra spaces. Whether the quotes wrap a single word or multiple words, the tool "hugs" the content without additional spacing.
- **Example:**  
  - Input:  
    ```
    I am exactly how they describe me: ' awesome '
    ```
  - Output:  
    ```
    I am exactly how they describe me: 'awesome'
    ```

#### 3. Text Case Modification

- **Purpose:**  
  Recognizes inline tags such as `(up)`, `(low)`, and `(cap)` to modify the case of the preceding word. It also supports numbered variants—for example, `(up, 3)` will convert the last three words before the tag to uppercase.

- **Simple Example:**  
  - **Input:**  
    ```
    hello (up)
    ```  
  - **Output:**  
    ```
    HELLO
    ```

- **Numbered Variant Example:**  
  - **Input:**  
    ```
    one two three four five (up, 3)
    ```  
  - **Output:**  
    ```
    one two THREE FOUR FIVE
    ```
  - **Explanation:**  
    In this example, the tag `(up, 3)` tells the program to convert the three words immediately preceding it to uppercase. Thus, "three", "four", and "five" become "THREE", "FOUR", and "FIVE", while "one" and "two" remain unchanged.

#### 4. Numeric Conversion

**Binary Conversion (`(bin)` tag):**  
- **Purpose:**  
  Converts a binary number (base 2) to its decimal (base 10) equivalent. When the `(bin)` tag immediately follows a binary string, the tool parses the preceding word as a binary number and replaces it with the corresponding decimal value.  
- **Example:**  
  - **Input:**  
    ```
    1010 (bin)
    ```  
  - **Output:**  
    ```
    10
    ```  
  *Explanation:* The binary number `1010` is converted to the decimal number `10`.

**Hexadecimal Conversion (`(hex)` tag):**  
- **Purpose:**  
  Converts a hexadecimal number (base 16) to its decimal (base 10) equivalent. When the `(hex)` tag immediately follows a hexadecimal string, the tool parses the preceding word as a hexadecimal number and replaces it with its decimal value.  
- **Example:**  
  - **Input:**  
    ```
    1A (hex)
    ```  
  - **Output:**  
    ```
    26
    ```  
  *Explanation:* The hexadecimal number `1A` is converted to the decimal number `26`.

#### 5. Article Correction (`ConvertAtoAn`)
- **Purpose:**  
  Automatically converts instances of the article "a" to "an" when the following word begins with a vowel sound.
- **How It Works:**  
  The function scans the text for occurrences of "a" (or "A") and checks if the next word starts with one of the vowels (`a`, `e`, `i`, `o`, `u` in either case). If so, it appends an "n" to change "a" to "an".
- **Example:**  
  - Input:  
    ```
    a apple
    ```
  - Output:  
    ```
    an apple
    ```

### Why Use go-reloaded?

In today’s world of data and text processing, clean and consistently formatted text is essential for readability and professional presentation. **go-reloaded** offers:
- **Automated Formatting:**  
  It eliminates manual editing by automatically correcting common formatting issues.
- **Multiple Functionalities:**  
  It not only fixes punctuation and quotes but also handles numerical conversions and grammatical corrections.
- **Simplicity and Efficiency:**  
  With an easy-to-use command-line interface, you simply supply the input and output filenames, and the tool processes the text efficiently.
- **Extensibility:**  
  Written in Go with a modular design, it’s easy to extend the functionality further or integrate with other systems.

In a nutshell, **go-reloaded** is designed to streamline the text formatting process, ensuring that your documents look professional and are easy to read—whether you're preparing data for analysis, publishing content, or simply cleaning up raw text.

## ✨ Features
- **Punctuation Formatting:**  
  Automatically adjusts spacing around punctuation, ensuring proper attachment to words and consistent spacing.
  
- **Quote Formatting:**  
  Corrects spacing around single quotes so that quoted text appears neatly enclosed.
  
- **Text Case Modification:**  
  Supports inline tags like `(up)`, `(low)`, `(cap)` and their numbered variants to change text case effortlessly.
  
- **Numeric Conversion:**  
  Converts binary (`(bin)`) and hexadecimal (`(hex)`) numbers to their decimal equivalents.
  
- **Article Correction:**  
  Automatically changes "a" to "an" when appropriate (e.g., before vowel sounds).

For detailed examples and explanations of each feature, please see the [About](#about) section.

## 📦 Installation

Steps to install and run **go-reloaded**:

1. **Clone the Repository:**
This command clones the repository to your local machine and navigates into the project directory
    ```bash
    git clone https://github.com/IbsYoussef/Go-Reloaded.git
    cd go-reloaded
    ```

2. **Build the project:** Build the executable by running:
    ```bash
    go build ./cmd
    ```
3. **Run the Project:** You can run the project directly without building a seperate executable:
    ```bash
    go run ./cmd <input_filename.txt> <output_filename.txt>
    ```
    - **Note**:
    By default, input files should be placed in the text-files folder, and output files will be created in the outputs folder.

4. **Dependencies** This project requires Go 1.16 or later. All dependencies are managed through Go Modules, so no additional package installations are required.

## 🚀 Usage
**Go-Reloaded** is designed to be run from the command line with two arguments: the input filename and the output filename. The program automatically assumes that:
- Input files are located in the `text-files` directory.
- Processed output files will be written to the `outputs` directory.

### Running the program
To run the program, use the following command:
```go
    go run ./cmd <input_filename> <output_filename>
```

## 📁 File structure
```
.
├── README.md
├── cmd
│   ├── main.go
│   ├── outputs
│   │   ├── result-1.txt
│   │   ├── result-2.txt
│   │   ├── result-3.txt
│   │   ├── result-4.txt
│   │   └── result-5.txt
│   └── text-files
│       ├── sample-1.txt
│       ├── sample-2.txt
│       ├── sample-3.txt
│       ├── sample-4.txt
│       └── sample-5.txt
├── go.mod
├── go.sum
├── internal
│   ├── textmod
│   │   ├── a_an.go
│   │   ├── bin.go
│   │   ├── case.go
│   │   ├── hex.go
│   │   ├── punctuation.go
│   │   └── textmod.go
│   └── utils
│       └── fileutils.go
└── unit_tests
    ├── a_an_test.go
    ├── bin_test.go
    ├── case_test.go
    ├── fileutils_test.go
    ├── hex_test.go
    └── punctuation_test.go

8 directories, 27 files
```

## ⚙️ Testing
To ensure that **go-reloaded** works as expected, a comprehensive suite of unit tests is provided in the `unit_tests` folder. You can run these tests to verify the functionality of different parts of the project.

### Running All Tests

To run all tests in the project with detailed (verbose) output, execute the following command from the root directory of the project:

- **Run all the tests** in the ```unit_tests``` folder:
```bash
go test -v ./unit_tests 
```

### Running Individual tests

- **Run tests in a specific file (e.g., ```case_tests.go```):
```bash
go test -v ./unit_tests/case_test.go 
```

## 🤝 Contributions
Contributions are welcome! If you'd like to help improve **go-reloaded**, please follow these steps:

1. **Fork the Repository:**  
   Click the "Fork" button at the top-right of the repository page to create your own copy of the project.

2. **Create a New Branch:**  
   Create a new branch for your feature or bug fix:
   ```bash
    git checkout -b feature-or-bugfix-description
   ```
3. **Make your Changes:**
Implement your changes and ensure that your code adheres to the project's style guidelines.
Tip: Write or update tests as needed.

4. **Commit and Push your Changes**:
Commit your changes with a clear, descriptive message and push your branch to your forked repository:
    ```bash
    git commit -m "Add: description of your changes"
    git push origin feature-or-bugfix-description
    ```
5. **Open a Pull Request**:
Open a pull request (PR) from your branch to the main repository. Please include a clear description of your changes and the motivation behind them.
If you're not sure about a major change, open an issue first to discuss your ideas.

Thank you for helping make go-reloaded even better!

## 🙏 Acknowledgements
- Created as part of my Go learning journey at 01 Founders

## 📄 License
This project is licensed under the [MIT License](LICENSE).

Acknowledgements
Special Thanks:
Thanks to all contributors, mentors, and peers who provided feedback and support during the development of go-reloaded.

Inspiration:
This project was inspired by best practices in Go development and the need for automated text formatting solutions.

Resources:

The MIT License
Various open-source projects and communities that encourage collaboration and learning.
Thank you for checking out go-reloaded! We hope this tool helps streamline your text processing tasks and that you find it both useful and easy to contribute to.