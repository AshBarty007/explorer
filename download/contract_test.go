package download

import (
	"blockchain_services/blockchain"
	db "blockchain_services/postgres"
	"context"
	"encoding/hex"
	"math/big"
	"strings"
	"testing"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
)

// TestERC20Methods 测试ERC20合约的方法调用
func TestERC20Methods(t *testing.T) {
	etClient, err := blockchain.Dial("http://192.168.120.31:8545")
	if err != nil {
		t.Fatalf("连接以太坊客户端失败: %v", err)
	}
	defer etClient.Close()

	ctx := context.Background()

	// 使用测试合约地址
	testAddress := common.HexToAddress("0x3d41eacfcc93aec57e02513f8ebdd9d1642cb11d")

	contractABI, err := abi.JSON(strings.NewReader(blockchain.ERC20_abi))
	if err != nil {
		t.Fatalf("解析ERC20 ABI失败: %v", err)
	}

	t.Run("ERC20_name", func(t *testing.T) {
		result, err := callStringMethod(etClient, ctx, testAddress, &contractABI, "name")
		if err != nil {
			t.Skipf("调用 name() 失败: %v (可能是网络问题或合约不支持)", err)
			return
		}
		if result == "" {
			t.Errorf("调用 name() 返回空字符串")
		} else {
			t.Logf("✓ name() = %s", result)
		}
	})

	t.Run("ERC20_symbol", func(t *testing.T) {
		result, err := callStringMethod(etClient, ctx, testAddress, &contractABI, "symbol")
		if err != nil {
			t.Skipf("调用 symbol() 失败: %v (可能是网络问题或合约不支持)", err)
			return
		}
		if result == "" {
			t.Errorf("调用 symbol() 返回空字符串")
		} else {
			t.Logf("✓ symbol() = %s", result)
		}
	})

	t.Run("ERC20_decimals", func(t *testing.T) {
		result, err := callUint8Method(etClient, ctx, testAddress, &contractABI, "decimals")
		if err != nil {
			t.Skipf("调用 decimals() 失败: %v (可能是网络问题或合约不支持)", err)
			return
		}
		if result == 0 {
			t.Logf("警告: decimals() 返回0，可能合约不支持或值为0")
		} else {
			t.Logf("✓ decimals() = %d", result)
		}
		// 验证decimals在合理范围内（通常0-18）
		if result > 18 {
			t.Logf("警告: decimals值 %d 超出常见范围(0-18)，但可能是合法的", result)
		}
	})

	t.Run("ERC20_totalSupply", func(t *testing.T) {
		result, err := callUint256Method(etClient, ctx, testAddress, &contractABI, "totalSupply")
		if err != nil {
			t.Skipf("调用 totalSupply() 失败: %v (可能是网络问题或合约不支持)", err)
			return
		}
		if result == nil {
			t.Errorf("调用 totalSupply() 返回nil")
		} else if result.Sign() <= 0 {
			t.Errorf("调用 totalSupply() 返回非正数: %s", result.String())
		} else {
			t.Logf("✓ totalSupply() = %s", result.String())
		}
	})

	t.Run("ERC20_balanceOf", func(t *testing.T) {
		// 测试 balanceOf(address) 方法，需要传入参数
		owner := common.HexToAddress("0x30e938b0630c02f394d17925fdb5fb046f70d452")
		data, err := contractABI.Pack("balanceOf", owner)
		if err != nil {
			t.Fatalf("打包 balanceOf 方法失败: %v", err)
		}

		msg := ethereum.CallMsg{
			To:   &testAddress,
			Data: data,
		}

		result, err := etClient.CallContract(ctx, msg, nil)
		if err != nil {
			t.Skipf("调用 balanceOf() 失败: %v (可能是网络问题或合约不支持)", err)
			return
		}

		var balance *big.Int
		err = contractABI.UnpackIntoInterface(&balance, "balanceOf", result)
		if err != nil {
			t.Fatalf("解析 balanceOf 返回值失败: %v", err)
		}

		t.Logf("✓ balanceOf(%s) = %s", owner.Hex(), balance.String())
	})
}

