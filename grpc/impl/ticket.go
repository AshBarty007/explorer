package impl

import (
	common "blockchain_services/config"
	"blockchain_services/ethclient"
	"blockchain_services/grpc/pb"
	"blockchain_services/ipfs"
	postgres "blockchain_services/postgres"
	"blockchain_services/redis"
	"context"
	"encoding/json"
	"fmt"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"log"
	"time"
)

type TicketsServer struct {
}

func NewTicketsServer() *TicketsServer {
	return &TicketsServer{}
}

func (t *TicketsServer) MainOrderTicket(ctx context.Context, req *pb.RequestMessage) (*pb.ResponseMessage, error) {

	logrus.Infof("已接收,  method id is %v", req.MessageType)

	rsp := new(pb.ResponseMessage)
	var err error

	switch req.MessageType {
	case pb.MessageType_TICKET_CHECK: //门票核销
		var data *pb.BlockChainTicketCheckResponse
		tmp := BuildCheckData(req.GetTicketCheck())
		input := fmt.Sprintf("%+v", tmp)
		txId, err := t.UploadToBlockChain(input)
		if err != nil {
			rsp.Code = common.RespStatusServiceError
			rsp.Message = common.MsgFlags[rsp.Code]
			break
		}
		data = &pb.BlockChainTicketCheckResponse{
			Code: common.RespStatusSuccess,
			Msg:  "OK",
			Data: &pb.BlockChainTicketCheckResponse_TxData{
				TxId: txId,
			},
		}

		rsp.Code = common.RespStatusSuccess
		rsp.MessageType = pb.MessageType_TICKET_CHECK
		rsp.Data = &pb.ResponseMessage_TicketCheck{TicketCheck: data}
	case pb.MessageType_CREATE_STOCK: //库存创建
		var data *pb.CreateStockResponse
		tmp := BuildTickets(req.GetCreateStock())
		input := fmt.Sprintf("%+v", tmp)
		txId, err := t.UploadToBlockChain(input)
		if err != nil {
			rsp.Code = common.RespStatusServiceError
			rsp.Message = common.MsgFlags[rsp.Code]
			break
		}
		data = &pb.CreateStockResponse{
			Code: common.RespStatusSuccess,
			Msg:  "OK",
			Data: &pb.TxData{
				TxId: txId,
			},
		}

		rsp.Code = common.RespStatusSuccess
		rsp.MessageType = pb.MessageType_CREATE_STOCK
		rsp.Data = &pb.ResponseMessage_CreateStock{CreateStock: data}
	case pb.MessageType_SET_SELF_PRICE: //设置价格策略
		var data *pb.SetSelfPriceInfoResponse
		tmp := BuildPriceInfo(req.GetSetSelfPriceInfo())
		input := fmt.Sprintf("%+v", tmp)
		txId, err := t.UploadToBlockChain(input)
		if err != nil {
			rsp.Code = common.RespStatusServiceError
			rsp.Message = common.MsgFlags[rsp.Code]
			break
		}
		data = &pb.SetSelfPriceInfoResponse{
			Code: common.RespStatusSuccess,
			Msg:  "OK",
			Data: &pb.TxData{
				TxId: txId,
			},
		}

		rsp.Code = common.RespStatusSuccess
		rsp.MessageType = pb.MessageType_SET_SELF_PRICE
		rsp.Data = &pb.ResponseMessage_SetSelfPriceInfo{SetSelfPriceInfo: data}
	case pb.MessageType_GENERATE_TICKET_ISSUANCE: //生成出票信息（票号/二维码/Sign签名）
		var data *pb.GenerateTicketIssuanceResponse
		tmp := BuildGenerateTicketIssuance(req.GetGenerateTicketIssuance())
		input := fmt.Sprintf("%+v", tmp)
		txId, err := t.UploadToBlockChain(input)
		if err != nil {
			rsp.Code = common.RespStatusServiceError
			rsp.Message = common.MsgFlags[rsp.Code]
			break
		}
		data = &pb.GenerateTicketIssuanceResponse{
			Code: common.RespStatusSuccess,
			Msg:  "OK",
			Data: &pb.GenerateTicketIssuanceResponse_GenerateData{
				TxId: txId,
			},
		}

		rsp.Code = common.RespStatusSuccess
		rsp.MessageType = pb.MessageType_GENERATE_TICKET_ISSUANCE
		rsp.Data = &pb.ResponseMessage_Generate_Ticket_Issuance{Generate_Ticket_Issuance: data}
	case pb.MessageType_ORDER_TICKET: //出票信息
		var data *pb.BlockTicketIssuanceResponse
		tmp := BuildTicketIssuance(req.GetTicketIssuance())
		input := fmt.Sprintf("%+v", tmp)
		txId, err := t.UploadToBlockChain(input)
		if err != nil {
			rsp.Code = common.RespStatusServiceError
			rsp.Message = common.MsgFlags[rsp.Code]
			break
		}
		data = &pb.BlockTicketIssuanceResponse{
			Code: common.RespStatusSuccess,
			Msg:  "OK",
			Data: &pb.TxData{
				TxId: txId,
			},
		}

		rsp.Code = common.RespStatusSuccess
		rsp.MessageType = pb.MessageType_ORDER_TICKET
		rsp.Data = &pb.ResponseMessage_TicketIssuance{TicketIssuance: data}
	case pb.MessageType_ORDER_INFO: //订单信息
		var data *pb.OrderInfoResponse
		tmp := BuildOrderInfo(req.GetOrderInfo())
		input := fmt.Sprintf("%+v", tmp)
		txId, err := t.UploadToBlockChain(input)
		if err != nil {
			rsp.Code = common.RespStatusServiceError
			rsp.Message = common.MsgFlags[rsp.Code]
			break
		}
		data = &pb.OrderInfoResponse{
			Code: common.RespStatusSuccess,
			Msg:  "OK",
			Data: &pb.TxData{
				TxId: txId,
			},
		}
		rsp.Code = common.RespStatusSuccess
		rsp.MessageType = pb.MessageType_ORDER_INFO
		rsp.Data = &pb.ResponseMessage_OrderInfo{OrderInfo: data}
	case pb.MessageType_DISTRIBUTION_ORDER_INFO: //创建分销单
		var data *pb.DistributionOrderInfoResponse
		tmp := BuildDistributionOrderInfo(req.GetDistributionOrderInfo())
		input := fmt.Sprintf("%+v", tmp)
		txId, err := t.UploadToBlockChain(input)
		if err != nil {
			rsp.Code = common.RespStatusServiceError
			rsp.Message = common.MsgFlags[rsp.Code]
			break
		}
		data = &pb.DistributionOrderInfoResponse{
			Code: common.RespStatusSuccess,
			Msg:  "OK",
			Data: &pb.TxData{
				TxId: txId,
			},
		}
		rsp.Code = common.RespStatusSuccess
		rsp.MessageType = pb.MessageType_ORDER_INFO
		rsp.Data = &pb.ResponseMessage_DistributionOrderInfoResponse{DistributionOrderInfoResponse: data}
	case pb.MessageType_DISTRIBUTE_REFUND_CREATE: //分销退订单
		var data *pb.DistributeRefundCreateResponse
		tmp := BuildDistributionRefundInfo(req.GetDistributeRefundCreate())
		input := fmt.Sprintf("%+v", tmp)
		txId, err := t.UploadToBlockChain(input)
		if err != nil {
			rsp.Code = common.RespStatusServiceError
			rsp.Message = common.MsgFlags[rsp.Code]
			break
		}
		data = &pb.DistributeRefundCreateResponse{
			Code: common.RespStatusSuccess,
			Msg:  "OK",
			Data: &pb.TxData{
				TxId: txId,
			},
		}

		rsp.Code = common.RespStatusSuccess
		rsp.MessageType = pb.MessageType_DISTRIBUTE_REFUND_CREATE
		rsp.Data = &pb.ResponseMessage_DistributeRefundCreate{DistributeRefundCreate: data}
	case pb.MessageType_BLOCK_ORDER_REFUND: //C端退票同步信息接口
		var data *pb.BlockOrderRefundResponse
		tmp := BuildRefundTocInfo(req.GetBlockOrderRefund())
		input := fmt.Sprintf("%+v", tmp)
		txId, err := t.UploadToBlockChain(input)
		if err != nil {
			rsp.Code = common.RespStatusServiceError
			rsp.Message = common.MsgFlags[rsp.Code]
			break
		}
		data = &pb.BlockOrderRefundResponse{
			Code: common.RespStatusSuccess,
			Msg:  "OK",
			Data: &pb.TxData{
				TxId: txId,
			},
		}

		rsp.Code = common.RespStatusSuccess
		rsp.MessageType = pb.MessageType_BLOCK_ORDER_REFUND
		rsp.Data = &pb.ResponseMessage_BlockOrderRefundResponse{BlockOrderRefundResponse: data}
	case pb.MessageType_TICKET_STATUS_PUSH: //门票变更状态推送，包括门票核验，定时的过期检查...
		var data *pb.TicketStatusPushResponse
		tmp := BuildTicketStatus(req.GetTicketStatus())
		input := fmt.Sprintf("%+v", tmp)
		txId, err := t.UploadToBlockChain(input)
		if err != nil {
			rsp.Code = common.RespStatusServiceError
			rsp.Message = common.MsgFlags[rsp.Code]
			break
		}
		data = &pb.TicketStatusPushResponse{
			Code: common.RespStatusSuccess,
			Msg:  "OK",
			Data: &pb.TxData{
				TxId: txId,
			},
		}

		rsp.Code = common.RespStatusSuccess
		rsp.MessageType = pb.MessageType_TICKET_STATUS_PUSH
		rsp.Data = &pb.ResponseMessage_TicketStatusPushResponse{TicketStatusPushResponse: data}
	case pb.MessageType_INSTRUMENT_TICKET_REQUEST: //交易所票务标的
		var data *pb.InstrumentTicketResponse
		tmp := BuildInstrument(req.GetInstrumentTicketRequest())
		//input := fmt.Sprintf("%+v", tmp)
		in, _ := json.MarshalIndent(tmp, "", " ")
		txId, err := t.UploadToBlockChain(string(in))
		if err != nil {
			rsp.Code = common.RespStatusServiceError
			rsp.Message = common.MsgFlags[rsp.Code]
			break
		}
		data = &pb.InstrumentTicketResponse{
			Code: common.RespStatusSuccess,
			Msg:  "OK",
			Data: &pb.InstrumentTicketData{
				InstrumentTicketHash: &pb.InstrumentTicketHash{
					InstrumentId: req.GetInstrumentTicketRequest().InstrumentId,
					TxHash:       txId,
					//BlockHeight:  int64(resp.BlockNumber),
					OnChainTime: time.Now().UTC().UnixMilli(),
					ChainName:   ChainName,
				},
			},
		}

		rsp.Code = common.RespStatusSuccess
		rsp.MessageType = pb.MessageType_INSTRUMENT_TICKET_REQUEST
		rsp.Data = &pb.ResponseMessage_InstrumentTicketResponse{InstrumentTicketResponse: data}
	default:
		err = fmt.Errorf("not find method")
		logrus.Error(err)
		rsp.Code = common.RespStatusServiceError
	}

	if err != nil {
		err = common.GrpcResponseError(rsp.Code)
		return rsp, err
	}

	return rsp, nil
}

