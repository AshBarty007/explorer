package ipfs

import "testing"

func TestUploadToIPFS(t *testing.T) {
	str := "{InstrumentID:1 ScenicID:2001 GoodsID:3001 ProductType:Ticket TicketType: ScenicProvinceCode:1 ScenicCityCode:2 ScenicAreaCode:3 BuyStartDate:2025-07-10T15:46:52+08:00 BuyEndDate:2025-07-17T15:48:00 UseStartDate:2025-07-11T15:46:52+08:00 UseEndDate:2025-07-18T15:46:52+08:00 ProjectChainUniqueId:111111 Instrument:0xc0001a8960}"
	ipfsResp, err := UploadToIPFS(str)
	if err != nil {
		return
	}
	t.Log("resp: ", ipfsResp)
}
