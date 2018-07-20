package main

import (
	"github.com/lightningnetwork/lnd/keychain"
	litecoinCfg "github.com/ltcsuite/ltcd/chaincfg"
	einsteiniumCfg "github.com/MatijaMitic/emc2d-chainconfig"
	litecoinWire "github.com/ltcsuite/ltcd/wire"
	"github.com/btcsuite/btcd/chaincfg"
	bitcoinCfg "github.com/btcsuite/btcd/chaincfg"
	"github.com/btcsuite/btcd/chaincfg/chainhash"
	bitcoinWire "github.com/btcsuite/btcd/wire"
)

// activeNetParams is a pointer to the parameters specific to the currently
// active bitcoin network.
var activeNetParams = bitcoinTestNetParams

// bitcoinNetParams couples the p2p parameters of a network with the
// corresponding RPC port of a daemon running on the particular network.
type bitcoinNetParams struct {
	*bitcoinCfg.Params
	rpcPort  string
	CoinType uint32
}

// litecoinNetParams couples the p2p parameters of a network with the
// corresponding RPC port of a daemon running on the particular network.
type litecoinNetParams struct {
	*litecoinCfg.Params
	rpcPort  string
	CoinType uint32
}

// einsteiniumNetParams couples the p2p parameters of a network with the
// corresponding RPC port of a daemon running on the particular network.
type einsteiniumNetParams struct {
	*litecoinCfg.Params //TODO cnahge to to einsteiniumCfg in ltcd/chaincfg
	rpcPort  string
	CoinType uint32
}

// bitcoinTestNetParams contains parameters specific to the 3rd version of the
// test network.
var bitcoinTestNetParams = bitcoinNetParams{
	Params:   &bitcoinCfg.TestNet3Params,
	rpcPort:  "18334",
	CoinType: keychain.CoinTypeTestnet,
}

// bitcoinMainNetParams contains parameters specific to the current Bitcoin
// mainnet.
var bitcoinMainNetParams = bitcoinNetParams{
	Params:   &bitcoinCfg.MainNetParams,
	rpcPort:  "8334",
	CoinType: keychain.CoinTypeBitcoin,
}

// bitcoinSimNetParams contains parameters specific to the simulation test
// network.
var bitcoinSimNetParams = bitcoinNetParams{
	Params:   &bitcoinCfg.SimNetParams,
	rpcPort:  "18556",
	CoinType: keychain.CoinTypeTestnet,
}

// litecoinTestNetParams contains parameters specific to the 4th version of the
// test network.
var litecoinTestNetParams = litecoinNetParams{
	Params:   &litecoinCfg.TestNet4Params,
	rpcPort:  "19334",
	CoinType: keychain.CoinTypeTestnet,
}

// litecoinMainNetParams contains the parameters specific to the current
// Litecoin mainnet.
var litecoinMainNetParams = litecoinNetParams{
	Params:   &litecoinCfg.MainNetParams,
	rpcPort:  "9334",
	CoinType: keychain.CoinTypeLitecoin,
}

// litecoinRegtestParams contains the parameters specific to the current
// Litecoin regtest.
var litecoinRegtestParams = litecoinNetParams{
	Params:   &litecoinCfg.RegressionNetParams,
	rpcPort:  "19334",
	CoinType: keychain.CoinTypeLitecoin,
}

// einsteiniumTestNetParams contains parameters specific to the 4th version of the
// Einsteinium test network.
var einsteiniumTestNetParams = einsteiniumNetParams{
	Params:   &einsteiniumCfg.TestNet4Params,
	rpcPort:  "31876",
	CoinType: keychain.CoinTypeTestnet,
}

// einsteiniumMainNetParams contains the parameters specific to the current
// Einsteinium mainnet.
var einsteiniumMainNetParams = einsteiniumNetParams{
	Params:   &einsteiniumCfg.MainNetParams,
	rpcPort:  "41876",
	CoinType: keychain.CointTypeEinsteinium,
}

// einsteiniumRegtestParams contains the parameters specific to the current
// Einsteinium regtest.
var einsteiniumRegtestParams = einsteiniumNetParams{
	Params:   &einsteiniumCfg.RegressionNetParams,
	rpcPort:  "31882",
	CoinType: keychain.CointTypeEinsteinium,
}

// regTestNetParams contains parameters specific to a local regtest network.
var regTestNetParams = bitcoinNetParams{
	Params:   &bitcoinCfg.RegressionNetParams,
	rpcPort:  "18334",
	CoinType: keychain.CoinTypeTestnet,
}

