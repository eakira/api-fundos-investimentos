package services

import (
	"api-fundos-investimentos/application/domain"
	"api-fundos-investimentos/configuration/resterrors"
	"testing"

	"github.com/go-faker/faker/v4"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestCreateTopicService(t *testing.T) {

	t.Run("when_sending_a_valid_topic_returns_success", func(t *testing.T) {
		_, service, queue, _ := InitServiceTest(t)

		topicDomain := domain.TopicDomain{}
		err := faker.FakeData(&topicDomain)
		assert.Nil(t, err)

		queue.EXPECT().CreateTopic(gomock.Any()).Return(
			nil,
		).AnyTimes()

		_, err = service.CreateTopicService(topicDomain)
		assert.Nil(t, err)
	})

	t.Run("when_sending_a_invalid_topic_returns_error", func(t *testing.T) {
		_, service, queue, _ := InitServiceTest(t)

		topicDomain := domain.TopicDomain{}
		err := faker.FakeData(&topicDomain)
		assert.Nil(t, err)

		queue.EXPECT().CreateTopic(gomock.Any()).Return(
			resterrors.NewInternalServerError("Erro pra teste"),
		).AnyTimes()

		_, err = service.CreateTopicService(topicDomain)
		assert.NotNil(t, err)

	})
}
