entryPt := "./cmd"

default:
    just --list

run:
    go run {{ entryPt }}
