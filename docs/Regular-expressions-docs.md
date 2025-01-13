# Regex Cheat Sheet for Go

## Table of Contents
- [Basic Characters](#basic-characters)
- [Anchors](#anchors)
- [Quantifiers](#quantifiers)
- [Character Classes](#character-classes)
- [Special Characters](#special-characters)
- [Groups and Lookarounds](#groups-and-lookarounds)
- [Escaping Characters](#escaping-characters)
- [Examples of Common Patterns](#examples-of-common-patterns)
- [Go Functions for Regex](#go-functions-for-regex)
- [Cheat Sheet Summary](#cheat-sheet-summary)

---

## **Basic Characters**
| Pattern   | Description                                  | Example         | Matches                  |
|-----------|----------------------------------------------|-----------------|--------------------------|
| `.`       | Any character (except newline)              | `a.b`           | `acb`, `a*b`            |
| `\d`     | Digit (0-9)                                 | `\d`            | `1`, `9`                |
| `\D`     | Non-digit                                   | `\D`            | `a`, `@`                |
| `\w`     | Word character (a-z, A-Z, 0-9, _)           | `\w`            | `a`, `9`, `_`           |
| `\W`     | Non-word character                          | `\W`            | `@`, `!`                |
| `\s`     | Whitespace (spaces, tabs)                  | `\s`            | Space, `\t`             |
| `\S`     | Non-whitespace                              | `\S`            | `a`, `9`, `!`           |

---

## **Anchors**
| Pattern   | Description                                  | Example         | Matches                  |
|-----------|----------------------------------------------|-----------------|--------------------------|
| `^`       | Start of string                             | `^a`            | `"abc"` matches `a`     |
| `$`       | End of string                               | `c$`            | `"abc"` matches `c`     |
| `\b`      | Word boundary                               | `\bword\b`      | Matches `word` as a full word |
| `\B`      | Non-word boundary                           | `\Bword\B`      | Matches `word` within other text |

---

## **Quantifiers**
| Pattern   | Description                                  | Example         | Matches                  |
|-----------|----------------------------------------------|-----------------|--------------------------|
| `*`       | 0 or more                                   | `a*`            | `""`, `a`, `aaa`        |
| `+`       | 1 or more                                   | `a+`            | `a`, `aa`, `aaa`        |
| `?`       | 0 or 1                                      | `a?`            | `""`, `a`               |
| `{n}`     | Exactly `n` repetitions                     | `a{3}`          | `aaa`                   |
| `{n,}`    | `n` or more repetitions                     | `a{2,}`         | `aa`, `aaa`, `aaaa`     |
| `{n,m}`   | Between `n` and `m` repetitions             | `a{1,3}`        | `a`, `aa`, `aaa`        |

---

## **Character Classes**
| Pattern   | Description                                  | Example         | Matches                  |
|-----------|----------------------------------------------|-----------------|--------------------------|
| `[abc]`   | Matches `a`, `b`, or `c`                    | `[aeiou]`       | `a`, `e`, `i`, etc.     |
| `[^abc]`  | Negation: matches anything except `a`, `b`, `c` | `[^aeiou]`  | Any consonant           |
| `[a-z]`   | Matches any lowercase letter                | `[a-z]`         | `a`, `b`, ..., `z`      |
| `[A-Z]`   | Matches any uppercase letter                | `[A-Z]`         | `A`, `B`, ..., `Z`      |
| `[0-9]`   | Matches any digit                           | `[0-9]`         | `0`, `1`, ..., `9`      |

---

## **Special Characters**
| Pattern   | Description                                  | Example         | Matches                  |
|-----------|----------------------------------------------|-----------------|--------------------------|
| `\.`      | Matches a literal `.`                       | `\.`            | `.`                      |
| `\\`      | Matches a literal `\`                       | `\\`            | `\`                      |
| `\|`      | Matches alternatives (OR)                  | `a|b`           | `a`, `b`                |
| `\(`      | Matches a literal `(`                      | `\(`            | `(`                      |
| `\)`      | Matches a literal `)`                      | `\)`            | `)`                      |

---

## **Groups and Lookarounds**
| Pattern        | Description                              | Example          | Matches                  |
|----------------|------------------------------------------|------------------|--------------------------|
| `(abc)`        | Capturing group                         | `(abc)`          | Captures `abc`          |
| `(?:abc)`      | Non-capturing group                     | `(?:abc)`        | Matches `abc` but doesn't capture |
| `(?=abc)`      | Positive lookahead                     | `\d(?=px)`       | `5` in `5px`            |
| `(?!abc)`      | Negative lookahead                     | `\d(?!px)`       | `5` in `5kg`            |
| `(?<=abc)`     | Positive lookbehind                    | `(?<=\$)\d+`     | `100` in `$100`         |
| `(?<!abc)`     | Negative lookbehind                    | `(?<!\$)\d+`     | `100` in `100kg`        |

---

## **Escaping Characters**
To match special regex characters literally (like `.` or `*`), prefix them with a backslash (`\`):
- Example: To match `1+1=2`, use `1\+1=2`.

---

## **Examples of Common Patterns**
| Pattern         | Description                              | Example Input    | Matches                  |
|-----------------|------------------------------------------|------------------|--------------------------|
| `\d{3}-\d{2}-\d{4}` | Social Security Number format         | `123-45-6789`    | `123-45-6789`           |
| `[a-z0-9._%+-]+@[a-z0-9.-]+\.[a-z]{2,}` | Email format    | `test@example.com` | `test@example.com`     |
| `https?://\S+`  | URL format                              | `http://go.dev`  | `http://go.dev`         |
| `^#\w+`         | Hashtags                                | `#golang`        | `#golang`               |

---

## **Go Functions for Regex**
Here are some key `regexp` functions in Go and how to use them:

### 1. **Compile or MustCompile**
- Compiles a regular expression pattern.
- `MustCompile` panics if the pattern is invalid.

```go
re := regexp.MustCompile(`\d+`)
fmt.Println(re.MatchString("123")) // Output: true
```

### 2. **MatchString**
- Checks if a string matches the pattern.

```go
fmt.Println(regexp.MustCompile(`a+b`).MatchString("aaab")) // Output: true
```

### 3. **FindString**
- Finds the first match of the pattern in a string.

```go
fmt.Println(regexp.MustCompile(`[a-z]+`).FindString("Go123Lang")) // Output: "o"
```

### 4. **FindAllString**
- Finds all matches of the pattern in a string.

```go
fmt.Println(regexp.MustCompile(`\d+`).FindAllString("123 456 789", -1)) // Output: ["123", "456", "789"]
```

### 5. **ReplaceAllString**
- Replaces matches of the pattern with a replacement string.

```go
fmt.Println(regexp.MustCompile(`\d+`).ReplaceAllString("abc123xyz", "###")) // Output: "abc###xyz"
```

### 6. **Split**
- Splits a string by matches of the pattern.

```go
fmt.Println(regexp.MustCompile(`[\s,]+`).Split("Go,Lang Python,Ruby", -1)) // Output: ["Go", "Lang", "Python", "Ruby"]
```

---

## **Cheat Sheet Summary**
1. **Start simple:** Match specific characters, then use ranges `[a-z]`, anchors (`^`, `$`), and quantifiers (`+`, `*`).
2. **Use groups** for capturing or combining patterns.
3. **Lookarounds** for advanced matching without consuming characters.

---
