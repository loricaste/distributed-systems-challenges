package main

import (
	"encoding/json"
	"log"
	"os"
	"sync"
	"sync/atomic"

	maelstrom "github.com/jepsen-io/maelstrom/demo/go"
)

// devi fare l'unmarshal di tutti i body :(

func main() {

	counter := &atomic.Int64{}

	n := maelstrom.NewNode()

	storedMessages := map[int64]any{}

	var mu sync.Mutex

	n.Handle("broadcast", func(msg maelstrom.Message) error {

		var body struct {
			Message int `json:"message"`
		}

		if err := json.Unmarshal(msg.Body, &body); err != nil {
			return err
		}

		mu.Lock()
		storedMessages[counter.Add(1)] = msg.message
		mu.Unlock()

		return n.Reply(msg, map[string]any{
			"type": "broadcast_ok",
		})
	})

	n.Handle("read", func(msg maelstrom.Message) error {

		mu.Lock()
		receivedMessages := []any{}
		for _, value := range storedMessages {
			receivedMessages = append(receivedMessages, value)
		}
		mu.Unlock()

		body := map[string]any{
			"type": "read_ok",
			"id":   receivedMessages,
		}

		return n.Reply(body)
	})

	n.Handle("topology", func(msg maelstrom.Message) error {

		body := map[string]any{
			"type": "topology_ok",
		}

		return n.Reply(msg, body)
	})

	// Execute the node's message loop. This will run until STDIN is closed.
	if err := n.Run(); err != nil {
		log.Printf("ERROR: %s", err)
		os.Exit(1)
	}
}
