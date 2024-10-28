package service

import (
	"github.com/Pilipupu/go-kit/framwork/message"
	"testing"
)

func TestService(t *testing.T) {

}

type TestRealService struct {
}

func (t TestRealService) Start() error {
	panic("implement me")
}

func (t TestRealService) handle(msg message.Message) {
	switch msg.(type) {

	}
}