// applyLitecoinParams applies the relevant chain configuration parameters that
// differ for litecoin to the chain parameters typed for btcsuite derivation.
// This function is used in place of using something like interface{} to
// abstract over _which_ chain (or fork) the parameters are for.
func applyLitecoinParams(params *bitcoinNetParams, litecoinParams *litecoinNetParams) {
	params.Name = litecoinParams.Name
	params.Net = bitcoinWire.BitcoinNet(litecoinParams.Net)
	params.DefaultPort = litecoinParams.DefaultPort
	params.CoinbaseMaturity = litecoinParams.CoinbaseMaturity

	copy(params.GenesisHash[:], litecoinParams.GenesisHash[:])

	// Address encoding magics
	params.PubKeyHashAddrID = litecoinParams.PubKeyHashAddrID
	params.ScriptHashAddrID = litecoinParams.ScriptHashAddrID
	params.PrivateKeyID = litecoinParams.PrivateKeyID
	params.WitnessPubKeyHashAddrID = litecoinParams.WitnessPubKeyHashAddrID
	params.WitnessScriptHashAddrID = litecoinParams.WitnessScriptHashAddrID
	params.Bech32HRPSegwit = litecoinParams.Bech32HRPSegwit

	copy(params.HDPrivateKeyID[:], litecoinParams.HDPrivateKeyID[:])
	copy(params.HDPublicKeyID[:], litecoinParams.HDPublicKeyID[:])

	params.HDCoinType = litecoinParams.HDCoinType

	checkPoints := make([]chaincfg.Checkpoint, len(litecoinParams.Checkpoints))
	for i := 0; i < len(litecoinParams.Checkpoints); i++ {
		var chainHash chainhash.Hash
		copy(chainHash[:], litecoinParams.Checkpoints[i].Hash[:])

		checkPoints[i] = chaincfg.Checkpoint{
			Height: litecoinParams.Checkpoints[i].Height,
			Hash:   &chainHash,
		}
	}
	params.Checkpoints = checkPoints

	params.rpcPort = litecoinParams.rpcPort
	params.CoinType = litecoinParams.CoinType
}

// applyEinsteiniumParams applies the relevant chain configuration parameters that
// differ for einsteinium to the chain parameters typed for btcsuite derivation.
// This function is used in place of using something like interface{} to
// abstract over _which_ chain (or fork) the parameters are for.
func applyEinsteiniumParams(params *bitcoinNetParams, einsteiniumParams *einsteiniumNetParams) {
	params.Name = einsteiniumParams.Name
	params.Net = bitcoinWire.BitcoinNet(einsteiniumParams.Net)
	params.DefaultPort = einsteiniumParams.DefaultPort
	params.CoinbaseMaturity = einsteiniumParams.CoinbaseMaturity

	copy(params.GenesisHash[:], einsteiniumParams.GenesisHash[:])

	// Address encoding magics
	params.PubKeyHashAddrID = einsteiniumParams.PubKeyHashAddrID
	params.ScriptHashAddrID = einsteiniumParams.ScriptHashAddrID
	params.PrivateKeyID = einsteiniumParams.PrivateKeyID
	params.WitnessPubKeyHashAddrID = einsteiniumParams.WitnessPubKeyHashAddrID
	params.WitnessScriptHashAddrID = einsteiniumParams.WitnessScriptHashAddrID
	params.Bech32HRPSegwit = einsteiniumParams.Bech32HRPSegwit

	copy(params.HDPrivateKeyID[:], einsteiniumParams.HDPrivateKeyID[:])
	copy(params.HDPublicKeyID[:], einsteiniumParams.HDPublicKeyID[:])

	params.HDCoinType = einsteiniumParams.HDCoinType

	checkPoints := make([]chaincfg.Checkpoint, len(einsteiniumParams.Checkpoints))
	for i := 0; i < len(einsteiniumParams.Checkpoints); i++ {
		var chainHash chainhash.Hash
		copy(chainHash[:], einsteiniumParams.Checkpoints[i].Hash[:])

		checkPoints[i] = chaincfg.Checkpoint{
			Height: einsteiniumParams.Checkpoints[i].Height,
			Hash:   &chainHash,
		}
	}
	params.Checkpoints = checkPoints

	params.rpcPort = einsteiniumParams.rpcPort
	params.CoinType = einsteiniumParams.CoinType
}

// isTestnet tests if the given params correspond to a testnet
// parameter configuration.
func isTestnet(params *bitcoinNetParams) bool {
	switch params.Params.Net {
	case bitcoinWire.TestNet3, bitcoinWire.BitcoinNet(litecoinWire.TestNet4):
		return true
	default:
		return false
	}
}
