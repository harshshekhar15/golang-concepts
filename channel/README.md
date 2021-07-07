## Channels

Learnings:

- fan-out: Multiple funcs reading from the same channel.
- fan-in: One func reads from multiple channels and fan-in the results.
- Buffered channels can be used to prevent deadlocks but are still fragile and cannot eliminate deadlocks completely.
- Explicit cancellation: Send in a channel explicitly to each go routine and if that channel is closed all the go routines should return.
- Explicit cancellation is used to prevent deadlocks.
- Bounded-parallelism: Bounding the no. of go routines being launched to a certain value.

-------------------------------------------------------------

Non buffered channel:

- Sending to a non buffered channel is blocked until there is go routine waiting to receive value from that channel.

Buffered channel:

- A go routine can send to buffered channel until it reaches it capacity, sending is blocked post that until a go routine consumes values from that channel.

### Blogs:

- https://blog.golang.org/pipelines