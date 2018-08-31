package config

type Experiments struct {
	FilestoreEnabled     bool
	UrlstoreEnabled      bool
	ShardingEnabled      bool
	Libp2pStreamMounting bool
	QUIC                 bool
	BitswapStrategyEnabled bool
	BitswapStrategy        string
	BitswapRRQRoundBurst   int
}
