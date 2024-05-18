# ASCII-Art README

## Overview

**ASCII-Art** is a program that converts a given string into a graphic representation using ASCII characters. The program supports various features, including different templates, output to files, color manipulation, and text alignment.

## Features

### 1. Basic ASCII-Art Generation

- **Description**: Convert a given string into an ASCII representation using a specified template.
- **Input**: A string containing numbers, letters, spaces, special characters, and `\n`.
- **Output**: ASCII art of the input string.
- **Templates**: Provided templates include `shadow`, `standard`, and `thinkertoy`.

### 2. Template Selection

- **Description**: Choose a specific template for ASCII conversion.
- **Usage**: `go run . [STRING] [BANNER]`
- **Example**: `go run . "hello" standard`

### 3. Output to File

- **Description**: Write the ASCII art output to a specified file.
- **Usage**: `go run . --output=<fileName.txt> [STRING] [BANNER]`
- **Example**: `go run . --output=banner.txt "hello" standard`

### 4. Color Manipulation

- **Description**: Apply colors to the ASCII art.
- **Usage**: `go run . --color=<color> <letters to be colored> [STRING]`
- **Example**: `go run . --color=red "h" "hello"`

### 5. Text Alignment

- **Description**: Change the alignment of the ASCII art output.
- **Types**: 
  - `center`
  - `left`
  - `right`
  - `justify`
- **Usage**: `go run . --align=<type> [STRING] [BANNER]`
- **Example**: `go run . --align=center "hello" standard`

## Instructions

### Usage Instructions

#### Basic Usage

- **Command**: `go run . [STRING]`
- **Example**: `go run . "Hello\nWorld"`

#### Template Selection

- **Command**: `go run . [STRING] [BANNER]`
- **Example**: `go run . "hello" shadow`

#### Output to File

- **Command**: `go run . --output=<fileName.txt> [STRING] [BANNER]`
- **Example**: `go run . --output=test.txt "hello" thinkertoy`

#### Color Manipulation

- **Command**: `go run . --color=<color> <letters to be colored> [STRING]`
- **Example**: `go run . --color=red "lo" "hello"`

#### Text Alignment

- **Command**: `go run . --align=<type> [STRING] [BANNER]`
- **Example**: `go run . --align=center "hello" standard`

### Banner Format

- **Height**: Each character has a height of 8 lines.
- **Separation**: Characters are separated by two new lines (`\n\n`).

#### Examples

**Space ('A'):**

```
...........
...........
..../\.....
.../..\....
.././\.\...
./.____.\..
/_/....\_\.
...........
...........
```

**Exclamation Mark ('7'):**

```
........
._____..
|___..|.
.../ /..
../ /...
./_/....
........
........

```


### Error Handling

- Any incorrect formats will return a usage message:
  - **Template Selection**: `Usage: go run . [STRING] [BANNER]`
  - **Output to File**: `Usage: go run . [OPTION] [STRING] [BANNER]`
  - **Color Manipulation**: `Usage: go run . [OPTION] [STRING]`
  - **Text Alignment**: `Usage: go run . [OPTION] [STRING] [BANNER]`

## Example Outputs

### Basic Output

```shell
$ go run . "Hello\n" standard | cat -e
 _    _          _   _          $
| |  | |        | | | |         $
| |__| |   ___  | | | |   ___   $
|  __  |  / _ \ | | | |  / _ \  $
| |  | | |  __/ | | | | | (_) | $
|_|  |_|  \___| |_| |_|  \___/  $
                                $
                                $
$
```

### Output to File

```shell
$ go run . --output=banner.txt "hello" standard
$ cat -e banner.txt
 _              _   _          $
| |            | | | |         $
| |__     ___  | | | |   ___   $
|  _ \   / _ \ | | | |  / _ \  $
| | | | |  __/ | | | | | (_) | $
|_| |_|  \___| |_| |_|  \___/  $
                               $
                               $
$
```

### Colored Output

```shell
$ go run . --color=red "h" "hello"
```

### Center Aligned Output

```shell
|$ go run . --align=center "hello" standard
```

## Conclusion

The ASCII-Art program in Go is a versatile tool for converting strings into ASCII art, supporting various features such as template selection, output to files, color manipulation, and text alignment. Ensure to follow the usage instructions and error handling guidelines for optimal performance.