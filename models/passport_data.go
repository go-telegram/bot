package models

// PassportFile https://core.telegram.org/bots/api#passportfile
type PassportFile struct {
	FileID       string `json:"file_id"`
	FileUniqueID string `json:"file_unique_id"`
	FileSize     int    `json:"file_size"`
	FileDate     int    `json:"file_date"`
}

// EncryptedPassportElement https://core.telegram.org/bots/api#encryptedpassportelement
type EncryptedPassportElement struct {
	Type        string         `json:"type"`
	Data        string         `json:"data,omitempty"`
	PhoneNumber string         `json:"phone_number,omitempty"`
	Email       string         `json:"email,omitempty"`
	Files       []PassportFile `json:"files,omitempty"`
	FrontSide   *PassportFile  `json:"front_side,omitempty"`
	ReverseSide *PassportFile  `json:"reverse_side,omitempty"`
	Selfie      *PassportFile  `json:"selfie,omitempty"`
	Translation []PassportFile `json:"translation,omitempty"`
	Hash        string         `json:"hash"`
}

// EncryptedCredentials https://core.telegram.org/bots/api#encryptedcredentials
type EncryptedCredentials struct {
	Data   string `json:"data"`
	Hash   string `json:"hash"`
	Secret string `json:"secret"`
}

// PassportData https://core.telegram.org/bots/api#passportdata
type PassportData struct {
	Data        []EncryptedPassportElement `json:"data"`
	Credentials EncryptedCredentials       `json:"credentials"`
}
