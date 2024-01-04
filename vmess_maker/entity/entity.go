package entity

type VmessJson struct {
	Log struct {
		Loglevel string `json:"loglevel"`
	} `json:"log"`
	Inbounds []Inbound `json:"inbounds"`
	Routing  struct {
		DomainStrategy string `json:"domainStrategy"`
		Rules          []Rule `json:"rules"`
	} `json:"routing"`
	Outbounds []Outbound `json:"outbounds"`
}

type Rule struct {
	Type        string   `json:"type"`
	IP          []string `json:"ip,omitempty"`
	OutboundTag string   `json:"outboundTag"`
	Domain      []string `json:"domain,omitempty"`
}

type Outbound struct {
	Tag      string `json:"tag"`
	Protocol string `json:"protocol"`
	Settings struct {
	} `json:"settings"`
}

type Inbound struct {
	Listen   interface{} `json:"listen"`
	Port     int         `json:"port"`
	Protocol string      `json:"protocol"`
	Settings struct {
		Clients []Client `json:"clients"`
	} `json:"settings"`
	Sniffing struct {
		DestOverride []string `json:"destOverride"`
		Enabled      bool     `json:"enabled"`
	} `json:"sniffing"`
	StreamSettings struct {
		Network  string `json:"network"`
		Security string `json:"security"`
		Sockopt  struct {
			AcceptProxyProtocol bool   `json:"acceptProxyProtocol"`
			Mark                int    `json:"mark"`
			TCPFastOpen         bool   `json:"tcpFastOpen"`
			Tproxy              string `json:"tproxy"`
		} `json:"sockopt"`
		TCPSettings struct {
			AcceptProxyProtocol bool `json:"acceptProxyProtocol"`
			Header              struct {
				Request struct {
					Headers struct {
						Host []string `json:"Host"`
					} `json:"headers"`
					Method string   `json:"method"`
					Path   []string `json:"path"`
				} `json:"request"`
				Response struct {
					Headers struct {
						Connection    []string `json:"Connection"`
						ContentLength []string `json:"Content-Length"`
						ContentType   []string `json:"Content-Type"`
					} `json:"headers"`
					Reason  string `json:"reason"`
					Status  string `json:"status"`
					Version string `json:"version"`
				} `json:"response"`
				Type string `json:"type"`
			} `json:"header"`
		} `json:"tcpSettings"`
	} `json:"streamSettings"`
	Tag string `json:"tag"`
}

type Client struct {
	Email string `json:"email"`
	ID    string `json:"id"`
}

type Setting struct {
	Port                   int      `json:"port"`
	BotToken               string   `json:"bot_token"`
	ChatID                 string   `json:"chat_id"`
	DynamicSubscription    bool     `json:"dynamic_subscription"`
	ChannelName            string   `json:"channel_name"`
	SendVNstat             bool     `json:"send_vnstat"`
	AggregateSubscriptions []string `json:"aggregate_subscriptions"`
	SendSubscriptions      bool     `json:"send_subscriptions"`
	SendConfiguration      string   `json:"send_configuration"`
	RandomHeader           bool     `json:"random_header"`
}
