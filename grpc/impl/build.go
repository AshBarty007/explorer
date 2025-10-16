package impl

import (
	"blockchain_services/grpc/pb"
	"fmt"
	"github.com/sirupsen/logrus"
	"strconv"
	"time"
)

const (
	TicketURI = `https://minio-test.shukeyun.com/blc-nft-pic/
	%E5%8D%83%E5%BA%93%E7%BD%91_%E8%A5%
	BF%E8%97%8F%E7%BA%B3%E6%9C%A8%E9%94%99%E9%9B%AA%E5%B1%B1%E5%9C
	%A3%E6%B9%96_%E6%91%84%E5%BD%B1%E5%9B%BE%E7%BC%96%E5%8F%B720585679%5B1%5D.png?
	X-Amz-Algorithm=AWS4-HMAC-SHA256&
	X-Amz-Credential=jFTZtwMhhLMK8gekGTr5%2F20240709%2F%2Fs3%2Faws4_request&
	X-Amz-Date=20240709T092559Z&X-Amz-Expires=432000&X-Amz-SignedHeaders=host
	&X-Amz-Signature=599ac29e4428780de83562b369e064d2c13a4a7d3343aa846bc5947dc1d8a3f6`

	ChainName = "智旅链"

	layout = "2006-01-02 15:04:05.999"
)

func BuildTickets(req *pb.CreateStockRequest) *TicketInfo {

	if req == nil || !CheckTicket(req) {
		logrus.Errorf("CreateStockRequest has nil field")
		return nil
	}

	// ticketsData := make([]*TicketData, 0, len(req.TicketIssuance.TicketData))

	// for _, v := range req.TicketIssuance.TicketData {

	// 	buyerInfo := make([]*BuyerInfo, len(req.TicketIssuance.TicketData))

	// 	for _, info := range v.BuyerInfo {
	// 		buyerInfo = append(buyerInfo, &BuyerInfo{
	// 			BuyerInfoIDName: info.IdName,
	// 			IDNumber:        info.IdNumber,
	// 		})
	// 	}

	// 	ticketData := &TicketData{
	// 		BuyerInfo:      buyerInfo,
	// 		Phone:          req.TicketIssuance.Phone,
	// 		SaleChannel:    v.SaleChannel.GetValue(),
	// 		OrderID:        v.OrderId,
	// 		OrderGroupID:   req.TicketIssuance.OrderGroupId,
	// 		PlayerNum:      v.PlayerNum.GetValue(),
	// 		IssuanceType:   v.IssuanceType.GetValue(),
	// 		Status:         v.Status.GetValue(),
	// 		PrintEncode:    v.PrintEncode,
	// 		EnterBeginTime: v.EnterBeginTime,
	// 		EnterEndTime:   v.EnterEndTime,
	// 		OverdueTime:    v.OverdueTime,
	// 		ProviderID:     strconv.FormatInt(v.ProviderId, 10),
	// 		StoreID:        strconv.FormatInt(v.StoreId, 10),
	// 		SellingPrice:   v.SellingPrice,
	// 		CancelCount:    v.CancelCount.GetValue(),
	// 		EnterTime:      v.EnterTime,
	// 		CheckedNum:     v.CheckedNum.GetValue(),
	// 		UsedCount:      v.UsedCount.GetValue(),
	// 		UsedDays:       v.UsedDays.GetValue(),
	// 		Account:        v.Account,
	// 		Org:            v.Org,
	// 	}

	// 	ticketsData = append(ticketsData, ticketData)

	// }

	return &TicketInfo{
		BasicInformation: BasicInformation{
			SimpleTicket: SimpleTicket{
				ScenicID:      strconv.FormatInt(req.SimpleTicket.ScenicId, 10),
				ScenicName:    req.SimpleTicket.TicketStock.ScenicName,
				SimpleName:    req.SimpleTicket.Name, //NOTE: 为了防止重名
				MarketPrice:   req.SimpleTicket.MarketPrice,
				ProType:       req.SimpleTicket.ProType.GetValue(),
				UseType:       req.SimpleTicket.UseType.GetValue(),
				TimeRestrict:  req.SimpleTicket.TimeRestrict.GetValue(),
				RestrictType:  req.SimpleTicket.RestrictType.GetValue(),
				RestrictWeek:  req.SimpleTicket.RestrictWeek,
				ValidityDay:   req.SimpleTicket.ValidityDay.GetValue(),
				IsActivate:    req.SimpleTicket.IsActivate.GetValue(),
				UseCount:      req.SimpleTicket.UseCount.GetValue(),
				AvailableDays: req.SimpleTicket.AvailableDays.GetValue(),
				ParkStatistic: req.SimpleTicket.ParkStatistic.GetValue(),
				OperatorID:    strconv.FormatInt(req.SimpleTicket.OperatorId, 10),
				TimeSharing: []TimeSharing{
					{
						TimeSharingID:        strconv.FormatInt(req.SimpleTicket.TicketGoods.TimeShareId, 10),
						TimeSharingBeginTime: req.SimpleTicket.TicketGoods.TimeSharing.BeginTime,
						TimeSharingEndTime:   req.SimpleTicket.TicketGoods.TimeSharing.EndTime,
					},
				},
				TicketGoods: []TicketGoods{
					{
						TicketGoodsID:   strconv.FormatInt(req.SimpleTicket.TicketGoods.Id, 10),
						GoodsName:       req.SimpleTicket.TicketGoods.GoodsName,
						TimeShareID:     strconv.FormatInt(req.SimpleTicket.TicketGoods.TimeShareId, 10),
						OverallDiscount: req.SimpleTicket.TicketGoods.OverallDiscount,
						BeginDiscount:   req.SimpleTicket.TicketGoods.BeginDiscount,
						EndDiscount:     req.SimpleTicket.TicketGoods.EndDiscount,
						TicketGoodsType: req.SimpleTicket.TicketGoods.Type.GetValue(), //NOTE: 为了防止重名
						PeopleNumber:    req.SimpleTicket.TicketGoods.PeopleNumber.GetValue(),
						MinPeople:       req.SimpleTicket.TicketGoods.MinPeople.GetValue(),
						MaxPeople:       req.SimpleTicket.TicketGoods.MaxPeople.GetValue(),
						RuleIssue: RuleIssue{
							RuleIssueName:      req.SimpleTicket.TicketGoods.RuleIssue.Name,            //NOTE: 为了防止重名
							RuleIssueWay:       req.SimpleTicket.TicketGoods.RuleIssue.Way.GetValue(),  //NOTE: 为了防止重名
							RuleIssueType:      req.SimpleTicket.TicketGoods.RuleIssue.Type.GetValue(), //NOTE: 为了防止重名
							IsRealName:         req.SimpleTicket.TicketGoods.RuleIssue.IsRealName.GetValue(),
							UseTime:            req.SimpleTicket.TicketGoods.RuleIssue.UseTime,
							RuleIssueBeginTime: req.SimpleTicket.TicketGoods.RuleIssue.BeginTime,
							RuleIssueEndTime:   req.SimpleTicket.TicketGoods.RuleIssue.EndTime,
							RealNameCheck:      req.SimpleTicket.TicketGoods.RuleIssue.RealNameCheck.GetValue(),
							OnlyOwnerBuy:       req.SimpleTicket.TicketGoods.RuleIssue.OnlyOwnerBuy.GetValue(),
							RightsCheck:        req.SimpleTicket.TicketGoods.RuleIssue.RightsCheck.GetValue(),
							RightsID:           strconv.FormatInt(req.SimpleTicket.TicketGoods.RuleIssue.RightsId, 10),
							NeedApproval:       req.SimpleTicket.TicketGoods.RuleIssue.NeedApproval.GetValue(),
							ApproveID:          strconv.FormatInt(req.SimpleTicket.TicketGoods.RuleIssue.ApproveId, 10),
							ApproveContent:     req.SimpleTicket.TicketGoods.RuleIssue.ApproveContent,
							RuleType:           req.SimpleTicket.TicketGoods.RuleIssue.RuleType.GetValue(),
							OnlyWindowSale:     req.SimpleTicket.TicketGoods.RuleIssue.OnlyOwnerBuy.GetValue(),
						},
						RuleCheck: RuleCheck{
							IdentityType:  req.SimpleTicket.TicketGoods.RuleCheck.IdentityType,
							RuleCheckName: req.SimpleTicket.TicketGoods.RuleCheck.Name, //NOTE: 为了防止重名
							ControlType:   req.SimpleTicket.TicketGoods.RuleCheck.ControlType.GetValue(),
							AdoptType:     req.SimpleTicket.TicketGoods.RuleCheck.AdoptType.GetValue(),
							IntervalTime:  req.SimpleTicket.TicketGoods.RuleCheck.IntervalTime.GetValue(),
							TimeShareBook: req.SimpleTicket.TicketGoods.RuleCheck.TimeShareBook.GetValue(),
							CheckPointIds: req.SimpleTicket.TicketGoods.CheckPointIdList,
						},
						RuleRetreat: RuleRetreat{
							RuleRetreatName: req.SimpleTicket.TicketGoods.RuleRetreat.Name, //NOTE: 为了防止重名
							IsRetreat:       req.SimpleTicket.TicketGoods.RuleRetreat.IsRetreat.GetValue(),
							DefaultRate:     req.SimpleTicket.TicketGoods.RuleRetreat.DefaultRate,
						},
					},
				},
				TicketStock: TicketStock{
					StockScenicID:   req.SimpleTicket.TicketStock.ScenicId,   //NOTE: 为了防止重名
					StockScenicName: req.SimpleTicket.TicketStock.ScenicName, //NOTE: 为了防止重名
					StockTicketID:   req.SimpleTicket.TicketStock.TicketId,
					TicketName:      req.SimpleTicket.TicketStock.TicketName,
					AccountID:       req.SimpleTicket.TicketStock.AccountId,
					TicketType:      req.SimpleTicket.TicketStock.TicketType.GetValue(),
					Nums:            req.SimpleTicket.TicketStock.Nums.GetValue(),
					TotalStock:      req.SimpleTicket.TicketStock.TotalStock.GetValue(),
					StockOperatorID: req.SimpleTicket.TicketStock.OperatorId, //NOTE: 为了防止重名
					BatchID:         strconv.FormatInt(req.BatchId, 10),
				},
			},
			IsExchange: req.IsExchange.GetValue(),
		},
		AdditionalInformation: AdditionalInformation{
			TicketData: TicketDataCreate{
				BuyerInfo: []*BuyerInfo{},
			},
			PriceInfo: []*PriceInfo{
				// DistributorID: req.SimpleTicket.TicketGoods.SetSelfPrice.DistributorId,
				// GoodsID:       req.SimpleTicket.TicketGoods.SetSelfPrice.GoodsId,
				// PriceDetailedInfo: PriceDetailedInfo{
				// 	SalePrice:      req.SimpleTicket.TicketGoods.SetSelfPrice.Price.SalePrice,
				// 	ComposePrice:   req.SimpleTicket.TicketGoods.SetSelfPrice.Price.ComposePrice,
				// 	CommissionRate: req.SimpleTicket.TicketGoods.SetSelfPrice.Price.CommissionRate,
				// 	IsCompose:      req.SimpleTicket.TicketGoods.SetSelfPrice.Price.IsCompose,
				// },
			},
			TicketCheck: []*VerifyTicket{},
			StockInfo: StockInfo{
				StockEnterBeginTime: req.SimpleTicket.TicketStock.EnterBeginTime, //NOTE: 为了防止重名
				StockEnterEndTime:   req.SimpleTicket.TicketStock.EnterEndTime,   //NOTE: 为了防止重名
				PurchaseBeginTime:   req.SimpleTicket.TicketStock.PurchaseBeginTime,
				PurchaseEndTime:     req.SimpleTicket.TicketStock.PurchaseEndTime,
			},
		},
	}
}

