package enums

//go:generate go run github.com/dmarkham/enumer -type=ValidatorColumnName -json -transform=snake-upper -output=./validatorcolumnname.gen.go
type ValidatorColumnName int

const (
	ValidatorColumnNameHealth ValidatorColumnName = iota
	ValidatorColumnNameAddress
	ValidatorColumnNameTotalTransactionCertificates
	ValidatorColumnNameTotalTransactionEffects
	ValidatorColumnNameHighestKnownCheckpoint
	ValidatorColumnNameLastExecutedCheckpoint
	ValidatorColumnNameCheckpointExecBacklog
	ValidatorColumnNameHighestSyncedCheckpoint
	ValidatorColumnNameCheckpointSyncBacklog
	ValidatorColumnNameCurrentEpoch
	ValidatorColumnNameUptime
	ValidatorColumnNameVersion
	ValidatorColumnNameCommit
	ValidatorColumnNameCountry
	ValidatorColumnNameCurrentRound
	ValidatorColumnNameHighestProcessedRound
	ValidatorColumnNameLastCommittedRound
	ValidatorColumnNamePrimaryNetworkPeers
	ValidatorColumnNameWorkerNetworkPeers
	ValidatorColumnNameSkippedConsensusTransactions
	ValidatorColumnNameTotalSignatureErrors
)
