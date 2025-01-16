package business

import "github.com/Pump-Elf-Ranch/per_apollo/config"

type ContractInfoService interface {
	ContractInfos() []*config.ContractInfo
}

type contractInfoService struct {
	cfg *config.Config
}

func NewContractInfoService(cfg *config.Config) ContractInfoService {
	return &contractInfoService{
		cfg: cfg,
	}
}

func (cis *contractInfoService) ContractInfos() []*config.ContractInfo {
	info := cis.cfg.ContractInfos
	return info
}
