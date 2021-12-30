package main

import (
	"context"
	"fmt"
	"strings"
	"sync/atomic"

	"github.com/segmentio/kafka-go"

	"github.com/ncarlier/feedpushr/v3/pkg/format"
	"github.com/ncarlier/feedpushr/v3/pkg/model"
)

var spec = model.Spec{
	Name: "kafka",
	Desc: "Send new articles to Kafka.",
	PropsSpec: []model.PropSpec{
		{
			Name: "brokers",
			Desc: "Brokers (comma-separated values)",
			Type: model.Text,
		},
		{
			Name: "topic",
			Desc: "Topic",
			Type: model.Text,
		},
		{
			Name: "format",
			Desc: "Payload format (internal JSON format if not provided)",
			Type: model.Textarea,
		},
	},
}

// KafkaOutputPlugin is the Kafka output plugin
type KafkaOutputPlugin struct{}

// Spec returns plugin spec
func (p *KafkaOutputPlugin) Spec() model.Spec {
	return spec
}

// Build creates kafka output provider instance
func (p *KafkaOutputPlugin) Build(def *model.OutputDef) (model.Output, error) {
	formatter, err := format.NewOutputFormatter(def)
	if err != nil {
		return nil, err
	}
	brokers := def.Props.Get("brokers")
	if brokers == "" {
		return nil, fmt.Errorf("missing brokers property")
	}
	topic := def.Props.Get("topic")
	if topic == "" {
		return nil, fmt.Errorf("missing topic property")
	}

	kafkaWriter := kafka.NewWriter(kafka.WriterConfig{
		Brokers:  strings.Split(brokers, ","),
		Topic:    topic,
		Balancer: &kafka.Hash{},
	})

	definition := *def
	definition.Spec = spec

	return &KafkaOutputProvider{
		definition:  definition,
		formatter:   formatter,
		brokers:     brokers,
		topic:       topic,
		kafkaWriter: kafkaWriter,
	}, nil
}

// KafkaOutputProvider output provider to send articles to kafka
type KafkaOutputProvider struct {
	definition  model.OutputDef
	formatter   format.Formatter
	brokers     string
	topic       string
	kafkaWriter *kafka.Writer
}

// Send sent an article as Tweet to a kafka timeline
func (op *KafkaOutputProvider) Send(article *model.Article) (bool, error) {
	b, err := op.formatter.Format(article)
	if err != nil {
		atomic.AddUint32(&op.definition.NbError, 1)
		return false, err
	}
	err = op.kafkaWriter.WriteMessages(
		context.Background(),
		kafka.Message{
			Value: b.Bytes(),
		},
	)
	if err != nil {
		atomic.AddUint32(&op.definition.NbError, 1)
		return false, nil
	}
	atomic.AddUint32(&op.definition.NbSuccess, 1)
	return true, err
}

// GetDef return filter definition
func (op *KafkaOutputProvider) GetDef() model.OutputDef {
	return op.definition
}

// GetPluginSpec returns plugin spec
func GetPluginSpec() model.PluginSpec {
	return model.PluginSpec{
		Spec: spec,
		Type: model.OutputPluginType,
	}
}

// GetOutputPlugin returns output plugin
func GetOutputPlugin() (op model.OutputPlugin, err error) {
	return &KafkaOutputPlugin{}, nil
}
