package model

type Anytls struct {
	Type              string `yaml:"type"`
	Name              string `yaml:"name"`
	Server            string `yaml:"server"`
	Port              int    `yaml:"port"`
	Password          string `yaml:"password,omitempty"`
	UDP               bool   `yaml:"udp,omitempty"`
	SNI               string `yaml:"sni,omitempty"`
	SkipCertVerify    bool   `yaml:"skip-cert-verify,omitempty"`
	Fingerprint       string `yaml:"fingerprint,omitempty"`
	ClientFingerprint string `yaml:"client-fingerprint,omitempty"`
}

func ProxyToAnytls(p Proxy) Anytls {
	return Anytls{
		Type:              "anytls",
		Name:              p.Name,
		Server:            p.Server,
		Port:              p.Port,
		Password:          p.Password,
		UDP:               p.UDP,
		SNI:               p.Sni,
		SkipCertVerify:    p.SkipCertVerify,
		ClientFingerprint: p.ClientFingerprint,
	}
}
