package kafka

import (
	"github.com/Shopify/sarama"
	"github.com/pkg/errors"
)

type Producer struct {
	brokers []string
	sarama.SyncProducer
}

func newSyncProducer(brokers []string, config *sarama.Config) (sarama.SyncProducer, error) {
	syncProducer, err := sarama.NewSyncProducer(brokers, config)

	if err != nil {
		return nil, errors.Wrap(err, "error with sync kafka-producer")
	}

	return syncProducer, nil
}

func NewProducer(brokers []string, config *sarama.Config) (*Producer, error) {
	syncProducer, err := newSyncProducer(brokers, config)
	if err != nil {
		return nil, errors.Wrap(err, "error with sync kafka-producer")
	}

	producer := &Producer{
		brokers:      brokers,
		SyncProducer: syncProducer,
	}

	return producer, nil
}

func NewProducerConfig() *sarama.Config {
	syncProducerConfig := sarama.NewConfig()

	// случайная партиция
	// syncProducerConfig.Producer.Partitioner = sarama.NewRandomPartitioner

	// по кругу
	// syncProducerConfig.Producer.Partitioner = sarama.NewRoundRobinPartitioner

	// по ключу
	// syncProducerConfig.Producer.Partitioner = sarama.NewHashPartitioner
	/**
	Кейсы:
		- одинаковые ключи в одной партиции
		- при cleanup.policy = compact останется только последнее сообщение по этому ключу
	*/
	syncProducerConfig.Producer.Partitioner = sarama.NewRoundRobinPartitioner

	syncProducerConfig.Producer.RequiredAcks = sarama.WaitForAll

	/*
	  Если хотим exactly once, то выставляем в true

	  У продюсера есть счетчик (count)
	  Каждое успешно отправленное сообщение учеличивает счетчик (count++)
	  Если продюсер не смог отправить сообщение, то счетчик не меняется и отправляется в таком виде в другом сообщение
	  Кафка это видит и начинает сравнивать (в том числе Key) сообщения с одниковыми счетчиками
	  Далее не дает отправить дубль, если Idempotent = true
	*/
	syncProducerConfig.Producer.Idempotent = true
	syncProducerConfig.Net.MaxOpenRequests = 1

	// Если хотим сжимать, то задаем нужный уровень кодировщику
	syncProducerConfig.Producer.CompressionLevel = sarama.CompressionLevelDefault

	syncProducerConfig.Producer.Return.Successes = true
	syncProducerConfig.Producer.Return.Errors = true

	// И сам кодировщик
	syncProducerConfig.Producer.Compression = sarama.CompressionGZIP

	return syncProducerConfig
}

func (k *Producer) Close() error {
	err := k.SyncProducer.Close()
	if err != nil {
		return errors.Wrap(err, "kafka.Connector.Close")
	}

	return nil
}