func BuildMetaData() *Metadata {
	return &Metadata{
		TokenURL: TicketURI,
	}
}

func BuildPriceInfo(req *pb.SetSelfPriceInfoRequest) *PriceInfo {

	if req == nil || !CheckPriceInfo(req) {
		logrus.Errorf("SetSelfPriceInfoRequest has nil field")
		return nil
	}
	return &PriceInfo{
		DistributorID: req.PriceInfoDetailsList.DistributorId,
		GoodsID:       req.PriceInfoDetailsList.GoodsId,
		PriceDetailedInfo: PriceDetailedInfo{
			SalePrice:      req.PriceInfoDetailsList.Price.SalePrice,
			ComposePrice:   req.PriceInfoDetailsList.Price.ComposePrice,
			CommissionRate: req.PriceInfoDetailsList.Price.CommissionRate,
			IsCompose:      req.PriceInfoDetailsList.Price.IsCompose,
			PriceID:        req.PriceInfoDetailsList.PriceId,
			GroupInfo: GroupInfo{
				AddGroupID: req.PriceInfoDetailsList.AddGroupId,
				DelGroupID: req.PriceInfoDetailsList.DelGroupId,
				GroupID:    req.PriceInfoDetailsList.GroupId,
			},
		},
	}

}

func BuildGenerateTicketIssuance(req *pb.GenerateTicketIssuanceRequest) *GenerateTicketNumberInfo {
	if req == nil {
		logrus.Errorf("GenerateTicketIssuanceRequest has nil field")
		return nil
	}

	ownerList := make([]*StockConsumeData, 0, len(req.StockBatchNumberData))

	list := make([]*GenerateTicketInfo, 0, len(req.GenerateTicketNumberInfoData))

	for _, v := range req.StockBatchNumberData {
		senderStockID := fmt.Sprintf("%v-%v-%v", v.StockBatchNumber, v.Account, v.Org)
		ownerList = append(ownerList, &StockConsumeData{
			StockBatchNumber: senderStockID,
			Sender:           v.Account,
			Amount:           v.Number,
		})

	}

	for _, v := range req.GenerateTicketNumberInfoData {

		ids := make([]int64, 0, len(v.GetCheckPointIds()))

		for _, id := range v.GetCheckPointIds() {
			ids = append(ids, id.GetValue())
		}
		data := &GenerateTicketInfo{
			EnterTime:     v.EnterTime,
			PlayerNum:     v.PlayerNum.GetValue(),
			Certificate:   v.Certificate,
			Rand:          v.Rand,
			ScenicID:      v.ScenicId,
			ProType:       v.ProType.GetValue(),
			TimeShareID:   v.TimeShareId.GetValue(),
			TimeShareBook: v.TimeShareBook.GetValue(),
			BeginTime:     v.BeginTime,
			EndTime:       v.EndTime,
			CheckPointIds: ids,
			UUID:          v.Uuid,
		}
		list = append(list, data)
	}
	return &GenerateTicketNumberInfo{
		GenerateTicketNumberInfo: list,
		StockConsumeData:         ownerList,
	}
}

func BuildTicketIssuance(req *pb.BlockTicketIssuanceRequest) *BlockTicketIssuanceData {

	if req == nil {
		logrus.Errorf("GenerateTicketIssuanceRequest has nil field")
		return nil
	}

	data := make([]*TicketData, 0, len(req.TicketData))

	for _, v := range req.TicketData {

		buyers := make([]*BuyerInfo, 0, len(v.BuyerInfo))

		for _, buyInfo := range v.BuyerInfo {
			buyer := &BuyerInfo{
				IDNumber:        buyInfo.IdNumber,
				BuyerInfoIDName: buyInfo.IdName,
			}
			buyers = append(buyers, buyer)
		}

		if v.TicketIssuanceSubInfo == nil {
			logrus.Errorf("TicketIssuanceSubInfo has nil field")
			return nil
		}
		ticket := &TicketData{
			TicketID:         v.Id,
			Phone:            req.Phone,
			SaleChannel:      v.SaleChannel.GetValue(),
			BuyerInfo:        buyers,
			OrderID:          v.OrderId,
			OrderGroupID:     req.OrderGroupId,
			PlayerNum:        v.PlayerNum.GetValue(),
			IssuanceType:     v.IssuanceType.GetValue(),
			Status:           v.Status.GetValue(),
			PrintEncode:      v.PrintEncode,
			EnterBeginTime:   v.EnterBeginTime,
			EnterEndTime:     v.EnterEndTime,
			OverdueTime:      v.OverdueTime,
			ProviderID:       strconv.FormatInt(v.ProviderId, 10),
			UserID:           strconv.FormatInt(v.BuyerId, 10),
			StoreID:          strconv.FormatInt(v.StoreId, 10),
			SellingPrice:     v.SellingPrice,
			CancelCount:      v.CancelCount.GetValue(),
			EnterTime:        v.EnterTime,
			CheckedNum:       v.CheckedNum.GetValue(),
			UsedCount:        v.UsedCount.Value,
			UsedDays:         v.UsedDays.GetValue(),
			StockBatchNumber: v.StockBatchNumber,
			Account:          v.Account,
			Org:              v.Org,
		}

		data = append(data, ticket)

	}

	return &BlockTicketIssuanceData{
		TicketData: data,
	}
}