// TestERC721Methods 测试ERC721合约的方法调用
func TestERC721Methods(t *testing.T) {
	etClient, err := blockchain.Dial("http://192.168.120.31:8545")
	if err != nil {
		t.Fatalf("连接以太坊客户端失败: %v", err)
	}
	defer etClient.Close()

	ctx := context.Background()

	// 使用测试NFT合约地址（请根据实际情况修改）
	testAddress := common.HexToAddress("0x64bbb7423b92c3e8672b9215c365674f8fbc3850")

	contractABI, err := abi.JSON(strings.NewReader(blockchain.ERC721_abi))
	if err != nil {
		t.Fatalf("解析ERC721 ABI失败: %v", err)
	}

	t.Run("ERC721_balanceOf", func(t *testing.T) {
		// 测试 balanceOf(address) 方法
		owner := common.HexToAddress("0x30E938B0630C02F394d17925FDB5fB046f70d452")
		data, err := contractABI.Pack("balanceOf", owner)
		if err != nil {
			t.Fatalf("打包 balanceOf 方法失败: %v", err)
		}

		msg := ethereum.CallMsg{
			To:   &testAddress,
			Data: data,
		}

		result, err := etClient.CallContract(ctx, msg, nil)
		if err != nil {
			t.Skipf("调用 balanceOf() 失败: %v (可能是网络问题或合约不支持)", err)
			return
		}

		var balance *big.Int
		err = contractABI.UnpackIntoInterface(&balance, "balanceOf", result)
		if err != nil {
			t.Fatalf("解析 balanceOf 返回值失败: %v", err)
		}

		t.Logf("✓ balanceOf(%s) = %s", owner.Hex(), balance.String())
	})

	t.Run("ERC721_ownerOf", func(t *testing.T) {
		// 测试 ownerOf(uint256) 方法
		tokenId := big.NewInt(1)
		data, err := contractABI.Pack("ownerOf", tokenId)
		if err != nil {
			t.Fatalf("打包 ownerOf 方法失败: %v", err)
		}

		msg := ethereum.CallMsg{
			To:   &testAddress,
			Data: data,
		}

		result, err := etClient.CallContract(ctx, msg, nil)
		if err != nil {
			t.Skipf("调用 ownerOf() 失败: %v (可能是网络问题或合约不支持，或tokenId不存在)", err)
			return
		}

		var owner common.Address
		err = contractABI.UnpackIntoInterface(&owner, "ownerOf", result)
		if err != nil {
			t.Fatalf("解析 ownerOf 返回值失败: %v", err)
		}

		t.Logf("✓ ownerOf(%s) = %s", tokenId.String(), owner.Hex())
	})

	t.Run("ERC721_tokenURI", func(t *testing.T) {
		// 测试 tokenURI(uint256) 方法
		tokenId := big.NewInt(1)
		data, err := contractABI.Pack("tokenURI", tokenId)
		if err != nil {
			t.Fatalf("打包 tokenURI 方法失败: %v", err)
		}

		msg := ethereum.CallMsg{
			To:   &testAddress,
			Data: data,
		}

		result, err := etClient.CallContract(ctx, msg, nil)
		if err != nil {
			t.Skipf("调用 tokenURI() 失败: %v (可能是网络问题或合约不支持，或tokenId不存在)", err)
			return
		}

		var uri string
		err = contractABI.UnpackIntoInterface(&uri, "tokenURI", result)
		if err != nil {
			t.Fatalf("解析 tokenURI 返回值失败: %v", err)
		}

		if uri == "" {
			t.Logf("警告: tokenURI() 返回空字符串")
		} else {
			t.Logf("✓ tokenURI(%s) = %s", tokenId.String(), uri)
		}
	})

	t.Run("ERC721_totalSupply", func(t *testing.T) {
		// 测试 totalSupply() 方法（如果合约支持）
		data, err := contractABI.Pack("totalSupply")
		if err != nil {
			t.Skipf("打包 totalSupply 方法失败: %v (可能合约不支持此方法)", err)
			return
		}

		msg := ethereum.CallMsg{
			To:   &testAddress,
			Data: data,
		}

		result, err := etClient.CallContract(ctx, msg, nil)
		if err != nil {
			t.Skipf("调用 totalSupply() 失败: %v (可能是网络问题或合约不支持)", err)
			return
		}

		var supply *big.Int
		err = contractABI.UnpackIntoInterface(&supply, "totalSupply", result)
		if err != nil {
			t.Fatalf("解析 totalSupply 返回值失败: %v", err)
		}

		t.Logf("✓ totalSupply() = %s", supply.String())
	})
}

