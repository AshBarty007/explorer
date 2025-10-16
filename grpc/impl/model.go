package impl

type Tickets struct {
	TokenID  string     `json:"token_id"`
	Owner    string     `json:"owner"`
	Slot     TicketInfo `json:"slot"`
	Balance  string     `json:"balance"`
	Metadata Metadata   `json:"metadata"`
}

type Metadata struct {
	TokenURL    string `json:"token_url"`
	Description string `json:"description"`
}

type TicketInfo struct {
	BasicInformation      BasicInformation      `json:"BasicInformation"`
	AdditionalInformation AdditionalInformation `json:"AdditionalInformation"`
}
type TimeSharing struct {
	TimeSharingID        string `json:"timeSharing_id"`
	TimeSharingBeginTime string `json:"timeSharing_begin_time"`
	TimeSharingEndTime   string `json:"timeSharing_end_time"`
}
type RuleIssue struct {
	RuleIssueName      string `json:"ruleIssue_name"`
	RuleIssueWay       int32  `json:"ruleIssue_way"`
	RuleIssueType      int32  `json:"ruleIssue_type"`
	IsRealName         int32  `json:"is_real_name"`
	UseTime            string `json:"use_time"`
	RuleIssueBeginTime string `json:"ruleIssue_begin_time"`
	RuleIssueEndTime   string `json:"ruleIssue_end_time"`
	RealNameCheck      int32  `json:"real_name_check"`
	OnlyOwnerBuy       int32  `json:"only_owner_buy"`
	RightsCheck        int32  `json:"rights_check"`
	RightsID           string `json:"rights_id"`
	NeedApproval       int32  `json:"need_approval"`
	ApproveID          string `json:"approve_id"`
	ApproveContent     string `json:"approve_content"`
	RuleType           int32  `json:"rule_type"`
	OnlyWindowSale     int32  `json:"only_window_sale"`
}
type RuleCheck struct {
	IdentityType  string   `json:"identity_type"`
	RuleCheckName string   `json:"ruleCheck_name"`
	ControlType   int32    `json:"control_type"`
	AdoptType     int32    `json:"adopt_type"`
	IntervalTime  int32    `json:"interval_time"`
	TimeShareBook int32    `json:"time_share_book"`
	CheckPointIds []string `json:"check_point_ids"`
}
type RuleRetreat struct {
	RuleRetreatName string  `json:"ruleRetreat_name"`
	IsRetreat       int32   `json:"is_retreat"`
	DefaultRate     float64 `json:"default_rate"`
}

type TicketGoods struct {
	TicketGoodsID   string      `json:"ticketGoods_id"`
	GoodsName       string      `json:"goods_name"`
	TimeShareID     string      `json:"time_share_id"`
	OverallDiscount float64     `json:"overall_discount"`
	BeginDiscount   float64     `json:"begin_discount"`
	EndDiscount     float64     `json:"end_discount"`
	TicketGoodsType int32       `json:"ticketGoods_type"`
	PeopleNumber    int32       `json:"people_number"`
	MinPeople       int32       `json:"min_people"`
	MaxPeople       int32       `json:"max_people"`
	RuleIssue       RuleIssue   `json:"RuleIssue"`
	RuleCheck       RuleCheck   `json:"RuleCheck"`
	RuleRetreat     RuleRetreat `json:"RuleRetreat"`
}
type TicketStock struct {
	StockScenicID   string `json:"stock_scenic_id"`
	StockTicketID   string `json:"stock_ticket_id"`
	StockScenicName string `json:"stock_scenic_name"`
	TicketName      string `json:"ticket_name"`
	AccountID       string `json:"account_id"`
	StockOperatorID string `json:"stock_operator_id"`
	TotalStock      int32  `json:"total_stock"`
	TicketType      int32  `json:"ticket_type"`
	Nums            int32  `json:"nums"`
	BatchID         string `json:"batch_id"`
}
type SimpleTicket struct {
	ScenicID      string        `json:"scenic_id"`
	ScenicName    string        `json:"scenic_name"`
	SimpleName    string        `json:"simple_name"`
	MarketPrice   float64       `json:"market_price"`
	ProType       int32         `json:"pro_type"`
	UseType       int32         `json:"use_type"`
	TimeRestrict  int32         `json:"time_restrict"`
	RestrictType  int32         `json:"restrict_type"`
	RestrictWeek  string        `json:"restrict_week"`
	ValidityDay   int32         `json:"validity_day"`
	IsActivate    int32         `json:"is_activate"`
	UseCount      int32         `json:"use_count"`
	AvailableDays int32         `json:"available_days"`
	ParkStatistic int32         `json:"park_statistic"`
	OperatorID    string        `json:"operator_id"`
	TimeSharing   []TimeSharing `json:"timeSharing"`
	TicketGoods   []TicketGoods `json:"ticketGoods"`
	TicketStock   TicketStock   `json:"TicketStock"`
}
type BasicInformation struct {
	SimpleTicket SimpleTicket `json:"SimpleTicket"`
	IsExchange   int32        `json:"is_exchange"`
}
type BuyerInfo struct {
	BuyerInfoIDName string `json:"buyerInfo_id_name"`
	IDNumber        string `json:"id_number"`
}