func BuildOrderInfo(req *pb.OrderInfoRequest) *OrderInfo {

	if req == nil {
		logrus.Errorf("OrderInfoRequest has nil field")
		return nil
	}

	data := make([]*OrderTab, 0, len(req.OrderTabData))

	for _, v := range req.OrderTabData {

		productInfo := make([]*OrderProductTicketData, 0, len(v.GetOrderProductTicketData()))

		for _, info := range v.OrderProductTicketData {

			tickets := make([]*OrderProductTicketRnData, 0, len(info.OrderProductTicketRnData))

			for _, ticket := range info.OrderProductTicketRnData {
				t := &OrderProductTicketRnData{
					ID:                      ticket.Id,
					OrderProductID:          ticket.OrderProductId,
					TicketNumber:            ticket.TicketNumber,
					TicketStatus:            ticket.TicketStatus,
					CommissionSettledStatus: ticket.CommissionSettledStatus,
					IsChain:                 ticket.IsChain,
					BillStatus:              ticket.BillStatus,
					IssueTicketType:         ticket.IssueTicketType,
				}
				tickets = append(tickets, t)
			}
			ticketData := &OrderProductTicketData{
				ScenicID:                 info.ScenicId,
				ScenicName:               info.ScenicName,
				TicketType:               info.TicketType,
				Day:                      info.Day,
				TimeShareID:              info.TimeShareId,
				TimeShare:                info.TimeShare,
				ParentProductID:          info.ParentProductId,
				CommissionType:           info.CommissionType,
				CommissionRate:           info.CommissionRate,
				CommissionAmount:         info.CommissionAmount,
				ActualComAmount:          info.ActualComAmount,
				BdsAccount:               info.BdsAccount,
				BdsOrg:                   info.BdsOrg,
				TicketTypeID:             info.TicketTypeId,
				TicketTypeSubID:          info.TicketTypeSubId,
				RealQuantity:             info.CommissionType,
				OrderProductTicketRnData: tickets,
			}
			productInfo = append(productInfo, ticketData)
		}

		orderInfo := &OrderTab{
			OrderID:                v.OrderId,
			OrderType:              v.OrderType,
			SellerID:               v.SellerId,
			SellerName:             v.SellerName,
			TotalAmount:            v.TotalAmount,
			PayType:                v.PayType,
			SourceType:             v.SourceType,
			OrderStatus:            v.OrderStatus,
			TradeNo:                v.TradeNo,
			MerchantID:             v.MerchantId,
			StoreID:                v.StoreId,
			AgentID:                v.StoreId,
			AgentName:              v.AgentName,
			CommissionSettledType:  v.CommissionSettledType,
			UserID:                 v.UserId,
			Username:               v.Username,
			PayTime:                v.PayTime,
			ModifyTime:             v.ModifyTime,
			PayPeople:              v.PayPeople,
			Nickname:               v.Nickname,
			MerchantNo:             v.MerchantNo,
			OrderProductTicketData: productInfo,
		}

		data = append(data, orderInfo)

	}

	return &OrderInfo{
		OrderGroupID:     req.OrderGroupId,
		OrderStatus:      req.OrderStatus,
		OrderType:        req.OrderType,
		TotalAmount:      req.TotalAmount,
		PayAmount:        req.PayAmount,
		PayType:          req.PayType,
		SourceType:       req.SourceType,
		StockCertificate: req.StockCertificate,
		TradeNo:          req.TradeNo,
		UserID:           req.UserId,
		Username:         req.Username,
		PayTime:          req.PayTime,
		UserPhone:        req.UserPhone,
		OrderTab:         data,
	}
}

func BuildCheckData(req *pb.BlockChainTicketCheckRequest) []*VerifyTicket {

	if req == nil {
		logrus.Errorf("GenerateTicketIssuanceRequest has nil field")
		return nil
	}
	verifyData := make([]*VerifyTicket, 0, len(req.List))
	for _, v := range req.List {

		if v.TicketStatus == nil {
			logrus.Errorf("GenerateTicketIssuanceRequest TicketStatus has nil field")
			return nil
		}

		verifyData = append(verifyData, &VerifyTicket{
			VerifyInfo: VerifyInfo{
				Account:          req.Account,
				Org:              req.Org,
				CheckType:        req.CheckType,
				TicketNumber:     v.TicketNumber,
				StockBatchNumber: v.StockBatchNumber,
				EnterTime:        v.EnterTime,
				CheckNumber:      v.CheckNumber,
				ScenicID:         v.ScenicId,
				IDName:           v.IdName,
				IDCard:           v.IdCard,
				QrCode:           v.QrCode,
				PointName:        v.PointName,
				PointID:          v.PointId,
				EquipmentName:    v.EquipmentName,
				EquipmentID:      v.EquipmentId,
				EquipmentType:    v.EquipmentType,
				UserID:           v.UserId,
				Username:         v.Username,
			},
			TicketStatus: TicketStatus{
				Status:     v.TicketStatus.TicketStatus,
				TicketID:   v.TicketStatus.TicketNumber,
				CheckedNum: v.TicketStatus.CheckedNum,
				UsedCount:  v.TicketStatus.UsedCount,
				UsedDays:   v.TicketStatus.UsedDays,
			},
		})
	}
	return verifyData
}

func BuildTicketStatus(req *pb.TicketStatusPushRequest) []*TimerTicketStatus {
	if req == nil {
		logrus.Errorf("TicketStatusPushRequest has nil field")
		return nil
	}
	list := make([]*TimerTicketStatus, 0, len(req.Data))

	for _, v := range req.Data {
		one := &TimerTicketStatus{
			Status:   v.TicketStatus,
			TicketID: v.TicketNumber,
		}

		list = append(list, one)
	}
	return list
}

func BuildDistributionOrderInfo(req *pb.DistributionOrderInfoRequest) *DistributionInfo {
	if req == nil {
		logrus.Errorf("DistributionOrderInfoRequest has nil field")
		return nil
	}
	tansferDatas := make([]*TansferDataTob, 0)
	orderTabData := make([]*OrderTabToBData, 0, len(req.OrderTabToBData))
	for _, v := range req.OrderTabToBData {
		data := &OrderTabToBData{
			OrderID:               v.OrderId,
			OrderType:             v.OrderType,
			SellerID:              v.SellerId,
			SellerName:            v.SellerName,
			TotalAmount:           v.TotalAmount,
			PayAmount:             v.PayAmount,
			PayType:               v.PayType,
			PayTime:               v.PayTime,
			SourceType:            v.SourceType,
			OrderStatus:           v.OrderStatus,
			TradeNo:               v.TradeNo,
			MerchantID:            v.MerchantId,
			StoreID:               v.StoreId,
			AgentID:               v.AgentId,
			AgentName:             v.AgentName,
			CommissionSettledType: v.CommissionSettledType,
			UserID:                v.UserId,
			Username:              v.Username,
			Nickname:              v.Nickname,
			MerchantNo:            v.MerchantNo,
		}
		orderTabData = append(orderTabData, data)
	}
	orderProductData := make([]*OrderTabDistributeData, 0, len(req.OrderTabDistributeData))
	for _, v := range req.OrderTabDistributeData {

		orderTabData := make([]*OrderProductDistributeData, 0, len(v.OrderProductDistributeData))

		for _, item := range v.OrderProductDistributeData {
			data := &OrderProductDistributeData{
				ScenicID:                 item.ScenicId,
				ScenicName:               item.ScenicName,
				DistributorTicketStockID: item.DistributorTicketStockId,
				BatchID:                  strconv.FormatInt(item.BatchId, 10),
				TicketType:               item.TicketType,
				DayBegin:                 item.DayBegin,
				DayEnd:                   item.DayEnd,
				TimeShare:                item.TimeShare,
				UsableNum:                item.UsableNum,
				OrderProductID:           item.OrderProductId,
				OrderID:                  item.OrderId,
				ProductID:                item.ProductId,
				ProductName:              item.ProductName,
				ProductType:              item.ProductType,
				ProductPrice:             item.ProductPrice,
				Num:                      item.Num,
				AvailableRatio:           item.AvailableRatio,
				AvailableTotalNum:        item.AvailableTotalNum,
				ExchangeFreezeNum:        item.ExchangeFreezeNum,
			}
			orderTabData = append(orderTabData, data)
			senderStockID := fmt.Sprintf("%v-%v-%v", item.StockBatchNumber, v.SellerAccount, v.SellerOrg)
			receiveStockID := fmt.Sprintf("%v-%v-%v", item.StockBatchNumber, v.BuyerAccount, v.BuyerOrg)
			tansferData := &TansferDataTob{
				SenderStockID:     senderStockID,
				ReceiveStockID:    receiveStockID,
				Sender:            v.SellerAccount,
				Receive:           v.BuyerAccount,
				Amount:            strconv.FormatInt(int64(item.Num), 10),
				AvailableRatio:    item.AvailableRatio,
				AvailableTotalNum: item.AvailableTotalNum,
				StockBatchNumber:  item.StockBatchNumber,
			}
			tansferDatas = append(tansferDatas, tansferData)
		}
		data := &OrderTabDistributeData{
			OrderID:                    v.OrderId,
			BuyerID:                    v.BuyerId,
			BuyerName:                  v.BuyerName,
			SellerID:                   v.SellerId,
			SellerName:                 v.SellerName,
			ServiceProviderID:          v.ServiceProviderId,
			ServiceProviderName:        v.ServiceProviderName,
			OrderProductDistributeData: orderTabData,
		}
		orderProductData = append(orderProductData, data)

	}
	distributionOrderInfo := &DistributionOrderInfo{
		OrderGroupID:           req.OrderGroupId,
		OrderStatus:            req.OrderStatus,
		OrderType:              req.OrderType,
		TotalAmount:            req.TotalAmount,
		PayType:                req.PayType,
		SourceType:             req.SourceType,
		StockCertificate:       req.StockCertificate,
		TradeNo:                req.TradeNo,
		UserID:                 req.UserId,
		Username:               req.Username,
		PayTime:                req.PayTime,
		CertID:                 req.CertId,
		UserPhone:              req.UserPhone,
		OrderTabToBData:        orderTabData,
		OrderTabDistributeData: orderProductData,
	}

	return &DistributionInfo{
		DistributionOrderInfo: distributionOrderInfo,
		TansferDatas:          tansferDatas,
	}
}

