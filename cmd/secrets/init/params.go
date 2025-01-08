package init

import (
	"errors"

	command "github.com/The-Cipher-Protocol/cipher-node/cmd"
	"github.com/The-Cipher-Protocol/cipher-node/secrets"
	"github.com/The-Cipher-Protocol/cipher-node/secrets/helper"
)

const (
	dataDirFlag = "datadir"
	configFlag  = "config"
	ecdsaFlag   = "ecdsa"
	//blsFlag                = "bls"
	networkFlag = "network"
	numFlag     = "num"
	//insecureLocalStoreFlag = "insecure"
)

var (
	errInvalidConfig   = errors.New("invalid secrets configuration")
	errInvalidParams   = errors.New("no config file or data directory passed in")
	errUnsupportedType = errors.New("unsupported secrets manager")
)

type initParams struct {
	dataDir          string
	configPath       string
	generatesECDSA   bool
	generatesBLS     bool
	generatesNetwork bool
	//insecureLocalStore bool

	secretsManager secrets.SecretsManager
	secretsConfig  *secrets.SecretsManagerConfig
}

func (ip *initParams) validateFlags() error {
	if ip.dataDir == "" && ip.configPath == "" {
		return errInvalidParams
	}

	return nil
}

func (ip *initParams) initSecrets() error {
	if err := ip.initSecretsManager(); err != nil {
		return err
	}

	if err := ip.initValidatorKey(); err != nil {
		return err
	}

	return ip.initNetworkingKey()
}

func (ip *initParams) initSecretsManager() error {
	var err error
	if ip.hasConfigPath() {
		if err = ip.parseConfig(); err != nil {
			return err
		}

		ip.secretsManager, err = helper.InitCloudSecretsManager(ip.secretsConfig)

		return err
	}

	return ip.initLocalSecretsManager()
}

func (ip *initParams) hasConfigPath() bool {
	return ip.configPath != ""
}

func (ip *initParams) parseConfig() error {
	secretsConfig, readErr := secrets.ReadConfig(ip.configPath)
	if readErr != nil {
		return errInvalidConfig
	}

	if !secrets.SupportedServiceManager(secretsConfig.Type) {
		return errUnsupportedType
	}

	ip.secretsConfig = secretsConfig

	return nil
}

func (ip *initParams) initLocalSecretsManager() error {
	// if !ip.insecureLocalStore {
	// 	//Storing secrets on a local file system should only be allowed with --insecure flag,
	// 	//to raise awareness that it should be only used in development/testing environments.
	// 	//Production setups should use one of the supported secrets managers
	// 	return errSecureLocalStoreNotImplemented
	// }

	// setup local secrets manager
	local, err := helper.SetupLocalSecretsManager(ip.dataDir)
	if err != nil {
		return err
	}

	ip.secretsManager = local

	return nil
}

func (ip *initParams) initValidatorKey() error {
	var err error

	if ip.generatesECDSA {
		if _, err = helper.InitECDSAValidatorKey(ip.secretsManager); err != nil {
			return err
		}
	}

	if ip.generatesBLS {
		if _, err = helper.InitBLSValidatorKey(ip.secretsManager); err != nil {
			return err
		}
	}

	return nil
}

func (ip *initParams) initNetworkingKey() error {
	if ip.generatesNetwork {
		if _, err := helper.InitNetworkingPrivateKey(ip.secretsManager); err != nil {
			return err
		}
	}

	return nil
}

// getResult gets keys from secret manager and return result to display
func (ip *initParams) getResult() (command.CommandResult, error) {
	var (
		res = &SecretsInitResult{}
		err error
	)

	if res.Address, err = helper.LoadValidatorAddress(ip.secretsManager); err != nil {
		return nil, err
	}
	// LoadValidatorKey and convert the private key to a hex string
	if ValidatorKey, err := helper.LoadValidatorKey(ip.secretsManager); err != nil {
		return nil, err
	} else if ValidatorKey != nil {
		res.ValidatorKey = helper.ConvertPrivateKeyToHex(ValidatorKey)
	}
	if res.NodeID, err = helper.LoadNodeID(ip.secretsManager); err != nil {
		return nil, err
	}

	return res, nil
}
