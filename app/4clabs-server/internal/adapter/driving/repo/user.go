package repo

import (
	"4clabs-server/app/4clabs-server/internal/adapter/data/model"
	"4clabs-server/app/4clabs-server/internal/adapter/data/query"
	"4clabs-server/app/4clabs-server/internal/data"
	"context"
	"github.com/google/uuid"
	"github.com/pkg/errors"
	"gorm.io/gorm"
	"strings"
)

type User struct {
	data *data.Data
}

func (u User) GetNonce(ctx context.Context, address string) (string, error) {
	us := query.Use(u.data.DB).User
	col, err := us.WithContext(ctx).Where(us.Address.Eq(address)).Select(us.Nonce).First()
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return "", errors.WithMessagef(err, "address : %s", address)
	}
	if errors.Is(err, gorm.ErrRecordNotFound) {
		nonce, err := u.generateNonce(ctx)
		if err != nil {
			return "", err
		}
		if err := us.WithContext(ctx).Create(&model.User{
			Address: address,
			Nonce:   nonce,
		}); err != nil {
			return "", err
		}
		return nonce, nil
	}
	return col.Nonce, nil
}
func (u User) generateNonce(ctx context.Context) (string, error) {
	uuid, err := uuid.NewUUID()
	if err != nil {
		return "", errors.WithStack(err)
	}
	newn := strings.Replace(uuid.String(), "-", "", -1)
	return newn, nil
}
func (u User) RefreshNonce(ctx context.Context, address string) error {
	nonce, err := u.generateNonce(ctx)
	if err != nil {
		return err
	}
	us := query.Use(u.data.DB).User
	_, err = us.WithContext(ctx).Where(us.Address.Eq(address)).Update(us.Nonce, nonce)
	if err != nil {
		return errors.WithMessagef(err, "address : %s", address)
	}
	return nil
}

func NewUser(data *data.Data) *User {
	return &User{data: data}
}
