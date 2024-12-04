package user_usecase

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/holiman/uint256"
	"github.com/tribeshq/tribes/internal/domain/entity"
)

type FindUserOutputDTO struct {
	Id                uint                    `json:"id"`
	Role              string                  `json:"role"`
	Address           common.Address          `json:"address"`
	SocialAccounts    []*entity.SocialAccount `json:"social_accounts"`
	InvestmentLimit   *uint256.Int            `json:"investment_limit,omitempty" gorm:"type:bigint"`
	DebtIssuanceLimit *uint256.Int            `json:"debt_issuance_limit,omitempty" gorm:"type:bigint"`
	CreatedAt         int64                   `json:"created_at"`
	UpdatedAt         int64                   `json:"updated_at"`
}