type TicketDataCreate struct {
	BuyerInfo        []*BuyerInfo `json:"BuyerInfo"`
	Phone            string       `json:"phone"`
	SaleChannel      int32        `json:"sale_channel"`
	OrderID          string       `json:"order_id"`
	OrderGroupID     string       `json:"order_group_id"`
	PlayerNum        int32        `json:"player_num"`
	IssuanceType     int32        `json:"issuance_type"`
	Status           int32        `json:"status"`
	TicketID         string       `json:"ticket_id"`
	PrintEncode      string       `json:"print_encode"`
	EnterBeginTime   string       `json:"enter_begin_time"`
	EnterEndTime     string       `json:"enter_end_time"`
	OverdueTime      string       `json:"overdue_time"`
	ProviderID       string       `json:"provider_id"`
	StoreID          string       `json:"store_id"`
	SellingPrice     float64      `json:"selling_price"`
	CancelCount      int32        `json:"cancel_count"`
	EnterTime        string       `json:"enter_time"`
	CheckedNum       int32        `json:"checked_num"`
	UsedCount        int32        `json:"used_count"`
	UsedDays         int32        `json:"used_days"`
	StockBatchNumber string       `json:"stock_batch_number"`
}
type TicketData struct {
	BuyerInfo        []*BuyerInfo `json:"BuyerInfo"`
	Phone            string       `json:"phone"`
	SaleChannel      int32        `json:"sale_channel"`
	OrderID          string       `json:"order_id"`
	OrderGroupID     string       `json:"order_group_id"`
	PlayerNum        int32        `json:"player_num"`
	IssuanceType     int32        `json:"issuance_type"`
	Status           int32        `json:"status"`
	TicketID         string       `json:"ticket_id"`
	PrintEncode      string       `json:"print_encode"`
	EnterBeginTime   string       `json:"enter_begin_time"`
	EnterEndTime     string       `json:"enter_end_time"`
	OverdueTime      string       `json:"overdue_time"`
	ProviderID       string       `json:"provider_id"`
	UserID           string       `json:"user_id"`
	StoreID          string       `json:"store_id"`
	SellingPrice     float64      `json:"selling_price"`
	CancelCount      int32        `json:"cancel_count"`
	EnterTime        string       `json:"enter_time"`
	CheckedNum       int32        `json:"checked_num"`
	UsedCount        int32        `json:"used_count"`
	UsedDays         int32        `json:"used_days"`
	StockBatchNumber string       `json:"stock_batch_number"`
	Account          string       `json:"account"`
	Org              string       `json:"org"`
}
type PriceDetailedInfo struct {
	PriceID        string    `json:"price_id"`
	SalePrice      float64   `json:"sale_price"`
	ComposePrice   float64   `json:"compose_price"`
	CommissionRate float64   `json:"commission_rate"`
	IsCompose      bool      `json:"is_compose"`
	GroupInfo      GroupInfo `json:"group"`
}

type GroupInfo struct {
	AddGroupID []string `json:"add_group_id"`
	DelGroupID []string `json:"del_group_id"`
	GroupID    []string `json:"group_id"`
}

type PriceInfo struct {
	DistributorID     string            `json:"distributor_id"`
	GoodsID           string            `json:"goods_id"`
	PriceDetailedInfo PriceDetailedInfo `json:"PriceDetailedInfo"`
}
type TicketCheck struct {
	TicketCheckWay    string `json:"ticketCheck_way"`
	PointName         string `json:"point_name"`
	PointID           string `json:"point_id"`
	EquipmentName     string `json:"equipment_name"`
	EquipmentID       string `json:"equipment_id"`
	EquipmentType     string `json:"equipment_type"`
	CheckTime         string `json:"check_time"`
	TicketCheckIDName string `json:"ticketCheck_id_name"`
	CheckType         int32  `json:"check_type"`
}
type AdditionalInformation struct {
	TicketData  TicketDataCreate `json:"TicketData"`
	PriceInfo   []*PriceInfo     `json:"PriceInfo"`
	TicketCheck []*VerifyTicket  `json:"TicketCheckData"`
	StockInfo   StockInfo        `json:"StockInfo"`
}

type StockInfo struct {
	PurchaseBeginTime   string `json:"purchase_begin_time"`
	PurchaseEndTime     string `json:"purchase_end_time"`
	StockEnterBeginTime string `json:"stock_enter_begin_time"`
	StockEnterEndTime   string `json:"stock_enter_end_time"`
}

type BlockTicketIssuanceData struct {
	TicketData []*TicketData `json:"TicketData"`
}

type GenerateTicketNumberInfo struct {
	StockConsumeData         []*StockConsumeData   `json:"tokenIds"`
	GenerateTicketNumberInfo []*GenerateTicketInfo `json:"ticketsData"`
}

type GenerateTicketInfo struct {
	EnterTime     string  `json:"enter_time"`
	PlayerNum     int32   `json:"player_num"`
	Certificate   string  `json:"certificate"`
	Rand          string  `json:"rand"`
	ScenicID      string  `json:"scenic_id"`
	ProType       int32   `json:"pro_type"`
	TimeShareID   int64   `json:"time_share_id"`
	TimeShareBook int32   `json:"time_share_book"`
	BeginTime     string  `json:"begin_time"`
	EndTime       string  `json:"end_time"`
	CheckPointIds []int64 `json:"check_point_ids"`
	UUID          string  `json:"uuid"`
}

type OrderInfo struct {
	OrderGroupID     string      `json:"order_group_id"`
	OrderStatus      string      `json:"order_status"`
	OrderType        string      `json:"order_type"`
	TotalAmount      float64     `json:"total_amount"`
	PayAmount        float64     `json:"pay_amount"`
	PayType          int32       `json:"pay_type"`
	SourceType       int32       `json:"source_type"`
	StockCertificate string      `json:"stock_certificate"`
	TradeNo          string      `json:"trade_no"`
	UserID           string      `json:"user_id"`
	Username         string      `json:"username"`
	PayTime          string      `json:"pay_time"`
	UserPhone        string      `json:"user_phone"`
	OrderTab         []*OrderTab `json:"OrderTab"`
}
type OrderProductTicketData struct {
	ScenicID                 int64                       `json:"scenic_id"`
	ScenicName               string                      `json:"scenic_name"`
	TicketType               int32                       `json:"ticket_type"`
	Day                      string                      `json:"day"`
	TimeShareID              int64                       `json:"time_share_id"`
	TimeShare                string                      `json:"time_share"`
	ParentProductID          int64                       `json:"parent_product_id"`
	CommissionType           int32                       `json:"commission_type"`
	CommissionRate           float64                     `json:"commission_rate"`
	CommissionAmount         float64                     `json:"commission_amount"`
	ActualComAmount          float64                     `json:"actual_com_amount"`
	BdsAccount               string                      `json:"bds_account"`
	BdsOrg                   string                      `json:"bds_org"`
	TicketTypeID             int64                       `json:"ticket_type_id"`
	TicketTypeSubID          int64                       `json:"ticket_type_sub_id"`
	RealQuantity             int32                       `json:"real_quantity"`
	OrderProductTicketRnData []*OrderProductTicketRnData `json:"OrderProductTicketRnData"`
}

