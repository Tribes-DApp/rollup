package inspect_handler

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/rollmelette/rollmelette"
	"github.com/tribeshq/tribes/internal/domain/entity"
	"github.com/tribeshq/tribes/internal/usecase/contract_usecase"
	"github.com/tribeshq/tribes/pkg/custom_type"
	"github.com/tribeshq/tribes/pkg/router"
)

type ContractInspectHandlers struct {
	ContractRepository entity.ContractRepository
}

func NewContractInspectHandlers(contractRepository entity.ContractRepository) *ContractInspectHandlers {
	return &ContractInspectHandlers{
		ContractRepository: contractRepository,
	}
}

func (h *ContractInspectHandlers) FindAllContractsHandler(ctx context.Context, env rollmelette.EnvInspector) error {
	findAllContracts := contract_usecase.NewFindAllContractsUseCase(h.ContractRepository)
	contracts, err := findAllContracts.Execute(ctx)
	if err != nil {
		return fmt.Errorf("failed to find all contracts: %w", err)
	}
	contractsBytes, err := json.Marshal(contracts)
	if err != nil {
		return fmt.Errorf("failed to marshal contracts: %w", err)
	}
	env.Report(contractsBytes)
	return nil
}

func (h *ContractInspectHandlers) FindContractBySymbolHandler(ctx context.Context, env rollmelette.EnvInspector) error {
	symbol := strings.ToUpper(router.PathValue(ctx, "symbol"))
	findOrderBySymbol := contract_usecase.NewFindContractBySymbolUseCase(h.ContractRepository)
	contract, err := findOrderBySymbol.Execute(ctx, &contract_usecase.FindContractBySymbolInputDTO{
		Symbol: symbol,
	})
	if err != nil {
		return err
	}
	contractBytes, err := json.Marshal(contract)
	if err != nil {
		return fmt.Errorf("failed to marshal contract: %w", err)
	}
	env.Report(contractBytes)
	return nil
}

func (h *ContractInspectHandlers) FindContractByAddressHandler(ctx context.Context, env rollmelette.EnvInspector) error {
	findOrderByAddress := contract_usecase.NewFindContractByAddressUseCase(h.ContractRepository)
	contract, err := findOrderByAddress.Execute(ctx, &contract_usecase.FindContractByAddressInputDTO{
		Address: custom_type.HexToAddress(router.PathValue(ctx, "address")),
	})
	if err != nil {
		return err
	}
	contractBytes, err := json.Marshal(contract)
	if err != nil {
		return fmt.Errorf("failed to marshal contract: %w", err)
	}
	env.Report(contractBytes)
	return nil
}
