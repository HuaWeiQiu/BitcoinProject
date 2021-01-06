package entity

type AddressType string

const (
	LEGACY   AddressType="legacy"
	P2SH_SEGWIT AddressType="p2sh-segwit"
	BECH32 AddressType="bech32"
)
func AddressTypes(addtype AddressType) string {
	switch addtype {
	case LEGACY:
		return "legacy"
	case P2SH_SEGWIT:
		return "p2sh-segwit"
	case BECH32:
		return "bech32"
	default:return ""
	}
}
