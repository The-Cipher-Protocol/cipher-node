package server

import (
	"github.com/nnlgsakib/neth/chain"
	"github.com/nnlgsakib/neth/consensus"
	consensusIBFT "github.com/nnlgsakib/neth/consensus/NLG-ibft"
	consensusDev "github.com/nnlgsakib/neth/consensus/dev"
	consensusDummy "github.com/nnlgsakib/neth/consensus/dummy"

	//consensusPolyBFT
	"github.com/nnlgsakib/neth/forkmanager"
	"github.com/nnlgsakib/neth/secrets"
	"github.com/nnlgsakib/neth/secrets/awsssm"
	"github.com/nnlgsakib/neth/secrets/gcpssm"
	"github.com/nnlgsakib/neth/secrets/hashicorpvault"
	"github.com/nnlgsakib/neth/secrets/local"
	"github.com/nnlgsakib/neth/state"
)

type GenesisFactoryHook func(config *chain.Chain, engineName string) func(*state.Transition) error

type ConsensusType string

type ForkManagerFactory func(forks *chain.Forks) error

type ForkManagerInitialParamsFactory func(config *chain.Chain) (*forkmanager.ForkParams, error)

const (
	DevConsensus  ConsensusType = "dev"
	IBFTConsensus ConsensusType = "ibft"
	//PolyBFTConsensus ConsensusType = consensusPolyBFT.ConsensusName
	DummyConsensus ConsensusType = "dummy"
)

var consensusBackends = map[ConsensusType]consensus.Factory{
	DevConsensus:  consensusDev.Factory,
	IBFTConsensus: consensusIBFT.Factory,
	//PolyBFTConsensus: consensusPolyBFT.Factory,
	DummyConsensus: consensusDummy.Factory,
}

// secretsManagerBackends defines the SecretManager factories for different
// secret management solutions
var secretsManagerBackends = map[secrets.SecretsManagerType]secrets.SecretsManagerFactory{
	secrets.Local:          local.SecretsManagerFactory,
	secrets.HashicorpVault: hashicorpvault.SecretsManagerFactory,
	secrets.AWSSSM:         awsssm.SecretsManagerFactory,
	secrets.GCPSSM:         gcpssm.SecretsManagerFactory,
}

var genesisCreationFactory = map[ConsensusType]GenesisFactoryHook{}

var forkManagerInitialParamsFactory = map[ConsensusType]ForkManagerInitialParamsFactory{
	// PolyBFTConsensus: consensusPolyBFT.ForkManagerInitialParamsFactory,
}

func ConsensusSupported(value string) bool {
	_, ok := consensusBackends[ConsensusType(value)]

	return ok
}
