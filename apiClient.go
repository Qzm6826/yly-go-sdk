package ylyOpenApi

type ApiClient struct {
	PrintService PrintService
	SetPrinter   SetPrinter
}

type PrintService struct {
	config *Config
}

type SetPrinter struct {
	config *Config
}
