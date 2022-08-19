package main

type configModel struct {
	ForwardRules []string `yaml:"rules"`
	ListenAddr   string   `yaml:"listen_addr.omitempty"`
	EnableSocks  bool     `yaml:"enable_socks5,omitempty"`
	SocksAddr    string   `yaml:"socks_addr,omitempty"`
}
