// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package bindings

import (
	"encoding/json"

	"github.com/ethereum-optimism/optimism/op-bindings/solc"
)

const SystemConfigStorageLayoutJSON = "{\"storage\":[{\"astId\":1000,\"contract\":\"contracts/L1/SystemConfig.sol:SystemConfig\",\"label\":\"_initialized\",\"offset\":0,\"slot\":\"0\",\"type\":\"t_uint8\"},{\"astId\":1001,\"contract\":\"contracts/L1/SystemConfig.sol:SystemConfig\",\"label\":\"_initializing\",\"offset\":1,\"slot\":\"0\",\"type\":\"t_bool\"},{\"astId\":1002,\"contract\":\"contracts/L1/SystemConfig.sol:SystemConfig\",\"label\":\"__gap\",\"offset\":0,\"slot\":\"1\",\"type\":\"t_array(t_uint256)1011_storage\"},{\"astId\":1003,\"contract\":\"contracts/L1/SystemConfig.sol:SystemConfig\",\"label\":\"_owner\",\"offset\":0,\"slot\":\"51\",\"type\":\"t_address\"},{\"astId\":1004,\"contract\":\"contracts/L1/SystemConfig.sol:SystemConfig\",\"label\":\"__gap\",\"offset\":0,\"slot\":\"52\",\"type\":\"t_array(t_uint256)1010_storage\"},{\"astId\":1005,\"contract\":\"contracts/L1/SystemConfig.sol:SystemConfig\",\"label\":\"overhead\",\"offset\":0,\"slot\":\"101\",\"type\":\"t_uint256\"},{\"astId\":1006,\"contract\":\"contracts/L1/SystemConfig.sol:SystemConfig\",\"label\":\"scalar\",\"offset\":0,\"slot\":\"102\",\"type\":\"t_uint256\"},{\"astId\":1007,\"contract\":\"contracts/L1/SystemConfig.sol:SystemConfig\",\"label\":\"batcherHash\",\"offset\":0,\"slot\":\"103\",\"type\":\"t_bytes32\"},{\"astId\":1008,\"contract\":\"contracts/L1/SystemConfig.sol:SystemConfig\",\"label\":\"unsafeBlockSigner\",\"offset\":0,\"slot\":\"104\",\"type\":\"t_address\"},{\"astId\":1009,\"contract\":\"contracts/L1/SystemConfig.sol:SystemConfig\",\"label\":\"gasLimit\",\"offset\":20,\"slot\":\"104\",\"type\":\"t_uint64\"}],\"types\":{\"t_address\":{\"encoding\":\"inplace\",\"label\":\"address\",\"numberOfBytes\":\"20\"},\"t_array(t_uint256)1010_storage\":{\"encoding\":\"inplace\",\"label\":\"uint256[49]\",\"numberOfBytes\":\"1568\"},\"t_array(t_uint256)1011_storage\":{\"encoding\":\"inplace\",\"label\":\"uint256[50]\",\"numberOfBytes\":\"1600\"},\"t_bool\":{\"encoding\":\"inplace\",\"label\":\"bool\",\"numberOfBytes\":\"1\"},\"t_bytes32\":{\"encoding\":\"inplace\",\"label\":\"bytes32\",\"numberOfBytes\":\"32\"},\"t_uint256\":{\"encoding\":\"inplace\",\"label\":\"uint256\",\"numberOfBytes\":\"32\"},\"t_uint64\":{\"encoding\":\"inplace\",\"label\":\"uint64\",\"numberOfBytes\":\"8\"},\"t_uint8\":{\"encoding\":\"inplace\",\"label\":\"uint8\",\"numberOfBytes\":\"1\"}}}"

var SystemConfigStorageLayout = new(solc.StorageLayout)

