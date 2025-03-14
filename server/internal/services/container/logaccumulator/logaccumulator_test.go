package logaccumulator

import (
	"crypto/rand"
	"math/big"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

const chars = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

func generateRandomString(size int) (string, error) {
	var result string
	for i := 0; i < size; i++ {
		randomIndex, err := rand.Int(rand.Reader, big.NewInt(int64(len(chars))))
		if err != nil {
			return "", err
		}
		result += string(chars[randomIndex.Int64()])
	}
	return result, nil
}

func TestLogAccumulatorSizeTriggeredFlush(t *testing.T) {
	chunks := []*LogChunk{}

	la := NewLogAccumulator(0, 0, func(ls []*LogChunk) {
		chunks = append(chunks, ls...)
	})

	msg, err := generateRandomString(7168)
	require.NoError(t, err)

	la.AddLog(LogEntry{
		ProducerID: "1",
		Message:    msg,
	})

	la.AddLog(LogEntry{
		ProducerID: "1",
		Message:    msg,
	})

	require.Equal(t, 1, len(chunks))

	la.AddLog(LogEntry{
		ProducerID: "1",
		Message:    msg,
	})

	require.Equal(t, 2, len(chunks))
	require.Equal(t, true, chunks[0].ChunkID != chunks[1].ChunkID)
}

func TestLogAccumulatorTimeTriggeredFlush(t *testing.T) {
	chunks := []*LogChunk{}

	la := NewLogAccumulator(0, 10*time.Millisecond, func(ls []*LogChunk) {
		chunks = append(chunks, ls...)
	})

	msg, err := generateRandomString(168)
	require.NoError(t, err)

	la.AddLog(LogEntry{
		ProducerID: "1",
		Message:    msg,
	})

	t.Run("elapsed flush interval should trigger a flush when there is content", func(t *testing.T) {
		require.Equal(t, 0, len(chunks))
	})

	t.Run("elapsed flush interval should trigger a flush when there is content", func(t *testing.T) {
		time.Sleep(11 * time.Millisecond)

		require.Equal(t, 1, len(chunks))
	})

	t.Run("no new chunk because there were no new writes", func(t *testing.T) {
		time.Sleep(11 * time.Millisecond)

		require.Equal(t, 1, len(chunks))
	})

	t.Run("new chunk cretaed because of new writes", func(t *testing.T) {
		la.AddLog(LogEntry{
			ProducerID: "1",
			Message:    msg,
		})

		time.Sleep(11 * time.Millisecond)
		require.Equal(t, 2, len(chunks))
	})

	t.Run("chunk id should stay the same until size limit is reached", func(t *testing.T) {
		require.Equal(t, true, chunks[0].ChunkID == chunks[1].ChunkID)
	})
}
