// Code generated by "enumer -type=ActiveValidatorColumnName -json -transform=snake-upper -output=./activevalidatorcolumnname.gen.go"; DO NOT EDIT.

package enums

import (
	"encoding/json"
	"fmt"
	"strings"
)

const _ActiveValidatorColumnNameName = "ACTIVE_VALIDATOR_COLUMN_NAME_NAMEACTIVE_VALIDATOR_COLUMN_NAME_NET_ADDRESSACTIVE_VALIDATOR_COLUMN_NAME_VOTING_POWERACTIVE_VALIDATOR_COLUMN_NAME_GAS_PRICEACTIVE_VALIDATOR_COLUMN_NAME_COMMISSION_RATEACTIVE_VALIDATOR_COLUMN_NAME_NEXT_EPOCH_STAKEACTIVE_VALIDATOR_COLUMN_NAME_NEXT_EPOCH_GAS_PRICEACTIVE_VALIDATOR_COLUMN_NAME_NEXT_EPOCH_COMMISSION_RATEACTIVE_VALIDATOR_COLUMN_NAME_STAKING_POOL_SUI_BALANCEACTIVE_VALIDATOR_COLUMN_NAME_REWARDS_POOLACTIVE_VALIDATOR_COLUMN_NAME_POOL_TOKEN_BALANCEACTIVE_VALIDATOR_COLUMN_NAME_PENDING_STAKE"

var _ActiveValidatorColumnNameIndex = [...]uint16{0, 33, 73, 114, 152, 196, 241, 290, 345, 398, 439, 486, 528}

const _ActiveValidatorColumnNameLowerName = "active_validator_column_name_nameactive_validator_column_name_net_addressactive_validator_column_name_voting_poweractive_validator_column_name_gas_priceactive_validator_column_name_commission_rateactive_validator_column_name_next_epoch_stakeactive_validator_column_name_next_epoch_gas_priceactive_validator_column_name_next_epoch_commission_rateactive_validator_column_name_staking_pool_sui_balanceactive_validator_column_name_rewards_poolactive_validator_column_name_pool_token_balanceactive_validator_column_name_pending_stake"

func (i ActiveValidatorColumnName) String() string {
	if i < 0 || i >= ActiveValidatorColumnName(len(_ActiveValidatorColumnNameIndex)-1) {
		return fmt.Sprintf("ActiveValidatorColumnName(%d)", i)
	}
	return _ActiveValidatorColumnNameName[_ActiveValidatorColumnNameIndex[i]:_ActiveValidatorColumnNameIndex[i+1]]
}

// An "invalid array index" compiler error signifies that the constant values have changed.
// Re-run the stringer command to generate them again.
func _ActiveValidatorColumnNameNoOp() {
	var x [1]struct{}
	_ = x[ActiveValidatorColumnNameName-(0)]
	_ = x[ActiveValidatorColumnNameNetAddress-(1)]
	_ = x[ActiveValidatorColumnNameVotingPower-(2)]
	_ = x[ActiveValidatorColumnNameGasPrice-(3)]
	_ = x[ActiveValidatorColumnNameCommissionRate-(4)]
	_ = x[ActiveValidatorColumnNameNextEpochStake-(5)]
	_ = x[ActiveValidatorColumnNameNextEpochGasPrice-(6)]
	_ = x[ActiveValidatorColumnNameNextEpochCommissionRate-(7)]
	_ = x[ActiveValidatorColumnNameStakingPoolSuiBalance-(8)]
	_ = x[ActiveValidatorColumnNameRewardsPool-(9)]
	_ = x[ActiveValidatorColumnNamePoolTokenBalance-(10)]
	_ = x[ActiveValidatorColumnNamePendingStake-(11)]
}

var _ActiveValidatorColumnNameValues = []ActiveValidatorColumnName{ActiveValidatorColumnNameName, ActiveValidatorColumnNameNetAddress, ActiveValidatorColumnNameVotingPower, ActiveValidatorColumnNameGasPrice, ActiveValidatorColumnNameCommissionRate, ActiveValidatorColumnNameNextEpochStake, ActiveValidatorColumnNameNextEpochGasPrice, ActiveValidatorColumnNameNextEpochCommissionRate, ActiveValidatorColumnNameStakingPoolSuiBalance, ActiveValidatorColumnNameRewardsPool, ActiveValidatorColumnNamePoolTokenBalance, ActiveValidatorColumnNamePendingStake}

