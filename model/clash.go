package model

type ClashType int

const (
	Clash ClashType = 1 + iota
	ClashMeta
)

func GetSupportProxyTypes(clashType ClashType) map[string]bool {
	if clashType == Clash {
		return map[string]bool{
			"ss":     true,
			"ssr":    true,
			"vmess":  true,
			"trojan": true,
			"socks5": true,
			"anytls": true,
		}
	}
	if clashType == ClashMeta {
		return map[string]bool{
			"ss":        true,
			"ssr":       true,
			"vmess":     true,
			"trojan":    true,
			"vless":     true,
			"hysteria":  true,
			"hysteria2": true,
			"socks5":    true,
			"anytls": 	 true,
		}
	}
	return nil
}
