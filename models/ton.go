package models

// TonWalletCard https://core.telegram.org/bots/api#tonwalletcard
//
// Describes a TON wallet card, which is shown on the profile page of a user to everyone.
type TonWalletCard struct {
	Address              string                 `json:"address"`
	Amount               string                 `json:"amount"`
	Stickers             []TonWalletCardSticker `json:"stickers"`
	BackgroundNftAddress string                 `json:"background_nft_address,omitempty"`
}

// TonWalletCardSticker https://core.telegram.org/bots/api#tonwalletcardsticker
//
// Describes a sticker on a TON wallet card.
type TonWalletCardSticker struct {
	Position AreaPosition `json:"position"`
	Nft      TonNft       `json:"nft"`
}

// TonNft https://core.telegram.org/bots/api#tonnft
//
// Represents an NFT in the TON blockchain.
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
//
// Represents an attribute of an NFT in the TON blockchain.
type TonNftAttribute struct {
	TraitType string `json:"trait_type"`
	Value     string `json:"value"`
}

// TonWalletTransfer https://core.telegram.org/bots/api#tonwallettransfer
//
// Describes a service message about a transfer with a TON wallet.
type TonWalletTransfer struct {
	TransactionID string `json:"transaction_id"`
	PeerAddress   string `json:"peer_address"`
	Amount        string `json:"amount"`
	Comment       string `json:"comment,omitempty"`
}
