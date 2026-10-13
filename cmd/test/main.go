package main

import (
	"context"
	"log"

	"github.com/Rodco-off/checkCode/internal/sandbox"
)

func main() {
	ctx := context.Background()

	runner, err := sandbox.NewRunner(ctx)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("container started:", runner.ContainerID())

	code := `func Solve(input string) string {
        return "echo: " + input
    }`

	out, err := runner.Run(ctx, code, "hello")
	log.Printf("out=%q\nerr=%v", out, err)
}
