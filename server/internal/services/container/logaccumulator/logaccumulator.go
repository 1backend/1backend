package logaccumulator

import (
	"bytes"
	"sync"
	"time"

	sdk "github.com/openorch/openorch/sdk/go"
)

type LogEntry struct {
	ProducerID string
	Message    string
}

type LogChunk struct {
	ChunkID     string
	ProducerID  string
	Buffer      bytes.Buffer
	LastFlush   time.Time
	HasNewWrite bool
}

type LogAccumulator struct {
	mu sync.Mutex

	// Active chunks per Producer ID
	chunks map[string]*LogChunk

	consumer func([]*LogChunk)

	maxChunkSize  int
	flushInterval time.Duration
}

func NewLogAccumulator(maxChunkSize int, flushInterval time.Duration, consumer func([]*LogChunk)) *LogAccumulator {
	if maxChunkSize == 0 {
		maxChunkSize = 8 * 1024 // 8 KB per batch default
	}

	if flushInterval == 0 {
		flushInterval = 1 * time.Second
	}

	la := &LogAccumulator{
		chunks:   make(map[string]*LogChunk),
		consumer: consumer,

		flushInterval: flushInterval,
		maxChunkSize:  maxChunkSize,
	}

	go la.autoFlush() // Background flushing

	return la
}

// Auto flush every FlushInterval duration
func (la *LogAccumulator) autoFlush() {
	ticker := time.NewTicker(la.flushInterval)
	defer ticker.Stop()

	for range ticker.C {
		la.Flush()
	}
}

// Flush logs that passed a certain interval (calls consumer only if there's data)
func (la *LogAccumulator) Flush() {
	la.mu.Lock()
	defer la.mu.Unlock()

	flushed := make([]*LogChunk, 0, len(la.chunks))

	for _, chunk := range la.chunks {
		if chunk.Buffer.Len() == 0 {
			continue
		}
		if time.Since(chunk.LastFlush) > la.flushInterval &&
			chunk.HasNewWrite {
			flushed = append(flushed, chunk)
			chunk.HasNewWrite = false
			chunk.LastFlush = time.Now()
		}
	}

	if len(flushed) > 0 {
		la.consumer(flushed)
	}
}

// Add log entry and manage chunk size
func (la *LogAccumulator) AddLog(entry LogEntry) {
	la.mu.Lock()
	defer la.mu.Unlock()

	chunk, exists := la.chunks[entry.ProducerID]

	if !exists {
		chunk = &LogChunk{
			ChunkID:     sdk.Id("colog"),
			ProducerID:  entry.ProducerID,
			HasNewWrite: true,
		}

		chunk.Buffer.WriteString(entry.Message)
		la.chunks[entry.ProducerID] = chunk

		return
	}

	if chunk.Buffer.Len()+len(entry.Message)+1 > la.maxChunkSize {
		la.flushSpecificKey(chunk.ProducerID)

		newChunk := &LogChunk{
			ChunkID:     sdk.Id("colog"),
			ProducerID:  entry.ProducerID,
			LastFlush:   time.Now(),
			HasNewWrite: true,
		}

		newChunk.Buffer.WriteString(entry.Message)

		la.chunks[entry.ProducerID] = newChunk

		return
	}

	chunk.HasNewWrite = true
	chunk.Buffer.WriteString(entry.Message + "\n")
}

func (la *LogAccumulator) flushSpecificKey(producerId string) {
	la.consumer([]*LogChunk{
		la.chunks[producerId],
	})
}
