package main

import (
	"log"
	"os"
	"strconv"

	maelstrom "github.com/jepsen-io/maelstrom/demo/go"
)

func main() {

	counter := 0

	n := maelstrom.NewNode()

	n.Handle("generate", func(msg maelstrom.Message) error {

		id := n.ID() + "-" + strconv.Itoa(counter)

		body := map[string]any{
			"type": "generate_ok",
			"id":   id,
		}

		counter++
		return n.Reply(msg, body)
	})

	// Execute the node's message loop. This will run until STDIN is closed.
	if err := n.Run(); err != nil {
		log.Printf("ERROR: %s", err)
		os.Exit(1)
	}
}