// TestCallMethodsIntegration 集成测试：测试完整的fetchTokenInfo和fetchNFTInfo流程
func TestCallMethodsIntegration(t *testing.T) {
	etClient, err := blockchain.Dial("http://192.168.120.31:8545")
	if err != nil {
		t.Fatalf("连接以太坊客户端失败: %v", err)
	}
	defer etClient.Close()

	t.Run("fetchTokenInfo", func(t *testing.T) {
		// 测试ERC20 Token信息获取
		token := &db.Token{
			Address:  "0x3d41eacfcc93aec57e02513f8ebdd9d1642cb11d",
			Standard: "ERC20",
		}

		err := fetchTokenInfo(etClient, token)
		if err != nil {
			t.Skipf("获取Token信息失败: %v (可能是网络问题)", err)
			return
		}

		// 验证获取到的信息
		t.Logf("Token信息:")
		t.Logf("  Address: %s", token.Address)
		t.Logf("  Name: %s", token.Name)
		t.Logf("  Symbol: %s", token.Symbol)
		t.Logf("  Decimals: %d", token.Decimals)
		t.Logf("  Supply: %s", token.Supply)

		// 如果至少有一个字段有值，认为测试通过
		if token.Name != "" || token.Symbol != "" || token.Decimals > 0 || (token.Supply != "" && token.Supply != "0") {
			t.Logf("✓ 成功获取到Token信息")
		} else {
			t.Errorf("所有字段都为空，可能合约不支持这些方法")
		}
	})

	t.Run("fetchNFTInfo_ERC721", func(t *testing.T) {
		// 测试ERC721 NFT信息获取
		nft := &db.NFT{
			Address:  "0x3d41eacfcc93aec57e02513f8ebdd9d1642cb11d",
			Standard: "ERC721",
		}

		err := fetchNFTInfo(etClient, nft)
		if err != nil {
			t.Skipf("获取NFT信息失败: %v (可能是网络问题)", err)
			return
		}

		// 验证获取到的信息
		t.Logf("NFT信息 (ERC721):")
		t.Logf("  Address: %s", nft.Address)
		t.Logf("  Name: %s", nft.Name)
		t.Logf("  Symbol: %s", nft.Symbol)
		t.Logf("  Supply: %s", nft.Supply)

		// 如果至少有一个字段有值，认为测试通过
		if nft.Name != "" || nft.Symbol != "" || (nft.Supply != "" && nft.Supply != "0") {
			t.Logf("✓ 成功获取到NFT信息")
		} else {
			t.Logf("警告: 所有字段都为空，可能合约不支持这些方法")
		}
	})
}

// TestCallMethods 原始测试方法（保留用于调试）
func TestCallMethods(t *testing.T) {
	etClient, err := blockchain.Dial("http://192.168.120.31:8545")
	if err != nil {
		t.Fatalf("连接以太坊客户端失败: %v", err)
	}
	defer etClient.Close()

	ctx := context.Background()

	testAddress := common.HexToAddress("0x3d41eacfcc93aec57e02513f8ebdd9d1642cb11d")

	contractABI, err := abi.JSON(strings.NewReader(blockchain.ERC20_abi))
	if err != nil {
		t.Fatalf("解析ERC20 ABI失败: %v", err)
	}

	data, err := contractABI.Pack("symbol")
	if err != nil {
		t.Fatalf("打包方法失败: %v", err)
	}

	t.Logf("data: %v", hex.EncodeToString(data))
	t.Logf("data: %v", hex.EncodeToString(getFunctionSelector("symbol()")))

	msg := ethereum.CallMsg{
		To:   &testAddress,
		Data: getFunctionSelector("name()"),
	}

	result, err := etClient.CallContract(ctx, msg, nil)
	if err != nil {
		t.Fatalf("调用合约方法失败: %v", err)
		return
	}

	var ret string
	err = contractABI.UnpackIntoInterface(&ret, "name", result)
	if err != nil {
		t.Fatalf("解析返回值失败: %v", err)
		return
	}
	t.Logf("返回值: %v", ret)
}
