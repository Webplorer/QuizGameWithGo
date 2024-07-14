# QuizGameWithGo
A simple Go-based quiz app that reads questions from a CSV file and quizzes you within a time limit.

## Features

- Reads quiz questions and answers from a CSV file.
- Timed quiz with a configurable time limit.
- Displays the final score after the quiz ends.

## Getting Started

### Prerequisites

- Go installed on your system. You can download it from [here](https://golang.org/dl/).

### Running the Application

1. Clone the repository:

    ```sh
    git clone https://github.com/Webplorer/QuizGameWithGo.git
    cd QuizGameWithGo
    ```

2. Build the application:

    ```sh
    go build -o quiz-app
    ```

3. Run the application:

    ```sh
    ./quiz-app -csv=problems.csv -limit=30
    ```

### Flags

- `-csv`: The path to the CSV file containing the quiz questions and answers. The default value is `problems.csv`.
- `-limit`: The time limit for the quiz in seconds. The default value is `30` seconds.

## CSV Format

The CSV file should contain questions and answers in the following format:

```csv
question,answer
What is 2+2?,4