func (t *TicketsServer) UploadToBlockChain(content string) (txHash string, err error) {
	type AddressTest struct {
		gorm.Model
		Address string
		Key     string
	}

	ctx := context.Background()

	rc := redis.NewRedisClient()
	exists, err := rc.Exists(ctx, common.EthParams.AdminAddr).Result()
	if err != nil {
		return "", err
	}
	defer rc.Close()

	var key string
	if exists == 1 {
		key, err = rc.Get(ctx, common.EthParams.AdminAddr).Result()
		if err != nil {
			return "", err
		}
	} else {
		err := postgres.InitPgConn()
		if err != nil {
			return "", err
		}

		var p AddressTest
		err = postgres.Db.First(&p, "address = ?", common.EthParams.AdminAddr).Error
		if err != nil {
			return "", err
		}

		err = rc.Set(ctx, common.EthParams.AdminAddr, p.Key, 120*time.Second).Err()
		if err != nil {
			return "", err
		}

		key = p.Key
	}
	log.Println("设置(或读取)redis的值: ", key)

	ipfsHash, err := ipfs.UploadToIPFS(content)
	if err != nil {
		return "", err
	}
	log.Println("已返回ipfs哈希: ", ipfsHash)

	txHash, err = bs_eth.Call(key, "EVIDENCE", ipfsHash)
	if err != nil {
		return "", err
	}
	log.Println("已返回交易哈希: ", txHash)

	return txHash, nil
}
