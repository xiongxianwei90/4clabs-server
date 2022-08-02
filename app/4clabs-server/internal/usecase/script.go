package usecase

import (
	"4clabs-server/app/4clabs-server/internal/ports"
	"context"
)

var _ ports.Script = (*Script)(nil)

type Script struct {
	script ports.Script
}

func NewScript(script ports.Script) *Script {
	return &Script{script: script}
}

func (s Script) RegisterUpdate(ctx context.Context, contactAddress string, tokenId string, userAddress string) error {
	return s.script.RegisterUpdate(ctx, contactAddress, tokenId, userAddress)
}