func BuildDistributionRefundInfo(req *pb.DistributeRefundCreateRequest) *DistributeReturnInfo {
	if req == nil {
		logrus.Errorf("DistributeRefundCreateRequest has nil field")
		return nil
	}
	tansferDatas := make([]*TansferData, 0)

	tansferMap := make(map[string]*AccountData, 0)

	orderTabData := make([]*OrderRefundGroup, 0, len(req.OrderRefundGroup))
	orderRefundData := make([]*OrderRefund, 0, len(req.OrderRefund))
	orderRefundProductDistribute := make([]*OrderRefundProductDistribute, 0, len(req.OrderRefundProductDistribute))
	for _, v := range req.OrderRefundGroup {
		tobData := &OrderRefundGroup{
			OrderRefundGroupID: v.OrderGroupId,
			OrderGroupID:       v.OrderGroupId,
			OrderRefundID:      v.OrderRefundId,
			CreateTime:         v.CreateTime,
		}
		orderTabData = append(orderTabData, tobData)
	}

	for _, v := range req.OrderRefund {
		orderData := &OrderRefund{
			RefundID:                v.RefundId,
			OrderID:                 v.OrderId,
			RefundAmount:            v.RefundAmount,
			RefundFee:               v.RefundFee,
			RefundStatus:            v.RefundStatus,
			RefundType:              v.RefundType,
			TradeNo:                 v.TradeNo,
			RefundTime:              v.RefundTime,
			CreateTime:              v.CreateTime,
			Remark:                  v.Remark,
			FailMessage:             v.FailMessage,
			UserID:                  v.UserId,
			Username:                v.Username,
			CommissionSettledStatus: v.CommissionSettledStatus,
			StockCertificate:        v.StockCertificate,
			ProductSkuName:          v.ProductSkuName,
		}
		orderRefundData = append(orderRefundData, orderData)

		tansferMap[v.RefundId] = &AccountData{
			From:         v.BuyerAccount,
			FromOrgMspID: v.BuyerOrg,
			To:           v.SellerAccount,
			ToOrgMspID:   v.SellerOrg,
		}
	}

	for _, v := range req.OrderRefundProductDistribute {
		if _, ok := tansferMap[v.RefundId]; !ok {
			logrus.Errorf("tansferMap not find key with RefundId := %v", v.RefundId)
			return nil
		}
		productDistribute := &OrderRefundProductDistribute{
			RefundID:                v.RefundId,
			OrderProductID:          v.OrderProductId,
			Num:                     v.Num,
			ProductID:               v.OrderProductId,
			ProductName:             v.ProductName,
			ProductSkuID:            v.ProductSkuId,
			ProductType:             v.ProductType,
			ProductPrice:            v.ProductPrice,
			DayBegin:                v.DayBegin,
			DayEnd:                  v.DayEnd,
			TimeShareID:             v.TimeShareId,
			TimeShare:               v.TimeShare,
			ScenicID:                v.ScenicId,
			ScenicName:              v.ScenicName,
			BatchID:                 v.BatchId,
			DistributeTicketStockID: v.DistributeTicketStockId,
		}
		orderRefundProductDistribute = append(orderRefundProductDistribute, productDistribute)

		senderTokenID := fmt.Sprintf("%v-%v-%v", v.BatchId, tansferMap[v.RefundId].From, tansferMap[v.RefundId].FromOrgMspID)
		receiveTokenID := fmt.Sprintf("%v-%v-%v", v.BatchId, tansferMap[v.RefundId].To, tansferMap[v.RefundId].ToOrgMspID)
		tansferData := &TansferData{
			SenderStockID:  senderTokenID,
			Sender:         tansferMap[v.RefundId].From,
			Receive:        tansferMap[v.RefundId].To,
			ReceiveStockID: receiveTokenID,
			Amount:         strconv.FormatInt(int64(v.Num), 10),
		}
		tansferDatas = append(tansferDatas, tansferData)
	}

	return &DistributeReturnInfo{
		DistributeRefundInfo: &DistributeRefundInfo{
			OrderRefundGroup:             orderTabData,
			OrderRefund:                  orderRefundData,
			OrderRefundProductDistribute: orderRefundProductDistribute,
		},
		TansferDatas: tansferDatas,
	}
}

func BuildRefundTocInfo(req *pb.BlockOrderRefundRequest) *OrderRefundInfoToC {
	if !CheckRefundTocInfo(req) {
		logrus.Errorf("BlockOrderRefundRequest has nil field")
		return nil
	}
	refundProductTicketsToC := make([]*RefundProductTicketToC, 0, len(req.GetBlockOrderRefundProductTicket()))
	for _, v := range req.GetBlockOrderRefundProductTicket() {
		list := make([]*StockConsumeData, 0)

		for _, stockInfo := range v.DistributorTicketStockTransferInfo {
			senderTokenID := fmt.Sprintf("%v-%v-%v", stockInfo.StockBatchNumber, stockInfo.Account, stockInfo.Org)
			data := &StockConsumeData{
				StockBatchNumber: senderTokenID,
				Sender:           stockInfo.Account,
				Amount:           stockInfo.Number,
			}
			list = append(list, data)
		}
		refundProductTicketToC := &RefundProductTicketToC{
			RefundID:       v.RefundId,
			OrderProductID: v.OrderProductId,
			TicketNumber:   v.TicketNumber,
			ProductID:      v.ProductId,
			ProductName:    v.ProductName,
			ProductSkuID:   v.ProductSkuId,
			ProductType:    v.ProductType,
			TicketType:     v.TicketType,
			Day:            v.Day,
			Name:           v.Name,
			Identity:       v.Identity,
			SourceType:     v.SourceType,
			RefundAmount:   v.RefundAmount,
			RefundFee:      v.RefundFee,
			RefundNum:      v.RefundNum,
			StockBatchInfo: list,
		}
		refundProductTicketsToC = append(refundProductTicketsToC, refundProductTicketToC)
	}
	return &OrderRefundInfoToC{
		RefundInfoToC: &RefundInfoToC{
			RefundID:                req.BlockRefundInfo.RefundId,
			OrderID:                 req.BlockRefundInfo.OrderId,
			RefundAmount:            req.BlockRefundInfo.RefundAmount,
			RefundFee:               req.BlockRefundInfo.RefundFee,
			RefundStatus:            req.BlockRefundInfo.RefundStatus,
			RefundType:              req.BlockRefundInfo.RefundType,
			TradeNo:                 req.BlockRefundInfo.TradeNo,
			RefundTime:              req.BlockRefundInfo.RefundTime,
			Remark:                  req.BlockRefundInfo.Remark,
			FailMessage:             req.BlockRefundInfo.FailMessage,
			CommissionSettledStatus: req.BlockRefundInfo.CommissionSettledStatus,
			StockCertificate:        req.BlockRefundInfo.StockCertificate,
			ProductSkuName:          req.BlockRefundInfo.ProductSkuName,
			UserID:                  req.BlockRefundInfo.UserId,
			Username:                req.BlockRefundInfo.Username,
			BillStatus:              req.BlockRefundInfo.BillStatus,
			OrderGroupID:            req.BlockRefundInfo.OrderGroupId,
		},
		RefundProductTicketToC: refundProductTicketsToC,
	}
}