type OrderProductTicketRnData struct {
	ID                      int64  `json:"id"`
	OrderProductID          int64  `json:"order_product_id"`
	TicketNumber            string `json:"ticket_number"`
	TicketStatus            int32  `json:"ticket_status"`
	CommissionSettledStatus int32  `json:"commission_settled_status"`
	IsChain                 int32  `json:"is_chain"`
	BillStatus              int32  `json:"bill_status"`
	IssueTicketType         int32  `json:"issue_ticket_type"`
}
type OrderTab struct {
	OrderID                string                    `json:"order_id"`
	OrderType              string                    `json:"order_type"`
	SellerID               int64                     `json:"seller_id"`
	SellerName             string                    `json:"seller_name"`
	TotalAmount            float64                   `json:"total_amount"`
	PayType                int32                     `json:"pay_type"`
	SourceType             int32                     `json:"source_type"`
	OrderStatus            string                    `json:"order_status"`
	TradeNo                string                    `json:"trade_no"`
	MerchantID             string                    `json:"merchant_id"`
	StoreID                int64                     `json:"store_id"`
	AgentID                int64                     `json:"agent_id"`
	AgentName              string                    `json:"agent_name"`
	CommissionSettledType  int32                     `json:"commission_settled_type"`
	UserID                 string                    `json:"user_id"`
	Username               string                    `json:"username"`
	PayTime                string                    `json:"pay_time"`
	ModifyTime             string                    `json:"modify_time"`
	PayPeople              int64                     `json:"pay_people"`
	Nickname               string                    `json:"nickname"`
	MerchantNo             string                    `json:"merchant_no"`
	OrderProductTicketData []*OrderProductTicketData `json:"OrderProductTicketData"`
}

type VerifyTicket struct {
	VerifyInfo   VerifyInfo   `json:"VerifyInfo"`
	TicketStatus TicketStatus `json:"VerifyStatus"`
}

type VerifyInfo struct {
	Account          string `json:"account"`
	Org              string `json:"org"`
	CheckType        int32  `json:"check_type"`
	TicketNumber     string `json:"ticket_number"`
	StockBatchNumber string `json:"stock_batch_number"`
	EnterTime        string `json:"enter_time"`
	CheckNumber      int32  `json:"check_number"`
	ScenicID         string `json:"scenic_id"`
	IDName           string `json:"id_name"`
	IDCard           string `json:"id_card"`
	QrCode           string `json:"qr_code"`
	PointName        string `json:"point_name"`
	PointID          string `json:"point_id"`
	EquipmentName    string `json:"equipment_name"`
	EquipmentID      string `json:"equipment_id"`
	EquipmentType    string `json:"equipment_type"`
	UserID           string `json:"user_id"`
	Username         string `json:"username"`
}

type TicketsStatus struct {
	TicketStatusData []*TicketStatus
}

type TicketStatus struct {
	// 门票状态
	Status int32 `json:"status"`
	// 票号
	TicketID string `json:"ticket_id"`
	// 已检票人数
	CheckedNum int32 `json:"checked_num"`
	// 已使用次数
	UsedCount int32 `json:"used_count"`
	// 已入园天数
	UsedDays int32 `json:"used_days"`
}

// ------- 迭代2---------

type DistributionInfo struct {
	DistributionOrderInfo *DistributionOrderInfo
	TansferDatas          []*TansferDataTob
}

type AccountData struct {
	From         string `json:"from"`
	FromOrgMspID string `json:"from_org_msp_id"`
	To           string `json:"to"`
	ToOrgMspID   string `json:"to_org_msp_id"`
}

type TansferData struct {
	SenderStockID  string `json:"sender_stock_id"`
	Sender         string `json:"sender"`
	Receive        string `json:"receive"`
	ReceiveStockID string `json:"receive_stock_id"`
	Amount         string `json:"amount"`
}

type TansferDataTob struct {
	SenderStockID     string `json:"sender_stock_id"`
	Sender            string `json:"sender"`
	Receive           string `json:"receive"`
	ReceiveStockID    string `json:"receive_stock_id"`
	Amount            string `json:"amount"`
	AvailableRatio    string `json:"available_ratio"`
	AvailableTotalNum string `json:"available_total_num"`
	StockBatchNumber  string `json:"stock_batch_number"`
}

type DistributionOrderInfo struct {
	OrderGroupID           string                    `json:"order_group_id"`
	OrderStatus            string                    `json:"order_status"`
	OrderType              string                    `json:"order_type"`
	TotalAmount            float64                   `json:"total_amount"`
	PayType                int32                     `json:"pay_type"`
	SourceType             int32                     `json:"source_type"`
	StockCertificate       string                    `json:"stock_certificate"`
	TradeNo                string                    `json:"trade_no"`
	UserID                 string                    `json:"user_id"`
	Username               string                    `json:"username"`
	PayTime                string                    `json:"pay_time"`
	CertID                 string                    `json:"cert_id"`
	UserPhone              string                    `json:"user_phone"`
	OrderTabToBData        []*OrderTabToBData        `json:"orderTabToBData"`
	OrderTabDistributeData []*OrderTabDistributeData `json:"orderTabDistributeData"`
}
type OrderTabToBData struct {
	OrderID               string  `json:"order_id"`
	OrderType             string  `json:"order_type"`
	SellerID              int64   `json:"seller_id"`
	SellerName            string  `json:"seller_name"`
	TotalAmount           float64 `json:"total_amount"`
	PayAmount             float64 `json:"pay_amount"`
	PayType               int32   `json:"pay_type"`
	PayTime               string  `json:"pay_time"`
	SourceType            int32   `json:"source_type"`
	OrderStatus           string  `json:"order_status"`
	TradeNo               string  `json:"trade_no"`
	MerchantID            string  `json:"merchant_id"`
	StoreID               int64   `json:"store_id"`
	AgentID               int64   `json:"agent_id"`
	AgentName             string  `json:"agent_name"`
	CommissionSettledType int32   `json:"commission_settled_type"`
	UserID                string  `json:"user_id"`
	Username              string  `json:"username"`
	Nickname              string  `json:"nickname"`
	MerchantNo            string  `json:"merchant_no"`
}
type OrderProductDistributeData struct {
	ScenicID                 int64   `json:"scenic_id"`
	ScenicName               string  `json:"scenic_name"`
	DistributorTicketStockID int64   `json:"distributor_ticket_stock_id"`
	BatchID                  string  `json:"batch_id"`
	TicketType               int32   `json:"ticket_type"`
	DayBegin                 string  `json:"day_begin"`
	DayEnd                   string  `json:"day_end"`
	TimeShare                string  `json:"time_share"`
	UsableNum                int32   `json:"usable_num"`
	OrderProductID           int64   `json:"order_product_id"`
	OrderID                  string  `json:"order_id"`
	ProductID                int64   `json:"product_id"`
	ProductName              string  `json:"product_name"`
	ProductSkuID             int64   `json:"product_sku_id"`
	ProductSkuName           string  `json:"product_sku_name"`
	ProductPrice             float64 `json:"product_price"`
	Num                      int32   `json:"num"`
	ProductType              int32   `json:"product_type"`
	AvailableRatio           string  `json:"available_ratio"`
	AvailableTotalNum        string  `json:"available_total_num"`
	ExchangeFreezeNum        string  `json:"exchange_freeze_num"`
}
type OrderTabDistributeData struct {
	OrderID                    string                        `json:"order_id"`
	BuyerID                    int64                         `json:"buyer_id"`
	BuyerName                  string                        `json:"buyer_name"`
	SellerID                   int64                         `json:"seller_id"`
	SellerName                 string                        `json:"seller_name"`
	ServiceProviderID          int64                         `json:"service_provider_id"`
	ServiceProviderName        string                        `json:"service_provider_name"`
	OrderProductDistributeData []*OrderProductDistributeData `json:"OrderProductDistributeData"`
}

