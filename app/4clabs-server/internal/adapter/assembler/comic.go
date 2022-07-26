package assembler

import (
	typs "4clabs-server/api/nft/v1"
	"4clabs-server/app/4clabs-server/internal/domain/entity"
)

func CoverComicNftToHttpDtos(nfts ...entity.ComicNft) []*typs.ComicNft {
	var result []*typs.ComicNft
	for _, n := range nfts {
		result = append(result, CoverComicNftToHttpDto(n))
	}
	return result
}

func CoverComicNftToHttpDto(nft entity.ComicNft) *typs.ComicNft {
	return &typs.ComicNft{
		TokenId: nft.TokenId,
		Owner:   nft.Owner,
		Summary: &typs.ComicWork{
			OriginNft:          CoverNftToHttpDto(&nft.Comic.Origin)[0],
			MintLimit:          nft.Comic.MintLimit,
			MintPrice:          float32(nft.Comic.MintPrice),
			Name:               nft.Comic.Name,
			CreatedAtTimestamp: uint32(nft.Comic.CreatedAt.Unix()),
			MinterAddress:      nft.Comic.UserAddress,
			ImageUrl:           nft.Comic.ImageUris,
		},
	}
}
