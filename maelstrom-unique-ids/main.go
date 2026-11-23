package main

import (
	"fmt"
	"log"
	"os"
	"sync/atomic"

	maelstrom "github.com/jepsen-io/maelstrom/demo/go"
)

func main() {

	counter := &atomic.Int64{}

	n := maelstrom.NewNode()

	n.Handle("generate", func(msg maelstrom.Message) error {

		id := fmt.Sprintf("%s-%d", n.ID(), counter.Add(1))

		body := map[string]any{
			"type": "generate_ok",
			"id":   id,
		}

		return n.Reply(msg, body)
	})

	// Execute the node's message loop. This will run until STDIN is closed.
	if err := n.Run(); err != nil {
		log.Printf("ERROR: %s", err)
		os.Exit(1)
	}
}