type DistributeReturnInfo struct {
	DistributeRefundInfo *DistributeRefundInfo
	TansferDatas         []*TansferData
}

type DistributeRefundInfo struct {
	OrderRefundGroup             []*OrderRefundGroup             `json:"orderRefundGroup"`
	OrderRefund                  []*OrderRefund                  `json:"orderRefund"`
	OrderRefundProductDistribute []*OrderRefundProductDistribute `json:"orderRefundProductDistribute"`
}
type OrderRefundGroup struct {
	OrderRefundGroupID string `json:"order_refund_group_id"`
	OrderGroupID       string `json:"order_group_id"`
	OrderRefundID      string `json:"order_refund_id"`
	CreateTime         string `json:"create_time"`
}
type OrderRefund struct {
	RefundID                string  `json:"refund_id"`
	OrderID                 string  `json:"order_id"`
	RefundAmount            float64 `json:"refund_amount"`
	RefundFee               float64 `json:"refund_fee"`
	RefundStatus            string  `json:"refund_status"`
	RefundType              string  `json:"refund_type"`
	TradeNo                 string  `json:"trade_no"`
	RefundTime              string  `json:"refund_time"`
	CreateTime              string  `json:"create_time"`
	Remark                  string  `json:"remark"`
	FailMessage             string  `json:"fail_message"`
	UserID                  string  `json:"user_id"`
	Username                string  `json:"username"`
	CommissionSettledStatus string  `json:"commission_settled_status"`
	StockCertificate        string  `json:"stock_certificate"`
	ProductSkuName          string  `json:"product_sku_name"`
}
type OrderRefundProductDistribute struct {
	RefundID                string `json:"refund_id"`
	OrderProductID          string `json:"order_product_id"`
	Num                     int32  `json:"num"`
	ProductID               string `json:"product_id"`
	ProductName             string `json:"product_name"`
	ProductSkuID            string `json:"product_sku_id"`
	ProductType             string `json:"product_type"`
	ProductPrice            string `json:"product_price"`
	DayBegin                string `json:"day_begin"`
	DayEnd                  string `json:"day_end"`
	TimeShareID             string `json:"time_share_id"`
	TimeShare               string `json:"time_share"`
	ScenicID                string `json:"scenic_id"`
	ScenicName              string `json:"scenic_name"`
	BatchID                 string `json:"batch_id"`
	DistributeTicketStockID string `json:"distribute_ticket_stock_id"`
}

type OrderRefundInfoToC struct {
	RefundInfoToC          *RefundInfoToC            `json:"refundInfoToC"`
	RefundProductTicketToC []*RefundProductTicketToC `json:"refundProductTicketToC"`
}
type RefundInfoToC struct {
	RefundID                string  `json:"refund_id"`
	OrderID                 string  `json:"order_id"`
	RefundAmount            float64 `json:"refund_amount"`
	RefundFee               float64 `json:"refund_fee"`
	RefundStatus            string  `json:"refund_status"`
	RefundType              int32   `json:"refund_type"`
	TradeNo                 string  `json:"trade_no"`
	RefundTime              string  `json:"refund_time"`
	Remark                  string  `json:"remark"`
	FailMessage             string  `json:"fail_message"`
	CommissionSettledStatus int32   `json:"commission_settled_status"`
	StockCertificate        string  `json:"stock_certificate"`
	ProductSkuName          string  `json:"product_sku_name"`
	UserID                  int64   `json:"user_id"`
	Username                string  `json:"username"`
	BillStatus              int32   `json:"bill_status"`
	OrderGroupID            string  `json:"order_group_id"`
}
type RefundProductTicketToC struct {
	RefundID       string              `json:"refund_id"`
	OrderProductID int64               `json:"order_product_id"`
	TicketNumber   string              `json:"ticket_number"`
	ProductID      int64               `json:"product_id"`
	ProductName    string              `json:"product_name"`
	ProductSkuID   int64               `json:"product_sku_id"`
	ProductType    int32               `json:"product_type"`
	TicketType     int32               `json:"ticket_type"`
	Day            string              `json:"day"`
	Name           string              `json:"name"`
	Identity       string              `json:"identity"`
	SourceType     int32               `json:"source_type"`
	RefundAmount   string              `json:"refund_amount"`
	RefundFee      string              `json:"refund_fee"`
	RefundNum      int32               `json:"refund_num"`
	StockBatchInfo []*StockConsumeData `json:"stock_batch_info"`
}

type StockConsumeData struct {
	StockBatchNumber string `json:"stock_batch_number"`
	Sender           string `json:"sender"`
	Amount           uint32 `json:"amount"`
}

type TimerTicketStatus struct {
	// 门票状态
	Status int32 `json:"status"`
	// 票号
	TicketID string `json:"ticket_id"`
}

type ActiveInfo struct {
	OrderID           string `json:"order_id"`
	BatchID           string `json:"batch_id"`
	Periods           string `json:"periods"`
	TotalPeriods      string `json:"total_periods"`
	TokenID           string `json:"token_id"`
	TradeNo           string `json:"trade_no"`
	Amount            string `json:"amount"`
	TotalRepayment    string `json:"total_repayment"`
	AvailableTotalNum string `json:"available_total_num"`
}

type UpdateStockInfo struct {
	StockInfo StockInfo `json:"StockInfo"`
}

