package config

type Config struct {
	ListenAddr         string `envconfig:"listen_addr"`
	InternalListenAddr string `envconfig:"internal_listen_addr"`
	DbUrl              string `envconfig:"db_url"`
	CDNBaseURL         string `envconfig:"cdn_base_url"`
}
