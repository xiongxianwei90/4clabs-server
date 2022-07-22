package assembler

import (
	typs "4clabs-server/api/nft/v1"
	"4clabs-server/app/4clabs-server/internal/domain/entity"
)

func CoverNftDetailToHttpDto(nft entity.NftDetail) *typs.Detail {
	var CoverPrice = func(p entity.PriceInfo) *typs.Stat_PriceInfo {
		return &typs.Stat_PriceInfo{
			TxHash:               p.TxHash,
			PriceToken:           float32(p.PriceToken),
			TokenSymbol:          p.TokenSymbol,
			TokenContractAddress: p.TokenContractAddress,
			PriceUsd:             float32(p.PriceUsd),
			Time:                 uint64(p.Time.Unix()),
		}
	}
	var CoverOwners = func(ows ...entity.OwnerStat) []*typs.Stat_OwnerStat {
		var result []*typs.Stat_OwnerStat
		for _, o := range ows {
			result = append(result, &typs.Stat_OwnerStat{
				Address:     o.Address,
				HoldingTime: uint64(o.HoldingTime.Seconds()),
			})
		}
		return result
	}

	return &typs.Detail{
		Summary: CoverNftToHttpDto(&nft.Summary)[0],
		Stat: &typs.Stat{
			LastUpdated:        nft.Stat.LastUpdated,
			SaleNum_7D:         nft.Stat.SaleNum7d,
			SaleNumAll:         nft.Stat.SaleNumAll,
			MaxPrice:           CoverPrice(nft.Stat.MaxPrice),
			MinPrice:           CoverPrice(nft.Stat.MinPrice),
			LastPrice:          CoverPrice(nft.Stat.LastPrice),
			PastOwners:         CoverOwners(nft.Stat.PastOwners...),
			CreateTime:         uint64(nft.Stat.CreateTime.Unix()),
			StartHoldingTime:   uint64(nft.Stat.StartHoldingTime.Unix()),
			LongestHoldingTime: uint64(nft.Stat.LongestHoldingTime.Seconds()),
		},
	}
}
func CoverComicToHttpDto(nfts ...entity.Comic) []*typs.ComicWork {
	var result []*typs.ComicWork
	for _, n := range nfts {
		result = append(result, &typs.ComicWork{
			OriginNft:          CoverNftToHttpDto(&n.Origin)[0],
			MintLimit:          n.MintLimit,
			MintPrice:          float32(n.MintPrice),
			Name:               n.Name,
			CreatedAtTimestamp: uint32(n.CreatedAt.Unix()),
			MinterAddress:      n.UserAddress,
			ImageUrl:           n.ImageUris,
		})
	}
	return result
}
func CoverNftToHttpDto(nfts ...*entity.Nft) []*typs.Summary {
	var result []*typs.Summary
	for _, n := range nfts {
		var ts []*typs.Trait
		for _, n := range n.Traits {
			ts = append(ts, &typs.Trait{
				Type:  n.Type,
				Value: n.Value,
			})
		}
		result = append(result, &typs.Summary{
			Blockchain:      n.Blockchain,
			CollectionName:  n.CollectionName,
			CollectionSlug:  n.CollectionSlug,
			ContractAddress: n.ContractAddress,
			TokenId:         n.TokenId,
			Name:            n.Name,
			Description:     n.Description,
			Image:           n.Image,
			AnimationUrl:    n.AnimationUrl,
			OwnerAddresses:  n.OwnerAddresses,
			Traits:          ts,
			Rarity: &typs.Rarity{
				Score: float32(n.Rarity.Score),
				Rank:  int32(n.Rarity.Rank),
				Total: int32(n.Rarity.Total),
			},
			Registered: n.Registered,
		})
	}
	return result
}

func CoverComicNftToHttpDto(nfts ...entity.ComicNft) []*typs.ComicNft {
	var result []*typs.ComicNft
	for _, n := range nfts {
		result = append(result, &typs.ComicNft{
			TokenId: n.TokenId,
			Summary: &typs.ComicWork{
				OriginNft:          CoverNftToHttpDto(&n.Comic.Origin)[0],
				MintLimit:          n.Comic.MintLimit,
				MintPrice:          float32(n.Comic.MintPrice),
				Name:               n.Comic.Name,
				CreatedAtTimestamp: uint32(n.Comic.CreatedAt.Unix()),
				MinterAddress:      n.Comic.UserAddress,
				ImageUrl:           n.Comic.ImageUris,
			},
		})
	}
	return result
}
