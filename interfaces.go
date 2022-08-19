package main

import (
	"net"

	"golang.org/x/net/proxy"
)

var proxyDialer proxy.Dialer = nil

func GetDialer(isSocks5 bool) proxy.Dialer {
	if !isSocks5 {
		return &net.Dialer{}
	}
	if proxyDialer != nil {
		return proxyDialer

	}
	var err error
	proxyDialer, err = proxy.SOCKS5("tcp", cfg.SocksAddr, nil, proxy.Direct)
	if err != nil {
		panic(err)
	}
	return proxyDialer
}
