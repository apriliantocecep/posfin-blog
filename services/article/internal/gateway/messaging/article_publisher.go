package messaging

import (
	sharedmessaging "github.com/apriliantocecep/posfin-blog/shared/messaging"
	sharedmodel "github.com/apriliantocecep/posfin-blog/shared/model"
	amqp "github.com/rabbitmq/amqp091-go"
)

type ArticlePublisher struct {
	sharedmessaging.Publisher[*sharedmodel.ArticleEvent]
}

func NewArticlePublisher(channel *amqp.Channel, queueName string, routingKey string) *ArticlePublisher {
	return &ArticlePublisher{
		Publisher: sharedmessaging.Publisher[*sharedmodel.ArticleEvent]{
			Channel:      channel,
			QueueName:    queueName,
			RoutingKey:   routingKey,
			ExchangeName: "article",
		},
	}
}
