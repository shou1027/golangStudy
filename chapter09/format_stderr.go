package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Fprint(os.Stderr, "北海道", "の降水確率は", 100, "%です。\n")

	fmt.Fprintln(os.Stderr, "北海道", "の降水確率は", 100, "%です。\n")

	fmt.Fprintf(os.Stderr, "%sの降水確率は%d%%です。\n", "北海道", 100)
}