type ProjectBidding struct {
	ProjectID            string `json:"project_id"`
	ParentProjectID      string `json:"parent_project_id"`
	ProjectType          string `json:"project_type"`
	ProjectNumber        string `json:"project_number"`
	ExchangeID           string `json:"exchange_id"`
	CompanyID            string `json:"company_id"`
	ProjectName          string `json:"project_name"`
	ProjectGrade         string `json:"project_grade"`
	ProjectDescription   string `json:"project_description"`
	ProjectAnnex         string `json:"project_annex"`
	ProjectStatus        string `json:"project_status"`
	NodeDescription      string `json:"node_description"`
	ProjectChainUniqueId string `json:"project_chain_unique_id"`
	ReviewTime           string `json:"review_time"`
}

// type InstrumentBiddingRule struct {
// 	InstrumentID       string   `json:"instrument_id"`
// 	InstrumentGrade    int32   `json:"instrument_grade"`
// 	BiddingStartTime   string  `json:"bidding_start_time"`
// 	BiddingEndTime     string  `json:"bidding_end_time"`
// 	BiddingDelayTime   int64   `json:"bidding_delay_time"`
// 	BiddingType        string  `json:"bidding_type"`
// 	BiddingStartPrice  float64 `json:"bidding_start_price"`
// 	BiddingDealRule    int32   `json:"bidding_deal_rule"`
// 	BiddingDeposit     float64 `json:"bidding_deposit"`
// 	BiddingChanges     float64 `json:"bidding_changes"`
// 	InstrumentQuantity int64   `json:"instrument_quantity"`
// 	BiddingDescription string  `json:"bidding_description"`
// }

// type InstrumentBiddingSales struct {
// 	InstrumentID       string  `json:"instrument_id"`
// 	InstrumentViews          int64  `json:"instrument_views"`
// 	RegistrationQuantity     int64  `json:"registration_quantity"`
// 	BidQuantity              int64  `json:"bid_quantity"`
// 	CurrentPrice             int64  `json:"current_price"`
// 	CurrentRemainingQuantity int64  `json:"current_remaining_quantity"`
// 	EstimatedStartTime       string `json:"estimated_start_time"`
// 	EstimatedEndTime         string `json:"estimated_end_time"`
// }

// type InstrumentTicket struct {
// 	InstrumentID       string  `json:"instrument_id"`
// 	ScenicID           int64  `json:"scenic_id"`
// 	GoodsID            int64  `json:"goods_id"`
// 	ProductType        string `json:"product_type"`
// 	TicketType         int32  `json:"ticket_type"`
// 	ScenicProvinceCode int32  `json:"scenic_province_code"`
// 	ScenicCityCode     int32  `json:"scenic_city_code"`
// 	ScenicAreaCode     int32  `json:"scenic_area_code"`
// 	BuyStartDate       string `json:"buy_start_date"`
// 	BuyEndDate         string `json:"buy_end_date"`
// 	UseStartDate       string `json:"use_start_date"`
// 	UseEndDate         string `json:"use_end_date"`
// }

type InstrumentData struct {
	InstrumentID         string          `json:"instrument_id"`
	ScenicID             string          `json:"scenic_id"`
	GoodsID              string          `json:"goods_id"`
	ProductType          string          `json:"product_type"`
	TicketType           string          `json:"ticket_type"`
	ScenicProvinceCode   string          `json:"scenic_province_code"`
	ScenicCityCode       string          `json:"scenic_city_code"`
	ScenicAreaCode       string          `json:"scenic_area_code"`
	BuyStartDate         string          `json:"buy_start_date"`
	BuyEndDate           string          `json:"buy_end_date"`
	UseStartDate         string          `json:"use_start_date"`
	UseEndDate           string          `json:"use_end_date"`
	ProjectChainUniqueId string          `json:"project_chain_unique_id"`
	Instrument           *InstrumentInfo `json:"instrument"`
}

type InstrumentBiddingRule struct {
	InstrumentID        string `json:"instrument_id"`
	InstrumentGrade     string `json:"instrument_grade"`
	BiddingStartTime    string `json:"bidding_start_time"`
	BiddingEndTime      string `json:"bidding_end_time"`
	BiddingDelayTime    string `json:"bidding_delay_time"`
	BiddingType         string `json:"bidding_type"`
	BiddingStartPrice   string `json:"bidding_start_price"`
	BiddingDealRule     string `json:"bidding_deal_rule"`
	BiddingDeposit      string `json:"bidding_deposit"`
	BiddingChanges      string `json:"bidding_changes"`
	InstrumentQuantity  string `json:"instrument_quantity"`
	BiddingDescription  string `json:"bidding_description"`
	MarketPrice         string `json:"market_price"`
	BiddingDepositRatio string `json:"bidding_deposit_ratio"`
	LowestPrice         string `json:"lowest_price"`
}
type InstrumentBiddingSales struct {
	InstrumentID             string `json:"instrument_id"`
	InstrumentViews          string `json:"instrument_views"`
	RegistrationQuantity     string `json:"registration_quantity"`
	BidQuantity              string `json:"bid_quantity"`
	CurrentPrice             string `json:"current_price"`
	CurrentRemainingQuantity string `json:"current_remaining_quantity"`
	EstimatedStartTime       string `json:"estimated_start_time"`
	EstimatedEndTime         string `json:"estimated_end_time"`
}
type InstrumentInfo struct {
	InstrumentID           string                  `json:"instrument_id"`
	ProjectID              string                  `json:"project_id"`
	ParentInstrumentID     string                  `json:"parent_instrument_id"`
	CompanyID              string                  `json:"company_id"`
	AssetType              string                  `json:"asset_type"`
	InstrumentType         string                  `json:"instrument_type"`
	InstrumentName         string                  `json:"instrument_name"`
	InstrumentImages       string                  `json:"instrument_images"`
	InstrumentVideo        string                  `json:"instrument_video"`
	InstrumentDescription  string                  `json:"instrument_description"`
	InstrumentAnnex        string                  `json:"instrument_annex"`
	ContactName            string                  `json:"contact_name"`
	ContactPhone           string                  `json:"contact_phone"`
	InstrumentStatus       string                  `json:"instrument_status"`
	InstrumentBiddingRule  *InstrumentBiddingRule  `json:"instrument_bidding_rule"`
	InstrumentBiddingSales *InstrumentBiddingSales `json:"instrument_bidding_sales"`
}

