package service

import (
	"context"
	"fmt"
	Logger "github.com/agitdevcenter/gopkg/logger"
	"github.com/segmentio/kafka-go"
	"strings"
	"sync"
	"time"
)

type Server struct {
	logger Logger.Logger
	debug  bool
}

func New() *Server {

	return &Server{}
}

func (k *Server) SetLogger(logger Logger.Logger) {
	k.logger = logger
}

func (k *Server) SetDebug(enabled bool) {
	k.debug = enabled
}

func (k *Server) Start(ctx context.Context, wg *sync.WaitGroup) func() error {
	return func() error {

		readerConfig := kafka.ReaderConfig{
			Brokers:         []string{":9092"},
			GroupID:         "reminder-balancing",
			Topic:           "reminder",
			MinBytes:        10e3,
			MaxBytes:        10e6,
			MaxWait:         1 * time.Second,
			ReadLagInterval: -1,
		}
		reader := kafka.NewReader(readerConfig)

		loop := true

		errorServer := make(chan error, 1)

		go func() {
			<-ctx.Done()
			if err := reader.Close(); err != nil {
				errorServer <- fmt.Errorf("error closing kafka reader : %+v", err)
			}
			loop = false
			close(errorServer)
			wg.Done()
		}()

		if k.debug {
			k.logger.Info("starting kafka reader")
		}

		kafkaContext := context.Background()

		for loop {
			msg, err := reader.FetchMessage(kafkaContext)
			if err != nil {
				if strings.Contains(err.Error(), "error while receiving message: EOF") {
					k.logger.Error(fmt.Sprintf("error while receiving message: %+v", err))
				}
				continue
			}
			if k.debug {
				k.logger.Info(fmt.Sprintf("Topic : %s | Partition : %d | Offset : %d | Time : %+v | Message : %s", msg.Topic, msg.Partition, msg.Offset, msg.Time, string(msg.Value)))
			}
			if err := reader.CommitMessages(ctx, msg); err != nil {
				loop = false
				k.logger.Error(fmt.Sprintf("error commiting kafka messages : %+v", err))
			}
		}

		if k.debug {
			k.logger.Info("kafka reader stopped")
		}

		err := <-errorServer
		wg.Wait()
		return err
	}
}
