package main

import (
	"fmt"
)

type AssetReader interface {
	download(assetLocation string) []*byte
	downloadAssets(assetList array, assetTypeLocation map)
}
