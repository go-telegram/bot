package models

// TonWalletCard https://core.telegram.org/bots/api#tonwalletcard
type TonWalletCard struct {
	Address              string                 `json:"address"`
	Amount               string                 `json:"amount"`
	Stickers             []TonWalletCardSticker `json:"stickers"`
	BackgroundNftAddress string                 `json:"background_nft_address,omitempty"`
}

// TonWalletCardSticker https://core.telegram.org/bots/api#tonwalletcardsticker
type TonWalletCardSticker struct {
	Position AreaPosition `json:"position"`
	Nft      TonNft       `json:"nft"`
}

// TonNft https://core.telegram.org/bots/api#tonnft
type TonNft struct {
	Address           string            `json:"address"`
	CollectionAddress string            `json:"collection_address,omitempty"`
	Index             string            `json:"index"`
	Name              string            `json:"name"`
	Description       string            `json:"description"`
	Attributes        []TonNftAttribute `json:"attributes,omitempty"`
	Gift              *UniqueGift       `json:"gift,omitempty"`
}

// TonNftAttribute https://core.telegram.org/bots/api#tonnftattribute
type TonNftAttribute struct {
	TraitType string `json:"trait_type"`
	Value     string `json:"value"`
}

// TonWalletTransfer https://core.telegram.org/bots/api#tonwallettransfer
type TonWalletTransfer struct {
	TransactionID string `json:"transaction_id"`
	PeerAddress   string `json:"peer_address"`
	Amount        string `json:"amount"`
	Comment       string `json:"comment,omitempty"`
}
