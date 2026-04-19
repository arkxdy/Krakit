package db

import "github.com/segmentio/kafka-go"

type KafkaClient struct {
	conn *kafka.Conn
}

func connectKafka(broker string) (*KafkaClient, error) {
	conn, err := kafka.Dial("tcp", broker)
	if err != nil {
		return nil, err
	}
	// Verify broker is reachable
	if _, err := conn.ReadPartitions(); err != nil {
		return nil, err
	}
	return &KafkaClient{conn: conn}, nil
}

func (c *KafkaClient) Close() {
	c.conn.Close()
}
