package redisapp

import (
	"context"
	"log"
	"time"

	"github.com/redis/go-redis/v9"
)

const (
	batchSize   = 100
	batchWindow = 2 * time.Second
)

var (
	buffer     []redis.XMessage
	lastInsert = time.Now()
)

type StreamConsumer struct {
	client       *redis.Client
	stream       string
	group        string
	consumerName string
}

func NewStreamConsumer(client *redis.Client, stream, group, consumerName string) *StreamConsumer {
	return &StreamConsumer{
		client:       client,
		stream:       stream,
		group:        group,
		consumerName: consumerName,
	}
}

func (c *StreamConsumer) CreateConsumerGroup() {
	ctx := context.Background()
	c.client.XGroupCreateMkStream(ctx, c.stream, c.group, "0")
}

func (c *StreamConsumer) Consume(handler func(ctx context.Context, data []map[string]string) error) error {
	ctx := context.Background()

	for {
		res, err := c.client.XReadGroup(ctx, &redis.XReadGroupArgs{
			Group:    c.group,
			Consumer: c.consumerName,
			Streams:  []string{c.stream, ">"},
			Count:    batchSize,
			Block:    time.Second * 5,
		}).Result()

		if err != nil && err != redis.Nil {
			log.Println("+++Error reading from stream:", err)
			continue
		}

		now := time.Now()

		for _, stream := range res {
			for _, msg := range stream.Messages {
				buffer = append(buffer, msg)
			}
		}

		// Điều kiện xử lý batch
		if len(buffer) >= batchSize || now.Sub(lastInsert) >= batchWindow {
			// Gọi handler xử lý cả batch
			if err := c.handlerBatch(ctx, buffer, handler); err != nil {
				log.Println("+++Handler batch error:", err)
				// Có thể retry hoặc xử lý theo logic riêng
			} else {
				// Xác nhận ack toàn bộ message đã xử lý thành công
				for _, msg := range buffer {
					c.client.XAck(ctx, c.stream, c.group, msg.ID)
				}

				buffer = nil // clear buffer
				lastInsert = time.Now()
			}
		}
	}
}

func (c *StreamConsumer) handlerBatch(ctx context.Context, msgs []redis.XMessage, handler func(ctx context.Context, data []map[string]string) error) error {
	// Chuyển đổi sang struct phù hợp
	var comments []map[string]string
	for _, msg := range msgs {
		cmt := ValuesToString(msg)
		comments = append(comments, cmt)
	}

	// Gọi insert batch (ví dụ MySQL: INSERT INTO ... VALUES (...), (...), ...)
	if len(comments) > 0 {
		return handler(ctx, comments)
	}

	return nil
}

func (c *StreamConsumer) RecoverPending(handler func(ctx context.Context, data []map[string]string) error) {
	var (
		streams []redis.XMessage
	)

	ctx := context.Background()

	res, err := c.client.XPendingExt(ctx, &redis.XPendingExtArgs{
		Stream:   c.stream,
		Group:    c.group,
		Start:    "-",
		End:      "+",
		Count:    100,
		Consumer: c.consumerName,
	}).Result()

	if err != nil {
		log.Println("+++XPendingExt error:", err)
		return
	}

	for _, pending := range res {
		msgRes, err := c.client.XRange(ctx, c.stream, pending.ID, pending.ID).Result()
		if err != nil || len(msgRes) == 0 {
			continue
		}
		streams = append(streams, msgRes[0])
	}

	if len(buffer) > 0 {
		if err := c.handlerBatch(ctx, buffer, handler); err != nil {
			log.Println("+++Batch recovery error:", err)
		} else {
			for _, msg := range buffer {
				c.client.XAck(ctx, c.stream, c.group, msg.ID)
			}
		}
	}
}

func ValuesToString(msg redis.XMessage) map[string]string {
	result := make(map[string]string)
	for k, v := range msg.Values {
		result[k] = v.(string)
	}
	return result
}
