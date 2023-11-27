package my

import (
	"encoding/json"
	"github.com/Shopify/sarama"
	"github.com/modern-questions-team-13/orange-stock-market/internal/infrastructure/kafka"
	"github.com/rs/zerolog/log"
)

type KafkaSender struct {
	producer sarama.SyncProducer
	topic    string
}

func NewKafkaSender(producer sarama.SyncProducer, topic string) *KafkaSender {
	return &KafkaSender{
		producer,
		topic,
	}
}

func (s *KafkaSender) SendMessage(message kafka.RequestMessage) error {
	kafkaMsg, err := s.buildMessage(message)
	if err != nil {
		log.Err(err).Msg("Send message marshal error")
		return err
	}

	partition, offset, err := s.producer.SendMessage(kafkaMsg)

	if err != nil {
		log.Err(err).Msg("Send message connector error")
		return err
	}

	log.Info().Int32("Partition", partition).Int64("Offset", offset).Msg("send")
	return nil
}

func (s *KafkaSender) buildMessage(message kafka.RequestMessage) (*sarama.ProducerMessage, error) {
	msg, err := json.Marshal(message)

	if err != nil {
		log.Err(err).Msg("Send message marshal error")
		return nil, err
	}

	return &sarama.ProducerMessage{
		Topic:     s.topic,
		Value:     sarama.ByteEncoder(msg),
		Partition: -1,
	}, nil
}
