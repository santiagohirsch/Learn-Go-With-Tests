package main

import (
	"os"
	"time"

	"Learn-Go-With-Tests/Math/clockface"
)

func main() {
	t := time.Now()
	clockface.SVGWriter(os.Stdout, t)
}
