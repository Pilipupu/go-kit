package service

import (
	"context"
	"github.com/Pilipupu/go-kit/framwork/component"
	"github.com/Pilipupu/go-kit/framwork/message"
)

type Service interface {
	component.Component
	handle(msg message.Message)
	Context() context.Context
}
