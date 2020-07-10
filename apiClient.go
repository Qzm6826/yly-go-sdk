package ylyOpenApi

type ApiClient struct {
	PrintService PrintService
	SetPrinter   SetPrinter
	config       Config
}

type PrintService struct {
	config *Config
}

type SetPrinter struct {
	config *Config
}

func NewClient(config Config) ApiClient{
	client := ApiClient{}
	client.SetConfig(config)
	return client
}

func (client *ApiClient) SetConfig(config Config) {
	client.config = config
	client.PrintService.config = &client.config
	client.SetPrinter.config = &client.config
}