var SystemConfigDeployedBin = "0x608060405234801561001057600080fd5b50600436106101005760003560e01c8063935f029e11610097578063f2fde38b11610066578063f2fde38b1461022e578063f45e65d814610241578063f68016b71461024a578063ffa1ad741461027657600080fd5b8063935f029e146101ec578063b40a817c146101ff578063c9b26f6114610212578063e81b2c6d1461022557600080fd5b806354fd4d50116100d357806354fd4d501461019e578063715018a6146101b35780638da5cb5b146101bb5780638f974d7f146101d957600080fd5b80630c18c1621461010557806318d13918146101215780631fd19ee11461013657806329477e861461017b575b600080fd5b61010e60655481565b6040519081526020015b60405180910390f35b61013461012f366004610cb2565b61027e565b005b6068546101569073ffffffffffffffffffffffffffffffffffffffff1681565b60405173ffffffffffffffffffffffffffffffffffffffff9091168152602001610118565b610185627a120081565b60405167ffffffffffffffff9091168152602001610118565b6101a6610324565b6040516101189190610d4e565b6101346103c7565b60335473ffffffffffffffffffffffffffffffffffffffff16610156565b6101346101e7366004610d79565b6103db565b6101346101fa366004610dd8565b61068c565b61013461020d366004610dfa565b610725565b610134610220366004610e15565b610817565b61010e60675481565b61013461023c366004610cb2565b610847565b61010e60665481565b6068546101859074010000000000000000000000000000000000000000900467ffffffffffffffff1681565b61010e600081565b61028661091a565b606880547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff831690811790915560408051602080820193909352815180820390930183528101905260035b60007f1d2b0bda21d56b8bd12d4f94ebacffdfb35f5e226f84b461103bb8beab6353be836040516103189190610d4e565b60405180910390a35050565b606061034f7f000000000000000000000000000000000000000000000000000000000000000061099b565b6103787f000000000000000000000000000000000000000000000000000000000000000061099b565b6103a17f000000000000000000000000000000000000000000000000000000000000000061099b565b6040516020016103b393929190610e2e565b604051602081830303815290604052905090565b6103cf61091a565b6103d96000610ad8565b565b600054610100900460ff16158080156103fb5750600054600160ff909116105b806104155750303b158015610415575060005460ff166001145b6104a6576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201527f647920696e697469616c697a656400000000000000000000000000000000000060648201526084015b60405180910390fd5b600080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00166001179055801561050457600080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00ff166101001790555b627a120067ffffffffffffffff8416101561057b576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601f60248201527f53797374656d436f6e6669673a20676173206c696d697420746f6f206c6f7700604482015260640161049d565b610583610b4f565b61058c87610847565b606586905560668590556067849055606880547fffffffff00000000000000000000000000000000000000000000000000000000167401000000000000000000000000000000000000000067ffffffffffffffff8616027fffffffffffffffffffffffff0000000000000000000000000000000000000000161773ffffffffffffffffffffffffffffffffffffffff8416179055801561068357600080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00ff169055604051600181527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15b50505050505050565b61069461091a565b606582905560668190556040805160208101849052908101829052600090606001604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08184030181529190529050600160007f1d2b0bda21d56b8bd12d4f94ebacffdfb35f5e226f84b461103bb8beab6353be836040516107189190610d4e565b60405180910390a3505050565b61072d61091a565b627a120067ffffffffffffffff821610156107a4576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601f60248201527f53797374656d436f6e6669673a20676173206c696d697420746f6f206c6f7700604482015260640161049d565b606880547fffffffff0000000000000000ffffffffffffffffffffffffffffffffffffffff167401000000000000000000000000000000000000000067ffffffffffffffff84169081029190911790915560408051602080820193909352815180820390930183528101905260026102e7565b61081f61091a565b60678190556040805160208082018490528251808303909101815290820190915260006102e7565b61084f61091a565b73ffffffffffffffffffffffffffffffffffffffff81166108f2576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201527f6464726573730000000000000000000000000000000000000000000000000000606482015260840161049d565b6108fb81610ad8565b50565b73ffffffffffffffffffffffffffffffffffffffff163b151590565b60335473ffffffffffffffffffffffffffffffffffffffff1633146103d9576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e6572604482015260640161049d565b6060816000036109de57505060408051808201909152600181527f3000000000000000000000000000000000000000000000000000000000000000602082015290565b8160005b8115610a0857806109f281610ed3565b9150610a019050600a83610f3a565b91506109e2565b60008167ffffffffffffffff811115610a2357610a23610f4e565b6040519080825280601f01601f191660200182016040528015610a4d576020820181803683370190505b5090505b8415610ad057610a62600183610f7d565b9150610a6f600a86610f94565b610a7a906030610fa8565b60f81b818381518110610a8f57610a8f610fc0565b60200101907effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916908160001a905350610ac9600a86610f3a565b9450610a51565b949350505050565b6033805473ffffffffffffffffffffffffffffffffffffffff8381167fffffffffffffffffffffffff0000000000000000000000000000000000000000831681179093556040519116919082907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e090600090a35050565b600054610100900460ff16610be6576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602b60248201527f496e697469616c697a61626c653a20636f6e7472616374206973206e6f74206960448201527f6e697469616c697a696e67000000000000000000000000000000000000000000606482015260840161049d565b6103d9600054610100900460ff16610c80576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602b60248201527f496e697469616c697a61626c653a20636f6e7472616374206973206e6f74206960448201527f6e697469616c697a696e67000000000000000000000000000000000000000000606482015260840161049d565b6103d933610ad8565b803573ffffffffffffffffffffffffffffffffffffffff81168114610cad57600080fd5b919050565b600060208284031215610cc457600080fd5b610ccd82610c89565b9392505050565b60005b83811015610cef578181015183820152602001610cd7565b83811115610cfe576000848401525b50505050565b60008151808452610d1c816020860160208601610cd4565b601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0169290920160200192915050565b602081526000610ccd6020830184610d04565b803567ffffffffffffffff81168114610cad57600080fd5b60008060008060008060c08789031215610d9257600080fd5b610d9b87610c89565b9550602087013594506040870135935060608701359250610dbe60808801610d61565b9150610dcc60a08801610c89565b90509295509295509295565b60008060408385031215610deb57600080fd5b50508035926020909101359150565b600060208284031215610e0c57600080fd5b610ccd82610d61565b600060208284031215610e2757600080fd5b5035919050565b60008451610e40818460208901610cd4565b80830190507f2e000000000000000000000000000000000000000000000000000000000000008082528551610e7c816001850160208a01610cd4565b60019201918201528351610e97816002840160208801610cd4565b0160020195945050505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b60007fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8203610f0457610f04610ea4565b5060010190565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601260045260246000fd5b600082610f4957610f49610f0b565b500490565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b600082821015610f8f57610f8f610ea4565b500390565b600082610fa357610fa3610f0b565b500690565b60008219821115610fbb57610fbb610ea4565b500190565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fdfea164736f6c634300080f000a"

func init() {
	if err := json.Unmarshal([]byte(SystemConfigStorageLayoutJSON), SystemConfigStorageLayout); err != nil {
		panic(err)
	}

	layouts["SystemConfig"] = SystemConfigStorageLayout
	deployedBytecodes["SystemConfig"] = SystemConfigDeployedBin
}
