package config

type GRPC interface {
	GRPCAddress() string
}

func (c Config) GRPCAddress() string {
	return c.GetString("grpc.address")
}
