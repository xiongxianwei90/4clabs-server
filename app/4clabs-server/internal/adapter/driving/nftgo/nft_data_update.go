package nftgo

import (
	"4clabs-server/app/4clabs-server/internal/adapter/data/model"
	"4clabs-server/app/4clabs-server/internal/adapter/data/query"
	forClabs "4clabs-server/app/4clabs-server/pkg/contract"
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"gorm.io/gorm"
	"math/big"
	"strconv"
	"strings"
)

func ComicNftUpdate(ctx context.Context, db *gorm.DB, rawurl string, contractAddress string, isIncrement bool) error {
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

	rnft := query.Use(db).ComicsNft
	totalDb, err := rnft.WithContext(ctx).Count()

	var start = total.Int64() - 1
	var end int64 = 0
	// 从后往前更新，是否增量更新，增量更新只更新合约和数据库相差数量
	if isIncrement {
		end = totalDb - 1
	}

	for i := start; i >= end; i-- {
		owner, err := instance.OwnerOf(nil, big.NewInt(i))
		if err != nil {
			break
		}
		comicWorksId, err := instance.ComicWorksIdWithTokenId(nil, big.NewInt(i))
		comicId := comicWorksId.Int64() + 1
		comicWorks, err := instance.ComicWorks(nil, big.NewInt(i))
		if err != nil {
			break
		}

		query := rnft.WithContext(ctx).
			Where(rnft.ID.Eq(int32(i)))
		_, err = query.First()
		if err == nil {
			query.Updates(map[string]interface{}{"owner": owner.String(), "comics_id": comicId, "author": comicWorks.Author.String()})
		} else {
			rnft.WithContext(ctx).Create(&model.ComicsNft{
				ID:       int32(i + 1),
				ComicsID: fmt.Sprint(int(comicId)),
				Owner:    owner.String(),
				Author:   comicWorks.Author.String(),
			})
		}
	}
	return nil
}

func ComicUpdate(ctx context.Context, db *gorm.DB, rawurl string, contractAddress string, isIncrement bool) error {
	client, err := ethclient.Dial(rawurl)
	if err != nil {
		return err
	}
	contract := common.HexToAddress(contractAddress)
	instance, err := forClabs.NewForClabs(contract, client)
	if err != nil {
		return err
	}

	rnft := query.Use(db).Comic
	totalDb, err := rnft.WithContext(ctx).Count()
	if err != nil {
		return err
	}

	total, err := instance.TotalComicWorksCount(nil)
	if err != nil {
		return err
	}

	var start = total.Int64() - 1
	var end int64 = 0

	// 从后往前更新，是否增量更新，增量更新只更新合约和数据库相差数量
	if isIncrement {
		end = totalDb - 1
	}

	println("total:", total.Int64(), "totalDb:", totalDb)

	for i := start; i >= end; i-- {
		comicWorks, err := instance.ComicWorksWithWorksId(nil, big.NewInt(i))
		if err != nil {
			break
		}

		query := rnft.WithContext(ctx).
			Where(rnft.ID.Eq(int32(i + 1)))
		_, err = query.First()

		price, _ := WeiToEth(comicWorks.Price)
		newData := map[string]interface{}{"id": i + 1, "token_id": comicWorks.BasedOnTokenId.String(), "contract_address": comicWorks.BasedOnContract.String(),
			"user_address": comicWorks.Author.String(), "mint_limit": comicWorks.Limit.Int64(),
			"mint_price": price, "imageURIs": strings.Join(comicWorks.ImageURIs, ","), "name": comicWorks.Name}
		if err == nil {
			query.Updates(newData)
		} else {
			rnft.WithContext(ctx).Create(&model.Comic{
				ID:              int32(i + 1),
				TokenID:         comicWorks.BasedOnTokenId.String(),
				Name:            comicWorks.Name,
				ContractAddress: comicWorks.BasedOnContract.String(),
				UserAddress:     comicWorks.Author.String(),
				MintLimit:       int32(comicWorks.Limit.Int64()),
				MintPrice:       price,
				ImageURIs:       strings.Join(comicWorks.ImageURIs, ","),
			})
		}

	}

	return nil
}

func EthToWei(val float64) *big.Int {
	bigval := new(big.Float)
	bigval.SetFloat64(val)
	coin := new(big.Float)
	coin.SetInt(big.NewInt(1000000000000000000))

	bigval.Mul(bigval, coin)

	result := new(big.Int)
	bigval.Int(result) // store converted number in result

	return result
}

func WeiToEth(val *big.Int) (float64, big.Accuracy) {
	bigval := new(big.Float)
	wei := new(big.Float)
	wei.SetInt(val)
	coin := new(big.Float)
	coin.SetInt(big.NewInt(1000000000000000000))
	bigval.Quo(wei, coin)
	return bigval.Float64()
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