var _ActiveValidatorColumnNameNameToValueMap = map[string]ActiveValidatorColumnName{
	_ActiveValidatorColumnNameName[0:33]:         ActiveValidatorColumnNameName,
	_ActiveValidatorColumnNameLowerName[0:33]:    ActiveValidatorColumnNameName,
	_ActiveValidatorColumnNameName[33:73]:        ActiveValidatorColumnNameNetAddress,
	_ActiveValidatorColumnNameLowerName[33:73]:   ActiveValidatorColumnNameNetAddress,
	_ActiveValidatorColumnNameName[73:114]:       ActiveValidatorColumnNameVotingPower,
	_ActiveValidatorColumnNameLowerName[73:114]:  ActiveValidatorColumnNameVotingPower,
	_ActiveValidatorColumnNameName[114:152]:      ActiveValidatorColumnNameGasPrice,
	_ActiveValidatorColumnNameLowerName[114:152]: ActiveValidatorColumnNameGasPrice,
	_ActiveValidatorColumnNameName[152:196]:      ActiveValidatorColumnNameCommissionRate,
	_ActiveValidatorColumnNameLowerName[152:196]: ActiveValidatorColumnNameCommissionRate,
	_ActiveValidatorColumnNameName[196:241]:      ActiveValidatorColumnNameNextEpochStake,
	_ActiveValidatorColumnNameLowerName[196:241]: ActiveValidatorColumnNameNextEpochStake,
	_ActiveValidatorColumnNameName[241:290]:      ActiveValidatorColumnNameNextEpochGasPrice,
	_ActiveValidatorColumnNameLowerName[241:290]: ActiveValidatorColumnNameNextEpochGasPrice,
	_ActiveValidatorColumnNameName[290:345]:      ActiveValidatorColumnNameNextEpochCommissionRate,
	_ActiveValidatorColumnNameLowerName[290:345]: ActiveValidatorColumnNameNextEpochCommissionRate,
	_ActiveValidatorColumnNameName[345:398]:      ActiveValidatorColumnNameStakingPoolSuiBalance,
	_ActiveValidatorColumnNameLowerName[345:398]: ActiveValidatorColumnNameStakingPoolSuiBalance,
	_ActiveValidatorColumnNameName[398:439]:      ActiveValidatorColumnNameRewardsPool,
	_ActiveValidatorColumnNameLowerName[398:439]: ActiveValidatorColumnNameRewardsPool,
	_ActiveValidatorColumnNameName[439:486]:      ActiveValidatorColumnNamePoolTokenBalance,
	_ActiveValidatorColumnNameLowerName[439:486]: ActiveValidatorColumnNamePoolTokenBalance,
	_ActiveValidatorColumnNameName[486:528]:      ActiveValidatorColumnNamePendingStake,
	_ActiveValidatorColumnNameLowerName[486:528]: ActiveValidatorColumnNamePendingStake,
}

var _ActiveValidatorColumnNameNames = []string{
	_ActiveValidatorColumnNameName[0:33],
	_ActiveValidatorColumnNameName[33:73],
	_ActiveValidatorColumnNameName[73:114],
	_ActiveValidatorColumnNameName[114:152],
	_ActiveValidatorColumnNameName[152:196],
	_ActiveValidatorColumnNameName[196:241],
	_ActiveValidatorColumnNameName[241:290],
	_ActiveValidatorColumnNameName[290:345],
	_ActiveValidatorColumnNameName[345:398],
	_ActiveValidatorColumnNameName[398:439],
	_ActiveValidatorColumnNameName[439:486],
	_ActiveValidatorColumnNameName[486:528],
}

// ActiveValidatorColumnNameString retrieves an enum value from the enum constants string name.
// Throws an error if the param is not part of the enum.
func ActiveValidatorColumnNameString(s string) (ActiveValidatorColumnName, error) {
	if val, ok := _ActiveValidatorColumnNameNameToValueMap[s]; ok {
		return val, nil
	}

	if val, ok := _ActiveValidatorColumnNameNameToValueMap[strings.ToLower(s)]; ok {
		return val, nil
	}
	return 0, fmt.Errorf("%s does not belong to ActiveValidatorColumnName values", s)
}

// ActiveValidatorColumnNameValues returns all values of the enum
func ActiveValidatorColumnNameValues() []ActiveValidatorColumnName {
	return _ActiveValidatorColumnNameValues
}

// ActiveValidatorColumnNameStrings returns a slice of all String values of the enum
func ActiveValidatorColumnNameStrings() []string {
	strs := make([]string, len(_ActiveValidatorColumnNameNames))
	copy(strs, _ActiveValidatorColumnNameNames)
	return strs
}

// IsAActiveValidatorColumnName returns "true" if the value is listed in the enum definition. "false" otherwise
func (i ActiveValidatorColumnName) IsAActiveValidatorColumnName() bool {
	for _, v := range _ActiveValidatorColumnNameValues {
		if i == v {
			return true
		}
	}
	return false
}

// MarshalJSON implements the json.Marshaler interface for ActiveValidatorColumnName
func (i ActiveValidatorColumnName) MarshalJSON() ([]byte, error) {
	return json.Marshal(i.String())
}

// UnmarshalJSON implements the json.Unmarshaler interface for ActiveValidatorColumnName
func (i *ActiveValidatorColumnName) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return fmt.Errorf("ActiveValidatorColumnName should be a string, got %s", data)
	}

	var err error
	*i, err = ActiveValidatorColumnNameString(s)
	return err
}
