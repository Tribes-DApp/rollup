package contract_usecase

import (
	"context"

	"github.com/tribeshq/tribes/internal/domain/entity"
)

type FindContractBySymbolInputDTO struct {
	Symbol string `json:"symbol"`
}

type FindContractBySymbolUseCase struct {
	ContractRepository entity.ContractRepository
}

func NewFindContractBySymbolUseCase(contractRepository entity.ContractRepository) *FindContractBySymbolUseCase {
	return &FindContractBySymbolUseCase{
		ContractRepository: contractRepository,
	}
}

func (s *FindContractBySymbolUseCase) Execute(ctx context.Context, input *FindContractBySymbolInputDTO) (*FindContractOutputDTO, error) {
	res, err := s.ContractRepository.FindContractBySymbol(ctx, input.Symbol)
	if err != nil {
		return nil, err
	}
	return &FindContractOutputDTO{
		Id:        res.Id,
		Symbol:    res.Symbol,
		Address:   res.Address,
		CreatedAt: res.CreatedAt,
		UpdatedAt: res.UpdatedAt,
	}, nil
}
