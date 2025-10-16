package pb

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"testing"
	"time"
)

func Test_GrpcInTestEnv(t *testing.T) {
	var opts []grpc.DialOption
	opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))
	conn, err := grpc.NewClient("192.168.120.32:9966", opts...)
	if err != nil {
		t.Error(err)
	}
	defer conn.Close()

	var data RequestMessage_InstrumentTicketRequest
	// 创建 Instrument 对象
	instrument := &Instrument{
		InstrumentId:          1,
		ProjectId:             1011,
		ParentInstrumentId:    0,
		CompanyId:             1001,
		AssetType:             "Test",
		InstrumentType:        "Test",
		InstrumentName:        "Test",
		InstrumentImages:      "Test",
		InstrumentVideo:       12345,
		InstrumentDescription: "Test",
		InstrumentAnnex:       "Test",
		ContactName:           "Test",
		ContactPhone:          "1234567890",
		InstrumentStatus:      "Test",
		InstrumentBiddingRule: &InstrumentBiddingRule{
			InstrumentId:        1,
			InstrumentGrade:     1,
			BiddingStartTime:    time.Now().Format(time.RFC3339),
			BiddingEndTime:      time.Now().Add(24 * time.Hour).Format(time.RFC3339),
			BiddingDelayTime:    30,
			BiddingType:         "Open",
			BiddingStartPrice:   "100000.0",
			BiddingDealRule:     1,
			BiddingDeposit:      "1000.0",
			BiddingChanges:      "100.0",
			InstrumentQuantity:  10,
			BiddingDescription:  "Starting auction for the property.",
			MarketPrice:         "88888",
			BiddingDepositRatio: "0.2",
			LowestPrice:         "666",
		},
		InstrumentBiddingSales: &InstrumentBiddingSales{
			InstrumentId:             1,
			InstrumentViews:          100,
			RegistrationQuantity:     5,
			BidQuantity:              10,
			CurrentPrice:             "105000.5555",
			CurrentRemainingQuantity: 5,
			EstimatedStartTime:       time.Now().Format(time.RFC3339),
			EstimatedEndTime:         time.Now().Add(48 * time.Hour).Format(time.RFC3339),
		},
	}
	// 创建 InstrumentTicketRequest 对象
	data.InstrumentTicketRequest = &InstrumentTicketRequest{
		InstrumentId:         1,
		ScenicId:             2001,
		GoodsId:              3001,
		ProductType:          "Ticket",
		TicketType:           "",
		ScenicProvinceCode:   1,
		ScenicCityCode:       2,
		ScenicAreaCode:       3,
		BuyStartDate:         time.Now().Format(time.RFC3339),
		BuyEndDate:           time.Now().Add(7 * 24 * time.Hour).Format(time.RFC3339),
		UseStartDate:         time.Now().Add(1 * 24 * time.Hour).Format(time.RFC3339),
		UseEndDate:           time.Now().Add(8 * 24 * time.Hour).Format(time.RFC3339),
		ProjectChainUniqueId: "111961",
		Instrument:           instrument,
	}

	in := &RequestMessage{
		RequestId:     "789895",
		SerialNumber:  64,
		ReferenceDate: "klznacvxcvzmnc.zxncnvbxc",
		MessageType:   16,
		Data:          &data,
	}
	ctx := context.Background()
	client := NewBlockChainTicketServerClient(conn)
	ticket, err := client.MainOrderTicket(ctx, in)
	if err != nil {
		t.Error(err)
	}

	fmt.Println(ticket)
}

func Test_GrpcInCanaryEnv(t *testing.T) {
	var opts []grpc.DialOption
	opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))
	conn, err := grpc.NewClient("192.168.10.127:9966", opts...)
	if err != nil {
		t.Error(err)
	}
	defer conn.Close()

	var data RequestMessage_InstrumentTicketRequest
	// 创建 Instrument 对象
	instrument := &Instrument{
		InstrumentId:          1,
		ProjectId:             1011,
		ParentInstrumentId:    0,
		CompanyId:             1001,
		AssetType:             "Canary",
		InstrumentType:        "Canary",
		InstrumentName:        "Canary",
		InstrumentImages:      "Canary",
		InstrumentVideo:       12345,
		InstrumentDescription: "Canary",
		InstrumentAnnex:       "Canary",
		ContactName:           "Canary",
		ContactPhone:          "Canary",
		InstrumentStatus:      "Canary",
		InstrumentBiddingRule: &InstrumentBiddingRule{
			InstrumentId:        1,
			InstrumentGrade:     1,
			BiddingStartTime:    time.Now().Format(time.RFC3339),
			BiddingEndTime:      time.Now().Add(24 * time.Hour).Format(time.RFC3339),
			BiddingDelayTime:    30,
			BiddingType:         "Open",
			BiddingStartPrice:   "100000.0",
			BiddingDealRule:     1,
			BiddingDeposit:      "1000.0",
			BiddingChanges:      "100.0",
			InstrumentQuantity:  10,
			BiddingDescription:  "Starting auction for the property.",
			MarketPrice:         "88888",
			BiddingDepositRatio: "0.2",
			LowestPrice:         "666",
		},
		InstrumentBiddingSales: &InstrumentBiddingSales{
			InstrumentId:             1,
			InstrumentViews:          100,
			RegistrationQuantity:     5,
			BidQuantity:              10,
			CurrentPrice:             "105000.5555",
			CurrentRemainingQuantity: 5,
			EstimatedStartTime:       time.Now().Format(time.RFC3339),
			EstimatedEndTime:         time.Now().Add(48 * time.Hour).Format(time.RFC3339),
		},
	}
	// 创建 InstrumentTicketRequest 对象
	data.InstrumentTicketRequest = &InstrumentTicketRequest{
		InstrumentId:         1,
		ScenicId:             2001,
		GoodsId:              3001,
		ProductType:          "Ticket",
		TicketType:           "",
		ScenicProvinceCode:   1,
		ScenicCityCode:       2,
		ScenicAreaCode:       3,
		BuyStartDate:         time.Now().Format(time.RFC3339),
		BuyEndDate:           time.Now().Add(7 * 24 * time.Hour).Format(time.RFC3339),
		UseStartDate:         time.Now().Add(1 * 24 * time.Hour).Format(time.RFC3339),
		UseEndDate:           time.Now().Add(8 * 24 * time.Hour).Format(time.RFC3339),
		ProjectChainUniqueId: "111961",
		Instrument:           instrument,
	}

	in := &RequestMessage{
		RequestId:     "789895",
		SerialNumber:  64,
		ReferenceDate: "klznacvxcvzmnc.zxncnvbxc",
		MessageType:   16,
		Data:          &data,
	}
	ctx := context.Background()
	client := NewBlockChainTicketServerClient(conn)
	ticket, err := client.MainOrderTicket(ctx, in)
	if err != nil {
		t.Error(err)
	}

	fmt.Println(ticket)
}
