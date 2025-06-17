package redisapp

import (
	"context"
	"log"
	"time"

	"github.com/redis/go-redis/v9"
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

func (c *StreamConsumer) CreateConsumerGroup() error {
	ctx := context.Background()
	return c.client.XGroupCreateMkStream(ctx, c.stream, c.group, "0").Err()
}

func (c *StreamConsumer) Consume(handler func(map[string]string) error) error {
	ctx := context.Background()

	for {
		res, err := c.client.XReadGroup(ctx, &redis.XReadGroupArgs{
			Group:    c.group,
			Consumer: c.consumerName,
			Streams:  []string{c.stream, ">"},
			Count:    10,
			Block:    time.Second * 5,
		}).Result()

		if err != nil && err != redis.Nil {
			log.Println("Error reading from stream:", err)
			continue
		}

		for _, stream := range res {
			for _, msg := range stream.Messages {
				err := handler(ValuesToString(msg))
				if err == nil {
					// Ack sau khi xử lý xong
					c.client.XAck(ctx, c.stream, c.group, msg.ID)
				} else {
					log.Println("Handler error:", err)
				}
			}
		}
	}
}

func (c *StreamConsumer) RecoverPending(handler func(map[string]string) error) {
	ctx := context.Background()

	res, err := c.client.XPendingExt(ctx, &redis.XPendingExtArgs{
		Stream:   c.stream,
		Group:    c.group,
		Start:    "-",
		End:      "+",
		Count:    10,
		Consumer: c.consumerName,
	}).Result()

	if err != nil {
		log.Println("XPendingExt error:", err)
		return
	}

	for _, pending := range res {
		msgRes, err := c.client.XRange(ctx, c.stream, pending.ID, pending.ID).Result()
		if err != nil || len(msgRes) == 0 {
			continue
		}

		msg := msgRes[0]
		err = handler(ValuesToString(msg))
		if err == nil {
			c.client.XAck(ctx, c.stream, c.group, msg.ID)
		} else {
			log.Println("Error recovering pending message:", err)
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
