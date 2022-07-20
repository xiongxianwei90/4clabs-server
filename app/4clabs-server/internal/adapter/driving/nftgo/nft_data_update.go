package nftgo

import (
	"4clabs-server/app/4clabs-server/internal/adapter/data/query"
	forClabs "4clabs-server/app/4clabs-server/pkg/contract"
	"context"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"gorm.io/gorm"
	"math/big"
	"strconv"
)

func ComicNftAllUpdate(ctx context.Context, db *gorm.DB, rawurl string, contractAddress string) error {
	client, err := ethclient.Dial(rawurl)
	if err != nil {
		return err
	}
	contract := common.HexToAddress(contractAddress)
	instance, err := forClabs.NewForClabs(contract, client)
	if err != nil {
		return err
	}
	total, err := instance.TotalSupply(nil)
	if err != nil {
		return err
	}
	for i := 0; i < int(total.Int64()); i++ {
		rnft := query.Use(db).ComicsNft
		owner, err := instance.OwnerOf(nil, big.NewInt(int64(i)))
		if err != nil {
			break
		}
		query := rnft.WithContext(ctx).
			Where(rnft.ID.Eq(int32(i)))
		query.Updates(map[string]interface{}{"owner": owner.String()})
	}
	return nil
}

func ComicNftUpdateById(ctx context.Context, db *gorm.DB, tokenId string, buyerAddress string) error {
	rnft := query.Use(db).ComicsNft

	id, err := strconv.Atoi(tokenId)
	if err != nil {
		return err
	}
	query := rnft.WithContext(ctx).
		Where(rnft.ID.Eq(int32(id)))
	query.Updates(map[string]interface{}{"owner": buyerAddress})
	return nil
}