func BuildRepayRollBack(req *pb.RepayRollBackRequest) []*ActiveInfo {
	if req == nil {
		logrus.Errorf("BuildRepayRollBack has nil field")
		return nil
	}
	list := make([]*ActiveInfo, 0, len(req.Data))

	for _, v := range req.Data {

		tokenID := fmt.Sprintf("%v-%v-%v", v.StockBatchNumber, v.Account, v.Org)
		one := &ActiveInfo{
			OrderID:           v.OrderGroupId,
			BatchID:           v.BatchId,
			Periods:           v.Periods,
			TotalPeriods:      v.TotalPeriods,
			TokenID:           tokenID,
			Amount:            v.Amount,
			TradeNo:           v.TradeNo,
			AvailableTotalNum: v.AvailableTotalNum,
			TotalRepayment:    v.TotalRepayment,
		}

		list = append(list, one)
	}
	return list
}

func BuildStockBatchInfoUpdate(req *pb.StockBatchInfoUpdateRequest) *UpdateStockInfo {
	if req == nil {
		logrus.Errorf("BuildStockBatchInfoUpdate has nil field")
		return nil
	}

	return &UpdateStockInfo{
		StockInfo: StockInfo{
			PurchaseBeginTime:   req.PurchaseBeginTime,
			PurchaseEndTime:     req.PurchaseEndTime,
			StockEnterBeginTime: req.EnterBeginTime,
			StockEnterEndTime:   req.EnterEndTime,
		},
	}
}

func BuildProjectBidding(req *pb.ProjectBiddingRequest) *ProjectBidding {
	if req == nil {
		logrus.Errorf("BuildStockBatchInfoUpdate has nil field")
		return nil
	}

	// 转换为时间字符串
	// timeFromTimestamp := tFromTimestamp.Format(layout)

	return &ProjectBidding{
		ProjectID:            strconv.FormatInt(req.ProjectId, 10),
		ParentProjectID:      strconv.FormatInt(req.ParentProjectId, 10),
		ProjectType:          req.ProjectType,
		ProjectNumber:        req.ProjectNumber,
		ExchangeID:           strconv.FormatInt(req.ExchangeId, 10),
		CompanyID:            strconv.FormatInt(req.CompanyId, 10),
		ProjectName:          req.ProjectName,
		ProjectGrade:         strconv.FormatInt(int64(req.ProjectGrade), 10),
		ProjectDescription:   req.ProjectDescription,
		ProjectAnnex:         req.ProjectAnnex,
		ProjectStatus:        req.ProjectStatus,
		NodeDescription:      req.NodeDescription,
		ProjectChainUniqueId: req.ProjectChainUniqueId,
		ReviewTime:           ParseUTC8Time(req.ReviewTime),
		// NodeDescription: req.
	}
}

func BuildInstrument(req *pb.InstrumentTicketRequest) *InstrumentData {
	if req == nil {
		logrus.Errorf("BuildStockBatchInfoUpdate has nil field")
		return nil
	}
	return &InstrumentData{
		InstrumentID:         strconv.FormatInt(req.InstrumentId, 10),
		ScenicID:             strconv.FormatInt(req.ScenicId, 10),
		GoodsID:              strconv.FormatInt(req.GoodsId, 10),
		ProductType:          req.ProductType,
		TicketType:           req.TicketType,
		ScenicProvinceCode:   strconv.FormatInt(int64(req.ScenicProvinceCode), 10),
		ScenicCityCode:       strconv.FormatInt(int64(req.ScenicCityCode), 10),
		ScenicAreaCode:       strconv.FormatInt(int64(req.ScenicAreaCode), 10),
		BuyStartDate:         req.BuyStartDate,
		BuyEndDate:           req.BuyEndDate,
		UseStartDate:         req.UseStartDate,
		UseEndDate:           req.UseEndDate,
		ProjectChainUniqueId: req.ProjectChainUniqueId,
		Instrument: &InstrumentInfo{
			InstrumentID:          strconv.FormatInt(req.Instrument.InstrumentId, 10),
			ProjectID:             strconv.FormatInt(req.Instrument.ProjectId, 10),
			ParentInstrumentID:    strconv.FormatInt(req.Instrument.ParentInstrumentId, 10),
			CompanyID:             strconv.FormatInt(req.Instrument.CompanyId, 10),
			AssetType:             req.Instrument.AssetType,
			InstrumentType:        req.Instrument.InstrumentType,
			InstrumentName:        req.Instrument.InstrumentName,
			InstrumentImages:      req.Instrument.InstrumentImages,
			InstrumentVideo:       strconv.FormatInt(req.Instrument.InstrumentVideo, 10),
			InstrumentDescription: req.Instrument.InstrumentDescription,
			InstrumentAnnex:       req.Instrument.InstrumentAnnex,
			ContactName:           req.Instrument.ContactName,
			ContactPhone:          req.Instrument.ContactPhone,
			InstrumentStatus:      req.Instrument.InstrumentStatus,
			InstrumentBiddingRule: &InstrumentBiddingRule{
				InstrumentID:        strconv.FormatInt(req.Instrument.InstrumentBiddingRule.InstrumentId, 10),
				InstrumentGrade:     strconv.FormatInt(int64(req.Instrument.InstrumentBiddingRule.InstrumentGrade), 10),
				BiddingStartTime:    req.Instrument.InstrumentBiddingRule.BiddingStartTime,
				BiddingEndTime:      req.Instrument.InstrumentBiddingRule.BiddingEndTime,
				BiddingDelayTime:    strconv.FormatInt(req.Instrument.InstrumentBiddingRule.BiddingDelayTime, 10),
				BiddingType:         req.Instrument.InstrumentBiddingRule.BiddingType,
				BiddingStartPrice:   req.Instrument.InstrumentBiddingRule.BiddingStartPrice,
				BiddingDealRule:     strconv.FormatInt(int64(req.Instrument.InstrumentBiddingRule.BiddingDealRule), 10),
				BiddingDeposit:      req.Instrument.InstrumentBiddingRule.BiddingDeposit,
				BiddingChanges:      req.Instrument.InstrumentBiddingRule.BiddingChanges,
				InstrumentQuantity:  strconv.FormatInt(req.Instrument.InstrumentBiddingRule.InstrumentQuantity, 10),
				BiddingDescription:  req.Instrument.InstrumentBiddingRule.BiddingDescription,
				MarketPrice:         req.Instrument.InstrumentBiddingRule.MarketPrice,
				BiddingDepositRatio: req.Instrument.InstrumentBiddingRule.BiddingDepositRatio,
				LowestPrice:         req.Instrument.InstrumentBiddingRule.LowestPrice,
			},
			InstrumentBiddingSales: &InstrumentBiddingSales{
				InstrumentID:             strconv.FormatInt(req.GetInstrument().GetInstrumentBiddingSales().GetInstrumentId(), 10),
				InstrumentViews:          strconv.FormatInt(req.GetInstrument().GetInstrumentBiddingSales().GetInstrumentViews(), 10),
				RegistrationQuantity:     strconv.FormatInt(req.GetInstrument().GetInstrumentBiddingSales().GetRegistrationQuantity(), 10),
				BidQuantity:              strconv.FormatInt(req.GetInstrument().GetInstrumentBiddingSales().GetBidQuantity(), 10),
				CurrentPrice:             req.GetInstrument().GetInstrumentBiddingSales().GetCurrentPrice(),
				CurrentRemainingQuantity: strconv.FormatInt(req.GetInstrument().GetInstrumentBiddingSales().GetCurrentRemainingQuantity(), 10),
				EstimatedStartTime:       req.GetInstrument().GetInstrumentBiddingSales().GetEstimatedStartTime(),
				EstimatedEndTime:         req.GetInstrument().GetInstrumentBiddingSales().GetEstimatedEndTime(),
			},
		},
	}
}