type StoreMarginOrder struct {
	ProjectID   string      `json:"projectId"`
	ProjectName string      `json:"projectName"`
	Exchange    *Exchange   `json:"exchange"`
	Instrument  *Instrument `json:"Instrument"`
}
type Exchange struct {
	ExchangeID   string `json:"exchangeId"`
	ExchangeName string `json:"exchangeName"`
}
type Bidbond struct {
	BidbondAmount string `json:"bidbondAmount"`
	OutTradeNo    string `json:"outTradeNo"`
	TradeNo       string `json:"tradeNo"`
	CreateTime    string `json:"createTime"`
	OrderID       string `json:"orderId"`
}
type Instrument struct {
	InstrumentID   string   `json:"instrumentId"`
	InstrumentName string   `json:"instrumentName"`
	SellerID       string   `json:"sellerId"`
	SellerName     string   `json:"sellerName"`
	BuyerID        string   `json:"buyerId"`
	BuyerName      string   `json:"buyerName"`
	Bidbond        *Bidbond `json:"bidbond"`
}

type InstrumentPayMent struct {
	InstrumentID   string          `json:"instrumentId"`
	InstrumentName string          `json:"instrumentName"`
	SellerID       string          `json:"sellerId"`
	SellerName     string          `json:"sellerName"`
	BuyerID        string          `json:"buyerId"`
	BuyerName      string          `json:"buyerName"`
	Bidbond        *BidbondPayment `json:"bidbond"`
}

type BidbondPayment struct {
	BidbondAmount  string `json:"bidbondAmount"`
	OutTradeNo     string `json:"outTradeNo"`
	TradeNo        string `json:"tradeNo"`
	BidNumber      string `json:"bidNumber"`
	PaymentTradeNo string `json:"paymentTradeNo"`
	PaymentTime    string `json:"paymentTime"`
}

type StoreMarginPayment struct {
	ProjectID   string             `json:"projectId"`
	ProjectName string             `json:"projectName"`
	Exchange    *Exchange          `json:"exchange"`
	Instrument  *InstrumentPayMent `json:"Instrument"`
}

type CreditRating struct {
	LegalName          string `json:"legalName"`
	RegistrationNumber string `json:"registrationNumber"`
	Score              string `json:"score"`
	Level              string `json:"level"`
	AssessmentTime     string `json:"assessmentTime"`
}

type LoanApplication struct {
	ApplicationNo              string `json:"applicationNo"`
	LegalName                  string `json:"legalName"`
	DebtorRegistrationNumber   string `json:"debtorRegistrationNumber"`
	OwnersEquity               string `json:"ownersEquity"`
	BusinessScope              string `json:"businessScope"`
	BusinessEconomicType       string `json:"businessEconomicType"`
	BusinessPhone              string `json:"businessPhone"`
	BusinessFax                string `json:"businessFax"`
	BusinessEmail              string `json:"businessEmail"`
	CreditorName               string `json:"creditorName"`
	CreditorRegistrationNumber string `json:"creditorRegistrationNumber"`
	FinancingPurpose           string `json:"financingPurpose"`
	Level                      string `json:"level"`
	ProposedCreditLimit        string `json:"proposedCreditLimit"`
	CreateTime                 string `json:"createTime"`
}

type LoanApproval struct {
	ApplicationNo              string `json:"applicationNo"`
	LegalName                  string `json:"legalName"`
	DebtorRegistrationNumber   string `json:"debtorRegistrationNumber"`
	OwnersEquity               string `json:"ownersEquity"`
	BusinessScope              string `json:"businessScope"`
	BusinessEconomicType       string `json:"businessEconomicType"`
	BusinessPhone              string `json:"businessPhone"`
	BusinessFax                string `json:"businessFax"`
	BusinessEmail              string `json:"businessEmail"`
	CreditorName               string `json:"creditorName"`
	CreditorRegistrationNumber string `json:"creditorRegistrationNumber"`
	FinancingPurpose           string `json:"financingPurpose"`
	Level                      string `json:"level"`
	ProposedCreditLimit        string `json:"proposedCreditLimit"`
	CreateTime                 string `json:"createTime"`
	CreditStatus               string `json:"creditStatus"`
	CreditLine                 string `json:"creditLine"`
	ExpiresTime                string `json:"expiresTime"`
	RepaymentType              string `json:"repaymentType"`
	Rate                       string `json:"rate"`
	Periods                    string `json:"periods"`
	MonthlyRepaymentDay        string `json:"monthlyRepaymentDay"`
	OverdueRate                string `json:"overdueRate"`
	RepaymentName              string `json:"repaymentName"`
	RepaymentBankAccount       string `json:"repaymentBankAccount"`
	RepaymentBankName          string `json:"repaymentBankName"`
	ModifyTime                 string `json:"modifyTime"`
}

type LoanApprovalFailed struct {
	ApplicationNo              string `json:"applicationNo"`
	LegalName                  string `json:"legalName"`
	DebtorRegistrationNumber   string `json:"debtorRegistrationNumber"`
	OwnersEquity               string `json:"ownersEquity"`
	BusinessScope              string `json:"businessScope"`
	BusinessEconomicType       string `json:"businessEconomicType"`
	BusinessPhone              string `json:"businessPhone"`
	BusinessFax                string `json:"businessFax"`
	BusinessEmail              string `json:"businessEmail"`
	CreditorName               string `json:"creditorName"`
	CreditorRegistrationNumber string `json:"creditorRegistrationNumber"`
	FinancingPurpose           string `json:"financingPurpose"`
	Level                      string `json:"level"`
	ProposedCreditLimit        string `json:"proposedCreditLimit"`
	CreateTime                 string `json:"createTime"`
	CreditStatus               string `json:"creditStatus"`
	CreditAuditRemark          string `json:"creditAuditRemark"`
	ModifyTime                 string `json:"modifyTime"`
}

