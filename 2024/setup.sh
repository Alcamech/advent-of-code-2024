#!/bin/bash

if [ "$#" -ne 3 ]; then
  echo "Usage: $0 <day> <language> <attempt>"
  echo "Example: $0 1 go 1"
  exit 1
fi

DAY=$1
LANG=$2
ATTEMPT=$3

if [ -z "$SESSION_COOKIE" ]; then
  echo "Error: SESSION_COOKIE environment variable is not set."
  echo "Please set it using: export SESSION_COOKIE=<your-session-cookie>"
  exit 1
fi

DAY_PADDED=$(printf "%02d" $DAY)

ATTEMPT_DIR=day${DAY_PADDED}/attempt${ATTEMPT}
mkdir -p ${ATTEMPT_DIR}

if [ "$LANG" == "go" ]; then
  cat <<EOF >${ATTEMPT_DIR}/main.go
package main

import "fmt"

func main() {
    fmt.Println("Hello, Advent of Code!")
}
EOF
elif [ "$LANG" == "python" ]; then
  cat <<EOF >${ATTEMPT_DIR}/main.py
def main():
    print("Hello, Advent of Code!")

if __name__ == "__main__":
    main()
EOF
elif [ "$LANG" == "c" ]; then
  cat <<EOF >${ATTEMPT_DIR}/main.c
#include <stdio.h>

int main() {
    printf("Hello, Advent of Code!\\n");
    return 0;
}
EOF
else
  echo "Unsupported language: $LANG"
  exit 1
fi

if [ "$ATTEMPT" -eq 1 ]; then
  INPUT_URL="https://adventofcode.com/2024/day/${DAY}/input"
  curl -s --cookie "session=${SESSION_COOKIE}" ${INPUT_URL} -o ${ATTEMPT_DIR}/input.txt
fi

echo "# Day ${DAY} - Attempt ${ATTEMPT}" >${ATTEMPT_DIR}/README.md
echo "Day ${DAY}, Attempt ${ATTEMPT} setup complete."