func BuildMarginOrder(req *pb.MarginOrderRequest) *StoreMarginOrder {
	return &StoreMarginOrder{
		ProjectID:   req.GetExchangeId(),
		ProjectName: req.GetProjectName(),
		Exchange: &Exchange{
			ExchangeID:   req.GetExchangeId(),
			ExchangeName: req.GetExchangeName(),
		},
		Instrument: &Instrument{
			InstrumentID:   req.GetInstrumentId(),
			InstrumentName: req.GetInstrumentName(),
			SellerID:       req.GetSellerId(),
			SellerName:     req.GetSellerName(),
			BuyerID:        req.GetBuyerId(),
			BuyerName:      req.GetBuyerName(),
			Bidbond: &Bidbond{
				BidbondAmount: req.GetBidbondAmount(),
				OutTradeNo:    req.GetOutTradeNo(),
				TradeNo:       req.GetTradeNo(),
				CreateTime:    req.GetCreateTime(),
				OrderID:       req.GetOrderId(),
			},
		},
	}
}

func BuildMarginOrderOver(req *pb.MarginOrderOverRequest) *StoreMarginPayment {
	return &StoreMarginPayment{
		ProjectID:   req.GetExchangeId(),
		ProjectName: req.GetExchangeName(),
		Exchange: &Exchange{
			ExchangeID:   req.GetExchangeId(),
			ExchangeName: req.GetExchangeName(),
		},
		Instrument: &InstrumentPayMent{
			InstrumentID:   req.GetInstrumentId(),
			InstrumentName: req.GetInstrumentName(),
			SellerID:       req.GetSellerId(),
			SellerName:     req.GetSellerName(),
			BuyerID:        req.GetBuyerId(),
			BuyerName:      req.GetBuyerName(),
			Bidbond: &BidbondPayment{
				BidbondAmount:  req.GetBidbondAmount(),
				OutTradeNo:     req.GetOutTradeNo(),
				TradeNo:        req.GetTradeNo(),
				PaymentTradeNo: req.GetPaymentTradeNo(),
				PaymentTime:    req.GetPaymentTime(),
				BidNumber:      req.GetBidNumber(),
			},
		},
	}
}

func BuildCreditAssessment(req *pb.CreditAssessmentRequest) *CreditRating {
	return &CreditRating{
		LegalName:          req.GetLegalName(),
		RegistrationNumber: req.GetRegistrationNumber(),
		Score:              req.GetScore(),
		Level:              req.GetLevel(),
		AssessmentTime:     req.GetAssessmentTime(),
	}
}

func BuildLoanApplication(req *pb.LoanRequestRequest) *LoanApplication {
	return &LoanApplication{
		ApplicationNo:              req.GetApplicationNo(),
		LegalName:                  req.GetLegalName(),
		DebtorRegistrationNumber:   req.GetDebtorRegistrationNumber(),
		OwnersEquity:               req.GetOwnersEquity(),
		BusinessScope:              req.GetBusinessScope(),
		BusinessEconomicType:       req.GetBusinessEconomicType(),
		BusinessPhone:              req.GetBusinessPhone(),
		BusinessFax:                req.GetBusinessFax(),
		BusinessEmail:              req.GetBusinessEmail(),
		CreditorName:               req.GetCreditorName(),
		CreditorRegistrationNumber: req.GetCreditorRegistrationNumber(),
		FinancingPurpose:           req.GetFinancingPurpose(),
		Level:                      req.GetLevel(),
		ProposedCreditLimit:        req.GetProposedCreditLimit(),
		CreateTime:                 req.GetCreateTime(),
	}
}

func BuildLoanApproval(req *pb.LoanReviewApprovedRequest) *LoanApproval {
	return &LoanApproval{
		ApplicationNo:              req.GetApplicationNo(),
		LegalName:                  req.GetLegalName(),
		DebtorRegistrationNumber:   req.GetDebtorRegistrationNumber(),
		OwnersEquity:               req.GetOwnersEquity(),
		BusinessScope:              req.GetBusinessScope(),
		BusinessEconomicType:       req.GetBusinessEconomicType(),
		BusinessPhone:              req.GetBusinessPhone(),
		BusinessFax:                req.GetBusinessFax(),
		BusinessEmail:              req.GetBusinessEmail(),
		CreditorName:               req.GetCreditorName(),
		CreditorRegistrationNumber: req.GetCreditorRegistrationNumber(),
		FinancingPurpose:           req.GetFinancingPurpose(),
		Level:                      req.GetLevel(),
		ProposedCreditLimit:        req.GetProposedCreditLimit(),
		CreateTime:                 req.GetCreateTime(),
		CreditStatus:               req.GetCreditStatus(),
		CreditLine:                 req.GetCreditLine(),
		ExpiresTime:                req.GetExpiresTime(),
		RepaymentType:              req.GetRepaymentType(),
		Rate:                       req.GetRate(),
		Periods:                    req.GetPeriods(),
		MonthlyRepaymentDay:        req.GetMonthlyRepaymentDay(),
		OverdueRate:                req.GetOverdueRate(),
		RepaymentName:              req.GetRepaymentName(),
		RepaymentBankAccount:       req.GetRepaymentBankAccount(),
		RepaymentBankName:          req.GetRepaymentBankName(),
		ModifyTime:                 req.GetModifyTime(),
	}
}

func BuildLoanApprovalFailed(req *pb.LoanReviewDeniedRequest) *LoanApprovalFailed {
	return &LoanApprovalFailed{
		ApplicationNo:              req.GetApplicationNo(),
		LegalName:                  req.GetLegalName(),
		DebtorRegistrationNumber:   req.GetDebtorRegistrationNumber(),
		OwnersEquity:               req.GetOwnersEquity(),
		BusinessScope:              req.GetBusinessScope(),
		BusinessEconomicType:       req.GetBusinessEconomicType(),
		BusinessPhone:              req.GetBusinessPhone(),
		BusinessFax:                req.GetBusinessFax(),
		BusinessEmail:              req.GetBusinessEmail(),
		CreditorName:               req.GetCreditorName(),
		CreditorRegistrationNumber: req.GetCreditorRegistrationNumber(),
		FinancingPurpose:           req.GetFinancingPurpose(),
		Level:                      req.GetLevel(),
		ProposedCreditLimit:        req.GetProposedCreditLimit(),
		CreateTime:                 req.GetCreateTime(),
		CreditStatus:               req.GetCreditStatus(),
		CreditAuditRemark:          req.GetCreditAuditRemark(),
		ModifyTime:                 req.GetModifyTime(),
	}
}

func BuildBorrowApplication(req *pb.LoanApplicationRequest) *BorrowApplication {
	return &BorrowApplication{
		OutTradeNo:                 req.GetOutTradeNo(),
		LoanAmount:                 req.GetLoanAmount(),
		DebtorName:                 req.GetDebtorName(),
		DebtorRegistrationNumber:   req.GetDebtorRegistrationNumber(),
		CreditorName:               req.GetCreditorName(),
		CreditorRegistrationNumber: req.GetCreditorRegistrationNumber(),
		ExchangeName:               req.GetExchangeName(),
		ExchangeRegistrationNumber: req.GetExchangeRegistrationNumber(),
		ReceiverCorporateName:      req.GetReceiverCorporateName(),
		ReceiverBankAccountNo:      req.GetReceiverBankAccountNo(),
		ReceiverBankName:           req.GetReceiverBankName(),
		TradeOrderFileURL:          req.GetTradeOrderFileUrl(),
		CautionMoneyFileURL:        req.GetCautionMoneyFileUrl(),
		PaymentBookFileURL:         req.GetPaymentBookFileUrl(),
		ApplyTime:                  req.GetApplyTime(),
	}
}

func BuildBorrowApproval(req *pb.LoanApprovalPaymentRequest) *BorrowApproval {
	return &BorrowApproval{
		OutTradeNo:                 req.GetOutTradeNo(),
		LoanAmount:                 req.GetLoanAmount(),
		DebtorName:                 req.GetDebtorName(),
		DebtorRegistrationNumber:   req.GetDebtorRegistrationNumber(),
		CreditorName:               req.GetCreditorName(),
		CreditorRegistrationNumber: req.GetCreditorRegistrationNumber(),
		ExchangeName:               req.GetExchangeName(),
		ExchangeRegistrationNumber: req.GetExchangeRegistrationNumber(),
		ReceiverCorporateName:      req.GetReceiverCorporateName(),
		ReceiverBankAccountNo:      req.GetReceiverBankAccountNo(),
		ReceiverBankName:           req.GetReceiverBankName(),
		TradeOrderFileURL:          req.GetTradeOrderFileUrl(),
		CautionMoneyFileURL:        req.GetCautionMoneyFileUrl(),
		PaymentBookFileURL:         req.GetPaymentBookFileUrl(),
		PayBankSlipFileURL:         req.GetPayBankSlipFileUrl(),
		PayBankTradeNo:             req.GetPayBankTradeNo(),
		PayCorporateName:           req.GetPayCorporateName(),
		PayBankAccountNo:           req.GetPayBankAccountNo(),
		PayBankName:                req.GetPayBankName(),
		PayAuditTime:               req.GetPayAuditTime(),
	}
}