type BorrowApplication struct {
	OutTradeNo                 string `json:"outTradeNo"`
	LoanAmount                 string `json:"loanAmount"`
	DebtorName                 string `json:"debtorName"`
	DebtorRegistrationNumber   string `json:"debtorRegistrationNumber"`
	CreditorName               string `json:"creditorName"`
	CreditorRegistrationNumber string `json:"creditorRegistrationNumber"`
	ExchangeName               string `json:"exchangeName"`
	ExchangeRegistrationNumber string `json:"exchangeRegistrationNumber"`
	ReceiverCorporateName      string `json:"receiverCorporateName"`
	ReceiverBankAccountNo      string `json:"receiverBankAccountNo"`
	ReceiverBankName           string `json:"receiverBankName"`
	TradeOrderFileURL          string `json:"tradeOrderFileUrl"`
	CautionMoneyFileURL        string `json:"cautionMoneyFileUrl"`
	PaymentBookFileURL         string `json:"paymentBookFileUrl"`
	ApplyTime                  string `json:"applyTime"`
}
type BorrowApproval struct {
	OutTradeNo                 string `json:"outTradeNo"`
	LoanAmount                 string `json:"loanAmount"`
	DebtorName                 string `json:"debtorName"`
	DebtorRegistrationNumber   string `json:"debtorRegistrationNumber"`
	CreditorName               string `json:"creditorName"`
	CreditorRegistrationNumber string `json:"creditorRegistrationNumber"`
	ExchangeName               string `json:"exchangeName"`
	ExchangeRegistrationNumber string `json:"exchangeRegistrationNumber"`
	ReceiverCorporateName      string `json:"receiverCorporateName"`
	ReceiverBankAccountNo      string `json:"receiverBankAccountNo"`
	ReceiverBankName           string `json:"receiverBankName"`
	TradeOrderFileURL          string `json:"tradeOrderFileUrl"`
	CautionMoneyFileURL        string `json:"cautionMoneyFileUrl"`
	PaymentBookFileURL         string `json:"paymentBookFileUrl"`
	PayBankSlipFileURL         string `json:"payBankSlipFileUrl"`
	PayBankTradeNo             string `json:"payBankTradeNo"`
	PayCorporateName           string `json:"payCorporateName"`
	PayBankAccountNo           string `json:"payBankAccountNo"`
	PayBankName                string `json:"payBankName"`
	PayAuditTime               string `json:"payAuditTime"`
}

type RepaymentOrder struct {
	RepaymentNo                string `json:"repaymentNo"`
	DebtorName                 string `json:"debtorName"`
	DebtorRegistrationNumber   string `json:"debtorRegistrationNumber"`
	CreditorName               string `json:"creditorName"`
	CreditorRegistrationNumber string `json:"creditorRegistrationNumber"`
	TotalPrincipal             string `json:"totalPrincipal"`
	RepaymentType              string `json:"repaymentType"`
	Rate                       string `json:"rate"`
	CreateTime                 string `json:"createTime"`
}

type RepaymentApplication struct {
	RepaymentNo                string `json:"repaymentNo"`
	DebtorName                 string `json:"debtorName"`
	DebtorRegistrationNumber   string `json:"debtorRegistrationNumber"`
	CreditorName               string `json:"creditorName"`
	CreditorRegistrationNumber string `json:"creditorRegistrationNumber"`
	TotalPrincipal             string `json:"totalPrincipal"`
	Principal                  string `json:"principal"`
	RepaymentType              string `json:"repaymentType"`
	Rate                       string `json:"rate"`
	Periods                    string `json:"periods"`
	RepaymentDate              string `json:"repaymentDate"`
	IsOverdue                  string `json:"isOverdue"`
	PenaltyInterest            string `json:"penaltyInterest"`
	Amount                     string `json:"amount"`
	BankSlip                   string `json:"bankSlip"`
	PayBankTradeNo             string `json:"payBankTradeNo"`
	PayUnitName                string `json:"payUnitName"`
	PayBankAccountNo           string `json:"payBankAccountNo"`
	PayBankName                string `json:"payBankName"`
	RepaymentName              string `json:"repaymentName"`
	RepaymentBankAccount       string `json:"repaymentBankAccount"`
	RepaymentBankName          string `json:"repaymentBankName"`
	CreateTime                 string `json:"createTime"`
}

type StoreRepaymentApproval struct {
	RepaymentNo                string `json:"repaymentNo"`
	DebtorName                 string `json:"debtorName"`
	DebtorRegistrationNumber   string `json:"debtorRegistrationNumber"`
	CreditorName               string `json:"creditorName"`
	CreditorRegistrationNumber string `json:"creditorRegistrationNumber"`
	TotalPrincipal             string `json:"totalPrincipal"`
	Principal                  string `json:"principal"`
	RepaymentType              string `json:"repaymentType"`
	Rate                       string `json:"rate"`
	Periods                    string `json:"periods"`
	RepaymentDate              string `json:"repaymentDate"`
	IsOverdue                  string `json:"isOverdue"`
	PenaltyInterest            string `json:"penaltyInterest"`
	Amount                     string `json:"amount"`
	BankSlip                   string `json:"bankSlip"`
	PayBankTradeNo             string `json:"payBankTradeNo"`
	PayUnitName                string `json:"payUnitName"`
	PayBankAccountNo           string `json:"payBankAccountNo"`
	PayBankName                string `json:"payBankName"`
	RepaymentName              string `json:"repaymentName"`
	RepaymentBankAccount       string `json:"repaymentBankAccount"`
	RepaymentBankName          string `json:"repaymentBankName"`
	RepaymentStatus            string `json:"repaymentStatus"`
	ModifyTime                 string `json:"modifyTime"`
}

type StoreRepaymentApprovalFailed struct {
	RepaymentNo                string `json:"repaymentNo"`
	DebtorName                 string `json:"debtorName"`
	DebtorRegistrationNumber   string `json:"debtorRegistrationNumber"`
	CreditorName               string `json:"creditorName"`
	CreditorRegistrationNumber string `json:"creditorRegistrationNumber"`
	TotalPrincipal             string `json:"totalPrincipal"`
	Principal                  string `json:"principal"`
	RepaymentType              string `json:"repaymentType"`
	Rate                       string `json:"rate"`
	Periods                    string `json:"periods"`
	RepaymentDate              string `json:"repaymentDate"`
	IsOverdue                  string `json:"isOverdue"`
	PenaltyInterest            string `json:"penaltyInterest"`
	Amount                     string `json:"amount"`
	BankSlip                   string `json:"bankSlip"`
	PayBankTradeNo             string `json:"payBankTradeNo"`
	PayUnitName                string `json:"payUnitName"`
	PayBankAccountNo           string `json:"payBankAccountNo"`
	PayBankName                string `json:"payBankName"`
	RepaymentName              string `json:"repaymentName"`
	RepaymentBankAccount       string `json:"repaymentBankAccount"`
	RepaymentBankName          string `json:"repaymentBankName"`
	RepaymentStatus            string `json:"repaymentStatus"`
	Remark                     string `json:"remark"`
	ModifyTime                 string `json:"modifyTime"`
}

