package assembler

import (
	typs "4clabs-server/api/nft/v1"
	"4clabs-server/app/4clabs-server/internal/domain/entity"
)

func CoverNftToHttpDto(nfts ...entity.Nft) []*typs.Summary {
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
		})
	}
	return result
}