func BuildRepaymentOrder(req *pb.PaymentScheduleGenerationRequest) *RepaymentOrder {
	return &RepaymentOrder{
		RepaymentNo:                req.GetOutTradeNo(),
		DebtorName:                 req.GetDebtorName(),
		DebtorRegistrationNumber:   req.GetDebtorRegistrationNumber(),
		CreditorName:               req.GetCreditorName(),
		CreditorRegistrationNumber: req.GetCreditorRegistrationNumber(),
		TotalPrincipal:             req.GetAmount(),
		RepaymentType:              req.GetRepaymentType(),
		Rate:                       req.GetRate(),
		CreateTime:                 req.GetCreateTime(),
	}
}

func BuildRepaymentApplication(req *pb.RepaymentApplicationRequest) *RepaymentApplication {
	return &RepaymentApplication{
		RepaymentNo:                req.GetOutTradeNo(),
		DebtorName:                 req.GetDebtorName(),
		DebtorRegistrationNumber:   req.GetDebtorRegistrationNumber(),
		CreditorName:               req.GetCreditorName(),
		CreditorRegistrationNumber: req.GetCreditorRegistrationNumber(),
		TotalPrincipal:             req.GetAmount(),
		Principal:                  req.GetPrincipal(),
		RepaymentType:              req.GetRepaymentType(),
		Rate:                       req.GetRate(),
		Periods:                    req.GetPeriods(),
		RepaymentDate:              req.GetRepaymentDate(),
		IsOverdue:                  req.GetIsOverdue(),
		PenaltyInterest:            req.GetPenaltyInterest(),
		Amount:                     req.GetAmount(),
		BankSlip:                   req.GetBankSlip(),
		PayBankTradeNo:             req.GetPayBankTradeNo(),
		PayUnitName:                req.GetPayUnitName(),
		PayBankAccountNo:           req.GetPayBankAccountNo(),
		PayBankName:                req.GetPayBankName(),
		RepaymentName:              req.GetRepaymentName(),
		RepaymentBankAccount:       req.GetRepaymentBankAccount(),
		RepaymentBankName:          req.GetRepaymentBankName(),
		CreateTime:                 req.GetApplyTime(),
	}
}

func BuildStoreRepaymentApproval(req *pb.RepaymentApprovedRequest) *StoreRepaymentApproval {
	return &StoreRepaymentApproval{
		RepaymentNo:                req.GetOutTradeNo(),
		DebtorName:                 req.GetDebtorName(),
		DebtorRegistrationNumber:   req.GetDebtorRegistrationNumber(),
		CreditorName:               req.GetCreditorName(),
		CreditorRegistrationNumber: req.GetCreditorRegistrationNumber(),
		TotalPrincipal:             req.GetAmount(),
		Principal:                  req.GetPrincipal(),
		RepaymentType:              req.GetRepaymentType(),
		Rate:                       req.GetRate(),
		Periods:                    req.GetPeriods(),
		RepaymentDate:              req.GetRepaymentDate(),
		IsOverdue:                  req.GetIsOverdue(),
		PenaltyInterest:            req.GetPenaltyInterest(),
		Amount:                     req.GetRepaymentAmount(),
		BankSlip:                   req.GetBankSlip(),
		PayBankTradeNo:             req.GetPayBankTradeNo(),
		PayUnitName:                req.GetPayUnitName(),
		PayBankAccountNo:           req.GetPayBankAccountNo(),
		PayBankName:                req.GetPayBankName(),
		RepaymentName:              req.GetRepaymentName(),
		RepaymentBankAccount:       req.GetRepaymentBankAccount(),
		RepaymentBankName:          req.GetRepaymentBankName(),
		RepaymentStatus:            req.GetApprovalResult(),
		ModifyTime:                 req.GetApprovalTime(),
	}
}

func BuildStoreRepaymentApprovalFailed(req *pb.RepaymentDeniedRequest) *StoreRepaymentApprovalFailed {
	return &StoreRepaymentApprovalFailed{
		RepaymentNo:                req.GetOutTradeNo(),
		DebtorName:                 req.GetDebtorName(),
		DebtorRegistrationNumber:   req.GetDebtorRegistrationNumber(),
		CreditorName:               req.GetCreditorName(),
		CreditorRegistrationNumber: req.GetCreditorRegistrationNumber(),
		TotalPrincipal:             req.GetAmount(),
		Principal:                  req.GetPrincipal(),
		RepaymentType:              req.GetRepaymentType(),
		Rate:                       req.GetRate(),
		Periods:                    req.GetPeriods(),
		RepaymentDate:              req.GetRepaymentDate(),
		IsOverdue:                  req.GetIsOverdue(),
		PenaltyInterest:            req.GetPenaltyInterest(),
		Amount:                     req.GetRepaymentAmount(),
		BankSlip:                   req.GetBankSlip(),
		PayBankTradeNo:             req.GetPayBankTradeNo(),
		PayUnitName:                req.GetPayUnitName(),
		PayBankAccountNo:           req.GetPayBankAccountNo(),
		PayBankName:                req.GetPayBankName(),
		RepaymentName:              req.GetRepaymentName(),
		RepaymentBankAccount:       req.GetRepaymentBankAccount(),
		RepaymentBankName:          req.GetRepaymentBankName(),
		RepaymentStatus:            req.GetApprovalResult(),
		Remark:                     req.GetRejectionReason(),
		ModifyTime:                 req.GetApprovalTime(),
	}
}

func BuildStoreReRepaymentApplication(req *pb.RepaymentReapplyRequest) *RepaymentReApplication {
	return &RepaymentReApplication{
		RepaymentNo:                req.GetOutTradeNo(),
		DebtorName:                 req.GetDebtorName(),
		DebtorRegistrationNumber:   req.GetDebtorRegistrationNumber(),
		CreditorName:               req.GetCreditorName(),
		CreditorRegistrationNumber: req.GetCreditorRegistrationNumber(),
		TotalPrincipal:             req.GetAmount(),
		Principal:                  req.GetPrincipal(),
		RepaymentType:              req.GetRepaymentType(),
		Rate:                       req.GetRate(),
		Periods:                    req.GetPeriods(),
		RepaymentDate:              req.GetRepaymentDate(),
		IsOverdue:                  req.GetIsOverdue(),
		PenaltyInterest:            req.GetPenaltyInterest(),
		Amount:                     req.GetRepaymentAmount(),
		BankSlip:                   req.GetBankSlip(),
		PayBankTradeNo:             req.GetPayBankTradeNo(),
		PayUnitName:                req.GetPayUnitName(),
		PayBankAccountNo:           req.GetPayBankAccountNo(),
		PayBankName:                req.GetPayBankName(),
		RepaymentName:              req.GetRepaymentName(),
		RepaymentBankAccount:       req.GetRepaymentBankAccount(),
		RepaymentBankName:          req.GetRepaymentBankName(),
		CreateTime:                 req.GetApplyTime(),
	}
}

func BuildStoreLoanReApplication(req *pb.LoanReapplyRequest) *LoanReApplication {
	return &LoanReApplication{
		ApplicationNo:              req.GetApplicationNo(),
		LegalName:                  req.GetLegalName(),
		DebtorRegistrationNumber:   req.GetDebtorRegistrationNumber(),
		OwnersEquity:               req.GetOwnersEquity(),
		BusinessScope:              req.GetBusinessScope(),
		BusinessEconomicType:       req.GetBusinessEconomicType(),
		BusinessPhone:              req.GetBusinessPhone(),
		BusinessFax:                req.GetBusinessFax(),
		BusinessEmail:              req.GetBusinessEmail(),
		CreditorName:               req.GetCreditorName(),
		CreditorRegistrationNumber: req.GetCreditorRegistrationNumber(),
		FinancingPurpose:           req.GetFinancingPurpose(),
		Level:                      req.GetLevel(),
		ProposedCreditLimit:        req.GetProposedCreditLimit(),
		CreateTime:                 req.GetCreateTime(),
	}
}

