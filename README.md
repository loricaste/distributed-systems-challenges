# Distributed Systems Challenges

My solutions to the [Fly.io distributed systems challenges](https://fly.io/dist-sys/), built against [Maelstrom](https://github.com/jepsen-io/maelstrom).

Maelstrom spawns a node, feeds it a workload over stdin/stdout and checks the result against a consistency model. Each challenge adds a stricter model, so the code has to survive partition, latency and node failure.

## Challenges

Checked when the challenge passes its own test suite.

- [ ] **#1 Echo** — speak the protocol, nothing else
- [ ] **#2 Unique ID generation** — globally unique IDs without coordination
- [ ] **#3a Broadcast** — single node
- [ ] **#3b Broadcast** — multi-node
- [ ] **#3c Broadcast** — fault tolerant
- [ ] **#3d Broadcast** — efficient (low message count, low latency)
- [ ] **#4 Grow-only counter** — CRDT counter, sequential consistency
- [ ] **#5a Kafka-style log** — single node
- [ ] **#5b Kafka-style log** — multi-node
- [ ] **#5c Kafka-style log** — efficient
- [ ] **#6 Totally-available transactions** — available under partition

## Layout

One folder per challenge, added as I go.

## Notes

Written for learning, not for production: the point is to see where consistency, availability and partition tolerance actually collide.
