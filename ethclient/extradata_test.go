package bs_eth

import (
	"bytes"
	"fmt"
	"log"
	"sort"
	"testing"

	"github.com/ethereum/go-ethereum/common"
)

func TestExtra(t *testing.T) {
	// 定义验证人地址（可以有多个）
	validatorAddrs := [][]byte{
		//common.HexToAddress("0x30E938B0630C02F394d17925FDB5fB046f70d452"),
		common.Hex2Bytes("013ce455fd10aff3f6ebf82a10270829b93b5e65aeebac604efcaba66e4c8a6091"),
		// 添加更多地址...
	}

	// 生成 extraData
	extraData, err := GenerateExtraData(validatorAddrs)
	if err != nil {
		log.Fatal("生成 extraData 失败:", err)
	}
	fmt.Printf("extraData: %s\n", common.Bytes2Hex(extraData))
}

// GenerateExtraData 生成 Clique 共识的 extraData
// 格式: 32 bytes vanity + N * 20 bytes signers + 65 bytes seal (0)
func GenerateExtraData(signers [][]byte) ([]byte, error) {
	// 1. vanity（32 bytes，这里设为全 0）
	vanity := make([]byte, 32)

	// 2. 验证人地址必须按字典序排序
	sortedSigners := make([][]byte, len(signers))
	copy(sortedSigners, signers)
	sort.Slice(sortedSigners, func(i, j int) bool {
		return bytes.Compare(sortedSigners[i][:], sortedSigners[j][:]) < 0
	})

	// 3. 拼接地址
	signersBytes := make([]byte, 0, len(sortedSigners)*20)
	for _, addr := range sortedSigners {
		signersBytes = append(signersBytes, addr...)
	}

	// 4. seal（65 bytes，创世块为全 0）
	seal := make([]byte, 65)

	// 5. 拼接最终 extraData
	extraData := make([]byte, 0, 32+len(signersBytes)+65)
	extraData = append(extraData, vanity...)
	extraData = append(extraData, signersBytes...)
	extraData = append(extraData, seal...)

	return extraData, nil
}
