package config

import (
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

const (
	ContractUnknownError    = 3000
	ContractAmountNotEnough = 3001
	ContractTicketExisted   = 3002
	ContractParseOjbError   = 3003
	ContractOnChainError    = 3004
	ContractParamIsNull     = 3005
	ContractParamError      = 3006
	ContractNotOwner        = 3007
	ContractObjNotFind      = 3008

	ContractInitError             = 4001
	ContractInitErrorInvalidToken = 4002
	ContractStopServer            = 4003

	RespStatusSuccess                     = 20000
	RespStatusUserNotExist                = 20001
	RespStatusNotEnoughBalanceCoins       = 20002
	RespStatusNotEnoughBalanceTickets     = 20003
	RespStatusNotEnoughBalanceRealTickets = 20004
	RespStatusNotEnoughAccount            = 20006
	RespStatusLowAuth                     = 20009
	RespStatusScenicIsExist               = 20010
	RespStatusScenicNotExist              = 20011
	RespStatusAlliancesNotExist           = 20012
	RespStatusScenicConfigNotYet          = 20013
	RespStatusCertRecordNotExist          = 20014
	RespStatusIssueRecordNotExist         = 20015

	RespStatusBadParams  = 30000
	RespStatusLoseParams = 30001

	RespStatusServiceError       = 50000
	RespStatusContractError      = 50001
	RespStatusDBError            = 50002
	RespStatusModifyConfigError  = 50003
	RespStatusNetWorkBuildError  = 50004
	RespStatusTicketNotFindError = 50005
)

var MsgFlags = map[int64]string{

	ContractUnknownError:    "未知错误",
	ContractAmountNotEnough: "票数量不足",
	ContractTicketExisted:   "门票已存在",
	ContractParseOjbError:   "解析JSON数据出错",
	ContractOnChainError:    "在链上操作失败",
	ContractParamIsNull:     "参数异常，字段或者对象缺失",
	ContractParamError:      "参数不合法",
	ContractNotOwner:        "无权限操作",
	ContractObjNotFind:      "参数异常，门票/订单不存在",

	ContractInitError:             "初始化失败",
	ContractInitErrorInvalidToken: "无效的身份",
	ContractStopServer:            "服务已关闭",

	RespStatusSuccess:                     "成功",
	RespStatusUserNotExist:                "用户不存在",
	RespStatusNotEnoughBalanceCoins:       "稳定币余额不足",
	RespStatusNotEnoughBalanceTickets:     "门票余额不足",
	RespStatusNotEnoughBalanceRealTickets: "门票余额不足, 补货中...",
	RespStatusNotEnoughAccount:            "可注册用户数量不足",
	RespStatusLowAuth:                     "权限不足",
	RespStatusScenicIsExist:               "景区已存在",
	RespStatusScenicNotExist:              "景区不存在",
	RespStatusAlliancesNotExist:           "联盟不存在(景区不属于任何联盟)",
	RespStatusScenicConfigNotYet:          "景区组织已启动或状态异常",
	RespStatusCertRecordNotExist:          "存证不存在",
	RespStatusIssueRecordNotExist:         "发行记录不存在",

	RespStatusBadParams:  "参数错误",
	RespStatusLoseParams: "缺少参数",

	RespStatusServiceError:      "服务错误",
	RespStatusContractError:     "合约错误",
	RespStatusDBError:           "数据库错误",
	RespStatusModifyConfigError: "服务修改配置失败",
	RespStatusNetWorkBuildError: "网络构建失败",

	RespStatusTicketNotFindError: "票不存在",
}

func GrpcResponseError(code int64) error {
	if code == RespStatusSuccess {
		return nil
	} else if msg, ok := MsgFlags[code]; ok {
		return status.Error(codes.Aborted, msg)
	}

	return status.Error(codes.Unknown, "未知错误")
}