type RepaymentReApplication struct {
	RepaymentNo                string `json:"repaymentNo"`
	DebtorName                 string `json:"debtorName"`
	DebtorRegistrationNumber   string `json:"debtorRegistrationNumber"`
	CreditorName               string `json:"creditorName"`
	CreditorRegistrationNumber string `json:"creditorRegistrationNumber"`
	TotalPrincipal             string `json:"totalPrincipal"`
	Principal                  string `json:"principal"`
	RepaymentType              string `json:"repaymentType"`
	Rate                       string `json:"rate"`
	Periods                    string `json:"periods"`
	RepaymentDate              string `json:"repaymentDate"`
	IsOverdue                  string `json:"isOverdue"`
	PenaltyInterest            string `json:"penaltyInterest"`
	Amount                     string `json:"amount"`
	BankSlip                   string `json:"bankSlip"`
	PayBankTradeNo             string `json:"payBankTradeNo"`
	PayUnitName                string `json:"payUnitName"`
	PayBankAccountNo           string `json:"payBankAccountNo"`
	PayBankName                string `json:"payBankName"`
	RepaymentName              string `json:"repaymentName"`
	RepaymentBankAccount       string `json:"repaymentBankAccount"`
	RepaymentBankName          string `json:"repaymentBankName"`
	CreateTime                 string `json:"createTime"`
}

type LoanReApplication struct {
	ApplicationNo              string `json:"applicationNo"`
	LegalName                  string `json:"legalName"`
	DebtorRegistrationNumber   string `json:"debtorRegistrationNumber"`
	OwnersEquity               string `json:"ownersEquity"`
	BusinessScope              string `json:"businessScope"`
	BusinessEconomicType       string `json:"businessEconomicType"`
	BusinessPhone              string `json:"businessPhone"`
	BusinessFax                string `json:"businessFax"`
	BusinessEmail              string `json:"businessEmail"`
	CreditorName               string `json:"creditorName"`
	CreditorRegistrationNumber string `json:"creditorRegistrationNumber"`
	FinancingPurpose           string `json:"financingPurpose"`
	Level                      string `json:"level"`
	ProposedCreditLimit        string `json:"proposedCreditLimit"`
	CreateTime                 string `json:"createTime"`
}

type InstrumentOrder struct {
	TradeID        string `json:"trade_id"`
	ExchangeID     string `json:"exchange_id"`
	SellerID       string `json:"seller_id"`
	BuyerID        string `json:"buyer_id"`
	TradeAmount    string `json:"trade_amount"`
	TradeType      string `json:"trade_type"`
	TradeStatus    string `json:"trade_status"`
	TradeDatetime  string `json:"trade_datetime"`
	CreateTime     string `json:"create_time"`
	UpdateTime     string `json:"update_time"`
	CompanyID      string `json:"company_id"`
	CompanyName    string `json:"company_name"`
	ProjectID      string `json:"project_id"`
	ProjectName    string `json:"project_name"`
	ExchangeName   string `json:"exchange_name"`
	InstrumentID   string `json:"instrument_id"`
	InstrumentName string `json:"instrument_name"`
	SettlementType string `json:"settlement_type"`
	SettlementId   string `json:"settlement_id"`
}

type ConvertToInvoice struct {
	ProjectID         string `json:"project_id"`         // 公告ID
	ProjectName       string `json:"project_name"`       // 公告名称
	ExchangeID        string `json:"exchange_id"`        // 交易所ID
	ExchangeName      string `json:"exchange_name"`      // 交易所名称
	SellerID          string `json:"seller_id"`          // 委托方ID
	SellerName        string `json:"seller_name"`        // 委托方名称
	InstrumentID      string `json:"instrument_id"`      // 标的物ID
	InstrumentName    string `json:"instrument_name"`    // 标的物名称
	BuyerID           string `json:"buyer_id"`           // 竞买人ID
	BuyerName         string `json:"buyer_name"`         // 竞买人名称
	BidbondAmount     string `json:"bidbond_amount"`     // 保证金金额
	TransactionAmount string `json:"transaction_amount"` // 成交总金额
	ConvertedAmount   string `json:"converted_amount"`   // 保证金转为价款金额
	Status            string `json:"status"`             // 状态：已转为价款
	UpdateTime        string `json:"update_time"`        // 更新时间
}

type RefundBalance struct {
	ProjectID         string `json:"project_id"`
	ProjectName       string `json:"project_name"`
	ExchangeID        string `json:"exchange_id"`
	ExchangeName      string `json:"exchange_name"`
	SellerID          string `json:"seller_id"`
	SellerName        string `json:"seller_name"`
	InstrumentID      string `json:"instrument_id"`
	InstrumentName    string `json:"instrument_name"`
	BuyerID           string `json:"buyer_id"`
	BuyerName         string `json:"buyer_name"`
	BidbondAmount     string `json:"bidbond_amount"`
	TransactionAmount string `json:"transaction_amount"`
	ConvertedAmount   string `json:"converted_amount"`
	RefundAmount      string `json:"refund_amount"`
	Status            string `json:"status"`
	UpdateTime        string `json:"update_time"`
}

type FullRefund struct {
	ProjectID      string `json:"project_id"`
	ProjectName    string `json:"project_name"`
	ExchangeID     string `json:"exchange_id"`
	ExchangeName   string `json:"exchange_name"`
	SellerID       string `json:"seller_id"`
	SellerName     string `json:"seller_name"`
	InstrumentID   string `json:"instrument_id"`
	InstrumentName string `json:"instrument_name"`
	BuyerID        string `json:"buyer_id"`
	BuyerName      string `json:"buyer_name"`
	BidbondAmount  string `json:"bidbond_amount"`
	RefundAmount   string `json:"refund_amount"`
	Status         string `json:"status"`
	UpdateTime     string `json:"update_time"`
}

type TradeCharge struct {
	TradeChargeID      string `json:"trade_charge_id"`
	TradeID            string `json:"trade_id"`
	ChargeObject       string `json:"charge_object"`
	CompanyName        string `json:"company_name"`
	ChargeAmount       string `json:"charge_amount"`
	PaymentStatus      string `json:"payment_status"`
	CreateTime         string `json:"create_time"`
	ExpiryDatetime     string `json:"expiry_datetime"`
	PayeeAccountName   string `json:"payee_account_name"`
	PayeeAccountNumber string `json:"payee_account_number"`
	PayeeBankName      string `json:"payee_bank_name"`
	TransactionFlowID  string `json:"transaction_flow_id"`
	PayerAccountName   string `json:"payer_account_name"`
	PayerAccountNumber string `json:"payer_account_number"`
	PayerBankName      string `json:"payer_bank_name"`
	FileURL            string `json:"file_url"`
	SubmitTime         string `json:"submit_time"`
	AuditResult        string `json:"audit_result"`
	Remark             string `json:"remark"`
	AuditTime          string `json:"audit_time"`
}