func BuildConvertToInvoice(req *pb.ConvertToInvoiceRequest) *ConvertToInvoice {
	return &ConvertToInvoice{
		ProjectID:         req.GetProjectId(),
		ProjectName:       req.GetProjectName(),
		ExchangeID:        req.GetExchangeId(),
		ExchangeName:      req.GetExchangeName(),
		SellerID:          req.GetSellerId(),
		SellerName:        req.GetSellerName(),
		InstrumentID:      req.GetInstrumentId(),
		InstrumentName:    req.GetInstrumentName(),
		BuyerID:           req.GetBuyerId(),
		BuyerName:         req.GetBuyerName(),
		BidbondAmount:     req.GetBidbondAmount(),
		TransactionAmount: req.GetTransactionAmount(),
		ConvertedAmount:   req.GetConvertedAmount(),
		Status:            req.GetStatus(),
		UpdateTime:        req.GetUpdateTime(),
	}
}

func BuildRefundBalance(req *pb.RefundBalanceRequest) *RefundBalance {
	return &RefundBalance{
		ProjectID:         req.GetProjectId(),
		ProjectName:       req.GetProjectName(),
		ExchangeID:        req.GetExchangeId(),
		ExchangeName:      req.GetExchangeName(),
		SellerID:          req.GetSellerId(),
		SellerName:        req.GetSellerName(),
		InstrumentID:      req.GetInstrumentId(),
		InstrumentName:    req.GetInstrumentName(),
		BuyerID:           req.GetBuyerId(),
		BuyerName:         req.GetBuyerName(),
		BidbondAmount:     req.GetBidbondAmount(),
		TransactionAmount: req.GetTransactionAmount(),
		ConvertedAmount:   req.GetConvertedAmount(),
		RefundAmount:      req.GetRefundAmount(),
		Status:            req.GetStatus(),
		UpdateTime:        req.GetUpdateTime(),
	}
}

func BuildFullRefund(req *pb.FullRefundRequest) *FullRefund {
	return &FullRefund{
		ProjectID:      req.GetProjectId(),
		ProjectName:    req.GetProjectName(),
		ExchangeID:     req.GetExchangeId(),
		ExchangeName:   req.GetExchangeName(),
		SellerID:       req.GetSellerId(),
		SellerName:     req.GetSellerName(),
		InstrumentID:   req.GetInstrumentId(),
		InstrumentName: req.GetInstrumentName(),
		BuyerID:        req.GetBuyerId(),
		BuyerName:      req.GetBuyerName(),
		BidbondAmount:  req.GetBidbondAmount(),
		RefundAmount:   req.GetRefundAmount(),
		Status:         req.GetStatus(),
		UpdateTime:     req.GetUpdateTime(),
	}
}

func BuildStoreTradeCharge(req *pb.TradeChargeRequest) *TradeCharge {
	return &TradeCharge{
		TradeChargeID:      strconv.FormatInt(req.GetTradeChargeId(), 10),
		TradeID:            strconv.FormatInt(req.GetTradeId(), 10),
		ChargeObject:       req.GetChargeObject(),
		CompanyName:        req.GetCompanyName(),
		ChargeAmount:       strconv.FormatFloat(req.GetChargeAmount(), 'f', 10, 64),
		PaymentStatus:      req.GetPaymentStatus(),
		CreateTime:         req.GetCreateTime(),
		ExpiryDatetime:     req.GetExpiryDatetime(),
		PayeeAccountName:   req.GetPayerAccountName(),
		PayeeAccountNumber: req.GetPayeeAccountNumber(),
		PayeeBankName:      req.GetPayeeBankName(),
		TransactionFlowID:  req.GetTransactionFlowId(),
		PayerAccountName:   req.GetPayeeAccountName(),
		PayerAccountNumber: req.GetPayeeAccountNumber(),
		PayerBankName:      req.GetPayerBankName(),
		FileURL:            req.GetFileUrl(),
		SubmitTime:         req.GetSubmitTime(),
		AuditResult:        req.GetAuditResult(),
		Remark:             req.GetRemark(),
		AuditTime:          req.GetAuditTime(),
	}
}

func BuildInstrumentOrder(req *pb.InstrumentOrderRequest) *InstrumentOrder {
	return &InstrumentOrder{
		TradeID:        strconv.FormatInt(req.GetTradeId(), 10),
		ExchangeID:     strconv.FormatInt(req.GetExchangeId(), 10),
		SellerID:       strconv.FormatInt(req.GetSellerId(), 10),
		BuyerID:        strconv.FormatInt(req.GetBuyerId(), 10),
		TradeAmount:    strconv.FormatFloat(req.GetTradeAmount(), 'f', 10, 64),
		TradeType:      req.GetTradeType(),
		TradeStatus:    req.GetTradeStatus(),
		TradeDatetime:  req.GetTradeDatetime(),
		CreateTime:     req.GetCreateTime(),
		UpdateTime:     req.GetUpdateTime(),
		CompanyID:      strconv.FormatInt(req.GetCompanyId(), 10),
		CompanyName:    req.GetCompanyName(),
		ProjectID:      strconv.FormatInt(req.GetProjectId(), 10),
		ProjectName:    req.GetProjectName(),
		ExchangeName:   req.GetExchangeName(),
		InstrumentID:   strconv.FormatInt(req.GetInstrumentId(), 10),
		InstrumentName: req.GetInstrumentName(),
		SettlementType: req.GetSettlementType(),
		SettlementId:   strconv.FormatInt(req.GetSettlementId(), 10),
	}
}

func CheckTicket(req *pb.CreateStockRequest) bool {
	if req.GetSimpleTicket() == nil ||
		req.GetSimpleTicket().GetTicketGoods() == nil ||
		req.GetSimpleTicket().GetTicketGoods().GetRuleIssue() == nil ||
		req.GetSimpleTicket().GetTicketGoods().GetRuleCheck() == nil ||
		req.GetSimpleTicket().GetTicketGoods().GetRuleRetreat() == nil ||
		req.GetSimpleTicket().GetTicketStock() == nil {
		return false
	}
	if req.GetSimpleTicket().GetTicketGoods().GetSetSelfPrice() == nil {
		req.GetSimpleTicket().GetTicketGoods().SetSelfPrice = &pb.SetSelfPrice{
			Price: &pb.PriceInfo{},
		}
	}
	if req.GetSimpleTicket().GetTicketGoods().GetTimeSharing() == nil {
		req.GetSimpleTicket().GetTicketGoods().TimeSharing = &pb.TimeSharing{}
	}

	if req.GetBlockChainTicketCheck() == nil {
		req.BlockChainTicketCheck = &pb.BlockChainTicketCheckRequest{}
	}

	if req.GetTicketIssuance() == nil {
		req.TicketIssuance = &pb.BlockTicketIssuanceRequest{}
	}

	return true
}

func CheckPriceInfo(req *pb.SetSelfPriceInfoRequest) bool {

	if req.PriceInfoDetailsList == nil || req.PriceInfoDetailsList.Price == nil {
		return false
	}

	if req.PriceInfoDetailsList.AddGroupId == nil {
		req.PriceInfoDetailsList.AddGroupId = []string{}
	}
	if req.PriceInfoDetailsList.DelGroupId == nil {
		req.PriceInfoDetailsList.DelGroupId = []string{}
	}
	if req.PriceInfoDetailsList.GroupId == nil {
		req.PriceInfoDetailsList.GroupId = []string{}
	}

	return true
}

func CheckRefundTocInfo(req *pb.BlockOrderRefundRequest) bool {
	if req == nil || req.BlockRefundInfo == nil {
		return false
	}
	return true

}

func ParseUTC8Time(timestampMilli int64) string {
	// 从时间戳转换回时间对象
	tFromTimestamp := time.Unix(0, timestampMilli*int64(time.Millisecond))
	// 设置为你需要的时区，例如北京时间（UTC+8）
	location, err := time.LoadLocation("Asia/Shanghai") // 根据需要修改时区
	if err != nil {
		// 处理错误
		return ""
	}
	tFromTimestamp = tFromTimestamp.In(location)

	// 转换为时间字符串

	return tFromTimestamp.Format(layout)
}

func ConverTime(timeStr string) int64 {

	// 设置为你需要的时区，例如北京时间（UTC+8）
	location, err := time.LoadLocation("Asia/Shanghai") // 根据需要修改时区
	if err != nil {
		// 处理错误
		return 0 // 或者其他适当的错误处理
	}
	// 解析时间字符串并指定时区
	t, _ := time.ParseInLocation(layout, timeStr, location)

	// 转换为时间字符串

	return t.UnixMilli()
}
