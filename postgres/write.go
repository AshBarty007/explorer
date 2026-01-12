package bsdb

import (
	"errors"
	"fmt"

	"gorm.io/gorm"
)

// WriteBlock 写入区块数据，如果已存在则跳过
func WriteBlock(block Block) error {
	// 检查区块是否已存在（通过 hash 或 number）
	var existingBlock Block
	result := Db.Where("hash = ? OR number = ?", block.Hash, block.Number).First(&existingBlock)
	if result.Error == nil {
		// 区块已存在，跳过
		return nil
	}
	if !errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return fmt.Errorf("查询区块失败: %w", result.Error)
	}

	// 创建新区块
	if err := Db.Create(&block).Error; err != nil {
		return fmt.Errorf("写入区块失败: %w", err)
	}
	return nil
}

// WriteTransaction 写入交易数据，如果已存在则跳过
func WriteTransaction(transaction Transaction) error {
	// 检查交易是否已存在（通过 hash）
	var existingTx Transaction
	result := Db.Where("hash = ?", transaction.Hash).First(&existingTx)
	if result.Error == nil {
		// 交易已存在，跳过
		return nil
	}
	if !errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return fmt.Errorf("查询交易失败: %w", result.Error)
	}

	// 创建新交易
	if err := Db.Create(&transaction).Error; err != nil {
		return fmt.Errorf("写入交易失败: %w", err)
	}
	return nil
}

// WriteReceipt 写入收据数据，如果已存在则跳过
func WriteReceipt(receipt Receipt) error {
	// 检查收据是否已存在（通过 transaction_hash）
	var existingReceipt Receipt
	result := Db.Where("transaction_hash = ?", receipt.TransactionHash).First(&existingReceipt)
	if result.Error == nil {
		// 收据已存在，跳过
		return nil
	}
	if !errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return fmt.Errorf("查询收据失败: %w", result.Error)
	}

	// 创建新收据
	if err := Db.Create(&receipt).Error; err != nil {
		return fmt.Errorf("写入收据失败: %w", err)
	}
	return nil
}

// WriteLog 写入日志数据
func WriteLog(log Log) error {
	// 日志可能重复，但通常每条日志都有唯一的组合（transaction_hash + log_index）
	// 检查日志是否已存在
	var existingLog Log
	result := Db.Where("transaction_hash = ? AND log_index = ?", log.TransactionHash, log.LogIndex).First(&existingLog)
	if result.Error == nil {
		// 日志已存在，跳过
		return nil
	}
	if !errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return fmt.Errorf("查询日志失败: %w", result.Error)
	}

	// 创建新日志
	if err := Db.Create(&log).Error; err != nil {
		return fmt.Errorf("写入日志失败: %w", err)
	}
	return nil
}

// WriteAccount写入账户数据，如果已存在则更新
func WriteAccount(account Account) error {
	// Account 表有唯一索引 idx_account_address，使用 FirstOrCreate
	var existingAccount Account
	result := Db.Where("address = ?", account.Address).First(&existingAccount)

	if result.Error == nil {
		// 账户已存在，更新信息
		updateData := Account{
			Type:             account.Type,
			Description:      account.Description,
			Balance:          account.Balance,
			TotalTransaction: account.TotalTransaction,
			Code:             account.Code,
			Token:            account.Token,
			NFT:              account.NFT,
		}

		// 更新首次交易时间（如果为空）
		if existingAccount.FirstTransactionAt == nil && account.FirstTransactionAt != nil {
			updateData.FirstTransactionAt = account.FirstTransactionAt
		}

		// 更新最后交易时间
		if account.LastTransactionAt != nil {
			updateData.LastTransactionAt = account.LastTransactionAt
		}

		if err := Db.Model(&existingAccount).Updates(updateData).Error; err != nil {
			return fmt.Errorf("更新账户失败: %w", err)
		}
		return nil
	}

	if !errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return fmt.Errorf("查询账户失败: %w", result.Error)
	}

	// 创建新账户
	if err := Db.Create(&account).Error; err != nil {
		return fmt.Errorf("写入账户失败: %w", err)
	}
	return nil
}

// WriteToken 写入Token数据，如果已存在则更新
func WriteToken(token Token) error {
	// 检查Token是否已存在（通过 address）
	var existingToken Token
	result := Db.Where("address = ?", token.Address).First(&existingToken)

	if result.Error == nil {
		// Token已存在，更新信息
		updateData := Token{
			Name:               token.Name,
			Symbol:             token.Symbol,
			Decimals:           token.Decimals,
			Supply:             token.Supply,
			Standard:           token.Standard,
			Description:        token.Description,
			Creator:            token.Creator,
			CreatedTime:        token.CreatedTime,
			CreatedHash:        token.CreatedHash,
			CreatedBlockNumber: token.CreatedBlockNumber,
		}

		if err := Db.Model(&existingToken).Updates(updateData).Error; err != nil {
			return fmt.Errorf("更新Token失败: %w", err)
		}
		return nil
	}

	if !errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return fmt.Errorf("查询Token失败: %w", result.Error)
	}

	// 创建新Token
	if err := Db.Create(&token).Error; err != nil {
		return fmt.Errorf("写入Token失败: %w", err)
	}
	return nil
}

// WriteNFT 写入NFT数据，如果已存在则更新
func WriteNFT(nft NFT) error {
	// 检查NFT是否已存在（通过 address）
	var existingNFT NFT
	result := Db.Where("address = ?", nft.Address).First(&existingNFT)

	if result.Error == nil {
		// NFT已存在，更新信息
		updateData := NFT{
			Name:               nft.Name,
			Symbol:             nft.Symbol,
			Supply:             nft.Supply,
			Standard:           nft.Standard,
			Description:        nft.Description,
			TokenUri:           nft.TokenUri,
			Metadata:           nft.Metadata,
			Creator:            nft.Creator,
			CreatedTime:        nft.CreatedTime,
			CreatedHash:        nft.CreatedHash,
			CreatedBlockNumber: nft.CreatedBlockNumber,
		}

		if err := Db.Model(&existingNFT).Updates(updateData).Error; err != nil {
			return fmt.Errorf("更新NFT失败: %w", err)
		}
		return nil
	}

	if !errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return fmt.Errorf("查询NFT失败: %w", result.Error)
	}

	// 创建新NFT
	if err := Db.Create(&nft).Error; err != nil {
		return fmt.Errorf("写入NFT失败: %w", err)
	}
	return nil
}

// WriteContract 写入合约数据，如果已存在则更新
func WriteContract(contract Contract) error {
	// 检查合约是否已存在（通过 address）
	var existingContract Contract
	result := Db.Where("address = ?", contract.Address).First(&existingContract)

	if result.Error == nil {
		// 合约已存在，更新信息
		updateData := Contract{
			Description: contract.Description,
		}

		if err := Db.Model(&existingContract).Updates(updateData).Error; err != nil {
			return fmt.Errorf("更新合约失败: %w", err)
		}
		return nil
	}

	if !errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return fmt.Errorf("查询合约失败: %w", result.Error)
	}

	// 创建新合约
	if err := Db.Create(&contract).Error; err != nil {
		return fmt.Errorf("写入合约失败: %w", err)
	}
	return nil
}