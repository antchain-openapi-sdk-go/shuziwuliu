// This file is auto-generated, don't edit it. Thanks.
package client

import (
	rpcutil "github.com/alibabacloud-go/tea-rpc-utils/service"
	util "github.com/alibabacloud-go/tea-utils/service"
	"github.com/alibabacloud-go/tea/tea"
	antchainutil "github.com/antchain-openapi-sdk-go/antchain-util/service"
)

/**
 * Model for initing client
 */
type Config struct {
	// accesskey id
	AccessKeyId *string `json:"accessKeyId,omitempty" xml:"accessKeyId,omitempty"`
	// accesskey secret
	AccessKeySecret *string `json:"accessKeySecret,omitempty" xml:"accessKeySecret,omitempty"`
	// security token
	SecurityToken *string `json:"securityToken,omitempty" xml:"securityToken,omitempty"`
	// http protocol
	Protocol *string `json:"protocol,omitempty" xml:"protocol,omitempty"`
	// read timeout
	ReadTimeout *int `json:"readTimeout,omitempty" xml:"readTimeout,omitempty"`
	// connect timeout
	ConnectTimeout *int `json:"connectTimeout,omitempty" xml:"connectTimeout,omitempty"`
	// http proxy
	HttpProxy *string `json:"httpProxy,omitempty" xml:"httpProxy,omitempty"`
	// https proxy
	HttpsProxy *string `json:"httpsProxy,omitempty" xml:"httpsProxy,omitempty"`
	// endpoint
	Endpoint *string `json:"endpoint,omitempty" xml:"endpoint,omitempty"`
	// proxy white list
	NoProxy *string `json:"noProxy,omitempty" xml:"noProxy,omitempty"`
	// max idle conns
	MaxIdleConns *int `json:"maxIdleConns,omitempty" xml:"maxIdleConns,omitempty"`
	// user agent
	UserAgent *string `json:"userAgent,omitempty" xml:"userAgent,omitempty"`
	// socks5 proxy
	Socks5Proxy *string `json:"socks5Proxy,omitempty" xml:"socks5Proxy,omitempty"`
	// socks5 network
	Socks5NetWork *string `json:"socks5NetWork,omitempty" xml:"socks5NetWork,omitempty"`
	// 长链接最大空闲时长
	MaxIdleTimeMillis *int `json:"maxIdleTimeMillis,omitempty" xml:"maxIdleTimeMillis,omitempty"`
	// 长链接最大连接时长
	KeepAliveDurationMillis *int `json:"keepAliveDurationMillis,omitempty" xml:"keepAliveDurationMillis,omitempty"`
	// 最大连接数（长链接最大总数）
	MaxRequests *int `json:"maxRequests,omitempty" xml:"maxRequests,omitempty"`
	// 每个目标主机的最大连接数（分主机域名的长链接最大总数
	MaxRequestsPerHost *int `json:"maxRequestsPerHost,omitempty" xml:"maxRequestsPerHost,omitempty"`
}

func (s Config) String() string {
	return tea.Prettify(s)
}

func (s Config) GoString() string {
	return s.String()
}

func (s *Config) SetAccessKeyId(v string) *Config {
	s.AccessKeyId = &v
	return s
}

func (s *Config) SetAccessKeySecret(v string) *Config {
	s.AccessKeySecret = &v
	return s
}

func (s *Config) SetSecurityToken(v string) *Config {
	s.SecurityToken = &v
	return s
}

func (s *Config) SetProtocol(v string) *Config {
	s.Protocol = &v
	return s
}

func (s *Config) SetReadTimeout(v int) *Config {
	s.ReadTimeout = &v
	return s
}

func (s *Config) SetConnectTimeout(v int) *Config {
	s.ConnectTimeout = &v
	return s
}

func (s *Config) SetHttpProxy(v string) *Config {
	s.HttpProxy = &v
	return s
}

func (s *Config) SetHttpsProxy(v string) *Config {
	s.HttpsProxy = &v
	return s
}

func (s *Config) SetEndpoint(v string) *Config {
	s.Endpoint = &v
	return s
}

func (s *Config) SetNoProxy(v string) *Config {
	s.NoProxy = &v
	return s
}

func (s *Config) SetMaxIdleConns(v int) *Config {
	s.MaxIdleConns = &v
	return s
}

func (s *Config) SetUserAgent(v string) *Config {
	s.UserAgent = &v
	return s
}

func (s *Config) SetSocks5Proxy(v string) *Config {
	s.Socks5Proxy = &v
	return s
}

func (s *Config) SetSocks5NetWork(v string) *Config {
	s.Socks5NetWork = &v
	return s
}

func (s *Config) SetMaxIdleTimeMillis(v int) *Config {
	s.MaxIdleTimeMillis = &v
	return s
}

func (s *Config) SetKeepAliveDurationMillis(v int) *Config {
	s.KeepAliveDurationMillis = &v
	return s
}

func (s *Config) SetMaxRequests(v int) *Config {
	s.MaxRequests = &v
	return s
}

func (s *Config) SetMaxRequestsPerHost(v int) *Config {
	s.MaxRequestsPerHost = &v
	return s
}

// 集装箱信息
type ContainerInfo struct {
	// 订舱单唯一标识
	BookingNo *string `json:"booking_no,omitempty" xml:"booking_no,omitempty"`
	// 集装箱唯一标识
	ContainerId *string `json:"container_id,omitempty" xml:"container_id,omitempty"`
	// 箱号
	ContainerNo *string `json:"container_no,omitempty" xml:"container_no,omitempty"`
	// 箱型
	ContainerType *string `json:"container_type,omitempty" xml:"container_type,omitempty"`
}

func (s ContainerInfo) String() string {
	return tea.Prettify(s)
}

func (s ContainerInfo) GoString() string {
	return s.String()
}

func (s *ContainerInfo) SetBookingNo(v string) *ContainerInfo {
	s.BookingNo = &v
	return s
}

func (s *ContainerInfo) SetContainerId(v string) *ContainerInfo {
	s.ContainerId = &v
	return s
}

func (s *ContainerInfo) SetContainerNo(v string) *ContainerInfo {
	s.ContainerNo = &v
	return s
}

func (s *ContainerInfo) SetContainerType(v string) *ContainerInfo {
	s.ContainerType = &v
	return s
}

// 货主支付方式
type PayAmount struct {
	// 支付金额（2位小数）
	Amount *string `json:"amount,omitempty" xml:"amount,omitempty" require:"true"`
	// 支付方式
	PayType *string `json:"pay_type,omitempty" xml:"pay_type,omitempty" require:"true"`
}

func (s PayAmount) String() string {
	return tea.Prettify(s)
}

func (s PayAmount) GoString() string {
	return s.String()
}

func (s *PayAmount) SetAmount(v string) *PayAmount {
	s.Amount = &v
	return s
}

func (s *PayAmount) SetPayType(v string) *PayAmount {
	s.PayType = &v
	return s
}

// 电子提单批次下提单明细（无效）
type EblDeatil struct {
	// 电子提单copy文件hash
	EblCopyPdfFileHash *string `json:"ebl_copy_pdf_file_hash,omitempty" xml:"ebl_copy_pdf_file_hash,omitempty" require:"true"`
	// 电子提单copy文件id
	EblCopyPdfFileId *string `json:"ebl_copy_pdf_file_id,omitempty" xml:"ebl_copy_pdf_file_id,omitempty" require:"true"`
	// 电子提单编号
	EblNo *string `json:"ebl_no,omitempty" xml:"ebl_no,omitempty" require:"true"`
}

func (s EblDeatil) String() string {
	return tea.Prettify(s)
}

func (s EblDeatil) GoString() string {
	return s.String()
}

func (s *EblDeatil) SetEblCopyPdfFileHash(v string) *EblDeatil {
	s.EblCopyPdfFileHash = &v
	return s
}

func (s *EblDeatil) SetEblCopyPdfFileId(v string) *EblDeatil {
	s.EblCopyPdfFileId = &v
	return s
}

func (s *EblDeatil) SetEblNo(v string) *EblDeatil {
	s.EblNo = &v
	return s
}

// 集装箱列表
type HouseBlContainerParam struct {
	// 操作动作,为空为新增或更新，UPDATE为更新，DELETE为删除
	Action *string `json:"action,omitempty" xml:"action,omitempty"`
	// 集装箱ID
	ContainerId *string `json:"container_id,omitempty" xml:"container_id,omitempty" require:"true"`
	// 箱号
	ContainerNo *string `json:"container_no,omitempty" xml:"container_no,omitempty"`
}

func (s HouseBlContainerParam) String() string {
	return tea.Prettify(s)
}

func (s HouseBlContainerParam) GoString() string {
	return s.String()
}

func (s *HouseBlContainerParam) SetAction(v string) *HouseBlContainerParam {
	s.Action = &v
	return s
}

func (s *HouseBlContainerParam) SetContainerId(v string) *HouseBlContainerParam {
	s.ContainerId = &v
	return s
}

func (s *HouseBlContainerParam) SetContainerNo(v string) *HouseBlContainerParam {
	s.ContainerNo = &v
	return s
}

// 订舱单
type MasterBlBookingParam struct {
	// 操作动作,为空为新增或更新，UPDATE为更新，DELETE为删除
	Action *string `json:"action,omitempty" xml:"action,omitempty"`
	// 订舱单号
	BookingNo *string `json:"booking_no,omitempty" xml:"booking_no,omitempty" require:"true"`
}

func (s MasterBlBookingParam) String() string {
	return tea.Prettify(s)
}

func (s MasterBlBookingParam) GoString() string {
	return s.String()
}

func (s *MasterBlBookingParam) SetAction(v string) *MasterBlBookingParam {
	s.Action = &v
	return s
}

func (s *MasterBlBookingParam) SetBookingNo(v string) *MasterBlBookingParam {
	s.BookingNo = &v
	return s
}

// 包含文件id、文件hash信息
type UploadFileInfo struct {
	// 文件id
	FileId *string `json:"file_id,omitempty" xml:"file_id,omitempty" require:"true"`
	// 文件hash
	FileHash *string `json:"file_hash,omitempty" xml:"file_hash,omitempty" require:"true"`
}

func (s UploadFileInfo) String() string {
	return tea.Prettify(s)
}

func (s UploadFileInfo) GoString() string {
	return s.String()
}

func (s *UploadFileInfo) SetFileId(v string) *UploadFileInfo {
	s.FileId = &v
	return s
}

func (s *UploadFileInfo) SetFileHash(v string) *UploadFileInfo {
	s.FileHash = &v
	return s
}

// 箱子信息
type VehicleContainerParam struct {
	// 操作动作,为空为新增或更新，UPDATE为更新，DELETE为删除
	Action *string `json:"action,omitempty" xml:"action,omitempty"`
	// 集装箱ID
	ContainerId *string `json:"container_id,omitempty" xml:"container_id,omitempty" require:"true"`
	// 箱号
	ContainerNo *string `json:"container_no,omitempty" xml:"container_no,omitempty"`
	// 封号
	SealNo *string `json:"seal_no,omitempty" xml:"seal_no,omitempty"`
}

func (s VehicleContainerParam) String() string {
	return tea.Prettify(s)
}

func (s VehicleContainerParam) GoString() string {
	return s.String()
}

func (s *VehicleContainerParam) SetAction(v string) *VehicleContainerParam {
	s.Action = &v
	return s
}

func (s *VehicleContainerParam) SetContainerId(v string) *VehicleContainerParam {
	s.ContainerId = &v
	return s
}

func (s *VehicleContainerParam) SetContainerNo(v string) *VehicleContainerParam {
	s.ContainerNo = &v
	return s
}

func (s *VehicleContainerParam) SetSealNo(v string) *VehicleContainerParam {
	s.SealNo = &v
	return s
}

// 资费项账单
type PayBillTariffParam struct {
	// 操作动作,为空为新增或更新，UPDATE为更新，DELETE为删除
	Action *string `json:"action,omitempty" xml:"action,omitempty"`
	// 账单金额 业务必填
	BillAmount *string `json:"bill_amount,omitempty" xml:"bill_amount,omitempty"`
	// 应付账单、应付资费项 多对多code
	PayBillTariffCode *string `json:"pay_bill_tariff_code,omitempty" xml:"pay_bill_tariff_code,omitempty" require:"true"`
	//  资费项金额 业务必填
	PayTariffAmount *string `json:"pay_tariff_amount,omitempty" xml:"pay_tariff_amount,omitempty"`
	// 应付资费项编号 业务必填
	PayTariffCode *string `json:"pay_tariff_code,omitempty" xml:"pay_tariff_code,omitempty"`
}

func (s PayBillTariffParam) String() string {
	return tea.Prettify(s)
}

func (s PayBillTariffParam) GoString() string {
	return s.String()
}

func (s *PayBillTariffParam) SetAction(v string) *PayBillTariffParam {
	s.Action = &v
	return s
}

func (s *PayBillTariffParam) SetBillAmount(v string) *PayBillTariffParam {
	s.BillAmount = &v
	return s
}

func (s *PayBillTariffParam) SetPayBillTariffCode(v string) *PayBillTariffParam {
	s.PayBillTariffCode = &v
	return s
}

func (s *PayBillTariffParam) SetPayTariffAmount(v string) *PayBillTariffParam {
	s.PayTariffAmount = &v
	return s
}

func (s *PayBillTariffParam) SetPayTariffCode(v string) *PayBillTariffParam {
	s.PayTariffCode = &v
	return s
}

// 物流金融信用流转流水信息
type StatementInfo struct {
	// 信用流转批次号
	BatchId *string `json:"batch_id,omitempty" xml:"batch_id,omitempty" require:"true"`
	// 全局唯一业务号
	OutBizNo *string `json:"out_biz_no,omitempty" xml:"out_biz_no,omitempty" require:"true"`
	// 信用流转凭证
	IssueId *string `json:"issue_id,omitempty" xml:"issue_id,omitempty" require:"true"`
	// 合同号（预留）
	ContractId *string `json:"contract_id,omitempty" xml:"contract_id,omitempty"`
	// 发行信用流转的运单号
	WaybillId *string `json:"waybill_id,omitempty" xml:"waybill_id,omitempty" require:"true"`
	// 发行信用流转的支付单号
	PayOrder *string `json:"pay_order,omitempty" xml:"pay_order,omitempty" require:"true"`
	// 金额信息
	CreditLimit *string `json:"credit_limit,omitempty" xml:"credit_limit,omitempty" require:"true"`
	// 流水类型
	StateType *string `json:"state_type,omitempty" xml:"state_type,omitempty" require:"true"`
	// 流水类型说明
	StateMsg *string `json:"state_msg,omitempty" xml:"state_msg,omitempty" require:"true"`
	// 凭证来源方did
	FromDid *string `json:"from_did,omitempty" xml:"from_did,omitempty" require:"true"`
	// 凭证流转方did
	ToDid *string `json:"to_did,omitempty" xml:"to_did,omitempty" require:"true"`
	// 信用凭证发起时间
	IssueDate *string `json:"issue_date,omitempty" xml:"issue_date,omitempty" require:"true"`
	// 信用凭证到期时间
	ExpireDate *string `json:"expire_date,omitempty" xml:"expire_date,omitempty" require:"true"`
}

func (s StatementInfo) String() string {
	return tea.Prettify(s)
}

func (s StatementInfo) GoString() string {
	return s.String()
}

func (s *StatementInfo) SetBatchId(v string) *StatementInfo {
	s.BatchId = &v
	return s
}

func (s *StatementInfo) SetOutBizNo(v string) *StatementInfo {
	s.OutBizNo = &v
	return s
}

func (s *StatementInfo) SetIssueId(v string) *StatementInfo {
	s.IssueId = &v
	return s
}

func (s *StatementInfo) SetContractId(v string) *StatementInfo {
	s.ContractId = &v
	return s
}

func (s *StatementInfo) SetWaybillId(v string) *StatementInfo {
	s.WaybillId = &v
	return s
}

func (s *StatementInfo) SetPayOrder(v string) *StatementInfo {
	s.PayOrder = &v
	return s
}

func (s *StatementInfo) SetCreditLimit(v string) *StatementInfo {
	s.CreditLimit = &v
	return s
}

func (s *StatementInfo) SetStateType(v string) *StatementInfo {
	s.StateType = &v
	return s
}

func (s *StatementInfo) SetStateMsg(v string) *StatementInfo {
	s.StateMsg = &v
	return s
}

func (s *StatementInfo) SetFromDid(v string) *StatementInfo {
	s.FromDid = &v
	return s
}

func (s *StatementInfo) SetToDid(v string) *StatementInfo {
	s.ToDid = &v
	return s
}

func (s *StatementInfo) SetIssueDate(v string) *StatementInfo {
	s.IssueDate = &v
	return s
}

func (s *StatementInfo) SetExpireDate(v string) *StatementInfo {
	s.ExpireDate = &v
	return s
}

// 信用凭证数据集合
type IssueTransferData struct {
	// 凭证id
	IssueId *string `json:"issue_id,omitempty" xml:"issue_id,omitempty" require:"true"`
	// 转出方did
	PayerDid *string `json:"payer_did,omitempty" xml:"payer_did,omitempty" require:"true"`
	// 接收方did
	RcvDid *string `json:"rcv_did,omitempty" xml:"rcv_did,omitempty" require:"true"`
}

func (s IssueTransferData) String() string {
	return tea.Prettify(s)
}

func (s IssueTransferData) GoString() string {
	return s.String()
}

func (s *IssueTransferData) SetIssueId(v string) *IssueTransferData {
	s.IssueId = &v
	return s
}

func (s *IssueTransferData) SetPayerDid(v string) *IssueTransferData {
	s.PayerDid = &v
	return s
}

func (s *IssueTransferData) SetRcvDid(v string) *IssueTransferData {
	s.RcvDid = &v
	return s
}

// 订舱单号
type CustomsOrderBookingParam struct {
	// 操作动作,为空为新增或更新，UPDATE为更新，DELETE为删除
	Action *string `json:"action,omitempty" xml:"action,omitempty"`
	// 订舱单号
	BookingNo *string `json:"booking_no,omitempty" xml:"booking_no,omitempty" require:"true"`
}

func (s CustomsOrderBookingParam) String() string {
	return tea.Prettify(s)
}

func (s CustomsOrderBookingParam) GoString() string {
	return s.String()
}

func (s *CustomsOrderBookingParam) SetAction(v string) *CustomsOrderBookingParam {
	s.Action = &v
	return s
}

func (s *CustomsOrderBookingParam) SetBookingNo(v string) *CustomsOrderBookingParam {
	s.BookingNo = &v
	return s
}

// 航运集装箱类型信息
type ContainerTypeInfo struct {
	// 箱型
	ContainerType *string `json:"container_type,omitempty" xml:"container_type,omitempty"`
	// 箱量
	ContainerVolume *string `json:"container_volume,omitempty" xml:"container_volume,omitempty"`
}

func (s ContainerTypeInfo) String() string {
	return tea.Prettify(s)
}

func (s ContainerTypeInfo) GoString() string {
	return s.String()
}

func (s *ContainerTypeInfo) SetContainerType(v string) *ContainerTypeInfo {
	s.ContainerType = &v
	return s
}

func (s *ContainerTypeInfo) SetContainerVolume(v string) *ContainerTypeInfo {
	s.ContainerVolume = &v
	return s
}

// 池融资凭证核验查询结果
type PfVoucherCheckResult struct {
	// 凭证id
	VoucherId *string `json:"voucher_id,omitempty" xml:"voucher_id,omitempty" require:"true"`
	// 凭证类型
	VoucherCategory *string `json:"voucher_category,omitempty" xml:"voucher_category,omitempty" require:"true"`
	// 状态；PASS:通过，NO_PASS 未通过
	Status *string `json:"status,omitempty" xml:"status,omitempty" require:"true"`
}

func (s PfVoucherCheckResult) String() string {
	return tea.Prettify(s)
}

func (s PfVoucherCheckResult) GoString() string {
	return s.String()
}

func (s *PfVoucherCheckResult) SetVoucherId(v string) *PfVoucherCheckResult {
	s.VoucherId = &v
	return s
}

func (s *PfVoucherCheckResult) SetVoucherCategory(v string) *PfVoucherCheckResult {
	s.VoucherCategory = &v
	return s
}

func (s *PfVoucherCheckResult) SetStatus(v string) *PfVoucherCheckResult {
	s.Status = &v
	return s
}

// 凭证id发行信息
type IssueIdInfo struct {
	// 信用流转凭证
	IssueId *string `json:"issue_id,omitempty" xml:"issue_id,omitempty" require:"true"`
	// 全局唯一业务号
	OutBizNo *string `json:"out_biz_no,omitempty" xml:"out_biz_no,omitempty" require:"true"`
	// 合同号（预留）
	ContractId *string `json:"contract_id,omitempty" xml:"contract_id,omitempty"`
	// 发行信用流转的运单号
	WaybillId *string `json:"waybill_id,omitempty" xml:"waybill_id,omitempty" require:"true"`
	// 支付订单
	PayOrder *string `json:"pay_order,omitempty" xml:"pay_order,omitempty" require:"true"`
	// 凭证金额
	TicketAmt *string `json:"ticket_amt,omitempty" xml:"ticket_amt,omitempty" require:"true"`
	// 信用凭证发起时间
	IssueDate *string `json:"issue_date,omitempty" xml:"issue_date,omitempty" require:"true"`
	// 信用凭证到期时间
	ExpireDate *string `json:"expire_date,omitempty" xml:"expire_date,omitempty" require:"true"`
	// 发行结果状态 -1:发行失败状态， 0:未完成状态， 1:已发行状态
	Status *int64 `json:"status,omitempty" xml:"status,omitempty" require:"true"`
	// 失败原因信息
	ErrMsg *string `json:"err_msg,omitempty" xml:"err_msg,omitempty" require:"true"`
}

func (s IssueIdInfo) String() string {
	return tea.Prettify(s)
}

func (s IssueIdInfo) GoString() string {
	return s.String()
}

func (s *IssueIdInfo) SetIssueId(v string) *IssueIdInfo {
	s.IssueId = &v
	return s
}

func (s *IssueIdInfo) SetOutBizNo(v string) *IssueIdInfo {
	s.OutBizNo = &v
	return s
}

func (s *IssueIdInfo) SetContractId(v string) *IssueIdInfo {
	s.ContractId = &v
	return s
}

func (s *IssueIdInfo) SetWaybillId(v string) *IssueIdInfo {
	s.WaybillId = &v
	return s
}

func (s *IssueIdInfo) SetPayOrder(v string) *IssueIdInfo {
	s.PayOrder = &v
	return s
}

func (s *IssueIdInfo) SetTicketAmt(v string) *IssueIdInfo {
	s.TicketAmt = &v
	return s
}

func (s *IssueIdInfo) SetIssueDate(v string) *IssueIdInfo {
	s.IssueDate = &v
	return s
}

func (s *IssueIdInfo) SetExpireDate(v string) *IssueIdInfo {
	s.ExpireDate = &v
	return s
}

func (s *IssueIdInfo) SetStatus(v int64) *IssueIdInfo {
	s.Status = &v
	return s
}

func (s *IssueIdInfo) SetErrMsg(v string) *IssueIdInfo {
	s.ErrMsg = &v
	return s
}

// so通知关联的订舱单
type SoNotifyBookingParam struct {
	// 操作动作,为空为新增或更新，UPDATE为更新，DELETE为删除
	Action *string `json:"action,omitempty" xml:"action,omitempty"`
	// 订舱单号
	BookingNo *string `json:"booking_no,omitempty" xml:"booking_no,omitempty" require:"true"`
	// 船公司  业务必填
	Carrier *string `json:"carrier,omitempty" xml:"carrier,omitempty"`
	// 箱型箱量 json格式 业务必填
	// [{"containerVolume":"33","containerType":"22"}]
	// containerVolume--箱量 业务必填
	// containerType--箱型 业务必填
	ContainerParams *string `json:"container_params,omitempty" xml:"container_params,omitempty"`
	// 截关时间
	CustomsClearance *int64 `json:"customs_clearance,omitempty" xml:"customs_clearance,omitempty"`
	// 场站
	Cy *string `json:"cy,omitempty" xml:"cy,omitempty"`
	// 截港时间
	CyClosing *int64 `json:"cy_closing,omitempty" xml:"cy_closing,omitempty"`
	// 预计船期
	Etd *int64 `json:"etd,omitempty" xml:"etd,omitempty"`
	// 港口
	Port *string `json:"port,omitempty" xml:"port,omitempty"`
	// 截单时间
	SiClosing *int64 `json:"si_closing,omitempty" xml:"si_closing,omitempty"`
	// 船名  业务必填
	Vessel *string `json:"vessel,omitempty" xml:"vessel,omitempty"`
	// 航次 业务必填
	Voyage *string `json:"voyage,omitempty" xml:"voyage,omitempty"`
}

func (s SoNotifyBookingParam) String() string {
	return tea.Prettify(s)
}

func (s SoNotifyBookingParam) GoString() string {
	return s.String()
}

func (s *SoNotifyBookingParam) SetAction(v string) *SoNotifyBookingParam {
	s.Action = &v
	return s
}

func (s *SoNotifyBookingParam) SetBookingNo(v string) *SoNotifyBookingParam {
	s.BookingNo = &v
	return s
}

func (s *SoNotifyBookingParam) SetCarrier(v string) *SoNotifyBookingParam {
	s.Carrier = &v
	return s
}

func (s *SoNotifyBookingParam) SetContainerParams(v string) *SoNotifyBookingParam {
	s.ContainerParams = &v
	return s
}

func (s *SoNotifyBookingParam) SetCustomsClearance(v int64) *SoNotifyBookingParam {
	s.CustomsClearance = &v
	return s
}

func (s *SoNotifyBookingParam) SetCy(v string) *SoNotifyBookingParam {
	s.Cy = &v
	return s
}

func (s *SoNotifyBookingParam) SetCyClosing(v int64) *SoNotifyBookingParam {
	s.CyClosing = &v
	return s
}

func (s *SoNotifyBookingParam) SetEtd(v int64) *SoNotifyBookingParam {
	s.Etd = &v
	return s
}

func (s *SoNotifyBookingParam) SetPort(v string) *SoNotifyBookingParam {
	s.Port = &v
	return s
}

func (s *SoNotifyBookingParam) SetSiClosing(v int64) *SoNotifyBookingParam {
	s.SiClosing = &v
	return s
}

func (s *SoNotifyBookingParam) SetVessel(v string) *SoNotifyBookingParam {
	s.Vessel = &v
	return s
}

func (s *SoNotifyBookingParam) SetVoyage(v string) *SoNotifyBookingParam {
	s.Voyage = &v
	return s
}

// 接口测试
type ApiTest struct {
	// 测试String
	TestString *string `json:"test_string,omitempty" xml:"test_string,omitempty" require:"true" maxLength:"100" minLength:"0"`
	// 测试Int
	TestInt *int64 `json:"test_int,omitempty" xml:"test_int,omitempty" require:"true" maximum:"100" minimum:"0"`
	// 测试Integer
	TestInteger *int64 `json:"test_integer,omitempty" xml:"test_integer,omitempty" require:"true" maximum:"100" minimum:"0"`
	// 测试Long
	TestLong *int64 `json:"test_long,omitempty" xml:"test_long,omitempty" require:"true" maximum:"100" minimum:"0"`
	// 测试Boolean
	TestBoolean *bool `json:"test_boolean,omitempty" xml:"test_boolean,omitempty" require:"true"`
	// 测试Date
	TestDate *string `json:"test_date,omitempty" xml:"test_date,omitempty" require:"true" pattern:"\\d{4}[-]\\d{1,2}[-]\\d{1,2}[T]\\d{2}:\\d{2}:\\d{2}([Z]|([\\.]\\d{1,9})?[\\+]\\d{2}[\\:]?\\d{2})"`
	// 凭证列表_stringList
	StringList []*string `json:"string_list,omitempty" xml:"string_list,omitempty" require:"true" type:"Repeated"`
	// 凭证列表_ints
	Ints []*int64 `json:"ints,omitempty" xml:"ints,omitempty" require:"true" type:"Repeated"`
	// 凭证列表_longs
	Longs []*int64 `json:"longs,omitempty" xml:"longs,omitempty" require:"true" type:"Repeated"`
	// 凭证列表_integerList
	IntegerList []*int64 `json:"integer_list,omitempty" xml:"integer_list,omitempty" require:"true" type:"Repeated"`
	// 凭证列表_longList
	LongList []*int64 `json:"long_list,omitempty" xml:"long_list,omitempty" require:"true" type:"Repeated"`
	// 凭证列表_booleanList
	BooleanList []*bool `json:"boolean_list,omitempty" xml:"boolean_list,omitempty" require:"true" type:"Repeated"`
	// 凭证列表_dateList
	DateList []*string `json:"date_list,omitempty" xml:"date_list,omitempty" require:"true" type:"Repeated"`
	// 凭证列表_apiTestList
	ApiTestList []*ApiTestInfo `json:"api_test_list,omitempty" xml:"api_test_list,omitempty" require:"true" type:"Repeated"`
	// 测试apiTestInfo
	ApiTestInfo *ApiTest `json:"api_test_info,omitempty" xml:"api_test_info,omitempty" require:"true"`
}

func (s ApiTest) String() string {
	return tea.Prettify(s)
}

func (s ApiTest) GoString() string {
	return s.String()
}

func (s *ApiTest) SetTestString(v string) *ApiTest {
	s.TestString = &v
	return s
}

func (s *ApiTest) SetTestInt(v int64) *ApiTest {
	s.TestInt = &v
	return s
}

func (s *ApiTest) SetTestInteger(v int64) *ApiTest {
	s.TestInteger = &v
	return s
}

func (s *ApiTest) SetTestLong(v int64) *ApiTest {
	s.TestLong = &v
	return s
}

func (s *ApiTest) SetTestBoolean(v bool) *ApiTest {
	s.TestBoolean = &v
	return s
}

func (s *ApiTest) SetTestDate(v string) *ApiTest {
	s.TestDate = &v
	return s
}

func (s *ApiTest) SetStringList(v []*string) *ApiTest {
	s.StringList = v
	return s
}

func (s *ApiTest) SetInts(v []*int64) *ApiTest {
	s.Ints = v
	return s
}

func (s *ApiTest) SetLongs(v []*int64) *ApiTest {
	s.Longs = v
	return s
}

func (s *ApiTest) SetIntegerList(v []*int64) *ApiTest {
	s.IntegerList = v
	return s
}

func (s *ApiTest) SetLongList(v []*int64) *ApiTest {
	s.LongList = v
	return s
}

func (s *ApiTest) SetBooleanList(v []*bool) *ApiTest {
	s.BooleanList = v
	return s
}

func (s *ApiTest) SetDateList(v []*string) *ApiTest {
	s.DateList = v
	return s
}

func (s *ApiTest) SetApiTestList(v []*ApiTestInfo) *ApiTest {
	s.ApiTestList = v
	return s
}

func (s *ApiTest) SetApiTestInfo(v *ApiTest) *ApiTest {
	s.ApiTestInfo = v
	return s
}

// 资费项账单
type ReceiptBillTariffParam struct {
	// 操作动作,为空为新增或更新，UPDATE为更新，DELETE为删除
	Action *string `json:"action,omitempty" xml:"action,omitempty"`
	//  账单金额 业务必填
	BillAmount *string `json:"bill_amount,omitempty" xml:"bill_amount,omitempty"`
	// 应收账单 、应收资费项 多对多关联code
	ReceiptBillTariffCode *string `json:"receipt_bill_tariff_code,omitempty" xml:"receipt_bill_tariff_code,omitempty" require:"true"`
	// 资费项金额 业务必填
	ReceiptTariffAmount *string `json:"receipt_tariff_amount,omitempty" xml:"receipt_tariff_amount,omitempty"`
	// 应收资费项编号 业务必填
	ReceiptTariffCode *string `json:"receipt_tariff_code,omitempty" xml:"receipt_tariff_code,omitempty"`
}

func (s ReceiptBillTariffParam) String() string {
	return tea.Prettify(s)
}

func (s ReceiptBillTariffParam) GoString() string {
	return s.String()
}

func (s *ReceiptBillTariffParam) SetAction(v string) *ReceiptBillTariffParam {
	s.Action = &v
	return s
}

func (s *ReceiptBillTariffParam) SetBillAmount(v string) *ReceiptBillTariffParam {
	s.BillAmount = &v
	return s
}

func (s *ReceiptBillTariffParam) SetReceiptBillTariffCode(v string) *ReceiptBillTariffParam {
	s.ReceiptBillTariffCode = &v
	return s
}

func (s *ReceiptBillTariffParam) SetReceiptTariffAmount(v string) *ReceiptBillTariffParam {
	s.ReceiptTariffAmount = &v
	return s
}

func (s *ReceiptBillTariffParam) SetReceiptTariffCode(v string) *ReceiptBillTariffParam {
	s.ReceiptTariffCode = &v
	return s
}

// 航运集装箱ID信息
type ContainerIdInfo struct {
	// 箱子唯一标识
	ContainerId *string `json:"container_id,omitempty" xml:"container_id,omitempty"`
	// 箱号
	ContainerNo *string `json:"container_no,omitempty" xml:"container_no,omitempty"`
}

func (s ContainerIdInfo) String() string {
	return tea.Prettify(s)
}

func (s ContainerIdInfo) GoString() string {
	return s.String()
}

func (s *ContainerIdInfo) SetContainerId(v string) *ContainerIdInfo {
	s.ContainerId = &v
	return s
}

func (s *ContainerIdInfo) SetContainerNo(v string) *ContainerIdInfo {
	s.ContainerNo = &v
	return s
}

// 货物信息
type GoodsInfo struct {
	// 货物ID [业务必填]
	GoodsId *string `json:"goods_id,omitempty" xml:"goods_id,omitempty"`
	// 唛头
	//
	//
	Marks *string `json:"marks,omitempty" xml:"marks,omitempty"`
	// 货物名称
	Goods *string `json:"goods,omitempty" xml:"goods,omitempty"`
	// 货物类型
	GoodsType *string `json:"goods_type,omitempty" xml:"goods_type,omitempty"`
	// 货物重量
	Weight *string `json:"weight,omitempty" xml:"weight,omitempty"`
	// 件数
	Number *string `json:"number,omitempty" xml:"number,omitempty"`
}

func (s GoodsInfo) String() string {
	return tea.Prettify(s)
}

func (s GoodsInfo) GoString() string {
	return s.String()
}

func (s *GoodsInfo) SetGoodsId(v string) *GoodsInfo {
	s.GoodsId = &v
	return s
}

func (s *GoodsInfo) SetMarks(v string) *GoodsInfo {
	s.Marks = &v
	return s
}

func (s *GoodsInfo) SetGoods(v string) *GoodsInfo {
	s.Goods = &v
	return s
}

func (s *GoodsInfo) SetGoodsType(v string) *GoodsInfo {
	s.GoodsType = &v
	return s
}

func (s *GoodsInfo) SetWeight(v string) *GoodsInfo {
	s.Weight = &v
	return s
}

func (s *GoodsInfo) SetNumber(v string) *GoodsInfo {
	s.Number = &v
	return s
}

// 授权上链文件
type AuthChainFile struct {
	// 签署文件的hash值
	SignFileHash *string `json:"sign_file_hash,omitempty" xml:"sign_file_hash,omitempty" require:"true"`
	// 上链事务唯一标识
	UploadChainTxCode *string `json:"upload_chain_tx_code,omitempty" xml:"upload_chain_tx_code,omitempty" require:"true"`
	// 蚂蚁区块链统一证据编号
	BaasUniqCode *string `json:"baas_uniq_code,omitempty" xml:"baas_uniq_code,omitempty" require:"true"`
	// 上链时间(13位毫秒级时间戳)
	UploadChainTime *string `json:"upload_chain_time,omitempty" xml:"upload_chain_time,omitempty" require:"true"`
	// 上链文件下载链接
	FileUrl *string `json:"file_url,omitempty" xml:"file_url,omitempty" require:"true"`
	// 上链文件名称，要求包含扩展名。文件格式允许: pdf, txt, doc, docx
	FileName *string `json:"file_name,omitempty" xml:"file_name,omitempty" require:"true"`
}

func (s AuthChainFile) String() string {
	return tea.Prettify(s)
}

func (s AuthChainFile) GoString() string {
	return s.String()
}

func (s *AuthChainFile) SetSignFileHash(v string) *AuthChainFile {
	s.SignFileHash = &v
	return s
}

func (s *AuthChainFile) SetUploadChainTxCode(v string) *AuthChainFile {
	s.UploadChainTxCode = &v
	return s
}

func (s *AuthChainFile) SetBaasUniqCode(v string) *AuthChainFile {
	s.BaasUniqCode = &v
	return s
}

func (s *AuthChainFile) SetUploadChainTime(v string) *AuthChainFile {
	s.UploadChainTime = &v
	return s
}

func (s *AuthChainFile) SetFileUrl(v string) *AuthChainFile {
	s.FileUrl = &v
	return s
}

func (s *AuthChainFile) SetFileName(v string) *AuthChainFile {
	s.FileName = &v
	return s
}

// 池融资授信额度信息
type PfCreditQuotaInfo struct {
	// 证件号
	CertNo *string `json:"cert_no,omitempty" xml:"cert_no,omitempty" require:"true"`
	// 证件类型
	CertType *string `json:"cert_type,omitempty" xml:"cert_type,omitempty" require:"true"`
	// 授信到期日期
	CreditEnd *string `json:"credit_end,omitempty" xml:"credit_end,omitempty" require:"true"`
	// 授信起始日期
	CreditStart *string `json:"credit_start,omitempty" xml:"credit_start,omitempty" require:"true"`
	// 额度编号
	QuotaNo *string `json:"quota_no,omitempty" xml:"quota_no,omitempty" require:"true"`
	// 剩余额度
	RemainingQuota *string `json:"remaining_quota,omitempty" xml:"remaining_quota,omitempty" require:"true"`
	// SON:放款账号loanAccNo
	// 还款账号repayAcctNo
	Remark *string `json:"remark,omitempty" xml:"remark,omitempty" require:"true"`
	// 额度状态：
	// 0、停用 / 1、启用  /  2、冻结
	Status *string `json:"status,omitempty" xml:"status,omitempty" require:"true"`
	// 授信额度
	TotalQuota *string `json:"total_quota,omitempty" xml:"total_quota,omitempty" require:"true"`
	// 数据更新时间
	UpdateTime *string `json:"update_time,omitempty" xml:"update_time,omitempty" require:"true"`
	// 总质押额度
	TotalPledgeQuota *string `json:"total_pledge_quota,omitempty" xml:"total_pledge_quota,omitempty" require:"true"`
	// 剩余质押额度
	RemainPledgeQuota *string `json:"remain_pledge_quota,omitempty" xml:"remain_pledge_quota,omitempty" require:"true"`
}

func (s PfCreditQuotaInfo) String() string {
	return tea.Prettify(s)
}

func (s PfCreditQuotaInfo) GoString() string {
	return s.String()
}

func (s *PfCreditQuotaInfo) SetCertNo(v string) *PfCreditQuotaInfo {
	s.CertNo = &v
	return s
}

func (s *PfCreditQuotaInfo) SetCertType(v string) *PfCreditQuotaInfo {
	s.CertType = &v
	return s
}

func (s *PfCreditQuotaInfo) SetCreditEnd(v string) *PfCreditQuotaInfo {
	s.CreditEnd = &v
	return s
}

func (s *PfCreditQuotaInfo) SetCreditStart(v string) *PfCreditQuotaInfo {
	s.CreditStart = &v
	return s
}

func (s *PfCreditQuotaInfo) SetQuotaNo(v string) *PfCreditQuotaInfo {
	s.QuotaNo = &v
	return s
}

func (s *PfCreditQuotaInfo) SetRemainingQuota(v string) *PfCreditQuotaInfo {
	s.RemainingQuota = &v
	return s
}

func (s *PfCreditQuotaInfo) SetRemark(v string) *PfCreditQuotaInfo {
	s.Remark = &v
	return s
}

func (s *PfCreditQuotaInfo) SetStatus(v string) *PfCreditQuotaInfo {
	s.Status = &v
	return s
}

func (s *PfCreditQuotaInfo) SetTotalQuota(v string) *PfCreditQuotaInfo {
	s.TotalQuota = &v
	return s
}

func (s *PfCreditQuotaInfo) SetUpdateTime(v string) *PfCreditQuotaInfo {
	s.UpdateTime = &v
	return s
}

func (s *PfCreditQuotaInfo) SetTotalPledgeQuota(v string) *PfCreditQuotaInfo {
	s.TotalPledgeQuota = &v
	return s
}

func (s *PfCreditQuotaInfo) SetRemainPledgeQuota(v string) *PfCreditQuotaInfo {
	s.RemainPledgeQuota = &v
	return s
}

// 应付账单发票关联项
type PayBillInvoiceParam struct {
	// 账单发票code
	PayBillInvoiceCode *string `json:"pay_bill_invoice_code,omitempty" xml:"pay_bill_invoice_code,omitempty" require:"true"`
	// 账单编号
	PayBillOrderCode *string `json:"pay_bill_order_code,omitempty" xml:"pay_bill_order_code,omitempty" require:"true"`
	// 账单金额
	PayBillAmount *string `json:"pay_bill_amount,omitempty" xml:"pay_bill_amount,omitempty" require:"true"`
	// 发票金额
	InvoiceAmount *string `json:"invoice_amount,omitempty" xml:"invoice_amount,omitempty" require:"true"`
	// 操作动作
	Action *string `json:"action,omitempty" xml:"action,omitempty"`
}

func (s PayBillInvoiceParam) String() string {
	return tea.Prettify(s)
}

func (s PayBillInvoiceParam) GoString() string {
	return s.String()
}

func (s *PayBillInvoiceParam) SetPayBillInvoiceCode(v string) *PayBillInvoiceParam {
	s.PayBillInvoiceCode = &v
	return s
}

func (s *PayBillInvoiceParam) SetPayBillOrderCode(v string) *PayBillInvoiceParam {
	s.PayBillOrderCode = &v
	return s
}

func (s *PayBillInvoiceParam) SetPayBillAmount(v string) *PayBillInvoiceParam {
	s.PayBillAmount = &v
	return s
}

func (s *PayBillInvoiceParam) SetInvoiceAmount(v string) *PayBillInvoiceParam {
	s.InvoiceAmount = &v
	return s
}

func (s *PayBillInvoiceParam) SetAction(v string) *PayBillInvoiceParam {
	s.Action = &v
	return s
}

// 资费项发票
type PayTariffInvoiceParam struct {
	// 操作动作,为空为新增或更新，UPDATE为更新，DELETE为删除
	Action *string `json:"action,omitempty" xml:"action,omitempty"`
	// 发票金额 业务必填
	InvoiceAmount *string `json:"invoice_amount,omitempty" xml:"invoice_amount,omitempty"`
	// 资费项金额 业务必填
	PayTariffAmount *string `json:"pay_tariff_amount,omitempty" xml:"pay_tariff_amount,omitempty"`
	// 资费单据编号 业务必填
	PayTariffCode *string `json:"pay_tariff_code,omitempty" xml:"pay_tariff_code,omitempty"`
	// 资费项发票code
	PayTariffInvoiceCode *string `json:"pay_tariff_invoice_code,omitempty" xml:"pay_tariff_invoice_code,omitempty" require:"true"`
}

func (s PayTariffInvoiceParam) String() string {
	return tea.Prettify(s)
}

func (s PayTariffInvoiceParam) GoString() string {
	return s.String()
}

func (s *PayTariffInvoiceParam) SetAction(v string) *PayTariffInvoiceParam {
	s.Action = &v
	return s
}

func (s *PayTariffInvoiceParam) SetInvoiceAmount(v string) *PayTariffInvoiceParam {
	s.InvoiceAmount = &v
	return s
}

func (s *PayTariffInvoiceParam) SetPayTariffAmount(v string) *PayTariffInvoiceParam {
	s.PayTariffAmount = &v
	return s
}

func (s *PayTariffInvoiceParam) SetPayTariffCode(v string) *PayTariffInvoiceParam {
	s.PayTariffCode = &v
	return s
}

func (s *PayTariffInvoiceParam) SetPayTariffInvoiceCode(v string) *PayTariffInvoiceParam {
	s.PayTariffInvoiceCode = &v
	return s
}

// 投保基本信息
type InsureBaseInfo struct {
	// 投保人姓名
	TbrName *string `json:"tbr_name,omitempty" xml:"tbr_name,omitempty" require:"true"`
	// 投保人证件号
	TbrIdNo *string `json:"tbr_id_no,omitempty" xml:"tbr_id_no,omitempty" require:"true"`
	// 投保人证件类型
	TbrIdType *string `json:"tbr_id_type,omitempty" xml:"tbr_id_type,omitempty" require:"true"`
	// 投保人联系电话
	TbrTel *string `json:"tbr_tel,omitempty" xml:"tbr_tel,omitempty" require:"true"`
	// 投保人地址
	TbrAddr *string `json:"tbr_addr,omitempty" xml:"tbr_addr,omitempty"`
	// 投保人邮箱
	TbrEmail *string `json:"tbr_email,omitempty" xml:"tbr_email,omitempty" require:"true"`
	// 被保险人姓名
	BbrName *string `json:"bbr_name,omitempty" xml:"bbr_name,omitempty" require:"true"`
	// 被保险人证件类型
	BbrIdType *string `json:"bbr_id_type,omitempty" xml:"bbr_id_type,omitempty" require:"true"`
	// 被保险人证件号码
	BbrIdNo *string `json:"bbr_id_no,omitempty" xml:"bbr_id_no,omitempty" require:"true"`
	// 被保险人联系电话
	BbrTel *string `json:"bbr_tel,omitempty" xml:"bbr_tel,omitempty" require:"true"`
	// 被保险人地址
	BbrAddr *string `json:"bbr_addr,omitempty" xml:"bbr_addr,omitempty"`
	// 含税保费(元)，小数点两位
	PreMium *string `json:"pre_mium,omitempty" xml:"pre_mium,omitempty" require:"true"`
	// 保险起期
	EffDate *string `json:"eff_date,omitempty" xml:"eff_date,omitempty" require:"true" pattern:"\\d{4}[-]\\d{1,2}[-]\\d{1,2}[T]\\d{2}:\\d{2}:\\d{2}([Z]|([\\.]\\d{1,9})?[\\+]\\d{2}[\\:]?\\d{2})"`
	// 保险止期
	TermDate *string `json:"term_date,omitempty" xml:"term_date,omitempty" require:"true" pattern:"\\d{4}[-]\\d{1,2}[-]\\d{1,2}[T]\\d{2}:\\d{2}:\\d{2}([Z]|([\\.]\\d{1,9})?[\\+]\\d{2}[\\:]?\\d{2})"`
	// 投保人证件类型有效起期
	IdenrifyPeriodStart *string `json:"idenrify_period_start,omitempty" xml:"idenrify_period_start,omitempty" pattern:"\\d{4}[-]\\d{1,2}[-]\\d{1,2}[T]\\d{2}:\\d{2}:\\d{2}([Z]|([\\.]\\d{1,9})?[\\+]\\d{2}[\\:]?\\d{2})"`
	// 投保人证件类型有效止期
	IdentifyPeriodEnd *string `json:"identify_period_end,omitempty" xml:"identify_period_end,omitempty" pattern:"\\d{4}[-]\\d{1,2}[-]\\d{1,2}[T]\\d{2}:\\d{2}:\\d{2}([Z]|([\\.]\\d{1,9})?[\\+]\\d{2}[\\:]?\\d{2})"`
}

func (s InsureBaseInfo) String() string {
	return tea.Prettify(s)
}

func (s InsureBaseInfo) GoString() string {
	return s.String()
}

func (s *InsureBaseInfo) SetTbrName(v string) *InsureBaseInfo {
	s.TbrName = &v
	return s
}

func (s *InsureBaseInfo) SetTbrIdNo(v string) *InsureBaseInfo {
	s.TbrIdNo = &v
	return s
}

func (s *InsureBaseInfo) SetTbrIdType(v string) *InsureBaseInfo {
	s.TbrIdType = &v
	return s
}

func (s *InsureBaseInfo) SetTbrTel(v string) *InsureBaseInfo {
	s.TbrTel = &v
	return s
}

func (s *InsureBaseInfo) SetTbrAddr(v string) *InsureBaseInfo {
	s.TbrAddr = &v
	return s
}

func (s *InsureBaseInfo) SetTbrEmail(v string) *InsureBaseInfo {
	s.TbrEmail = &v
	return s
}

func (s *InsureBaseInfo) SetBbrName(v string) *InsureBaseInfo {
	s.BbrName = &v
	return s
}

func (s *InsureBaseInfo) SetBbrIdType(v string) *InsureBaseInfo {
	s.BbrIdType = &v
	return s
}

func (s *InsureBaseInfo) SetBbrIdNo(v string) *InsureBaseInfo {
	s.BbrIdNo = &v
	return s
}

func (s *InsureBaseInfo) SetBbrTel(v string) *InsureBaseInfo {
	s.BbrTel = &v
	return s
}

func (s *InsureBaseInfo) SetBbrAddr(v string) *InsureBaseInfo {
	s.BbrAddr = &v
	return s
}

func (s *InsureBaseInfo) SetPreMium(v string) *InsureBaseInfo {
	s.PreMium = &v
	return s
}

func (s *InsureBaseInfo) SetEffDate(v string) *InsureBaseInfo {
	s.EffDate = &v
	return s
}

func (s *InsureBaseInfo) SetTermDate(v string) *InsureBaseInfo {
	s.TermDate = &v
	return s
}

func (s *InsureBaseInfo) SetIdenrifyPeriodStart(v string) *InsureBaseInfo {
	s.IdenrifyPeriodStart = &v
	return s
}

func (s *InsureBaseInfo) SetIdentifyPeriodEnd(v string) *InsureBaseInfo {
	s.IdentifyPeriodEnd = &v
	return s
}

// 电子回单查询，具体凭证数据
type ScpTicketIssueData struct {
	// 凭证对应的司机/货主的did
	Did *string `json:"did,omitempty" xml:"did,omitempty" require:"true"`
	//
	// 凭证id
	IssueId *string `json:"issue_id,omitempty" xml:"issue_id,omitempty" require:"true"`
}

func (s ScpTicketIssueData) String() string {
	return tea.Prettify(s)
}

func (s ScpTicketIssueData) GoString() string {
	return s.String()
}

func (s *ScpTicketIssueData) SetDid(v string) *ScpTicketIssueData {
	s.Did = &v
	return s
}

func (s *ScpTicketIssueData) SetIssueId(v string) *ScpTicketIssueData {
	s.IssueId = &v
	return s
}

// 货物列表
type HouseBlGoodsParam struct {
	// 预计备货时间
	CargoReadyDate *int64 `json:"cargo_ready_date,omitempty" xml:"cargo_ready_date,omitempty"`
	// 危险品页号
	DgPageNo *string `json:"dg_page_no,omitempty" xml:"dg_page_no,omitempty"`
	// 危险品级别
	DgType *string `json:"dg_type,omitempty" xml:"dg_type,omitempty"`
	// 危险品闪点
	FlashPoint *string `json:"flash_point,omitempty" xml:"flash_point,omitempty"`
	// 货物名称 业务必填
	Goods *string `json:"goods,omitempty" xml:"goods,omitempty"`
	// 货物中文名
	GoodsCn *string `json:"goods_cn,omitempty" xml:"goods_cn,omitempty"`
	// 货物类型 业务必填
	GoodsType *string `json:"goods_type,omitempty" xml:"goods_type,omitempty"`
	// HS CODE
	HsCodes []*string `json:"hs_codes,omitempty" xml:"hs_codes,omitempty" type:"Repeated"`
	// 唛头
	Marks *string `json:"marks,omitempty" xml:"marks,omitempty"`
	// 件数 业务必填
	Number *string `json:"number,omitempty" xml:"number,omitempty"`
	// 包装类型
	PackageType *string `json:"package_type,omitempty" xml:"package_type,omitempty"`
	// 实际件数
	RealNumber *string `json:"real_number,omitempty" xml:"real_number,omitempty"`
	// 实际体积
	RealVolume *string `json:"real_volume,omitempty" xml:"real_volume,omitempty"`
	// 实际重量
	RealWeight *string `json:"real_weight,omitempty" xml:"real_weight,omitempty"`
	// 危险品联合国编号
	UnNo *string `json:"un_no,omitempty" xml:"un_no,omitempty"`
	// 委托体积 业务必填
	Volume *string `json:"volume,omitempty" xml:"volume,omitempty"`
	// 委托重量 业务必填
	Weight *string `json:"weight,omitempty" xml:"weight,omitempty"`
}

func (s HouseBlGoodsParam) String() string {
	return tea.Prettify(s)
}

func (s HouseBlGoodsParam) GoString() string {
	return s.String()
}

func (s *HouseBlGoodsParam) SetCargoReadyDate(v int64) *HouseBlGoodsParam {
	s.CargoReadyDate = &v
	return s
}

func (s *HouseBlGoodsParam) SetDgPageNo(v string) *HouseBlGoodsParam {
	s.DgPageNo = &v
	return s
}

func (s *HouseBlGoodsParam) SetDgType(v string) *HouseBlGoodsParam {
	s.DgType = &v
	return s
}

func (s *HouseBlGoodsParam) SetFlashPoint(v string) *HouseBlGoodsParam {
	s.FlashPoint = &v
	return s
}

func (s *HouseBlGoodsParam) SetGoods(v string) *HouseBlGoodsParam {
	s.Goods = &v
	return s
}

func (s *HouseBlGoodsParam) SetGoodsCn(v string) *HouseBlGoodsParam {
	s.GoodsCn = &v
	return s
}

func (s *HouseBlGoodsParam) SetGoodsType(v string) *HouseBlGoodsParam {
	s.GoodsType = &v
	return s
}

func (s *HouseBlGoodsParam) SetHsCodes(v []*string) *HouseBlGoodsParam {
	s.HsCodes = v
	return s
}

func (s *HouseBlGoodsParam) SetMarks(v string) *HouseBlGoodsParam {
	s.Marks = &v
	return s
}

func (s *HouseBlGoodsParam) SetNumber(v string) *HouseBlGoodsParam {
	s.Number = &v
	return s
}

func (s *HouseBlGoodsParam) SetPackageType(v string) *HouseBlGoodsParam {
	s.PackageType = &v
	return s
}

func (s *HouseBlGoodsParam) SetRealNumber(v string) *HouseBlGoodsParam {
	s.RealNumber = &v
	return s
}

func (s *HouseBlGoodsParam) SetRealVolume(v string) *HouseBlGoodsParam {
	s.RealVolume = &v
	return s
}

func (s *HouseBlGoodsParam) SetRealWeight(v string) *HouseBlGoodsParam {
	s.RealWeight = &v
	return s
}

func (s *HouseBlGoodsParam) SetUnNo(v string) *HouseBlGoodsParam {
	s.UnNo = &v
	return s
}

func (s *HouseBlGoodsParam) SetVolume(v string) *HouseBlGoodsParam {
	s.Volume = &v
	return s
}

func (s *HouseBlGoodsParam) SetWeight(v string) *HouseBlGoodsParam {
	s.Weight = &v
	return s
}

// 电子提单批次下提单明细
type EblDetail struct {
	// 电子提单copy文件hash
	EblCopyPdfFileHash *string `json:"ebl_copy_pdf_file_hash,omitempty" xml:"ebl_copy_pdf_file_hash,omitempty" require:"true"`
	// 电子提单copy文件id
	EblCopyPdfFileId *string `json:"ebl_copy_pdf_file_id,omitempty" xml:"ebl_copy_pdf_file_id,omitempty" require:"true"`
	// 电子提单编号
	EblNo *string `json:"ebl_no,omitempty" xml:"ebl_no,omitempty" require:"true"`
}

func (s EblDetail) String() string {
	return tea.Prettify(s)
}

func (s EblDetail) GoString() string {
	return s.String()
}

func (s *EblDetail) SetEblCopyPdfFileHash(v string) *EblDetail {
	s.EblCopyPdfFileHash = &v
	return s
}

func (s *EblDetail) SetEblCopyPdfFileId(v string) *EblDetail {
	s.EblCopyPdfFileId = &v
	return s
}

func (s *EblDetail) SetEblNo(v string) *EblDetail {
	s.EblNo = &v
	return s
}

// 历史数据
type UploadFinancingParam struct {
	// 订舱单量（票）
	BookingCount *int64 `json:"booking_count,omitempty" xml:"booking_count,omitempty" require:"true"`
	// 唯一标识
	Code *string `json:"code,omitempty" xml:"code,omitempty" require:"true"`
	// 结束日期
	EndDate *string `json:"end_date,omitempty" xml:"end_date,omitempty" require:"true"`
	// 货代did
	ForwarderDid *string `json:"forwarder_did,omitempty" xml:"forwarder_did,omitempty" require:"true"`
	// 开始日期
	StartDate *string `json:"start_date,omitempty" xml:"start_date,omitempty" require:"true"`
	// 箱量【数量，不区分箱型，20GP是1TEU，40GP是2TEU】
	Teu *int64 `json:"teu,omitempty" xml:"teu,omitempty" require:"true"`
	// 运输总额
	Amounts *string `json:"amounts,omitempty" xml:"amounts,omitempty" require:"true"`
}

func (s UploadFinancingParam) String() string {
	return tea.Prettify(s)
}

func (s UploadFinancingParam) GoString() string {
	return s.String()
}

func (s *UploadFinancingParam) SetBookingCount(v int64) *UploadFinancingParam {
	s.BookingCount = &v
	return s
}

func (s *UploadFinancingParam) SetCode(v string) *UploadFinancingParam {
	s.Code = &v
	return s
}

func (s *UploadFinancingParam) SetEndDate(v string) *UploadFinancingParam {
	s.EndDate = &v
	return s
}

func (s *UploadFinancingParam) SetForwarderDid(v string) *UploadFinancingParam {
	s.ForwarderDid = &v
	return s
}

func (s *UploadFinancingParam) SetStartDate(v string) *UploadFinancingParam {
	s.StartDate = &v
	return s
}

func (s *UploadFinancingParam) SetTeu(v int64) *UploadFinancingParam {
	s.Teu = &v
	return s
}

func (s *UploadFinancingParam) SetAmounts(v string) *UploadFinancingParam {
	s.Amounts = &v
	return s
}

// 库存货物
type InventoryCargo struct {
	// 序号，在同一次库存申报请求中，序号保持不重复，不能小于等于0
	InventoryIndex *int64 `json:"inventory_index,omitempty" xml:"inventory_index,omitempty" require:"true"`
	// sku品名
	Sku *string `json:"sku,omitempty" xml:"sku,omitempty" require:"true" maxLength:"200"`
	// 商品名称
	//
	CargoName *string `json:"cargo_name,omitempty" xml:"cargo_name,omitempty" maxLength:"200"`
	// 商品单品重量(kg)
	CargoWeight *string `json:"cargo_weight,omitempty" xml:"cargo_weight,omitempty" maxLength:"50"`
	// 商品外扩长宽高(cm)
	CargoDimensions *string `json:"cargo_dimensions,omitempty" xml:"cargo_dimensions,omitempty" maxLength:"200"`
	// 商品单品货物价值(元),最多支持2位小数
	CargoWorth *string `json:"cargo_worth,omitempty" xml:"cargo_worth,omitempty" maxLength:"30"`
	// 当前库存货物数量
	CurrentInventoryCargoNum *int64 `json:"current_inventory_cargo_num,omitempty" xml:"current_inventory_cargo_num,omitempty" require:"true"`
	// 客户代码
	//
	CustomerCode *string `json:"customer_code,omitempty" xml:"customer_code,omitempty" require:"true" maxLength:"50"`
	// 关联保单号,需要仓储CP做拆分计算
	PolicyNo *string `json:"policy_no,omitempty" xml:"policy_no,omitempty" maxLength:"64"`
	// 入库时间, yyyy-MM-dd HH:mm:ss，需要仓储CP做拆分计算
	//
	StockinDate *string `json:"stockin_date,omitempty" xml:"stockin_date,omitempty"`
	// 时区,仓储CP上报入库时间所属的时区
	Timezone *string `json:"timezone,omitempty" xml:"timezone,omitempty" maxLength:"16"`
}

func (s InventoryCargo) String() string {
	return tea.Prettify(s)
}

func (s InventoryCargo) GoString() string {
	return s.String()
}

func (s *InventoryCargo) SetInventoryIndex(v int64) *InventoryCargo {
	s.InventoryIndex = &v
	return s
}

func (s *InventoryCargo) SetSku(v string) *InventoryCargo {
	s.Sku = &v
	return s
}

func (s *InventoryCargo) SetCargoName(v string) *InventoryCargo {
	s.CargoName = &v
	return s
}

func (s *InventoryCargo) SetCargoWeight(v string) *InventoryCargo {
	s.CargoWeight = &v
	return s
}

func (s *InventoryCargo) SetCargoDimensions(v string) *InventoryCargo {
	s.CargoDimensions = &v
	return s
}

func (s *InventoryCargo) SetCargoWorth(v string) *InventoryCargo {
	s.CargoWorth = &v
	return s
}

func (s *InventoryCargo) SetCurrentInventoryCargoNum(v int64) *InventoryCargo {
	s.CurrentInventoryCargoNum = &v
	return s
}

func (s *InventoryCargo) SetCustomerCode(v string) *InventoryCargo {
	s.CustomerCode = &v
	return s
}

func (s *InventoryCargo) SetPolicyNo(v string) *InventoryCargo {
	s.PolicyNo = &v
	return s
}

func (s *InventoryCargo) SetStockinDate(v string) *InventoryCargo {
	s.StockinDate = &v
	return s
}

func (s *InventoryCargo) SetTimezone(v string) *InventoryCargo {
	s.Timezone = &v
	return s
}

// 签署方
type AuthParty struct {
	// 签署方名称
	SignPartyName *string `json:"sign_party_name,omitempty" xml:"sign_party_name,omitempty" require:"true"`
	// 签署方证件类型，可以填写的枚举类：IDENTIFICATION_CARD，表示身份证
	SignPartyCertType *string `json:"sign_party_cert_type,omitempty" xml:"sign_party_cert_type,omitempty" require:"true"`
	// 签署方证件号码
	SignPartyCertNum *string `json:"sign_party_cert_num,omitempty" xml:"sign_party_cert_num,omitempty" require:"true"`
	// 签署结果（必填，FINISH,FAIL,REFUSE三者选一，分别表示签署完成、失败和拒签）
	SignResult *string `json:"sign_result,omitempty" xml:"sign_result,omitempty" require:"true"`
	// 签署失败或拒签原因（失败或拒签时必填）
	SignFailReason *string `json:"sign_fail_reason,omitempty" xml:"sign_fail_reason,omitempty"`
	// 签署时间(13位毫秒时间戳)
	SignTime *string `json:"sign_time,omitempty" xml:"sign_time,omitempty" require:"true"`
}

func (s AuthParty) String() string {
	return tea.Prettify(s)
}

func (s AuthParty) GoString() string {
	return s.String()
}

func (s *AuthParty) SetSignPartyName(v string) *AuthParty {
	s.SignPartyName = &v
	return s
}

func (s *AuthParty) SetSignPartyCertType(v string) *AuthParty {
	s.SignPartyCertType = &v
	return s
}

func (s *AuthParty) SetSignPartyCertNum(v string) *AuthParty {
	s.SignPartyCertNum = &v
	return s
}

func (s *AuthParty) SetSignResult(v string) *AuthParty {
	s.SignResult = &v
	return s
}

func (s *AuthParty) SetSignFailReason(v string) *AuthParty {
	s.SignFailReason = &v
	return s
}

func (s *AuthParty) SetSignTime(v string) *AuthParty {
	s.SignTime = &v
	return s
}

// 链上hash
type TxDto struct {
	// 链上凭证
	TxCode *string `json:"tx_code,omitempty" xml:"tx_code,omitempty" require:"true"`
	// 链上存储结构对应类型
	DataType *string `json:"data_type,omitempty" xml:"data_type,omitempty" require:"true"`
}

func (s TxDto) String() string {
	return tea.Prettify(s)
}

func (s TxDto) GoString() string {
	return s.String()
}

func (s *TxDto) SetTxCode(v string) *TxDto {
	s.TxCode = &v
	return s
}

func (s *TxDto) SetDataType(v string) *TxDto {
	s.DataType = &v
	return s
}

// 接口测试
type ApiTestInfo struct {
	// 测试String
	TestString *string `json:"test_string,omitempty" xml:"test_string,omitempty" require:"true" maxLength:"100" minLength:"0"`
	// 测试Int
	TestInt *int64 `json:"test_int,omitempty" xml:"test_int,omitempty" require:"true" maximum:"100" minimum:"0"`
	// 测试Integer
	TestInteger *int64 `json:"test_integer,omitempty" xml:"test_integer,omitempty" require:"true" maximum:"100" minimum:"0"`
	// 测试Long
	TestLong *int64 `json:"test_long,omitempty" xml:"test_long,omitempty" require:"true" maximum:"100" minimum:"0"`
	// 测试Boolean
	TestBoolean *bool `json:"test_boolean,omitempty" xml:"test_boolean,omitempty" require:"true"`
	// 测试Date
	TestDate *string `json:"test_date,omitempty" xml:"test_date,omitempty" require:"true" pattern:"\\d{4}[-]\\d{1,2}[-]\\d{1,2}[T]\\d{2}:\\d{2}:\\d{2}([Z]|([\\.]\\d{1,9})?[\\+]\\d{2}[\\:]?\\d{2})"`
}

func (s ApiTestInfo) String() string {
	return tea.Prettify(s)
}

func (s ApiTestInfo) GoString() string {
	return s.String()
}

func (s *ApiTestInfo) SetTestString(v string) *ApiTestInfo {
	s.TestString = &v
	return s
}

func (s *ApiTestInfo) SetTestInt(v int64) *ApiTestInfo {
	s.TestInt = &v
	return s
}

func (s *ApiTestInfo) SetTestInteger(v int64) *ApiTestInfo {
	s.TestInteger = &v
	return s
}

func (s *ApiTestInfo) SetTestLong(v int64) *ApiTestInfo {
	s.TestLong = &v
	return s
}

func (s *ApiTestInfo) SetTestBoolean(v bool) *ApiTestInfo {
	s.TestBoolean = &v
	return s
}

func (s *ApiTestInfo) SetTestDate(v string) *ApiTestInfo {
	s.TestDate = &v
	return s
}

// 用户凭证信息
type UserIssueId struct {
	// 凭证id
	IssueId *string `json:"issue_id,omitempty" xml:"issue_id,omitempty" require:"true"`
	// 凭证余额
	BalanceAmt *string `json:"balance_amt,omitempty" xml:"balance_amt,omitempty" require:"true"`
}

func (s UserIssueId) String() string {
	return tea.Prettify(s)
}

func (s UserIssueId) GoString() string {
	return s.String()
}

func (s *UserIssueId) SetIssueId(v string) *UserIssueId {
	s.IssueId = &v
	return s
}

func (s *UserIssueId) SetBalanceAmt(v string) *UserIssueId {
	s.BalanceAmt = &v
	return s
}

// 附加条款
type MainItemAdd struct {
	// 附加条款代码
	MainItemAddCode *string `json:"main_item_add_code,omitempty" xml:"main_item_add_code,omitempty"`
	// 附加条款内容
	MainItemAddContent *string `json:"main_item_add_content,omitempty" xml:"main_item_add_content,omitempty"`
}

func (s MainItemAdd) String() string {
	return tea.Prettify(s)
}

func (s MainItemAdd) GoString() string {
	return s.String()
}

func (s *MainItemAdd) SetMainItemAddCode(v string) *MainItemAdd {
	s.MainItemAddCode = &v
	return s
}

func (s *MainItemAdd) SetMainItemAddContent(v string) *MainItemAdd {
	s.MainItemAddContent = &v
	return s
}

// 承运人责任险保险标的信息
type InsureCarrierObjectInfo struct {
	// 厂牌型号
	CpModel *string `json:"cp_model,omitempty" xml:"cp_model,omitempty" require:"true"`
	// 车架号
	FrameNo *string `json:"frame_no,omitempty" xml:"frame_no,omitempty" require:"true"`
	// 车牌号码
	LicenseNo *string `json:"license_no,omitempty" xml:"license_no,omitempty" require:"true"`
	// 吨位
	TonNage *string `json:"ton_nage,omitempty" xml:"ton_nage,omitempty" require:"true"`
	// 行驶证车主
	DrivPer *string `json:"driv_per,omitempty" xml:"driv_per,omitempty" require:"true"`
	// 运营证号
	RunNo *string `json:"run_no,omitempty" xml:"run_no,omitempty" require:"true"`
	// 运输货物
	TsCarGo *string `json:"ts_car_go,omitempty" xml:"ts_car_go,omitempty" require:"true"`
}

func (s InsureCarrierObjectInfo) String() string {
	return tea.Prettify(s)
}

func (s InsureCarrierObjectInfo) GoString() string {
	return s.String()
}

func (s *InsureCarrierObjectInfo) SetCpModel(v string) *InsureCarrierObjectInfo {
	s.CpModel = &v
	return s
}

func (s *InsureCarrierObjectInfo) SetFrameNo(v string) *InsureCarrierObjectInfo {
	s.FrameNo = &v
	return s
}

func (s *InsureCarrierObjectInfo) SetLicenseNo(v string) *InsureCarrierObjectInfo {
	s.LicenseNo = &v
	return s
}

func (s *InsureCarrierObjectInfo) SetTonNage(v string) *InsureCarrierObjectInfo {
	s.TonNage = &v
	return s
}

func (s *InsureCarrierObjectInfo) SetDrivPer(v string) *InsureCarrierObjectInfo {
	s.DrivPer = &v
	return s
}

func (s *InsureCarrierObjectInfo) SetRunNo(v string) *InsureCarrierObjectInfo {
	s.RunNo = &v
	return s
}

func (s *InsureCarrierObjectInfo) SetTsCarGo(v string) *InsureCarrierObjectInfo {
	s.TsCarGo = &v
	return s
}

// 航运订舱单号信息
type BookingNoInfo struct {
	// 订舱单唯一标识
	BookingNo *string `json:"booking_no,omitempty" xml:"booking_no,omitempty"`
	// 订舱号
	BkgNo *string `json:"bkg_no,omitempty" xml:"bkg_no,omitempty"`
}

func (s BookingNoInfo) String() string {
	return tea.Prettify(s)
}

func (s BookingNoInfo) GoString() string {
	return s.String()
}

func (s *BookingNoInfo) SetBookingNo(v string) *BookingNoInfo {
	s.BookingNo = &v
	return s
}

func (s *BookingNoInfo) SetBkgNo(v string) *BookingNoInfo {
	s.BkgNo = &v
	return s
}

// 入库货物
type StockinCargo struct {
	// 入库序号，在同一次入库请求中，入库序号保持不重复，不能小于0
	StockinIndex *int64 `json:"stockin_index,omitempty" xml:"stockin_index,omitempty" require:"true"`
	// sku品名
	//
	Sku *string `json:"sku,omitempty" xml:"sku,omitempty" require:"true" maxLength:"200"`
	// 商品名称
	CargoName *string `json:"cargo_name,omitempty" xml:"cargo_name,omitempty" maxLength:"200"`
	// 商品单品重量(kg)
	CargoWeight *string `json:"cargo_weight,omitempty" xml:"cargo_weight,omitempty" maxLength:"50"`
	// 商品外扩长宽高(cm)
	CargoDimensions *string `json:"cargo_dimensions,omitempty" xml:"cargo_dimensions,omitempty" maxLength:"200"`
	// 商品单品货物价值(元),，最多支持2位小数
	CargoWorth *string `json:"cargo_worth,omitempty" xml:"cargo_worth,omitempty" maxLength:"30"`
	// 箱号
	ContainerNo *string `json:"container_no,omitempty" xml:"container_no,omitempty" maxLength:"50"`
	// 实际入库件数
	ActualStockinNum *int64 `json:"actual_stockin_num,omitempty" xml:"actual_stockin_num,omitempty" require:"true"`
}

func (s StockinCargo) String() string {
	return tea.Prettify(s)
}

func (s StockinCargo) GoString() string {
	return s.String()
}

func (s *StockinCargo) SetStockinIndex(v int64) *StockinCargo {
	s.StockinIndex = &v
	return s
}

func (s *StockinCargo) SetSku(v string) *StockinCargo {
	s.Sku = &v
	return s
}

func (s *StockinCargo) SetCargoName(v string) *StockinCargo {
	s.CargoName = &v
	return s
}

func (s *StockinCargo) SetCargoWeight(v string) *StockinCargo {
	s.CargoWeight = &v
	return s
}

func (s *StockinCargo) SetCargoDimensions(v string) *StockinCargo {
	s.CargoDimensions = &v
	return s
}

func (s *StockinCargo) SetCargoWorth(v string) *StockinCargo {
	s.CargoWorth = &v
	return s
}

func (s *StockinCargo) SetContainerNo(v string) *StockinCargo {
	s.ContainerNo = &v
	return s
}

func (s *StockinCargo) SetActualStockinNum(v int64) *StockinCargo {
	s.ActualStockinNum = &v
	return s
}

// 提单货物数据
type MasterBlGoodsDto struct {
	// 唛头
	Marks *string `json:"marks,omitempty" xml:"marks,omitempty"`
	// 货物
	Goods *string `json:"goods,omitempty" xml:"goods,omitempty" require:"true"`
	// 货物类型
	GoodsType *string `json:"goods_type,omitempty" xml:"goods_type,omitempty" require:"true"`
	// 包装类型
	PackageType *string `json:"package_type,omitempty" xml:"package_type,omitempty"`
	// 委托件数
	Number *string `json:"number,omitempty" xml:"number,omitempty" require:"true"`
	// 委托重量
	Weight *string `json:"weight,omitempty" xml:"weight,omitempty" require:"true"`
	// 委托体积
	Volume *string `json:"volume,omitempty" xml:"volume,omitempty" require:"true"`
}

func (s MasterBlGoodsDto) String() string {
	return tea.Prettify(s)
}

func (s MasterBlGoodsDto) GoString() string {
	return s.String()
}

func (s *MasterBlGoodsDto) SetMarks(v string) *MasterBlGoodsDto {
	s.Marks = &v
	return s
}

func (s *MasterBlGoodsDto) SetGoods(v string) *MasterBlGoodsDto {
	s.Goods = &v
	return s
}

func (s *MasterBlGoodsDto) SetGoodsType(v string) *MasterBlGoodsDto {
	s.GoodsType = &v
	return s
}

func (s *MasterBlGoodsDto) SetPackageType(v string) *MasterBlGoodsDto {
	s.PackageType = &v
	return s
}

func (s *MasterBlGoodsDto) SetNumber(v string) *MasterBlGoodsDto {
	s.Number = &v
	return s
}

func (s *MasterBlGoodsDto) SetWeight(v string) *MasterBlGoodsDto {
	s.Weight = &v
	return s
}

func (s *MasterBlGoodsDto) SetVolume(v string) *MasterBlGoodsDto {
	s.Volume = &v
	return s
}

// 资费项发票
type ReceiptTariffInvoiceParam struct {
	// 操作动作,为空为新增或更新，UPDATE为更新，DELETE为删除
	Action *string `json:"action,omitempty" xml:"action,omitempty"`
	// 发票金额 业务必填
	InvoiceAmount *string `json:"invoice_amount,omitempty" xml:"invoice_amount,omitempty"`
	// 资费项金额 业务必填
	ReceiptTariffAmount *string `json:"receipt_tariff_amount,omitempty" xml:"receipt_tariff_amount,omitempty"`
	// 资费单据编号 业务必填
	ReceiptTariffCode *string `json:"receipt_tariff_code,omitempty" xml:"receipt_tariff_code,omitempty"`
	// 资费项发票code
	ReceiptTariffInvoiceCode *string `json:"receipt_tariff_invoice_code,omitempty" xml:"receipt_tariff_invoice_code,omitempty" require:"true"`
}

func (s ReceiptTariffInvoiceParam) String() string {
	return tea.Prettify(s)
}

func (s ReceiptTariffInvoiceParam) GoString() string {
	return s.String()
}

func (s *ReceiptTariffInvoiceParam) SetAction(v string) *ReceiptTariffInvoiceParam {
	s.Action = &v
	return s
}

func (s *ReceiptTariffInvoiceParam) SetInvoiceAmount(v string) *ReceiptTariffInvoiceParam {
	s.InvoiceAmount = &v
	return s
}

func (s *ReceiptTariffInvoiceParam) SetReceiptTariffAmount(v string) *ReceiptTariffInvoiceParam {
	s.ReceiptTariffAmount = &v
	return s
}

func (s *ReceiptTariffInvoiceParam) SetReceiptTariffCode(v string) *ReceiptTariffInvoiceParam {
	s.ReceiptTariffCode = &v
	return s
}

func (s *ReceiptTariffInvoiceParam) SetReceiptTariffInvoiceCode(v string) *ReceiptTariffInvoiceParam {
	s.ReceiptTariffInvoiceCode = &v
	return s
}

// 上传booking信息
type UploadOrderBooking struct {
	// 订舱单号
	BookingNo *string `json:"booking_no,omitempty" xml:"booking_no,omitempty" require:"true"`
	// 集装箱号  json字符串上传
	ContainerNos *string `json:"container_nos,omitempty" xml:"container_nos,omitempty" require:"true"`
}

func (s UploadOrderBooking) String() string {
	return tea.Prettify(s)
}

func (s UploadOrderBooking) GoString() string {
	return s.String()
}

func (s *UploadOrderBooking) SetBookingNo(v string) *UploadOrderBooking {
	s.BookingNo = &v
	return s
}

func (s *UploadOrderBooking) SetContainerNos(v string) *UploadOrderBooking {
	s.ContainerNos = &v
	return s
}

// 轨迹核验结果
type TrackCheckResult struct {
	// 轨迹核验状态code
	TrackCheckStatus *string `json:"track_check_status,omitempty" xml:"track_check_status,omitempty"`
	// 轨迹核验结果描述
	TrackCheckStatusMsg *string `json:"track_check_status_msg,omitempty" xml:"track_check_status_msg,omitempty"`
}

func (s TrackCheckResult) String() string {
	return tea.Prettify(s)
}

func (s TrackCheckResult) GoString() string {
	return s.String()
}

func (s *TrackCheckResult) SetTrackCheckStatus(v string) *TrackCheckResult {
	s.TrackCheckStatus = &v
	return s
}

func (s *TrackCheckResult) SetTrackCheckStatusMsg(v string) *TrackCheckResult {
	s.TrackCheckStatusMsg = &v
	return s
}

// 货物
type ContainerGoodsParam struct {
	// 预计备货时间
	CargoReadyDate *int64 `json:"cargo_ready_date,omitempty" xml:"cargo_ready_date,omitempty"`
	// 危险品页号
	DgPageNo *string `json:"dg_page_no,omitempty" xml:"dg_page_no,omitempty"`
	// 危险品级别
	DgType *string `json:"dg_type,omitempty" xml:"dg_type,omitempty"`
	// 危险品闪点
	FlashPoint *string `json:"flash_point,omitempty" xml:"flash_point,omitempty"`
	// 货物
	Goods *string `json:"goods,omitempty" xml:"goods,omitempty"`
	// 货物中文名
	GoodsCn *string `json:"goods_cn,omitempty" xml:"goods_cn,omitempty"`
	// 货物类型 业务必填
	GoodsType *string `json:"goods_type,omitempty" xml:"goods_type,omitempty"`
	// HS CODE
	HsCodes []*string `json:"hs_codes,omitempty" xml:"hs_codes,omitempty" type:"Repeated"`
	// 唛头
	Marks *string `json:"marks,omitempty" xml:"marks,omitempty"`
	// 件数 业务必填
	Number *string `json:"number,omitempty" xml:"number,omitempty"`
	// 包装类型
	PackageType *string `json:"package_type,omitempty" xml:"package_type,omitempty"`
	// 实际件数
	RealNumber *string `json:"real_number,omitempty" xml:"real_number,omitempty"`
	// 实际体积
	RealVolume *string `json:"real_volume,omitempty" xml:"real_volume,omitempty"`
	// 实际重量
	RealWeight *string `json:"real_weight,omitempty" xml:"real_weight,omitempty"`
	// 危险品联合国编号
	UnNo *string `json:"un_no,omitempty" xml:"un_no,omitempty"`
	// 体积 业务必填
	Volume *string `json:"volume,omitempty" xml:"volume,omitempty"`
	// 毛重 业务必填
	Weight *string `json:"weight,omitempty" xml:"weight,omitempty"`
}

func (s ContainerGoodsParam) String() string {
	return tea.Prettify(s)
}

func (s ContainerGoodsParam) GoString() string {
	return s.String()
}

func (s *ContainerGoodsParam) SetCargoReadyDate(v int64) *ContainerGoodsParam {
	s.CargoReadyDate = &v
	return s
}

func (s *ContainerGoodsParam) SetDgPageNo(v string) *ContainerGoodsParam {
	s.DgPageNo = &v
	return s
}

func (s *ContainerGoodsParam) SetDgType(v string) *ContainerGoodsParam {
	s.DgType = &v
	return s
}

func (s *ContainerGoodsParam) SetFlashPoint(v string) *ContainerGoodsParam {
	s.FlashPoint = &v
	return s
}

func (s *ContainerGoodsParam) SetGoods(v string) *ContainerGoodsParam {
	s.Goods = &v
	return s
}

func (s *ContainerGoodsParam) SetGoodsCn(v string) *ContainerGoodsParam {
	s.GoodsCn = &v
	return s
}

func (s *ContainerGoodsParam) SetGoodsType(v string) *ContainerGoodsParam {
	s.GoodsType = &v
	return s
}

func (s *ContainerGoodsParam) SetHsCodes(v []*string) *ContainerGoodsParam {
	s.HsCodes = v
	return s
}

func (s *ContainerGoodsParam) SetMarks(v string) *ContainerGoodsParam {
	s.Marks = &v
	return s
}

func (s *ContainerGoodsParam) SetNumber(v string) *ContainerGoodsParam {
	s.Number = &v
	return s
}

func (s *ContainerGoodsParam) SetPackageType(v string) *ContainerGoodsParam {
	s.PackageType = &v
	return s
}

func (s *ContainerGoodsParam) SetRealNumber(v string) *ContainerGoodsParam {
	s.RealNumber = &v
	return s
}

func (s *ContainerGoodsParam) SetRealVolume(v string) *ContainerGoodsParam {
	s.RealVolume = &v
	return s
}

func (s *ContainerGoodsParam) SetRealWeight(v string) *ContainerGoodsParam {
	s.RealWeight = &v
	return s
}

func (s *ContainerGoodsParam) SetUnNo(v string) *ContainerGoodsParam {
	s.UnNo = &v
	return s
}

func (s *ContainerGoodsParam) SetVolume(v string) *ContainerGoodsParam {
	s.Volume = &v
	return s
}

func (s *ContainerGoodsParam) SetWeight(v string) *ContainerGoodsParam {
	s.Weight = &v
	return s
}

// 电子提单变更状态明细
type EblStatusDetail struct {
	// 当前提单状态
	CurrentEblStatus *string `json:"current_ebl_status,omitempty" xml:"current_ebl_status,omitempty" require:"true"`
	// 电子提单编号
	EblNo *string `json:"ebl_no,omitempty" xml:"ebl_no,omitempty" require:"true"`
	// 下一个提单状态
	NextEblStatus *string `json:"next_ebl_status,omitempty" xml:"next_ebl_status,omitempty" require:"true"`
}

func (s EblStatusDetail) String() string {
	return tea.Prettify(s)
}

func (s EblStatusDetail) GoString() string {
	return s.String()
}

func (s *EblStatusDetail) SetCurrentEblStatus(v string) *EblStatusDetail {
	s.CurrentEblStatus = &v
	return s
}

func (s *EblStatusDetail) SetEblNo(v string) *EblStatusDetail {
	s.EblNo = &v
	return s
}

func (s *EblStatusDetail) SetNextEblStatus(v string) *EblStatusDetail {
	s.NextEblStatus = &v
	return s
}

// 应付资费项
type PayTariffInfo struct {
	// 托运单号 [业务必填]
	OrderNo *string `json:"order_no,omitempty" xml:"order_no,omitempty"`
	// 应付资费项code [业务必填]
	//
	//
	PayTariffCode *string `json:"pay_tariff_code,omitempty" xml:"pay_tariff_code,omitempty"`
	// 应付资费项项目 [业务必填]
	//
	//
	PayTariffProject *string `json:"pay_tariff_project,omitempty" xml:"pay_tariff_project,omitempty"`
	// 资费项中文描述 [业务必填]
	//
	//
	PayTariffDesc *string `json:"pay_tariff_desc,omitempty" xml:"pay_tariff_desc,omitempty"`
	// 币种 [业务必填]
	Currency *string `json:"currency,omitempty" xml:"currency,omitempty"`
	// 资费项含税价 [业务必填]
	//
	//
	PriceIncludingTax *string `json:"price_including_tax,omitempty" xml:"price_including_tax,omitempty"`
	// 订舱单唯一性标识 [业务必填]
	BookingNo *string `json:"booking_no,omitempty" xml:"booking_no,omitempty"`
	// 订舱单 [业务必填]
	BkgNo *string `json:"bkg_no,omitempty" xml:"bkg_no,omitempty"`
}

func (s PayTariffInfo) String() string {
	return tea.Prettify(s)
}

func (s PayTariffInfo) GoString() string {
	return s.String()
}

func (s *PayTariffInfo) SetOrderNo(v string) *PayTariffInfo {
	s.OrderNo = &v
	return s
}

func (s *PayTariffInfo) SetPayTariffCode(v string) *PayTariffInfo {
	s.PayTariffCode = &v
	return s
}

func (s *PayTariffInfo) SetPayTariffProject(v string) *PayTariffInfo {
	s.PayTariffProject = &v
	return s
}

func (s *PayTariffInfo) SetPayTariffDesc(v string) *PayTariffInfo {
	s.PayTariffDesc = &v
	return s
}

func (s *PayTariffInfo) SetCurrency(v string) *PayTariffInfo {
	s.Currency = &v
	return s
}

func (s *PayTariffInfo) SetPriceIncludingTax(v string) *PayTariffInfo {
	s.PriceIncludingTax = &v
	return s
}

func (s *PayTariffInfo) SetBookingNo(v string) *PayTariffInfo {
	s.BookingNo = &v
	return s
}

func (s *PayTariffInfo) SetBkgNo(v string) *PayTariffInfo {
	s.BkgNo = &v
	return s
}

// FinishWaybillOrderReq
type FinishWaybillOrderReq struct {
	// 运费
	AllFreight *string `json:"all_freight,omitempty" xml:"all_freight,omitempty"`
	// 回单押金
	BackFee *string `json:"back_fee,omitempty" xml:"back_fee,omitempty"`
	// 货主支付运费金额
	ConsignorFreightAmount *string `json:"consignor_freight_amount,omitempty" xml:"consignor_freight_amount,omitempty"`
	// 运费增项
	FreightIncr *string `json:"freight_incr,omitempty" xml:"freight_incr,omitempty"`
	// 运费扣减
	LossFee *string `json:"loss_fee,omitempty" xml:"loss_fee,omitempty"`
	// 平台did
	PlatformDid *string `json:"platform_did,omitempty" xml:"platform_did,omitempty" require:"true"`
	// 运单id
	TaxWaybillId *string `json:"tax_waybill_id,omitempty" xml:"tax_waybill_id,omitempty" require:"true"`
}

func (s FinishWaybillOrderReq) String() string {
	return tea.Prettify(s)
}

func (s FinishWaybillOrderReq) GoString() string {
	return s.String()
}

func (s *FinishWaybillOrderReq) SetAllFreight(v string) *FinishWaybillOrderReq {
	s.AllFreight = &v
	return s
}

func (s *FinishWaybillOrderReq) SetBackFee(v string) *FinishWaybillOrderReq {
	s.BackFee = &v
	return s
}

func (s *FinishWaybillOrderReq) SetConsignorFreightAmount(v string) *FinishWaybillOrderReq {
	s.ConsignorFreightAmount = &v
	return s
}

func (s *FinishWaybillOrderReq) SetFreightIncr(v string) *FinishWaybillOrderReq {
	s.FreightIncr = &v
	return s
}

func (s *FinishWaybillOrderReq) SetLossFee(v string) *FinishWaybillOrderReq {
	s.LossFee = &v
	return s
}

func (s *FinishWaybillOrderReq) SetPlatformDid(v string) *FinishWaybillOrderReq {
	s.PlatformDid = &v
	return s
}

func (s *FinishWaybillOrderReq) SetTaxWaybillId(v string) *FinishWaybillOrderReq {
	s.TaxWaybillId = &v
	return s
}

// 运单号-司机运费
type WaybillAmount struct {
	// 运单金额（2位小数）
	Amount *string `json:"amount,omitempty" xml:"amount,omitempty" require:"true"`
	// 运单号
	WaybillId *string `json:"waybill_id,omitempty" xml:"waybill_id,omitempty" require:"true"`
}

func (s WaybillAmount) String() string {
	return tea.Prettify(s)
}

func (s WaybillAmount) GoString() string {
	return s.String()
}

func (s *WaybillAmount) SetAmount(v string) *WaybillAmount {
	s.Amount = &v
	return s
}

func (s *WaybillAmount) SetWaybillId(v string) *WaybillAmount {
	s.WaybillId = &v
	return s
}

// 电子回单查询凭证数据
type ScpTicketIssueDataParam struct {
	// 凭证id
	IssueId *string `json:"issue_id,omitempty" xml:"issue_id,omitempty" require:"true"`
	// 凭证对应的司机/货主的did
	Did *string `json:"did,omitempty" xml:"did,omitempty" require:"true"`
}

func (s ScpTicketIssueDataParam) String() string {
	return tea.Prettify(s)
}

func (s ScpTicketIssueDataParam) GoString() string {
	return s.String()
}

func (s *ScpTicketIssueDataParam) SetIssueId(v string) *ScpTicketIssueDataParam {
	s.IssueId = &v
	return s
}

func (s *ScpTicketIssueDataParam) SetDid(v string) *ScpTicketIssueDataParam {
	s.Did = &v
	return s
}

// saas模式发行信息
type SaasIssueApplyInfo struct {
	// 货源订单
	CargoOrder *string `json:"cargo_order,omitempty" xml:"cargo_order,omitempty"`
	// 合同号
	ContractId *string `json:"contract_id,omitempty" xml:"contract_id,omitempty"`
	// 全局唯一业务单号
	OutBizNo *string `json:"out_biz_no,omitempty" xml:"out_biz_no,omitempty" require:"true"`
	// 支付单号
	PayOrder *string `json:"pay_order,omitempty" xml:"pay_order,omitempty" require:"true"`
	// 运单号
	WaybillId *string `json:"waybill_id,omitempty" xml:"waybill_id,omitempty" require:"true"`
	// 司机did
	DriverDid *string `json:"driver_did,omitempty" xml:"driver_did,omitempty" require:"true"`
	// 发行费
	Freight *string `json:"freight,omitempty" xml:"freight,omitempty" require:"true"`
	// 到期时间戳
	ExpireDate *string `json:"expire_date,omitempty" xml:"expire_date,omitempty" require:"true"`
}

func (s SaasIssueApplyInfo) String() string {
	return tea.Prettify(s)
}

func (s SaasIssueApplyInfo) GoString() string {
	return s.String()
}

func (s *SaasIssueApplyInfo) SetCargoOrder(v string) *SaasIssueApplyInfo {
	s.CargoOrder = &v
	return s
}

func (s *SaasIssueApplyInfo) SetContractId(v string) *SaasIssueApplyInfo {
	s.ContractId = &v
	return s
}

func (s *SaasIssueApplyInfo) SetOutBizNo(v string) *SaasIssueApplyInfo {
	s.OutBizNo = &v
	return s
}

func (s *SaasIssueApplyInfo) SetPayOrder(v string) *SaasIssueApplyInfo {
	s.PayOrder = &v
	return s
}

func (s *SaasIssueApplyInfo) SetWaybillId(v string) *SaasIssueApplyInfo {
	s.WaybillId = &v
	return s
}

func (s *SaasIssueApplyInfo) SetDriverDid(v string) *SaasIssueApplyInfo {
	s.DriverDid = &v
	return s
}

func (s *SaasIssueApplyInfo) SetFreight(v string) *SaasIssueApplyInfo {
	s.Freight = &v
	return s
}

func (s *SaasIssueApplyInfo) SetExpireDate(v string) *SaasIssueApplyInfo {
	s.ExpireDate = &v
	return s
}

// 应收资费项
type ReceiptTariffInfo struct {
	// 托运单号 [业务必填]
	OrderNo *string `json:"order_no,omitempty" xml:"order_no,omitempty"`
	// 应收资费项code [业务必填]
	//
	//
	ReceiptTariffCode *string `json:"receipt_tariff_code,omitempty" xml:"receipt_tariff_code,omitempty"`
	// 应收资费项项目 [业务必填]
	ReceiptTariffProject *string `json:"receipt_tariff_project,omitempty" xml:"receipt_tariff_project,omitempty"`
	// 资费项中文描述 [业务必填]
	//
	//
	ReceiptTariffDesc *string `json:"receipt_tariff_desc,omitempty" xml:"receipt_tariff_desc,omitempty"`
	// 币种 [业务必填]
	Currency *string `json:"currency,omitempty" xml:"currency,omitempty"`
	// 资费项含税价 [业务必填]
	//
	//
	PriceIncludingTax *string `json:"price_including_tax,omitempty" xml:"price_including_tax,omitempty"`
	// 订舱单唯一标识 [业务必填]
	BookingNo *string `json:"booking_no,omitempty" xml:"booking_no,omitempty"`
	// 订舱号 [业务必填]
	BkgNo *string `json:"bkg_no,omitempty" xml:"bkg_no,omitempty"`
}

func (s ReceiptTariffInfo) String() string {
	return tea.Prettify(s)
}

func (s ReceiptTariffInfo) GoString() string {
	return s.String()
}

func (s *ReceiptTariffInfo) SetOrderNo(v string) *ReceiptTariffInfo {
	s.OrderNo = &v
	return s
}

func (s *ReceiptTariffInfo) SetReceiptTariffCode(v string) *ReceiptTariffInfo {
	s.ReceiptTariffCode = &v
	return s
}

func (s *ReceiptTariffInfo) SetReceiptTariffProject(v string) *ReceiptTariffInfo {
	s.ReceiptTariffProject = &v
	return s
}

func (s *ReceiptTariffInfo) SetReceiptTariffDesc(v string) *ReceiptTariffInfo {
	s.ReceiptTariffDesc = &v
	return s
}

func (s *ReceiptTariffInfo) SetCurrency(v string) *ReceiptTariffInfo {
	s.Currency = &v
	return s
}

func (s *ReceiptTariffInfo) SetPriceIncludingTax(v string) *ReceiptTariffInfo {
	s.PriceIncludingTax = &v
	return s
}

func (s *ReceiptTariffInfo) SetBookingNo(v string) *ReceiptTariffInfo {
	s.BookingNo = &v
	return s
}

func (s *ReceiptTariffInfo) SetBkgNo(v string) *ReceiptTariffInfo {
	s.BkgNo = &v
	return s
}

// 货源单号-货主运费
type CargoAmount struct {
	// 货运单对应金额（2位小数）
	Amount *string `json:"amount,omitempty" xml:"amount,omitempty" require:"true"`
	// 货源单号
	CargoOrder *string `json:"cargo_order,omitempty" xml:"cargo_order,omitempty" require:"true"`
}

func (s CargoAmount) String() string {
	return tea.Prettify(s)
}

func (s CargoAmount) GoString() string {
	return s.String()
}

func (s *CargoAmount) SetAmount(v string) *CargoAmount {
	s.Amount = &v
	return s
}

func (s *CargoAmount) SetCargoOrder(v string) *CargoAmount {
	s.CargoOrder = &v
	return s
}

// 凭证开立申请信息
type IssueApplyInfo struct {
	// 货源订单
	CargoOrder *string `json:"cargo_order,omitempty" xml:"cargo_order,omitempty"`
	// 合同号（预留）
	ContractId *string `json:"contract_id,omitempty" xml:"contract_id,omitempty"`
	// 凭证到期时间
	ExpireDate *string `json:"expire_date,omitempty" xml:"expire_date,omitempty" require:"true"`
	// 支付单运费，运费最多精确到小数点后2位
	Freight *string `json:"freight,omitempty" xml:"freight,omitempty" require:"true"`
	// 全局唯一业务号
	OutBizNo *string `json:"out_biz_no,omitempty" xml:"out_biz_no,omitempty" require:"true"`
	// 支付订单
	PayOrder *string `json:"pay_order,omitempty" xml:"pay_order,omitempty" require:"true"`
	// 运单id
	WaybillId *string `json:"waybill_id,omitempty" xml:"waybill_id,omitempty" require:"true"`
	// 司机did
	DriverDid *string `json:"driver_did,omitempty" xml:"driver_did,omitempty"`
}

func (s IssueApplyInfo) String() string {
	return tea.Prettify(s)
}

func (s IssueApplyInfo) GoString() string {
	return s.String()
}

func (s *IssueApplyInfo) SetCargoOrder(v string) *IssueApplyInfo {
	s.CargoOrder = &v
	return s
}

func (s *IssueApplyInfo) SetContractId(v string) *IssueApplyInfo {
	s.ContractId = &v
	return s
}

func (s *IssueApplyInfo) SetExpireDate(v string) *IssueApplyInfo {
	s.ExpireDate = &v
	return s
}

func (s *IssueApplyInfo) SetFreight(v string) *IssueApplyInfo {
	s.Freight = &v
	return s
}

func (s *IssueApplyInfo) SetOutBizNo(v string) *IssueApplyInfo {
	s.OutBizNo = &v
	return s
}

func (s *IssueApplyInfo) SetPayOrder(v string) *IssueApplyInfo {
	s.PayOrder = &v
	return s
}

func (s *IssueApplyInfo) SetWaybillId(v string) *IssueApplyInfo {
	s.WaybillId = &v
	return s
}

func (s *IssueApplyInfo) SetDriverDid(v string) *IssueApplyInfo {
	s.DriverDid = &v
	return s
}

// 物流轨迹位置
type LogisticLocation struct {
	// 结构化地址信息,规则遵循：国家、省份、城市、区县、城镇、乡村、街道、门牌号码、屋邨、大厦
	Address *string `json:"address,omitempty" xml:"address,omitempty"`
	// 行政区划代码
	CityCode *string `json:"city_code,omitempty" xml:"city_code,omitempty"`
	// 纬度
	//
	Lat *string `json:"lat,omitempty" xml:"lat,omitempty" require:"true"`
	// 经度
	Lon *string `json:"lon,omitempty" xml:"lon,omitempty" require:"true"`
	// 轨迹时间戳
	TrackTime *int64 `json:"track_time,omitempty" xml:"track_time,omitempty" require:"true"`
}

func (s LogisticLocation) String() string {
	return tea.Prettify(s)
}

func (s LogisticLocation) GoString() string {
	return s.String()
}

func (s *LogisticLocation) SetAddress(v string) *LogisticLocation {
	s.Address = &v
	return s
}

func (s *LogisticLocation) SetCityCode(v string) *LogisticLocation {
	s.CityCode = &v
	return s
}

func (s *LogisticLocation) SetLat(v string) *LogisticLocation {
	s.Lat = &v
	return s
}

func (s *LogisticLocation) SetLon(v string) *LogisticLocation {
	s.Lon = &v
	return s
}

func (s *LogisticLocation) SetTrackTime(v int64) *LogisticLocation {
	s.TrackTime = &v
	return s
}

// 集装箱列表
type MasterBlContainerParam struct {
	// 操作动作,为空为新增或更新，UPDATE为更新，DELETE为删除
	Action *string `json:"action,omitempty" xml:"action,omitempty"`
	// 集装箱ID
	ContainerId *string `json:"container_id,omitempty" xml:"container_id,omitempty" require:"true"`
	// 箱号
	ContainerNo *string `json:"container_no,omitempty" xml:"container_no,omitempty"`
}

func (s MasterBlContainerParam) String() string {
	return tea.Prettify(s)
}

func (s MasterBlContainerParam) GoString() string {
	return s.String()
}

func (s *MasterBlContainerParam) SetAction(v string) *MasterBlContainerParam {
	s.Action = &v
	return s
}

func (s *MasterBlContainerParam) SetContainerId(v string) *MasterBlContainerParam {
	s.ContainerId = &v
	return s
}

func (s *MasterBlContainerParam) SetContainerNo(v string) *MasterBlContainerParam {
	s.ContainerNo = &v
	return s
}

// 订舱单
type HouseBlBookingParam struct {
	// 操作动作,为空为新增或更新，UPDATE为更新，DELETE为删除
	Action *string `json:"action,omitempty" xml:"action,omitempty"`
	// 订舱单号
	BookingNo *string `json:"booking_no,omitempty" xml:"booking_no,omitempty" require:"true"`
}

func (s HouseBlBookingParam) String() string {
	return tea.Prettify(s)
}

func (s HouseBlBookingParam) GoString() string {
	return s.String()
}

func (s *HouseBlBookingParam) SetAction(v string) *HouseBlBookingParam {
	s.Action = &v
	return s
}

func (s *HouseBlBookingParam) SetBookingNo(v string) *HouseBlBookingParam {
	s.BookingNo = &v
	return s
}

// 上传订单总金额
type UploadOrderAmount struct {
	// 币种
	Currency *string `json:"currency,omitempty" xml:"currency,omitempty" require:"true"`
	// 总金额
	TotalAmount *string `json:"total_amount,omitempty" xml:"total_amount,omitempty" require:"true"`
}

func (s UploadOrderAmount) String() string {
	return tea.Prettify(s)
}

func (s UploadOrderAmount) GoString() string {
	return s.String()
}

func (s *UploadOrderAmount) SetCurrency(v string) *UploadOrderAmount {
	s.Currency = &v
	return s
}

func (s *UploadOrderAmount) SetTotalAmount(v string) *UploadOrderAmount {
	s.TotalAmount = &v
	return s
}

// 货物列表
type MasterBlGoodsParam struct {
	// 预计备货时间
	CargoReadyDate *int64 `json:"cargo_ready_date,omitempty" xml:"cargo_ready_date,omitempty"`
	// 危险品页号
	DgPageNo *string `json:"dg_page_no,omitempty" xml:"dg_page_no,omitempty"`
	// 危险品级别
	DgType *string `json:"dg_type,omitempty" xml:"dg_type,omitempty"`
	// 危险品闪点
	FlashPoint *string `json:"flash_point,omitempty" xml:"flash_point,omitempty"`
	// 货物名称 业务必填
	Goods *string `json:"goods,omitempty" xml:"goods,omitempty"`
	// 货物中文名
	GoodsCn *string `json:"goods_cn,omitempty" xml:"goods_cn,omitempty"`
	// 货物类型 业务必填
	GoodsType *string `json:"goods_type,omitempty" xml:"goods_type,omitempty"`
	// 毛重 业务必填
	GrossWeight *string `json:"gross_weight,omitempty" xml:"gross_weight,omitempty"`
	// HS CODE
	HsCodes []*string `json:"hs_codes,omitempty" xml:"hs_codes,omitempty" type:"Repeated"`
	// 唛头
	Marks *string `json:"marks,omitempty" xml:"marks,omitempty"`
	// 件数 业务必填
	Number *string `json:"number,omitempty" xml:"number,omitempty"`
	// 包装类型
	PackageType *string `json:"package_type,omitempty" xml:"package_type,omitempty"`
	// 实际件数
	RealNumber *string `json:"real_number,omitempty" xml:"real_number,omitempty"`
	// 实际体积
	RealVolume *string `json:"real_volume,omitempty" xml:"real_volume,omitempty"`
	// 实际重量
	RealWeight *string `json:"real_weight,omitempty" xml:"real_weight,omitempty"`
	// 危险品联合国编号
	UnNo *string `json:"un_no,omitempty" xml:"un_no,omitempty"`
	// 委托体积 业务必填
	Volume *string `json:"volume,omitempty" xml:"volume,omitempty"`
	// 委托重量 业务必填
	Weight *string `json:"weight,omitempty" xml:"weight,omitempty"`
}

func (s MasterBlGoodsParam) String() string {
	return tea.Prettify(s)
}

func (s MasterBlGoodsParam) GoString() string {
	return s.String()
}

func (s *MasterBlGoodsParam) SetCargoReadyDate(v int64) *MasterBlGoodsParam {
	s.CargoReadyDate = &v
	return s
}

func (s *MasterBlGoodsParam) SetDgPageNo(v string) *MasterBlGoodsParam {
	s.DgPageNo = &v
	return s
}

func (s *MasterBlGoodsParam) SetDgType(v string) *MasterBlGoodsParam {
	s.DgType = &v
	return s
}

func (s *MasterBlGoodsParam) SetFlashPoint(v string) *MasterBlGoodsParam {
	s.FlashPoint = &v
	return s
}

func (s *MasterBlGoodsParam) SetGoods(v string) *MasterBlGoodsParam {
	s.Goods = &v
	return s
}

func (s *MasterBlGoodsParam) SetGoodsCn(v string) *MasterBlGoodsParam {
	s.GoodsCn = &v
	return s
}

func (s *MasterBlGoodsParam) SetGoodsType(v string) *MasterBlGoodsParam {
	s.GoodsType = &v
	return s
}

func (s *MasterBlGoodsParam) SetGrossWeight(v string) *MasterBlGoodsParam {
	s.GrossWeight = &v
	return s
}

func (s *MasterBlGoodsParam) SetHsCodes(v []*string) *MasterBlGoodsParam {
	s.HsCodes = v
	return s
}

func (s *MasterBlGoodsParam) SetMarks(v string) *MasterBlGoodsParam {
	s.Marks = &v
	return s
}

func (s *MasterBlGoodsParam) SetNumber(v string) *MasterBlGoodsParam {
	s.Number = &v
	return s
}

func (s *MasterBlGoodsParam) SetPackageType(v string) *MasterBlGoodsParam {
	s.PackageType = &v
	return s
}

func (s *MasterBlGoodsParam) SetRealNumber(v string) *MasterBlGoodsParam {
	s.RealNumber = &v
	return s
}

func (s *MasterBlGoodsParam) SetRealVolume(v string) *MasterBlGoodsParam {
	s.RealVolume = &v
	return s
}

func (s *MasterBlGoodsParam) SetRealWeight(v string) *MasterBlGoodsParam {
	s.RealWeight = &v
	return s
}

func (s *MasterBlGoodsParam) SetUnNo(v string) *MasterBlGoodsParam {
	s.UnNo = &v
	return s
}

func (s *MasterBlGoodsParam) SetVolume(v string) *MasterBlGoodsParam {
	s.Volume = &v
	return s
}

func (s *MasterBlGoodsParam) SetWeight(v string) *MasterBlGoodsParam {
	s.Weight = &v
	return s
}

// A+模式发行信息
type IssueApplyInfoPlus struct {
	// 订单中的BookingNo，以英文逗号分割
	BookingNo *string `json:"booking_no,omitempty" xml:"booking_no,omitempty" require:"true"`
	// 船公司did
	CarrierDid *string `json:"carrier_did,omitempty" xml:"carrier_did,omitempty" require:"true"`
	// BookingNo中的箱号，以英文逗号分割
	ContainerNo *string `json:"container_no,omitempty" xml:"container_no,omitempty" require:"true"`
	// 到期时间戳
	ExpireDate *string `json:"expire_date,omitempty" xml:"expire_date,omitempty" require:"true"`
	// 发行金额，精确到小数点后2位
	IssueAmt *string `json:"issue_amt,omitempty" xml:"issue_amt,omitempty" require:"true"`
	// 全局唯一业务号
	OutBizNo *string `json:"out_biz_no,omitempty" xml:"out_biz_no,omitempty" require:"true"`
	// 支付单号
	OutOrderNo *string `json:"out_order_no,omitempty" xml:"out_order_no,omitempty" require:"true"`
	// 运单订单id
	WaybillId *string `json:"waybill_id,omitempty" xml:"waybill_id,omitempty" require:"true"`
}

func (s IssueApplyInfoPlus) String() string {
	return tea.Prettify(s)
}

func (s IssueApplyInfoPlus) GoString() string {
	return s.String()
}

func (s *IssueApplyInfoPlus) SetBookingNo(v string) *IssueApplyInfoPlus {
	s.BookingNo = &v
	return s
}

func (s *IssueApplyInfoPlus) SetCarrierDid(v string) *IssueApplyInfoPlus {
	s.CarrierDid = &v
	return s
}

func (s *IssueApplyInfoPlus) SetContainerNo(v string) *IssueApplyInfoPlus {
	s.ContainerNo = &v
	return s
}

func (s *IssueApplyInfoPlus) SetExpireDate(v string) *IssueApplyInfoPlus {
	s.ExpireDate = &v
	return s
}

func (s *IssueApplyInfoPlus) SetIssueAmt(v string) *IssueApplyInfoPlus {
	s.IssueAmt = &v
	return s
}

func (s *IssueApplyInfoPlus) SetOutBizNo(v string) *IssueApplyInfoPlus {
	s.OutBizNo = &v
	return s
}

func (s *IssueApplyInfoPlus) SetOutOrderNo(v string) *IssueApplyInfoPlus {
	s.OutOrderNo = &v
	return s
}

func (s *IssueApplyInfoPlus) SetWaybillId(v string) *IssueApplyInfoPlus {
	s.WaybillId = &v
	return s
}

// 电子提单变更状态明细（无效）
type EblStatusDeatil struct {
	// 当前提单状态
	CurrentEblStatus *string `json:"current_ebl_status,omitempty" xml:"current_ebl_status,omitempty" require:"true"`
	// 电子提单编号
	EblNo *string `json:"ebl_no,omitempty" xml:"ebl_no,omitempty" require:"true"`
	// 下一个提单状态
	NextEblStatus *string `json:"next_ebl_status,omitempty" xml:"next_ebl_status,omitempty" require:"true"`
}

func (s EblStatusDeatil) String() string {
	return tea.Prettify(s)
}

func (s EblStatusDeatil) GoString() string {
	return s.String()
}

func (s *EblStatusDeatil) SetCurrentEblStatus(v string) *EblStatusDeatil {
	s.CurrentEblStatus = &v
	return s
}

func (s *EblStatusDeatil) SetEblNo(v string) *EblStatusDeatil {
	s.EblNo = &v
	return s
}

func (s *EblStatusDeatil) SetNextEblStatus(v string) *EblStatusDeatil {
	s.NextEblStatus = &v
	return s
}

type CreateReceivableBillRequest struct {
	// OAuth模式下的授权token
	AuthToken         *string `json:"auth_token,omitempty" xml:"auth_token,omitempty"`
	ProductInstanceId *string `json:"product_instance_id,omitempty" xml:"product_instance_id,omitempty"`
	// 测试
	Test *AuthParty `json:"test,omitempty" xml:"test,omitempty"`
	// 的撒
	Status *ApiTest `json:"status,omitempty" xml:"status,omitempty" require:"true"`
}

func (s CreateReceivableBillRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateReceivableBillRequest) GoString() string {
	return s.String()
}

func (s *CreateReceivableBillRequest) SetAuthToken(v string) *CreateReceivableBillRequest {
	s.AuthToken = &v
	return s
}

func (s *CreateReceivableBillRequest) SetProductInstanceId(v string) *CreateReceivableBillRequest {
	s.ProductInstanceId = &v
	return s
}

func (s *CreateReceivableBillRequest) SetTest(v *AuthParty) *CreateReceivableBillRequest {
	s.Test = v
	return s
}

func (s *CreateReceivableBillRequest) SetStatus(v *ApiTest) *CreateReceivableBillRequest {
	s.Status = v
	return s
}

type CreateReceivableBillResponse struct {
	// 请求唯一ID，用于链路跟踪和问题排查
	ReqMsgId *string `json:"req_msg_id,omitempty" xml:"req_msg_id,omitempty"`
	// 结果码，一般OK表示调用成功
	ResultCode *string `json:"result_code,omitempty" xml:"result_code,omitempty"`
	// 异常信息的文本描述
	ResultMsg *string `json:"result_msg,omitempty" xml:"result_msg,omitempty"`
	// 返回结果
	TxCodes []*TxDto `json:"tx_codes,omitempty" xml:"tx_codes,omitempty" type:"Repeated"`
}

func (s CreateReceivableBillResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateReceivableBillResponse) GoString() string {
	return s.String()
}

func (s *CreateReceivableBillResponse) SetReqMsgId(v string) *CreateReceivableBillResponse {
	s.ReqMsgId = &v
	return s
}

func (s *CreateReceivableBillResponse) SetResultCode(v string) *CreateReceivableBillResponse {
	s.ResultCode = &v
	return s
}

func (s *CreateReceivableBillResponse) SetResultMsg(v string) *CreateReceivableBillResponse {
	s.ResultMsg = &v
	return s
}

func (s *CreateReceivableBillResponse) SetTxCodes(v []*TxDto) *CreateReceivableBillResponse {
	s.TxCodes = v
	return s
}

type CreateStandardVoucherRequest struct {
	// OAuth模式下的授权token
	AuthToken         *string `json:"auth_token,omitempty" xml:"auth_token,omitempty"`
	ProductInstanceId *string `json:"product_instance_id,omitempty" xml:"product_instance_id,omitempty"`
	// 数据标识
	BusinessCode *string `json:"business_code,omitempty" xml:"business_code,omitempty" require:"true" maxLength:"10" minLength:"0"`
	// 资产类型
	AssetType *string `json:"asset_type,omitempty" xml:"asset_type,omitempty" require:"true" maxLength:"10" minLength:"0"`
	// 发行金额_Integer
	AmountInt *int64 `json:"amount_int,omitempty" xml:"amount_int,omitempty" require:"true" maximum:"10" minimum:"0"`
	// 发行金额_Long
	AmountLong *int64 `json:"amount_long,omitempty" xml:"amount_long,omitempty" require:"true" maximum:"10" minimum:"0"`
	// 账户是否存在
	Exist *bool `json:"exist,omitempty" xml:"exist,omitempty" require:"true"`
	// 发行时间
	IssueTime *string `json:"issue_time,omitempty" xml:"issue_time,omitempty" require:"true" pattern:"\\d{4}[-]\\d{1,2}[-]\\d{1,2}[T]\\d{2}:\\d{2}:\\d{2}([Z]|([\\.]\\d{1,9})?[\\+]\\d{2}[\\:]?\\d{2})"`
	// 签署方
	ApiTest *ApiTest `json:"api_test,omitempty" xml:"api_test,omitempty" require:"true"`
	// 凭证列表_voucherArray
	VoucherArray []*int64 `json:"voucher_array,omitempty" xml:"voucher_array,omitempty" require:"true" type:"Repeated"`
	// 凭证列表_voucherList
	VoucherList []*string `json:"voucher_list,omitempty" xml:"voucher_list,omitempty" require:"true" type:"Repeated"`
	// 凭证列表_apiTestList
	ApiTestList []*ApiTest `json:"api_test_list,omitempty" xml:"api_test_list,omitempty" require:"true" type:"Repeated"`
	// 凭证列表_booleanList
	BooleanList []*bool `json:"boolean_list,omitempty" xml:"boolean_list,omitempty" require:"true" type:"Repeated"`
	// 凭证列表_dateList
	DateList []*string `json:"date_list,omitempty" xml:"date_list,omitempty" require:"true" type:"Repeated"`
}

func (s CreateStandardVoucherRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateStandardVoucherRequest) GoString() string {
	return s.String()
}

func (s *CreateStandardVoucherRequest) SetAuthToken(v string) *CreateStandardVoucherRequest {
	s.AuthToken = &v
	return s
}

func (s *CreateStandardVoucherRequest) SetProductInstanceId(v string) *CreateStandardVoucherRequest {
	s.ProductInstanceId = &v
	return s
}

func (s *CreateStandardVoucherRequest) SetBusinessCode(v string) *CreateStandardVoucherRequest {
	s.BusinessCode = &v
	return s
}

func (s *CreateStandardVoucherRequest) SetAssetType(v string) *CreateStandardVoucherRequest {
	s.AssetType = &v
	return s
}

func (s *CreateStandardVoucherRequest) SetAmountInt(v int64) *CreateStandardVoucherRequest {
	s.AmountInt = &v
	return s
}

func (s *CreateStandardVoucherRequest) SetAmountLong(v int64) *CreateStandardVoucherRequest {
	s.AmountLong = &v
	return s
}

func (s *CreateStandardVoucherRequest) SetExist(v bool) *CreateStandardVoucherRequest {
	s.Exist = &v
	return s
}

func (s *CreateStandardVoucherRequest) SetIssueTime(v string) *CreateStandardVoucherRequest {
	s.IssueTime = &v
	return s
}

func (s *CreateStandardVoucherRequest) SetApiTest(v *ApiTest) *CreateStandardVoucherRequest {
	s.ApiTest = v
	return s
}

func (s *CreateStandardVoucherRequest) SetVoucherArray(v []*int64) *CreateStandardVoucherRequest {
	s.VoucherArray = v
	return s
}

func (s *CreateStandardVoucherRequest) SetVoucherList(v []*string) *CreateStandardVoucherRequest {
	s.VoucherList = v
	return s
}

func (s *CreateStandardVoucherRequest) SetApiTestList(v []*ApiTest) *CreateStandardVoucherRequest {
	s.ApiTestList = v
	return s
}

func (s *CreateStandardVoucherRequest) SetBooleanList(v []*bool) *CreateStandardVoucherRequest {
	s.BooleanList = v
	return s
}

func (s *CreateStandardVoucherRequest) SetDateList(v []*string) *CreateStandardVoucherRequest {
	s.DateList = v
	return s
}

type CreateStandardVoucherResponse struct {
	// 请求唯一ID，用于链路跟踪和问题排查
	ReqMsgId *string `json:"req_msg_id,omitempty" xml:"req_msg_id,omitempty"`
	// 结果码，一般OK表示调用成功
	ResultCode *string `json:"result_code,omitempty" xml:"result_code,omitempty"`
	// 异常信息的文本描述
	ResultMsg *string `json:"result_msg,omitempty" xml:"result_msg,omitempty"`
	// 编码
	Code *string `json:"code,omitempty" xml:"code,omitempty" require:"true"`
}

func (s CreateStandardVoucherResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateStandardVoucherResponse) GoString() string {
	return s.String()
}

func (s *CreateStandardVoucherResponse) SetReqMsgId(v string) *CreateStandardVoucherResponse {
	s.ReqMsgId = &v
	return s
}

func (s *CreateStandardVoucherResponse) SetResultCode(v string) *CreateStandardVoucherResponse {
	s.ResultCode = &v
	return s
}

func (s *CreateStandardVoucherResponse) SetResultMsg(v string) *CreateStandardVoucherResponse {
	s.ResultMsg = &v
	return s
}

func (s *CreateStandardVoucherResponse) SetCode(v string) *CreateStandardVoucherResponse {
	s.Code = &v
	return s
}

type CreatePlatformDidRequest struct {
	// OAuth模式下的授权token
	AuthToken         *string `json:"auth_token,omitempty" xml:"auth_token,omitempty"`
	ProductInstanceId *string `json:"product_instance_id,omitempty" xml:"product_instance_id,omitempty"`
	// 企业名称
	EpCertName *string `json:"ep_cert_name,omitempty" xml:"ep_cert_name,omitempty" require:"true"`
	// 企业信用号码
	EpCertNo *string `json:"ep_cert_no,omitempty" xml:"ep_cert_no,omitempty" require:"true"`
	// 企业法人姓名
	LegalPersonCertName *string `json:"legal_person_cert_name,omitempty" xml:"legal_person_cert_name,omitempty" require:"true"`
	// 企业法人身份证号码
	LegalPersonCertNo *string `json:"legal_person_cert_no,omitempty" xml:"legal_person_cert_no,omitempty" require:"true"`
	// 扩展字段
	ExtensionInfo *string `json:"extension_info,omitempty" xml:"extension_info,omitempty"`
}

func (s CreatePlatformDidRequest) String() string {
	return tea.Prettify(s)
}

func (s CreatePlatformDidRequest) GoString() string {
	return s.String()
}

func (s *CreatePlatformDidRequest) SetAuthToken(v string) *CreatePlatformDidRequest {
	s.AuthToken = &v
	return s
}

func (s *CreatePlatformDidRequest) SetProductInstanceId(v string) *CreatePlatformDidRequest {
	s.ProductInstanceId = &v
	return s
}

func (s *CreatePlatformDidRequest) SetEpCertName(v string) *CreatePlatformDidRequest {
	s.EpCertName = &v
	return s
}

func (s *CreatePlatformDidRequest) SetEpCertNo(v string) *CreatePlatformDidRequest {
	s.EpCertNo = &v
	return s
}

func (s *CreatePlatformDidRequest) SetLegalPersonCertName(v string) *CreatePlatformDidRequest {
	s.LegalPersonCertName = &v
	return s
}

func (s *CreatePlatformDidRequest) SetLegalPersonCertNo(v string) *CreatePlatformDidRequest {
	s.LegalPersonCertNo = &v
	return s
}

func (s *CreatePlatformDidRequest) SetExtensionInfo(v string) *CreatePlatformDidRequest {
	s.ExtensionInfo = &v
	return s
}

type CreatePlatformDidResponse struct {
	// 请求唯一ID，用于链路跟踪和问题排查
	ReqMsgId *string `json:"req_msg_id,omitempty" xml:"req_msg_id,omitempty"`
	// 结果码，一般OK表示调用成功
	ResultCode *string `json:"result_code,omitempty" xml:"result_code,omitempty"`
	// 异常信息的文本描述
	ResultMsg *string `json:"result_msg,omitempty" xml:"result_msg,omitempty"`
	// 无车承运平台DIS
	Did *string `json:"did,omitempty" xml:"did,omitempty"`
}

func (s CreatePlatformDidResponse) String() string {
	return tea.Prettify(s)
}

func (s CreatePlatformDidResponse) GoString() string {
	return s.String()
}

func (s *CreatePlatformDidResponse) SetReqMsgId(v string) *CreatePlatformDidResponse {
	s.ReqMsgId = &v
	return s
}

func (s *CreatePlatformDidResponse) SetResultCode(v string) *CreatePlatformDidResponse {
	s.ResultCode = &v
	return s
}

func (s *CreatePlatformDidResponse) SetResultMsg(v string) *CreatePlatformDidResponse {
	s.ResultMsg = &v
	return s
}

func (s *CreatePlatformDidResponse) SetDid(v string) *CreatePlatformDidResponse {
	s.Did = &v
	return s
}

type CreateAgentDidRequest struct {
	// OAuth模式下的授权token
	AuthToken         *string `json:"auth_token,omitempty" xml:"auth_token,omitempty"`
	ProductInstanceId *string `json:"product_instance_id,omitempty" xml:"product_instance_id,omitempty"`
	// 创建did的代理did
	AgentDid *string `json:"agent_did,omitempty" xml:"agent_did,omitempty" require:"true"`
	// 企业名称
	EpCertName *string `json:"ep_cert_name,omitempty" xml:"ep_cert_name,omitempty" require:"true"`
	// 企业信用号码
	EpCertNo *string `json:"ep_cert_no,omitempty" xml:"ep_cert_no,omitempty" require:"true"`
	// 扩展字段
	ExtensionInfo *string `json:"extension_info,omitempty" xml:"extension_info,omitempty"`
	// 企业法人姓名
	LegalPersonCertName *string `json:"legal_person_cert_name,omitempty" xml:"legal_person_cert_name,omitempty" require:"true"`
	// 企业法人身份证号码
	LegalPersonCertNo *string `json:"legal_person_cert_no,omitempty" xml:"legal_person_cert_no,omitempty" require:"true"`
}

func (s CreateAgentDidRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateAgentDidRequest) GoString() string {
	return s.String()
}

func (s *CreateAgentDidRequest) SetAuthToken(v string) *CreateAgentDidRequest {
	s.AuthToken = &v
	return s
}

func (s *CreateAgentDidRequest) SetProductInstanceId(v string) *CreateAgentDidRequest {
	s.ProductInstanceId = &v
	return s
}

func (s *CreateAgentDidRequest) SetAgentDid(v string) *CreateAgentDidRequest {
	s.AgentDid = &v
	return s
}

func (s *CreateAgentDidRequest) SetEpCertName(v string) *CreateAgentDidRequest {
	s.EpCertName = &v
	return s
}

func (s *CreateAgentDidRequest) SetEpCertNo(v string) *CreateAgentDidRequest {
	s.EpCertNo = &v
	return s
}

func (s *CreateAgentDidRequest) SetExtensionInfo(v string) *CreateAgentDidRequest {
	s.ExtensionInfo = &v
	return s
}

func (s *CreateAgentDidRequest) SetLegalPersonCertName(v string) *CreateAgentDidRequest {
	s.LegalPersonCertName = &v
	return s
}

func (s *CreateAgentDidRequest) SetLegalPersonCertNo(v string) *CreateAgentDidRequest {
	s.LegalPersonCertNo = &v
	return s
}

type CreateAgentDidResponse struct {
	// 请求唯一ID，用于链路跟踪和问题排查
	ReqMsgId *string `json:"req_msg_id,omitempty" xml:"req_msg_id,omitempty"`
	// 结果码，一般OK表示调用成功
	ResultCode *string `json:"result_code,omitempty" xml:"result_code,omitempty"`
	// 异常信息的文本描述
	ResultMsg *string `json:"result_msg,omitempty" xml:"result_msg,omitempty"`
	// 无车承运平台did
	Did *string `json:"did,omitempty" xml:"did,omitempty"`
}

func (s CreateAgentDidResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateAgentDidResponse) GoString() string {
	return s.String()
}

func (s *CreateAgentDidResponse) SetReqMsgId(v string) *CreateAgentDidResponse {
	s.ReqMsgId = &v
	return s
}

func (s *CreateAgentDidResponse) SetResultCode(v string) *CreateAgentDidResponse {
	s.ResultCode = &v
	return s
}

func (s *CreateAgentDidResponse) SetResultMsg(v string) *CreateAgentDidResponse {
	s.ResultMsg = &v
	return s
}

func (s *CreateAgentDidResponse) SetDid(v string) *CreateAgentDidResponse {
	s.Did = &v
	return s
}

type CreateConsignorDisRequest struct {
	// OAuth模式下的授权token
	AuthToken         *string `json:"auth_token,omitempty" xml:"auth_token,omitempty"`
	ProductInstanceId *string `json:"product_instance_id,omitempty" xml:"product_instance_id,omitempty"`
	// 企业名称
	EpCertName *string `json:"ep_cert_name,omitempty" xml:"ep_cert_name,omitempty" require:"true"`
	// 企业信用号码
	EpCertNo *string `json:"ep_cert_no,omitempty" xml:"ep_cert_no,omitempty" require:"true"`
	// 扩展字段
	ExtensionInfo *string `json:"extension_info,omitempty" xml:"extension_info,omitempty"`
	// 企业法人姓名
	LegalPersonCertName *string `json:"legal_person_cert_name,omitempty" xml:"legal_person_cert_name,omitempty"`
	// 企业法人身份证号码
	LegalPersonCertNo *string `json:"legal_person_cert_no,omitempty" xml:"legal_person_cert_no,omitempty"`
	// 所属承运平台did
	PlatformDid *string `json:"platform_did,omitempty" xml:"platform_did,omitempty" require:"true"`
	// 是否核验货主四要素，如果为true  法人姓名和法人身份证号为必填
	CheckAll *bool `json:"check_all,omitempty" xml:"check_all,omitempty" require:"true"`
}

func (s CreateConsignorDisRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateConsignorDisRequest) GoString() string {
	return s.String()
}

func (s *CreateConsignorDisRequest) SetAuthToken(v string) *CreateConsignorDisRequest {
	s.AuthToken = &v
	return s
}

func (s *CreateConsignorDisRequest) SetProductInstanceId(v string) *CreateConsignorDisRequest {
	s.ProductInstanceId = &v
	return s
}

func (s *CreateConsignorDisRequest) SetEpCertName(v string) *CreateConsignorDisRequest {
	s.EpCertName = &v
	return s
}

func (s *CreateConsignorDisRequest) SetEpCertNo(v string) *CreateConsignorDisRequest {
	s.EpCertNo = &v
	return s
}

func (s *CreateConsignorDisRequest) SetExtensionInfo(v string) *CreateConsignorDisRequest {
	s.ExtensionInfo = &v
	return s
}

func (s *CreateConsignorDisRequest) SetLegalPersonCertName(v string) *CreateConsignorDisRequest {
	s.LegalPersonCertName = &v
	return s
}

func (s *CreateConsignorDisRequest) SetLegalPersonCertNo(v string) *CreateConsignorDisRequest {
	s.LegalPersonCertNo = &v
	return s
}

func (s *CreateConsignorDisRequest) SetPlatformDid(v string) *CreateConsignorDisRequest {
	s.PlatformDid = &v
	return s
}

func (s *CreateConsignorDisRequest) SetCheckAll(v bool) *CreateConsignorDisRequest {
	s.CheckAll = &v
	return s
}

type CreateConsignorDisResponse struct {
	// 请求唯一ID，用于链路跟踪和问题排查
	ReqMsgId *string `json:"req_msg_id,omitempty" xml:"req_msg_id,omitempty"`
	// 结果码，一般OK表示调用成功
	ResultCode *string `json:"result_code,omitempty" xml:"result_code,omitempty"`
	// 异常信息的文本描述
	ResultMsg *string `json:"result_msg,omitempty" xml:"result_msg,omitempty"`
	// 货主did信息
	Did *string `json:"did,omitempty" xml:"did,omitempty"`
}

func (s CreateConsignorDisResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateConsignorDisResponse) GoString() string {
	return s.String()
}

func (s *CreateConsignorDisResponse) SetReqMsgId(v string) *CreateConsignorDisResponse {
	s.ReqMsgId = &v
	return s
}

func (s *CreateConsignorDisResponse) SetResultCode(v string) *CreateConsignorDisResponse {
	s.ResultCode = &v
	return s
}

func (s *CreateConsignorDisResponse) SetResultMsg(v string) *CreateConsignorDisResponse {
	s.ResultMsg = &v
	return s
}

func (s *CreateConsignorDisResponse) SetDid(v string) *CreateConsignorDisResponse {
	s.Did = &v
	return s
}

type CreateDriverDisRequest struct {
	// OAuth模式下的授权token
	AuthToken         *string `json:"auth_token,omitempty" xml:"auth_token,omitempty"`
	ProductInstanceId *string `json:"product_instance_id,omitempty" xml:"product_instance_id,omitempty"`
	// 司机身份证号码
	CertNo *string `json:"cert_no,omitempty" xml:"cert_no,omitempty" require:"true"`
	// 扩展字段
	ExtensionInfo *string `json:"extension_info,omitempty" xml:"extension_info,omitempty"`
	// 司机手机号码
	Mobile *string `json:"mobile,omitempty" xml:"mobile,omitempty" require:"true"`
	// 司机姓名
	Name *string `json:"name,omitempty" xml:"name,omitempty" require:"true"`
	// 所属承运平台did
	PlatformDid *string `json:"platform_did,omitempty" xml:"platform_did,omitempty" require:"true"`
}

func (s CreateDriverDisRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateDriverDisRequest) GoString() string {
	return s.String()
}

func (s *CreateDriverDisRequest) SetAuthToken(v string) *CreateDriverDisRequest {
	s.AuthToken = &v
	return s
}

func (s *CreateDriverDisRequest) SetProductInstanceId(v string) *CreateDriverDisRequest {
	s.ProductInstanceId = &v
	return s
}

func (s *CreateDriverDisRequest) SetCertNo(v string) *CreateDriverDisRequest {
	s.CertNo = &v
	return s
}

func (s *CreateDriverDisRequest) SetExtensionInfo(v string) *CreateDriverDisRequest {
	s.ExtensionInfo = &v
	return s
}

func (s *CreateDriverDisRequest) SetMobile(v string) *CreateDriverDisRequest {
	s.Mobile = &v
	return s
}

func (s *CreateDriverDisRequest) SetName(v string) *CreateDriverDisRequest {
	s.Name = &v
	return s
}

func (s *CreateDriverDisRequest) SetPlatformDid(v string) *CreateDriverDisRequest {
	s.PlatformDid = &v
	return s
}

type CreateDriverDisResponse struct {
	// 请求唯一ID，用于链路跟踪和问题排查
	ReqMsgId *string `json:"req_msg_id,omitempty" xml:"req_msg_id,omitempty"`
	// 结果码，一般OK表示调用成功
	ResultCode *string `json:"result_code,omitempty" xml:"result_code,omitempty"`
	// 异常信息的文本描述
	ResultMsg *string `json:"result_msg,omitempty" xml:"result_msg,omitempty"`
	// 司机did
	Did *string `json:"did,omitempty" xml:"did,omitempty"`
}

func (s CreateDriverDisResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateDriverDisResponse) GoString() string {
	return s.String()
}

func (s *CreateDriverDisResponse) SetReqMsgId(v string) *CreateDriverDisResponse {
	s.ReqMsgId = &v
	return s
}

func (s *CreateDriverDisResponse) SetResultCode(v string) *CreateDriverDisResponse {
	s.ResultCode = &v
	return s
}

func (s *CreateDriverDisResponse) SetResultMsg(v string) *CreateDriverDisResponse {
	s.ResultMsg = &v
	return s
}

func (s *CreateDriverDisResponse) SetDid(v string) *CreateDriverDisResponse {
	s.Did = &v
	return s
}

type CreateCargoOrderRequest struct {
	// OAuth模式下的授权token
	AuthToken         *string `json:"auth_token,omitempty" xml:"auth_token,omitempty"`
	ProductInstanceId *string `json:"product_instance_id,omitempty" xml:"product_instance_id,omitempty"`
	// 运费
	AllFreight *string `json:"all_freight,omitempty" xml:"all_freight,omitempty" require:"true"`
	// 货物行业编码
	CargoBusinessCode *string `json:"cargo_business_code,omitempty" xml:"cargo_business_code,omitempty"`
	// 货物商品编码
	CargoCode *string `json:"cargo_code,omitempty" xml:"cargo_code,omitempty"`
	// 货运险金额
	CargoInsuranceMoney *string `json:"cargo_insurance_money,omitempty" xml:"cargo_insurance_money,omitempty"`
	// 货物名称
	CargoName *string `json:"cargo_name,omitempty" xml:"cargo_name,omitempty" require:"true"`
	// 货源单号
	CargoOrder *string `json:"cargo_order,omitempty" xml:"cargo_order,omitempty" require:"true"`
	// 货物类型
	CargoType *string `json:"cargo_type,omitempty" xml:"cargo_type,omitempty" require:"true"`
	// 货物单位
	CargoUnit *string `json:"cargo_unit,omitempty" xml:"cargo_unit,omitempty"`
	// 货物体积，单位（方）
	CargoVolume *string `json:"cargo_volume,omitempty" xml:"cargo_volume,omitempty"`
	// 货主did
	ConsignorDid *string `json:"consignor_did,omitempty" xml:"consignor_did,omitempty" require:"true"`
	// 卸货地，XX省-XX市-XX区
	DeliveryPlace *string `json:"delivery_place,omitempty" xml:"delivery_place,omitempty" require:"true"`
	// 装货地，XX省-XX市-XX区
	LoadingPlace *string `json:"loading_place,omitempty" xml:"loading_place,omitempty" require:"true"`
	// 所属承运平台did
	PlatformDid *string `json:"platform_did,omitempty" xml:"platform_did,omitempty" require:"true"`
	// 货源单创建时间
	StartTime *int64 `json:"start_time,omitempty" xml:"start_time,omitempty" require:"true"`
	// 联系人电话
	UserPhone *string `json:"user_phone,omitempty" xml:"user_phone,omitempty"`
	// 货物重量，单位（吨）
	Weight *string `json:"weight,omitempty" xml:"weight,omitempty" require:"true"`
}

func (s CreateCargoOrderRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateCargoOrderRequest) GoString() string {
	return s.String()
}

func (s *CreateCargoOrderRequest) SetAuthToken(v string) *CreateCargoOrderRequest {
	s.AuthToken = &v
	return s
}

func (s *CreateCargoOrderRequest) SetProductInstanceId(v string) *CreateCargoOrderRequest {
	s.ProductInstanceId = &v
	return s
}

func (s *CreateCargoOrderRequest) SetAllFreight(v string) *CreateCargoOrderRequest {
	s.AllFreight = &v
	return s
}

func (s *CreateCargoOrderRequest) SetCargoBusinessCode(v string) *CreateCargoOrderRequest {
	s.CargoBusinessCode = &v
	return s
}

func (s *CreateCargoOrderRequest) SetCargoCode(v string) *CreateCargoOrderRequest {
	s.CargoCode = &v
	return s
}

func (s *CreateCargoOrderRequest) SetCargoInsuranceMoney(v string) *CreateCargoOrderRequest {
	s.CargoInsuranceMoney = &v
	return s
}

func (s *CreateCargoOrderRequest) SetCargoName(v string) *CreateCargoOrderRequest {
	s.CargoName = &v
	return s
}

func (s *CreateCargoOrderRequest) SetCargoOrder(v string) *CreateCargoOrderRequest {
	s.CargoOrder = &v
	return s
}

func (s *CreateCargoOrderRequest) SetCargoType(v string) *CreateCargoOrderRequest {
	s.CargoType = &v
	return s
}

func (s *CreateCargoOrderRequest) SetCargoUnit(v string) *CreateCargoOrderRequest {
	s.CargoUnit = &v
	return s
}

func (s *CreateCargoOrderRequest) SetCargoVolume(v string) *CreateCargoOrderRequest {
	s.CargoVolume = &v
	return s
}

func (s *CreateCargoOrderRequest) SetConsignorDid(v string) *CreateCargoOrderRequest {
	s.ConsignorDid = &v
	return s
}

func (s *CreateCargoOrderRequest) SetDeliveryPlace(v string) *CreateCargoOrderRequest {
	s.DeliveryPlace = &v
	return s
}

func (s *CreateCargoOrderRequest) SetLoadingPlace(v string) *CreateCargoOrderRequest {
	s.LoadingPlace = &v
	return s
}

func (s *CreateCargoOrderRequest) SetPlatformDid(v string) *CreateCargoOrderRequest {
	s.PlatformDid = &v
	return s
}

func (s *CreateCargoOrderRequest) SetStartTime(v int64) *CreateCargoOrderRequest {
	s.StartTime = &v
	return s
}

func (s *CreateCargoOrderRequest) SetUserPhone(v string) *CreateCargoOrderRequest {
	s.UserPhone = &v
	return s
}

func (s *CreateCargoOrderRequest) SetWeight(v string) *CreateCargoOrderRequest {
	s.Weight = &v
	return s
}

type CreateCargoOrderResponse struct {
	// 请求唯一ID，用于链路跟踪和问题排查
	ReqMsgId *string `json:"req_msg_id,omitempty" xml:"req_msg_id,omitempty"`
	// 结果码，一般OK表示调用成功
	ResultCode *string `json:"result_code,omitempty" xml:"result_code,omitempty"`
	// 异常信息的文本描述
	ResultMsg *string `json:"result_msg,omitempty" xml:"result_msg,omitempty"`
	// 货源链上凭证
	TxCode *string `json:"tx_code,omitempty" xml:"tx_code,omitempty"`
}

func (s CreateCargoOrderResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateCargoOrderResponse) GoString() string {
	return s.String()
}

func (s *CreateCargoOrderResponse) SetReqMsgId(v string) *CreateCargoOrderResponse {
	s.ReqMsgId = &v
	return s
}

func (s *CreateCargoOrderResponse) SetResultCode(v string) *CreateCargoOrderResponse {
	s.ResultCode = &v
	return s
}

func (s *CreateCargoOrderResponse) SetResultMsg(v string) *CreateCargoOrderResponse {
	s.ResultMsg = &v
	return s
}

func (s *CreateCargoOrderResponse) SetTxCode(v string) *CreateCargoOrderResponse {
	s.TxCode = &v
	return s
}

type CreateCargoPayRequest struct {
	// OAuth模式下的授权token
	AuthToken         *string `json:"auth_token,omitempty" xml:"auth_token,omitempty"`
	ProductInstanceId *string `json:"product_instance_id,omitempty" xml:"product_instance_id,omitempty"`
	//  收款方银行账号
	BankAccountOfPayee *string `json:"bank_account_of_payee,omitempty" xml:"bank_account_of_payee,omitempty"`
	// 付款方银行账号,货主付款的开户银行账号
	BankAccountOfPayer *string `json:"bank_account_of_payer,omitempty" xml:"bank_account_of_payer,omitempty"`
	// 银行流水号
	BankNo *string `json:"bank_no,omitempty" xml:"bank_no,omitempty" require:"true"`
	// 收款方开户行（平台收款）
	// 取值：
	// MYBank_CloudCapital_2,  ## 云资金2.0
	BankOfPayee *string `json:"bank_of_payee,omitempty" xml:"bank_of_payee,omitempty"`
	// MYBank_CloudCapital_2,  ## 云资金2.0
	BankOfPayer *string `json:"bank_of_payer,omitempty" xml:"bank_of_payer,omitempty"`
	// 货源单号
	CargoOrder *string `json:"cargo_order,omitempty" xml:"cargo_order,omitempty" require:"true"`
	// 核验渠道，请按取值约束值填入。
	//
	// YBank_CloudCapital_2,  ## 云资金2.0
	CheckChannel *string `json:"check_channel,omitempty" xml:"check_channel,omitempty"`
	// CNY、USD
	// 币种
	Currency *string `json:"currency,omitempty" xml:"currency,omitempty"`
	// 费用类型 (运费、调度费、服务费)
	ExpenseType *string `json:"expense_type,omitempty" xml:"expense_type,omitempty" require:"true"`
	// 收款方名称
	Payee *string `json:"payee,omitempty" xml:"payee,omitempty"`
	// 付款方名称
	Payer *string `json:"payer,omitempty" xml:"payer,omitempty"`
	// 费用金额
	Payment *string `json:"payment,omitempty" xml:"payment,omitempty" require:"true"`
	// 是否核验
	PayCheck *bool `json:"pay_check,omitempty" xml:"pay_check,omitempty"`
	// 付款方did
	PayDid *string `json:"pay_did,omitempty" xml:"pay_did,omitempty" require:"true"`
	// 平台支付单号
	PayId *string `json:"pay_id,omitempty" xml:"pay_id,omitempty" require:"true"`
	// 支付备注
	PayNote *string `json:"pay_note,omitempty" xml:"pay_note,omitempty"`
	// 支付时间
	PayTime *int64 `json:"pay_time,omitempty" xml:"pay_time,omitempty" require:"true"`
	// 支付类型
	PayType *string `json:"pay_type,omitempty" xml:"pay_type,omitempty"`
	// 收款方did
	RecvDid *string `json:"recv_did,omitempty" xml:"recv_did,omitempty" require:"true"`
	// 支持关联单个运单和多个运单
	TaxWaybillIds []*string `json:"tax_waybill_ids,omitempty" xml:"tax_waybill_ids,omitempty" require:"true" type:"Repeated"`
}

func (s CreateCargoPayRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateCargoPayRequest) GoString() string {
	return s.String()
}

func (s *CreateCargoPayRequest) SetAuthToken(v string) *CreateCargoPayRequest {
	s.AuthToken = &v
	return s
}

func (s *CreateCargoPayRequest) SetProductInstanceId(v string) *CreateCargoPayRequest {
	s.ProductInstanceId = &v
	return s
}

func (s *CreateCargoPayRequest) SetBankAccountOfPayee(v string) *CreateCargoPayRequest {
	s.BankAccountOfPayee = &v
	return s
}

func (s *CreateCargoPayRequest) SetBankAccountOfPayer(v string) *CreateCargoPayRequest {
	s.BankAccountOfPayer = &v
	return s
}

func (s *CreateCargoPayRequest) SetBankNo(v string) *CreateCargoPayRequest {
	s.BankNo = &v
	return s
}

func (s *CreateCargoPayRequest) SetBankOfPayee(v string) *CreateCargoPayRequest {
	s.BankOfPayee = &v
	return s
}

func (s *CreateCargoPayRequest) SetBankOfPayer(v string) *CreateCargoPayRequest {
	s.BankOfPayer = &v
	return s
}

func (s *CreateCargoPayRequest) SetCargoOrder(v string) *CreateCargoPayRequest {
	s.CargoOrder = &v
	return s
}

func (s *CreateCargoPayRequest) SetCheckChannel(v string) *CreateCargoPayRequest {
	s.CheckChannel = &v
	return s
}

func (s *CreateCargoPayRequest) SetCurrency(v string) *CreateCargoPayRequest {
	s.Currency = &v
	return s
}

func (s *CreateCargoPayRequest) SetExpenseType(v string) *CreateCargoPayRequest {
	s.ExpenseType = &v
	return s
}

func (s *CreateCargoPayRequest) SetPayee(v string) *CreateCargoPayRequest {
	s.Payee = &v
	return s
}

func (s *CreateCargoPayRequest) SetPayer(v string) *CreateCargoPayRequest {
	s.Payer = &v
	return s
}

func (s *CreateCargoPayRequest) SetPayment(v string) *CreateCargoPayRequest {
	s.Payment = &v
	return s
}

func (s *CreateCargoPayRequest) SetPayCheck(v bool) *CreateCargoPayRequest {
	s.PayCheck = &v
	return s
}

func (s *CreateCargoPayRequest) SetPayDid(v string) *CreateCargoPayRequest {
	s.PayDid = &v
	return s
}

func (s *CreateCargoPayRequest) SetPayId(v string) *CreateCargoPayRequest {
	s.PayId = &v
	return s
}

func (s *CreateCargoPayRequest) SetPayNote(v string) *CreateCargoPayRequest {
	s.PayNote = &v
	return s
}

func (s *CreateCargoPayRequest) SetPayTime(v int64) *CreateCargoPayRequest {
	s.PayTime = &v
	return s
}

func (s *CreateCargoPayRequest) SetPayType(v string) *CreateCargoPayRequest {
	s.PayType = &v
	return s
}

func (s *CreateCargoPayRequest) SetRecvDid(v string) *CreateCargoPayRequest {
	s.RecvDid = &v
	return s
}

func (s *CreateCargoPayRequest) SetTaxWaybillIds(v []*string) *CreateCargoPayRequest {
	s.TaxWaybillIds = v
	return s
}

type CreateCargoPayResponse struct {
	// 请求唯一ID，用于链路跟踪和问题排查
	ReqMsgId *string `json:"req_msg_id,omitempty" xml:"req_msg_id,omitempty"`
	// 结果码，一般OK表示调用成功
	ResultCode *string `json:"result_code,omitempty" xml:"result_code,omitempty"`
	// 异常信息的文本描述
	ResultMsg *string `json:"result_msg,omitempty" xml:"result_msg,omitempty"`
	// 货源支付链上凭证
	TxCode *string `json:"tx_code,omitempty" xml:"tx_code,omitempty"`
}

func (s CreateCargoPayResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateCargoPayResponse) GoString() string {
	return s.String()
}

func (s *CreateCargoPayResponse) SetReqMsgId(v string) *CreateCargoPayResponse {
	s.ReqMsgId = &v
	return s
}

func (s *CreateCargoPayResponse) SetResultCode(v string) *CreateCargoPayResponse {
	s.ResultCode = &v
	return s
}

func (s *CreateCargoPayResponse) SetResultMsg(v string) *CreateCargoPayResponse {
	s.ResultMsg = &v
	return s
}

func (s *CreateCargoPayResponse) SetTxCode(v string) *CreateCargoPayResponse {
	s.TxCode = &v
	return s
}

type CreateWaybillOrderRequest struct {
	// OAuth模式下的授权token
	AuthToken         *string `json:"auth_token,omitempty" xml:"auth_token,omitempty"`
	ProductInstanceId *string `json:"product_instance_id,omitempty" xml:"product_instance_id,omitempty"`
	// 用户链上账户 与司机did 二选一填写
	AccountId *string `json:"account_id,omitempty" xml:"account_id,omitempty"`
	// 运费，单位（元），平台支付给司机的运费金额
	AllFreight *string `json:"all_freight,omitempty" xml:"all_freight,omitempty"`
	// 回单押金
	BackFee *string `json:"back_fee,omitempty" xml:"back_fee,omitempty"`
	// 业务类型
	BusinessType *string `json:"business_type,omitempty" xml:"business_type,omitempty"`
	// 货物行业编码
	CargoBusinessCode *string `json:"cargo_business_code,omitempty" xml:"cargo_business_code,omitempty"`
	// 货物商品编码
	CargoCode *string `json:"cargo_code,omitempty" xml:"cargo_code,omitempty"`
	// 货运险金额
	CargoInsuranceMoney *string `json:"cargo_insurance_money,omitempty" xml:"cargo_insurance_money,omitempty"`
	// 货源单号，关联货主订单
	CargoOrder *string `json:"cargo_order,omitempty" xml:"cargo_order,omitempty"`
	// 货源单号数组
	CargoOrders []*string `json:"cargo_orders,omitempty" xml:"cargo_orders,omitempty" type:"Repeated"`
	// 货物单位
	CargoUnit *string `json:"cargo_unit,omitempty" xml:"cargo_unit,omitempty"`
	// 货物体积，单位（方）
	CargoVolume *string `json:"cargo_volume,omitempty" xml:"cargo_volume,omitempty"`
	// 货物重量，单位（吨）
	CargoWeight *string `json:"cargo_weight,omitempty" xml:"cargo_weight,omitempty"`
	// 车牌颜色，黄色、蓝色、绿色
	CartBadgeColor *string `json:"cart_badge_color,omitempty" xml:"cart_badge_color,omitempty" require:"true"`
	// 车牌号
	//
	CartBadgeNo *string `json:"cart_badge_no,omitempty" xml:"cart_badge_no,omitempty" require:"true"`
	// 承运商did
	CarCaptainDid *string `json:"car_captain_did,omitempty" xml:"car_captain_did,omitempty"`
	// 货主运费金额，单位（元），货主支付给平台的运费金额
	ConsignorFreightAmount *string `json:"consignor_freight_amount,omitempty" xml:"consignor_freight_amount,omitempty"`
	// 建单时间，13位毫秒级时间戳
	//
	CreatedTime *int64 `json:"created_time,omitempty" xml:"created_time,omitempty" require:"true"`
	// 到达门点时间，13位毫秒级时间戳
	DestDoorsEndTime *int64 `json:"dest_doors_end_time,omitempty" xml:"dest_doors_end_time,omitempty"`
	// 门点城市CODE，6位区域行政编码
	DoorsCityCode *string `json:"doors_city_code,omitempty" xml:"doors_city_code,omitempty"`
	// 门点城市名称
	DoorsCityName *string `json:"doors_city_name,omitempty" xml:"doors_city_name,omitempty"`
	// 门点区县CODE，6位区域行政编码
	DoorsCountyCode *string `json:"doors_county_code,omitempty" xml:"doors_county_code,omitempty"`
	// 门点区县名称
	DoorsCountyName *string `json:"doors_county_name,omitempty" xml:"doors_county_name,omitempty"`
	// 门点行政区划代码
	DoorsDivisionCode *string `json:"doors_division_code,omitempty" xml:"doors_division_code,omitempty"`
	// 门点省份CODE，6位区域行政编码
	DoorsProvinceCode *string `json:"doors_province_code,omitempty" xml:"doors_province_code,omitempty"`
	// 门点省份名称
	DoorsProvinceName *string `json:"doors_province_name,omitempty" xml:"doors_province_name,omitempty"`
	// 发货方名称
	//
	Drawee *string `json:"drawee,omitempty" xml:"drawee,omitempty" require:"true"`
	// 发货方纳税人识别号
	//
	DraweeTaxNo *string `json:"drawee_tax_no,omitempty" xml:"drawee_tax_no,omitempty" require:"true"`
	// 司机分布式身份
	DriverDid *string `json:"driver_did,omitempty" xml:"driver_did,omitempty" require:"true"`
	// 司机姓名 已填司机分布式身份的情况下可不填
	//
	DriverName *string `json:"driver_name,omitempty" xml:"driver_name,omitempty"`
	// 目的地详细地址
	EndAddress *string `json:"end_address,omitempty" xml:"end_address,omitempty"`
	// 目的地城市CODE，6位区域行政编码
	//
	EndCityCode *string `json:"end_city_code,omitempty" xml:"end_city_code,omitempty" require:"true"`
	// 目的地城市名称
	//
	EndCityName *string `json:"end_city_name,omitempty" xml:"end_city_name,omitempty" require:"true"`
	// 目的地区县CODE，6位区域行政编码
	EndCountyCode *string `json:"end_county_code,omitempty" xml:"end_county_code,omitempty"`
	// 目的地区县名称
	//
	EndCountyName *string `json:"end_county_name,omitempty" xml:"end_county_name,omitempty"`
	// 结束行政区划代码，12位区域行政编码，（长途运输须准确提供前6位，如不能提供后6位可补0；短途运输须准确提供12位）
	//
	EndDivisionCode *string `json:"end_division_code,omitempty" xml:"end_division_code,omitempty" require:"true"`
	// 目的地省份CODE，6位区域行政编码
	//
	EndProvinceCode *string `json:"end_province_code,omitempty" xml:"end_province_code,omitempty" require:"true"`
	// 目的地省份名称
	//
	EndProvinceName *string `json:"end_province_name,omitempty" xml:"end_province_name,omitempty" require:"true"`
	// 目的地街道CODE，12区域行政编码
	EndStreetCode *string `json:"end_street_code,omitempty" xml:"end_street_code,omitempty"`
	// 目的地街道名称
	EndStreetName *string `json:"end_street_name,omitempty" xml:"end_street_name,omitempty"`
	// 运费增项
	//
	FreightIncr *string `json:"freight_incr,omitempty" xml:"freight_incr,omitempty"`
	// 货物数量
	//
	GoodsAmount *int64 `json:"goods_amount,omitempty" xml:"goods_amount,omitempty"`
	// 货物数量单位类型
	//
	GoodsAmountType *string `json:"goods_amount_type,omitempty" xml:"goods_amount_type,omitempty"`
	// 货物名称
	//
	GoodsName *string `json:"goods_name,omitempty" xml:"goods_name,omitempty" require:"true"`
	// 司机身份证号 已填司机分布式身份的情况下可不填
	IdCard *string `json:"id_card,omitempty" xml:"id_card,omitempty"`
	// 运费扣减
	//
	LossFee *string `json:"loss_fee,omitempty" xml:"loss_fee,omitempty"`
	// 司机手机号 已填司机分布式身份的情况下可不填
	//
	MobileNo *string `json:"mobile_no,omitempty" xml:"mobile_no,omitempty"`
	// 系统识别id 网商识别号
	//
	PartnerId *string `json:"partner_id,omitempty" xml:"partner_id,omitempty"`
	// 无车承运平台分布式数字身份，缺省时为自己的分布式数字身份
	//
	PlatformDid *string `json:"platform_did,omitempty" xml:"platform_did,omitempty" require:"true"`
	// 预付款金额
	//
	Prepayments *string `json:"prepayments,omitempty" xml:"prepayments,omitempty"`
	// 线下预付ETC
	//
	PrepaymentsBuyEtc *string `json:"prepayments_buy_etc,omitempty" xml:"prepayments_buy_etc,omitempty"`
	// 线下气款金额
	//
	PrepaymentsBuyGas *string `json:"prepayments_buy_gas,omitempty" xml:"prepayments_buy_gas,omitempty"`
	// 运单id
	//
	TaxWaybillId *string `json:"tax_waybill_id,omitempty" xml:"tax_waybill_id,omitempty" require:"true"`
	// 线下油款预付
	//
	PrepaymentsBuyOil *string `json:"prepayments_buy_oil,omitempty" xml:"prepayments_buy_oil,omitempty"`
	// 预付ETC卡金额
	//
	PrepaymentsEtccard *string `json:"prepayments_etccard,omitempty" xml:"prepayments_etccard,omitempty"`
	// 预付油卡金额
	//
	PrepaymentsOilcard *string `json:"prepayments_oilcard,omitempty" xml:"prepayments_oilcard,omitempty"`
	// 油卡赠送金额
	//
	PresentAmountOil *string `json:"present_amount_oil,omitempty" xml:"present_amount_oil,omitempty"`
	// 起始地详细地址
	StartAddress *string `json:"start_address,omitempty" xml:"start_address,omitempty"`
	// 起始地城市CODE，6位区域行政编码
	StartCityCode *string `json:"start_city_code,omitempty" xml:"start_city_code,omitempty" require:"true"`
	// 起始地城市名称
	//
	StartCityName *string `json:"start_city_name,omitempty" xml:"start_city_name,omitempty" require:"true"`
	// 起始地区县CODE，6位区域行政编码
	StartCountyCode *string `json:"start_county_code,omitempty" xml:"start_county_code,omitempty"`
	// 起始地区县名称
	//
	StartCountyName *string `json:"start_county_name,omitempty" xml:"start_county_name,omitempty"`
	// 起始行政区划代码，12位区域行政编码（长途运输须准确提供前6位，如不能提供后6位可补0；短途运输须准确提供12位）
	//
	StartDivisionCode *string `json:"start_division_code,omitempty" xml:"start_division_code,omitempty" require:"true"`
	// 起始地省份CODE，6位区域行政编码
	StartProvinceCode *string `json:"start_province_code,omitempty" xml:"start_province_code,omitempty" require:"true"`
	// 起始地省份名称
	//
	StartProvinceName *string `json:"start_province_name,omitempty" xml:"start_province_name,omitempty" require:"true"`
	// 起始地街道CODE，12区域行政编码
	StartStreetCode *string `json:"start_street_code,omitempty" xml:"start_street_code,omitempty"`
	// 起始地街道名称
	StartStreetName *string `json:"start_street_name,omitempty" xml:"start_street_name,omitempty"`
	// 起运时间戳，13位毫秒级时间戳
	StartTime *int64 `json:"start_time,omitempty" xml:"start_time,omitempty"`
	// 车长，可以填：不限车长或者1.8，2.7，3.8，4.2，5，6.2，6.8，7.7，8.2，8.7，9.6，11.7，12.5，13，13.7，15，16，17.5等不超过2位小数的数字
	TruckLength *string `json:"truck_length,omitempty" xml:"truck_length,omitempty"`
	// 可以填：不限车型，平板，高栏，厢式，集装箱，自卸，冷藏，保温，高低板，面包车，棉被车，爬梯车，飞翼车
	TruckType *string `json:"truck_type,omitempty" xml:"truck_type,omitempty"`
	// 运输单价
	//
	UnitPrice *string `json:"unit_price,omitempty" xml:"unit_price,omitempty"`
}

func (s CreateWaybillOrderRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateWaybillOrderRequest) GoString() string {
	return s.String()
}

func (s *CreateWaybillOrderRequest) SetAuthToken(v string) *CreateWaybillOrderRequest {
	s.AuthToken = &v
	return s
}

func (s *CreateWaybillOrderRequest) SetProductInstanceId(v string) *CreateWaybillOrderRequest {
	s.ProductInstanceId = &v
	return s
}

func (s *CreateWaybillOrderRequest) SetAccountId(v string) *CreateWaybillOrderRequest {
	s.AccountId = &v
	return s
}

func (s *CreateWaybillOrderRequest) SetAllFreight(v string) *CreateWaybillOrderRequest {
	s.AllFreight = &v
	return s
}

func (s *CreateWaybillOrderRequest) SetBackFee(v string) *CreateWaybillOrderRequest {
	s.BackFee = &v
	return s
}

func (s *CreateWaybillOrderRequest) SetBusinessType(v string) *CreateWaybillOrderRequest {
	s.BusinessType = &v
	return s
}

func (s *CreateWaybillOrderRequest) SetCargoBusinessCode(v string) *CreateWaybillOrderRequest {
	s.CargoBusinessCode = &v
	return s
}

func (s *CreateWaybillOrderRequest) SetCargoCode(v string) *CreateWaybillOrderRequest {
	s.CargoCode = &v
	return s
}

func (s *CreateWaybillOrderRequest) SetCargoInsuranceMoney(v string) *CreateWaybillOrderRequest {
	s.CargoInsuranceMoney = &v
	return s
}

func (s *CreateWaybillOrderRequest) SetCargoOrder(v string) *CreateWaybillOrderRequest {
	s.CargoOrder = &v
	return s
}

func (s *CreateWaybillOrderRequest) SetCargoOrders(v []*string) *CreateWaybillOrderRequest {
	s.CargoOrders = v
	return s
}

func (s *CreateWaybillOrderRequest) SetCargoUnit(v string) *CreateWaybillOrderRequest {
	s.CargoUnit = &v
	return s
}

func (s *CreateWaybillOrderRequest) SetCargoVolume(v string) *CreateWaybillOrderRequest {
	s.CargoVolume = &v
	return s
}

func (s *CreateWaybillOrderRequest) SetCargoWeight(v string) *CreateWaybillOrderRequest {
	s.CargoWeight = &v
	return s
}

func (s *CreateWaybillOrderRequest) SetCartBadgeColor(v string) *CreateWaybillOrderRequest {
	s.CartBadgeColor = &v
	return s
}

func (s *CreateWaybillOrderRequest) SetCartBadgeNo(v string) *CreateWaybillOrderRequest {
	s.CartBadgeNo = &v
	return s
}

func (s *CreateWaybillOrderRequest) SetCarCaptainDid(v string) *CreateWaybillOrderRequest {
	s.CarCaptainDid = &v
	return s
}

func (s *CreateWaybillOrderRequest) SetConsignorFreightAmount(v string) *CreateWaybillOrderRequest {
	s.ConsignorFreightAmount = &v
	return s
}

func (s *CreateWaybillOrderRequest) SetCreatedTime(v int64) *CreateWaybillOrderRequest {
	s.CreatedTime = &v
	return s
}

func (s *CreateWaybillOrderRequest) SetDestDoorsEndTime(v int64) *CreateWaybillOrderRequest {
	s.DestDoorsEndTime = &v
	return s
}

func (s *CreateWaybillOrderRequest) SetDoorsCityCode(v string) *CreateWaybillOrderRequest {
	s.DoorsCityCode = &v
	return s
}

func (s *CreateWaybillOrderRequest) SetDoorsCityName(v string) *CreateWaybillOrderRequest {
	s.DoorsCityName = &v
	return s
}

func (s *CreateWaybillOrderRequest) SetDoorsCountyCode(v string) *CreateWaybillOrderRequest {
	s.DoorsCountyCode = &v
	return s
}

func (s *CreateWaybillOrderRequest) SetDoorsCountyName(v string) *CreateWaybillOrderRequest {
	s.DoorsCountyName = &v
	return s
}

func (s *CreateWaybillOrderRequest) SetDoorsDivisionCode(v string) *CreateWaybillOrderRequest {
	s.DoorsDivisionCode = &v
	return s
}

func (s *CreateWaybillOrderRequest) SetDoorsProvinceCode(v string) *CreateWaybillOrderRequest {
	s.DoorsProvinceCode = &v
	return s
}

func (s *CreateWaybillOrderRequest) SetDoorsProvinceName(v string) *CreateWaybillOrderRequest {
	s.DoorsProvinceName = &v
	return s
}

func (s *CreateWaybillOrderRequest) SetDrawee(v string) *CreateWaybillOrderRequest {
	s.Drawee = &v
	return s
}

func (s *CreateWaybillOrderRequest) SetDraweeTaxNo(v string) *CreateWaybillOrderRequest {
	s.DraweeTaxNo = &v
	return s
}

func (s *CreateWaybillOrderRequest) SetDriverDid(v string) *CreateWaybillOrderRequest {
	s.DriverDid = &v
	return s
}

func (s *CreateWaybillOrderRequest) SetDriverName(v string) *CreateWaybillOrderRequest {
	s.DriverName = &v
	return s
}

func (s *CreateWaybillOrderRequest) SetEndAddress(v string) *CreateWaybillOrderRequest {
	s.EndAddress = &v
	return s
}

func (s *CreateWaybillOrderRequest) SetEndCityCode(v string) *CreateWaybillOrderRequest {
	s.EndCityCode = &v
	return s
}

func (s *CreateWaybillOrderRequest) SetEndCityName(v string) *CreateWaybillOrderRequest {
	s.EndCityName = &v
	return s
}

func (s *CreateWaybillOrderRequest) SetEndCountyCode(v string) *CreateWaybillOrderRequest {
	s.EndCountyCode = &v
	return s
}

func (s *CreateWaybillOrderRequest) SetEndCountyName(v string) *CreateWaybillOrderRequest {
	s.EndCountyName = &v
	return s
}

func (s *CreateWaybillOrderRequest) SetEndDivisionCode(v string) *CreateWaybillOrderRequest {
	s.EndDivisionCode = &v
	return s
}

func (s *CreateWaybillOrderRequest) SetEndProvinceCode(v string) *CreateWaybillOrderRequest {
	s.EndProvinceCode = &v
	return s
}

func (s *CreateWaybillOrderRequest) SetEndProvinceName(v string) *CreateWaybillOrderRequest {
	s.EndProvinceName = &v
	return s
}

func (s *CreateWaybillOrderRequest) SetEndStreetCode(v string) *CreateWaybillOrderRequest {
	s.EndStreetCode = &v
	return s
}

func (s *CreateWaybillOrderRequest) SetEndStreetName(v string) *CreateWaybillOrderRequest {
	s.EndStreetName = &v
	return s
}

func (s *CreateWaybillOrderRequest) SetFreightIncr(v string) *CreateWaybillOrderRequest {
	s.FreightIncr = &v
	return s
}

func (s *CreateWaybillOrderRequest) SetGoodsAmount(v int64) *CreateWaybillOrderRequest {
	s.GoodsAmount = &v
	return s
}

func (s *CreateWaybillOrderRequest) SetGoodsAmountType(v string) *CreateWaybillOrderRequest {
	s.GoodsAmountType = &v
	return s
}

func (s *CreateWaybillOrderRequest) SetGoodsName(v string) *CreateWaybillOrderRequest {
	s.GoodsName = &v
	return s
}

func (s *CreateWaybillOrderRequest) SetIdCard(v string) *CreateWaybillOrderRequest {
	s.IdCard = &v
	return s
}

func (s *CreateWaybillOrderRequest) SetLossFee(v string) *CreateWaybillOrderRequest {
	s.LossFee = &v
	return s
}

func (s *CreateWaybillOrderRequest) SetMobileNo(v string) *CreateWaybillOrderRequest {
	s.MobileNo = &v
	return s
}

func (s *CreateWaybillOrderRequest) SetPartnerId(v string) *CreateWaybillOrderRequest {
	s.PartnerId = &v
	return s
}

func (s *CreateWaybillOrderRequest) SetPlatformDid(v string) *CreateWaybillOrderRequest {
	s.PlatformDid = &v
	return s
}

func (s *CreateWaybillOrderRequest) SetPrepayments(v string) *CreateWaybillOrderRequest {
	s.Prepayments = &v
	return s
}

func (s *CreateWaybillOrderRequest) SetPrepaymentsBuyEtc(v string) *CreateWaybillOrderRequest {
	s.PrepaymentsBuyEtc = &v
	return s
}

func (s *CreateWaybillOrderRequest) SetPrepaymentsBuyGas(v string) *CreateWaybillOrderRequest {
	s.PrepaymentsBuyGas = &v
	return s
}

func (s *CreateWaybillOrderRequest) SetTaxWaybillId(v string) *CreateWaybillOrderRequest {
	s.TaxWaybillId = &v
	return s
}

func (s *CreateWaybillOrderRequest) SetPrepaymentsBuyOil(v string) *CreateWaybillOrderRequest {
	s.PrepaymentsBuyOil = &v
	return s
}

func (s *CreateWaybillOrderRequest) SetPrepaymentsEtccard(v string) *CreateWaybillOrderRequest {
	s.PrepaymentsEtccard = &v
	return s
}

func (s *CreateWaybillOrderRequest) SetPrepaymentsOilcard(v string) *CreateWaybillOrderRequest {
	s.PrepaymentsOilcard = &v
	return s
}

func (s *CreateWaybillOrderRequest) SetPresentAmountOil(v string) *CreateWaybillOrderRequest {
	s.PresentAmountOil = &v
	return s
}

func (s *CreateWaybillOrderRequest) SetStartAddress(v string) *CreateWaybillOrderRequest {
	s.StartAddress = &v
	return s
}

func (s *CreateWaybillOrderRequest) SetStartCityCode(v string) *CreateWaybillOrderRequest {
	s.StartCityCode = &v
	return s
}

func (s *CreateWaybillOrderRequest) SetStartCityName(v string) *CreateWaybillOrderRequest {
	s.StartCityName = &v
	return s
}

func (s *CreateWaybillOrderRequest) SetStartCountyCode(v string) *CreateWaybillOrderRequest {
	s.StartCountyCode = &v
	return s
}

func (s *CreateWaybillOrderRequest) SetStartCountyName(v string) *CreateWaybillOrderRequest {
	s.StartCountyName = &v
	return s
}

func (s *CreateWaybillOrderRequest) SetStartDivisionCode(v string) *CreateWaybillOrderRequest {
	s.StartDivisionCode = &v
	return s
}

func (s *CreateWaybillOrderRequest) SetStartProvinceCode(v string) *CreateWaybillOrderRequest {
	s.StartProvinceCode = &v
	return s
}

func (s *CreateWaybillOrderRequest) SetStartProvinceName(v string) *CreateWaybillOrderRequest {
	s.StartProvinceName = &v
	return s
}

func (s *CreateWaybillOrderRequest) SetStartStreetCode(v string) *CreateWaybillOrderRequest {
	s.StartStreetCode = &v
	return s
}

func (s *CreateWaybillOrderRequest) SetStartStreetName(v string) *CreateWaybillOrderRequest {
	s.StartStreetName = &v
	return s
}

func (s *CreateWaybillOrderRequest) SetStartTime(v int64) *CreateWaybillOrderRequest {
	s.StartTime = &v
	return s
}

func (s *CreateWaybillOrderRequest) SetTruckLength(v string) *CreateWaybillOrderRequest {
	s.TruckLength = &v
	return s
}

func (s *CreateWaybillOrderRequest) SetTruckType(v string) *CreateWaybillOrderRequest {
	s.TruckType = &v
	return s
}

func (s *CreateWaybillOrderRequest) SetUnitPrice(v string) *CreateWaybillOrderRequest {
	s.UnitPrice = &v
	return s
}

type CreateWaybillOrderResponse struct {
	// 请求唯一ID，用于链路跟踪和问题排查
	ReqMsgId *string `json:"req_msg_id,omitempty" xml:"req_msg_id,omitempty"`
	// 结果码，一般OK表示调用成功
	ResultCode *string `json:"result_code,omitempty" xml:"result_code,omitempty"`
	// 异常信息的文本描述
	ResultMsg *string `json:"result_msg,omitempty" xml:"result_msg,omitempty"`
	// 链上凭证
	TxCode *string `json:"tx_code,omitempty" xml:"tx_code,omitempty"`
}

func (s CreateWaybillOrderResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateWaybillOrderResponse) GoString() string {
	return s.String()
}

func (s *CreateWaybillOrderResponse) SetReqMsgId(v string) *CreateWaybillOrderResponse {
	s.ReqMsgId = &v
	return s
}

func (s *CreateWaybillOrderResponse) SetResultCode(v string) *CreateWaybillOrderResponse {
	s.ResultCode = &v
	return s
}

func (s *CreateWaybillOrderResponse) SetResultMsg(v string) *CreateWaybillOrderResponse {
	s.ResultMsg = &v
	return s
}

func (s *CreateWaybillOrderResponse) SetTxCode(v string) *CreateWaybillOrderResponse {
	s.TxCode = &v
	return s
}

type CreateWaybillPayRequest struct {
	// OAuth模式下的授权token
	AuthToken         *string `json:"auth_token,omitempty" xml:"auth_token,omitempty"`
	ProductInstanceId *string `json:"product_instance_id,omitempty" xml:"product_instance_id,omitempty"`
	// 确认到账时间
	//
	ArriveTime *int64 `json:"arrive_time,omitempty" xml:"arrive_time,omitempty" require:"true"`
	// 关联银行流水号
	//
	BankSn *string `json:"bank_sn,omitempty" xml:"bank_sn,omitempty" require:"true"`
	// 付款金额
	//
	PayAmount *string `json:"pay_amount,omitempty" xml:"pay_amount,omitempty" require:"true"`
	// 付款方银行账号
	//
	PayBankCardNo *string `json:"pay_bank_card_no,omitempty" xml:"pay_bank_card_no,omitempty" require:"true"`
	// 付款方开户行
	//
	PayBankName *string `json:"pay_bank_name,omitempty" xml:"pay_bank_name,omitempty" require:"true"`
	// 是否核验
	PayCheck *bool `json:"pay_check,omitempty" xml:"pay_check,omitempty"`
	// 付款方did
	//
	PayDid *string `json:"pay_did,omitempty" xml:"pay_did,omitempty" require:"true"`
	// 付款方名称
	//
	PayName *string `json:"pay_name,omitempty" xml:"pay_name,omitempty" require:"true"`
	// 付款时间戳
	//
	PayTime *int64 `json:"pay_time,omitempty" xml:"pay_time,omitempty" require:"true"`
	// 支付类型
	//
	PayTypeNew *string `json:"pay_type_new,omitempty" xml:"pay_type_new,omitempty" require:"true"`
	// 支付方式
	//
	PayWay *string `json:"pay_way,omitempty" xml:"pay_way,omitempty" require:"true"`
	// 无车承运平台分布式数字身份，缺省为自己的分布式数字身份
	//
	PlatformDid *string `json:"platform_did,omitempty" xml:"platform_did,omitempty" require:"true"`
	// 支付订单ID
	//
	PosInfoId *string `json:"pos_info_id,omitempty" xml:"pos_info_id,omitempty" require:"true"`
	// 实际付款方开户行
	//
	RealPayBank *string `json:"real_pay_bank,omitempty" xml:"real_pay_bank,omitempty" require:"true"`
	// 实际付款方银行账号
	//
	RealPayBankCardNo *string `json:"real_pay_bank_card_no,omitempty" xml:"real_pay_bank_card_no,omitempty" require:"true"`
	// 实际付款方名称
	//
	RealPayName *string `json:"real_pay_name,omitempty" xml:"real_pay_name,omitempty" require:"true"`
	// 收款方银行账号
	//
	RecvBankCardNo *string `json:"recv_bank_card_no,omitempty" xml:"recv_bank_card_no,omitempty" require:"true"`
	// 收款方开户行
	//
	RecvBankName *string `json:"recv_bank_name,omitempty" xml:"recv_bank_name,omitempty" require:"true"`
	// 收款方did
	//
	RecvDid *string `json:"recv_did,omitempty" xml:"recv_did,omitempty" require:"true"`
	// 收款方名称
	//
	RecvName *string `json:"recv_name,omitempty" xml:"recv_name,omitempty" require:"true"`
	// 关联的运单ID
	//
	WaybillId *string `json:"waybill_id,omitempty" xml:"waybill_id,omitempty" require:"true"`
	// MYBank_CloudCapital_1,  ## 云资金1.0
	// MYBank_CloudCapital_2,  ## 云资金2.0
	CheckChannel *string `json:"check_channel,omitempty" xml:"check_channel,omitempty"`
	// CNY, USD
	Currency *string `json:"currency,omitempty" xml:"currency,omitempty"`
}

func (s CreateWaybillPayRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateWaybillPayRequest) GoString() string {
	return s.String()
}

func (s *CreateWaybillPayRequest) SetAuthToken(v string) *CreateWaybillPayRequest {
	s.AuthToken = &v
	return s
}

func (s *CreateWaybillPayRequest) SetProductInstanceId(v string) *CreateWaybillPayRequest {
	s.ProductInstanceId = &v
	return s
}

func (s *CreateWaybillPayRequest) SetArriveTime(v int64) *CreateWaybillPayRequest {
	s.ArriveTime = &v
	return s
}

func (s *CreateWaybillPayRequest) SetBankSn(v string) *CreateWaybillPayRequest {
	s.BankSn = &v
	return s
}

func (s *CreateWaybillPayRequest) SetPayAmount(v string) *CreateWaybillPayRequest {
	s.PayAmount = &v
	return s
}

func (s *CreateWaybillPayRequest) SetPayBankCardNo(v string) *CreateWaybillPayRequest {
	s.PayBankCardNo = &v
	return s
}

func (s *CreateWaybillPayRequest) SetPayBankName(v string) *CreateWaybillPayRequest {
	s.PayBankName = &v
	return s
}

func (s *CreateWaybillPayRequest) SetPayCheck(v bool) *CreateWaybillPayRequest {
	s.PayCheck = &v
	return s
}

func (s *CreateWaybillPayRequest) SetPayDid(v string) *CreateWaybillPayRequest {
	s.PayDid = &v
	return s
}

func (s *CreateWaybillPayRequest) SetPayName(v string) *CreateWaybillPayRequest {
	s.PayName = &v
	return s
}

func (s *CreateWaybillPayRequest) SetPayTime(v int64) *CreateWaybillPayRequest {
	s.PayTime = &v
	return s
}

func (s *CreateWaybillPayRequest) SetPayTypeNew(v string) *CreateWaybillPayRequest {
	s.PayTypeNew = &v
	return s
}

func (s *CreateWaybillPayRequest) SetPayWay(v string) *CreateWaybillPayRequest {
	s.PayWay = &v
	return s
}

func (s *CreateWaybillPayRequest) SetPlatformDid(v string) *CreateWaybillPayRequest {
	s.PlatformDid = &v
	return s
}

func (s *CreateWaybillPayRequest) SetPosInfoId(v string) *CreateWaybillPayRequest {
	s.PosInfoId = &v
	return s
}

func (s *CreateWaybillPayRequest) SetRealPayBank(v string) *CreateWaybillPayRequest {
	s.RealPayBank = &v
	return s
}

func (s *CreateWaybillPayRequest) SetRealPayBankCardNo(v string) *CreateWaybillPayRequest {
	s.RealPayBankCardNo = &v
	return s
}

func (s *CreateWaybillPayRequest) SetRealPayName(v string) *CreateWaybillPayRequest {
	s.RealPayName = &v
	return s
}

func (s *CreateWaybillPayRequest) SetRecvBankCardNo(v string) *CreateWaybillPayRequest {
	s.RecvBankCardNo = &v
	return s
}

func (s *CreateWaybillPayRequest) SetRecvBankName(v string) *CreateWaybillPayRequest {
	s.RecvBankName = &v
	return s
}

func (s *CreateWaybillPayRequest) SetRecvDid(v string) *CreateWaybillPayRequest {
	s.RecvDid = &v
	return s
}

func (s *CreateWaybillPayRequest) SetRecvName(v string) *CreateWaybillPayRequest {
	s.RecvName = &v
	return s
}

func (s *CreateWaybillPayRequest) SetWaybillId(v string) *CreateWaybillPayRequest {
	s.WaybillId = &v
	return s
}

func (s *CreateWaybillPayRequest) SetCheckChannel(v string) *CreateWaybillPayRequest {
	s.CheckChannel = &v
	return s
}

func (s *CreateWaybillPayRequest) SetCurrency(v string) *CreateWaybillPayRequest {
	s.Currency = &v
	return s
}

type CreateWaybillPayResponse struct {
	// 请求唯一ID，用于链路跟踪和问题排查
	ReqMsgId *string `json:"req_msg_id,omitempty" xml:"req_msg_id,omitempty"`
	// 结果码，一般OK表示调用成功
	ResultCode *string `json:"result_code,omitempty" xml:"result_code,omitempty"`
	// 异常信息的文本描述
	ResultMsg *string `json:"result_msg,omitempty" xml:"result_msg,omitempty"`
	// 链上凭证
	//
	TxCode *string `json:"tx_code,omitempty" xml:"tx_code,omitempty"`
}

func (s CreateWaybillPayResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateWaybillPayResponse) GoString() string {
	return s.String()
}

func (s *CreateWaybillPayResponse) SetReqMsgId(v string) *CreateWaybillPayResponse {
	s.ReqMsgId = &v
	return s
}

func (s *CreateWaybillPayResponse) SetResultCode(v string) *CreateWaybillPayResponse {
	s.ResultCode = &v
	return s
}

func (s *CreateWaybillPayResponse) SetResultMsg(v string) *CreateWaybillPayResponse {
	s.ResultMsg = &v
	return s
}

func (s *CreateWaybillPayResponse) SetTxCode(v string) *CreateWaybillPayResponse {
	s.TxCode = &v
	return s
}

type ImportWaybillLocationRequest struct {
	// OAuth模式下的授权token
	AuthToken         *string `json:"auth_token,omitempty" xml:"auth_token,omitempty"`
	ProductInstanceId *string `json:"product_instance_id,omitempty" xml:"product_instance_id,omitempty"`
	// 描述信息
	//
	Desc *string `json:"desc,omitempty" xml:"desc,omitempty"`
	// 运单轨迹信息
	Location []*LogisticLocation `json:"location,omitempty" xml:"location,omitempty" require:"true" type:"Repeated"`
	// 所属平台分布式数字身份，缺省时为自己的分布式数字身份
	//
	PlatformDid *string `json:"platform_did,omitempty" xml:"platform_did,omitempty" require:"true"`
	// 运单id
	//
	TaxWaybillId *string `json:"tax_waybill_id,omitempty" xml:"tax_waybill_id,omitempty" require:"true"`
}

func (s ImportWaybillLocationRequest) String() string {
	return tea.Prettify(s)
}

func (s ImportWaybillLocationRequest) GoString() string {
	return s.String()
}

func (s *ImportWaybillLocationRequest) SetAuthToken(v string) *ImportWaybillLocationRequest {
	s.AuthToken = &v
	return s
}

func (s *ImportWaybillLocationRequest) SetProductInstanceId(v string) *ImportWaybillLocationRequest {
	s.ProductInstanceId = &v
	return s
}

func (s *ImportWaybillLocationRequest) SetDesc(v string) *ImportWaybillLocationRequest {
	s.Desc = &v
	return s
}

func (s *ImportWaybillLocationRequest) SetLocation(v []*LogisticLocation) *ImportWaybillLocationRequest {
	s.Location = v
	return s
}

func (s *ImportWaybillLocationRequest) SetPlatformDid(v string) *ImportWaybillLocationRequest {
	s.PlatformDid = &v
	return s
}

func (s *ImportWaybillLocationRequest) SetTaxWaybillId(v string) *ImportWaybillLocationRequest {
	s.TaxWaybillId = &v
	return s
}

type ImportWaybillLocationResponse struct {
	// 请求唯一ID，用于链路跟踪和问题排查
	ReqMsgId *string `json:"req_msg_id,omitempty" xml:"req_msg_id,omitempty"`
	// 结果码，一般OK表示调用成功
	ResultCode *string `json:"result_code,omitempty" xml:"result_code,omitempty"`
	// 异常信息的文本描述
	ResultMsg *string `json:"result_msg,omitempty" xml:"result_msg,omitempty"`
}

func (s ImportWaybillLocationResponse) String() string {
	return tea.Prettify(s)
}

func (s ImportWaybillLocationResponse) GoString() string {
	return s.String()
}

func (s *ImportWaybillLocationResponse) SetReqMsgId(v string) *ImportWaybillLocationResponse {
	s.ReqMsgId = &v
	return s
}

func (s *ImportWaybillLocationResponse) SetResultCode(v string) *ImportWaybillLocationResponse {
	s.ResultCode = &v
	return s
}

func (s *ImportWaybillLocationResponse) SetResultMsg(v string) *ImportWaybillLocationResponse {
	s.ResultMsg = &v
	return s
}

type CreateWaybillBillRequest struct {
	// OAuth模式下的授权token
	AuthToken         *string `json:"auth_token,omitempty" xml:"auth_token,omitempty"`
	ProductInstanceId *string `json:"product_instance_id,omitempty" xml:"product_instance_id,omitempty"`
	// 受票方名称
	//
	Drawee *string `json:"drawee,omitempty" xml:"drawee,omitempty" require:"true"`
	// 受票方纳税人识别号
	//
	DraweeTaxNo *string `json:"drawee_tax_no,omitempty" xml:"drawee_tax_no,omitempty" require:"true"`
	// 发票代码
	//
	Lzfpdm *string `json:"lzfpdm,omitempty" xml:"lzfpdm,omitempty"`
	// 发票号码
	//
	Lzfphm *string `json:"lzfphm,omitempty" xml:"lzfphm,omitempty"`
	// 开票时间戳
	//
	OpenTime *int64 `json:"open_time,omitempty" xml:"open_time,omitempty" require:"true"`
	// 所属平台的分布式数字身份，缺省时为自己的分布式数字身份
	//
	PlatformDid *string `json:"platform_did,omitempty" xml:"platform_did,omitempty" require:"true"`
	// 关联的运单ID
	//
	WaybillId *string `json:"waybill_id,omitempty" xml:"waybill_id,omitempty" require:"true"`
}

func (s CreateWaybillBillRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateWaybillBillRequest) GoString() string {
	return s.String()
}

func (s *CreateWaybillBillRequest) SetAuthToken(v string) *CreateWaybillBillRequest {
	s.AuthToken = &v
	return s
}

func (s *CreateWaybillBillRequest) SetProductInstanceId(v string) *CreateWaybillBillRequest {
	s.ProductInstanceId = &v
	return s
}

func (s *CreateWaybillBillRequest) SetDrawee(v string) *CreateWaybillBillRequest {
	s.Drawee = &v
	return s
}

func (s *CreateWaybillBillRequest) SetDraweeTaxNo(v string) *CreateWaybillBillRequest {
	s.DraweeTaxNo = &v
	return s
}

func (s *CreateWaybillBillRequest) SetLzfpdm(v string) *CreateWaybillBillRequest {
	s.Lzfpdm = &v
	return s
}

func (s *CreateWaybillBillRequest) SetLzfphm(v string) *CreateWaybillBillRequest {
	s.Lzfphm = &v
	return s
}

func (s *CreateWaybillBillRequest) SetOpenTime(v int64) *CreateWaybillBillRequest {
	s.OpenTime = &v
	return s
}

func (s *CreateWaybillBillRequest) SetPlatformDid(v string) *CreateWaybillBillRequest {
	s.PlatformDid = &v
	return s
}

func (s *CreateWaybillBillRequest) SetWaybillId(v string) *CreateWaybillBillRequest {
	s.WaybillId = &v
	return s
}

type CreateWaybillBillResponse struct {
	// 请求唯一ID，用于链路跟踪和问题排查
	ReqMsgId *string `json:"req_msg_id,omitempty" xml:"req_msg_id,omitempty"`
	// 结果码，一般OK表示调用成功
	ResultCode *string `json:"result_code,omitempty" xml:"result_code,omitempty"`
	// 异常信息的文本描述
	ResultMsg *string `json:"result_msg,omitempty" xml:"result_msg,omitempty"`
	// 链上凭证
	//
	TxCode *string `json:"tx_code,omitempty" xml:"tx_code,omitempty"`
}

func (s CreateWaybillBillResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateWaybillBillResponse) GoString() string {
	return s.String()
}

func (s *CreateWaybillBillResponse) SetReqMsgId(v string) *CreateWaybillBillResponse {
	s.ReqMsgId = &v
	return s
}

func (s *CreateWaybillBillResponse) SetResultCode(v string) *CreateWaybillBillResponse {
	s.ResultCode = &v
	return s
}

func (s *CreateWaybillBillResponse) SetResultMsg(v string) *CreateWaybillBillResponse {
	s.ResultMsg = &v
	return s
}

func (s *CreateWaybillBillResponse) SetTxCode(v string) *CreateWaybillBillResponse {
	s.TxCode = &v
	return s
}

type QueryWaybillStatusRequest struct {
	// OAuth模式下的授权token
	AuthToken         *string `json:"auth_token,omitempty" xml:"auth_token,omitempty"`
	ProductInstanceId *string `json:"product_instance_id,omitempty" xml:"product_instance_id,omitempty"`
	// 无车承运平台分布式数字身份，缺省时为自己的分布式数字身份
	//
	PlatformDid *string `json:"platform_did,omitempty" xml:"platform_did,omitempty"`
	// 运单ID
	//
	TaxWaybillId *string `json:"tax_waybill_id,omitempty" xml:"tax_waybill_id,omitempty" require:"true"`
}

func (s QueryWaybillStatusRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryWaybillStatusRequest) GoString() string {
	return s.String()
}

func (s *QueryWaybillStatusRequest) SetAuthToken(v string) *QueryWaybillStatusRequest {
	s.AuthToken = &v
	return s
}

func (s *QueryWaybillStatusRequest) SetProductInstanceId(v string) *QueryWaybillStatusRequest {
	s.ProductInstanceId = &v
	return s
}

func (s *QueryWaybillStatusRequest) SetPlatformDid(v string) *QueryWaybillStatusRequest {
	s.PlatformDid = &v
	return s
}

func (s *QueryWaybillStatusRequest) SetTaxWaybillId(v string) *QueryWaybillStatusRequest {
	s.TaxWaybillId = &v
	return s
}

type QueryWaybillStatusResponse struct {
	// 请求唯一ID，用于链路跟踪和问题排查
	ReqMsgId *string `json:"req_msg_id,omitempty" xml:"req_msg_id,omitempty"`
	// 结果码，一般OK表示调用成功
	ResultCode *string `json:"result_code,omitempty" xml:"result_code,omitempty"`
	// 异常信息的文本描述
	ResultMsg *string `json:"result_msg,omitempty" xml:"result_msg,omitempty"`
	// 运单状态
	//
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// 轨迹核验结果
	TrackCheckResult *TrackCheckResult `json:"track_check_result,omitempty" xml:"track_check_result,omitempty"`
}

func (s QueryWaybillStatusResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryWaybillStatusResponse) GoString() string {
	return s.String()
}

func (s *QueryWaybillStatusResponse) SetReqMsgId(v string) *QueryWaybillStatusResponse {
	s.ReqMsgId = &v
	return s
}

func (s *QueryWaybillStatusResponse) SetResultCode(v string) *QueryWaybillStatusResponse {
	s.ResultCode = &v
	return s
}

func (s *QueryWaybillStatusResponse) SetResultMsg(v string) *QueryWaybillStatusResponse {
	s.ResultMsg = &v
	return s
}

func (s *QueryWaybillStatusResponse) SetStatus(v string) *QueryWaybillStatusResponse {
	s.Status = &v
	return s
}

func (s *QueryWaybillStatusResponse) SetTrackCheckResult(v *TrackCheckResult) *QueryWaybillStatusResponse {
	s.TrackCheckResult = v
	return s
}

type FinishWaybillOrderRequest struct {
	// OAuth模式下的授权token
	AuthToken         *string `json:"auth_token,omitempty" xml:"auth_token,omitempty"`
	ProductInstanceId *string `json:"product_instance_id,omitempty" xml:"product_instance_id,omitempty"`
	// 链上账号 与司机did 二选一填写
	//
	AccountId *string `json:"account_id,omitempty" xml:"account_id,omitempty"`
	// 运费，单位（元），平台支付给司机的运费
	AllFreight *string `json:"all_freight,omitempty" xml:"all_freight,omitempty" require:"true"`
	// 回单押金
	//
	BackFee *string `json:"back_fee,omitempty" xml:"back_fee,omitempty"`
	// 业务类型
	BusinessType *string `json:"business_type,omitempty" xml:"business_type,omitempty"`
	// 货物行业编码
	CargoBusinessCode *string `json:"cargo_business_code,omitempty" xml:"cargo_business_code,omitempty"`
	// 货物商品编码
	CargoCode *string `json:"cargo_code,omitempty" xml:"cargo_code,omitempty"`
	// 货物运费险
	CargoInsuranceMoney *string `json:"cargo_insurance_money,omitempty" xml:"cargo_insurance_money,omitempty"`
	// 货源单号
	//
	CargoOrder *string `json:"cargo_order,omitempty" xml:"cargo_order,omitempty"`
	// 货物单位
	CargoUnit *string `json:"cargo_unit,omitempty" xml:"cargo_unit,omitempty"`
	// 货物体积，单位（方）
	CargoVolume *string `json:"cargo_volume,omitempty" xml:"cargo_volume,omitempty"`
	// 货物重量，单位（吨）
	CargoWeight *string `json:"cargo_weight,omitempty" xml:"cargo_weight,omitempty"`
	// 车牌颜色，黄色、蓝色、绿色
	CartBadgeColor *string `json:"cart_badge_color,omitempty" xml:"cart_badge_color,omitempty" require:"true"`
	// 车牌号
	//
	CartBadgeNo *string `json:"cart_badge_no,omitempty" xml:"cart_badge_no,omitempty" require:"true"`
	// 承运商did
	CarCaptainDid *string `json:"car_captain_did,omitempty" xml:"car_captain_did,omitempty"`
	// 货主运费金额，货主支付给平台的运费金额
	ConsignorFreightAmount *string `json:"consignor_freight_amount,omitempty" xml:"consignor_freight_amount,omitempty" require:"true"`
	// 建单时间，13位毫秒级时间戳
	CreatedTime *int64 `json:"created_time,omitempty" xml:"created_time,omitempty" require:"true"`
	// 到达门点时间，13位毫秒级时间戳
	DestDoorsEndTime *int64 `json:"dest_doors_end_time,omitempty" xml:"dest_doors_end_time,omitempty"`
	// 门点城市CODE，6位区域行政编码
	DoorsCityCode *string `json:"doors_city_code,omitempty" xml:"doors_city_code,omitempty"`
	// 门点城市名称
	DoorsCityName *string `json:"doors_city_name,omitempty" xml:"doors_city_name,omitempty"`
	// 门点区县CODE，6位区域行政编码
	DoorsCountyCode *string `json:"doors_county_code,omitempty" xml:"doors_county_code,omitempty"`
	// 门点区县名称
	DoorsCountyName *string `json:"doors_county_name,omitempty" xml:"doors_county_name,omitempty"`
	// 门点行政区划代码，12位区域行政编码，（长途运输须准确提供前6位，如不能提供后6位可补0；短途运输须准确提供12位）
	DoorsDivisionCode *string `json:"doors_division_code,omitempty" xml:"doors_division_code,omitempty"`
	// 门点省份CODE，6位区域行政编
	DoorsProvinceCode *string `json:"doors_province_code,omitempty" xml:"doors_province_code,omitempty"`
	// 门点省份名称
	DoorsProvinceName *string `json:"doors_province_name,omitempty" xml:"doors_province_name,omitempty"`
	// 发货方名称
	//
	Drawee *string `json:"drawee,omitempty" xml:"drawee,omitempty" require:"true"`
	// 发货方纳税人识别号
	//
	DraweeTaxNo *string `json:"drawee_tax_no,omitempty" xml:"drawee_tax_no,omitempty" require:"true"`
	// 司机分布式数字身份
	//
	DriverDid *string `json:"driver_did,omitempty" xml:"driver_did,omitempty" require:"true"`
	// 司机姓名 已填司机分布式身份的情况下可不填
	//
	DriverName *string `json:"driver_name,omitempty" xml:"driver_name,omitempty"`
	// 目的地详细地址
	EndAddress *string `json:"end_address,omitempty" xml:"end_address,omitempty"`
	// 目的地城市CODE，6位区域行政编
	EndCityCode *string `json:"end_city_code,omitempty" xml:"end_city_code,omitempty" require:"true"`
	// 目的地城市名称
	//
	EndCityName *string `json:"end_city_name,omitempty" xml:"end_city_name,omitempty" require:"true"`
	// 目的地区县CODE，6位区域行政编
	EndCountyCode *string `json:"end_county_code,omitempty" xml:"end_county_code,omitempty"`
	// 目的地区县名称
	//
	EndCountyName *string `json:"end_county_name,omitempty" xml:"end_county_name,omitempty"`
	// 结束行政区划代码 ，12位区域行政编，（长途运输须准确提供前6位，如不能提供后6位可补0；短途运输须准确提供12位）
	//
	EndDivisionCode *string `json:"end_division_code,omitempty" xml:"end_division_code,omitempty" require:"true"`
	// 目的地省份CODE，6位区域行政编
	//
	EndProvinceCode *string `json:"end_province_code,omitempty" xml:"end_province_code,omitempty" require:"true"`
	// 目的地省份名称
	//
	EndProvinceName *string `json:"end_province_name,omitempty" xml:"end_province_name,omitempty" require:"true"`
	// 目的地街道CODE，12位行政区域编码
	EndStreetCode *string `json:"end_street_code,omitempty" xml:"end_street_code,omitempty"`
	// 目的地街道名称
	EndStreetName *string `json:"end_street_name,omitempty" xml:"end_street_name,omitempty"`
	// 终结时间，13位毫秒级时间戳
	EndTime *int64 `json:"end_time,omitempty" xml:"end_time,omitempty" require:"true"`
	// 运费增项
	//
	FreightIncr *string `json:"freight_incr,omitempty" xml:"freight_incr,omitempty"`
	// 货物数量
	//
	GoodsAmount *int64 `json:"goods_amount,omitempty" xml:"goods_amount,omitempty"`
	// 货物数量单位类型
	//
	GoodsAmountType *string `json:"goods_amount_type,omitempty" xml:"goods_amount_type,omitempty"`
	// 货物名称
	//
	GoodsName *string `json:"goods_name,omitempty" xml:"goods_name,omitempty" require:"true"`
	// 司机身份证号 已填司机分布式身份的情况下可不填
	//
	IdCard *string `json:"id_card,omitempty" xml:"id_card,omitempty"`
	// 运费扣减
	//
	LossFee *string `json:"loss_fee,omitempty" xml:"loss_fee,omitempty"`
	// 司机手机号 已填司机分布式身份的情况下可不填
	//
	MobileNo *string `json:"mobile_no,omitempty" xml:"mobile_no,omitempty"`
	// 系统识别id 网商识别号
	//
	PartnerId *string `json:"partner_id,omitempty" xml:"partner_id,omitempty"`
	// 是否进行资金验证
	//
	PayCheck *bool `json:"pay_check,omitempty" xml:"pay_check,omitempty"`
	// 无车承运平台分布式数字身份，缺省时为自己的分布式数字身份
	//
	PlatformDid *string `json:"platform_did,omitempty" xml:"platform_did,omitempty" require:"true"`
	// 预付款金额
	//
	Prepayments *string `json:"prepayments,omitempty" xml:"prepayments,omitempty"`
	// 线下预付ETC
	//
	PrepaymentsBuyEtc *string `json:"prepayments_buy_etc,omitempty" xml:"prepayments_buy_etc,omitempty"`
	// 线下气款金额
	//
	PrepaymentsBuyGas *string `json:"prepayments_buy_gas,omitempty" xml:"prepayments_buy_gas,omitempty"`
	// 线下油款预付
	//
	PrepaymentsBuyOil *string `json:"prepayments_buy_oil,omitempty" xml:"prepayments_buy_oil,omitempty"`
	// 预付ETC卡金额
	//
	PrepaymentsEtccard *string `json:"prepayments_etccard,omitempty" xml:"prepayments_etccard,omitempty"`
	// 预付油卡金额
	//
	PrepaymentsOilcard *string `json:"prepayments_oilcard,omitempty" xml:"prepayments_oilcard,omitempty"`
	// 油卡赠送金额
	//
	PresentAmountOil *string `json:"present_amount_oil,omitempty" xml:"present_amount_oil,omitempty"`
	// 起始地详细地址
	StartAddress *string `json:"start_address,omitempty" xml:"start_address,omitempty"`
	// 起始地CODE
	// ，6位区域行政编
	StartCityCode *string `json:"start_city_code,omitempty" xml:"start_city_code,omitempty" require:"true"`
	// 起始地城市名称
	//
	StartCityName *string `json:"start_city_name,omitempty" xml:"start_city_name,omitempty" require:"true"`
	// 起始地区县CODE
	// ，6位区域行政编
	StartCountyCode *string `json:"start_county_code,omitempty" xml:"start_county_code,omitempty"`
	// 起始地区县名称
	//
	StartCountyName *string `json:"start_county_name,omitempty" xml:"start_county_name,omitempty"`
	// 起始行政区划代码，12位区域行政编，（长途运输须准确提供前6位，如不能提供后6位可补0；短途运输须准确提供12位）
	//
	StartDivisionCode *string `json:"start_division_code,omitempty" xml:"start_division_code,omitempty" require:"true"`
	// 起始地省份CODE
	// ，6位区域行政编
	StartProvinceCode *string `json:"start_province_code,omitempty" xml:"start_province_code,omitempty" require:"true"`
	// 起始地省份名称
	//
	StartProvinceName *string `json:"start_province_name,omitempty" xml:"start_province_name,omitempty" require:"true"`
	// 起始地街道CODE，12位区域行政编
	StartStreetCode *string `json:"start_street_code,omitempty" xml:"start_street_code,omitempty"`
	// 起始地街道名称
	StartStreetName *string `json:"start_street_name,omitempty" xml:"start_street_name,omitempty"`
	// 起运时间，13位毫秒级时间戳
	//
	StartTime *int64 `json:"start_time,omitempty" xml:"start_time,omitempty" require:"true"`
	// 运单ID
	//
	TaxWaybillId *string `json:"tax_waybill_id,omitempty" xml:"tax_waybill_id,omitempty" require:"true"`
	// 车长，可以填：不限车长或者1.8，2.7，3.8，4.2，5，6.2，6.8，7.7，8.2，8.7，9.6，11.7，12.5，13，13.7，15，16，17.5等不超过2位小数的数字
	TruckLength *string `json:"truck_length,omitempty" xml:"truck_length,omitempty"`
	// 车型，可以填写：不限车型，平板，高栏，厢式，集装箱，自卸，冷藏，保温，高低板，面包车，棉被车，爬梯车，飞翼车
	TruckType *string `json:"truck_type,omitempty" xml:"truck_type,omitempty"`
	// 运输单价
	//
	UnitPrice *string `json:"unit_price,omitempty" xml:"unit_price,omitempty"`
}

func (s FinishWaybillOrderRequest) String() string {
	return tea.Prettify(s)
}

func (s FinishWaybillOrderRequest) GoString() string {
	return s.String()
}

func (s *FinishWaybillOrderRequest) SetAuthToken(v string) *FinishWaybillOrderRequest {
	s.AuthToken = &v
	return s
}

func (s *FinishWaybillOrderRequest) SetProductInstanceId(v string) *FinishWaybillOrderRequest {
	s.ProductInstanceId = &v
	return s
}

func (s *FinishWaybillOrderRequest) SetAccountId(v string) *FinishWaybillOrderRequest {
	s.AccountId = &v
	return s
}

func (s *FinishWaybillOrderRequest) SetAllFreight(v string) *FinishWaybillOrderRequest {
	s.AllFreight = &v
	return s
}

func (s *FinishWaybillOrderRequest) SetBackFee(v string) *FinishWaybillOrderRequest {
	s.BackFee = &v
	return s
}

func (s *FinishWaybillOrderRequest) SetBusinessType(v string) *FinishWaybillOrderRequest {
	s.BusinessType = &v
	return s
}

func (s *FinishWaybillOrderRequest) SetCargoBusinessCode(v string) *FinishWaybillOrderRequest {
	s.CargoBusinessCode = &v
	return s
}

func (s *FinishWaybillOrderRequest) SetCargoCode(v string) *FinishWaybillOrderRequest {
	s.CargoCode = &v
	return s
}

func (s *FinishWaybillOrderRequest) SetCargoInsuranceMoney(v string) *FinishWaybillOrderRequest {
	s.CargoInsuranceMoney = &v
	return s
}

func (s *FinishWaybillOrderRequest) SetCargoOrder(v string) *FinishWaybillOrderRequest {
	s.CargoOrder = &v
	return s
}

func (s *FinishWaybillOrderRequest) SetCargoUnit(v string) *FinishWaybillOrderRequest {
	s.CargoUnit = &v
	return s
}

func (s *FinishWaybillOrderRequest) SetCargoVolume(v string) *FinishWaybillOrderRequest {
	s.CargoVolume = &v
	return s
}

func (s *FinishWaybillOrderRequest) SetCargoWeight(v string) *FinishWaybillOrderRequest {
	s.CargoWeight = &v
	return s
}

func (s *FinishWaybillOrderRequest) SetCartBadgeColor(v string) *FinishWaybillOrderRequest {
	s.CartBadgeColor = &v
	return s
}

func (s *FinishWaybillOrderRequest) SetCartBadgeNo(v string) *FinishWaybillOrderRequest {
	s.CartBadgeNo = &v
	return s
}

func (s *FinishWaybillOrderRequest) SetCarCaptainDid(v string) *FinishWaybillOrderRequest {
	s.CarCaptainDid = &v
	return s
}

func (s *FinishWaybillOrderRequest) SetConsignorFreightAmount(v string) *FinishWaybillOrderRequest {
	s.ConsignorFreightAmount = &v
	return s
}

func (s *FinishWaybillOrderRequest) SetCreatedTime(v int64) *FinishWaybillOrderRequest {
	s.CreatedTime = &v
	return s
}

func (s *FinishWaybillOrderRequest) SetDestDoorsEndTime(v int64) *FinishWaybillOrderRequest {
	s.DestDoorsEndTime = &v
	return s
}

func (s *FinishWaybillOrderRequest) SetDoorsCityCode(v string) *FinishWaybillOrderRequest {
	s.DoorsCityCode = &v
	return s
}

func (s *FinishWaybillOrderRequest) SetDoorsCityName(v string) *FinishWaybillOrderRequest {
	s.DoorsCityName = &v
	return s
}

func (s *FinishWaybillOrderRequest) SetDoorsCountyCode(v string) *FinishWaybillOrderRequest {
	s.DoorsCountyCode = &v
	return s
}

func (s *FinishWaybillOrderRequest) SetDoorsCountyName(v string) *FinishWaybillOrderRequest {
	s.DoorsCountyName = &v
	return s
}

func (s *FinishWaybillOrderRequest) SetDoorsDivisionCode(v string) *FinishWaybillOrderRequest {
	s.DoorsDivisionCode = &v
	return s
}

func (s *FinishWaybillOrderRequest) SetDoorsProvinceCode(v string) *FinishWaybillOrderRequest {
	s.DoorsProvinceCode = &v
	return s
}

func (s *FinishWaybillOrderRequest) SetDoorsProvinceName(v string) *FinishWaybillOrderRequest {
	s.DoorsProvinceName = &v
	return s
}

func (s *FinishWaybillOrderRequest) SetDrawee(v string) *FinishWaybillOrderRequest {
	s.Drawee = &v
	return s
}

func (s *FinishWaybillOrderRequest) SetDraweeTaxNo(v string) *FinishWaybillOrderRequest {
	s.DraweeTaxNo = &v
	return s
}

func (s *FinishWaybillOrderRequest) SetDriverDid(v string) *FinishWaybillOrderRequest {
	s.DriverDid = &v
	return s
}

func (s *FinishWaybillOrderRequest) SetDriverName(v string) *FinishWaybillOrderRequest {
	s.DriverName = &v
	return s
}

func (s *FinishWaybillOrderRequest) SetEndAddress(v string) *FinishWaybillOrderRequest {
	s.EndAddress = &v
	return s
}

func (s *FinishWaybillOrderRequest) SetEndCityCode(v string) *FinishWaybillOrderRequest {
	s.EndCityCode = &v
	return s
}

func (s *FinishWaybillOrderRequest) SetEndCityName(v string) *FinishWaybillOrderRequest {
	s.EndCityName = &v
	return s
}

func (s *FinishWaybillOrderRequest) SetEndCountyCode(v string) *FinishWaybillOrderRequest {
	s.EndCountyCode = &v
	return s
}

func (s *FinishWaybillOrderRequest) SetEndCountyName(v string) *FinishWaybillOrderRequest {
	s.EndCountyName = &v
	return s
}

func (s *FinishWaybillOrderRequest) SetEndDivisionCode(v string) *FinishWaybillOrderRequest {
	s.EndDivisionCode = &v
	return s
}

func (s *FinishWaybillOrderRequest) SetEndProvinceCode(v string) *FinishWaybillOrderRequest {
	s.EndProvinceCode = &v
	return s
}

func (s *FinishWaybillOrderRequest) SetEndProvinceName(v string) *FinishWaybillOrderRequest {
	s.EndProvinceName = &v
	return s
}

func (s *FinishWaybillOrderRequest) SetEndStreetCode(v string) *FinishWaybillOrderRequest {
	s.EndStreetCode = &v
	return s
}

func (s *FinishWaybillOrderRequest) SetEndStreetName(v string) *FinishWaybillOrderRequest {
	s.EndStreetName = &v
	return s
}

func (s *FinishWaybillOrderRequest) SetEndTime(v int64) *FinishWaybillOrderRequest {
	s.EndTime = &v
	return s
}

func (s *FinishWaybillOrderRequest) SetFreightIncr(v string) *FinishWaybillOrderRequest {
	s.FreightIncr = &v
	return s
}

func (s *FinishWaybillOrderRequest) SetGoodsAmount(v int64) *FinishWaybillOrderRequest {
	s.GoodsAmount = &v
	return s
}

func (s *FinishWaybillOrderRequest) SetGoodsAmountType(v string) *FinishWaybillOrderRequest {
	s.GoodsAmountType = &v
	return s
}

func (s *FinishWaybillOrderRequest) SetGoodsName(v string) *FinishWaybillOrderRequest {
	s.GoodsName = &v
	return s
}

func (s *FinishWaybillOrderRequest) SetIdCard(v string) *FinishWaybillOrderRequest {
	s.IdCard = &v
	return s
}

func (s *FinishWaybillOrderRequest) SetLossFee(v string) *FinishWaybillOrderRequest {
	s.LossFee = &v
	return s
}

func (s *FinishWaybillOrderRequest) SetMobileNo(v string) *FinishWaybillOrderRequest {
	s.MobileNo = &v
	return s
}

func (s *FinishWaybillOrderRequest) SetPartnerId(v string) *FinishWaybillOrderRequest {
	s.PartnerId = &v
	return s
}

func (s *FinishWaybillOrderRequest) SetPayCheck(v bool) *FinishWaybillOrderRequest {
	s.PayCheck = &v
	return s
}

func (s *FinishWaybillOrderRequest) SetPlatformDid(v string) *FinishWaybillOrderRequest {
	s.PlatformDid = &v
	return s
}

func (s *FinishWaybillOrderRequest) SetPrepayments(v string) *FinishWaybillOrderRequest {
	s.Prepayments = &v
	return s
}

func (s *FinishWaybillOrderRequest) SetPrepaymentsBuyEtc(v string) *FinishWaybillOrderRequest {
	s.PrepaymentsBuyEtc = &v
	return s
}

func (s *FinishWaybillOrderRequest) SetPrepaymentsBuyGas(v string) *FinishWaybillOrderRequest {
	s.PrepaymentsBuyGas = &v
	return s
}

func (s *FinishWaybillOrderRequest) SetPrepaymentsBuyOil(v string) *FinishWaybillOrderRequest {
	s.PrepaymentsBuyOil = &v
	return s
}

func (s *FinishWaybillOrderRequest) SetPrepaymentsEtccard(v string) *FinishWaybillOrderRequest {
	s.PrepaymentsEtccard = &v
	return s
}

func (s *FinishWaybillOrderRequest) SetPrepaymentsOilcard(v string) *FinishWaybillOrderRequest {
	s.PrepaymentsOilcard = &v
	return s
}

func (s *FinishWaybillOrderRequest) SetPresentAmountOil(v string) *FinishWaybillOrderRequest {
	s.PresentAmountOil = &v
	return s
}

func (s *FinishWaybillOrderRequest) SetStartAddress(v string) *FinishWaybillOrderRequest {
	s.StartAddress = &v
	return s
}

func (s *FinishWaybillOrderRequest) SetStartCityCode(v string) *FinishWaybillOrderRequest {
	s.StartCityCode = &v
	return s
}

func (s *FinishWaybillOrderRequest) SetStartCityName(v string) *FinishWaybillOrderRequest {
	s.StartCityName = &v
	return s
}

func (s *FinishWaybillOrderRequest) SetStartCountyCode(v string) *FinishWaybillOrderRequest {
	s.StartCountyCode = &v
	return s
}

func (s *FinishWaybillOrderRequest) SetStartCountyName(v string) *FinishWaybillOrderRequest {
	s.StartCountyName = &v
	return s
}

func (s *FinishWaybillOrderRequest) SetStartDivisionCode(v string) *FinishWaybillOrderRequest {
	s.StartDivisionCode = &v
	return s
}

func (s *FinishWaybillOrderRequest) SetStartProvinceCode(v string) *FinishWaybillOrderRequest {
	s.StartProvinceCode = &v
	return s
}

func (s *FinishWaybillOrderRequest) SetStartProvinceName(v string) *FinishWaybillOrderRequest {
	s.StartProvinceName = &v
	return s
}

func (s *FinishWaybillOrderRequest) SetStartStreetCode(v string) *FinishWaybillOrderRequest {
	s.StartStreetCode = &v
	return s
}

func (s *FinishWaybillOrderRequest) SetStartStreetName(v string) *FinishWaybillOrderRequest {
	s.StartStreetName = &v
	return s
}

func (s *FinishWaybillOrderRequest) SetStartTime(v int64) *FinishWaybillOrderRequest {
	s.StartTime = &v
	return s
}

func (s *FinishWaybillOrderRequest) SetTaxWaybillId(v string) *FinishWaybillOrderRequest {
	s.TaxWaybillId = &v
	return s
}

func (s *FinishWaybillOrderRequest) SetTruckLength(v string) *FinishWaybillOrderRequest {
	s.TruckLength = &v
	return s
}

func (s *FinishWaybillOrderRequest) SetTruckType(v string) *FinishWaybillOrderRequest {
	s.TruckType = &v
	return s
}

func (s *FinishWaybillOrderRequest) SetUnitPrice(v string) *FinishWaybillOrderRequest {
	s.UnitPrice = &v
	return s
}

type FinishWaybillOrderResponse struct {
	// 请求唯一ID，用于链路跟踪和问题排查
	ReqMsgId *string `json:"req_msg_id,omitempty" xml:"req_msg_id,omitempty"`
	// 结果码，一般OK表示调用成功
	ResultCode *string `json:"result_code,omitempty" xml:"result_code,omitempty"`
	// 异常信息的文本描述
	ResultMsg *string `json:"result_msg,omitempty" xml:"result_msg,omitempty"`
	// 链上凭证
	//
	//
	TxCode *string `json:"tx_code,omitempty" xml:"tx_code,omitempty"`
}

func (s FinishWaybillOrderResponse) String() string {
	return tea.Prettify(s)
}

func (s FinishWaybillOrderResponse) GoString() string {
	return s.String()
}

func (s *FinishWaybillOrderResponse) SetReqMsgId(v string) *FinishWaybillOrderResponse {
	s.ReqMsgId = &v
	return s
}

func (s *FinishWaybillOrderResponse) SetResultCode(v string) *FinishWaybillOrderResponse {
	s.ResultCode = &v
	return s
}

func (s *FinishWaybillOrderResponse) SetResultMsg(v string) *FinishWaybillOrderResponse {
	s.ResultMsg = &v
	return s
}

func (s *FinishWaybillOrderResponse) SetTxCode(v string) *FinishWaybillOrderResponse {
	s.TxCode = &v
	return s
}

type CreateCargoPayorderRequest struct {
	// OAuth模式下的授权token
	AuthToken         *string `json:"auth_token,omitempty" xml:"auth_token,omitempty"`
	ProductInstanceId *string `json:"product_instance_id,omitempty" xml:"product_instance_id,omitempty"`
	// 收款方银行账号
	BankAccountOfPayee *string `json:"bank_account_of_payee,omitempty" xml:"bank_account_of_payee,omitempty"`
	// 付款方银行账号,货主付款的开户银行账号
	BankAccountOfPayer *string `json:"bank_account_of_payer,omitempty" xml:"bank_account_of_payer,omitempty"`
	// 银行流水号，云资金校验时 必传；
	//
	BankNo *string `json:"bank_no,omitempty" xml:"bank_no,omitempty"`
	// 收款方开户行（平台收款）
	//
	// MYBank_CloudCapital_2,  ## 云资金2.0
	//
	//
	BankOfPayee *string `json:"bank_of_payee,omitempty" xml:"bank_of_payee,omitempty"`
	// 付款方开户行(货主付款)  取值
	//
	// MYBank_CloudCapital_2,  ## 云资金2.0
	BankOfPayer *string `json:"bank_of_payer,omitempty" xml:"bank_of_payer,omitempty"`
	// 账单编号
	BillId *string `json:"bill_id,omitempty" xml:"bill_id,omitempty"`
	// 货源单号
	//
	CargoOrder *string `json:"cargo_order,omitempty" xml:"cargo_order,omitempty" require:"true"`
	// 核验渠道，请按取值约束值填入
	// MYBank_CloudCapital_2,  ## 云资金2.0
	CheckChannel *string `json:"check_channel,omitempty" xml:"check_channel,omitempty"`
	// 币种
	// 按下列取值：
	// CNY、USD
	//
	//
	Currency *string `json:"currency,omitempty" xml:"currency,omitempty"`
	// 费用类型 (运费、调度费、服务费)
	//
	ExpenseType *string `json:"expense_type,omitempty" xml:"expense_type,omitempty" require:"true"`
	// 收款方名称
	Payee *string `json:"payee,omitempty" xml:"payee,omitempty"`
	// 付款方名称
	Payer *string `json:"payer,omitempty" xml:"payer,omitempty"`
	// 费用金额（运单维度）
	//
	Payment *string `json:"payment,omitempty" xml:"payment,omitempty" require:"true"`
	// 是否核验
	PayCheck *bool `json:"pay_check,omitempty" xml:"pay_check,omitempty"`
	// 付款方did
	//
	PayDid *string `json:"pay_did,omitempty" xml:"pay_did,omitempty" require:"true"`
	// 支付单号
	//
	PayId *string `json:"pay_id,omitempty" xml:"pay_id,omitempty" require:"true"`
	// 支付备注
	//
	PayNote *string `json:"pay_note,omitempty" xml:"pay_note,omitempty"`
	// 支付时间
	//
	PayTime *int64 `json:"pay_time,omitempty" xml:"pay_time,omitempty" require:"true"`
	// 支付类型
	//
	PayType *string `json:"pay_type,omitempty" xml:"pay_type,omitempty"`
	// 所属平台did
	PlatformDid *string `json:"platform_did,omitempty" xml:"platform_did,omitempty" require:"true"`
	// 收款方did
	RecvDid *string `json:"recv_did,omitempty" xml:"recv_did,omitempty" require:"true"`
	// 请求唯一标识
	RequestId *string `json:"request_id,omitempty" xml:"request_id,omitempty" require:"true"`
	// 运单号
	TaxWaybillId *string `json:"tax_waybill_id,omitempty" xml:"tax_waybill_id,omitempty" require:"true"`
}

func (s CreateCargoPayorderRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateCargoPayorderRequest) GoString() string {
	return s.String()
}

func (s *CreateCargoPayorderRequest) SetAuthToken(v string) *CreateCargoPayorderRequest {
	s.AuthToken = &v
	return s
}

func (s *CreateCargoPayorderRequest) SetProductInstanceId(v string) *CreateCargoPayorderRequest {
	s.ProductInstanceId = &v
	return s
}

func (s *CreateCargoPayorderRequest) SetBankAccountOfPayee(v string) *CreateCargoPayorderRequest {
	s.BankAccountOfPayee = &v
	return s
}

func (s *CreateCargoPayorderRequest) SetBankAccountOfPayer(v string) *CreateCargoPayorderRequest {
	s.BankAccountOfPayer = &v
	return s
}

func (s *CreateCargoPayorderRequest) SetBankNo(v string) *CreateCargoPayorderRequest {
	s.BankNo = &v
	return s
}

func (s *CreateCargoPayorderRequest) SetBankOfPayee(v string) *CreateCargoPayorderRequest {
	s.BankOfPayee = &v
	return s
}

func (s *CreateCargoPayorderRequest) SetBankOfPayer(v string) *CreateCargoPayorderRequest {
	s.BankOfPayer = &v
	return s
}

func (s *CreateCargoPayorderRequest) SetBillId(v string) *CreateCargoPayorderRequest {
	s.BillId = &v
	return s
}

func (s *CreateCargoPayorderRequest) SetCargoOrder(v string) *CreateCargoPayorderRequest {
	s.CargoOrder = &v
	return s
}

func (s *CreateCargoPayorderRequest) SetCheckChannel(v string) *CreateCargoPayorderRequest {
	s.CheckChannel = &v
	return s
}

func (s *CreateCargoPayorderRequest) SetCurrency(v string) *CreateCargoPayorderRequest {
	s.Currency = &v
	return s
}

func (s *CreateCargoPayorderRequest) SetExpenseType(v string) *CreateCargoPayorderRequest {
	s.ExpenseType = &v
	return s
}

func (s *CreateCargoPayorderRequest) SetPayee(v string) *CreateCargoPayorderRequest {
	s.Payee = &v
	return s
}

func (s *CreateCargoPayorderRequest) SetPayer(v string) *CreateCargoPayorderRequest {
	s.Payer = &v
	return s
}

func (s *CreateCargoPayorderRequest) SetPayment(v string) *CreateCargoPayorderRequest {
	s.Payment = &v
	return s
}

func (s *CreateCargoPayorderRequest) SetPayCheck(v bool) *CreateCargoPayorderRequest {
	s.PayCheck = &v
	return s
}

func (s *CreateCargoPayorderRequest) SetPayDid(v string) *CreateCargoPayorderRequest {
	s.PayDid = &v
	return s
}

func (s *CreateCargoPayorderRequest) SetPayId(v string) *CreateCargoPayorderRequest {
	s.PayId = &v
	return s
}

func (s *CreateCargoPayorderRequest) SetPayNote(v string) *CreateCargoPayorderRequest {
	s.PayNote = &v
	return s
}

func (s *CreateCargoPayorderRequest) SetPayTime(v int64) *CreateCargoPayorderRequest {
	s.PayTime = &v
	return s
}

func (s *CreateCargoPayorderRequest) SetPayType(v string) *CreateCargoPayorderRequest {
	s.PayType = &v
	return s
}

func (s *CreateCargoPayorderRequest) SetPlatformDid(v string) *CreateCargoPayorderRequest {
	s.PlatformDid = &v
	return s
}

func (s *CreateCargoPayorderRequest) SetRecvDid(v string) *CreateCargoPayorderRequest {
	s.RecvDid = &v
	return s
}

func (s *CreateCargoPayorderRequest) SetRequestId(v string) *CreateCargoPayorderRequest {
	s.RequestId = &v
	return s
}

func (s *CreateCargoPayorderRequest) SetTaxWaybillId(v string) *CreateCargoPayorderRequest {
	s.TaxWaybillId = &v
	return s
}

type CreateCargoPayorderResponse struct {
	// 请求唯一ID，用于链路跟踪和问题排查
	ReqMsgId *string `json:"req_msg_id,omitempty" xml:"req_msg_id,omitempty"`
	// 结果码，一般OK表示调用成功
	ResultCode *string `json:"result_code,omitempty" xml:"result_code,omitempty"`
	// 异常信息的文本描述
	ResultMsg *string `json:"result_msg,omitempty" xml:"result_msg,omitempty"`
	// 货源支付链上凭证
	//
	TxCode *string `json:"tx_code,omitempty" xml:"tx_code,omitempty"`
}

func (s CreateCargoPayorderResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateCargoPayorderResponse) GoString() string {
	return s.String()
}

func (s *CreateCargoPayorderResponse) SetReqMsgId(v string) *CreateCargoPayorderResponse {
	s.ReqMsgId = &v
	return s
}

func (s *CreateCargoPayorderResponse) SetResultCode(v string) *CreateCargoPayorderResponse {
	s.ResultCode = &v
	return s
}

func (s *CreateCargoPayorderResponse) SetResultMsg(v string) *CreateCargoPayorderResponse {
	s.ResultMsg = &v
	return s
}

func (s *CreateCargoPayorderResponse) SetTxCode(v string) *CreateCargoPayorderResponse {
	s.TxCode = &v
	return s
}

type SaveWaybillOrderRequest struct {
	// OAuth模式下的授权token
	AuthToken         *string `json:"auth_token,omitempty" xml:"auth_token,omitempty"`
	ProductInstanceId *string `json:"product_instance_id,omitempty" xml:"product_instance_id,omitempty"`
	// 链上账号 与司机did 二选一填写
	//
	AccountId *string `json:"account_id,omitempty" xml:"account_id,omitempty"`
	// 运费，单位（元），平台支付给司机的运费
	AllFreight *string `json:"all_freight,omitempty" xml:"all_freight,omitempty" require:"true"`
	// 回单押金
	//
	BackFee *string `json:"back_fee,omitempty" xml:"back_fee,omitempty"`
	// 业务类型
	BusinessType *string `json:"business_type,omitempty" xml:"business_type,omitempty"`
	// 货物行业编码
	//
	CargoBusinessCode *string `json:"cargo_business_code,omitempty" xml:"cargo_business_code,omitempty"`
	// 货物商品编码
	//
	CargoCode *string `json:"cargo_code,omitempty" xml:"cargo_code,omitempty"`
	// 货物运费险
	//
	CargoInsuranceMoney *string `json:"cargo_insurance_money,omitempty" xml:"cargo_insurance_money,omitempty"`
	// 货源单号
	//
	CargoOrder *string `json:"cargo_order,omitempty" xml:"cargo_order,omitempty"`
	// 货物单位
	//
	CargoUnit *string `json:"cargo_unit,omitempty" xml:"cargo_unit,omitempty"`
	// 货物体积，单位（方）
	//
	CargoVolume *string `json:"cargo_volume,omitempty" xml:"cargo_volume,omitempty"`
	// 货物重量，单位（吨）
	CargoWeight *string `json:"cargo_weight,omitempty" xml:"cargo_weight,omitempty"`
	// 车牌颜色，黄色、蓝色、绿色
	CartBadgeColor *string `json:"cart_badge_color,omitempty" xml:"cart_badge_color,omitempty" require:"true"`
	// 车牌号
	//
	CartBadgeNo *string `json:"cart_badge_no,omitempty" xml:"cart_badge_no,omitempty" require:"true"`
	// 承运商did
	CarCaptainDid *string `json:"car_captain_did,omitempty" xml:"car_captain_did,omitempty"`
	// 货主运费金额，货主支付给平台的运费
	//
	ConsignorFreightAmount *string `json:"consignor_freight_amount,omitempty" xml:"consignor_freight_amount,omitempty" require:"true"`
	// 建单时间，13位毫秒级时间戳
	CreatedTime *int64 `json:"created_time,omitempty" xml:"created_time,omitempty" require:"true"`
	// 到达门点时间，13位毫秒级时间戳
	DestDoorsEndTime *int64 `json:"dest_doors_end_time,omitempty" xml:"dest_doors_end_time,omitempty"`
	// 门点城市CODE，6位区域行政编码
	DoorsCityCode *string `json:"doors_city_code,omitempty" xml:"doors_city_code,omitempty"`
	// 门点城市名称
	DoorsCityName *string `json:"doors_city_name,omitempty" xml:"doors_city_name,omitempty"`
	// 门点区县CODE，6位区域行政编码
	DoorsCountyCode *string `json:"doors_county_code,omitempty" xml:"doors_county_code,omitempty"`
	// 门点区县名称
	DoorsCountyName *string `json:"doors_county_name,omitempty" xml:"doors_county_name,omitempty"`
	// 门点行政区划代码，12位区域行政编码（长途运输须准确提供前6位，如不能提供后6位可补0；短途运输须准确提供12位）
	DoorsDivisionCode *string `json:"doors_division_code,omitempty" xml:"doors_division_code,omitempty"`
	// 门点省份CODE，6位区域行政编码
	DoorsProvinceCode *string `json:"doors_province_code,omitempty" xml:"doors_province_code,omitempty"`
	// 门点省份名称
	DoorsProvinceName *string `json:"doors_province_name,omitempty" xml:"doors_province_name,omitempty"`
	// 发货方名称
	//
	Drawee *string `json:"drawee,omitempty" xml:"drawee,omitempty" require:"true"`
	// 发货方纳税人识别号
	//
	DraweeTaxNo *string `json:"drawee_tax_no,omitempty" xml:"drawee_tax_no,omitempty" require:"true"`
	// 司机分布式数字身份
	//
	DriverDid *string `json:"driver_did,omitempty" xml:"driver_did,omitempty" require:"true"`
	// 司机姓名 已填司机分布式身份的情况下可不填
	//
	DriverName *string `json:"driver_name,omitempty" xml:"driver_name,omitempty"`
	// 目的地详细地址
	EndAddress *string `json:"end_address,omitempty" xml:"end_address,omitempty"`
	// 目的地城市CODE，6位区域行政编码
	//
	EndCityCode *string `json:"end_city_code,omitempty" xml:"end_city_code,omitempty" require:"true"`
	// 目的地城市名称
	//
	EndCityName *string `json:"end_city_name,omitempty" xml:"end_city_name,omitempty" require:"true"`
	// 目的地区县CODE，6位区域行政编码
	EndCountyCode *string `json:"end_county_code,omitempty" xml:"end_county_code,omitempty"`
	// 目的地区县名称
	//
	EndCountyName *string `json:"end_county_name,omitempty" xml:"end_county_name,omitempty"`
	// 结束行政区划代码（长途运输须准确提供前6位，如不能提供后6位可补0；短途运输须准确提供12位）
	//
	EndDivisionCode *string `json:"end_division_code,omitempty" xml:"end_division_code,omitempty" require:"true"`
	// 目的地省份CODE，6位区域行政编码
	EndProvinceCode *string `json:"end_province_code,omitempty" xml:"end_province_code,omitempty" require:"true"`
	// 目的地省份名称
	//
	EndProvinceName *string `json:"end_province_name,omitempty" xml:"end_province_name,omitempty" require:"true"`
	// 目的地街道CODE，12位区域行政编码
	EndStreetCode *string `json:"end_street_code,omitempty" xml:"end_street_code,omitempty"`
	// 目的地街道名称
	EndStreetName *string `json:"end_street_name,omitempty" xml:"end_street_name,omitempty"`
	// 运达时间，13位毫秒级时间戳
	EndTime *int64 `json:"end_time,omitempty" xml:"end_time,omitempty" require:"true"`
	// 运费增项
	//
	FreightIncr *string `json:"freight_incr,omitempty" xml:"freight_incr,omitempty"`
	// 货物数量
	//
	GoodsAmount *string `json:"goods_amount,omitempty" xml:"goods_amount,omitempty"`
	// 货物数量单位类型
	//
	GoodsAmountType *string `json:"goods_amount_type,omitempty" xml:"goods_amount_type,omitempty"`
	// 货物名称
	//
	GoodsName *string `json:"goods_name,omitempty" xml:"goods_name,omitempty" require:"true"`
	// 司机身份证号 已填司机分布式身份的情况下可不填
	//
	IdCard *string `json:"id_card,omitempty" xml:"id_card,omitempty"`
	// 运费扣减
	//
	LossFee *string `json:"loss_fee,omitempty" xml:"loss_fee,omitempty"`
	// 司机手机号 已填司机分布式身份的情况下可不填
	//
	MobileNo *string `json:"mobile_no,omitempty" xml:"mobile_no,omitempty"`
	// 系统识别id 网商识别号
	//
	PartnerId *string `json:"partner_id,omitempty" xml:"partner_id,omitempty"`
	// 是否进行资金验证
	//
	PayCheck *bool `json:"pay_check,omitempty" xml:"pay_check,omitempty"`
	// 无车承运平台分布式数字身份，缺省时为自己的分布式数字身份
	//
	PlatformDid *string `json:"platform_did,omitempty" xml:"platform_did,omitempty" require:"true"`
	// 预付款金额
	//
	Prepayments *string `json:"prepayments,omitempty" xml:"prepayments,omitempty"`
	// 线下预付ETC
	//
	PrepaymentsBuyEtc *string `json:"prepayments_buy_etc,omitempty" xml:"prepayments_buy_etc,omitempty"`
	// 线下气款金额
	//
	PrepaymentsBuyGas *string `json:"prepayments_buy_gas,omitempty" xml:"prepayments_buy_gas,omitempty"`
	// 线下油款预付
	//
	PrepaymentsBuyOil *string `json:"prepayments_buy_oil,omitempty" xml:"prepayments_buy_oil,omitempty"`
	// 预付ETC卡金额
	//
	PrepaymentsEtccard *string `json:"prepayments_etccard,omitempty" xml:"prepayments_etccard,omitempty"`
	// 预付油卡金额
	//
	PrepaymentsOilcard *string `json:"prepayments_oilcard,omitempty" xml:"prepayments_oilcard,omitempty"`
	// 油卡赠送金额
	//
	PresentAmountOil *string `json:"present_amount_oil,omitempty" xml:"present_amount_oil,omitempty"`
	// 起始地详细地址
	StartAddress *string `json:"start_address,omitempty" xml:"start_address,omitempty"`
	// 起始地城市CODE，6位区域行政编码
	//
	StartCityCode *string `json:"start_city_code,omitempty" xml:"start_city_code,omitempty" require:"true"`
	// 起始地城市名称
	//
	StartCityName *string `json:"start_city_name,omitempty" xml:"start_city_name,omitempty" require:"true"`
	// 起始地区县CODE，6位区域行政编码
	//
	//
	StartCountyCode *string `json:"start_county_code,omitempty" xml:"start_county_code,omitempty"`
	// 起始地区县名称
	//
	StartCountyName *string `json:"start_county_name,omitempty" xml:"start_county_name,omitempty"`
	// 起始行政区划代码（长途运输须准确提供前6位，如不能提供后6位可补0；短途运输须准确提供12位）
	//
	StartDivisionCode *string `json:"start_division_code,omitempty" xml:"start_division_code,omitempty" require:"true"`
	// 起始地省份CODE，6位区域行政编码
	//
	//
	StartProvinceCode *string `json:"start_province_code,omitempty" xml:"start_province_code,omitempty" require:"true"`
	// 起始地省份名称
	//
	StartProvinceName *string `json:"start_province_name,omitempty" xml:"start_province_name,omitempty" require:"true"`
	// 起始地街道CODE，12位区域行政编码
	StartStreetCode *string `json:"start_street_code,omitempty" xml:"start_street_code,omitempty"`
	// 起始地街道名称
	StartStreetName *string `json:"start_street_name,omitempty" xml:"start_street_name,omitempty"`
	// 起运时间，13位毫秒级时间戳
	StartTime *int64 `json:"start_time,omitempty" xml:"start_time,omitempty" require:"true"`
	// 运单ID
	//
	TaxWaybillId *string `json:"tax_waybill_id,omitempty" xml:"tax_waybill_id,omitempty" require:"true"`
	// 车长，可以填：不限车长或者1.8，2.7，3.8，4.2，5，6.2，6.8，7.7，8.2，8.7，9.6，11.7，12.5，13，13.7，15，16，17.5等不超过2位小数的数字
	TruckLength *string `json:"truck_length,omitempty" xml:"truck_length,omitempty"`
	// 车型
	TruckType *string `json:"truck_type,omitempty" xml:"truck_type,omitempty"`
	// 运输单价
	//
	UnitPrice *string `json:"unit_price,omitempty" xml:"unit_price,omitempty"`
}

func (s SaveWaybillOrderRequest) String() string {
	return tea.Prettify(s)
}

func (s SaveWaybillOrderRequest) GoString() string {
	return s.String()
}

func (s *SaveWaybillOrderRequest) SetAuthToken(v string) *SaveWaybillOrderRequest {
	s.AuthToken = &v
	return s
}

func (s *SaveWaybillOrderRequest) SetProductInstanceId(v string) *SaveWaybillOrderRequest {
	s.ProductInstanceId = &v
	return s
}

func (s *SaveWaybillOrderRequest) SetAccountId(v string) *SaveWaybillOrderRequest {
	s.AccountId = &v
	return s
}

func (s *SaveWaybillOrderRequest) SetAllFreight(v string) *SaveWaybillOrderRequest {
	s.AllFreight = &v
	return s
}

func (s *SaveWaybillOrderRequest) SetBackFee(v string) *SaveWaybillOrderRequest {
	s.BackFee = &v
	return s
}

func (s *SaveWaybillOrderRequest) SetBusinessType(v string) *SaveWaybillOrderRequest {
	s.BusinessType = &v
	return s
}

func (s *SaveWaybillOrderRequest) SetCargoBusinessCode(v string) *SaveWaybillOrderRequest {
	s.CargoBusinessCode = &v
	return s
}

func (s *SaveWaybillOrderRequest) SetCargoCode(v string) *SaveWaybillOrderRequest {
	s.CargoCode = &v
	return s
}

func (s *SaveWaybillOrderRequest) SetCargoInsuranceMoney(v string) *SaveWaybillOrderRequest {
	s.CargoInsuranceMoney = &v
	return s
}

func (s *SaveWaybillOrderRequest) SetCargoOrder(v string) *SaveWaybillOrderRequest {
	s.CargoOrder = &v
	return s
}

func (s *SaveWaybillOrderRequest) SetCargoUnit(v string) *SaveWaybillOrderRequest {
	s.CargoUnit = &v
	return s
}

func (s *SaveWaybillOrderRequest) SetCargoVolume(v string) *SaveWaybillOrderRequest {
	s.CargoVolume = &v
	return s
}

func (s *SaveWaybillOrderRequest) SetCargoWeight(v string) *SaveWaybillOrderRequest {
	s.CargoWeight = &v
	return s
}

func (s *SaveWaybillOrderRequest) SetCartBadgeColor(v string) *SaveWaybillOrderRequest {
	s.CartBadgeColor = &v
	return s
}

func (s *SaveWaybillOrderRequest) SetCartBadgeNo(v string) *SaveWaybillOrderRequest {
	s.CartBadgeNo = &v
	return s
}

func (s *SaveWaybillOrderRequest) SetCarCaptainDid(v string) *SaveWaybillOrderRequest {
	s.CarCaptainDid = &v
	return s
}

func (s *SaveWaybillOrderRequest) SetConsignorFreightAmount(v string) *SaveWaybillOrderRequest {
	s.ConsignorFreightAmount = &v
	return s
}

func (s *SaveWaybillOrderRequest) SetCreatedTime(v int64) *SaveWaybillOrderRequest {
	s.CreatedTime = &v
	return s
}

func (s *SaveWaybillOrderRequest) SetDestDoorsEndTime(v int64) *SaveWaybillOrderRequest {
	s.DestDoorsEndTime = &v
	return s
}

func (s *SaveWaybillOrderRequest) SetDoorsCityCode(v string) *SaveWaybillOrderRequest {
	s.DoorsCityCode = &v
	return s
}

func (s *SaveWaybillOrderRequest) SetDoorsCityName(v string) *SaveWaybillOrderRequest {
	s.DoorsCityName = &v
	return s
}

func (s *SaveWaybillOrderRequest) SetDoorsCountyCode(v string) *SaveWaybillOrderRequest {
	s.DoorsCountyCode = &v
	return s
}

func (s *SaveWaybillOrderRequest) SetDoorsCountyName(v string) *SaveWaybillOrderRequest {
	s.DoorsCountyName = &v
	return s
}

func (s *SaveWaybillOrderRequest) SetDoorsDivisionCode(v string) *SaveWaybillOrderRequest {
	s.DoorsDivisionCode = &v
	return s
}

func (s *SaveWaybillOrderRequest) SetDoorsProvinceCode(v string) *SaveWaybillOrderRequest {
	s.DoorsProvinceCode = &v
	return s
}

func (s *SaveWaybillOrderRequest) SetDoorsProvinceName(v string) *SaveWaybillOrderRequest {
	s.DoorsProvinceName = &v
	return s
}

func (s *SaveWaybillOrderRequest) SetDrawee(v string) *SaveWaybillOrderRequest {
	s.Drawee = &v
	return s
}

func (s *SaveWaybillOrderRequest) SetDraweeTaxNo(v string) *SaveWaybillOrderRequest {
	s.DraweeTaxNo = &v
	return s
}

func (s *SaveWaybillOrderRequest) SetDriverDid(v string) *SaveWaybillOrderRequest {
	s.DriverDid = &v
	return s
}

func (s *SaveWaybillOrderRequest) SetDriverName(v string) *SaveWaybillOrderRequest {
	s.DriverName = &v
	return s
}

func (s *SaveWaybillOrderRequest) SetEndAddress(v string) *SaveWaybillOrderRequest {
	s.EndAddress = &v
	return s
}

func (s *SaveWaybillOrderRequest) SetEndCityCode(v string) *SaveWaybillOrderRequest {
	s.EndCityCode = &v
	return s
}

func (s *SaveWaybillOrderRequest) SetEndCityName(v string) *SaveWaybillOrderRequest {
	s.EndCityName = &v
	return s
}

func (s *SaveWaybillOrderRequest) SetEndCountyCode(v string) *SaveWaybillOrderRequest {
	s.EndCountyCode = &v
	return s
}

func (s *SaveWaybillOrderRequest) SetEndCountyName(v string) *SaveWaybillOrderRequest {
	s.EndCountyName = &v
	return s
}

func (s *SaveWaybillOrderRequest) SetEndDivisionCode(v string) *SaveWaybillOrderRequest {
	s.EndDivisionCode = &v
	return s
}

func (s *SaveWaybillOrderRequest) SetEndProvinceCode(v string) *SaveWaybillOrderRequest {
	s.EndProvinceCode = &v
	return s
}

func (s *SaveWaybillOrderRequest) SetEndProvinceName(v string) *SaveWaybillOrderRequest {
	s.EndProvinceName = &v
	return s
}

func (s *SaveWaybillOrderRequest) SetEndStreetCode(v string) *SaveWaybillOrderRequest {
	s.EndStreetCode = &v
	return s
}

func (s *SaveWaybillOrderRequest) SetEndStreetName(v string) *SaveWaybillOrderRequest {
	s.EndStreetName = &v
	return s
}

func (s *SaveWaybillOrderRequest) SetEndTime(v int64) *SaveWaybillOrderRequest {
	s.EndTime = &v
	return s
}

func (s *SaveWaybillOrderRequest) SetFreightIncr(v string) *SaveWaybillOrderRequest {
	s.FreightIncr = &v
	return s
}

func (s *SaveWaybillOrderRequest) SetGoodsAmount(v string) *SaveWaybillOrderRequest {
	s.GoodsAmount = &v
	return s
}

func (s *SaveWaybillOrderRequest) SetGoodsAmountType(v string) *SaveWaybillOrderRequest {
	s.GoodsAmountType = &v
	return s
}

func (s *SaveWaybillOrderRequest) SetGoodsName(v string) *SaveWaybillOrderRequest {
	s.GoodsName = &v
	return s
}

func (s *SaveWaybillOrderRequest) SetIdCard(v string) *SaveWaybillOrderRequest {
	s.IdCard = &v
	return s
}

func (s *SaveWaybillOrderRequest) SetLossFee(v string) *SaveWaybillOrderRequest {
	s.LossFee = &v
	return s
}

func (s *SaveWaybillOrderRequest) SetMobileNo(v string) *SaveWaybillOrderRequest {
	s.MobileNo = &v
	return s
}

func (s *SaveWaybillOrderRequest) SetPartnerId(v string) *SaveWaybillOrderRequest {
	s.PartnerId = &v
	return s
}

func (s *SaveWaybillOrderRequest) SetPayCheck(v bool) *SaveWaybillOrderRequest {
	s.PayCheck = &v
	return s
}

func (s *SaveWaybillOrderRequest) SetPlatformDid(v string) *SaveWaybillOrderRequest {
	s.PlatformDid = &v
	return s
}

func (s *SaveWaybillOrderRequest) SetPrepayments(v string) *SaveWaybillOrderRequest {
	s.Prepayments = &v
	return s
}

func (s *SaveWaybillOrderRequest) SetPrepaymentsBuyEtc(v string) *SaveWaybillOrderRequest {
	s.PrepaymentsBuyEtc = &v
	return s
}

func (s *SaveWaybillOrderRequest) SetPrepaymentsBuyGas(v string) *SaveWaybillOrderRequest {
	s.PrepaymentsBuyGas = &v
	return s
}

func (s *SaveWaybillOrderRequest) SetPrepaymentsBuyOil(v string) *SaveWaybillOrderRequest {
	s.PrepaymentsBuyOil = &v
	return s
}

func (s *SaveWaybillOrderRequest) SetPrepaymentsEtccard(v string) *SaveWaybillOrderRequest {
	s.PrepaymentsEtccard = &v
	return s
}

func (s *SaveWaybillOrderRequest) SetPrepaymentsOilcard(v string) *SaveWaybillOrderRequest {
	s.PrepaymentsOilcard = &v
	return s
}

func (s *SaveWaybillOrderRequest) SetPresentAmountOil(v string) *SaveWaybillOrderRequest {
	s.PresentAmountOil = &v
	return s
}

func (s *SaveWaybillOrderRequest) SetStartAddress(v string) *SaveWaybillOrderRequest {
	s.StartAddress = &v
	return s
}

func (s *SaveWaybillOrderRequest) SetStartCityCode(v string) *SaveWaybillOrderRequest {
	s.StartCityCode = &v
	return s
}

func (s *SaveWaybillOrderRequest) SetStartCityName(v string) *SaveWaybillOrderRequest {
	s.StartCityName = &v
	return s
}

func (s *SaveWaybillOrderRequest) SetStartCountyCode(v string) *SaveWaybillOrderRequest {
	s.StartCountyCode = &v
	return s
}

func (s *SaveWaybillOrderRequest) SetStartCountyName(v string) *SaveWaybillOrderRequest {
	s.StartCountyName = &v
	return s
}

func (s *SaveWaybillOrderRequest) SetStartDivisionCode(v string) *SaveWaybillOrderRequest {
	s.StartDivisionCode = &v
	return s
}

func (s *SaveWaybillOrderRequest) SetStartProvinceCode(v string) *SaveWaybillOrderRequest {
	s.StartProvinceCode = &v
	return s
}

func (s *SaveWaybillOrderRequest) SetStartProvinceName(v string) *SaveWaybillOrderRequest {
	s.StartProvinceName = &v
	return s
}

func (s *SaveWaybillOrderRequest) SetStartStreetCode(v string) *SaveWaybillOrderRequest {
	s.StartStreetCode = &v
	return s
}

func (s *SaveWaybillOrderRequest) SetStartStreetName(v string) *SaveWaybillOrderRequest {
	s.StartStreetName = &v
	return s
}

func (s *SaveWaybillOrderRequest) SetStartTime(v int64) *SaveWaybillOrderRequest {
	s.StartTime = &v
	return s
}

func (s *SaveWaybillOrderRequest) SetTaxWaybillId(v string) *SaveWaybillOrderRequest {
	s.TaxWaybillId = &v
	return s
}

func (s *SaveWaybillOrderRequest) SetTruckLength(v string) *SaveWaybillOrderRequest {
	s.TruckLength = &v
	return s
}

func (s *SaveWaybillOrderRequest) SetTruckType(v string) *SaveWaybillOrderRequest {
	s.TruckType = &v
	return s
}

func (s *SaveWaybillOrderRequest) SetUnitPrice(v string) *SaveWaybillOrderRequest {
	s.UnitPrice = &v
	return s
}

type SaveWaybillOrderResponse struct {
	// 请求唯一ID，用于链路跟踪和问题排查
	ReqMsgId *string `json:"req_msg_id,omitempty" xml:"req_msg_id,omitempty"`
	// 结果码，一般OK表示调用成功
	ResultCode *string `json:"result_code,omitempty" xml:"result_code,omitempty"`
	// 异常信息的文本描述
	ResultMsg *string `json:"result_msg,omitempty" xml:"result_msg,omitempty"`
	// 链上凭证
	//
	TxCode *string `json:"tx_code,omitempty" xml:"tx_code,omitempty"`
}

func (s SaveWaybillOrderResponse) String() string {
	return tea.Prettify(s)
}

func (s SaveWaybillOrderResponse) GoString() string {
	return s.String()
}

func (s *SaveWaybillOrderResponse) SetReqMsgId(v string) *SaveWaybillOrderResponse {
	s.ReqMsgId = &v
	return s
}

func (s *SaveWaybillOrderResponse) SetResultCode(v string) *SaveWaybillOrderResponse {
	s.ResultCode = &v
	return s
}

func (s *SaveWaybillOrderResponse) SetResultMsg(v string) *SaveWaybillOrderResponse {
	s.ResultMsg = &v
	return s
}

func (s *SaveWaybillOrderResponse) SetTxCode(v string) *SaveWaybillOrderResponse {
	s.TxCode = &v
	return s
}

type CloseWaybillOrderRequest struct {
	// OAuth模式下的授权token
	AuthToken         *string `json:"auth_token,omitempty" xml:"auth_token,omitempty"`
	ProductInstanceId *string `json:"product_instance_id,omitempty" xml:"product_instance_id,omitempty"`
	// 无车承运平台分布式数字身份
	PlatformDid *string `json:"platform_did,omitempty" xml:"platform_did,omitempty" require:"true"`
	// 运单ID
	//
	TaxWaybillId *string `json:"tax_waybill_id,omitempty" xml:"tax_waybill_id,omitempty" require:"true"`
}

func (s CloseWaybillOrderRequest) String() string {
	return tea.Prettify(s)
}

func (s CloseWaybillOrderRequest) GoString() string {
	return s.String()
}

func (s *CloseWaybillOrderRequest) SetAuthToken(v string) *CloseWaybillOrderRequest {
	s.AuthToken = &v
	return s
}

func (s *CloseWaybillOrderRequest) SetProductInstanceId(v string) *CloseWaybillOrderRequest {
	s.ProductInstanceId = &v
	return s
}

func (s *CloseWaybillOrderRequest) SetPlatformDid(v string) *CloseWaybillOrderRequest {
	s.PlatformDid = &v
	return s
}

func (s *CloseWaybillOrderRequest) SetTaxWaybillId(v string) *CloseWaybillOrderRequest {
	s.TaxWaybillId = &v
	return s
}

type CloseWaybillOrderResponse struct {
	// 请求唯一ID，用于链路跟踪和问题排查
	ReqMsgId *string `json:"req_msg_id,omitempty" xml:"req_msg_id,omitempty"`
	// 结果码，一般OK表示调用成功
	ResultCode *string `json:"result_code,omitempty" xml:"result_code,omitempty"`
	// 异常信息的文本描述
	ResultMsg *string `json:"result_msg,omitempty" xml:"result_msg,omitempty"`
	// 链上凭证
	//
	TxCode *string `json:"tx_code,omitempty" xml:"tx_code,omitempty"`
}

func (s CloseWaybillOrderResponse) String() string {
	return tea.Prettify(s)
}

func (s CloseWaybillOrderResponse) GoString() string {
	return s.String()
}

func (s *CloseWaybillOrderResponse) SetReqMsgId(v string) *CloseWaybillOrderResponse {
	s.ReqMsgId = &v
	return s
}

func (s *CloseWaybillOrderResponse) SetResultCode(v string) *CloseWaybillOrderResponse {
	s.ResultCode = &v
	return s
}

func (s *CloseWaybillOrderResponse) SetResultMsg(v string) *CloseWaybillOrderResponse {
	s.ResultMsg = &v
	return s
}

func (s *CloseWaybillOrderResponse) SetTxCode(v string) *CloseWaybillOrderResponse {
	s.TxCode = &v
	return s
}

type FinishFinanceWaybillRequest struct {
	// OAuth模式下的授权token
	AuthToken         *string `json:"auth_token,omitempty" xml:"auth_token,omitempty"`
	ProductInstanceId *string `json:"product_instance_id,omitempty" xml:"product_instance_id,omitempty"`
	// 运费
	AllFreight *string `json:"all_freight,omitempty" xml:"all_freight,omitempty"`
	// 回单押金
	BackFee *string `json:"back_fee,omitempty" xml:"back_fee,omitempty"`
	// 货主支付运费金额
	ConsignorFreightAmount *string `json:"consignor_freight_amount,omitempty" xml:"consignor_freight_amount,omitempty" require:"true"`
	// 运费增项
	FreightIncr *string `json:"freight_incr,omitempty" xml:"freight_incr,omitempty"`
	// 运费扣减
	LossFee *string `json:"loss_fee,omitempty" xml:"loss_fee,omitempty"`
	// 无车承运平台分布式数字身份，缺省时为自己的分布式数字身份
	PlatformDid *string `json:"platform_did,omitempty" xml:"platform_did,omitempty" require:"true"`
	// 运单id
	TaxWaybillId *string `json:"tax_waybill_id,omitempty" xml:"tax_waybill_id,omitempty" require:"true"`
}

func (s FinishFinanceWaybillRequest) String() string {
	return tea.Prettify(s)
}

func (s FinishFinanceWaybillRequest) GoString() string {
	return s.String()
}

func (s *FinishFinanceWaybillRequest) SetAuthToken(v string) *FinishFinanceWaybillRequest {
	s.AuthToken = &v
	return s
}

func (s *FinishFinanceWaybillRequest) SetProductInstanceId(v string) *FinishFinanceWaybillRequest {
	s.ProductInstanceId = &v
	return s
}

func (s *FinishFinanceWaybillRequest) SetAllFreight(v string) *FinishFinanceWaybillRequest {
	s.AllFreight = &v
	return s
}

func (s *FinishFinanceWaybillRequest) SetBackFee(v string) *FinishFinanceWaybillRequest {
	s.BackFee = &v
	return s
}

func (s *FinishFinanceWaybillRequest) SetConsignorFreightAmount(v string) *FinishFinanceWaybillRequest {
	s.ConsignorFreightAmount = &v
	return s
}

func (s *FinishFinanceWaybillRequest) SetFreightIncr(v string) *FinishFinanceWaybillRequest {
	s.FreightIncr = &v
	return s
}

func (s *FinishFinanceWaybillRequest) SetLossFee(v string) *FinishFinanceWaybillRequest {
	s.LossFee = &v
	return s
}

func (s *FinishFinanceWaybillRequest) SetPlatformDid(v string) *FinishFinanceWaybillRequest {
	s.PlatformDid = &v
	return s
}

func (s *FinishFinanceWaybillRequest) SetTaxWaybillId(v string) *FinishFinanceWaybillRequest {
	s.TaxWaybillId = &v
	return s
}

type FinishFinanceWaybillResponse struct {
	// 请求唯一ID，用于链路跟踪和问题排查
	ReqMsgId *string `json:"req_msg_id,omitempty" xml:"req_msg_id,omitempty"`
	// 结果码，一般OK表示调用成功
	ResultCode *string `json:"result_code,omitempty" xml:"result_code,omitempty"`
	// 异常信息的文本描述
	ResultMsg *string `json:"result_msg,omitempty" xml:"result_msg,omitempty"`
	// 链上凭证
	TxCode *string `json:"tx_code,omitempty" xml:"tx_code,omitempty"`
}

func (s FinishFinanceWaybillResponse) String() string {
	return tea.Prettify(s)
}

func (s FinishFinanceWaybillResponse) GoString() string {
	return s.String()
}

func (s *FinishFinanceWaybillResponse) SetReqMsgId(v string) *FinishFinanceWaybillResponse {
	s.ReqMsgId = &v
	return s
}

func (s *FinishFinanceWaybillResponse) SetResultCode(v string) *FinishFinanceWaybillResponse {
	s.ResultCode = &v
	return s
}

func (s *FinishFinanceWaybillResponse) SetResultMsg(v string) *FinishFinanceWaybillResponse {
	s.ResultMsg = &v
	return s
}

func (s *FinishFinanceWaybillResponse) SetTxCode(v string) *FinishFinanceWaybillResponse {
	s.TxCode = &v
	return s
}

type FinishFinanceTransportRequest struct {
	// OAuth模式下的授权token
	AuthToken         *string `json:"auth_token,omitempty" xml:"auth_token,omitempty"`
	ProductInstanceId *string `json:"product_instance_id,omitempty" xml:"product_instance_id,omitempty"`
	// 运费
	AllFreight *string `json:"all_freight,omitempty" xml:"all_freight,omitempty"`
	// 回单押金
	BackFee *string `json:"back_fee,omitempty" xml:"back_fee,omitempty"`
	// 货主支付运费金额
	ConsignorFreightAmount *string `json:"consignor_freight_amount,omitempty" xml:"consignor_freight_amount,omitempty" require:"true"`
	// 到达门点时间
	DestDoorsEndTime *int64 `json:"dest_doors_end_time,omitempty" xml:"dest_doors_end_time,omitempty"`
	// 终结时间
	EndTime *int64 `json:"end_time,omitempty" xml:"end_time,omitempty" require:"true"`
	// 运费增项
	FreightIncr *string `json:"freight_incr,omitempty" xml:"freight_incr,omitempty"`
	// 运费扣减
	LossFee *string `json:"loss_fee,omitempty" xml:"loss_fee,omitempty"`
	// 无车承运平台分布式数字身份，缺省时为自己的分布式数字身份
	PlatformDid *string `json:"platform_did,omitempty" xml:"platform_did,omitempty" require:"true"`
	// 运单id
	TaxWaybillId *string `json:"tax_waybill_id,omitempty" xml:"tax_waybill_id,omitempty" require:"true"`
}

func (s FinishFinanceTransportRequest) String() string {
	return tea.Prettify(s)
}

func (s FinishFinanceTransportRequest) GoString() string {
	return s.String()
}

func (s *FinishFinanceTransportRequest) SetAuthToken(v string) *FinishFinanceTransportRequest {
	s.AuthToken = &v
	return s
}

func (s *FinishFinanceTransportRequest) SetProductInstanceId(v string) *FinishFinanceTransportRequest {
	s.ProductInstanceId = &v
	return s
}

func (s *FinishFinanceTransportRequest) SetAllFreight(v string) *FinishFinanceTransportRequest {
	s.AllFreight = &v
	return s
}

func (s *FinishFinanceTransportRequest) SetBackFee(v string) *FinishFinanceTransportRequest {
	s.BackFee = &v
	return s
}

func (s *FinishFinanceTransportRequest) SetConsignorFreightAmount(v string) *FinishFinanceTransportRequest {
	s.ConsignorFreightAmount = &v
	return s
}

func (s *FinishFinanceTransportRequest) SetDestDoorsEndTime(v int64) *FinishFinanceTransportRequest {
	s.DestDoorsEndTime = &v
	return s
}

func (s *FinishFinanceTransportRequest) SetEndTime(v int64) *FinishFinanceTransportRequest {
	s.EndTime = &v
	return s
}

func (s *FinishFinanceTransportRequest) SetFreightIncr(v string) *FinishFinanceTransportRequest {
	s.FreightIncr = &v
	return s
}

func (s *FinishFinanceTransportRequest) SetLossFee(v string) *FinishFinanceTransportRequest {
	s.LossFee = &v
	return s
}

func (s *FinishFinanceTransportRequest) SetPlatformDid(v string) *FinishFinanceTransportRequest {
	s.PlatformDid = &v
	return s
}

func (s *FinishFinanceTransportRequest) SetTaxWaybillId(v string) *FinishFinanceTransportRequest {
	s.TaxWaybillId = &v
	return s
}

type FinishFinanceTransportResponse struct {
	// 请求唯一ID，用于链路跟踪和问题排查
	ReqMsgId *string `json:"req_msg_id,omitempty" xml:"req_msg_id,omitempty"`
	// 结果码，一般OK表示调用成功
	ResultCode *string `json:"result_code,omitempty" xml:"result_code,omitempty"`
	// 异常信息的文本描述
	ResultMsg *string `json:"result_msg,omitempty" xml:"result_msg,omitempty"`
	// 货源支付链上凭证
	//
	TxCode *string `json:"tx_code,omitempty" xml:"tx_code,omitempty"`
}

func (s FinishFinanceTransportResponse) String() string {
	return tea.Prettify(s)
}

func (s FinishFinanceTransportResponse) GoString() string {
	return s.String()
}

func (s *FinishFinanceTransportResponse) SetReqMsgId(v string) *FinishFinanceTransportResponse {
	s.ReqMsgId = &v
	return s
}

func (s *FinishFinanceTransportResponse) SetResultCode(v string) *FinishFinanceTransportResponse {
	s.ResultCode = &v
	return s
}

func (s *FinishFinanceTransportResponse) SetResultMsg(v string) *FinishFinanceTransportResponse {
	s.ResultMsg = &v
	return s
}

func (s *FinishFinanceTransportResponse) SetTxCode(v string) *FinishFinanceTransportResponse {
	s.TxCode = &v
	return s
}

type UpdateFinanceWaybillRequest struct {
	// OAuth模式下的授权token
	AuthToken         *string `json:"auth_token,omitempty" xml:"auth_token,omitempty"`
	ProductInstanceId *string `json:"product_instance_id,omitempty" xml:"product_instance_id,omitempty"`
	// 运费，单位（元），平台支付给司机的运费
	AllFreight *string `json:"all_freight,omitempty" xml:"all_freight,omitempty"`
	// 回单押金
	BackFee *string `json:"back_fee,omitempty" xml:"back_fee,omitempty"`
	// 业务类型
	BusinessType *string `json:"business_type,omitempty" xml:"business_type,omitempty"`
	// 货物行业编码
	CargoBusinessCode *string `json:"cargo_business_code,omitempty" xml:"cargo_business_code,omitempty"`
	// 货物商品编码
	CargoCode *string `json:"cargo_code,omitempty" xml:"cargo_code,omitempty"`
	// 货物运费险
	CargoInsuranceMoney *string `json:"cargo_insurance_money,omitempty" xml:"cargo_insurance_money,omitempty"`
	// 货源单号
	CargoOrder *string `json:"cargo_order,omitempty" xml:"cargo_order,omitempty"`
	// 货物单位
	CargoUnit *string `json:"cargo_unit,omitempty" xml:"cargo_unit,omitempty"`
	// 货物体积，单位（方）
	CargoVolume *string `json:"cargo_volume,omitempty" xml:"cargo_volume,omitempty"`
	// 货物重量，单位（吨）
	CargoWeight *string `json:"cargo_weight,omitempty" xml:"cargo_weight,omitempty"`
	// 车牌颜色，黄色、蓝色、绿色
	CartBadgeColor *string `json:"cart_badge_color,omitempty" xml:"cart_badge_color,omitempty"`
	// 车牌号
	CartBadgeNo *string `json:"cart_badge_no,omitempty" xml:"cart_badge_no,omitempty"`
	// 承运商did
	CarCaptainDid *string `json:"car_captain_did,omitempty" xml:"car_captain_did,omitempty"`
	// 货主支付运费金额，货主支付给平台的运费金额
	ConsignorFreightAmount *string `json:"consignor_freight_amount,omitempty" xml:"consignor_freight_amount,omitempty"`
	// 建单时间，13位毫秒级时间戳
	CreatedTime *int64 `json:"created_time,omitempty" xml:"created_time,omitempty"`
	// 到达门点时间，13位毫秒级时间戳
	DestDoorsEndTime *string `json:"dest_doors_end_time,omitempty" xml:"dest_doors_end_time,omitempty"`
	// 门点城市CODE，6位区域行政编码
	DoorsCityCode *string `json:"doors_city_code,omitempty" xml:"doors_city_code,omitempty"`
	// 门点城市名称
	DoorsCityName *string `json:"doors_city_name,omitempty" xml:"doors_city_name,omitempty"`
	// 门点区县CODE，6位区域行政编码
	DoorsCountyCode *string `json:"doors_county_code,omitempty" xml:"doors_county_code,omitempty"`
	// 门点区县名称
	DoorsCountyName *string `json:"doors_county_name,omitempty" xml:"doors_county_name,omitempty"`
	// 门点行政区划代码，12位区域行政编码，（长途运输须准确提供前6位，如不能提供后6位可补0；短途运输须准确提供12位）
	DoorsDivisionCode *string `json:"doors_division_code,omitempty" xml:"doors_division_code,omitempty"`
	// 门点省份CODE，6位行政区域编码
	DoorsProvinceCode *string `json:"doors_province_code,omitempty" xml:"doors_province_code,omitempty"`
	// 门点省份名称
	DoorsProvinceName *string `json:"doors_province_name,omitempty" xml:"doors_province_name,omitempty"`
	// 发货方名称
	Drawee *string `json:"drawee,omitempty" xml:"drawee,omitempty"`
	// 发货方纳税人识别号
	DraweeTaxNo *string `json:"drawee_tax_no,omitempty" xml:"drawee_tax_no,omitempty"`
	// 司机分布式数字身份
	DriverDid *string `json:"driver_did,omitempty" xml:"driver_did,omitempty"`
	// 目的地详细地址
	EndAddress *string `json:"end_address,omitempty" xml:"end_address,omitempty"`
	// 目的地城市CODE，6位区域行政编码
	EndCityCode *string `json:"end_city_code,omitempty" xml:"end_city_code,omitempty"`
	// 目的地城市名称
	EndCityName *string `json:"end_city_name,omitempty" xml:"end_city_name,omitempty"`
	// 目的地区县CODE
	EndCountyCode *string `json:"end_county_code,omitempty" xml:"end_county_code,omitempty"`
	// 目的地区县名称，6位区域行政编码
	EndCountyName *string `json:"end_county_name,omitempty" xml:"end_county_name,omitempty"`
	// 结束行政区划代码，12位区域行政编码，（长途运输须准确提供前6位，如不能提供后6位可补0；短途运输须准确提供12位）
	EndDivisionCode *string `json:"end_division_code,omitempty" xml:"end_division_code,omitempty"`
	// 目的地省份CODE，6位区域行政编码
	EndProvinceCode *string `json:"end_province_code,omitempty" xml:"end_province_code,omitempty"`
	// 目的地省份名称
	EndProvinceName *string `json:"end_province_name,omitempty" xml:"end_province_name,omitempty"`
	// 目的地街道CODE，12位区域行政编码
	EndStreetCode *string `json:"end_street_code,omitempty" xml:"end_street_code,omitempty"`
	// 目的地街道名称
	EndStreetName *string `json:"end_street_name,omitempty" xml:"end_street_name,omitempty"`
	// 终结时间，13位毫秒级时间戳
	EndTime *int64 `json:"end_time,omitempty" xml:"end_time,omitempty"`
	// 运费增项
	FreightIncr *string `json:"freight_incr,omitempty" xml:"freight_incr,omitempty"`
	// 货物数量
	GoodsAmount *int64 `json:"goods_amount,omitempty" xml:"goods_amount,omitempty"`
	// 货物数量单位类型
	GoodsAmountType *string `json:"goods_amount_type,omitempty" xml:"goods_amount_type,omitempty"`
	// 货物名称
	GoodsName *string `json:"goods_name,omitempty" xml:"goods_name,omitempty"`
	// 运费扣减
	LossFee *string `json:"loss_fee,omitempty" xml:"loss_fee,omitempty"`
	// 网商识别号
	PartnerId *string `json:"partner_id,omitempty" xml:"partner_id,omitempty"`
	// 无车承运平台分布式数字身份，缺省时为自己的分布式数字身份
	//
	PlatformDid *string `json:"platform_did,omitempty" xml:"platform_did,omitempty" require:"true"`
	// 预付款金额
	Prepayments *string `json:"prepayments,omitempty" xml:"prepayments,omitempty"`
	// 线下预付ETC
	PrepaymentsBuyEtc *string `json:"prepayments_buy_etc,omitempty" xml:"prepayments_buy_etc,omitempty"`
	// 线下气款金额
	PrepaymentsBuyGas *string `json:"prepayments_buy_gas,omitempty" xml:"prepayments_buy_gas,omitempty"`
	// 线下油款预付
	PrepaymentsBuyOil *string `json:"prepayments_buy_oil,omitempty" xml:"prepayments_buy_oil,omitempty"`
	// 预付ETC卡金额
	PrepaymentsEtccard *string `json:"prepayments_etccard,omitempty" xml:"prepayments_etccard,omitempty"`
	// 预付油卡金额
	PrepaymentsOilcard *string `json:"prepayments_oilcard,omitempty" xml:"prepayments_oilcard,omitempty"`
	// 油卡赠送金额
	PresentAmountOil *string `json:"present_amount_oil,omitempty" xml:"present_amount_oil,omitempty"`
	// 起始地详细地址
	StartAddress *string `json:"start_address,omitempty" xml:"start_address,omitempty"`
	// 起始地CODE，6位区域行政编码
	StartCityCode *string `json:"start_city_code,omitempty" xml:"start_city_code,omitempty"`
	// 起始地城市名称
	StartCityName *string `json:"start_city_name,omitempty" xml:"start_city_name,omitempty"`
	// 起始地区县CODE，6位区域行政编码
	StartCountyCode *string `json:"start_county_code,omitempty" xml:"start_county_code,omitempty"`
	// 起始地区县名称
	StartCountyName *string `json:"start_county_name,omitempty" xml:"start_county_name,omitempty"`
	// 起始行政区划代码，12位区域行政编码（长途运输须准确提供前6位，如不能提供后6位可补0；短途运输须准确提供12位）
	StartDivisionCode *string `json:"start_division_code,omitempty" xml:"start_division_code,omitempty"`
	// 起始地省份CODE，6位区域行政编码
	StartProvinceCode *string `json:"start_province_code,omitempty" xml:"start_province_code,omitempty"`
	// 起始地省份名称
	StartProvinceName *string `json:"start_province_name,omitempty" xml:"start_province_name,omitempty"`
	// 起始地街道CODE，12位区域行政编码
	StartStreetCode *string `json:"start_street_code,omitempty" xml:"start_street_code,omitempty"`
	// 起始地街道名称
	StartStreetName *string `json:"start_street_name,omitempty" xml:"start_street_name,omitempty"`
	// 起运时间戳，13位毫秒级时间戳
	StartTime *int64 `json:"start_time,omitempty" xml:"start_time,omitempty"`
	// 运单id
	TaxWaybillId *string `json:"tax_waybill_id,omitempty" xml:"tax_waybill_id,omitempty" require:"true"`
	// 车长，可以填：不限车长或者1.8，2.7，3.8，4.2，5，6.2，6.8，7.7，8.2，8.7，9.6，11.7，12.5，13，13.7，15，16，17.5等不超过2位小数的数字
	TruckLength *string `json:"truck_length,omitempty" xml:"truck_length,omitempty"`
	// 车型，可以填写：不限车型，平板，高栏，厢式，集装箱，自卸，冷藏，保温，高低板，面包车，棉被车，爬梯车，飞翼车
	TruckType *string `json:"truck_type,omitempty" xml:"truck_type,omitempty"`
	// 运输单价
	UnitPrice *string `json:"unit_price,omitempty" xml:"unit_price,omitempty"`
}

func (s UpdateFinanceWaybillRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateFinanceWaybillRequest) GoString() string {
	return s.String()
}

func (s *UpdateFinanceWaybillRequest) SetAuthToken(v string) *UpdateFinanceWaybillRequest {
	s.AuthToken = &v
	return s
}

func (s *UpdateFinanceWaybillRequest) SetProductInstanceId(v string) *UpdateFinanceWaybillRequest {
	s.ProductInstanceId = &v
	return s
}

func (s *UpdateFinanceWaybillRequest) SetAllFreight(v string) *UpdateFinanceWaybillRequest {
	s.AllFreight = &v
	return s
}

func (s *UpdateFinanceWaybillRequest) SetBackFee(v string) *UpdateFinanceWaybillRequest {
	s.BackFee = &v
	return s
}

func (s *UpdateFinanceWaybillRequest) SetBusinessType(v string) *UpdateFinanceWaybillRequest {
	s.BusinessType = &v
	return s
}

func (s *UpdateFinanceWaybillRequest) SetCargoBusinessCode(v string) *UpdateFinanceWaybillRequest {
	s.CargoBusinessCode = &v
	return s
}

func (s *UpdateFinanceWaybillRequest) SetCargoCode(v string) *UpdateFinanceWaybillRequest {
	s.CargoCode = &v
	return s
}

func (s *UpdateFinanceWaybillRequest) SetCargoInsuranceMoney(v string) *UpdateFinanceWaybillRequest {
	s.CargoInsuranceMoney = &v
	return s
}

func (s *UpdateFinanceWaybillRequest) SetCargoOrder(v string) *UpdateFinanceWaybillRequest {
	s.CargoOrder = &v
	return s
}

func (s *UpdateFinanceWaybillRequest) SetCargoUnit(v string) *UpdateFinanceWaybillRequest {
	s.CargoUnit = &v
	return s
}

func (s *UpdateFinanceWaybillRequest) SetCargoVolume(v string) *UpdateFinanceWaybillRequest {
	s.CargoVolume = &v
	return s
}

func (s *UpdateFinanceWaybillRequest) SetCargoWeight(v string) *UpdateFinanceWaybillRequest {
	s.CargoWeight = &v
	return s
}

func (s *UpdateFinanceWaybillRequest) SetCartBadgeColor(v string) *UpdateFinanceWaybillRequest {
	s.CartBadgeColor = &v
	return s
}

func (s *UpdateFinanceWaybillRequest) SetCartBadgeNo(v string) *UpdateFinanceWaybillRequest {
	s.CartBadgeNo = &v
	return s
}

func (s *UpdateFinanceWaybillRequest) SetCarCaptainDid(v string) *UpdateFinanceWaybillRequest {
	s.CarCaptainDid = &v
	return s
}

func (s *UpdateFinanceWaybillRequest) SetConsignorFreightAmount(v string) *UpdateFinanceWaybillRequest {
	s.ConsignorFreightAmount = &v
	return s
}

func (s *UpdateFinanceWaybillRequest) SetCreatedTime(v int64) *UpdateFinanceWaybillRequest {
	s.CreatedTime = &v
	return s
}

func (s *UpdateFinanceWaybillRequest) SetDestDoorsEndTime(v string) *UpdateFinanceWaybillRequest {
	s.DestDoorsEndTime = &v
	return s
}

func (s *UpdateFinanceWaybillRequest) SetDoorsCityCode(v string) *UpdateFinanceWaybillRequest {
	s.DoorsCityCode = &v
	return s
}

func (s *UpdateFinanceWaybillRequest) SetDoorsCityName(v string) *UpdateFinanceWaybillRequest {
	s.DoorsCityName = &v
	return s
}

func (s *UpdateFinanceWaybillRequest) SetDoorsCountyCode(v string) *UpdateFinanceWaybillRequest {
	s.DoorsCountyCode = &v
	return s
}

func (s *UpdateFinanceWaybillRequest) SetDoorsCountyName(v string) *UpdateFinanceWaybillRequest {
	s.DoorsCountyName = &v
	return s
}

func (s *UpdateFinanceWaybillRequest) SetDoorsDivisionCode(v string) *UpdateFinanceWaybillRequest {
	s.DoorsDivisionCode = &v
	return s
}

func (s *UpdateFinanceWaybillRequest) SetDoorsProvinceCode(v string) *UpdateFinanceWaybillRequest {
	s.DoorsProvinceCode = &v
	return s
}

func (s *UpdateFinanceWaybillRequest) SetDoorsProvinceName(v string) *UpdateFinanceWaybillRequest {
	s.DoorsProvinceName = &v
	return s
}

func (s *UpdateFinanceWaybillRequest) SetDrawee(v string) *UpdateFinanceWaybillRequest {
	s.Drawee = &v
	return s
}

func (s *UpdateFinanceWaybillRequest) SetDraweeTaxNo(v string) *UpdateFinanceWaybillRequest {
	s.DraweeTaxNo = &v
	return s
}

func (s *UpdateFinanceWaybillRequest) SetDriverDid(v string) *UpdateFinanceWaybillRequest {
	s.DriverDid = &v
	return s
}

func (s *UpdateFinanceWaybillRequest) SetEndAddress(v string) *UpdateFinanceWaybillRequest {
	s.EndAddress = &v
	return s
}

func (s *UpdateFinanceWaybillRequest) SetEndCityCode(v string) *UpdateFinanceWaybillRequest {
	s.EndCityCode = &v
	return s
}

func (s *UpdateFinanceWaybillRequest) SetEndCityName(v string) *UpdateFinanceWaybillRequest {
	s.EndCityName = &v
	return s
}

func (s *UpdateFinanceWaybillRequest) SetEndCountyCode(v string) *UpdateFinanceWaybillRequest {
	s.EndCountyCode = &v
	return s
}

func (s *UpdateFinanceWaybillRequest) SetEndCountyName(v string) *UpdateFinanceWaybillRequest {
	s.EndCountyName = &v
	return s
}

func (s *UpdateFinanceWaybillRequest) SetEndDivisionCode(v string) *UpdateFinanceWaybillRequest {
	s.EndDivisionCode = &v
	return s
}

func (s *UpdateFinanceWaybillRequest) SetEndProvinceCode(v string) *UpdateFinanceWaybillRequest {
	s.EndProvinceCode = &v
	return s
}

func (s *UpdateFinanceWaybillRequest) SetEndProvinceName(v string) *UpdateFinanceWaybillRequest {
	s.EndProvinceName = &v
	return s
}

func (s *UpdateFinanceWaybillRequest) SetEndStreetCode(v string) *UpdateFinanceWaybillRequest {
	s.EndStreetCode = &v
	return s
}

func (s *UpdateFinanceWaybillRequest) SetEndStreetName(v string) *UpdateFinanceWaybillRequest {
	s.EndStreetName = &v
	return s
}

func (s *UpdateFinanceWaybillRequest) SetEndTime(v int64) *UpdateFinanceWaybillRequest {
	s.EndTime = &v
	return s
}

func (s *UpdateFinanceWaybillRequest) SetFreightIncr(v string) *UpdateFinanceWaybillRequest {
	s.FreightIncr = &v
	return s
}

func (s *UpdateFinanceWaybillRequest) SetGoodsAmount(v int64) *UpdateFinanceWaybillRequest {
	s.GoodsAmount = &v
	return s
}

func (s *UpdateFinanceWaybillRequest) SetGoodsAmountType(v string) *UpdateFinanceWaybillRequest {
	s.GoodsAmountType = &v
	return s
}

func (s *UpdateFinanceWaybillRequest) SetGoodsName(v string) *UpdateFinanceWaybillRequest {
	s.GoodsName = &v
	return s
}

func (s *UpdateFinanceWaybillRequest) SetLossFee(v string) *UpdateFinanceWaybillRequest {
	s.LossFee = &v
	return s
}

func (s *UpdateFinanceWaybillRequest) SetPartnerId(v string) *UpdateFinanceWaybillRequest {
	s.PartnerId = &v
	return s
}

func (s *UpdateFinanceWaybillRequest) SetPlatformDid(v string) *UpdateFinanceWaybillRequest {
	s.PlatformDid = &v
	return s
}

func (s *UpdateFinanceWaybillRequest) SetPrepayments(v string) *UpdateFinanceWaybillRequest {
	s.Prepayments = &v
	return s
}

func (s *UpdateFinanceWaybillRequest) SetPrepaymentsBuyEtc(v string) *UpdateFinanceWaybillRequest {
	s.PrepaymentsBuyEtc = &v
	return s
}

func (s *UpdateFinanceWaybillRequest) SetPrepaymentsBuyGas(v string) *UpdateFinanceWaybillRequest {
	s.PrepaymentsBuyGas = &v
	return s
}

func (s *UpdateFinanceWaybillRequest) SetPrepaymentsBuyOil(v string) *UpdateFinanceWaybillRequest {
	s.PrepaymentsBuyOil = &v
	return s
}

func (s *UpdateFinanceWaybillRequest) SetPrepaymentsEtccard(v string) *UpdateFinanceWaybillRequest {
	s.PrepaymentsEtccard = &v
	return s
}

func (s *UpdateFinanceWaybillRequest) SetPrepaymentsOilcard(v string) *UpdateFinanceWaybillRequest {
	s.PrepaymentsOilcard = &v
	return s
}

func (s *UpdateFinanceWaybillRequest) SetPresentAmountOil(v string) *UpdateFinanceWaybillRequest {
	s.PresentAmountOil = &v
	return s
}

func (s *UpdateFinanceWaybillRequest) SetStartAddress(v string) *UpdateFinanceWaybillRequest {
	s.StartAddress = &v
	return s
}

func (s *UpdateFinanceWaybillRequest) SetStartCityCode(v string) *UpdateFinanceWaybillRequest {
	s.StartCityCode = &v
	return s
}

func (s *UpdateFinanceWaybillRequest) SetStartCityName(v string) *UpdateFinanceWaybillRequest {
	s.StartCityName = &v
	return s
}

func (s *UpdateFinanceWaybillRequest) SetStartCountyCode(v string) *UpdateFinanceWaybillRequest {
	s.StartCountyCode = &v
	return s
}

func (s *UpdateFinanceWaybillRequest) SetStartCountyName(v string) *UpdateFinanceWaybillRequest {
	s.StartCountyName = &v
	return s
}

func (s *UpdateFinanceWaybillRequest) SetStartDivisionCode(v string) *UpdateFinanceWaybillRequest {
	s.StartDivisionCode = &v
	return s
}

func (s *UpdateFinanceWaybillRequest) SetStartProvinceCode(v string) *UpdateFinanceWaybillRequest {
	s.StartProvinceCode = &v
	return s
}

func (s *UpdateFinanceWaybillRequest) SetStartProvinceName(v string) *UpdateFinanceWaybillRequest {
	s.StartProvinceName = &v
	return s
}

func (s *UpdateFinanceWaybillRequest) SetStartStreetCode(v string) *UpdateFinanceWaybillRequest {
	s.StartStreetCode = &v
	return s
}

func (s *UpdateFinanceWaybillRequest) SetStartStreetName(v string) *UpdateFinanceWaybillRequest {
	s.StartStreetName = &v
	return s
}

func (s *UpdateFinanceWaybillRequest) SetStartTime(v int64) *UpdateFinanceWaybillRequest {
	s.StartTime = &v
	return s
}

func (s *UpdateFinanceWaybillRequest) SetTaxWaybillId(v string) *UpdateFinanceWaybillRequest {
	s.TaxWaybillId = &v
	return s
}

func (s *UpdateFinanceWaybillRequest) SetTruckLength(v string) *UpdateFinanceWaybillRequest {
	s.TruckLength = &v
	return s
}

func (s *UpdateFinanceWaybillRequest) SetTruckType(v string) *UpdateFinanceWaybillRequest {
	s.TruckType = &v
	return s
}

func (s *UpdateFinanceWaybillRequest) SetUnitPrice(v string) *UpdateFinanceWaybillRequest {
	s.UnitPrice = &v
	return s
}

type UpdateFinanceWaybillResponse struct {
	// 请求唯一ID，用于链路跟踪和问题排查
	ReqMsgId *string `json:"req_msg_id,omitempty" xml:"req_msg_id,omitempty"`
	// 结果码，一般OK表示调用成功
	ResultCode *string `json:"result_code,omitempty" xml:"result_code,omitempty"`
	// 异常信息的文本描述
	ResultMsg *string `json:"result_msg,omitempty" xml:"result_msg,omitempty"`
	// 货源支付链上凭证
	//
	TxCode *string `json:"tx_code,omitempty" xml:"tx_code,omitempty"`
}

func (s UpdateFinanceWaybillResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateFinanceWaybillResponse) GoString() string {
	return s.String()
}

func (s *UpdateFinanceWaybillResponse) SetReqMsgId(v string) *UpdateFinanceWaybillResponse {
	s.ReqMsgId = &v
	return s
}

func (s *UpdateFinanceWaybillResponse) SetResultCode(v string) *UpdateFinanceWaybillResponse {
	s.ResultCode = &v
	return s
}

func (s *UpdateFinanceWaybillResponse) SetResultMsg(v string) *UpdateFinanceWaybillResponse {
	s.ResultMsg = &v
	return s
}

func (s *UpdateFinanceWaybillResponse) SetTxCode(v string) *UpdateFinanceWaybillResponse {
	s.TxCode = &v
	return s
}

type StartFinanceWaybillRequest struct {
	// OAuth模式下的授权token
	AuthToken         *string `json:"auth_token,omitempty" xml:"auth_token,omitempty"`
	ProductInstanceId *string `json:"product_instance_id,omitempty" xml:"product_instance_id,omitempty"`
	// 货物运费险
	CargoInsuranceMoney *string `json:"cargo_insurance_money,omitempty" xml:"cargo_insurance_money,omitempty"`
	// 车牌颜色
	CartBadgeColor *string `json:"cart_badge_color,omitempty" xml:"cart_badge_color,omitempty" require:"true"`
	// 车牌号
	CartBadgeNo *string `json:"cart_badge_no,omitempty" xml:"cart_badge_no,omitempty" require:"true"`
	// 承运商did
	CarCaptainDid *string `json:"car_captain_did,omitempty" xml:"car_captain_did,omitempty"`
	// 司机分布式数字身份
	DriverDid *string `json:"driver_did,omitempty" xml:"driver_did,omitempty" require:"true"`
	// 无车承运平台分布式数字身份，缺省时为自己的分布式数字身份
	PlatformDid *string `json:"platform_did,omitempty" xml:"platform_did,omitempty" require:"true"`
	// 起运时间，13位毫秒级时间戳
	StartTime *int64 `json:"start_time,omitempty" xml:"start_time,omitempty" require:"true"`
	// 运单id
	TaxWaybillId *string `json:"tax_waybill_id,omitempty" xml:"tax_waybill_id,omitempty" require:"true"`
}

func (s StartFinanceWaybillRequest) String() string {
	return tea.Prettify(s)
}

func (s StartFinanceWaybillRequest) GoString() string {
	return s.String()
}

func (s *StartFinanceWaybillRequest) SetAuthToken(v string) *StartFinanceWaybillRequest {
	s.AuthToken = &v
	return s
}

func (s *StartFinanceWaybillRequest) SetProductInstanceId(v string) *StartFinanceWaybillRequest {
	s.ProductInstanceId = &v
	return s
}

func (s *StartFinanceWaybillRequest) SetCargoInsuranceMoney(v string) *StartFinanceWaybillRequest {
	s.CargoInsuranceMoney = &v
	return s
}

func (s *StartFinanceWaybillRequest) SetCartBadgeColor(v string) *StartFinanceWaybillRequest {
	s.CartBadgeColor = &v
	return s
}

func (s *StartFinanceWaybillRequest) SetCartBadgeNo(v string) *StartFinanceWaybillRequest {
	s.CartBadgeNo = &v
	return s
}

func (s *StartFinanceWaybillRequest) SetCarCaptainDid(v string) *StartFinanceWaybillRequest {
	s.CarCaptainDid = &v
	return s
}

func (s *StartFinanceWaybillRequest) SetDriverDid(v string) *StartFinanceWaybillRequest {
	s.DriverDid = &v
	return s
}

func (s *StartFinanceWaybillRequest) SetPlatformDid(v string) *StartFinanceWaybillRequest {
	s.PlatformDid = &v
	return s
}

func (s *StartFinanceWaybillRequest) SetStartTime(v int64) *StartFinanceWaybillRequest {
	s.StartTime = &v
	return s
}

func (s *StartFinanceWaybillRequest) SetTaxWaybillId(v string) *StartFinanceWaybillRequest {
	s.TaxWaybillId = &v
	return s
}

type StartFinanceWaybillResponse struct {
	// 请求唯一ID，用于链路跟踪和问题排查
	ReqMsgId *string `json:"req_msg_id,omitempty" xml:"req_msg_id,omitempty"`
	// 结果码，一般OK表示调用成功
	ResultCode *string `json:"result_code,omitempty" xml:"result_code,omitempty"`
	// 异常信息的文本描述
	ResultMsg *string `json:"result_msg,omitempty" xml:"result_msg,omitempty"`
	// 货源支付链上凭证
	//
	TxCode *string `json:"tx_code,omitempty" xml:"tx_code,omitempty"`
}

func (s StartFinanceWaybillResponse) String() string {
	return tea.Prettify(s)
}

func (s StartFinanceWaybillResponse) GoString() string {
	return s.String()
}

func (s *StartFinanceWaybillResponse) SetReqMsgId(v string) *StartFinanceWaybillResponse {
	s.ReqMsgId = &v
	return s
}

func (s *StartFinanceWaybillResponse) SetResultCode(v string) *StartFinanceWaybillResponse {
	s.ResultCode = &v
	return s
}

func (s *StartFinanceWaybillResponse) SetResultMsg(v string) *StartFinanceWaybillResponse {
	s.ResultMsg = &v
	return s
}

func (s *StartFinanceWaybillResponse) SetTxCode(v string) *StartFinanceWaybillResponse {
	s.TxCode = &v
	return s
}

type CreateCaptainDisRequest struct {
	// OAuth模式下的授权token
	AuthToken         *string `json:"auth_token,omitempty" xml:"auth_token,omitempty"`
	ProductInstanceId *string `json:"product_instance_id,omitempty" xml:"product_instance_id,omitempty"`
	// 身份证号码
	CertNo *string `json:"cert_no,omitempty" xml:"cert_no,omitempty" require:"true"`
	// 扩展字段
	ExtensionInfo *string `json:"extension_info,omitempty" xml:"extension_info,omitempty"`
	// 手机号码
	Mobile *string `json:"mobile,omitempty" xml:"mobile,omitempty" require:"true"`
	// 姓名
	Name *string `json:"name,omitempty" xml:"name,omitempty" require:"true"`
	// 所属平台did
	PlatformDid *string `json:"platform_did,omitempty" xml:"platform_did,omitempty" require:"true"`
}

func (s CreateCaptainDisRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateCaptainDisRequest) GoString() string {
	return s.String()
}

func (s *CreateCaptainDisRequest) SetAuthToken(v string) *CreateCaptainDisRequest {
	s.AuthToken = &v
	return s
}

func (s *CreateCaptainDisRequest) SetProductInstanceId(v string) *CreateCaptainDisRequest {
	s.ProductInstanceId = &v
	return s
}

func (s *CreateCaptainDisRequest) SetCertNo(v string) *CreateCaptainDisRequest {
	s.CertNo = &v
	return s
}

func (s *CreateCaptainDisRequest) SetExtensionInfo(v string) *CreateCaptainDisRequest {
	s.ExtensionInfo = &v
	return s
}

func (s *CreateCaptainDisRequest) SetMobile(v string) *CreateCaptainDisRequest {
	s.Mobile = &v
	return s
}

func (s *CreateCaptainDisRequest) SetName(v string) *CreateCaptainDisRequest {
	s.Name = &v
	return s
}

func (s *CreateCaptainDisRequest) SetPlatformDid(v string) *CreateCaptainDisRequest {
	s.PlatformDid = &v
	return s
}

type CreateCaptainDisResponse struct {
	// 请求唯一ID，用于链路跟踪和问题排查
	ReqMsgId *string `json:"req_msg_id,omitempty" xml:"req_msg_id,omitempty"`
	// 结果码，一般OK表示调用成功
	ResultCode *string `json:"result_code,omitempty" xml:"result_code,omitempty"`
	// 异常信息的文本描述
	ResultMsg *string `json:"result_msg,omitempty" xml:"result_msg,omitempty"`
	// 车队长did
	Did *string `json:"did,omitempty" xml:"did,omitempty"`
}

func (s CreateCaptainDisResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateCaptainDisResponse) GoString() string {
	return s.String()
}

func (s *CreateCaptainDisResponse) SetReqMsgId(v string) *CreateCaptainDisResponse {
	s.ReqMsgId = &v
	return s
}

func (s *CreateCaptainDisResponse) SetResultCode(v string) *CreateCaptainDisResponse {
	s.ResultCode = &v
	return s
}

func (s *CreateCaptainDisResponse) SetResultMsg(v string) *CreateCaptainDisResponse {
	s.ResultMsg = &v
	return s
}

func (s *CreateCaptainDisResponse) SetDid(v string) *CreateCaptainDisResponse {
	s.Did = &v
	return s
}

type CreateCargowaybillBillRequest struct {
	// OAuth模式下的授权token
	AuthToken         *string `json:"auth_token,omitempty" xml:"auth_token,omitempty"`
	ProductInstanceId *string `json:"product_instance_id,omitempty" xml:"product_instance_id,omitempty"`
	// 货主账单总额（两位小数）
	BillAmount *string `json:"bill_amount,omitempty" xml:"bill_amount,omitempty" require:"true"`
	// 账单code（唯一标识）
	BillCode *string `json:"bill_code,omitempty" xml:"bill_code,omitempty" require:"true"`
	// 账单生成时间(毫秒)
	BillCreateTime *int64 `json:"bill_create_time,omitempty" xml:"bill_create_time,omitempty" require:"true"`
	// 账单期限
	BillDeadline *string `json:"bill_deadline,omitempty" xml:"bill_deadline,omitempty" require:"true"`
	// 付款方did
	BillPayerDid *string `json:"bill_payer_did,omitempty" xml:"bill_payer_did,omitempty" require:"true"`
	// 货源单号-货主运费列表
	CargoAmounts []*CargoAmount `json:"cargo_amounts,omitempty" xml:"cargo_amounts,omitempty" require:"true" type:"Repeated"`
	// 币种
	Currency *string `json:"currency,omitempty" xml:"currency,omitempty" require:"true"`
	// 货主支付方式-支付金额列表
	PayAmounts []*PayAmount `json:"pay_amounts,omitempty" xml:"pay_amounts,omitempty" require:"true" type:"Repeated"`
	// 平台did
	PlatformDid *string `json:"platform_did,omitempty" xml:"platform_did,omitempty" require:"true"`
	// 预计货主付款日期
	PreConsignorPayDate *int64 `json:"pre_consignor_pay_date,omitempty" xml:"pre_consignor_pay_date,omitempty" require:"true"`
	// 运单号-司机运费列表
	WaybillAmounts []*WaybillAmount `json:"waybill_amounts,omitempty" xml:"waybill_amounts,omitempty" require:"true" type:"Repeated"`
	// 账单到期日期
	Deadline *int64 `json:"deadline,omitempty" xml:"deadline,omitempty" require:"true"`
	// 运单上传状态，可填写：已完成、未完成
	WaybillUploadStatus *string `json:"waybill_upload_status,omitempty" xml:"waybill_upload_status,omitempty" require:"true"`
}

func (s CreateCargowaybillBillRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateCargowaybillBillRequest) GoString() string {
	return s.String()
}

func (s *CreateCargowaybillBillRequest) SetAuthToken(v string) *CreateCargowaybillBillRequest {
	s.AuthToken = &v
	return s
}

func (s *CreateCargowaybillBillRequest) SetProductInstanceId(v string) *CreateCargowaybillBillRequest {
	s.ProductInstanceId = &v
	return s
}

func (s *CreateCargowaybillBillRequest) SetBillAmount(v string) *CreateCargowaybillBillRequest {
	s.BillAmount = &v
	return s
}

func (s *CreateCargowaybillBillRequest) SetBillCode(v string) *CreateCargowaybillBillRequest {
	s.BillCode = &v
	return s
}

func (s *CreateCargowaybillBillRequest) SetBillCreateTime(v int64) *CreateCargowaybillBillRequest {
	s.BillCreateTime = &v
	return s
}

func (s *CreateCargowaybillBillRequest) SetBillDeadline(v string) *CreateCargowaybillBillRequest {
	s.BillDeadline = &v
	return s
}

func (s *CreateCargowaybillBillRequest) SetBillPayerDid(v string) *CreateCargowaybillBillRequest {
	s.BillPayerDid = &v
	return s
}

func (s *CreateCargowaybillBillRequest) SetCargoAmounts(v []*CargoAmount) *CreateCargowaybillBillRequest {
	s.CargoAmounts = v
	return s
}

func (s *CreateCargowaybillBillRequest) SetCurrency(v string) *CreateCargowaybillBillRequest {
	s.Currency = &v
	return s
}

func (s *CreateCargowaybillBillRequest) SetPayAmounts(v []*PayAmount) *CreateCargowaybillBillRequest {
	s.PayAmounts = v
	return s
}

func (s *CreateCargowaybillBillRequest) SetPlatformDid(v string) *CreateCargowaybillBillRequest {
	s.PlatformDid = &v
	return s
}

func (s *CreateCargowaybillBillRequest) SetPreConsignorPayDate(v int64) *CreateCargowaybillBillRequest {
	s.PreConsignorPayDate = &v
	return s
}

func (s *CreateCargowaybillBillRequest) SetWaybillAmounts(v []*WaybillAmount) *CreateCargowaybillBillRequest {
	s.WaybillAmounts = v
	return s
}

func (s *CreateCargowaybillBillRequest) SetDeadline(v int64) *CreateCargowaybillBillRequest {
	s.Deadline = &v
	return s
}

func (s *CreateCargowaybillBillRequest) SetWaybillUploadStatus(v string) *CreateCargowaybillBillRequest {
	s.WaybillUploadStatus = &v
	return s
}

type CreateCargowaybillBillResponse struct {
	// 请求唯一ID，用于链路跟踪和问题排查
	ReqMsgId *string `json:"req_msg_id,omitempty" xml:"req_msg_id,omitempty"`
	// 结果码，一般OK表示调用成功
	ResultCode *string `json:"result_code,omitempty" xml:"result_code,omitempty"`
	// 异常信息的文本描述
	ResultMsg *string `json:"result_msg,omitempty" xml:"result_msg,omitempty"`
	// 链上凭证
	TxCode *string `json:"tx_code,omitempty" xml:"tx_code,omitempty"`
}

func (s CreateCargowaybillBillResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateCargowaybillBillResponse) GoString() string {
	return s.String()
}

func (s *CreateCargowaybillBillResponse) SetReqMsgId(v string) *CreateCargowaybillBillResponse {
	s.ReqMsgId = &v
	return s
}

func (s *CreateCargowaybillBillResponse) SetResultCode(v string) *CreateCargowaybillBillResponse {
	s.ResultCode = &v
	return s
}

func (s *CreateCargowaybillBillResponse) SetResultMsg(v string) *CreateCargowaybillBillResponse {
	s.ResultMsg = &v
	return s
}

func (s *CreateCargowaybillBillResponse) SetTxCode(v string) *CreateCargowaybillBillResponse {
	s.TxCode = &v
	return s
}

type ConfirmCargowaybillBillRequest struct {
	// OAuth模式下的授权token
	AuthToken         *string `json:"auth_token,omitempty" xml:"auth_token,omitempty"`
	ProductInstanceId *string `json:"product_instance_id,omitempty" xml:"product_instance_id,omitempty"`
	// 账单金额（两位小数）
	BillAmount *string `json:"bill_amount,omitempty" xml:"bill_amount,omitempty" require:"true"`
	// 账单code（账单唯一标识）
	BillCode *string `json:"bill_code,omitempty" xml:"bill_code,omitempty" require:"true"`
	// 账单确认货主did
	BillConsignorDid *string `json:"bill_consignor_did,omitempty" xml:"bill_consignor_did,omitempty" require:"true"`
	// 账单确认日期（毫秒）
	BillSureDate *int64 `json:"bill_sure_date,omitempty" xml:"bill_sure_date,omitempty"`
	// 是否结算
	WhetherSettle *bool `json:"whether_settle,omitempty" xml:"whether_settle,omitempty" require:"true"`
	// 平台did
	PlatformDid *string `json:"platform_did,omitempty" xml:"platform_did,omitempty" require:"true"`
}

func (s ConfirmCargowaybillBillRequest) String() string {
	return tea.Prettify(s)
}

func (s ConfirmCargowaybillBillRequest) GoString() string {
	return s.String()
}

func (s *ConfirmCargowaybillBillRequest) SetAuthToken(v string) *ConfirmCargowaybillBillRequest {
	s.AuthToken = &v
	return s
}

func (s *ConfirmCargowaybillBillRequest) SetProductInstanceId(v string) *ConfirmCargowaybillBillRequest {
	s.ProductInstanceId = &v
	return s
}

func (s *ConfirmCargowaybillBillRequest) SetBillAmount(v string) *ConfirmCargowaybillBillRequest {
	s.BillAmount = &v
	return s
}

func (s *ConfirmCargowaybillBillRequest) SetBillCode(v string) *ConfirmCargowaybillBillRequest {
	s.BillCode = &v
	return s
}

func (s *ConfirmCargowaybillBillRequest) SetBillConsignorDid(v string) *ConfirmCargowaybillBillRequest {
	s.BillConsignorDid = &v
	return s
}

func (s *ConfirmCargowaybillBillRequest) SetBillSureDate(v int64) *ConfirmCargowaybillBillRequest {
	s.BillSureDate = &v
	return s
}

func (s *ConfirmCargowaybillBillRequest) SetWhetherSettle(v bool) *ConfirmCargowaybillBillRequest {
	s.WhetherSettle = &v
	return s
}

func (s *ConfirmCargowaybillBillRequest) SetPlatformDid(v string) *ConfirmCargowaybillBillRequest {
	s.PlatformDid = &v
	return s
}

type ConfirmCargowaybillBillResponse struct {
	// 请求唯一ID，用于链路跟踪和问题排查
	ReqMsgId *string `json:"req_msg_id,omitempty" xml:"req_msg_id,omitempty"`
	// 结果码，一般OK表示调用成功
	ResultCode *string `json:"result_code,omitempty" xml:"result_code,omitempty"`
	// 异常信息的文本描述
	ResultMsg *string `json:"result_msg,omitempty" xml:"result_msg,omitempty"`
	// 链上凭证
	TxCode *string `json:"tx_code,omitempty" xml:"tx_code,omitempty"`
}

func (s ConfirmCargowaybillBillResponse) String() string {
	return tea.Prettify(s)
}

func (s ConfirmCargowaybillBillResponse) GoString() string {
	return s.String()
}

func (s *ConfirmCargowaybillBillResponse) SetReqMsgId(v string) *ConfirmCargowaybillBillResponse {
	s.ReqMsgId = &v
	return s
}

func (s *ConfirmCargowaybillBillResponse) SetResultCode(v string) *ConfirmCargowaybillBillResponse {
	s.ResultCode = &v
	return s
}

func (s *ConfirmCargowaybillBillResponse) SetResultMsg(v string) *ConfirmCargowaybillBillResponse {
	s.ResultMsg = &v
	return s
}

func (s *ConfirmCargowaybillBillResponse) SetTxCode(v string) *ConfirmCargowaybillBillResponse {
	s.TxCode = &v
	return s
}

type CreateCargowaybillBillsettleRequest struct {
	// OAuth模式下的授权token
	AuthToken         *string `json:"auth_token,omitempty" xml:"auth_token,omitempty"`
	ProductInstanceId *string `json:"product_instance_id,omitempty" xml:"product_instance_id,omitempty"`
	// 平台did
	PlatformDid *string `json:"platform_did,omitempty" xml:"platform_did,omitempty" require:"true"`
	// 账单code（账单唯一标识）
	BillCode *string `json:"bill_code,omitempty" xml:"bill_code,omitempty" require:"true"`
	// 结算金额（2位小数）
	SettleAmount *string `json:"settle_amount,omitempty" xml:"settle_amount,omitempty" require:"true"`
	// 结算状态（只有2种状态：部分结算、已结清）
	SettleStatus *string `json:"settle_status,omitempty" xml:"settle_status,omitempty" require:"true"`
}

func (s CreateCargowaybillBillsettleRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateCargowaybillBillsettleRequest) GoString() string {
	return s.String()
}

func (s *CreateCargowaybillBillsettleRequest) SetAuthToken(v string) *CreateCargowaybillBillsettleRequest {
	s.AuthToken = &v
	return s
}

func (s *CreateCargowaybillBillsettleRequest) SetProductInstanceId(v string) *CreateCargowaybillBillsettleRequest {
	s.ProductInstanceId = &v
	return s
}

func (s *CreateCargowaybillBillsettleRequest) SetPlatformDid(v string) *CreateCargowaybillBillsettleRequest {
	s.PlatformDid = &v
	return s
}

func (s *CreateCargowaybillBillsettleRequest) SetBillCode(v string) *CreateCargowaybillBillsettleRequest {
	s.BillCode = &v
	return s
}

func (s *CreateCargowaybillBillsettleRequest) SetSettleAmount(v string) *CreateCargowaybillBillsettleRequest {
	s.SettleAmount = &v
	return s
}

func (s *CreateCargowaybillBillsettleRequest) SetSettleStatus(v string) *CreateCargowaybillBillsettleRequest {
	s.SettleStatus = &v
	return s
}

type CreateCargowaybillBillsettleResponse struct {
	// 请求唯一ID，用于链路跟踪和问题排查
	ReqMsgId *string `json:"req_msg_id,omitempty" xml:"req_msg_id,omitempty"`
	// 结果码，一般OK表示调用成功
	ResultCode *string `json:"result_code,omitempty" xml:"result_code,omitempty"`
	// 异常信息的文本描述
	ResultMsg *string `json:"result_msg,omitempty" xml:"result_msg,omitempty"`
	// 链上凭证
	TxCode *string `json:"tx_code,omitempty" xml:"tx_code,omitempty"`
}

func (s CreateCargowaybillBillsettleResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateCargowaybillBillsettleResponse) GoString() string {
	return s.String()
}

func (s *CreateCargowaybillBillsettleResponse) SetReqMsgId(v string) *CreateCargowaybillBillsettleResponse {
	s.ReqMsgId = &v
	return s
}

func (s *CreateCargowaybillBillsettleResponse) SetResultCode(v string) *CreateCargowaybillBillsettleResponse {
	s.ResultCode = &v
	return s
}

func (s *CreateCargowaybillBillsettleResponse) SetResultMsg(v string) *CreateCargowaybillBillsettleResponse {
	s.ResultMsg = &v
	return s
}

func (s *CreateCargowaybillBillsettleResponse) SetTxCode(v string) *CreateCargowaybillBillsettleResponse {
	s.TxCode = &v
	return s
}

type UpdateWaybillorderPlatformdidRequest struct {
	// OAuth模式下的授权token
	AuthToken         *string `json:"auth_token,omitempty" xml:"auth_token,omitempty"`
	ProductInstanceId *string `json:"product_instance_id,omitempty" xml:"product_instance_id,omitempty"`
	// 新平台diid
	NewPlatformDid *string `json:"new_platform_did,omitempty" xml:"new_platform_did,omitempty" require:"true"`
	// 老平台did
	OldPlatformDid *string `json:"old_platform_did,omitempty" xml:"old_platform_did,omitempty" require:"true"`
	// 运单号
	TaxWaybillId *string `json:"tax_waybill_id,omitempty" xml:"tax_waybill_id,omitempty" require:"true"`
}

func (s UpdateWaybillorderPlatformdidRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateWaybillorderPlatformdidRequest) GoString() string {
	return s.String()
}

func (s *UpdateWaybillorderPlatformdidRequest) SetAuthToken(v string) *UpdateWaybillorderPlatformdidRequest {
	s.AuthToken = &v
	return s
}

func (s *UpdateWaybillorderPlatformdidRequest) SetProductInstanceId(v string) *UpdateWaybillorderPlatformdidRequest {
	s.ProductInstanceId = &v
	return s
}

func (s *UpdateWaybillorderPlatformdidRequest) SetNewPlatformDid(v string) *UpdateWaybillorderPlatformdidRequest {
	s.NewPlatformDid = &v
	return s
}

func (s *UpdateWaybillorderPlatformdidRequest) SetOldPlatformDid(v string) *UpdateWaybillorderPlatformdidRequest {
	s.OldPlatformDid = &v
	return s
}

func (s *UpdateWaybillorderPlatformdidRequest) SetTaxWaybillId(v string) *UpdateWaybillorderPlatformdidRequest {
	s.TaxWaybillId = &v
	return s
}

type UpdateWaybillorderPlatformdidResponse struct {
	// 请求唯一ID，用于链路跟踪和问题排查
	ReqMsgId *string `json:"req_msg_id,omitempty" xml:"req_msg_id,omitempty"`
	// 结果码，一般OK表示调用成功
	ResultCode *string `json:"result_code,omitempty" xml:"result_code,omitempty"`
	// 异常信息的文本描述
	ResultMsg *string `json:"result_msg,omitempty" xml:"result_msg,omitempty"`
	// 链上凭证
	TxCode *string `json:"tx_code,omitempty" xml:"tx_code,omitempty"`
}

func (s UpdateWaybillorderPlatformdidResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateWaybillorderPlatformdidResponse) GoString() string {
	return s.String()
}

func (s *UpdateWaybillorderPlatformdidResponse) SetReqMsgId(v string) *UpdateWaybillorderPlatformdidResponse {
	s.ReqMsgId = &v
	return s
}

func (s *UpdateWaybillorderPlatformdidResponse) SetResultCode(v string) *UpdateWaybillorderPlatformdidResponse {
	s.ResultCode = &v
	return s
}

func (s *UpdateWaybillorderPlatformdidResponse) SetResultMsg(v string) *UpdateWaybillorderPlatformdidResponse {
	s.ResultMsg = &v
	return s
}

func (s *UpdateWaybillorderPlatformdidResponse) SetTxCode(v string) *UpdateWaybillorderPlatformdidResponse {
	s.TxCode = &v
	return s
}

type UpdateCargowaybillBillRequest struct {
	// OAuth模式下的授权token
	AuthToken         *string `json:"auth_token,omitempty" xml:"auth_token,omitempty"`
	ProductInstanceId *string `json:"product_instance_id,omitempty" xml:"product_instance_id,omitempty"`
	// 货主账单总额（两位小数）
	BillAmount *string `json:"bill_amount,omitempty" xml:"bill_amount,omitempty" require:"true"`
	// 账单code（唯一标识）
	BillCode *string `json:"bill_code,omitempty" xml:"bill_code,omitempty" require:"true"`
	// 账单生成时间(毫秒)
	BillCreateTime *int64 `json:"bill_create_time,omitempty" xml:"bill_create_time,omitempty" require:"true"`
	// 账单期限
	BillDeadline *string `json:"bill_deadline,omitempty" xml:"bill_deadline,omitempty" require:"true"`
	// 付款方did
	BillPayerDid *string `json:"bill_payer_did,omitempty" xml:"bill_payer_did,omitempty" require:"true"`
	// 货源单号-货主运费列表
	CargoAmounts []*CargoAmount `json:"cargo_amounts,omitempty" xml:"cargo_amounts,omitempty" require:"true" type:"Repeated"`
	// 币种
	Currency *string `json:"currency,omitempty" xml:"currency,omitempty" require:"true"`
	// 账单到期日期
	Deadline *int64 `json:"deadline,omitempty" xml:"deadline,omitempty" require:"true"`
	// 货主支付方式-支付金额列表
	PayAmounts []*PayAmount `json:"pay_amounts,omitempty" xml:"pay_amounts,omitempty" require:"true" type:"Repeated"`
	// 平台did
	PlatformDid *string `json:"platform_did,omitempty" xml:"platform_did,omitempty" require:"true"`
	// 预计货主付款日期
	PreConsignorPayDate *int64 `json:"pre_consignor_pay_date,omitempty" xml:"pre_consignor_pay_date,omitempty" require:"true"`
	// 运单号-司机运费列表
	WaybillAmounts []*WaybillAmount `json:"waybill_amounts,omitempty" xml:"waybill_amounts,omitempty" type:"Repeated"`
	// 运单上传状态，可填写：已完成、未完成
	WaybillUploadStatus *string `json:"waybill_upload_status,omitempty" xml:"waybill_upload_status,omitempty" require:"true"`
}

func (s UpdateCargowaybillBillRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateCargowaybillBillRequest) GoString() string {
	return s.String()
}

func (s *UpdateCargowaybillBillRequest) SetAuthToken(v string) *UpdateCargowaybillBillRequest {
	s.AuthToken = &v
	return s
}

func (s *UpdateCargowaybillBillRequest) SetProductInstanceId(v string) *UpdateCargowaybillBillRequest {
	s.ProductInstanceId = &v
	return s
}

func (s *UpdateCargowaybillBillRequest) SetBillAmount(v string) *UpdateCargowaybillBillRequest {
	s.BillAmount = &v
	return s
}

func (s *UpdateCargowaybillBillRequest) SetBillCode(v string) *UpdateCargowaybillBillRequest {
	s.BillCode = &v
	return s
}

func (s *UpdateCargowaybillBillRequest) SetBillCreateTime(v int64) *UpdateCargowaybillBillRequest {
	s.BillCreateTime = &v
	return s
}

func (s *UpdateCargowaybillBillRequest) SetBillDeadline(v string) *UpdateCargowaybillBillRequest {
	s.BillDeadline = &v
	return s
}

func (s *UpdateCargowaybillBillRequest) SetBillPayerDid(v string) *UpdateCargowaybillBillRequest {
	s.BillPayerDid = &v
	return s
}

func (s *UpdateCargowaybillBillRequest) SetCargoAmounts(v []*CargoAmount) *UpdateCargowaybillBillRequest {
	s.CargoAmounts = v
	return s
}

func (s *UpdateCargowaybillBillRequest) SetCurrency(v string) *UpdateCargowaybillBillRequest {
	s.Currency = &v
	return s
}

func (s *UpdateCargowaybillBillRequest) SetDeadline(v int64) *UpdateCargowaybillBillRequest {
	s.Deadline = &v
	return s
}

func (s *UpdateCargowaybillBillRequest) SetPayAmounts(v []*PayAmount) *UpdateCargowaybillBillRequest {
	s.PayAmounts = v
	return s
}

func (s *UpdateCargowaybillBillRequest) SetPlatformDid(v string) *UpdateCargowaybillBillRequest {
	s.PlatformDid = &v
	return s
}

func (s *UpdateCargowaybillBillRequest) SetPreConsignorPayDate(v int64) *UpdateCargowaybillBillRequest {
	s.PreConsignorPayDate = &v
	return s
}

func (s *UpdateCargowaybillBillRequest) SetWaybillAmounts(v []*WaybillAmount) *UpdateCargowaybillBillRequest {
	s.WaybillAmounts = v
	return s
}

func (s *UpdateCargowaybillBillRequest) SetWaybillUploadStatus(v string) *UpdateCargowaybillBillRequest {
	s.WaybillUploadStatus = &v
	return s
}

type UpdateCargowaybillBillResponse struct {
	// 请求唯一ID，用于链路跟踪和问题排查
	ReqMsgId *string `json:"req_msg_id,omitempty" xml:"req_msg_id,omitempty"`
	// 结果码，一般OK表示调用成功
	ResultCode *string `json:"result_code,omitempty" xml:"result_code,omitempty"`
	// 异常信息的文本描述
	ResultMsg *string `json:"result_msg,omitempty" xml:"result_msg,omitempty"`
	// 链上凭证
	TxCode *string `json:"tx_code,omitempty" xml:"tx_code,omitempty"`
}

func (s UpdateCargowaybillBillResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateCargowaybillBillResponse) GoString() string {
	return s.String()
}

func (s *UpdateCargowaybillBillResponse) SetReqMsgId(v string) *UpdateCargowaybillBillResponse {
	s.ReqMsgId = &v
	return s
}

func (s *UpdateCargowaybillBillResponse) SetResultCode(v string) *UpdateCargowaybillBillResponse {
	s.ResultCode = &v
	return s
}

func (s *UpdateCargowaybillBillResponse) SetResultMsg(v string) *UpdateCargowaybillBillResponse {
	s.ResultMsg = &v
	return s
}

func (s *UpdateCargowaybillBillResponse) SetTxCode(v string) *UpdateCargowaybillBillResponse {
	s.TxCode = &v
	return s
}

type CreateDisDidRequest struct {
	// OAuth模式下的授权token
	AuthToken         *string `json:"auth_token,omitempty" xml:"auth_token,omitempty"`
	ProductInstanceId *string `json:"product_instance_id,omitempty" xml:"product_instance_id,omitempty"`
	// 个人身份证号。当组织类型为个人时，此字段为必填项
	CertNo *string `json:"cert_no,omitempty" xml:"cert_no,omitempty"`
	// 企业名称。当组织类型为企业时，此字段为必填项
	EpCertName *string `json:"ep_cert_name,omitempty" xml:"ep_cert_name,omitempty"`
	// 企业信用号码。当组织类型为企业时，此字段为必填项
	EpCertNo *string `json:"ep_cert_no,omitempty" xml:"ep_cert_no,omitempty"`
	// 扩展字段
	ExtensionInfo *string `json:"extension_info,omitempty" xml:"extension_info,omitempty"`
	// 企业法人姓名。当申请企业类型网络货运平台或者子平台时，此字段为必填项
	LegalPersonCertName *string `json:"legal_person_cert_name,omitempty" xml:"legal_person_cert_name,omitempty"`
	// 企业法人身份证号码。当申请企业类型网络货运平台或者子平台时，此字段为必填项
	LegalPersonCertNo *string `json:"legal_person_cert_no,omitempty" xml:"legal_person_cert_no,omitempty"`
	// 个人手机号码。当组织类型为个人时，此字段为必填项
	Mobile *string `json:"mobile,omitempty" xml:"mobile,omitempty"`
	// 个人姓名。当组织类型为个人时，此字段为必填项
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// 组织类型。企业或者个人，二选一填入
	OrganizationType *string `json:"organization_type,omitempty" xml:"organization_type,omitempty" require:"true"`
	// 所属平台did。如果为空时，表示创建根平台，允许申请网络货运平台或者3pl角色。创建除根平台外的其他身份时，所属平台did必须填写。
	PlatformDid *string `json:"platform_did,omitempty" xml:"platform_did,omitempty"`
	// 角色类型。
	// 当组织类型为个人时，可填角色：货主、司机、承运商；
	// 当组织类型为企业时，可填角色：网络货运平台、道路运输企业/3pl、货主、子平台、承运商
	RoleType *string `json:"role_type,omitempty" xml:"role_type,omitempty" require:"true"`
}

func (s CreateDisDidRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateDisDidRequest) GoString() string {
	return s.String()
}

func (s *CreateDisDidRequest) SetAuthToken(v string) *CreateDisDidRequest {
	s.AuthToken = &v
	return s
}

func (s *CreateDisDidRequest) SetProductInstanceId(v string) *CreateDisDidRequest {
	s.ProductInstanceId = &v
	return s
}

func (s *CreateDisDidRequest) SetCertNo(v string) *CreateDisDidRequest {
	s.CertNo = &v
	return s
}

func (s *CreateDisDidRequest) SetEpCertName(v string) *CreateDisDidRequest {
	s.EpCertName = &v
	return s
}

func (s *CreateDisDidRequest) SetEpCertNo(v string) *CreateDisDidRequest {
	s.EpCertNo = &v
	return s
}

func (s *CreateDisDidRequest) SetExtensionInfo(v string) *CreateDisDidRequest {
	s.ExtensionInfo = &v
	return s
}

func (s *CreateDisDidRequest) SetLegalPersonCertName(v string) *CreateDisDidRequest {
	s.LegalPersonCertName = &v
	return s
}

func (s *CreateDisDidRequest) SetLegalPersonCertNo(v string) *CreateDisDidRequest {
	s.LegalPersonCertNo = &v
	return s
}

func (s *CreateDisDidRequest) SetMobile(v string) *CreateDisDidRequest {
	s.Mobile = &v
	return s
}

func (s *CreateDisDidRequest) SetName(v string) *CreateDisDidRequest {
	s.Name = &v
	return s
}

func (s *CreateDisDidRequest) SetOrganizationType(v string) *CreateDisDidRequest {
	s.OrganizationType = &v
	return s
}

func (s *CreateDisDidRequest) SetPlatformDid(v string) *CreateDisDidRequest {
	s.PlatformDid = &v
	return s
}

func (s *CreateDisDidRequest) SetRoleType(v string) *CreateDisDidRequest {
	s.RoleType = &v
	return s
}

type CreateDisDidResponse struct {
	// 请求唯一ID，用于链路跟踪和问题排查
	ReqMsgId *string `json:"req_msg_id,omitempty" xml:"req_msg_id,omitempty"`
	// 结果码，一般OK表示调用成功
	ResultCode *string `json:"result_code,omitempty" xml:"result_code,omitempty"`
	// 异常信息的文本描述
	ResultMsg *string `json:"result_msg,omitempty" xml:"result_msg,omitempty"`
	// 分布式数字身份did
	Did *string `json:"did,omitempty" xml:"did,omitempty"`
	// 组织类型
	OrganizationType *string `json:"organization_type,omitempty" xml:"organization_type,omitempty"`
	// 现阶段此did下的所有授予的角色
	RoleTypes []*string `json:"role_types,omitempty" xml:"role_types,omitempty" type:"Repeated"`
}

func (s CreateDisDidResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateDisDidResponse) GoString() string {
	return s.String()
}

func (s *CreateDisDidResponse) SetReqMsgId(v string) *CreateDisDidResponse {
	s.ReqMsgId = &v
	return s
}

func (s *CreateDisDidResponse) SetResultCode(v string) *CreateDisDidResponse {
	s.ResultCode = &v
	return s
}

func (s *CreateDisDidResponse) SetResultMsg(v string) *CreateDisDidResponse {
	s.ResultMsg = &v
	return s
}

func (s *CreateDisDidResponse) SetDid(v string) *CreateDisDidResponse {
	s.Did = &v
	return s
}

func (s *CreateDisDidResponse) SetOrganizationType(v string) *CreateDisDidResponse {
	s.OrganizationType = &v
	return s
}

func (s *CreateDisDidResponse) SetRoleTypes(v []*string) *CreateDisDidResponse {
	s.RoleTypes = v
	return s
}

type UploadTransportContractRequest struct {
	// OAuth模式下的授权token
	AuthToken         *string `json:"auth_token,omitempty" xml:"auth_token,omitempty"`
	ProductInstanceId *string `json:"product_instance_id,omitempty" xml:"product_instance_id,omitempty"`
	// 货主did，一般为合同甲方的链上数字身份
	ConsignorDid *string `json:"consignor_did,omitempty" xml:"consignor_did,omitempty" require:"true"`
	// 运输合同生效日期，格式要求yyyy-MM-dd
	ContractEffectiveDate *string `json:"contract_effective_date,omitempty" xml:"contract_effective_date,omitempty" require:"true"`
	// 运输合同到期日期，要求格式yyyy-MM-dd
	ContractExpiresDate *string `json:"contract_expires_date,omitempty" xml:"contract_expires_date,omitempty" require:"true"`
	// 影像件文件信息列表，可以包含多个文件，每个文件需要有文件id和文件hash  (请求蚂蚁影像上传接口获取的文件id和文件hash)。影像文件格式要求：bmp,jpg,jpeg,gif,psd,png,tiff,tga,eps,pdf
	FileInfos []*UploadFileInfo `json:"file_infos,omitempty" xml:"file_infos,omitempty" require:"true" type:"Repeated"`
	// 3plDid，一般为合同乙方的链上数字身份
	ThirdPartyLogisticsDid *string `json:"third_party_logistics_did,omitempty" xml:"third_party_logistics_did,omitempty" require:"true"`
	// 运输合同编号
	TransportContractCode *string `json:"transport_contract_code,omitempty" xml:"transport_contract_code,omitempty" require:"true"`
}

func (s UploadTransportContractRequest) String() string {
	return tea.Prettify(s)
}

func (s UploadTransportContractRequest) GoString() string {
	return s.String()
}

func (s *UploadTransportContractRequest) SetAuthToken(v string) *UploadTransportContractRequest {
	s.AuthToken = &v
	return s
}

func (s *UploadTransportContractRequest) SetProductInstanceId(v string) *UploadTransportContractRequest {
	s.ProductInstanceId = &v
	return s
}

func (s *UploadTransportContractRequest) SetConsignorDid(v string) *UploadTransportContractRequest {
	s.ConsignorDid = &v
	return s
}

func (s *UploadTransportContractRequest) SetContractEffectiveDate(v string) *UploadTransportContractRequest {
	s.ContractEffectiveDate = &v
	return s
}

func (s *UploadTransportContractRequest) SetContractExpiresDate(v string) *UploadTransportContractRequest {
	s.ContractExpiresDate = &v
	return s
}

func (s *UploadTransportContractRequest) SetFileInfos(v []*UploadFileInfo) *UploadTransportContractRequest {
	s.FileInfos = v
	return s
}

func (s *UploadTransportContractRequest) SetThirdPartyLogisticsDid(v string) *UploadTransportContractRequest {
	s.ThirdPartyLogisticsDid = &v
	return s
}

func (s *UploadTransportContractRequest) SetTransportContractCode(v string) *UploadTransportContractRequest {
	s.TransportContractCode = &v
	return s
}

type UploadTransportContractResponse struct {
	// 请求唯一ID，用于链路跟踪和问题排查
	ReqMsgId *string `json:"req_msg_id,omitempty" xml:"req_msg_id,omitempty"`
	// 结果码，一般OK表示调用成功
	ResultCode *string `json:"result_code,omitempty" xml:"result_code,omitempty"`
	// 异常信息的文本描述
	ResultMsg *string `json:"result_msg,omitempty" xml:"result_msg,omitempty"`
	// 链上凭证
	TxCode *string `json:"tx_code,omitempty" xml:"tx_code,omitempty"`
}

func (s UploadTransportContractResponse) String() string {
	return tea.Prettify(s)
}

func (s UploadTransportContractResponse) GoString() string {
	return s.String()
}

func (s *UploadTransportContractResponse) SetReqMsgId(v string) *UploadTransportContractResponse {
	s.ReqMsgId = &v
	return s
}

func (s *UploadTransportContractResponse) SetResultCode(v string) *UploadTransportContractResponse {
	s.ResultCode = &v
	return s
}

func (s *UploadTransportContractResponse) SetResultMsg(v string) *UploadTransportContractResponse {
	s.ResultMsg = &v
	return s
}

func (s *UploadTransportContractResponse) SetTxCode(v string) *UploadTransportContractResponse {
	s.TxCode = &v
	return s
}

type UploadTransportRouteRequest struct {
	// OAuth模式下的授权token
	AuthToken         *string `json:"auth_token,omitempty" xml:"auth_token,omitempty"`
	ProductInstanceId *string `json:"product_instance_id,omitempty" xml:"product_instance_id,omitempty"`
	// 目的地省市区，要求格式 XX省-XX市-XX区，比如四川省-成都市-青白江区。当线路类型为STOCK_IN（即入库物流）时，此字段必填
	EndAddress *string `json:"end_address,omitempty" xml:"end_address,omitempty"`
	// 目的地详细地址，街道村社区道路楼宇门牌号。当线路类型为STOCK_IN（即入库物流）时，此字段必填
	EndDetailedAddress *string `json:"end_detailed_address,omitempty" xml:"end_detailed_address,omitempty"`
	// 起始地省市区，要求格式 XX省-XX市-XX区。比如浙江省-杭州市-余杭区。当线路类型为STOCK_OUT（即出库物流）时，此字段必填
	StartAddress *string `json:"start_address,omitempty" xml:"start_address,omitempty"`
	// 起始地详细地址，街道村社区道路楼宇门牌号。当线路类型为STOCK_OUT（即出库物流）时，此字段必填
	StartDetailedAddress *string `json:"start_detailed_address,omitempty" xml:"start_detailed_address,omitempty"`
	// 3plDid
	ThirdPartyLogisticsDid *string `json:"third_party_logistics_did,omitempty" xml:"third_party_logistics_did,omitempty" require:"true"`
	// 运输合同编号
	TransportContractCode *string `json:"transport_contract_code,omitempty" xml:"transport_contract_code,omitempty" require:"true"`
	// 运输线路编码
	TransportRouteCode *string `json:"transport_route_code,omitempty" xml:"transport_route_code,omitempty" require:"true"`
	// 线路类型，以下二选一填写（可填STOCK_OUT、STOCK_IN）。注：以上分别表示出库物流、入库物流
	RouteType *string `json:"route_type,omitempty" xml:"route_type,omitempty" require:"true"`
}

func (s UploadTransportRouteRequest) String() string {
	return tea.Prettify(s)
}

func (s UploadTransportRouteRequest) GoString() string {
	return s.String()
}

func (s *UploadTransportRouteRequest) SetAuthToken(v string) *UploadTransportRouteRequest {
	s.AuthToken = &v
	return s
}

func (s *UploadTransportRouteRequest) SetProductInstanceId(v string) *UploadTransportRouteRequest {
	s.ProductInstanceId = &v
	return s
}

func (s *UploadTransportRouteRequest) SetEndAddress(v string) *UploadTransportRouteRequest {
	s.EndAddress = &v
	return s
}

func (s *UploadTransportRouteRequest) SetEndDetailedAddress(v string) *UploadTransportRouteRequest {
	s.EndDetailedAddress = &v
	return s
}

func (s *UploadTransportRouteRequest) SetStartAddress(v string) *UploadTransportRouteRequest {
	s.StartAddress = &v
	return s
}

func (s *UploadTransportRouteRequest) SetStartDetailedAddress(v string) *UploadTransportRouteRequest {
	s.StartDetailedAddress = &v
	return s
}

func (s *UploadTransportRouteRequest) SetThirdPartyLogisticsDid(v string) *UploadTransportRouteRequest {
	s.ThirdPartyLogisticsDid = &v
	return s
}

func (s *UploadTransportRouteRequest) SetTransportContractCode(v string) *UploadTransportRouteRequest {
	s.TransportContractCode = &v
	return s
}

func (s *UploadTransportRouteRequest) SetTransportRouteCode(v string) *UploadTransportRouteRequest {
	s.TransportRouteCode = &v
	return s
}

func (s *UploadTransportRouteRequest) SetRouteType(v string) *UploadTransportRouteRequest {
	s.RouteType = &v
	return s
}

type UploadTransportRouteResponse struct {
	// 请求唯一ID，用于链路跟踪和问题排查
	ReqMsgId *string `json:"req_msg_id,omitempty" xml:"req_msg_id,omitempty"`
	// 结果码，一般OK表示调用成功
	ResultCode *string `json:"result_code,omitempty" xml:"result_code,omitempty"`
	// 异常信息的文本描述
	ResultMsg *string `json:"result_msg,omitempty" xml:"result_msg,omitempty"`
	// 链上凭证
	TxCode *string `json:"tx_code,omitempty" xml:"tx_code,omitempty"`
}

func (s UploadTransportRouteResponse) String() string {
	return tea.Prettify(s)
}

func (s UploadTransportRouteResponse) GoString() string {
	return s.String()
}

func (s *UploadTransportRouteResponse) SetReqMsgId(v string) *UploadTransportRouteResponse {
	s.ReqMsgId = &v
	return s
}

func (s *UploadTransportRouteResponse) SetResultCode(v string) *UploadTransportRouteResponse {
	s.ResultCode = &v
	return s
}

func (s *UploadTransportRouteResponse) SetResultMsg(v string) *UploadTransportRouteResponse {
	s.ResultMsg = &v
	return s
}

func (s *UploadTransportRouteResponse) SetTxCode(v string) *UploadTransportRouteResponse {
	s.TxCode = &v
	return s
}

type CreateTransportWaybillRequest struct {
	// OAuth模式下的授权token
	AuthToken         *string `json:"auth_token,omitempty" xml:"auth_token,omitempty"`
	ProductInstanceId *string `json:"product_instance_id,omitempty" xml:"product_instance_id,omitempty"`
	// 运费，3pl支付给下承运商的运费金额，单位（元），要求格式为不超过二位小数
	AllFreight *string `json:"all_freight,omitempty" xml:"all_freight,omitempty" require:"true"`
	// 货物名称
	CargoName *string `json:"cargo_name,omitempty" xml:"cargo_name,omitempty" require:"true"`
	// 货物体积，单位（方）。货物重量、货物体积二选一填写
	CargoVolume *string `json:"cargo_volume,omitempty" xml:"cargo_volume,omitempty"`
	// 货物重量，单位（吨）。货物重量、货物体积二选一填写
	CargoWeight *string `json:"cargo_weight,omitempty" xml:"cargo_weight,omitempty"`
	// 下游承运商did，一般为下一级承运商数字身份
	CarrierDid *string `json:"carrier_did,omitempty" xml:"carrier_did,omitempty" require:"true"`
	// 车牌颜色，需填写黄色、蓝色、绿色中的一种
	CarBadgeColor *string `json:"car_badge_color,omitempty" xml:"car_badge_color,omitempty" require:"true"`
	// 车牌号，承运车牌号
	CarBadgeNo *string `json:"car_badge_no,omitempty" xml:"car_badge_no,omitempty" require:"true"`
	// 货主did，一般为合同甲方的链上数字身份
	ConsignorDid *string `json:"consignor_did,omitempty" xml:"consignor_did,omitempty" require:"true"`
	// 创建时间，13位毫秒级时间戳
	CreateTime *int64 `json:"create_time,omitempty" xml:"create_time,omitempty" require:"true"`
	// 实际承运司机did，实际承运司机的链上数字身份
	DriverDid *string `json:"driver_did,omitempty" xml:"driver_did,omitempty" require:"true"`
	// 目的地省市区，要求格式 XX省-XX市-XX区，比如四川省-成都市-青白江区
	EndAddress *string `json:"end_address,omitempty" xml:"end_address,omitempty" require:"true"`
	// 目的地详细地址，街道村社区道路楼宇门牌号
	EndDetailedAddress *string `json:"end_detailed_address,omitempty" xml:"end_detailed_address,omitempty" require:"true"`
	// 货物数量
	GoodsAmount *int64 `json:"goods_amount,omitempty" xml:"goods_amount,omitempty"`
	// 货物数量单位类型
	GoodsAmountType *string `json:"goods_amount_type,omitempty" xml:"goods_amount_type,omitempty"`
	// 起始地省市区，要求格式 XX省-XX市-XX区。比如浙江省-杭州市-余杭区
	StartAddress *string `json:"start_address,omitempty" xml:"start_address,omitempty" require:"true"`
	// 起始地详细地址，街道村社区道路楼宇门牌号
	StartDetailedAddress *string `json:"start_detailed_address,omitempty" xml:"start_detailed_address,omitempty" require:"true"`
	// 客户系统内运单编号
	TaxWaybillId *string `json:"tax_waybill_id,omitempty" xml:"tax_waybill_id,omitempty" require:"true"`
	// 3plDid，一般为合同乙方的链上数字身份
	ThirdPartyLogisticsDid *string `json:"third_party_logistics_did,omitempty" xml:"third_party_logistics_did,omitempty" require:"true"`
	// 所属合同编号
	TransportContractCode *string `json:"transport_contract_code,omitempty" xml:"transport_contract_code,omitempty"`
	// 所属运输线路编码
	TransportRouteCode *string `json:"transport_route_code,omitempty" xml:"transport_route_code,omitempty"`
}

func (s CreateTransportWaybillRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateTransportWaybillRequest) GoString() string {
	return s.String()
}

func (s *CreateTransportWaybillRequest) SetAuthToken(v string) *CreateTransportWaybillRequest {
	s.AuthToken = &v
	return s
}

func (s *CreateTransportWaybillRequest) SetProductInstanceId(v string) *CreateTransportWaybillRequest {
	s.ProductInstanceId = &v
	return s
}

func (s *CreateTransportWaybillRequest) SetAllFreight(v string) *CreateTransportWaybillRequest {
	s.AllFreight = &v
	return s
}

func (s *CreateTransportWaybillRequest) SetCargoName(v string) *CreateTransportWaybillRequest {
	s.CargoName = &v
	return s
}

func (s *CreateTransportWaybillRequest) SetCargoVolume(v string) *CreateTransportWaybillRequest {
	s.CargoVolume = &v
	return s
}

func (s *CreateTransportWaybillRequest) SetCargoWeight(v string) *CreateTransportWaybillRequest {
	s.CargoWeight = &v
	return s
}

func (s *CreateTransportWaybillRequest) SetCarrierDid(v string) *CreateTransportWaybillRequest {
	s.CarrierDid = &v
	return s
}

func (s *CreateTransportWaybillRequest) SetCarBadgeColor(v string) *CreateTransportWaybillRequest {
	s.CarBadgeColor = &v
	return s
}

func (s *CreateTransportWaybillRequest) SetCarBadgeNo(v string) *CreateTransportWaybillRequest {
	s.CarBadgeNo = &v
	return s
}

func (s *CreateTransportWaybillRequest) SetConsignorDid(v string) *CreateTransportWaybillRequest {
	s.ConsignorDid = &v
	return s
}

func (s *CreateTransportWaybillRequest) SetCreateTime(v int64) *CreateTransportWaybillRequest {
	s.CreateTime = &v
	return s
}

func (s *CreateTransportWaybillRequest) SetDriverDid(v string) *CreateTransportWaybillRequest {
	s.DriverDid = &v
	return s
}

func (s *CreateTransportWaybillRequest) SetEndAddress(v string) *CreateTransportWaybillRequest {
	s.EndAddress = &v
	return s
}

func (s *CreateTransportWaybillRequest) SetEndDetailedAddress(v string) *CreateTransportWaybillRequest {
	s.EndDetailedAddress = &v
	return s
}

func (s *CreateTransportWaybillRequest) SetGoodsAmount(v int64) *CreateTransportWaybillRequest {
	s.GoodsAmount = &v
	return s
}

func (s *CreateTransportWaybillRequest) SetGoodsAmountType(v string) *CreateTransportWaybillRequest {
	s.GoodsAmountType = &v
	return s
}

func (s *CreateTransportWaybillRequest) SetStartAddress(v string) *CreateTransportWaybillRequest {
	s.StartAddress = &v
	return s
}

func (s *CreateTransportWaybillRequest) SetStartDetailedAddress(v string) *CreateTransportWaybillRequest {
	s.StartDetailedAddress = &v
	return s
}

func (s *CreateTransportWaybillRequest) SetTaxWaybillId(v string) *CreateTransportWaybillRequest {
	s.TaxWaybillId = &v
	return s
}

func (s *CreateTransportWaybillRequest) SetThirdPartyLogisticsDid(v string) *CreateTransportWaybillRequest {
	s.ThirdPartyLogisticsDid = &v
	return s
}

func (s *CreateTransportWaybillRequest) SetTransportContractCode(v string) *CreateTransportWaybillRequest {
	s.TransportContractCode = &v
	return s
}

func (s *CreateTransportWaybillRequest) SetTransportRouteCode(v string) *CreateTransportWaybillRequest {
	s.TransportRouteCode = &v
	return s
}

type CreateTransportWaybillResponse struct {
	// 请求唯一ID，用于链路跟踪和问题排查
	ReqMsgId *string `json:"req_msg_id,omitempty" xml:"req_msg_id,omitempty"`
	// 结果码，一般OK表示调用成功
	ResultCode *string `json:"result_code,omitempty" xml:"result_code,omitempty"`
	// 异常信息的文本描述
	ResultMsg *string `json:"result_msg,omitempty" xml:"result_msg,omitempty"`
	// 链上凭证
	TxCode *string `json:"tx_code,omitempty" xml:"tx_code,omitempty"`
}

func (s CreateTransportWaybillResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateTransportWaybillResponse) GoString() string {
	return s.String()
}

func (s *CreateTransportWaybillResponse) SetReqMsgId(v string) *CreateTransportWaybillResponse {
	s.ReqMsgId = &v
	return s
}

func (s *CreateTransportWaybillResponse) SetResultCode(v string) *CreateTransportWaybillResponse {
	s.ResultCode = &v
	return s
}

func (s *CreateTransportWaybillResponse) SetResultMsg(v string) *CreateTransportWaybillResponse {
	s.ResultMsg = &v
	return s
}

func (s *CreateTransportWaybillResponse) SetTxCode(v string) *CreateTransportWaybillResponse {
	s.TxCode = &v
	return s
}

type UpdateWaybillActionRequest struct {
	// OAuth模式下的授权token
	AuthToken         *string `json:"auth_token,omitempty" xml:"auth_token,omitempty"`
	ProductInstanceId *string `json:"product_instance_id,omitempty" xml:"product_instance_id,omitempty"`
	// 操作类型，以下二选一填写：运单起运、运输完成
	OperateAction *string `json:"operate_action,omitempty" xml:"operate_action,omitempty" require:"true"`
	// 运单起运或运输完成的时间，要求为13位毫秒级时间戳
	OperateActionTime *int64 `json:"operate_action_time,omitempty" xml:"operate_action_time,omitempty" require:"true"`
	// 客户系统内运单编号
	TaxWaybillId *string `json:"tax_waybill_id,omitempty" xml:"tax_waybill_id,omitempty" require:"true"`
	// 3plDid，一般为合同乙方的链上数字身份
	ThirdPartyLogisticsDid *string `json:"third_party_logistics_did,omitempty" xml:"third_party_logistics_did,omitempty" require:"true"`
}

func (s UpdateWaybillActionRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateWaybillActionRequest) GoString() string {
	return s.String()
}

func (s *UpdateWaybillActionRequest) SetAuthToken(v string) *UpdateWaybillActionRequest {
	s.AuthToken = &v
	return s
}

func (s *UpdateWaybillActionRequest) SetProductInstanceId(v string) *UpdateWaybillActionRequest {
	s.ProductInstanceId = &v
	return s
}

func (s *UpdateWaybillActionRequest) SetOperateAction(v string) *UpdateWaybillActionRequest {
	s.OperateAction = &v
	return s
}

func (s *UpdateWaybillActionRequest) SetOperateActionTime(v int64) *UpdateWaybillActionRequest {
	s.OperateActionTime = &v
	return s
}

func (s *UpdateWaybillActionRequest) SetTaxWaybillId(v string) *UpdateWaybillActionRequest {
	s.TaxWaybillId = &v
	return s
}

func (s *UpdateWaybillActionRequest) SetThirdPartyLogisticsDid(v string) *UpdateWaybillActionRequest {
	s.ThirdPartyLogisticsDid = &v
	return s
}

type UpdateWaybillActionResponse struct {
	// 请求唯一ID，用于链路跟踪和问题排查
	ReqMsgId *string `json:"req_msg_id,omitempty" xml:"req_msg_id,omitempty"`
	// 结果码，一般OK表示调用成功
	ResultCode *string `json:"result_code,omitempty" xml:"result_code,omitempty"`
	// 异常信息的文本描述
	ResultMsg *string `json:"result_msg,omitempty" xml:"result_msg,omitempty"`
	// 链上凭证
	TxCode *string `json:"tx_code,omitempty" xml:"tx_code,omitempty"`
}

func (s UpdateWaybillActionResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateWaybillActionResponse) GoString() string {
	return s.String()
}

func (s *UpdateWaybillActionResponse) SetReqMsgId(v string) *UpdateWaybillActionResponse {
	s.ReqMsgId = &v
	return s
}

func (s *UpdateWaybillActionResponse) SetResultCode(v string) *UpdateWaybillActionResponse {
	s.ResultCode = &v
	return s
}

func (s *UpdateWaybillActionResponse) SetResultMsg(v string) *UpdateWaybillActionResponse {
	s.ResultMsg = &v
	return s
}

func (s *UpdateWaybillActionResponse) SetTxCode(v string) *UpdateWaybillActionResponse {
	s.TxCode = &v
	return s
}

type UpdateTransportWaybillRequest struct {
	// OAuth模式下的授权token
	AuthToken         *string `json:"auth_token,omitempty" xml:"auth_token,omitempty"`
	ProductInstanceId *string `json:"product_instance_id,omitempty" xml:"product_instance_id,omitempty"`
	// 运费，3pl支付给下承运商的运费金额，单位（元），要求格式为不超过二位小数
	AllFreight *string `json:"all_freight,omitempty" xml:"all_freight,omitempty"`
	// 货物名称
	CargoName *string `json:"cargo_name,omitempty" xml:"cargo_name,omitempty"`
	// 单位（方），货物体积
	CargoVolume *string `json:"cargo_volume,omitempty" xml:"cargo_volume,omitempty"`
	// 单位（吨），货物重量
	CargoWeight *string `json:"cargo_weight,omitempty" xml:"cargo_weight,omitempty"`
	// 下游承运商did，一般为下一级承运商数字身份
	CarrierDid *string `json:"carrier_did,omitempty" xml:"carrier_did,omitempty"`
	// 车牌颜色，需填写黄色、蓝色、绿色中的一种
	CarBadgeColor *string `json:"car_badge_color,omitempty" xml:"car_badge_color,omitempty"`
	// 车牌号，承运车牌号
	CarBadgeNo *string `json:"car_badge_no,omitempty" xml:"car_badge_no,omitempty"`
	// 货主did，一般为合同甲方的链上数字身份
	ConsignorDid *string `json:"consignor_did,omitempty" xml:"consignor_did,omitempty"`
	// 创建时间，13位毫秒级时间戳
	CreateTime *int64 `json:"create_time,omitempty" xml:"create_time,omitempty"`
	// 实际承运司机did，实际承运司机的链上数字身份
	DriverDid *string `json:"driver_did,omitempty" xml:"driver_did,omitempty"`
	// 目的地省市区，要求格式 XX省-XX市-XX区，比如四川省-成都市-青白江区
	EndAddress *string `json:"end_address,omitempty" xml:"end_address,omitempty"`
	// 目的地详细地址，街道村社区道路楼宇门牌号
	EndDetailedAddress *string `json:"end_detailed_address,omitempty" xml:"end_detailed_address,omitempty"`
	// 到达时间，13位毫秒级时间戳
	EndTime *int64 `json:"end_time,omitempty" xml:"end_time,omitempty"`
	// 货物数量
	GoodsAmount *int64 `json:"goods_amount,omitempty" xml:"goods_amount,omitempty"`
	// 货物数量单位类型
	GoodsAmountType *string `json:"goods_amount_type,omitempty" xml:"goods_amount_type,omitempty"`
	// 起始地省市区，要求格式 XX省-XX市-XX区。比如浙江省-杭州市-余杭区
	StartAddress *string `json:"start_address,omitempty" xml:"start_address,omitempty"`
	// 起始地详细地址，街道村社区道路楼宇门牌号
	StartDetailedAddress *string `json:"start_detailed_address,omitempty" xml:"start_detailed_address,omitempty"`
	// 起运时间，13位毫秒级时间戳
	StartTime *int64 `json:"start_time,omitempty" xml:"start_time,omitempty"`
	// 客户系统内运单编号
	TaxWaybillId *string `json:"tax_waybill_id,omitempty" xml:"tax_waybill_id,omitempty" require:"true"`
	// 3plDid，一般为合同乙方的链上数字身份
	ThirdPartyLogisticsDid *string `json:"third_party_logistics_did,omitempty" xml:"third_party_logistics_did,omitempty" require:"true"`
	// 所属合同编号
	TransportContractCode *string `json:"transport_contract_code,omitempty" xml:"transport_contract_code,omitempty"`
	// 所属运输线路编码
	TransportRouteCode *string `json:"transport_route_code,omitempty" xml:"transport_route_code,omitempty"`
}

func (s UpdateTransportWaybillRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateTransportWaybillRequest) GoString() string {
	return s.String()
}

func (s *UpdateTransportWaybillRequest) SetAuthToken(v string) *UpdateTransportWaybillRequest {
	s.AuthToken = &v
	return s
}

func (s *UpdateTransportWaybillRequest) SetProductInstanceId(v string) *UpdateTransportWaybillRequest {
	s.ProductInstanceId = &v
	return s
}

func (s *UpdateTransportWaybillRequest) SetAllFreight(v string) *UpdateTransportWaybillRequest {
	s.AllFreight = &v
	return s
}

func (s *UpdateTransportWaybillRequest) SetCargoName(v string) *UpdateTransportWaybillRequest {
	s.CargoName = &v
	return s
}

func (s *UpdateTransportWaybillRequest) SetCargoVolume(v string) *UpdateTransportWaybillRequest {
	s.CargoVolume = &v
	return s
}

func (s *UpdateTransportWaybillRequest) SetCargoWeight(v string) *UpdateTransportWaybillRequest {
	s.CargoWeight = &v
	return s
}

func (s *UpdateTransportWaybillRequest) SetCarrierDid(v string) *UpdateTransportWaybillRequest {
	s.CarrierDid = &v
	return s
}

func (s *UpdateTransportWaybillRequest) SetCarBadgeColor(v string) *UpdateTransportWaybillRequest {
	s.CarBadgeColor = &v
	return s
}

func (s *UpdateTransportWaybillRequest) SetCarBadgeNo(v string) *UpdateTransportWaybillRequest {
	s.CarBadgeNo = &v
	return s
}

func (s *UpdateTransportWaybillRequest) SetConsignorDid(v string) *UpdateTransportWaybillRequest {
	s.ConsignorDid = &v
	return s
}

func (s *UpdateTransportWaybillRequest) SetCreateTime(v int64) *UpdateTransportWaybillRequest {
	s.CreateTime = &v
	return s
}

func (s *UpdateTransportWaybillRequest) SetDriverDid(v string) *UpdateTransportWaybillRequest {
	s.DriverDid = &v
	return s
}

func (s *UpdateTransportWaybillRequest) SetEndAddress(v string) *UpdateTransportWaybillRequest {
	s.EndAddress = &v
	return s
}

func (s *UpdateTransportWaybillRequest) SetEndDetailedAddress(v string) *UpdateTransportWaybillRequest {
	s.EndDetailedAddress = &v
	return s
}

func (s *UpdateTransportWaybillRequest) SetEndTime(v int64) *UpdateTransportWaybillRequest {
	s.EndTime = &v
	return s
}

func (s *UpdateTransportWaybillRequest) SetGoodsAmount(v int64) *UpdateTransportWaybillRequest {
	s.GoodsAmount = &v
	return s
}

func (s *UpdateTransportWaybillRequest) SetGoodsAmountType(v string) *UpdateTransportWaybillRequest {
	s.GoodsAmountType = &v
	return s
}

func (s *UpdateTransportWaybillRequest) SetStartAddress(v string) *UpdateTransportWaybillRequest {
	s.StartAddress = &v
	return s
}

func (s *UpdateTransportWaybillRequest) SetStartDetailedAddress(v string) *UpdateTransportWaybillRequest {
	s.StartDetailedAddress = &v
	return s
}

func (s *UpdateTransportWaybillRequest) SetStartTime(v int64) *UpdateTransportWaybillRequest {
	s.StartTime = &v
	return s
}

func (s *UpdateTransportWaybillRequest) SetTaxWaybillId(v string) *UpdateTransportWaybillRequest {
	s.TaxWaybillId = &v
	return s
}

func (s *UpdateTransportWaybillRequest) SetThirdPartyLogisticsDid(v string) *UpdateTransportWaybillRequest {
	s.ThirdPartyLogisticsDid = &v
	return s
}

func (s *UpdateTransportWaybillRequest) SetTransportContractCode(v string) *UpdateTransportWaybillRequest {
	s.TransportContractCode = &v
	return s
}

func (s *UpdateTransportWaybillRequest) SetTransportRouteCode(v string) *UpdateTransportWaybillRequest {
	s.TransportRouteCode = &v
	return s
}

type UpdateTransportWaybillResponse struct {
	// 请求唯一ID，用于链路跟踪和问题排查
	ReqMsgId *string `json:"req_msg_id,omitempty" xml:"req_msg_id,omitempty"`
	// 结果码，一般OK表示调用成功
	ResultCode *string `json:"result_code,omitempty" xml:"result_code,omitempty"`
	// 异常信息的文本描述
	ResultMsg *string `json:"result_msg,omitempty" xml:"result_msg,omitempty"`
	// 链上凭证
	TxCode *string `json:"tx_code,omitempty" xml:"tx_code,omitempty"`
}

func (s UpdateTransportWaybillResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateTransportWaybillResponse) GoString() string {
	return s.String()
}

func (s *UpdateTransportWaybillResponse) SetReqMsgId(v string) *UpdateTransportWaybillResponse {
	s.ReqMsgId = &v
	return s
}

func (s *UpdateTransportWaybillResponse) SetResultCode(v string) *UpdateTransportWaybillResponse {
	s.ResultCode = &v
	return s
}

func (s *UpdateTransportWaybillResponse) SetResultMsg(v string) *UpdateTransportWaybillResponse {
	s.ResultMsg = &v
	return s
}

func (s *UpdateTransportWaybillResponse) SetTxCode(v string) *UpdateTransportWaybillResponse {
	s.TxCode = &v
	return s
}

type UploadTransportReceiptRequest struct {
	// OAuth模式下的授权token
	AuthToken         *string `json:"auth_token,omitempty" xml:"auth_token,omitempty"`
	ProductInstanceId *string `json:"product_instance_id,omitempty" xml:"product_instance_id,omitempty"`
	// 货主did，一般为合同甲方的链上数字身份
	ConsignorDid *string `json:"consignor_did,omitempty" xml:"consignor_did,omitempty" require:"true"`
	// 请求蚂蚁影像上传接口上传文件获取的文件id和文件hash信息。影像文件格式：bmp,jpg,jpeg,gif,psd,png,tiff,tga,eps,pdf
	FileInfo *UploadFileInfo `json:"file_info,omitempty" xml:"file_info,omitempty" require:"true"`
	// 回单id，客户编辑的唯一回单编码
	ReceiptId *string `json:"receipt_id,omitempty" xml:"receipt_id,omitempty" require:"true"`
	// 客户系统内运单编号
	TaxWaybillId *string `json:"tax_waybill_id,omitempty" xml:"tax_waybill_id,omitempty" require:"true"`
	// 3plDid，一般为合同乙方的链上数字身份
	ThirdPartyLogisticsDid *string `json:"third_party_logistics_did,omitempty" xml:"third_party_logistics_did,omitempty" require:"true"`
}

func (s UploadTransportReceiptRequest) String() string {
	return tea.Prettify(s)
}

func (s UploadTransportReceiptRequest) GoString() string {
	return s.String()
}

func (s *UploadTransportReceiptRequest) SetAuthToken(v string) *UploadTransportReceiptRequest {
	s.AuthToken = &v
	return s
}

func (s *UploadTransportReceiptRequest) SetProductInstanceId(v string) *UploadTransportReceiptRequest {
	s.ProductInstanceId = &v
	return s
}

func (s *UploadTransportReceiptRequest) SetConsignorDid(v string) *UploadTransportReceiptRequest {
	s.ConsignorDid = &v
	return s
}

func (s *UploadTransportReceiptRequest) SetFileInfo(v *UploadFileInfo) *UploadTransportReceiptRequest {
	s.FileInfo = v
	return s
}

func (s *UploadTransportReceiptRequest) SetReceiptId(v string) *UploadTransportReceiptRequest {
	s.ReceiptId = &v
	return s
}

func (s *UploadTransportReceiptRequest) SetTaxWaybillId(v string) *UploadTransportReceiptRequest {
	s.TaxWaybillId = &v
	return s
}

func (s *UploadTransportReceiptRequest) SetThirdPartyLogisticsDid(v string) *UploadTransportReceiptRequest {
	s.ThirdPartyLogisticsDid = &v
	return s
}

type UploadTransportReceiptResponse struct {
	// 请求唯一ID，用于链路跟踪和问题排查
	ReqMsgId *string `json:"req_msg_id,omitempty" xml:"req_msg_id,omitempty"`
	// 结果码，一般OK表示调用成功
	ResultCode *string `json:"result_code,omitempty" xml:"result_code,omitempty"`
	// 异常信息的文本描述
	ResultMsg *string `json:"result_msg,omitempty" xml:"result_msg,omitempty"`
	// 链上凭证
	TxCode *string `json:"tx_code,omitempty" xml:"tx_code,omitempty"`
}

func (s UploadTransportReceiptResponse) String() string {
	return tea.Prettify(s)
}

func (s UploadTransportReceiptResponse) GoString() string {
	return s.String()
}

func (s *UploadTransportReceiptResponse) SetReqMsgId(v string) *UploadTransportReceiptResponse {
	s.ReqMsgId = &v
	return s
}

func (s *UploadTransportReceiptResponse) SetResultCode(v string) *UploadTransportReceiptResponse {
	s.ResultCode = &v
	return s
}

func (s *UploadTransportReceiptResponse) SetResultMsg(v string) *UploadTransportReceiptResponse {
	s.ResultMsg = &v
	return s
}

func (s *UploadTransportReceiptResponse) SetTxCode(v string) *UploadTransportReceiptResponse {
	s.TxCode = &v
	return s
}

type CreateBillReceivablebillRequest struct {
	// OAuth模式下的授权token
	AuthToken         *string `json:"auth_token,omitempty" xml:"auth_token,omitempty"`
	ProductInstanceId *string `json:"product_instance_id,omitempty" xml:"product_instance_id,omitempty"`
	// 账单总金额，单位（元），周期内应向发货方收取的运费总金额，不超过2位小数的数字
	BillAmount *string `json:"bill_amount,omitempty" xml:"bill_amount,omitempty" require:"true"`
	// 账单生成时间，13位毫秒级时间戳
	BillCreateTime *int64 `json:"bill_create_time,omitempty" xml:"bill_create_time,omitempty" require:"true"`
	// 账单期限，单位（天），合同约定的结算周期
	BillDeadline *int64 `json:"bill_deadline,omitempty" xml:"bill_deadline,omitempty" require:"true"`
	// 账单id，客户生成的账单唯一编号
	BillId *string `json:"bill_id,omitempty" xml:"bill_id,omitempty" require:"true"`
	// 收款方did，账单的收款方数字身份
	BillPayeeDid *string `json:"bill_payee_did,omitempty" xml:"bill_payee_did,omitempty" require:"true"`
	// 付款方did，账单的付款方数字身份
	BillPayerDid *string `json:"bill_payer_did,omitempty" xml:"bill_payer_did,omitempty" require:"true"`
	// 账单起始日期，13位毫秒级时间戳
	BillStartTime *int64 `json:"bill_start_time,omitempty" xml:"bill_start_time,omitempty" require:"true"`
	// 关联合同编号，账单关联的合同编号，如为合同物流请填写
	//
	ContractCode *string `json:"contract_code,omitempty" xml:"contract_code,omitempty"`
	// 账单到期日期，13位毫秒级时间戳
	Deadline *int64 `json:"deadline,omitempty" xml:"deadline,omitempty" require:"true"`
	// 账单关联运单号数组，元素个数不能超过1000个
	WaybillIds []*string `json:"waybill_ids,omitempty" xml:"waybill_ids,omitempty" require:"true" type:"Repeated"`
}

func (s CreateBillReceivablebillRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateBillReceivablebillRequest) GoString() string {
	return s.String()
}

func (s *CreateBillReceivablebillRequest) SetAuthToken(v string) *CreateBillReceivablebillRequest {
	s.AuthToken = &v
	return s
}

func (s *CreateBillReceivablebillRequest) SetProductInstanceId(v string) *CreateBillReceivablebillRequest {
	s.ProductInstanceId = &v
	return s
}

func (s *CreateBillReceivablebillRequest) SetBillAmount(v string) *CreateBillReceivablebillRequest {
	s.BillAmount = &v
	return s
}

func (s *CreateBillReceivablebillRequest) SetBillCreateTime(v int64) *CreateBillReceivablebillRequest {
	s.BillCreateTime = &v
	return s
}

func (s *CreateBillReceivablebillRequest) SetBillDeadline(v int64) *CreateBillReceivablebillRequest {
	s.BillDeadline = &v
	return s
}

func (s *CreateBillReceivablebillRequest) SetBillId(v string) *CreateBillReceivablebillRequest {
	s.BillId = &v
	return s
}

func (s *CreateBillReceivablebillRequest) SetBillPayeeDid(v string) *CreateBillReceivablebillRequest {
	s.BillPayeeDid = &v
	return s
}

func (s *CreateBillReceivablebillRequest) SetBillPayerDid(v string) *CreateBillReceivablebillRequest {
	s.BillPayerDid = &v
	return s
}

func (s *CreateBillReceivablebillRequest) SetBillStartTime(v int64) *CreateBillReceivablebillRequest {
	s.BillStartTime = &v
	return s
}

func (s *CreateBillReceivablebillRequest) SetContractCode(v string) *CreateBillReceivablebillRequest {
	s.ContractCode = &v
	return s
}

func (s *CreateBillReceivablebillRequest) SetDeadline(v int64) *CreateBillReceivablebillRequest {
	s.Deadline = &v
	return s
}

func (s *CreateBillReceivablebillRequest) SetWaybillIds(v []*string) *CreateBillReceivablebillRequest {
	s.WaybillIds = v
	return s
}

type CreateBillReceivablebillResponse struct {
	// 请求唯一ID，用于链路跟踪和问题排查
	ReqMsgId *string `json:"req_msg_id,omitempty" xml:"req_msg_id,omitempty"`
	// 结果码，一般OK表示调用成功
	ResultCode *string `json:"result_code,omitempty" xml:"result_code,omitempty"`
	// 异常信息的文本描述
	ResultMsg *string `json:"result_msg,omitempty" xml:"result_msg,omitempty"`
	// 链上凭证
	TxCode *string `json:"tx_code,omitempty" xml:"tx_code,omitempty"`
}

func (s CreateBillReceivablebillResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateBillReceivablebillResponse) GoString() string {
	return s.String()
}

func (s *CreateBillReceivablebillResponse) SetReqMsgId(v string) *CreateBillReceivablebillResponse {
	s.ReqMsgId = &v
	return s
}

func (s *CreateBillReceivablebillResponse) SetResultCode(v string) *CreateBillReceivablebillResponse {
	s.ResultCode = &v
	return s
}

func (s *CreateBillReceivablebillResponse) SetResultMsg(v string) *CreateBillReceivablebillResponse {
	s.ResultMsg = &v
	return s
}

func (s *CreateBillReceivablebillResponse) SetTxCode(v string) *CreateBillReceivablebillResponse {
	s.TxCode = &v
	return s
}

type UpdateReceivablebillStatusRequest struct {
	// OAuth模式下的授权token
	AuthToken         *string `json:"auth_token,omitempty" xml:"auth_token,omitempty"`
	ProductInstanceId *string `json:"product_instance_id,omitempty" xml:"product_instance_id,omitempty"`
	// 账单id，客户生成的账单唯一编号
	BillId *string `json:"bill_id,omitempty" xml:"bill_id,omitempty" require:"true"`
	// 收款方did，账单的收款方数字身份
	BillPayeeDid *string `json:"bill_payee_did,omitempty" xml:"bill_payee_did,omitempty" require:"true"`
	// 账单后续所可能产生的状态，以下三选一填写：账单确认、部分结算、已结清
	//
	//
	UpdateStatus *string `json:"update_status,omitempty" xml:"update_status,omitempty" require:"true"`
	// 更新状态时间，13位毫秒级时间戳
	UpdateStatusTime *int64 `json:"update_status_time,omitempty" xml:"update_status_time,omitempty" require:"true"`
}

func (s UpdateReceivablebillStatusRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateReceivablebillStatusRequest) GoString() string {
	return s.String()
}

func (s *UpdateReceivablebillStatusRequest) SetAuthToken(v string) *UpdateReceivablebillStatusRequest {
	s.AuthToken = &v
	return s
}

func (s *UpdateReceivablebillStatusRequest) SetProductInstanceId(v string) *UpdateReceivablebillStatusRequest {
	s.ProductInstanceId = &v
	return s
}

func (s *UpdateReceivablebillStatusRequest) SetBillId(v string) *UpdateReceivablebillStatusRequest {
	s.BillId = &v
	return s
}

func (s *UpdateReceivablebillStatusRequest) SetBillPayeeDid(v string) *UpdateReceivablebillStatusRequest {
	s.BillPayeeDid = &v
	return s
}

func (s *UpdateReceivablebillStatusRequest) SetUpdateStatus(v string) *UpdateReceivablebillStatusRequest {
	s.UpdateStatus = &v
	return s
}

func (s *UpdateReceivablebillStatusRequest) SetUpdateStatusTime(v int64) *UpdateReceivablebillStatusRequest {
	s.UpdateStatusTime = &v
	return s
}

type UpdateReceivablebillStatusResponse struct {
	// 请求唯一ID，用于链路跟踪和问题排查
	ReqMsgId *string `json:"req_msg_id,omitempty" xml:"req_msg_id,omitempty"`
	// 结果码，一般OK表示调用成功
	ResultCode *string `json:"result_code,omitempty" xml:"result_code,omitempty"`
	// 异常信息的文本描述
	ResultMsg *string `json:"result_msg,omitempty" xml:"result_msg,omitempty"`
	// 链上凭证
	TxCode *string `json:"tx_code,omitempty" xml:"tx_code,omitempty"`
}

func (s UpdateReceivablebillStatusResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateReceivablebillStatusResponse) GoString() string {
	return s.String()
}

func (s *UpdateReceivablebillStatusResponse) SetReqMsgId(v string) *UpdateReceivablebillStatusResponse {
	s.ReqMsgId = &v
	return s
}

func (s *UpdateReceivablebillStatusResponse) SetResultCode(v string) *UpdateReceivablebillStatusResponse {
	s.ResultCode = &v
	return s
}

func (s *UpdateReceivablebillStatusResponse) SetResultMsg(v string) *UpdateReceivablebillStatusResponse {
	s.ResultMsg = &v
	return s
}

func (s *UpdateReceivablebillStatusResponse) SetTxCode(v string) *UpdateReceivablebillStatusResponse {
	s.TxCode = &v
	return s
}

type UpdateBillReceivablebillRequest struct {
	// OAuth模式下的授权token
	AuthToken         *string `json:"auth_token,omitempty" xml:"auth_token,omitempty"`
	ProductInstanceId *string `json:"product_instance_id,omitempty" xml:"product_instance_id,omitempty"`
	// 账单总金额，单位（元），周期内应向发货方收取的运费总金额，不超过2位小数的数字
	BillAmount *string `json:"bill_amount,omitempty" xml:"bill_amount,omitempty"`
	// 账单生成时间，13位毫秒级时间戳
	BillCreateTime *int64 `json:"bill_create_time,omitempty" xml:"bill_create_time,omitempty"`
	// 账单期限，单位（天），合同约定的结算周期
	BillDeadline *int64 `json:"bill_deadline,omitempty" xml:"bill_deadline,omitempty"`
	// 账单id，客户生成的账单唯一编号
	BillId *string `json:"bill_id,omitempty" xml:"bill_id,omitempty" require:"true"`
	// 收款方did，账单的收款方数字身份
	BillPayeeDid *string `json:"bill_payee_did,omitempty" xml:"bill_payee_did,omitempty" require:"true"`
	// 付款方did，账单的付款方数字身份
	BillPayerDid *string `json:"bill_payer_did,omitempty" xml:"bill_payer_did,omitempty"`
	// 账单起始日期，13位毫秒级时间戳
	BillStartTime *int64 `json:"bill_start_time,omitempty" xml:"bill_start_time,omitempty"`
	// 关联合同编号，账单关联的合同编号，如为合同物流请填写
	ContractCode *string `json:"contract_code,omitempty" xml:"contract_code,omitempty"`
	// 账单到期日期，13位毫秒级时间戳
	Deadline *int64 `json:"deadline,omitempty" xml:"deadline,omitempty"`
	// 账单关联运单号数组，元素个数不能超过1000个
	WaybillIds []*string `json:"waybill_ids,omitempty" xml:"waybill_ids,omitempty" type:"Repeated"`
}

func (s UpdateBillReceivablebillRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateBillReceivablebillRequest) GoString() string {
	return s.String()
}

func (s *UpdateBillReceivablebillRequest) SetAuthToken(v string) *UpdateBillReceivablebillRequest {
	s.AuthToken = &v
	return s
}

func (s *UpdateBillReceivablebillRequest) SetProductInstanceId(v string) *UpdateBillReceivablebillRequest {
	s.ProductInstanceId = &v
	return s
}

func (s *UpdateBillReceivablebillRequest) SetBillAmount(v string) *UpdateBillReceivablebillRequest {
	s.BillAmount = &v
	return s
}

func (s *UpdateBillReceivablebillRequest) SetBillCreateTime(v int64) *UpdateBillReceivablebillRequest {
	s.BillCreateTime = &v
	return s
}

func (s *UpdateBillReceivablebillRequest) SetBillDeadline(v int64) *UpdateBillReceivablebillRequest {
	s.BillDeadline = &v
	return s
}

func (s *UpdateBillReceivablebillRequest) SetBillId(v string) *UpdateBillReceivablebillRequest {
	s.BillId = &v
	return s
}

func (s *UpdateBillReceivablebillRequest) SetBillPayeeDid(v string) *UpdateBillReceivablebillRequest {
	s.BillPayeeDid = &v
	return s
}

func (s *UpdateBillReceivablebillRequest) SetBillPayerDid(v string) *UpdateBillReceivablebillRequest {
	s.BillPayerDid = &v
	return s
}

func (s *UpdateBillReceivablebillRequest) SetBillStartTime(v int64) *UpdateBillReceivablebillRequest {
	s.BillStartTime = &v
	return s
}

func (s *UpdateBillReceivablebillRequest) SetContractCode(v string) *UpdateBillReceivablebillRequest {
	s.ContractCode = &v
	return s
}

func (s *UpdateBillReceivablebillRequest) SetDeadline(v int64) *UpdateBillReceivablebillRequest {
	s.Deadline = &v
	return s
}

func (s *UpdateBillReceivablebillRequest) SetWaybillIds(v []*string) *UpdateBillReceivablebillRequest {
	s.WaybillIds = v
	return s
}

type UpdateBillReceivablebillResponse struct {
	// 请求唯一ID，用于链路跟踪和问题排查
	ReqMsgId *string `json:"req_msg_id,omitempty" xml:"req_msg_id,omitempty"`
	// 结果码，一般OK表示调用成功
	ResultCode *string `json:"result_code,omitempty" xml:"result_code,omitempty"`
	// 异常信息的文本描述
	ResultMsg *string `json:"result_msg,omitempty" xml:"result_msg,omitempty"`
	// 链上凭证
	TxCode *string `json:"tx_code,omitempty" xml:"tx_code,omitempty"`
}

func (s UpdateBillReceivablebillResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateBillReceivablebillResponse) GoString() string {
	return s.String()
}

func (s *UpdateBillReceivablebillResponse) SetReqMsgId(v string) *UpdateBillReceivablebillResponse {
	s.ReqMsgId = &v
	return s
}

func (s *UpdateBillReceivablebillResponse) SetResultCode(v string) *UpdateBillReceivablebillResponse {
	s.ResultCode = &v
	return s
}

func (s *UpdateBillReceivablebillResponse) SetResultMsg(v string) *UpdateBillReceivablebillResponse {
	s.ResultMsg = &v
	return s
}

func (s *UpdateBillReceivablebillResponse) SetTxCode(v string) *UpdateBillReceivablebillResponse {
	s.TxCode = &v
	return s
}

type CreateHighwayInvoiceRequest struct {
	// OAuth模式下的授权token
	AuthToken         *string `json:"auth_token,omitempty" xml:"auth_token,omitempty"`
	ProductInstanceId *string `json:"product_instance_id,omitempty" xml:"product_instance_id,omitempty"`
	// 发票号码，8位发票号码
	InvoiceNo *string `json:"invoice_no,omitempty" xml:"invoice_no,omitempty" require:"true"`
	// 发票代码，10位或者12位发票代码
	InvoiceCode *string `json:"invoice_code,omitempty" xml:"invoice_code,omitempty" require:"true"`
	// 发票校验码后6位
	VerifyCode *string `json:"verify_code,omitempty" xml:"verify_code,omitempty"`
	// 收款方did，出票方，需要出票收钱的
	PayeeDid *string `json:"payee_did,omitempty" xml:"payee_did,omitempty" require:"true"`
	// 付款方did，收票方，需要收票付钱的
	PayerDid *string `json:"payer_did,omitempty" xml:"payer_did,omitempty" require:"true"`
	// 发票含税金额， 不超过2位小数的数字
	InvoiceContainsTax *string `json:"invoice_contains_tax,omitempty" xml:"invoice_contains_tax,omitempty" require:"true"`
	// 发票不含税金额，不超过2位小数的数字
	InvoiceWithoutTax *string `json:"invoice_without_tax,omitempty" xml:"invoice_without_tax,omitempty" require:"true"`
	// 开票日期，格式为yyyy-mm-dd
	InvoiceDate *string `json:"invoice_date,omitempty" xml:"invoice_date,omitempty" require:"true"`
	// 发票类型，以下二选一填写：应收发票、应付发票
	InvoiceType *string `json:"invoice_type,omitempty" xml:"invoice_type,omitempty" require:"true"`
	// 影像件ID，发票上传蚂蚁oss的文件ID。注意：影像件id和影像件hash，必须都填写或都不填，不可只填其中一项
	FileId *string `json:"file_id,omitempty" xml:"file_id,omitempty"`
	// 影像件hash，上传蚂蚁oss的文件hash。注意：影像件id和影像件hash，必须都填写或都不填，不可只填其中一项
	FileHash *string `json:"file_hash,omitempty" xml:"file_hash,omitempty"`
	// 发票归属平台did，用以区分是哪家平台/企业的发票
	PlatformDid *string `json:"platform_did,omitempty" xml:"platform_did,omitempty" require:"true"`
	// 发票种类，以下四选一填写（可填VAT_SPECIAL_INVOICE、VAT_NORMAL_INVOICE、VAT_ROLL_INVOICE、ELECTRONIC_INVOICE）。注：以上分别表示增值税专票、增值税普票、增值税卷票、电子发票
	Type *string `json:"type,omitempty" xml:"type,omitempty" require:"true"`
}

func (s CreateHighwayInvoiceRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateHighwayInvoiceRequest) GoString() string {
	return s.String()
}

func (s *CreateHighwayInvoiceRequest) SetAuthToken(v string) *CreateHighwayInvoiceRequest {
	s.AuthToken = &v
	return s
}

func (s *CreateHighwayInvoiceRequest) SetProductInstanceId(v string) *CreateHighwayInvoiceRequest {
	s.ProductInstanceId = &v
	return s
}

func (s *CreateHighwayInvoiceRequest) SetInvoiceNo(v string) *CreateHighwayInvoiceRequest {
	s.InvoiceNo = &v
	return s
}

func (s *CreateHighwayInvoiceRequest) SetInvoiceCode(v string) *CreateHighwayInvoiceRequest {
	s.InvoiceCode = &v
	return s
}

func (s *CreateHighwayInvoiceRequest) SetVerifyCode(v string) *CreateHighwayInvoiceRequest {
	s.VerifyCode = &v
	return s
}

func (s *CreateHighwayInvoiceRequest) SetPayeeDid(v string) *CreateHighwayInvoiceRequest {
	s.PayeeDid = &v
	return s
}

func (s *CreateHighwayInvoiceRequest) SetPayerDid(v string) *CreateHighwayInvoiceRequest {
	s.PayerDid = &v
	return s
}

func (s *CreateHighwayInvoiceRequest) SetInvoiceContainsTax(v string) *CreateHighwayInvoiceRequest {
	s.InvoiceContainsTax = &v
	return s
}

func (s *CreateHighwayInvoiceRequest) SetInvoiceWithoutTax(v string) *CreateHighwayInvoiceRequest {
	s.InvoiceWithoutTax = &v
	return s
}

func (s *CreateHighwayInvoiceRequest) SetInvoiceDate(v string) *CreateHighwayInvoiceRequest {
	s.InvoiceDate = &v
	return s
}

func (s *CreateHighwayInvoiceRequest) SetInvoiceType(v string) *CreateHighwayInvoiceRequest {
	s.InvoiceType = &v
	return s
}

func (s *CreateHighwayInvoiceRequest) SetFileId(v string) *CreateHighwayInvoiceRequest {
	s.FileId = &v
	return s
}

func (s *CreateHighwayInvoiceRequest) SetFileHash(v string) *CreateHighwayInvoiceRequest {
	s.FileHash = &v
	return s
}

func (s *CreateHighwayInvoiceRequest) SetPlatformDid(v string) *CreateHighwayInvoiceRequest {
	s.PlatformDid = &v
	return s
}

func (s *CreateHighwayInvoiceRequest) SetType(v string) *CreateHighwayInvoiceRequest {
	s.Type = &v
	return s
}

type CreateHighwayInvoiceResponse struct {
	// 请求唯一ID，用于链路跟踪和问题排查
	ReqMsgId *string `json:"req_msg_id,omitempty" xml:"req_msg_id,omitempty"`
	// 结果码，一般OK表示调用成功
	ResultCode *string `json:"result_code,omitempty" xml:"result_code,omitempty"`
	// 异常信息的文本描述
	ResultMsg *string `json:"result_msg,omitempty" xml:"result_msg,omitempty"`
	// 链上凭证
	TxCode *string `json:"tx_code,omitempty" xml:"tx_code,omitempty"`
}

func (s CreateHighwayInvoiceResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateHighwayInvoiceResponse) GoString() string {
	return s.String()
}

func (s *CreateHighwayInvoiceResponse) SetReqMsgId(v string) *CreateHighwayInvoiceResponse {
	s.ReqMsgId = &v
	return s
}

func (s *CreateHighwayInvoiceResponse) SetResultCode(v string) *CreateHighwayInvoiceResponse {
	s.ResultCode = &v
	return s
}

func (s *CreateHighwayInvoiceResponse) SetResultMsg(v string) *CreateHighwayInvoiceResponse {
	s.ResultMsg = &v
	return s
}

func (s *CreateHighwayInvoiceResponse) SetTxCode(v string) *CreateHighwayInvoiceResponse {
	s.TxCode = &v
	return s
}

type QueryWaybillInfoRequest struct {
	// OAuth模式下的授权token
	AuthToken         *string `json:"auth_token,omitempty" xml:"auth_token,omitempty"`
	ProductInstanceId *string `json:"product_instance_id,omitempty" xml:"product_instance_id,omitempty"`
	// 运单id
	TaxWaybillId *string `json:"tax_waybill_id,omitempty" xml:"tax_waybill_id,omitempty" require:"true"`
	// 填写无车承运平台did或者3plDid
	Did *string `json:"did,omitempty" xml:"did,omitempty" require:"true"`
}

func (s QueryWaybillInfoRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryWaybillInfoRequest) GoString() string {
	return s.String()
}

func (s *QueryWaybillInfoRequest) SetAuthToken(v string) *QueryWaybillInfoRequest {
	s.AuthToken = &v
	return s
}

func (s *QueryWaybillInfoRequest) SetProductInstanceId(v string) *QueryWaybillInfoRequest {
	s.ProductInstanceId = &v
	return s
}

func (s *QueryWaybillInfoRequest) SetTaxWaybillId(v string) *QueryWaybillInfoRequest {
	s.TaxWaybillId = &v
	return s
}

func (s *QueryWaybillInfoRequest) SetDid(v string) *QueryWaybillInfoRequest {
	s.Did = &v
	return s
}

type QueryWaybillInfoResponse struct {
	// 请求唯一ID，用于链路跟踪和问题排查
	ReqMsgId *string `json:"req_msg_id,omitempty" xml:"req_msg_id,omitempty"`
	// 结果码，一般OK表示调用成功
	ResultCode *string `json:"result_code,omitempty" xml:"result_code,omitempty"`
	// 异常信息的文本描述
	ResultMsg *string `json:"result_msg,omitempty" xml:"result_msg,omitempty"`
	// json串形式的运单信息
	Waybill *string `json:"waybill,omitempty" xml:"waybill,omitempty"`
}

func (s QueryWaybillInfoResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryWaybillInfoResponse) GoString() string {
	return s.String()
}

func (s *QueryWaybillInfoResponse) SetReqMsgId(v string) *QueryWaybillInfoResponse {
	s.ReqMsgId = &v
	return s
}

func (s *QueryWaybillInfoResponse) SetResultCode(v string) *QueryWaybillInfoResponse {
	s.ResultCode = &v
	return s
}

func (s *QueryWaybillInfoResponse) SetResultMsg(v string) *QueryWaybillInfoResponse {
	s.ResultMsg = &v
	return s
}

func (s *QueryWaybillInfoResponse) SetWaybill(v string) *QueryWaybillInfoResponse {
	s.Waybill = &v
	return s
}

type QueryBusinessInstancestatusRequest struct {
	// OAuth模式下的授权token
	AuthToken         *string `json:"auth_token,omitempty" xml:"auth_token,omitempty"`
	ProductInstanceId *string `json:"product_instance_id,omitempty" xml:"product_instance_id,omitempty"`
	// 实例id
	InstanceId *string `json:"instance_id,omitempty" xml:"instance_id,omitempty" require:"true"`
	// 实例对应的租户id
	InstanceTenant *string `json:"instance_tenant,omitempty" xml:"instance_tenant,omitempty" require:"true"`
}

func (s QueryBusinessInstancestatusRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryBusinessInstancestatusRequest) GoString() string {
	return s.String()
}

func (s *QueryBusinessInstancestatusRequest) SetAuthToken(v string) *QueryBusinessInstancestatusRequest {
	s.AuthToken = &v
	return s
}

func (s *QueryBusinessInstancestatusRequest) SetProductInstanceId(v string) *QueryBusinessInstancestatusRequest {
	s.ProductInstanceId = &v
	return s
}

func (s *QueryBusinessInstancestatusRequest) SetInstanceId(v string) *QueryBusinessInstancestatusRequest {
	s.InstanceId = &v
	return s
}

func (s *QueryBusinessInstancestatusRequest) SetInstanceTenant(v string) *QueryBusinessInstancestatusRequest {
	s.InstanceTenant = &v
	return s
}

type QueryBusinessInstancestatusResponse struct {
	// 请求唯一ID，用于链路跟踪和问题排查
	ReqMsgId *string `json:"req_msg_id,omitempty" xml:"req_msg_id,omitempty"`
	// 结果码，一般OK表示调用成功
	ResultCode *string `json:"result_code,omitempty" xml:"result_code,omitempty"`
	// 异常信息的文本描述
	ResultMsg *string `json:"result_msg,omitempty" xml:"result_msg,omitempty"`
	// 实例id
	InstanceId *string `json:"instance_id,omitempty" xml:"instance_id,omitempty"`
	// 实例对应的租户id
	InstanceTenant *string `json:"instance_tenant,omitempty" xml:"instance_tenant,omitempty"`
	// STARTED---运行中
	// STOPPED--已停服
	// RELEASED--已释放
	InstanceStatus *string `json:"instance_status,omitempty" xml:"instance_status,omitempty"`
}

func (s QueryBusinessInstancestatusResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryBusinessInstancestatusResponse) GoString() string {
	return s.String()
}

func (s *QueryBusinessInstancestatusResponse) SetReqMsgId(v string) *QueryBusinessInstancestatusResponse {
	s.ReqMsgId = &v
	return s
}

func (s *QueryBusinessInstancestatusResponse) SetResultCode(v string) *QueryBusinessInstancestatusResponse {
	s.ResultCode = &v
	return s
}

func (s *QueryBusinessInstancestatusResponse) SetResultMsg(v string) *QueryBusinessInstancestatusResponse {
	s.ResultMsg = &v
	return s
}

func (s *QueryBusinessInstancestatusResponse) SetInstanceId(v string) *QueryBusinessInstancestatusResponse {
	s.InstanceId = &v
	return s
}

func (s *QueryBusinessInstancestatusResponse) SetInstanceTenant(v string) *QueryBusinessInstancestatusResponse {
	s.InstanceTenant = &v
	return s
}

func (s *QueryBusinessInstancestatusResponse) SetInstanceStatus(v string) *QueryBusinessInstancestatusResponse {
	s.InstanceStatus = &v
	return s
}

type OpenCreditDriverRequest struct {
	// OAuth模式下的授权token
	AuthToken         *string `json:"auth_token,omitempty" xml:"auth_token,omitempty"`
	ProductInstanceId *string `json:"product_instance_id,omitempty" xml:"product_instance_id,omitempty"`
	// 司机云资金商户ID
	AccountId *string `json:"account_id,omitempty" xml:"account_id,omitempty" require:"true"`
	// 开通结果回调url
	CallbackUrl *string `json:"callback_url,omitempty" xml:"callback_url,omitempty" require:"true"`
	// 司机分布式数字身份
	DriverDid *string `json:"driver_did,omitempty" xml:"driver_did,omitempty" require:"true"`
	// 集团平台方分布式数字身份
	GroupPlatformDid *string `json:"group_platform_did,omitempty" xml:"group_platform_did,omitempty" require:"true"`
	// 产品id 目前填写 PRODUCT_MYBANK
	ProductId *string `json:"product_id,omitempty" xml:"product_id,omitempty" require:"true"`
}

func (s OpenCreditDriverRequest) String() string {
	return tea.Prettify(s)
}

func (s OpenCreditDriverRequest) GoString() string {
	return s.String()
}

func (s *OpenCreditDriverRequest) SetAuthToken(v string) *OpenCreditDriverRequest {
	s.AuthToken = &v
	return s
}

func (s *OpenCreditDriverRequest) SetProductInstanceId(v string) *OpenCreditDriverRequest {
	s.ProductInstanceId = &v
	return s
}

func (s *OpenCreditDriverRequest) SetAccountId(v string) *OpenCreditDriverRequest {
	s.AccountId = &v
	return s
}

func (s *OpenCreditDriverRequest) SetCallbackUrl(v string) *OpenCreditDriverRequest {
	s.CallbackUrl = &v
	return s
}

func (s *OpenCreditDriverRequest) SetDriverDid(v string) *OpenCreditDriverRequest {
	s.DriverDid = &v
	return s
}

func (s *OpenCreditDriverRequest) SetGroupPlatformDid(v string) *OpenCreditDriverRequest {
	s.GroupPlatformDid = &v
	return s
}

func (s *OpenCreditDriverRequest) SetProductId(v string) *OpenCreditDriverRequest {
	s.ProductId = &v
	return s
}

type OpenCreditDriverResponse struct {
	// 请求唯一ID，用于链路跟踪和问题排查
	ReqMsgId *string `json:"req_msg_id,omitempty" xml:"req_msg_id,omitempty"`
	// 结果码，一般OK表示调用成功
	ResultCode *string `json:"result_code,omitempty" xml:"result_code,omitempty"`
	// 异常信息的文本描述
	ResultMsg *string `json:"result_msg,omitempty" xml:"result_msg,omitempty"`
	// 签署开通申请id
	ApplyId *string `json:"apply_id,omitempty" xml:"apply_id,omitempty"`
	// 协议签署开通url
	SignUrl *string `json:"sign_url,omitempty" xml:"sign_url,omitempty"`
	// 签署状态
	// -1:签署开通失败, 0:未签署开通, 1:已签署开通
	Status *int64 `json:"status,omitempty" xml:"status,omitempty"`
}

func (s OpenCreditDriverResponse) String() string {
	return tea.Prettify(s)
}

func (s OpenCreditDriverResponse) GoString() string {
	return s.String()
}

func (s *OpenCreditDriverResponse) SetReqMsgId(v string) *OpenCreditDriverResponse {
	s.ReqMsgId = &v
	return s
}

func (s *OpenCreditDriverResponse) SetResultCode(v string) *OpenCreditDriverResponse {
	s.ResultCode = &v
	return s
}

func (s *OpenCreditDriverResponse) SetResultMsg(v string) *OpenCreditDriverResponse {
	s.ResultMsg = &v
	return s
}

func (s *OpenCreditDriverResponse) SetApplyId(v string) *OpenCreditDriverResponse {
	s.ApplyId = &v
	return s
}

func (s *OpenCreditDriverResponse) SetSignUrl(v string) *OpenCreditDriverResponse {
	s.SignUrl = &v
	return s
}

func (s *OpenCreditDriverResponse) SetStatus(v int64) *OpenCreditDriverResponse {
	s.Status = &v
	return s
}

type QueryCreditDriverRequest struct {
	// OAuth模式下的授权token
	AuthToken         *string `json:"auth_token,omitempty" xml:"auth_token,omitempty"`
	ProductInstanceId *string `json:"product_instance_id,omitempty" xml:"product_instance_id,omitempty"`
	// 申请id
	ApplyId *string `json:"apply_id,omitempty" xml:"apply_id,omitempty"`
	// 司机分布式数字身份
	Did *string `json:"did,omitempty" xml:"did,omitempty" require:"true"`
	// 集团平台方分布式数字身份
	GroupPlatformDid *string `json:"group_platform_did,omitempty" xml:"group_platform_did,omitempty" require:"true"`
	// 产品id  目前填写PRODUCT_MYBANK
	ProductId *string `json:"product_id,omitempty" xml:"product_id,omitempty" require:"true"`
}

func (s QueryCreditDriverRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryCreditDriverRequest) GoString() string {
	return s.String()
}

func (s *QueryCreditDriverRequest) SetAuthToken(v string) *QueryCreditDriverRequest {
	s.AuthToken = &v
	return s
}

func (s *QueryCreditDriverRequest) SetProductInstanceId(v string) *QueryCreditDriverRequest {
	s.ProductInstanceId = &v
	return s
}

func (s *QueryCreditDriverRequest) SetApplyId(v string) *QueryCreditDriverRequest {
	s.ApplyId = &v
	return s
}

func (s *QueryCreditDriverRequest) SetDid(v string) *QueryCreditDriverRequest {
	s.Did = &v
	return s
}

func (s *QueryCreditDriverRequest) SetGroupPlatformDid(v string) *QueryCreditDriverRequest {
	s.GroupPlatformDid = &v
	return s
}

func (s *QueryCreditDriverRequest) SetProductId(v string) *QueryCreditDriverRequest {
	s.ProductId = &v
	return s
}

type QueryCreditDriverResponse struct {
	// 请求唯一ID，用于链路跟踪和问题排查
	ReqMsgId *string `json:"req_msg_id,omitempty" xml:"req_msg_id,omitempty"`
	// 结果码，一般OK表示调用成功
	ResultCode *string `json:"result_code,omitempty" xml:"result_code,omitempty"`
	// 异常信息的文本描述
	ResultMsg *string `json:"result_msg,omitempty" xml:"result_msg,omitempty"`
	// 开通失败信息
	QueryMsg *string `json:"query_msg,omitempty" xml:"query_msg,omitempty"`
	// 开通状态
	// -1:失败状态， 0:未完成状态， 1:已完成状态
	Status *int64 `json:"status,omitempty" xml:"status,omitempty"`
}

func (s QueryCreditDriverResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryCreditDriverResponse) GoString() string {
	return s.String()
}

func (s *QueryCreditDriverResponse) SetReqMsgId(v string) *QueryCreditDriverResponse {
	s.ReqMsgId = &v
	return s
}

func (s *QueryCreditDriverResponse) SetResultCode(v string) *QueryCreditDriverResponse {
	s.ResultCode = &v
	return s
}

func (s *QueryCreditDriverResponse) SetResultMsg(v string) *QueryCreditDriverResponse {
	s.ResultMsg = &v
	return s
}

func (s *QueryCreditDriverResponse) SetQueryMsg(v string) *QueryCreditDriverResponse {
	s.QueryMsg = &v
	return s
}

func (s *QueryCreditDriverResponse) SetStatus(v int64) *QueryCreditDriverResponse {
	s.Status = &v
	return s
}

type QueryCreditConsignorRequest struct {
	// OAuth模式下的授权token
	AuthToken         *string `json:"auth_token,omitempty" xml:"auth_token,omitempty"`
	ProductInstanceId *string `json:"product_instance_id,omitempty" xml:"product_instance_id,omitempty"`
	// 申请id
	ApplyId *string `json:"apply_id,omitempty" xml:"apply_id,omitempty"`
	// 货主分布式数字身份
	Did *string `json:"did,omitempty" xml:"did,omitempty" require:"true"`
	// 集团平台方分布式数字身份
	GroupPlatformDid *string `json:"group_platform_did,omitempty" xml:"group_platform_did,omitempty" require:"true"`
	// 产品id
	// A模式： PRODUCT_MYBANK,
	// A+模式： PRODUCT_MYBANK_A_PLUS,
	// B模式： PRODUCT_MYBANK_B,
	ProductId *string `json:"product_id,omitempty" xml:"product_id,omitempty" require:"true"`
}

func (s QueryCreditConsignorRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryCreditConsignorRequest) GoString() string {
	return s.String()
}

func (s *QueryCreditConsignorRequest) SetAuthToken(v string) *QueryCreditConsignorRequest {
	s.AuthToken = &v
	return s
}

func (s *QueryCreditConsignorRequest) SetProductInstanceId(v string) *QueryCreditConsignorRequest {
	s.ProductInstanceId = &v
	return s
}

func (s *QueryCreditConsignorRequest) SetApplyId(v string) *QueryCreditConsignorRequest {
	s.ApplyId = &v
	return s
}

func (s *QueryCreditConsignorRequest) SetDid(v string) *QueryCreditConsignorRequest {
	s.Did = &v
	return s
}

func (s *QueryCreditConsignorRequest) SetGroupPlatformDid(v string) *QueryCreditConsignorRequest {
	s.GroupPlatformDid = &v
	return s
}

func (s *QueryCreditConsignorRequest) SetProductId(v string) *QueryCreditConsignorRequest {
	s.ProductId = &v
	return s
}

type QueryCreditConsignorResponse struct {
	// 请求唯一ID，用于链路跟踪和问题排查
	ReqMsgId *string `json:"req_msg_id,omitempty" xml:"req_msg_id,omitempty"`
	// 结果码，一般OK表示调用成功
	ResultCode *string `json:"result_code,omitempty" xml:"result_code,omitempty"`
	// 异常信息的文本描述
	ResultMsg *string `json:"result_msg,omitempty" xml:"result_msg,omitempty"`
	// 开通失败信息
	QueryMsg *string `json:"query_msg,omitempty" xml:"query_msg,omitempty"`
	// 开通状态
	// -1:失败状态， 0:未完成状态， 1:已完成状态
	Status *int64 `json:"status,omitempty" xml:"status,omitempty"`
}

func (s QueryCreditConsignorResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryCreditConsignorResponse) GoString() string {
	return s.String()
}

func (s *QueryCreditConsignorResponse) SetReqMsgId(v string) *QueryCreditConsignorResponse {
	s.ReqMsgId = &v
	return s
}

func (s *QueryCreditConsignorResponse) SetResultCode(v string) *QueryCreditConsignorResponse {
	s.ResultCode = &v
	return s
}

func (s *QueryCreditConsignorResponse) SetResultMsg(v string) *QueryCreditConsignorResponse {
	s.ResultMsg = &v
	return s
}

func (s *QueryCreditConsignorResponse) SetQueryMsg(v string) *QueryCreditConsignorResponse {
	s.QueryMsg = &v
	return s
}

func (s *QueryCreditConsignorResponse) SetStatus(v int64) *QueryCreditConsignorResponse {
	s.Status = &v
	return s
}

type QueryCreditBalanceRequest struct {
	// OAuth模式下的授权token
	AuthToken         *string `json:"auth_token,omitempty" xml:"auth_token,omitempty"`
	ProductInstanceId *string `json:"product_instance_id,omitempty" xml:"product_instance_id,omitempty"`
	// 货主分布式数字身份
	ConsignorDid *string `json:"consignor_did,omitempty" xml:"consignor_did,omitempty" require:"true"`
	// 支付单运费，运费最多精确到小数点后2位
	Freight *string `json:"freight,omitempty" xml:"freight,omitempty" require:"true"`
	// 集团平台方分布式数字身份
	GroupPlatformDid *string `json:"group_platform_did,omitempty" xml:"group_platform_did,omitempty" require:"true"`
	// 业务发起方分布式数字身份
	PlatformDid *string `json:"platform_did,omitempty" xml:"platform_did,omitempty" require:"true"`
	// 产品id
	// A模式：PRODUCT_MYBANK,
	// A+模式：PRODUCT_MYBANK_A_PLUS,
	// B模式：PRODUCT_MYBANK_B,
	//
	ProductId *string `json:"product_id,omitempty" xml:"product_id,omitempty" require:"true"`
}

func (s QueryCreditBalanceRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryCreditBalanceRequest) GoString() string {
	return s.String()
}

func (s *QueryCreditBalanceRequest) SetAuthToken(v string) *QueryCreditBalanceRequest {
	s.AuthToken = &v
	return s
}

func (s *QueryCreditBalanceRequest) SetProductInstanceId(v string) *QueryCreditBalanceRequest {
	s.ProductInstanceId = &v
	return s
}

func (s *QueryCreditBalanceRequest) SetConsignorDid(v string) *QueryCreditBalanceRequest {
	s.ConsignorDid = &v
	return s
}

func (s *QueryCreditBalanceRequest) SetFreight(v string) *QueryCreditBalanceRequest {
	s.Freight = &v
	return s
}

func (s *QueryCreditBalanceRequest) SetGroupPlatformDid(v string) *QueryCreditBalanceRequest {
	s.GroupPlatformDid = &v
	return s
}

func (s *QueryCreditBalanceRequest) SetPlatformDid(v string) *QueryCreditBalanceRequest {
	s.PlatformDid = &v
	return s
}

func (s *QueryCreditBalanceRequest) SetProductId(v string) *QueryCreditBalanceRequest {
	s.ProductId = &v
	return s
}

type QueryCreditBalanceResponse struct {
	// 请求唯一ID，用于链路跟踪和问题排查
	ReqMsgId *string `json:"req_msg_id,omitempty" xml:"req_msg_id,omitempty"`
	// 结果码，一般OK表示调用成功
	ResultCode *string `json:"result_code,omitempty" xml:"result_code,omitempty"`
	// 异常信息的文本描述
	ResultMsg *string `json:"result_msg,omitempty" xml:"result_msg,omitempty"`
	// 是否有信用流转额度
	HasBalance *bool `json:"has_balance,omitempty" xml:"has_balance,omitempty"`
}

func (s QueryCreditBalanceResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryCreditBalanceResponse) GoString() string {
	return s.String()
}

func (s *QueryCreditBalanceResponse) SetReqMsgId(v string) *QueryCreditBalanceResponse {
	s.ReqMsgId = &v
	return s
}

func (s *QueryCreditBalanceResponse) SetResultCode(v string) *QueryCreditBalanceResponse {
	s.ResultCode = &v
	return s
}

func (s *QueryCreditBalanceResponse) SetResultMsg(v string) *QueryCreditBalanceResponse {
	s.ResultMsg = &v
	return s
}

func (s *QueryCreditBalanceResponse) SetHasBalance(v bool) *QueryCreditBalanceResponse {
	s.HasBalance = &v
	return s
}

type UploadCreditIssueRequest struct {
	// OAuth模式下的授权token
	AuthToken         *string `json:"auth_token,omitempty" xml:"auth_token,omitempty"`
	ProductInstanceId *string `json:"product_instance_id,omitempty" xml:"product_instance_id,omitempty"`
	// 批次号
	BatchId *string `json:"batch_id,omitempty" xml:"batch_id,omitempty" require:"true"`
	// 发行结果回调url
	CallbackUrl *string `json:"callback_url,omitempty" xml:"callback_url,omitempty" require:"true"`
	// 货主分布式数字身份
	ConsignorDid *string `json:"consignor_did,omitempty" xml:"consignor_did,omitempty" require:"true"`
	// 集团平台方分布式数字身份
	GroupPlatformDid *string `json:"group_platform_did,omitempty" xml:"group_platform_did,omitempty" require:"true"`
	// 发行信用流转信息列表
	IssueApplyInfos []*IssueApplyInfo `json:"issue_apply_infos,omitempty" xml:"issue_apply_infos,omitempty" require:"true" type:"Repeated"`
	// 业务发起方分布式数字身份
	PlatformDid *string `json:"platform_did,omitempty" xml:"platform_did,omitempty" require:"true"`
	// 产品id 目前填PRODUCT_MYBANK
	ProductId *string `json:"product_id,omitempty" xml:"product_id,omitempty" require:"true"`
}

func (s UploadCreditIssueRequest) String() string {
	return tea.Prettify(s)
}

func (s UploadCreditIssueRequest) GoString() string {
	return s.String()
}

func (s *UploadCreditIssueRequest) SetAuthToken(v string) *UploadCreditIssueRequest {
	s.AuthToken = &v
	return s
}

func (s *UploadCreditIssueRequest) SetProductInstanceId(v string) *UploadCreditIssueRequest {
	s.ProductInstanceId = &v
	return s
}

func (s *UploadCreditIssueRequest) SetBatchId(v string) *UploadCreditIssueRequest {
	s.BatchId = &v
	return s
}

func (s *UploadCreditIssueRequest) SetCallbackUrl(v string) *UploadCreditIssueRequest {
	s.CallbackUrl = &v
	return s
}

func (s *UploadCreditIssueRequest) SetConsignorDid(v string) *UploadCreditIssueRequest {
	s.ConsignorDid = &v
	return s
}

func (s *UploadCreditIssueRequest) SetGroupPlatformDid(v string) *UploadCreditIssueRequest {
	s.GroupPlatformDid = &v
	return s
}

func (s *UploadCreditIssueRequest) SetIssueApplyInfos(v []*IssueApplyInfo) *UploadCreditIssueRequest {
	s.IssueApplyInfos = v
	return s
}

func (s *UploadCreditIssueRequest) SetPlatformDid(v string) *UploadCreditIssueRequest {
	s.PlatformDid = &v
	return s
}

func (s *UploadCreditIssueRequest) SetProductId(v string) *UploadCreditIssueRequest {
	s.ProductId = &v
	return s
}

type UploadCreditIssueResponse struct {
	// 请求唯一ID，用于链路跟踪和问题排查
	ReqMsgId *string `json:"req_msg_id,omitempty" xml:"req_msg_id,omitempty"`
	// 结果码，一般OK表示调用成功
	ResultCode *string `json:"result_code,omitempty" xml:"result_code,omitempty"`
	// 异常信息的文本描述
	ResultMsg *string `json:"result_msg,omitempty" xml:"result_msg,omitempty"`
	// 批次号
	BatchId *string `json:"batch_id,omitempty" xml:"batch_id,omitempty"`
	// 批次状态
	BatchIdStatus *int64 `json:"batch_id_status,omitempty" xml:"batch_id_status,omitempty"`
	// 发行url
	IssueUrl *string `json:"issue_url,omitempty" xml:"issue_url,omitempty"`
}

func (s UploadCreditIssueResponse) String() string {
	return tea.Prettify(s)
}

func (s UploadCreditIssueResponse) GoString() string {
	return s.String()
}

func (s *UploadCreditIssueResponse) SetReqMsgId(v string) *UploadCreditIssueResponse {
	s.ReqMsgId = &v
	return s
}

func (s *UploadCreditIssueResponse) SetResultCode(v string) *UploadCreditIssueResponse {
	s.ResultCode = &v
	return s
}

func (s *UploadCreditIssueResponse) SetResultMsg(v string) *UploadCreditIssueResponse {
	s.ResultMsg = &v
	return s
}

func (s *UploadCreditIssueResponse) SetBatchId(v string) *UploadCreditIssueResponse {
	s.BatchId = &v
	return s
}

func (s *UploadCreditIssueResponse) SetBatchIdStatus(v int64) *UploadCreditIssueResponse {
	s.BatchIdStatus = &v
	return s
}

func (s *UploadCreditIssueResponse) SetIssueUrl(v string) *UploadCreditIssueResponse {
	s.IssueUrl = &v
	return s
}

type QueryCreditIssuebatchstatusRequest struct {
	// OAuth模式下的授权token
	AuthToken         *string `json:"auth_token,omitempty" xml:"auth_token,omitempty"`
	ProductInstanceId *string `json:"product_instance_id,omitempty" xml:"product_instance_id,omitempty"`
	// 批次号
	BatchId *string `json:"batch_id,omitempty" xml:"batch_id,omitempty" require:"true"`
	// 集团平台方分布式数字身份
	GroupPlatformDid *string `json:"group_platform_did,omitempty" xml:"group_platform_did,omitempty" require:"true"`
	// 业务发起方分布式数字身份
	PlatformDid *string `json:"platform_did,omitempty" xml:"platform_did,omitempty" require:"true"`
	// 产品id
	// A模式：PRODUCT_MYBANK，
	// A+模式：PRODUCT_MYBANK_A_PLUS，
	// A模式：PRODUCT_MYBANK_B，
	ProductId *string `json:"product_id,omitempty" xml:"product_id,omitempty" require:"true"`
}

func (s QueryCreditIssuebatchstatusRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryCreditIssuebatchstatusRequest) GoString() string {
	return s.String()
}

func (s *QueryCreditIssuebatchstatusRequest) SetAuthToken(v string) *QueryCreditIssuebatchstatusRequest {
	s.AuthToken = &v
	return s
}

func (s *QueryCreditIssuebatchstatusRequest) SetProductInstanceId(v string) *QueryCreditIssuebatchstatusRequest {
	s.ProductInstanceId = &v
	return s
}

func (s *QueryCreditIssuebatchstatusRequest) SetBatchId(v string) *QueryCreditIssuebatchstatusRequest {
	s.BatchId = &v
	return s
}

func (s *QueryCreditIssuebatchstatusRequest) SetGroupPlatformDid(v string) *QueryCreditIssuebatchstatusRequest {
	s.GroupPlatformDid = &v
	return s
}

func (s *QueryCreditIssuebatchstatusRequest) SetPlatformDid(v string) *QueryCreditIssuebatchstatusRequest {
	s.PlatformDid = &v
	return s
}

func (s *QueryCreditIssuebatchstatusRequest) SetProductId(v string) *QueryCreditIssuebatchstatusRequest {
	s.ProductId = &v
	return s
}

type QueryCreditIssuebatchstatusResponse struct {
	// 请求唯一ID，用于链路跟踪和问题排查
	ReqMsgId *string `json:"req_msg_id,omitempty" xml:"req_msg_id,omitempty"`
	// 结果码，一般OK表示调用成功
	ResultCode *string `json:"result_code,omitempty" xml:"result_code,omitempty"`
	// 异常信息的文本描述
	ResultMsg *string `json:"result_msg,omitempty" xml:"result_msg,omitempty"`
	// 批次号
	BatchId *string `json:"batch_id,omitempty" xml:"batch_id,omitempty"`
	// 批次状态
	BatchStatus *int64 `json:"batch_status,omitempty" xml:"batch_status,omitempty"`
}

func (s QueryCreditIssuebatchstatusResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryCreditIssuebatchstatusResponse) GoString() string {
	return s.String()
}

func (s *QueryCreditIssuebatchstatusResponse) SetReqMsgId(v string) *QueryCreditIssuebatchstatusResponse {
	s.ReqMsgId = &v
	return s
}

func (s *QueryCreditIssuebatchstatusResponse) SetResultCode(v string) *QueryCreditIssuebatchstatusResponse {
	s.ResultCode = &v
	return s
}

func (s *QueryCreditIssuebatchstatusResponse) SetResultMsg(v string) *QueryCreditIssuebatchstatusResponse {
	s.ResultMsg = &v
	return s
}

func (s *QueryCreditIssuebatchstatusResponse) SetBatchId(v string) *QueryCreditIssuebatchstatusResponse {
	s.BatchId = &v
	return s
}

func (s *QueryCreditIssuebatchstatusResponse) SetBatchStatus(v int64) *QueryCreditIssuebatchstatusResponse {
	s.BatchStatus = &v
	return s
}

type CancelCreditIssuebatchRequest struct {
	// OAuth模式下的授权token
	AuthToken         *string `json:"auth_token,omitempty" xml:"auth_token,omitempty"`
	ProductInstanceId *string `json:"product_instance_id,omitempty" xml:"product_instance_id,omitempty"`
	// 批次号
	BatchId *string `json:"batch_id,omitempty" xml:"batch_id,omitempty" require:"true"`
	// 货主分布式数字身份
	ConsignorDid *string `json:"consignor_did,omitempty" xml:"consignor_did,omitempty" require:"true"`
	// 集团平台方分布式数字身份
	GroupPlatformDid *string `json:"group_platform_did,omitempty" xml:"group_platform_did,omitempty" require:"true"`
	// 业务发起方分布式数字身份
	PlatformDid *string `json:"platform_did,omitempty" xml:"platform_did,omitempty" require:"true"`
	// 产品id
	// A模式：PRODUCT_MYBANK，
	// A+模式：PRODUCT_MYBANK_A_PLUS，
	// B模式：PRODUCT_MYBANK_B
	ProductId *string `json:"product_id,omitempty" xml:"product_id,omitempty" require:"true"`
}

func (s CancelCreditIssuebatchRequest) String() string {
	return tea.Prettify(s)
}

func (s CancelCreditIssuebatchRequest) GoString() string {
	return s.String()
}

func (s *CancelCreditIssuebatchRequest) SetAuthToken(v string) *CancelCreditIssuebatchRequest {
	s.AuthToken = &v
	return s
}

func (s *CancelCreditIssuebatchRequest) SetProductInstanceId(v string) *CancelCreditIssuebatchRequest {
	s.ProductInstanceId = &v
	return s
}

func (s *CancelCreditIssuebatchRequest) SetBatchId(v string) *CancelCreditIssuebatchRequest {
	s.BatchId = &v
	return s
}

func (s *CancelCreditIssuebatchRequest) SetConsignorDid(v string) *CancelCreditIssuebatchRequest {
	s.ConsignorDid = &v
	return s
}

func (s *CancelCreditIssuebatchRequest) SetGroupPlatformDid(v string) *CancelCreditIssuebatchRequest {
	s.GroupPlatformDid = &v
	return s
}

func (s *CancelCreditIssuebatchRequest) SetPlatformDid(v string) *CancelCreditIssuebatchRequest {
	s.PlatformDid = &v
	return s
}

func (s *CancelCreditIssuebatchRequest) SetProductId(v string) *CancelCreditIssuebatchRequest {
	s.ProductId = &v
	return s
}

type CancelCreditIssuebatchResponse struct {
	// 请求唯一ID，用于链路跟踪和问题排查
	ReqMsgId *string `json:"req_msg_id,omitempty" xml:"req_msg_id,omitempty"`
	// 结果码，一般OK表示调用成功
	ResultCode *string `json:"result_code,omitempty" xml:"result_code,omitempty"`
	// 异常信息的文本描述
	ResultMsg *string `json:"result_msg,omitempty" xml:"result_msg,omitempty"`
	// 批次号
	BatchId *string `json:"batch_id,omitempty" xml:"batch_id,omitempty"`
}

func (s CancelCreditIssuebatchResponse) String() string {
	return tea.Prettify(s)
}

func (s CancelCreditIssuebatchResponse) GoString() string {
	return s.String()
}

func (s *CancelCreditIssuebatchResponse) SetReqMsgId(v string) *CancelCreditIssuebatchResponse {
	s.ReqMsgId = &v
	return s
}

func (s *CancelCreditIssuebatchResponse) SetResultCode(v string) *CancelCreditIssuebatchResponse {
	s.ResultCode = &v
	return s
}

func (s *CancelCreditIssuebatchResponse) SetResultMsg(v string) *CancelCreditIssuebatchResponse {
	s.ResultMsg = &v
	return s
}

func (s *CancelCreditIssuebatchResponse) SetBatchId(v string) *CancelCreditIssuebatchResponse {
	s.BatchId = &v
	return s
}

type QueryCreditIssuebyidRequest struct {
	// OAuth模式下的授权token
	AuthToken         *string `json:"auth_token,omitempty" xml:"auth_token,omitempty"`
	ProductInstanceId *string `json:"product_instance_id,omitempty" xml:"product_instance_id,omitempty"`
	// 信用流转发行批次号
	BatchId *string `json:"batch_id,omitempty" xml:"batch_id,omitempty" require:"true"`
	// 集团平台方分布式数字身份
	GroupPlatformDid *string `json:"group_platform_did,omitempty" xml:"group_platform_did,omitempty" require:"true"`
	// 是否只返回已发行凭证信息
	IssuedOnly *bool `json:"issued_only,omitempty" xml:"issued_only,omitempty" require:"true"`
	// 信用流转发行凭证Id
	IssueId *string `json:"issue_id,omitempty" xml:"issue_id,omitempty"`
	// 页数 从1开始
	PageNum *int64 `json:"page_num,omitempty" xml:"page_num,omitempty" require:"true"`
	// 每页显示数量 不超过100
	PageSize *int64 `json:"page_size,omitempty" xml:"page_size,omitempty" require:"true"`
	// 业务发起方分布式数字身份
	PlatformDid *string `json:"platform_did,omitempty" xml:"platform_did,omitempty" require:"true"`
	// 产品id
	// A模式：PRODUCT_MYBANK，
	// A+模式：PRODUCT_MYBANK_A_PLUS，
	// B模式：PRODUCT_MYBANK_B
	ProductId *string `json:"product_id,omitempty" xml:"product_id,omitempty" require:"true"`
}

func (s QueryCreditIssuebyidRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryCreditIssuebyidRequest) GoString() string {
	return s.String()
}

func (s *QueryCreditIssuebyidRequest) SetAuthToken(v string) *QueryCreditIssuebyidRequest {
	s.AuthToken = &v
	return s
}

func (s *QueryCreditIssuebyidRequest) SetProductInstanceId(v string) *QueryCreditIssuebyidRequest {
	s.ProductInstanceId = &v
	return s
}

func (s *QueryCreditIssuebyidRequest) SetBatchId(v string) *QueryCreditIssuebyidRequest {
	s.BatchId = &v
	return s
}

func (s *QueryCreditIssuebyidRequest) SetGroupPlatformDid(v string) *QueryCreditIssuebyidRequest {
	s.GroupPlatformDid = &v
	return s
}

func (s *QueryCreditIssuebyidRequest) SetIssuedOnly(v bool) *QueryCreditIssuebyidRequest {
	s.IssuedOnly = &v
	return s
}

func (s *QueryCreditIssuebyidRequest) SetIssueId(v string) *QueryCreditIssuebyidRequest {
	s.IssueId = &v
	return s
}

func (s *QueryCreditIssuebyidRequest) SetPageNum(v int64) *QueryCreditIssuebyidRequest {
	s.PageNum = &v
	return s
}

func (s *QueryCreditIssuebyidRequest) SetPageSize(v int64) *QueryCreditIssuebyidRequest {
	s.PageSize = &v
	return s
}

func (s *QueryCreditIssuebyidRequest) SetPlatformDid(v string) *QueryCreditIssuebyidRequest {
	s.PlatformDid = &v
	return s
}

func (s *QueryCreditIssuebyidRequest) SetProductId(v string) *QueryCreditIssuebyidRequest {
	s.ProductId = &v
	return s
}

type QueryCreditIssuebyidResponse struct {
	// 请求唯一ID，用于链路跟踪和问题排查
	ReqMsgId *string `json:"req_msg_id,omitempty" xml:"req_msg_id,omitempty"`
	// 结果码，一般OK表示调用成功
	ResultCode *string `json:"result_code,omitempty" xml:"result_code,omitempty"`
	// 异常信息的文本描述
	ResultMsg *string `json:"result_msg,omitempty" xml:"result_msg,omitempty"`
	// 信用流转批次号
	BatchId *string `json:"batch_id,omitempty" xml:"batch_id,omitempty"`
	// 批次状态
	BatchIdStatus *string `json:"batch_id_status,omitempty" xml:"batch_id_status,omitempty"`
	// 凭证发行者did
	Did *string `json:"did,omitempty" xml:"did,omitempty"`
	// 发行凭证列表
	IssueIds []*IssueIdInfo `json:"issue_ids,omitempty" xml:"issue_ids,omitempty" type:"Repeated"`
	// 页数 从1开始
	PageNum *int64 `json:"page_num,omitempty" xml:"page_num,omitempty"`
	// 每页显示数量
	PageSize *int64 `json:"page_size,omitempty" xml:"page_size,omitempty"`
	// 数据总量
	TotalCount *int64 `json:"total_count,omitempty" xml:"total_count,omitempty"`
}

func (s QueryCreditIssuebyidResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryCreditIssuebyidResponse) GoString() string {
	return s.String()
}

func (s *QueryCreditIssuebyidResponse) SetReqMsgId(v string) *QueryCreditIssuebyidResponse {
	s.ReqMsgId = &v
	return s
}

func (s *QueryCreditIssuebyidResponse) SetResultCode(v string) *QueryCreditIssuebyidResponse {
	s.ResultCode = &v
	return s
}

func (s *QueryCreditIssuebyidResponse) SetResultMsg(v string) *QueryCreditIssuebyidResponse {
	s.ResultMsg = &v
	return s
}

func (s *QueryCreditIssuebyidResponse) SetBatchId(v string) *QueryCreditIssuebyidResponse {
	s.BatchId = &v
	return s
}

func (s *QueryCreditIssuebyidResponse) SetBatchIdStatus(v string) *QueryCreditIssuebyidResponse {
	s.BatchIdStatus = &v
	return s
}

func (s *QueryCreditIssuebyidResponse) SetDid(v string) *QueryCreditIssuebyidResponse {
	s.Did = &v
	return s
}

func (s *QueryCreditIssuebyidResponse) SetIssueIds(v []*IssueIdInfo) *QueryCreditIssuebyidResponse {
	s.IssueIds = v
	return s
}

func (s *QueryCreditIssuebyidResponse) SetPageNum(v int64) *QueryCreditIssuebyidResponse {
	s.PageNum = &v
	return s
}

func (s *QueryCreditIssuebyidResponse) SetPageSize(v int64) *QueryCreditIssuebyidResponse {
	s.PageSize = &v
	return s
}

func (s *QueryCreditIssuebyidResponse) SetTotalCount(v int64) *QueryCreditIssuebyidResponse {
	s.TotalCount = &v
	return s
}

type QueryCreditIssuebytimeRequest struct {
	// OAuth模式下的授权token
	AuthToken         *string `json:"auth_token,omitempty" xml:"auth_token,omitempty"`
	ProductInstanceId *string `json:"product_instance_id,omitempty" xml:"product_instance_id,omitempty"`
	// 指定查询者分布式数字身份
	Did *string `json:"did,omitempty" xml:"did,omitempty" require:"true"`
	// 查询截止时间
	EndDate *string `json:"end_date,omitempty" xml:"end_date,omitempty" require:"true"`
	// 集团平台方分布式数字身份
	GroupPlatformDid *string `json:"group_platform_did,omitempty" xml:"group_platform_did,omitempty" require:"true"`
	// 是否只返回已发行凭证信息
	IssuedOnly *bool `json:"issued_only,omitempty" xml:"issued_only,omitempty" require:"true"`
	// 页数 从1开始
	PageNum *int64 `json:"page_num,omitempty" xml:"page_num,omitempty" require:"true"`
	// 每页显示数量 不超过100
	PageSize *int64 `json:"page_size,omitempty" xml:"page_size,omitempty" require:"true"`
	// 业务发起方分布式数字身份
	PlatformDid *string `json:"platform_did,omitempty" xml:"platform_did,omitempty" require:"true"`
	// 产品id 目前填PRODUCT_MYBANK
	ProductId *string `json:"product_id,omitempty" xml:"product_id,omitempty" require:"true"`
	// 查询起始时间
	StartDate *string `json:"start_date,omitempty" xml:"start_date,omitempty" require:"true"`
}

func (s QueryCreditIssuebytimeRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryCreditIssuebytimeRequest) GoString() string {
	return s.String()
}

func (s *QueryCreditIssuebytimeRequest) SetAuthToken(v string) *QueryCreditIssuebytimeRequest {
	s.AuthToken = &v
	return s
}

func (s *QueryCreditIssuebytimeRequest) SetProductInstanceId(v string) *QueryCreditIssuebytimeRequest {
	s.ProductInstanceId = &v
	return s
}

func (s *QueryCreditIssuebytimeRequest) SetDid(v string) *QueryCreditIssuebytimeRequest {
	s.Did = &v
	return s
}

func (s *QueryCreditIssuebytimeRequest) SetEndDate(v string) *QueryCreditIssuebytimeRequest {
	s.EndDate = &v
	return s
}

func (s *QueryCreditIssuebytimeRequest) SetGroupPlatformDid(v string) *QueryCreditIssuebytimeRequest {
	s.GroupPlatformDid = &v
	return s
}

func (s *QueryCreditIssuebytimeRequest) SetIssuedOnly(v bool) *QueryCreditIssuebytimeRequest {
	s.IssuedOnly = &v
	return s
}

func (s *QueryCreditIssuebytimeRequest) SetPageNum(v int64) *QueryCreditIssuebytimeRequest {
	s.PageNum = &v
	return s
}

func (s *QueryCreditIssuebytimeRequest) SetPageSize(v int64) *QueryCreditIssuebytimeRequest {
	s.PageSize = &v
	return s
}

func (s *QueryCreditIssuebytimeRequest) SetPlatformDid(v string) *QueryCreditIssuebytimeRequest {
	s.PlatformDid = &v
	return s
}

func (s *QueryCreditIssuebytimeRequest) SetProductId(v string) *QueryCreditIssuebytimeRequest {
	s.ProductId = &v
	return s
}

func (s *QueryCreditIssuebytimeRequest) SetStartDate(v string) *QueryCreditIssuebytimeRequest {
	s.StartDate = &v
	return s
}

type QueryCreditIssuebytimeResponse struct {
	// 请求唯一ID，用于链路跟踪和问题排查
	ReqMsgId *string `json:"req_msg_id,omitempty" xml:"req_msg_id,omitempty"`
	// 结果码，一般OK表示调用成功
	ResultCode *string `json:"result_code,omitempty" xml:"result_code,omitempty"`
	// 异常信息的文本描述
	ResultMsg *string `json:"result_msg,omitempty" xml:"result_msg,omitempty"`
	// 信用流转批次号
	BatchId *string `json:"batch_id,omitempty" xml:"batch_id,omitempty"`
	// 批次状态
	BatchIdStatus *int64 `json:"batch_id_status,omitempty" xml:"batch_id_status,omitempty"`
	// 凭证发行者did
	Did *string `json:"did,omitempty" xml:"did,omitempty"`
	// 凭证列表
	IssueIds []*IssueIdInfo `json:"issue_ids,omitempty" xml:"issue_ids,omitempty" type:"Repeated"`
	// 页数 从1开始
	PageNum *int64 `json:"page_num,omitempty" xml:"page_num,omitempty"`
	// 每页显示数量
	PageSize *int64 `json:"page_size,omitempty" xml:"page_size,omitempty"`
	// 数据总量
	TotalCount *int64 `json:"total_count,omitempty" xml:"total_count,omitempty"`
}

func (s QueryCreditIssuebytimeResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryCreditIssuebytimeResponse) GoString() string {
	return s.String()
}

func (s *QueryCreditIssuebytimeResponse) SetReqMsgId(v string) *QueryCreditIssuebytimeResponse {
	s.ReqMsgId = &v
	return s
}

func (s *QueryCreditIssuebytimeResponse) SetResultCode(v string) *QueryCreditIssuebytimeResponse {
	s.ResultCode = &v
	return s
}

func (s *QueryCreditIssuebytimeResponse) SetResultMsg(v string) *QueryCreditIssuebytimeResponse {
	s.ResultMsg = &v
	return s
}

func (s *QueryCreditIssuebytimeResponse) SetBatchId(v string) *QueryCreditIssuebytimeResponse {
	s.BatchId = &v
	return s
}

func (s *QueryCreditIssuebytimeResponse) SetBatchIdStatus(v int64) *QueryCreditIssuebytimeResponse {
	s.BatchIdStatus = &v
	return s
}

func (s *QueryCreditIssuebytimeResponse) SetDid(v string) *QueryCreditIssuebytimeResponse {
	s.Did = &v
	return s
}

func (s *QueryCreditIssuebytimeResponse) SetIssueIds(v []*IssueIdInfo) *QueryCreditIssuebytimeResponse {
	s.IssueIds = v
	return s
}

func (s *QueryCreditIssuebytimeResponse) SetPageNum(v int64) *QueryCreditIssuebytimeResponse {
	s.PageNum = &v
	return s
}

func (s *QueryCreditIssuebytimeResponse) SetPageSize(v int64) *QueryCreditIssuebytimeResponse {
	s.PageSize = &v
	return s
}

func (s *QueryCreditIssuebytimeResponse) SetTotalCount(v int64) *QueryCreditIssuebytimeResponse {
	s.TotalCount = &v
	return s
}

type QueryCreditUserissueRequest struct {
	// OAuth模式下的授权token
	AuthToken         *string `json:"auth_token,omitempty" xml:"auth_token,omitempty"`
	ProductInstanceId *string `json:"product_instance_id,omitempty" xml:"product_instance_id,omitempty"`
	// 被查询方分布式数字身份
	Did *string `json:"did,omitempty" xml:"did,omitempty" require:"true"`
	// 集团平台方分布式数字身份
	GroupPlatformDid *string `json:"group_platform_did,omitempty" xml:"group_platform_did,omitempty" require:"true"`
	// 页数
	PageNum *int64 `json:"page_num,omitempty" xml:"page_num,omitempty" require:"true"`
	// 每页显示的最大条数
	PageSize *int64 `json:"page_size,omitempty" xml:"page_size,omitempty" require:"true"`
	// 业务发起方分布式数字身份
	PlatformDid *string `json:"platform_did,omitempty" xml:"platform_did,omitempty" require:"true"`
	// 产品id
	// A模式：PRODUCT_MYBANK， A+模式：PRODUCT_MYBANK_A_PLUS， B模式：PRODUCT_MYBANK_B
	ProductId *string `json:"product_id,omitempty" xml:"product_id,omitempty" require:"true"`
}

func (s QueryCreditUserissueRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryCreditUserissueRequest) GoString() string {
	return s.String()
}

func (s *QueryCreditUserissueRequest) SetAuthToken(v string) *QueryCreditUserissueRequest {
	s.AuthToken = &v
	return s
}

func (s *QueryCreditUserissueRequest) SetProductInstanceId(v string) *QueryCreditUserissueRequest {
	s.ProductInstanceId = &v
	return s
}

func (s *QueryCreditUserissueRequest) SetDid(v string) *QueryCreditUserissueRequest {
	s.Did = &v
	return s
}

func (s *QueryCreditUserissueRequest) SetGroupPlatformDid(v string) *QueryCreditUserissueRequest {
	s.GroupPlatformDid = &v
	return s
}

func (s *QueryCreditUserissueRequest) SetPageNum(v int64) *QueryCreditUserissueRequest {
	s.PageNum = &v
	return s
}

func (s *QueryCreditUserissueRequest) SetPageSize(v int64) *QueryCreditUserissueRequest {
	s.PageSize = &v
	return s
}

func (s *QueryCreditUserissueRequest) SetPlatformDid(v string) *QueryCreditUserissueRequest {
	s.PlatformDid = &v
	return s
}

func (s *QueryCreditUserissueRequest) SetProductId(v string) *QueryCreditUserissueRequest {
	s.ProductId = &v
	return s
}

type QueryCreditUserissueResponse struct {
	// 请求唯一ID，用于链路跟踪和问题排查
	ReqMsgId *string `json:"req_msg_id,omitempty" xml:"req_msg_id,omitempty"`
	// 结果码，一般OK表示调用成功
	ResultCode *string `json:"result_code,omitempty" xml:"result_code,omitempty"`
	// 异常信息的文本描述
	ResultMsg *string `json:"result_msg,omitempty" xml:"result_msg,omitempty"`
	// 凭证所有者分布式数字身份
	Did *string `json:"did,omitempty" xml:"did,omitempty"`
	// 信用流转凭证列表
	IssueIds []*UserIssueId `json:"issue_ids,omitempty" xml:"issue_ids,omitempty" type:"Repeated"`
	// 页数
	PageNum *int64 `json:"page_num,omitempty" xml:"page_num,omitempty"`
	// 每页显示最大条数
	PageSize *int64 `json:"page_size,omitempty" xml:"page_size,omitempty"`
	// 总数
	TotalCount *int64 `json:"total_count,omitempty" xml:"total_count,omitempty"`
}

func (s QueryCreditUserissueResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryCreditUserissueResponse) GoString() string {
	return s.String()
}

func (s *QueryCreditUserissueResponse) SetReqMsgId(v string) *QueryCreditUserissueResponse {
	s.ReqMsgId = &v
	return s
}

func (s *QueryCreditUserissueResponse) SetResultCode(v string) *QueryCreditUserissueResponse {
	s.ResultCode = &v
	return s
}

func (s *QueryCreditUserissueResponse) SetResultMsg(v string) *QueryCreditUserissueResponse {
	s.ResultMsg = &v
	return s
}

func (s *QueryCreditUserissueResponse) SetDid(v string) *QueryCreditUserissueResponse {
	s.Did = &v
	return s
}

func (s *QueryCreditUserissueResponse) SetIssueIds(v []*UserIssueId) *QueryCreditUserissueResponse {
	s.IssueIds = v
	return s
}

func (s *QueryCreditUserissueResponse) SetPageNum(v int64) *QueryCreditUserissueResponse {
	s.PageNum = &v
	return s
}

func (s *QueryCreditUserissueResponse) SetPageSize(v int64) *QueryCreditUserissueResponse {
	s.PageSize = &v
	return s
}

func (s *QueryCreditUserissueResponse) SetTotalCount(v int64) *QueryCreditUserissueResponse {
	s.TotalCount = &v
	return s
}

type QueryCreditStatementRequest struct {
	// OAuth模式下的授权token
	AuthToken         *string `json:"auth_token,omitempty" xml:"auth_token,omitempty"`
	ProductInstanceId *string `json:"product_instance_id,omitempty" xml:"product_instance_id,omitempty"`
	// 需查询对象分布式数字身份不能为空
	Did *string `json:"did,omitempty" xml:"did,omitempty" require:"true"`
	// 查询截止时间
	EndDate *string `json:"end_date,omitempty" xml:"end_date,omitempty" require:"true"`
	// 集团平台方分布式数字身份
	GroupPlatformDid *string `json:"group_platform_did,omitempty" xml:"group_platform_did,omitempty" require:"true"`
	// 页数 从1开始
	PageNum *int64 `json:"page_num,omitempty" xml:"page_num,omitempty" require:"true"`
	// 每页显示数量 最多100
	PageSize *int64 `json:"page_size,omitempty" xml:"page_size,omitempty" require:"true"`
	// 业务发起方分布式数字身份
	PlatformDid *string `json:"platform_did,omitempty" xml:"platform_did,omitempty" require:"true"`
	// 产品id
	// A模式：PRODUCT_MYBANK， A+模式：PRODUCT_MYBANK_A_PLUS， B模式：PRODUCT_MYBANK_B
	ProductId *string `json:"product_id,omitempty" xml:"product_id,omitempty" require:"true"`
	// 查询起始时间
	StartDate *string `json:"start_date,omitempty" xml:"start_date,omitempty" require:"true"`
}

func (s QueryCreditStatementRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryCreditStatementRequest) GoString() string {
	return s.String()
}

func (s *QueryCreditStatementRequest) SetAuthToken(v string) *QueryCreditStatementRequest {
	s.AuthToken = &v
	return s
}

func (s *QueryCreditStatementRequest) SetProductInstanceId(v string) *QueryCreditStatementRequest {
	s.ProductInstanceId = &v
	return s
}

func (s *QueryCreditStatementRequest) SetDid(v string) *QueryCreditStatementRequest {
	s.Did = &v
	return s
}

func (s *QueryCreditStatementRequest) SetEndDate(v string) *QueryCreditStatementRequest {
	s.EndDate = &v
	return s
}

func (s *QueryCreditStatementRequest) SetGroupPlatformDid(v string) *QueryCreditStatementRequest {
	s.GroupPlatformDid = &v
	return s
}

func (s *QueryCreditStatementRequest) SetPageNum(v int64) *QueryCreditStatementRequest {
	s.PageNum = &v
	return s
}

func (s *QueryCreditStatementRequest) SetPageSize(v int64) *QueryCreditStatementRequest {
	s.PageSize = &v
	return s
}

func (s *QueryCreditStatementRequest) SetPlatformDid(v string) *QueryCreditStatementRequest {
	s.PlatformDid = &v
	return s
}

func (s *QueryCreditStatementRequest) SetProductId(v string) *QueryCreditStatementRequest {
	s.ProductId = &v
	return s
}

func (s *QueryCreditStatementRequest) SetStartDate(v string) *QueryCreditStatementRequest {
	s.StartDate = &v
	return s
}

type QueryCreditStatementResponse struct {
	// 请求唯一ID，用于链路跟踪和问题排查
	ReqMsgId *string `json:"req_msg_id,omitempty" xml:"req_msg_id,omitempty"`
	// 结果码，一般OK表示调用成功
	ResultCode *string `json:"result_code,omitempty" xml:"result_code,omitempty"`
	// 异常信息的文本描述
	ResultMsg *string `json:"result_msg,omitempty" xml:"result_msg,omitempty"`
	// 页数 从1开始
	PageNum *int64 `json:"page_num,omitempty" xml:"page_num,omitempty"`
	// 每页显示数量
	PageSize *int64 `json:"page_size,omitempty" xml:"page_size,omitempty"`
	// 流水列表
	StatementInfos []*StatementInfo `json:"statement_infos,omitempty" xml:"statement_infos,omitempty" type:"Repeated"`
	// 数据总量
	TotalCount *int64 `json:"total_count,omitempty" xml:"total_count,omitempty"`
}

func (s QueryCreditStatementResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryCreditStatementResponse) GoString() string {
	return s.String()
}

func (s *QueryCreditStatementResponse) SetReqMsgId(v string) *QueryCreditStatementResponse {
	s.ReqMsgId = &v
	return s
}

func (s *QueryCreditStatementResponse) SetResultCode(v string) *QueryCreditStatementResponse {
	s.ResultCode = &v
	return s
}

func (s *QueryCreditStatementResponse) SetResultMsg(v string) *QueryCreditStatementResponse {
	s.ResultMsg = &v
	return s
}

func (s *QueryCreditStatementResponse) SetPageNum(v int64) *QueryCreditStatementResponse {
	s.PageNum = &v
	return s
}

func (s *QueryCreditStatementResponse) SetPageSize(v int64) *QueryCreditStatementResponse {
	s.PageSize = &v
	return s
}

func (s *QueryCreditStatementResponse) SetStatementInfos(v []*StatementInfo) *QueryCreditStatementResponse {
	s.StatementInfos = v
	return s
}

func (s *QueryCreditStatementResponse) SetTotalCount(v int64) *QueryCreditStatementResponse {
	s.TotalCount = &v
	return s
}

type CreateCreditIssuetransferRequest struct {
	// OAuth模式下的授权token
	AuthToken         *string `json:"auth_token,omitempty" xml:"auth_token,omitempty"`
	ProductInstanceId *string `json:"product_instance_id,omitempty" xml:"product_instance_id,omitempty"`
	// 转让结果回调url
	CallbackUrl *string `json:"callback_url,omitempty" xml:"callback_url,omitempty" require:"true"`
	// 集团平台方分布式数字身份
	GroupPlatformDid *string `json:"group_platform_did,omitempty" xml:"group_platform_did,omitempty" require:"true"`
	// 信用流转凭证
	IssueId *string `json:"issue_id,omitempty" xml:"issue_id,omitempty" require:"true"`
	// 全局业务号，保证唯一性，如拆分转让业务单号
	OutBizNo *string `json:"out_biz_no,omitempty" xml:"out_biz_no,omitempty" require:"true"`
	// 转入方分布式数字身份
	PayeeDid *string `json:"payee_did,omitempty" xml:"payee_did,omitempty" require:"true"`
	// 转出方分布式数字身份
	PayerDid *string `json:"payer_did,omitempty" xml:"payer_did,omitempty" require:"true"`
	// 业务发起方分布式数字身份
	PlatformDid *string `json:"platform_did,omitempty" xml:"platform_did,omitempty" require:"true"`
	// 产品id 目前使用PRODUCT_MYBANK
	ProductId *string `json:"product_id,omitempty" xml:"product_id,omitempty" require:"true"`
	// 转让金额
	TxAmt *string `json:"tx_amt,omitempty" xml:"tx_amt,omitempty" require:"true"`
}

func (s CreateCreditIssuetransferRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateCreditIssuetransferRequest) GoString() string {
	return s.String()
}

func (s *CreateCreditIssuetransferRequest) SetAuthToken(v string) *CreateCreditIssuetransferRequest {
	s.AuthToken = &v
	return s
}

func (s *CreateCreditIssuetransferRequest) SetProductInstanceId(v string) *CreateCreditIssuetransferRequest {
	s.ProductInstanceId = &v
	return s
}

func (s *CreateCreditIssuetransferRequest) SetCallbackUrl(v string) *CreateCreditIssuetransferRequest {
	s.CallbackUrl = &v
	return s
}

func (s *CreateCreditIssuetransferRequest) SetGroupPlatformDid(v string) *CreateCreditIssuetransferRequest {
	s.GroupPlatformDid = &v
	return s
}

func (s *CreateCreditIssuetransferRequest) SetIssueId(v string) *CreateCreditIssuetransferRequest {
	s.IssueId = &v
	return s
}

func (s *CreateCreditIssuetransferRequest) SetOutBizNo(v string) *CreateCreditIssuetransferRequest {
	s.OutBizNo = &v
	return s
}

func (s *CreateCreditIssuetransferRequest) SetPayeeDid(v string) *CreateCreditIssuetransferRequest {
	s.PayeeDid = &v
	return s
}

func (s *CreateCreditIssuetransferRequest) SetPayerDid(v string) *CreateCreditIssuetransferRequest {
	s.PayerDid = &v
	return s
}

func (s *CreateCreditIssuetransferRequest) SetPlatformDid(v string) *CreateCreditIssuetransferRequest {
	s.PlatformDid = &v
	return s
}

func (s *CreateCreditIssuetransferRequest) SetProductId(v string) *CreateCreditIssuetransferRequest {
	s.ProductId = &v
	return s
}

func (s *CreateCreditIssuetransferRequest) SetTxAmt(v string) *CreateCreditIssuetransferRequest {
	s.TxAmt = &v
	return s
}

type CreateCreditIssuetransferResponse struct {
	// 请求唯一ID，用于链路跟踪和问题排查
	ReqMsgId *string `json:"req_msg_id,omitempty" xml:"req_msg_id,omitempty"`
	// 结果码，一般OK表示调用成功
	ResultCode *string `json:"result_code,omitempty" xml:"result_code,omitempty"`
	// 异常信息的文本描述
	ResultMsg *string `json:"result_msg,omitempty" xml:"result_msg,omitempty"`
	// 转让申请成功金额
	SuccessApplyAmt *string `json:"success_apply_amt,omitempty" xml:"success_apply_amt,omitempty"`
	// 请求时传入的全局业务号
	OutBizNo *string `json:"out_biz_no,omitempty" xml:"out_biz_no,omitempty"`
}

func (s CreateCreditIssuetransferResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateCreditIssuetransferResponse) GoString() string {
	return s.String()
}

func (s *CreateCreditIssuetransferResponse) SetReqMsgId(v string) *CreateCreditIssuetransferResponse {
	s.ReqMsgId = &v
	return s
}

func (s *CreateCreditIssuetransferResponse) SetResultCode(v string) *CreateCreditIssuetransferResponse {
	s.ResultCode = &v
	return s
}

func (s *CreateCreditIssuetransferResponse) SetResultMsg(v string) *CreateCreditIssuetransferResponse {
	s.ResultMsg = &v
	return s
}

func (s *CreateCreditIssuetransferResponse) SetSuccessApplyAmt(v string) *CreateCreditIssuetransferResponse {
	s.SuccessApplyAmt = &v
	return s
}

func (s *CreateCreditIssuetransferResponse) SetOutBizNo(v string) *CreateCreditIssuetransferResponse {
	s.OutBizNo = &v
	return s
}

type QueryCreditIssuetransferRequest struct {
	// OAuth模式下的授权token
	AuthToken         *string `json:"auth_token,omitempty" xml:"auth_token,omitempty"`
	ProductInstanceId *string `json:"product_instance_id,omitempty" xml:"product_instance_id,omitempty"`
	// 集团平台方分布式数字身份
	GroupPlatformDid *string `json:"group_platform_did,omitempty" xml:"group_platform_did,omitempty" require:"true"`
	// 全局业务号，保证唯一性，如拆分转让业务单号
	OutBizNo *string `json:"out_biz_no,omitempty" xml:"out_biz_no,omitempty" require:"true"`
	// 业务发起方分布式数字身份
	PlatformDid *string `json:"platform_did,omitempty" xml:"platform_did,omitempty" require:"true"`
	// 产品id 目前使用PRODUCT_MYBANK
	ProductId *string `json:"product_id,omitempty" xml:"product_id,omitempty" require:"true"`
}

func (s QueryCreditIssuetransferRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryCreditIssuetransferRequest) GoString() string {
	return s.String()
}

func (s *QueryCreditIssuetransferRequest) SetAuthToken(v string) *QueryCreditIssuetransferRequest {
	s.AuthToken = &v
	return s
}

func (s *QueryCreditIssuetransferRequest) SetProductInstanceId(v string) *QueryCreditIssuetransferRequest {
	s.ProductInstanceId = &v
	return s
}

func (s *QueryCreditIssuetransferRequest) SetGroupPlatformDid(v string) *QueryCreditIssuetransferRequest {
	s.GroupPlatformDid = &v
	return s
}

func (s *QueryCreditIssuetransferRequest) SetOutBizNo(v string) *QueryCreditIssuetransferRequest {
	s.OutBizNo = &v
	return s
}

func (s *QueryCreditIssuetransferRequest) SetPlatformDid(v string) *QueryCreditIssuetransferRequest {
	s.PlatformDid = &v
	return s
}

func (s *QueryCreditIssuetransferRequest) SetProductId(v string) *QueryCreditIssuetransferRequest {
	s.ProductId = &v
	return s
}

type QueryCreditIssuetransferResponse struct {
	// 请求唯一ID，用于链路跟踪和问题排查
	ReqMsgId *string `json:"req_msg_id,omitempty" xml:"req_msg_id,omitempty"`
	// 结果码，一般OK表示调用成功
	ResultCode *string `json:"result_code,omitempty" xml:"result_code,omitempty"`
	// 异常信息的文本描述
	ResultMsg *string `json:"result_msg,omitempty" xml:"result_msg,omitempty"`
	// 错误信息
	ErrMsg *string `json:"err_msg,omitempty" xml:"err_msg,omitempty"`
	// 转让失败金额
	FailAmt *string `json:"fail_amt,omitempty" xml:"fail_amt,omitempty"`
	// 信用凭证号
	IssueId *string `json:"issue_id,omitempty" xml:"issue_id,omitempty"`
	// 请求时传入的全局业务号
	OutBizNo *string `json:"out_biz_no,omitempty" xml:"out_biz_no,omitempty"`
	// 转入方分布式数字身份
	PayeeDid *string `json:"payee_did,omitempty" xml:"payee_did,omitempty"`
	// 转出方分布式数字身份
	PayerDid *string `json:"payer_did,omitempty" xml:"payer_did,omitempty"`
	// 产品id
	ProductId *string `json:"product_id,omitempty" xml:"product_id,omitempty"`
	// 转让结果状态 -1:转让失败状态， 0:转让未完成状态， 1:转让部分成功状态，2:转让成功状态
	Status *int64 `json:"status,omitempty" xml:"status,omitempty"`
	// 转让成功金额
	SuccessAmt *string `json:"success_amt,omitempty" xml:"success_amt,omitempty"`
}

func (s QueryCreditIssuetransferResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryCreditIssuetransferResponse) GoString() string {
	return s.String()
}

func (s *QueryCreditIssuetransferResponse) SetReqMsgId(v string) *QueryCreditIssuetransferResponse {
	s.ReqMsgId = &v
	return s
}

func (s *QueryCreditIssuetransferResponse) SetResultCode(v string) *QueryCreditIssuetransferResponse {
	s.ResultCode = &v
	return s
}

func (s *QueryCreditIssuetransferResponse) SetResultMsg(v string) *QueryCreditIssuetransferResponse {
	s.ResultMsg = &v
	return s
}

func (s *QueryCreditIssuetransferResponse) SetErrMsg(v string) *QueryCreditIssuetransferResponse {
	s.ErrMsg = &v
	return s
}

func (s *QueryCreditIssuetransferResponse) SetFailAmt(v string) *QueryCreditIssuetransferResponse {
	s.FailAmt = &v
	return s
}

func (s *QueryCreditIssuetransferResponse) SetIssueId(v string) *QueryCreditIssuetransferResponse {
	s.IssueId = &v
	return s
}

func (s *QueryCreditIssuetransferResponse) SetOutBizNo(v string) *QueryCreditIssuetransferResponse {
	s.OutBizNo = &v
	return s
}

func (s *QueryCreditIssuetransferResponse) SetPayeeDid(v string) *QueryCreditIssuetransferResponse {
	s.PayeeDid = &v
	return s
}

func (s *QueryCreditIssuetransferResponse) SetPayerDid(v string) *QueryCreditIssuetransferResponse {
	s.PayerDid = &v
	return s
}

func (s *QueryCreditIssuetransferResponse) SetProductId(v string) *QueryCreditIssuetransferResponse {
	s.ProductId = &v
	return s
}

func (s *QueryCreditIssuetransferResponse) SetStatus(v int64) *QueryCreditIssuetransferResponse {
	s.Status = &v
	return s
}

func (s *QueryCreditIssuetransferResponse) SetSuccessAmt(v string) *QueryCreditIssuetransferResponse {
	s.SuccessAmt = &v
	return s
}

type CreateCreditIssuefinanceRequest struct {
	// OAuth模式下的授权token
	AuthToken         *string `json:"auth_token,omitempty" xml:"auth_token,omitempty"`
	ProductInstanceId *string `json:"product_instance_id,omitempty" xml:"product_instance_id,omitempty"`
	// 转让结果回调url
	CallbackUrl *string `json:"callback_url,omitempty" xml:"callback_url,omitempty" require:"true"`
	// 集团平台方分布式数字身份
	GroupPlatformDid *string `json:"group_platform_did,omitempty" xml:"group_platform_did,omitempty" require:"true"`
	// 信用流转凭证
	IssueId *string `json:"issue_id,omitempty" xml:"issue_id,omitempty" require:"true"`
	// 融资申请方分布式数字身份
	LoanerDid *string `json:"loaner_did,omitempty" xml:"loaner_did,omitempty" require:"true"`
	// 全局业务号，保证唯一性，如融资业务单号
	OutBizNo *string `json:"out_biz_no,omitempty" xml:"out_biz_no,omitempty" require:"true"`
	// 业务发起方分布式数字身份
	PlatformDid *string `json:"platform_did,omitempty" xml:"platform_did,omitempty" require:"true"`
	// 产品id
	// A模式：PRODUCT_MYBANK， A+模式：PRODUCT_MYBANK_A_PLUS， B模式：PRODUCT_MYBANK_B
	ProductId *string `json:"product_id,omitempty" xml:"product_id,omitempty" require:"true"`
	// 融资金额
	TxAmt *string `json:"tx_amt,omitempty" xml:"tx_amt,omitempty" require:"true"`
}

func (s CreateCreditIssuefinanceRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateCreditIssuefinanceRequest) GoString() string {
	return s.String()
}

func (s *CreateCreditIssuefinanceRequest) SetAuthToken(v string) *CreateCreditIssuefinanceRequest {
	s.AuthToken = &v
	return s
}

func (s *CreateCreditIssuefinanceRequest) SetProductInstanceId(v string) *CreateCreditIssuefinanceRequest {
	s.ProductInstanceId = &v
	return s
}

func (s *CreateCreditIssuefinanceRequest) SetCallbackUrl(v string) *CreateCreditIssuefinanceRequest {
	s.CallbackUrl = &v
	return s
}

func (s *CreateCreditIssuefinanceRequest) SetGroupPlatformDid(v string) *CreateCreditIssuefinanceRequest {
	s.GroupPlatformDid = &v
	return s
}

func (s *CreateCreditIssuefinanceRequest) SetIssueId(v string) *CreateCreditIssuefinanceRequest {
	s.IssueId = &v
	return s
}

func (s *CreateCreditIssuefinanceRequest) SetLoanerDid(v string) *CreateCreditIssuefinanceRequest {
	s.LoanerDid = &v
	return s
}

func (s *CreateCreditIssuefinanceRequest) SetOutBizNo(v string) *CreateCreditIssuefinanceRequest {
	s.OutBizNo = &v
	return s
}

func (s *CreateCreditIssuefinanceRequest) SetPlatformDid(v string) *CreateCreditIssuefinanceRequest {
	s.PlatformDid = &v
	return s
}

func (s *CreateCreditIssuefinanceRequest) SetProductId(v string) *CreateCreditIssuefinanceRequest {
	s.ProductId = &v
	return s
}

func (s *CreateCreditIssuefinanceRequest) SetTxAmt(v string) *CreateCreditIssuefinanceRequest {
	s.TxAmt = &v
	return s
}

type CreateCreditIssuefinanceResponse struct {
	// 请求唯一ID，用于链路跟踪和问题排查
	ReqMsgId *string `json:"req_msg_id,omitempty" xml:"req_msg_id,omitempty"`
	// 结果码，一般OK表示调用成功
	ResultCode *string `json:"result_code,omitempty" xml:"result_code,omitempty"`
	// 异常信息的文本描述
	ResultMsg *string `json:"result_msg,omitempty" xml:"result_msg,omitempty"`
	// 融资申请失败，回转至无车承运平台金额
	FailApplyAmt *string `json:"fail_apply_amt,omitempty" xml:"fail_apply_amt,omitempty"`
	// 请求时传入的全局业务号
	OutBizNo *string `json:"out_biz_no,omitempty" xml:"out_biz_no,omitempty"`
	// 融资申请成功金额
	SuccessApplyAmt *string `json:"success_apply_amt,omitempty" xml:"success_apply_amt,omitempty"`
}

func (s CreateCreditIssuefinanceResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateCreditIssuefinanceResponse) GoString() string {
	return s.String()
}

func (s *CreateCreditIssuefinanceResponse) SetReqMsgId(v string) *CreateCreditIssuefinanceResponse {
	s.ReqMsgId = &v
	return s
}

func (s *CreateCreditIssuefinanceResponse) SetResultCode(v string) *CreateCreditIssuefinanceResponse {
	s.ResultCode = &v
	return s
}

func (s *CreateCreditIssuefinanceResponse) SetResultMsg(v string) *CreateCreditIssuefinanceResponse {
	s.ResultMsg = &v
	return s
}

func (s *CreateCreditIssuefinanceResponse) SetFailApplyAmt(v string) *CreateCreditIssuefinanceResponse {
	s.FailApplyAmt = &v
	return s
}

func (s *CreateCreditIssuefinanceResponse) SetOutBizNo(v string) *CreateCreditIssuefinanceResponse {
	s.OutBizNo = &v
	return s
}

func (s *CreateCreditIssuefinanceResponse) SetSuccessApplyAmt(v string) *CreateCreditIssuefinanceResponse {
	s.SuccessApplyAmt = &v
	return s
}

type QueryCreditIssuefinanceRequest struct {
	// OAuth模式下的授权token
	AuthToken         *string `json:"auth_token,omitempty" xml:"auth_token,omitempty"`
	ProductInstanceId *string `json:"product_instance_id,omitempty" xml:"product_instance_id,omitempty"`
	// 集团平台方分布式数字身份
	GroupPlatformDid *string `json:"group_platform_did,omitempty" xml:"group_platform_did,omitempty" require:"true"`
	// 融资申请方分布式数字身份
	LoanerDid *string `json:"loaner_did,omitempty" xml:"loaner_did,omitempty" require:"true"`
	// 全局业务号，保证唯一性，如融资业务单号
	OutBizNo *string `json:"out_biz_no,omitempty" xml:"out_biz_no,omitempty" require:"true"`
	// 业务发起方分布式数字身份
	PlatformDid *string `json:"platform_did,omitempty" xml:"platform_did,omitempty" require:"true"`
	// 产品id
	// A模式：PRODUCT_MYBANK， A+模式：PRODUCT_MYBANK_A_PLUS， B模式：PRODUCT_MYBANK_B
	ProductId *string `json:"product_id,omitempty" xml:"product_id,omitempty" require:"true"`
}

func (s QueryCreditIssuefinanceRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryCreditIssuefinanceRequest) GoString() string {
	return s.String()
}

func (s *QueryCreditIssuefinanceRequest) SetAuthToken(v string) *QueryCreditIssuefinanceRequest {
	s.AuthToken = &v
	return s
}

func (s *QueryCreditIssuefinanceRequest) SetProductInstanceId(v string) *QueryCreditIssuefinanceRequest {
	s.ProductInstanceId = &v
	return s
}

func (s *QueryCreditIssuefinanceRequest) SetGroupPlatformDid(v string) *QueryCreditIssuefinanceRequest {
	s.GroupPlatformDid = &v
	return s
}

func (s *QueryCreditIssuefinanceRequest) SetLoanerDid(v string) *QueryCreditIssuefinanceRequest {
	s.LoanerDid = &v
	return s
}

func (s *QueryCreditIssuefinanceRequest) SetOutBizNo(v string) *QueryCreditIssuefinanceRequest {
	s.OutBizNo = &v
	return s
}

func (s *QueryCreditIssuefinanceRequest) SetPlatformDid(v string) *QueryCreditIssuefinanceRequest {
	s.PlatformDid = &v
	return s
}

func (s *QueryCreditIssuefinanceRequest) SetProductId(v string) *QueryCreditIssuefinanceRequest {
	s.ProductId = &v
	return s
}

type QueryCreditIssuefinanceResponse struct {
	// 请求唯一ID，用于链路跟踪和问题排查
	ReqMsgId *string `json:"req_msg_id,omitempty" xml:"req_msg_id,omitempty"`
	// 结果码，一般OK表示调用成功
	ResultCode *string `json:"result_code,omitempty" xml:"result_code,omitempty"`
	// 异常信息的文本描述
	ResultMsg *string `json:"result_msg,omitempty" xml:"result_msg,omitempty"`
	// 错误信息
	ErrMsg *string `json:"err_msg,omitempty" xml:"err_msg,omitempty"`
	// 融资失败，回转至无车承运平台失败金额
	FailAmt *string `json:"fail_amt,omitempty" xml:"fail_amt,omitempty"`
	// 融资失败，回转至无车承运平台成功金额
	FailTransferAmt *string `json:"fail_transfer_amt,omitempty" xml:"fail_transfer_amt,omitempty"`
	// 请求时传入的全局业务号
	OutBizNo *string `json:"out_biz_no,omitempty" xml:"out_biz_no,omitempty"`
	// 融资结果状态 -1:融资未完成， 0:融资完成， 1:融资全部失败，逆流转回平台全部成功，2:融资部分成功，逆流转回平台全部成功，3:逆流转回平台发生失败，需人工介入
	Status *int64 `json:"status,omitempty" xml:"status,omitempty"`
	// 融资成功金额
	SuccessAmt *string `json:"success_amt,omitempty" xml:"success_amt,omitempty"`
}

func (s QueryCreditIssuefinanceResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryCreditIssuefinanceResponse) GoString() string {
	return s.String()
}

func (s *QueryCreditIssuefinanceResponse) SetReqMsgId(v string) *QueryCreditIssuefinanceResponse {
	s.ReqMsgId = &v
	return s
}

func (s *QueryCreditIssuefinanceResponse) SetResultCode(v string) *QueryCreditIssuefinanceResponse {
	s.ResultCode = &v
	return s
}

func (s *QueryCreditIssuefinanceResponse) SetResultMsg(v string) *QueryCreditIssuefinanceResponse {
	s.ResultMsg = &v
	return s
}

func (s *QueryCreditIssuefinanceResponse) SetErrMsg(v string) *QueryCreditIssuefinanceResponse {
	s.ErrMsg = &v
	return s
}

func (s *QueryCreditIssuefinanceResponse) SetFailAmt(v string) *QueryCreditIssuefinanceResponse {
	s.FailAmt = &v
	return s
}

func (s *QueryCreditIssuefinanceResponse) SetFailTransferAmt(v string) *QueryCreditIssuefinanceResponse {
	s.FailTransferAmt = &v
	return s
}

func (s *QueryCreditIssuefinanceResponse) SetOutBizNo(v string) *QueryCreditIssuefinanceResponse {
	s.OutBizNo = &v
	return s
}

func (s *QueryCreditIssuefinanceResponse) SetStatus(v int64) *QueryCreditIssuefinanceResponse {
	s.Status = &v
	return s
}

func (s *QueryCreditIssuefinanceResponse) SetSuccessAmt(v string) *QueryCreditIssuefinanceResponse {
	s.SuccessAmt = &v
	return s
}

type QueryCreditIssuereceivableRequest struct {
	// OAuth模式下的授权token
	AuthToken         *string `json:"auth_token,omitempty" xml:"auth_token,omitempty"`
	ProductInstanceId *string `json:"product_instance_id,omitempty" xml:"product_instance_id,omitempty"`
	// 集团平台方分布式数字身份
	GroupPlatformDid *string `json:"group_platform_did,omitempty" xml:"group_platform_did,omitempty" require:"true"`
	// 信用凭证id
	IssueId *string `json:"issue_id,omitempty" xml:"issue_id,omitempty" require:"true"`
	// 业务发起方分布式数字身份
	PlatformDid *string `json:"platform_did,omitempty" xml:"platform_did,omitempty" require:"true"`
	// PRODUCT_MYBANK  、PRODUCT_MYBANK_B
	ProductId *string `json:"product_id,omitempty" xml:"product_id,omitempty" require:"true"`
}

func (s QueryCreditIssuereceivableRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryCreditIssuereceivableRequest) GoString() string {
	return s.String()
}

func (s *QueryCreditIssuereceivableRequest) SetAuthToken(v string) *QueryCreditIssuereceivableRequest {
	s.AuthToken = &v
	return s
}

func (s *QueryCreditIssuereceivableRequest) SetProductInstanceId(v string) *QueryCreditIssuereceivableRequest {
	s.ProductInstanceId = &v
	return s
}

func (s *QueryCreditIssuereceivableRequest) SetGroupPlatformDid(v string) *QueryCreditIssuereceivableRequest {
	s.GroupPlatformDid = &v
	return s
}

func (s *QueryCreditIssuereceivableRequest) SetIssueId(v string) *QueryCreditIssuereceivableRequest {
	s.IssueId = &v
	return s
}

func (s *QueryCreditIssuereceivableRequest) SetPlatformDid(v string) *QueryCreditIssuereceivableRequest {
	s.PlatformDid = &v
	return s
}

func (s *QueryCreditIssuereceivableRequest) SetProductId(v string) *QueryCreditIssuereceivableRequest {
	s.ProductId = &v
	return s
}

type QueryCreditIssuereceivableResponse struct {
	// 请求唯一ID，用于链路跟踪和问题排查
	ReqMsgId *string `json:"req_msg_id,omitempty" xml:"req_msg_id,omitempty"`
	// 结果码，一般OK表示调用成功
	ResultCode *string `json:"result_code,omitempty" xml:"result_code,omitempty"`
	// 异常信息的文本描述
	ResultMsg *string `json:"result_msg,omitempty" xml:"result_msg,omitempty"`
	// 已清分金额
	ClearAmt *string `json:"clear_amt,omitempty" xml:"clear_amt,omitempty"`
	// 代偿金额
	CompensateAmt *string `json:"compensate_amt,omitempty" xml:"compensate_amt,omitempty"`
	// 错误信息
	ErrMsg *string `json:"err_msg,omitempty" xml:"err_msg,omitempty"`
	// 完成日期 long类型字符串
	FinishDate *string `json:"finish_date,omitempty" xml:"finish_date,omitempty"`
	// 是否代偿,true为代偿，false为不代偿
	HasCompensate *bool `json:"has_compensate,omitempty" xml:"has_compensate,omitempty"`
	// 信用凭证id
	IssueId *string `json:"issue_id,omitempty" xml:"issue_id,omitempty"`
	// 全局唯一业务号
	OutBizNo *string `json:"out_biz_no,omitempty" xml:"out_biz_no,omitempty"`
	// 外部订单号，也就是凭证发行时传递的支付订单
	OutOrderNo *string `json:"out_order_no,omitempty" xml:"out_order_no,omitempty"`
	// 产品id--PRODUCT_MYBANK 、PRODUCT_MYBANK_B;
	ProductId *string `json:"product_id,omitempty" xml:"product_id,omitempty"`
	// 清分结果状态--- 0：未清分， 1：清分完成
	Status *int64 `json:"status,omitempty" xml:"status,omitempty"`
	// 凭证总金额
	TotalAmt *string `json:"total_amt,omitempty" xml:"total_amt,omitempty"`
	// 清分类型--主动清分 、 到期清分
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s QueryCreditIssuereceivableResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryCreditIssuereceivableResponse) GoString() string {
	return s.String()
}

func (s *QueryCreditIssuereceivableResponse) SetReqMsgId(v string) *QueryCreditIssuereceivableResponse {
	s.ReqMsgId = &v
	return s
}

func (s *QueryCreditIssuereceivableResponse) SetResultCode(v string) *QueryCreditIssuereceivableResponse {
	s.ResultCode = &v
	return s
}

func (s *QueryCreditIssuereceivableResponse) SetResultMsg(v string) *QueryCreditIssuereceivableResponse {
	s.ResultMsg = &v
	return s
}

func (s *QueryCreditIssuereceivableResponse) SetClearAmt(v string) *QueryCreditIssuereceivableResponse {
	s.ClearAmt = &v
	return s
}

func (s *QueryCreditIssuereceivableResponse) SetCompensateAmt(v string) *QueryCreditIssuereceivableResponse {
	s.CompensateAmt = &v
	return s
}

func (s *QueryCreditIssuereceivableResponse) SetErrMsg(v string) *QueryCreditIssuereceivableResponse {
	s.ErrMsg = &v
	return s
}

func (s *QueryCreditIssuereceivableResponse) SetFinishDate(v string) *QueryCreditIssuereceivableResponse {
	s.FinishDate = &v
	return s
}

func (s *QueryCreditIssuereceivableResponse) SetHasCompensate(v bool) *QueryCreditIssuereceivableResponse {
	s.HasCompensate = &v
	return s
}

func (s *QueryCreditIssuereceivableResponse) SetIssueId(v string) *QueryCreditIssuereceivableResponse {
	s.IssueId = &v
	return s
}

func (s *QueryCreditIssuereceivableResponse) SetOutBizNo(v string) *QueryCreditIssuereceivableResponse {
	s.OutBizNo = &v
	return s
}

func (s *QueryCreditIssuereceivableResponse) SetOutOrderNo(v string) *QueryCreditIssuereceivableResponse {
	s.OutOrderNo = &v
	return s
}

func (s *QueryCreditIssuereceivableResponse) SetProductId(v string) *QueryCreditIssuereceivableResponse {
	s.ProductId = &v
	return s
}

func (s *QueryCreditIssuereceivableResponse) SetStatus(v int64) *QueryCreditIssuereceivableResponse {
	s.Status = &v
	return s
}

func (s *QueryCreditIssuereceivableResponse) SetTotalAmt(v string) *QueryCreditIssuereceivableResponse {
	s.TotalAmt = &v
	return s
}

func (s *QueryCreditIssuereceivableResponse) SetType(v string) *QueryCreditIssuereceivableResponse {
	s.Type = &v
	return s
}

type QueryCreditIssueamountRequest struct {
	// OAuth模式下的授权token
	AuthToken         *string `json:"auth_token,omitempty" xml:"auth_token,omitempty"`
	ProductInstanceId *string `json:"product_instance_id,omitempty" xml:"product_instance_id,omitempty"`
	// 指定查询者分布式数字身份
	Did *string `json:"did,omitempty" xml:"did,omitempty" require:"true"`
	// 集团平台方分布式数字身份
	GroupPlatformDid *string `json:"group_platform_did,omitempty" xml:"group_platform_did,omitempty" require:"true"`
	// 信用凭证id
	IssueId *string `json:"issue_id,omitempty" xml:"issue_id,omitempty" require:"true"`
	// 业务发起方分布式数字身份
	PlatformDid *string `json:"platform_did,omitempty" xml:"platform_did,omitempty" require:"true"`
	// 产品id 目前填写PRODUCT_MYBANK
	ProductId *string `json:"product_id,omitempty" xml:"product_id,omitempty" require:"true"`
}

func (s QueryCreditIssueamountRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryCreditIssueamountRequest) GoString() string {
	return s.String()
}

func (s *QueryCreditIssueamountRequest) SetAuthToken(v string) *QueryCreditIssueamountRequest {
	s.AuthToken = &v
	return s
}

func (s *QueryCreditIssueamountRequest) SetProductInstanceId(v string) *QueryCreditIssueamountRequest {
	s.ProductInstanceId = &v
	return s
}

func (s *QueryCreditIssueamountRequest) SetDid(v string) *QueryCreditIssueamountRequest {
	s.Did = &v
	return s
}

func (s *QueryCreditIssueamountRequest) SetGroupPlatformDid(v string) *QueryCreditIssueamountRequest {
	s.GroupPlatformDid = &v
	return s
}

func (s *QueryCreditIssueamountRequest) SetIssueId(v string) *QueryCreditIssueamountRequest {
	s.IssueId = &v
	return s
}

func (s *QueryCreditIssueamountRequest) SetPlatformDid(v string) *QueryCreditIssueamountRequest {
	s.PlatformDid = &v
	return s
}

func (s *QueryCreditIssueamountRequest) SetProductId(v string) *QueryCreditIssueamountRequest {
	s.ProductId = &v
	return s
}

type QueryCreditIssueamountResponse struct {
	// 请求唯一ID，用于链路跟踪和问题排查
	ReqMsgId *string `json:"req_msg_id,omitempty" xml:"req_msg_id,omitempty"`
	// 结果码，一般OK表示调用成功
	ResultCode *string `json:"result_code,omitempty" xml:"result_code,omitempty"`
	// 异常信息的文本描述
	ResultMsg *string `json:"result_msg,omitempty" xml:"result_msg,omitempty"`
	// 剩余金额
	BalanceAmt *string `json:"balance_amt,omitempty" xml:"balance_amt,omitempty"`
	// 已清分金额
	CashAmt *string `json:"cash_amt,omitempty" xml:"cash_amt,omitempty"`
	// 欠的滞纳金金额
	DebtIntAmt *string `json:"debt_int_amt,omitempty" xml:"debt_int_amt,omitempty"`
	// 欠款本金金额
	DebtPrinAmt *string `json:"debt_prin_amt,omitempty" xml:"debt_prin_amt,omitempty"`
	// 该凭证下的总欠款金额
	DebtTotalAmt *string `json:"debt_total_amt,omitempty" xml:"debt_total_amt,omitempty"`
	// 信用凭证到期时间
	ExpireDate *string `json:"expire_date,omitempty" xml:"expire_date,omitempty"`
	// 冻结的金额
	FreezeAmt *string `json:"freeze_amt,omitempty" xml:"freeze_amt,omitempty"`
	// 信用凭证id
	IssueId *string `json:"issue_id,omitempty" xml:"issue_id,omitempty"`
	// 累计已还滞纳金金额
	PaidIntAmt *string `json:"paid_int_amt,omitempty" xml:"paid_int_amt,omitempty"`
	// 凭证总金额
	TicketAmt *string `json:"ticket_amt,omitempty" xml:"ticket_amt,omitempty"`
}

func (s QueryCreditIssueamountResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryCreditIssueamountResponse) GoString() string {
	return s.String()
}

func (s *QueryCreditIssueamountResponse) SetReqMsgId(v string) *QueryCreditIssueamountResponse {
	s.ReqMsgId = &v
	return s
}

func (s *QueryCreditIssueamountResponse) SetResultCode(v string) *QueryCreditIssueamountResponse {
	s.ResultCode = &v
	return s
}

func (s *QueryCreditIssueamountResponse) SetResultMsg(v string) *QueryCreditIssueamountResponse {
	s.ResultMsg = &v
	return s
}

func (s *QueryCreditIssueamountResponse) SetBalanceAmt(v string) *QueryCreditIssueamountResponse {
	s.BalanceAmt = &v
	return s
}

func (s *QueryCreditIssueamountResponse) SetCashAmt(v string) *QueryCreditIssueamountResponse {
	s.CashAmt = &v
	return s
}

func (s *QueryCreditIssueamountResponse) SetDebtIntAmt(v string) *QueryCreditIssueamountResponse {
	s.DebtIntAmt = &v
	return s
}

func (s *QueryCreditIssueamountResponse) SetDebtPrinAmt(v string) *QueryCreditIssueamountResponse {
	s.DebtPrinAmt = &v
	return s
}

func (s *QueryCreditIssueamountResponse) SetDebtTotalAmt(v string) *QueryCreditIssueamountResponse {
	s.DebtTotalAmt = &v
	return s
}

func (s *QueryCreditIssueamountResponse) SetExpireDate(v string) *QueryCreditIssueamountResponse {
	s.ExpireDate = &v
	return s
}

func (s *QueryCreditIssueamountResponse) SetFreezeAmt(v string) *QueryCreditIssueamountResponse {
	s.FreezeAmt = &v
	return s
}

func (s *QueryCreditIssueamountResponse) SetIssueId(v string) *QueryCreditIssueamountResponse {
	s.IssueId = &v
	return s
}

func (s *QueryCreditIssueamountResponse) SetPaidIntAmt(v string) *QueryCreditIssueamountResponse {
	s.PaidIntAmt = &v
	return s
}

func (s *QueryCreditIssueamountResponse) SetTicketAmt(v string) *QueryCreditIssueamountResponse {
	s.TicketAmt = &v
	return s
}

type CallbackCreditCommonRequest struct {
	// OAuth模式下的授权token
	AuthToken         *string `json:"auth_token,omitempty" xml:"auth_token,omitempty"`
	ProductInstanceId *string `json:"product_instance_id,omitempty" xml:"product_instance_id,omitempty"`
	// 回调数据，根据不同msg_method 返回不同结构
	MsgData *string `json:"msg_data,omitempty" xml:"msg_data,omitempty" require:"true"`
	// 回调方法类型
	MsgMethod *string `json:"msg_method,omitempty" xml:"msg_method,omitempty" require:"true"`
}

func (s CallbackCreditCommonRequest) String() string {
	return tea.Prettify(s)
}

func (s CallbackCreditCommonRequest) GoString() string {
	return s.String()
}

func (s *CallbackCreditCommonRequest) SetAuthToken(v string) *CallbackCreditCommonRequest {
	s.AuthToken = &v
	return s
}

func (s *CallbackCreditCommonRequest) SetProductInstanceId(v string) *CallbackCreditCommonRequest {
	s.ProductInstanceId = &v
	return s
}

func (s *CallbackCreditCommonRequest) SetMsgData(v string) *CallbackCreditCommonRequest {
	s.MsgData = &v
	return s
}

func (s *CallbackCreditCommonRequest) SetMsgMethod(v string) *CallbackCreditCommonRequest {
	s.MsgMethod = &v
	return s
}

type CallbackCreditCommonResponse struct {
	// 请求唯一ID，用于链路跟踪和问题排查
	ReqMsgId *string `json:"req_msg_id,omitempty" xml:"req_msg_id,omitempty"`
	// 结果码，一般OK表示调用成功
	ResultCode *string `json:"result_code,omitempty" xml:"result_code,omitempty"`
	// 异常信息的文本描述
	ResultMsg *string `json:"result_msg,omitempty" xml:"result_msg,omitempty"`
	// 是否已接收
	Received *bool `json:"received,omitempty" xml:"received,omitempty"`
}

func (s CallbackCreditCommonResponse) String() string {
	return tea.Prettify(s)
}

func (s CallbackCreditCommonResponse) GoString() string {
	return s.String()
}

func (s *CallbackCreditCommonResponse) SetReqMsgId(v string) *CallbackCreditCommonResponse {
	s.ReqMsgId = &v
	return s
}

func (s *CallbackCreditCommonResponse) SetResultCode(v string) *CallbackCreditCommonResponse {
	s.ResultCode = &v
	return s
}

func (s *CallbackCreditCommonResponse) SetResultMsg(v string) *CallbackCreditCommonResponse {
	s.ResultMsg = &v
	return s
}

func (s *CallbackCreditCommonResponse) SetReceived(v bool) *CallbackCreditCommonResponse {
	s.Received = &v
	return s
}

type ApplyCreditIssueclearRequest struct {
	// OAuth模式下的授权token
	AuthToken         *string `json:"auth_token,omitempty" xml:"auth_token,omitempty"`
	ProductInstanceId *string `json:"product_instance_id,omitempty" xml:"product_instance_id,omitempty"`
	// 申请日期，不输入则是当前时间
	ApplyDate *int64 `json:"apply_date,omitempty" xml:"apply_date,omitempty"`
	// 清分金额，当前只支持全额清分，不填就是默认全额清分
	ClearAmt *string `json:"clear_amt,omitempty" xml:"clear_amt,omitempty"`
	// 主动清分方分布式数字身份
	ClearDid *string `json:"clear_did,omitempty" xml:"clear_did,omitempty" require:"true"`
	// 扩展字段
	ExtInfo *string `json:"ext_info,omitempty" xml:"ext_info,omitempty"`
	// 集团平台方分布式数字身份
	GroupPlatformDid *string `json:"group_platform_did,omitempty" xml:"group_platform_did,omitempty" require:"true"`
	// 凭证id
	IssueId *string `json:"issue_id,omitempty" xml:"issue_id,omitempty" require:"true"`
	// 全局唯一业务流水号
	OutBizNo *string `json:"out_biz_no,omitempty" xml:"out_biz_no,omitempty" require:"true"`
	// 外部订单号，此为支付单号
	OutOrderNo *string `json:"out_order_no,omitempty" xml:"out_order_no,omitempty" require:"true"`
	// 业务发起方分布式数字身份
	PlatformDid *string `json:"platform_did,omitempty" xml:"platform_did,omitempty" require:"true"`
	// 产品id
	// A模式：PRODUCT_MYBANK， A+模式：PRODUCT_MYBANK_A_PLUS， B模式：PRODUCT_MYBANK_B
	ProductId *string `json:"product_id,omitempty" xml:"product_id,omitempty" require:"true"`
}

func (s ApplyCreditIssueclearRequest) String() string {
	return tea.Prettify(s)
}

func (s ApplyCreditIssueclearRequest) GoString() string {
	return s.String()
}

func (s *ApplyCreditIssueclearRequest) SetAuthToken(v string) *ApplyCreditIssueclearRequest {
	s.AuthToken = &v
	return s
}

func (s *ApplyCreditIssueclearRequest) SetProductInstanceId(v string) *ApplyCreditIssueclearRequest {
	s.ProductInstanceId = &v
	return s
}

func (s *ApplyCreditIssueclearRequest) SetApplyDate(v int64) *ApplyCreditIssueclearRequest {
	s.ApplyDate = &v
	return s
}

func (s *ApplyCreditIssueclearRequest) SetClearAmt(v string) *ApplyCreditIssueclearRequest {
	s.ClearAmt = &v
	return s
}

func (s *ApplyCreditIssueclearRequest) SetClearDid(v string) *ApplyCreditIssueclearRequest {
	s.ClearDid = &v
	return s
}

func (s *ApplyCreditIssueclearRequest) SetExtInfo(v string) *ApplyCreditIssueclearRequest {
	s.ExtInfo = &v
	return s
}

func (s *ApplyCreditIssueclearRequest) SetGroupPlatformDid(v string) *ApplyCreditIssueclearRequest {
	s.GroupPlatformDid = &v
	return s
}

func (s *ApplyCreditIssueclearRequest) SetIssueId(v string) *ApplyCreditIssueclearRequest {
	s.IssueId = &v
	return s
}

func (s *ApplyCreditIssueclearRequest) SetOutBizNo(v string) *ApplyCreditIssueclearRequest {
	s.OutBizNo = &v
	return s
}

func (s *ApplyCreditIssueclearRequest) SetOutOrderNo(v string) *ApplyCreditIssueclearRequest {
	s.OutOrderNo = &v
	return s
}

func (s *ApplyCreditIssueclearRequest) SetPlatformDid(v string) *ApplyCreditIssueclearRequest {
	s.PlatformDid = &v
	return s
}

func (s *ApplyCreditIssueclearRequest) SetProductId(v string) *ApplyCreditIssueclearRequest {
	s.ProductId = &v
	return s
}

type ApplyCreditIssueclearResponse struct {
	// 请求唯一ID，用于链路跟踪和问题排查
	ReqMsgId *string `json:"req_msg_id,omitempty" xml:"req_msg_id,omitempty"`
	// 结果码，一般OK表示调用成功
	ResultCode *string `json:"result_code,omitempty" xml:"result_code,omitempty"`
	// 异常信息的文本描述
	ResultMsg *string `json:"result_msg,omitempty" xml:"result_msg,omitempty"`
	// 全局唯一业务流水号
	OutBizNo *string `json:"out_biz_no,omitempty" xml:"out_biz_no,omitempty"`
}

func (s ApplyCreditIssueclearResponse) String() string {
	return tea.Prettify(s)
}

func (s ApplyCreditIssueclearResponse) GoString() string {
	return s.String()
}

func (s *ApplyCreditIssueclearResponse) SetReqMsgId(v string) *ApplyCreditIssueclearResponse {
	s.ReqMsgId = &v
	return s
}

func (s *ApplyCreditIssueclearResponse) SetResultCode(v string) *ApplyCreditIssueclearResponse {
	s.ResultCode = &v
	return s
}

func (s *ApplyCreditIssueclearResponse) SetResultMsg(v string) *ApplyCreditIssueclearResponse {
	s.ResultMsg = &v
	return s
}

func (s *ApplyCreditIssueclearResponse) SetOutBizNo(v string) *ApplyCreditIssueclearResponse {
	s.OutBizNo = &v
	return s
}

type SendCreditProxyRequest struct {
	// OAuth模式下的授权token
	AuthToken         *string `json:"auth_token,omitempty" xml:"auth_token,omitempty"`
	ProductInstanceId *string `json:"product_instance_id,omitempty" xml:"product_instance_id,omitempty"`
	// 代理请求方法名
	ApplyMethod *string `json:"apply_method,omitempty" xml:"apply_method,omitempty" require:"true"`
	// 代理请求数据
	ApplyData *string `json:"apply_data,omitempty" xml:"apply_data,omitempty" require:"true"`
}

func (s SendCreditProxyRequest) String() string {
	return tea.Prettify(s)
}

func (s SendCreditProxyRequest) GoString() string {
	return s.String()
}

func (s *SendCreditProxyRequest) SetAuthToken(v string) *SendCreditProxyRequest {
	s.AuthToken = &v
	return s
}

func (s *SendCreditProxyRequest) SetProductInstanceId(v string) *SendCreditProxyRequest {
	s.ProductInstanceId = &v
	return s
}

func (s *SendCreditProxyRequest) SetApplyMethod(v string) *SendCreditProxyRequest {
	s.ApplyMethod = &v
	return s
}

func (s *SendCreditProxyRequest) SetApplyData(v string) *SendCreditProxyRequest {
	s.ApplyData = &v
	return s
}

type SendCreditProxyResponse struct {
	// 请求唯一ID，用于链路跟踪和问题排查
	ReqMsgId *string `json:"req_msg_id,omitempty" xml:"req_msg_id,omitempty"`
	// 结果码，一般OK表示调用成功
	ResultCode *string `json:"result_code,omitempty" xml:"result_code,omitempty"`
	// 异常信息的文本描述
	ResultMsg *string `json:"result_msg,omitempty" xml:"result_msg,omitempty"`
}

func (s SendCreditProxyResponse) String() string {
	return tea.Prettify(s)
}

func (s SendCreditProxyResponse) GoString() string {
	return s.String()
}

func (s *SendCreditProxyResponse) SetReqMsgId(v string) *SendCreditProxyResponse {
	s.ReqMsgId = &v
	return s
}

func (s *SendCreditProxyResponse) SetResultCode(v string) *SendCreditProxyResponse {
	s.ResultCode = &v
	return s
}

func (s *SendCreditProxyResponse) SetResultMsg(v string) *SendCreditProxyResponse {
	s.ResultMsg = &v
	return s
}

type CheckCreditWaybillRequest struct {
	// OAuth模式下的授权token
	AuthToken         *string `json:"auth_token,omitempty" xml:"auth_token,omitempty"`
	ProductInstanceId *string `json:"product_instance_id,omitempty" xml:"product_instance_id,omitempty"`
	// 集团平台方分布式数字身份
	GroupPlatformDid *string `json:"group_platform_did,omitempty" xml:"group_platform_did,omitempty" require:"true"`
	// 业务发起方分布式数字身
	PlatformDid *string `json:"platform_did,omitempty" xml:"platform_did,omitempty" require:"true"`
	// 产品id
	// A模式：PRODUCT_MYBANK， A+模式：PRODUCT_MYBANK_A_PLUS， B模式：PRODUCT_MYBANK_B
	ProductId *string `json:"product_id,omitempty" xml:"product_id,omitempty" require:"true"`
	// 运单列表
	WaybillList []*string `json:"waybill_list,omitempty" xml:"waybill_list,omitempty" require:"true" type:"Repeated"`
}

func (s CheckCreditWaybillRequest) String() string {
	return tea.Prettify(s)
}

func (s CheckCreditWaybillRequest) GoString() string {
	return s.String()
}

func (s *CheckCreditWaybillRequest) SetAuthToken(v string) *CheckCreditWaybillRequest {
	s.AuthToken = &v
	return s
}

func (s *CheckCreditWaybillRequest) SetProductInstanceId(v string) *CheckCreditWaybillRequest {
	s.ProductInstanceId = &v
	return s
}

func (s *CheckCreditWaybillRequest) SetGroupPlatformDid(v string) *CheckCreditWaybillRequest {
	s.GroupPlatformDid = &v
	return s
}

func (s *CheckCreditWaybillRequest) SetPlatformDid(v string) *CheckCreditWaybillRequest {
	s.PlatformDid = &v
	return s
}

func (s *CheckCreditWaybillRequest) SetProductId(v string) *CheckCreditWaybillRequest {
	s.ProductId = &v
	return s
}

func (s *CheckCreditWaybillRequest) SetWaybillList(v []*string) *CheckCreditWaybillRequest {
	s.WaybillList = v
	return s
}

type CheckCreditWaybillResponse struct {
	// 请求唯一ID，用于链路跟踪和问题排查
	ReqMsgId *string `json:"req_msg_id,omitempty" xml:"req_msg_id,omitempty"`
	// 结果码，一般OK表示调用成功
	ResultCode *string `json:"result_code,omitempty" xml:"result_code,omitempty"`
	// 异常信息的文本描述
	ResultMsg *string `json:"result_msg,omitempty" xml:"result_msg,omitempty"`
	// 核验结果列表,  格式为 运单号:核验状态， 核验状态包含  0:运单未知状态   1:通过   2:不通过   3:未核验   4:没有指定信息
	CheckResult []*string `json:"check_result,omitempty" xml:"check_result,omitempty" type:"Repeated"`
}

func (s CheckCreditWaybillResponse) String() string {
	return tea.Prettify(s)
}

func (s CheckCreditWaybillResponse) GoString() string {
	return s.String()
}

func (s *CheckCreditWaybillResponse) SetReqMsgId(v string) *CheckCreditWaybillResponse {
	s.ReqMsgId = &v
	return s
}

func (s *CheckCreditWaybillResponse) SetResultCode(v string) *CheckCreditWaybillResponse {
	s.ResultCode = &v
	return s
}

func (s *CheckCreditWaybillResponse) SetResultMsg(v string) *CheckCreditWaybillResponse {
	s.ResultMsg = &v
	return s
}

func (s *CheckCreditWaybillResponse) SetCheckResult(v []*string) *CheckCreditWaybillResponse {
	s.CheckResult = v
	return s
}

type ReopenCreditDriverRequest struct {
	// OAuth模式下的授权token
	AuthToken         *string `json:"auth_token,omitempty" xml:"auth_token,omitempty"`
	ProductInstanceId *string `json:"product_instance_id,omitempty" xml:"product_instance_id,omitempty"`
	// 产品id
	ProductId *string `json:"product_id,omitempty" xml:"product_id,omitempty" require:"true"`
	// 集团平台did
	GroupPlatformDid *string `json:"group_platform_did,omitempty" xml:"group_platform_did,omitempty" require:"true"`
	// 司机did
	DriverDid *string `json:"driver_did,omitempty" xml:"driver_did,omitempty" require:"true"`
	// 云资金商户id
	AccountId *string `json:"account_id,omitempty" xml:"account_id,omitempty" require:"true"`
	// 回调地址
	CallbackUrl *string `json:"callback_url,omitempty" xml:"callback_url,omitempty" require:"true"`
}

func (s ReopenCreditDriverRequest) String() string {
	return tea.Prettify(s)
}

func (s ReopenCreditDriverRequest) GoString() string {
	return s.String()
}

func (s *ReopenCreditDriverRequest) SetAuthToken(v string) *ReopenCreditDriverRequest {
	s.AuthToken = &v
	return s
}

func (s *ReopenCreditDriverRequest) SetProductInstanceId(v string) *ReopenCreditDriverRequest {
	s.ProductInstanceId = &v
	return s
}

func (s *ReopenCreditDriverRequest) SetProductId(v string) *ReopenCreditDriverRequest {
	s.ProductId = &v
	return s
}

func (s *ReopenCreditDriverRequest) SetGroupPlatformDid(v string) *ReopenCreditDriverRequest {
	s.GroupPlatformDid = &v
	return s
}

func (s *ReopenCreditDriverRequest) SetDriverDid(v string) *ReopenCreditDriverRequest {
	s.DriverDid = &v
	return s
}

func (s *ReopenCreditDriverRequest) SetAccountId(v string) *ReopenCreditDriverRequest {
	s.AccountId = &v
	return s
}

func (s *ReopenCreditDriverRequest) SetCallbackUrl(v string) *ReopenCreditDriverRequest {
	s.CallbackUrl = &v
	return s
}

type ReopenCreditDriverResponse struct {
	// 请求唯一ID，用于链路跟踪和问题排查
	ReqMsgId *string `json:"req_msg_id,omitempty" xml:"req_msg_id,omitempty"`
	// 结果码，一般OK表示调用成功
	ResultCode *string `json:"result_code,omitempty" xml:"result_code,omitempty"`
	// 异常信息的文本描述
	ResultMsg *string `json:"result_msg,omitempty" xml:"result_msg,omitempty"`
	// 申请id
	ApplyId *string `json:"apply_id,omitempty" xml:"apply_id,omitempty"`
	// 会员注册状态
	Status *int64 `json:"status,omitempty" xml:"status,omitempty"`
}

func (s ReopenCreditDriverResponse) String() string {
	return tea.Prettify(s)
}

func (s ReopenCreditDriverResponse) GoString() string {
	return s.String()
}

func (s *ReopenCreditDriverResponse) SetReqMsgId(v string) *ReopenCreditDriverResponse {
	s.ReqMsgId = &v
	return s
}

func (s *ReopenCreditDriverResponse) SetResultCode(v string) *ReopenCreditDriverResponse {
	s.ResultCode = &v
	return s
}

func (s *ReopenCreditDriverResponse) SetResultMsg(v string) *ReopenCreditDriverResponse {
	s.ResultMsg = &v
	return s
}

func (s *ReopenCreditDriverResponse) SetApplyId(v string) *ReopenCreditDriverResponse {
	s.ApplyId = &v
	return s
}

func (s *ReopenCreditDriverResponse) SetStatus(v int64) *ReopenCreditDriverResponse {
	s.Status = &v
	return s
}

type UploadCreditAuthorizationRequest struct {
	// OAuth模式下的授权token
	AuthToken         *string `json:"auth_token,omitempty" xml:"auth_token,omitempty"`
	ProductInstanceId *string `json:"product_instance_id,omitempty" xml:"product_instance_id,omitempty"`
	// 签署货主统一社会信用代码
	ConsignorCertNo *string `json:"consignor_cert_no,omitempty" xml:"consignor_cert_no,omitempty" require:"true"`
	// 签署货主did
	ConsignorDid *string `json:"consignor_did,omitempty" xml:"consignor_did,omitempty" require:"true"`
	// 签署货主企业名称
	ConsignorName *string `json:"consignor_name,omitempty" xml:"consignor_name,omitempty" require:"true"`
	// 授权的货主平台账号
	ConsignorPlatformAccount *string `json:"consignor_platform_account,omitempty" xml:"consignor_platform_account,omitempty" require:"true"`
	// 协议到期日
	ExpireDate *string `json:"expire_date,omitempty" xml:"expire_date,omitempty" require:"true"`
	// 集团平台did
	GroupPlatformDid *string `json:"group_platform_did,omitempty" xml:"group_platform_did,omitempty" require:"true"`
	// 是否包含账号授权条款
	HasAuthorizationClause *bool `json:"has_authorization_clause,omitempty" xml:"has_authorization_clause,omitempty" require:"true"`
	// 签署子公司统一社会信用代码
	PlatformCertNo *string `json:"platform_cert_no,omitempty" xml:"platform_cert_no,omitempty" require:"true"`
	// 签署子公司did
	PlatformDid *string `json:"platform_did,omitempty" xml:"platform_did,omitempty" require:"true"`
	// 签署子公司企业名称
	//
	PlatformName *string `json:"platform_name,omitempty" xml:"platform_name,omitempty" require:"true"`
	// 线下协议影像件id与文件哈希, 多id以逗号分隔, 最多支持10个
	ProtocolImgId *string `json:"protocol_img_id,omitempty" xml:"protocol_img_id,omitempty" require:"true"`
	// 协议名称
	ProtocolName *string `json:"protocol_name,omitempty" xml:"protocol_name,omitempty" require:"true"`
	// 线下协议编号
	ProtocolNo *string `json:"protocol_no,omitempty" xml:"protocol_no,omitempty" require:"true"`
	// 签署时间
	SignDate *string `json:"sign_date,omitempty" xml:"sign_date,omitempty" require:"true"`
	// 产品id
	ProductId *string `json:"product_id,omitempty" xml:"product_id,omitempty" require:"true"`
}

func (s UploadCreditAuthorizationRequest) String() string {
	return tea.Prettify(s)
}

func (s UploadCreditAuthorizationRequest) GoString() string {
	return s.String()
}

func (s *UploadCreditAuthorizationRequest) SetAuthToken(v string) *UploadCreditAuthorizationRequest {
	s.AuthToken = &v
	return s
}

func (s *UploadCreditAuthorizationRequest) SetProductInstanceId(v string) *UploadCreditAuthorizationRequest {
	s.ProductInstanceId = &v
	return s
}

func (s *UploadCreditAuthorizationRequest) SetConsignorCertNo(v string) *UploadCreditAuthorizationRequest {
	s.ConsignorCertNo = &v
	return s
}

func (s *UploadCreditAuthorizationRequest) SetConsignorDid(v string) *UploadCreditAuthorizationRequest {
	s.ConsignorDid = &v
	return s
}

func (s *UploadCreditAuthorizationRequest) SetConsignorName(v string) *UploadCreditAuthorizationRequest {
	s.ConsignorName = &v
	return s
}

func (s *UploadCreditAuthorizationRequest) SetConsignorPlatformAccount(v string) *UploadCreditAuthorizationRequest {
	s.ConsignorPlatformAccount = &v
	return s
}

func (s *UploadCreditAuthorizationRequest) SetExpireDate(v string) *UploadCreditAuthorizationRequest {
	s.ExpireDate = &v
	return s
}

func (s *UploadCreditAuthorizationRequest) SetGroupPlatformDid(v string) *UploadCreditAuthorizationRequest {
	s.GroupPlatformDid = &v
	return s
}

func (s *UploadCreditAuthorizationRequest) SetHasAuthorizationClause(v bool) *UploadCreditAuthorizationRequest {
	s.HasAuthorizationClause = &v
	return s
}

func (s *UploadCreditAuthorizationRequest) SetPlatformCertNo(v string) *UploadCreditAuthorizationRequest {
	s.PlatformCertNo = &v
	return s
}

func (s *UploadCreditAuthorizationRequest) SetPlatformDid(v string) *UploadCreditAuthorizationRequest {
	s.PlatformDid = &v
	return s
}

func (s *UploadCreditAuthorizationRequest) SetPlatformName(v string) *UploadCreditAuthorizationRequest {
	s.PlatformName = &v
	return s
}

func (s *UploadCreditAuthorizationRequest) SetProtocolImgId(v string) *UploadCreditAuthorizationRequest {
	s.ProtocolImgId = &v
	return s
}

func (s *UploadCreditAuthorizationRequest) SetProtocolName(v string) *UploadCreditAuthorizationRequest {
	s.ProtocolName = &v
	return s
}

func (s *UploadCreditAuthorizationRequest) SetProtocolNo(v string) *UploadCreditAuthorizationRequest {
	s.ProtocolNo = &v
	return s
}

func (s *UploadCreditAuthorizationRequest) SetSignDate(v string) *UploadCreditAuthorizationRequest {
	s.SignDate = &v
	return s
}

func (s *UploadCreditAuthorizationRequest) SetProductId(v string) *UploadCreditAuthorizationRequest {
	s.ProductId = &v
	return s
}

type UploadCreditAuthorizationResponse struct {
	// 请求唯一ID，用于链路跟踪和问题排查
	ReqMsgId *string `json:"req_msg_id,omitempty" xml:"req_msg_id,omitempty"`
	// 结果码，一般OK表示调用成功
	ResultCode *string `json:"result_code,omitempty" xml:"result_code,omitempty"`
	// 异常信息的文本描述
	ResultMsg *string `json:"result_msg,omitempty" xml:"result_msg,omitempty"`
	// 协议链上地址
	TxHash *string `json:"tx_hash,omitempty" xml:"tx_hash,omitempty"`
}

func (s UploadCreditAuthorizationResponse) String() string {
	return tea.Prettify(s)
}

func (s UploadCreditAuthorizationResponse) GoString() string {
	return s.String()
}

func (s *UploadCreditAuthorizationResponse) SetReqMsgId(v string) *UploadCreditAuthorizationResponse {
	s.ReqMsgId = &v
	return s
}

func (s *UploadCreditAuthorizationResponse) SetResultCode(v string) *UploadCreditAuthorizationResponse {
	s.ResultCode = &v
	return s
}

func (s *UploadCreditAuthorizationResponse) SetResultMsg(v string) *UploadCreditAuthorizationResponse {
	s.ResultMsg = &v
	return s
}

func (s *UploadCreditAuthorizationResponse) SetTxHash(v string) *UploadCreditAuthorizationResponse {
	s.TxHash = &v
	return s
}

type UploadCreditConfirmRequest struct {
	// OAuth模式下的授权token
	AuthToken         *string `json:"auth_token,omitempty" xml:"auth_token,omitempty"`
	ProductInstanceId *string `json:"product_instance_id,omitempty" xml:"product_instance_id,omitempty"`
	// 货主云资金商户id
	AccountId *string `json:"account_id,omitempty" xml:"account_id,omitempty" require:"true"`
	// 确认时间
	ConfirmDate *string `json:"confirm_date,omitempty" xml:"confirm_date,omitempty" require:"true"`
	// 应收转让通知编号
	ConfirmNo *string `json:"confirm_no,omitempty" xml:"confirm_no,omitempty" require:"true"`
	// 签署货主did
	ConsignorDid *string `json:"consignor_did,omitempty" xml:"consignor_did,omitempty" require:"true"`
	// 签署的货主平台账号
	ConsignorPlatformAccount *string `json:"consignor_platform_account,omitempty" xml:"consignor_platform_account,omitempty" require:"true"`
	// 协议到期日
	ExpireDate *string `json:"expire_date,omitempty" xml:"expire_date,omitempty" require:"true"`
	// 集团平台did
	GroupPlatformDid *string `json:"group_platform_did,omitempty" xml:"group_platform_did,omitempty" require:"true"`
	// 签署子公司did
	PlatformDid *string `json:"platform_did,omitempty" xml:"platform_did,omitempty" require:"true"`
	// 协议名称
	ProtocolName *string `json:"protocol_name,omitempty" xml:"protocol_name,omitempty" require:"true"`
	// 协议PDF文件, 多个文件可用逗号分隔, 最多10个
	ProtocolPdfId *string `json:"protocol_pdf_id,omitempty" xml:"protocol_pdf_id,omitempty" require:"true"`
	// 产品id
	ProductId *string `json:"product_id,omitempty" xml:"product_id,omitempty" require:"true"`
}

func (s UploadCreditConfirmRequest) String() string {
	return tea.Prettify(s)
}

func (s UploadCreditConfirmRequest) GoString() string {
	return s.String()
}

func (s *UploadCreditConfirmRequest) SetAuthToken(v string) *UploadCreditConfirmRequest {
	s.AuthToken = &v
	return s
}

func (s *UploadCreditConfirmRequest) SetProductInstanceId(v string) *UploadCreditConfirmRequest {
	s.ProductInstanceId = &v
	return s
}

func (s *UploadCreditConfirmRequest) SetAccountId(v string) *UploadCreditConfirmRequest {
	s.AccountId = &v
	return s
}

func (s *UploadCreditConfirmRequest) SetConfirmDate(v string) *UploadCreditConfirmRequest {
	s.ConfirmDate = &v
	return s
}

func (s *UploadCreditConfirmRequest) SetConfirmNo(v string) *UploadCreditConfirmRequest {
	s.ConfirmNo = &v
	return s
}

func (s *UploadCreditConfirmRequest) SetConsignorDid(v string) *UploadCreditConfirmRequest {
	s.ConsignorDid = &v
	return s
}

func (s *UploadCreditConfirmRequest) SetConsignorPlatformAccount(v string) *UploadCreditConfirmRequest {
	s.ConsignorPlatformAccount = &v
	return s
}

func (s *UploadCreditConfirmRequest) SetExpireDate(v string) *UploadCreditConfirmRequest {
	s.ExpireDate = &v
	return s
}

func (s *UploadCreditConfirmRequest) SetGroupPlatformDid(v string) *UploadCreditConfirmRequest {
	s.GroupPlatformDid = &v
	return s
}

func (s *UploadCreditConfirmRequest) SetPlatformDid(v string) *UploadCreditConfirmRequest {
	s.PlatformDid = &v
	return s
}

func (s *UploadCreditConfirmRequest) SetProtocolName(v string) *UploadCreditConfirmRequest {
	s.ProtocolName = &v
	return s
}

func (s *UploadCreditConfirmRequest) SetProtocolPdfId(v string) *UploadCreditConfirmRequest {
	s.ProtocolPdfId = &v
	return s
}

func (s *UploadCreditConfirmRequest) SetProductId(v string) *UploadCreditConfirmRequest {
	s.ProductId = &v
	return s
}

type UploadCreditConfirmResponse struct {
	// 请求唯一ID，用于链路跟踪和问题排查
	ReqMsgId *string `json:"req_msg_id,omitempty" xml:"req_msg_id,omitempty"`
	// 结果码，一般OK表示调用成功
	ResultCode *string `json:"result_code,omitempty" xml:"result_code,omitempty"`
	// 异常信息的文本描述
	ResultMsg *string `json:"result_msg,omitempty" xml:"result_msg,omitempty"`
	// 协议链上地址
	TxHash *string `json:"tx_hash,omitempty" xml:"tx_hash,omitempty"`
}

func (s UploadCreditConfirmResponse) String() string {
	return tea.Prettify(s)
}

func (s UploadCreditConfirmResponse) GoString() string {
	return s.String()
}

func (s *UploadCreditConfirmResponse) SetReqMsgId(v string) *UploadCreditConfirmResponse {
	s.ReqMsgId = &v
	return s
}

func (s *UploadCreditConfirmResponse) SetResultCode(v string) *UploadCreditConfirmResponse {
	s.ResultCode = &v
	return s
}

func (s *UploadCreditConfirmResponse) SetResultMsg(v string) *UploadCreditConfirmResponse {
	s.ResultMsg = &v
	return s
}

func (s *UploadCreditConfirmResponse) SetTxHash(v string) *UploadCreditConfirmResponse {
	s.TxHash = &v
	return s
}

type BatchcreateCreditmodeIssueRequest struct {
	// OAuth模式下的授权token
	AuthToken         *string `json:"auth_token,omitempty" xml:"auth_token,omitempty"`
	ProductInstanceId *string `json:"product_instance_id,omitempty" xml:"product_instance_id,omitempty"`
	// 批次id
	BatchId *string `json:"batch_id,omitempty" xml:"batch_id,omitempty" require:"true"`
	// 回调地址
	CallbackUrl *string `json:"callback_url,omitempty" xml:"callback_url,omitempty" require:"true"`
	// 集团平台分布式数字身份
	GroupPlatformDid *string `json:"group_platform_did,omitempty" xml:"group_platform_did,omitempty" require:"true"`
	// 发行列表
	IssueApplyInfos []*IssueApplyInfo `json:"issue_apply_infos,omitempty" xml:"issue_apply_infos,omitempty" require:"true" type:"Repeated"`
	// 业务平台方分布式数字身份
	PlatformDid *string `json:"platform_did,omitempty" xml:"platform_did,omitempty" require:"true"`
	// 产品id
	ProductId *string `json:"product_id,omitempty" xml:"product_id,omitempty" require:"true"`
}

func (s BatchcreateCreditmodeIssueRequest) String() string {
	return tea.Prettify(s)
}

func (s BatchcreateCreditmodeIssueRequest) GoString() string {
	return s.String()
}

func (s *BatchcreateCreditmodeIssueRequest) SetAuthToken(v string) *BatchcreateCreditmodeIssueRequest {
	s.AuthToken = &v
	return s
}

func (s *BatchcreateCreditmodeIssueRequest) SetProductInstanceId(v string) *BatchcreateCreditmodeIssueRequest {
	s.ProductInstanceId = &v
	return s
}

func (s *BatchcreateCreditmodeIssueRequest) SetBatchId(v string) *BatchcreateCreditmodeIssueRequest {
	s.BatchId = &v
	return s
}

func (s *BatchcreateCreditmodeIssueRequest) SetCallbackUrl(v string) *BatchcreateCreditmodeIssueRequest {
	s.CallbackUrl = &v
	return s
}

func (s *BatchcreateCreditmodeIssueRequest) SetGroupPlatformDid(v string) *BatchcreateCreditmodeIssueRequest {
	s.GroupPlatformDid = &v
	return s
}

func (s *BatchcreateCreditmodeIssueRequest) SetIssueApplyInfos(v []*IssueApplyInfo) *BatchcreateCreditmodeIssueRequest {
	s.IssueApplyInfos = v
	return s
}

func (s *BatchcreateCreditmodeIssueRequest) SetPlatformDid(v string) *BatchcreateCreditmodeIssueRequest {
	s.PlatformDid = &v
	return s
}

func (s *BatchcreateCreditmodeIssueRequest) SetProductId(v string) *BatchcreateCreditmodeIssueRequest {
	s.ProductId = &v
	return s
}

type BatchcreateCreditmodeIssueResponse struct {
	// 请求唯一ID，用于链路跟踪和问题排查
	ReqMsgId *string `json:"req_msg_id,omitempty" xml:"req_msg_id,omitempty"`
	// 结果码，一般OK表示调用成功
	ResultCode *string `json:"result_code,omitempty" xml:"result_code,omitempty"`
	// 异常信息的文本描述
	ResultMsg *string `json:"result_msg,omitempty" xml:"result_msg,omitempty"`
	// 批次id
	BatchId *string `json:"batch_id,omitempty" xml:"batch_id,omitempty"`
	// 批次状态
	BatchStatus *string `json:"batch_status,omitempty" xml:"batch_status,omitempty"`
}

func (s BatchcreateCreditmodeIssueResponse) String() string {
	return tea.Prettify(s)
}

func (s BatchcreateCreditmodeIssueResponse) GoString() string {
	return s.String()
}

func (s *BatchcreateCreditmodeIssueResponse) SetReqMsgId(v string) *BatchcreateCreditmodeIssueResponse {
	s.ReqMsgId = &v
	return s
}

func (s *BatchcreateCreditmodeIssueResponse) SetResultCode(v string) *BatchcreateCreditmodeIssueResponse {
	s.ResultCode = &v
	return s
}

func (s *BatchcreateCreditmodeIssueResponse) SetResultMsg(v string) *BatchcreateCreditmodeIssueResponse {
	s.ResultMsg = &v
	return s
}

func (s *BatchcreateCreditmodeIssueResponse) SetBatchId(v string) *BatchcreateCreditmodeIssueResponse {
	s.BatchId = &v
	return s
}

func (s *BatchcreateCreditmodeIssueResponse) SetBatchStatus(v string) *BatchcreateCreditmodeIssueResponse {
	s.BatchStatus = &v
	return s
}

type ApplyCreditmodeIssueclearRequest struct {
	// OAuth模式下的授权token
	AuthToken         *string `json:"auth_token,omitempty" xml:"auth_token,omitempty"`
	ProductInstanceId *string `json:"product_instance_id,omitempty" xml:"product_instance_id,omitempty"`
	// 产品id
	ProductId *string `json:"product_id,omitempty" xml:"product_id,omitempty" require:"true"`
	// 集团平台方分布式数字身份
	GroupPlatformDid *string `json:"group_platform_did,omitempty" xml:"group_platform_did,omitempty" require:"true"`
	// 业务平台方分布式数字身份
	PlatformDid *string `json:"platform_did,omitempty" xml:"platform_did,omitempty" require:"true"`
	// 主动清分方分布式数字身份
	ClearDid *string `json:"clear_did,omitempty" xml:"clear_did,omitempty" require:"true"`
	// 凭证id
	IssueId *string `json:"issue_id,omitempty" xml:"issue_id,omitempty" require:"true"`
	// 全局业务号
	OutBizNo *string `json:"out_biz_no,omitempty" xml:"out_biz_no,omitempty" require:"true"`
	// 支付单号
	OutOrderNo *string `json:"out_order_no,omitempty" xml:"out_order_no,omitempty" require:"true"`
	// 清分金额
	ClearAmt *string `json:"clear_amt,omitempty" xml:"clear_amt,omitempty"`
	// 模式  B:b模式
	ModeType *string `json:"mode_type,omitempty" xml:"mode_type,omitempty" require:"true"`
	// 申请日期
	ApplyDate *string `json:"apply_date,omitempty" xml:"apply_date,omitempty"`
	// 扩展字段
	ExtInfo *string `json:"ext_info,omitempty" xml:"ext_info,omitempty"`
}

func (s ApplyCreditmodeIssueclearRequest) String() string {
	return tea.Prettify(s)
}

func (s ApplyCreditmodeIssueclearRequest) GoString() string {
	return s.String()
}

func (s *ApplyCreditmodeIssueclearRequest) SetAuthToken(v string) *ApplyCreditmodeIssueclearRequest {
	s.AuthToken = &v
	return s
}

func (s *ApplyCreditmodeIssueclearRequest) SetProductInstanceId(v string) *ApplyCreditmodeIssueclearRequest {
	s.ProductInstanceId = &v
	return s
}

func (s *ApplyCreditmodeIssueclearRequest) SetProductId(v string) *ApplyCreditmodeIssueclearRequest {
	s.ProductId = &v
	return s
}

func (s *ApplyCreditmodeIssueclearRequest) SetGroupPlatformDid(v string) *ApplyCreditmodeIssueclearRequest {
	s.GroupPlatformDid = &v
	return s
}

func (s *ApplyCreditmodeIssueclearRequest) SetPlatformDid(v string) *ApplyCreditmodeIssueclearRequest {
	s.PlatformDid = &v
	return s
}

func (s *ApplyCreditmodeIssueclearRequest) SetClearDid(v string) *ApplyCreditmodeIssueclearRequest {
	s.ClearDid = &v
	return s
}

func (s *ApplyCreditmodeIssueclearRequest) SetIssueId(v string) *ApplyCreditmodeIssueclearRequest {
	s.IssueId = &v
	return s
}

func (s *ApplyCreditmodeIssueclearRequest) SetOutBizNo(v string) *ApplyCreditmodeIssueclearRequest {
	s.OutBizNo = &v
	return s
}

func (s *ApplyCreditmodeIssueclearRequest) SetOutOrderNo(v string) *ApplyCreditmodeIssueclearRequest {
	s.OutOrderNo = &v
	return s
}

func (s *ApplyCreditmodeIssueclearRequest) SetClearAmt(v string) *ApplyCreditmodeIssueclearRequest {
	s.ClearAmt = &v
	return s
}

func (s *ApplyCreditmodeIssueclearRequest) SetModeType(v string) *ApplyCreditmodeIssueclearRequest {
	s.ModeType = &v
	return s
}

func (s *ApplyCreditmodeIssueclearRequest) SetApplyDate(v string) *ApplyCreditmodeIssueclearRequest {
	s.ApplyDate = &v
	return s
}

func (s *ApplyCreditmodeIssueclearRequest) SetExtInfo(v string) *ApplyCreditmodeIssueclearRequest {
	s.ExtInfo = &v
	return s
}

type ApplyCreditmodeIssueclearResponse struct {
	// 请求唯一ID，用于链路跟踪和问题排查
	ReqMsgId *string `json:"req_msg_id,omitempty" xml:"req_msg_id,omitempty"`
	// 结果码，一般OK表示调用成功
	ResultCode *string `json:"result_code,omitempty" xml:"result_code,omitempty"`
	// 异常信息的文本描述
	ResultMsg *string `json:"result_msg,omitempty" xml:"result_msg,omitempty"`
	// 全局业务号
	OutBizNo *string `json:"out_biz_no,omitempty" xml:"out_biz_no,omitempty"`
}

func (s ApplyCreditmodeIssueclearResponse) String() string {
	return tea.Prettify(s)
}

func (s ApplyCreditmodeIssueclearResponse) GoString() string {
	return s.String()
}

func (s *ApplyCreditmodeIssueclearResponse) SetReqMsgId(v string) *ApplyCreditmodeIssueclearResponse {
	s.ReqMsgId = &v
	return s
}

func (s *ApplyCreditmodeIssueclearResponse) SetResultCode(v string) *ApplyCreditmodeIssueclearResponse {
	s.ResultCode = &v
	return s
}

func (s *ApplyCreditmodeIssueclearResponse) SetResultMsg(v string) *ApplyCreditmodeIssueclearResponse {
	s.ResultMsg = &v
	return s
}

func (s *ApplyCreditmodeIssueclearResponse) SetOutBizNo(v string) *ApplyCreditmodeIssueclearResponse {
	s.OutBizNo = &v
	return s
}

type UploadCreditIssuebysaasRequest struct {
	// OAuth模式下的授权token
	AuthToken         *string `json:"auth_token,omitempty" xml:"auth_token,omitempty"`
	ProductInstanceId *string `json:"product_instance_id,omitempty" xml:"product_instance_id,omitempty"`
	// 批次id
	BatchId *string `json:"batch_id,omitempty" xml:"batch_id,omitempty" require:"true"`
	// 发行结果回调地址
	CallbackUrl *string `json:"callback_url,omitempty" xml:"callback_url,omitempty" require:"true"`
	// 集团平台did
	GroupPlatformDid *string `json:"group_platform_did,omitempty" xml:"group_platform_did,omitempty" require:"true"`
	// 发行信息
	IssueApplyInfos []*SaasIssueApplyInfo `json:"issue_apply_infos,omitempty" xml:"issue_apply_infos,omitempty" require:"true" type:"Repeated"`
	// 凭证发行方did
	PlatformDid *string `json:"platform_did,omitempty" xml:"platform_did,omitempty" require:"true"`
	// 产品id
	ProductId *string `json:"product_id,omitempty" xml:"product_id,omitempty" require:"true"`
}

func (s UploadCreditIssuebysaasRequest) String() string {
	return tea.Prettify(s)
}

func (s UploadCreditIssuebysaasRequest) GoString() string {
	return s.String()
}

func (s *UploadCreditIssuebysaasRequest) SetAuthToken(v string) *UploadCreditIssuebysaasRequest {
	s.AuthToken = &v
	return s
}

func (s *UploadCreditIssuebysaasRequest) SetProductInstanceId(v string) *UploadCreditIssuebysaasRequest {
	s.ProductInstanceId = &v
	return s
}

func (s *UploadCreditIssuebysaasRequest) SetBatchId(v string) *UploadCreditIssuebysaasRequest {
	s.BatchId = &v
	return s
}

func (s *UploadCreditIssuebysaasRequest) SetCallbackUrl(v string) *UploadCreditIssuebysaasRequest {
	s.CallbackUrl = &v
	return s
}

func (s *UploadCreditIssuebysaasRequest) SetGroupPlatformDid(v string) *UploadCreditIssuebysaasRequest {
	s.GroupPlatformDid = &v
	return s
}

func (s *UploadCreditIssuebysaasRequest) SetIssueApplyInfos(v []*SaasIssueApplyInfo) *UploadCreditIssuebysaasRequest {
	s.IssueApplyInfos = v
	return s
}

func (s *UploadCreditIssuebysaasRequest) SetPlatformDid(v string) *UploadCreditIssuebysaasRequest {
	s.PlatformDid = &v
	return s
}

func (s *UploadCreditIssuebysaasRequest) SetProductId(v string) *UploadCreditIssuebysaasRequest {
	s.ProductId = &v
	return s
}

type UploadCreditIssuebysaasResponse struct {
	// 请求唯一ID，用于链路跟踪和问题排查
	ReqMsgId *string `json:"req_msg_id,omitempty" xml:"req_msg_id,omitempty"`
	// 结果码，一般OK表示调用成功
	ResultCode *string `json:"result_code,omitempty" xml:"result_code,omitempty"`
	// 异常信息的文本描述
	ResultMsg *string `json:"result_msg,omitempty" xml:"result_msg,omitempty"`
	// 批次id
	BatchId *string `json:"batch_id,omitempty" xml:"batch_id,omitempty"`
	// 批次状态
	BatchStatus *int64 `json:"batch_status,omitempty" xml:"batch_status,omitempty"`
	// 发行url
	IssueUrl *string `json:"issue_url,omitempty" xml:"issue_url,omitempty"`
}

func (s UploadCreditIssuebysaasResponse) String() string {
	return tea.Prettify(s)
}

func (s UploadCreditIssuebysaasResponse) GoString() string {
	return s.String()
}

func (s *UploadCreditIssuebysaasResponse) SetReqMsgId(v string) *UploadCreditIssuebysaasResponse {
	s.ReqMsgId = &v
	return s
}

func (s *UploadCreditIssuebysaasResponse) SetResultCode(v string) *UploadCreditIssuebysaasResponse {
	s.ResultCode = &v
	return s
}

func (s *UploadCreditIssuebysaasResponse) SetResultMsg(v string) *UploadCreditIssuebysaasResponse {
	s.ResultMsg = &v
	return s
}

func (s *UploadCreditIssuebysaasResponse) SetBatchId(v string) *UploadCreditIssuebysaasResponse {
	s.BatchId = &v
	return s
}

func (s *UploadCreditIssuebysaasResponse) SetBatchStatus(v int64) *UploadCreditIssuebysaasResponse {
	s.BatchStatus = &v
	return s
}

func (s *UploadCreditIssuebysaasResponse) SetIssueUrl(v string) *UploadCreditIssuebysaasResponse {
	s.IssueUrl = &v
	return s
}

type CancelCreditIssuebatchbysaasRequest struct {
	// OAuth模式下的授权token
	AuthToken         *string `json:"auth_token,omitempty" xml:"auth_token,omitempty"`
	ProductInstanceId *string `json:"product_instance_id,omitempty" xml:"product_instance_id,omitempty"`
	// 批次号
	BatchId *string `json:"batch_id,omitempty" xml:"batch_id,omitempty" require:"true"`
	// 集团平台did
	GroupPlatformDid *string `json:"group_platform_did,omitempty" xml:"group_platform_did,omitempty" require:"true"`
	// 凭证发行方did
	PlatformDid *string `json:"platform_did,omitempty" xml:"platform_did,omitempty" require:"true"`
	// 产品id
	ProductId *string `json:"product_id,omitempty" xml:"product_id,omitempty" require:"true"`
}

func (s CancelCreditIssuebatchbysaasRequest) String() string {
	return tea.Prettify(s)
}

func (s CancelCreditIssuebatchbysaasRequest) GoString() string {
	return s.String()
}

func (s *CancelCreditIssuebatchbysaasRequest) SetAuthToken(v string) *CancelCreditIssuebatchbysaasRequest {
	s.AuthToken = &v
	return s
}

func (s *CancelCreditIssuebatchbysaasRequest) SetProductInstanceId(v string) *CancelCreditIssuebatchbysaasRequest {
	s.ProductInstanceId = &v
	return s
}

func (s *CancelCreditIssuebatchbysaasRequest) SetBatchId(v string) *CancelCreditIssuebatchbysaasRequest {
	s.BatchId = &v
	return s
}

func (s *CancelCreditIssuebatchbysaasRequest) SetGroupPlatformDid(v string) *CancelCreditIssuebatchbysaasRequest {
	s.GroupPlatformDid = &v
	return s
}

func (s *CancelCreditIssuebatchbysaasRequest) SetPlatformDid(v string) *CancelCreditIssuebatchbysaasRequest {
	s.PlatformDid = &v
	return s
}

func (s *CancelCreditIssuebatchbysaasRequest) SetProductId(v string) *CancelCreditIssuebatchbysaasRequest {
	s.ProductId = &v
	return s
}

type CancelCreditIssuebatchbysaasResponse struct {
	// 请求唯一ID，用于链路跟踪和问题排查
	ReqMsgId *string `json:"req_msg_id,omitempty" xml:"req_msg_id,omitempty"`
	// 结果码，一般OK表示调用成功
	ResultCode *string `json:"result_code,omitempty" xml:"result_code,omitempty"`
	// 异常信息的文本描述
	ResultMsg *string `json:"result_msg,omitempty" xml:"result_msg,omitempty"`
	// 批次id
	BatchId *string `json:"batch_id,omitempty" xml:"batch_id,omitempty"`
}

func (s CancelCreditIssuebatchbysaasResponse) String() string {
	return tea.Prettify(s)
}

func (s CancelCreditIssuebatchbysaasResponse) GoString() string {
	return s.String()
}

func (s *CancelCreditIssuebatchbysaasResponse) SetReqMsgId(v string) *CancelCreditIssuebatchbysaasResponse {
	s.ReqMsgId = &v
	return s
}

func (s *CancelCreditIssuebatchbysaasResponse) SetResultCode(v string) *CancelCreditIssuebatchbysaasResponse {
	s.ResultCode = &v
	return s
}

func (s *CancelCreditIssuebatchbysaasResponse) SetResultMsg(v string) *CancelCreditIssuebatchbysaasResponse {
	s.ResultMsg = &v
	return s
}

func (s *CancelCreditIssuebatchbysaasResponse) SetBatchId(v string) *CancelCreditIssuebatchbysaasResponse {
	s.BatchId = &v
	return s
}

type QueryCreditBalancebysaasRequest struct {
	// OAuth模式下的授权token
	AuthToken         *string `json:"auth_token,omitempty" xml:"auth_token,omitempty"`
	ProductInstanceId *string `json:"product_instance_id,omitempty" xml:"product_instance_id,omitempty"`
	// 产品id
	ProductId *string `json:"product_id,omitempty" xml:"product_id,omitempty" require:"true"`
	// 集团平台did
	GroupPlatformDid *string `json:"group_platform_did,omitempty" xml:"group_platform_did,omitempty" require:"true"`
	// 查询者did
	PlatformDid *string `json:"platform_did,omitempty" xml:"platform_did,omitempty" require:"true"`
	// 发行金额
	Freight *string `json:"freight,omitempty" xml:"freight,omitempty" require:"true"`
}

func (s QueryCreditBalancebysaasRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryCreditBalancebysaasRequest) GoString() string {
	return s.String()
}

func (s *QueryCreditBalancebysaasRequest) SetAuthToken(v string) *QueryCreditBalancebysaasRequest {
	s.AuthToken = &v
	return s
}

func (s *QueryCreditBalancebysaasRequest) SetProductInstanceId(v string) *QueryCreditBalancebysaasRequest {
	s.ProductInstanceId = &v
	return s
}

func (s *QueryCreditBalancebysaasRequest) SetProductId(v string) *QueryCreditBalancebysaasRequest {
	s.ProductId = &v
	return s
}

func (s *QueryCreditBalancebysaasRequest) SetGroupPlatformDid(v string) *QueryCreditBalancebysaasRequest {
	s.GroupPlatformDid = &v
	return s
}

func (s *QueryCreditBalancebysaasRequest) SetPlatformDid(v string) *QueryCreditBalancebysaasRequest {
	s.PlatformDid = &v
	return s
}

func (s *QueryCreditBalancebysaasRequest) SetFreight(v string) *QueryCreditBalancebysaasRequest {
	s.Freight = &v
	return s
}

type QueryCreditBalancebysaasResponse struct {
	// 请求唯一ID，用于链路跟踪和问题排查
	ReqMsgId *string `json:"req_msg_id,omitempty" xml:"req_msg_id,omitempty"`
	// 结果码，一般OK表示调用成功
	ResultCode *string `json:"result_code,omitempty" xml:"result_code,omitempty"`
	// 异常信息的文本描述
	ResultMsg *string `json:"result_msg,omitempty" xml:"result_msg,omitempty"`
	// 是否有信用流转额度
	HasBalance *bool `json:"has_balance,omitempty" xml:"has_balance,omitempty"`
}

func (s QueryCreditBalancebysaasResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryCreditBalancebysaasResponse) GoString() string {
	return s.String()
}

func (s *QueryCreditBalancebysaasResponse) SetReqMsgId(v string) *QueryCreditBalancebysaasResponse {
	s.ReqMsgId = &v
	return s
}

func (s *QueryCreditBalancebysaasResponse) SetResultCode(v string) *QueryCreditBalancebysaasResponse {
	s.ResultCode = &v
	return s
}

func (s *QueryCreditBalancebysaasResponse) SetResultMsg(v string) *QueryCreditBalancebysaasResponse {
	s.ResultMsg = &v
	return s
}

func (s *QueryCreditBalancebysaasResponse) SetHasBalance(v bool) *QueryCreditBalancebysaasResponse {
	s.HasBalance = &v
	return s
}

type GetCreditIssuescpticketRequest struct {
	// OAuth模式下的授权token
	AuthToken         *string `json:"auth_token,omitempty" xml:"auth_token,omitempty"`
	ProductInstanceId *string `json:"product_instance_id,omitempty" xml:"product_instance_id,omitempty"`
	// 申请唯一流水号
	ApplyId *string `json:"apply_id,omitempty" xml:"apply_id,omitempty" require:"true"`
	// 回调地址
	CallbackUrl *string `json:"callback_url,omitempty" xml:"callback_url,omitempty" require:"true"`
	// 集团平台方分布式数字身份
	GroupPlatformDid *string `json:"group_platform_did,omitempty" xml:"group_platform_did,omitempty" require:"true"`
	// 业务发起方分布式数字身
	PlatformDid *string `json:"platform_did,omitempty" xml:"platform_did,omitempty" require:"true"`
	// 产品id
	ProductId *string `json:"product_id,omitempty" xml:"product_id,omitempty" require:"true"`
	// 电子回单查询凭证数据
	ScpTicketIssueData []*ScpTicketIssueData `json:"scp_ticket_issue_data,omitempty" xml:"scp_ticket_issue_data,omitempty" require:"true" type:"Repeated"`
}

func (s GetCreditIssuescpticketRequest) String() string {
	return tea.Prettify(s)
}

func (s GetCreditIssuescpticketRequest) GoString() string {
	return s.String()
}

func (s *GetCreditIssuescpticketRequest) SetAuthToken(v string) *GetCreditIssuescpticketRequest {
	s.AuthToken = &v
	return s
}

func (s *GetCreditIssuescpticketRequest) SetProductInstanceId(v string) *GetCreditIssuescpticketRequest {
	s.ProductInstanceId = &v
	return s
}

func (s *GetCreditIssuescpticketRequest) SetApplyId(v string) *GetCreditIssuescpticketRequest {
	s.ApplyId = &v
	return s
}

func (s *GetCreditIssuescpticketRequest) SetCallbackUrl(v string) *GetCreditIssuescpticketRequest {
	s.CallbackUrl = &v
	return s
}

func (s *GetCreditIssuescpticketRequest) SetGroupPlatformDid(v string) *GetCreditIssuescpticketRequest {
	s.GroupPlatformDid = &v
	return s
}

func (s *GetCreditIssuescpticketRequest) SetPlatformDid(v string) *GetCreditIssuescpticketRequest {
	s.PlatformDid = &v
	return s
}

func (s *GetCreditIssuescpticketRequest) SetProductId(v string) *GetCreditIssuescpticketRequest {
	s.ProductId = &v
	return s
}

func (s *GetCreditIssuescpticketRequest) SetScpTicketIssueData(v []*ScpTicketIssueData) *GetCreditIssuescpticketRequest {
	s.ScpTicketIssueData = v
	return s
}

type GetCreditIssuescpticketResponse struct {
	// 请求唯一ID，用于链路跟踪和问题排查
	ReqMsgId *string `json:"req_msg_id,omitempty" xml:"req_msg_id,omitempty"`
	// 结果码，一般OK表示调用成功
	ResultCode *string `json:"result_code,omitempty" xml:"result_code,omitempty"`
	// 异常信息的文本描述
	ResultMsg *string `json:"result_msg,omitempty" xml:"result_msg,omitempty"`
	// 申请唯一流水号
	ApplyId *string `json:"apply_id,omitempty" xml:"apply_id,omitempty"`
	// 集团平台方分布式数字身份
	GroupPlatformDid *string `json:"group_platform_did,omitempty" xml:"group_platform_did,omitempty"`
	// 业务发起方分布式数字身
	PlatformDid *string `json:"platform_did,omitempty" xml:"platform_did,omitempty"`
	// 产品id
	ProductId *string `json:"product_id,omitempty" xml:"product_id,omitempty"`
}

func (s GetCreditIssuescpticketResponse) String() string {
	return tea.Prettify(s)
}

func (s GetCreditIssuescpticketResponse) GoString() string {
	return s.String()
}

func (s *GetCreditIssuescpticketResponse) SetReqMsgId(v string) *GetCreditIssuescpticketResponse {
	s.ReqMsgId = &v
	return s
}

func (s *GetCreditIssuescpticketResponse) SetResultCode(v string) *GetCreditIssuescpticketResponse {
	s.ResultCode = &v
	return s
}

func (s *GetCreditIssuescpticketResponse) SetResultMsg(v string) *GetCreditIssuescpticketResponse {
	s.ResultMsg = &v
	return s
}

func (s *GetCreditIssuescpticketResponse) SetApplyId(v string) *GetCreditIssuescpticketResponse {
	s.ApplyId = &v
	return s
}

func (s *GetCreditIssuescpticketResponse) SetGroupPlatformDid(v string) *GetCreditIssuescpticketResponse {
	s.GroupPlatformDid = &v
	return s
}

func (s *GetCreditIssuescpticketResponse) SetPlatformDid(v string) *GetCreditIssuescpticketResponse {
	s.PlatformDid = &v
	return s
}

func (s *GetCreditIssuescpticketResponse) SetProductId(v string) *GetCreditIssuescpticketResponse {
	s.ProductId = &v
	return s
}

type QueryCreditIssuescpticketresultRequest struct {
	// OAuth模式下的授权token
	AuthToken         *string `json:"auth_token,omitempty" xml:"auth_token,omitempty"`
	ProductInstanceId *string `json:"product_instance_id,omitempty" xml:"product_instance_id,omitempty"`
	// 申请唯一流水号
	ApplyId *string `json:"apply_id,omitempty" xml:"apply_id,omitempty" require:"true"`
	// 集团平台方分布式数字身份
	GroupPlatformDid *string `json:"group_platform_did,omitempty" xml:"group_platform_did,omitempty" require:"true"`
	// 业务发起方分布式数字身
	PlatformDid *string `json:"platform_did,omitempty" xml:"platform_did,omitempty" require:"true"`
	// 产品id
	ProductId *string `json:"product_id,omitempty" xml:"product_id,omitempty" require:"true"`
}

func (s QueryCreditIssuescpticketresultRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryCreditIssuescpticketresultRequest) GoString() string {
	return s.String()
}

func (s *QueryCreditIssuescpticketresultRequest) SetAuthToken(v string) *QueryCreditIssuescpticketresultRequest {
	s.AuthToken = &v
	return s
}

func (s *QueryCreditIssuescpticketresultRequest) SetProductInstanceId(v string) *QueryCreditIssuescpticketresultRequest {
	s.ProductInstanceId = &v
	return s
}

func (s *QueryCreditIssuescpticketresultRequest) SetApplyId(v string) *QueryCreditIssuescpticketresultRequest {
	s.ApplyId = &v
	return s
}

func (s *QueryCreditIssuescpticketresultRequest) SetGroupPlatformDid(v string) *QueryCreditIssuescpticketresultRequest {
	s.GroupPlatformDid = &v
	return s
}

func (s *QueryCreditIssuescpticketresultRequest) SetPlatformDid(v string) *QueryCreditIssuescpticketresultRequest {
	s.PlatformDid = &v
	return s
}

func (s *QueryCreditIssuescpticketresultRequest) SetProductId(v string) *QueryCreditIssuescpticketresultRequest {
	s.ProductId = &v
	return s
}

type QueryCreditIssuescpticketresultResponse struct {
	// 请求唯一ID，用于链路跟踪和问题排查
	ReqMsgId *string `json:"req_msg_id,omitempty" xml:"req_msg_id,omitempty"`
	// 结果码，一般OK表示调用成功
	ResultCode *string `json:"result_code,omitempty" xml:"result_code,omitempty"`
	// 异常信息的文本描述
	ResultMsg *string `json:"result_msg,omitempty" xml:"result_msg,omitempty"`
	// 申请唯一流水号
	ApplyId *string `json:"apply_id,omitempty" xml:"apply_id,omitempty"`
	// 电子回单下载地址,链接有有效期
	FileDownloadUrl *string `json:"file_download_url,omitempty" xml:"file_download_url,omitempty"`
	// 集团平台方分布式数字身份
	//
	GroupPlatformDid *string `json:"group_platform_did,omitempty" xml:"group_platform_did,omitempty"`
	// 业务发起方分布式数字身
	PlatformDid *string `json:"platform_did,omitempty" xml:"platform_did,omitempty"`
	// 产品id
	//
	ProductId *string `json:"product_id,omitempty" xml:"product_id,omitempty"`
	// 状态
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
}

func (s QueryCreditIssuescpticketresultResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryCreditIssuescpticketresultResponse) GoString() string {
	return s.String()
}

func (s *QueryCreditIssuescpticketresultResponse) SetReqMsgId(v string) *QueryCreditIssuescpticketresultResponse {
	s.ReqMsgId = &v
	return s
}

func (s *QueryCreditIssuescpticketresultResponse) SetResultCode(v string) *QueryCreditIssuescpticketresultResponse {
	s.ResultCode = &v
	return s
}

func (s *QueryCreditIssuescpticketresultResponse) SetResultMsg(v string) *QueryCreditIssuescpticketresultResponse {
	s.ResultMsg = &v
	return s
}

func (s *QueryCreditIssuescpticketresultResponse) SetApplyId(v string) *QueryCreditIssuescpticketresultResponse {
	s.ApplyId = &v
	return s
}

func (s *QueryCreditIssuescpticketresultResponse) SetFileDownloadUrl(v string) *QueryCreditIssuescpticketresultResponse {
	s.FileDownloadUrl = &v
	return s
}

func (s *QueryCreditIssuescpticketresultResponse) SetGroupPlatformDid(v string) *QueryCreditIssuescpticketresultResponse {
	s.GroupPlatformDid = &v
	return s
}

func (s *QueryCreditIssuescpticketresultResponse) SetPlatformDid(v string) *QueryCreditIssuescpticketresultResponse {
	s.PlatformDid = &v
	return s
}

func (s *QueryCreditIssuescpticketresultResponse) SetProductId(v string) *QueryCreditIssuescpticketresultResponse {
	s.ProductId = &v
	return s
}

func (s *QueryCreditIssuescpticketresultResponse) SetStatus(v string) *QueryCreditIssuescpticketresultResponse {
	s.Status = &v
	return s
}

type UploadCreditAplusissueRequest struct {
	// OAuth模式下的授权token
	AuthToken         *string `json:"auth_token,omitempty" xml:"auth_token,omitempty"`
	ProductInstanceId *string `json:"product_instance_id,omitempty" xml:"product_instance_id,omitempty"`
	// 批次id
	BatchId *string `json:"batch_id,omitempty" xml:"batch_id,omitempty" require:"true"`
	// 发行结果信息回调地址
	CallbackUrl *string `json:"callback_url,omitempty" xml:"callback_url,omitempty" require:"true"`
	// 货代did
	ConsignorDid *string `json:"consignor_did,omitempty" xml:"consignor_did,omitempty" require:"true"`
	// 发行时直接进行快速提现：true ，
	// 发行后自主触发提现：false （默认）
	EasyFinance *bool `json:"easy_finance,omitempty" xml:"easy_finance,omitempty"`
	// 集团平台did
	GroupPlatformDid *string `json:"group_platform_did,omitempty" xml:"group_platform_did,omitempty" require:"true"`
	// 发行信息列表
	IssueApplyInfos []*IssueApplyInfoPlus `json:"issue_apply_infos,omitempty" xml:"issue_apply_infos,omitempty" require:"true" type:"Repeated"`
	// 业务承接方did
	PlatformDid *string `json:"platform_did,omitempty" xml:"platform_did,omitempty" require:"true"`
	// 产品id
	ProductId *string `json:"product_id,omitempty" xml:"product_id,omitempty" require:"true"`
	// 登录模式，返回的发行url指定使用何种模式登录。ORG:企业支付宝登录（缺省时默认ORG模式），USER_ACCOUNT:个人支付宝登录。
	LoginType *string `json:"login_type,omitempty" xml:"login_type,omitempty"`
}

func (s UploadCreditAplusissueRequest) String() string {
	return tea.Prettify(s)
}

func (s UploadCreditAplusissueRequest) GoString() string {
	return s.String()
}

func (s *UploadCreditAplusissueRequest) SetAuthToken(v string) *UploadCreditAplusissueRequest {
	s.AuthToken = &v
	return s
}

func (s *UploadCreditAplusissueRequest) SetProductInstanceId(v string) *UploadCreditAplusissueRequest {
	s.ProductInstanceId = &v
	return s
}

func (s *UploadCreditAplusissueRequest) SetBatchId(v string) *UploadCreditAplusissueRequest {
	s.BatchId = &v
	return s
}

func (s *UploadCreditAplusissueRequest) SetCallbackUrl(v string) *UploadCreditAplusissueRequest {
	s.CallbackUrl = &v
	return s
}

func (s *UploadCreditAplusissueRequest) SetConsignorDid(v string) *UploadCreditAplusissueRequest {
	s.ConsignorDid = &v
	return s
}

func (s *UploadCreditAplusissueRequest) SetEasyFinance(v bool) *UploadCreditAplusissueRequest {
	s.EasyFinance = &v
	return s
}

func (s *UploadCreditAplusissueRequest) SetGroupPlatformDid(v string) *UploadCreditAplusissueRequest {
	s.GroupPlatformDid = &v
	return s
}

func (s *UploadCreditAplusissueRequest) SetIssueApplyInfos(v []*IssueApplyInfoPlus) *UploadCreditAplusissueRequest {
	s.IssueApplyInfos = v
	return s
}

func (s *UploadCreditAplusissueRequest) SetPlatformDid(v string) *UploadCreditAplusissueRequest {
	s.PlatformDid = &v
	return s
}

func (s *UploadCreditAplusissueRequest) SetProductId(v string) *UploadCreditAplusissueRequest {
	s.ProductId = &v
	return s
}

func (s *UploadCreditAplusissueRequest) SetLoginType(v string) *UploadCreditAplusissueRequest {
	s.LoginType = &v
	return s
}

type UploadCreditAplusissueResponse struct {
	// 请求唯一ID，用于链路跟踪和问题排查
	ReqMsgId *string `json:"req_msg_id,omitempty" xml:"req_msg_id,omitempty"`
	// 结果码，一般OK表示调用成功
	ResultCode *string `json:"result_code,omitempty" xml:"result_code,omitempty"`
	// 异常信息的文本描述
	ResultMsg *string `json:"result_msg,omitempty" xml:"result_msg,omitempty"`
	// 批次id
	BatchId *string `json:"batch_id,omitempty" xml:"batch_id,omitempty"`
	// 批次状态 -1:失效， 0:初始化完成，1:进行中，2:部分成功，3:成功
	BatchIdStatus *int64 `json:"batch_id_status,omitempty" xml:"batch_id_status,omitempty"`
	// 凭证发行跳转url
	IssueUrl *string `json:"issue_url,omitempty" xml:"issue_url,omitempty"`
}

func (s UploadCreditAplusissueResponse) String() string {
	return tea.Prettify(s)
}

func (s UploadCreditAplusissueResponse) GoString() string {
	return s.String()
}

func (s *UploadCreditAplusissueResponse) SetReqMsgId(v string) *UploadCreditAplusissueResponse {
	s.ReqMsgId = &v
	return s
}

func (s *UploadCreditAplusissueResponse) SetResultCode(v string) *UploadCreditAplusissueResponse {
	s.ResultCode = &v
	return s
}

func (s *UploadCreditAplusissueResponse) SetResultMsg(v string) *UploadCreditAplusissueResponse {
	s.ResultMsg = &v
	return s
}

func (s *UploadCreditAplusissueResponse) SetBatchId(v string) *UploadCreditAplusissueResponse {
	s.BatchId = &v
	return s
}

func (s *UploadCreditAplusissueResponse) SetBatchIdStatus(v int64) *UploadCreditAplusissueResponse {
	s.BatchIdStatus = &v
	return s
}

func (s *UploadCreditAplusissueResponse) SetIssueUrl(v string) *UploadCreditAplusissueResponse {
	s.IssueUrl = &v
	return s
}

type QueryCreditCreditamountRequest struct {
	// OAuth模式下的授权token
	AuthToken         *string `json:"auth_token,omitempty" xml:"auth_token,omitempty"`
	ProductInstanceId *string `json:"product_instance_id,omitempty" xml:"product_instance_id,omitempty"`
	// 集团平台did
	GroupPlatformDid *string `json:"group_platform_did,omitempty" xml:"group_platform_did,omitempty" require:"true"`
	// 平台did
	PlatformDid *string `json:"platform_did,omitempty" xml:"platform_did,omitempty" require:"true"`
	// 产品ID
	ProductId *string `json:"product_id,omitempty" xml:"product_id,omitempty" require:"true"`
}

func (s QueryCreditCreditamountRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryCreditCreditamountRequest) GoString() string {
	return s.String()
}

func (s *QueryCreditCreditamountRequest) SetAuthToken(v string) *QueryCreditCreditamountRequest {
	s.AuthToken = &v
	return s
}

func (s *QueryCreditCreditamountRequest) SetProductInstanceId(v string) *QueryCreditCreditamountRequest {
	s.ProductInstanceId = &v
	return s
}

func (s *QueryCreditCreditamountRequest) SetGroupPlatformDid(v string) *QueryCreditCreditamountRequest {
	s.GroupPlatformDid = &v
	return s
}

func (s *QueryCreditCreditamountRequest) SetPlatformDid(v string) *QueryCreditCreditamountRequest {
	s.PlatformDid = &v
	return s
}

func (s *QueryCreditCreditamountRequest) SetProductId(v string) *QueryCreditCreditamountRequest {
	s.ProductId = &v
	return s
}

type QueryCreditCreditamountResponse struct {
	// 请求唯一ID，用于链路跟踪和问题排查
	ReqMsgId *string `json:"req_msg_id,omitempty" xml:"req_msg_id,omitempty"`
	// 结果码，一般OK表示调用成功
	ResultCode *string `json:"result_code,omitempty" xml:"result_code,omitempty"`
	// 异常信息的文本描述
	ResultMsg *string `json:"result_msg,omitempty" xml:"result_msg,omitempty"`
	// 可用额度
	CreditAmount *string `json:"credit_amount,omitempty" xml:"credit_amount,omitempty"`
	// 集团平台did
	GroupPlatformDid *string `json:"group_platform_did,omitempty" xml:"group_platform_did,omitempty"`
	// 平台did
	PlatformDid *string `json:"platform_did,omitempty" xml:"platform_did,omitempty"`
	// 产品ID
	ProductId *string `json:"product_id,omitempty" xml:"product_id,omitempty"`
}

func (s QueryCreditCreditamountResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryCreditCreditamountResponse) GoString() string {
	return s.String()
}

func (s *QueryCreditCreditamountResponse) SetReqMsgId(v string) *QueryCreditCreditamountResponse {
	s.ReqMsgId = &v
	return s
}

func (s *QueryCreditCreditamountResponse) SetResultCode(v string) *QueryCreditCreditamountResponse {
	s.ResultCode = &v
	return s
}

func (s *QueryCreditCreditamountResponse) SetResultMsg(v string) *QueryCreditCreditamountResponse {
	s.ResultMsg = &v
	return s
}

func (s *QueryCreditCreditamountResponse) SetCreditAmount(v string) *QueryCreditCreditamountResponse {
	s.CreditAmount = &v
	return s
}

func (s *QueryCreditCreditamountResponse) SetGroupPlatformDid(v string) *QueryCreditCreditamountResponse {
	s.GroupPlatformDid = &v
	return s
}

func (s *QueryCreditCreditamountResponse) SetPlatformDid(v string) *QueryCreditCreditamountResponse {
	s.PlatformDid = &v
	return s
}

func (s *QueryCreditCreditamountResponse) SetProductId(v string) *QueryCreditCreditamountResponse {
	s.ProductId = &v
	return s
}

type CreateCreditCommonsignRequest struct {
	// OAuth模式下的授权token
	AuthToken         *string `json:"auth_token,omitempty" xml:"auth_token,omitempty"`
	ProductInstanceId *string `json:"product_instance_id,omitempty" xml:"product_instance_id,omitempty"`
	// 账户号，account_type=ACCOUNT_BANK_NO时填写外部银行卡号， account_type=ACCOUNT_CLOUD_FUND时填写云资金商户ID
	AccountId *string `json:"account_id,omitempty" xml:"account_id,omitempty" require:"true"`
	// 账户类型，ACCOUNT_BANK_NO:外部银行卡，ACCOUNT_CLOUD_FUND:云资金商户ID
	AccountType *string `json:"account_type,omitempty" xml:"account_type,omitempty" require:"true"`
	// 开户行联行号，当帐户类型为ACCOUNT_BANK_NO时必填
	BankCnaps *string `json:"bank_cnaps,omitempty" xml:"bank_cnaps,omitempty"`
	// 开户行名称，当帐户类型为ACCOUNT_BANK_NO时必填
	BankName *string `json:"bank_name,omitempty" xml:"bank_name,omitempty"`
	// 签约结果回调地址
	CallbackUrl *string `json:"callback_url,omitempty" xml:"callback_url,omitempty" require:"true"`
	// 驾驶证号码，当帐户类型为ACCOUNT_BANK_NO且签约方sign_did类型为个人时必填
	DrivingLicense *string `json:"driving_license,omitempty" xml:"driving_license,omitempty"`
	// 驾驶证影像件ID，当帐户类型为ACCOUNT_BANK_NO且签约方sign_did类型为个人时必填
	DrivingLicenseFileId *string `json:"driving_license_file_id,omitempty" xml:"driving_license_file_id,omitempty"`
	// 集团平台did
	GroupPlatformDid *string `json:"group_platform_did,omitempty" xml:"group_platform_did,omitempty" require:"true"`
	// 法定代表人名称，签约方sign_did类型为企业时必填
	LegalName *string `json:"legal_name,omitempty" xml:"legal_name,omitempty"`
	// 产品id
	ProductId *string `json:"product_id,omitempty" xml:"product_id,omitempty" require:"true"`
	// 签约方did
	SignDid *string `json:"sign_did,omitempty" xml:"sign_did,omitempty" require:"true"`
	// 行驶证号码，当帐户类型为ACCOUNT_BANK_NO且签约方sign_did类型为个人时必填
	VehicleLicense *string `json:"vehicle_license,omitempty" xml:"vehicle_license,omitempty"`
	// 行驶证影像件ID，当帐户类型为ACCOUNT_BANK_NO且签约方sign_did类型为个人时必填
	VehicleLicenseFileId *string `json:"vehicle_license_file_id,omitempty" xml:"vehicle_license_file_id,omitempty"`
}

func (s CreateCreditCommonsignRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateCreditCommonsignRequest) GoString() string {
	return s.String()
}

func (s *CreateCreditCommonsignRequest) SetAuthToken(v string) *CreateCreditCommonsignRequest {
	s.AuthToken = &v
	return s
}

func (s *CreateCreditCommonsignRequest) SetProductInstanceId(v string) *CreateCreditCommonsignRequest {
	s.ProductInstanceId = &v
	return s
}

func (s *CreateCreditCommonsignRequest) SetAccountId(v string) *CreateCreditCommonsignRequest {
	s.AccountId = &v
	return s
}

func (s *CreateCreditCommonsignRequest) SetAccountType(v string) *CreateCreditCommonsignRequest {
	s.AccountType = &v
	return s
}

func (s *CreateCreditCommonsignRequest) SetBankCnaps(v string) *CreateCreditCommonsignRequest {
	s.BankCnaps = &v
	return s
}

func (s *CreateCreditCommonsignRequest) SetBankName(v string) *CreateCreditCommonsignRequest {
	s.BankName = &v
	return s
}

func (s *CreateCreditCommonsignRequest) SetCallbackUrl(v string) *CreateCreditCommonsignRequest {
	s.CallbackUrl = &v
	return s
}

func (s *CreateCreditCommonsignRequest) SetDrivingLicense(v string) *CreateCreditCommonsignRequest {
	s.DrivingLicense = &v
	return s
}

func (s *CreateCreditCommonsignRequest) SetDrivingLicenseFileId(v string) *CreateCreditCommonsignRequest {
	s.DrivingLicenseFileId = &v
	return s
}

func (s *CreateCreditCommonsignRequest) SetGroupPlatformDid(v string) *CreateCreditCommonsignRequest {
	s.GroupPlatformDid = &v
	return s
}

func (s *CreateCreditCommonsignRequest) SetLegalName(v string) *CreateCreditCommonsignRequest {
	s.LegalName = &v
	return s
}

func (s *CreateCreditCommonsignRequest) SetProductId(v string) *CreateCreditCommonsignRequest {
	s.ProductId = &v
	return s
}

func (s *CreateCreditCommonsignRequest) SetSignDid(v string) *CreateCreditCommonsignRequest {
	s.SignDid = &v
	return s
}

func (s *CreateCreditCommonsignRequest) SetVehicleLicense(v string) *CreateCreditCommonsignRequest {
	s.VehicleLicense = &v
	return s
}

func (s *CreateCreditCommonsignRequest) SetVehicleLicenseFileId(v string) *CreateCreditCommonsignRequest {
	s.VehicleLicenseFileId = &v
	return s
}

type CreateCreditCommonsignResponse struct {
	// 请求唯一ID，用于链路跟踪和问题排查
	ReqMsgId *string `json:"req_msg_id,omitempty" xml:"req_msg_id,omitempty"`
	// 结果码，一般OK表示调用成功
	ResultCode *string `json:"result_code,omitempty" xml:"result_code,omitempty"`
	// 异常信息的文本描述
	ResultMsg *string `json:"result_msg,omitempty" xml:"result_msg,omitempty"`
	// 申请id
	ApplyId *string `json:"apply_id,omitempty" xml:"apply_id,omitempty"`
	// 签约状态，-1:签约失败，0:未签约，1:已签约，2:签约中
	Status *int64 `json:"status,omitempty" xml:"status,omitempty"`
}

func (s CreateCreditCommonsignResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateCreditCommonsignResponse) GoString() string {
	return s.String()
}

func (s *CreateCreditCommonsignResponse) SetReqMsgId(v string) *CreateCreditCommonsignResponse {
	s.ReqMsgId = &v
	return s
}

func (s *CreateCreditCommonsignResponse) SetResultCode(v string) *CreateCreditCommonsignResponse {
	s.ResultCode = &v
	return s
}

func (s *CreateCreditCommonsignResponse) SetResultMsg(v string) *CreateCreditCommonsignResponse {
	s.ResultMsg = &v
	return s
}

func (s *CreateCreditCommonsignResponse) SetApplyId(v string) *CreateCreditCommonsignResponse {
	s.ApplyId = &v
	return s
}

func (s *CreateCreditCommonsignResponse) SetStatus(v int64) *CreateCreditCommonsignResponse {
	s.Status = &v
	return s
}

type QueryCreditCommonsignRequest struct {
	// OAuth模式下的授权token
	AuthToken         *string `json:"auth_token,omitempty" xml:"auth_token,omitempty"`
	ProductInstanceId *string `json:"product_instance_id,omitempty" xml:"product_instance_id,omitempty"`
	// 产品id
	ProductId *string `json:"product_id,omitempty" xml:"product_id,omitempty" require:"true"`
	// 集团平台did
	GroupPlatformDid *string `json:"group_platform_did,omitempty" xml:"group_platform_did,omitempty" require:"true"`
	// 签约方did
	SignDid *string `json:"sign_did,omitempty" xml:"sign_did,omitempty" require:"true"`
	// 申请id，不填则查询当前签约状态，填写则查询apply_id对应的签约申请结果
	ApplyId *string `json:"apply_id,omitempty" xml:"apply_id,omitempty"`
}

func (s QueryCreditCommonsignRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryCreditCommonsignRequest) GoString() string {
	return s.String()
}

func (s *QueryCreditCommonsignRequest) SetAuthToken(v string) *QueryCreditCommonsignRequest {
	s.AuthToken = &v
	return s
}

func (s *QueryCreditCommonsignRequest) SetProductInstanceId(v string) *QueryCreditCommonsignRequest {
	s.ProductInstanceId = &v
	return s
}

func (s *QueryCreditCommonsignRequest) SetProductId(v string) *QueryCreditCommonsignRequest {
	s.ProductId = &v
	return s
}

func (s *QueryCreditCommonsignRequest) SetGroupPlatformDid(v string) *QueryCreditCommonsignRequest {
	s.GroupPlatformDid = &v
	return s
}

func (s *QueryCreditCommonsignRequest) SetSignDid(v string) *QueryCreditCommonsignRequest {
	s.SignDid = &v
	return s
}

func (s *QueryCreditCommonsignRequest) SetApplyId(v string) *QueryCreditCommonsignRequest {
	s.ApplyId = &v
	return s
}

type QueryCreditCommonsignResponse struct {
	// 请求唯一ID，用于链路跟踪和问题排查
	ReqMsgId *string `json:"req_msg_id,omitempty" xml:"req_msg_id,omitempty"`
	// 结果码，一般OK表示调用成功
	ResultCode *string `json:"result_code,omitempty" xml:"result_code,omitempty"`
	// 异常信息的文本描述
	ResultMsg *string `json:"result_msg,omitempty" xml:"result_msg,omitempty"`
	// 签约结果，-1:签约失败，0:未签约，1:已签约，2:签约中
	Status *int64 `json:"status,omitempty" xml:"status,omitempty"`
	// 查询信息
	QueryMsg *string `json:"query_msg,omitempty" xml:"query_msg,omitempty"`
}

func (s QueryCreditCommonsignResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryCreditCommonsignResponse) GoString() string {
	return s.String()
}

func (s *QueryCreditCommonsignResponse) SetReqMsgId(v string) *QueryCreditCommonsignResponse {
	s.ReqMsgId = &v
	return s
}

func (s *QueryCreditCommonsignResponse) SetResultCode(v string) *QueryCreditCommonsignResponse {
	s.ResultCode = &v
	return s
}

func (s *QueryCreditCommonsignResponse) SetResultMsg(v string) *QueryCreditCommonsignResponse {
	s.ResultMsg = &v
	return s
}

func (s *QueryCreditCommonsignResponse) SetStatus(v int64) *QueryCreditCommonsignResponse {
	s.Status = &v
	return s
}

func (s *QueryCreditCommonsignResponse) SetQueryMsg(v string) *QueryCreditCommonsignResponse {
	s.QueryMsg = &v
	return s
}

type BatchcreateCreditIssueRequest struct {
	// OAuth模式下的授权token
	AuthToken         *string `json:"auth_token,omitempty" xml:"auth_token,omitempty"`
	ProductInstanceId *string `json:"product_instance_id,omitempty" xml:"product_instance_id,omitempty"`
	// 产品id
	ProductId *string `json:"product_id,omitempty" xml:"product_id,omitempty" require:"true"`
	// 集团平台did
	GroupPlatformDid *string `json:"group_platform_did,omitempty" xml:"group_platform_did,omitempty" require:"true"`
	// 凭证发行方分布式身份
	CreateDid *string `json:"create_did,omitempty" xml:"create_did,omitempty" require:"true"`
	// 凭证接收方分布式身份
	ReceiveDid *string `json:"receive_did,omitempty" xml:"receive_did,omitempty" require:"true"`
	// 批次id
	BatchId *string `json:"batch_id,omitempty" xml:"batch_id,omitempty" require:"true"`
	// 全局业务唯一号
	OutBizNo *string `json:"out_biz_no,omitempty" xml:"out_biz_no,omitempty" require:"true"`
	// 支付订单号
	PayOrder *string `json:"pay_order,omitempty" xml:"pay_order,omitempty" require:"true"`
	// 支付单运费总额，运费最多精确到小数点后2位
	Freight *string `json:"freight,omitempty" xml:"freight,omitempty" require:"true"`
	// 发行后自动贴现，true: 自动贴现，false: 不自动贴现 （默认）
	EasyFinance *bool `json:"easy_finance,omitempty" xml:"easy_finance,omitempty"`
	// 凭证到期时间
	ExpireDate *string `json:"expire_date,omitempty" xml:"expire_date,omitempty" require:"true"`
	// 发行结果回调地址
	CallbackUrl *string `json:"callback_url,omitempty" xml:"callback_url,omitempty" require:"true"`
	// 需合并发行的运单id列表
	WaybillIds []*string `json:"waybill_ids,omitempty" xml:"waybill_ids,omitempty" require:"true" type:"Repeated"`
}

func (s BatchcreateCreditIssueRequest) String() string {
	return tea.Prettify(s)
}

func (s BatchcreateCreditIssueRequest) GoString() string {
	return s.String()
}

func (s *BatchcreateCreditIssueRequest) SetAuthToken(v string) *BatchcreateCreditIssueRequest {
	s.AuthToken = &v
	return s
}

func (s *BatchcreateCreditIssueRequest) SetProductInstanceId(v string) *BatchcreateCreditIssueRequest {
	s.ProductInstanceId = &v
	return s
}

func (s *BatchcreateCreditIssueRequest) SetProductId(v string) *BatchcreateCreditIssueRequest {
	s.ProductId = &v
	return s
}

func (s *BatchcreateCreditIssueRequest) SetGroupPlatformDid(v string) *BatchcreateCreditIssueRequest {
	s.GroupPlatformDid = &v
	return s
}

func (s *BatchcreateCreditIssueRequest) SetCreateDid(v string) *BatchcreateCreditIssueRequest {
	s.CreateDid = &v
	return s
}

func (s *BatchcreateCreditIssueRequest) SetReceiveDid(v string) *BatchcreateCreditIssueRequest {
	s.ReceiveDid = &v
	return s
}

func (s *BatchcreateCreditIssueRequest) SetBatchId(v string) *BatchcreateCreditIssueRequest {
	s.BatchId = &v
	return s
}

func (s *BatchcreateCreditIssueRequest) SetOutBizNo(v string) *BatchcreateCreditIssueRequest {
	s.OutBizNo = &v
	return s
}

func (s *BatchcreateCreditIssueRequest) SetPayOrder(v string) *BatchcreateCreditIssueRequest {
	s.PayOrder = &v
	return s
}

func (s *BatchcreateCreditIssueRequest) SetFreight(v string) *BatchcreateCreditIssueRequest {
	s.Freight = &v
	return s
}

func (s *BatchcreateCreditIssueRequest) SetEasyFinance(v bool) *BatchcreateCreditIssueRequest {
	s.EasyFinance = &v
	return s
}

func (s *BatchcreateCreditIssueRequest) SetExpireDate(v string) *BatchcreateCreditIssueRequest {
	s.ExpireDate = &v
	return s
}

func (s *BatchcreateCreditIssueRequest) SetCallbackUrl(v string) *BatchcreateCreditIssueRequest {
	s.CallbackUrl = &v
	return s
}

func (s *BatchcreateCreditIssueRequest) SetWaybillIds(v []*string) *BatchcreateCreditIssueRequest {
	s.WaybillIds = v
	return s
}

type BatchcreateCreditIssueResponse struct {
	// 请求唯一ID，用于链路跟踪和问题排查
	ReqMsgId *string `json:"req_msg_id,omitempty" xml:"req_msg_id,omitempty"`
	// 结果码，一般OK表示调用成功
	ResultCode *string `json:"result_code,omitempty" xml:"result_code,omitempty"`
	// 异常信息的文本描述
	ResultMsg *string `json:"result_msg,omitempty" xml:"result_msg,omitempty"`
	// 批次id
	BatchId *string `json:"batch_id,omitempty" xml:"batch_id,omitempty"`
	// 批次状态
	BatchStatus *int64 `json:"batch_status,omitempty" xml:"batch_status,omitempty"`
}

func (s BatchcreateCreditIssueResponse) String() string {
	return tea.Prettify(s)
}

func (s BatchcreateCreditIssueResponse) GoString() string {
	return s.String()
}

func (s *BatchcreateCreditIssueResponse) SetReqMsgId(v string) *BatchcreateCreditIssueResponse {
	s.ReqMsgId = &v
	return s
}

func (s *BatchcreateCreditIssueResponse) SetResultCode(v string) *BatchcreateCreditIssueResponse {
	s.ResultCode = &v
	return s
}

func (s *BatchcreateCreditIssueResponse) SetResultMsg(v string) *BatchcreateCreditIssueResponse {
	s.ResultMsg = &v
	return s
}

func (s *BatchcreateCreditIssueResponse) SetBatchId(v string) *BatchcreateCreditIssueResponse {
	s.BatchId = &v
	return s
}

func (s *BatchcreateCreditIssueResponse) SetBatchStatus(v int64) *BatchcreateCreditIssueResponse {
	s.BatchStatus = &v
	return s
}

type UploadCreditIssuebatchRequest struct {
	// OAuth模式下的授权token
	AuthToken         *string `json:"auth_token,omitempty" xml:"auth_token,omitempty"`
	ProductInstanceId *string `json:"product_instance_id,omitempty" xml:"product_instance_id,omitempty"`
	// 批次id
	BatchId *string `json:"batch_id,omitempty" xml:"batch_id,omitempty" require:"true"`
	// 回调地址
	CallbackUrl *string `json:"callback_url,omitempty" xml:"callback_url,omitempty" require:"true"`
	// 凭证发行方did
	CreateDid *string `json:"create_did,omitempty" xml:"create_did,omitempty" require:"true"`
	// 发行后自动贴现，true: 自动贴现，false: 不自动贴现 （默认）
	EasyFinance *bool `json:"easy_finance,omitempty" xml:"easy_finance,omitempty"`
	// 凭证到期时间
	ExpireDate *string `json:"expire_date,omitempty" xml:"expire_date,omitempty" require:"true"`
	// 支付单运费总额，运费最多精确到小数点后2位
	Freight *string `json:"freight,omitempty" xml:"freight,omitempty" require:"true"`
	// 集团平台did
	GroupPlatformDid *string `json:"group_platform_did,omitempty" xml:"group_platform_did,omitempty" require:"true"`
	// 全局业务唯一编号
	OutBizNo *string `json:"out_biz_no,omitempty" xml:"out_biz_no,omitempty" require:"true"`
	// 支付订单号
	PayOrder *string `json:"pay_order,omitempty" xml:"pay_order,omitempty" require:"true"`
	// 产品id
	ProductId *string `json:"product_id,omitempty" xml:"product_id,omitempty" require:"true"`
	// 凭证接收方did
	ReceiveDid *string `json:"receive_did,omitempty" xml:"receive_did,omitempty" require:"true"`
	// 运单合并列表
	WaybillIds []*string `json:"waybill_ids,omitempty" xml:"waybill_ids,omitempty" require:"true" type:"Repeated"`
}

func (s UploadCreditIssuebatchRequest) String() string {
	return tea.Prettify(s)
}

func (s UploadCreditIssuebatchRequest) GoString() string {
	return s.String()
}

func (s *UploadCreditIssuebatchRequest) SetAuthToken(v string) *UploadCreditIssuebatchRequest {
	s.AuthToken = &v
	return s
}

func (s *UploadCreditIssuebatchRequest) SetProductInstanceId(v string) *UploadCreditIssuebatchRequest {
	s.ProductInstanceId = &v
	return s
}

func (s *UploadCreditIssuebatchRequest) SetBatchId(v string) *UploadCreditIssuebatchRequest {
	s.BatchId = &v
	return s
}

func (s *UploadCreditIssuebatchRequest) SetCallbackUrl(v string) *UploadCreditIssuebatchRequest {
	s.CallbackUrl = &v
	return s
}

func (s *UploadCreditIssuebatchRequest) SetCreateDid(v string) *UploadCreditIssuebatchRequest {
	s.CreateDid = &v
	return s
}

func (s *UploadCreditIssuebatchRequest) SetEasyFinance(v bool) *UploadCreditIssuebatchRequest {
	s.EasyFinance = &v
	return s
}

func (s *UploadCreditIssuebatchRequest) SetExpireDate(v string) *UploadCreditIssuebatchRequest {
	s.ExpireDate = &v
	return s
}

func (s *UploadCreditIssuebatchRequest) SetFreight(v string) *UploadCreditIssuebatchRequest {
	s.Freight = &v
	return s
}

func (s *UploadCreditIssuebatchRequest) SetGroupPlatformDid(v string) *UploadCreditIssuebatchRequest {
	s.GroupPlatformDid = &v
	return s
}

func (s *UploadCreditIssuebatchRequest) SetOutBizNo(v string) *UploadCreditIssuebatchRequest {
	s.OutBizNo = &v
	return s
}

func (s *UploadCreditIssuebatchRequest) SetPayOrder(v string) *UploadCreditIssuebatchRequest {
	s.PayOrder = &v
	return s
}

func (s *UploadCreditIssuebatchRequest) SetProductId(v string) *UploadCreditIssuebatchRequest {
	s.ProductId = &v
	return s
}

func (s *UploadCreditIssuebatchRequest) SetReceiveDid(v string) *UploadCreditIssuebatchRequest {
	s.ReceiveDid = &v
	return s
}

func (s *UploadCreditIssuebatchRequest) SetWaybillIds(v []*string) *UploadCreditIssuebatchRequest {
	s.WaybillIds = v
	return s
}

type UploadCreditIssuebatchResponse struct {
	// 请求唯一ID，用于链路跟踪和问题排查
	ReqMsgId *string `json:"req_msg_id,omitempty" xml:"req_msg_id,omitempty"`
	// 结果码，一般OK表示调用成功
	ResultCode *string `json:"result_code,omitempty" xml:"result_code,omitempty"`
	// 异常信息的文本描述
	ResultMsg *string `json:"result_msg,omitempty" xml:"result_msg,omitempty"`
	// 批次id
	BatchId *string `json:"batch_id,omitempty" xml:"batch_id,omitempty"`
	// 批次状态
	BatchStatus *int64 `json:"batch_status,omitempty" xml:"batch_status,omitempty"`
	// 发行链接
	IssueUrl *string `json:"issue_url,omitempty" xml:"issue_url,omitempty"`
}

func (s UploadCreditIssuebatchResponse) String() string {
	return tea.Prettify(s)
}

func (s UploadCreditIssuebatchResponse) GoString() string {
	return s.String()
}

func (s *UploadCreditIssuebatchResponse) SetReqMsgId(v string) *UploadCreditIssuebatchResponse {
	s.ReqMsgId = &v
	return s
}

func (s *UploadCreditIssuebatchResponse) SetResultCode(v string) *UploadCreditIssuebatchResponse {
	s.ResultCode = &v
	return s
}

func (s *UploadCreditIssuebatchResponse) SetResultMsg(v string) *UploadCreditIssuebatchResponse {
	s.ResultMsg = &v
	return s
}

func (s *UploadCreditIssuebatchResponse) SetBatchId(v string) *UploadCreditIssuebatchResponse {
	s.BatchId = &v
	return s
}

func (s *UploadCreditIssuebatchResponse) SetBatchStatus(v int64) *UploadCreditIssuebatchResponse {
	s.BatchStatus = &v
	return s
}

func (s *UploadCreditIssuebatchResponse) SetIssueUrl(v string) *UploadCreditIssuebatchResponse {
	s.IssueUrl = &v
	return s
}

type GetIssueTransferfileRequest struct {
	// OAuth模式下的授权token
	AuthToken         *string `json:"auth_token,omitempty" xml:"auth_token,omitempty"`
	ProductInstanceId *string `json:"product_instance_id,omitempty" xml:"product_instance_id,omitempty"`
	// 申请唯一流水号
	ApplyId *string `json:"apply_id,omitempty" xml:"apply_id,omitempty" require:"true"`
	// 回调地址
	//
	CallbackUrl *string `json:"callback_url,omitempty" xml:"callback_url,omitempty" require:"true"`
	// 集团平台方分布式数字身份
	GroupPlatformDid *string `json:"group_platform_did,omitempty" xml:"group_platform_did,omitempty" require:"true"`
	// 信用流转凭证集合，最大长度支持30
	IssueTransferDatas []*IssueTransferData `json:"issue_transfer_datas,omitempty" xml:"issue_transfer_datas,omitempty" require:"true" type:"Repeated"`
	// 业务发起方分布式数字身份
	PlatformDid *string `json:"platform_did,omitempty" xml:"platform_did,omitempty" require:"true"`
	// 产品id
	//
	ProductId *string `json:"product_id,omitempty" xml:"product_id,omitempty" require:"true"`
}

func (s GetIssueTransferfileRequest) String() string {
	return tea.Prettify(s)
}

func (s GetIssueTransferfileRequest) GoString() string {
	return s.String()
}

func (s *GetIssueTransferfileRequest) SetAuthToken(v string) *GetIssueTransferfileRequest {
	s.AuthToken = &v
	return s
}

func (s *GetIssueTransferfileRequest) SetProductInstanceId(v string) *GetIssueTransferfileRequest {
	s.ProductInstanceId = &v
	return s
}

func (s *GetIssueTransferfileRequest) SetApplyId(v string) *GetIssueTransferfileRequest {
	s.ApplyId = &v
	return s
}

func (s *GetIssueTransferfileRequest) SetCallbackUrl(v string) *GetIssueTransferfileRequest {
	s.CallbackUrl = &v
	return s
}

func (s *GetIssueTransferfileRequest) SetGroupPlatformDid(v string) *GetIssueTransferfileRequest {
	s.GroupPlatformDid = &v
	return s
}

func (s *GetIssueTransferfileRequest) SetIssueTransferDatas(v []*IssueTransferData) *GetIssueTransferfileRequest {
	s.IssueTransferDatas = v
	return s
}

func (s *GetIssueTransferfileRequest) SetPlatformDid(v string) *GetIssueTransferfileRequest {
	s.PlatformDid = &v
	return s
}

func (s *GetIssueTransferfileRequest) SetProductId(v string) *GetIssueTransferfileRequest {
	s.ProductId = &v
	return s
}

type GetIssueTransferfileResponse struct {
	// 请求唯一ID，用于链路跟踪和问题排查
	ReqMsgId *string `json:"req_msg_id,omitempty" xml:"req_msg_id,omitempty"`
	// 结果码，一般OK表示调用成功
	ResultCode *string `json:"result_code,omitempty" xml:"result_code,omitempty"`
	// 异常信息的文本描述
	ResultMsg *string `json:"result_msg,omitempty" xml:"result_msg,omitempty"`
	// 申请唯一流水号
	ApplyId *string `json:"apply_id,omitempty" xml:"apply_id,omitempty"`
	// 集团平台方分布式数字身份
	GroupPlatformDid *string `json:"group_platform_did,omitempty" xml:"group_platform_did,omitempty"`
	// 业务发起方分布式数字身
	PlatformDid *string `json:"platform_did,omitempty" xml:"platform_did,omitempty"`
	// 产品id
	ProductId *string `json:"product_id,omitempty" xml:"product_id,omitempty"`
}

func (s GetIssueTransferfileResponse) String() string {
	return tea.Prettify(s)
}

func (s GetIssueTransferfileResponse) GoString() string {
	return s.String()
}

func (s *GetIssueTransferfileResponse) SetReqMsgId(v string) *GetIssueTransferfileResponse {
	s.ReqMsgId = &v
	return s
}

func (s *GetIssueTransferfileResponse) SetResultCode(v string) *GetIssueTransferfileResponse {
	s.ResultCode = &v
	return s
}

func (s *GetIssueTransferfileResponse) SetResultMsg(v string) *GetIssueTransferfileResponse {
	s.ResultMsg = &v
	return s
}

func (s *GetIssueTransferfileResponse) SetApplyId(v string) *GetIssueTransferfileResponse {
	s.ApplyId = &v
	return s
}

func (s *GetIssueTransferfileResponse) SetGroupPlatformDid(v string) *GetIssueTransferfileResponse {
	s.GroupPlatformDid = &v
	return s
}

func (s *GetIssueTransferfileResponse) SetPlatformDid(v string) *GetIssueTransferfileResponse {
	s.PlatformDid = &v
	return s
}

func (s *GetIssueTransferfileResponse) SetProductId(v string) *GetIssueTransferfileResponse {
	s.ProductId = &v
	return s
}

type ApplyInsurancepolicyZhonghuacaixianRequest struct {
	// OAuth模式下的授权token
	AuthToken         *string `json:"auth_token,omitempty" xml:"auth_token,omitempty"`
	ProductInstanceId *string `json:"product_instance_id,omitempty" xml:"product_instance_id,omitempty"`
	// 被保险人地址
	BbrAddr *string `json:"bbr_addr,omitempty" xml:"bbr_addr,omitempty" maxLength:"80"`
	// 被保险人证件号码
	BbrIdNo *string `json:"bbr_id_no,omitempty" xml:"bbr_id_no,omitempty" require:"true" maxLength:"40"`
	// 被保险人证件类型。按照如下值填写
	// 464001	身份证
	// 464002	护照
	// 464003	其他
	// 464004	组织机构代码
	// 464005	军人证
	// 464006	工商注册号码
	// 464007	统一社会信用代码
	// 464008	临时身份证
	// 464009	外国护照
	// 464010	中国人民武装警察身份证件
	// 464011	军官证
	// 464012	企业营业执照号码
	// 464013	统一社会信用代码（五证合一号码）
	// 464014	个体工商户营业执照号码
	// 464015	户口本
	// 464016	其他类境内个人身份有效证件
	// 464017	其他类境外个人身份有效证件
	// 464018	税务登记证
	// 464019	金融许可证号码
	// 464020	国家主管部门颁外国驻华机构批文号码
	// 464021	其他类境外机构代码
	BbrIdType *string `json:"bbr_id_type,omitempty" xml:"bbr_id_type,omitempty" require:"true" maxLength:"40"`
	// 被保险人姓名
	BbrName *string `json:"bbr_name,omitempty" xml:"bbr_name,omitempty" require:"true" maxLength:"40"`
	// 被保险人联系电话
	BbrTel *string `json:"bbr_tel,omitempty" xml:"bbr_tel,omitempty" require:"true" maxLength:"40"`
	// 运费,四舍五入精确到小数点两位。系统将根据运费和费率计算含税保费，计算的保费结果为四舍五入，精确到两位小数
	Carriage *string `json:"carriage,omitempty" xml:"carriage,omitempty" require:"true" maxLength:"20"`
	// 货物名称
	CarGo *string `json:"car_go,omitempty" xml:"car_go,omitempty" require:"true" maxLength:"40"`
	// 厂牌型号
	CpModel *string `json:"cp_model,omitempty" xml:"cp_model,omitempty" require:"true" maxLength:"80"`
	// 目的地
	Destination *string `json:"destination,omitempty" xml:"destination,omitempty" maxLength:"200"`
	// 行驶证车主
	DrivPer *string `json:"driv_per,omitempty" xml:"driv_per,omitempty" require:"true" maxLength:"40"`
	// 保险起期，精确到天；最短起保时间为次日0点，最长延时起保时间为次日0点后24h
	EffDate *string `json:"eff_date,omitempty" xml:"eff_date,omitempty" require:"true" pattern:"\\d{4}[-]\\d{1,2}[-]\\d{1,2}[T]\\d{2}:\\d{2}:\\d{2}([Z]|([\\.]\\d{1,9})?[\\+]\\d{2}[\\:]?\\d{2})"`
	// 发动机编号
	EngineNo *string `json:"engine_no,omitempty" xml:"engine_no,omitempty" require:"true" maxLength:"40"`
	// 车架号
	FrameNo *string `json:"frame_no,omitempty" xml:"frame_no,omitempty" require:"true" maxLength:"40"`
	// 运单所属集团分布式身份标识
	GroupPlatformDid *string `json:"group_platform_did,omitempty" xml:"group_platform_did,omitempty" require:"true" maxLength:"80"`
	// 投保人证件类型有效止期
	IdentifyPeriodEnd *string `json:"identify_period_end,omitempty" xml:"identify_period_end,omitempty" pattern:"\\d{4}[-]\\d{1,2}[-]\\d{1,2}[T]\\d{2}:\\d{2}:\\d{2}([Z]|([\\.]\\d{1,9})?[\\+]\\d{2}[\\:]?\\d{2})"`
	// 投保人证件类型有效起期
	IdentifyPeriodStart *string `json:"identify_period_start,omitempty" xml:"identify_period_start,omitempty" pattern:"\\d{4}[-]\\d{1,2}[-]\\d{1,2}[T]\\d{2}:\\d{2}:\\d{2}([Z]|([\\.]\\d{1,9})?[\\+]\\d{2}[\\:]?\\d{2})"`
	// 车牌号码
	LicenseNo *string `json:"license_no,omitempty" xml:"license_no,omitempty" require:"true" maxLength:"40"`
	// 运单所属平台分布式身份标识
	PlatformDid *string `json:"platform_did,omitempty" xml:"platform_did,omitempty" require:"true" maxLength:"80"`
	// 运营证号
	RunNo *string `json:"run_no,omitempty" xml:"run_no,omitempty" require:"true" maxLength:"40"`
	// 起运地
	StartPlace *string `json:"start_place,omitempty" xml:"start_place,omitempty" maxLength:"200"`
	// 投保人地址
	TbrAddr *string `json:"tbr_addr,omitempty" xml:"tbr_addr,omitempty" maxLength:"80"`
	// 投保人邮箱
	TbrEmail *string `json:"tbr_email,omitempty" xml:"tbr_email,omitempty" maxLength:"40"`
	// 投保人证件号
	TbrIdNo *string `json:"tbr_id_no,omitempty" xml:"tbr_id_no,omitempty" require:"true" maxLength:"40"`
	// 投保人证件类型，按照如下状态进行填写
	// 464001	身份证
	// 464002	护照
	// 464003	其他
	// 464004	组织机构代码
	// 464005	军人证
	// 464006	工商注册号码
	// 464007	统一社会信用代码
	// 464008	临时身份证
	// 464009	外国护照
	// 464010	中国人民武装警察身份证件
	// 464011	军官证
	// 464012	企业营业执照号码
	// 464013	统一社会信用代码（五证合一号码）
	// 464014	个体工商户营业执照号码
	// 464015	户口本
	// 464016	其他类境内个人身份有效证件
	// 464017	其他类境外个人身份有效证件
	// 464018	税务登记证
	// 464019	金融许可证号码
	// 464020	国家主管部门颁外国驻华机构批文号码
	// 464021	其他类境外机构代码
	TbrIdType *string `json:"tbr_id_type,omitempty" xml:"tbr_id_type,omitempty" require:"true" maxLength:"40"`
	// 投保人姓名
	TbrName *string `json:"tbr_name,omitempty" xml:"tbr_name,omitempty" require:"true" maxLength:"40"`
	// 投保人联系电话
	TbrTel *string `json:"tbr_tel,omitempty" xml:"tbr_tel,omitempty" require:"true" maxLength:"40"`
	// 保险止期，保期时间间隔为1-30天。
	TermDate *string `json:"term_date,omitempty" xml:"term_date,omitempty" require:"true" pattern:"\\d{4}[-]\\d{1,2}[-]\\d{1,2}[T]\\d{2}:\\d{2}:\\d{2}([Z]|([\\.]\\d{1,9})?[\\+]\\d{2}[\\:]?\\d{2})"`
	// 吨位
	TonNage *string `json:"ton_nage,omitempty" xml:"ton_nage,omitempty" require:"true" maxLength:"20"`
	// 交易流水号,全局唯一；格式为 yyyyMMdd+身份标识+其他编码。系统会根据该流水号做防重、幂等判断逻辑。当极端场景中，系统会返回处理中状态，客户端应该保持该流水号不变，并使用原来的请求再次发送请求，系统会根据幂等逻辑返回处理结果
	TradeNo *string `json:"trade_no,omitempty" xml:"trade_no,omitempty" require:"true" maxLength:"32" minLength:"32"`
	// 运输货物
	TsCarGo *string `json:"ts_car_go,omitempty" xml:"ts_car_go,omitempty" require:"true" maxLength:"80"`
	// 运单id。通过运单创建接口上传运单时指定的运单标识。系统会根据该标识查询运单相关信息做投保业务校验
	WaybillId *string `json:"waybill_id,omitempty" xml:"waybill_id,omitempty" require:"true" maxLength:"128"`
	// 重量
	Weight *string `json:"weight,omitempty" xml:"weight,omitempty" maxLength:"20"`
}

func (s ApplyInsurancepolicyZhonghuacaixianRequest) String() string {
	return tea.Prettify(s)
}

func (s ApplyInsurancepolicyZhonghuacaixianRequest) GoString() string {
	return s.String()
}

func (s *ApplyInsurancepolicyZhonghuacaixianRequest) SetAuthToken(v string) *ApplyInsurancepolicyZhonghuacaixianRequest {
	s.AuthToken = &v
	return s
}

func (s *ApplyInsurancepolicyZhonghuacaixianRequest) SetProductInstanceId(v string) *ApplyInsurancepolicyZhonghuacaixianRequest {
	s.ProductInstanceId = &v
	return s
}

func (s *ApplyInsurancepolicyZhonghuacaixianRequest) SetBbrAddr(v string) *ApplyInsurancepolicyZhonghuacaixianRequest {
	s.BbrAddr = &v
	return s
}

func (s *ApplyInsurancepolicyZhonghuacaixianRequest) SetBbrIdNo(v string) *ApplyInsurancepolicyZhonghuacaixianRequest {
	s.BbrIdNo = &v
	return s
}

func (s *ApplyInsurancepolicyZhonghuacaixianRequest) SetBbrIdType(v string) *ApplyInsurancepolicyZhonghuacaixianRequest {
	s.BbrIdType = &v
	return s
}

func (s *ApplyInsurancepolicyZhonghuacaixianRequest) SetBbrName(v string) *ApplyInsurancepolicyZhonghuacaixianRequest {
	s.BbrName = &v
	return s
}

func (s *ApplyInsurancepolicyZhonghuacaixianRequest) SetBbrTel(v string) *ApplyInsurancepolicyZhonghuacaixianRequest {
	s.BbrTel = &v
	return s
}

func (s *ApplyInsurancepolicyZhonghuacaixianRequest) SetCarriage(v string) *ApplyInsurancepolicyZhonghuacaixianRequest {
	s.Carriage = &v
	return s
}

func (s *ApplyInsurancepolicyZhonghuacaixianRequest) SetCarGo(v string) *ApplyInsurancepolicyZhonghuacaixianRequest {
	s.CarGo = &v
	return s
}

func (s *ApplyInsurancepolicyZhonghuacaixianRequest) SetCpModel(v string) *ApplyInsurancepolicyZhonghuacaixianRequest {
	s.CpModel = &v
	return s
}

func (s *ApplyInsurancepolicyZhonghuacaixianRequest) SetDestination(v string) *ApplyInsurancepolicyZhonghuacaixianRequest {
	s.Destination = &v
	return s
}

func (s *ApplyInsurancepolicyZhonghuacaixianRequest) SetDrivPer(v string) *ApplyInsurancepolicyZhonghuacaixianRequest {
	s.DrivPer = &v
	return s
}

func (s *ApplyInsurancepolicyZhonghuacaixianRequest) SetEffDate(v string) *ApplyInsurancepolicyZhonghuacaixianRequest {
	s.EffDate = &v
	return s
}

func (s *ApplyInsurancepolicyZhonghuacaixianRequest) SetEngineNo(v string) *ApplyInsurancepolicyZhonghuacaixianRequest {
	s.EngineNo = &v
	return s
}

func (s *ApplyInsurancepolicyZhonghuacaixianRequest) SetFrameNo(v string) *ApplyInsurancepolicyZhonghuacaixianRequest {
	s.FrameNo = &v
	return s
}

func (s *ApplyInsurancepolicyZhonghuacaixianRequest) SetGroupPlatformDid(v string) *ApplyInsurancepolicyZhonghuacaixianRequest {
	s.GroupPlatformDid = &v
	return s
}

func (s *ApplyInsurancepolicyZhonghuacaixianRequest) SetIdentifyPeriodEnd(v string) *ApplyInsurancepolicyZhonghuacaixianRequest {
	s.IdentifyPeriodEnd = &v
	return s
}

func (s *ApplyInsurancepolicyZhonghuacaixianRequest) SetIdentifyPeriodStart(v string) *ApplyInsurancepolicyZhonghuacaixianRequest {
	s.IdentifyPeriodStart = &v
	return s
}

func (s *ApplyInsurancepolicyZhonghuacaixianRequest) SetLicenseNo(v string) *ApplyInsurancepolicyZhonghuacaixianRequest {
	s.LicenseNo = &v
	return s
}

func (s *ApplyInsurancepolicyZhonghuacaixianRequest) SetPlatformDid(v string) *ApplyInsurancepolicyZhonghuacaixianRequest {
	s.PlatformDid = &v
	return s
}

func (s *ApplyInsurancepolicyZhonghuacaixianRequest) SetRunNo(v string) *ApplyInsurancepolicyZhonghuacaixianRequest {
	s.RunNo = &v
	return s
}

func (s *ApplyInsurancepolicyZhonghuacaixianRequest) SetStartPlace(v string) *ApplyInsurancepolicyZhonghuacaixianRequest {
	s.StartPlace = &v
	return s
}

func (s *ApplyInsurancepolicyZhonghuacaixianRequest) SetTbrAddr(v string) *ApplyInsurancepolicyZhonghuacaixianRequest {
	s.TbrAddr = &v
	return s
}

func (s *ApplyInsurancepolicyZhonghuacaixianRequest) SetTbrEmail(v string) *ApplyInsurancepolicyZhonghuacaixianRequest {
	s.TbrEmail = &v
	return s
}

func (s *ApplyInsurancepolicyZhonghuacaixianRequest) SetTbrIdNo(v string) *ApplyInsurancepolicyZhonghuacaixianRequest {
	s.TbrIdNo = &v
	return s
}

func (s *ApplyInsurancepolicyZhonghuacaixianRequest) SetTbrIdType(v string) *ApplyInsurancepolicyZhonghuacaixianRequest {
	s.TbrIdType = &v
	return s
}

func (s *ApplyInsurancepolicyZhonghuacaixianRequest) SetTbrName(v string) *ApplyInsurancepolicyZhonghuacaixianRequest {
	s.TbrName = &v
	return s
}

func (s *ApplyInsurancepolicyZhonghuacaixianRequest) SetTbrTel(v string) *ApplyInsurancepolicyZhonghuacaixianRequest {
	s.TbrTel = &v
	return s
}

func (s *ApplyInsurancepolicyZhonghuacaixianRequest) SetTermDate(v string) *ApplyInsurancepolicyZhonghuacaixianRequest {
	s.TermDate = &v
	return s
}

func (s *ApplyInsurancepolicyZhonghuacaixianRequest) SetTonNage(v string) *ApplyInsurancepolicyZhonghuacaixianRequest {
	s.TonNage = &v
	return s
}

func (s *ApplyInsurancepolicyZhonghuacaixianRequest) SetTradeNo(v string) *ApplyInsurancepolicyZhonghuacaixianRequest {
	s.TradeNo = &v
	return s
}

func (s *ApplyInsurancepolicyZhonghuacaixianRequest) SetTsCarGo(v string) *ApplyInsurancepolicyZhonghuacaixianRequest {
	s.TsCarGo = &v
	return s
}

func (s *ApplyInsurancepolicyZhonghuacaixianRequest) SetWaybillId(v string) *ApplyInsurancepolicyZhonghuacaixianRequest {
	s.WaybillId = &v
	return s
}

func (s *ApplyInsurancepolicyZhonghuacaixianRequest) SetWeight(v string) *ApplyInsurancepolicyZhonghuacaixianRequest {
	s.Weight = &v
	return s
}

type ApplyInsurancepolicyZhonghuacaixianResponse struct {
	// 请求唯一ID，用于链路跟踪和问题排查
	ReqMsgId *string `json:"req_msg_id,omitempty" xml:"req_msg_id,omitempty"`
	// 结果码，一般OK表示调用成功
	ResultCode *string `json:"result_code,omitempty" xml:"result_code,omitempty"`
	// 异常信息的文本描述
	ResultMsg *string `json:"result_msg,omitempty" xml:"result_msg,omitempty"`
	// 保单号
	PolNo *string `json:"pol_no,omitempty" xml:"pol_no,omitempty"`
	// 电子保单下载地址
	PolUrl *string `json:"pol_url,omitempty" xml:"pol_url,omitempty"`
	// 含税保费
	PreMium *string `json:"pre_mium,omitempty" xml:"pre_mium,omitempty"`
}

func (s ApplyInsurancepolicyZhonghuacaixianResponse) String() string {
	return tea.Prettify(s)
}

func (s ApplyInsurancepolicyZhonghuacaixianResponse) GoString() string {
	return s.String()
}

func (s *ApplyInsurancepolicyZhonghuacaixianResponse) SetReqMsgId(v string) *ApplyInsurancepolicyZhonghuacaixianResponse {
	s.ReqMsgId = &v
	return s
}

func (s *ApplyInsurancepolicyZhonghuacaixianResponse) SetResultCode(v string) *ApplyInsurancepolicyZhonghuacaixianResponse {
	s.ResultCode = &v
	return s
}

func (s *ApplyInsurancepolicyZhonghuacaixianResponse) SetResultMsg(v string) *ApplyInsurancepolicyZhonghuacaixianResponse {
	s.ResultMsg = &v
	return s
}

func (s *ApplyInsurancepolicyZhonghuacaixianResponse) SetPolNo(v string) *ApplyInsurancepolicyZhonghuacaixianResponse {
	s.PolNo = &v
	return s
}

func (s *ApplyInsurancepolicyZhonghuacaixianResponse) SetPolUrl(v string) *ApplyInsurancepolicyZhonghuacaixianResponse {
	s.PolUrl = &v
	return s
}

func (s *ApplyInsurancepolicyZhonghuacaixianResponse) SetPreMium(v string) *ApplyInsurancepolicyZhonghuacaixianResponse {
	s.PreMium = &v
	return s
}

type CancelInsurancepolicyZhonghuacaixianRequest struct {
	// OAuth模式下的授权token
	AuthToken         *string `json:"auth_token,omitempty" xml:"auth_token,omitempty"`
	ProductInstanceId *string `json:"product_instance_id,omitempty" xml:"product_instance_id,omitempty"`
	// 保单号，投保成功后返回的
	PolNo *string `json:"pol_no,omitempty" xml:"pol_no,omitempty" require:"true" maxLength:"60"`
	// 含税保费，精确到小数点后两位
	PreMium *string `json:"pre_mium,omitempty" xml:"pre_mium,omitempty" require:"true"`
	// 全局唯一交易流水号
	TradeNo *string `json:"trade_no,omitempty" xml:"trade_no,omitempty" require:"true" maxLength:"32" minLength:"32"`
}

func (s CancelInsurancepolicyZhonghuacaixianRequest) String() string {
	return tea.Prettify(s)
}

func (s CancelInsurancepolicyZhonghuacaixianRequest) GoString() string {
	return s.String()
}

func (s *CancelInsurancepolicyZhonghuacaixianRequest) SetAuthToken(v string) *CancelInsurancepolicyZhonghuacaixianRequest {
	s.AuthToken = &v
	return s
}

func (s *CancelInsurancepolicyZhonghuacaixianRequest) SetProductInstanceId(v string) *CancelInsurancepolicyZhonghuacaixianRequest {
	s.ProductInstanceId = &v
	return s
}

func (s *CancelInsurancepolicyZhonghuacaixianRequest) SetPolNo(v string) *CancelInsurancepolicyZhonghuacaixianRequest {
	s.PolNo = &v
	return s
}

func (s *CancelInsurancepolicyZhonghuacaixianRequest) SetPreMium(v string) *CancelInsurancepolicyZhonghuacaixianRequest {
	s.PreMium = &v
	return s
}

func (s *CancelInsurancepolicyZhonghuacaixianRequest) SetTradeNo(v string) *CancelInsurancepolicyZhonghuacaixianRequest {
	s.TradeNo = &v
	return s
}

type CancelInsurancepolicyZhonghuacaixianResponse struct {
	// 请求唯一ID，用于链路跟踪和问题排查
	ReqMsgId *string `json:"req_msg_id,omitempty" xml:"req_msg_id,omitempty"`
	// 结果码，一般OK表示调用成功
	ResultCode *string `json:"result_code,omitempty" xml:"result_code,omitempty"`
	// 异常信息的文本描述
	ResultMsg *string `json:"result_msg,omitempty" xml:"result_msg,omitempty"`
	// 退保成功的对应的保单号
	PolNo *string `json:"pol_no,omitempty" xml:"pol_no,omitempty"`
}

func (s CancelInsurancepolicyZhonghuacaixianResponse) String() string {
	return tea.Prettify(s)
}

func (s CancelInsurancepolicyZhonghuacaixianResponse) GoString() string {
	return s.String()
}

func (s *CancelInsurancepolicyZhonghuacaixianResponse) SetReqMsgId(v string) *CancelInsurancepolicyZhonghuacaixianResponse {
	s.ReqMsgId = &v
	return s
}

func (s *CancelInsurancepolicyZhonghuacaixianResponse) SetResultCode(v string) *CancelInsurancepolicyZhonghuacaixianResponse {
	s.ResultCode = &v
	return s
}

func (s *CancelInsurancepolicyZhonghuacaixianResponse) SetResultMsg(v string) *CancelInsurancepolicyZhonghuacaixianResponse {
	s.ResultMsg = &v
	return s
}

func (s *CancelInsurancepolicyZhonghuacaixianResponse) SetPolNo(v string) *CancelInsurancepolicyZhonghuacaixianResponse {
	s.PolNo = &v
	return s
}

type PushInsurancenotifyClaimRequest struct {
	// OAuth模式下的授权token
	AuthToken         *string `json:"auth_token,omitempty" xml:"auth_token,omitempty"`
	ProductInstanceId *string `json:"product_instance_id,omitempty" xml:"product_instance_id,omitempty"`
	// 通知类型：
	// claims_pay:理赔支付信息
	// claims_info: 理赔信息
	Category *string `json:"category,omitempty" xml:"category,omitempty" require:"true" maxLength:"16"`
	// 通知内容
	Content *string `json:"content,omitempty" xml:"content,omitempty" require:"true" maxLength:"3000"`
}

func (s PushInsurancenotifyClaimRequest) String() string {
	return tea.Prettify(s)
}

func (s PushInsurancenotifyClaimRequest) GoString() string {
	return s.String()
}

func (s *PushInsurancenotifyClaimRequest) SetAuthToken(v string) *PushInsurancenotifyClaimRequest {
	s.AuthToken = &v
	return s
}

func (s *PushInsurancenotifyClaimRequest) SetProductInstanceId(v string) *PushInsurancenotifyClaimRequest {
	s.ProductInstanceId = &v
	return s
}

func (s *PushInsurancenotifyClaimRequest) SetCategory(v string) *PushInsurancenotifyClaimRequest {
	s.Category = &v
	return s
}

func (s *PushInsurancenotifyClaimRequest) SetContent(v string) *PushInsurancenotifyClaimRequest {
	s.Content = &v
	return s
}

type PushInsurancenotifyClaimResponse struct {
	// 请求唯一ID，用于链路跟踪和问题排查
	ReqMsgId *string `json:"req_msg_id,omitempty" xml:"req_msg_id,omitempty"`
	// 结果码，一般OK表示调用成功
	ResultCode *string `json:"result_code,omitempty" xml:"result_code,omitempty"`
	// 异常信息的文本描述
	ResultMsg *string `json:"result_msg,omitempty" xml:"result_msg,omitempty"`
	// 业务返回报文
	Data *string `json:"data,omitempty" xml:"data,omitempty"`
}

func (s PushInsurancenotifyClaimResponse) String() string {
	return tea.Prettify(s)
}

func (s PushInsurancenotifyClaimResponse) GoString() string {
	return s.String()
}

func (s *PushInsurancenotifyClaimResponse) SetReqMsgId(v string) *PushInsurancenotifyClaimResponse {
	s.ReqMsgId = &v
	return s
}

func (s *PushInsurancenotifyClaimResponse) SetResultCode(v string) *PushInsurancenotifyClaimResponse {
	s.ResultCode = &v
	return s
}

func (s *PushInsurancenotifyClaimResponse) SetResultMsg(v string) *PushInsurancenotifyClaimResponse {
	s.ResultMsg = &v
	return s
}

func (s *PushInsurancenotifyClaimResponse) SetData(v string) *PushInsurancenotifyClaimResponse {
	s.Data = &v
	return s
}

type ApplyInsuranceFileurlRequest struct {
	// OAuth模式下的授权token
	AuthToken         *string `json:"auth_token,omitempty" xml:"auth_token,omitempty"`
	ProductInstanceId *string `json:"product_instance_id,omitempty" xml:"product_instance_id,omitempty"`
	// 上传的文件名，要求文件名后缀必须以 _yyyyMMdd结尾
	FileName *string `json:"file_name,omitempty" xml:"file_name,omitempty" require:"true" maxLength:"80"`
}

func (s ApplyInsuranceFileurlRequest) String() string {
	return tea.Prettify(s)
}

func (s ApplyInsuranceFileurlRequest) GoString() string {
	return s.String()
}

func (s *ApplyInsuranceFileurlRequest) SetAuthToken(v string) *ApplyInsuranceFileurlRequest {
	s.AuthToken = &v
	return s
}

func (s *ApplyInsuranceFileurlRequest) SetProductInstanceId(v string) *ApplyInsuranceFileurlRequest {
	s.ProductInstanceId = &v
	return s
}

func (s *ApplyInsuranceFileurlRequest) SetFileName(v string) *ApplyInsuranceFileurlRequest {
	s.FileName = &v
	return s
}

type ApplyInsuranceFileurlResponse struct {
	// 请求唯一ID，用于链路跟踪和问题排查
	ReqMsgId *string `json:"req_msg_id,omitempty" xml:"req_msg_id,omitempty"`
	// 结果码，一般OK表示调用成功
	ResultCode *string `json:"result_code,omitempty" xml:"result_code,omitempty"`
	// 异常信息的文本描述
	ResultMsg *string `json:"result_msg,omitempty" xml:"result_msg,omitempty"`
	// 用于上传文件的连接
	FileUrl *string `json:"file_url,omitempty" xml:"file_url,omitempty"`
}

func (s ApplyInsuranceFileurlResponse) String() string {
	return tea.Prettify(s)
}

func (s ApplyInsuranceFileurlResponse) GoString() string {
	return s.String()
}

func (s *ApplyInsuranceFileurlResponse) SetReqMsgId(v string) *ApplyInsuranceFileurlResponse {
	s.ReqMsgId = &v
	return s
}

func (s *ApplyInsuranceFileurlResponse) SetResultCode(v string) *ApplyInsuranceFileurlResponse {
	s.ResultCode = &v
	return s
}

func (s *ApplyInsuranceFileurlResponse) SetResultMsg(v string) *ApplyInsuranceFileurlResponse {
	s.ResultMsg = &v
	return s
}

func (s *ApplyInsuranceFileurlResponse) SetFileUrl(v string) *ApplyInsuranceFileurlResponse {
	s.FileUrl = &v
	return s
}

type ApplyInsurancepolicyUniversalRequest struct {
	// OAuth模式下的授权token
	AuthToken         *string `json:"auth_token,omitempty" xml:"auth_token,omitempty"`
	ProductInstanceId *string `json:"product_instance_id,omitempty" xml:"product_instance_id,omitempty"`
	// 投保人开户银行，当打印发票选择为“1”时，必填
	AccountBankName *string `json:"account_bank_name,omitempty" xml:"account_bank_name,omitempty" maxLength:"100"`
	// 投保人银行账户，当打印发票选择为“1”时，必填
	BankAccountInformation *string `json:"bank_account_information,omitempty" xml:"bank_account_information,omitempty" maxLength:"50"`
	// 被保人地址
	BbrAddr *string `json:"bbr_addr,omitempty" xml:"bbr_addr,omitempty" maxLength:"200"`
	// 被保人did（被保人所在链上分布式数字身份）
	BbrDid *string `json:"bbr_did,omitempty" xml:"bbr_did,omitempty" maxLength:"200"`
	// 被保人证件号
	BbrIdNo *string `json:"bbr_id_no,omitempty" xml:"bbr_id_no,omitempty" require:"true" maxLength:"50"`
	// 被保人证件类型，01-居民身份证, 02-驾驶执照, 03-营业执照, 99-其他
	//
	BbrIdType *string `json:"bbr_id_type,omitempty" xml:"bbr_id_type,omitempty" require:"true" maxLength:"2"`
	// 被保人姓名
	BbrName *string `json:"bbr_name,omitempty" xml:"bbr_name,omitempty" require:"true" maxLength:"100"`
	// 被保人电话号
	BbrTel *string `json:"bbr_tel,omitempty" xml:"bbr_tel,omitempty" require:"true" maxLength:"20"`
	// 被保人类型, 1-个人 2-团队
	BbrType *string `json:"bbr_type,omitempty" xml:"bbr_type,omitempty" require:"true" maxLength:"1"`
	// 货物名称
	CargoName *string `json:"cargo_name,omitempty" xml:"cargo_name,omitempty" maxLength:"100"`
	// 货物数量
	CargoQuantity *string `json:"cargo_quantity,omitempty" xml:"cargo_quantity,omitempty" require:"true" maxLength:"20"`
	// 货物数量单位
	CargoQuantityUnit *string `json:"cargo_quantity_unit,omitempty" xml:"cargo_quantity_unit,omitempty" require:"true" maxLength:"20"`
	// 货物吨位
	CargoTonnage *string `json:"cargo_tonnage,omitempty" xml:"cargo_tonnage,omitempty" maxLength:"20"`
	// 货物类型，普货，手机/电子产品/家用电器，精密仪器，家具，易碎品，大宗散货，生鲜（瓜果蔬菜），棉纱
	CargoType *string `json:"cargo_type,omitempty" xml:"cargo_type,omitempty" require:"true" maxLength:"50"`
	// 货物体积
	CargoVolume *string `json:"cargo_volume,omitempty" xml:"cargo_volume,omitempty" maxLength:"20"`
	// 货值
	CargoWorth *string `json:"cargo_worth,omitempty" xml:"cargo_worth,omitempty" require:"true" maxLength:"50"`
	// 车长
	CarLength *string `json:"car_length,omitempty" xml:"car_length,omitempty" maxLength:"10"`
	// 车型
	CarModel *string `json:"car_model,omitempty" xml:"car_model,omitempty" maxLength:"20"`
	// 目的地，格式为"省-市-区"
	Destination *string `json:"destination,omitempty" xml:"destination,omitempty" require:"true" maxLength:"100"`
	// 司机姓名
	DriverName *string `json:"driver_name,omitempty" xml:"driver_name,omitempty" require:"true" maxLength:"50"`
	// 司机联系方式
	DriverTel *string `json:"driver_tel,omitempty" xml:"driver_tel,omitempty" require:"true" maxLength:"20"`
	// 保险起始时间, 时间格式 "yyyy-MM-dd HH:mm:ss", 预计起运日期
	EffDate *string `json:"eff_date,omitempty" xml:"eff_date,omitempty" require:"true" pattern:"\\d{4}[-]\\d{1,2}[-]\\d{1,2}[T]\\d{2}:\\d{2}:\\d{2}([Z]|([\\.]\\d{1,9})?[\\+]\\d{2}[\\:]?\\d{2})"`
	// 保司, 01-华泰, 02-平安, 03-亚太
	ExternalChannelCode *string `json:"external_channel_code,omitempty" xml:"external_channel_code,omitempty" require:"true" maxLength:"2"`
	// 险种, 01-承运人平台责任险
	//
	ExternalProductCode *string `json:"external_product_code,omitempty" xml:"external_product_code,omitempty" require:"true" maxLength:"2"`
	// 车架号
	FrameNo *string `json:"frame_no,omitempty" xml:"frame_no,omitempty" maxLength:"50"`
	// 投保额，整数以元为单位
	//
	InsuredAmount *string `json:"insured_amount,omitempty" xml:"insured_amount,omitempty" require:"true" maxLength:"12"`
	// 车牌号
	LicenseNo *string `json:"license_no,omitempty" xml:"license_no,omitempty" require:"true" maxLength:"20"`
	// 车牌颜色，01-黄牌，02-蓝牌
	LicensePlateColor *string `json:"license_plate_color,omitempty" xml:"license_plate_color,omitempty" maxLength:"2"`
	// 税务登记证/纳税人识别号，当打印发票选择为“1”时，必填
	NsrIdentifier *string `json:"nsr_identifier,omitempty" xml:"nsr_identifier,omitempty" maxLength:"50"`
	// 接单时间, 时间格式 "yyyy-MM-dd HH:mm:ss"
	OrderTime *string `json:"order_time,omitempty" xml:"order_time,omitempty" require:"true" pattern:"\\d{4}[-]\\d{1,2}[-]\\d{1,2}[T]\\d{2}:\\d{2}:\\d{2}([Z]|([\\.]\\d{1,9})?[\\+]\\d{2}[\\:]?\\d{2})"`
	// 包装方式, 01-箱装、02-袋装、03-包装、04-桶装、05-捆扎包装、06-裸装、07-散装、08-其他包装
	Packing *string `json:"packing,omitempty" xml:"packing,omitempty" maxLength:"100"`
	// 是否打印发票，默认0-不开票  1-专票 2-电子发票
	PrintInvoice *string `json:"print_invoice,omitempty" xml:"print_invoice,omitempty" require:"true" maxLength:"1"`
	// 费率，按实际业务商讨费率执行, 但最终结果以保司计算结果为准，0到1之间，最多6位小数
	Rate *string `json:"rate,omitempty" xml:"rate,omitempty" require:"true" maxLength:"8"`
	// 起运地，格式为"省-市-区"
	StartPlace *string `json:"start_place,omitempty" xml:"start_place,omitempty" require:"true" maxLength:"100"`
	// 投保人地址，当打印发票选择为“1”时，必填
	TbrAddr *string `json:"tbr_addr,omitempty" xml:"tbr_addr,omitempty" maxLength:"200"`
	// 投保人-企业联系人联系方式
	TbrCorporateContract *string `json:"tbr_corporate_contract,omitempty" xml:"tbr_corporate_contract,omitempty" maxLength:"20"`
	// 投保人did（投保人所在链上分布式数字身份）
	TbrDid *string `json:"tbr_did,omitempty" xml:"tbr_did,omitempty" maxLength:"200"`
	// 投保人邮箱，用以接受电子发票的邮箱地址
	TbrEmail *string `json:"tbr_email,omitempty" xml:"tbr_email,omitempty" maxLength:"50"`
	// 投保人证件号
	TbrIdNo *string `json:"tbr_id_no,omitempty" xml:"tbr_id_no,omitempty" require:"true" maxLength:"50"`
	// 投保人证件类型，01-居民身份证, 02-驾驶执照, 03-营业执照, 99-其他
	TbrIdType *string `json:"tbr_id_type,omitempty" xml:"tbr_id_type,omitempty" require:"true" maxLength:"2"`
	// 投保人名称
	TbrName *string `json:"tbr_name,omitempty" xml:"tbr_name,omitempty" require:"true" maxLength:"100"`
	// 投保人电话号
	TbrTel *string `json:"tbr_tel,omitempty" xml:"tbr_tel,omitempty" require:"true" maxLength:"20"`
	// 投保人类型, 1-个人 2-团队
	TbrType *string `json:"tbr_type,omitempty" xml:"tbr_type,omitempty" require:"true" maxLength:"1"`
	// 保险截止时间, 时间格式 "yyyy-MM-dd HH:mm:ss", 预计到达日期，保险止期间隔为1-30天
	TermDate *string `json:"term_date,omitempty" xml:"term_date,omitempty" require:"true" pattern:"\\d{4}[-]\\d{1,2}[-]\\d{1,2}[T]\\d{2}:\\d{2}:\\d{2}([Z]|([\\.]\\d{1,9})?[\\+]\\d{2}[\\:]?\\d{2})"`
	// 专票邮寄地址
	TicketAddr *string `json:"ticket_addr,omitempty" xml:"ticket_addr,omitempty" maxLength:"200"`
	// 订单号, 年月日+唯一字符ID
	TradeNo *string `json:"trade_no,omitempty" xml:"trade_no,omitempty" require:"true" maxLength:"50"`
	// 中转地，格式为"省-市-区"
	TransitPoint *string `json:"transit_point,omitempty" xml:"transit_point,omitempty" maxLength:"100"`
	// 运单id（平台方运单id）
	WaybillId *string `json:"waybill_id,omitempty" xml:"waybill_id,omitempty" require:"true" maxLength:"50"`
}

func (s ApplyInsurancepolicyUniversalRequest) String() string {
	return tea.Prettify(s)
}

func (s ApplyInsurancepolicyUniversalRequest) GoString() string {
	return s.String()
}

func (s *ApplyInsurancepolicyUniversalRequest) SetAuthToken(v string) *ApplyInsurancepolicyUniversalRequest {
	s.AuthToken = &v
	return s
}

func (s *ApplyInsurancepolicyUniversalRequest) SetProductInstanceId(v string) *ApplyInsurancepolicyUniversalRequest {
	s.ProductInstanceId = &v
	return s
}

func (s *ApplyInsurancepolicyUniversalRequest) SetAccountBankName(v string) *ApplyInsurancepolicyUniversalRequest {
	s.AccountBankName = &v
	return s
}

func (s *ApplyInsurancepolicyUniversalRequest) SetBankAccountInformation(v string) *ApplyInsurancepolicyUniversalRequest {
	s.BankAccountInformation = &v
	return s
}

func (s *ApplyInsurancepolicyUniversalRequest) SetBbrAddr(v string) *ApplyInsurancepolicyUniversalRequest {
	s.BbrAddr = &v
	return s
}

func (s *ApplyInsurancepolicyUniversalRequest) SetBbrDid(v string) *ApplyInsurancepolicyUniversalRequest {
	s.BbrDid = &v
	return s
}

func (s *ApplyInsurancepolicyUniversalRequest) SetBbrIdNo(v string) *ApplyInsurancepolicyUniversalRequest {
	s.BbrIdNo = &v
	return s
}

func (s *ApplyInsurancepolicyUniversalRequest) SetBbrIdType(v string) *ApplyInsurancepolicyUniversalRequest {
	s.BbrIdType = &v
	return s
}

func (s *ApplyInsurancepolicyUniversalRequest) SetBbrName(v string) *ApplyInsurancepolicyUniversalRequest {
	s.BbrName = &v
	return s
}

func (s *ApplyInsurancepolicyUniversalRequest) SetBbrTel(v string) *ApplyInsurancepolicyUniversalRequest {
	s.BbrTel = &v
	return s
}

func (s *ApplyInsurancepolicyUniversalRequest) SetBbrType(v string) *ApplyInsurancepolicyUniversalRequest {
	s.BbrType = &v
	return s
}

func (s *ApplyInsurancepolicyUniversalRequest) SetCargoName(v string) *ApplyInsurancepolicyUniversalRequest {
	s.CargoName = &v
	return s
}

func (s *ApplyInsurancepolicyUniversalRequest) SetCargoQuantity(v string) *ApplyInsurancepolicyUniversalRequest {
	s.CargoQuantity = &v
	return s
}

func (s *ApplyInsurancepolicyUniversalRequest) SetCargoQuantityUnit(v string) *ApplyInsurancepolicyUniversalRequest {
	s.CargoQuantityUnit = &v
	return s
}

func (s *ApplyInsurancepolicyUniversalRequest) SetCargoTonnage(v string) *ApplyInsurancepolicyUniversalRequest {
	s.CargoTonnage = &v
	return s
}

func (s *ApplyInsurancepolicyUniversalRequest) SetCargoType(v string) *ApplyInsurancepolicyUniversalRequest {
	s.CargoType = &v
	return s
}

func (s *ApplyInsurancepolicyUniversalRequest) SetCargoVolume(v string) *ApplyInsurancepolicyUniversalRequest {
	s.CargoVolume = &v
	return s
}

func (s *ApplyInsurancepolicyUniversalRequest) SetCargoWorth(v string) *ApplyInsurancepolicyUniversalRequest {
	s.CargoWorth = &v
	return s
}

func (s *ApplyInsurancepolicyUniversalRequest) SetCarLength(v string) *ApplyInsurancepolicyUniversalRequest {
	s.CarLength = &v
	return s
}

func (s *ApplyInsurancepolicyUniversalRequest) SetCarModel(v string) *ApplyInsurancepolicyUniversalRequest {
	s.CarModel = &v
	return s
}

func (s *ApplyInsurancepolicyUniversalRequest) SetDestination(v string) *ApplyInsurancepolicyUniversalRequest {
	s.Destination = &v
	return s
}

func (s *ApplyInsurancepolicyUniversalRequest) SetDriverName(v string) *ApplyInsurancepolicyUniversalRequest {
	s.DriverName = &v
	return s
}

func (s *ApplyInsurancepolicyUniversalRequest) SetDriverTel(v string) *ApplyInsurancepolicyUniversalRequest {
	s.DriverTel = &v
	return s
}

func (s *ApplyInsurancepolicyUniversalRequest) SetEffDate(v string) *ApplyInsurancepolicyUniversalRequest {
	s.EffDate = &v
	return s
}

func (s *ApplyInsurancepolicyUniversalRequest) SetExternalChannelCode(v string) *ApplyInsurancepolicyUniversalRequest {
	s.ExternalChannelCode = &v
	return s
}

func (s *ApplyInsurancepolicyUniversalRequest) SetExternalProductCode(v string) *ApplyInsurancepolicyUniversalRequest {
	s.ExternalProductCode = &v
	return s
}

func (s *ApplyInsurancepolicyUniversalRequest) SetFrameNo(v string) *ApplyInsurancepolicyUniversalRequest {
	s.FrameNo = &v
	return s
}

func (s *ApplyInsurancepolicyUniversalRequest) SetInsuredAmount(v string) *ApplyInsurancepolicyUniversalRequest {
	s.InsuredAmount = &v
	return s
}

func (s *ApplyInsurancepolicyUniversalRequest) SetLicenseNo(v string) *ApplyInsurancepolicyUniversalRequest {
	s.LicenseNo = &v
	return s
}

func (s *ApplyInsurancepolicyUniversalRequest) SetLicensePlateColor(v string) *ApplyInsurancepolicyUniversalRequest {
	s.LicensePlateColor = &v
	return s
}

func (s *ApplyInsurancepolicyUniversalRequest) SetNsrIdentifier(v string) *ApplyInsurancepolicyUniversalRequest {
	s.NsrIdentifier = &v
	return s
}

func (s *ApplyInsurancepolicyUniversalRequest) SetOrderTime(v string) *ApplyInsurancepolicyUniversalRequest {
	s.OrderTime = &v
	return s
}

func (s *ApplyInsurancepolicyUniversalRequest) SetPacking(v string) *ApplyInsurancepolicyUniversalRequest {
	s.Packing = &v
	return s
}

func (s *ApplyInsurancepolicyUniversalRequest) SetPrintInvoice(v string) *ApplyInsurancepolicyUniversalRequest {
	s.PrintInvoice = &v
	return s
}

func (s *ApplyInsurancepolicyUniversalRequest) SetRate(v string) *ApplyInsurancepolicyUniversalRequest {
	s.Rate = &v
	return s
}

func (s *ApplyInsurancepolicyUniversalRequest) SetStartPlace(v string) *ApplyInsurancepolicyUniversalRequest {
	s.StartPlace = &v
	return s
}

func (s *ApplyInsurancepolicyUniversalRequest) SetTbrAddr(v string) *ApplyInsurancepolicyUniversalRequest {
	s.TbrAddr = &v
	return s
}

func (s *ApplyInsurancepolicyUniversalRequest) SetTbrCorporateContract(v string) *ApplyInsurancepolicyUniversalRequest {
	s.TbrCorporateContract = &v
	return s
}

func (s *ApplyInsurancepolicyUniversalRequest) SetTbrDid(v string) *ApplyInsurancepolicyUniversalRequest {
	s.TbrDid = &v
	return s
}

func (s *ApplyInsurancepolicyUniversalRequest) SetTbrEmail(v string) *ApplyInsurancepolicyUniversalRequest {
	s.TbrEmail = &v
	return s
}

func (s *ApplyInsurancepolicyUniversalRequest) SetTbrIdNo(v string) *ApplyInsurancepolicyUniversalRequest {
	s.TbrIdNo = &v
	return s
}

func (s *ApplyInsurancepolicyUniversalRequest) SetTbrIdType(v string) *ApplyInsurancepolicyUniversalRequest {
	s.TbrIdType = &v
	return s
}

func (s *ApplyInsurancepolicyUniversalRequest) SetTbrName(v string) *ApplyInsurancepolicyUniversalRequest {
	s.TbrName = &v
	return s
}

func (s *ApplyInsurancepolicyUniversalRequest) SetTbrTel(v string) *ApplyInsurancepolicyUniversalRequest {
	s.TbrTel = &v
	return s
}

func (s *ApplyInsurancepolicyUniversalRequest) SetTbrType(v string) *ApplyInsurancepolicyUniversalRequest {
	s.TbrType = &v
	return s
}

func (s *ApplyInsurancepolicyUniversalRequest) SetTermDate(v string) *ApplyInsurancepolicyUniversalRequest {
	s.TermDate = &v
	return s
}

func (s *ApplyInsurancepolicyUniversalRequest) SetTicketAddr(v string) *ApplyInsurancepolicyUniversalRequest {
	s.TicketAddr = &v
	return s
}

func (s *ApplyInsurancepolicyUniversalRequest) SetTradeNo(v string) *ApplyInsurancepolicyUniversalRequest {
	s.TradeNo = &v
	return s
}

func (s *ApplyInsurancepolicyUniversalRequest) SetTransitPoint(v string) *ApplyInsurancepolicyUniversalRequest {
	s.TransitPoint = &v
	return s
}

func (s *ApplyInsurancepolicyUniversalRequest) SetWaybillId(v string) *ApplyInsurancepolicyUniversalRequest {
	s.WaybillId = &v
	return s
}

type ApplyInsurancepolicyUniversalResponse struct {
	// 请求唯一ID，用于链路跟踪和问题排查
	ReqMsgId *string `json:"req_msg_id,omitempty" xml:"req_msg_id,omitempty"`
	// 结果码，一般OK表示调用成功
	ResultCode *string `json:"result_code,omitempty" xml:"result_code,omitempty"`
	// 异常信息的文本描述
	ResultMsg *string `json:"result_msg,omitempty" xml:"result_msg,omitempty"`
	// 保单号
	PolNo *string `json:"pol_no,omitempty" xml:"pol_no,omitempty"`
	// 电子保单下载地址
	PolUrl *string `json:"pol_url,omitempty" xml:"pol_url,omitempty"`
	// 总保费
	TotalPremium *string `json:"total_premium,omitempty" xml:"total_premium,omitempty"`
	// 订单号
	TradeNo *string `json:"trade_no,omitempty" xml:"trade_no,omitempty"`
}

func (s ApplyInsurancepolicyUniversalResponse) String() string {
	return tea.Prettify(s)
}

func (s ApplyInsurancepolicyUniversalResponse) GoString() string {
	return s.String()
}

func (s *ApplyInsurancepolicyUniversalResponse) SetReqMsgId(v string) *ApplyInsurancepolicyUniversalResponse {
	s.ReqMsgId = &v
	return s
}

func (s *ApplyInsurancepolicyUniversalResponse) SetResultCode(v string) *ApplyInsurancepolicyUniversalResponse {
	s.ResultCode = &v
	return s
}

func (s *ApplyInsurancepolicyUniversalResponse) SetResultMsg(v string) *ApplyInsurancepolicyUniversalResponse {
	s.ResultMsg = &v
	return s
}

func (s *ApplyInsurancepolicyUniversalResponse) SetPolNo(v string) *ApplyInsurancepolicyUniversalResponse {
	s.PolNo = &v
	return s
}

func (s *ApplyInsurancepolicyUniversalResponse) SetPolUrl(v string) *ApplyInsurancepolicyUniversalResponse {
	s.PolUrl = &v
	return s
}

func (s *ApplyInsurancepolicyUniversalResponse) SetTotalPremium(v string) *ApplyInsurancepolicyUniversalResponse {
	s.TotalPremium = &v
	return s
}

func (s *ApplyInsurancepolicyUniversalResponse) SetTradeNo(v string) *ApplyInsurancepolicyUniversalResponse {
	s.TradeNo = &v
	return s
}

type CancelInsurancepolicyUniversalRequest struct {
	// OAuth模式下的授权token
	AuthToken         *string `json:"auth_token,omitempty" xml:"auth_token,omitempty"`
	ProductInstanceId *string `json:"product_instance_id,omitempty" xml:"product_instance_id,omitempty"`
	// 投保订单号
	ApplyTradeNo *string `json:"apply_trade_no,omitempty" xml:"apply_trade_no,omitempty" require:"true" maxLength:"50"`
	// 保司, 01-华泰, 03-亚太
	ExternalChannelCode *string `json:"external_channel_code,omitempty" xml:"external_channel_code,omitempty" require:"true" maxLength:"2"`
	// 险种, 01-承运人平台责任险
	//
	ExternalProductCode *string `json:"external_product_code,omitempty" xml:"external_product_code,omitempty" require:"true" maxLength:"2"`
	// 保单号
	PolNo *string `json:"pol_no,omitempty" xml:"pol_no,omitempty" require:"true" maxLength:"50"`
	// 退保全局唯一流水号
	TradeNo *string `json:"trade_no,omitempty" xml:"trade_no,omitempty" require:"true" maxLength:"50"`
}

func (s CancelInsurancepolicyUniversalRequest) String() string {
	return tea.Prettify(s)
}

func (s CancelInsurancepolicyUniversalRequest) GoString() string {
	return s.String()
}

func (s *CancelInsurancepolicyUniversalRequest) SetAuthToken(v string) *CancelInsurancepolicyUniversalRequest {
	s.AuthToken = &v
	return s
}

func (s *CancelInsurancepolicyUniversalRequest) SetProductInstanceId(v string) *CancelInsurancepolicyUniversalRequest {
	s.ProductInstanceId = &v
	return s
}

func (s *CancelInsurancepolicyUniversalRequest) SetApplyTradeNo(v string) *CancelInsurancepolicyUniversalRequest {
	s.ApplyTradeNo = &v
	return s
}

func (s *CancelInsurancepolicyUniversalRequest) SetExternalChannelCode(v string) *CancelInsurancepolicyUniversalRequest {
	s.ExternalChannelCode = &v
	return s
}

func (s *CancelInsurancepolicyUniversalRequest) SetExternalProductCode(v string) *CancelInsurancepolicyUniversalRequest {
	s.ExternalProductCode = &v
	return s
}

func (s *CancelInsurancepolicyUniversalRequest) SetPolNo(v string) *CancelInsurancepolicyUniversalRequest {
	s.PolNo = &v
	return s
}

func (s *CancelInsurancepolicyUniversalRequest) SetTradeNo(v string) *CancelInsurancepolicyUniversalRequest {
	s.TradeNo = &v
	return s
}

type CancelInsurancepolicyUniversalResponse struct {
	// 请求唯一ID，用于链路跟踪和问题排查
	ReqMsgId *string `json:"req_msg_id,omitempty" xml:"req_msg_id,omitempty"`
	// 结果码，一般OK表示调用成功
	ResultCode *string `json:"result_code,omitempty" xml:"result_code,omitempty"`
	// 异常信息的文本描述
	ResultMsg *string `json:"result_msg,omitempty" xml:"result_msg,omitempty"`
	// 投保订单号
	ApplyTradeNo *string `json:"apply_trade_no,omitempty" xml:"apply_trade_no,omitempty"`
	// 保单号
	PolNo *string `json:"pol_no,omitempty" xml:"pol_no,omitempty"`
	// 退还保费
	SendBackPremium *string `json:"send_back_premium,omitempty" xml:"send_back_premium,omitempty"`
	// 成功退保时间, 时间格式 "yyyy-MM-dd HH:mm:ss"
	SuccessSurrenderTime *string `json:"success_surrender_time,omitempty" xml:"success_surrender_time,omitempty"`
	// 退保生效时间, 时间格式 "yyyy-MM-dd HH:mm:ss"
	SurrenderEffectiveTime *string `json:"surrender_effective_time,omitempty" xml:"surrender_effective_time,omitempty"`
}

func (s CancelInsurancepolicyUniversalResponse) String() string {
	return tea.Prettify(s)
}

func (s CancelInsurancepolicyUniversalResponse) GoString() string {
	return s.String()
}

func (s *CancelInsurancepolicyUniversalResponse) SetReqMsgId(v string) *CancelInsurancepolicyUniversalResponse {
	s.ReqMsgId = &v
	return s
}

func (s *CancelInsurancepolicyUniversalResponse) SetResultCode(v string) *CancelInsurancepolicyUniversalResponse {
	s.ResultCode = &v
	return s
}

func (s *CancelInsurancepolicyUniversalResponse) SetResultMsg(v string) *CancelInsurancepolicyUniversalResponse {
	s.ResultMsg = &v
	return s
}

func (s *CancelInsurancepolicyUniversalResponse) SetApplyTradeNo(v string) *CancelInsurancepolicyUniversalResponse {
	s.ApplyTradeNo = &v
	return s
}

func (s *CancelInsurancepolicyUniversalResponse) SetPolNo(v string) *CancelInsurancepolicyUniversalResponse {
	s.PolNo = &v
	return s
}

func (s *CancelInsurancepolicyUniversalResponse) SetSendBackPremium(v string) *CancelInsurancepolicyUniversalResponse {
	s.SendBackPremium = &v
	return s
}

func (s *CancelInsurancepolicyUniversalResponse) SetSuccessSurrenderTime(v string) *CancelInsurancepolicyUniversalResponse {
	s.SuccessSurrenderTime = &v
	return s
}

func (s *CancelInsurancepolicyUniversalResponse) SetSurrenderEffectiveTime(v string) *CancelInsurancepolicyUniversalResponse {
	s.SurrenderEffectiveTime = &v
	return s
}

type PushInsuranceOlpRequest struct {
	// OAuth模式下的授权token
	AuthToken         *string `json:"auth_token,omitempty" xml:"auth_token,omitempty"`
	ProductInstanceId *string `json:"product_instance_id,omitempty" xml:"product_instance_id,omitempty"`
	// 被保人证件号
	BbrIdNo *string `json:"bbr_id_no,omitempty" xml:"bbr_id_no,omitempty" require:"true" maxLength:"50"`
	// 被保人名称
	BbrName *string `json:"bbr_name,omitempty" xml:"bbr_name,omitempty" require:"true" maxLength:"100"`
	// 被保人类型, 1-个人 2-企业
	//
	BbrType *string `json:"bbr_type,omitempty" xml:"bbr_type,omitempty" require:"true" maxLength:"10"`
	// 保险公司编码
	ChannelCode *string `json:"channel_code,omitempty" xml:"channel_code,omitempty" require:"true" maxLength:"50"`
	// 保险公司名称
	ChannelName *string `json:"channel_name,omitempty" xml:"channel_name,omitempty" require:"true" maxLength:"100"`
	// 理赔单状态，ALREADY_RISK: 已出险，NOT_RISK: 未出险，CLAIMED: 已理赔
	ClaimStatus *string `json:"claim_status,omitempty" xml:"claim_status,omitempty" require:"true" maxLength:"12"`
	// 投保时间
	InsureDate *string `json:"insure_date,omitempty" xml:"insure_date,omitempty" require:"true" pattern:"\\d{4}[-]\\d{1,2}[-]\\d{1,2}[T]\\d{2}:\\d{2}:\\d{2}([Z]|([\\.]\\d{1,9})?[\\+]\\d{2}[\\:]?\\d{2})"`
	// 保险止期
	InsureEnd *string `json:"insure_end,omitempty" xml:"insure_end,omitempty" pattern:"\\d{4}[-]\\d{1,2}[-]\\d{1,2}[T]\\d{2}:\\d{2}:\\d{2}([Z]|([\\.]\\d{1,9})?[\\+]\\d{2}[\\:]?\\d{2})"`
	// 保险起期
	InsureStart *string `json:"insure_start,omitempty" xml:"insure_start,omitempty" pattern:"\\d{4}[-]\\d{1,2}[-]\\d{1,2}[T]\\d{2}:\\d{2}:\\d{2}([Z]|([\\.]\\d{1,9})?[\\+]\\d{2}[\\:]?\\d{2})"`
	// 存证平台编码
	PlatformConsumerCode *string `json:"platform_consumer_code,omitempty" xml:"platform_consumer_code,omitempty" require:"true" maxLength:"50"`
	// 存证平台名称
	PlatformConsumerName *string `json:"platform_consumer_name,omitempty" xml:"platform_consumer_name,omitempty" require:"true" maxLength:"100"`
	// 保单文件id，可支持多个，逗号隔开
	PolFileIds *string `json:"pol_file_ids,omitempty" xml:"pol_file_ids,omitempty" maxLength:"3200"`
	// 保单号
	PolNo *string `json:"pol_no,omitempty" xml:"pol_no,omitempty" require:"true" maxLength:"100"`
	// 保单状态，INSURED: 已投保，SURRENDERED: 已退保,
	PolStatus *string `json:"pol_status,omitempty" xml:"pol_status,omitempty" require:"true" maxLength:"12"`
	// 电子保单url地址
	PolUrl *string `json:"pol_url,omitempty" xml:"pol_url,omitempty" maxLength:"500"`
	// 保费
	Premium *string `json:"premium,omitempty" xml:"premium,omitempty" require:"true" maxLength:"20"`
	// 险种代码
	ProductCode *string `json:"product_code,omitempty" xml:"product_code,omitempty" require:"true" maxLength:"50"`
	// 险种名称
	ProductName *string `json:"product_name,omitempty" xml:"product_name,omitempty" require:"true" maxLength:"100"`
	// 退保时间
	SurrenderTime *string `json:"surrender_time,omitempty" xml:"surrender_time,omitempty" pattern:"\\d{4}[-]\\d{1,2}[-]\\d{1,2}[T]\\d{2}:\\d{2}:\\d{2}([Z]|([\\.]\\d{1,9})?[\\+]\\d{2}[\\:]?\\d{2})"`
	// 投保人证件号
	TbrIdNo *string `json:"tbr_id_no,omitempty" xml:"tbr_id_no,omitempty" require:"true" maxLength:"50"`
	// 投保人名称
	TbrName *string `json:"tbr_name,omitempty" xml:"tbr_name,omitempty" require:"true" maxLength:"100"`
	// 投保人类型, 1-个人 2-企业
	TbrType *string `json:"tbr_type,omitempty" xml:"tbr_type,omitempty" require:"true" maxLength:"10"`
	// 交易流水号
	TradeNo *string `json:"trade_no,omitempty" xml:"trade_no,omitempty" maxLength:"200"`
}

func (s PushInsuranceOlpRequest) String() string {
	return tea.Prettify(s)
}

func (s PushInsuranceOlpRequest) GoString() string {
	return s.String()
}

func (s *PushInsuranceOlpRequest) SetAuthToken(v string) *PushInsuranceOlpRequest {
	s.AuthToken = &v
	return s
}

func (s *PushInsuranceOlpRequest) SetProductInstanceId(v string) *PushInsuranceOlpRequest {
	s.ProductInstanceId = &v
	return s
}

func (s *PushInsuranceOlpRequest) SetBbrIdNo(v string) *PushInsuranceOlpRequest {
	s.BbrIdNo = &v
	return s
}

func (s *PushInsuranceOlpRequest) SetBbrName(v string) *PushInsuranceOlpRequest {
	s.BbrName = &v
	return s
}

func (s *PushInsuranceOlpRequest) SetBbrType(v string) *PushInsuranceOlpRequest {
	s.BbrType = &v
	return s
}

func (s *PushInsuranceOlpRequest) SetChannelCode(v string) *PushInsuranceOlpRequest {
	s.ChannelCode = &v
	return s
}

func (s *PushInsuranceOlpRequest) SetChannelName(v string) *PushInsuranceOlpRequest {
	s.ChannelName = &v
	return s
}

func (s *PushInsuranceOlpRequest) SetClaimStatus(v string) *PushInsuranceOlpRequest {
	s.ClaimStatus = &v
	return s
}

func (s *PushInsuranceOlpRequest) SetInsureDate(v string) *PushInsuranceOlpRequest {
	s.InsureDate = &v
	return s
}

func (s *PushInsuranceOlpRequest) SetInsureEnd(v string) *PushInsuranceOlpRequest {
	s.InsureEnd = &v
	return s
}

func (s *PushInsuranceOlpRequest) SetInsureStart(v string) *PushInsuranceOlpRequest {
	s.InsureStart = &v
	return s
}

func (s *PushInsuranceOlpRequest) SetPlatformConsumerCode(v string) *PushInsuranceOlpRequest {
	s.PlatformConsumerCode = &v
	return s
}

func (s *PushInsuranceOlpRequest) SetPlatformConsumerName(v string) *PushInsuranceOlpRequest {
	s.PlatformConsumerName = &v
	return s
}

func (s *PushInsuranceOlpRequest) SetPolFileIds(v string) *PushInsuranceOlpRequest {
	s.PolFileIds = &v
	return s
}

func (s *PushInsuranceOlpRequest) SetPolNo(v string) *PushInsuranceOlpRequest {
	s.PolNo = &v
	return s
}

func (s *PushInsuranceOlpRequest) SetPolStatus(v string) *PushInsuranceOlpRequest {
	s.PolStatus = &v
	return s
}

func (s *PushInsuranceOlpRequest) SetPolUrl(v string) *PushInsuranceOlpRequest {
	s.PolUrl = &v
	return s
}

func (s *PushInsuranceOlpRequest) SetPremium(v string) *PushInsuranceOlpRequest {
	s.Premium = &v
	return s
}

func (s *PushInsuranceOlpRequest) SetProductCode(v string) *PushInsuranceOlpRequest {
	s.ProductCode = &v
	return s
}

func (s *PushInsuranceOlpRequest) SetProductName(v string) *PushInsuranceOlpRequest {
	s.ProductName = &v
	return s
}

func (s *PushInsuranceOlpRequest) SetSurrenderTime(v string) *PushInsuranceOlpRequest {
	s.SurrenderTime = &v
	return s
}

func (s *PushInsuranceOlpRequest) SetTbrIdNo(v string) *PushInsuranceOlpRequest {
	s.TbrIdNo = &v
	return s
}

func (s *PushInsuranceOlpRequest) SetTbrName(v string) *PushInsuranceOlpRequest {
	s.TbrName = &v
	return s
}

func (s *PushInsuranceOlpRequest) SetTbrType(v string) *PushInsuranceOlpRequest {
	s.TbrType = &v
	return s
}

func (s *PushInsuranceOlpRequest) SetTradeNo(v string) *PushInsuranceOlpRequest {
	s.TradeNo = &v
	return s
}

type PushInsuranceOlpResponse struct {
	// 请求唯一ID，用于链路跟踪和问题排查
	ReqMsgId *string `json:"req_msg_id,omitempty" xml:"req_msg_id,omitempty"`
	// 结果码，一般OK表示调用成功
	ResultCode *string `json:"result_code,omitempty" xml:"result_code,omitempty"`
	// 异常信息的文本描述
	ResultMsg *string `json:"result_msg,omitempty" xml:"result_msg,omitempty"`
	// 返回状态
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
}

func (s PushInsuranceOlpResponse) String() string {
	return tea.Prettify(s)
}

func (s PushInsuranceOlpResponse) GoString() string {
	return s.String()
}

func (s *PushInsuranceOlpResponse) SetReqMsgId(v string) *PushInsuranceOlpResponse {
	s.ReqMsgId = &v
	return s
}

func (s *PushInsuranceOlpResponse) SetResultCode(v string) *PushInsuranceOlpResponse {
	s.ResultCode = &v
	return s
}

func (s *PushInsuranceOlpResponse) SetResultMsg(v string) *PushInsuranceOlpResponse {
	s.ResultMsg = &v
	return s
}

func (s *PushInsuranceOlpResponse) SetStatus(v string) *PushInsuranceOlpResponse {
	s.Status = &v
	return s
}

type UpdateInsuranceOlpRequest struct {
	// OAuth模式下的授权token
	AuthToken         *string `json:"auth_token,omitempty" xml:"auth_token,omitempty"`
	ProductInstanceId *string `json:"product_instance_id,omitempty" xml:"product_instance_id,omitempty"`
	// 被保人证件号
	BbrIdNo *string `json:"bbr_id_no,omitempty" xml:"bbr_id_no,omitempty" maxLength:"50"`
	// 被保人名称
	BbrName *string `json:"bbr_name,omitempty" xml:"bbr_name,omitempty" maxLength:"100"`
	// 被保人类型, 1-个人 2-企业
	BbrType *string `json:"bbr_type,omitempty" xml:"bbr_type,omitempty" maxLength:"10"`
	// 保险公司编码
	ChannelCode *string `json:"channel_code,omitempty" xml:"channel_code,omitempty" maxLength:"50"`
	// 保险公司名称
	ChannelName *string `json:"channel_name,omitempty" xml:"channel_name,omitempty" maxLength:"100"`
	// 理赔单状态，ALREADY_RISK: 已出险，NOT_RISK: 未出险，CLAIMED: 已理赔
	ClaimStatus *string `json:"claim_status,omitempty" xml:"claim_status,omitempty" maxLength:"12"`
	// 投保时间
	InsureDate *string `json:"insure_date,omitempty" xml:"insure_date,omitempty" pattern:"\\d{4}[-]\\d{1,2}[-]\\d{1,2}[T]\\d{2}:\\d{2}:\\d{2}([Z]|([\\.]\\d{1,9})?[\\+]\\d{2}[\\:]?\\d{2})"`
	// 保险止期
	InsureEnd *string `json:"insure_end,omitempty" xml:"insure_end,omitempty" pattern:"\\d{4}[-]\\d{1,2}[-]\\d{1,2}[T]\\d{2}:\\d{2}:\\d{2}([Z]|([\\.]\\d{1,9})?[\\+]\\d{2}[\\:]?\\d{2})"`
	// 保险起期
	InsureStart *string `json:"insure_start,omitempty" xml:"insure_start,omitempty" pattern:"\\d{4}[-]\\d{1,2}[-]\\d{1,2}[T]\\d{2}:\\d{2}:\\d{2}([Z]|([\\.]\\d{1,9})?[\\+]\\d{2}[\\:]?\\d{2})"`
	// 存证平台编码
	PlatformConsumerCode *string `json:"platform_consumer_code,omitempty" xml:"platform_consumer_code,omitempty" require:"true" maxLength:"50"`
	// 存证平台名称
	PlatformConsumerName *string `json:"platform_consumer_name,omitempty" xml:"platform_consumer_name,omitempty" maxLength:"100"`
	// 保单文件id，可支持多个，逗号隔开
	PolFileIds *string `json:"pol_file_ids,omitempty" xml:"pol_file_ids,omitempty" maxLength:"3200"`
	// 保单号
	PolNo *string `json:"pol_no,omitempty" xml:"pol_no,omitempty" require:"true" maxLength:"100"`
	// 保单状态，INSURED: 已投保，SURRENDERED: 已退保,
	PolStatus *string `json:"pol_status,omitempty" xml:"pol_status,omitempty" maxLength:"12"`
	// 电子保单url地址
	PolUrl *string `json:"pol_url,omitempty" xml:"pol_url,omitempty" maxLength:"500"`
	// 保费
	//
	Premium *string `json:"premium,omitempty" xml:"premium,omitempty" maxLength:"20"`
	// 险种代码
	//
	ProductCode *string `json:"product_code,omitempty" xml:"product_code,omitempty" maxLength:"50"`
	// 险种名称
	//
	ProductName *string `json:"product_name,omitempty" xml:"product_name,omitempty" maxLength:"100"`
	// 退保时间
	SurrenderTime *string `json:"surrender_time,omitempty" xml:"surrender_time,omitempty" pattern:"\\d{4}[-]\\d{1,2}[-]\\d{1,2}[T]\\d{2}:\\d{2}:\\d{2}([Z]|([\\.]\\d{1,9})?[\\+]\\d{2}[\\:]?\\d{2})"`
	// 投保人证件号
	TbrIdNo *string `json:"tbr_id_no,omitempty" xml:"tbr_id_no,omitempty" maxLength:"50"`
	// 投保人名称
	TbrName *string `json:"tbr_name,omitempty" xml:"tbr_name,omitempty" maxLength:"100"`
	// 投保人类型, 1-个人 2-企业
	TbrType *string `json:"tbr_type,omitempty" xml:"tbr_type,omitempty" maxLength:"10"`
	// 交易流水号
	TradeNo *string `json:"trade_no,omitempty" xml:"trade_no,omitempty" maxLength:"200"`
}

func (s UpdateInsuranceOlpRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateInsuranceOlpRequest) GoString() string {
	return s.String()
}

func (s *UpdateInsuranceOlpRequest) SetAuthToken(v string) *UpdateInsuranceOlpRequest {
	s.AuthToken = &v
	return s
}

func (s *UpdateInsuranceOlpRequest) SetProductInstanceId(v string) *UpdateInsuranceOlpRequest {
	s.ProductInstanceId = &v
	return s
}

func (s *UpdateInsuranceOlpRequest) SetBbrIdNo(v string) *UpdateInsuranceOlpRequest {
	s.BbrIdNo = &v
	return s
}

func (s *UpdateInsuranceOlpRequest) SetBbrName(v string) *UpdateInsuranceOlpRequest {
	s.BbrName = &v
	return s
}

func (s *UpdateInsuranceOlpRequest) SetBbrType(v string) *UpdateInsuranceOlpRequest {
	s.BbrType = &v
	return s
}

func (s *UpdateInsuranceOlpRequest) SetChannelCode(v string) *UpdateInsuranceOlpRequest {
	s.ChannelCode = &v
	return s
}

func (s *UpdateInsuranceOlpRequest) SetChannelName(v string) *UpdateInsuranceOlpRequest {
	s.ChannelName = &v
	return s
}

func (s *UpdateInsuranceOlpRequest) SetClaimStatus(v string) *UpdateInsuranceOlpRequest {
	s.ClaimStatus = &v
	return s
}

func (s *UpdateInsuranceOlpRequest) SetInsureDate(v string) *UpdateInsuranceOlpRequest {
	s.InsureDate = &v
	return s
}

func (s *UpdateInsuranceOlpRequest) SetInsureEnd(v string) *UpdateInsuranceOlpRequest {
	s.InsureEnd = &v
	return s
}

func (s *UpdateInsuranceOlpRequest) SetInsureStart(v string) *UpdateInsuranceOlpRequest {
	s.InsureStart = &v
	return s
}

func (s *UpdateInsuranceOlpRequest) SetPlatformConsumerCode(v string) *UpdateInsuranceOlpRequest {
	s.PlatformConsumerCode = &v
	return s
}

func (s *UpdateInsuranceOlpRequest) SetPlatformConsumerName(v string) *UpdateInsuranceOlpRequest {
	s.PlatformConsumerName = &v
	return s
}

func (s *UpdateInsuranceOlpRequest) SetPolFileIds(v string) *UpdateInsuranceOlpRequest {
	s.PolFileIds = &v
	return s
}

func (s *UpdateInsuranceOlpRequest) SetPolNo(v string) *UpdateInsuranceOlpRequest {
	s.PolNo = &v
	return s
}

func (s *UpdateInsuranceOlpRequest) SetPolStatus(v string) *UpdateInsuranceOlpRequest {
	s.PolStatus = &v
	return s
}

func (s *UpdateInsuranceOlpRequest) SetPolUrl(v string) *UpdateInsuranceOlpRequest {
	s.PolUrl = &v
	return s
}

func (s *UpdateInsuranceOlpRequest) SetPremium(v string) *UpdateInsuranceOlpRequest {
	s.Premium = &v
	return s
}

func (s *UpdateInsuranceOlpRequest) SetProductCode(v string) *UpdateInsuranceOlpRequest {
	s.ProductCode = &v
	return s
}

func (s *UpdateInsuranceOlpRequest) SetProductName(v string) *UpdateInsuranceOlpRequest {
	s.ProductName = &v
	return s
}

func (s *UpdateInsuranceOlpRequest) SetSurrenderTime(v string) *UpdateInsuranceOlpRequest {
	s.SurrenderTime = &v
	return s
}

func (s *UpdateInsuranceOlpRequest) SetTbrIdNo(v string) *UpdateInsuranceOlpRequest {
	s.TbrIdNo = &v
	return s
}

func (s *UpdateInsuranceOlpRequest) SetTbrName(v string) *UpdateInsuranceOlpRequest {
	s.TbrName = &v
	return s
}

func (s *UpdateInsuranceOlpRequest) SetTbrType(v string) *UpdateInsuranceOlpRequest {
	s.TbrType = &v
	return s
}

func (s *UpdateInsuranceOlpRequest) SetTradeNo(v string) *UpdateInsuranceOlpRequest {
	s.TradeNo = &v
	return s
}

type UpdateInsuranceOlpResponse struct {
	// 请求唯一ID，用于链路跟踪和问题排查
	ReqMsgId *string `json:"req_msg_id,omitempty" xml:"req_msg_id,omitempty"`
	// 结果码，一般OK表示调用成功
	ResultCode *string `json:"result_code,omitempty" xml:"result_code,omitempty"`
	// 异常信息的文本描述
	ResultMsg *string `json:"result_msg,omitempty" xml:"result_msg,omitempty"`
	// 返回状态
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
}

func (s UpdateInsuranceOlpResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateInsuranceOlpResponse) GoString() string {
	return s.String()
}

func (s *UpdateInsuranceOlpResponse) SetReqMsgId(v string) *UpdateInsuranceOlpResponse {
	s.ReqMsgId = &v
	return s
}

func (s *UpdateInsuranceOlpResponse) SetResultCode(v string) *UpdateInsuranceOlpResponse {
	s.ResultCode = &v
	return s
}

func (s *UpdateInsuranceOlpResponse) SetResultMsg(v string) *UpdateInsuranceOlpResponse {
	s.ResultMsg = &v
	return s
}

func (s *UpdateInsuranceOlpResponse) SetStatus(v string) *UpdateInsuranceOlpResponse {
	s.Status = &v
	return s
}

type PushInsuranceReppolicyRequest struct {
	// OAuth模式下的授权token
	AuthToken         *string `json:"auth_token,omitempty" xml:"auth_token,omitempty"`
	ProductInstanceId *string `json:"product_instance_id,omitempty" xml:"product_instance_id,omitempty"`
	// 直保人租户ID
	PlatformCode *string `json:"platform_code,omitempty" xml:"platform_code,omitempty" require:"true" maxLength:"50"`
	// 直保人
	DirectInsurerName *string `json:"direct_insurer_name,omitempty" xml:"direct_insurer_name,omitempty" require:"true" maxLength:"200"`
	// 保单号
	PolNo *string `json:"pol_no,omitempty" xml:"pol_no,omitempty" require:"true" maxLength:"100"`
	// 共保标志，1:是 0:否
	JointInsuranceLogo *string `json:"joint_insurance_logo,omitempty" xml:"joint_insurance_logo,omitempty" maxLength:"5"`
	// 共保比例（%）
	JointInsuranceRate *string `json:"joint_insurance_rate,omitempty" xml:"joint_insurance_rate,omitempty" maxLength:"5"`
	// 险种代码
	ProductCode *string `json:"product_code,omitempty" xml:"product_code,omitempty" require:"true" maxLength:"20"`
	// 险种名称
	ProductName *string `json:"product_name,omitempty" xml:"product_name,omitempty" require:"true" maxLength:"100"`
	// 被保险人
	BbrName *string `json:"bbr_name,omitempty" xml:"bbr_name,omitempty" require:"true" maxLength:"200"`
	// 保险起期
	InsureStart *string `json:"insure_start,omitempty" xml:"insure_start,omitempty" require:"true" pattern:"\\d{4}[-]\\d{1,2}[-]\\d{1,2}[T]\\d{2}:\\d{2}:\\d{2}([Z]|([\\.]\\d{1,9})?[\\+]\\d{2}[\\:]?\\d{2})"`
	// 保险止期
	InsureEnd *string `json:"insure_end,omitempty" xml:"insure_end,omitempty" require:"true" pattern:"\\d{4}[-]\\d{1,2}[-]\\d{1,2}[T]\\d{2}:\\d{2}:\\d{2}([Z]|([\\.]\\d{1,9})?[\\+]\\d{2}[\\:]?\\d{2})"`
	// 币种，CNY:人民币, USD:美元，JPY:日元，HKD:港元，EUR:欧元，GBP英镑
	Currency *string `json:"currency,omitempty" xml:"currency,omitempty" require:"true" maxLength:"10"`
	// 保费
	Premium *string `json:"premium,omitempty" xml:"premium,omitempty" require:"true" maxLength:"20"`
	// 保额
	InsuranceAmount *string `json:"insurance_amount,omitempty" xml:"insurance_amount,omitempty" maxLength:"20"`
	// 固定保额分出比例（%）
	FixedInsuranceAmountSplitRate *string `json:"fixed_insurance_amount_split_rate,omitempty" xml:"fixed_insurance_amount_split_rate,omitempty" maxLength:"5"`
	// 固定保费分出比例（%）
	FixedPremiumSplitRate *string `json:"fixed_premium_split_rate,omitempty" xml:"fixed_premium_split_rate,omitempty" require:"true" maxLength:"5"`
	// 临分保额分出比例（%）
	FacultativeInsuranceAmountSplitRate *string `json:"facultative_insurance_amount_split_rate,omitempty" xml:"facultative_insurance_amount_split_rate,omitempty" maxLength:"5"`
	// 临分保费分出比例（%）
	FacultativePremiumSplitRate *string `json:"facultative_premium_split_rate,omitempty" xml:"facultative_premium_split_rate,omitempty" require:"true" maxLength:"5"`
	// 分出保费
	SplitPremium *string `json:"split_premium,omitempty" xml:"split_premium,omitempty" require:"true" maxLength:"20"`
	// 再保人
	Reinsurer *string `json:"reinsurer,omitempty" xml:"reinsurer,omitempty" require:"true" maxLength:"100"`
	// 摊回手续费比例（%）
	AmortizeFeeRate *string `json:"amortize_fee_rate,omitempty" xml:"amortize_fee_rate,omitempty" require:"true" maxLength:"5"`
	// 摊回手续费
	AmortizeFee *string `json:"amortize_fee,omitempty" xml:"amortize_fee,omitempty" require:"true" maxLength:"20"`
	// 再保合同名称
	ReinsuranceContractName *string `json:"reinsurance_contract_name,omitempty" xml:"reinsurance_contract_name,omitempty" maxLength:"200"`
	// 是否有超赔临分，1:是 0:否
	OverPayFacultative *string `json:"over_pay_facultative,omitempty" xml:"over_pay_facultative,omitempty" maxLength:"5"`
	// 项目名称
	ProjectName *string `json:"project_name,omitempty" xml:"project_name,omitempty" maxLength:"100"`
	// 接收方租户id
	ReceiverPlatformId *string `json:"receiver_platform_id,omitempty" xml:"receiver_platform_id,omitempty" require:"true" maxLength:"50"`
}

func (s PushInsuranceReppolicyRequest) String() string {
	return tea.Prettify(s)
}

func (s PushInsuranceReppolicyRequest) GoString() string {
	return s.String()
}

func (s *PushInsuranceReppolicyRequest) SetAuthToken(v string) *PushInsuranceReppolicyRequest {
	s.AuthToken = &v
	return s
}

func (s *PushInsuranceReppolicyRequest) SetProductInstanceId(v string) *PushInsuranceReppolicyRequest {
	s.ProductInstanceId = &v
	return s
}

func (s *PushInsuranceReppolicyRequest) SetPlatformCode(v string) *PushInsuranceReppolicyRequest {
	s.PlatformCode = &v
	return s
}

func (s *PushInsuranceReppolicyRequest) SetDirectInsurerName(v string) *PushInsuranceReppolicyRequest {
	s.DirectInsurerName = &v
	return s
}

func (s *PushInsuranceReppolicyRequest) SetPolNo(v string) *PushInsuranceReppolicyRequest {
	s.PolNo = &v
	return s
}

func (s *PushInsuranceReppolicyRequest) SetJointInsuranceLogo(v string) *PushInsuranceReppolicyRequest {
	s.JointInsuranceLogo = &v
	return s
}

func (s *PushInsuranceReppolicyRequest) SetJointInsuranceRate(v string) *PushInsuranceReppolicyRequest {
	s.JointInsuranceRate = &v
	return s
}

func (s *PushInsuranceReppolicyRequest) SetProductCode(v string) *PushInsuranceReppolicyRequest {
	s.ProductCode = &v
	return s
}

func (s *PushInsuranceReppolicyRequest) SetProductName(v string) *PushInsuranceReppolicyRequest {
	s.ProductName = &v
	return s
}

func (s *PushInsuranceReppolicyRequest) SetBbrName(v string) *PushInsuranceReppolicyRequest {
	s.BbrName = &v
	return s
}

func (s *PushInsuranceReppolicyRequest) SetInsureStart(v string) *PushInsuranceReppolicyRequest {
	s.InsureStart = &v
	return s
}

func (s *PushInsuranceReppolicyRequest) SetInsureEnd(v string) *PushInsuranceReppolicyRequest {
	s.InsureEnd = &v
	return s
}

func (s *PushInsuranceReppolicyRequest) SetCurrency(v string) *PushInsuranceReppolicyRequest {
	s.Currency = &v
	return s
}

func (s *PushInsuranceReppolicyRequest) SetPremium(v string) *PushInsuranceReppolicyRequest {
	s.Premium = &v
	return s
}

func (s *PushInsuranceReppolicyRequest) SetInsuranceAmount(v string) *PushInsuranceReppolicyRequest {
	s.InsuranceAmount = &v
	return s
}

func (s *PushInsuranceReppolicyRequest) SetFixedInsuranceAmountSplitRate(v string) *PushInsuranceReppolicyRequest {
	s.FixedInsuranceAmountSplitRate = &v
	return s
}

func (s *PushInsuranceReppolicyRequest) SetFixedPremiumSplitRate(v string) *PushInsuranceReppolicyRequest {
	s.FixedPremiumSplitRate = &v
	return s
}

func (s *PushInsuranceReppolicyRequest) SetFacultativeInsuranceAmountSplitRate(v string) *PushInsuranceReppolicyRequest {
	s.FacultativeInsuranceAmountSplitRate = &v
	return s
}

func (s *PushInsuranceReppolicyRequest) SetFacultativePremiumSplitRate(v string) *PushInsuranceReppolicyRequest {
	s.FacultativePremiumSplitRate = &v
	return s
}

func (s *PushInsuranceReppolicyRequest) SetSplitPremium(v string) *PushInsuranceReppolicyRequest {
	s.SplitPremium = &v
	return s
}

func (s *PushInsuranceReppolicyRequest) SetReinsurer(v string) *PushInsuranceReppolicyRequest {
	s.Reinsurer = &v
	return s
}

func (s *PushInsuranceReppolicyRequest) SetAmortizeFeeRate(v string) *PushInsuranceReppolicyRequest {
	s.AmortizeFeeRate = &v
	return s
}

func (s *PushInsuranceReppolicyRequest) SetAmortizeFee(v string) *PushInsuranceReppolicyRequest {
	s.AmortizeFee = &v
	return s
}

func (s *PushInsuranceReppolicyRequest) SetReinsuranceContractName(v string) *PushInsuranceReppolicyRequest {
	s.ReinsuranceContractName = &v
	return s
}

func (s *PushInsuranceReppolicyRequest) SetOverPayFacultative(v string) *PushInsuranceReppolicyRequest {
	s.OverPayFacultative = &v
	return s
}

func (s *PushInsuranceReppolicyRequest) SetProjectName(v string) *PushInsuranceReppolicyRequest {
	s.ProjectName = &v
	return s
}

func (s *PushInsuranceReppolicyRequest) SetReceiverPlatformId(v string) *PushInsuranceReppolicyRequest {
	s.ReceiverPlatformId = &v
	return s
}

type PushInsuranceReppolicyResponse struct {
	// 请求唯一ID，用于链路跟踪和问题排查
	ReqMsgId *string `json:"req_msg_id,omitempty" xml:"req_msg_id,omitempty"`
	// 结果码，一般OK表示调用成功
	ResultCode *string `json:"result_code,omitempty" xml:"result_code,omitempty"`
	// 异常信息的文本描述
	ResultMsg *string `json:"result_msg,omitempty" xml:"result_msg,omitempty"`
	// 返回状态
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
}

func (s PushInsuranceReppolicyResponse) String() string {
	return tea.Prettify(s)
}

func (s PushInsuranceReppolicyResponse) GoString() string {
	return s.String()
}

func (s *PushInsuranceReppolicyResponse) SetReqMsgId(v string) *PushInsuranceReppolicyResponse {
	s.ReqMsgId = &v
	return s
}

func (s *PushInsuranceReppolicyResponse) SetResultCode(v string) *PushInsuranceReppolicyResponse {
	s.ResultCode = &v
	return s
}

func (s *PushInsuranceReppolicyResponse) SetResultMsg(v string) *PushInsuranceReppolicyResponse {
	s.ResultMsg = &v
	return s
}

func (s *PushInsuranceReppolicyResponse) SetStatus(v string) *PushInsuranceReppolicyResponse {
	s.Status = &v
	return s
}

type PushInsuranceRepcorrectRequest struct {
	// OAuth模式下的授权token
	AuthToken         *string `json:"auth_token,omitempty" xml:"auth_token,omitempty"`
	ProductInstanceId *string `json:"product_instance_id,omitempty" xml:"product_instance_id,omitempty"`
	// 直保人租户ID
	PlatformCode *string `json:"platform_code,omitempty" xml:"platform_code,omitempty" require:"true" maxLength:"50"`
	// 直保人
	DirectInsurerName *string `json:"direct_insurer_name,omitempty" xml:"direct_insurer_name,omitempty" require:"true" maxLength:"200"`
	// 批单类型 1:批增 2: 批减
	ApprovalType *string `json:"approval_type,omitempty" xml:"approval_type,omitempty" require:"true" maxLength:"20"`
	// 批单号
	ApprovalNo *string `json:"approval_no,omitempty" xml:"approval_no,omitempty" require:"true" maxLength:"100"`
	// 保单号
	PolNo *string `json:"pol_no,omitempty" xml:"pol_no,omitempty" require:"true" maxLength:"100"`
	// 共保标志，1:是 0:否
	//
	JointInsuranceLogo *string `json:"joint_insurance_logo,omitempty" xml:"joint_insurance_logo,omitempty" maxLength:"5"`
	// 共保比例（%）
	JointInsuranceRate *string `json:"joint_insurance_rate,omitempty" xml:"joint_insurance_rate,omitempty" maxLength:"5"`
	// 险种代码
	ProductCode *string `json:"product_code,omitempty" xml:"product_code,omitempty" require:"true" maxLength:"20"`
	// 险种名称
	//
	ProductName *string `json:"product_name,omitempty" xml:"product_name,omitempty" require:"true" maxLength:"100"`
	// 被保险人
	BbrName *string `json:"bbr_name,omitempty" xml:"bbr_name,omitempty" require:"true" maxLength:"200"`
	// 保险起期
	InsureStart *string `json:"insure_start,omitempty" xml:"insure_start,omitempty" require:"true" pattern:"\\d{4}[-]\\d{1,2}[-]\\d{1,2}[T]\\d{2}:\\d{2}:\\d{2}([Z]|([\\.]\\d{1,9})?[\\+]\\d{2}[\\:]?\\d{2})"`
	// 保险止期
	//
	InsureEnd *string `json:"insure_end,omitempty" xml:"insure_end,omitempty" require:"true" pattern:"\\d{4}[-]\\d{1,2}[-]\\d{1,2}[T]\\d{2}:\\d{2}:\\d{2}([Z]|([\\.]\\d{1,9})?[\\+]\\d{2}[\\:]?\\d{2})"`
	// 币种，CNY:人民币, USD:美元，JPY:日元，HKD:港元，EUR:欧元，GBP英镑
	Currency *string `json:"currency,omitempty" xml:"currency,omitempty" require:"true" maxLength:"10"`
	// 保额
	InsuranceAmount *string `json:"insurance_amount,omitempty" xml:"insurance_amount,omitempty" maxLength:"20"`
	// 批改保费
	CorrectPremium *string `json:"correct_premium,omitempty" xml:"correct_premium,omitempty" require:"true" maxLength:"20"`
	// 固定保额分出比例（%）
	FixedInsuranceAmountSplitRate *string `json:"fixed_insurance_amount_split_rate,omitempty" xml:"fixed_insurance_amount_split_rate,omitempty" maxLength:"5"`
	// 固定保费分出比例（%）
	FixedPremiumSplitRate *string `json:"fixed_premium_split_rate,omitempty" xml:"fixed_premium_split_rate,omitempty" require:"true" maxLength:"5"`
	// 临分保额分出比例（%）
	FacultativeInsuranceAmountSplitRate *string `json:"facultative_insurance_amount_split_rate,omitempty" xml:"facultative_insurance_amount_split_rate,omitempty" maxLength:"5"`
	// 临分保费分出比例（%）
	FacultativePremiumSplitRate *string `json:"facultative_premium_split_rate,omitempty" xml:"facultative_premium_split_rate,omitempty" require:"true" maxLength:"5"`
	// 分出保费
	SplitPremium *string `json:"split_premium,omitempty" xml:"split_premium,omitempty" require:"true" maxLength:"20"`
	// 再保人
	//
	Reinsurer *string `json:"reinsurer,omitempty" xml:"reinsurer,omitempty" require:"true" maxLength:"100"`
	// 摊回手续费比例（%）
	AmortizeFeeRate *string `json:"amortize_fee_rate,omitempty" xml:"amortize_fee_rate,omitempty" require:"true" maxLength:"5"`
	// 摊回手续费
	//
	AmortizeFee *string `json:"amortize_fee,omitempty" xml:"amortize_fee,omitempty" require:"true" maxLength:"20"`
	// 再保合同名称
	ReinsuranceContractName *string `json:"reinsurance_contract_name,omitempty" xml:"reinsurance_contract_name,omitempty" maxLength:"200"`
	// 是否有超赔临分，1:是 0:否
	OverPayFacultative *string `json:"over_pay_facultative,omitempty" xml:"over_pay_facultative,omitempty" maxLength:"5"`
	// 批改原因
	//
	CorrectReason *string `json:"correct_reason,omitempty" xml:"correct_reason,omitempty" maxLength:"500"`
	// 项目名称
	ProjectName *string `json:"project_name,omitempty" xml:"project_name,omitempty" maxLength:"100"`
	// 接收方租户id
	//
	ReceiverPlatformId *string `json:"receiver_platform_id,omitempty" xml:"receiver_platform_id,omitempty" require:"true" maxLength:"50"`
}

func (s PushInsuranceRepcorrectRequest) String() string {
	return tea.Prettify(s)
}

func (s PushInsuranceRepcorrectRequest) GoString() string {
	return s.String()
}

func (s *PushInsuranceRepcorrectRequest) SetAuthToken(v string) *PushInsuranceRepcorrectRequest {
	s.AuthToken = &v
	return s
}

func (s *PushInsuranceRepcorrectRequest) SetProductInstanceId(v string) *PushInsuranceRepcorrectRequest {
	s.ProductInstanceId = &v
	return s
}

func (s *PushInsuranceRepcorrectRequest) SetPlatformCode(v string) *PushInsuranceRepcorrectRequest {
	s.PlatformCode = &v
	return s
}

func (s *PushInsuranceRepcorrectRequest) SetDirectInsurerName(v string) *PushInsuranceRepcorrectRequest {
	s.DirectInsurerName = &v
	return s
}

func (s *PushInsuranceRepcorrectRequest) SetApprovalType(v string) *PushInsuranceRepcorrectRequest {
	s.ApprovalType = &v
	return s
}

func (s *PushInsuranceRepcorrectRequest) SetApprovalNo(v string) *PushInsuranceRepcorrectRequest {
	s.ApprovalNo = &v
	return s
}

func (s *PushInsuranceRepcorrectRequest) SetPolNo(v string) *PushInsuranceRepcorrectRequest {
	s.PolNo = &v
	return s
}

func (s *PushInsuranceRepcorrectRequest) SetJointInsuranceLogo(v string) *PushInsuranceRepcorrectRequest {
	s.JointInsuranceLogo = &v
	return s
}

func (s *PushInsuranceRepcorrectRequest) SetJointInsuranceRate(v string) *PushInsuranceRepcorrectRequest {
	s.JointInsuranceRate = &v
	return s
}

func (s *PushInsuranceRepcorrectRequest) SetProductCode(v string) *PushInsuranceRepcorrectRequest {
	s.ProductCode = &v
	return s
}

func (s *PushInsuranceRepcorrectRequest) SetProductName(v string) *PushInsuranceRepcorrectRequest {
	s.ProductName = &v
	return s
}

func (s *PushInsuranceRepcorrectRequest) SetBbrName(v string) *PushInsuranceRepcorrectRequest {
	s.BbrName = &v
	return s
}

func (s *PushInsuranceRepcorrectRequest) SetInsureStart(v string) *PushInsuranceRepcorrectRequest {
	s.InsureStart = &v
	return s
}

func (s *PushInsuranceRepcorrectRequest) SetInsureEnd(v string) *PushInsuranceRepcorrectRequest {
	s.InsureEnd = &v
	return s
}

func (s *PushInsuranceRepcorrectRequest) SetCurrency(v string) *PushInsuranceRepcorrectRequest {
	s.Currency = &v
	return s
}

func (s *PushInsuranceRepcorrectRequest) SetInsuranceAmount(v string) *PushInsuranceRepcorrectRequest {
	s.InsuranceAmount = &v
	return s
}

func (s *PushInsuranceRepcorrectRequest) SetCorrectPremium(v string) *PushInsuranceRepcorrectRequest {
	s.CorrectPremium = &v
	return s
}

func (s *PushInsuranceRepcorrectRequest) SetFixedInsuranceAmountSplitRate(v string) *PushInsuranceRepcorrectRequest {
	s.FixedInsuranceAmountSplitRate = &v
	return s
}

func (s *PushInsuranceRepcorrectRequest) SetFixedPremiumSplitRate(v string) *PushInsuranceRepcorrectRequest {
	s.FixedPremiumSplitRate = &v
	return s
}

func (s *PushInsuranceRepcorrectRequest) SetFacultativeInsuranceAmountSplitRate(v string) *PushInsuranceRepcorrectRequest {
	s.FacultativeInsuranceAmountSplitRate = &v
	return s
}

func (s *PushInsuranceRepcorrectRequest) SetFacultativePremiumSplitRate(v string) *PushInsuranceRepcorrectRequest {
	s.FacultativePremiumSplitRate = &v
	return s
}

func (s *PushInsuranceRepcorrectRequest) SetSplitPremium(v string) *PushInsuranceRepcorrectRequest {
	s.SplitPremium = &v
	return s
}

func (s *PushInsuranceRepcorrectRequest) SetReinsurer(v string) *PushInsuranceRepcorrectRequest {
	s.Reinsurer = &v
	return s
}

func (s *PushInsuranceRepcorrectRequest) SetAmortizeFeeRate(v string) *PushInsuranceRepcorrectRequest {
	s.AmortizeFeeRate = &v
	return s
}

func (s *PushInsuranceRepcorrectRequest) SetAmortizeFee(v string) *PushInsuranceRepcorrectRequest {
	s.AmortizeFee = &v
	return s
}

func (s *PushInsuranceRepcorrectRequest) SetReinsuranceContractName(v string) *PushInsuranceRepcorrectRequest {
	s.ReinsuranceContractName = &v
	return s
}

func (s *PushInsuranceRepcorrectRequest) SetOverPayFacultative(v string) *PushInsuranceRepcorrectRequest {
	s.OverPayFacultative = &v
	return s
}

func (s *PushInsuranceRepcorrectRequest) SetCorrectReason(v string) *PushInsuranceRepcorrectRequest {
	s.CorrectReason = &v
	return s
}

func (s *PushInsuranceRepcorrectRequest) SetProjectName(v string) *PushInsuranceRepcorrectRequest {
	s.ProjectName = &v
	return s
}

func (s *PushInsuranceRepcorrectRequest) SetReceiverPlatformId(v string) *PushInsuranceRepcorrectRequest {
	s.ReceiverPlatformId = &v
	return s
}

type PushInsuranceRepcorrectResponse struct {
	// 请求唯一ID，用于链路跟踪和问题排查
	ReqMsgId *string `json:"req_msg_id,omitempty" xml:"req_msg_id,omitempty"`
	// 结果码，一般OK表示调用成功
	ResultCode *string `json:"result_code,omitempty" xml:"result_code,omitempty"`
	// 异常信息的文本描述
	ResultMsg *string `json:"result_msg,omitempty" xml:"result_msg,omitempty"`
	// 返回状态
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
}

func (s PushInsuranceRepcorrectResponse) String() string {
	return tea.Prettify(s)
}

func (s PushInsuranceRepcorrectResponse) GoString() string {
	return s.String()
}

func (s *PushInsuranceRepcorrectResponse) SetReqMsgId(v string) *PushInsuranceRepcorrectResponse {
	s.ReqMsgId = &v
	return s
}

func (s *PushInsuranceRepcorrectResponse) SetResultCode(v string) *PushInsuranceRepcorrectResponse {
	s.ResultCode = &v
	return s
}

func (s *PushInsuranceRepcorrectResponse) SetResultMsg(v string) *PushInsuranceRepcorrectResponse {
	s.ResultMsg = &v
	return s
}

func (s *PushInsuranceRepcorrectResponse) SetStatus(v string) *PushInsuranceRepcorrectResponse {
	s.Status = &v
	return s
}

type ApplyInsuranceCbrfRequest struct {
	// OAuth模式下的授权token
	AuthToken         *string `json:"auth_token,omitempty" xml:"auth_token,omitempty"`
	ProductInstanceId *string `json:"product_instance_id,omitempty" xml:"product_instance_id,omitempty"`
	// 物流公司业务线的简称
	BusinessId *string `json:"business_id,omitempty" xml:"business_id,omitempty" maxLength:"30" minLength:"0"`
	// 保司编码
	InsuranceCode *string `json:"insurance_code,omitempty" xml:"insurance_code,omitempty" require:"true" maxLength:"8"`
	// 险种编码
	ProductCode *string `json:"product_code,omitempty" xml:"product_code,omitempty" require:"true" maxLength:"2"`
	// 起保时间
	InsuranceStart *string `json:"insurance_start,omitempty" xml:"insurance_start,omitempty" require:"true" pattern:"\\d{4}[-]\\d{1,2}[-]\\d{1,2}[T]\\d{2}:\\d{2}:\\d{2}([Z]|([\\.]\\d{1,9})?[\\+]\\d{2}[\\:]?\\d{2})"`
	// 交付航司确认时间
	DeliveryTime *string `json:"delivery_time,omitempty" xml:"delivery_time,omitempty" require:"true" pattern:"\\d{4}[-]\\d{1,2}[-]\\d{1,2}[T]\\d{2}:\\d{2}:\\d{2}([Z]|([\\.]\\d{1,9})?[\\+]\\d{2}[\\:]?\\d{2})"`
	// 客户订单编号
	RelatedOderNo *string `json:"related_oder_no,omitempty" xml:"related_oder_no,omitempty" require:"true" maxLength:"256" minLength:"2"`
	// 目的国，2位iso缩写
	DestCountry *string `json:"dest_country,omitempty" xml:"dest_country,omitempty" require:"true" maxLength:"256" minLength:"2"`
	// 商家唯一脱敏的编码
	BusinessMerchantId *string `json:"business_merchant_id,omitempty" xml:"business_merchant_id,omitempty" maxLength:"256"`
	// 货物的揽收时间
	CollectionTime *string `json:"collection_time,omitempty" xml:"collection_time,omitempty" maxLength:"50"`
	// 货物名称
	GoodName *string `json:"good_name,omitempty" xml:"good_name,omitempty" maxLength:"200"`
	// 货值(美金)，货物的美金商品价值
	GoodValue *string `json:"good_value,omitempty" xml:"good_value,omitempty" maxLength:"50"`
}

func (s ApplyInsuranceCbrfRequest) String() string {
	return tea.Prettify(s)
}

func (s ApplyInsuranceCbrfRequest) GoString() string {
	return s.String()
}

func (s *ApplyInsuranceCbrfRequest) SetAuthToken(v string) *ApplyInsuranceCbrfRequest {
	s.AuthToken = &v
	return s
}

func (s *ApplyInsuranceCbrfRequest) SetProductInstanceId(v string) *ApplyInsuranceCbrfRequest {
	s.ProductInstanceId = &v
	return s
}

func (s *ApplyInsuranceCbrfRequest) SetBusinessId(v string) *ApplyInsuranceCbrfRequest {
	s.BusinessId = &v
	return s
}

func (s *ApplyInsuranceCbrfRequest) SetInsuranceCode(v string) *ApplyInsuranceCbrfRequest {
	s.InsuranceCode = &v
	return s
}

func (s *ApplyInsuranceCbrfRequest) SetProductCode(v string) *ApplyInsuranceCbrfRequest {
	s.ProductCode = &v
	return s
}

func (s *ApplyInsuranceCbrfRequest) SetInsuranceStart(v string) *ApplyInsuranceCbrfRequest {
	s.InsuranceStart = &v
	return s
}

func (s *ApplyInsuranceCbrfRequest) SetDeliveryTime(v string) *ApplyInsuranceCbrfRequest {
	s.DeliveryTime = &v
	return s
}

func (s *ApplyInsuranceCbrfRequest) SetRelatedOderNo(v string) *ApplyInsuranceCbrfRequest {
	s.RelatedOderNo = &v
	return s
}

func (s *ApplyInsuranceCbrfRequest) SetDestCountry(v string) *ApplyInsuranceCbrfRequest {
	s.DestCountry = &v
	return s
}

func (s *ApplyInsuranceCbrfRequest) SetBusinessMerchantId(v string) *ApplyInsuranceCbrfRequest {
	s.BusinessMerchantId = &v
	return s
}

func (s *ApplyInsuranceCbrfRequest) SetCollectionTime(v string) *ApplyInsuranceCbrfRequest {
	s.CollectionTime = &v
	return s
}

func (s *ApplyInsuranceCbrfRequest) SetGoodName(v string) *ApplyInsuranceCbrfRequest {
	s.GoodName = &v
	return s
}

func (s *ApplyInsuranceCbrfRequest) SetGoodValue(v string) *ApplyInsuranceCbrfRequest {
	s.GoodValue = &v
	return s
}

type ApplyInsuranceCbrfResponse struct {
	// 请求唯一ID，用于链路跟踪和问题排查
	ReqMsgId *string `json:"req_msg_id,omitempty" xml:"req_msg_id,omitempty"`
	// 结果码，一般OK表示调用成功
	ResultCode *string `json:"result_code,omitempty" xml:"result_code,omitempty"`
	// 异常信息的文本描述
	ResultMsg *string `json:"result_msg,omitempty" xml:"result_msg,omitempty"`
	// 投保响应编码
	InsuredReceiptNo *string `json:"insured_receipt_no,omitempty" xml:"insured_receipt_no,omitempty"`
	// 投保时的标的订单号
	RelatedOrderNo *string `json:"related_order_no,omitempty" xml:"related_order_no,omitempty"`
	// 保司出具的保单编号
	PolicyNo *string `json:"policy_no,omitempty" xml:"policy_no,omitempty"`
	// 返回时间
	ResponseTime *string `json:"response_time,omitempty" xml:"response_time,omitempty" pattern:"\\d{4}[-]\\d{1,2}[-]\\d{1,2}[T]\\d{2}:\\d{2}:\\d{2}([Z]|([\\.]\\d{1,9})?[\\+]\\d{2}[\\:]?\\d{2})"`
	// 幂等标识；true:幂等结果；false: 非幂等结果
	IdemFlag *bool `json:"idem_flag,omitempty" xml:"idem_flag,omitempty"`
}

func (s ApplyInsuranceCbrfResponse) String() string {
	return tea.Prettify(s)
}

func (s ApplyInsuranceCbrfResponse) GoString() string {
	return s.String()
}

func (s *ApplyInsuranceCbrfResponse) SetReqMsgId(v string) *ApplyInsuranceCbrfResponse {
	s.ReqMsgId = &v
	return s
}

func (s *ApplyInsuranceCbrfResponse) SetResultCode(v string) *ApplyInsuranceCbrfResponse {
	s.ResultCode = &v
	return s
}

func (s *ApplyInsuranceCbrfResponse) SetResultMsg(v string) *ApplyInsuranceCbrfResponse {
	s.ResultMsg = &v
	return s
}

func (s *ApplyInsuranceCbrfResponse) SetInsuredReceiptNo(v string) *ApplyInsuranceCbrfResponse {
	s.InsuredReceiptNo = &v
	return s
}

func (s *ApplyInsuranceCbrfResponse) SetRelatedOrderNo(v string) *ApplyInsuranceCbrfResponse {
	s.RelatedOrderNo = &v
	return s
}

func (s *ApplyInsuranceCbrfResponse) SetPolicyNo(v string) *ApplyInsuranceCbrfResponse {
	s.PolicyNo = &v
	return s
}

func (s *ApplyInsuranceCbrfResponse) SetResponseTime(v string) *ApplyInsuranceCbrfResponse {
	s.ResponseTime = &v
	return s
}

func (s *ApplyInsuranceCbrfResponse) SetIdemFlag(v bool) *ApplyInsuranceCbrfResponse {
	s.IdemFlag = &v
	return s
}

type RepayInsuranceCbrfRequest struct {
	// OAuth模式下的授权token
	AuthToken         *string `json:"auth_token,omitempty" xml:"auth_token,omitempty"`
	ProductInstanceId *string `json:"product_instance_id,omitempty" xml:"product_instance_id,omitempty"`
	// 物流公司业务线的简称
	BusinessId *string `json:"business_id,omitempty" xml:"business_id,omitempty" maxLength:"50"`
	// 保司编码
	InsuranceCode *string `json:"insurance_code,omitempty" xml:"insurance_code,omitempty" require:"true" maxLength:"16"`
	//  险种
	ProductCode *string `json:"product_code,omitempty" xml:"product_code,omitempty" require:"true" maxLength:"2"`
	// 客户的订单编号
	RelatedOrderNo *string `json:"related_order_no,omitempty" xml:"related_order_no,omitempty" require:"true" maxLength:"256"`
	// 保单号
	PolicyNo *string `json:"policy_no,omitempty" xml:"policy_no,omitempty" require:"true" maxLength:"50"`
	// 投保响应编码
	InsuredReceiptNo *string `json:"insured_receipt_no,omitempty" xml:"insured_receipt_no,omitempty" require:"true" maxLength:"40"`
	// 2位ISO国家编码
	DestCountry *string `json:"dest_country,omitempty" xml:"dest_country,omitempty" require:"true" maxLength:"256"`
	// 理赔时间
	ClaimTime *string `json:"claim_time,omitempty" xml:"claim_time,omitempty" require:"true" pattern:"\\d{4}[-]\\d{1,2}[-]\\d{1,2}[T]\\d{2}:\\d{2}:\\d{2}([Z]|([\\.]\\d{1,9})?[\\+]\\d{2}[\\:]?\\d{2})"`
	// 理赔金额，单位为元，依据实际情况计算的理赔金额，最多小数点后2位
	ClaimAmount *string `json:"claim_amount,omitempty" xml:"claim_amount,omitempty" require:"true" maxLength:"20"`
	// 达到国内仓库时间
	DwaTime *string `json:"dwa_time,omitempty" xml:"dwa_time,omitempty"`
}

func (s RepayInsuranceCbrfRequest) String() string {
	return tea.Prettify(s)
}

func (s RepayInsuranceCbrfRequest) GoString() string {
	return s.String()
}

func (s *RepayInsuranceCbrfRequest) SetAuthToken(v string) *RepayInsuranceCbrfRequest {
	s.AuthToken = &v
	return s
}

func (s *RepayInsuranceCbrfRequest) SetProductInstanceId(v string) *RepayInsuranceCbrfRequest {
	s.ProductInstanceId = &v
	return s
}

func (s *RepayInsuranceCbrfRequest) SetBusinessId(v string) *RepayInsuranceCbrfRequest {
	s.BusinessId = &v
	return s
}

func (s *RepayInsuranceCbrfRequest) SetInsuranceCode(v string) *RepayInsuranceCbrfRequest {
	s.InsuranceCode = &v
	return s
}

func (s *RepayInsuranceCbrfRequest) SetProductCode(v string) *RepayInsuranceCbrfRequest {
	s.ProductCode = &v
	return s
}

func (s *RepayInsuranceCbrfRequest) SetRelatedOrderNo(v string) *RepayInsuranceCbrfRequest {
	s.RelatedOrderNo = &v
	return s
}

func (s *RepayInsuranceCbrfRequest) SetPolicyNo(v string) *RepayInsuranceCbrfRequest {
	s.PolicyNo = &v
	return s
}

func (s *RepayInsuranceCbrfRequest) SetInsuredReceiptNo(v string) *RepayInsuranceCbrfRequest {
	s.InsuredReceiptNo = &v
	return s
}

func (s *RepayInsuranceCbrfRequest) SetDestCountry(v string) *RepayInsuranceCbrfRequest {
	s.DestCountry = &v
	return s
}

func (s *RepayInsuranceCbrfRequest) SetClaimTime(v string) *RepayInsuranceCbrfRequest {
	s.ClaimTime = &v
	return s
}

func (s *RepayInsuranceCbrfRequest) SetClaimAmount(v string) *RepayInsuranceCbrfRequest {
	s.ClaimAmount = &v
	return s
}

func (s *RepayInsuranceCbrfRequest) SetDwaTime(v string) *RepayInsuranceCbrfRequest {
	s.DwaTime = &v
	return s
}

type RepayInsuranceCbrfResponse struct {
	// 请求唯一ID，用于链路跟踪和问题排查
	ReqMsgId *string `json:"req_msg_id,omitempty" xml:"req_msg_id,omitempty"`
	// 结果码，一般OK表示调用成功
	ResultCode *string `json:"result_code,omitempty" xml:"result_code,omitempty"`
	// 异常信息的文本描述
	ResultMsg *string `json:"result_msg,omitempty" xml:"result_msg,omitempty"`
	// 赔案号
	ClaimNo *string `json:"claim_no,omitempty" xml:"claim_no,omitempty"`
	// 客户投保时的标的订单号
	RelatedOrderNo *string `json:"related_order_no,omitempty" xml:"related_order_no,omitempty"`
	// 保单号
	PolicyNo *string `json:"policy_no,omitempty" xml:"policy_no,omitempty"`
	// 幂等标识。true：幂等结果;false：非幂等结果
	IdemFlag *bool `json:"idem_flag,omitempty" xml:"idem_flag,omitempty"`
}

func (s RepayInsuranceCbrfResponse) String() string {
	return tea.Prettify(s)
}

func (s RepayInsuranceCbrfResponse) GoString() string {
	return s.String()
}

func (s *RepayInsuranceCbrfResponse) SetReqMsgId(v string) *RepayInsuranceCbrfResponse {
	s.ReqMsgId = &v
	return s
}

func (s *RepayInsuranceCbrfResponse) SetResultCode(v string) *RepayInsuranceCbrfResponse {
	s.ResultCode = &v
	return s
}

func (s *RepayInsuranceCbrfResponse) SetResultMsg(v string) *RepayInsuranceCbrfResponse {
	s.ResultMsg = &v
	return s
}

func (s *RepayInsuranceCbrfResponse) SetClaimNo(v string) *RepayInsuranceCbrfResponse {
	s.ClaimNo = &v
	return s
}

func (s *RepayInsuranceCbrfResponse) SetRelatedOrderNo(v string) *RepayInsuranceCbrfResponse {
	s.RelatedOrderNo = &v
	return s
}

func (s *RepayInsuranceCbrfResponse) SetPolicyNo(v string) *RepayInsuranceCbrfResponse {
	s.PolicyNo = &v
	return s
}

func (s *RepayInsuranceCbrfResponse) SetIdemFlag(v bool) *RepayInsuranceCbrfResponse {
	s.IdemFlag = &v
	return s
}

type ApplyInsuranceCbecRequest struct {
	// OAuth模式下的授权token
	AuthToken         *string `json:"auth_token,omitempty" xml:"auth_token,omitempty"`
	ProductInstanceId *string `json:"product_instance_id,omitempty" xml:"product_instance_id,omitempty"`
	// 投保交易号，调用方生成的唯一编码
	TradeNo *string `json:"trade_no,omitempty" xml:"trade_no,omitempty" require:"true" maxLength:"50"`
	// 保司编码，CPIC---太保
	ExternalChannelCode *string `json:"external_channel_code,omitempty" xml:"external_channel_code,omitempty" require:"true" maxLength:"64"`
	// 险种编码，03--跨境出口货运险
	ExternalProductCode *string `json:"external_product_code,omitempty" xml:"external_product_code,omitempty" require:"true" maxLength:"64"`
	// 投保人姓名
	TbrName *string `json:"tbr_name,omitempty" xml:"tbr_name,omitempty" require:"true" maxLength:"100"`
	// 投保人证件类型，03--营业执照
	TbrIdType *string `json:"tbr_id_type,omitempty" xml:"tbr_id_type,omitempty" maxLength:"40"`
	// 投保人证件号码
	TbrIdNo *string `json:"tbr_id_no,omitempty" xml:"tbr_id_no,omitempty" maxLength:"800"`
	// 被保人姓名
	BbrName *string `json:"bbr_name,omitempty" xml:"bbr_name,omitempty" require:"true" maxLength:"100"`
	// 被保人证件类型，01--居民身份证，03--营业执照
	BbrIdType *string `json:"bbr_id_type,omitempty" xml:"bbr_id_type,omitempty" maxLength:"40"`
	// 被保人证件号码
	BbrIdNo *string `json:"bbr_id_no,omitempty" xml:"bbr_id_no,omitempty" maxLength:"800"`
	// 出库单号/航次号/运单号
	RelatedOrderNo *string `json:"related_order_no,omitempty" xml:"related_order_no,omitempty" require:"true" maxLength:"200"`
	// 包装及数量
	CargoQuantity *string `json:"cargo_quantity,omitempty" xml:"cargo_quantity,omitempty" require:"true" maxLength:"20"`
	// 货物名称
	CargoName *string `json:"cargo_name,omitempty" xml:"cargo_name,omitempty" require:"true" maxLength:"100"`
	// 包装代码,参考保司提供样例-包装代码
	PackingCode *string `json:"packing_code,omitempty" xml:"packing_code,omitempty" require:"true" maxLength:"2"`
	// 货物类型代码,货物类型代码,参考保司提供样例-货物类型代码
	CargoTypeCode *string `json:"cargo_type_code,omitempty" xml:"cargo_type_code,omitempty" require:"true" maxLength:"4"`
	// 航行区域代码,参考保司提供样例-航行区域代码
	FlightAreaCode *string `json:"flight_area_code,omitempty" xml:"flight_area_code,omitempty" require:"true" maxLength:"10"`
	// 运输方式代码,参考保司提供样例-运输方式
	TransportTypeCode *string `json:"transport_type_code,omitempty" xml:"transport_type_code,omitempty" require:"true" maxLength:"2"`
	// 运输工具名称
	TransportMeansName *string `json:"transport_means_name,omitempty" xml:"transport_means_name,omitempty" require:"true" maxLength:"30"`
	// 航次号
	Voyage *string `json:"voyage,omitempty" xml:"voyage,omitempty" maxLength:"30"`
	// 出发地
	StartPlace *string `json:"start_place,omitempty" xml:"start_place,omitempty" require:"true" maxLength:"200"`
	// 中转地
	TransitPoint *string `json:"transit_point,omitempty" xml:"transit_point,omitempty" maxLength:"100"`
	// 目的地
	Destination *string `json:"destination,omitempty" xml:"destination,omitempty" require:"true" maxLength:"200"`
	// 理赔代理地代码，参考保司提供样例-理赔代理地
	ClaimAgentCode *string `json:"claim_agent_code,omitempty" xml:"claim_agent_code,omitempty" require:"true" maxLength:"30"`
	// 主险条款代码
	MainItemCode *string `json:"main_item_code,omitempty" xml:"main_item_code,omitempty" require:"true" maxLength:"12"`
	// 主险条款内容
	MainItemContent *string `json:"main_item_content,omitempty" xml:"main_item_content,omitempty" require:"true" maxLength:"500"`
	// 附加条款集合
	MainItemAdds []*MainItemAdd `json:"main_item_adds,omitempty" xml:"main_item_adds,omitempty" type:"Repeated"`
	// 特别约定
	Specialize *string `json:"specialize,omitempty" xml:"specialize,omitempty" maxLength:"500"`
	// 申报货物价值,，最多兼容2位小数，单位（元）
	CargoWorth *string `json:"cargo_worth,omitempty" xml:"cargo_worth,omitempty" require:"true" maxLength:"50"`
	// 费率，最多兼容9位小数
	Rate *string `json:"rate,omitempty" xml:"rate,omitempty" require:"true"`
	// 保费，最多兼容2位小数，单位（元）
	Premium *string `json:"premium,omitempty" xml:"premium,omitempty" require:"true"`
	// 保险起期，日期格式yyyy-mm-dd
	InsureStart *string `json:"insure_start,omitempty" xml:"insure_start,omitempty" require:"true"`
	// 起运日期，日期格式yyyy-mm-dd
	SaleDate *string `json:"sale_date,omitempty" xml:"sale_date,omitempty" require:"true"`
}

func (s ApplyInsuranceCbecRequest) String() string {
	return tea.Prettify(s)
}

func (s ApplyInsuranceCbecRequest) GoString() string {
	return s.String()
}

func (s *ApplyInsuranceCbecRequest) SetAuthToken(v string) *ApplyInsuranceCbecRequest {
	s.AuthToken = &v
	return s
}

func (s *ApplyInsuranceCbecRequest) SetProductInstanceId(v string) *ApplyInsuranceCbecRequest {
	s.ProductInstanceId = &v
	return s
}

func (s *ApplyInsuranceCbecRequest) SetTradeNo(v string) *ApplyInsuranceCbecRequest {
	s.TradeNo = &v
	return s
}

func (s *ApplyInsuranceCbecRequest) SetExternalChannelCode(v string) *ApplyInsuranceCbecRequest {
	s.ExternalChannelCode = &v
	return s
}

func (s *ApplyInsuranceCbecRequest) SetExternalProductCode(v string) *ApplyInsuranceCbecRequest {
	s.ExternalProductCode = &v
	return s
}

func (s *ApplyInsuranceCbecRequest) SetTbrName(v string) *ApplyInsuranceCbecRequest {
	s.TbrName = &v
	return s
}

func (s *ApplyInsuranceCbecRequest) SetTbrIdType(v string) *ApplyInsuranceCbecRequest {
	s.TbrIdType = &v
	return s
}

func (s *ApplyInsuranceCbecRequest) SetTbrIdNo(v string) *ApplyInsuranceCbecRequest {
	s.TbrIdNo = &v
	return s
}

func (s *ApplyInsuranceCbecRequest) SetBbrName(v string) *ApplyInsuranceCbecRequest {
	s.BbrName = &v
	return s
}

func (s *ApplyInsuranceCbecRequest) SetBbrIdType(v string) *ApplyInsuranceCbecRequest {
	s.BbrIdType = &v
	return s
}

func (s *ApplyInsuranceCbecRequest) SetBbrIdNo(v string) *ApplyInsuranceCbecRequest {
	s.BbrIdNo = &v
	return s
}

func (s *ApplyInsuranceCbecRequest) SetRelatedOrderNo(v string) *ApplyInsuranceCbecRequest {
	s.RelatedOrderNo = &v
	return s
}

func (s *ApplyInsuranceCbecRequest) SetCargoQuantity(v string) *ApplyInsuranceCbecRequest {
	s.CargoQuantity = &v
	return s
}

func (s *ApplyInsuranceCbecRequest) SetCargoName(v string) *ApplyInsuranceCbecRequest {
	s.CargoName = &v
	return s
}

func (s *ApplyInsuranceCbecRequest) SetPackingCode(v string) *ApplyInsuranceCbecRequest {
	s.PackingCode = &v
	return s
}

func (s *ApplyInsuranceCbecRequest) SetCargoTypeCode(v string) *ApplyInsuranceCbecRequest {
	s.CargoTypeCode = &v
	return s
}

func (s *ApplyInsuranceCbecRequest) SetFlightAreaCode(v string) *ApplyInsuranceCbecRequest {
	s.FlightAreaCode = &v
	return s
}

func (s *ApplyInsuranceCbecRequest) SetTransportTypeCode(v string) *ApplyInsuranceCbecRequest {
	s.TransportTypeCode = &v
	return s
}

func (s *ApplyInsuranceCbecRequest) SetTransportMeansName(v string) *ApplyInsuranceCbecRequest {
	s.TransportMeansName = &v
	return s
}

func (s *ApplyInsuranceCbecRequest) SetVoyage(v string) *ApplyInsuranceCbecRequest {
	s.Voyage = &v
	return s
}

func (s *ApplyInsuranceCbecRequest) SetStartPlace(v string) *ApplyInsuranceCbecRequest {
	s.StartPlace = &v
	return s
}

func (s *ApplyInsuranceCbecRequest) SetTransitPoint(v string) *ApplyInsuranceCbecRequest {
	s.TransitPoint = &v
	return s
}

func (s *ApplyInsuranceCbecRequest) SetDestination(v string) *ApplyInsuranceCbecRequest {
	s.Destination = &v
	return s
}

func (s *ApplyInsuranceCbecRequest) SetClaimAgentCode(v string) *ApplyInsuranceCbecRequest {
	s.ClaimAgentCode = &v
	return s
}

func (s *ApplyInsuranceCbecRequest) SetMainItemCode(v string) *ApplyInsuranceCbecRequest {
	s.MainItemCode = &v
	return s
}

func (s *ApplyInsuranceCbecRequest) SetMainItemContent(v string) *ApplyInsuranceCbecRequest {
	s.MainItemContent = &v
	return s
}

func (s *ApplyInsuranceCbecRequest) SetMainItemAdds(v []*MainItemAdd) *ApplyInsuranceCbecRequest {
	s.MainItemAdds = v
	return s
}

func (s *ApplyInsuranceCbecRequest) SetSpecialize(v string) *ApplyInsuranceCbecRequest {
	s.Specialize = &v
	return s
}

func (s *ApplyInsuranceCbecRequest) SetCargoWorth(v string) *ApplyInsuranceCbecRequest {
	s.CargoWorth = &v
	return s
}

func (s *ApplyInsuranceCbecRequest) SetRate(v string) *ApplyInsuranceCbecRequest {
	s.Rate = &v
	return s
}

func (s *ApplyInsuranceCbecRequest) SetPremium(v string) *ApplyInsuranceCbecRequest {
	s.Premium = &v
	return s
}

func (s *ApplyInsuranceCbecRequest) SetInsureStart(v string) *ApplyInsuranceCbecRequest {
	s.InsureStart = &v
	return s
}

func (s *ApplyInsuranceCbecRequest) SetSaleDate(v string) *ApplyInsuranceCbecRequest {
	s.SaleDate = &v
	return s
}

type ApplyInsuranceCbecResponse struct {
	// 请求唯一ID，用于链路跟踪和问题排查
	ReqMsgId *string `json:"req_msg_id,omitempty" xml:"req_msg_id,omitempty"`
	// 结果码，一般OK表示调用成功
	ResultCode *string `json:"result_code,omitempty" xml:"result_code,omitempty"`
	// 异常信息的文本描述
	ResultMsg *string `json:"result_msg,omitempty" xml:"result_msg,omitempty"`
	// 投保交易号
	TradeNo *string `json:"trade_no,omitempty" xml:"trade_no,omitempty"`
	// 保单号，保司生成的投保业务号
	PolicyNo *string `json:"policy_no,omitempty" xml:"policy_no,omitempty"`
	// 保费
	Premium *string `json:"premium,omitempty" xml:"premium,omitempty"`
}

func (s ApplyInsuranceCbecResponse) String() string {
	return tea.Prettify(s)
}

func (s ApplyInsuranceCbecResponse) GoString() string {
	return s.String()
}

func (s *ApplyInsuranceCbecResponse) SetReqMsgId(v string) *ApplyInsuranceCbecResponse {
	s.ReqMsgId = &v
	return s
}

func (s *ApplyInsuranceCbecResponse) SetResultCode(v string) *ApplyInsuranceCbecResponse {
	s.ResultCode = &v
	return s
}

func (s *ApplyInsuranceCbecResponse) SetResultMsg(v string) *ApplyInsuranceCbecResponse {
	s.ResultMsg = &v
	return s
}

func (s *ApplyInsuranceCbecResponse) SetTradeNo(v string) *ApplyInsuranceCbecResponse {
	s.TradeNo = &v
	return s
}

func (s *ApplyInsuranceCbecResponse) SetPolicyNo(v string) *ApplyInsuranceCbecResponse {
	s.PolicyNo = &v
	return s
}

func (s *ApplyInsuranceCbecResponse) SetPremium(v string) *ApplyInsuranceCbecResponse {
	s.Premium = &v
	return s
}

type ApplyInsuranceStockinRequest struct {
	// OAuth模式下的授权token
	AuthToken         *string `json:"auth_token,omitempty" xml:"auth_token,omitempty"`
	ProductInstanceId *string `json:"product_instance_id,omitempty" xml:"product_instance_id,omitempty"`
	// 入库交易流水号，保持唯一
	TradeNo *string `json:"trade_no,omitempty" xml:"trade_no,omitempty" require:"true" maxLength:"50"`
	// 入库单号，可参考格式：年月日+唯一字符
	StockinNo *string `json:"stockin_no,omitempty" xml:"stockin_no,omitempty" require:"true" maxLength:"50"`
	// 入库时间，格式：yyyy-MM-dd HH:mm:ss
	StockinDate *string `json:"stockin_date,omitempty" xml:"stockin_date,omitempty" require:"true"`
	// 时区
	Timezone *string `json:"timezone,omitempty" xml:"timezone,omitempty" require:"true" maxLength:"16"`
	// 仓库代码
	RepositoryCode *string `json:"repository_code,omitempty" xml:"repository_code,omitempty" require:"true" maxLength:"50"`
	// 客户代码
	CustomerCode *string `json:"customer_code,omitempty" xml:"customer_code,omitempty" require:"true" maxLength:"50"`
	// 保单号,国内起运时投保产生的保单号
	PolicyNo *string `json:"policy_no,omitempty" xml:"policy_no,omitempty" maxLength:"64"`
	// 入库货物列表
	StockinCargos []*StockinCargo `json:"stockin_cargos,omitempty" xml:"stockin_cargos,omitempty" require:"true" type:"Repeated"`
}

func (s ApplyInsuranceStockinRequest) String() string {
	return tea.Prettify(s)
}

func (s ApplyInsuranceStockinRequest) GoString() string {
	return s.String()
}

func (s *ApplyInsuranceStockinRequest) SetAuthToken(v string) *ApplyInsuranceStockinRequest {
	s.AuthToken = &v
	return s
}

func (s *ApplyInsuranceStockinRequest) SetProductInstanceId(v string) *ApplyInsuranceStockinRequest {
	s.ProductInstanceId = &v
	return s
}

func (s *ApplyInsuranceStockinRequest) SetTradeNo(v string) *ApplyInsuranceStockinRequest {
	s.TradeNo = &v
	return s
}

func (s *ApplyInsuranceStockinRequest) SetStockinNo(v string) *ApplyInsuranceStockinRequest {
	s.StockinNo = &v
	return s
}

func (s *ApplyInsuranceStockinRequest) SetStockinDate(v string) *ApplyInsuranceStockinRequest {
	s.StockinDate = &v
	return s
}

func (s *ApplyInsuranceStockinRequest) SetTimezone(v string) *ApplyInsuranceStockinRequest {
	s.Timezone = &v
	return s
}

func (s *ApplyInsuranceStockinRequest) SetRepositoryCode(v string) *ApplyInsuranceStockinRequest {
	s.RepositoryCode = &v
	return s
}

func (s *ApplyInsuranceStockinRequest) SetCustomerCode(v string) *ApplyInsuranceStockinRequest {
	s.CustomerCode = &v
	return s
}

func (s *ApplyInsuranceStockinRequest) SetPolicyNo(v string) *ApplyInsuranceStockinRequest {
	s.PolicyNo = &v
	return s
}

func (s *ApplyInsuranceStockinRequest) SetStockinCargos(v []*StockinCargo) *ApplyInsuranceStockinRequest {
	s.StockinCargos = v
	return s
}

type ApplyInsuranceStockinResponse struct {
	// 请求唯一ID，用于链路跟踪和问题排查
	ReqMsgId *string `json:"req_msg_id,omitempty" xml:"req_msg_id,omitempty"`
	// 结果码，一般OK表示调用成功
	ResultCode *string `json:"result_code,omitempty" xml:"result_code,omitempty"`
	// 异常信息的文本描述
	ResultMsg *string `json:"result_msg,omitempty" xml:"result_msg,omitempty"`
	// 入库交易流水号
	TradeNo *string `json:"trade_no,omitempty" xml:"trade_no,omitempty"`
	// 入库单号
	//
	//
	StockinNo *string `json:"stockin_no,omitempty" xml:"stockin_no,omitempty"`
}

func (s ApplyInsuranceStockinResponse) String() string {
	return tea.Prettify(s)
}

func (s ApplyInsuranceStockinResponse) GoString() string {
	return s.String()
}

func (s *ApplyInsuranceStockinResponse) SetReqMsgId(v string) *ApplyInsuranceStockinResponse {
	s.ReqMsgId = &v
	return s
}

func (s *ApplyInsuranceStockinResponse) SetResultCode(v string) *ApplyInsuranceStockinResponse {
	s.ResultCode = &v
	return s
}

func (s *ApplyInsuranceStockinResponse) SetResultMsg(v string) *ApplyInsuranceStockinResponse {
	s.ResultMsg = &v
	return s
}

func (s *ApplyInsuranceStockinResponse) SetTradeNo(v string) *ApplyInsuranceStockinResponse {
	s.TradeNo = &v
	return s
}

func (s *ApplyInsuranceStockinResponse) SetStockinNo(v string) *ApplyInsuranceStockinResponse {
	s.StockinNo = &v
	return s
}

type ApplyInsuranceInventoryRequest struct {
	// OAuth模式下的授权token
	AuthToken         *string `json:"auth_token,omitempty" xml:"auth_token,omitempty"`
	ProductInstanceId *string `json:"product_instance_id,omitempty" xml:"product_instance_id,omitempty"`
	// 库存申报交易号，调用方生成的唯一编码
	TradeNo *string `json:"trade_no,omitempty" xml:"trade_no,omitempty" require:"true" maxLength:"50"`
	// 库存申报编号
	InventoryNo *string `json:"inventory_no,omitempty" xml:"inventory_no,omitempty" require:"true" maxLength:"50"`
	// 库存查询时间，yyyy-mm-dd，精确到（天）
	InventoryQueryDate *string `json:"inventory_query_date,omitempty" xml:"inventory_query_date,omitempty" require:"true"`
	// 仓库代码
	//
	//
	RepositoryCode *string `json:"repository_code,omitempty" xml:"repository_code,omitempty" require:"true" maxLength:"50"`
	// 库存货物列表
	InventoryCargos []*InventoryCargo `json:"inventory_cargos,omitempty" xml:"inventory_cargos,omitempty" require:"true" type:"Repeated"`
}

func (s ApplyInsuranceInventoryRequest) String() string {
	return tea.Prettify(s)
}

func (s ApplyInsuranceInventoryRequest) GoString() string {
	return s.String()
}

func (s *ApplyInsuranceInventoryRequest) SetAuthToken(v string) *ApplyInsuranceInventoryRequest {
	s.AuthToken = &v
	return s
}

func (s *ApplyInsuranceInventoryRequest) SetProductInstanceId(v string) *ApplyInsuranceInventoryRequest {
	s.ProductInstanceId = &v
	return s
}

func (s *ApplyInsuranceInventoryRequest) SetTradeNo(v string) *ApplyInsuranceInventoryRequest {
	s.TradeNo = &v
	return s
}

func (s *ApplyInsuranceInventoryRequest) SetInventoryNo(v string) *ApplyInsuranceInventoryRequest {
	s.InventoryNo = &v
	return s
}

func (s *ApplyInsuranceInventoryRequest) SetInventoryQueryDate(v string) *ApplyInsuranceInventoryRequest {
	s.InventoryQueryDate = &v
	return s
}

func (s *ApplyInsuranceInventoryRequest) SetRepositoryCode(v string) *ApplyInsuranceInventoryRequest {
	s.RepositoryCode = &v
	return s
}

func (s *ApplyInsuranceInventoryRequest) SetInventoryCargos(v []*InventoryCargo) *ApplyInsuranceInventoryRequest {
	s.InventoryCargos = v
	return s
}

type ApplyInsuranceInventoryResponse struct {
	// 请求唯一ID，用于链路跟踪和问题排查
	ReqMsgId *string `json:"req_msg_id,omitempty" xml:"req_msg_id,omitempty"`
	// 结果码，一般OK表示调用成功
	ResultCode *string `json:"result_code,omitempty" xml:"result_code,omitempty"`
	// 异常信息的文本描述
	ResultMsg *string `json:"result_msg,omitempty" xml:"result_msg,omitempty"`
	// 库存申报编号
	InventoryNo *string `json:"inventory_no,omitempty" xml:"inventory_no,omitempty"`
	// 库存申报交易流水号
	TradeNo *string `json:"trade_no,omitempty" xml:"trade_no,omitempty"`
}

func (s ApplyInsuranceInventoryResponse) String() string {
	return tea.Prettify(s)
}

func (s ApplyInsuranceInventoryResponse) GoString() string {
	return s.String()
}

func (s *ApplyInsuranceInventoryResponse) SetReqMsgId(v string) *ApplyInsuranceInventoryResponse {
	s.ReqMsgId = &v
	return s
}

func (s *ApplyInsuranceInventoryResponse) SetResultCode(v string) *ApplyInsuranceInventoryResponse {
	s.ResultCode = &v
	return s
}

func (s *ApplyInsuranceInventoryResponse) SetResultMsg(v string) *ApplyInsuranceInventoryResponse {
	s.ResultMsg = &v
	return s
}

func (s *ApplyInsuranceInventoryResponse) SetInventoryNo(v string) *ApplyInsuranceInventoryResponse {
	s.InventoryNo = &v
	return s
}

func (s *ApplyInsuranceInventoryResponse) SetTradeNo(v string) *ApplyInsuranceInventoryResponse {
	s.TradeNo = &v
	return s
}

type PushAuthSigninfoRequest struct {
	// OAuth模式下的授权token
	AuthToken         *string `json:"auth_token,omitempty" xml:"auth_token,omitempty"`
	ProductInstanceId *string `json:"product_instance_id,omitempty" xml:"product_instance_id,omitempty"`
	// 签署流程id
	SignProcessId *string `json:"sign_process_id,omitempty" xml:"sign_process_id,omitempty" require:"true"`
	// 授权关系类型id
	AuthRelTypeId *string `json:"auth_rel_type_id,omitempty" xml:"auth_rel_type_id,omitempty" require:"true"`
	// 某某主题
	SignTheme *string `json:"sign_theme,omitempty" xml:"sign_theme,omitempty"`
	// 发起方名称
	InitiatorName *string `json:"initiator_name,omitempty" xml:"initiator_name,omitempty" require:"true"`
	// 发起方证件类型，可以填写的枚举类型：UNIFIED_SOCIAL_CREDIT_CODE,  BUSINESS_LICENSE_NUMBER。分别表示统一社会信用代码和工商注册号
	InitiatorCertType *string `json:"initiator_cert_type,omitempty" xml:"initiator_cert_type,omitempty" require:"true"`
	// 发起方证件号码
	InitiatorCertNum *string `json:"initiator_cert_num,omitempty" xml:"initiator_cert_num,omitempty" require:"true"`
	// 合同模板hash
	ContTmplHash *string `json:"cont_tmpl_hash,omitempty" xml:"cont_tmpl_hash,omitempty" require:"true"`
	// 签署方信息列表
	AuthPartys []*AuthParty `json:"auth_partys,omitempty" xml:"auth_partys,omitempty" require:"true" type:"Repeated"`
	// 上链文件信息列表
	AuthChainFiles []*AuthChainFile `json:"auth_chain_files,omitempty" xml:"auth_chain_files,omitempty" require:"true" type:"Repeated"`
	// 动态字段1
	DynaField1 *string `json:"dyna_field1,omitempty" xml:"dyna_field1,omitempty"`
	// 动态字段2
	DynaField2 *string `json:"dyna_field2,omitempty" xml:"dyna_field2,omitempty"`
	// 动态字段3
	DynaField3 *string `json:"dyna_field3,omitempty" xml:"dyna_field3,omitempty"`
	// 动态字段4
	DynaField4 *string `json:"dyna_field4,omitempty" xml:"dyna_field4,omitempty"`
	// 动态字段5
	DynaField5 *string `json:"dyna_field5,omitempty" xml:"dyna_field5,omitempty"`
	// 动态字段6
	DynaField6 *string `json:"dyna_field6,omitempty" xml:"dyna_field6,omitempty"`
	// 动态字段7
	DynaField7 *string `json:"dyna_field7,omitempty" xml:"dyna_field7,omitempty"`
	// 动态字段8
	DynaField8 *string `json:"dyna_field8,omitempty" xml:"dyna_field8,omitempty"`
	// 动态字段9
	DynaField9 *string `json:"dyna_field9,omitempty" xml:"dyna_field9,omitempty"`
	// 动态字段10
	DynaField10 *string `json:"dyna_field10,omitempty" xml:"dyna_field10,omitempty"`
	// 动态字段11
	DynaField11 *string `json:"dyna_field11,omitempty" xml:"dyna_field11,omitempty"`
	// 动态字段12
	DynaField12 *string `json:"dyna_field12,omitempty" xml:"dyna_field12,omitempty"`
	// 动态字段13
	DynaField13 *string `json:"dyna_field13,omitempty" xml:"dyna_field13,omitempty"`
	// 动态字段14
	DynaField14 *string `json:"dyna_field14,omitempty" xml:"dyna_field14,omitempty"`
	// 动态字段15
	DynaField15 *string `json:"dyna_field15,omitempty" xml:"dyna_field15,omitempty"`
	// 动态字段16
	DynaField16 *string `json:"dyna_field16,omitempty" xml:"dyna_field16,omitempty"`
	// 动态字段17
	DynaField17 *string `json:"dyna_field17,omitempty" xml:"dyna_field17,omitempty"`
	// 动态字段18
	DynaField18 *string `json:"dyna_field18,omitempty" xml:"dyna_field18,omitempty"`
	// 动态字段19
	DynaField19 *string `json:"dyna_field19,omitempty" xml:"dyna_field19,omitempty"`
	// 动态字段20
	DynaField20 *string `json:"dyna_field20,omitempty" xml:"dyna_field20,omitempty"`
}

func (s PushAuthSigninfoRequest) String() string {
	return tea.Prettify(s)
}

func (s PushAuthSigninfoRequest) GoString() string {
	return s.String()
}

func (s *PushAuthSigninfoRequest) SetAuthToken(v string) *PushAuthSigninfoRequest {
	s.AuthToken = &v
	return s
}

func (s *PushAuthSigninfoRequest) SetProductInstanceId(v string) *PushAuthSigninfoRequest {
	s.ProductInstanceId = &v
	return s
}

func (s *PushAuthSigninfoRequest) SetSignProcessId(v string) *PushAuthSigninfoRequest {
	s.SignProcessId = &v
	return s
}

func (s *PushAuthSigninfoRequest) SetAuthRelTypeId(v string) *PushAuthSigninfoRequest {
	s.AuthRelTypeId = &v
	return s
}

func (s *PushAuthSigninfoRequest) SetSignTheme(v string) *PushAuthSigninfoRequest {
	s.SignTheme = &v
	return s
}

func (s *PushAuthSigninfoRequest) SetInitiatorName(v string) *PushAuthSigninfoRequest {
	s.InitiatorName = &v
	return s
}

func (s *PushAuthSigninfoRequest) SetInitiatorCertType(v string) *PushAuthSigninfoRequest {
	s.InitiatorCertType = &v
	return s
}

func (s *PushAuthSigninfoRequest) SetInitiatorCertNum(v string) *PushAuthSigninfoRequest {
	s.InitiatorCertNum = &v
	return s
}

func (s *PushAuthSigninfoRequest) SetContTmplHash(v string) *PushAuthSigninfoRequest {
	s.ContTmplHash = &v
	return s
}

func (s *PushAuthSigninfoRequest) SetAuthPartys(v []*AuthParty) *PushAuthSigninfoRequest {
	s.AuthPartys = v
	return s
}

func (s *PushAuthSigninfoRequest) SetAuthChainFiles(v []*AuthChainFile) *PushAuthSigninfoRequest {
	s.AuthChainFiles = v
	return s
}

func (s *PushAuthSigninfoRequest) SetDynaField1(v string) *PushAuthSigninfoRequest {
	s.DynaField1 = &v
	return s
}

func (s *PushAuthSigninfoRequest) SetDynaField2(v string) *PushAuthSigninfoRequest {
	s.DynaField2 = &v
	return s
}

func (s *PushAuthSigninfoRequest) SetDynaField3(v string) *PushAuthSigninfoRequest {
	s.DynaField3 = &v
	return s
}

func (s *PushAuthSigninfoRequest) SetDynaField4(v string) *PushAuthSigninfoRequest {
	s.DynaField4 = &v
	return s
}

func (s *PushAuthSigninfoRequest) SetDynaField5(v string) *PushAuthSigninfoRequest {
	s.DynaField5 = &v
	return s
}

func (s *PushAuthSigninfoRequest) SetDynaField6(v string) *PushAuthSigninfoRequest {
	s.DynaField6 = &v
	return s
}

func (s *PushAuthSigninfoRequest) SetDynaField7(v string) *PushAuthSigninfoRequest {
	s.DynaField7 = &v
	return s
}

func (s *PushAuthSigninfoRequest) SetDynaField8(v string) *PushAuthSigninfoRequest {
	s.DynaField8 = &v
	return s
}

func (s *PushAuthSigninfoRequest) SetDynaField9(v string) *PushAuthSigninfoRequest {
	s.DynaField9 = &v
	return s
}

func (s *PushAuthSigninfoRequest) SetDynaField10(v string) *PushAuthSigninfoRequest {
	s.DynaField10 = &v
	return s
}

func (s *PushAuthSigninfoRequest) SetDynaField11(v string) *PushAuthSigninfoRequest {
	s.DynaField11 = &v
	return s
}

func (s *PushAuthSigninfoRequest) SetDynaField12(v string) *PushAuthSigninfoRequest {
	s.DynaField12 = &v
	return s
}

func (s *PushAuthSigninfoRequest) SetDynaField13(v string) *PushAuthSigninfoRequest {
	s.DynaField13 = &v
	return s
}

func (s *PushAuthSigninfoRequest) SetDynaField14(v string) *PushAuthSigninfoRequest {
	s.DynaField14 = &v
	return s
}

func (s *PushAuthSigninfoRequest) SetDynaField15(v string) *PushAuthSigninfoRequest {
	s.DynaField15 = &v
	return s
}

func (s *PushAuthSigninfoRequest) SetDynaField16(v string) *PushAuthSigninfoRequest {
	s.DynaField16 = &v
	return s
}

func (s *PushAuthSigninfoRequest) SetDynaField17(v string) *PushAuthSigninfoRequest {
	s.DynaField17 = &v
	return s
}

func (s *PushAuthSigninfoRequest) SetDynaField18(v string) *PushAuthSigninfoRequest {
	s.DynaField18 = &v
	return s
}

func (s *PushAuthSigninfoRequest) SetDynaField19(v string) *PushAuthSigninfoRequest {
	s.DynaField19 = &v
	return s
}

func (s *PushAuthSigninfoRequest) SetDynaField20(v string) *PushAuthSigninfoRequest {
	s.DynaField20 = &v
	return s
}

type PushAuthSigninfoResponse struct {
	// 请求唯一ID，用于链路跟踪和问题排查
	ReqMsgId *string `json:"req_msg_id,omitempty" xml:"req_msg_id,omitempty"`
	// 结果码，一般OK表示调用成功
	ResultCode *string `json:"result_code,omitempty" xml:"result_code,omitempty"`
	// 异常信息的文本描述
	ResultMsg *string `json:"result_msg,omitempty" xml:"result_msg,omitempty"`
	// 推送成功
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
}

func (s PushAuthSigninfoResponse) String() string {
	return tea.Prettify(s)
}

func (s PushAuthSigninfoResponse) GoString() string {
	return s.String()
}

func (s *PushAuthSigninfoResponse) SetReqMsgId(v string) *PushAuthSigninfoResponse {
	s.ReqMsgId = &v
	return s
}

func (s *PushAuthSigninfoResponse) SetResultCode(v string) *PushAuthSigninfoResponse {
	s.ResultCode = &v
	return s
}

func (s *PushAuthSigninfoResponse) SetResultMsg(v string) *PushAuthSigninfoResponse {
	s.ResultMsg = &v
	return s
}

func (s *PushAuthSigninfoResponse) SetStatus(v string) *PushAuthSigninfoResponse {
	s.Status = &v
	return s
}

type QueryPfPaymentRequest struct {
	// OAuth模式下的授权token
	AuthToken         *string `json:"auth_token,omitempty" xml:"auth_token,omitempty"`
	ProductInstanceId *string `json:"product_instance_id,omitempty" xml:"product_instance_id,omitempty"`
	// 业务参考号
	BussRefrNo *string `json:"buss_refr_no,omitempty" xml:"buss_refr_no,omitempty" require:"true" maxLength:"100"`
	// 支用id
	FinancingId *string `json:"financing_id,omitempty" xml:"financing_id,omitempty" require:"true" maxLength:"32"`
}

func (s QueryPfPaymentRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryPfPaymentRequest) GoString() string {
	return s.String()
}

func (s *QueryPfPaymentRequest) SetAuthToken(v string) *QueryPfPaymentRequest {
	s.AuthToken = &v
	return s
}

func (s *QueryPfPaymentRequest) SetProductInstanceId(v string) *QueryPfPaymentRequest {
	s.ProductInstanceId = &v
	return s
}

func (s *QueryPfPaymentRequest) SetBussRefrNo(v string) *QueryPfPaymentRequest {
	s.BussRefrNo = &v
	return s
}

func (s *QueryPfPaymentRequest) SetFinancingId(v string) *QueryPfPaymentRequest {
	s.FinancingId = &v
	return s
}

type QueryPfPaymentResponse struct {
	// 请求唯一ID，用于链路跟踪和问题排查
	ReqMsgId *string `json:"req_msg_id,omitempty" xml:"req_msg_id,omitempty"`
	// 结果码，一般OK表示调用成功
	ResultCode *string `json:"result_code,omitempty" xml:"result_code,omitempty"`
	// 异常信息的文本描述
	ResultMsg *string `json:"result_msg,omitempty" xml:"result_msg,omitempty"`
	// 入账状态: 10-受理 20-在途 00-成功 99-失败
	AccountinStatus *string `json:"accountin_status,omitempty" xml:"accountin_status,omitempty"`
	// 交易时间
	TradeTime *string `json:"trade_time,omitempty" xml:"trade_time,omitempty"`
}

func (s QueryPfPaymentResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryPfPaymentResponse) GoString() string {
	return s.String()
}

func (s *QueryPfPaymentResponse) SetReqMsgId(v string) *QueryPfPaymentResponse {
	s.ReqMsgId = &v
	return s
}

func (s *QueryPfPaymentResponse) SetResultCode(v string) *QueryPfPaymentResponse {
	s.ResultCode = &v
	return s
}

func (s *QueryPfPaymentResponse) SetResultMsg(v string) *QueryPfPaymentResponse {
	s.ResultMsg = &v
	return s
}

func (s *QueryPfPaymentResponse) SetAccountinStatus(v string) *QueryPfPaymentResponse {
	s.AccountinStatus = &v
	return s
}

func (s *QueryPfPaymentResponse) SetTradeTime(v string) *QueryPfPaymentResponse {
	s.TradeTime = &v
	return s
}

type QueryPfIouRequest struct {
	// OAuth模式下的授权token
	AuthToken         *string `json:"auth_token,omitempty" xml:"auth_token,omitempty"`
	ProductInstanceId *string `json:"product_instance_id,omitempty" xml:"product_instance_id,omitempty"`
	// 客户号
	CustomerNo *string `json:"customer_no,omitempty" xml:"customer_no,omitempty" maxLength:"20"`
	// 借据Id
	DebitId *string `json:"debit_id,omitempty" xml:"debit_id,omitempty" maxLength:"16"`
	// 支用Id
	FinancingId *string `json:"financing_id,omitempty" xml:"financing_id,omitempty" require:"true" maxLength:"32"`
}

func (s QueryPfIouRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryPfIouRequest) GoString() string {
	return s.String()
}

func (s *QueryPfIouRequest) SetAuthToken(v string) *QueryPfIouRequest {
	s.AuthToken = &v
	return s
}

func (s *QueryPfIouRequest) SetProductInstanceId(v string) *QueryPfIouRequest {
	s.ProductInstanceId = &v
	return s
}

func (s *QueryPfIouRequest) SetCustomerNo(v string) *QueryPfIouRequest {
	s.CustomerNo = &v
	return s
}

func (s *QueryPfIouRequest) SetDebitId(v string) *QueryPfIouRequest {
	s.DebitId = &v
	return s
}

func (s *QueryPfIouRequest) SetFinancingId(v string) *QueryPfIouRequest {
	s.FinancingId = &v
	return s
}

type QueryPfIouResponse struct {
	// 请求唯一ID，用于链路跟踪和问题排查
	ReqMsgId *string `json:"req_msg_id,omitempty" xml:"req_msg_id,omitempty"`
	// 结果码，一般OK表示调用成功
	ResultCode *string `json:"result_code,omitempty" xml:"result_code,omitempty"`
	// 异常信息的文本描述
	ResultMsg *string `json:"result_msg,omitempty" xml:"result_msg,omitempty"`
	// 执行年利率
	AnnualInterestRate *string `json:"annual_interest_rate,omitempty" xml:"annual_interest_rate,omitempty"`
	// 借据Id
	CreditId *string `json:"credit_id,omitempty" xml:"credit_id,omitempty"`
	// 借据到期日
	InterestBearingEnd *string `json:"interest_bearing_end,omitempty" xml:"interest_bearing_end,omitempty"`
	// 借据起息日
	InterestBearingStart *string `json:"interest_bearing_start,omitempty" xml:"interest_bearing_start,omitempty"`
	// 发放金额
	IssuedAmount *string `json:"issued_amount,omitempty" xml:"issued_amount,omitempty"`
	// 贷款性质 0-正常 1-展期 2-一类逾期 3-二类逾期 4-呆滞 5-呆帐
	LoanNature *string `json:"loan_nature,omitempty" xml:"loan_nature,omitempty"`
	// 贷款状态 销户=结清 0-正常 1-销户 5-已发放未复核入账
	LoanStatus *string `json:"loan_status,omitempty" xml:"loan_status,omitempty"`
	// 下次结息日期
	NextParsingDate *string `json:"next_parsing_date,omitempty" xml:"next_parsing_date,omitempty"`
	// 逾期计息方式 0-逾期利率 1-逾期罚息比例 2-协议违约利率
	OdiCalType *string `json:"odi_cal_type,omitempty" xml:"odi_cal_type,omitempty"`
	// 逾期罚息浮动比率
	OpiFloatingRatio *string `json:"opi_floating_ratio,omitempty" xml:"opi_floating_ratio,omitempty"`
	// 贷款入账账号
	PayAccount *string `json:"pay_account,omitempty" xml:"pay_account,omitempty"`
	// 本金余额
	PrincipalBalance *string `json:"principal_balance,omitempty" xml:"principal_balance,omitempty"`
	// 还款账号
	RepayAccount *string `json:"repay_account,omitempty" xml:"repay_account,omitempty"`
}

func (s QueryPfIouResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryPfIouResponse) GoString() string {
	return s.String()
}

func (s *QueryPfIouResponse) SetReqMsgId(v string) *QueryPfIouResponse {
	s.ReqMsgId = &v
	return s
}

func (s *QueryPfIouResponse) SetResultCode(v string) *QueryPfIouResponse {
	s.ResultCode = &v
	return s
}

func (s *QueryPfIouResponse) SetResultMsg(v string) *QueryPfIouResponse {
	s.ResultMsg = &v
	return s
}

func (s *QueryPfIouResponse) SetAnnualInterestRate(v string) *QueryPfIouResponse {
	s.AnnualInterestRate = &v
	return s
}

func (s *QueryPfIouResponse) SetCreditId(v string) *QueryPfIouResponse {
	s.CreditId = &v
	return s
}

func (s *QueryPfIouResponse) SetInterestBearingEnd(v string) *QueryPfIouResponse {
	s.InterestBearingEnd = &v
	return s
}

func (s *QueryPfIouResponse) SetInterestBearingStart(v string) *QueryPfIouResponse {
	s.InterestBearingStart = &v
	return s
}

func (s *QueryPfIouResponse) SetIssuedAmount(v string) *QueryPfIouResponse {
	s.IssuedAmount = &v
	return s
}

func (s *QueryPfIouResponse) SetLoanNature(v string) *QueryPfIouResponse {
	s.LoanNature = &v
	return s
}

func (s *QueryPfIouResponse) SetLoanStatus(v string) *QueryPfIouResponse {
	s.LoanStatus = &v
	return s
}

func (s *QueryPfIouResponse) SetNextParsingDate(v string) *QueryPfIouResponse {
	s.NextParsingDate = &v
	return s
}

func (s *QueryPfIouResponse) SetOdiCalType(v string) *QueryPfIouResponse {
	s.OdiCalType = &v
	return s
}

func (s *QueryPfIouResponse) SetOpiFloatingRatio(v string) *QueryPfIouResponse {
	s.OpiFloatingRatio = &v
	return s
}

func (s *QueryPfIouResponse) SetPayAccount(v string) *QueryPfIouResponse {
	s.PayAccount = &v
	return s
}

func (s *QueryPfIouResponse) SetPrincipalBalance(v string) *QueryPfIouResponse {
	s.PrincipalBalance = &v
	return s
}

func (s *QueryPfIouResponse) SetRepayAccount(v string) *QueryPfIouResponse {
	s.RepayAccount = &v
	return s
}

type QueryPfQuotaRequest struct {
	// OAuth模式下的授权token
	AuthToken         *string `json:"auth_token,omitempty" xml:"auth_token,omitempty"`
	ProductInstanceId *string `json:"product_instance_id,omitempty" xml:"product_instance_id,omitempty"`
	// 项目标识
	ProjectId *string `json:"project_id,omitempty" xml:"project_id,omitempty" require:"true" maxLength:"64"`
	// 证件号
	CertNo *string `json:"cert_no,omitempty" xml:"cert_no,omitempty" require:"true" maxLength:"20"`
	// 证件类型;050 统一社会信用证代码
	CertType *string `json:"cert_type,omitempty" xml:"cert_type,omitempty" require:"true" maxLength:"3"`
	// 银行端客户号
	CustomerNo *string `json:"customer_no,omitempty" xml:"customer_no,omitempty" require:"true" maxLength:"20"`
	// 融资主体did
	FinancingSubjectDid *string `json:"financing_subject_did,omitempty" xml:"financing_subject_did,omitempty" require:"true"`
}

func (s QueryPfQuotaRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryPfQuotaRequest) GoString() string {
	return s.String()
}

func (s *QueryPfQuotaRequest) SetAuthToken(v string) *QueryPfQuotaRequest {
	s.AuthToken = &v
	return s
}

func (s *QueryPfQuotaRequest) SetProductInstanceId(v string) *QueryPfQuotaRequest {
	s.ProductInstanceId = &v
	return s
}

func (s *QueryPfQuotaRequest) SetProjectId(v string) *QueryPfQuotaRequest {
	s.ProjectId = &v
	return s
}

func (s *QueryPfQuotaRequest) SetCertNo(v string) *QueryPfQuotaRequest {
	s.CertNo = &v
	return s
}

func (s *QueryPfQuotaRequest) SetCertType(v string) *QueryPfQuotaRequest {
	s.CertType = &v
	return s
}

func (s *QueryPfQuotaRequest) SetCustomerNo(v string) *QueryPfQuotaRequest {
	s.CustomerNo = &v
	return s
}

func (s *QueryPfQuotaRequest) SetFinancingSubjectDid(v string) *QueryPfQuotaRequest {
	s.FinancingSubjectDid = &v
	return s
}

type QueryPfQuotaResponse struct {
	// 请求唯一ID，用于链路跟踪和问题排查
	ReqMsgId *string `json:"req_msg_id,omitempty" xml:"req_msg_id,omitempty"`
	// 结果码，一般OK表示调用成功
	ResultCode *string `json:"result_code,omitempty" xml:"result_code,omitempty"`
	// 异常信息的文本描述
	ResultMsg *string `json:"result_msg,omitempty" xml:"result_msg,omitempty"`
	// 授信额度信息
	QuotaInfo []*PfCreditQuotaInfo `json:"quota_info,omitempty" xml:"quota_info,omitempty" type:"Repeated"`
}

func (s QueryPfQuotaResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryPfQuotaResponse) GoString() string {
	return s.String()
}

func (s *QueryPfQuotaResponse) SetReqMsgId(v string) *QueryPfQuotaResponse {
	s.ReqMsgId = &v
	return s
}

func (s *QueryPfQuotaResponse) SetResultCode(v string) *QueryPfQuotaResponse {
	s.ResultCode = &v
	return s
}

func (s *QueryPfQuotaResponse) SetResultMsg(v string) *QueryPfQuotaResponse {
	s.ResultMsg = &v
	return s
}

func (s *QueryPfQuotaResponse) SetQuotaInfo(v []*PfCreditQuotaInfo) *QueryPfQuotaResponse {
	s.QuotaInfo = v
	return s
}

type ApplyPfWaybillfinancingRequest struct {
	// OAuth模式下的授权token
	AuthToken         *string `json:"auth_token,omitempty" xml:"auth_token,omitempty"`
	ProductInstanceId *string `json:"product_instance_id,omitempty" xml:"product_instance_id,omitempty"`
	// 请求号，以时间串yyyyMMdd 开头，要求该请求号在请求方系统内唯一；同时该字段也是幂等字段
	RequestNo *string `json:"request_no,omitempty" xml:"request_no,omitempty" require:"true" maxLength:"23"`
	// 项目标识；与对接同学确认对应的标识值
	ProjectId *string `json:"project_id,omitempty" xml:"project_id,omitempty" require:"true" maxLength:"64"`
	// 承运商did
	CarrierDid *string `json:"carrier_did,omitempty" xml:"carrier_did,omitempty" require:"true" maxLength:"80"`
	// 融资支用金额；总长度最长20位，保留2位小数，四舍五入
	FinancingAmount *string `json:"financing_amount,omitempty" xml:"financing_amount,omitempty" require:"true"`
	// 支用币种，支持 CNY,USD
	FinancingCurrency *string `json:"financing_currency,omitempty" xml:"financing_currency,omitempty" require:"true"`
	// 融资授信主体公司分布式数字身份
	FinancingSubjectDid *string `json:"financing_subject_did,omitempty" xml:"financing_subject_did,omitempty" require:"true" maxLength:"80"`
	// 贷款期限，值为2~6，单位为月
	LoanTerm *string `json:"loan_term,omitempty" xml:"loan_term,omitempty" require:"true" maxLength:"6" minLength:"2"`
	// 收款方开户行总行联行号
	PayeeBankNumber *string `json:"payee_bank_number,omitempty" xml:"payee_bank_number,omitempty" require:"true" maxLength:"32"`
	// 收款方开户行名称
	PayeeBcb *string `json:"payee_bcb,omitempty" xml:"payee_bcb,omitempty" require:"true" maxLength:"200"`
	// 收款方开户行银行卡号
	PayeeBcbCardNo *string `json:"payee_bcb_card_no,omitempty" xml:"payee_bcb_card_no,omitempty" require:"true"`
	// 收款方证件号
	//
	//
	PayeeIdNumber *string `json:"payee_id_number,omitempty" xml:"payee_id_number,omitempty" require:"true" maxLength:"40"`
	// 收款方证件类型
	PayeeIdType *string `json:"payee_id_type,omitempty" xml:"payee_id_type,omitempty" require:"true" maxLength:"5"`
	// 收款方名称
	PayeeName *string `json:"payee_name,omitempty" xml:"payee_name,omitempty" require:"true" maxLength:"200"`
	// 0政府投标 1经营周转 2支付货款 3采购机票，一般默认填 2
	Purpose *string `json:"purpose,omitempty" xml:"purpose,omitempty" require:"true"`
	// 银行端的Ukey签名；使用方调用接口前使用银行Ukey做签名，并将签名后的结果填入该字段；一期，该字段可不传，使用方通过登录网上银行使用网银进行确认
	Signature *string `json:"signature,omitempty" xml:"signature,omitempty" maxLength:"800"`
	// 转账附言
	// 1: 工资、奖金收入
	// 2：稿费、演出费等劳务收入
	// 3：债券、期货、信托等投资的本金和收益
	// 4：个人债券或产权转让收益
	// 该字段建议填写 2
	TransferPostscript *string `json:"transfer_postscript,omitempty" xml:"transfer_postscript,omitempty" require:"true" maxLength:"3"`
	// 凭证类型，支持 WAYBILL,PAYABLE
	VoucherCategory *string `json:"voucher_category,omitempty" xml:"voucher_category,omitempty" require:"true"`
	// 支用凭证ids，支持多个，逗号隔开；
	VoucherIds *string `json:"voucher_ids,omitempty" xml:"voucher_ids,omitempty" require:"true" maxLength:"2000"`
	// 8位发票号，支持多个，逗号分隔
	VoucherInvoiceCodes *string `json:"voucher_invoice_codes,omitempty" xml:"voucher_invoice_codes,omitempty" maxLength:"800"`
	// 提款确认书hash
	ConfirmationHash *string `json:"confirmation_hash,omitempty" xml:"confirmation_hash,omitempty"`
	// 提款确认书 osskey
	ConfirmationOssKey *string `json:"confirmation_oss_key,omitempty" xml:"confirmation_oss_key,omitempty"`
	// 签名公钥
	SignaturePubKey *string `json:"signature_pub_key,omitempty" xml:"signature_pub_key,omitempty"`
	// 网银操作员账号
	//
	//
	OnlineBankOperatorAccount *string `json:"online_bank_operator_account,omitempty" xml:"online_bank_operator_account,omitempty"`
}

func (s ApplyPfWaybillfinancingRequest) String() string {
	return tea.Prettify(s)
}

func (s ApplyPfWaybillfinancingRequest) GoString() string {
	return s.String()
}

func (s *ApplyPfWaybillfinancingRequest) SetAuthToken(v string) *ApplyPfWaybillfinancingRequest {
	s.AuthToken = &v
	return s
}

func (s *ApplyPfWaybillfinancingRequest) SetProductInstanceId(v string) *ApplyPfWaybillfinancingRequest {
	s.ProductInstanceId = &v
	return s
}

func (s *ApplyPfWaybillfinancingRequest) SetRequestNo(v string) *ApplyPfWaybillfinancingRequest {
	s.RequestNo = &v
	return s
}

func (s *ApplyPfWaybillfinancingRequest) SetProjectId(v string) *ApplyPfWaybillfinancingRequest {
	s.ProjectId = &v
	return s
}

func (s *ApplyPfWaybillfinancingRequest) SetCarrierDid(v string) *ApplyPfWaybillfinancingRequest {
	s.CarrierDid = &v
	return s
}

func (s *ApplyPfWaybillfinancingRequest) SetFinancingAmount(v string) *ApplyPfWaybillfinancingRequest {
	s.FinancingAmount = &v
	return s
}

func (s *ApplyPfWaybillfinancingRequest) SetFinancingCurrency(v string) *ApplyPfWaybillfinancingRequest {
	s.FinancingCurrency = &v
	return s
}

func (s *ApplyPfWaybillfinancingRequest) SetFinancingSubjectDid(v string) *ApplyPfWaybillfinancingRequest {
	s.FinancingSubjectDid = &v
	return s
}

func (s *ApplyPfWaybillfinancingRequest) SetLoanTerm(v string) *ApplyPfWaybillfinancingRequest {
	s.LoanTerm = &v
	return s
}

func (s *ApplyPfWaybillfinancingRequest) SetPayeeBankNumber(v string) *ApplyPfWaybillfinancingRequest {
	s.PayeeBankNumber = &v
	return s
}

func (s *ApplyPfWaybillfinancingRequest) SetPayeeBcb(v string) *ApplyPfWaybillfinancingRequest {
	s.PayeeBcb = &v
	return s
}

func (s *ApplyPfWaybillfinancingRequest) SetPayeeBcbCardNo(v string) *ApplyPfWaybillfinancingRequest {
	s.PayeeBcbCardNo = &v
	return s
}

func (s *ApplyPfWaybillfinancingRequest) SetPayeeIdNumber(v string) *ApplyPfWaybillfinancingRequest {
	s.PayeeIdNumber = &v
	return s
}

func (s *ApplyPfWaybillfinancingRequest) SetPayeeIdType(v string) *ApplyPfWaybillfinancingRequest {
	s.PayeeIdType = &v
	return s
}

func (s *ApplyPfWaybillfinancingRequest) SetPayeeName(v string) *ApplyPfWaybillfinancingRequest {
	s.PayeeName = &v
	return s
}

func (s *ApplyPfWaybillfinancingRequest) SetPurpose(v string) *ApplyPfWaybillfinancingRequest {
	s.Purpose = &v
	return s
}

func (s *ApplyPfWaybillfinancingRequest) SetSignature(v string) *ApplyPfWaybillfinancingRequest {
	s.Signature = &v
	return s
}

func (s *ApplyPfWaybillfinancingRequest) SetTransferPostscript(v string) *ApplyPfWaybillfinancingRequest {
	s.TransferPostscript = &v
	return s
}

func (s *ApplyPfWaybillfinancingRequest) SetVoucherCategory(v string) *ApplyPfWaybillfinancingRequest {
	s.VoucherCategory = &v
	return s
}

func (s *ApplyPfWaybillfinancingRequest) SetVoucherIds(v string) *ApplyPfWaybillfinancingRequest {
	s.VoucherIds = &v
	return s
}

func (s *ApplyPfWaybillfinancingRequest) SetVoucherInvoiceCodes(v string) *ApplyPfWaybillfinancingRequest {
	s.VoucherInvoiceCodes = &v
	return s
}

func (s *ApplyPfWaybillfinancingRequest) SetConfirmationHash(v string) *ApplyPfWaybillfinancingRequest {
	s.ConfirmationHash = &v
	return s
}

func (s *ApplyPfWaybillfinancingRequest) SetConfirmationOssKey(v string) *ApplyPfWaybillfinancingRequest {
	s.ConfirmationOssKey = &v
	return s
}

func (s *ApplyPfWaybillfinancingRequest) SetSignaturePubKey(v string) *ApplyPfWaybillfinancingRequest {
	s.SignaturePubKey = &v
	return s
}

func (s *ApplyPfWaybillfinancingRequest) SetOnlineBankOperatorAccount(v string) *ApplyPfWaybillfinancingRequest {
	s.OnlineBankOperatorAccount = &v
	return s
}

type ApplyPfWaybillfinancingResponse struct {
	// 请求唯一ID，用于链路跟踪和问题排查
	ReqMsgId *string `json:"req_msg_id,omitempty" xml:"req_msg_id,omitempty"`
	// 结果码，一般OK表示调用成功
	ResultCode *string `json:"result_code,omitempty" xml:"result_code,omitempty"`
	// 异常信息的文本描述
	ResultMsg *string `json:"result_msg,omitempty" xml:"result_msg,omitempty"`
	// 融资支用Id
	FinancingId *string `json:"financing_id,omitempty" xml:"financing_id,omitempty"`
	// 申请状态，
	// 融资申请失败 IN_VALID
	// 融资申请失败 F_APPLY_FAIL
	// 融资申请成功 F_APPLY_SUC
	// 融资申请成功 COMMIT
	// 融资核验中 VERIFYING
	// 融资核验失败 VERIFY_FAILURE
	// 融资订单生成成功 P_WITHDRAW
	// 融资成功 A_WITHDRAW
	// 融资失败 EXPIRED
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
}

func (s ApplyPfWaybillfinancingResponse) String() string {
	return tea.Prettify(s)
}

func (s ApplyPfWaybillfinancingResponse) GoString() string {
	return s.String()
}

func (s *ApplyPfWaybillfinancingResponse) SetReqMsgId(v string) *ApplyPfWaybillfinancingResponse {
	s.ReqMsgId = &v
	return s
}

func (s *ApplyPfWaybillfinancingResponse) SetResultCode(v string) *ApplyPfWaybillfinancingResponse {
	s.ResultCode = &v
	return s
}

func (s *ApplyPfWaybillfinancingResponse) SetResultMsg(v string) *ApplyPfWaybillfinancingResponse {
	s.ResultMsg = &v
	return s
}

func (s *ApplyPfWaybillfinancingResponse) SetFinancingId(v string) *ApplyPfWaybillfinancingResponse {
	s.FinancingId = &v
	return s
}

func (s *ApplyPfWaybillfinancingResponse) SetStatus(v string) *ApplyPfWaybillfinancingResponse {
	s.Status = &v
	return s
}

type PushPfPledgeRequest struct {
	// OAuth模式下的授权token
	AuthToken         *string `json:"auth_token,omitempty" xml:"auth_token,omitempty"`
	ProductInstanceId *string `json:"product_instance_id,omitempty" xml:"product_instance_id,omitempty"`
	// 项目标识，可联系对接同学提供
	ProjectId *string `json:"project_id,omitempty" xml:"project_id,omitempty" require:"true" maxLength:"64"`
	// 账单标识
	BillId *string `json:"bill_id,omitempty" xml:"bill_id,omitempty" require:"true" maxLength:"128" minLength:"1"`
	// 融资主体did
	FinancingSubjectDid *string `json:"financing_subject_did,omitempty" xml:"financing_subject_did,omitempty" require:"true" maxLength:"128" minLength:"1"`
	// 请求号；以yyyyMMdd 时间串开头的32位字符串；该字符串需要保持请求系统内唯一，系统会以该请求号作为幂等处理
	RequestNo *string `json:"request_no,omitempty" xml:"request_no,omitempty" require:"true" maxLength:"32" minLength:"16"`
}

func (s PushPfPledgeRequest) String() string {
	return tea.Prettify(s)
}

func (s PushPfPledgeRequest) GoString() string {
	return s.String()
}

func (s *PushPfPledgeRequest) SetAuthToken(v string) *PushPfPledgeRequest {
	s.AuthToken = &v
	return s
}

func (s *PushPfPledgeRequest) SetProductInstanceId(v string) *PushPfPledgeRequest {
	s.ProductInstanceId = &v
	return s
}

func (s *PushPfPledgeRequest) SetProjectId(v string) *PushPfPledgeRequest {
	s.ProjectId = &v
	return s
}

func (s *PushPfPledgeRequest) SetBillId(v string) *PushPfPledgeRequest {
	s.BillId = &v
	return s
}

func (s *PushPfPledgeRequest) SetFinancingSubjectDid(v string) *PushPfPledgeRequest {
	s.FinancingSubjectDid = &v
	return s
}

func (s *PushPfPledgeRequest) SetRequestNo(v string) *PushPfPledgeRequest {
	s.RequestNo = &v
	return s
}

type PushPfPledgeResponse struct {
	// 请求唯一ID，用于链路跟踪和问题排查
	ReqMsgId *string `json:"req_msg_id,omitempty" xml:"req_msg_id,omitempty"`
	// 结果码，一般OK表示调用成功
	ResultCode *string `json:"result_code,omitempty" xml:"result_code,omitempty"`
	// 异常信息的文本描述
	ResultMsg *string `json:"result_msg,omitempty" xml:"result_msg,omitempty"`
	// 质押标识，用于账单质押推送状态查询
	PledgeId *string `json:"pledge_id,omitempty" xml:"pledge_id,omitempty"`
	// 已提交  COMMIT
	// 推送中 PUSHING
	// 推送成功 PUSH_SUC
	// 推送失败 PUSH_FAIL
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// 描述
	Desc *string `json:"desc,omitempty" xml:"desc,omitempty"`
}

func (s PushPfPledgeResponse) String() string {
	return tea.Prettify(s)
}

func (s PushPfPledgeResponse) GoString() string {
	return s.String()
}

func (s *PushPfPledgeResponse) SetReqMsgId(v string) *PushPfPledgeResponse {
	s.ReqMsgId = &v
	return s
}

func (s *PushPfPledgeResponse) SetResultCode(v string) *PushPfPledgeResponse {
	s.ResultCode = &v
	return s
}

func (s *PushPfPledgeResponse) SetResultMsg(v string) *PushPfPledgeResponse {
	s.ResultMsg = &v
	return s
}

func (s *PushPfPledgeResponse) SetPledgeId(v string) *PushPfPledgeResponse {
	s.PledgeId = &v
	return s
}

func (s *PushPfPledgeResponse) SetStatus(v string) *PushPfPledgeResponse {
	s.Status = &v
	return s
}

func (s *PushPfPledgeResponse) SetDesc(v string) *PushPfPledgeResponse {
	s.Desc = &v
	return s
}

type QueryPfPledgeRequest struct {
	// OAuth模式下的授权token
	AuthToken         *string `json:"auth_token,omitempty" xml:"auth_token,omitempty"`
	ProductInstanceId *string `json:"product_instance_id,omitempty" xml:"product_instance_id,omitempty"`
	// 质押id
	PledgeId *string `json:"pledge_id,omitempty" xml:"pledge_id,omitempty" require:"true" maxLength:"32" minLength:"16"`
}

func (s QueryPfPledgeRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryPfPledgeRequest) GoString() string {
	return s.String()
}

func (s *QueryPfPledgeRequest) SetAuthToken(v string) *QueryPfPledgeRequest {
	s.AuthToken = &v
	return s
}

func (s *QueryPfPledgeRequest) SetProductInstanceId(v string) *QueryPfPledgeRequest {
	s.ProductInstanceId = &v
	return s
}

func (s *QueryPfPledgeRequest) SetPledgeId(v string) *QueryPfPledgeRequest {
	s.PledgeId = &v
	return s
}

type QueryPfPledgeResponse struct {
	// 请求唯一ID，用于链路跟踪和问题排查
	ReqMsgId *string `json:"req_msg_id,omitempty" xml:"req_msg_id,omitempty"`
	// 结果码，一般OK表示调用成功
	ResultCode *string `json:"result_code,omitempty" xml:"result_code,omitempty"`
	// 异常信息的文本描述
	ResultMsg *string `json:"result_msg,omitempty" xml:"result_msg,omitempty"`
	// 账单id
	BillId *string `json:"bill_id,omitempty" xml:"bill_id,omitempty"`
	// 融资主体did
	FinancingSubjectDid *string `json:"financing_subject_did,omitempty" xml:"financing_subject_did,omitempty"`
	// 已提交  COMMIT
	// 推送中 PUSHING
	// 推送成功 PUSH_SUC
	// 推送失败 PUSH_FAIL
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// 状态描述
	Desc *string `json:"desc,omitempty" xml:"desc,omitempty"`
}

func (s QueryPfPledgeResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryPfPledgeResponse) GoString() string {
	return s.String()
}

func (s *QueryPfPledgeResponse) SetReqMsgId(v string) *QueryPfPledgeResponse {
	s.ReqMsgId = &v
	return s
}

func (s *QueryPfPledgeResponse) SetResultCode(v string) *QueryPfPledgeResponse {
	s.ResultCode = &v
	return s
}

func (s *QueryPfPledgeResponse) SetResultMsg(v string) *QueryPfPledgeResponse {
	s.ResultMsg = &v
	return s
}

func (s *QueryPfPledgeResponse) SetBillId(v string) *QueryPfPledgeResponse {
	s.BillId = &v
	return s
}

func (s *QueryPfPledgeResponse) SetFinancingSubjectDid(v string) *QueryPfPledgeResponse {
	s.FinancingSubjectDid = &v
	return s
}

func (s *QueryPfPledgeResponse) SetStatus(v string) *QueryPfPledgeResponse {
	s.Status = &v
	return s
}

func (s *QueryPfPledgeResponse) SetDesc(v string) *QueryPfPledgeResponse {
	s.Desc = &v
	return s
}

type QueryPfFinancingRequest struct {
	// OAuth模式下的授权token
	AuthToken         *string `json:"auth_token,omitempty" xml:"auth_token,omitempty"`
	ProductInstanceId *string `json:"product_instance_id,omitempty" xml:"product_instance_id,omitempty"`
	// 支用Id
	FinancingId *string `json:"financing_id,omitempty" xml:"financing_id,omitempty" require:"true"`
}

func (s QueryPfFinancingRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryPfFinancingRequest) GoString() string {
	return s.String()
}

func (s *QueryPfFinancingRequest) SetAuthToken(v string) *QueryPfFinancingRequest {
	s.AuthToken = &v
	return s
}

func (s *QueryPfFinancingRequest) SetProductInstanceId(v string) *QueryPfFinancingRequest {
	s.ProductInstanceId = &v
	return s
}

func (s *QueryPfFinancingRequest) SetFinancingId(v string) *QueryPfFinancingRequest {
	s.FinancingId = &v
	return s
}

type QueryPfFinancingResponse struct {
	// 请求唯一ID，用于链路跟踪和问题排查
	ReqMsgId *string `json:"req_msg_id,omitempty" xml:"req_msg_id,omitempty"`
	// 结果码，一般OK表示调用成功
	ResultCode *string `json:"result_code,omitempty" xml:"result_code,omitempty"`
	// 异常信息的文本描述
	ResultMsg *string `json:"result_msg,omitempty" xml:"result_msg,omitempty"`
	// 支用Id
	FinancingId *string `json:"financing_id,omitempty" xml:"financing_id,omitempty"`
	// 状态，说明如下：
	// 融资申请失败 IN_VALID;
	// 融资申请失败 F_APPLY_FAIL;
	// 融资申请成功 F_APPLY_SUC;
	// 融资申请成功 COMMIT;
	// 融资核验中 VERIFYING;
	// 融资核验失败 VERIFY_FAILURE;
	// 融资订单生成成功 P_WITHDRAW;
	// 融资成功 A_WITHDRAW;
	// 融资失败 EXPIRED
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// 描述
	Desc *string `json:"desc,omitempty" xml:"desc,omitempty"`
}

func (s QueryPfFinancingResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryPfFinancingResponse) GoString() string {
	return s.String()
}

func (s *QueryPfFinancingResponse) SetReqMsgId(v string) *QueryPfFinancingResponse {
	s.ReqMsgId = &v
	return s
}

func (s *QueryPfFinancingResponse) SetResultCode(v string) *QueryPfFinancingResponse {
	s.ResultCode = &v
	return s
}

func (s *QueryPfFinancingResponse) SetResultMsg(v string) *QueryPfFinancingResponse {
	s.ResultMsg = &v
	return s
}

func (s *QueryPfFinancingResponse) SetFinancingId(v string) *QueryPfFinancingResponse {
	s.FinancingId = &v
	return s
}

func (s *QueryPfFinancingResponse) SetStatus(v string) *QueryPfFinancingResponse {
	s.Status = &v
	return s
}

func (s *QueryPfFinancingResponse) SetDesc(v string) *QueryPfFinancingResponse {
	s.Desc = &v
	return s
}

type CheckPfVoucherRequest struct {
	// OAuth模式下的授权token
	AuthToken         *string `json:"auth_token,omitempty" xml:"auth_token,omitempty"`
	ProductInstanceId *string `json:"product_instance_id,omitempty" xml:"product_instance_id,omitempty"`
	// 融资主体did
	FinancingSubjectDid *string `json:"financing_subject_did,omitempty" xml:"financing_subject_did,omitempty" require:"true" maxLength:"80"`
	// 运单所属平台did
	PlatformDid *string `json:"platform_did,omitempty" xml:"platform_did,omitempty" require:"true" maxLength:"80"`
	// 凭证类型，支持 WAYBILL(运单)
	VoucherCategory *string `json:"voucher_category,omitempty" xml:"voucher_category,omitempty" require:"true" maxLength:"100" minLength:"1"`
	// 支持多个，逗号隔开
	VoucherIds *string `json:"voucher_ids,omitempty" xml:"voucher_ids,omitempty" require:"true" maxLength:"3000" minLength:"1"`
}

func (s CheckPfVoucherRequest) String() string {
	return tea.Prettify(s)
}

func (s CheckPfVoucherRequest) GoString() string {
	return s.String()
}

func (s *CheckPfVoucherRequest) SetAuthToken(v string) *CheckPfVoucherRequest {
	s.AuthToken = &v
	return s
}

func (s *CheckPfVoucherRequest) SetProductInstanceId(v string) *CheckPfVoucherRequest {
	s.ProductInstanceId = &v
	return s
}

func (s *CheckPfVoucherRequest) SetFinancingSubjectDid(v string) *CheckPfVoucherRequest {
	s.FinancingSubjectDid = &v
	return s
}

func (s *CheckPfVoucherRequest) SetPlatformDid(v string) *CheckPfVoucherRequest {
	s.PlatformDid = &v
	return s
}

func (s *CheckPfVoucherRequest) SetVoucherCategory(v string) *CheckPfVoucherRequest {
	s.VoucherCategory = &v
	return s
}

func (s *CheckPfVoucherRequest) SetVoucherIds(v string) *CheckPfVoucherRequest {
	s.VoucherIds = &v
	return s
}

type CheckPfVoucherResponse struct {
	// 请求唯一ID，用于链路跟踪和问题排查
	ReqMsgId *string `json:"req_msg_id,omitempty" xml:"req_msg_id,omitempty"`
	// 结果码，一般OK表示调用成功
	ResultCode *string `json:"result_code,omitempty" xml:"result_code,omitempty"`
	// 异常信息的文本描述
	ResultMsg *string `json:"result_msg,omitempty" xml:"result_msg,omitempty"`
	// 核验结果
	CheckResults []*PfVoucherCheckResult `json:"check_results,omitempty" xml:"check_results,omitempty" type:"Repeated"`
}

func (s CheckPfVoucherResponse) String() string {
	return tea.Prettify(s)
}

func (s CheckPfVoucherResponse) GoString() string {
	return s.String()
}

func (s *CheckPfVoucherResponse) SetReqMsgId(v string) *CheckPfVoucherResponse {
	s.ReqMsgId = &v
	return s
}

func (s *CheckPfVoucherResponse) SetResultCode(v string) *CheckPfVoucherResponse {
	s.ResultCode = &v
	return s
}

func (s *CheckPfVoucherResponse) SetResultMsg(v string) *CheckPfVoucherResponse {
	s.ResultMsg = &v
	return s
}

func (s *CheckPfVoucherResponse) SetCheckResults(v []*PfVoucherCheckResult) *CheckPfVoucherResponse {
	s.CheckResults = v
	return s
}

type ApplyPfConfirmationRequest struct {
	// OAuth模式下的授权token
	AuthToken         *string `json:"auth_token,omitempty" xml:"auth_token,omitempty"`
	ProductInstanceId *string `json:"product_instance_id,omitempty" xml:"product_instance_id,omitempty"`
	// 请求号
	RequestNo *string `json:"request_no,omitempty" xml:"request_no,omitempty" require:"true"`
	// 项目标识，可联系对接同学获取
	ProjectId *string `json:"project_id,omitempty" xml:"project_id,omitempty" require:"true" maxLength:"64"`
	// 客户号
	ClientNo *string `json:"client_no,omitempty" xml:"client_no,omitempty" require:"true"`
	// 融资主体did
	FinancingSubjectDid *string `json:"financing_subject_did,omitempty" xml:"financing_subject_did,omitempty" require:"true"`
	// 借款人证件类型
	BorrowerCardType *string `json:"borrower_card_type,omitempty" xml:"borrower_card_type,omitempty" require:"true"`
	// 借款人证件号码
	//
	//
	BorrowerCardNo *string `json:"borrower_card_no,omitempty" xml:"borrower_card_no,omitempty" require:"true"`
	// 融资金额
	//
	//
	FinancingAmount *string `json:"financing_amount,omitempty" xml:"financing_amount,omitempty" require:"true"`
	// 支用币种
	//
	//
	Currency *string `json:"currency,omitempty" xml:"currency,omitempty" require:"true"`
	// 贷款期限（月）
	LoanTerm *string `json:"loan_term,omitempty" xml:"loan_term,omitempty" require:"true"`
	// 收款方开户行总行联行号
	//
	//
	PayeeBankUnionNumber *string `json:"payee_bank_union_number,omitempty" xml:"payee_bank_union_number,omitempty" require:"true"`
	// 收款方开户行名称
	//
	//
	PayeeBankName *string `json:"payee_bank_name,omitempty" xml:"payee_bank_name,omitempty" require:"true"`
	// 收款方户名
	//
	//
	PayeeName *string `json:"payee_name,omitempty" xml:"payee_name,omitempty" require:"true"`
	// 收款方银行卡账号
	//
	//
	PayeeBankNo *string `json:"payee_bank_no,omitempty" xml:"payee_bank_no,omitempty" require:"true"`
	// 收款人证件类型
	//
	//
	PayeeCardType *string `json:"payee_card_type,omitempty" xml:"payee_card_type,omitempty"`
	// 收款人证件号码
	//
	//
	PayeeCardNo *string `json:"payee_card_no,omitempty" xml:"payee_card_no,omitempty"`
	// 贷款用途
	//
	//
	Purpose *string `json:"purpose,omitempty" xml:"purpose,omitempty" require:"true"`
	// 附言
	Postscript *string `json:"postscript,omitempty" xml:"postscript,omitempty"`
}

func (s ApplyPfConfirmationRequest) String() string {
	return tea.Prettify(s)
}

func (s ApplyPfConfirmationRequest) GoString() string {
	return s.String()
}

func (s *ApplyPfConfirmationRequest) SetAuthToken(v string) *ApplyPfConfirmationRequest {
	s.AuthToken = &v
	return s
}

func (s *ApplyPfConfirmationRequest) SetProductInstanceId(v string) *ApplyPfConfirmationRequest {
	s.ProductInstanceId = &v
	return s
}

func (s *ApplyPfConfirmationRequest) SetRequestNo(v string) *ApplyPfConfirmationRequest {
	s.RequestNo = &v
	return s
}

func (s *ApplyPfConfirmationRequest) SetProjectId(v string) *ApplyPfConfirmationRequest {
	s.ProjectId = &v
	return s
}

func (s *ApplyPfConfirmationRequest) SetClientNo(v string) *ApplyPfConfirmationRequest {
	s.ClientNo = &v
	return s
}

func (s *ApplyPfConfirmationRequest) SetFinancingSubjectDid(v string) *ApplyPfConfirmationRequest {
	s.FinancingSubjectDid = &v
	return s
}

func (s *ApplyPfConfirmationRequest) SetBorrowerCardType(v string) *ApplyPfConfirmationRequest {
	s.BorrowerCardType = &v
	return s
}

func (s *ApplyPfConfirmationRequest) SetBorrowerCardNo(v string) *ApplyPfConfirmationRequest {
	s.BorrowerCardNo = &v
	return s
}

func (s *ApplyPfConfirmationRequest) SetFinancingAmount(v string) *ApplyPfConfirmationRequest {
	s.FinancingAmount = &v
	return s
}

func (s *ApplyPfConfirmationRequest) SetCurrency(v string) *ApplyPfConfirmationRequest {
	s.Currency = &v
	return s
}

func (s *ApplyPfConfirmationRequest) SetLoanTerm(v string) *ApplyPfConfirmationRequest {
	s.LoanTerm = &v
	return s
}

func (s *ApplyPfConfirmationRequest) SetPayeeBankUnionNumber(v string) *ApplyPfConfirmationRequest {
	s.PayeeBankUnionNumber = &v
	return s
}

func (s *ApplyPfConfirmationRequest) SetPayeeBankName(v string) *ApplyPfConfirmationRequest {
	s.PayeeBankName = &v
	return s
}

func (s *ApplyPfConfirmationRequest) SetPayeeName(v string) *ApplyPfConfirmationRequest {
	s.PayeeName = &v
	return s
}

func (s *ApplyPfConfirmationRequest) SetPayeeBankNo(v string) *ApplyPfConfirmationRequest {
	s.PayeeBankNo = &v
	return s
}

func (s *ApplyPfConfirmationRequest) SetPayeeCardType(v string) *ApplyPfConfirmationRequest {
	s.PayeeCardType = &v
	return s
}

func (s *ApplyPfConfirmationRequest) SetPayeeCardNo(v string) *ApplyPfConfirmationRequest {
	s.PayeeCardNo = &v
	return s
}

func (s *ApplyPfConfirmationRequest) SetPurpose(v string) *ApplyPfConfirmationRequest {
	s.Purpose = &v
	return s
}

func (s *ApplyPfConfirmationRequest) SetPostscript(v string) *ApplyPfConfirmationRequest {
	s.Postscript = &v
	return s
}

type ApplyPfConfirmationResponse struct {
	// 请求唯一ID，用于链路跟踪和问题排查
	ReqMsgId *string `json:"req_msg_id,omitempty" xml:"req_msg_id,omitempty"`
	// 结果码，一般OK表示调用成功
	ResultCode *string `json:"result_code,omitempty" xml:"result_code,omitempty"`
	// 异常信息的文本描述
	ResultMsg *string `json:"result_msg,omitempty" xml:"result_msg,omitempty"`
	// 融资提款确认书 hash
	//
	ConfirmationHash *string `json:"confirmation_hash,omitempty" xml:"confirmation_hash,omitempty"`
	// 融资提款确认书 OSSkey
	//
	//
	ConfirmationOssKey *string `json:"confirmation_oss_key,omitempty" xml:"confirmation_oss_key,omitempty"`
	// 贷款起始日期，格式：yyyyMMdd
	//
	//
	LoanStartTime *string `json:"loan_start_time,omitempty" xml:"loan_start_time,omitempty"`
	// 贷款到期日期，格式：yyyyMMdd
	//
	//
	LoanEndTime *string `json:"loan_end_time,omitempty" xml:"loan_end_time,omitempty"`
	// 还款方式
	//
	//
	Repayment *string `json:"repayment,omitempty" xml:"repayment,omitempty"`
	// 贷款利率
	EtrdLnIntRt *string `json:"etrd_ln_int_rt,omitempty" xml:"etrd_ln_int_rt,omitempty"`
}

func (s ApplyPfConfirmationResponse) String() string {
	return tea.Prettify(s)
}

func (s ApplyPfConfirmationResponse) GoString() string {
	return s.String()
}

func (s *ApplyPfConfirmationResponse) SetReqMsgId(v string) *ApplyPfConfirmationResponse {
	s.ReqMsgId = &v
	return s
}

func (s *ApplyPfConfirmationResponse) SetResultCode(v string) *ApplyPfConfirmationResponse {
	s.ResultCode = &v
	return s
}

func (s *ApplyPfConfirmationResponse) SetResultMsg(v string) *ApplyPfConfirmationResponse {
	s.ResultMsg = &v
	return s
}

func (s *ApplyPfConfirmationResponse) SetConfirmationHash(v string) *ApplyPfConfirmationResponse {
	s.ConfirmationHash = &v
	return s
}

func (s *ApplyPfConfirmationResponse) SetConfirmationOssKey(v string) *ApplyPfConfirmationResponse {
	s.ConfirmationOssKey = &v
	return s
}

func (s *ApplyPfConfirmationResponse) SetLoanStartTime(v string) *ApplyPfConfirmationResponse {
	s.LoanStartTime = &v
	return s
}

func (s *ApplyPfConfirmationResponse) SetLoanEndTime(v string) *ApplyPfConfirmationResponse {
	s.LoanEndTime = &v
	return s
}

func (s *ApplyPfConfirmationResponse) SetRepayment(v string) *ApplyPfConfirmationResponse {
	s.Repayment = &v
	return s
}

func (s *ApplyPfConfirmationResponse) SetEtrdLnIntRt(v string) *ApplyPfConfirmationResponse {
	s.EtrdLnIntRt = &v
	return s
}

type ApplyPfFinancingqualificationRequest struct {
	// OAuth模式下的授权token
	AuthToken         *string `json:"auth_token,omitempty" xml:"auth_token,omitempty"`
	ProductInstanceId *string `json:"product_instance_id,omitempty" xml:"product_instance_id,omitempty"`
	// 项目标识，可联系对接同学获取
	ProjectId *string `json:"project_id,omitempty" xml:"project_id,omitempty" require:"true" maxLength:"64"`
	// 客户号
	ClientNo *string `json:"client_no,omitempty" xml:"client_no,omitempty" require:"true"`
	// 融资主体did
	FinancingSubjectDid *string `json:"financing_subject_did,omitempty" xml:"financing_subject_did,omitempty" require:"true"`
	// 借款人证件类型
	BorrowerCardType *string `json:"borrower_card_type,omitempty" xml:"borrower_card_type,omitempty" require:"true"`
	// 借款人证件号码
	//
	//
	BorrowerCardNo *string `json:"borrower_card_no,omitempty" xml:"borrower_card_no,omitempty" require:"true"`
}

func (s ApplyPfFinancingqualificationRequest) String() string {
	return tea.Prettify(s)
}

func (s ApplyPfFinancingqualificationRequest) GoString() string {
	return s.String()
}

func (s *ApplyPfFinancingqualificationRequest) SetAuthToken(v string) *ApplyPfFinancingqualificationRequest {
	s.AuthToken = &v
	return s
}

func (s *ApplyPfFinancingqualificationRequest) SetProductInstanceId(v string) *ApplyPfFinancingqualificationRequest {
	s.ProductInstanceId = &v
	return s
}

func (s *ApplyPfFinancingqualificationRequest) SetProjectId(v string) *ApplyPfFinancingqualificationRequest {
	s.ProjectId = &v
	return s
}

func (s *ApplyPfFinancingqualificationRequest) SetClientNo(v string) *ApplyPfFinancingqualificationRequest {
	s.ClientNo = &v
	return s
}

func (s *ApplyPfFinancingqualificationRequest) SetFinancingSubjectDid(v string) *ApplyPfFinancingqualificationRequest {
	s.FinancingSubjectDid = &v
	return s
}

func (s *ApplyPfFinancingqualificationRequest) SetBorrowerCardType(v string) *ApplyPfFinancingqualificationRequest {
	s.BorrowerCardType = &v
	return s
}

func (s *ApplyPfFinancingqualificationRequest) SetBorrowerCardNo(v string) *ApplyPfFinancingqualificationRequest {
	s.BorrowerCardNo = &v
	return s
}

type ApplyPfFinancingqualificationResponse struct {
	// 请求唯一ID，用于链路跟踪和问题排查
	ReqMsgId *string `json:"req_msg_id,omitempty" xml:"req_msg_id,omitempty"`
	// 结果码，一般OK表示调用成功
	ResultCode *string `json:"result_code,omitempty" xml:"result_code,omitempty"`
	// 异常信息的文本描述
	ResultMsg *string `json:"result_msg,omitempty" xml:"result_msg,omitempty"`
	// 申请流水号
	AplSeqNo *string `json:"apl_seq_no,omitempty" xml:"apl_seq_no,omitempty"`
	// 受理标志
	// 0-受理失败
	// 1-受理成功
	AcceptanceFlag *string `json:"acceptance_flag,omitempty" xml:"acceptance_flag,omitempty"`
}

func (s ApplyPfFinancingqualificationResponse) String() string {
	return tea.Prettify(s)
}

func (s ApplyPfFinancingqualificationResponse) GoString() string {
	return s.String()
}

func (s *ApplyPfFinancingqualificationResponse) SetReqMsgId(v string) *ApplyPfFinancingqualificationResponse {
	s.ReqMsgId = &v
	return s
}

func (s *ApplyPfFinancingqualificationResponse) SetResultCode(v string) *ApplyPfFinancingqualificationResponse {
	s.ResultCode = &v
	return s
}

func (s *ApplyPfFinancingqualificationResponse) SetResultMsg(v string) *ApplyPfFinancingqualificationResponse {
	s.ResultMsg = &v
	return s
}

func (s *ApplyPfFinancingqualificationResponse) SetAplSeqNo(v string) *ApplyPfFinancingqualificationResponse {
	s.AplSeqNo = &v
	return s
}

func (s *ApplyPfFinancingqualificationResponse) SetAcceptanceFlag(v string) *ApplyPfFinancingqualificationResponse {
	s.AcceptanceFlag = &v
	return s
}

type QueryPfFinancingqualificationRequest struct {
	// OAuth模式下的授权token
	AuthToken         *string `json:"auth_token,omitempty" xml:"auth_token,omitempty"`
	ProductInstanceId *string `json:"product_instance_id,omitempty" xml:"product_instance_id,omitempty"`
	// 项目标识，可联系对接同学获取
	ProjectId *string `json:"project_id,omitempty" xml:"project_id,omitempty" require:"true" maxLength:"64"`
	// 客户号
	ClientNo *string `json:"client_no,omitempty" xml:"client_no,omitempty" require:"true"`
	// 融资主体did
	FinancingSubjectDid *string `json:"financing_subject_did,omitempty" xml:"financing_subject_did,omitempty" require:"true"`
	// 借款人证件类型
	BorrowerCardType *string `json:"borrower_card_type,omitempty" xml:"borrower_card_type,omitempty" require:"true"`
	// 借款人证件号码
	//
	//
	BorrowerCardNo *string `json:"borrower_card_no,omitempty" xml:"borrower_card_no,omitempty" require:"true"`
	// 申请流水号
	AplSeqNo *string `json:"apl_seq_no,omitempty" xml:"apl_seq_no,omitempty" require:"true"`
}

func (s QueryPfFinancingqualificationRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryPfFinancingqualificationRequest) GoString() string {
	return s.String()
}

func (s *QueryPfFinancingqualificationRequest) SetAuthToken(v string) *QueryPfFinancingqualificationRequest {
	s.AuthToken = &v
	return s
}

func (s *QueryPfFinancingqualificationRequest) SetProductInstanceId(v string) *QueryPfFinancingqualificationRequest {
	s.ProductInstanceId = &v
	return s
}

func (s *QueryPfFinancingqualificationRequest) SetProjectId(v string) *QueryPfFinancingqualificationRequest {
	s.ProjectId = &v
	return s
}

func (s *QueryPfFinancingqualificationRequest) SetClientNo(v string) *QueryPfFinancingqualificationRequest {
	s.ClientNo = &v
	return s
}

func (s *QueryPfFinancingqualificationRequest) SetFinancingSubjectDid(v string) *QueryPfFinancingqualificationRequest {
	s.FinancingSubjectDid = &v
	return s
}

func (s *QueryPfFinancingqualificationRequest) SetBorrowerCardType(v string) *QueryPfFinancingqualificationRequest {
	s.BorrowerCardType = &v
	return s
}

func (s *QueryPfFinancingqualificationRequest) SetBorrowerCardNo(v string) *QueryPfFinancingqualificationRequest {
	s.BorrowerCardNo = &v
	return s
}

func (s *QueryPfFinancingqualificationRequest) SetAplSeqNo(v string) *QueryPfFinancingqualificationRequest {
	s.AplSeqNo = &v
	return s
}

type QueryPfFinancingqualificationResponse struct {
	// 请求唯一ID，用于链路跟踪和问题排查
	ReqMsgId *string `json:"req_msg_id,omitempty" xml:"req_msg_id,omitempty"`
	// 结果码，一般OK表示调用成功
	ResultCode *string `json:"result_code,omitempty" xml:"result_code,omitempty"`
	// 异常信息的文本描述
	ResultMsg *string `json:"result_msg,omitempty" xml:"result_msg,omitempty"`
	// 请求号
	AplSeqNo *string `json:"apl_seq_no,omitempty" xml:"apl_seq_no,omitempty"`
	// 申请日期
	AppDate *string `json:"app_date,omitempty" xml:"app_date,omitempty"`
	// 审批通过时间
	//
	//
	ApproveDate *string `json:"approve_date,omitempty" xml:"approve_date,omitempty"`
	// 额度协议文件编号
	//
	//
	AmntAgrmntFn *string `json:"amnt_agrmnt_fn,omitempty" xml:"amnt_agrmnt_fn,omitempty"`
	// 处理状态
	// A01-审批中
	// A02-审批失败
	// A03-审批通过
	// A04-审批作废
	ProceStatus *string `json:"proce_status,omitempty" xml:"proce_status,omitempty"`
	// 失败原因
	FailRslt *string `json:"fail_rslt,omitempty" xml:"fail_rslt,omitempty"`
	// 错误码
	GenReason *string `json:"gen_reason,omitempty" xml:"gen_reason,omitempty"`
}

func (s QueryPfFinancingqualificationResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryPfFinancingqualificationResponse) GoString() string {
	return s.String()
}

func (s *QueryPfFinancingqualificationResponse) SetReqMsgId(v string) *QueryPfFinancingqualificationResponse {
	s.ReqMsgId = &v
	return s
}

func (s *QueryPfFinancingqualificationResponse) SetResultCode(v string) *QueryPfFinancingqualificationResponse {
	s.ResultCode = &v
	return s
}

func (s *QueryPfFinancingqualificationResponse) SetResultMsg(v string) *QueryPfFinancingqualificationResponse {
	s.ResultMsg = &v
	return s
}

func (s *QueryPfFinancingqualificationResponse) SetAplSeqNo(v string) *QueryPfFinancingqualificationResponse {
	s.AplSeqNo = &v
	return s
}

func (s *QueryPfFinancingqualificationResponse) SetAppDate(v string) *QueryPfFinancingqualificationResponse {
	s.AppDate = &v
	return s
}

func (s *QueryPfFinancingqualificationResponse) SetApproveDate(v string) *QueryPfFinancingqualificationResponse {
	s.ApproveDate = &v
	return s
}

func (s *QueryPfFinancingqualificationResponse) SetAmntAgrmntFn(v string) *QueryPfFinancingqualificationResponse {
	s.AmntAgrmntFn = &v
	return s
}

func (s *QueryPfFinancingqualificationResponse) SetProceStatus(v string) *QueryPfFinancingqualificationResponse {
	s.ProceStatus = &v
	return s
}

func (s *QueryPfFinancingqualificationResponse) SetFailRslt(v string) *QueryPfFinancingqualificationResponse {
	s.FailRslt = &v
	return s
}

func (s *QueryPfFinancingqualificationResponse) SetGenReason(v string) *QueryPfFinancingqualificationResponse {
	s.GenReason = &v
	return s
}

type CallbackPfDefinpfRequest struct {
	// OAuth模式下的授权token
	AuthToken         *string `json:"auth_token,omitempty" xml:"auth_token,omitempty"`
	ProductInstanceId *string `json:"product_instance_id,omitempty" xml:"product_instance_id,omitempty"`
	// 业务类型
	BizType *string `json:"biz_type,omitempty" xml:"biz_type,omitempty" require:"true"`
	// 回调内容
	ContentInfo *string `json:"content_info,omitempty" xml:"content_info,omitempty" require:"true"`
}

func (s CallbackPfDefinpfRequest) String() string {
	return tea.Prettify(s)
}

func (s CallbackPfDefinpfRequest) GoString() string {
	return s.String()
}

func (s *CallbackPfDefinpfRequest) SetAuthToken(v string) *CallbackPfDefinpfRequest {
	s.AuthToken = &v
	return s
}

func (s *CallbackPfDefinpfRequest) SetProductInstanceId(v string) *CallbackPfDefinpfRequest {
	s.ProductInstanceId = &v
	return s
}

func (s *CallbackPfDefinpfRequest) SetBizType(v string) *CallbackPfDefinpfRequest {
	s.BizType = &v
	return s
}

func (s *CallbackPfDefinpfRequest) SetContentInfo(v string) *CallbackPfDefinpfRequest {
	s.ContentInfo = &v
	return s
}

type CallbackPfDefinpfResponse struct {
	// 请求唯一ID，用于链路跟踪和问题排查
	ReqMsgId *string `json:"req_msg_id,omitempty" xml:"req_msg_id,omitempty"`
	// 结果码，一般OK表示调用成功
	ResultCode *string `json:"result_code,omitempty" xml:"result_code,omitempty"`
	// 异常信息的文本描述
	ResultMsg *string `json:"result_msg,omitempty" xml:"result_msg,omitempty"`
	// 是否成功
	Result *bool `json:"result,omitempty" xml:"result,omitempty"`
	// 错误信息
	ErrorMsg *string `json:"error_msg,omitempty" xml:"error_msg,omitempty"`
	// 返回结果
	Response *string `json:"response,omitempty" xml:"response,omitempty"`
}

func (s CallbackPfDefinpfResponse) String() string {
	return tea.Prettify(s)
}

func (s CallbackPfDefinpfResponse) GoString() string {
	return s.String()
}

func (s *CallbackPfDefinpfResponse) SetReqMsgId(v string) *CallbackPfDefinpfResponse {
	s.ReqMsgId = &v
	return s
}

func (s *CallbackPfDefinpfResponse) SetResultCode(v string) *CallbackPfDefinpfResponse {
	s.ResultCode = &v
	return s
}

func (s *CallbackPfDefinpfResponse) SetResultMsg(v string) *CallbackPfDefinpfResponse {
	s.ResultMsg = &v
	return s
}

func (s *CallbackPfDefinpfResponse) SetResult(v bool) *CallbackPfDefinpfResponse {
	s.Result = &v
	return s
}

func (s *CallbackPfDefinpfResponse) SetErrorMsg(v string) *CallbackPfDefinpfResponse {
	s.ErrorMsg = &v
	return s
}

func (s *CallbackPfDefinpfResponse) SetResponse(v string) *CallbackPfDefinpfResponse {
	s.Response = &v
	return s
}

type CreateDidForwarderRequest struct {
	// OAuth模式下的授权token
	AuthToken         *string `json:"auth_token,omitempty" xml:"auth_token,omitempty"`
	ProductInstanceId *string `json:"product_instance_id,omitempty" xml:"product_instance_id,omitempty"`
	// 代理did
	AgentDid *string `json:"agent_did,omitempty" xml:"agent_did,omitempty" require:"true"`
	// 企业名称
	EpCertName *string `json:"ep_cert_name,omitempty" xml:"ep_cert_name,omitempty" require:"true"`
	// 企业证件号
	EpCertNo *string `json:"ep_cert_no,omitempty" xml:"ep_cert_no,omitempty" require:"true"`
	// 扩展字段
	ExtensionInfo *string `json:"extension_info,omitempty" xml:"extension_info,omitempty"`
	// 法人姓名
	LegalPersonCertName *string `json:"legal_person_cert_name,omitempty" xml:"legal_person_cert_name,omitempty"`
	// 法人身份证号
	LegalPersonCertNo *string `json:"legal_person_cert_no,omitempty" xml:"legal_person_cert_no,omitempty"`
}

func (s CreateDidForwarderRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateDidForwarderRequest) GoString() string {
	return s.String()
}

func (s *CreateDidForwarderRequest) SetAuthToken(v string) *CreateDidForwarderRequest {
	s.AuthToken = &v
	return s
}

func (s *CreateDidForwarderRequest) SetProductInstanceId(v string) *CreateDidForwarderRequest {
	s.ProductInstanceId = &v
	return s
}

func (s *CreateDidForwarderRequest) SetAgentDid(v string) *CreateDidForwarderRequest {
	s.AgentDid = &v
	return s
}

func (s *CreateDidForwarderRequest) SetEpCertName(v string) *CreateDidForwarderRequest {
	s.EpCertName = &v
	return s
}

func (s *CreateDidForwarderRequest) SetEpCertNo(v string) *CreateDidForwarderRequest {
	s.EpCertNo = &v
	return s
}

func (s *CreateDidForwarderRequest) SetExtensionInfo(v string) *CreateDidForwarderRequest {
	s.ExtensionInfo = &v
	return s
}

func (s *CreateDidForwarderRequest) SetLegalPersonCertName(v string) *CreateDidForwarderRequest {
	s.LegalPersonCertName = &v
	return s
}

func (s *CreateDidForwarderRequest) SetLegalPersonCertNo(v string) *CreateDidForwarderRequest {
	s.LegalPersonCertNo = &v
	return s
}

type CreateDidForwarderResponse struct {
	// 请求唯一ID，用于链路跟踪和问题排查
	ReqMsgId *string `json:"req_msg_id,omitempty" xml:"req_msg_id,omitempty"`
	// 结果码，一般OK表示调用成功
	ResultCode *string `json:"result_code,omitempty" xml:"result_code,omitempty"`
	// 异常信息的文本描述
	ResultMsg *string `json:"result_msg,omitempty" xml:"result_msg,omitempty"`
	// 货代did
	Did *string `json:"did,omitempty" xml:"did,omitempty"`
}

func (s CreateDidForwarderResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateDidForwarderResponse) GoString() string {
	return s.String()
}

func (s *CreateDidForwarderResponse) SetReqMsgId(v string) *CreateDidForwarderResponse {
	s.ReqMsgId = &v
	return s
}

func (s *CreateDidForwarderResponse) SetResultCode(v string) *CreateDidForwarderResponse {
	s.ResultCode = &v
	return s
}

func (s *CreateDidForwarderResponse) SetResultMsg(v string) *CreateDidForwarderResponse {
	s.ResultMsg = &v
	return s
}

func (s *CreateDidForwarderResponse) SetDid(v string) *CreateDidForwarderResponse {
	s.Did = &v
	return s
}

type CreateDidSaasplatformRequest struct {
	// OAuth模式下的授权token
	AuthToken         *string `json:"auth_token,omitempty" xml:"auth_token,omitempty"`
	ProductInstanceId *string `json:"product_instance_id,omitempty" xml:"product_instance_id,omitempty"`
	// 企业名称
	//
	EpCertName *string `json:"ep_cert_name,omitempty" xml:"ep_cert_name,omitempty" require:"true"`
	// 企业证件号
	EpCertNo *string `json:"ep_cert_no,omitempty" xml:"ep_cert_no,omitempty" require:"true"`
	// 扩展字段
	ExtensionInfo *string `json:"extension_info,omitempty" xml:"extension_info,omitempty"`
	// 法人姓名
	LegalPersonCertName *string `json:"legal_person_cert_name,omitempty" xml:"legal_person_cert_name,omitempty" require:"true"`
	// 法人身份证
	LegalPersonCertNo *string `json:"legal_person_cert_no,omitempty" xml:"legal_person_cert_no,omitempty" require:"true"`
}

func (s CreateDidSaasplatformRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateDidSaasplatformRequest) GoString() string {
	return s.String()
}

func (s *CreateDidSaasplatformRequest) SetAuthToken(v string) *CreateDidSaasplatformRequest {
	s.AuthToken = &v
	return s
}

func (s *CreateDidSaasplatformRequest) SetProductInstanceId(v string) *CreateDidSaasplatformRequest {
	s.ProductInstanceId = &v
	return s
}

func (s *CreateDidSaasplatformRequest) SetEpCertName(v string) *CreateDidSaasplatformRequest {
	s.EpCertName = &v
	return s
}

func (s *CreateDidSaasplatformRequest) SetEpCertNo(v string) *CreateDidSaasplatformRequest {
	s.EpCertNo = &v
	return s
}

func (s *CreateDidSaasplatformRequest) SetExtensionInfo(v string) *CreateDidSaasplatformRequest {
	s.ExtensionInfo = &v
	return s
}

func (s *CreateDidSaasplatformRequest) SetLegalPersonCertName(v string) *CreateDidSaasplatformRequest {
	s.LegalPersonCertName = &v
	return s
}

func (s *CreateDidSaasplatformRequest) SetLegalPersonCertNo(v string) *CreateDidSaasplatformRequest {
	s.LegalPersonCertNo = &v
	return s
}

type CreateDidSaasplatformResponse struct {
	// 请求唯一ID，用于链路跟踪和问题排查
	ReqMsgId *string `json:"req_msg_id,omitempty" xml:"req_msg_id,omitempty"`
	// 结果码，一般OK表示调用成功
	ResultCode *string `json:"result_code,omitempty" xml:"result_code,omitempty"`
	// 异常信息的文本描述
	ResultMsg *string `json:"result_msg,omitempty" xml:"result_msg,omitempty"`
	// saas平台did
	Did *string `json:"did,omitempty" xml:"did,omitempty"`
}

func (s CreateDidSaasplatformResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateDidSaasplatformResponse) GoString() string {
	return s.String()
}

func (s *CreateDidSaasplatformResponse) SetReqMsgId(v string) *CreateDidSaasplatformResponse {
	s.ReqMsgId = &v
	return s
}

func (s *CreateDidSaasplatformResponse) SetResultCode(v string) *CreateDidSaasplatformResponse {
	s.ResultCode = &v
	return s
}

func (s *CreateDidSaasplatformResponse) SetResultMsg(v string) *CreateDidSaasplatformResponse {
	s.ResultMsg = &v
	return s
}

func (s *CreateDidSaasplatformResponse) SetDid(v string) *CreateDidSaasplatformResponse {
	s.Did = &v
	return s
}

type CreateDidClientRequest struct {
	// OAuth模式下的授权token
	AuthToken         *string `json:"auth_token,omitempty" xml:"auth_token,omitempty"`
	ProductInstanceId *string `json:"product_instance_id,omitempty" xml:"product_instance_id,omitempty"`
	// 代理did
	AgentDid *string `json:"agent_did,omitempty" xml:"agent_did,omitempty" require:"true"`
	// 企业名称
	EpCertName *string `json:"ep_cert_name,omitempty" xml:"ep_cert_name,omitempty" require:"true"`
	// 企业证件号
	EpCertNo *string `json:"ep_cert_no,omitempty" xml:"ep_cert_no,omitempty" require:"true"`
	// 扩展字段
	ExtensionInfo *string `json:"extension_info,omitempty" xml:"extension_info,omitempty"`
	// 法人姓名
	LegalPersonCertName *string `json:"legal_person_cert_name,omitempty" xml:"legal_person_cert_name,omitempty"`
	// 法人身份证号
	LegalPersonCertNo *string `json:"legal_person_cert_no,omitempty" xml:"legal_person_cert_no,omitempty"`
}

func (s CreateDidClientRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateDidClientRequest) GoString() string {
	return s.String()
}

func (s *CreateDidClientRequest) SetAuthToken(v string) *CreateDidClientRequest {
	s.AuthToken = &v
	return s
}

func (s *CreateDidClientRequest) SetProductInstanceId(v string) *CreateDidClientRequest {
	s.ProductInstanceId = &v
	return s
}

func (s *CreateDidClientRequest) SetAgentDid(v string) *CreateDidClientRequest {
	s.AgentDid = &v
	return s
}

func (s *CreateDidClientRequest) SetEpCertName(v string) *CreateDidClientRequest {
	s.EpCertName = &v
	return s
}

func (s *CreateDidClientRequest) SetEpCertNo(v string) *CreateDidClientRequest {
	s.EpCertNo = &v
	return s
}

func (s *CreateDidClientRequest) SetExtensionInfo(v string) *CreateDidClientRequest {
	s.ExtensionInfo = &v
	return s
}

func (s *CreateDidClientRequest) SetLegalPersonCertName(v string) *CreateDidClientRequest {
	s.LegalPersonCertName = &v
	return s
}

func (s *CreateDidClientRequest) SetLegalPersonCertNo(v string) *CreateDidClientRequest {
	s.LegalPersonCertNo = &v
	return s
}

type CreateDidClientResponse struct {
	// 请求唯一ID，用于链路跟踪和问题排查
	ReqMsgId *string `json:"req_msg_id,omitempty" xml:"req_msg_id,omitempty"`
	// 结果码，一般OK表示调用成功
	ResultCode *string `json:"result_code,omitempty" xml:"result_code,omitempty"`
	// 异常信息的文本描述
	ResultMsg *string `json:"result_msg,omitempty" xml:"result_msg,omitempty"`
	// 直客did
	Did *string `json:"did,omitempty" xml:"did,omitempty"`
}

func (s CreateDidClientResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateDidClientResponse) GoString() string {
	return s.String()
}

func (s *CreateDidClientResponse) SetReqMsgId(v string) *CreateDidClientResponse {
	s.ReqMsgId = &v
	return s
}

func (s *CreateDidClientResponse) SetResultCode(v string) *CreateDidClientResponse {
	s.ResultCode = &v
	return s
}

func (s *CreateDidClientResponse) SetResultMsg(v string) *CreateDidClientResponse {
	s.ResultMsg = &v
	return s
}

func (s *CreateDidClientResponse) SetDid(v string) *CreateDidClientResponse {
	s.Did = &v
	return s
}

type SaveBizOrderRequest struct {
	// OAuth模式下的授权token
	AuthToken         *string `json:"auth_token,omitempty" xml:"auth_token,omitempty"`
	ProductInstanceId *string `json:"product_instance_id,omitempty" xml:"product_instance_id,omitempty"`
	// 操作动作,为空为新增或更新，UPDATE为更新，DELETE为删除
	Action *string `json:"action,omitempty" xml:"action,omitempty"`
	// 提单要求
	BlRequest *string `json:"bl_request,omitempty" xml:"bl_request,omitempty"`
	// 订舱单号
	BookingNo *string `json:"booking_no,omitempty" xml:"booking_no,omitempty"`
	// 船公司
	Carrier *string `json:"carrier,omitempty" xml:"carrier,omitempty"`
	// 客户did  业务必填
	ClientDid *string `json:"client_did,omitempty" xml:"client_did,omitempty"`
	// 收货人
	Consignee *string `json:"consignee,omitempty" xml:"consignee,omitempty"`
	// 约号
	Contract *string `json:"contract,omitempty" xml:"contract,omitempty"`
	// 报关
	CustomsClearance *string `json:"customs_clearance,omitempty" xml:"customs_clearance,omitempty"`
	// 目的地 业务必填
	//
	DeliveryPlace *string `json:"delivery_place,omitempty" xml:"delivery_place,omitempty"`
	// 运输条款
	DeliveryTerms *string `json:"delivery_terms,omitempty" xml:"delivery_terms,omitempty"`
	// 卸货港. 业务必填
	//
	DischargePort *string `json:"discharge_port,omitempty" xml:"discharge_port,omitempty"`
	// 货代did
	ForwarderDid *string `json:"forwarder_did,omitempty" xml:"forwarder_did,omitempty" require:"true"`
	// 运费
	Freight *string `json:"freight,omitempty" xml:"freight,omitempty"`
	// 保险
	Insurance *string `json:"insurance,omitempty" xml:"insurance,omitempty"`
	// 放单地点
	IssuePlace *string `json:"issue_place,omitempty" xml:"issue_place,omitempty"`
	// 起运港. 业务必填
	//
	LoadingPort *string `json:"loading_port,omitempty" xml:"loading_port,omitempty"`
	// 裝卸方式
	Movement *string `json:"movement,omitempty" xml:"movement,omitempty"`
	// 通知方
	NotifyParty *string `json:"notify_party,omitempty" xml:"notify_party,omitempty"`
	// 订单号
	OrderNo *string `json:"order_no,omitempty" xml:"order_no,omitempty" require:"true"`
	// 付费方式 业务必填
	//
	PaymentTerms *string `json:"payment_terms,omitempty" xml:"payment_terms,omitempty"`
	// 拖车
	PickUp *string `json:"pick_up,omitempty" xml:"pick_up,omitempty"`
	// 收货地点
	ReceiptPlace *string `json:"receipt_place,omitempty" xml:"receipt_place,omitempty"`
	// 备注
	Remark *string `json:"remark,omitempty" xml:"remark,omitempty"`
	// 发货人
	Shipper *string `json:"shipper,omitempty" xml:"shipper,omitempty"`
	// 任务单号
	TaskOrder *string `json:"task_order,omitempty" xml:"task_order,omitempty"`
	// 运输方式
	Transportation *string `json:"transportation,omitempty" xml:"transportation,omitempty"`
	// 船名
	Vessel *string `json:"vessel,omitempty" xml:"vessel,omitempty"`
	// 航次
	Voyage *string `json:"voyage,omitempty" xml:"voyage,omitempty"`
}

func (s SaveBizOrderRequest) String() string {
	return tea.Prettify(s)
}

func (s SaveBizOrderRequest) GoString() string {
	return s.String()
}

func (s *SaveBizOrderRequest) SetAuthToken(v string) *SaveBizOrderRequest {
	s.AuthToken = &v
	return s
}

func (s *SaveBizOrderRequest) SetProductInstanceId(v string) *SaveBizOrderRequest {
	s.ProductInstanceId = &v
	return s
}

func (s *SaveBizOrderRequest) SetAction(v string) *SaveBizOrderRequest {
	s.Action = &v
	return s
}

func (s *SaveBizOrderRequest) SetBlRequest(v string) *SaveBizOrderRequest {
	s.BlRequest = &v
	return s
}

func (s *SaveBizOrderRequest) SetBookingNo(v string) *SaveBizOrderRequest {
	s.BookingNo = &v
	return s
}

func (s *SaveBizOrderRequest) SetCarrier(v string) *SaveBizOrderRequest {
	s.Carrier = &v
	return s
}

func (s *SaveBizOrderRequest) SetClientDid(v string) *SaveBizOrderRequest {
	s.ClientDid = &v
	return s
}

func (s *SaveBizOrderRequest) SetConsignee(v string) *SaveBizOrderRequest {
	s.Consignee = &v
	return s
}

func (s *SaveBizOrderRequest) SetContract(v string) *SaveBizOrderRequest {
	s.Contract = &v
	return s
}

func (s *SaveBizOrderRequest) SetCustomsClearance(v string) *SaveBizOrderRequest {
	s.CustomsClearance = &v
	return s
}

func (s *SaveBizOrderRequest) SetDeliveryPlace(v string) *SaveBizOrderRequest {
	s.DeliveryPlace = &v
	return s
}

func (s *SaveBizOrderRequest) SetDeliveryTerms(v string) *SaveBizOrderRequest {
	s.DeliveryTerms = &v
	return s
}

func (s *SaveBizOrderRequest) SetDischargePort(v string) *SaveBizOrderRequest {
	s.DischargePort = &v
	return s
}

func (s *SaveBizOrderRequest) SetForwarderDid(v string) *SaveBizOrderRequest {
	s.ForwarderDid = &v
	return s
}

func (s *SaveBizOrderRequest) SetFreight(v string) *SaveBizOrderRequest {
	s.Freight = &v
	return s
}

func (s *SaveBizOrderRequest) SetInsurance(v string) *SaveBizOrderRequest {
	s.Insurance = &v
	return s
}

func (s *SaveBizOrderRequest) SetIssuePlace(v string) *SaveBizOrderRequest {
	s.IssuePlace = &v
	return s
}

func (s *SaveBizOrderRequest) SetLoadingPort(v string) *SaveBizOrderRequest {
	s.LoadingPort = &v
	return s
}

func (s *SaveBizOrderRequest) SetMovement(v string) *SaveBizOrderRequest {
	s.Movement = &v
	return s
}

func (s *SaveBizOrderRequest) SetNotifyParty(v string) *SaveBizOrderRequest {
	s.NotifyParty = &v
	return s
}

func (s *SaveBizOrderRequest) SetOrderNo(v string) *SaveBizOrderRequest {
	s.OrderNo = &v
	return s
}

func (s *SaveBizOrderRequest) SetPaymentTerms(v string) *SaveBizOrderRequest {
	s.PaymentTerms = &v
	return s
}

func (s *SaveBizOrderRequest) SetPickUp(v string) *SaveBizOrderRequest {
	s.PickUp = &v
	return s
}

func (s *SaveBizOrderRequest) SetReceiptPlace(v string) *SaveBizOrderRequest {
	s.ReceiptPlace = &v
	return s
}

func (s *SaveBizOrderRequest) SetRemark(v string) *SaveBizOrderRequest {
	s.Remark = &v
	return s
}

func (s *SaveBizOrderRequest) SetShipper(v string) *SaveBizOrderRequest {
	s.Shipper = &v
	return s
}

func (s *SaveBizOrderRequest) SetTaskOrder(v string) *SaveBizOrderRequest {
	s.TaskOrder = &v
	return s
}

func (s *SaveBizOrderRequest) SetTransportation(v string) *SaveBizOrderRequest {
	s.Transportation = &v
	return s
}

func (s *SaveBizOrderRequest) SetVessel(v string) *SaveBizOrderRequest {
	s.Vessel = &v
	return s
}

func (s *SaveBizOrderRequest) SetVoyage(v string) *SaveBizOrderRequest {
	s.Voyage = &v
	return s
}

type SaveBizOrderResponse struct {
	// 请求唯一ID，用于链路跟踪和问题排查
	ReqMsgId *string `json:"req_msg_id,omitempty" xml:"req_msg_id,omitempty"`
	// 结果码，一般OK表示调用成功
	ResultCode *string `json:"result_code,omitempty" xml:"result_code,omitempty"`
	// 异常信息的文本描述
	ResultMsg *string `json:"result_msg,omitempty" xml:"result_msg,omitempty"`
	// 链上凭证
	TxCodes []*TxDto `json:"tx_codes,omitempty" xml:"tx_codes,omitempty" type:"Repeated"`
}

func (s SaveBizOrderResponse) String() string {
	return tea.Prettify(s)
}

func (s SaveBizOrderResponse) GoString() string {
	return s.String()
}

func (s *SaveBizOrderResponse) SetReqMsgId(v string) *SaveBizOrderResponse {
	s.ReqMsgId = &v
	return s
}

func (s *SaveBizOrderResponse) SetResultCode(v string) *SaveBizOrderResponse {
	s.ResultCode = &v
	return s
}

func (s *SaveBizOrderResponse) SetResultMsg(v string) *SaveBizOrderResponse {
	s.ResultMsg = &v
	return s
}

func (s *SaveBizOrderResponse) SetTxCodes(v []*TxDto) *SaveBizOrderResponse {
	s.TxCodes = v
	return s
}

type SaveBizConsignorderRequest struct {
	// OAuth模式下的授权token
	AuthToken         *string `json:"auth_token,omitempty" xml:"auth_token,omitempty"`
	ProductInstanceId *string `json:"product_instance_id,omitempty" xml:"product_instance_id,omitempty"`
	// 操作动作,为空为新增或更新，UPDATE为更新，DELETE为删除
	Action *string `json:"action,omitempty" xml:"action,omitempty"`
	// 托单code
	ConsignOrderCode *string `json:"consign_order_code,omitempty" xml:"consign_order_code,omitempty" require:"true"`
	// 托单文件hash  业务必填
	FileHash *string `json:"file_hash,omitempty" xml:"file_hash,omitempty"`
	// 托单文件id  业务必填
	FileId *string `json:"file_id,omitempty" xml:"file_id,omitempty"`
	// 货代did
	ForwarderDid *string `json:"forwarder_did,omitempty" xml:"forwarder_did,omitempty" require:"true"`
	// 订单号
	OrderNo *string `json:"order_no,omitempty" xml:"order_no,omitempty" require:"true"`
}

func (s SaveBizConsignorderRequest) String() string {
	return tea.Prettify(s)
}

func (s SaveBizConsignorderRequest) GoString() string {
	return s.String()
}

func (s *SaveBizConsignorderRequest) SetAuthToken(v string) *SaveBizConsignorderRequest {
	s.AuthToken = &v
	return s
}

func (s *SaveBizConsignorderRequest) SetProductInstanceId(v string) *SaveBizConsignorderRequest {
	s.ProductInstanceId = &v
	return s
}

func (s *SaveBizConsignorderRequest) SetAction(v string) *SaveBizConsignorderRequest {
	s.Action = &v
	return s
}

func (s *SaveBizConsignorderRequest) SetConsignOrderCode(v string) *SaveBizConsignorderRequest {
	s.ConsignOrderCode = &v
	return s
}

func (s *SaveBizConsignorderRequest) SetFileHash(v string) *SaveBizConsignorderRequest {
	s.FileHash = &v
	return s
}

func (s *SaveBizConsignorderRequest) SetFileId(v string) *SaveBizConsignorderRequest {
	s.FileId = &v
	return s
}

func (s *SaveBizConsignorderRequest) SetForwarderDid(v string) *SaveBizConsignorderRequest {
	s.ForwarderDid = &v
	return s
}

func (s *SaveBizConsignorderRequest) SetOrderNo(v string) *SaveBizConsignorderRequest {
	s.OrderNo = &v
	return s
}

type SaveBizConsignorderResponse struct {
	// 请求唯一ID，用于链路跟踪和问题排查
	ReqMsgId *string `json:"req_msg_id,omitempty" xml:"req_msg_id,omitempty"`
	// 结果码，一般OK表示调用成功
	ResultCode *string `json:"result_code,omitempty" xml:"result_code,omitempty"`
	// 异常信息的文本描述
	ResultMsg *string `json:"result_msg,omitempty" xml:"result_msg,omitempty"`
	// 链上凭证
	TxCodes []*TxDto `json:"tx_codes,omitempty" xml:"tx_codes,omitempty" type:"Repeated"`
}

func (s SaveBizConsignorderResponse) String() string {
	return tea.Prettify(s)
}

func (s SaveBizConsignorderResponse) GoString() string {
	return s.String()
}

func (s *SaveBizConsignorderResponse) SetReqMsgId(v string) *SaveBizConsignorderResponse {
	s.ReqMsgId = &v
	return s
}

func (s *SaveBizConsignorderResponse) SetResultCode(v string) *SaveBizConsignorderResponse {
	s.ResultCode = &v
	return s
}

func (s *SaveBizConsignorderResponse) SetResultMsg(v string) *SaveBizConsignorderResponse {
	s.ResultMsg = &v
	return s
}

func (s *SaveBizConsignorderResponse) SetTxCodes(v []*TxDto) *SaveBizConsignorderResponse {
	s.TxCodes = v
	return s
}

type SaveBizGoodsRequest struct {
	// OAuth模式下的授权token
	AuthToken         *string `json:"auth_token,omitempty" xml:"auth_token,omitempty"`
	ProductInstanceId *string `json:"product_instance_id,omitempty" xml:"product_instance_id,omitempty"`
	// 操作动作,为空为新增或更新，UPDATE为更新，DELETE为删除
	Action *string `json:"action,omitempty" xml:"action,omitempty"`
	// 预计备货时间 毫秒值单位
	CargoReadyDate *int64 `json:"cargo_ready_date,omitempty" xml:"cargo_ready_date,omitempty"`
	// 危险品页号
	DgPageNo *string `json:"dg_page_no,omitempty" xml:"dg_page_no,omitempty"`
	// 危险品级别
	DgType *string `json:"dg_type,omitempty" xml:"dg_type,omitempty"`
	// 危险品闪点
	FlashPoint *string `json:"flash_point,omitempty" xml:"flash_point,omitempty"`
	// 货代did
	ForwarderDid *string `json:"forwarder_did,omitempty" xml:"forwarder_did,omitempty" require:"true"`
	// 货物 业务必填
	//
	Goods *string `json:"goods,omitempty" xml:"goods,omitempty"`
	// 中文品名
	GoodsCn *string `json:"goods_cn,omitempty" xml:"goods_cn,omitempty"`
	// 货物ID
	GoodsId *string `json:"goods_id,omitempty" xml:"goods_id,omitempty" require:"true"`
	// 货物类型（业务必填）
	GoodsType *string `json:"goods_type,omitempty" xml:"goods_type,omitempty"`
	// HS CODE
	HsCodes []*string `json:"hs_codes,omitempty" xml:"hs_codes,omitempty" type:"Repeated"`
	// 唛头
	Marks *string `json:"marks,omitempty" xml:"marks,omitempty"`
	// 委托件数 业务必填
	//
	Number *string `json:"number,omitempty" xml:"number,omitempty"`
	// 订单号
	OrderNo *string `json:"order_no,omitempty" xml:"order_no,omitempty" require:"true"`
	// 包装类型
	PackageType *string `json:"package_type,omitempty" xml:"package_type,omitempty"`
	// 实际件数
	RealNumber *string `json:"real_number,omitempty" xml:"real_number,omitempty"`
	// 实际体积
	RealVolume *string `json:"real_volume,omitempty" xml:"real_volume,omitempty"`
	// 实际重量
	RealWeight *string `json:"real_weight,omitempty" xml:"real_weight,omitempty"`
	// 危险品联合国编号
	UnNo *string `json:"un_no,omitempty" xml:"un_no,omitempty"`
	// 委托体积（业务必填）
	Volume *string `json:"volume,omitempty" xml:"volume,omitempty"`
	// 委托重量（业务必填）
	Weight *string `json:"weight,omitempty" xml:"weight,omitempty"`
}

func (s SaveBizGoodsRequest) String() string {
	return tea.Prettify(s)
}

func (s SaveBizGoodsRequest) GoString() string {
	return s.String()
}

func (s *SaveBizGoodsRequest) SetAuthToken(v string) *SaveBizGoodsRequest {
	s.AuthToken = &v
	return s
}

func (s *SaveBizGoodsRequest) SetProductInstanceId(v string) *SaveBizGoodsRequest {
	s.ProductInstanceId = &v
	return s
}

func (s *SaveBizGoodsRequest) SetAction(v string) *SaveBizGoodsRequest {
	s.Action = &v
	return s
}

func (s *SaveBizGoodsRequest) SetCargoReadyDate(v int64) *SaveBizGoodsRequest {
	s.CargoReadyDate = &v
	return s
}

func (s *SaveBizGoodsRequest) SetDgPageNo(v string) *SaveBizGoodsRequest {
	s.DgPageNo = &v
	return s
}

func (s *SaveBizGoodsRequest) SetDgType(v string) *SaveBizGoodsRequest {
	s.DgType = &v
	return s
}

func (s *SaveBizGoodsRequest) SetFlashPoint(v string) *SaveBizGoodsRequest {
	s.FlashPoint = &v
	return s
}

func (s *SaveBizGoodsRequest) SetForwarderDid(v string) *SaveBizGoodsRequest {
	s.ForwarderDid = &v
	return s
}

func (s *SaveBizGoodsRequest) SetGoods(v string) *SaveBizGoodsRequest {
	s.Goods = &v
	return s
}

func (s *SaveBizGoodsRequest) SetGoodsCn(v string) *SaveBizGoodsRequest {
	s.GoodsCn = &v
	return s
}

func (s *SaveBizGoodsRequest) SetGoodsId(v string) *SaveBizGoodsRequest {
	s.GoodsId = &v
	return s
}

func (s *SaveBizGoodsRequest) SetGoodsType(v string) *SaveBizGoodsRequest {
	s.GoodsType = &v
	return s
}

func (s *SaveBizGoodsRequest) SetHsCodes(v []*string) *SaveBizGoodsRequest {
	s.HsCodes = v
	return s
}

func (s *SaveBizGoodsRequest) SetMarks(v string) *SaveBizGoodsRequest {
	s.Marks = &v
	return s
}

func (s *SaveBizGoodsRequest) SetNumber(v string) *SaveBizGoodsRequest {
	s.Number = &v
	return s
}

func (s *SaveBizGoodsRequest) SetOrderNo(v string) *SaveBizGoodsRequest {
	s.OrderNo = &v
	return s
}

func (s *SaveBizGoodsRequest) SetPackageType(v string) *SaveBizGoodsRequest {
	s.PackageType = &v
	return s
}

func (s *SaveBizGoodsRequest) SetRealNumber(v string) *SaveBizGoodsRequest {
	s.RealNumber = &v
	return s
}

func (s *SaveBizGoodsRequest) SetRealVolume(v string) *SaveBizGoodsRequest {
	s.RealVolume = &v
	return s
}

func (s *SaveBizGoodsRequest) SetRealWeight(v string) *SaveBizGoodsRequest {
	s.RealWeight = &v
	return s
}

func (s *SaveBizGoodsRequest) SetUnNo(v string) *SaveBizGoodsRequest {
	s.UnNo = &v
	return s
}

func (s *SaveBizGoodsRequest) SetVolume(v string) *SaveBizGoodsRequest {
	s.Volume = &v
	return s
}

func (s *SaveBizGoodsRequest) SetWeight(v string) *SaveBizGoodsRequest {
	s.Weight = &v
	return s
}

type SaveBizGoodsResponse struct {
	// 请求唯一ID，用于链路跟踪和问题排查
	ReqMsgId *string `json:"req_msg_id,omitempty" xml:"req_msg_id,omitempty"`
	// 结果码，一般OK表示调用成功
	ResultCode *string `json:"result_code,omitempty" xml:"result_code,omitempty"`
	// 异常信息的文本描述
	ResultMsg *string `json:"result_msg,omitempty" xml:"result_msg,omitempty"`
	// 链上凭证
	TxCodes []*TxDto `json:"tx_codes,omitempty" xml:"tx_codes,omitempty" type:"Repeated"`
}

func (s SaveBizGoodsResponse) String() string {
	return tea.Prettify(s)
}

func (s SaveBizGoodsResponse) GoString() string {
	return s.String()
}

func (s *SaveBizGoodsResponse) SetReqMsgId(v string) *SaveBizGoodsResponse {
	s.ReqMsgId = &v
	return s
}

func (s *SaveBizGoodsResponse) SetResultCode(v string) *SaveBizGoodsResponse {
	s.ResultCode = &v
	return s
}

func (s *SaveBizGoodsResponse) SetResultMsg(v string) *SaveBizGoodsResponse {
	s.ResultMsg = &v
	return s
}

func (s *SaveBizGoodsResponse) SetTxCodes(v []*TxDto) *SaveBizGoodsResponse {
	s.TxCodes = v
	return s
}

type SaveBizSonotifyRequest struct {
	// OAuth模式下的授权token
	AuthToken         *string `json:"auth_token,omitempty" xml:"auth_token,omitempty"`
	ProductInstanceId *string `json:"product_instance_id,omitempty" xml:"product_instance_id,omitempty"`
	// 操作动作,为空为新增或更新，UPDATE为更新，DELETE为删除
	Action *string `json:"action,omitempty" xml:"action,omitempty"`
	// so通知关联的订舱单 (业务必填)
	BookingParams []*SoNotifyBookingParam `json:"booking_params,omitempty" xml:"booking_params,omitempty" type:"Repeated"`
	// 联系人
	ContactName *string `json:"contact_name,omitempty" xml:"contact_name,omitempty"`
	// 联系方式
	ContactType *string `json:"contact_type,omitempty" xml:"contact_type,omitempty"`
	// 卸货港 业务必填
	//
	DischargePort *string `json:"discharge_port,omitempty" xml:"discharge_port,omitempty"`
	// 货代did
	ForwarderDid *string `json:"forwarder_did,omitempty" xml:"forwarder_did,omitempty" require:"true"`
	// 起运港 业务必填
	//
	LoadingPort *string `json:"loading_port,omitempty" xml:"loading_port,omitempty"`
	//
	// 订单号
	OrderNo *string `json:"order_no,omitempty" xml:"order_no,omitempty" require:"true"`
	// soNotify 唯一标识code
	SoNotifyCode *string `json:"so_notify_code,omitempty" xml:"so_notify_code,omitempty" require:"true"`
}

func (s SaveBizSonotifyRequest) String() string {
	return tea.Prettify(s)
}

func (s SaveBizSonotifyRequest) GoString() string {
	return s.String()
}

func (s *SaveBizSonotifyRequest) SetAuthToken(v string) *SaveBizSonotifyRequest {
	s.AuthToken = &v
	return s
}

func (s *SaveBizSonotifyRequest) SetProductInstanceId(v string) *SaveBizSonotifyRequest {
	s.ProductInstanceId = &v
	return s
}

func (s *SaveBizSonotifyRequest) SetAction(v string) *SaveBizSonotifyRequest {
	s.Action = &v
	return s
}

func (s *SaveBizSonotifyRequest) SetBookingParams(v []*SoNotifyBookingParam) *SaveBizSonotifyRequest {
	s.BookingParams = v
	return s
}

func (s *SaveBizSonotifyRequest) SetContactName(v string) *SaveBizSonotifyRequest {
	s.ContactName = &v
	return s
}

func (s *SaveBizSonotifyRequest) SetContactType(v string) *SaveBizSonotifyRequest {
	s.ContactType = &v
	return s
}

func (s *SaveBizSonotifyRequest) SetDischargePort(v string) *SaveBizSonotifyRequest {
	s.DischargePort = &v
	return s
}

func (s *SaveBizSonotifyRequest) SetForwarderDid(v string) *SaveBizSonotifyRequest {
	s.ForwarderDid = &v
	return s
}

func (s *SaveBizSonotifyRequest) SetLoadingPort(v string) *SaveBizSonotifyRequest {
	s.LoadingPort = &v
	return s
}

func (s *SaveBizSonotifyRequest) SetOrderNo(v string) *SaveBizSonotifyRequest {
	s.OrderNo = &v
	return s
}

func (s *SaveBizSonotifyRequest) SetSoNotifyCode(v string) *SaveBizSonotifyRequest {
	s.SoNotifyCode = &v
	return s
}

type SaveBizSonotifyResponse struct {
	// 请求唯一ID，用于链路跟踪和问题排查
	ReqMsgId *string `json:"req_msg_id,omitempty" xml:"req_msg_id,omitempty"`
	// 结果码，一般OK表示调用成功
	ResultCode *string `json:"result_code,omitempty" xml:"result_code,omitempty"`
	// 异常信息的文本描述
	ResultMsg *string `json:"result_msg,omitempty" xml:"result_msg,omitempty"`
	// 链上凭证
	TxCodes []*TxDto `json:"tx_codes,omitempty" xml:"tx_codes,omitempty" type:"Repeated"`
}

func (s SaveBizSonotifyResponse) String() string {
	return tea.Prettify(s)
}

func (s SaveBizSonotifyResponse) GoString() string {
	return s.String()
}

func (s *SaveBizSonotifyResponse) SetReqMsgId(v string) *SaveBizSonotifyResponse {
	s.ReqMsgId = &v
	return s
}

func (s *SaveBizSonotifyResponse) SetResultCode(v string) *SaveBizSonotifyResponse {
	s.ResultCode = &v
	return s
}

func (s *SaveBizSonotifyResponse) SetResultMsg(v string) *SaveBizSonotifyResponse {
	s.ResultMsg = &v
	return s
}

func (s *SaveBizSonotifyResponse) SetTxCodes(v []*TxDto) *SaveBizSonotifyResponse {
	s.TxCodes = v
	return s
}

type SaveBizBookingorderRequest struct {
	// OAuth模式下的授权token
	AuthToken         *string `json:"auth_token,omitempty" xml:"auth_token,omitempty"`
	ProductInstanceId *string `json:"product_instance_id,omitempty" xml:"product_instance_id,omitempty"`
	// 操作动作,为空为新增或更新，UPDATE为更新，DELETE为删除
	Action *string `json:"action,omitempty" xml:"action,omitempty"`
	// 订舱单号
	BookingNo *string `json:"booking_no,omitempty" xml:"booking_no,omitempty" require:"true"`
	// 船公司 业务必填
	Carrier *string `json:"carrier,omitempty" xml:"carrier,omitempty"`
	// 截关时间
	CustomsClearance *int64 `json:"customs_clearance,omitempty" xml:"customs_clearance,omitempty"`
	// 场站 业务必填
	Cy *string `json:"cy,omitempty" xml:"cy,omitempty"`
	// 截港时间 毫秒值单位
	CyClosing *int64 `json:"cy_closing,omitempty" xml:"cy_closing,omitempty"`
	// 目的地
	DeliveryPlace *string `json:"delivery_place,omitempty" xml:"delivery_place,omitempty"`
	// 卸货港
	DischargePort *string `json:"discharge_port,omitempty" xml:"discharge_port,omitempty"`
	// 预计船期 毫秒值单位
	Etd *int64 `json:"etd,omitempty" xml:"etd,omitempty"`
	// 货代did
	ForwarderDid *string `json:"forwarder_did,omitempty" xml:"forwarder_did,omitempty" require:"true"`
	// 起运港
	LoadingPort *string `json:"loading_port,omitempty" xml:"loading_port,omitempty"`
	// 订单号
	OrderNo *string `json:"order_no,omitempty" xml:"order_no,omitempty" require:"true"`
	// 截单时间  毫秒值单位
	SiClosing *int64 `json:"si_closing,omitempty" xml:"si_closing,omitempty"`
	// 特殊字段无要求非必填
	Source *string `json:"source,omitempty" xml:"source,omitempty"`
	// 船名 业务必填
	Vessel *string `json:"vessel,omitempty" xml:"vessel,omitempty"`
	// 航次 业务必填
	Voyage *string `json:"voyage,omitempty" xml:"voyage,omitempty"`
	// 订舱号
	BkgNo *string `json:"bkg_no,omitempty" xml:"bkg_no,omitempty" require:"true"`
	// 确认时间
	ConfirmTime *int64 `json:"confirm_time,omitempty" xml:"confirm_time,omitempty"`
}

func (s SaveBizBookingorderRequest) String() string {
	return tea.Prettify(s)
}

func (s SaveBizBookingorderRequest) GoString() string {
	return s.String()
}

func (s *SaveBizBookingorderRequest) SetAuthToken(v string) *SaveBizBookingorderRequest {
	s.AuthToken = &v
	return s
}

func (s *SaveBizBookingorderRequest) SetProductInstanceId(v string) *SaveBizBookingorderRequest {
	s.ProductInstanceId = &v
	return s
}

func (s *SaveBizBookingorderRequest) SetAction(v string) *SaveBizBookingorderRequest {
	s.Action = &v
	return s
}

func (s *SaveBizBookingorderRequest) SetBookingNo(v string) *SaveBizBookingorderRequest {
	s.BookingNo = &v
	return s
}

func (s *SaveBizBookingorderRequest) SetCarrier(v string) *SaveBizBookingorderRequest {
	s.Carrier = &v
	return s
}

func (s *SaveBizBookingorderRequest) SetCustomsClearance(v int64) *SaveBizBookingorderRequest {
	s.CustomsClearance = &v
	return s
}

func (s *SaveBizBookingorderRequest) SetCy(v string) *SaveBizBookingorderRequest {
	s.Cy = &v
	return s
}

func (s *SaveBizBookingorderRequest) SetCyClosing(v int64) *SaveBizBookingorderRequest {
	s.CyClosing = &v
	return s
}

func (s *SaveBizBookingorderRequest) SetDeliveryPlace(v string) *SaveBizBookingorderRequest {
	s.DeliveryPlace = &v
	return s
}

func (s *SaveBizBookingorderRequest) SetDischargePort(v string) *SaveBizBookingorderRequest {
	s.DischargePort = &v
	return s
}

func (s *SaveBizBookingorderRequest) SetEtd(v int64) *SaveBizBookingorderRequest {
	s.Etd = &v
	return s
}

func (s *SaveBizBookingorderRequest) SetForwarderDid(v string) *SaveBizBookingorderRequest {
	s.ForwarderDid = &v
	return s
}

func (s *SaveBizBookingorderRequest) SetLoadingPort(v string) *SaveBizBookingorderRequest {
	s.LoadingPort = &v
	return s
}

func (s *SaveBizBookingorderRequest) SetOrderNo(v string) *SaveBizBookingorderRequest {
	s.OrderNo = &v
	return s
}

func (s *SaveBizBookingorderRequest) SetSiClosing(v int64) *SaveBizBookingorderRequest {
	s.SiClosing = &v
	return s
}

func (s *SaveBizBookingorderRequest) SetSource(v string) *SaveBizBookingorderRequest {
	s.Source = &v
	return s
}

func (s *SaveBizBookingorderRequest) SetVessel(v string) *SaveBizBookingorderRequest {
	s.Vessel = &v
	return s
}

func (s *SaveBizBookingorderRequest) SetVoyage(v string) *SaveBizBookingorderRequest {
	s.Voyage = &v
	return s
}

func (s *SaveBizBookingorderRequest) SetBkgNo(v string) *SaveBizBookingorderRequest {
	s.BkgNo = &v
	return s
}

func (s *SaveBizBookingorderRequest) SetConfirmTime(v int64) *SaveBizBookingorderRequest {
	s.ConfirmTime = &v
	return s
}

type SaveBizBookingorderResponse struct {
	// 请求唯一ID，用于链路跟踪和问题排查
	ReqMsgId *string `json:"req_msg_id,omitempty" xml:"req_msg_id,omitempty"`
	// 结果码，一般OK表示调用成功
	ResultCode *string `json:"result_code,omitempty" xml:"result_code,omitempty"`
	// 异常信息的文本描述
	ResultMsg *string `json:"result_msg,omitempty" xml:"result_msg,omitempty"`
	// 链上凭证
	TxCodes []*TxDto `json:"tx_codes,omitempty" xml:"tx_codes,omitempty" type:"Repeated"`
}

func (s SaveBizBookingorderResponse) String() string {
	return tea.Prettify(s)
}

func (s SaveBizBookingorderResponse) GoString() string {
	return s.String()
}

func (s *SaveBizBookingorderResponse) SetReqMsgId(v string) *SaveBizBookingorderResponse {
	s.ReqMsgId = &v
	return s
}

func (s *SaveBizBookingorderResponse) SetResultCode(v string) *SaveBizBookingorderResponse {
	s.ResultCode = &v
	return s
}

func (s *SaveBizBookingorderResponse) SetResultMsg(v string) *SaveBizBookingorderResponse {
	s.ResultMsg = &v
	return s
}

func (s *SaveBizBookingorderResponse) SetTxCodes(v []*TxDto) *SaveBizBookingorderResponse {
	s.TxCodes = v
	return s
}

type SaveBizContainerRequest struct {
	// OAuth模式下的授权token
	AuthToken         *string `json:"auth_token,omitempty" xml:"auth_token,omitempty"`
	ProductInstanceId *string `json:"product_instance_id,omitempty" xml:"product_instance_id,omitempty"`
	// 操作动作,为空为新增或更新，UPDATE为更新，DELETE为删除
	Action *string `json:"action,omitempty" xml:"action,omitempty"`
	// 订舱单号
	//
	BookingNo *string `json:"booking_no,omitempty" xml:"booking_no,omitempty" require:"true"`
	// 集装箱ID
	ContainerId *string `json:"container_id,omitempty" xml:"container_id,omitempty" require:"true"`
	// 箱号 业务必填
	ContainerNo *string `json:"container_no,omitempty" xml:"container_no,omitempty"`
	//  箱型  业务必填
	ContainerType *string `json:"container_type,omitempty" xml:"container_type,omitempty"`
	// 货代did
	ForwarderDid *string `json:"forwarder_did,omitempty" xml:"forwarder_did,omitempty" require:"true"`
	// 货物列表
	GoodsList []*ContainerGoodsParam `json:"goods_list,omitempty" xml:"goods_list,omitempty" type:"Repeated"`
	// 订单号
	OrderNo *string `json:"order_no,omitempty" xml:"order_no,omitempty" require:"true"`
	// 备注
	Remark *string `json:"remark,omitempty" xml:"remark,omitempty"`
	// 封铅号
	SealNo *string `json:"seal_no,omitempty" xml:"seal_no,omitempty"`
	// 是否SOC
	SocFlag *string `json:"soc_flag,omitempty" xml:"soc_flag,omitempty"`
	// 特殊字段无要求非必填
	Source *string `json:"source,omitempty" xml:"source,omitempty"`
}

func (s SaveBizContainerRequest) String() string {
	return tea.Prettify(s)
}

func (s SaveBizContainerRequest) GoString() string {
	return s.String()
}

func (s *SaveBizContainerRequest) SetAuthToken(v string) *SaveBizContainerRequest {
	s.AuthToken = &v
	return s
}

func (s *SaveBizContainerRequest) SetProductInstanceId(v string) *SaveBizContainerRequest {
	s.ProductInstanceId = &v
	return s
}

func (s *SaveBizContainerRequest) SetAction(v string) *SaveBizContainerRequest {
	s.Action = &v
	return s
}

func (s *SaveBizContainerRequest) SetBookingNo(v string) *SaveBizContainerRequest {
	s.BookingNo = &v
	return s
}

func (s *SaveBizContainerRequest) SetContainerId(v string) *SaveBizContainerRequest {
	s.ContainerId = &v
	return s
}

func (s *SaveBizContainerRequest) SetContainerNo(v string) *SaveBizContainerRequest {
	s.ContainerNo = &v
	return s
}

func (s *SaveBizContainerRequest) SetContainerType(v string) *SaveBizContainerRequest {
	s.ContainerType = &v
	return s
}

func (s *SaveBizContainerRequest) SetForwarderDid(v string) *SaveBizContainerRequest {
	s.ForwarderDid = &v
	return s
}

func (s *SaveBizContainerRequest) SetGoodsList(v []*ContainerGoodsParam) *SaveBizContainerRequest {
	s.GoodsList = v
	return s
}

func (s *SaveBizContainerRequest) SetOrderNo(v string) *SaveBizContainerRequest {
	s.OrderNo = &v
	return s
}

func (s *SaveBizContainerRequest) SetRemark(v string) *SaveBizContainerRequest {
	s.Remark = &v
	return s
}

func (s *SaveBizContainerRequest) SetSealNo(v string) *SaveBizContainerRequest {
	s.SealNo = &v
	return s
}

func (s *SaveBizContainerRequest) SetSocFlag(v string) *SaveBizContainerRequest {
	s.SocFlag = &v
	return s
}

func (s *SaveBizContainerRequest) SetSource(v string) *SaveBizContainerRequest {
	s.Source = &v
	return s
}

type SaveBizContainerResponse struct {
	// 请求唯一ID，用于链路跟踪和问题排查
	ReqMsgId *string `json:"req_msg_id,omitempty" xml:"req_msg_id,omitempty"`
	// 结果码，一般OK表示调用成功
	ResultCode *string `json:"result_code,omitempty" xml:"result_code,omitempty"`
	// 异常信息的文本描述
	ResultMsg *string `json:"result_msg,omitempty" xml:"result_msg,omitempty"`
	// 链上凭证
	TxCodes []*TxDto `json:"tx_codes,omitempty" xml:"tx_codes,omitempty" type:"Repeated"`
}

func (s SaveBizContainerResponse) String() string {
	return tea.Prettify(s)
}

func (s SaveBizContainerResponse) GoString() string {
	return s.String()
}

func (s *SaveBizContainerResponse) SetReqMsgId(v string) *SaveBizContainerResponse {
	s.ReqMsgId = &v
	return s
}

func (s *SaveBizContainerResponse) SetResultCode(v string) *SaveBizContainerResponse {
	s.ResultCode = &v
	return s
}

func (s *SaveBizContainerResponse) SetResultMsg(v string) *SaveBizContainerResponse {
	s.ResultMsg = &v
	return s
}

func (s *SaveBizContainerResponse) SetTxCodes(v []*TxDto) *SaveBizContainerResponse {
	s.TxCodes = v
	return s
}

type SaveBizCustomsorderRequest struct {
	// OAuth模式下的授权token
	AuthToken         *string `json:"auth_token,omitempty" xml:"auth_token,omitempty"`
	ProductInstanceId *string `json:"product_instance_id,omitempty" xml:"product_instance_id,omitempty"`
	// 操作动作,为空为新增或更新，UPDATE为更新，DELETE为删除
	Action *string `json:"action,omitempty" xml:"action,omitempty"`
	// 订舱单号
	BookingParams []*CustomsOrderBookingParam `json:"booking_params,omitempty" xml:"booking_params,omitempty" type:"Repeated"`
	// 报关代理
	Broker *string `json:"broker,omitempty" xml:"broker,omitempty"`
	// 集装箱ID
	ContainerId *string `json:"container_id,omitempty" xml:"container_id,omitempty"`
	// 箱号
	ContainerNo *string `json:"container_no,omitempty" xml:"container_no,omitempty"`
	// 报关单号
	CustomsCode *string `json:"customs_code,omitempty" xml:"customs_code,omitempty" require:"true"`
	//  出口人
	Exporter *string `json:"exporter,omitempty" xml:"exporter,omitempty"`
	// 货代did
	//
	ForwarderDid *string `json:"forwarder_did,omitempty" xml:"forwarder_did,omitempty" require:"true"`
	// 货物名称
	Goods *string `json:"goods,omitempty" xml:"goods,omitempty"`
	// 毛重
	GrossWeight *string `json:"gross_weight,omitempty" xml:"gross_weight,omitempty"`
	// 订单号
	//
	OrderNo *string `json:"order_no,omitempty" xml:"order_no,omitempty" require:"true"`
	// 件数
	PackagesNo *string `json:"packages_no,omitempty" xml:"packages_no,omitempty"`
	// 报关状态  APPROVED--通关，UNAPPROVED-未通关
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// 航名 业务必填
	Vessel *string `json:"vessel,omitempty" xml:"vessel,omitempty" require:"true"`
	// 航次 业务必填
	Voyage *string `json:"voyage,omitempty" xml:"voyage,omitempty"`
}

func (s SaveBizCustomsorderRequest) String() string {
	return tea.Prettify(s)
}

func (s SaveBizCustomsorderRequest) GoString() string {
	return s.String()
}

func (s *SaveBizCustomsorderRequest) SetAuthToken(v string) *SaveBizCustomsorderRequest {
	s.AuthToken = &v
	return s
}

func (s *SaveBizCustomsorderRequest) SetProductInstanceId(v string) *SaveBizCustomsorderRequest {
	s.ProductInstanceId = &v
	return s
}

func (s *SaveBizCustomsorderRequest) SetAction(v string) *SaveBizCustomsorderRequest {
	s.Action = &v
	return s
}

func (s *SaveBizCustomsorderRequest) SetBookingParams(v []*CustomsOrderBookingParam) *SaveBizCustomsorderRequest {
	s.BookingParams = v
	return s
}

func (s *SaveBizCustomsorderRequest) SetBroker(v string) *SaveBizCustomsorderRequest {
	s.Broker = &v
	return s
}

func (s *SaveBizCustomsorderRequest) SetContainerId(v string) *SaveBizCustomsorderRequest {
	s.ContainerId = &v
	return s
}

func (s *SaveBizCustomsorderRequest) SetContainerNo(v string) *SaveBizCustomsorderRequest {
	s.ContainerNo = &v
	return s
}

func (s *SaveBizCustomsorderRequest) SetCustomsCode(v string) *SaveBizCustomsorderRequest {
	s.CustomsCode = &v
	return s
}

func (s *SaveBizCustomsorderRequest) SetExporter(v string) *SaveBizCustomsorderRequest {
	s.Exporter = &v
	return s
}

func (s *SaveBizCustomsorderRequest) SetForwarderDid(v string) *SaveBizCustomsorderRequest {
	s.ForwarderDid = &v
	return s
}

func (s *SaveBizCustomsorderRequest) SetGoods(v string) *SaveBizCustomsorderRequest {
	s.Goods = &v
	return s
}

func (s *SaveBizCustomsorderRequest) SetGrossWeight(v string) *SaveBizCustomsorderRequest {
	s.GrossWeight = &v
	return s
}

func (s *SaveBizCustomsorderRequest) SetOrderNo(v string) *SaveBizCustomsorderRequest {
	s.OrderNo = &v
	return s
}

func (s *SaveBizCustomsorderRequest) SetPackagesNo(v string) *SaveBizCustomsorderRequest {
	s.PackagesNo = &v
	return s
}

func (s *SaveBizCustomsorderRequest) SetStatus(v string) *SaveBizCustomsorderRequest {
	s.Status = &v
	return s
}

func (s *SaveBizCustomsorderRequest) SetVessel(v string) *SaveBizCustomsorderRequest {
	s.Vessel = &v
	return s
}

func (s *SaveBizCustomsorderRequest) SetVoyage(v string) *SaveBizCustomsorderRequest {
	s.Voyage = &v
	return s
}

type SaveBizCustomsorderResponse struct {
	// 请求唯一ID，用于链路跟踪和问题排查
	ReqMsgId *string `json:"req_msg_id,omitempty" xml:"req_msg_id,omitempty"`
	// 结果码，一般OK表示调用成功
	ResultCode *string `json:"result_code,omitempty" xml:"result_code,omitempty"`
	// 异常信息的文本描述
	ResultMsg *string `json:"result_msg,omitempty" xml:"result_msg,omitempty"`
	// 链上凭证
	TxCodes []*TxDto `json:"tx_codes,omitempty" xml:"tx_codes,omitempty" type:"Repeated"`
}

func (s SaveBizCustomsorderResponse) String() string {
	return tea.Prettify(s)
}

func (s SaveBizCustomsorderResponse) GoString() string {
	return s.String()
}

func (s *SaveBizCustomsorderResponse) SetReqMsgId(v string) *SaveBizCustomsorderResponse {
	s.ReqMsgId = &v
	return s
}

func (s *SaveBizCustomsorderResponse) SetResultCode(v string) *SaveBizCustomsorderResponse {
	s.ResultCode = &v
	return s
}

func (s *SaveBizCustomsorderResponse) SetResultMsg(v string) *SaveBizCustomsorderResponse {
	s.ResultMsg = &v
	return s
}

func (s *SaveBizCustomsorderResponse) SetTxCodes(v []*TxDto) *SaveBizCustomsorderResponse {
	s.TxCodes = v
	return s
}

type SaveBizVehicleRequest struct {
	// OAuth模式下的授权token
	AuthToken         *string `json:"auth_token,omitempty" xml:"auth_token,omitempty"`
	ProductInstanceId *string `json:"product_instance_id,omitempty" xml:"product_instance_id,omitempty"`
	// 操作动作,为空为新增或更新，UPDATE为更新，DELETE为删除
	Action *string `json:"action,omitempty" xml:"action,omitempty"`
	// 箱子信息 业务必填
	ContainerParams []*VehicleContainerParam `json:"container_params,omitempty" xml:"container_params,omitempty" type:"Repeated"`
	// 货代did
	ForwarderDid *string `json:"forwarder_did,omitempty" xml:"forwarder_did,omitempty" require:"true"`
	// 订单号
	OrderNo *string `json:"order_no,omitempty" xml:"order_no,omitempty" require:"true"`
	// 拖车单号
	VehicleCode *string `json:"vehicle_code,omitempty" xml:"vehicle_code,omitempty" require:"true"`
	// 车牌号
	VehicleNo *string `json:"vehicle_no,omitempty" xml:"vehicle_no,omitempty"`
}

func (s SaveBizVehicleRequest) String() string {
	return tea.Prettify(s)
}

func (s SaveBizVehicleRequest) GoString() string {
	return s.String()
}

func (s *SaveBizVehicleRequest) SetAuthToken(v string) *SaveBizVehicleRequest {
	s.AuthToken = &v
	return s
}

func (s *SaveBizVehicleRequest) SetProductInstanceId(v string) *SaveBizVehicleRequest {
	s.ProductInstanceId = &v
	return s
}

func (s *SaveBizVehicleRequest) SetAction(v string) *SaveBizVehicleRequest {
	s.Action = &v
	return s
}

func (s *SaveBizVehicleRequest) SetContainerParams(v []*VehicleContainerParam) *SaveBizVehicleRequest {
	s.ContainerParams = v
	return s
}

func (s *SaveBizVehicleRequest) SetForwarderDid(v string) *SaveBizVehicleRequest {
	s.ForwarderDid = &v
	return s
}

func (s *SaveBizVehicleRequest) SetOrderNo(v string) *SaveBizVehicleRequest {
	s.OrderNo = &v
	return s
}

func (s *SaveBizVehicleRequest) SetVehicleCode(v string) *SaveBizVehicleRequest {
	s.VehicleCode = &v
	return s
}

func (s *SaveBizVehicleRequest) SetVehicleNo(v string) *SaveBizVehicleRequest {
	s.VehicleNo = &v
	return s
}

type SaveBizVehicleResponse struct {
	// 请求唯一ID，用于链路跟踪和问题排查
	ReqMsgId *string `json:"req_msg_id,omitempty" xml:"req_msg_id,omitempty"`
	// 结果码，一般OK表示调用成功
	ResultCode *string `json:"result_code,omitempty" xml:"result_code,omitempty"`
	// 异常信息的文本描述
	ResultMsg *string `json:"result_msg,omitempty" xml:"result_msg,omitempty"`
	// 链上凭证
	TxCodes []*TxDto `json:"tx_codes,omitempty" xml:"tx_codes,omitempty" type:"Repeated"`
}

func (s SaveBizVehicleResponse) String() string {
	return tea.Prettify(s)
}

func (s SaveBizVehicleResponse) GoString() string {
	return s.String()
}

func (s *SaveBizVehicleResponse) SetReqMsgId(v string) *SaveBizVehicleResponse {
	s.ReqMsgId = &v
	return s
}

func (s *SaveBizVehicleResponse) SetResultCode(v string) *SaveBizVehicleResponse {
	s.ResultCode = &v
	return s
}

func (s *SaveBizVehicleResponse) SetResultMsg(v string) *SaveBizVehicleResponse {
	s.ResultMsg = &v
	return s
}

func (s *SaveBizVehicleResponse) SetTxCodes(v []*TxDto) *SaveBizVehicleResponse {
	s.TxCodes = v
	return s
}

type SaveBizMasterblRequest struct {
	// OAuth模式下的授权token
	AuthToken         *string `json:"auth_token,omitempty" xml:"auth_token,omitempty"`
	ProductInstanceId *string `json:"product_instance_id,omitempty" xml:"product_instance_id,omitempty"`
	// 操作动作,为空为新增或更新，UPDATE为更新，DELETE为删除
	Action *string `json:"action,omitempty" xml:"action,omitempty"`
	// 订舱单
	BlBookingParams []*MasterBlBookingParam `json:"bl_booking_params,omitempty" xml:"bl_booking_params,omitempty" type:"Repeated"`
	// 提单要求
	BlRequest *string `json:"bl_request,omitempty" xml:"bl_request,omitempty"`
	// 提单类型
	BlType *string `json:"bl_type,omitempty" xml:"bl_type,omitempty"`
	// 船公司 业务必填
	Carrier *string `json:"carrier,omitempty" xml:"carrier,omitempty"`
	// 收货人 业务必填
	//
	Consignee *string `json:"consignee,omitempty" xml:"consignee,omitempty"`
	// 集装箱列表 业务必填
	//
	ContainerParams []*MasterBlContainerParam `json:"container_params,omitempty" xml:"container_params,omitempty" type:"Repeated"`
	// 约号
	Contract *string `json:"contract,omitempty" xml:"contract,omitempty"`
	// 目的地. 业务必填
	//
	DeliveryPlace *string `json:"delivery_place,omitempty" xml:"delivery_place,omitempty"`
	// 运输条款
	DeliveryTerms *string `json:"delivery_terms,omitempty" xml:"delivery_terms,omitempty"`
	// 卸货港. 业务必填
	//
	DischargePort *string `json:"discharge_port,omitempty" xml:"discharge_port,omitempty"`
	//
	// 货代did
	ForwarderDid *string `json:"forwarder_did,omitempty" xml:"forwarder_did,omitempty" require:"true"`
	// 运费/约号
	Freight *string `json:"freight,omitempty" xml:"freight,omitempty"`
	// 货物列表. 业务必填
	//
	GoodsParams []*MasterBlGoodsParam `json:"goods_params,omitempty" xml:"goods_params,omitempty" type:"Repeated"`
	// 出单日期
	IssueDate *int64 `json:"issue_date,omitempty" xml:"issue_date,omitempty"`
	// 放单地点
	IssuePlace *string `json:"issue_place,omitempty" xml:"issue_place,omitempty"`
	// 起运港  业务必填
	//
	LoadingPort *string `json:"loading_port,omitempty" xml:"loading_port,omitempty"`
	// master提单号
	MasterBlNo *string `json:"master_bl_no,omitempty" xml:"master_bl_no,omitempty" require:"true"`
	// 裝卸方式
	Movement *string `json:"movement,omitempty" xml:"movement,omitempty"`
	// 通知方
	NotifyParty *string `json:"notify_party,omitempty" xml:"notify_party,omitempty"`
	// 开船日期
	OnBoardDate *int64 `json:"on_board_date,omitempty" xml:"on_board_date,omitempty"`
	// 船状态
	OnBoardStatus *string `json:"on_board_status,omitempty" xml:"on_board_status,omitempty"`
	//
	// 订单号
	OrderNo *string `json:"order_no,omitempty" xml:"order_no,omitempty" require:"true"`
	// 付费方式  业务必填
	//
	PaymentTerms *string `json:"payment_terms,omitempty" xml:"payment_terms,omitempty"`
	// 前程运输
	PrCarriage *string `json:"pr_carriage,omitempty" xml:"pr_carriage,omitempty"`
	// 其他
	Remark *string `json:"remark,omitempty" xml:"remark,omitempty"`
	//  发货人. 业务必填
	//
	Shipper *string `json:"shipper,omitempty" xml:"shipper,omitempty"`
	// 特殊字段无要求非必填
	Source *string `json:"source,omitempty" xml:"source,omitempty"`
	// 运输方式
	Transportation *string `json:"transportation,omitempty" xml:"transportation,omitempty"`
	// 航名 业务必填
	Vessel *string `json:"vessel,omitempty" xml:"vessel,omitempty"`
	// 航次 业务必填
	Voyage *string `json:"voyage,omitempty" xml:"voyage,omitempty"`
}

func (s SaveBizMasterblRequest) String() string {
	return tea.Prettify(s)
}

func (s SaveBizMasterblRequest) GoString() string {
	return s.String()
}

func (s *SaveBizMasterblRequest) SetAuthToken(v string) *SaveBizMasterblRequest {
	s.AuthToken = &v
	return s
}

func (s *SaveBizMasterblRequest) SetProductInstanceId(v string) *SaveBizMasterblRequest {
	s.ProductInstanceId = &v
	return s
}

func (s *SaveBizMasterblRequest) SetAction(v string) *SaveBizMasterblRequest {
	s.Action = &v
	return s
}

func (s *SaveBizMasterblRequest) SetBlBookingParams(v []*MasterBlBookingParam) *SaveBizMasterblRequest {
	s.BlBookingParams = v
	return s
}

func (s *SaveBizMasterblRequest) SetBlRequest(v string) *SaveBizMasterblRequest {
	s.BlRequest = &v
	return s
}

func (s *SaveBizMasterblRequest) SetBlType(v string) *SaveBizMasterblRequest {
	s.BlType = &v
	return s
}

func (s *SaveBizMasterblRequest) SetCarrier(v string) *SaveBizMasterblRequest {
	s.Carrier = &v
	return s
}

func (s *SaveBizMasterblRequest) SetConsignee(v string) *SaveBizMasterblRequest {
	s.Consignee = &v
	return s
}

func (s *SaveBizMasterblRequest) SetContainerParams(v []*MasterBlContainerParam) *SaveBizMasterblRequest {
	s.ContainerParams = v
	return s
}

func (s *SaveBizMasterblRequest) SetContract(v string) *SaveBizMasterblRequest {
	s.Contract = &v
	return s
}

func (s *SaveBizMasterblRequest) SetDeliveryPlace(v string) *SaveBizMasterblRequest {
	s.DeliveryPlace = &v
	return s
}

func (s *SaveBizMasterblRequest) SetDeliveryTerms(v string) *SaveBizMasterblRequest {
	s.DeliveryTerms = &v
	return s
}

func (s *SaveBizMasterblRequest) SetDischargePort(v string) *SaveBizMasterblRequest {
	s.DischargePort = &v
	return s
}

func (s *SaveBizMasterblRequest) SetForwarderDid(v string) *SaveBizMasterblRequest {
	s.ForwarderDid = &v
	return s
}

func (s *SaveBizMasterblRequest) SetFreight(v string) *SaveBizMasterblRequest {
	s.Freight = &v
	return s
}

func (s *SaveBizMasterblRequest) SetGoodsParams(v []*MasterBlGoodsParam) *SaveBizMasterblRequest {
	s.GoodsParams = v
	return s
}

func (s *SaveBizMasterblRequest) SetIssueDate(v int64) *SaveBizMasterblRequest {
	s.IssueDate = &v
	return s
}

func (s *SaveBizMasterblRequest) SetIssuePlace(v string) *SaveBizMasterblRequest {
	s.IssuePlace = &v
	return s
}

func (s *SaveBizMasterblRequest) SetLoadingPort(v string) *SaveBizMasterblRequest {
	s.LoadingPort = &v
	return s
}

func (s *SaveBizMasterblRequest) SetMasterBlNo(v string) *SaveBizMasterblRequest {
	s.MasterBlNo = &v
	return s
}

func (s *SaveBizMasterblRequest) SetMovement(v string) *SaveBizMasterblRequest {
	s.Movement = &v
	return s
}

func (s *SaveBizMasterblRequest) SetNotifyParty(v string) *SaveBizMasterblRequest {
	s.NotifyParty = &v
	return s
}

func (s *SaveBizMasterblRequest) SetOnBoardDate(v int64) *SaveBizMasterblRequest {
	s.OnBoardDate = &v
	return s
}

func (s *SaveBizMasterblRequest) SetOnBoardStatus(v string) *SaveBizMasterblRequest {
	s.OnBoardStatus = &v
	return s
}

func (s *SaveBizMasterblRequest) SetOrderNo(v string) *SaveBizMasterblRequest {
	s.OrderNo = &v
	return s
}

func (s *SaveBizMasterblRequest) SetPaymentTerms(v string) *SaveBizMasterblRequest {
	s.PaymentTerms = &v
	return s
}

func (s *SaveBizMasterblRequest) SetPrCarriage(v string) *SaveBizMasterblRequest {
	s.PrCarriage = &v
	return s
}

func (s *SaveBizMasterblRequest) SetRemark(v string) *SaveBizMasterblRequest {
	s.Remark = &v
	return s
}

func (s *SaveBizMasterblRequest) SetShipper(v string) *SaveBizMasterblRequest {
	s.Shipper = &v
	return s
}

func (s *SaveBizMasterblRequest) SetSource(v string) *SaveBizMasterblRequest {
	s.Source = &v
	return s
}

func (s *SaveBizMasterblRequest) SetTransportation(v string) *SaveBizMasterblRequest {
	s.Transportation = &v
	return s
}

func (s *SaveBizMasterblRequest) SetVessel(v string) *SaveBizMasterblRequest {
	s.Vessel = &v
	return s
}

func (s *SaveBizMasterblRequest) SetVoyage(v string) *SaveBizMasterblRequest {
	s.Voyage = &v
	return s
}

type SaveBizMasterblResponse struct {
	// 请求唯一ID，用于链路跟踪和问题排查
	ReqMsgId *string `json:"req_msg_id,omitempty" xml:"req_msg_id,omitempty"`
	// 结果码，一般OK表示调用成功
	ResultCode *string `json:"result_code,omitempty" xml:"result_code,omitempty"`
	// 异常信息的文本描述
	ResultMsg *string `json:"result_msg,omitempty" xml:"result_msg,omitempty"`
	// 链上凭证
	TxCodes []*TxDto `json:"tx_codes,omitempty" xml:"tx_codes,omitempty" type:"Repeated"`
}

func (s SaveBizMasterblResponse) String() string {
	return tea.Prettify(s)
}

func (s SaveBizMasterblResponse) GoString() string {
	return s.String()
}

func (s *SaveBizMasterblResponse) SetReqMsgId(v string) *SaveBizMasterblResponse {
	s.ReqMsgId = &v
	return s
}

func (s *SaveBizMasterblResponse) SetResultCode(v string) *SaveBizMasterblResponse {
	s.ResultCode = &v
	return s
}

func (s *SaveBizMasterblResponse) SetResultMsg(v string) *SaveBizMasterblResponse {
	s.ResultMsg = &v
	return s
}

func (s *SaveBizMasterblResponse) SetTxCodes(v []*TxDto) *SaveBizMasterblResponse {
	s.TxCodes = v
	return s
}

type FinishBizAuditRequest struct {
	// OAuth模式下的授权token
	AuthToken         *string `json:"auth_token,omitempty" xml:"auth_token,omitempty"`
	ProductInstanceId *string `json:"product_instance_id,omitempty" xml:"product_instance_id,omitempty"`
	//
	// 货代did
	ForwarderDid *string `json:"forwarder_did,omitempty" xml:"forwarder_did,omitempty" require:"true"`
	//
	// 订单号
	OrderNo *string `json:"order_no,omitempty" xml:"order_no,omitempty" require:"true"`
}

func (s FinishBizAuditRequest) String() string {
	return tea.Prettify(s)
}

func (s FinishBizAuditRequest) GoString() string {
	return s.String()
}

func (s *FinishBizAuditRequest) SetAuthToken(v string) *FinishBizAuditRequest {
	s.AuthToken = &v
	return s
}

func (s *FinishBizAuditRequest) SetProductInstanceId(v string) *FinishBizAuditRequest {
	s.ProductInstanceId = &v
	return s
}

func (s *FinishBizAuditRequest) SetForwarderDid(v string) *FinishBizAuditRequest {
	s.ForwarderDid = &v
	return s
}

func (s *FinishBizAuditRequest) SetOrderNo(v string) *FinishBizAuditRequest {
	s.OrderNo = &v
	return s
}

type FinishBizAuditResponse struct {
	// 请求唯一ID，用于链路跟踪和问题排查
	ReqMsgId *string `json:"req_msg_id,omitempty" xml:"req_msg_id,omitempty"`
	// 结果码，一般OK表示调用成功
	ResultCode *string `json:"result_code,omitempty" xml:"result_code,omitempty"`
	// 异常信息的文本描述
	ResultMsg *string `json:"result_msg,omitempty" xml:"result_msg,omitempty"`
	// 链上凭证
	TxCodes []*TxDto `json:"tx_codes,omitempty" xml:"tx_codes,omitempty" type:"Repeated"`
}

func (s FinishBizAuditResponse) String() string {
	return tea.Prettify(s)
}

func (s FinishBizAuditResponse) GoString() string {
	return s.String()
}

func (s *FinishBizAuditResponse) SetReqMsgId(v string) *FinishBizAuditResponse {
	s.ReqMsgId = &v
	return s
}

func (s *FinishBizAuditResponse) SetResultCode(v string) *FinishBizAuditResponse {
	s.ResultCode = &v
	return s
}

func (s *FinishBizAuditResponse) SetResultMsg(v string) *FinishBizAuditResponse {
	s.ResultMsg = &v
	return s
}

func (s *FinishBizAuditResponse) SetTxCodes(v []*TxDto) *FinishBizAuditResponse {
	s.TxCodes = v
	return s
}

type SaveBizHouseblRequest struct {
	// OAuth模式下的授权token
	AuthToken         *string `json:"auth_token,omitempty" xml:"auth_token,omitempty"`
	ProductInstanceId *string `json:"product_instance_id,omitempty" xml:"product_instance_id,omitempty"`
	// 操作动作,为空为新增或更新，UPDATE为更新，DELETE为删除
	Action *string `json:"action,omitempty" xml:"action,omitempty"`
	// 订舱单
	BlBookingParams []*HouseBlBookingParam `json:"bl_booking_params,omitempty" xml:"bl_booking_params,omitempty" type:"Repeated"`
	// 提单签发单位
	BlIssuingAgency *string `json:"bl_issuing_agency,omitempty" xml:"bl_issuing_agency,omitempty"`
	// 提单要求
	BlRequest *string `json:"bl_request,omitempty" xml:"bl_request,omitempty"`
	// 提单类型
	BlType *string `json:"bl_type,omitempty" xml:"bl_type,omitempty"`
	// 船公司 业务必填
	Carrier *string `json:"carrier,omitempty" xml:"carrier,omitempty"`
	// 收货人 业务必填
	//
	Consignee *string `json:"consignee,omitempty" xml:"consignee,omitempty"`
	// 集装箱列表 业务必填
	//
	ContainerParams []*HouseBlContainerParam `json:"container_params,omitempty" xml:"container_params,omitempty" type:"Repeated"`
	// 约号
	Contract *string `json:"contract,omitempty" xml:"contract,omitempty"`
	// 目的地  业务必填
	//
	DeliveryPlace *string `json:"delivery_place,omitempty" xml:"delivery_place,omitempty"`
	// 运输条款
	DeliveryTerms *string `json:"delivery_terms,omitempty" xml:"delivery_terms,omitempty"`
	// 卸货港 业务必填
	//
	DischargePort *string `json:"discharge_port,omitempty" xml:"discharge_port,omitempty"`
	// 货代did
	ForwarderDid *string `json:"forwarder_did,omitempty" xml:"forwarder_did,omitempty" require:"true"`
	// 运费
	Freight *string `json:"freight,omitempty" xml:"freight,omitempty"`
	// 货物列表 业务必填
	//
	GoodsParams []*HouseBlGoodsParam `json:"goods_params,omitempty" xml:"goods_params,omitempty" type:"Repeated"`
	// house提单号
	HouseBlNo *string `json:"house_bl_no,omitempty" xml:"house_bl_no,omitempty" require:"true"`
	// 出单日期
	IssueDate *int64 `json:"issue_date,omitempty" xml:"issue_date,omitempty"`
	// 放单地点
	IssuePlace *string `json:"issue_place,omitempty" xml:"issue_place,omitempty"`
	// 起运港 业务必填
	//
	LoadingPort *string `json:"loading_port,omitempty" xml:"loading_port,omitempty"`
	// master提单号
	MasterBlNo *string `json:"master_bl_no,omitempty" xml:"master_bl_no,omitempty"`
	// 裝卸方式
	Movement *string `json:"movement,omitempty" xml:"movement,omitempty"`
	// 通知方
	NotifyParty *string `json:"notify_party,omitempty" xml:"notify_party,omitempty"`
	// 开船日期
	OnBoardDate *int64 `json:"on_board_date,omitempty" xml:"on_board_date,omitempty"`
	// 订单号
	OrderNo *string `json:"order_no,omitempty" xml:"order_no,omitempty" require:"true"`
	// 付费方式 业务必填
	//
	PaymentTerms *string `json:"payment_terms,omitempty" xml:"payment_terms,omitempty"`
	// 其他
	Remark *string `json:"remark,omitempty" xml:"remark,omitempty"`
	// 发货人 业务必填
	//
	Shipper *string `json:"shipper,omitempty" xml:"shipper,omitempty"`
	// 运输方式
	Transportation *string `json:"transportation,omitempty" xml:"transportation,omitempty"`
	// 航名 业务必填
	Vessel *string `json:"vessel,omitempty" xml:"vessel,omitempty"`
	// 航次 业务必填
	Voyage *string `json:"voyage,omitempty" xml:"voyage,omitempty"`
}

func (s SaveBizHouseblRequest) String() string {
	return tea.Prettify(s)
}

func (s SaveBizHouseblRequest) GoString() string {
	return s.String()
}

func (s *SaveBizHouseblRequest) SetAuthToken(v string) *SaveBizHouseblRequest {
	s.AuthToken = &v
	return s
}

func (s *SaveBizHouseblRequest) SetProductInstanceId(v string) *SaveBizHouseblRequest {
	s.ProductInstanceId = &v
	return s
}

func (s *SaveBizHouseblRequest) SetAction(v string) *SaveBizHouseblRequest {
	s.Action = &v
	return s
}

func (s *SaveBizHouseblRequest) SetBlBookingParams(v []*HouseBlBookingParam) *SaveBizHouseblRequest {
	s.BlBookingParams = v
	return s
}

func (s *SaveBizHouseblRequest) SetBlIssuingAgency(v string) *SaveBizHouseblRequest {
	s.BlIssuingAgency = &v
	return s
}

func (s *SaveBizHouseblRequest) SetBlRequest(v string) *SaveBizHouseblRequest {
	s.BlRequest = &v
	return s
}

func (s *SaveBizHouseblRequest) SetBlType(v string) *SaveBizHouseblRequest {
	s.BlType = &v
	return s
}

func (s *SaveBizHouseblRequest) SetCarrier(v string) *SaveBizHouseblRequest {
	s.Carrier = &v
	return s
}

func (s *SaveBizHouseblRequest) SetConsignee(v string) *SaveBizHouseblRequest {
	s.Consignee = &v
	return s
}

func (s *SaveBizHouseblRequest) SetContainerParams(v []*HouseBlContainerParam) *SaveBizHouseblRequest {
	s.ContainerParams = v
	return s
}

func (s *SaveBizHouseblRequest) SetContract(v string) *SaveBizHouseblRequest {
	s.Contract = &v
	return s
}

func (s *SaveBizHouseblRequest) SetDeliveryPlace(v string) *SaveBizHouseblRequest {
	s.DeliveryPlace = &v
	return s
}

func (s *SaveBizHouseblRequest) SetDeliveryTerms(v string) *SaveBizHouseblRequest {
	s.DeliveryTerms = &v
	return s
}

func (s *SaveBizHouseblRequest) SetDischargePort(v string) *SaveBizHouseblRequest {
	s.DischargePort = &v
	return s
}

func (s *SaveBizHouseblRequest) SetForwarderDid(v string) *SaveBizHouseblRequest {
	s.ForwarderDid = &v
	return s
}

func (s *SaveBizHouseblRequest) SetFreight(v string) *SaveBizHouseblRequest {
	s.Freight = &v
	return s
}

func (s *SaveBizHouseblRequest) SetGoodsParams(v []*HouseBlGoodsParam) *SaveBizHouseblRequest {
	s.GoodsParams = v
	return s
}

func (s *SaveBizHouseblRequest) SetHouseBlNo(v string) *SaveBizHouseblRequest {
	s.HouseBlNo = &v
	return s
}

func (s *SaveBizHouseblRequest) SetIssueDate(v int64) *SaveBizHouseblRequest {
	s.IssueDate = &v
	return s
}

func (s *SaveBizHouseblRequest) SetIssuePlace(v string) *SaveBizHouseblRequest {
	s.IssuePlace = &v
	return s
}

func (s *SaveBizHouseblRequest) SetLoadingPort(v string) *SaveBizHouseblRequest {
	s.LoadingPort = &v
	return s
}

func (s *SaveBizHouseblRequest) SetMasterBlNo(v string) *SaveBizHouseblRequest {
	s.MasterBlNo = &v
	return s
}

func (s *SaveBizHouseblRequest) SetMovement(v string) *SaveBizHouseblRequest {
	s.Movement = &v
	return s
}

func (s *SaveBizHouseblRequest) SetNotifyParty(v string) *SaveBizHouseblRequest {
	s.NotifyParty = &v
	return s
}

func (s *SaveBizHouseblRequest) SetOnBoardDate(v int64) *SaveBizHouseblRequest {
	s.OnBoardDate = &v
	return s
}

func (s *SaveBizHouseblRequest) SetOrderNo(v string) *SaveBizHouseblRequest {
	s.OrderNo = &v
	return s
}

func (s *SaveBizHouseblRequest) SetPaymentTerms(v string) *SaveBizHouseblRequest {
	s.PaymentTerms = &v
	return s
}

func (s *SaveBizHouseblRequest) SetRemark(v string) *SaveBizHouseblRequest {
	s.Remark = &v
	return s
}

func (s *SaveBizHouseblRequest) SetShipper(v string) *SaveBizHouseblRequest {
	s.Shipper = &v
	return s
}

func (s *SaveBizHouseblRequest) SetTransportation(v string) *SaveBizHouseblRequest {
	s.Transportation = &v
	return s
}

func (s *SaveBizHouseblRequest) SetVessel(v string) *SaveBizHouseblRequest {
	s.Vessel = &v
	return s
}

func (s *SaveBizHouseblRequest) SetVoyage(v string) *SaveBizHouseblRequest {
	s.Voyage = &v
	return s
}

type SaveBizHouseblResponse struct {
	// 请求唯一ID，用于链路跟踪和问题排查
	ReqMsgId *string `json:"req_msg_id,omitempty" xml:"req_msg_id,omitempty"`
	// 结果码，一般OK表示调用成功
	ResultCode *string `json:"result_code,omitempty" xml:"result_code,omitempty"`
	// 异常信息的文本描述
	ResultMsg *string `json:"result_msg,omitempty" xml:"result_msg,omitempty"`
	// 链上凭证
	TxCodes []*TxDto `json:"tx_codes,omitempty" xml:"tx_codes,omitempty" type:"Repeated"`
}

func (s SaveBizHouseblResponse) String() string {
	return tea.Prettify(s)
}

func (s SaveBizHouseblResponse) GoString() string {
	return s.String()
}

func (s *SaveBizHouseblResponse) SetReqMsgId(v string) *SaveBizHouseblResponse {
	s.ReqMsgId = &v
	return s
}

func (s *SaveBizHouseblResponse) SetResultCode(v string) *SaveBizHouseblResponse {
	s.ResultCode = &v
	return s
}

func (s *SaveBizHouseblResponse) SetResultMsg(v string) *SaveBizHouseblResponse {
	s.ResultMsg = &v
	return s
}

func (s *SaveBizHouseblResponse) SetTxCodes(v []*TxDto) *SaveBizHouseblResponse {
	s.TxCodes = v
	return s
}

type CreateBillPaybillorderRequest struct {
	// OAuth模式下的授权token
	AuthToken         *string `json:"auth_token,omitempty" xml:"auth_token,omitempty"`
	ProductInstanceId *string `json:"product_instance_id,omitempty" xml:"product_instance_id,omitempty"`
	// 申请时间
	ApplyDate *int64 `json:"apply_date,omitempty" xml:"apply_date,omitempty"`
	// 银行收款账户
	BankReceiptAccount *string `json:"bank_receipt_account,omitempty" xml:"bank_receipt_account,omitempty"`
	// 币种
	Currency *string `json:"currency,omitempty" xml:"currency,omitempty"`
	// 开户行
	DepositBank *string `json:"deposit_bank,omitempty" xml:"deposit_bank,omitempty"`
	// 货代did
	ForwarderDid *string `json:"forwarder_did,omitempty" xml:"forwarder_did,omitempty" require:"true"`
	// 应付总额
	PayAmount *string `json:"pay_amount,omitempty" xml:"pay_amount,omitempty"`
	// 付款单编号
	PayBillOrderCode *string `json:"pay_bill_order_code,omitempty" xml:"pay_bill_order_code,omitempty" require:"true"`
	// 应付账单资费项 业务必填
	PayBillTariffParams []*PayBillTariffParam `json:"pay_bill_tariff_params,omitempty" xml:"pay_bill_tariff_params,omitempty" type:"Repeated"`
	// 付款公司
	PayCompany *string `json:"pay_company,omitempty" xml:"pay_company,omitempty"`
	// 付款公司企业信用号
	//
	PayCompanyCertNo *string `json:"pay_company_cert_no,omitempty" xml:"pay_company_cert_no,omitempty"`
	// 付款公司did
	PayCompanyDid *string `json:"pay_company_did,omitempty" xml:"pay_company_did,omitempty"`
	// 付款期限
	PayDeadline *string `json:"pay_deadline,omitempty" xml:"pay_deadline,omitempty"`
	// 收款客户
	ReceiptClient *string `json:"receipt_client,omitempty" xml:"receipt_client,omitempty"`
	// 收款客户企业信用号
	//
	ReceiptClientCertNo *string `json:"receipt_client_cert_no,omitempty" xml:"receipt_client_cert_no,omitempty"`
	// 收款客户did
	ReceiptClientDid *string `json:"receipt_client_did,omitempty" xml:"receipt_client_did,omitempty"`
	// 对应应收资费项code
	ReceiptTariffCodes []*string `json:"receipt_tariff_codes,omitempty" xml:"receipt_tariff_codes,omitempty" type:"Repeated"`
	// 结算公司
	SettleCompany *string `json:"settle_company,omitempty" xml:"settle_company,omitempty"`
	// 结算公司企业信用号
	//
	SettleCompanyCertNo *string `json:"settle_company_cert_no,omitempty" xml:"settle_company_cert_no,omitempty"`
	// 结算公司did
	SettleCompanyDid *string `json:"settle_company_did,omitempty" xml:"settle_company_did,omitempty"`
	// 账单结算周期
	SettleCycle *string `json:"settle_cycle,omitempty" xml:"settle_cycle,omitempty"`
}

func (s CreateBillPaybillorderRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateBillPaybillorderRequest) GoString() string {
	return s.String()
}

func (s *CreateBillPaybillorderRequest) SetAuthToken(v string) *CreateBillPaybillorderRequest {
	s.AuthToken = &v
	return s
}

func (s *CreateBillPaybillorderRequest) SetProductInstanceId(v string) *CreateBillPaybillorderRequest {
	s.ProductInstanceId = &v
	return s
}

func (s *CreateBillPaybillorderRequest) SetApplyDate(v int64) *CreateBillPaybillorderRequest {
	s.ApplyDate = &v
	return s
}

func (s *CreateBillPaybillorderRequest) SetBankReceiptAccount(v string) *CreateBillPaybillorderRequest {
	s.BankReceiptAccount = &v
	return s
}

func (s *CreateBillPaybillorderRequest) SetCurrency(v string) *CreateBillPaybillorderRequest {
	s.Currency = &v
	return s
}

func (s *CreateBillPaybillorderRequest) SetDepositBank(v string) *CreateBillPaybillorderRequest {
	s.DepositBank = &v
	return s
}

func (s *CreateBillPaybillorderRequest) SetForwarderDid(v string) *CreateBillPaybillorderRequest {
	s.ForwarderDid = &v
	return s
}

func (s *CreateBillPaybillorderRequest) SetPayAmount(v string) *CreateBillPaybillorderRequest {
	s.PayAmount = &v
	return s
}

func (s *CreateBillPaybillorderRequest) SetPayBillOrderCode(v string) *CreateBillPaybillorderRequest {
	s.PayBillOrderCode = &v
	return s
}

func (s *CreateBillPaybillorderRequest) SetPayBillTariffParams(v []*PayBillTariffParam) *CreateBillPaybillorderRequest {
	s.PayBillTariffParams = v
	return s
}

func (s *CreateBillPaybillorderRequest) SetPayCompany(v string) *CreateBillPaybillorderRequest {
	s.PayCompany = &v
	return s
}

func (s *CreateBillPaybillorderRequest) SetPayCompanyCertNo(v string) *CreateBillPaybillorderRequest {
	s.PayCompanyCertNo = &v
	return s
}

func (s *CreateBillPaybillorderRequest) SetPayCompanyDid(v string) *CreateBillPaybillorderRequest {
	s.PayCompanyDid = &v
	return s
}

func (s *CreateBillPaybillorderRequest) SetPayDeadline(v string) *CreateBillPaybillorderRequest {
	s.PayDeadline = &v
	return s
}

func (s *CreateBillPaybillorderRequest) SetReceiptClient(v string) *CreateBillPaybillorderRequest {
	s.ReceiptClient = &v
	return s
}

func (s *CreateBillPaybillorderRequest) SetReceiptClientCertNo(v string) *CreateBillPaybillorderRequest {
	s.ReceiptClientCertNo = &v
	return s
}

func (s *CreateBillPaybillorderRequest) SetReceiptClientDid(v string) *CreateBillPaybillorderRequest {
	s.ReceiptClientDid = &v
	return s
}

func (s *CreateBillPaybillorderRequest) SetReceiptTariffCodes(v []*string) *CreateBillPaybillorderRequest {
	s.ReceiptTariffCodes = v
	return s
}

func (s *CreateBillPaybillorderRequest) SetSettleCompany(v string) *CreateBillPaybillorderRequest {
	s.SettleCompany = &v
	return s
}

func (s *CreateBillPaybillorderRequest) SetSettleCompanyCertNo(v string) *CreateBillPaybillorderRequest {
	s.SettleCompanyCertNo = &v
	return s
}

func (s *CreateBillPaybillorderRequest) SetSettleCompanyDid(v string) *CreateBillPaybillorderRequest {
	s.SettleCompanyDid = &v
	return s
}

func (s *CreateBillPaybillorderRequest) SetSettleCycle(v string) *CreateBillPaybillorderRequest {
	s.SettleCycle = &v
	return s
}

type CreateBillPaybillorderResponse struct {
	// 请求唯一ID，用于链路跟踪和问题排查
	ReqMsgId *string `json:"req_msg_id,omitempty" xml:"req_msg_id,omitempty"`
	// 结果码，一般OK表示调用成功
	ResultCode *string `json:"result_code,omitempty" xml:"result_code,omitempty"`
	// 异常信息的文本描述
	ResultMsg *string `json:"result_msg,omitempty" xml:"result_msg,omitempty"`
	// 链上凭证hash
	TxCodes []*TxDto `json:"tx_codes,omitempty" xml:"tx_codes,omitempty" type:"Repeated"`
}

func (s CreateBillPaybillorderResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateBillPaybillorderResponse) GoString() string {
	return s.String()
}

func (s *CreateBillPaybillorderResponse) SetReqMsgId(v string) *CreateBillPaybillorderResponse {
	s.ReqMsgId = &v
	return s
}

func (s *CreateBillPaybillorderResponse) SetResultCode(v string) *CreateBillPaybillorderResponse {
	s.ResultCode = &v
	return s
}

func (s *CreateBillPaybillorderResponse) SetResultMsg(v string) *CreateBillPaybillorderResponse {
	s.ResultMsg = &v
	return s
}

func (s *CreateBillPaybillorderResponse) SetTxCodes(v []*TxDto) *CreateBillPaybillorderResponse {
	s.TxCodes = v
	return s
}

type CreateBillReceiptbillorderRequest struct {
	// OAuth模式下的授权token
	AuthToken         *string `json:"auth_token,omitempty" xml:"auth_token,omitempty"`
	ProductInstanceId *string `json:"product_instance_id,omitempty" xml:"product_instance_id,omitempty"`
	// 实际收款结算公司
	ActualSettleCompany *string `json:"actual_settle_company,omitempty" xml:"actual_settle_company,omitempty"`
	// 实际收款结算公司企业信用号
	ActualSettleCompanyCertNo *string `json:"actual_settle_company_cert_no,omitempty" xml:"actual_settle_company_cert_no,omitempty"`
	// 实际收款结算公司did
	ActualSettleCompanyDid *string `json:"actual_settle_company_did,omitempty" xml:"actual_settle_company_did,omitempty"`
	// 确认时间  业务必填
	ApplyDate *int64 `json:"apply_date,omitempty" xml:"apply_date,omitempty"`
	// 揽货类型
	ClientType *string `json:"client_type,omitempty" xml:"client_type,omitempty"`
	// 币种 业务必填
	Currency *string `json:"currency,omitempty" xml:"currency,omitempty"`
	// 货代did
	ForwarderDid *string `json:"forwarder_did,omitempty" xml:"forwarder_did,omitempty" require:"true"`
	// 开票抬头
	InvoiceTitle *string `json:"invoice_title,omitempty" xml:"invoice_title,omitempty"`
	// 含税价
	PriceIncludingTax *string `json:"price_including_tax,omitempty" xml:"price_including_tax,omitempty"`
	// 收款账号
	ReceiptAccount *string `json:"receipt_account,omitempty" xml:"receipt_account,omitempty"`
	// 收款总额 业务必填
	ReceiptAmount *string `json:"receipt_amount,omitempty" xml:"receipt_amount,omitempty"`
	// 收款账单编号
	ReceiptBillOrderCode *string `json:"receipt_bill_order_code,omitempty" xml:"receipt_bill_order_code,omitempty" require:"true"`
	// 应收资费项账单 业务必填
	ReceiptBillTariffParams []*ReceiptBillTariffParam `json:"receipt_bill_tariff_params,omitempty" xml:"receipt_bill_tariff_params,omitempty" type:"Repeated"`
	// 结算客户名称
	SettleClient *string `json:"settle_client,omitempty" xml:"settle_client,omitempty"`
	// 结算客户企业信用号
	SettleClientCertNo *string `json:"settle_client_cert_no,omitempty" xml:"settle_client_cert_no,omitempty"`
	// 结算客户名称did
	SettleClientDid *string `json:"settle_client_did,omitempty" xml:"settle_client_did,omitempty"`
	// 结算公司
	SettleCompany *string `json:"settle_company,omitempty" xml:"settle_company,omitempty"`
	// 结算公司企业信用号
	SettleCompanyCertNo *string `json:"settle_company_cert_no,omitempty" xml:"settle_company_cert_no,omitempty"`
	// 结算公司did
	SettleCompanyDid *string `json:"settle_company_did,omitempty" xml:"settle_company_did,omitempty"`
	// 账单结算周期
	SettleCycle *string `json:"settle_cycle,omitempty" xml:"settle_cycle,omitempty"`
	// 税金
	Taxes *string `json:"taxes,omitempty" xml:"taxes,omitempty"`
	// 未税价
	UntaxedPrice *string `json:"untaxed_price,omitempty" xml:"untaxed_price,omitempty"`
}

func (s CreateBillReceiptbillorderRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateBillReceiptbillorderRequest) GoString() string {
	return s.String()
}

func (s *CreateBillReceiptbillorderRequest) SetAuthToken(v string) *CreateBillReceiptbillorderRequest {
	s.AuthToken = &v
	return s
}

func (s *CreateBillReceiptbillorderRequest) SetProductInstanceId(v string) *CreateBillReceiptbillorderRequest {
	s.ProductInstanceId = &v
	return s
}

func (s *CreateBillReceiptbillorderRequest) SetActualSettleCompany(v string) *CreateBillReceiptbillorderRequest {
	s.ActualSettleCompany = &v
	return s
}

func (s *CreateBillReceiptbillorderRequest) SetActualSettleCompanyCertNo(v string) *CreateBillReceiptbillorderRequest {
	s.ActualSettleCompanyCertNo = &v
	return s
}

func (s *CreateBillReceiptbillorderRequest) SetActualSettleCompanyDid(v string) *CreateBillReceiptbillorderRequest {
	s.ActualSettleCompanyDid = &v
	return s
}

func (s *CreateBillReceiptbillorderRequest) SetApplyDate(v int64) *CreateBillReceiptbillorderRequest {
	s.ApplyDate = &v
	return s
}

func (s *CreateBillReceiptbillorderRequest) SetClientType(v string) *CreateBillReceiptbillorderRequest {
	s.ClientType = &v
	return s
}

func (s *CreateBillReceiptbillorderRequest) SetCurrency(v string) *CreateBillReceiptbillorderRequest {
	s.Currency = &v
	return s
}

func (s *CreateBillReceiptbillorderRequest) SetForwarderDid(v string) *CreateBillReceiptbillorderRequest {
	s.ForwarderDid = &v
	return s
}

func (s *CreateBillReceiptbillorderRequest) SetInvoiceTitle(v string) *CreateBillReceiptbillorderRequest {
	s.InvoiceTitle = &v
	return s
}

func (s *CreateBillReceiptbillorderRequest) SetPriceIncludingTax(v string) *CreateBillReceiptbillorderRequest {
	s.PriceIncludingTax = &v
	return s
}

func (s *CreateBillReceiptbillorderRequest) SetReceiptAccount(v string) *CreateBillReceiptbillorderRequest {
	s.ReceiptAccount = &v
	return s
}

func (s *CreateBillReceiptbillorderRequest) SetReceiptAmount(v string) *CreateBillReceiptbillorderRequest {
	s.ReceiptAmount = &v
	return s
}

func (s *CreateBillReceiptbillorderRequest) SetReceiptBillOrderCode(v string) *CreateBillReceiptbillorderRequest {
	s.ReceiptBillOrderCode = &v
	return s
}

func (s *CreateBillReceiptbillorderRequest) SetReceiptBillTariffParams(v []*ReceiptBillTariffParam) *CreateBillReceiptbillorderRequest {
	s.ReceiptBillTariffParams = v
	return s
}

func (s *CreateBillReceiptbillorderRequest) SetSettleClient(v string) *CreateBillReceiptbillorderRequest {
	s.SettleClient = &v
	return s
}

func (s *CreateBillReceiptbillorderRequest) SetSettleClientCertNo(v string) *CreateBillReceiptbillorderRequest {
	s.SettleClientCertNo = &v
	return s
}

func (s *CreateBillReceiptbillorderRequest) SetSettleClientDid(v string) *CreateBillReceiptbillorderRequest {
	s.SettleClientDid = &v
	return s
}

func (s *CreateBillReceiptbillorderRequest) SetSettleCompany(v string) *CreateBillReceiptbillorderRequest {
	s.SettleCompany = &v
	return s
}

func (s *CreateBillReceiptbillorderRequest) SetSettleCompanyCertNo(v string) *CreateBillReceiptbillorderRequest {
	s.SettleCompanyCertNo = &v
	return s
}

func (s *CreateBillReceiptbillorderRequest) SetSettleCompanyDid(v string) *CreateBillReceiptbillorderRequest {
	s.SettleCompanyDid = &v
	return s
}

func (s *CreateBillReceiptbillorderRequest) SetSettleCycle(v string) *CreateBillReceiptbillorderRequest {
	s.SettleCycle = &v
	return s
}

func (s *CreateBillReceiptbillorderRequest) SetTaxes(v string) *CreateBillReceiptbillorderRequest {
	s.Taxes = &v
	return s
}

func (s *CreateBillReceiptbillorderRequest) SetUntaxedPrice(v string) *CreateBillReceiptbillorderRequest {
	s.UntaxedPrice = &v
	return s
}

type CreateBillReceiptbillorderResponse struct {
	// 请求唯一ID，用于链路跟踪和问题排查
	ReqMsgId *string `json:"req_msg_id,omitempty" xml:"req_msg_id,omitempty"`
	// 结果码，一般OK表示调用成功
	ResultCode *string `json:"result_code,omitempty" xml:"result_code,omitempty"`
	// 异常信息的文本描述
	ResultMsg *string `json:"result_msg,omitempty" xml:"result_msg,omitempty"`
	// 链上凭证
	TxCodes []*TxDto `json:"tx_codes,omitempty" xml:"tx_codes,omitempty" type:"Repeated"`
}

func (s CreateBillReceiptbillorderResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateBillReceiptbillorderResponse) GoString() string {
	return s.String()
}

func (s *CreateBillReceiptbillorderResponse) SetReqMsgId(v string) *CreateBillReceiptbillorderResponse {
	s.ReqMsgId = &v
	return s
}

func (s *CreateBillReceiptbillorderResponse) SetResultCode(v string) *CreateBillReceiptbillorderResponse {
	s.ResultCode = &v
	return s
}

func (s *CreateBillReceiptbillorderResponse) SetResultMsg(v string) *CreateBillReceiptbillorderResponse {
	s.ResultMsg = &v
	return s
}

func (s *CreateBillReceiptbillorderResponse) SetTxCodes(v []*TxDto) *CreateBillReceiptbillorderResponse {
	s.TxCodes = v
	return s
}

type SaveBillPaybilltariffRequest struct {
	// OAuth模式下的授权token
	AuthToken         *string `json:"auth_token,omitempty" xml:"auth_token,omitempty"`
	ProductInstanceId *string `json:"product_instance_id,omitempty" xml:"product_instance_id,omitempty"`
	// 操作动作,为空为新增或更新，UPDATE为更新，DELETE为删除
	Action *string `json:"action,omitempty" xml:"action,omitempty"`
	// 计费数量 业务必填
	BillingNumber *string `json:"billing_number,omitempty" xml:"billing_number,omitempty"`
	// 计费类型 业务必填
	BillingType *string `json:"billing_type,omitempty" xml:"billing_type,omitempty"`
	// 航运订舱号[业务必填]
	BkgNo *string `json:"bkg_no,omitempty" xml:"bkg_no,omitempty"`
	// 订舱单号[业务必填]
	BookingNo *string `json:"booking_no,omitempty" xml:"booking_no,omitempty"`
	// 币种 业务必填
	Currency *string `json:"currency,omitempty" xml:"currency,omitempty"`
	// 货代did
	ForwarderDid *string `json:"forwarder_did,omitempty" xml:"forwarder_did,omitempty" require:"true"`
	// 收支  INCOME--收入， EXPENSES--支出
	// 业务必填
	IncomeOrExpenses *string `json:"income_or_expenses,omitempty" xml:"income_or_expenses,omitempty"`
	// 电子发票网址
	InvoiceUrl *string `json:"invoice_url,omitempty" xml:"invoice_url,omitempty"`
	// 工作单号 业务必填
	OrderNo *string `json:"order_no,omitempty" xml:"order_no,omitempty"`
	// 资费单据编号
	PayTariffCode *string `json:"pay_tariff_code,omitempty" xml:"pay_tariff_code,omitempty" require:"true"`
	// 资费项中文描述 业务必填
	PayTariffDesc *string `json:"pay_tariff_desc,omitempty" xml:"pay_tariff_desc,omitempty"`
	// 资费项目 业务必填
	PayTariffProject *string `json:"pay_tariff_project,omitempty" xml:"pay_tariff_project,omitempty"`
	// 含税价 业务必填
	PriceIncludingTax *string `json:"price_including_tax,omitempty" xml:"price_including_tax,omitempty"`
	// 结算客户名称
	SettleClient *string `json:"settle_client,omitempty" xml:"settle_client,omitempty"`
	// 结算客户企业信用号
	SettleClientCertNo *string `json:"settle_client_cert_no,omitempty" xml:"settle_client_cert_no,omitempty"`
	// 结算客户did
	SettleClientDid *string `json:"settle_client_did,omitempty" xml:"settle_client_did,omitempty"`
	// 税金
	Taxes *string `json:"taxes,omitempty" xml:"taxes,omitempty"`
	// 未税价 业务必填
	UntaxedPrice *string `json:"untaxed_price,omitempty" xml:"untaxed_price,omitempty"`
	// 确认时间
	ConfirmTime *int64 `json:"confirm_time,omitempty" xml:"confirm_time,omitempty"`
}

func (s SaveBillPaybilltariffRequest) String() string {
	return tea.Prettify(s)
}

func (s SaveBillPaybilltariffRequest) GoString() string {
	return s.String()
}

func (s *SaveBillPaybilltariffRequest) SetAuthToken(v string) *SaveBillPaybilltariffRequest {
	s.AuthToken = &v
	return s
}

func (s *SaveBillPaybilltariffRequest) SetProductInstanceId(v string) *SaveBillPaybilltariffRequest {
	s.ProductInstanceId = &v
	return s
}

func (s *SaveBillPaybilltariffRequest) SetAction(v string) *SaveBillPaybilltariffRequest {
	s.Action = &v
	return s
}

func (s *SaveBillPaybilltariffRequest) SetBillingNumber(v string) *SaveBillPaybilltariffRequest {
	s.BillingNumber = &v
	return s
}

func (s *SaveBillPaybilltariffRequest) SetBillingType(v string) *SaveBillPaybilltariffRequest {
	s.BillingType = &v
	return s
}

func (s *SaveBillPaybilltariffRequest) SetBkgNo(v string) *SaveBillPaybilltariffRequest {
	s.BkgNo = &v
	return s
}

func (s *SaveBillPaybilltariffRequest) SetBookingNo(v string) *SaveBillPaybilltariffRequest {
	s.BookingNo = &v
	return s
}

func (s *SaveBillPaybilltariffRequest) SetCurrency(v string) *SaveBillPaybilltariffRequest {
	s.Currency = &v
	return s
}

func (s *SaveBillPaybilltariffRequest) SetForwarderDid(v string) *SaveBillPaybilltariffRequest {
	s.ForwarderDid = &v
	return s
}

func (s *SaveBillPaybilltariffRequest) SetIncomeOrExpenses(v string) *SaveBillPaybilltariffRequest {
	s.IncomeOrExpenses = &v
	return s
}

func (s *SaveBillPaybilltariffRequest) SetInvoiceUrl(v string) *SaveBillPaybilltariffRequest {
	s.InvoiceUrl = &v
	return s
}

func (s *SaveBillPaybilltariffRequest) SetOrderNo(v string) *SaveBillPaybilltariffRequest {
	s.OrderNo = &v
	return s
}

func (s *SaveBillPaybilltariffRequest) SetPayTariffCode(v string) *SaveBillPaybilltariffRequest {
	s.PayTariffCode = &v
	return s
}

func (s *SaveBillPaybilltariffRequest) SetPayTariffDesc(v string) *SaveBillPaybilltariffRequest {
	s.PayTariffDesc = &v
	return s
}

func (s *SaveBillPaybilltariffRequest) SetPayTariffProject(v string) *SaveBillPaybilltariffRequest {
	s.PayTariffProject = &v
	return s
}

func (s *SaveBillPaybilltariffRequest) SetPriceIncludingTax(v string) *SaveBillPaybilltariffRequest {
	s.PriceIncludingTax = &v
	return s
}

func (s *SaveBillPaybilltariffRequest) SetSettleClient(v string) *SaveBillPaybilltariffRequest {
	s.SettleClient = &v
	return s
}

func (s *SaveBillPaybilltariffRequest) SetSettleClientCertNo(v string) *SaveBillPaybilltariffRequest {
	s.SettleClientCertNo = &v
	return s
}

func (s *SaveBillPaybilltariffRequest) SetSettleClientDid(v string) *SaveBillPaybilltariffRequest {
	s.SettleClientDid = &v
	return s
}

func (s *SaveBillPaybilltariffRequest) SetTaxes(v string) *SaveBillPaybilltariffRequest {
	s.Taxes = &v
	return s
}

func (s *SaveBillPaybilltariffRequest) SetUntaxedPrice(v string) *SaveBillPaybilltariffRequest {
	s.UntaxedPrice = &v
	return s
}

func (s *SaveBillPaybilltariffRequest) SetConfirmTime(v int64) *SaveBillPaybilltariffRequest {
	s.ConfirmTime = &v
	return s
}

type SaveBillPaybilltariffResponse struct {
	// 请求唯一ID，用于链路跟踪和问题排查
	ReqMsgId *string `json:"req_msg_id,omitempty" xml:"req_msg_id,omitempty"`
	// 结果码，一般OK表示调用成功
	ResultCode *string `json:"result_code,omitempty" xml:"result_code,omitempty"`
	// 异常信息的文本描述
	ResultMsg *string `json:"result_msg,omitempty" xml:"result_msg,omitempty"`
	// 链上凭证
	TxCodes []*TxDto `json:"tx_codes,omitempty" xml:"tx_codes,omitempty" type:"Repeated"`
}

func (s SaveBillPaybilltariffResponse) String() string {
	return tea.Prettify(s)
}

func (s SaveBillPaybilltariffResponse) GoString() string {
	return s.String()
}

func (s *SaveBillPaybilltariffResponse) SetReqMsgId(v string) *SaveBillPaybilltariffResponse {
	s.ReqMsgId = &v
	return s
}

func (s *SaveBillPaybilltariffResponse) SetResultCode(v string) *SaveBillPaybilltariffResponse {
	s.ResultCode = &v
	return s
}

func (s *SaveBillPaybilltariffResponse) SetResultMsg(v string) *SaveBillPaybilltariffResponse {
	s.ResultMsg = &v
	return s
}

func (s *SaveBillPaybilltariffResponse) SetTxCodes(v []*TxDto) *SaveBillPaybilltariffResponse {
	s.TxCodes = v
	return s
}

type SaveBillReceiptbilltariffRequest struct {
	// OAuth模式下的授权token
	AuthToken         *string `json:"auth_token,omitempty" xml:"auth_token,omitempty"`
	ProductInstanceId *string `json:"product_instance_id,omitempty" xml:"product_instance_id,omitempty"`
	// 操作动作,为空为新增或更新，UPDATE为更新，DELETE为删除
	Action *string `json:"action,omitempty" xml:"action,omitempty"`
	// 计费数量  业务必填
	BillingNumber *string `json:"billing_number,omitempty" xml:"billing_number,omitempty"`
	// 计费类型 业务必填
	BillingType *string `json:"billing_type,omitempty" xml:"billing_type,omitempty"`
	// 航运订舱号[业务必填]
	BkgNo *string `json:"bkg_no,omitempty" xml:"bkg_no,omitempty"`
	// 订舱单号[业务必填]
	BookingNo *string `json:"booking_no,omitempty" xml:"booking_no,omitempty"`
	// 币种 [业务必填]
	Currency *string `json:"currency,omitempty" xml:"currency,omitempty"`
	// 货代did
	ForwarderDid *string `json:"forwarder_did,omitempty" xml:"forwarder_did,omitempty" require:"true"`
	// 收支  INCOME--收入， EXPENSES--支出
	// 业务必填
	IncomeOrExpenses *string `json:"income_or_expenses,omitempty" xml:"income_or_expenses,omitempty"`
	// 电子发票网址
	InvoiceUrl *string `json:"invoice_url,omitempty" xml:"invoice_url,omitempty"`
	// 工作单号 业务必填
	OrderNo *string `json:"order_no,omitempty" xml:"order_no,omitempty"`
	// 含税价 业务必填
	PriceIncludingTax *string `json:"price_including_tax,omitempty" xml:"price_including_tax,omitempty"`
	// 资费单据编号
	ReceiptTariffCode *string `json:"receipt_tariff_code,omitempty" xml:"receipt_tariff_code,omitempty" require:"true"`
	// 中文描述 业务必填
	ReceiptTariffDesc *string `json:"receipt_tariff_desc,omitempty" xml:"receipt_tariff_desc,omitempty"`
	// 资费项目 业务必填
	ReceiptTariffProject *string `json:"receipt_tariff_project,omitempty" xml:"receipt_tariff_project,omitempty"`
	// 结算客户名称
	SettleClient *string `json:"settle_client,omitempty" xml:"settle_client,omitempty"`
	// 结算客户企业信用号
	SettleClientCertNo *string `json:"settle_client_cert_no,omitempty" xml:"settle_client_cert_no,omitempty"`
	// 结算客户did
	SettleClientDid *string `json:"settle_client_did,omitempty" xml:"settle_client_did,omitempty"`
	// 税金
	Taxes *string `json:"taxes,omitempty" xml:"taxes,omitempty"`
	// 未税价 业务必填
	UntaxedPrice *string `json:"untaxed_price,omitempty" xml:"untaxed_price,omitempty"`
	// 确认时间
	ConfirmTime *int64 `json:"confirm_time,omitempty" xml:"confirm_time,omitempty"`
}

func (s SaveBillReceiptbilltariffRequest) String() string {
	return tea.Prettify(s)
}

func (s SaveBillReceiptbilltariffRequest) GoString() string {
	return s.String()
}

func (s *SaveBillReceiptbilltariffRequest) SetAuthToken(v string) *SaveBillReceiptbilltariffRequest {
	s.AuthToken = &v
	return s
}

func (s *SaveBillReceiptbilltariffRequest) SetProductInstanceId(v string) *SaveBillReceiptbilltariffRequest {
	s.ProductInstanceId = &v
	return s
}

func (s *SaveBillReceiptbilltariffRequest) SetAction(v string) *SaveBillReceiptbilltariffRequest {
	s.Action = &v
	return s
}

func (s *SaveBillReceiptbilltariffRequest) SetBillingNumber(v string) *SaveBillReceiptbilltariffRequest {
	s.BillingNumber = &v
	return s
}

func (s *SaveBillReceiptbilltariffRequest) SetBillingType(v string) *SaveBillReceiptbilltariffRequest {
	s.BillingType = &v
	return s
}

func (s *SaveBillReceiptbilltariffRequest) SetBkgNo(v string) *SaveBillReceiptbilltariffRequest {
	s.BkgNo = &v
	return s
}

func (s *SaveBillReceiptbilltariffRequest) SetBookingNo(v string) *SaveBillReceiptbilltariffRequest {
	s.BookingNo = &v
	return s
}

func (s *SaveBillReceiptbilltariffRequest) SetCurrency(v string) *SaveBillReceiptbilltariffRequest {
	s.Currency = &v
	return s
}

func (s *SaveBillReceiptbilltariffRequest) SetForwarderDid(v string) *SaveBillReceiptbilltariffRequest {
	s.ForwarderDid = &v
	return s
}

func (s *SaveBillReceiptbilltariffRequest) SetIncomeOrExpenses(v string) *SaveBillReceiptbilltariffRequest {
	s.IncomeOrExpenses = &v
	return s
}

func (s *SaveBillReceiptbilltariffRequest) SetInvoiceUrl(v string) *SaveBillReceiptbilltariffRequest {
	s.InvoiceUrl = &v
	return s
}

func (s *SaveBillReceiptbilltariffRequest) SetOrderNo(v string) *SaveBillReceiptbilltariffRequest {
	s.OrderNo = &v
	return s
}

func (s *SaveBillReceiptbilltariffRequest) SetPriceIncludingTax(v string) *SaveBillReceiptbilltariffRequest {
	s.PriceIncludingTax = &v
	return s
}

func (s *SaveBillReceiptbilltariffRequest) SetReceiptTariffCode(v string) *SaveBillReceiptbilltariffRequest {
	s.ReceiptTariffCode = &v
	return s
}

func (s *SaveBillReceiptbilltariffRequest) SetReceiptTariffDesc(v string) *SaveBillReceiptbilltariffRequest {
	s.ReceiptTariffDesc = &v
	return s
}

func (s *SaveBillReceiptbilltariffRequest) SetReceiptTariffProject(v string) *SaveBillReceiptbilltariffRequest {
	s.ReceiptTariffProject = &v
	return s
}

func (s *SaveBillReceiptbilltariffRequest) SetSettleClient(v string) *SaveBillReceiptbilltariffRequest {
	s.SettleClient = &v
	return s
}

func (s *SaveBillReceiptbilltariffRequest) SetSettleClientCertNo(v string) *SaveBillReceiptbilltariffRequest {
	s.SettleClientCertNo = &v
	return s
}

func (s *SaveBillReceiptbilltariffRequest) SetSettleClientDid(v string) *SaveBillReceiptbilltariffRequest {
	s.SettleClientDid = &v
	return s
}

func (s *SaveBillReceiptbilltariffRequest) SetTaxes(v string) *SaveBillReceiptbilltariffRequest {
	s.Taxes = &v
	return s
}

func (s *SaveBillReceiptbilltariffRequest) SetUntaxedPrice(v string) *SaveBillReceiptbilltariffRequest {
	s.UntaxedPrice = &v
	return s
}

func (s *SaveBillReceiptbilltariffRequest) SetConfirmTime(v int64) *SaveBillReceiptbilltariffRequest {
	s.ConfirmTime = &v
	return s
}

type SaveBillReceiptbilltariffResponse struct {
	// 请求唯一ID，用于链路跟踪和问题排查
	ReqMsgId *string `json:"req_msg_id,omitempty" xml:"req_msg_id,omitempty"`
	// 结果码，一般OK表示调用成功
	ResultCode *string `json:"result_code,omitempty" xml:"result_code,omitempty"`
	// 异常信息的文本描述
	ResultMsg *string `json:"result_msg,omitempty" xml:"result_msg,omitempty"`
	// 链上凭证
	TxCodes []*TxDto `json:"tx_codes,omitempty" xml:"tx_codes,omitempty" type:"Repeated"`
}

func (s SaveBillReceiptbilltariffResponse) String() string {
	return tea.Prettify(s)
}

func (s SaveBillReceiptbilltariffResponse) GoString() string {
	return s.String()
}

func (s *SaveBillReceiptbilltariffResponse) SetReqMsgId(v string) *SaveBillReceiptbilltariffResponse {
	s.ReqMsgId = &v
	return s
}

func (s *SaveBillReceiptbilltariffResponse) SetResultCode(v string) *SaveBillReceiptbilltariffResponse {
	s.ResultCode = &v
	return s
}

func (s *SaveBillReceiptbilltariffResponse) SetResultMsg(v string) *SaveBillReceiptbilltariffResponse {
	s.ResultMsg = &v
	return s
}

func (s *SaveBillReceiptbilltariffResponse) SetTxCodes(v []*TxDto) *SaveBillReceiptbilltariffResponse {
	s.TxCodes = v
	return s
}

type VerifyBillPaybillRequest struct {
	// OAuth模式下的授权token
	AuthToken         *string `json:"auth_token,omitempty" xml:"auth_token,omitempty"`
	ProductInstanceId *string `json:"product_instance_id,omitempty" xml:"product_instance_id,omitempty"`
	// 操作动作,为空为新增或更新，UPDATE为更新，DELETE为删除
	Action *string `json:"action,omitempty" xml:"action,omitempty"`
	// 货代did
	ForwarderDid *string `json:"forwarder_did,omitempty" xml:"forwarder_did,omitempty" require:"true"`
	// 应付资费项编号
	PayTariffCode *string `json:"pay_tariff_code,omitempty" xml:"pay_tariff_code,omitempty" require:"true"`
	// 核销金额 业务必填
	VerifyAmount *string `json:"verify_amount,omitempty" xml:"verify_amount,omitempty"`
	// 核销应付资费项编号
	VerifyPayTariffCode *string `json:"verify_pay_tariff_code,omitempty" xml:"verify_pay_tariff_code,omitempty" require:"true"`
	// 核销状态 WAIT_VERIFY-待核销，PART_VERIFY-部分核销，COMPLETE_VERIFY-完成核销
	// 业务必填
	VerifyStatus *string `json:"verify_status,omitempty" xml:"verify_status,omitempty"`
}

func (s VerifyBillPaybillRequest) String() string {
	return tea.Prettify(s)
}

func (s VerifyBillPaybillRequest) GoString() string {
	return s.String()
}

func (s *VerifyBillPaybillRequest) SetAuthToken(v string) *VerifyBillPaybillRequest {
	s.AuthToken = &v
	return s
}

func (s *VerifyBillPaybillRequest) SetProductInstanceId(v string) *VerifyBillPaybillRequest {
	s.ProductInstanceId = &v
	return s
}

func (s *VerifyBillPaybillRequest) SetAction(v string) *VerifyBillPaybillRequest {
	s.Action = &v
	return s
}

func (s *VerifyBillPaybillRequest) SetForwarderDid(v string) *VerifyBillPaybillRequest {
	s.ForwarderDid = &v
	return s
}

func (s *VerifyBillPaybillRequest) SetPayTariffCode(v string) *VerifyBillPaybillRequest {
	s.PayTariffCode = &v
	return s
}

func (s *VerifyBillPaybillRequest) SetVerifyAmount(v string) *VerifyBillPaybillRequest {
	s.VerifyAmount = &v
	return s
}

func (s *VerifyBillPaybillRequest) SetVerifyPayTariffCode(v string) *VerifyBillPaybillRequest {
	s.VerifyPayTariffCode = &v
	return s
}

func (s *VerifyBillPaybillRequest) SetVerifyStatus(v string) *VerifyBillPaybillRequest {
	s.VerifyStatus = &v
	return s
}

type VerifyBillPaybillResponse struct {
	// 请求唯一ID，用于链路跟踪和问题排查
	ReqMsgId *string `json:"req_msg_id,omitempty" xml:"req_msg_id,omitempty"`
	// 结果码，一般OK表示调用成功
	ResultCode *string `json:"result_code,omitempty" xml:"result_code,omitempty"`
	// 异常信息的文本描述
	ResultMsg *string `json:"result_msg,omitempty" xml:"result_msg,omitempty"`
	// 链上凭证
	TxCodes []*TxDto `json:"tx_codes,omitempty" xml:"tx_codes,omitempty" type:"Repeated"`
}

func (s VerifyBillPaybillResponse) String() string {
	return tea.Prettify(s)
}

func (s VerifyBillPaybillResponse) GoString() string {
	return s.String()
}

func (s *VerifyBillPaybillResponse) SetReqMsgId(v string) *VerifyBillPaybillResponse {
	s.ReqMsgId = &v
	return s
}

func (s *VerifyBillPaybillResponse) SetResultCode(v string) *VerifyBillPaybillResponse {
	s.ResultCode = &v
	return s
}

func (s *VerifyBillPaybillResponse) SetResultMsg(v string) *VerifyBillPaybillResponse {
	s.ResultMsg = &v
	return s
}

func (s *VerifyBillPaybillResponse) SetTxCodes(v []*TxDto) *VerifyBillPaybillResponse {
	s.TxCodes = v
	return s
}

type VerifyBillReceiptbillorderRequest struct {
	// OAuth模式下的授权token
	AuthToken         *string `json:"auth_token,omitempty" xml:"auth_token,omitempty"`
	ProductInstanceId *string `json:"product_instance_id,omitempty" xml:"product_instance_id,omitempty"`
	// 操作动作,为空为新增或更新，UPDATE为更新，DELETE为删除
	Action *string `json:"action,omitempty" xml:"action,omitempty"`
	// 货代did
	ForwarderDid *string `json:"forwarder_did,omitempty" xml:"forwarder_did,omitempty" require:"true"`
	// 应收资费项编号
	ReceiptTariffCode *string `json:"receipt_tariff_code,omitempty" xml:"receipt_tariff_code,omitempty" require:"true"`
	// 核销金额 业务必填
	VerifyAmount *string `json:"verify_amount,omitempty" xml:"verify_amount,omitempty"`
	// 核销应收资费项编号
	VerifyReceiptTariffCode *string `json:"verify_receipt_tariff_code,omitempty" xml:"verify_receipt_tariff_code,omitempty" require:"true"`
	// 核销状态 WAIT_VERIFY-待核销，PART_VERIFY-部分核销，COMPLETE_VERIFY-完成核销 业务必填
	VerifyStatus *string `json:"verify_status,omitempty" xml:"verify_status,omitempty"`
}

func (s VerifyBillReceiptbillorderRequest) String() string {
	return tea.Prettify(s)
}

func (s VerifyBillReceiptbillorderRequest) GoString() string {
	return s.String()
}

func (s *VerifyBillReceiptbillorderRequest) SetAuthToken(v string) *VerifyBillReceiptbillorderRequest {
	s.AuthToken = &v
	return s
}

func (s *VerifyBillReceiptbillorderRequest) SetProductInstanceId(v string) *VerifyBillReceiptbillorderRequest {
	s.ProductInstanceId = &v
	return s
}

func (s *VerifyBillReceiptbillorderRequest) SetAction(v string) *VerifyBillReceiptbillorderRequest {
	s.Action = &v
	return s
}

func (s *VerifyBillReceiptbillorderRequest) SetForwarderDid(v string) *VerifyBillReceiptbillorderRequest {
	s.ForwarderDid = &v
	return s
}

func (s *VerifyBillReceiptbillorderRequest) SetReceiptTariffCode(v string) *VerifyBillReceiptbillorderRequest {
	s.ReceiptTariffCode = &v
	return s
}

func (s *VerifyBillReceiptbillorderRequest) SetVerifyAmount(v string) *VerifyBillReceiptbillorderRequest {
	s.VerifyAmount = &v
	return s
}

func (s *VerifyBillReceiptbillorderRequest) SetVerifyReceiptTariffCode(v string) *VerifyBillReceiptbillorderRequest {
	s.VerifyReceiptTariffCode = &v
	return s
}

func (s *VerifyBillReceiptbillorderRequest) SetVerifyStatus(v string) *VerifyBillReceiptbillorderRequest {
	s.VerifyStatus = &v
	return s
}

type VerifyBillReceiptbillorderResponse struct {
	// 请求唯一ID，用于链路跟踪和问题排查
	ReqMsgId *string `json:"req_msg_id,omitempty" xml:"req_msg_id,omitempty"`
	// 结果码，一般OK表示调用成功
	ResultCode *string `json:"result_code,omitempty" xml:"result_code,omitempty"`
	// 异常信息的文本描述
	ResultMsg *string `json:"result_msg,omitempty" xml:"result_msg,omitempty"`
	// 链上凭证
	TxCodes []*TxDto `json:"tx_codes,omitempty" xml:"tx_codes,omitempty" type:"Repeated"`
}

func (s VerifyBillReceiptbillorderResponse) String() string {
	return tea.Prettify(s)
}

func (s VerifyBillReceiptbillorderResponse) GoString() string {
	return s.String()
}

func (s *VerifyBillReceiptbillorderResponse) SetReqMsgId(v string) *VerifyBillReceiptbillorderResponse {
	s.ReqMsgId = &v
	return s
}

func (s *VerifyBillReceiptbillorderResponse) SetResultCode(v string) *VerifyBillReceiptbillorderResponse {
	s.ResultCode = &v
	return s
}

func (s *VerifyBillReceiptbillorderResponse) SetResultMsg(v string) *VerifyBillReceiptbillorderResponse {
	s.ResultMsg = &v
	return s
}

func (s *VerifyBillReceiptbillorderResponse) SetTxCodes(v []*TxDto) *VerifyBillReceiptbillorderResponse {
	s.TxCodes = v
	return s
}

type UpdateBillPaybillorderRequest struct {
	// OAuth模式下的授权token
	AuthToken         *string `json:"auth_token,omitempty" xml:"auth_token,omitempty"`
	ProductInstanceId *string `json:"product_instance_id,omitempty" xml:"product_instance_id,omitempty"`
	// 操作动作,为空为新增或更新，UPDATE为更新，DELETE为删除
	//
	Action *string `json:"action,omitempty" xml:"action,omitempty"`
	// 申请时间 毫秒值 业务必填
	ApplyDate *int64 `json:"apply_date,omitempty" xml:"apply_date,omitempty"`
	// 银行收款账户
	BankReceiptAccount *string `json:"bank_receipt_account,omitempty" xml:"bank_receipt_account,omitempty"`
	// 币种 业务必填
	Currency *string `json:"currency,omitempty" xml:"currency,omitempty"`
	// 开户行
	DepositBank *string `json:"deposit_bank,omitempty" xml:"deposit_bank,omitempty"`
	// 账单到期日
	ExpireDate *int64 `json:"expire_date,omitempty" xml:"expire_date,omitempty"`
	// 货代did
	ForwarderDid *string `json:"forwarder_did,omitempty" xml:"forwarder_did,omitempty" require:"true"`
	// 应付总额 业务必填
	PayAmount *string `json:"pay_amount,omitempty" xml:"pay_amount,omitempty"`
	// 付款单编号
	PayBillOrderCode *string `json:"pay_bill_order_code,omitempty" xml:"pay_bill_order_code,omitempty" require:"true"`
	// 应付账单资费项 业务必填
	PayBillTariffParams []*PayBillTariffParam `json:"pay_bill_tariff_params,omitempty" xml:"pay_bill_tariff_params,omitempty" type:"Repeated"`
	// 付款公司
	PayCompany *string `json:"pay_company,omitempty" xml:"pay_company,omitempty"`
	// 付款公司企业信用号
	PayCompanyCertNo *string `json:"pay_company_cert_no,omitempty" xml:"pay_company_cert_no,omitempty"`
	// 付款公司did
	PayCompanyDid *string `json:"pay_company_did,omitempty" xml:"pay_company_did,omitempty"`
	// 付款期限
	PayDeadline *string `json:"pay_deadline,omitempty" xml:"pay_deadline,omitempty"`
	// 收款客户[业务必填]
	ReceiptClient *string `json:"receipt_client,omitempty" xml:"receipt_client,omitempty"`
	// 收款客户企业信用号[业务必填]
	ReceiptClientCertNo *string `json:"receipt_client_cert_no,omitempty" xml:"receipt_client_cert_no,omitempty"`
	// 收款客户did[业务必填]
	ReceiptClientDid *string `json:"receipt_client_did,omitempty" xml:"receipt_client_did,omitempty"`
	// 对应应收资费项code
	ReceiptTariffCodes []*string `json:"receipt_tariff_codes,omitempty" xml:"receipt_tariff_codes,omitempty" type:"Repeated"`
	// 结算公司
	SettleCompany *string `json:"settle_company,omitempty" xml:"settle_company,omitempty"`
	// 结算公司企业信用号
	SettleCompanyCertNo *string `json:"settle_company_cert_no,omitempty" xml:"settle_company_cert_no,omitempty"`
	// 结算公司did [业务必填]
	SettleCompanyDid *string `json:"settle_company_did,omitempty" xml:"settle_company_did,omitempty"`
	// 账单结算周期
	SettleCycle *string `json:"settle_cycle,omitempty" xml:"settle_cycle,omitempty"`
	// 结算状态 ：  SETTLED(已结算) ,  UNSETTLE（未结算）[业务必填]
	SettleStatus *string `json:"settle_status,omitempty" xml:"settle_status,omitempty"`
	// 确认时间
	ConfirmTime *int64 `json:"confirm_time,omitempty" xml:"confirm_time,omitempty"`
}

func (s UpdateBillPaybillorderRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateBillPaybillorderRequest) GoString() string {
	return s.String()
}

func (s *UpdateBillPaybillorderRequest) SetAuthToken(v string) *UpdateBillPaybillorderRequest {
	s.AuthToken = &v
	return s
}

func (s *UpdateBillPaybillorderRequest) SetProductInstanceId(v string) *UpdateBillPaybillorderRequest {
	s.ProductInstanceId = &v
	return s
}

func (s *UpdateBillPaybillorderRequest) SetAction(v string) *UpdateBillPaybillorderRequest {
	s.Action = &v
	return s
}

func (s *UpdateBillPaybillorderRequest) SetApplyDate(v int64) *UpdateBillPaybillorderRequest {
	s.ApplyDate = &v
	return s
}

func (s *UpdateBillPaybillorderRequest) SetBankReceiptAccount(v string) *UpdateBillPaybillorderRequest {
	s.BankReceiptAccount = &v
	return s
}

func (s *UpdateBillPaybillorderRequest) SetCurrency(v string) *UpdateBillPaybillorderRequest {
	s.Currency = &v
	return s
}

func (s *UpdateBillPaybillorderRequest) SetDepositBank(v string) *UpdateBillPaybillorderRequest {
	s.DepositBank = &v
	return s
}

func (s *UpdateBillPaybillorderRequest) SetExpireDate(v int64) *UpdateBillPaybillorderRequest {
	s.ExpireDate = &v
	return s
}

func (s *UpdateBillPaybillorderRequest) SetForwarderDid(v string) *UpdateBillPaybillorderRequest {
	s.ForwarderDid = &v
	return s
}

func (s *UpdateBillPaybillorderRequest) SetPayAmount(v string) *UpdateBillPaybillorderRequest {
	s.PayAmount = &v
	return s
}

func (s *UpdateBillPaybillorderRequest) SetPayBillOrderCode(v string) *UpdateBillPaybillorderRequest {
	s.PayBillOrderCode = &v
	return s
}

func (s *UpdateBillPaybillorderRequest) SetPayBillTariffParams(v []*PayBillTariffParam) *UpdateBillPaybillorderRequest {
	s.PayBillTariffParams = v
	return s
}

func (s *UpdateBillPaybillorderRequest) SetPayCompany(v string) *UpdateBillPaybillorderRequest {
	s.PayCompany = &v
	return s
}

func (s *UpdateBillPaybillorderRequest) SetPayCompanyCertNo(v string) *UpdateBillPaybillorderRequest {
	s.PayCompanyCertNo = &v
	return s
}

func (s *UpdateBillPaybillorderRequest) SetPayCompanyDid(v string) *UpdateBillPaybillorderRequest {
	s.PayCompanyDid = &v
	return s
}

func (s *UpdateBillPaybillorderRequest) SetPayDeadline(v string) *UpdateBillPaybillorderRequest {
	s.PayDeadline = &v
	return s
}

func (s *UpdateBillPaybillorderRequest) SetReceiptClient(v string) *UpdateBillPaybillorderRequest {
	s.ReceiptClient = &v
	return s
}

func (s *UpdateBillPaybillorderRequest) SetReceiptClientCertNo(v string) *UpdateBillPaybillorderRequest {
	s.ReceiptClientCertNo = &v
	return s
}

func (s *UpdateBillPaybillorderRequest) SetReceiptClientDid(v string) *UpdateBillPaybillorderRequest {
	s.ReceiptClientDid = &v
	return s
}

func (s *UpdateBillPaybillorderRequest) SetReceiptTariffCodes(v []*string) *UpdateBillPaybillorderRequest {
	s.ReceiptTariffCodes = v
	return s
}

func (s *UpdateBillPaybillorderRequest) SetSettleCompany(v string) *UpdateBillPaybillorderRequest {
	s.SettleCompany = &v
	return s
}

func (s *UpdateBillPaybillorderRequest) SetSettleCompanyCertNo(v string) *UpdateBillPaybillorderRequest {
	s.SettleCompanyCertNo = &v
	return s
}

func (s *UpdateBillPaybillorderRequest) SetSettleCompanyDid(v string) *UpdateBillPaybillorderRequest {
	s.SettleCompanyDid = &v
	return s
}

func (s *UpdateBillPaybillorderRequest) SetSettleCycle(v string) *UpdateBillPaybillorderRequest {
	s.SettleCycle = &v
	return s
}

func (s *UpdateBillPaybillorderRequest) SetSettleStatus(v string) *UpdateBillPaybillorderRequest {
	s.SettleStatus = &v
	return s
}

func (s *UpdateBillPaybillorderRequest) SetConfirmTime(v int64) *UpdateBillPaybillorderRequest {
	s.ConfirmTime = &v
	return s
}

type UpdateBillPaybillorderResponse struct {
	// 请求唯一ID，用于链路跟踪和问题排查
	ReqMsgId *string `json:"req_msg_id,omitempty" xml:"req_msg_id,omitempty"`
	// 结果码，一般OK表示调用成功
	ResultCode *string `json:"result_code,omitempty" xml:"result_code,omitempty"`
	// 异常信息的文本描述
	ResultMsg *string `json:"result_msg,omitempty" xml:"result_msg,omitempty"`
	// 链上凭证
	TxCodes []*TxDto `json:"tx_codes,omitempty" xml:"tx_codes,omitempty" type:"Repeated"`
}

func (s UpdateBillPaybillorderResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateBillPaybillorderResponse) GoString() string {
	return s.String()
}

func (s *UpdateBillPaybillorderResponse) SetReqMsgId(v string) *UpdateBillPaybillorderResponse {
	s.ReqMsgId = &v
	return s
}

func (s *UpdateBillPaybillorderResponse) SetResultCode(v string) *UpdateBillPaybillorderResponse {
	s.ResultCode = &v
	return s
}

func (s *UpdateBillPaybillorderResponse) SetResultMsg(v string) *UpdateBillPaybillorderResponse {
	s.ResultMsg = &v
	return s
}

func (s *UpdateBillPaybillorderResponse) SetTxCodes(v []*TxDto) *UpdateBillPaybillorderResponse {
	s.TxCodes = v
	return s
}

type UpdateBillReceiptbillorderRequest struct {
	// OAuth模式下的授权token
	AuthToken         *string `json:"auth_token,omitempty" xml:"auth_token,omitempty"`
	ProductInstanceId *string `json:"product_instance_id,omitempty" xml:"product_instance_id,omitempty"`
	// 操作动作,为空为新增或更新，UPDATE为更新，DELETE为删除
	Action *string `json:"action,omitempty" xml:"action,omitempty"`
	// 实际收款结算公司
	ActualSettleCompany *string `json:"actual_settle_company,omitempty" xml:"actual_settle_company,omitempty"`
	// 实际收款结算公司企业信用号
	ActualSettleCompanyCertNo *string `json:"actual_settle_company_cert_no,omitempty" xml:"actual_settle_company_cert_no,omitempty"`
	// 实际收款结算公司did [业务必填]
	ActualSettleCompanyDid *string `json:"actual_settle_company_did,omitempty" xml:"actual_settle_company_did,omitempty"`
	// 确认时间
	ApplyDate *int64 `json:"apply_date,omitempty" xml:"apply_date,omitempty"`
	// 揽货类型
	ClientType *string `json:"client_type,omitempty" xml:"client_type,omitempty"`
	// 币种
	Currency *string `json:"currency,omitempty" xml:"currency,omitempty"`
	// 账单到期日
	ExpireDate *int64 `json:"expire_date,omitempty" xml:"expire_date,omitempty"`
	// 货代did
	ForwarderDid *string `json:"forwarder_did,omitempty" xml:"forwarder_did,omitempty" require:"true"`
	// 开票抬头
	InvoiceTitle *string `json:"invoice_title,omitempty" xml:"invoice_title,omitempty"`
	// 含税价
	PriceIncludingTax *string `json:"price_including_tax,omitempty" xml:"price_including_tax,omitempty"`
	// 收款账号
	ReceiptAccount *string `json:"receipt_account,omitempty" xml:"receipt_account,omitempty"`
	// 收款总额
	ReceiptAmount *string `json:"receipt_amount,omitempty" xml:"receipt_amount,omitempty"`
	// 收款账单编号
	ReceiptBillOrderCode *string `json:"receipt_bill_order_code,omitempty" xml:"receipt_bill_order_code,omitempty" require:"true"`
	// 应收资费项账单 业务必填
	ReceiptBillTariffParams []*ReceiptBillTariffParam `json:"receipt_bill_tariff_params,omitempty" xml:"receipt_bill_tariff_params,omitempty" type:"Repeated"`
	// 结算客户名称
	SettleClient *string `json:"settle_client,omitempty" xml:"settle_client,omitempty"`
	// 结算客户名称企业信用号
	SettleClientCertNo *string `json:"settle_client_cert_no,omitempty" xml:"settle_client_cert_no,omitempty"`
	// 结算客户did
	SettleClientDid *string `json:"settle_client_did,omitempty" xml:"settle_client_did,omitempty"`
	// 结算公司 [业务必填]
	SettleCompany *string `json:"settle_company,omitempty" xml:"settle_company,omitempty"`
	// 结算公司企业信用号[业务必填]
	SettleCompanyCertNo *string `json:"settle_company_cert_no,omitempty" xml:"settle_company_cert_no,omitempty"`
	// 结算公司did[业务必填]
	SettleCompanyDid *string `json:"settle_company_did,omitempty" xml:"settle_company_did,omitempty"`
	// 账单结算周期
	SettleCycle *string `json:"settle_cycle,omitempty" xml:"settle_cycle,omitempty"`
	// SETTLED(已结算) UNSETTLE（未结算）[业务必填]
	SettleStatus *string `json:"settle_status,omitempty" xml:"settle_status,omitempty"`
	// 税金
	Taxes *string `json:"taxes,omitempty" xml:"taxes,omitempty"`
	// 未税价
	UntaxedPrice *string `json:"untaxed_price,omitempty" xml:"untaxed_price,omitempty"`
	// 确认时间
	ConfirmTime *int64 `json:"confirm_time,omitempty" xml:"confirm_time,omitempty"`
}

func (s UpdateBillReceiptbillorderRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateBillReceiptbillorderRequest) GoString() string {
	return s.String()
}

func (s *UpdateBillReceiptbillorderRequest) SetAuthToken(v string) *UpdateBillReceiptbillorderRequest {
	s.AuthToken = &v
	return s
}

func (s *UpdateBillReceiptbillorderRequest) SetProductInstanceId(v string) *UpdateBillReceiptbillorderRequest {
	s.ProductInstanceId = &v
	return s
}

func (s *UpdateBillReceiptbillorderRequest) SetAction(v string) *UpdateBillReceiptbillorderRequest {
	s.Action = &v
	return s
}

func (s *UpdateBillReceiptbillorderRequest) SetActualSettleCompany(v string) *UpdateBillReceiptbillorderRequest {
	s.ActualSettleCompany = &v
	return s
}

func (s *UpdateBillReceiptbillorderRequest) SetActualSettleCompanyCertNo(v string) *UpdateBillReceiptbillorderRequest {
	s.ActualSettleCompanyCertNo = &v
	return s
}

func (s *UpdateBillReceiptbillorderRequest) SetActualSettleCompanyDid(v string) *UpdateBillReceiptbillorderRequest {
	s.ActualSettleCompanyDid = &v
	return s
}

func (s *UpdateBillReceiptbillorderRequest) SetApplyDate(v int64) *UpdateBillReceiptbillorderRequest {
	s.ApplyDate = &v
	return s
}

func (s *UpdateBillReceiptbillorderRequest) SetClientType(v string) *UpdateBillReceiptbillorderRequest {
	s.ClientType = &v
	return s
}

func (s *UpdateBillReceiptbillorderRequest) SetCurrency(v string) *UpdateBillReceiptbillorderRequest {
	s.Currency = &v
	return s
}

func (s *UpdateBillReceiptbillorderRequest) SetExpireDate(v int64) *UpdateBillReceiptbillorderRequest {
	s.ExpireDate = &v
	return s
}

func (s *UpdateBillReceiptbillorderRequest) SetForwarderDid(v string) *UpdateBillReceiptbillorderRequest {
	s.ForwarderDid = &v
	return s
}

func (s *UpdateBillReceiptbillorderRequest) SetInvoiceTitle(v string) *UpdateBillReceiptbillorderRequest {
	s.InvoiceTitle = &v
	return s
}

func (s *UpdateBillReceiptbillorderRequest) SetPriceIncludingTax(v string) *UpdateBillReceiptbillorderRequest {
	s.PriceIncludingTax = &v
	return s
}

func (s *UpdateBillReceiptbillorderRequest) SetReceiptAccount(v string) *UpdateBillReceiptbillorderRequest {
	s.ReceiptAccount = &v
	return s
}

func (s *UpdateBillReceiptbillorderRequest) SetReceiptAmount(v string) *UpdateBillReceiptbillorderRequest {
	s.ReceiptAmount = &v
	return s
}

func (s *UpdateBillReceiptbillorderRequest) SetReceiptBillOrderCode(v string) *UpdateBillReceiptbillorderRequest {
	s.ReceiptBillOrderCode = &v
	return s
}

func (s *UpdateBillReceiptbillorderRequest) SetReceiptBillTariffParams(v []*ReceiptBillTariffParam) *UpdateBillReceiptbillorderRequest {
	s.ReceiptBillTariffParams = v
	return s
}

func (s *UpdateBillReceiptbillorderRequest) SetSettleClient(v string) *UpdateBillReceiptbillorderRequest {
	s.SettleClient = &v
	return s
}

func (s *UpdateBillReceiptbillorderRequest) SetSettleClientCertNo(v string) *UpdateBillReceiptbillorderRequest {
	s.SettleClientCertNo = &v
	return s
}

func (s *UpdateBillReceiptbillorderRequest) SetSettleClientDid(v string) *UpdateBillReceiptbillorderRequest {
	s.SettleClientDid = &v
	return s
}

func (s *UpdateBillReceiptbillorderRequest) SetSettleCompany(v string) *UpdateBillReceiptbillorderRequest {
	s.SettleCompany = &v
	return s
}

func (s *UpdateBillReceiptbillorderRequest) SetSettleCompanyCertNo(v string) *UpdateBillReceiptbillorderRequest {
	s.SettleCompanyCertNo = &v
	return s
}

func (s *UpdateBillReceiptbillorderRequest) SetSettleCompanyDid(v string) *UpdateBillReceiptbillorderRequest {
	s.SettleCompanyDid = &v
	return s
}

func (s *UpdateBillReceiptbillorderRequest) SetSettleCycle(v string) *UpdateBillReceiptbillorderRequest {
	s.SettleCycle = &v
	return s
}

func (s *UpdateBillReceiptbillorderRequest) SetSettleStatus(v string) *UpdateBillReceiptbillorderRequest {
	s.SettleStatus = &v
	return s
}

func (s *UpdateBillReceiptbillorderRequest) SetTaxes(v string) *UpdateBillReceiptbillorderRequest {
	s.Taxes = &v
	return s
}

func (s *UpdateBillReceiptbillorderRequest) SetUntaxedPrice(v string) *UpdateBillReceiptbillorderRequest {
	s.UntaxedPrice = &v
	return s
}

func (s *UpdateBillReceiptbillorderRequest) SetConfirmTime(v int64) *UpdateBillReceiptbillorderRequest {
	s.ConfirmTime = &v
	return s
}

type UpdateBillReceiptbillorderResponse struct {
	// 请求唯一ID，用于链路跟踪和问题排查
	ReqMsgId *string `json:"req_msg_id,omitempty" xml:"req_msg_id,omitempty"`
	// 结果码，一般OK表示调用成功
	ResultCode *string `json:"result_code,omitempty" xml:"result_code,omitempty"`
	// 异常信息的文本描述
	ResultMsg *string `json:"result_msg,omitempty" xml:"result_msg,omitempty"`
	// 链上凭证
	TxCodes []*TxDto `json:"tx_codes,omitempty" xml:"tx_codes,omitempty" type:"Repeated"`
}

func (s UpdateBillReceiptbillorderResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateBillReceiptbillorderResponse) GoString() string {
	return s.String()
}

func (s *UpdateBillReceiptbillorderResponse) SetReqMsgId(v string) *UpdateBillReceiptbillorderResponse {
	s.ReqMsgId = &v
	return s
}

func (s *UpdateBillReceiptbillorderResponse) SetResultCode(v string) *UpdateBillReceiptbillorderResponse {
	s.ResultCode = &v
	return s
}

func (s *UpdateBillReceiptbillorderResponse) SetResultMsg(v string) *UpdateBillReceiptbillorderResponse {
	s.ResultMsg = &v
	return s
}

func (s *UpdateBillReceiptbillorderResponse) SetTxCodes(v []*TxDto) *UpdateBillReceiptbillorderResponse {
	s.TxCodes = v
	return s
}

type SaveBillPayinvoiceRequest struct {
	// OAuth模式下的授权token
	AuthToken         *string `json:"auth_token,omitempty" xml:"auth_token,omitempty"`
	ProductInstanceId *string `json:"product_instance_id,omitempty" xml:"product_instance_id,omitempty"`
	// 操作动作,为空为新增或更新，UPDATE为更新，DELETE为删除
	Action *string `json:"action,omitempty" xml:"action,omitempty"`
	// 总金额 业务必填
	Amount *string `json:"amount,omitempty" xml:"amount,omitempty"`
	// 校验码后六位
	CheckCode *string `json:"check_code,omitempty" xml:"check_code,omitempty"`
	// 币种 CNY/USD [业务必填]
	Currency *string `json:"currency,omitempty" xml:"currency,omitempty"`
	// 美元金额
	DollarAmount *string `json:"dollar_amount,omitempty" xml:"dollar_amount,omitempty"`
	// 开票方名称  业务必填
	DrawerName *string `json:"drawer_name,omitempty" xml:"drawer_name,omitempty"`
	// 开票纳税人识别号 业务必填
	DrawerTaxpayerCode *string `json:"drawer_taxpayer_code,omitempty" xml:"drawer_taxpayer_code,omitempty"`
	// 货代did
	ForwarderDid *string `json:"forwarder_did,omitempty" xml:"forwarder_did,omitempty" require:"true"`
	// 发票唯一标识
	InvoiceCode *string `json:"invoice_code,omitempty" xml:"invoice_code,omitempty" require:"true"`
	// 发票抬头did[业务必填]
	InvoiceHeaderDid *string `json:"invoice_header_did,omitempty" xml:"invoice_header_did,omitempty"`
	// 发票抬头企业名称[业务必填]
	InvoiceHeaderName *string `json:"invoice_header_name,omitempty" xml:"invoice_header_name,omitempty"`
	// 发票抬头企业信用代码[业务必填]
	InvoiceHeaderSocialNo *string `json:"invoice_header_social_no,omitempty" xml:"invoice_header_social_no,omitempty"`
	// 发票号码 业务必填
	InvoiceNumber *string `json:"invoice_number,omitempty" xml:"invoice_number,omitempty"`
	// 发票税务号
	InvoiceTaxCode *string `json:"invoice_tax_code,omitempty" xml:"invoice_tax_code,omitempty"`
	// 开票类型 业务必填
	InvoiceType *string `json:"invoice_type,omitempty" xml:"invoice_type,omitempty"`
	// 开票日期 业务必填
	MakeInvoiceDate *int64 `json:"make_invoice_date,omitempty" xml:"make_invoice_date,omitempty"`
	// 账单关联项目
	PayBillInvoiceParams []*PayBillInvoiceParam `json:"pay_bill_invoice_params,omitempty" xml:"pay_bill_invoice_params,omitempty" type:"Repeated"`
	// 资费项发票 业务必填
	PayTariffInvoiceParams []*PayTariffInvoiceParam `json:"pay_tariff_invoice_params,omitempty" xml:"pay_tariff_invoice_params,omitempty" type:"Repeated"`
	// 不含税金额 业务必填
	UntaxedPrice *string `json:"untaxed_price,omitempty" xml:"untaxed_price,omitempty"`
}

func (s SaveBillPayinvoiceRequest) String() string {
	return tea.Prettify(s)
}

func (s SaveBillPayinvoiceRequest) GoString() string {
	return s.String()
}

func (s *SaveBillPayinvoiceRequest) SetAuthToken(v string) *SaveBillPayinvoiceRequest {
	s.AuthToken = &v
	return s
}

func (s *SaveBillPayinvoiceRequest) SetProductInstanceId(v string) *SaveBillPayinvoiceRequest {
	s.ProductInstanceId = &v
	return s
}

func (s *SaveBillPayinvoiceRequest) SetAction(v string) *SaveBillPayinvoiceRequest {
	s.Action = &v
	return s
}

func (s *SaveBillPayinvoiceRequest) SetAmount(v string) *SaveBillPayinvoiceRequest {
	s.Amount = &v
	return s
}

func (s *SaveBillPayinvoiceRequest) SetCheckCode(v string) *SaveBillPayinvoiceRequest {
	s.CheckCode = &v
	return s
}

func (s *SaveBillPayinvoiceRequest) SetCurrency(v string) *SaveBillPayinvoiceRequest {
	s.Currency = &v
	return s
}

func (s *SaveBillPayinvoiceRequest) SetDollarAmount(v string) *SaveBillPayinvoiceRequest {
	s.DollarAmount = &v
	return s
}

func (s *SaveBillPayinvoiceRequest) SetDrawerName(v string) *SaveBillPayinvoiceRequest {
	s.DrawerName = &v
	return s
}

func (s *SaveBillPayinvoiceRequest) SetDrawerTaxpayerCode(v string) *SaveBillPayinvoiceRequest {
	s.DrawerTaxpayerCode = &v
	return s
}

func (s *SaveBillPayinvoiceRequest) SetForwarderDid(v string) *SaveBillPayinvoiceRequest {
	s.ForwarderDid = &v
	return s
}

func (s *SaveBillPayinvoiceRequest) SetInvoiceCode(v string) *SaveBillPayinvoiceRequest {
	s.InvoiceCode = &v
	return s
}

func (s *SaveBillPayinvoiceRequest) SetInvoiceHeaderDid(v string) *SaveBillPayinvoiceRequest {
	s.InvoiceHeaderDid = &v
	return s
}

func (s *SaveBillPayinvoiceRequest) SetInvoiceHeaderName(v string) *SaveBillPayinvoiceRequest {
	s.InvoiceHeaderName = &v
	return s
}

func (s *SaveBillPayinvoiceRequest) SetInvoiceHeaderSocialNo(v string) *SaveBillPayinvoiceRequest {
	s.InvoiceHeaderSocialNo = &v
	return s
}

func (s *SaveBillPayinvoiceRequest) SetInvoiceNumber(v string) *SaveBillPayinvoiceRequest {
	s.InvoiceNumber = &v
	return s
}

func (s *SaveBillPayinvoiceRequest) SetInvoiceTaxCode(v string) *SaveBillPayinvoiceRequest {
	s.InvoiceTaxCode = &v
	return s
}

func (s *SaveBillPayinvoiceRequest) SetInvoiceType(v string) *SaveBillPayinvoiceRequest {
	s.InvoiceType = &v
	return s
}

func (s *SaveBillPayinvoiceRequest) SetMakeInvoiceDate(v int64) *SaveBillPayinvoiceRequest {
	s.MakeInvoiceDate = &v
	return s
}

func (s *SaveBillPayinvoiceRequest) SetPayBillInvoiceParams(v []*PayBillInvoiceParam) *SaveBillPayinvoiceRequest {
	s.PayBillInvoiceParams = v
	return s
}

func (s *SaveBillPayinvoiceRequest) SetPayTariffInvoiceParams(v []*PayTariffInvoiceParam) *SaveBillPayinvoiceRequest {
	s.PayTariffInvoiceParams = v
	return s
}

func (s *SaveBillPayinvoiceRequest) SetUntaxedPrice(v string) *SaveBillPayinvoiceRequest {
	s.UntaxedPrice = &v
	return s
}

type SaveBillPayinvoiceResponse struct {
	// 请求唯一ID，用于链路跟踪和问题排查
	ReqMsgId *string `json:"req_msg_id,omitempty" xml:"req_msg_id,omitempty"`
	// 结果码，一般OK表示调用成功
	ResultCode *string `json:"result_code,omitempty" xml:"result_code,omitempty"`
	// 异常信息的文本描述
	ResultMsg *string `json:"result_msg,omitempty" xml:"result_msg,omitempty"`
	// 链上凭证
	TxCodes []*TxDto `json:"tx_codes,omitempty" xml:"tx_codes,omitempty" type:"Repeated"`
}

func (s SaveBillPayinvoiceResponse) String() string {
	return tea.Prettify(s)
}

func (s SaveBillPayinvoiceResponse) GoString() string {
	return s.String()
}

func (s *SaveBillPayinvoiceResponse) SetReqMsgId(v string) *SaveBillPayinvoiceResponse {
	s.ReqMsgId = &v
	return s
}

func (s *SaveBillPayinvoiceResponse) SetResultCode(v string) *SaveBillPayinvoiceResponse {
	s.ResultCode = &v
	return s
}

func (s *SaveBillPayinvoiceResponse) SetResultMsg(v string) *SaveBillPayinvoiceResponse {
	s.ResultMsg = &v
	return s
}

func (s *SaveBillPayinvoiceResponse) SetTxCodes(v []*TxDto) *SaveBillPayinvoiceResponse {
	s.TxCodes = v
	return s
}

type SaveBillReceiptinvoiceRequest struct {
	// OAuth模式下的授权token
	AuthToken         *string `json:"auth_token,omitempty" xml:"auth_token,omitempty"`
	ProductInstanceId *string `json:"product_instance_id,omitempty" xml:"product_instance_id,omitempty"`
	// 操作动作,为空为新增或更新，UPDATE为更新，DELETE为删除
	Action *string `json:"action,omitempty" xml:"action,omitempty"`
	// 总金额 业务必填
	Amount *string `json:"amount,omitempty" xml:"amount,omitempty"`
	// 校验码后六位
	CheckCode *string `json:"check_code,omitempty" xml:"check_code,omitempty"`
	// 开票方名称 业务必填
	DrawerName *string `json:"drawer_name,omitempty" xml:"drawer_name,omitempty"`
	// 开票纳税人识别号 业务必填
	DrawerTaxpayerCode *string `json:"drawer_taxpayer_code,omitempty" xml:"drawer_taxpayer_code,omitempty"`
	// 货代did
	ForwarderDid *string `json:"forwarder_did,omitempty" xml:"forwarder_did,omitempty" require:"true"`
	// 应收发票唯一标识
	InvoiceCode *string `json:"invoice_code,omitempty" xml:"invoice_code,omitempty" require:"true"`
	// 发票号码 业务必填
	InvoiceNumber *string `json:"invoice_number,omitempty" xml:"invoice_number,omitempty"`
	// 开票类型 业务必填
	InvoiceType *string `json:"invoice_type,omitempty" xml:"invoice_type,omitempty"`
	// 开票日期 毫秒值 业务必填
	MakeInvoiceDate *int64 `json:"make_invoice_date,omitempty" xml:"make_invoice_date,omitempty"`
	// 资费项发票 业务必填
	ReceiptTariffInvoiceParams []*ReceiptTariffInvoiceParam `json:"receipt_tariff_invoice_params,omitempty" xml:"receipt_tariff_invoice_params,omitempty" type:"Repeated"`
	// 不含税金额 业务必填
	UntaxedPrice *string `json:"untaxed_price,omitempty" xml:"untaxed_price,omitempty"`
	// 发票税务号
	InvoiceTaxCode *string `json:"invoice_tax_code,omitempty" xml:"invoice_tax_code,omitempty"`
}

func (s SaveBillReceiptinvoiceRequest) String() string {
	return tea.Prettify(s)
}

func (s SaveBillReceiptinvoiceRequest) GoString() string {
	return s.String()
}

func (s *SaveBillReceiptinvoiceRequest) SetAuthToken(v string) *SaveBillReceiptinvoiceRequest {
	s.AuthToken = &v
	return s
}

func (s *SaveBillReceiptinvoiceRequest) SetProductInstanceId(v string) *SaveBillReceiptinvoiceRequest {
	s.ProductInstanceId = &v
	return s
}

func (s *SaveBillReceiptinvoiceRequest) SetAction(v string) *SaveBillReceiptinvoiceRequest {
	s.Action = &v
	return s
}

func (s *SaveBillReceiptinvoiceRequest) SetAmount(v string) *SaveBillReceiptinvoiceRequest {
	s.Amount = &v
	return s
}

func (s *SaveBillReceiptinvoiceRequest) SetCheckCode(v string) *SaveBillReceiptinvoiceRequest {
	s.CheckCode = &v
	return s
}

func (s *SaveBillReceiptinvoiceRequest) SetDrawerName(v string) *SaveBillReceiptinvoiceRequest {
	s.DrawerName = &v
	return s
}

func (s *SaveBillReceiptinvoiceRequest) SetDrawerTaxpayerCode(v string) *SaveBillReceiptinvoiceRequest {
	s.DrawerTaxpayerCode = &v
	return s
}

func (s *SaveBillReceiptinvoiceRequest) SetForwarderDid(v string) *SaveBillReceiptinvoiceRequest {
	s.ForwarderDid = &v
	return s
}

func (s *SaveBillReceiptinvoiceRequest) SetInvoiceCode(v string) *SaveBillReceiptinvoiceRequest {
	s.InvoiceCode = &v
	return s
}

func (s *SaveBillReceiptinvoiceRequest) SetInvoiceNumber(v string) *SaveBillReceiptinvoiceRequest {
	s.InvoiceNumber = &v
	return s
}

func (s *SaveBillReceiptinvoiceRequest) SetInvoiceType(v string) *SaveBillReceiptinvoiceRequest {
	s.InvoiceType = &v
	return s
}

func (s *SaveBillReceiptinvoiceRequest) SetMakeInvoiceDate(v int64) *SaveBillReceiptinvoiceRequest {
	s.MakeInvoiceDate = &v
	return s
}

func (s *SaveBillReceiptinvoiceRequest) SetReceiptTariffInvoiceParams(v []*ReceiptTariffInvoiceParam) *SaveBillReceiptinvoiceRequest {
	s.ReceiptTariffInvoiceParams = v
	return s
}

func (s *SaveBillReceiptinvoiceRequest) SetUntaxedPrice(v string) *SaveBillReceiptinvoiceRequest {
	s.UntaxedPrice = &v
	return s
}

func (s *SaveBillReceiptinvoiceRequest) SetInvoiceTaxCode(v string) *SaveBillReceiptinvoiceRequest {
	s.InvoiceTaxCode = &v
	return s
}

type SaveBillReceiptinvoiceResponse struct {
	// 请求唯一ID，用于链路跟踪和问题排查
	ReqMsgId *string `json:"req_msg_id,omitempty" xml:"req_msg_id,omitempty"`
	// 结果码，一般OK表示调用成功
	ResultCode *string `json:"result_code,omitempty" xml:"result_code,omitempty"`
	// 异常信息的文本描述
	ResultMsg *string `json:"result_msg,omitempty" xml:"result_msg,omitempty"`
	// 链上凭证
	TxCodes []*TxDto `json:"tx_codes,omitempty" xml:"tx_codes,omitempty" type:"Repeated"`
}

func (s SaveBillReceiptinvoiceResponse) String() string {
	return tea.Prettify(s)
}

func (s SaveBillReceiptinvoiceResponse) GoString() string {
	return s.String()
}

func (s *SaveBillReceiptinvoiceResponse) SetReqMsgId(v string) *SaveBillReceiptinvoiceResponse {
	s.ReqMsgId = &v
	return s
}

func (s *SaveBillReceiptinvoiceResponse) SetResultCode(v string) *SaveBillReceiptinvoiceResponse {
	s.ResultCode = &v
	return s
}

func (s *SaveBillReceiptinvoiceResponse) SetResultMsg(v string) *SaveBillReceiptinvoiceResponse {
	s.ResultMsg = &v
	return s
}

func (s *SaveBillReceiptinvoiceResponse) SetTxCodes(v []*TxDto) *SaveBillReceiptinvoiceResponse {
	s.TxCodes = v
	return s
}

type UploadBizFinancingRequest struct {
	// OAuth模式下的授权token
	AuthToken         *string `json:"auth_token,omitempty" xml:"auth_token,omitempty"`
	ProductInstanceId *string `json:"product_instance_id,omitempty" xml:"product_instance_id,omitempty"`
	// 融资数据
	FinancingData []*UploadFinancingParam `json:"financing_data,omitempty" xml:"financing_data,omitempty" require:"true" type:"Repeated"`
}

func (s UploadBizFinancingRequest) String() string {
	return tea.Prettify(s)
}

func (s UploadBizFinancingRequest) GoString() string {
	return s.String()
}

func (s *UploadBizFinancingRequest) SetAuthToken(v string) *UploadBizFinancingRequest {
	s.AuthToken = &v
	return s
}

func (s *UploadBizFinancingRequest) SetProductInstanceId(v string) *UploadBizFinancingRequest {
	s.ProductInstanceId = &v
	return s
}

func (s *UploadBizFinancingRequest) SetFinancingData(v []*UploadFinancingParam) *UploadBizFinancingRequest {
	s.FinancingData = v
	return s
}

type UploadBizFinancingResponse struct {
	// 请求唯一ID，用于链路跟踪和问题排查
	ReqMsgId *string `json:"req_msg_id,omitempty" xml:"req_msg_id,omitempty"`
	// 结果码，一般OK表示调用成功
	ResultCode *string `json:"result_code,omitempty" xml:"result_code,omitempty"`
	// 异常信息的文本描述
	ResultMsg *string `json:"result_msg,omitempty" xml:"result_msg,omitempty"`
	// 上链hash
	TxCodes []*TxDto `json:"tx_codes,omitempty" xml:"tx_codes,omitempty" type:"Repeated"`
}

func (s UploadBizFinancingResponse) String() string {
	return tea.Prettify(s)
}

func (s UploadBizFinancingResponse) GoString() string {
	return s.String()
}

func (s *UploadBizFinancingResponse) SetReqMsgId(v string) *UploadBizFinancingResponse {
	s.ReqMsgId = &v
	return s
}

func (s *UploadBizFinancingResponse) SetResultCode(v string) *UploadBizFinancingResponse {
	s.ResultCode = &v
	return s
}

func (s *UploadBizFinancingResponse) SetResultMsg(v string) *UploadBizFinancingResponse {
	s.ResultMsg = &v
	return s
}

func (s *UploadBizFinancingResponse) SetTxCodes(v []*TxDto) *UploadBizFinancingResponse {
	s.TxCodes = v
	return s
}

type UploadBizOrderRequest struct {
	// OAuth模式下的授权token
	AuthToken         *string `json:"auth_token,omitempty" xml:"auth_token,omitempty"`
	ProductInstanceId *string `json:"product_instance_id,omitempty" xml:"product_instance_id,omitempty"`
	// 船公司did
	CarrierDid *string `json:"carrier_did,omitempty" xml:"carrier_did,omitempty" require:"true"`
	// 尾款金额
	DeadlineAmount *string `json:"deadline_amount,omitempty" xml:"deadline_amount,omitempty" require:"true"`
	// 货代did
	ForwarderDid *string `json:"forwarder_did,omitempty" xml:"forwarder_did,omitempty" require:"true"`
	// 订单总额
	OrderAmounts []*UploadOrderAmount `json:"order_amounts,omitempty" xml:"order_amounts,omitempty" require:"true" type:"Repeated"`
	// 订单booking信息
	OrderBookings []*UploadOrderBooking `json:"order_bookings,omitempty" xml:"order_bookings,omitempty" require:"true" type:"Repeated"`
	// 订单号
	OrderNo *string `json:"order_no,omitempty" xml:"order_no,omitempty" require:"true"`
	// 结算did
	SettleDid *string `json:"settle_did,omitempty" xml:"settle_did,omitempty" require:"true"`
}

func (s UploadBizOrderRequest) String() string {
	return tea.Prettify(s)
}

func (s UploadBizOrderRequest) GoString() string {
	return s.String()
}

func (s *UploadBizOrderRequest) SetAuthToken(v string) *UploadBizOrderRequest {
	s.AuthToken = &v
	return s
}

func (s *UploadBizOrderRequest) SetProductInstanceId(v string) *UploadBizOrderRequest {
	s.ProductInstanceId = &v
	return s
}

func (s *UploadBizOrderRequest) SetCarrierDid(v string) *UploadBizOrderRequest {
	s.CarrierDid = &v
	return s
}

func (s *UploadBizOrderRequest) SetDeadlineAmount(v string) *UploadBizOrderRequest {
	s.DeadlineAmount = &v
	return s
}

func (s *UploadBizOrderRequest) SetForwarderDid(v string) *UploadBizOrderRequest {
	s.ForwarderDid = &v
	return s
}

func (s *UploadBizOrderRequest) SetOrderAmounts(v []*UploadOrderAmount) *UploadBizOrderRequest {
	s.OrderAmounts = v
	return s
}

func (s *UploadBizOrderRequest) SetOrderBookings(v []*UploadOrderBooking) *UploadBizOrderRequest {
	s.OrderBookings = v
	return s
}

func (s *UploadBizOrderRequest) SetOrderNo(v string) *UploadBizOrderRequest {
	s.OrderNo = &v
	return s
}

func (s *UploadBizOrderRequest) SetSettleDid(v string) *UploadBizOrderRequest {
	s.SettleDid = &v
	return s
}

type UploadBizOrderResponse struct {
	// 请求唯一ID，用于链路跟踪和问题排查
	ReqMsgId *string `json:"req_msg_id,omitempty" xml:"req_msg_id,omitempty"`
	// 结果码，一般OK表示调用成功
	ResultCode *string `json:"result_code,omitempty" xml:"result_code,omitempty"`
	// 异常信息的文本描述
	ResultMsg *string `json:"result_msg,omitempty" xml:"result_msg,omitempty"`
	// 上链hash
	TxCodes []*TxDto `json:"tx_codes,omitempty" xml:"tx_codes,omitempty" type:"Repeated"`
}

func (s UploadBizOrderResponse) String() string {
	return tea.Prettify(s)
}

func (s UploadBizOrderResponse) GoString() string {
	return s.String()
}

func (s *UploadBizOrderResponse) SetReqMsgId(v string) *UploadBizOrderResponse {
	s.ReqMsgId = &v
	return s
}

func (s *UploadBizOrderResponse) SetResultCode(v string) *UploadBizOrderResponse {
	s.ResultCode = &v
	return s
}

func (s *UploadBizOrderResponse) SetResultMsg(v string) *UploadBizOrderResponse {
	s.ResultMsg = &v
	return s
}

func (s *UploadBizOrderResponse) SetTxCodes(v []*TxDto) *UploadBizOrderResponse {
	s.TxCodes = v
	return s
}

type CreateDidCarrierRequest struct {
	// OAuth模式下的授权token
	AuthToken         *string `json:"auth_token,omitempty" xml:"auth_token,omitempty"`
	ProductInstanceId *string `json:"product_instance_id,omitempty" xml:"product_instance_id,omitempty"`
	// 代理did
	AgentDid *string `json:"agent_did,omitempty" xml:"agent_did,omitempty" require:"true"`
	// 企业名称
	EpCertName *string `json:"ep_cert_name,omitempty" xml:"ep_cert_name,omitempty" require:"true"`
	// 企业证件号
	EpCertNo *string `json:"ep_cert_no,omitempty" xml:"ep_cert_no,omitempty" require:"true"`
	// 扩展字段
	ExtensionInfo *string `json:"extension_info,omitempty" xml:"extension_info,omitempty"`
	// 法人姓名
	LegalPersonCertName *string `json:"legal_person_cert_name,omitempty" xml:"legal_person_cert_name,omitempty"`
	// 法人身份证
	LegalPersonCertNo *string `json:"legal_person_cert_no,omitempty" xml:"legal_person_cert_no,omitempty"`
	// 船公司编号
	Scac *string `json:"scac,omitempty" xml:"scac,omitempty"`
}

func (s CreateDidCarrierRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateDidCarrierRequest) GoString() string {
	return s.String()
}

func (s *CreateDidCarrierRequest) SetAuthToken(v string) *CreateDidCarrierRequest {
	s.AuthToken = &v
	return s
}

func (s *CreateDidCarrierRequest) SetProductInstanceId(v string) *CreateDidCarrierRequest {
	s.ProductInstanceId = &v
	return s
}

func (s *CreateDidCarrierRequest) SetAgentDid(v string) *CreateDidCarrierRequest {
	s.AgentDid = &v
	return s
}

func (s *CreateDidCarrierRequest) SetEpCertName(v string) *CreateDidCarrierRequest {
	s.EpCertName = &v
	return s
}

func (s *CreateDidCarrierRequest) SetEpCertNo(v string) *CreateDidCarrierRequest {
	s.EpCertNo = &v
	return s
}

func (s *CreateDidCarrierRequest) SetExtensionInfo(v string) *CreateDidCarrierRequest {
	s.ExtensionInfo = &v
	return s
}

func (s *CreateDidCarrierRequest) SetLegalPersonCertName(v string) *CreateDidCarrierRequest {
	s.LegalPersonCertName = &v
	return s
}

func (s *CreateDidCarrierRequest) SetLegalPersonCertNo(v string) *CreateDidCarrierRequest {
	s.LegalPersonCertNo = &v
	return s
}

func (s *CreateDidCarrierRequest) SetScac(v string) *CreateDidCarrierRequest {
	s.Scac = &v
	return s
}

type CreateDidCarrierResponse struct {
	// 请求唯一ID，用于链路跟踪和问题排查
	ReqMsgId *string `json:"req_msg_id,omitempty" xml:"req_msg_id,omitempty"`
	// 结果码，一般OK表示调用成功
	ResultCode *string `json:"result_code,omitempty" xml:"result_code,omitempty"`
	// 异常信息的文本描述
	ResultMsg *string `json:"result_msg,omitempty" xml:"result_msg,omitempty"`
	// 船公司did
	Did *string `json:"did,omitempty" xml:"did,omitempty"`
}

func (s CreateDidCarrierResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateDidCarrierResponse) GoString() string {
	return s.String()
}

func (s *CreateDidCarrierResponse) SetReqMsgId(v string) *CreateDidCarrierResponse {
	s.ReqMsgId = &v
	return s
}

func (s *CreateDidCarrierResponse) SetResultCode(v string) *CreateDidCarrierResponse {
	s.ResultCode = &v
	return s
}

func (s *CreateDidCarrierResponse) SetResultMsg(v string) *CreateDidCarrierResponse {
	s.ResultMsg = &v
	return s
}

func (s *CreateDidCarrierResponse) SetDid(v string) *CreateDidCarrierResponse {
	s.Did = &v
	return s
}

type AuthSysForwarderRequest struct {
	// OAuth模式下的授权token
	AuthToken         *string `json:"auth_token,omitempty" xml:"auth_token,omitempty"`
	ProductInstanceId *string `json:"product_instance_id,omitempty" xml:"product_instance_id,omitempty"`
	// 客户端id
	ClientId *string `json:"client_id,omitempty" xml:"client_id,omitempty" require:"true"`
	// 密钥
	ClientSecret *string `json:"client_secret,omitempty" xml:"client_secret,omitempty" require:"true"`
	// 货代did
	IntRefId *string `json:"int_ref_id,omitempty" xml:"int_ref_id,omitempty" require:"true"`
	// For HOME application internal
	TenantId *string `json:"tenant_id,omitempty" xml:"tenant_id,omitempty" require:"true"`
	// 渠道source
	Source *string `json:"source,omitempty" xml:"source,omitempty" require:"true"`
}

func (s AuthSysForwarderRequest) String() string {
	return tea.Prettify(s)
}

func (s AuthSysForwarderRequest) GoString() string {
	return s.String()
}

func (s *AuthSysForwarderRequest) SetAuthToken(v string) *AuthSysForwarderRequest {
	s.AuthToken = &v
	return s
}

func (s *AuthSysForwarderRequest) SetProductInstanceId(v string) *AuthSysForwarderRequest {
	s.ProductInstanceId = &v
	return s
}

func (s *AuthSysForwarderRequest) SetClientId(v string) *AuthSysForwarderRequest {
	s.ClientId = &v
	return s
}

func (s *AuthSysForwarderRequest) SetClientSecret(v string) *AuthSysForwarderRequest {
	s.ClientSecret = &v
	return s
}

func (s *AuthSysForwarderRequest) SetIntRefId(v string) *AuthSysForwarderRequest {
	s.IntRefId = &v
	return s
}

func (s *AuthSysForwarderRequest) SetTenantId(v string) *AuthSysForwarderRequest {
	s.TenantId = &v
	return s
}

func (s *AuthSysForwarderRequest) SetSource(v string) *AuthSysForwarderRequest {
	s.Source = &v
	return s
}

type AuthSysForwarderResponse struct {
	// 请求唯一ID，用于链路跟踪和问题排查
	ReqMsgId *string `json:"req_msg_id,omitempty" xml:"req_msg_id,omitempty"`
	// 结果码，一般OK表示调用成功
	ResultCode *string `json:"result_code,omitempty" xml:"result_code,omitempty"`
	// 异常信息的文本描述
	ResultMsg *string `json:"result_msg,omitempty" xml:"result_msg,omitempty"`
	// 推送结果
	Result *bool `json:"result,omitempty" xml:"result,omitempty"`
}

func (s AuthSysForwarderResponse) String() string {
	return tea.Prettify(s)
}

func (s AuthSysForwarderResponse) GoString() string {
	return s.String()
}

func (s *AuthSysForwarderResponse) SetReqMsgId(v string) *AuthSysForwarderResponse {
	s.ReqMsgId = &v
	return s
}

func (s *AuthSysForwarderResponse) SetResultCode(v string) *AuthSysForwarderResponse {
	s.ResultCode = &v
	return s
}

func (s *AuthSysForwarderResponse) SetResultMsg(v string) *AuthSysForwarderResponse {
	s.ResultMsg = &v
	return s
}

func (s *AuthSysForwarderResponse) SetResult(v bool) *AuthSysForwarderResponse {
	s.Result = &v
	return s
}

type QueryBizMasterblRequest struct {
	// OAuth模式下的授权token
	AuthToken         *string `json:"auth_token,omitempty" xml:"auth_token,omitempty"`
	ProductInstanceId *string `json:"product_instance_id,omitempty" xml:"product_instance_id,omitempty"`
	// 货代did
	ForwarderDid *string `json:"forwarder_did,omitempty" xml:"forwarder_did,omitempty" require:"true"`
	// master提单号
	MasterBlNo *string `json:"master_bl_no,omitempty" xml:"master_bl_no,omitempty" require:"true"`
}

func (s QueryBizMasterblRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryBizMasterblRequest) GoString() string {
	return s.String()
}

func (s *QueryBizMasterblRequest) SetAuthToken(v string) *QueryBizMasterblRequest {
	s.AuthToken = &v
	return s
}

func (s *QueryBizMasterblRequest) SetProductInstanceId(v string) *QueryBizMasterblRequest {
	s.ProductInstanceId = &v
	return s
}

func (s *QueryBizMasterblRequest) SetForwarderDid(v string) *QueryBizMasterblRequest {
	s.ForwarderDid = &v
	return s
}

func (s *QueryBizMasterblRequest) SetMasterBlNo(v string) *QueryBizMasterblRequest {
	s.MasterBlNo = &v
	return s
}

type QueryBizMasterblResponse struct {
	// 请求唯一ID，用于链路跟踪和问题排查
	ReqMsgId *string `json:"req_msg_id,omitempty" xml:"req_msg_id,omitempty"`
	// 结果码，一般OK表示调用成功
	ResultCode *string `json:"result_code,omitempty" xml:"result_code,omitempty"`
	// 异常信息的文本描述
	ResultMsg *string `json:"result_msg,omitempty" xml:"result_msg,omitempty"`
	// 船公司
	Carrier *string `json:"carrier,omitempty" xml:"carrier,omitempty"`
	// 收货人
	Consignee *string `json:"consignee,omitempty" xml:"consignee,omitempty"`
	// 目的地
	DeliveryPlace *string `json:"delivery_place,omitempty" xml:"delivery_place,omitempty"`
	// 卸货港
	DischargePort *string `json:"discharge_port,omitempty" xml:"discharge_port,omitempty"`
	// 货代did
	ForwarderDid *string `json:"forwarder_did,omitempty" xml:"forwarder_did,omitempty"`
	// 运费
	Freight *string `json:"freight,omitempty" xml:"freight,omitempty"`
	// 提单货物数据列表
	GoodsDtos []*MasterBlGoodsDto `json:"goods_dtos,omitempty" xml:"goods_dtos,omitempty" type:"Repeated"`
	// 起运港
	LoadingPort *string `json:"loading_port,omitempty" xml:"loading_port,omitempty"`
	// master提单号
	MasterBlNo *string `json:"master_bl_no,omitempty" xml:"master_bl_no,omitempty"`
	// 通知方
	NotifyParty *string `json:"notify_party,omitempty" xml:"notify_party,omitempty"`
	// 前程运输
	PrCarriage *string `json:"pr_carriage,omitempty" xml:"pr_carriage,omitempty"`
	// 发货人
	Shipper *string `json:"shipper,omitempty" xml:"shipper,omitempty"`
	// 航名
	Vessel *string `json:"vessel,omitempty" xml:"vessel,omitempty"`
	// 航次
	Voyage *string `json:"voyage,omitempty" xml:"voyage,omitempty"`
}

func (s QueryBizMasterblResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryBizMasterblResponse) GoString() string {
	return s.String()
}

func (s *QueryBizMasterblResponse) SetReqMsgId(v string) *QueryBizMasterblResponse {
	s.ReqMsgId = &v
	return s
}

func (s *QueryBizMasterblResponse) SetResultCode(v string) *QueryBizMasterblResponse {
	s.ResultCode = &v
	return s
}

func (s *QueryBizMasterblResponse) SetResultMsg(v string) *QueryBizMasterblResponse {
	s.ResultMsg = &v
	return s
}

func (s *QueryBizMasterblResponse) SetCarrier(v string) *QueryBizMasterblResponse {
	s.Carrier = &v
	return s
}

func (s *QueryBizMasterblResponse) SetConsignee(v string) *QueryBizMasterblResponse {
	s.Consignee = &v
	return s
}

func (s *QueryBizMasterblResponse) SetDeliveryPlace(v string) *QueryBizMasterblResponse {
	s.DeliveryPlace = &v
	return s
}

func (s *QueryBizMasterblResponse) SetDischargePort(v string) *QueryBizMasterblResponse {
	s.DischargePort = &v
	return s
}

func (s *QueryBizMasterblResponse) SetForwarderDid(v string) *QueryBizMasterblResponse {
	s.ForwarderDid = &v
	return s
}

func (s *QueryBizMasterblResponse) SetFreight(v string) *QueryBizMasterblResponse {
	s.Freight = &v
	return s
}

func (s *QueryBizMasterblResponse) SetGoodsDtos(v []*MasterBlGoodsDto) *QueryBizMasterblResponse {
	s.GoodsDtos = v
	return s
}

func (s *QueryBizMasterblResponse) SetLoadingPort(v string) *QueryBizMasterblResponse {
	s.LoadingPort = &v
	return s
}

func (s *QueryBizMasterblResponse) SetMasterBlNo(v string) *QueryBizMasterblResponse {
	s.MasterBlNo = &v
	return s
}

func (s *QueryBizMasterblResponse) SetNotifyParty(v string) *QueryBizMasterblResponse {
	s.NotifyParty = &v
	return s
}

func (s *QueryBizMasterblResponse) SetPrCarriage(v string) *QueryBizMasterblResponse {
	s.PrCarriage = &v
	return s
}

func (s *QueryBizMasterblResponse) SetShipper(v string) *QueryBizMasterblResponse {
	s.Shipper = &v
	return s
}

func (s *QueryBizMasterblResponse) SetVessel(v string) *QueryBizMasterblResponse {
	s.Vessel = &v
	return s
}

func (s *QueryBizMasterblResponse) SetVoyage(v string) *QueryBizMasterblResponse {
	s.Voyage = &v
	return s
}

type SaveBizPayinvoicefileRequest struct {
	// OAuth模式下的授权token
	AuthToken         *string `json:"auth_token,omitempty" xml:"auth_token,omitempty"`
	ProductInstanceId *string `json:"product_instance_id,omitempty" xml:"product_instance_id,omitempty"`
	// 操作动作,为空为新增或更新，UPDATE为更新，DELETE为删除，INSERT为新增
	Action *string `json:"action,omitempty" xml:"action,omitempty"`
	// 发票文件hash [业务必填]
	FileHash *string `json:"file_hash,omitempty" xml:"file_hash,omitempty"`
	// 发票文件ID [业务必填]
	FileId *string `json:"file_id,omitempty" xml:"file_id,omitempty"`
	// 货代DID
	ForwarderDid *string `json:"forwarder_did,omitempty" xml:"forwarder_did,omitempty" require:"true"`
	// 应付发票code
	InvoiceCode *string `json:"invoice_code,omitempty" xml:"invoice_code,omitempty" require:"true"`
}

func (s SaveBizPayinvoicefileRequest) String() string {
	return tea.Prettify(s)
}

func (s SaveBizPayinvoicefileRequest) GoString() string {
	return s.String()
}

func (s *SaveBizPayinvoicefileRequest) SetAuthToken(v string) *SaveBizPayinvoicefileRequest {
	s.AuthToken = &v
	return s
}

func (s *SaveBizPayinvoicefileRequest) SetProductInstanceId(v string) *SaveBizPayinvoicefileRequest {
	s.ProductInstanceId = &v
	return s
}

func (s *SaveBizPayinvoicefileRequest) SetAction(v string) *SaveBizPayinvoicefileRequest {
	s.Action = &v
	return s
}

func (s *SaveBizPayinvoicefileRequest) SetFileHash(v string) *SaveBizPayinvoicefileRequest {
	s.FileHash = &v
	return s
}

func (s *SaveBizPayinvoicefileRequest) SetFileId(v string) *SaveBizPayinvoicefileRequest {
	s.FileId = &v
	return s
}

func (s *SaveBizPayinvoicefileRequest) SetForwarderDid(v string) *SaveBizPayinvoicefileRequest {
	s.ForwarderDid = &v
	return s
}

func (s *SaveBizPayinvoicefileRequest) SetInvoiceCode(v string) *SaveBizPayinvoicefileRequest {
	s.InvoiceCode = &v
	return s
}

type SaveBizPayinvoicefileResponse struct {
	// 请求唯一ID，用于链路跟踪和问题排查
	ReqMsgId *string `json:"req_msg_id,omitempty" xml:"req_msg_id,omitempty"`
	// 结果码，一般OK表示调用成功
	ResultCode *string `json:"result_code,omitempty" xml:"result_code,omitempty"`
	// 异常信息的文本描述
	ResultMsg *string `json:"result_msg,omitempty" xml:"result_msg,omitempty"`
	// 链上凭证
	TxCodes []*TxDto `json:"tx_codes,omitempty" xml:"tx_codes,omitempty" type:"Repeated"`
}

func (s SaveBizPayinvoicefileResponse) String() string {
	return tea.Prettify(s)
}

func (s SaveBizPayinvoicefileResponse) GoString() string {
	return s.String()
}

func (s *SaveBizPayinvoicefileResponse) SetReqMsgId(v string) *SaveBizPayinvoicefileResponse {
	s.ReqMsgId = &v
	return s
}

func (s *SaveBizPayinvoicefileResponse) SetResultCode(v string) *SaveBizPayinvoicefileResponse {
	s.ResultCode = &v
	return s
}

func (s *SaveBizPayinvoicefileResponse) SetResultMsg(v string) *SaveBizPayinvoicefileResponse {
	s.ResultMsg = &v
	return s
}

func (s *SaveBizPayinvoicefileResponse) SetTxCodes(v []*TxDto) *SaveBizPayinvoicefileResponse {
	s.TxCodes = v
	return s
}

type SaveBiznewOrderRequest struct {
	// OAuth模式下的授权token
	AuthToken         *string `json:"auth_token,omitempty" xml:"auth_token,omitempty"`
	ProductInstanceId *string `json:"product_instance_id,omitempty" xml:"product_instance_id,omitempty"`
	// 托运订单号
	OrderNo *string `json:"order_no,omitempty" xml:"order_no,omitempty" require:"true"`
	// 托运人did [业务必填]
	ConsignorDid *string `json:"consignor_did,omitempty" xml:"consignor_did,omitempty"`
	// 承运人did [业务必填]
	CarrierDid *string `json:"carrier_did,omitempty" xml:"carrier_did,omitempty"`
	// 订舱单号列表
	BookingInfoList []*BookingNoInfo `json:"booking_info_list,omitempty" xml:"booking_info_list,omitempty" type:"Repeated"`
	// 拖车单号
	VehicleCodeList []*string `json:"vehicle_code_list,omitempty" xml:"vehicle_code_list,omitempty" type:"Repeated"`
	// 报关单号
	CustomsCodeList []*string `json:"customs_code_list,omitempty" xml:"customs_code_list,omitempty" type:"Repeated"`
	// 订单确认 CREATE/FINISH (创建/确认)  [业务必填]
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
}

func (s SaveBiznewOrderRequest) String() string {
	return tea.Prettify(s)
}

func (s SaveBiznewOrderRequest) GoString() string {
	return s.String()
}

func (s *SaveBiznewOrderRequest) SetAuthToken(v string) *SaveBiznewOrderRequest {
	s.AuthToken = &v
	return s
}

func (s *SaveBiznewOrderRequest) SetProductInstanceId(v string) *SaveBiznewOrderRequest {
	s.ProductInstanceId = &v
	return s
}

func (s *SaveBiznewOrderRequest) SetOrderNo(v string) *SaveBiznewOrderRequest {
	s.OrderNo = &v
	return s
}

func (s *SaveBiznewOrderRequest) SetConsignorDid(v string) *SaveBiznewOrderRequest {
	s.ConsignorDid = &v
	return s
}

func (s *SaveBiznewOrderRequest) SetCarrierDid(v string) *SaveBiznewOrderRequest {
	s.CarrierDid = &v
	return s
}

func (s *SaveBiznewOrderRequest) SetBookingInfoList(v []*BookingNoInfo) *SaveBiznewOrderRequest {
	s.BookingInfoList = v
	return s
}

func (s *SaveBiznewOrderRequest) SetVehicleCodeList(v []*string) *SaveBiznewOrderRequest {
	s.VehicleCodeList = v
	return s
}

func (s *SaveBiznewOrderRequest) SetCustomsCodeList(v []*string) *SaveBiznewOrderRequest {
	s.CustomsCodeList = v
	return s
}

func (s *SaveBiznewOrderRequest) SetStatus(v string) *SaveBiznewOrderRequest {
	s.Status = &v
	return s
}

type SaveBiznewOrderResponse struct {
	// 请求唯一ID，用于链路跟踪和问题排查
	ReqMsgId *string `json:"req_msg_id,omitempty" xml:"req_msg_id,omitempty"`
	// 结果码，一般OK表示调用成功
	ResultCode *string `json:"result_code,omitempty" xml:"result_code,omitempty"`
	// 异常信息的文本描述
	ResultMsg *string `json:"result_msg,omitempty" xml:"result_msg,omitempty"`
	// 链上凭证
	TxCodes []*TxDto `json:"tx_codes,omitempty" xml:"tx_codes,omitempty" type:"Repeated"`
}

func (s SaveBiznewOrderResponse) String() string {
	return tea.Prettify(s)
}

func (s SaveBiznewOrderResponse) GoString() string {
	return s.String()
}

func (s *SaveBiznewOrderResponse) SetReqMsgId(v string) *SaveBiznewOrderResponse {
	s.ReqMsgId = &v
	return s
}

func (s *SaveBiznewOrderResponse) SetResultCode(v string) *SaveBiznewOrderResponse {
	s.ResultCode = &v
	return s
}

func (s *SaveBiznewOrderResponse) SetResultMsg(v string) *SaveBiznewOrderResponse {
	s.ResultMsg = &v
	return s
}

func (s *SaveBiznewOrderResponse) SetTxCodes(v []*TxDto) *SaveBiznewOrderResponse {
	s.TxCodes = v
	return s
}

type SaveBiznewBookingRequest struct {
	// OAuth模式下的授权token
	AuthToken         *string `json:"auth_token,omitempty" xml:"auth_token,omitempty"`
	ProductInstanceId *string `json:"product_instance_id,omitempty" xml:"product_instance_id,omitempty"`
	// 托运订单号
	//
	//
	OrderNo *string `json:"order_no,omitempty" xml:"order_no,omitempty" require:"true"`
	// 订舱单唯一标识
	BookingNo *string `json:"booking_no,omitempty" xml:"booking_no,omitempty" require:"true"`
	// 订舱号 [业务必填]
	BkgNo *string `json:"bkg_no,omitempty" xml:"bkg_no,omitempty"`
	// 货代did
	ForwarderDid *string `json:"forwarder_did,omitempty" xml:"forwarder_did,omitempty" require:"true"`
	// 收货人
	Consignee *string `json:"consignee,omitempty" xml:"consignee,omitempty"`
	// 船公司 [业务必填]
	Carrier *string `json:"carrier,omitempty" xml:"carrier,omitempty"`
	// 船名 [业务必填]
	Vessel *string `json:"vessel,omitempty" xml:"vessel,omitempty"`
	// 航次 [业务必填]
	Voyage *string `json:"voyage,omitempty" xml:"voyage,omitempty"`
	// 起运港 [业务必填]
	LoadingPort *string `json:"loading_port,omitempty" xml:"loading_port,omitempty"`
	// 卸货港 [业务必填]
	DischargePort *string `json:"discharge_port,omitempty" xml:"discharge_port,omitempty"`
	// 箱型箱量 [业务必填]
	ContainerTypeInfoList []*ContainerTypeInfo `json:"container_type_info_list,omitempty" xml:"container_type_info_list,omitempty" type:"Repeated"`
	// 截关时间（秒时间戳）
	CustomsClearance *int64 `json:"customs_clearance,omitempty" xml:"customs_clearance,omitempty"`
	// 截港时间（秒时间戳）
	CyClosing *int64 `json:"cy_closing,omitempty" xml:"cy_closing,omitempty"`
	// 截单时间 (秒时间戳)
	SiClosing *int64 `json:"si_closing,omitempty" xml:"si_closing,omitempty"`
	// 预计船期（秒时间戳）
	Etd *int64 `json:"etd,omitempty" xml:"etd,omitempty"`
	// 订舱费总金额 [业务必填]
	BkgTotalAmount *string `json:"bkg_total_amount,omitempty" xml:"bkg_total_amount,omitempty"`
	// 订舱费金额 [业务必填]
	BkgAmount *string `json:"bkg_amount,omitempty" xml:"bkg_amount,omitempty"`
	// 港杂费金额 [业务必填]
	PortCharges *string `json:"port_charges,omitempty" xml:"port_charges,omitempty"`
	// 币种 CNY/USD [业务必填]
	//
	//
	Currency *string `json:"currency,omitempty" xml:"currency,omitempty"`
	// 订单确认 CREATE/FINISH (创建/确认) [业务必填]
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
}

func (s SaveBiznewBookingRequest) String() string {
	return tea.Prettify(s)
}

func (s SaveBiznewBookingRequest) GoString() string {
	return s.String()
}

func (s *SaveBiznewBookingRequest) SetAuthToken(v string) *SaveBiznewBookingRequest {
	s.AuthToken = &v
	return s
}

func (s *SaveBiznewBookingRequest) SetProductInstanceId(v string) *SaveBiznewBookingRequest {
	s.ProductInstanceId = &v
	return s
}

func (s *SaveBiznewBookingRequest) SetOrderNo(v string) *SaveBiznewBookingRequest {
	s.OrderNo = &v
	return s
}

func (s *SaveBiznewBookingRequest) SetBookingNo(v string) *SaveBiznewBookingRequest {
	s.BookingNo = &v
	return s
}

func (s *SaveBiznewBookingRequest) SetBkgNo(v string) *SaveBiznewBookingRequest {
	s.BkgNo = &v
	return s
}

func (s *SaveBiznewBookingRequest) SetForwarderDid(v string) *SaveBiznewBookingRequest {
	s.ForwarderDid = &v
	return s
}

func (s *SaveBiznewBookingRequest) SetConsignee(v string) *SaveBiznewBookingRequest {
	s.Consignee = &v
	return s
}

func (s *SaveBiznewBookingRequest) SetCarrier(v string) *SaveBiznewBookingRequest {
	s.Carrier = &v
	return s
}

func (s *SaveBiznewBookingRequest) SetVessel(v string) *SaveBiznewBookingRequest {
	s.Vessel = &v
	return s
}

func (s *SaveBiznewBookingRequest) SetVoyage(v string) *SaveBiznewBookingRequest {
	s.Voyage = &v
	return s
}

func (s *SaveBiznewBookingRequest) SetLoadingPort(v string) *SaveBiznewBookingRequest {
	s.LoadingPort = &v
	return s
}

func (s *SaveBiznewBookingRequest) SetDischargePort(v string) *SaveBiznewBookingRequest {
	s.DischargePort = &v
	return s
}

func (s *SaveBiznewBookingRequest) SetContainerTypeInfoList(v []*ContainerTypeInfo) *SaveBiznewBookingRequest {
	s.ContainerTypeInfoList = v
	return s
}

func (s *SaveBiznewBookingRequest) SetCustomsClearance(v int64) *SaveBiznewBookingRequest {
	s.CustomsClearance = &v
	return s
}

func (s *SaveBiznewBookingRequest) SetCyClosing(v int64) *SaveBiznewBookingRequest {
	s.CyClosing = &v
	return s
}

func (s *SaveBiznewBookingRequest) SetSiClosing(v int64) *SaveBiznewBookingRequest {
	s.SiClosing = &v
	return s
}

func (s *SaveBiznewBookingRequest) SetEtd(v int64) *SaveBiznewBookingRequest {
	s.Etd = &v
	return s
}

func (s *SaveBiznewBookingRequest) SetBkgTotalAmount(v string) *SaveBiznewBookingRequest {
	s.BkgTotalAmount = &v
	return s
}

func (s *SaveBiznewBookingRequest) SetBkgAmount(v string) *SaveBiznewBookingRequest {
	s.BkgAmount = &v
	return s
}

func (s *SaveBiznewBookingRequest) SetPortCharges(v string) *SaveBiznewBookingRequest {
	s.PortCharges = &v
	return s
}

func (s *SaveBiznewBookingRequest) SetCurrency(v string) *SaveBiznewBookingRequest {
	s.Currency = &v
	return s
}

func (s *SaveBiznewBookingRequest) SetStatus(v string) *SaveBiznewBookingRequest {
	s.Status = &v
	return s
}

type SaveBiznewBookingResponse struct {
	// 请求唯一ID，用于链路跟踪和问题排查
	ReqMsgId *string `json:"req_msg_id,omitempty" xml:"req_msg_id,omitempty"`
	// 结果码，一般OK表示调用成功
	ResultCode *string `json:"result_code,omitempty" xml:"result_code,omitempty"`
	// 异常信息的文本描述
	ResultMsg *string `json:"result_msg,omitempty" xml:"result_msg,omitempty"`
	// 链上存证
	TxCodes []*TxDto `json:"tx_codes,omitempty" xml:"tx_codes,omitempty" type:"Repeated"`
}

func (s SaveBiznewBookingResponse) String() string {
	return tea.Prettify(s)
}

func (s SaveBiznewBookingResponse) GoString() string {
	return s.String()
}

func (s *SaveBiznewBookingResponse) SetReqMsgId(v string) *SaveBiznewBookingResponse {
	s.ReqMsgId = &v
	return s
}

func (s *SaveBiznewBookingResponse) SetResultCode(v string) *SaveBiznewBookingResponse {
	s.ResultCode = &v
	return s
}

func (s *SaveBiznewBookingResponse) SetResultMsg(v string) *SaveBiznewBookingResponse {
	s.ResultMsg = &v
	return s
}

func (s *SaveBiznewBookingResponse) SetTxCodes(v []*TxDto) *SaveBiznewBookingResponse {
	s.TxCodes = v
	return s
}

type SaveBiznewVehicleRequest struct {
	// OAuth模式下的授权token
	AuthToken         *string `json:"auth_token,omitempty" xml:"auth_token,omitempty"`
	ProductInstanceId *string `json:"product_instance_id,omitempty" xml:"product_instance_id,omitempty"`
	// 托运订单号
	//
	//
	OrderNo *string `json:"order_no,omitempty" xml:"order_no,omitempty" require:"true"`
	// 拖车单号
	//
	//
	VehicleCode *string `json:"vehicle_code,omitempty" xml:"vehicle_code,omitempty" require:"true"`
	// 货代did
	ForwarderDid *string `json:"forwarder_did,omitempty" xml:"forwarder_did,omitempty" require:"true"`
	// 拖车总金额 [业务必填]
	VehicleTotalAmount *string `json:"vehicle_total_amount,omitempty" xml:"vehicle_total_amount,omitempty"`
	// 拖车运费金额 [业务必填]
	VehicleAmount *string `json:"vehicle_amount,omitempty" xml:"vehicle_amount,omitempty"`
	// 拖车杂费金额 [业务必填]
	VehicleCharges *string `json:"vehicle_charges,omitempty" xml:"vehicle_charges,omitempty"`
	// 币种 [业务必填]
	Currency *string `json:"currency,omitempty" xml:"currency,omitempty"`
	// 箱号箱ID
	//
	//
	ContainerIdInfoList []*ContainerIdInfo `json:"container_id_info_list,omitempty" xml:"container_id_info_list,omitempty" type:"Repeated"`
}

func (s SaveBiznewVehicleRequest) String() string {
	return tea.Prettify(s)
}

func (s SaveBiznewVehicleRequest) GoString() string {
	return s.String()
}

func (s *SaveBiznewVehicleRequest) SetAuthToken(v string) *SaveBiznewVehicleRequest {
	s.AuthToken = &v
	return s
}

func (s *SaveBiznewVehicleRequest) SetProductInstanceId(v string) *SaveBiznewVehicleRequest {
	s.ProductInstanceId = &v
	return s
}

func (s *SaveBiznewVehicleRequest) SetOrderNo(v string) *SaveBiznewVehicleRequest {
	s.OrderNo = &v
	return s
}

func (s *SaveBiznewVehicleRequest) SetVehicleCode(v string) *SaveBiznewVehicleRequest {
	s.VehicleCode = &v
	return s
}

func (s *SaveBiznewVehicleRequest) SetForwarderDid(v string) *SaveBiznewVehicleRequest {
	s.ForwarderDid = &v
	return s
}

func (s *SaveBiznewVehicleRequest) SetVehicleTotalAmount(v string) *SaveBiznewVehicleRequest {
	s.VehicleTotalAmount = &v
	return s
}

func (s *SaveBiznewVehicleRequest) SetVehicleAmount(v string) *SaveBiznewVehicleRequest {
	s.VehicleAmount = &v
	return s
}

func (s *SaveBiznewVehicleRequest) SetVehicleCharges(v string) *SaveBiznewVehicleRequest {
	s.VehicleCharges = &v
	return s
}

func (s *SaveBiznewVehicleRequest) SetCurrency(v string) *SaveBiznewVehicleRequest {
	s.Currency = &v
	return s
}

func (s *SaveBiznewVehicleRequest) SetContainerIdInfoList(v []*ContainerIdInfo) *SaveBiznewVehicleRequest {
	s.ContainerIdInfoList = v
	return s
}

type SaveBiznewVehicleResponse struct {
	// 请求唯一ID，用于链路跟踪和问题排查
	ReqMsgId *string `json:"req_msg_id,omitempty" xml:"req_msg_id,omitempty"`
	// 结果码，一般OK表示调用成功
	ResultCode *string `json:"result_code,omitempty" xml:"result_code,omitempty"`
	// 异常信息的文本描述
	ResultMsg *string `json:"result_msg,omitempty" xml:"result_msg,omitempty"`
	// 链上存证
	TxCodes []*TxDto `json:"tx_codes,omitempty" xml:"tx_codes,omitempty" type:"Repeated"`
}

func (s SaveBiznewVehicleResponse) String() string {
	return tea.Prettify(s)
}

func (s SaveBiznewVehicleResponse) GoString() string {
	return s.String()
}

func (s *SaveBiznewVehicleResponse) SetReqMsgId(v string) *SaveBiznewVehicleResponse {
	s.ReqMsgId = &v
	return s
}

func (s *SaveBiznewVehicleResponse) SetResultCode(v string) *SaveBiznewVehicleResponse {
	s.ResultCode = &v
	return s
}

func (s *SaveBiznewVehicleResponse) SetResultMsg(v string) *SaveBiznewVehicleResponse {
	s.ResultMsg = &v
	return s
}

func (s *SaveBiznewVehicleResponse) SetTxCodes(v []*TxDto) *SaveBiznewVehicleResponse {
	s.TxCodes = v
	return s
}

type SaveBiznewCustomsRequest struct {
	// OAuth模式下的授权token
	AuthToken         *string `json:"auth_token,omitempty" xml:"auth_token,omitempty"`
	ProductInstanceId *string `json:"product_instance_id,omitempty" xml:"product_instance_id,omitempty"`
	// 托运订单号
	OrderNo *string `json:"order_no,omitempty" xml:"order_no,omitempty" require:"true"`
	// 报关单号
	//
	//
	CustomsCode *string `json:"customs_code,omitempty" xml:"customs_code,omitempty" require:"true"`
	// 货代did
	ForwarderDid *string `json:"forwarder_did,omitempty" xml:"forwarder_did,omitempty" require:"true"`
	// 报关代理
	//
	//
	Broker *string `json:"broker,omitempty" xml:"broker,omitempty"`
	// 出口人
	Exporter *string `json:"exporter,omitempty" xml:"exporter,omitempty"`
	// 船名 [业务必填]
	Vessel *string `json:"vessel,omitempty" xml:"vessel,omitempty"`
	// 航次 [业务必填]
	Voyage *string `json:"voyage,omitempty" xml:"voyage,omitempty"`
	// 报关状态
	// APPROVED--通关，UNAPPROVED-未通关
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// 集装箱唯一标识
	ContainerId *string `json:"container_id,omitempty" xml:"container_id,omitempty"`
	// 箱号
	ContainerNo *string `json:"container_no,omitempty" xml:"container_no,omitempty"`
	// 货物名称
	Goods *string `json:"goods,omitempty" xml:"goods,omitempty"`
	// 毛重
	GrossWeight *string `json:"gross_weight,omitempty" xml:"gross_weight,omitempty"`
	// 件数
	PackagesNo *string `json:"packages_no,omitempty" xml:"packages_no,omitempty"`
	// 报关总金额 [业务必填]
	CustomsTotalAmount *string `json:"customs_total_amount,omitempty" xml:"customs_total_amount,omitempty"`
	// 报关运费金额 [业务必填]
	CustomsAmount *string `json:"customs_amount,omitempty" xml:"customs_amount,omitempty"`
	// 报关杂费金额 [业务必填]
	CustomsCharges *string `json:"customs_charges,omitempty" xml:"customs_charges,omitempty"`
	// 币种 [业务必填]
	Currency *string `json:"currency,omitempty" xml:"currency,omitempty"`
	// 订舱单号列表 [业务必填]
	BookingInfoList []*BookingNoInfo `json:"booking_info_list,omitempty" xml:"booking_info_list,omitempty" type:"Repeated"`
}

func (s SaveBiznewCustomsRequest) String() string {
	return tea.Prettify(s)
}

func (s SaveBiznewCustomsRequest) GoString() string {
	return s.String()
}

func (s *SaveBiznewCustomsRequest) SetAuthToken(v string) *SaveBiznewCustomsRequest {
	s.AuthToken = &v
	return s
}

func (s *SaveBiznewCustomsRequest) SetProductInstanceId(v string) *SaveBiznewCustomsRequest {
	s.ProductInstanceId = &v
	return s
}

func (s *SaveBiznewCustomsRequest) SetOrderNo(v string) *SaveBiznewCustomsRequest {
	s.OrderNo = &v
	return s
}

func (s *SaveBiznewCustomsRequest) SetCustomsCode(v string) *SaveBiznewCustomsRequest {
	s.CustomsCode = &v
	return s
}

func (s *SaveBiznewCustomsRequest) SetForwarderDid(v string) *SaveBiznewCustomsRequest {
	s.ForwarderDid = &v
	return s
}

func (s *SaveBiznewCustomsRequest) SetBroker(v string) *SaveBiznewCustomsRequest {
	s.Broker = &v
	return s
}

func (s *SaveBiznewCustomsRequest) SetExporter(v string) *SaveBiznewCustomsRequest {
	s.Exporter = &v
	return s
}

func (s *SaveBiznewCustomsRequest) SetVessel(v string) *SaveBiznewCustomsRequest {
	s.Vessel = &v
	return s
}

func (s *SaveBiznewCustomsRequest) SetVoyage(v string) *SaveBiznewCustomsRequest {
	s.Voyage = &v
	return s
}

func (s *SaveBiznewCustomsRequest) SetStatus(v string) *SaveBiznewCustomsRequest {
	s.Status = &v
	return s
}

func (s *SaveBiznewCustomsRequest) SetContainerId(v string) *SaveBiznewCustomsRequest {
	s.ContainerId = &v
	return s
}

func (s *SaveBiznewCustomsRequest) SetContainerNo(v string) *SaveBiznewCustomsRequest {
	s.ContainerNo = &v
	return s
}

func (s *SaveBiznewCustomsRequest) SetGoods(v string) *SaveBiznewCustomsRequest {
	s.Goods = &v
	return s
}

func (s *SaveBiznewCustomsRequest) SetGrossWeight(v string) *SaveBiznewCustomsRequest {
	s.GrossWeight = &v
	return s
}

func (s *SaveBiznewCustomsRequest) SetPackagesNo(v string) *SaveBiznewCustomsRequest {
	s.PackagesNo = &v
	return s
}

func (s *SaveBiznewCustomsRequest) SetCustomsTotalAmount(v string) *SaveBiznewCustomsRequest {
	s.CustomsTotalAmount = &v
	return s
}

func (s *SaveBiznewCustomsRequest) SetCustomsAmount(v string) *SaveBiznewCustomsRequest {
	s.CustomsAmount = &v
	return s
}

func (s *SaveBiznewCustomsRequest) SetCustomsCharges(v string) *SaveBiznewCustomsRequest {
	s.CustomsCharges = &v
	return s
}

func (s *SaveBiznewCustomsRequest) SetCurrency(v string) *SaveBiznewCustomsRequest {
	s.Currency = &v
	return s
}

func (s *SaveBiznewCustomsRequest) SetBookingInfoList(v []*BookingNoInfo) *SaveBiznewCustomsRequest {
	s.BookingInfoList = v
	return s
}

type SaveBiznewCustomsResponse struct {
	// 请求唯一ID，用于链路跟踪和问题排查
	ReqMsgId *string `json:"req_msg_id,omitempty" xml:"req_msg_id,omitempty"`
	// 结果码，一般OK表示调用成功
	ResultCode *string `json:"result_code,omitempty" xml:"result_code,omitempty"`
	// 异常信息的文本描述
	ResultMsg *string `json:"result_msg,omitempty" xml:"result_msg,omitempty"`
	// 链上存证
	TxCodes []*TxDto `json:"tx_codes,omitempty" xml:"tx_codes,omitempty" type:"Repeated"`
}

func (s SaveBiznewCustomsResponse) String() string {
	return tea.Prettify(s)
}

func (s SaveBiznewCustomsResponse) GoString() string {
	return s.String()
}

func (s *SaveBiznewCustomsResponse) SetReqMsgId(v string) *SaveBiznewCustomsResponse {
	s.ReqMsgId = &v
	return s
}

func (s *SaveBiznewCustomsResponse) SetResultCode(v string) *SaveBiznewCustomsResponse {
	s.ResultCode = &v
	return s
}

func (s *SaveBiznewCustomsResponse) SetResultMsg(v string) *SaveBiznewCustomsResponse {
	s.ResultMsg = &v
	return s
}

func (s *SaveBiznewCustomsResponse) SetTxCodes(v []*TxDto) *SaveBiznewCustomsResponse {
	s.TxCodes = v
	return s
}

type SaveBiznewMasterRequest struct {
	// OAuth模式下的授权token
	AuthToken         *string `json:"auth_token,omitempty" xml:"auth_token,omitempty"`
	ProductInstanceId *string `json:"product_instance_id,omitempty" xml:"product_instance_id,omitempty"`
	// 托运订单号
	OrderNo *string `json:"order_no,omitempty" xml:"order_no,omitempty" require:"true"`
	// master提单号或House提单号
	MasterBlNo *string `json:"master_bl_no,omitempty" xml:"master_bl_no,omitempty" require:"true"`
	// 提单类型  master/house
	//
	//
	Type *string `json:"type,omitempty" xml:"type,omitempty" require:"true"`
	// 提单签发单位 [业务必填]
	//
	//
	SignUnitName *string `json:"sign_unit_name,omitempty" xml:"sign_unit_name,omitempty"`
	// 货代did
	ForwarderDid *string `json:"forwarder_did,omitempty" xml:"forwarder_did,omitempty" require:"true"`
	// 发货人 [业务必填]
	Shipper *string `json:"shipper,omitempty" xml:"shipper,omitempty"`
	// 出口人 [业务必填]
	//
	//
	Consignee *string `json:"consignee,omitempty" xml:"consignee,omitempty"`
	// 船公司 [业务必填]
	Carrier *string `json:"carrier,omitempty" xml:"carrier,omitempty"`
	// 船名 [业务必填]
	Vessel *string `json:"vessel,omitempty" xml:"vessel,omitempty"`
	// 航次 [业务必填]
	Voyage *string `json:"voyage,omitempty" xml:"voyage,omitempty"`
	// 起运港 [业务必填]
	LoadingPort *string `json:"loading_port,omitempty" xml:"loading_port,omitempty"`
	// 卸货港 [业务必填]
	DischargePort *string `json:"discharge_port,omitempty" xml:"discharge_port,omitempty"`
	// 目的地 [业务必填]
	DeliveryPlace *string `json:"delivery_place,omitempty" xml:"delivery_place,omitempty"`
	// 开船时间 (秒时间戳)
	OnBoardDate *int64 `json:"on_board_date,omitempty" xml:"on_board_date,omitempty"`
	// 船状态 ATA （已到港） ，ATD （已离港 ），UNATD （未离港）
	OnBoardStatus *string `json:"on_board_status,omitempty" xml:"on_board_status,omitempty"`
	// 订舱单号列表 [业务必填]
	BookingInfoList []*BookingNoInfo `json:"booking_info_list,omitempty" xml:"booking_info_list,omitempty" type:"Repeated"`
	// 集装箱信息 [业务必填]
	ContainerInfoList []*ContainerInfo `json:"container_info_list,omitempty" xml:"container_info_list,omitempty" type:"Repeated"`
	// 货物列表
	//
	//
	GoodsInfoList []*GoodsInfo `json:"goods_info_list,omitempty" xml:"goods_info_list,omitempty" type:"Repeated"`
}

func (s SaveBiznewMasterRequest) String() string {
	return tea.Prettify(s)
}

func (s SaveBiznewMasterRequest) GoString() string {
	return s.String()
}

func (s *SaveBiznewMasterRequest) SetAuthToken(v string) *SaveBiznewMasterRequest {
	s.AuthToken = &v
	return s
}

func (s *SaveBiznewMasterRequest) SetProductInstanceId(v string) *SaveBiznewMasterRequest {
	s.ProductInstanceId = &v
	return s
}

func (s *SaveBiznewMasterRequest) SetOrderNo(v string) *SaveBiznewMasterRequest {
	s.OrderNo = &v
	return s
}

func (s *SaveBiznewMasterRequest) SetMasterBlNo(v string) *SaveBiznewMasterRequest {
	s.MasterBlNo = &v
	return s
}

func (s *SaveBiznewMasterRequest) SetType(v string) *SaveBiznewMasterRequest {
	s.Type = &v
	return s
}

func (s *SaveBiznewMasterRequest) SetSignUnitName(v string) *SaveBiznewMasterRequest {
	s.SignUnitName = &v
	return s
}

func (s *SaveBiznewMasterRequest) SetForwarderDid(v string) *SaveBiznewMasterRequest {
	s.ForwarderDid = &v
	return s
}

func (s *SaveBiznewMasterRequest) SetShipper(v string) *SaveBiznewMasterRequest {
	s.Shipper = &v
	return s
}

func (s *SaveBiznewMasterRequest) SetConsignee(v string) *SaveBiznewMasterRequest {
	s.Consignee = &v
	return s
}

func (s *SaveBiznewMasterRequest) SetCarrier(v string) *SaveBiznewMasterRequest {
	s.Carrier = &v
	return s
}

func (s *SaveBiznewMasterRequest) SetVessel(v string) *SaveBiznewMasterRequest {
	s.Vessel = &v
	return s
}

func (s *SaveBiznewMasterRequest) SetVoyage(v string) *SaveBiznewMasterRequest {
	s.Voyage = &v
	return s
}

func (s *SaveBiznewMasterRequest) SetLoadingPort(v string) *SaveBiznewMasterRequest {
	s.LoadingPort = &v
	return s
}

func (s *SaveBiznewMasterRequest) SetDischargePort(v string) *SaveBiznewMasterRequest {
	s.DischargePort = &v
	return s
}

func (s *SaveBiznewMasterRequest) SetDeliveryPlace(v string) *SaveBiznewMasterRequest {
	s.DeliveryPlace = &v
	return s
}

func (s *SaveBiznewMasterRequest) SetOnBoardDate(v int64) *SaveBiznewMasterRequest {
	s.OnBoardDate = &v
	return s
}

func (s *SaveBiznewMasterRequest) SetOnBoardStatus(v string) *SaveBiznewMasterRequest {
	s.OnBoardStatus = &v
	return s
}

func (s *SaveBiznewMasterRequest) SetBookingInfoList(v []*BookingNoInfo) *SaveBiznewMasterRequest {
	s.BookingInfoList = v
	return s
}

func (s *SaveBiznewMasterRequest) SetContainerInfoList(v []*ContainerInfo) *SaveBiznewMasterRequest {
	s.ContainerInfoList = v
	return s
}

func (s *SaveBiznewMasterRequest) SetGoodsInfoList(v []*GoodsInfo) *SaveBiznewMasterRequest {
	s.GoodsInfoList = v
	return s
}

type SaveBiznewMasterResponse struct {
	// 请求唯一ID，用于链路跟踪和问题排查
	ReqMsgId *string `json:"req_msg_id,omitempty" xml:"req_msg_id,omitempty"`
	// 结果码，一般OK表示调用成功
	ResultCode *string `json:"result_code,omitempty" xml:"result_code,omitempty"`
	// 异常信息的文本描述
	ResultMsg *string `json:"result_msg,omitempty" xml:"result_msg,omitempty"`
	// 链上存证
	TxCodes []*TxDto `json:"tx_codes,omitempty" xml:"tx_codes,omitempty" type:"Repeated"`
}

func (s SaveBiznewMasterResponse) String() string {
	return tea.Prettify(s)
}

func (s SaveBiznewMasterResponse) GoString() string {
	return s.String()
}

func (s *SaveBiznewMasterResponse) SetReqMsgId(v string) *SaveBiznewMasterResponse {
	s.ReqMsgId = &v
	return s
}

func (s *SaveBiznewMasterResponse) SetResultCode(v string) *SaveBiznewMasterResponse {
	s.ResultCode = &v
	return s
}

func (s *SaveBiznewMasterResponse) SetResultMsg(v string) *SaveBiznewMasterResponse {
	s.ResultMsg = &v
	return s
}

func (s *SaveBiznewMasterResponse) SetTxCodes(v []*TxDto) *SaveBiznewMasterResponse {
	s.TxCodes = v
	return s
}

type SaveBiznewPaybillorderRequest struct {
	// OAuth模式下的授权token
	AuthToken         *string `json:"auth_token,omitempty" xml:"auth_token,omitempty"`
	ProductInstanceId *string `json:"product_instance_id,omitempty" xml:"product_instance_id,omitempty"`
	// 应付账单号
	//
	//
	PayBillOrderCode *string `json:"pay_bill_order_code,omitempty" xml:"pay_bill_order_code,omitempty" require:"true"`
	// 货代did
	ForwarderDid *string `json:"forwarder_did,omitempty" xml:"forwarder_did,omitempty" require:"true"`
	// 应付方did [业务必填]
	//
	//
	SettleCompanyDid *string `json:"settle_company_did,omitempty" xml:"settle_company_did,omitempty"`
	// 收款方did [业务必填]
	//
	//
	ReceiptClientDid *string `json:"receipt_client_did,omitempty" xml:"receipt_client_did,omitempty"`
	// 付款金额 [业务必填]
	//
	//
	PayAmount *string `json:"pay_amount,omitempty" xml:"pay_amount,omitempty"`
	// 申请时间 （秒时间戳） [业务必填]
	//
	//
	ApplyDate *int64 `json:"apply_date,omitempty" xml:"apply_date,omitempty"`
	// 账单到期时间（秒时间戳） [业务必填]
	ExpireDate *int64 `json:"expire_date,omitempty" xml:"expire_date,omitempty"`
	// 账单状态  SETTLED 已结算 UNSETTLE 未结算 (业务必填)
	SettleStatus *string `json:"settle_status,omitempty" xml:"settle_status,omitempty"`
	// 应付资费项 (业务必填)
	PayTariffList []*PayTariffInfo `json:"pay_tariff_list,omitempty" xml:"pay_tariff_list,omitempty" type:"Repeated"`
	// 币种 [业务必填]
	Currency *string `json:"currency,omitempty" xml:"currency,omitempty"`
}

func (s SaveBiznewPaybillorderRequest) String() string {
	return tea.Prettify(s)
}

func (s SaveBiznewPaybillorderRequest) GoString() string {
	return s.String()
}

func (s *SaveBiznewPaybillorderRequest) SetAuthToken(v string) *SaveBiznewPaybillorderRequest {
	s.AuthToken = &v
	return s
}

func (s *SaveBiznewPaybillorderRequest) SetProductInstanceId(v string) *SaveBiznewPaybillorderRequest {
	s.ProductInstanceId = &v
	return s
}

func (s *SaveBiznewPaybillorderRequest) SetPayBillOrderCode(v string) *SaveBiznewPaybillorderRequest {
	s.PayBillOrderCode = &v
	return s
}

func (s *SaveBiznewPaybillorderRequest) SetForwarderDid(v string) *SaveBiznewPaybillorderRequest {
	s.ForwarderDid = &v
	return s
}

func (s *SaveBiznewPaybillorderRequest) SetSettleCompanyDid(v string) *SaveBiznewPaybillorderRequest {
	s.SettleCompanyDid = &v
	return s
}

func (s *SaveBiznewPaybillorderRequest) SetReceiptClientDid(v string) *SaveBiznewPaybillorderRequest {
	s.ReceiptClientDid = &v
	return s
}

func (s *SaveBiznewPaybillorderRequest) SetPayAmount(v string) *SaveBiznewPaybillorderRequest {
	s.PayAmount = &v
	return s
}

func (s *SaveBiznewPaybillorderRequest) SetApplyDate(v int64) *SaveBiznewPaybillorderRequest {
	s.ApplyDate = &v
	return s
}

func (s *SaveBiznewPaybillorderRequest) SetExpireDate(v int64) *SaveBiznewPaybillorderRequest {
	s.ExpireDate = &v
	return s
}

func (s *SaveBiznewPaybillorderRequest) SetSettleStatus(v string) *SaveBiznewPaybillorderRequest {
	s.SettleStatus = &v
	return s
}

func (s *SaveBiznewPaybillorderRequest) SetPayTariffList(v []*PayTariffInfo) *SaveBiznewPaybillorderRequest {
	s.PayTariffList = v
	return s
}

func (s *SaveBiznewPaybillorderRequest) SetCurrency(v string) *SaveBiznewPaybillorderRequest {
	s.Currency = &v
	return s
}

type SaveBiznewPaybillorderResponse struct {
	// 请求唯一ID，用于链路跟踪和问题排查
	ReqMsgId *string `json:"req_msg_id,omitempty" xml:"req_msg_id,omitempty"`
	// 结果码，一般OK表示调用成功
	ResultCode *string `json:"result_code,omitempty" xml:"result_code,omitempty"`
	// 异常信息的文本描述
	ResultMsg *string `json:"result_msg,omitempty" xml:"result_msg,omitempty"`
	// 链上存证
	TxCodes []*TxDto `json:"tx_codes,omitempty" xml:"tx_codes,omitempty" type:"Repeated"`
}

func (s SaveBiznewPaybillorderResponse) String() string {
	return tea.Prettify(s)
}

func (s SaveBiznewPaybillorderResponse) GoString() string {
	return s.String()
}

func (s *SaveBiznewPaybillorderResponse) SetReqMsgId(v string) *SaveBiznewPaybillorderResponse {
	s.ReqMsgId = &v
	return s
}

func (s *SaveBiznewPaybillorderResponse) SetResultCode(v string) *SaveBiznewPaybillorderResponse {
	s.ResultCode = &v
	return s
}

func (s *SaveBiznewPaybillorderResponse) SetResultMsg(v string) *SaveBiznewPaybillorderResponse {
	s.ResultMsg = &v
	return s
}

func (s *SaveBiznewPaybillorderResponse) SetTxCodes(v []*TxDto) *SaveBiznewPaybillorderResponse {
	s.TxCodes = v
	return s
}

type SaveBiznewReceiptbillorderRequest struct {
	// OAuth模式下的授权token
	AuthToken         *string `json:"auth_token,omitempty" xml:"auth_token,omitempty"`
	ProductInstanceId *string `json:"product_instance_id,omitempty" xml:"product_instance_id,omitempty"`
	// 应收账单号
	//
	//
	ReceiptBillOrderCode *string `json:"receipt_bill_order_code,omitempty" xml:"receipt_bill_order_code,omitempty" require:"true"`
	// 货代did
	ForwarderDid *string `json:"forwarder_did,omitempty" xml:"forwarder_did,omitempty" require:"true"`
	// 应付方did [业务必填]
	//
	//
	SettleCompanyDid *string `json:"settle_company_did,omitempty" xml:"settle_company_did,omitempty"`
	// 收款方did [业务必填]
	//
	//
	ReceiptClientDid *string `json:"receipt_client_did,omitempty" xml:"receipt_client_did,omitempty"`
	// 收款金额 [业务必填]
	//
	//
	ReceiptAmount *string `json:"receipt_amount,omitempty" xml:"receipt_amount,omitempty"`
	// 申请时间 (秒时间戳) [业务必填]
	ApplyDate *int64 `json:"apply_date,omitempty" xml:"apply_date,omitempty"`
	// 账单到期日 (秒时间戳) [业务必填]
	//
	//
	ExpireDate *int64 `json:"expire_date,omitempty" xml:"expire_date,omitempty"`
	// 账单状态  SETTLED 已结算 UNSETTLE 未结算  [业务必填]
	SettleStatus *string `json:"settle_status,omitempty" xml:"settle_status,omitempty"`
	// 应收资费项 [业务必填]
	//
	//
	ReceiptTariffList []*ReceiptTariffInfo `json:"receipt_tariff_list,omitempty" xml:"receipt_tariff_list,omitempty" type:"Repeated"`
	// 币种 [业务必填]
	Currency *string `json:"currency,omitempty" xml:"currency,omitempty"`
}

func (s SaveBiznewReceiptbillorderRequest) String() string {
	return tea.Prettify(s)
}

func (s SaveBiznewReceiptbillorderRequest) GoString() string {
	return s.String()
}

func (s *SaveBiznewReceiptbillorderRequest) SetAuthToken(v string) *SaveBiznewReceiptbillorderRequest {
	s.AuthToken = &v
	return s
}

func (s *SaveBiznewReceiptbillorderRequest) SetProductInstanceId(v string) *SaveBiznewReceiptbillorderRequest {
	s.ProductInstanceId = &v
	return s
}

func (s *SaveBiznewReceiptbillorderRequest) SetReceiptBillOrderCode(v string) *SaveBiznewReceiptbillorderRequest {
	s.ReceiptBillOrderCode = &v
	return s
}

func (s *SaveBiznewReceiptbillorderRequest) SetForwarderDid(v string) *SaveBiznewReceiptbillorderRequest {
	s.ForwarderDid = &v
	return s
}

func (s *SaveBiznewReceiptbillorderRequest) SetSettleCompanyDid(v string) *SaveBiznewReceiptbillorderRequest {
	s.SettleCompanyDid = &v
	return s
}

func (s *SaveBiznewReceiptbillorderRequest) SetReceiptClientDid(v string) *SaveBiznewReceiptbillorderRequest {
	s.ReceiptClientDid = &v
	return s
}

func (s *SaveBiznewReceiptbillorderRequest) SetReceiptAmount(v string) *SaveBiznewReceiptbillorderRequest {
	s.ReceiptAmount = &v
	return s
}

func (s *SaveBiznewReceiptbillorderRequest) SetApplyDate(v int64) *SaveBiznewReceiptbillorderRequest {
	s.ApplyDate = &v
	return s
}

func (s *SaveBiznewReceiptbillorderRequest) SetExpireDate(v int64) *SaveBiznewReceiptbillorderRequest {
	s.ExpireDate = &v
	return s
}

func (s *SaveBiznewReceiptbillorderRequest) SetSettleStatus(v string) *SaveBiznewReceiptbillorderRequest {
	s.SettleStatus = &v
	return s
}

func (s *SaveBiznewReceiptbillorderRequest) SetReceiptTariffList(v []*ReceiptTariffInfo) *SaveBiznewReceiptbillorderRequest {
	s.ReceiptTariffList = v
	return s
}

func (s *SaveBiznewReceiptbillorderRequest) SetCurrency(v string) *SaveBiznewReceiptbillorderRequest {
	s.Currency = &v
	return s
}

type SaveBiznewReceiptbillorderResponse struct {
	// 请求唯一ID，用于链路跟踪和问题排查
	ReqMsgId *string `json:"req_msg_id,omitempty" xml:"req_msg_id,omitempty"`
	// 结果码，一般OK表示调用成功
	ResultCode *string `json:"result_code,omitempty" xml:"result_code,omitempty"`
	// 异常信息的文本描述
	ResultMsg *string `json:"result_msg,omitempty" xml:"result_msg,omitempty"`
	// 链上存证
	TxCodes []*TxDto `json:"tx_codes,omitempty" xml:"tx_codes,omitempty" type:"Repeated"`
}

func (s SaveBiznewReceiptbillorderResponse) String() string {
	return tea.Prettify(s)
}

func (s SaveBiznewReceiptbillorderResponse) GoString() string {
	return s.String()
}

func (s *SaveBiznewReceiptbillorderResponse) SetReqMsgId(v string) *SaveBiznewReceiptbillorderResponse {
	s.ReqMsgId = &v
	return s
}

func (s *SaveBiznewReceiptbillorderResponse) SetResultCode(v string) *SaveBiznewReceiptbillorderResponse {
	s.ResultCode = &v
	return s
}

func (s *SaveBiznewReceiptbillorderResponse) SetResultMsg(v string) *SaveBiznewReceiptbillorderResponse {
	s.ResultMsg = &v
	return s
}

func (s *SaveBiznewReceiptbillorderResponse) SetTxCodes(v []*TxDto) *SaveBiznewReceiptbillorderResponse {
	s.TxCodes = v
	return s
}

type SaveBiznewInvoiceRequest struct {
	// OAuth模式下的授权token
	AuthToken         *string `json:"auth_token,omitempty" xml:"auth_token,omitempty"`
	ProductInstanceId *string `json:"product_instance_id,omitempty" xml:"product_instance_id,omitempty"`
	// 货代did
	ForwarderDid *string `json:"forwarder_did,omitempty" xml:"forwarder_did,omitempty" require:"true"`
	// 发票唯一标识
	InvoiceCode *string `json:"invoice_code,omitempty" xml:"invoice_code,omitempty" require:"true"`
	// 发票税务号 [业务必填]
	InvoiceTaxCode *string `json:"invoice_tax_code,omitempty" xml:"invoice_tax_code,omitempty"`
	// 发票号 [业务必填]
	InvoiceNumber *string `json:"invoice_number,omitempty" xml:"invoice_number,omitempty"`
	// 发票类型 [业务必填]
	InvoiceType *string `json:"invoice_type,omitempty" xml:"invoice_type,omitempty"`
	// 开票类型 pay/receipt (应付/应收) [业务必填]
	//
	//
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
	// 校验码
	//
	//
	CheckCode *string `json:"check_code,omitempty" xml:"check_code,omitempty"`
	// 开票方名称 [业务必填]
	//
	//
	DrawerName *string `json:"drawer_name,omitempty" xml:"drawer_name,omitempty"`
	// 开票纳税人识别号 [业务必填]
	//
	//
	DrawerTaxpayerCode *string `json:"drawer_taxpayer_code,omitempty" xml:"drawer_taxpayer_code,omitempty"`
	// 受票方企业名称 [业务必填]
	//
	//
	InvoiceHeaderName *string `json:"invoice_header_name,omitempty" xml:"invoice_header_name,omitempty"`
	// 受票方企业信用证代码 [业务必填]
	//
	//
	InvoiceHeaderSocialNo *string `json:"invoice_header_social_no,omitempty" xml:"invoice_header_social_no,omitempty"`
	// 开票时间 [业务必填]
	//
	//
	MakeInvoiceDate *int64 `json:"make_invoice_date,omitempty" xml:"make_invoice_date,omitempty"`
	// 发票金额 [业务必填]
	//
	//
	Amount *string `json:"amount,omitempty" xml:"amount,omitempty"`
	// 币种 [业务必填]
	Currency *string `json:"currency,omitempty" xml:"currency,omitempty"`
	// 美元
	DollarAmount *string `json:"dollar_amount,omitempty" xml:"dollar_amount,omitempty"`
	// 不含税金额 [业务必填]
	UntaxedPrice *string `json:"untaxed_price,omitempty" xml:"untaxed_price,omitempty"`
	// 发票文件ID [业务必填]
	FileId *string `json:"file_id,omitempty" xml:"file_id,omitempty"`
	// 发票文件hash [业务必填]
	FileHash *string `json:"file_hash,omitempty" xml:"file_hash,omitempty"`
}

func (s SaveBiznewInvoiceRequest) String() string {
	return tea.Prettify(s)
}

func (s SaveBiznewInvoiceRequest) GoString() string {
	return s.String()
}

func (s *SaveBiznewInvoiceRequest) SetAuthToken(v string) *SaveBiznewInvoiceRequest {
	s.AuthToken = &v
	return s
}

func (s *SaveBiznewInvoiceRequest) SetProductInstanceId(v string) *SaveBiznewInvoiceRequest {
	s.ProductInstanceId = &v
	return s
}

func (s *SaveBiznewInvoiceRequest) SetForwarderDid(v string) *SaveBiznewInvoiceRequest {
	s.ForwarderDid = &v
	return s
}

func (s *SaveBiznewInvoiceRequest) SetInvoiceCode(v string) *SaveBiznewInvoiceRequest {
	s.InvoiceCode = &v
	return s
}

func (s *SaveBiznewInvoiceRequest) SetInvoiceTaxCode(v string) *SaveBiznewInvoiceRequest {
	s.InvoiceTaxCode = &v
	return s
}

func (s *SaveBiznewInvoiceRequest) SetInvoiceNumber(v string) *SaveBiznewInvoiceRequest {
	s.InvoiceNumber = &v
	return s
}

func (s *SaveBiznewInvoiceRequest) SetInvoiceType(v string) *SaveBiznewInvoiceRequest {
	s.InvoiceType = &v
	return s
}

func (s *SaveBiznewInvoiceRequest) SetType(v string) *SaveBiznewInvoiceRequest {
	s.Type = &v
	return s
}

func (s *SaveBiznewInvoiceRequest) SetCheckCode(v string) *SaveBiznewInvoiceRequest {
	s.CheckCode = &v
	return s
}

func (s *SaveBiznewInvoiceRequest) SetDrawerName(v string) *SaveBiznewInvoiceRequest {
	s.DrawerName = &v
	return s
}

func (s *SaveBiznewInvoiceRequest) SetDrawerTaxpayerCode(v string) *SaveBiznewInvoiceRequest {
	s.DrawerTaxpayerCode = &v
	return s
}

func (s *SaveBiznewInvoiceRequest) SetInvoiceHeaderName(v string) *SaveBiznewInvoiceRequest {
	s.InvoiceHeaderName = &v
	return s
}

func (s *SaveBiznewInvoiceRequest) SetInvoiceHeaderSocialNo(v string) *SaveBiznewInvoiceRequest {
	s.InvoiceHeaderSocialNo = &v
	return s
}

func (s *SaveBiznewInvoiceRequest) SetMakeInvoiceDate(v int64) *SaveBiznewInvoiceRequest {
	s.MakeInvoiceDate = &v
	return s
}

func (s *SaveBiznewInvoiceRequest) SetAmount(v string) *SaveBiznewInvoiceRequest {
	s.Amount = &v
	return s
}

func (s *SaveBiznewInvoiceRequest) SetCurrency(v string) *SaveBiznewInvoiceRequest {
	s.Currency = &v
	return s
}

func (s *SaveBiznewInvoiceRequest) SetDollarAmount(v string) *SaveBiznewInvoiceRequest {
	s.DollarAmount = &v
	return s
}

func (s *SaveBiznewInvoiceRequest) SetUntaxedPrice(v string) *SaveBiznewInvoiceRequest {
	s.UntaxedPrice = &v
	return s
}

func (s *SaveBiznewInvoiceRequest) SetFileId(v string) *SaveBiznewInvoiceRequest {
	s.FileId = &v
	return s
}

func (s *SaveBiznewInvoiceRequest) SetFileHash(v string) *SaveBiznewInvoiceRequest {
	s.FileHash = &v
	return s
}

type SaveBiznewInvoiceResponse struct {
	// 请求唯一ID，用于链路跟踪和问题排查
	ReqMsgId *string `json:"req_msg_id,omitempty" xml:"req_msg_id,omitempty"`
	// 结果码，一般OK表示调用成功
	ResultCode *string `json:"result_code,omitempty" xml:"result_code,omitempty"`
	// 异常信息的文本描述
	ResultMsg *string `json:"result_msg,omitempty" xml:"result_msg,omitempty"`
	// 链上存证
	TxCodes []*TxDto `json:"tx_codes,omitempty" xml:"tx_codes,omitempty" type:"Repeated"`
}

func (s SaveBiznewInvoiceResponse) String() string {
	return tea.Prettify(s)
}

func (s SaveBiznewInvoiceResponse) GoString() string {
	return s.String()
}

func (s *SaveBiznewInvoiceResponse) SetReqMsgId(v string) *SaveBiznewInvoiceResponse {
	s.ReqMsgId = &v
	return s
}

func (s *SaveBiznewInvoiceResponse) SetResultCode(v string) *SaveBiznewInvoiceResponse {
	s.ResultCode = &v
	return s
}

func (s *SaveBiznewInvoiceResponse) SetResultMsg(v string) *SaveBiznewInvoiceResponse {
	s.ResultMsg = &v
	return s
}

func (s *SaveBiznewInvoiceResponse) SetTxCodes(v []*TxDto) *SaveBiznewInvoiceResponse {
	s.TxCodes = v
	return s
}

type UploadShippingEblRequest struct {
	// OAuth模式下的授权token
	AuthToken         *string `json:"auth_token,omitempty" xml:"auth_token,omitempty"`
	ProductInstanceId *string `json:"product_instance_id,omitempty" xml:"product_instance_id,omitempty"`
	// 船公司名称
	CarrierName *string `json:"carrier_name,omitempty" xml:"carrier_name,omitempty" require:"true"`
	// 收货人did
	ConsigneeDid *string `json:"consignee_did,omitempty" xml:"consignee_did,omitempty" require:"true"`
	// 电子提单类型
	EblCategory *string `json:"ebl_category,omitempty" xml:"ebl_category,omitempty" require:"true"`
	// 电子提单copy文件hash
	EblCopyPdfFileHash *string `json:"ebl_copy_pdf_file_hash,omitempty" xml:"ebl_copy_pdf_file_hash,omitempty" require:"true"`
	// copy电子提单pdf文件id
	EblCopyPdfFileId *string `json:"ebl_copy_pdf_file_id,omitempty" xml:"ebl_copy_pdf_file_id,omitempty" require:"true"`
	// 电子提单编号
	EblNo *string `json:"ebl_no,omitempty" xml:"ebl_no,omitempty" require:"true"`
	// 电子甜的原文件hash
	EblOriginalPdfFileHash *string `json:"ebl_original_pdf_file_hash,omitempty" xml:"ebl_original_pdf_file_hash,omitempty"`
	// 原电子提单pdf文件id
	EblOriginalPdfFileId *string `json:"ebl_original_pdf_file_id,omitempty" xml:"ebl_original_pdf_file_id,omitempty"`
	// 议付行did
	NegotiatingBankDid *string `json:"negotiating_bank_did,omitempty" xml:"negotiating_bank_did,omitempty" require:"true"`
	// 托运人did
	ShipperDid *string `json:"shipper_did,omitempty" xml:"shipper_did,omitempty" require:"true"`
}

func (s UploadShippingEblRequest) String() string {
	return tea.Prettify(s)
}

func (s UploadShippingEblRequest) GoString() string {
	return s.String()
}

func (s *UploadShippingEblRequest) SetAuthToken(v string) *UploadShippingEblRequest {
	s.AuthToken = &v
	return s
}

func (s *UploadShippingEblRequest) SetProductInstanceId(v string) *UploadShippingEblRequest {
	s.ProductInstanceId = &v
	return s
}

func (s *UploadShippingEblRequest) SetCarrierName(v string) *UploadShippingEblRequest {
	s.CarrierName = &v
	return s
}

func (s *UploadShippingEblRequest) SetConsigneeDid(v string) *UploadShippingEblRequest {
	s.ConsigneeDid = &v
	return s
}

func (s *UploadShippingEblRequest) SetEblCategory(v string) *UploadShippingEblRequest {
	s.EblCategory = &v
	return s
}

func (s *UploadShippingEblRequest) SetEblCopyPdfFileHash(v string) *UploadShippingEblRequest {
	s.EblCopyPdfFileHash = &v
	return s
}

func (s *UploadShippingEblRequest) SetEblCopyPdfFileId(v string) *UploadShippingEblRequest {
	s.EblCopyPdfFileId = &v
	return s
}

func (s *UploadShippingEblRequest) SetEblNo(v string) *UploadShippingEblRequest {
	s.EblNo = &v
	return s
}

func (s *UploadShippingEblRequest) SetEblOriginalPdfFileHash(v string) *UploadShippingEblRequest {
	s.EblOriginalPdfFileHash = &v
	return s
}

func (s *UploadShippingEblRequest) SetEblOriginalPdfFileId(v string) *UploadShippingEblRequest {
	s.EblOriginalPdfFileId = &v
	return s
}

func (s *UploadShippingEblRequest) SetNegotiatingBankDid(v string) *UploadShippingEblRequest {
	s.NegotiatingBankDid = &v
	return s
}

func (s *UploadShippingEblRequest) SetShipperDid(v string) *UploadShippingEblRequest {
	s.ShipperDid = &v
	return s
}

type UploadShippingEblResponse struct {
	// 请求唯一ID，用于链路跟踪和问题排查
	ReqMsgId *string `json:"req_msg_id,omitempty" xml:"req_msg_id,omitempty"`
	// 结果码，一般OK表示调用成功
	ResultCode *string `json:"result_code,omitempty" xml:"result_code,omitempty"`
	// 异常信息的文本描述
	ResultMsg *string `json:"result_msg,omitempty" xml:"result_msg,omitempty"`
}

func (s UploadShippingEblResponse) String() string {
	return tea.Prettify(s)
}

func (s UploadShippingEblResponse) GoString() string {
	return s.String()
}

func (s *UploadShippingEblResponse) SetReqMsgId(v string) *UploadShippingEblResponse {
	s.ReqMsgId = &v
	return s
}

func (s *UploadShippingEblResponse) SetResultCode(v string) *UploadShippingEblResponse {
	s.ResultCode = &v
	return s
}

func (s *UploadShippingEblResponse) SetResultMsg(v string) *UploadShippingEblResponse {
	s.ResultMsg = &v
	return s
}

type UploadShippingEblbatchRequest struct {
	// OAuth模式下的授权token
	AuthToken         *string `json:"auth_token,omitempty" xml:"auth_token,omitempty"`
	ProductInstanceId *string `json:"product_instance_id,omitempty" xml:"product_instance_id,omitempty"`
	// 代理人did
	AgentDid *string `json:"agent_did,omitempty" xml:"agent_did,omitempty"`
	// 是否为代理操作
	AgentFlag *bool `json:"agent_flag,omitempty" xml:"agent_flag,omitempty" require:"true"`
	// 格式：carrierName_yyyyMMddHHmmss_5位随机数字；全局唯一
	BatchNo *string `json:"batch_no,omitempty" xml:"batch_no,omitempty" require:"true"`
	// 船公司名称
	CarrierName *string `json:"carrier_name,omitempty" xml:"carrier_name,omitempty" require:"true"`
	// 收货人did
	ConsigneeDid *string `json:"consignee_did,omitempty" xml:"consignee_did,omitempty" require:"true"`
	// 电子提单类型
	EblCategory *string `json:"ebl_category,omitempty" xml:"ebl_category,omitempty" require:"true"`
	// 批次下的提单明细
	//
	//
	EblDetails []*EblDetail `json:"ebl_details,omitempty" xml:"ebl_details,omitempty" require:"true" type:"Repeated"`
	// 议付行did
	NegotiatingBankDid *string `json:"negotiating_bank_did,omitempty" xml:"negotiating_bank_did,omitempty" require:"true"`
	// 托运人did
	ShipperDid *string `json:"shipper_did,omitempty" xml:"shipper_did,omitempty" require:"true"`
}

func (s UploadShippingEblbatchRequest) String() string {
	return tea.Prettify(s)
}

func (s UploadShippingEblbatchRequest) GoString() string {
	return s.String()
}

func (s *UploadShippingEblbatchRequest) SetAuthToken(v string) *UploadShippingEblbatchRequest {
	s.AuthToken = &v
	return s
}

func (s *UploadShippingEblbatchRequest) SetProductInstanceId(v string) *UploadShippingEblbatchRequest {
	s.ProductInstanceId = &v
	return s
}

func (s *UploadShippingEblbatchRequest) SetAgentDid(v string) *UploadShippingEblbatchRequest {
	s.AgentDid = &v
	return s
}

func (s *UploadShippingEblbatchRequest) SetAgentFlag(v bool) *UploadShippingEblbatchRequest {
	s.AgentFlag = &v
	return s
}

func (s *UploadShippingEblbatchRequest) SetBatchNo(v string) *UploadShippingEblbatchRequest {
	s.BatchNo = &v
	return s
}

func (s *UploadShippingEblbatchRequest) SetCarrierName(v string) *UploadShippingEblbatchRequest {
	s.CarrierName = &v
	return s
}

func (s *UploadShippingEblbatchRequest) SetConsigneeDid(v string) *UploadShippingEblbatchRequest {
	s.ConsigneeDid = &v
	return s
}

func (s *UploadShippingEblbatchRequest) SetEblCategory(v string) *UploadShippingEblbatchRequest {
	s.EblCategory = &v
	return s
}

func (s *UploadShippingEblbatchRequest) SetEblDetails(v []*EblDetail) *UploadShippingEblbatchRequest {
	s.EblDetails = v
	return s
}

func (s *UploadShippingEblbatchRequest) SetNegotiatingBankDid(v string) *UploadShippingEblbatchRequest {
	s.NegotiatingBankDid = &v
	return s
}

func (s *UploadShippingEblbatchRequest) SetShipperDid(v string) *UploadShippingEblbatchRequest {
	s.ShipperDid = &v
	return s
}

type UploadShippingEblbatchResponse struct {
	// 请求唯一ID，用于链路跟踪和问题排查
	ReqMsgId *string `json:"req_msg_id,omitempty" xml:"req_msg_id,omitempty"`
	// 结果码，一般OK表示调用成功
	ResultCode *string `json:"result_code,omitempty" xml:"result_code,omitempty"`
	// 异常信息的文本描述
	ResultMsg *string `json:"result_msg,omitempty" xml:"result_msg,omitempty"`
}

func (s UploadShippingEblbatchResponse) String() string {
	return tea.Prettify(s)
}

func (s UploadShippingEblbatchResponse) GoString() string {
	return s.String()
}

func (s *UploadShippingEblbatchResponse) SetReqMsgId(v string) *UploadShippingEblbatchResponse {
	s.ReqMsgId = &v
	return s
}

func (s *UploadShippingEblbatchResponse) SetResultCode(v string) *UploadShippingEblbatchResponse {
	s.ResultCode = &v
	return s
}

func (s *UploadShippingEblbatchResponse) SetResultMsg(v string) *UploadShippingEblbatchResponse {
	s.ResultMsg = &v
	return s
}

type UpdateShippingEblbatchstatusRequest struct {
	// OAuth模式下的授权token
	AuthToken         *string `json:"auth_token,omitempty" xml:"auth_token,omitempty"`
	ProductInstanceId *string `json:"product_instance_id,omitempty" xml:"product_instance_id,omitempty"`
	// 电子提单批次号
	BatchNo *string `json:"batch_no,omitempty" xml:"batch_no,omitempty" require:"true"`
	// 批次下的提单状态变更明细
	EblStatusDetails []*EblStatusDetail `json:"ebl_status_details,omitempty" xml:"ebl_status_details,omitempty" require:"true" type:"Repeated"`
}

func (s UpdateShippingEblbatchstatusRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateShippingEblbatchstatusRequest) GoString() string {
	return s.String()
}

func (s *UpdateShippingEblbatchstatusRequest) SetAuthToken(v string) *UpdateShippingEblbatchstatusRequest {
	s.AuthToken = &v
	return s
}

func (s *UpdateShippingEblbatchstatusRequest) SetProductInstanceId(v string) *UpdateShippingEblbatchstatusRequest {
	s.ProductInstanceId = &v
	return s
}

func (s *UpdateShippingEblbatchstatusRequest) SetBatchNo(v string) *UpdateShippingEblbatchstatusRequest {
	s.BatchNo = &v
	return s
}

func (s *UpdateShippingEblbatchstatusRequest) SetEblStatusDetails(v []*EblStatusDetail) *UpdateShippingEblbatchstatusRequest {
	s.EblStatusDetails = v
	return s
}

type UpdateShippingEblbatchstatusResponse struct {
	// 请求唯一ID，用于链路跟踪和问题排查
	ReqMsgId *string `json:"req_msg_id,omitempty" xml:"req_msg_id,omitempty"`
	// 结果码，一般OK表示调用成功
	ResultCode *string `json:"result_code,omitempty" xml:"result_code,omitempty"`
	// 异常信息的文本描述
	ResultMsg *string `json:"result_msg,omitempty" xml:"result_msg,omitempty"`
}

func (s UpdateShippingEblbatchstatusResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateShippingEblbatchstatusResponse) GoString() string {
	return s.String()
}

func (s *UpdateShippingEblbatchstatusResponse) SetReqMsgId(v string) *UpdateShippingEblbatchstatusResponse {
	s.ReqMsgId = &v
	return s
}

func (s *UpdateShippingEblbatchstatusResponse) SetResultCode(v string) *UpdateShippingEblbatchstatusResponse {
	s.ResultCode = &v
	return s
}

func (s *UpdateShippingEblbatchstatusResponse) SetResultMsg(v string) *UpdateShippingEblbatchstatusResponse {
	s.ResultMsg = &v
	return s
}

type Client struct {
	Endpoint                *string
	RegionId                *string
	AccessKeyId             *string
	AccessKeySecret         *string
	Protocol                *string
	UserAgent               *string
	ReadTimeout             *int
	ConnectTimeout          *int
	HttpProxy               *string
	HttpsProxy              *string
	Socks5Proxy             *string
	Socks5NetWork           *string
	NoProxy                 *string
	MaxIdleConns            *int
	SecurityToken           *string
	MaxIdleTimeMillis       *int
	KeepAliveDurationMillis *int
	MaxRequests             *int
	MaxRequestsPerHost      *int
}

/**
 * Init client with Config
 * @param config config contains the necessary information to create a client
 */
func NewClient(config *Config) (*Client, error) {
	client := new(Client)
	err := client.Init(config)
	return client, err
}

func (client *Client) Init(config *Config) (_err error) {
	if tea.BoolValue(util.IsUnset(tea.ToMap(config))) {
		_err = tea.NewSDKError(map[string]interface{}{
			"code":    "ParameterMissing",
			"message": "'config' can not be unset",
		})
		return _err
	}

	client.AccessKeyId = config.AccessKeyId
	client.AccessKeySecret = config.AccessKeySecret
	client.SecurityToken = config.SecurityToken
	client.Endpoint = config.Endpoint
	client.Protocol = config.Protocol
	client.UserAgent = config.UserAgent
	client.ReadTimeout = util.DefaultNumber(config.ReadTimeout, tea.Int(20000))
	client.ConnectTimeout = util.DefaultNumber(config.ConnectTimeout, tea.Int(20000))
	client.HttpProxy = config.HttpProxy
	client.HttpsProxy = config.HttpsProxy
	client.NoProxy = config.NoProxy
	client.Socks5Proxy = config.Socks5Proxy
	client.Socks5NetWork = config.Socks5NetWork
	client.MaxIdleConns = util.DefaultNumber(config.MaxIdleConns, tea.Int(60000))
	client.MaxIdleTimeMillis = util.DefaultNumber(config.MaxIdleTimeMillis, tea.Int(5))
	client.KeepAliveDurationMillis = util.DefaultNumber(config.KeepAliveDurationMillis, tea.Int(5000))
	client.MaxRequests = util.DefaultNumber(config.MaxRequests, tea.Int(100))
	client.MaxRequestsPerHost = util.DefaultNumber(config.MaxRequestsPerHost, tea.Int(100))
	return nil
}

/**
 * Encapsulate the request and invoke the network
 * @param action api name
 * @param protocol http or https
 * @param method e.g. GET
 * @param pathname pathname of every api
 * @param request which contains request params
 * @param runtime which controls some details of call api, such as retry times
 * @return the response
 */
func (client *Client) DoRequest(version *string, action *string, protocol *string, method *string, pathname *string, request map[string]interface{}, headers map[string]*string, runtime *util.RuntimeOptions) (_result map[string]interface{}, _err error) {
	_err = tea.Validate(runtime)
	if _err != nil {
		return _result, _err
	}
	_runtime := map[string]interface{}{
		"timeouted":               "retry",
		"readTimeout":             tea.IntValue(util.DefaultNumber(runtime.ReadTimeout, client.ReadTimeout)),
		"connectTimeout":          tea.IntValue(util.DefaultNumber(runtime.ConnectTimeout, client.ConnectTimeout)),
		"httpProxy":               tea.StringValue(util.DefaultString(runtime.HttpProxy, client.HttpProxy)),
		"httpsProxy":              tea.StringValue(util.DefaultString(runtime.HttpsProxy, client.HttpsProxy)),
		"noProxy":                 tea.StringValue(util.DefaultString(runtime.NoProxy, client.NoProxy)),
		"maxIdleConns":            tea.IntValue(util.DefaultNumber(runtime.MaxIdleConns, client.MaxIdleConns)),
		"maxIdleTimeMillis":       tea.IntValue(client.MaxIdleTimeMillis),
		"keepAliveDurationMillis": tea.IntValue(client.KeepAliveDurationMillis),
		"maxRequests":             tea.IntValue(client.MaxRequests),
		"maxRequestsPerHost":      tea.IntValue(client.MaxRequestsPerHost),
		"retry": map[string]interface{}{
			"retryable":   tea.BoolValue(runtime.Autoretry),
			"maxAttempts": tea.IntValue(util.DefaultNumber(runtime.MaxAttempts, tea.Int(3))),
		},
		"backoff": map[string]interface{}{
			"policy": tea.StringValue(util.DefaultString(runtime.BackoffPolicy, tea.String("no"))),
			"period": tea.IntValue(util.DefaultNumber(runtime.BackoffPeriod, tea.Int(1))),
		},
		"ignoreSSL": tea.BoolValue(runtime.IgnoreSSL),
	}

	_resp := make(map[string]interface{})
	for _retryTimes := 0; tea.BoolValue(tea.AllowRetry(_runtime["retry"], tea.Int(_retryTimes))); _retryTimes++ {
		if _retryTimes > 0 {
			_backoffTime := tea.GetBackoffTime(_runtime["backoff"], tea.Int(_retryTimes))
			if tea.IntValue(_backoffTime) > 0 {
				tea.Sleep(_backoffTime)
			}
		}

		_resp, _err = func() (map[string]interface{}, error) {
			request_ := tea.NewRequest()
			request_.Protocol = util.DefaultString(client.Protocol, protocol)
			request_.Method = method
			request_.Pathname = pathname
			request_.Query = map[string]*string{
				"method":           action,
				"version":          version,
				"sign_type":        tea.String("HmacSHA1"),
				"req_time":         antchainutil.GetTimestamp(),
				"req_msg_id":       antchainutil.GetNonce(),
				"access_key":       client.AccessKeyId,
				"base_sdk_version": tea.String("TeaSDK-2.0"),
				"sdk_version":      tea.String("1.3.106"),
			}
			if !tea.BoolValue(util.Empty(client.SecurityToken)) {
				request_.Query["security_token"] = client.SecurityToken
			}

			request_.Headers = tea.Merge(map[string]*string{
				"host":       util.DefaultString(client.Endpoint, tea.String("openapi.antchain.antgroup.com")),
				"user-agent": util.GetUserAgent(client.UserAgent),
			}, headers)
			tmp := util.AnyifyMapValue(rpcutil.Query(request))
			request_.Body = tea.ToReader(util.ToFormString(tmp))
			request_.Headers["content-type"] = tea.String("application/x-www-form-urlencoded")
			signedParam := tea.Merge(request_.Query,
				rpcutil.Query(request))
			request_.Query["sign"] = antchainutil.GetSignature(signedParam, client.AccessKeySecret)
			response_, _err := tea.DoRequest(request_, _runtime)
			if _err != nil {
				return _result, _err
			}
			raw, _err := util.ReadAsString(response_.Body)
			if _err != nil {
				return _result, _err
			}

			obj := util.ParseJSON(raw)
			res := util.AssertAsMap(obj)
			resp := util.AssertAsMap(res["response"])
			if tea.BoolValue(antchainutil.HasError(raw, client.AccessKeySecret)) {
				_err = tea.NewSDKError(map[string]interface{}{
					"message": resp["result_msg"],
					"data":    resp,
					"code":    resp["result_code"],
				})
				return _result, _err
			}

			_result = resp
			return _result, _err
		}()
		if !tea.BoolValue(tea.Retryable(_err)) {
			break
		}
	}

	return _resp, _err
}

/**
 * Description: 标准化应收账单
 * Summary: 标准化应收账单
 */
func (client *Client) CreateReceivableBill(request *CreateReceivableBillRequest) (_result *CreateReceivableBillResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateReceivableBillResponse{}
	_body, _err := client.CreateReceivableBillEx(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * Description: 标准化应收账单
 * Summary: 标准化应收账单
 */
func (client *Client) CreateReceivableBillEx(request *CreateReceivableBillRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *CreateReceivableBillResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &CreateReceivableBillResponse{}
	_body, _err := client.DoRequest(tea.String("1.0"), tea.String("digital.logistic.receivable.bill.create"), tea.String("HTTPS"), tea.String("POST"), tea.String("/gateway.do"), tea.ToMap(request), headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * Description: 凭证发行
 * Summary: 凭证发行
 */
func (client *Client) CreateStandardVoucher(request *CreateStandardVoucherRequest) (_result *CreateStandardVoucherResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateStandardVoucherResponse{}
	_body, _err := client.CreateStandardVoucherEx(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * Description: 凭证发行
 * Summary: 凭证发行
 */
func (client *Client) CreateStandardVoucherEx(request *CreateStandardVoucherRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *CreateStandardVoucherResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &CreateStandardVoucherResponse{}
	_body, _err := client.DoRequest(tea.String("1.0"), tea.String("digital.logistic.standard.voucher.create"), tea.String("HTTPS"), tea.String("POST"), tea.String("/gateway.do"), tea.ToMap(request), headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * Description: 无车承运平台为自身申请DIS分布式数字身份
 * Summary: 无车承运平台DIS分布式数字身份申请
 */
func (client *Client) CreatePlatformDid(request *CreatePlatformDidRequest) (_result *CreatePlatformDidResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreatePlatformDidResponse{}
	_body, _err := client.CreatePlatformDidEx(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * Description: 无车承运平台为自身申请DIS分布式数字身份
 * Summary: 无车承运平台DIS分布式数字身份申请
 */
func (client *Client) CreatePlatformDidEx(request *CreatePlatformDidRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *CreatePlatformDidResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &CreatePlatformDidResponse{}
	_body, _err := client.DoRequest(tea.String("1.0"), tea.String("digital.logistic.platform.did.create"), tea.String("HTTPS"), tea.String("POST"), tea.String("/gateway.do"), tea.ToMap(request), headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * Description: 代理申请无车承运平台的DIS分布式数字身份。无车saas平台需要为平台中的各个无车承运平台创建dis时需要使用代理创建的模式为其创建分布式数字身份。代理申请分布式数字身份的前置条件为自身需已有分布式数字身份。
 * Summary: 无车承运平台DIS分布式数字身份代理申请
 */
func (client *Client) CreateAgentDid(request *CreateAgentDidRequest) (_result *CreateAgentDidResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateAgentDidResponse{}
	_body, _err := client.CreateAgentDidEx(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * Description: 代理申请无车承运平台的DIS分布式数字身份。无车saas平台需要为平台中的各个无车承运平台创建dis时需要使用代理创建的模式为其创建分布式数字身份。代理申请分布式数字身份的前置条件为自身需已有分布式数字身份。
 * Summary: 无车承运平台DIS分布式数字身份代理申请
 */
func (client *Client) CreateAgentDidEx(request *CreateAgentDidRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *CreateAgentDidResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &CreateAgentDidResponse{}
	_body, _err := client.DoRequest(tea.String("1.0"), tea.String("digital.logistic.agent.did.create"), tea.String("HTTPS"), tea.String("POST"), tea.String("/gateway.do"), tea.ToMap(request), headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * Description: 无车承运平台为货主申请联盟中唯一标识货主的DIS分布式数字身份,会对企业信息做核验，同一个企业多次申请dis将会返回相同的分布式数字身份。为货主申请分布式数字身份的前置条件为无车承运平台需已有分布式数字身份。
 * Summary: 货主DIS分布式数字身份申请
 */
func (client *Client) CreateConsignorDis(request *CreateConsignorDisRequest) (_result *CreateConsignorDisResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateConsignorDisResponse{}
	_body, _err := client.CreateConsignorDisEx(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * Description: 无车承运平台为货主申请联盟中唯一标识货主的DIS分布式数字身份,会对企业信息做核验，同一个企业多次申请dis将会返回相同的分布式数字身份。为货主申请分布式数字身份的前置条件为无车承运平台需已有分布式数字身份。
 * Summary: 货主DIS分布式数字身份申请
 */
func (client *Client) CreateConsignorDisEx(request *CreateConsignorDisRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *CreateConsignorDisResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &CreateConsignorDisResponse{}
	_body, _err := client.DoRequest(tea.String("1.0"), tea.String("digital.logistic.consignor.dis.create"), tea.String("HTTPS"), tea.String("POST"), tea.String("/gateway.do"), tea.ToMap(request), headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * Description: 无车承运平台为司机申请DIS分布式数字身份,会对司机做实人认证，同一个司机多次申请dis会返回相同的分布式数字身份。为司机申请分布式数字身份的前置条件为无车承运平台需已有分布式数字身份。
 * Summary: 司机DIS分布式数字身份申请
 */
func (client *Client) CreateDriverDis(request *CreateDriverDisRequest) (_result *CreateDriverDisResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateDriverDisResponse{}
	_body, _err := client.CreateDriverDisEx(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * Description: 无车承运平台为司机申请DIS分布式数字身份,会对司机做实人认证，同一个司机多次申请dis会返回相同的分布式数字身份。为司机申请分布式数字身份的前置条件为无车承运平台需已有分布式数字身份。
 * Summary: 司机DIS分布式数字身份申请
 */
func (client *Client) CreateDriverDisEx(request *CreateDriverDisRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *CreateDriverDisResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &CreateDriverDisResponse{}
	_body, _err := client.DoRequest(tea.String("1.0"), tea.String("digital.logistic.driver.dis.create"), tea.String("HTTPS"), tea.String("POST"), tea.String("/gateway.do"), tea.ToMap(request), headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * Description: 货源订单创建，货主发货给平台时的订单信息
 * Summary: 货源订单创建
 */
func (client *Client) CreateCargoOrder(request *CreateCargoOrderRequest) (_result *CreateCargoOrderResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateCargoOrderResponse{}
	_body, _err := client.CreateCargoOrderEx(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * Description: 货源订单创建，货主发货给平台时的订单信息
 * Summary: 货源订单创建
 */
func (client *Client) CreateCargoOrderEx(request *CreateCargoOrderRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *CreateCargoOrderResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &CreateCargoOrderResponse{}
	_body, _err := client.DoRequest(tea.String("1.0"), tea.String("digital.logistic.cargo.order.create"), tea.String("HTTPS"), tea.String("POST"), tea.String("/gateway.do"), tea.ToMap(request), headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * Description: 货源支付订单创建,货源订单需存在。
 * Summary: 货源支付订单创建
 */
func (client *Client) CreateCargoPay(request *CreateCargoPayRequest) (_result *CreateCargoPayResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateCargoPayResponse{}
	_body, _err := client.CreateCargoPayEx(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * Description: 货源支付订单创建,货源订单需存在。
 * Summary: 货源支付订单创建
 */
func (client *Client) CreateCargoPayEx(request *CreateCargoPayRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *CreateCargoPayResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &CreateCargoPayResponse{}
	_body, _err := client.DoRequest(tea.String("1.0"), tea.String("digital.logistic.cargo.pay.create"), tea.String("HTTPS"), tea.String("POST"), tea.String("/gateway.do"), tea.ToMap(request), headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * Description: 物流平台将运单指派给下游（承运商、司机）时产生
 * Summary: 运单创建
 */
func (client *Client) CreateWaybillOrder(request *CreateWaybillOrderRequest) (_result *CreateWaybillOrderResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateWaybillOrderResponse{}
	_body, _err := client.CreateWaybillOrderEx(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * Description: 物流平台将运单指派给下游（承运商、司机）时产生
 * Summary: 运单创建
 */
func (client *Client) CreateWaybillOrderEx(request *CreateWaybillOrderRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *CreateWaybillOrderResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &CreateWaybillOrderResponse{}
	_body, _err := client.DoRequest(tea.String("1.0"), tea.String("digital.logistic.waybill.order.create"), tea.String("HTTPS"), tea.String("POST"), tea.String("/gateway.do"), tea.ToMap(request), headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * Description: 创建物流支付订单，运单需已存在并且未结束
 * Summary: 运单支付订单创建
 */
func (client *Client) CreateWaybillPay(request *CreateWaybillPayRequest) (_result *CreateWaybillPayResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateWaybillPayResponse{}
	_body, _err := client.CreateWaybillPayEx(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * Description: 创建物流支付订单，运单需已存在并且未结束
 * Summary: 运单支付订单创建
 */
func (client *Client) CreateWaybillPayEx(request *CreateWaybillPayRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *CreateWaybillPayResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &CreateWaybillPayResponse{}
	_body, _err := client.DoRequest(tea.String("1.0"), tea.String("digital.logistic.waybill.pay.create"), tea.String("HTTPS"), tea.String("POST"), tea.String("/gateway.do"), tea.ToMap(request), headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * Description: 用于上报第三方的轨迹位置信息，运单需已存在并且未结束
 * Summary: 第三方位置信息上报
 */
func (client *Client) ImportWaybillLocation(request *ImportWaybillLocationRequest) (_result *ImportWaybillLocationResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ImportWaybillLocationResponse{}
	_body, _err := client.ImportWaybillLocationEx(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * Description: 用于上报第三方的轨迹位置信息，运单需已存在并且未结束
 * Summary: 第三方位置信息上报
 */
func (client *Client) ImportWaybillLocationEx(request *ImportWaybillLocationRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ImportWaybillLocationResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &ImportWaybillLocationResponse{}
	_body, _err := client.DoRequest(tea.String("1.0"), tea.String("digital.logistic.waybill.location.import"), tea.String("HTTPS"), tea.String("POST"), tea.String("/gateway.do"), tea.ToMap(request), headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * Description: 发票订单创建
 * Summary: 发票订单创建
 */
func (client *Client) CreateWaybillBill(request *CreateWaybillBillRequest) (_result *CreateWaybillBillResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateWaybillBillResponse{}
	_body, _err := client.CreateWaybillBillEx(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * Description: 发票订单创建
 * Summary: 发票订单创建
 */
func (client *Client) CreateWaybillBillEx(request *CreateWaybillBillRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *CreateWaybillBillResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &CreateWaybillBillResponse{}
	_body, _err := client.DoRequest(tea.String("1.0"), tea.String("digital.logistic.waybill.bill.create"), tea.String("HTTPS"), tea.String("POST"), tea.String("/gateway.do"), tea.ToMap(request), headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * Description: 查询物流运单状态
 * Summary: 查询运单状态
 */
func (client *Client) QueryWaybillStatus(request *QueryWaybillStatusRequest) (_result *QueryWaybillStatusResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &QueryWaybillStatusResponse{}
	_body, _err := client.QueryWaybillStatusEx(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * Description: 查询物流运单状态
 * Summary: 查询运单状态
 */
func (client *Client) QueryWaybillStatusEx(request *QueryWaybillStatusRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *QueryWaybillStatusResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &QueryWaybillStatusResponse{}
	_body, _err := client.DoRequest(tea.String("1.0"), tea.String("digital.logistic.waybill.status.query"), tea.String("HTTPS"), tea.String("POST"), tea.String("/gateway.do"), tea.ToMap(request), headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * Description: 运单完成表示整个运单的生命周期结束的最终状态，一般为和司机已进行确认结算后调用
完成后的运单无法再调用支付订单与轨迹位置上传接口。
 * Summary: 运单完成
*/
func (client *Client) FinishWaybillOrder(request *FinishWaybillOrderRequest) (_result *FinishWaybillOrderResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &FinishWaybillOrderResponse{}
	_body, _err := client.FinishWaybillOrderEx(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * Description: 运单完成表示整个运单的生命周期结束的最终状态，一般为和司机已进行确认结算后调用
完成后的运单无法再调用支付订单与轨迹位置上传接口。
 * Summary: 运单完成
*/
func (client *Client) FinishWaybillOrderEx(request *FinishWaybillOrderRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *FinishWaybillOrderResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &FinishWaybillOrderResponse{}
	_body, _err := client.DoRequest(tea.String("1.0"), tea.String("digital.logistic.waybill.order.finish"), tea.String("HTTPS"), tea.String("POST"), tea.String("/gateway.do"), tea.ToMap(request), headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * Description: 货源支付订单创建,货源订单需存在。
 * Summary: 货源支付订单创建
 */
func (client *Client) CreateCargoPayorder(request *CreateCargoPayorderRequest) (_result *CreateCargoPayorderResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateCargoPayorderResponse{}
	_body, _err := client.CreateCargoPayorderEx(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * Description: 货源支付订单创建,货源订单需存在。
 * Summary: 货源支付订单创建
 */
func (client *Client) CreateCargoPayorderEx(request *CreateCargoPayorderRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *CreateCargoPayorderResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &CreateCargoPayorderResponse{}
	_body, _err := client.DoRequest(tea.String("1.0"), tea.String("digital.logistic.cargo.payorder.create"), tea.String("HTTPS"), tea.String("POST"), tea.String("/gateway.do"), tea.ToMap(request), headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * Description: 当货物到达目的地时，调用此接口
 * Summary: 运单运输完成
 */
func (client *Client) SaveWaybillOrder(request *SaveWaybillOrderRequest) (_result *SaveWaybillOrderResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &SaveWaybillOrderResponse{}
	_body, _err := client.SaveWaybillOrderEx(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * Description: 当货物到达目的地时，调用此接口
 * Summary: 运单运输完成
 */
func (client *Client) SaveWaybillOrderEx(request *SaveWaybillOrderRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *SaveWaybillOrderResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &SaveWaybillOrderResponse{}
	_body, _err := client.DoRequest(tea.String("1.0"), tea.String("digital.logistic.waybill.order.save"), tea.String("HTTPS"), tea.String("POST"), tea.String("/gateway.do"), tea.ToMap(request), headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * Description: 运单关闭
 * Summary: 运单关闭
 */
func (client *Client) CloseWaybillOrder(request *CloseWaybillOrderRequest) (_result *CloseWaybillOrderResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CloseWaybillOrderResponse{}
	_body, _err := client.CloseWaybillOrderEx(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * Description: 运单关闭
 * Summary: 运单关闭
 */
func (client *Client) CloseWaybillOrderEx(request *CloseWaybillOrderRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *CloseWaybillOrderResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &CloseWaybillOrderResponse{}
	_body, _err := client.DoRequest(tea.String("1.0"), tea.String("digital.logistic.waybill.order.close"), tea.String("HTTPS"), tea.String("POST"), tea.String("/gateway.do"), tea.ToMap(request), headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * Description: 完成物流运单
 * Summary: 完成物流运单
 */
func (client *Client) FinishFinanceWaybill(request *FinishFinanceWaybillRequest) (_result *FinishFinanceWaybillResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &FinishFinanceWaybillResponse{}
	_body, _err := client.FinishFinanceWaybillEx(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * Description: 完成物流运单
 * Summary: 完成物流运单
 */
func (client *Client) FinishFinanceWaybillEx(request *FinishFinanceWaybillRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *FinishFinanceWaybillResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &FinishFinanceWaybillResponse{}
	_body, _err := client.DoRequest(tea.String("1.0"), tea.String("digital.logistic.finance.waybill.finish"), tea.String("HTTPS"), tea.String("POST"), tea.String("/gateway.do"), tea.ToMap(request), headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * Description: 运输完成
 * Summary: 运输完成
 */
func (client *Client) FinishFinanceTransport(request *FinishFinanceTransportRequest) (_result *FinishFinanceTransportResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &FinishFinanceTransportResponse{}
	_body, _err := client.FinishFinanceTransportEx(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * Description: 运输完成
 * Summary: 运输完成
 */
func (client *Client) FinishFinanceTransportEx(request *FinishFinanceTransportRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *FinishFinanceTransportResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &FinishFinanceTransportResponse{}
	_body, _err := client.DoRequest(tea.String("1.0"), tea.String("digital.logistic.finance.transport.finish"), tea.String("HTTPS"), tea.String("POST"), tea.String("/gateway.do"), tea.ToMap(request), headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * Description: 对运单信息项需更新时调用
 * Summary: 更新运单
 */
func (client *Client) UpdateFinanceWaybill(request *UpdateFinanceWaybillRequest) (_result *UpdateFinanceWaybillResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpdateFinanceWaybillResponse{}
	_body, _err := client.UpdateFinanceWaybillEx(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * Description: 对运单信息项需更新时调用
 * Summary: 更新运单
 */
func (client *Client) UpdateFinanceWaybillEx(request *UpdateFinanceWaybillRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *UpdateFinanceWaybillResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &UpdateFinanceWaybillResponse{}
	_body, _err := client.DoRequest(tea.String("1.0"), tea.String("digital.logistic.finance.waybill.update"), tea.String("HTTPS"), tea.String("POST"), tea.String("/gateway.do"), tea.ToMap(request), headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * Description: 当司机到达货主装货地后，调用此接口
 * Summary: 起运运单
 */
func (client *Client) StartFinanceWaybill(request *StartFinanceWaybillRequest) (_result *StartFinanceWaybillResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &StartFinanceWaybillResponse{}
	_body, _err := client.StartFinanceWaybillEx(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * Description: 当司机到达货主装货地后，调用此接口
 * Summary: 起运运单
 */
func (client *Client) StartFinanceWaybillEx(request *StartFinanceWaybillRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *StartFinanceWaybillResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &StartFinanceWaybillResponse{}
	_body, _err := client.DoRequest(tea.String("1.0"), tea.String("digital.logistic.finance.waybill.start"), tea.String("HTTPS"), tea.String("POST"), tea.String("/gateway.do"), tea.ToMap(request), headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * Description: 创建承运商账户
 * Summary: 创建承运商账户
 */
func (client *Client) CreateCaptainDis(request *CreateCaptainDisRequest) (_result *CreateCaptainDisResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateCaptainDisResponse{}
	_body, _err := client.CreateCaptainDisEx(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * Description: 创建承运商账户
 * Summary: 创建承运商账户
 */
func (client *Client) CreateCaptainDisEx(request *CreateCaptainDisRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *CreateCaptainDisResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &CreateCaptainDisResponse{}
	_body, _err := client.DoRequest(tea.String("1.0"), tea.String("digital.logistic.captain.dis.create"), tea.String("HTTPS"), tea.String("POST"), tea.String("/gateway.do"), tea.ToMap(request), headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * Description: 创建货主账单
 * Summary: 货主账单创建
 */
func (client *Client) CreateCargowaybillBill(request *CreateCargowaybillBillRequest) (_result *CreateCargowaybillBillResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateCargowaybillBillResponse{}
	_body, _err := client.CreateCargowaybillBillEx(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * Description: 创建货主账单
 * Summary: 货主账单创建
 */
func (client *Client) CreateCargowaybillBillEx(request *CreateCargowaybillBillRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *CreateCargowaybillBillResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &CreateCargowaybillBillResponse{}
	_body, _err := client.DoRequest(tea.String("1.0"), tea.String("digital.logistic.cargowaybill.bill.create"), tea.String("HTTPS"), tea.String("POST"), tea.String("/gateway.do"), tea.ToMap(request), headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * Description: 确认货主账单
 * Summary: 货主账单确认
 */
func (client *Client) ConfirmCargowaybillBill(request *ConfirmCargowaybillBillRequest) (_result *ConfirmCargowaybillBillResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ConfirmCargowaybillBillResponse{}
	_body, _err := client.ConfirmCargowaybillBillEx(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * Description: 确认货主账单
 * Summary: 货主账单确认
 */
func (client *Client) ConfirmCargowaybillBillEx(request *ConfirmCargowaybillBillRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ConfirmCargowaybillBillResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &ConfirmCargowaybillBillResponse{}
	_body, _err := client.DoRequest(tea.String("1.0"), tea.String("digital.logistic.cargowaybill.bill.confirm"), tea.String("HTTPS"), tea.String("POST"), tea.String("/gateway.do"), tea.ToMap(request), headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * Description: 结算货主账单
 * Summary: 货主账单结算
 */
func (client *Client) CreateCargowaybillBillsettle(request *CreateCargowaybillBillsettleRequest) (_result *CreateCargowaybillBillsettleResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateCargowaybillBillsettleResponse{}
	_body, _err := client.CreateCargowaybillBillsettleEx(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * Description: 结算货主账单
 * Summary: 货主账单结算
 */
func (client *Client) CreateCargowaybillBillsettleEx(request *CreateCargowaybillBillsettleRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *CreateCargowaybillBillsettleResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &CreateCargowaybillBillsettleResponse{}
	_body, _err := client.DoRequest(tea.String("1.0"), tea.String("digital.logistic.cargowaybill.billsettle.create"), tea.String("HTTPS"), tea.String("POST"), tea.String("/gateway.do"), tea.ToMap(request), headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * Description: 运单平台did更新
 * Summary: 运单平台did更新
 */
func (client *Client) UpdateWaybillorderPlatformdid(request *UpdateWaybillorderPlatformdidRequest) (_result *UpdateWaybillorderPlatformdidResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpdateWaybillorderPlatformdidResponse{}
	_body, _err := client.UpdateWaybillorderPlatformdidEx(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * Description: 运单平台did更新
 * Summary: 运单平台did更新
 */
func (client *Client) UpdateWaybillorderPlatformdidEx(request *UpdateWaybillorderPlatformdidRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *UpdateWaybillorderPlatformdidResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &UpdateWaybillorderPlatformdidResponse{}
	_body, _err := client.DoRequest(tea.String("1.0"), tea.String("digital.logistic.waybillorder.platformdid.update"), tea.String("HTTPS"), tea.String("POST"), tea.String("/gateway.do"), tea.ToMap(request), headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * Description: 更新货主账单
 * Summary: 货主账单更新
 */
func (client *Client) UpdateCargowaybillBill(request *UpdateCargowaybillBillRequest) (_result *UpdateCargowaybillBillResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpdateCargowaybillBillResponse{}
	_body, _err := client.UpdateCargowaybillBillEx(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * Description: 更新货主账单
 * Summary: 货主账单更新
 */
func (client *Client) UpdateCargowaybillBillEx(request *UpdateCargowaybillBillRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *UpdateCargowaybillBillResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &UpdateCargowaybillBillResponse{}
	_body, _err := client.DoRequest(tea.String("1.0"), tea.String("digital.logistic.cargowaybill.bill.update"), tea.String("HTTPS"), tea.String("POST"), tea.String("/gateway.do"), tea.ToMap(request), headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * Description: 分布式申请did账户集合，可通过此接口申请个人类型角色为货主、承运商、司机的分布式数字身份did；可申请企业类型角色为货主、网络货运平台、道路运输企业/3pl、承运商、子平台的分布式数字身份did。注：接口可允许多次调用，但每次调用只允许申请一个角色，不允许一次调用申请多个角色
 * Summary: 分布式数字身份申请did集合
 */
func (client *Client) CreateDisDid(request *CreateDisDidRequest) (_result *CreateDisDidResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateDisDidResponse{}
	_body, _err := client.CreateDisDidEx(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * Description: 分布式申请did账户集合，可通过此接口申请个人类型角色为货主、承运商、司机的分布式数字身份did；可申请企业类型角色为货主、网络货运平台、道路运输企业/3pl、承运商、子平台的分布式数字身份did。注：接口可允许多次调用，但每次调用只允许申请一个角色，不允许一次调用申请多个角色
 * Summary: 分布式数字身份申请did集合
 */
func (client *Client) CreateDisDidEx(request *CreateDisDidRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *CreateDisDidResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &CreateDisDidResponse{}
	_body, _err := client.DoRequest(tea.String("1.0"), tea.String("digital.logistic.dis.did.create"), tea.String("HTTPS"), tea.String("POST"), tea.String("/gateway.do"), tea.ToMap(request), headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * Description: 3PL运输合同上传
 * Summary: 3PL运输合同上传
 */
func (client *Client) UploadTransportContract(request *UploadTransportContractRequest) (_result *UploadTransportContractResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UploadTransportContractResponse{}
	_body, _err := client.UploadTransportContractEx(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * Description: 3PL运输合同上传
 * Summary: 3PL运输合同上传
 */
func (client *Client) UploadTransportContractEx(request *UploadTransportContractRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *UploadTransportContractResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &UploadTransportContractResponse{}
	_body, _err := client.DoRequest(tea.String("1.0"), tea.String("digital.logistic.transport.contract.upload"), tea.String("HTTPS"), tea.String("POST"), tea.String("/gateway.do"), tea.ToMap(request), headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * Description: 3PL运输线路上传
 * Summary: 3PL运输线路上传
 */
func (client *Client) UploadTransportRoute(request *UploadTransportRouteRequest) (_result *UploadTransportRouteResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UploadTransportRouteResponse{}
	_body, _err := client.UploadTransportRouteEx(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * Description: 3PL运输线路上传
 * Summary: 3PL运输线路上传
 */
func (client *Client) UploadTransportRouteEx(request *UploadTransportRouteRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *UploadTransportRouteResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &UploadTransportRouteResponse{}
	_body, _err := client.DoRequest(tea.String("1.0"), tea.String("digital.logistic.transport.route.upload"), tea.String("HTTPS"), tea.String("POST"), tea.String("/gateway.do"), tea.ToMap(request), headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * Description: 3PL运单创建
 * Summary: 3PL运单创建
 */
func (client *Client) CreateTransportWaybill(request *CreateTransportWaybillRequest) (_result *CreateTransportWaybillResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateTransportWaybillResponse{}
	_body, _err := client.CreateTransportWaybillEx(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * Description: 3PL运单创建
 * Summary: 3PL运单创建
 */
func (client *Client) CreateTransportWaybillEx(request *CreateTransportWaybillRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *CreateTransportWaybillResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &CreateTransportWaybillResponse{}
	_body, _err := client.DoRequest(tea.String("1.0"), tea.String("digital.logistic.transport.waybill.create"), tea.String("HTTPS"), tea.String("POST"), tea.String("/gateway.do"), tea.ToMap(request), headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * Description: 3PL运单状态更新
 * Summary: 3PL运单状态更新
 */
func (client *Client) UpdateWaybillAction(request *UpdateWaybillActionRequest) (_result *UpdateWaybillActionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpdateWaybillActionResponse{}
	_body, _err := client.UpdateWaybillActionEx(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * Description: 3PL运单状态更新
 * Summary: 3PL运单状态更新
 */
func (client *Client) UpdateWaybillActionEx(request *UpdateWaybillActionRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *UpdateWaybillActionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &UpdateWaybillActionResponse{}
	_body, _err := client.DoRequest(tea.String("1.0"), tea.String("digital.logistic.waybill.action.update"), tea.String("HTTPS"), tea.String("POST"), tea.String("/gateway.do"), tea.ToMap(request), headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * Description: 3PL运单修改
 * Summary: 3PL运单修改
 */
func (client *Client) UpdateTransportWaybill(request *UpdateTransportWaybillRequest) (_result *UpdateTransportWaybillResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpdateTransportWaybillResponse{}
	_body, _err := client.UpdateTransportWaybillEx(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * Description: 3PL运单修改
 * Summary: 3PL运单修改
 */
func (client *Client) UpdateTransportWaybillEx(request *UpdateTransportWaybillRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *UpdateTransportWaybillResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &UpdateTransportWaybillResponse{}
	_body, _err := client.DoRequest(tea.String("1.0"), tea.String("digital.logistic.transport.waybill.update"), tea.String("HTTPS"), tea.String("POST"), tea.String("/gateway.do"), tea.ToMap(request), headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * Description: 3PL回单上传
 * Summary: 3PL回单上传
 */
func (client *Client) UploadTransportReceipt(request *UploadTransportReceiptRequest) (_result *UploadTransportReceiptResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UploadTransportReceiptResponse{}
	_body, _err := client.UploadTransportReceiptEx(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * Description: 3PL回单上传
 * Summary: 3PL回单上传
 */
func (client *Client) UploadTransportReceiptEx(request *UploadTransportReceiptRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *UploadTransportReceiptResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &UploadTransportReceiptResponse{}
	_body, _err := client.DoRequest(tea.String("1.0"), tea.String("digital.logistic.transport.receipt.upload"), tea.String("HTTPS"), tea.String("POST"), tea.String("/gateway.do"), tea.ToMap(request), headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * Description: 应收账单创建
 * Summary: 应收账单创建
 */
func (client *Client) CreateBillReceivablebill(request *CreateBillReceivablebillRequest) (_result *CreateBillReceivablebillResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateBillReceivablebillResponse{}
	_body, _err := client.CreateBillReceivablebillEx(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * Description: 应收账单创建
 * Summary: 应收账单创建
 */
func (client *Client) CreateBillReceivablebillEx(request *CreateBillReceivablebillRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *CreateBillReceivablebillResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &CreateBillReceivablebillResponse{}
	_body, _err := client.DoRequest(tea.String("1.0"), tea.String("digital.logistic.bill.receivablebill.create"), tea.String("HTTPS"), tea.String("POST"), tea.String("/gateway.do"), tea.ToMap(request), headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * Description: 应收账单状态更新
 * Summary: 应收账单状态更新
 */
func (client *Client) UpdateReceivablebillStatus(request *UpdateReceivablebillStatusRequest) (_result *UpdateReceivablebillStatusResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpdateReceivablebillStatusResponse{}
	_body, _err := client.UpdateReceivablebillStatusEx(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * Description: 应收账单状态更新
 * Summary: 应收账单状态更新
 */
func (client *Client) UpdateReceivablebillStatusEx(request *UpdateReceivablebillStatusRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *UpdateReceivablebillStatusResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &UpdateReceivablebillStatusResponse{}
	_body, _err := client.DoRequest(tea.String("1.0"), tea.String("digital.logistic.receivablebill.status.update"), tea.String("HTTPS"), tea.String("POST"), tea.String("/gateway.do"), tea.ToMap(request), headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * Description: 应收账单修改
 * Summary: 应收账单修改
 */
func (client *Client) UpdateBillReceivablebill(request *UpdateBillReceivablebillRequest) (_result *UpdateBillReceivablebillResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpdateBillReceivablebillResponse{}
	_body, _err := client.UpdateBillReceivablebillEx(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * Description: 应收账单修改
 * Summary: 应收账单修改
 */
func (client *Client) UpdateBillReceivablebillEx(request *UpdateBillReceivablebillRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *UpdateBillReceivablebillResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &UpdateBillReceivablebillResponse{}
	_body, _err := client.DoRequest(tea.String("1.0"), tea.String("digital.logistic.bill.receivablebill.update"), tea.String("HTTPS"), tea.String("POST"), tea.String("/gateway.do"), tea.ToMap(request), headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * Description: 路运发票创建
 * Summary: 路运发票创建
 */
func (client *Client) CreateHighwayInvoice(request *CreateHighwayInvoiceRequest) (_result *CreateHighwayInvoiceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateHighwayInvoiceResponse{}
	_body, _err := client.CreateHighwayInvoiceEx(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * Description: 路运发票创建
 * Summary: 路运发票创建
 */
func (client *Client) CreateHighwayInvoiceEx(request *CreateHighwayInvoiceRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *CreateHighwayInvoiceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &CreateHighwayInvoiceResponse{}
	_body, _err := client.DoRequest(tea.String("1.0"), tea.String("digital.logistic.highway.invoice.create"), tea.String("HTTPS"), tea.String("POST"), tea.String("/gateway.do"), tea.ToMap(request), headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * Description: 查询运单信息
 * Summary: 运单信息查询
 */
func (client *Client) QueryWaybillInfo(request *QueryWaybillInfoRequest) (_result *QueryWaybillInfoResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &QueryWaybillInfoResponse{}
	_body, _err := client.QueryWaybillInfoEx(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * Description: 查询运单信息
 * Summary: 运单信息查询
 */
func (client *Client) QueryWaybillInfoEx(request *QueryWaybillInfoRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *QueryWaybillInfoResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &QueryWaybillInfoResponse{}
	_body, _err := client.DoRequest(tea.String("1.0"), tea.String("digital.logistic.waybill.info.query"), tea.String("HTTPS"), tea.String("POST"), tea.String("/gateway.do"), tea.ToMap(request), headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * Description: 根据实例id，查询实例状态
 * Summary: 实例状态查询
 */
func (client *Client) QueryBusinessInstancestatus(request *QueryBusinessInstancestatusRequest) (_result *QueryBusinessInstancestatusResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &QueryBusinessInstancestatusResponse{}
	_body, _err := client.QueryBusinessInstancestatusEx(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * Description: 根据实例id，查询实例状态
 * Summary: 实例状态查询
 */
func (client *Client) QueryBusinessInstancestatusEx(request *QueryBusinessInstancestatusRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *QueryBusinessInstancestatusResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &QueryBusinessInstancestatusResponse{}
	_body, _err := client.DoRequest(tea.String("1.0"), tea.String("digital.logistic.business.instancestatus.query"), tea.String("HTTPS"), tea.String("POST"), tea.String("/gateway.do"), tea.ToMap(request), headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * Description: 司机信用流转协议签署并开通流转能力
 * Summary: 司机信用流转协议签署并开通流转能力
 */
func (client *Client) OpenCreditDriver(request *OpenCreditDriverRequest) (_result *OpenCreditDriverResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &OpenCreditDriverResponse{}
	_body, _err := client.OpenCreditDriverEx(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * Description: 司机信用流转协议签署并开通流转能力
 * Summary: 司机信用流转协议签署并开通流转能力
 */
func (client *Client) OpenCreditDriverEx(request *OpenCreditDriverRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *OpenCreditDriverResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &OpenCreditDriverResponse{}
	_body, _err := client.DoRequest(tea.String("1.0"), tea.String("digital.logistic.credit.driver.open"), tea.String("HTTPS"), tea.String("POST"), tea.String("/gateway.do"), tea.ToMap(request), headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * Description: 查询司机信用流转能力开通情况
 * Summary: 查询司机信用流转能力开通情况
 */
func (client *Client) QueryCreditDriver(request *QueryCreditDriverRequest) (_result *QueryCreditDriverResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &QueryCreditDriverResponse{}
	_body, _err := client.QueryCreditDriverEx(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * Description: 查询司机信用流转能力开通情况
 * Summary: 查询司机信用流转能力开通情况
 */
func (client *Client) QueryCreditDriverEx(request *QueryCreditDriverRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *QueryCreditDriverResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &QueryCreditDriverResponse{}
	_body, _err := client.DoRequest(tea.String("1.0"), tea.String("digital.logistic.credit.driver.query"), tea.String("HTTPS"), tea.String("POST"), tea.String("/gateway.do"), tea.ToMap(request), headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * Description: 查询货主信用流转能力开通情况
 * Summary: 查询货主信用流转能力开通情况
 */
func (client *Client) QueryCreditConsignor(request *QueryCreditConsignorRequest) (_result *QueryCreditConsignorResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &QueryCreditConsignorResponse{}
	_body, _err := client.QueryCreditConsignorEx(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * Description: 查询货主信用流转能力开通情况
 * Summary: 查询货主信用流转能力开通情况
 */
func (client *Client) QueryCreditConsignorEx(request *QueryCreditConsignorRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *QueryCreditConsignorResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &QueryCreditConsignorResponse{}
	_body, _err := client.DoRequest(tea.String("1.0"), tea.String("digital.logistic.credit.consignor.query"), tea.String("HTTPS"), tea.String("POST"), tea.String("/gateway.do"), tea.ToMap(request), headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * Description: 查询货主信用流转额度
 * Summary: 查询货主信用流转额度
 */
func (client *Client) QueryCreditBalance(request *QueryCreditBalanceRequest) (_result *QueryCreditBalanceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &QueryCreditBalanceResponse{}
	_body, _err := client.QueryCreditBalanceEx(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * Description: 查询货主信用流转额度
 * Summary: 查询货主信用流转额度
 */
func (client *Client) QueryCreditBalanceEx(request *QueryCreditBalanceRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *QueryCreditBalanceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &QueryCreditBalanceResponse{}
	_body, _err := client.DoRequest(tea.String("1.0"), tea.String("digital.logistic.credit.balance.query"), tea.String("HTTPS"), tea.String("POST"), tea.String("/gateway.do"), tea.ToMap(request), headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * Description: 发行信用流转批次信息上传（只做信息上传，供真正web端发行时使用）
 * Summary: 发行信用流转批次信息上传
 */
func (client *Client) UploadCreditIssue(request *UploadCreditIssueRequest) (_result *UploadCreditIssueResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UploadCreditIssueResponse{}
	_body, _err := client.UploadCreditIssueEx(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * Description: 发行信用流转批次信息上传（只做信息上传，供真正web端发行时使用）
 * Summary: 发行信用流转批次信息上传
 */
func (client *Client) UploadCreditIssueEx(request *UploadCreditIssueRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *UploadCreditIssueResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &UploadCreditIssueResponse{}
	_body, _err := client.DoRequest(tea.String("1.0"), tea.String("digital.logistic.credit.issue.upload"), tea.String("HTTPS"), tea.String("POST"), tea.String("/gateway.do"), tea.ToMap(request), headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * Description: 信用流转批次状态查询
 * Summary: 信用流转批次状态查询
 */
func (client *Client) QueryCreditIssuebatchstatus(request *QueryCreditIssuebatchstatusRequest) (_result *QueryCreditIssuebatchstatusResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &QueryCreditIssuebatchstatusResponse{}
	_body, _err := client.QueryCreditIssuebatchstatusEx(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * Description: 信用流转批次状态查询
 * Summary: 信用流转批次状态查询
 */
func (client *Client) QueryCreditIssuebatchstatusEx(request *QueryCreditIssuebatchstatusRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *QueryCreditIssuebatchstatusResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &QueryCreditIssuebatchstatusResponse{}
	_body, _err := client.DoRequest(tea.String("1.0"), tea.String("digital.logistic.credit.issuebatchstatus.query"), tea.String("HTTPS"), tea.String("POST"), tea.String("/gateway.do"), tea.ToMap(request), headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * Description: 支付批次撤销
 * Summary: 支付批次撤销
 */
func (client *Client) CancelCreditIssuebatch(request *CancelCreditIssuebatchRequest) (_result *CancelCreditIssuebatchResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CancelCreditIssuebatchResponse{}
	_body, _err := client.CancelCreditIssuebatchEx(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * Description: 支付批次撤销
 * Summary: 支付批次撤销
 */
func (client *Client) CancelCreditIssuebatchEx(request *CancelCreditIssuebatchRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *CancelCreditIssuebatchResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &CancelCreditIssuebatchResponse{}
	_body, _err := client.DoRequest(tea.String("1.0"), tea.String("digital.logistic.credit.issuebatch.cancel"), tea.String("HTTPS"), tea.String("POST"), tea.String("/gateway.do"), tea.ToMap(request), headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * Description: 根据id查询信用凭证信息
 * Summary: 根据id查询信用凭证信息
 */
func (client *Client) QueryCreditIssuebyid(request *QueryCreditIssuebyidRequest) (_result *QueryCreditIssuebyidResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &QueryCreditIssuebyidResponse{}
	_body, _err := client.QueryCreditIssuebyidEx(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * Description: 根据id查询信用凭证信息
 * Summary: 根据id查询信用凭证信息
 */
func (client *Client) QueryCreditIssuebyidEx(request *QueryCreditIssuebyidRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *QueryCreditIssuebyidResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &QueryCreditIssuebyidResponse{}
	_body, _err := client.DoRequest(tea.String("1.0"), tea.String("digital.logistic.credit.issuebyid.query"), tea.String("HTTPS"), tea.String("POST"), tea.String("/gateway.do"), tea.ToMap(request), headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * Description: 根据时间查询信用凭证信息
 * Summary: 根据时间查询信用凭证信息
 */
func (client *Client) QueryCreditIssuebytime(request *QueryCreditIssuebytimeRequest) (_result *QueryCreditIssuebytimeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &QueryCreditIssuebytimeResponse{}
	_body, _err := client.QueryCreditIssuebytimeEx(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * Description: 根据时间查询信用凭证信息
 * Summary: 根据时间查询信用凭证信息
 */
func (client *Client) QueryCreditIssuebytimeEx(request *QueryCreditIssuebytimeRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *QueryCreditIssuebytimeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &QueryCreditIssuebytimeResponse{}
	_body, _err := client.DoRequest(tea.String("1.0"), tea.String("digital.logistic.credit.issuebytime.query"), tea.String("HTTPS"), tea.String("POST"), tea.String("/gateway.do"), tea.ToMap(request), headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * Description: 查询用户持有的信用凭证信息
 * Summary: 查询用户持有的信用凭证信息
 */
func (client *Client) QueryCreditUserissue(request *QueryCreditUserissueRequest) (_result *QueryCreditUserissueResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &QueryCreditUserissueResponse{}
	_body, _err := client.QueryCreditUserissueEx(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * Description: 查询用户持有的信用凭证信息
 * Summary: 查询用户持有的信用凭证信息
 */
func (client *Client) QueryCreditUserissueEx(request *QueryCreditUserissueRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *QueryCreditUserissueResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &QueryCreditUserissueResponse{}
	_body, _err := client.DoRequest(tea.String("1.0"), tea.String("digital.logistic.credit.userissue.query"), tea.String("HTTPS"), tea.String("POST"), tea.String("/gateway.do"), tea.ToMap(request), headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * Description: 查询用户信用流转流水记录
 * Summary: 查询用户信用流转流水记录
 */
func (client *Client) QueryCreditStatement(request *QueryCreditStatementRequest) (_result *QueryCreditStatementResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &QueryCreditStatementResponse{}
	_body, _err := client.QueryCreditStatementEx(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * Description: 查询用户信用流转流水记录
 * Summary: 查询用户信用流转流水记录
 */
func (client *Client) QueryCreditStatementEx(request *QueryCreditStatementRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *QueryCreditStatementResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &QueryCreditStatementResponse{}
	_body, _err := client.DoRequest(tea.String("1.0"), tea.String("digital.logistic.credit.statement.query"), tea.String("HTTPS"), tea.String("POST"), tea.String("/gateway.do"), tea.ToMap(request), headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * Description: 信用凭证拆分转让申请
 * Summary: 信用凭证拆分转让申请
 */
func (client *Client) CreateCreditIssuetransfer(request *CreateCreditIssuetransferRequest) (_result *CreateCreditIssuetransferResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateCreditIssuetransferResponse{}
	_body, _err := client.CreateCreditIssuetransferEx(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * Description: 信用凭证拆分转让申请
 * Summary: 信用凭证拆分转让申请
 */
func (client *Client) CreateCreditIssuetransferEx(request *CreateCreditIssuetransferRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *CreateCreditIssuetransferResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &CreateCreditIssuetransferResponse{}
	_body, _err := client.DoRequest(tea.String("1.0"), tea.String("digital.logistic.credit.issuetransfer.create"), tea.String("HTTPS"), tea.String("POST"), tea.String("/gateway.do"), tea.ToMap(request), headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * Description: 信用凭证拆分转让结果查询
 * Summary: 信用凭证拆分转让结果查询
 */
func (client *Client) QueryCreditIssuetransfer(request *QueryCreditIssuetransferRequest) (_result *QueryCreditIssuetransferResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &QueryCreditIssuetransferResponse{}
	_body, _err := client.QueryCreditIssuetransferEx(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * Description: 信用凭证拆分转让结果查询
 * Summary: 信用凭证拆分转让结果查询
 */
func (client *Client) QueryCreditIssuetransferEx(request *QueryCreditIssuetransferRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *QueryCreditIssuetransferResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &QueryCreditIssuetransferResponse{}
	_body, _err := client.DoRequest(tea.String("1.0"), tea.String("digital.logistic.credit.issuetransfer.query"), tea.String("HTTPS"), tea.String("POST"), tea.String("/gateway.do"), tea.ToMap(request), headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * Description: 信用凭证融资申请
 * Summary: 信用凭证融资申请
 */
func (client *Client) CreateCreditIssuefinance(request *CreateCreditIssuefinanceRequest) (_result *CreateCreditIssuefinanceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateCreditIssuefinanceResponse{}
	_body, _err := client.CreateCreditIssuefinanceEx(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * Description: 信用凭证融资申请
 * Summary: 信用凭证融资申请
 */
func (client *Client) CreateCreditIssuefinanceEx(request *CreateCreditIssuefinanceRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *CreateCreditIssuefinanceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &CreateCreditIssuefinanceResponse{}
	_body, _err := client.DoRequest(tea.String("1.0"), tea.String("digital.logistic.credit.issuefinance.create"), tea.String("HTTPS"), tea.String("POST"), tea.String("/gateway.do"), tea.ToMap(request), headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * Description: 信用凭证融资结果查询
 * Summary: 信用凭证融资结果查询
 */
func (client *Client) QueryCreditIssuefinance(request *QueryCreditIssuefinanceRequest) (_result *QueryCreditIssuefinanceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &QueryCreditIssuefinanceResponse{}
	_body, _err := client.QueryCreditIssuefinanceEx(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * Description: 信用凭证融资结果查询
 * Summary: 信用凭证融资结果查询
 */
func (client *Client) QueryCreditIssuefinanceEx(request *QueryCreditIssuefinanceRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *QueryCreditIssuefinanceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &QueryCreditIssuefinanceResponse{}
	_body, _err := client.DoRequest(tea.String("1.0"), tea.String("digital.logistic.credit.issuefinance.query"), tea.String("HTTPS"), tea.String("POST"), tea.String("/gateway.do"), tea.ToMap(request), headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * Description: 信用凭证清分信息查询
 * Summary: 信用凭证清分信息查询
 */
func (client *Client) QueryCreditIssuereceivable(request *QueryCreditIssuereceivableRequest) (_result *QueryCreditIssuereceivableResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &QueryCreditIssuereceivableResponse{}
	_body, _err := client.QueryCreditIssuereceivableEx(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * Description: 信用凭证清分信息查询
 * Summary: 信用凭证清分信息查询
 */
func (client *Client) QueryCreditIssuereceivableEx(request *QueryCreditIssuereceivableRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *QueryCreditIssuereceivableResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &QueryCreditIssuereceivableResponse{}
	_body, _err := client.DoRequest(tea.String("1.0"), tea.String("digital.logistic.credit.issuereceivable.query"), tea.String("HTTPS"), tea.String("POST"), tea.String("/gateway.do"), tea.ToMap(request), headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * Description: 信用凭证资产查询
 * Summary: 信用凭证资产查询
 */
func (client *Client) QueryCreditIssueamount(request *QueryCreditIssueamountRequest) (_result *QueryCreditIssueamountResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &QueryCreditIssueamountResponse{}
	_body, _err := client.QueryCreditIssueamountEx(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * Description: 信用凭证资产查询
 * Summary: 信用凭证资产查询
 */
func (client *Client) QueryCreditIssueamountEx(request *QueryCreditIssueamountRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *QueryCreditIssueamountResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &QueryCreditIssueamountResponse{}
	_body, _err := client.DoRequest(tea.String("1.0"), tea.String("digital.logistic.credit.issueamount.query"), tea.String("HTTPS"), tea.String("POST"), tea.String("/gateway.do"), tea.ToMap(request), headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * Description: 物流金融统一回调接口
 * Summary: 物流金融统一回调接口
 */
func (client *Client) CallbackCreditCommon(request *CallbackCreditCommonRequest) (_result *CallbackCreditCommonResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CallbackCreditCommonResponse{}
	_body, _err := client.CallbackCreditCommonEx(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * Description: 物流金融统一回调接口
 * Summary: 物流金融统一回调接口
 */
func (client *Client) CallbackCreditCommonEx(request *CallbackCreditCommonRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *CallbackCreditCommonResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &CallbackCreditCommonResponse{}
	_body, _err := client.DoRequest(tea.String("1.0"), tea.String("digital.logistic.credit.common.callback"), tea.String("HTTPS"), tea.String("POST"), tea.String("/gateway.do"), tea.ToMap(request), headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * Description: 信用凭证主动清分
 * Summary: 信用凭证主动清分
 */
func (client *Client) ApplyCreditIssueclear(request *ApplyCreditIssueclearRequest) (_result *ApplyCreditIssueclearResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ApplyCreditIssueclearResponse{}
	_body, _err := client.ApplyCreditIssueclearEx(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * Description: 信用凭证主动清分
 * Summary: 信用凭证主动清分
 */
func (client *Client) ApplyCreditIssueclearEx(request *ApplyCreditIssueclearRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ApplyCreditIssueclearResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &ApplyCreditIssueclearResponse{}
	_body, _err := client.DoRequest(tea.String("1.0"), tea.String("digital.logistic.credit.issueclear.apply"), tea.String("HTTPS"), tea.String("POST"), tea.String("/gateway.do"), tea.ToMap(request), headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * Description: 物流金融代理消息
 * Summary: 物流金融代理消息
 */
func (client *Client) SendCreditProxy(request *SendCreditProxyRequest) (_result *SendCreditProxyResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &SendCreditProxyResponse{}
	_body, _err := client.SendCreditProxyEx(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * Description: 物流金融代理消息
 * Summary: 物流金融代理消息
 */
func (client *Client) SendCreditProxyEx(request *SendCreditProxyRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *SendCreditProxyResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &SendCreditProxyResponse{}
	_body, _err := client.DoRequest(tea.String("1.0"), tea.String("digital.logistic.credit.proxy.send"), tea.String("HTTPS"), tea.String("POST"), tea.String("/gateway.do"), tea.ToMap(request), headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * Description: 运单信用流转核验结果查询
 * Summary: 运单信用流转核验结果查询
 */
func (client *Client) CheckCreditWaybill(request *CheckCreditWaybillRequest) (_result *CheckCreditWaybillResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CheckCreditWaybillResponse{}
	_body, _err := client.CheckCreditWaybillEx(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * Description: 运单信用流转核验结果查询
 * Summary: 运单信用流转核验结果查询
 */
func (client *Client) CheckCreditWaybillEx(request *CheckCreditWaybillRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *CheckCreditWaybillResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &CheckCreditWaybillResponse{}
	_body, _err := client.DoRequest(tea.String("1.0"), tea.String("digital.logistic.credit.waybill.check"), tea.String("HTTPS"), tea.String("POST"), tea.String("/gateway.do"), tea.ToMap(request), headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * Description: 物流金融信用流转司机换绑
 * Summary: 物流金融信用流转司机换绑
 */
func (client *Client) ReopenCreditDriver(request *ReopenCreditDriverRequest) (_result *ReopenCreditDriverResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ReopenCreditDriverResponse{}
	_body, _err := client.ReopenCreditDriverEx(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * Description: 物流金融信用流转司机换绑
 * Summary: 物流金融信用流转司机换绑
 */
func (client *Client) ReopenCreditDriverEx(request *ReopenCreditDriverRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ReopenCreditDriverResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &ReopenCreditDriverResponse{}
	_body, _err := client.DoRequest(tea.String("1.0"), tea.String("digital.logistic.credit.driver.reopen"), tea.String("HTTPS"), tea.String("POST"), tea.String("/gateway.do"), tea.ToMap(request), headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * Description: 线下协议授权关系上传
 * Summary: 线下协议授权关系上传
 */
func (client *Client) UploadCreditAuthorization(request *UploadCreditAuthorizationRequest) (_result *UploadCreditAuthorizationResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UploadCreditAuthorizationResponse{}
	_body, _err := client.UploadCreditAuthorizationEx(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * Description: 线下协议授权关系上传
 * Summary: 线下协议授权关系上传
 */
func (client *Client) UploadCreditAuthorizationEx(request *UploadCreditAuthorizationRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *UploadCreditAuthorizationResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &UploadCreditAuthorizationResponse{}
	_body, _err := client.DoRequest(tea.String("1.0"), tea.String("digital.logistic.credit.authorization.upload"), tea.String("HTTPS"), tea.String("POST"), tea.String("/gateway.do"), tea.ToMap(request), headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * Description: 线上应收转让确认关系
 * Summary: 线上应收转让确认关系
 */
func (client *Client) UploadCreditConfirm(request *UploadCreditConfirmRequest) (_result *UploadCreditConfirmResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UploadCreditConfirmResponse{}
	_body, _err := client.UploadCreditConfirmEx(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * Description: 线上应收转让确认关系
 * Summary: 线上应收转让确认关系
 */
func (client *Client) UploadCreditConfirmEx(request *UploadCreditConfirmRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *UploadCreditConfirmResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &UploadCreditConfirmResponse{}
	_body, _err := client.DoRequest(tea.String("1.0"), tea.String("digital.logistic.credit.confirm.upload"), tea.String("HTTPS"), tea.String("POST"), tea.String("/gateway.do"), tea.ToMap(request), headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * Description: 物流信用流转凭证批量发行
 * Summary: 物流信用流转凭证批量发行
 */
func (client *Client) BatchcreateCreditmodeIssue(request *BatchcreateCreditmodeIssueRequest) (_result *BatchcreateCreditmodeIssueResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &BatchcreateCreditmodeIssueResponse{}
	_body, _err := client.BatchcreateCreditmodeIssueEx(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * Description: 物流信用流转凭证批量发行
 * Summary: 物流信用流转凭证批量发行
 */
func (client *Client) BatchcreateCreditmodeIssueEx(request *BatchcreateCreditmodeIssueRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *BatchcreateCreditmodeIssueResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &BatchcreateCreditmodeIssueResponse{}
	_body, _err := client.DoRequest(tea.String("1.0"), tea.String("digital.logistic.creditmode.issue.batchcreate"), tea.String("HTTPS"), tea.String("POST"), tea.String("/gateway.do"), tea.ToMap(request), headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * Description: 物流信用流转按模式主动清分
 * Summary: 物流信用流转按模式主动清分
 */
func (client *Client) ApplyCreditmodeIssueclear(request *ApplyCreditmodeIssueclearRequest) (_result *ApplyCreditmodeIssueclearResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ApplyCreditmodeIssueclearResponse{}
	_body, _err := client.ApplyCreditmodeIssueclearEx(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * Description: 物流信用流转按模式主动清分
 * Summary: 物流信用流转按模式主动清分
 */
func (client *Client) ApplyCreditmodeIssueclearEx(request *ApplyCreditmodeIssueclearRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ApplyCreditmodeIssueclearResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &ApplyCreditmodeIssueclearResponse{}
	_body, _err := client.DoRequest(tea.String("1.0"), tea.String("digital.logistic.creditmode.issueclear.apply"), tea.String("HTTPS"), tea.String("POST"), tea.String("/gateway.do"), tea.ToMap(request), headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * Description: 信用流转发行信息上传SAAS版
 * Summary: 信用流转发行信息上传SAAS版
 */
func (client *Client) UploadCreditIssuebysaas(request *UploadCreditIssuebysaasRequest) (_result *UploadCreditIssuebysaasResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UploadCreditIssuebysaasResponse{}
	_body, _err := client.UploadCreditIssuebysaasEx(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * Description: 信用流转发行信息上传SAAS版
 * Summary: 信用流转发行信息上传SAAS版
 */
func (client *Client) UploadCreditIssuebysaasEx(request *UploadCreditIssuebysaasRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *UploadCreditIssuebysaasResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &UploadCreditIssuebysaasResponse{}
	_body, _err := client.DoRequest(tea.String("1.0"), tea.String("digital.logistic.credit.issuebysaas.upload"), tea.String("HTTPS"), tea.String("POST"), tea.String("/gateway.do"), tea.ToMap(request), headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * Description: 信用流转B模式发行批次撤销SAAS版
 * Summary: 信用流转B模式发行批次撤销SAAS版
 */
func (client *Client) CancelCreditIssuebatchbysaas(request *CancelCreditIssuebatchbysaasRequest) (_result *CancelCreditIssuebatchbysaasResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CancelCreditIssuebatchbysaasResponse{}
	_body, _err := client.CancelCreditIssuebatchbysaasEx(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * Description: 信用流转B模式发行批次撤销SAAS版
 * Summary: 信用流转B模式发行批次撤销SAAS版
 */
func (client *Client) CancelCreditIssuebatchbysaasEx(request *CancelCreditIssuebatchbysaasRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *CancelCreditIssuebatchbysaasResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &CancelCreditIssuebatchbysaasResponse{}
	_body, _err := client.DoRequest(tea.String("1.0"), tea.String("digital.logistic.credit.issuebatchbysaas.cancel"), tea.String("HTTPS"), tea.String("POST"), tea.String("/gateway.do"), tea.ToMap(request), headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * Description: 信用流转可用额度查询SAAS版
 * Summary: 信用流转可用额度查询SAAS版
 */
func (client *Client) QueryCreditBalancebysaas(request *QueryCreditBalancebysaasRequest) (_result *QueryCreditBalancebysaasResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &QueryCreditBalancebysaasResponse{}
	_body, _err := client.QueryCreditBalancebysaasEx(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * Description: 信用流转可用额度查询SAAS版
 * Summary: 信用流转可用额度查询SAAS版
 */
func (client *Client) QueryCreditBalancebysaasEx(request *QueryCreditBalancebysaasRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *QueryCreditBalancebysaasResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &QueryCreditBalancebysaasResponse{}
	_body, _err := client.DoRequest(tea.String("1.0"), tea.String("digital.logistic.credit.balancebysaas.query"), tea.String("HTTPS"), tea.String("POST"), tea.String("/gateway.do"), tea.ToMap(request), headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * Description: 信用凭证电子回单获取
 * Summary: 信用凭证电子回单获取
 */
func (client *Client) GetCreditIssuescpticket(request *GetCreditIssuescpticketRequest) (_result *GetCreditIssuescpticketResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetCreditIssuescpticketResponse{}
	_body, _err := client.GetCreditIssuescpticketEx(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * Description: 信用凭证电子回单获取
 * Summary: 信用凭证电子回单获取
 */
func (client *Client) GetCreditIssuescpticketEx(request *GetCreditIssuescpticketRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetCreditIssuescpticketResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &GetCreditIssuescpticketResponse{}
	_body, _err := client.DoRequest(tea.String("1.0"), tea.String("digital.logistic.credit.issuescpticket.get"), tea.String("HTTPS"), tea.String("POST"), tea.String("/gateway.do"), tea.ToMap(request), headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * Description: 信用凭证电子回单获取结果查询接口
 * Summary: 信用凭证电子回单获取结果查询接口
 */
func (client *Client) QueryCreditIssuescpticketresult(request *QueryCreditIssuescpticketresultRequest) (_result *QueryCreditIssuescpticketresultResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &QueryCreditIssuescpticketresultResponse{}
	_body, _err := client.QueryCreditIssuescpticketresultEx(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * Description: 信用凭证电子回单获取结果查询接口
 * Summary: 信用凭证电子回单获取结果查询接口
 */
func (client *Client) QueryCreditIssuescpticketresultEx(request *QueryCreditIssuescpticketresultRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *QueryCreditIssuescpticketresultResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &QueryCreditIssuescpticketresultResponse{}
	_body, _err := client.DoRequest(tea.String("1.0"), tea.String("digital.logistic.credit.issuescpticketresult.query"), tea.String("HTTPS"), tea.String("POST"), tea.String("/gateway.do"), tea.ToMap(request), headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * Description: 信用流转A+模式发行信息上传
 * Summary: 信用流转A+模式发行信息上传
 */
func (client *Client) UploadCreditAplusissue(request *UploadCreditAplusissueRequest) (_result *UploadCreditAplusissueResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UploadCreditAplusissueResponse{}
	_body, _err := client.UploadCreditAplusissueEx(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * Description: 信用流转A+模式发行信息上传
 * Summary: 信用流转A+模式发行信息上传
 */
func (client *Client) UploadCreditAplusissueEx(request *UploadCreditAplusissueRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *UploadCreditAplusissueResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &UploadCreditAplusissueResponse{}
	_body, _err := client.DoRequest(tea.String("1.0"), tea.String("digital.logistic.credit.aplusissue.upload"), tea.String("HTTPS"), tea.String("POST"), tea.String("/gateway.do"), tea.ToMap(request), headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * Description: 信用流转可用额度查询接口
 * Summary: 信用流转可用额度查询接口
 */
func (client *Client) QueryCreditCreditamount(request *QueryCreditCreditamountRequest) (_result *QueryCreditCreditamountResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &QueryCreditCreditamountResponse{}
	_body, _err := client.QueryCreditCreditamountEx(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * Description: 信用流转可用额度查询接口
 * Summary: 信用流转可用额度查询接口
 */
func (client *Client) QueryCreditCreditamountEx(request *QueryCreditCreditamountRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *QueryCreditCreditamountResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &QueryCreditCreditamountResponse{}
	_body, _err := client.DoRequest(tea.String("1.0"), tea.String("digital.logistic.credit.creditamount.query"), tea.String("HTTPS"), tea.String("POST"), tea.String("/gateway.do"), tea.ToMap(request), headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * Description: 信用流转非授信通用签约接口
 * Summary: 信用流转非授信通用签约接口
 */
func (client *Client) CreateCreditCommonsign(request *CreateCreditCommonsignRequest) (_result *CreateCreditCommonsignResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateCreditCommonsignResponse{}
	_body, _err := client.CreateCreditCommonsignEx(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * Description: 信用流转非授信通用签约接口
 * Summary: 信用流转非授信通用签约接口
 */
func (client *Client) CreateCreditCommonsignEx(request *CreateCreditCommonsignRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *CreateCreditCommonsignResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &CreateCreditCommonsignResponse{}
	_body, _err := client.DoRequest(tea.String("1.0"), tea.String("digital.logistic.credit.commonsign.create"), tea.String("HTTPS"), tea.String("POST"), tea.String("/gateway.do"), tea.ToMap(request), headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * Description: 信用流转非授信通用签约查询接口
 * Summary: 信用流转非授信通用签约查询接口
 */
func (client *Client) QueryCreditCommonsign(request *QueryCreditCommonsignRequest) (_result *QueryCreditCommonsignResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &QueryCreditCommonsignResponse{}
	_body, _err := client.QueryCreditCommonsignEx(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * Description: 信用流转非授信通用签约查询接口
 * Summary: 信用流转非授信通用签约查询接口
 */
func (client *Client) QueryCreditCommonsignEx(request *QueryCreditCommonsignRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *QueryCreditCommonsignResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &QueryCreditCommonsignResponse{}
	_body, _err := client.DoRequest(tea.String("1.0"), tea.String("digital.logistic.credit.commonsign.query"), tea.String("HTTPS"), tea.String("POST"), tea.String("/gateway.do"), tea.ToMap(request), headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * Description: 信用流转凭证合并发行接口
 * Summary: 信用流转凭证合并发行接口
 */
func (client *Client) BatchcreateCreditIssue(request *BatchcreateCreditIssueRequest) (_result *BatchcreateCreditIssueResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &BatchcreateCreditIssueResponse{}
	_body, _err := client.BatchcreateCreditIssueEx(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * Description: 信用流转凭证合并发行接口
 * Summary: 信用流转凭证合并发行接口
 */
func (client *Client) BatchcreateCreditIssueEx(request *BatchcreateCreditIssueRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *BatchcreateCreditIssueResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &BatchcreateCreditIssueResponse{}
	_body, _err := client.DoRequest(tea.String("1.0"), tea.String("digital.logistic.credit.issue.batchcreate"), tea.String("HTTPS"), tea.String("POST"), tea.String("/gateway.do"), tea.ToMap(request), headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * Description: 信用流转凭证合并发行信息上传接口
 * Summary: 信用流转凭证合并发行信息上传接口
 */
func (client *Client) UploadCreditIssuebatch(request *UploadCreditIssuebatchRequest) (_result *UploadCreditIssuebatchResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UploadCreditIssuebatchResponse{}
	_body, _err := client.UploadCreditIssuebatchEx(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * Description: 信用流转凭证合并发行信息上传接口
 * Summary: 信用流转凭证合并发行信息上传接口
 */
func (client *Client) UploadCreditIssuebatchEx(request *UploadCreditIssuebatchRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *UploadCreditIssuebatchResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &UploadCreditIssuebatchResponse{}
	_body, _err := client.DoRequest(tea.String("1.0"), tea.String("digital.logistic.credit.issuebatch.upload"), tea.String("HTTPS"), tea.String("POST"), tea.String("/gateway.do"), tea.ToMap(request), headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * Description: 信用凭证电子回单获取(新)，明确了凭证的转出方did和转入方did
 * Summary: 信用凭证电子回单获取(新)
 */
func (client *Client) GetIssueTransferfile(request *GetIssueTransferfileRequest) (_result *GetIssueTransferfileResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetIssueTransferfileResponse{}
	_body, _err := client.GetIssueTransferfileEx(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * Description: 信用凭证电子回单获取(新)，明确了凭证的转出方did和转入方did
 * Summary: 信用凭证电子回单获取(新)
 */
func (client *Client) GetIssueTransferfileEx(request *GetIssueTransferfileRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetIssueTransferfileResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &GetIssueTransferfileResponse{}
	_body, _err := client.DoRequest(tea.String("1.0"), tea.String("digital.logistic.issue.transferfile.get"), tea.String("HTTPS"), tea.String("POST"), tea.String("/gateway.do"), tea.ToMap(request), headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * Description: 保险服务，投保接口。接口提供幂等逻辑，请求后接口会返回成功、失败、处理中。其中处理中是极端场景，需要客户端使用相同的业务流水号发起重试，以免造成重复投保
 * Summary: 投保接口-中华财险-承运人责任险
 */
func (client *Client) ApplyInsurancepolicyZhonghuacaixian(request *ApplyInsurancepolicyZhonghuacaixianRequest) (_result *ApplyInsurancepolicyZhonghuacaixianResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ApplyInsurancepolicyZhonghuacaixianResponse{}
	_body, _err := client.ApplyInsurancepolicyZhonghuacaixianEx(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * Description: 保险服务，投保接口。接口提供幂等逻辑，请求后接口会返回成功、失败、处理中。其中处理中是极端场景，需要客户端使用相同的业务流水号发起重试，以免造成重复投保
 * Summary: 投保接口-中华财险-承运人责任险
 */
func (client *Client) ApplyInsurancepolicyZhonghuacaixianEx(request *ApplyInsurancepolicyZhonghuacaixianRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ApplyInsurancepolicyZhonghuacaixianResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &ApplyInsurancepolicyZhonghuacaixianResponse{}
	_body, _err := client.DoRequest(tea.String("1.0"), tea.String("digital.logistic.insurancepolicy.zhonghuacaixian.apply"), tea.String("HTTPS"), tea.String("POST"), tea.String("/gateway.do"), tea.ToMap(request), headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * Description: 数字物流保险，退保接口
 * Summary: 退保接口-中华财险-承运人责任险
 */
func (client *Client) CancelInsurancepolicyZhonghuacaixian(request *CancelInsurancepolicyZhonghuacaixianRequest) (_result *CancelInsurancepolicyZhonghuacaixianResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CancelInsurancepolicyZhonghuacaixianResponse{}
	_body, _err := client.CancelInsurancepolicyZhonghuacaixianEx(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * Description: 数字物流保险，退保接口
 * Summary: 退保接口-中华财险-承运人责任险
 */
func (client *Client) CancelInsurancepolicyZhonghuacaixianEx(request *CancelInsurancepolicyZhonghuacaixianRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *CancelInsurancepolicyZhonghuacaixianResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &CancelInsurancepolicyZhonghuacaixianResponse{}
	_body, _err := client.DoRequest(tea.String("1.0"), tea.String("digital.logistic.insurancepolicy.zhonghuacaixian.cancel"), tea.String("HTTPS"), tea.String("POST"), tea.String("/gateway.do"), tea.ToMap(request), headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * Description: 保险理赔通知
 * Summary: 保险理赔通知
 */
func (client *Client) PushInsurancenotifyClaim(request *PushInsurancenotifyClaimRequest) (_result *PushInsurancenotifyClaimResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &PushInsurancenotifyClaimResponse{}
	_body, _err := client.PushInsurancenotifyClaimEx(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * Description: 保险理赔通知
 * Summary: 保险理赔通知
 */
func (client *Client) PushInsurancenotifyClaimEx(request *PushInsurancenotifyClaimRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *PushInsurancenotifyClaimResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &PushInsurancenotifyClaimResponse{}
	_body, _err := client.DoRequest(tea.String("1.0"), tea.String("digital.logistic.insurancenotify.claim.push"), tea.String("HTTPS"), tea.String("POST"), tea.String("/gateway.do"), tea.ToMap(request), headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * Description: 系统会生成上传文件的连接，使用方使用该链接进行文件上传。该链接具有一定的时效性，因此建议按需按时申请使用
 * Summary: 上传文件链接申请
 */
func (client *Client) ApplyInsuranceFileurl(request *ApplyInsuranceFileurlRequest) (_result *ApplyInsuranceFileurlResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ApplyInsuranceFileurlResponse{}
	_body, _err := client.ApplyInsuranceFileurlEx(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * Description: 系统会生成上传文件的连接，使用方使用该链接进行文件上传。该链接具有一定的时效性，因此建议按需按时申请使用
 * Summary: 上传文件链接申请
 */
func (client *Client) ApplyInsuranceFileurlEx(request *ApplyInsuranceFileurlRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ApplyInsuranceFileurlResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &ApplyInsuranceFileurlResponse{}
	_body, _err := client.DoRequest(tea.String("1.0"), tea.String("digital.logistic.insurance.fileurl.apply"), tea.String("HTTPS"), tea.String("POST"), tea.String("/gateway.do"), tea.ToMap(request), headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * Description: 保险服务-平台责任险投保接口。根据运输货物货值生成相应的保费。
接口提供幂等逻辑，请求后接口会返回成功、失败、处理中。其中处理中是极端场景，需要客户端使用相同的业务流水号发起重试，以免造成重复投保
 * Summary: 投保接口-承运人平台责任险
*/
func (client *Client) ApplyInsurancepolicyUniversal(request *ApplyInsurancepolicyUniversalRequest) (_result *ApplyInsurancepolicyUniversalResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ApplyInsurancepolicyUniversalResponse{}
	_body, _err := client.ApplyInsurancepolicyUniversalEx(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * Description: 保险服务-平台责任险投保接口。根据运输货物货值生成相应的保费。
接口提供幂等逻辑，请求后接口会返回成功、失败、处理中。其中处理中是极端场景，需要客户端使用相同的业务流水号发起重试，以免造成重复投保
 * Summary: 投保接口-承运人平台责任险
*/
func (client *Client) ApplyInsurancepolicyUniversalEx(request *ApplyInsurancepolicyUniversalRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ApplyInsurancepolicyUniversalResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &ApplyInsurancepolicyUniversalResponse{}
	_body, _err := client.DoRequest(tea.String("1.0"), tea.String("digital.logistic.insurancepolicy.universal.apply"), tea.String("HTTPS"), tea.String("POST"), tea.String("/gateway.do"), tea.ToMap(request), headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * Description: 数字物流保险，退保接口
 * Summary: 退保接口-承运人平台责任险
 */
func (client *Client) CancelInsurancepolicyUniversal(request *CancelInsurancepolicyUniversalRequest) (_result *CancelInsurancepolicyUniversalResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CancelInsurancepolicyUniversalResponse{}
	_body, _err := client.CancelInsurancepolicyUniversalEx(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * Description: 数字物流保险，退保接口
 * Summary: 退保接口-承运人平台责任险
 */
func (client *Client) CancelInsurancepolicyUniversalEx(request *CancelInsurancepolicyUniversalRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *CancelInsurancepolicyUniversalResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &CancelInsurancepolicyUniversalResponse{}
	_body, _err := client.DoRequest(tea.String("1.0"), tea.String("digital.logistic.insurancepolicy.universal.cancel"), tea.String("HTTPS"), tea.String("POST"), tea.String("/gateway.do"), tea.ToMap(request), headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * Description: 利用区块链智能合约进行离线保单存证
 * Summary: 离线保单存证推送
 */
func (client *Client) PushInsuranceOlp(request *PushInsuranceOlpRequest) (_result *PushInsuranceOlpResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &PushInsuranceOlpResponse{}
	_body, _err := client.PushInsuranceOlpEx(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * Description: 利用区块链智能合约进行离线保单存证
 * Summary: 离线保单存证推送
 */
func (client *Client) PushInsuranceOlpEx(request *PushInsuranceOlpRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *PushInsuranceOlpResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &PushInsuranceOlpResponse{}
	_body, _err := client.DoRequest(tea.String("1.0"), tea.String("digital.logistic.insurance.olp.push"), tea.String("HTTPS"), tea.String("POST"), tea.String("/gateway.do"), tea.ToMap(request), headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * Description: 利用区块链智能合约进行离线保单信息更新追踪
 * Summary: 离线保单更新
 */
func (client *Client) UpdateInsuranceOlp(request *UpdateInsuranceOlpRequest) (_result *UpdateInsuranceOlpResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpdateInsuranceOlpResponse{}
	_body, _err := client.UpdateInsuranceOlpEx(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * Description: 利用区块链智能合约进行离线保单信息更新追踪
 * Summary: 离线保单更新
 */
func (client *Client) UpdateInsuranceOlpEx(request *UpdateInsuranceOlpRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *UpdateInsuranceOlpResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &UpdateInsuranceOlpResponse{}
	_body, _err := client.DoRequest(tea.String("1.0"), tea.String("digital.logistic.insurance.olp.update"), tea.String("HTTPS"), tea.String("POST"), tea.String("/gateway.do"), tea.ToMap(request), headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * Description: 再保分保单推送接口
 * Summary: 再保分保单推送
 */
func (client *Client) PushInsuranceReppolicy(request *PushInsuranceReppolicyRequest) (_result *PushInsuranceReppolicyResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &PushInsuranceReppolicyResponse{}
	_body, _err := client.PushInsuranceReppolicyEx(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * Description: 再保分保单推送接口
 * Summary: 再保分保单推送
 */
func (client *Client) PushInsuranceReppolicyEx(request *PushInsuranceReppolicyRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *PushInsuranceReppolicyResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &PushInsuranceReppolicyResponse{}
	_body, _err := client.DoRequest(tea.String("1.0"), tea.String("digital.logistic.insurance.reppolicy.push"), tea.String("HTTPS"), tea.String("POST"), tea.String("/gateway.do"), tea.ToMap(request), headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * Description: 再保批改单推送
 * Summary: 再保批改单推送
 */
func (client *Client) PushInsuranceRepcorrect(request *PushInsuranceRepcorrectRequest) (_result *PushInsuranceRepcorrectResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &PushInsuranceRepcorrectResponse{}
	_body, _err := client.PushInsuranceRepcorrectEx(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * Description: 再保批改单推送
 * Summary: 再保批改单推送
 */
func (client *Client) PushInsuranceRepcorrectEx(request *PushInsuranceRepcorrectRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *PushInsuranceRepcorrectResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &PushInsuranceRepcorrectResponse{}
	_body, _err := client.DoRequest(tea.String("1.0"), tea.String("digital.logistic.insurance.repcorrect.push"), tea.String("HTTPS"), tea.String("POST"), tea.String("/gateway.do"), tea.ToMap(request), headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * Description: 跨境运费险投保申请
 * Summary: 跨境运费险投保
 */
func (client *Client) ApplyInsuranceCbrf(request *ApplyInsuranceCbrfRequest) (_result *ApplyInsuranceCbrfResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ApplyInsuranceCbrfResponse{}
	_body, _err := client.ApplyInsuranceCbrfEx(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * Description: 跨境运费险投保申请
 * Summary: 跨境运费险投保
 */
func (client *Client) ApplyInsuranceCbrfEx(request *ApplyInsuranceCbrfRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ApplyInsuranceCbrfResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &ApplyInsuranceCbrfResponse{}
	_body, _err := client.DoRequest(tea.String("1.0"), tea.String("digital.logistic.insurance.cbrf.apply"), tea.String("HTTPS"), tea.String("POST"), tea.String("/gateway.do"), tea.ToMap(request), headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * Description: 跨境运费险理赔
 * Summary: 跨境运费险理赔
 */
func (client *Client) RepayInsuranceCbrf(request *RepayInsuranceCbrfRequest) (_result *RepayInsuranceCbrfResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &RepayInsuranceCbrfResponse{}
	_body, _err := client.RepayInsuranceCbrfEx(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * Description: 跨境运费险理赔
 * Summary: 跨境运费险理赔
 */
func (client *Client) RepayInsuranceCbrfEx(request *RepayInsuranceCbrfRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *RepayInsuranceCbrfResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &RepayInsuranceCbrfResponse{}
	_body, _err := client.DoRequest(tea.String("1.0"), tea.String("digital.logistic.insurance.cbrf.repay"), tea.String("HTTPS"), tea.String("POST"), tea.String("/gateway.do"), tea.ToMap(request), headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * Description: 跨境出口货运险投保
 * Summary: 跨境出口货运险投保
 */
func (client *Client) ApplyInsuranceCbec(request *ApplyInsuranceCbecRequest) (_result *ApplyInsuranceCbecResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ApplyInsuranceCbecResponse{}
	_body, _err := client.ApplyInsuranceCbecEx(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * Description: 跨境出口货运险投保
 * Summary: 跨境出口货运险投保
 */
func (client *Client) ApplyInsuranceCbecEx(request *ApplyInsuranceCbecRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ApplyInsuranceCbecResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &ApplyInsuranceCbecResponse{}
	_body, _err := client.DoRequest(tea.String("1.0"), tea.String("digital.logistic.insurance.cbec.apply"), tea.String("HTTPS"), tea.String("POST"), tea.String("/gateway.do"), tea.ToMap(request), headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * Description: 货物入库申报
 * Summary: 货物入库申报
 */
func (client *Client) ApplyInsuranceStockin(request *ApplyInsuranceStockinRequest) (_result *ApplyInsuranceStockinResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ApplyInsuranceStockinResponse{}
	_body, _err := client.ApplyInsuranceStockinEx(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * Description: 货物入库申报
 * Summary: 货物入库申报
 */
func (client *Client) ApplyInsuranceStockinEx(request *ApplyInsuranceStockinRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ApplyInsuranceStockinResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &ApplyInsuranceStockinResponse{}
	_body, _err := client.DoRequest(tea.String("1.0"), tea.String("digital.logistic.insurance.stockin.apply"), tea.String("HTTPS"), tea.String("POST"), tea.String("/gateway.do"), tea.ToMap(request), headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * Description: 货物库存申报
 * Summary: 货物库存申报
 */
func (client *Client) ApplyInsuranceInventory(request *ApplyInsuranceInventoryRequest) (_result *ApplyInsuranceInventoryResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ApplyInsuranceInventoryResponse{}
	_body, _err := client.ApplyInsuranceInventoryEx(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * Description: 货物库存申报
 * Summary: 货物库存申报
 */
func (client *Client) ApplyInsuranceInventoryEx(request *ApplyInsuranceInventoryRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ApplyInsuranceInventoryResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &ApplyInsuranceInventoryResponse{}
	_body, _err := client.DoRequest(tea.String("1.0"), tea.String("digital.logistic.insurance.inventory.apply"), tea.String("HTTPS"), tea.String("POST"), tea.String("/gateway.do"), tea.ToMap(request), headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * Description: 授权签署信息推送
 * Summary: 授权签署信息推送
 */
func (client *Client) PushAuthSigninfo(request *PushAuthSigninfoRequest) (_result *PushAuthSigninfoResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &PushAuthSigninfoResponse{}
	_body, _err := client.PushAuthSigninfoEx(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * Description: 授权签署信息推送
 * Summary: 授权签署信息推送
 */
func (client *Client) PushAuthSigninfoEx(request *PushAuthSigninfoRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *PushAuthSigninfoResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &PushAuthSigninfoResponse{}
	_body, _err := client.DoRequest(tea.String("1.0"), tea.String("digital.logistic.auth.signinfo.push"), tea.String("HTTPS"), tea.String("POST"), tea.String("/gateway.do"), tea.ToMap(request), headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * Description: 池融资委托支付入账查询
 * Summary: 委托支付入账查询
 */
func (client *Client) QueryPfPayment(request *QueryPfPaymentRequest) (_result *QueryPfPaymentResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &QueryPfPaymentResponse{}
	_body, _err := client.QueryPfPaymentEx(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * Description: 池融资委托支付入账查询
 * Summary: 委托支付入账查询
 */
func (client *Client) QueryPfPaymentEx(request *QueryPfPaymentRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *QueryPfPaymentResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &QueryPfPaymentResponse{}
	_body, _err := client.DoRequest(tea.String("1.0"), tea.String("digital.logistic.pf.payment.query"), tea.String("HTTPS"), tea.String("POST"), tea.String("/gateway.do"), tea.ToMap(request), headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * Description: 池融资借据信息查询
 * Summary: 借据信息查询
 */
func (client *Client) QueryPfIou(request *QueryPfIouRequest) (_result *QueryPfIouResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &QueryPfIouResponse{}
	_body, _err := client.QueryPfIouEx(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * Description: 池融资借据信息查询
 * Summary: 借据信息查询
 */
func (client *Client) QueryPfIouEx(request *QueryPfIouRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *QueryPfIouResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &QueryPfIouResponse{}
	_body, _err := client.DoRequest(tea.String("1.0"), tea.String("digital.logistic.pf.iou.query"), tea.String("HTTPS"), tea.String("POST"), tea.String("/gateway.do"), tea.ToMap(request), headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * Description: 池融资授信额度查询
 * Summary: 授信额度查询
 */
func (client *Client) QueryPfQuota(request *QueryPfQuotaRequest) (_result *QueryPfQuotaResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &QueryPfQuotaResponse{}
	_body, _err := client.QueryPfQuotaEx(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * Description: 池融资授信额度查询
 * Summary: 授信额度查询
 */
func (client *Client) QueryPfQuotaEx(request *QueryPfQuotaRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *QueryPfQuotaResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &QueryPfQuotaResponse{}
	_body, _err := client.DoRequest(tea.String("1.0"), tea.String("digital.logistic.pf.quota.query"), tea.String("HTTPS"), tea.String("POST"), tea.String("/gateway.do"), tea.ToMap(request), headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * Description: 池融资支用申请
 * Summary: 池融资支用申请
 */
func (client *Client) ApplyPfWaybillfinancing(request *ApplyPfWaybillfinancingRequest) (_result *ApplyPfWaybillfinancingResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ApplyPfWaybillfinancingResponse{}
	_body, _err := client.ApplyPfWaybillfinancingEx(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * Description: 池融资支用申请
 * Summary: 池融资支用申请
 */
func (client *Client) ApplyPfWaybillfinancingEx(request *ApplyPfWaybillfinancingRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ApplyPfWaybillfinancingResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &ApplyPfWaybillfinancingResponse{}
	_body, _err := client.DoRequest(tea.String("1.0"), tea.String("digital.logistic.pf.waybillfinancing.apply"), tea.String("HTTPS"), tea.String("POST"), tea.String("/gateway.do"), tea.ToMap(request), headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * Description: 池融资账单质押，用于入池
 * Summary: 池融资账单质押
 */
func (client *Client) PushPfPledge(request *PushPfPledgeRequest) (_result *PushPfPledgeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &PushPfPledgeResponse{}
	_body, _err := client.PushPfPledgeEx(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * Description: 池融资账单质押，用于入池
 * Summary: 池融资账单质押
 */
func (client *Client) PushPfPledgeEx(request *PushPfPledgeRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *PushPfPledgeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &PushPfPledgeResponse{}
	_body, _err := client.DoRequest(tea.String("1.0"), tea.String("digital.logistic.pf.pledge.push"), tea.String("HTTPS"), tea.String("POST"), tea.String("/gateway.do"), tea.ToMap(request), headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * Description: 入池账单质押状态查询
 * Summary: 质押状态查询
 */
func (client *Client) QueryPfPledge(request *QueryPfPledgeRequest) (_result *QueryPfPledgeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &QueryPfPledgeResponse{}
	_body, _err := client.QueryPfPledgeEx(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * Description: 入池账单质押状态查询
 * Summary: 质押状态查询
 */
func (client *Client) QueryPfPledgeEx(request *QueryPfPledgeRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *QueryPfPledgeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &QueryPfPledgeResponse{}
	_body, _err := client.DoRequest(tea.String("1.0"), tea.String("digital.logistic.pf.pledge.query"), tea.String("HTTPS"), tea.String("POST"), tea.String("/gateway.do"), tea.ToMap(request), headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * Description: 池融资融资支用状态查询
 * Summary: 池融资融资支用状态查询
 */
func (client *Client) QueryPfFinancing(request *QueryPfFinancingRequest) (_result *QueryPfFinancingResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &QueryPfFinancingResponse{}
	_body, _err := client.QueryPfFinancingEx(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * Description: 池融资融资支用状态查询
 * Summary: 池融资融资支用状态查询
 */
func (client *Client) QueryPfFinancingEx(request *QueryPfFinancingRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *QueryPfFinancingResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &QueryPfFinancingResponse{}
	_body, _err := client.DoRequest(tea.String("1.0"), tea.String("digital.logistic.pf.financing.query"), tea.String("HTTPS"), tea.String("POST"), tea.String("/gateway.do"), tea.ToMap(request), headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * Description: 池融资凭证核验结果查询
 * Summary: 池融资凭证核验结果查询
 */
func (client *Client) CheckPfVoucher(request *CheckPfVoucherRequest) (_result *CheckPfVoucherResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CheckPfVoucherResponse{}
	_body, _err := client.CheckPfVoucherEx(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * Description: 池融资凭证核验结果查询
 * Summary: 池融资凭证核验结果查询
 */
func (client *Client) CheckPfVoucherEx(request *CheckPfVoucherRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *CheckPfVoucherResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &CheckPfVoucherResponse{}
	_body, _err := client.DoRequest(tea.String("1.0"), tea.String("digital.logistic.pf.voucher.check"), tea.String("HTTPS"), tea.String("POST"), tea.String("/gateway.do"), tea.ToMap(request), headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * Description: 池融资提款确认书申请
 * Summary: 提款确认书申请
 */
func (client *Client) ApplyPfConfirmation(request *ApplyPfConfirmationRequest) (_result *ApplyPfConfirmationResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ApplyPfConfirmationResponse{}
	_body, _err := client.ApplyPfConfirmationEx(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * Description: 池融资提款确认书申请
 * Summary: 提款确认书申请
 */
func (client *Client) ApplyPfConfirmationEx(request *ApplyPfConfirmationRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ApplyPfConfirmationResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &ApplyPfConfirmationResponse{}
	_body, _err := client.DoRequest(tea.String("1.0"), tea.String("digital.logistic.pf.confirmation.apply"), tea.String("HTTPS"), tea.String("POST"), tea.String("/gateway.do"), tea.ToMap(request), headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * Description: 池融资融资资格申请
 * Summary: 池融资融资资格申请
 */
func (client *Client) ApplyPfFinancingqualification(request *ApplyPfFinancingqualificationRequest) (_result *ApplyPfFinancingqualificationResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ApplyPfFinancingqualificationResponse{}
	_body, _err := client.ApplyPfFinancingqualificationEx(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * Description: 池融资融资资格申请
 * Summary: 池融资融资资格申请
 */
func (client *Client) ApplyPfFinancingqualificationEx(request *ApplyPfFinancingqualificationRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ApplyPfFinancingqualificationResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &ApplyPfFinancingqualificationResponse{}
	_body, _err := client.DoRequest(tea.String("1.0"), tea.String("digital.logistic.pf.financingqualification.apply"), tea.String("HTTPS"), tea.String("POST"), tea.String("/gateway.do"), tea.ToMap(request), headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * Description: 池融资融资资格申请结果查询
 * Summary: 池融资融资资格申请结果查询
 */
func (client *Client) QueryPfFinancingqualification(request *QueryPfFinancingqualificationRequest) (_result *QueryPfFinancingqualificationResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &QueryPfFinancingqualificationResponse{}
	_body, _err := client.QueryPfFinancingqualificationEx(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * Description: 池融资融资资格申请结果查询
 * Summary: 池融资融资资格申请结果查询
 */
func (client *Client) QueryPfFinancingqualificationEx(request *QueryPfFinancingqualificationRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *QueryPfFinancingqualificationResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &QueryPfFinancingqualificationResponse{}
	_body, _err := client.DoRequest(tea.String("1.0"), tea.String("digital.logistic.pf.financingqualification.query"), tea.String("HTTPS"), tea.String("POST"), tea.String("/gateway.do"), tea.ToMap(request), headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * Description: 池融资主站回调金融云接口
 * Summary: 池融资主站回调金融云接口
 */
func (client *Client) CallbackPfDefinpf(request *CallbackPfDefinpfRequest) (_result *CallbackPfDefinpfResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CallbackPfDefinpfResponse{}
	_body, _err := client.CallbackPfDefinpfEx(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * Description: 池融资主站回调金融云接口
 * Summary: 池融资主站回调金融云接口
 */
func (client *Client) CallbackPfDefinpfEx(request *CallbackPfDefinpfRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *CallbackPfDefinpfResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &CallbackPfDefinpfResponse{}
	_body, _err := client.DoRequest(tea.String("1.0"), tea.String("digital.logistic.pf.definpf.callback"), tea.String("HTTPS"), tea.String("POST"), tea.String("/gateway.do"), tea.ToMap(request), headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * Description: 创建货代did
 * Summary: 创建货代did
 */
func (client *Client) CreateDidForwarder(request *CreateDidForwarderRequest) (_result *CreateDidForwarderResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateDidForwarderResponse{}
	_body, _err := client.CreateDidForwarderEx(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * Description: 创建货代did
 * Summary: 创建货代did
 */
func (client *Client) CreateDidForwarderEx(request *CreateDidForwarderRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *CreateDidForwarderResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &CreateDidForwarderResponse{}
	_body, _err := client.DoRequest(tea.String("1.0"), tea.String("digital.logistic.did.forwarder.create"), tea.String("HTTPS"), tea.String("POST"), tea.String("/gateway.do"), tea.ToMap(request), headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * Description: 创建saas平台did
 * Summary: 创建saas平台did
 */
func (client *Client) CreateDidSaasplatform(request *CreateDidSaasplatformRequest) (_result *CreateDidSaasplatformResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateDidSaasplatformResponse{}
	_body, _err := client.CreateDidSaasplatformEx(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * Description: 创建saas平台did
 * Summary: 创建saas平台did
 */
func (client *Client) CreateDidSaasplatformEx(request *CreateDidSaasplatformRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *CreateDidSaasplatformResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &CreateDidSaasplatformResponse{}
	_body, _err := client.DoRequest(tea.String("1.0"), tea.String("digital.logistic.did.saasplatform.create"), tea.String("HTTPS"), tea.String("POST"), tea.String("/gateway.do"), tea.ToMap(request), headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * Description: 创建直客did
 * Summary: 创建直客did
 */
func (client *Client) CreateDidClient(request *CreateDidClientRequest) (_result *CreateDidClientResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateDidClientResponse{}
	_body, _err := client.CreateDidClientEx(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * Description: 创建直客did
 * Summary: 创建直客did
 */
func (client *Client) CreateDidClientEx(request *CreateDidClientRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *CreateDidClientResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &CreateDidClientResponse{}
	_body, _err := client.DoRequest(tea.String("1.0"), tea.String("digital.logistic.did.client.create"), tea.String("HTTPS"), tea.String("POST"), tea.String("/gateway.do"), tea.ToMap(request), headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * Description: 保存订单
 * Summary: 保存订单
 */
func (client *Client) SaveBizOrder(request *SaveBizOrderRequest) (_result *SaveBizOrderResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &SaveBizOrderResponse{}
	_body, _err := client.SaveBizOrderEx(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * Description: 保存订单
 * Summary: 保存订单
 */
func (client *Client) SaveBizOrderEx(request *SaveBizOrderRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *SaveBizOrderResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &SaveBizOrderResponse{}
	_body, _err := client.DoRequest(tea.String("1.0"), tea.String("digital.logistic.biz.order.save"), tea.String("HTTPS"), tea.String("POST"), tea.String("/gateway.do"), tea.ToMap(request), headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * Description: 保存托单文件
 * Summary: 保存托单文件
 */
func (client *Client) SaveBizConsignorder(request *SaveBizConsignorderRequest) (_result *SaveBizConsignorderResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &SaveBizConsignorderResponse{}
	_body, _err := client.SaveBizConsignorderEx(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * Description: 保存托单文件
 * Summary: 保存托单文件
 */
func (client *Client) SaveBizConsignorderEx(request *SaveBizConsignorderRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *SaveBizConsignorderResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &SaveBizConsignorderResponse{}
	_body, _err := client.DoRequest(tea.String("1.0"), tea.String("digital.logistic.biz.consignorder.save"), tea.String("HTTPS"), tea.String("POST"), tea.String("/gateway.do"), tea.ToMap(request), headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * Description: 保存货物
 * Summary: 保存货物
 */
func (client *Client) SaveBizGoods(request *SaveBizGoodsRequest) (_result *SaveBizGoodsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &SaveBizGoodsResponse{}
	_body, _err := client.SaveBizGoodsEx(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * Description: 保存货物
 * Summary: 保存货物
 */
func (client *Client) SaveBizGoodsEx(request *SaveBizGoodsRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *SaveBizGoodsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &SaveBizGoodsResponse{}
	_body, _err := client.DoRequest(tea.String("1.0"), tea.String("digital.logistic.biz.goods.save"), tea.String("HTTPS"), tea.String("POST"), tea.String("/gateway.do"), tea.ToMap(request), headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * Description: 保存SO入货通知
 * Summary: 保存SO入货通知
 */
func (client *Client) SaveBizSonotify(request *SaveBizSonotifyRequest) (_result *SaveBizSonotifyResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &SaveBizSonotifyResponse{}
	_body, _err := client.SaveBizSonotifyEx(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * Description: 保存SO入货通知
 * Summary: 保存SO入货通知
 */
func (client *Client) SaveBizSonotifyEx(request *SaveBizSonotifyRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *SaveBizSonotifyResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &SaveBizSonotifyResponse{}
	_body, _err := client.DoRequest(tea.String("1.0"), tea.String("digital.logistic.biz.sonotify.save"), tea.String("HTTPS"), tea.String("POST"), tea.String("/gateway.do"), tea.ToMap(request), headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * Description: 保存订舱单
 * Summary: 保存订舱单
 */
func (client *Client) SaveBizBookingorder(request *SaveBizBookingorderRequest) (_result *SaveBizBookingorderResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &SaveBizBookingorderResponse{}
	_body, _err := client.SaveBizBookingorderEx(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * Description: 保存订舱单
 * Summary: 保存订舱单
 */
func (client *Client) SaveBizBookingorderEx(request *SaveBizBookingorderRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *SaveBizBookingorderResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &SaveBizBookingorderResponse{}
	_body, _err := client.DoRequest(tea.String("1.0"), tea.String("digital.logistic.biz.bookingorder.save"), tea.String("HTTPS"), tea.String("POST"), tea.String("/gateway.do"), tea.ToMap(request), headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * Description: 保存集装箱
 * Summary: 保存集装箱
 */
func (client *Client) SaveBizContainer(request *SaveBizContainerRequest) (_result *SaveBizContainerResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &SaveBizContainerResponse{}
	_body, _err := client.SaveBizContainerEx(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * Description: 保存集装箱
 * Summary: 保存集装箱
 */
func (client *Client) SaveBizContainerEx(request *SaveBizContainerRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *SaveBizContainerResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &SaveBizContainerResponse{}
	_body, _err := client.DoRequest(tea.String("1.0"), tea.String("digital.logistic.biz.container.save"), tea.String("HTTPS"), tea.String("POST"), tea.String("/gateway.do"), tea.ToMap(request), headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * Description: 保存报关单
 * Summary: 保存报关单
 */
func (client *Client) SaveBizCustomsorder(request *SaveBizCustomsorderRequest) (_result *SaveBizCustomsorderResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &SaveBizCustomsorderResponse{}
	_body, _err := client.SaveBizCustomsorderEx(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * Description: 保存报关单
 * Summary: 保存报关单
 */
func (client *Client) SaveBizCustomsorderEx(request *SaveBizCustomsorderRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *SaveBizCustomsorderResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &SaveBizCustomsorderResponse{}
	_body, _err := client.DoRequest(tea.String("1.0"), tea.String("digital.logistic.biz.customsorder.save"), tea.String("HTTPS"), tea.String("POST"), tea.String("/gateway.do"), tea.ToMap(request), headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * Description: 保存拖车单
 * Summary: 保存拖车单
 */
func (client *Client) SaveBizVehicle(request *SaveBizVehicleRequest) (_result *SaveBizVehicleResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &SaveBizVehicleResponse{}
	_body, _err := client.SaveBizVehicleEx(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * Description: 保存拖车单
 * Summary: 保存拖车单
 */
func (client *Client) SaveBizVehicleEx(request *SaveBizVehicleRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *SaveBizVehicleResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &SaveBizVehicleResponse{}
	_body, _err := client.DoRequest(tea.String("1.0"), tea.String("digital.logistic.biz.vehicle.save"), tea.String("HTTPS"), tea.String("POST"), tea.String("/gateway.do"), tea.ToMap(request), headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * Description: 保存master提单
 * Summary: 保存master提单
 */
func (client *Client) SaveBizMasterbl(request *SaveBizMasterblRequest) (_result *SaveBizMasterblResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &SaveBizMasterblResponse{}
	_body, _err := client.SaveBizMasterblEx(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * Description: 保存master提单
 * Summary: 保存master提单
 */
func (client *Client) SaveBizMasterblEx(request *SaveBizMasterblRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *SaveBizMasterblResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &SaveBizMasterblResponse{}
	_body, _err := client.DoRequest(tea.String("1.0"), tea.String("digital.logistic.biz.masterbl.save"), tea.String("HTTPS"), tea.String("POST"), tea.String("/gateway.do"), tea.ToMap(request), headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * Description: 内审完成
 * Summary: 内审完成
 */
func (client *Client) FinishBizAudit(request *FinishBizAuditRequest) (_result *FinishBizAuditResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &FinishBizAuditResponse{}
	_body, _err := client.FinishBizAuditEx(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * Description: 内审完成
 * Summary: 内审完成
 */
func (client *Client) FinishBizAuditEx(request *FinishBizAuditRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *FinishBizAuditResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &FinishBizAuditResponse{}
	_body, _err := client.DoRequest(tea.String("1.0"), tea.String("digital.logistic.biz.audit.finish"), tea.String("HTTPS"), tea.String("POST"), tea.String("/gateway.do"), tea.ToMap(request), headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * Description: 保存house提单
 * Summary: 保存house提单
 */
func (client *Client) SaveBizHousebl(request *SaveBizHouseblRequest) (_result *SaveBizHouseblResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &SaveBizHouseblResponse{}
	_body, _err := client.SaveBizHouseblEx(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * Description: 保存house提单
 * Summary: 保存house提单
 */
func (client *Client) SaveBizHouseblEx(request *SaveBizHouseblRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *SaveBizHouseblResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &SaveBizHouseblResponse{}
	_body, _err := client.DoRequest(tea.String("1.0"), tea.String("digital.logistic.biz.housebl.save"), tea.String("HTTPS"), tea.String("POST"), tea.String("/gateway.do"), tea.ToMap(request), headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * Description: 创建应付账单
 * Summary: 创建应付账单(已下)
 */
func (client *Client) CreateBillPaybillorder(request *CreateBillPaybillorderRequest) (_result *CreateBillPaybillorderResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateBillPaybillorderResponse{}
	_body, _err := client.CreateBillPaybillorderEx(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * Description: 创建应付账单
 * Summary: 创建应付账单(已下)
 */
func (client *Client) CreateBillPaybillorderEx(request *CreateBillPaybillorderRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *CreateBillPaybillorderResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &CreateBillPaybillorderResponse{}
	_body, _err := client.DoRequest(tea.String("1.0"), tea.String("digital.logistic.bill.paybillorder.create"), tea.String("HTTPS"), tea.String("POST"), tea.String("/gateway.do"), tea.ToMap(request), headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * Description: 创建应收账单
 * Summary: 创建应收账单(已下)
 */
func (client *Client) CreateBillReceiptbillorder(request *CreateBillReceiptbillorderRequest) (_result *CreateBillReceiptbillorderResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateBillReceiptbillorderResponse{}
	_body, _err := client.CreateBillReceiptbillorderEx(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * Description: 创建应收账单
 * Summary: 创建应收账单(已下)
 */
func (client *Client) CreateBillReceiptbillorderEx(request *CreateBillReceiptbillorderRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *CreateBillReceiptbillorderResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &CreateBillReceiptbillorderResponse{}
	_body, _err := client.DoRequest(tea.String("1.0"), tea.String("digital.logistic.bill.receiptbillorder.create"), tea.String("HTTPS"), tea.String("POST"), tea.String("/gateway.do"), tea.ToMap(request), headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * Description: 保存应付资费项
 * Summary: 保存应付资费项
 */
func (client *Client) SaveBillPaybilltariff(request *SaveBillPaybilltariffRequest) (_result *SaveBillPaybilltariffResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &SaveBillPaybilltariffResponse{}
	_body, _err := client.SaveBillPaybilltariffEx(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * Description: 保存应付资费项
 * Summary: 保存应付资费项
 */
func (client *Client) SaveBillPaybilltariffEx(request *SaveBillPaybilltariffRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *SaveBillPaybilltariffResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &SaveBillPaybilltariffResponse{}
	_body, _err := client.DoRequest(tea.String("1.0"), tea.String("digital.logistic.bill.paybilltariff.save"), tea.String("HTTPS"), tea.String("POST"), tea.String("/gateway.do"), tea.ToMap(request), headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * Description: 保存应收资费项
 * Summary: 保存应收资费项
 */
func (client *Client) SaveBillReceiptbilltariff(request *SaveBillReceiptbilltariffRequest) (_result *SaveBillReceiptbilltariffResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &SaveBillReceiptbilltariffResponse{}
	_body, _err := client.SaveBillReceiptbilltariffEx(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * Description: 保存应收资费项
 * Summary: 保存应收资费项
 */
func (client *Client) SaveBillReceiptbilltariffEx(request *SaveBillReceiptbilltariffRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *SaveBillReceiptbilltariffResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &SaveBillReceiptbilltariffResponse{}
	_body, _err := client.DoRequest(tea.String("1.0"), tea.String("digital.logistic.bill.receiptbilltariff.save"), tea.String("HTTPS"), tea.String("POST"), tea.String("/gateway.do"), tea.ToMap(request), headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * Description: 核销应付资费项
 * Summary: 核销应付资费项
 */
func (client *Client) VerifyBillPaybill(request *VerifyBillPaybillRequest) (_result *VerifyBillPaybillResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &VerifyBillPaybillResponse{}
	_body, _err := client.VerifyBillPaybillEx(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * Description: 核销应付资费项
 * Summary: 核销应付资费项
 */
func (client *Client) VerifyBillPaybillEx(request *VerifyBillPaybillRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *VerifyBillPaybillResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &VerifyBillPaybillResponse{}
	_body, _err := client.DoRequest(tea.String("1.0"), tea.String("digital.logistic.bill.paybill.verify"), tea.String("HTTPS"), tea.String("POST"), tea.String("/gateway.do"), tea.ToMap(request), headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * Description: 核销应收资费项
 * Summary: 核销应收资费项
 */
func (client *Client) VerifyBillReceiptbillorder(request *VerifyBillReceiptbillorderRequest) (_result *VerifyBillReceiptbillorderResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &VerifyBillReceiptbillorderResponse{}
	_body, _err := client.VerifyBillReceiptbillorderEx(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * Description: 核销应收资费项
 * Summary: 核销应收资费项
 */
func (client *Client) VerifyBillReceiptbillorderEx(request *VerifyBillReceiptbillorderRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *VerifyBillReceiptbillorderResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &VerifyBillReceiptbillorderResponse{}
	_body, _err := client.DoRequest(tea.String("1.0"), tea.String("digital.logistic.bill.receiptbillorder.verify"), tea.String("HTTPS"), tea.String("POST"), tea.String("/gateway.do"), tea.ToMap(request), headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * Description: 更新应付账单
 * Summary: 更新应付账单
 */
func (client *Client) UpdateBillPaybillorder(request *UpdateBillPaybillorderRequest) (_result *UpdateBillPaybillorderResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpdateBillPaybillorderResponse{}
	_body, _err := client.UpdateBillPaybillorderEx(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * Description: 更新应付账单
 * Summary: 更新应付账单
 */
func (client *Client) UpdateBillPaybillorderEx(request *UpdateBillPaybillorderRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *UpdateBillPaybillorderResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &UpdateBillPaybillorderResponse{}
	_body, _err := client.DoRequest(tea.String("1.0"), tea.String("digital.logistic.bill.paybillorder.update"), tea.String("HTTPS"), tea.String("POST"), tea.String("/gateway.do"), tea.ToMap(request), headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * Description: 更新应收账单
 * Summary: 更新应收账单
 */
func (client *Client) UpdateBillReceiptbillorder(request *UpdateBillReceiptbillorderRequest) (_result *UpdateBillReceiptbillorderResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpdateBillReceiptbillorderResponse{}
	_body, _err := client.UpdateBillReceiptbillorderEx(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * Description: 更新应收账单
 * Summary: 更新应收账单
 */
func (client *Client) UpdateBillReceiptbillorderEx(request *UpdateBillReceiptbillorderRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *UpdateBillReceiptbillorderResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &UpdateBillReceiptbillorderResponse{}
	_body, _err := client.DoRequest(tea.String("1.0"), tea.String("digital.logistic.bill.receiptbillorder.update"), tea.String("HTTPS"), tea.String("POST"), tea.String("/gateway.do"), tea.ToMap(request), headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * Description: 保存应付发票
 * Summary: 保存应付发票
 */
func (client *Client) SaveBillPayinvoice(request *SaveBillPayinvoiceRequest) (_result *SaveBillPayinvoiceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &SaveBillPayinvoiceResponse{}
	_body, _err := client.SaveBillPayinvoiceEx(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * Description: 保存应付发票
 * Summary: 保存应付发票
 */
func (client *Client) SaveBillPayinvoiceEx(request *SaveBillPayinvoiceRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *SaveBillPayinvoiceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &SaveBillPayinvoiceResponse{}
	_body, _err := client.DoRequest(tea.String("1.0"), tea.String("digital.logistic.bill.payinvoice.save"), tea.String("HTTPS"), tea.String("POST"), tea.String("/gateway.do"), tea.ToMap(request), headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * Description: 保存应收发票
 * Summary: 保存应收发票
 */
func (client *Client) SaveBillReceiptinvoice(request *SaveBillReceiptinvoiceRequest) (_result *SaveBillReceiptinvoiceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &SaveBillReceiptinvoiceResponse{}
	_body, _err := client.SaveBillReceiptinvoiceEx(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * Description: 保存应收发票
 * Summary: 保存应收发票
 */
func (client *Client) SaveBillReceiptinvoiceEx(request *SaveBillReceiptinvoiceRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *SaveBillReceiptinvoiceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &SaveBillReceiptinvoiceResponse{}
	_body, _err := client.DoRequest(tea.String("1.0"), tea.String("digital.logistic.bill.receiptinvoice.save"), tea.String("HTTPS"), tea.String("POST"), tea.String("/gateway.do"), tea.ToMap(request), headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * Description: 上传历史数据
 * Summary: 上传历史数据
 */
func (client *Client) UploadBizFinancing(request *UploadBizFinancingRequest) (_result *UploadBizFinancingResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UploadBizFinancingResponse{}
	_body, _err := client.UploadBizFinancingEx(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * Description: 上传历史数据
 * Summary: 上传历史数据
 */
func (client *Client) UploadBizFinancingEx(request *UploadBizFinancingRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *UploadBizFinancingResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &UploadBizFinancingResponse{}
	_body, _err := client.DoRequest(tea.String("1.0"), tea.String("digital.logistic.biz.financing.upload"), tea.String("HTTPS"), tea.String("POST"), tea.String("/gateway.do"), tea.ToMap(request), headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * Description: 上传订单
 * Summary: 上传订单
 */
func (client *Client) UploadBizOrder(request *UploadBizOrderRequest) (_result *UploadBizOrderResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UploadBizOrderResponse{}
	_body, _err := client.UploadBizOrderEx(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * Description: 上传订单
 * Summary: 上传订单
 */
func (client *Client) UploadBizOrderEx(request *UploadBizOrderRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *UploadBizOrderResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &UploadBizOrderResponse{}
	_body, _err := client.DoRequest(tea.String("1.0"), tea.String("digital.logistic.biz.order.upload"), tea.String("HTTPS"), tea.String("POST"), tea.String("/gateway.do"), tea.ToMap(request), headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * Description: 创建船公司did账户
 * Summary: 创建船公司did账户
 */
func (client *Client) CreateDidCarrier(request *CreateDidCarrierRequest) (_result *CreateDidCarrierResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateDidCarrierResponse{}
	_body, _err := client.CreateDidCarrierEx(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * Description: 创建船公司did账户
 * Summary: 创建船公司did账户
 */
func (client *Client) CreateDidCarrierEx(request *CreateDidCarrierRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *CreateDidCarrierResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &CreateDidCarrierResponse{}
	_body, _err := client.DoRequest(tea.String("1.0"), tea.String("digital.logistic.did.carrier.create"), tea.String("HTTPS"), tea.String("POST"), tea.String("/gateway.do"), tea.ToMap(request), headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * Description:  货代授权
 * Summary:  货代授权
 */
func (client *Client) AuthSysForwarder(request *AuthSysForwarderRequest) (_result *AuthSysForwarderResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &AuthSysForwarderResponse{}
	_body, _err := client.AuthSysForwarderEx(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * Description:  货代授权
 * Summary:  货代授权
 */
func (client *Client) AuthSysForwarderEx(request *AuthSysForwarderRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *AuthSysForwarderResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &AuthSysForwarderResponse{}
	_body, _err := client.DoRequest(tea.String("1.0"), tea.String("digital.logistic.sys.forwarder.auth"), tea.String("HTTPS"), tea.String("POST"), tea.String("/gateway.do"), tea.ToMap(request), headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * Description: master提单信息查询
 * Summary: master提单信息查询
 */
func (client *Client) QueryBizMasterbl(request *QueryBizMasterblRequest) (_result *QueryBizMasterblResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &QueryBizMasterblResponse{}
	_body, _err := client.QueryBizMasterblEx(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * Description: master提单信息查询
 * Summary: master提单信息查询
 */
func (client *Client) QueryBizMasterblEx(request *QueryBizMasterblRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *QueryBizMasterblResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &QueryBizMasterblResponse{}
	_body, _err := client.DoRequest(tea.String("1.0"), tea.String("digital.logistic.biz.masterbl.query"), tea.String("HTTPS"), tea.String("POST"), tea.String("/gateway.do"), tea.ToMap(request), headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * Description: 应付发票文件上传接口
 * Summary: 应付发票文件上传接口
 */
func (client *Client) SaveBizPayinvoicefile(request *SaveBizPayinvoicefileRequest) (_result *SaveBizPayinvoicefileResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &SaveBizPayinvoicefileResponse{}
	_body, _err := client.SaveBizPayinvoicefileEx(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * Description: 应付发票文件上传接口
 * Summary: 应付发票文件上传接口
 */
func (client *Client) SaveBizPayinvoicefileEx(request *SaveBizPayinvoicefileRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *SaveBizPayinvoicefileResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &SaveBizPayinvoicefileResponse{}
	_body, _err := client.DoRequest(tea.String("1.0"), tea.String("digital.logistic.biz.payinvoicefile.save"), tea.String("HTTPS"), tea.String("POST"), tea.String("/gateway.do"), tea.ToMap(request), headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * Description: 托运订单保存接口
 * Summary: 托运订单保存接口
 */
func (client *Client) SaveBiznewOrder(request *SaveBiznewOrderRequest) (_result *SaveBiznewOrderResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &SaveBiznewOrderResponse{}
	_body, _err := client.SaveBiznewOrderEx(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * Description: 托运订单保存接口
 * Summary: 托运订单保存接口
 */
func (client *Client) SaveBiznewOrderEx(request *SaveBiznewOrderRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *SaveBiznewOrderResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &SaveBiznewOrderResponse{}
	_body, _err := client.DoRequest(tea.String("1.0"), tea.String("digital.logistic.biznew.order.save"), tea.String("HTTPS"), tea.String("POST"), tea.String("/gateway.do"), tea.ToMap(request), headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * Description: 订舱单创建接口
 * Summary: 订舱单创建接口
 */
func (client *Client) SaveBiznewBooking(request *SaveBiznewBookingRequest) (_result *SaveBiznewBookingResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &SaveBiznewBookingResponse{}
	_body, _err := client.SaveBiznewBookingEx(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * Description: 订舱单创建接口
 * Summary: 订舱单创建接口
 */
func (client *Client) SaveBiznewBookingEx(request *SaveBiznewBookingRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *SaveBiznewBookingResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &SaveBiznewBookingResponse{}
	_body, _err := client.DoRequest(tea.String("1.0"), tea.String("digital.logistic.biznew.booking.save"), tea.String("HTTPS"), tea.String("POST"), tea.String("/gateway.do"), tea.ToMap(request), headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * Description: 拖车单保存接口
 * Summary: 拖车单保存接口
 */
func (client *Client) SaveBiznewVehicle(request *SaveBiznewVehicleRequest) (_result *SaveBiznewVehicleResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &SaveBiznewVehicleResponse{}
	_body, _err := client.SaveBiznewVehicleEx(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * Description: 拖车单保存接口
 * Summary: 拖车单保存接口
 */
func (client *Client) SaveBiznewVehicleEx(request *SaveBiznewVehicleRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *SaveBiznewVehicleResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &SaveBiznewVehicleResponse{}
	_body, _err := client.DoRequest(tea.String("1.0"), tea.String("digital.logistic.biznew.vehicle.save"), tea.String("HTTPS"), tea.String("POST"), tea.String("/gateway.do"), tea.ToMap(request), headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * Description: 报关单保存接口
 * Summary: 报关单保存接口
 */
func (client *Client) SaveBiznewCustoms(request *SaveBiznewCustomsRequest) (_result *SaveBiznewCustomsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &SaveBiznewCustomsResponse{}
	_body, _err := client.SaveBiznewCustomsEx(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * Description: 报关单保存接口
 * Summary: 报关单保存接口
 */
func (client *Client) SaveBiznewCustomsEx(request *SaveBiznewCustomsRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *SaveBiznewCustomsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &SaveBiznewCustomsResponse{}
	_body, _err := client.DoRequest(tea.String("1.0"), tea.String("digital.logistic.biznew.customs.save"), tea.String("HTTPS"), tea.String("POST"), tea.String("/gateway.do"), tea.ToMap(request), headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * Description: 航运提单保存接口
 * Summary: 航运提单保存接口
 */
func (client *Client) SaveBiznewMaster(request *SaveBiznewMasterRequest) (_result *SaveBiznewMasterResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &SaveBiznewMasterResponse{}
	_body, _err := client.SaveBiznewMasterEx(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * Description: 航运提单保存接口
 * Summary: 航运提单保存接口
 */
func (client *Client) SaveBiznewMasterEx(request *SaveBiznewMasterRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *SaveBiznewMasterResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &SaveBiznewMasterResponse{}
	_body, _err := client.DoRequest(tea.String("1.0"), tea.String("digital.logistic.biznew.master.save"), tea.String("HTTPS"), tea.String("POST"), tea.String("/gateway.do"), tea.ToMap(request), headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * Description: 应付账单保存接口
 * Summary: 应付账单保存接口
 */
func (client *Client) SaveBiznewPaybillorder(request *SaveBiznewPaybillorderRequest) (_result *SaveBiznewPaybillorderResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &SaveBiznewPaybillorderResponse{}
	_body, _err := client.SaveBiznewPaybillorderEx(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * Description: 应付账单保存接口
 * Summary: 应付账单保存接口
 */
func (client *Client) SaveBiznewPaybillorderEx(request *SaveBiznewPaybillorderRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *SaveBiznewPaybillorderResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &SaveBiznewPaybillorderResponse{}
	_body, _err := client.DoRequest(tea.String("1.0"), tea.String("digital.logistic.biznew.paybillorder.save"), tea.String("HTTPS"), tea.String("POST"), tea.String("/gateway.do"), tea.ToMap(request), headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * Description: 应收账单保存接口
 * Summary: 应收账单保存接口
 */
func (client *Client) SaveBiznewReceiptbillorder(request *SaveBiznewReceiptbillorderRequest) (_result *SaveBiznewReceiptbillorderResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &SaveBiznewReceiptbillorderResponse{}
	_body, _err := client.SaveBiznewReceiptbillorderEx(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * Description: 应收账单保存接口
 * Summary: 应收账单保存接口
 */
func (client *Client) SaveBiznewReceiptbillorderEx(request *SaveBiznewReceiptbillorderRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *SaveBiznewReceiptbillorderResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &SaveBiznewReceiptbillorderResponse{}
	_body, _err := client.DoRequest(tea.String("1.0"), tea.String("digital.logistic.biznew.receiptbillorder.save"), tea.String("HTTPS"), tea.String("POST"), tea.String("/gateway.do"), tea.ToMap(request), headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * Description: 航运发票保存接口
 * Summary: 航运发票保存接口
 */
func (client *Client) SaveBiznewInvoice(request *SaveBiznewInvoiceRequest) (_result *SaveBiznewInvoiceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &SaveBiznewInvoiceResponse{}
	_body, _err := client.SaveBiznewInvoiceEx(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * Description: 航运发票保存接口
 * Summary: 航运发票保存接口
 */
func (client *Client) SaveBiznewInvoiceEx(request *SaveBiznewInvoiceRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *SaveBiznewInvoiceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &SaveBiznewInvoiceResponse{}
	_body, _err := client.DoRequest(tea.String("1.0"), tea.String("digital.logistic.biznew.invoice.save"), tea.String("HTTPS"), tea.String("POST"), tea.String("/gateway.do"), tea.ToMap(request), headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * Description: 上传电子提单
 * Summary: 上传电子提单
 */
func (client *Client) UploadShippingEbl(request *UploadShippingEblRequest) (_result *UploadShippingEblResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UploadShippingEblResponse{}
	_body, _err := client.UploadShippingEblEx(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * Description: 上传电子提单
 * Summary: 上传电子提单
 */
func (client *Client) UploadShippingEblEx(request *UploadShippingEblRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *UploadShippingEblResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &UploadShippingEblResponse{}
	_body, _err := client.DoRequest(tea.String("1.0"), tea.String("digital.logistic.shipping.ebl.upload"), tea.String("HTTPS"), tea.String("POST"), tea.String("/gateway.do"), tea.ToMap(request), headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * Description: 电子提单批次上传，一个I批次下传递多个提单
 * Summary: 电子提单批次上传
 */
func (client *Client) UploadShippingEblbatch(request *UploadShippingEblbatchRequest) (_result *UploadShippingEblbatchResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UploadShippingEblbatchResponse{}
	_body, _err := client.UploadShippingEblbatchEx(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * Description: 电子提单批次上传，一个I批次下传递多个提单
 * Summary: 电子提单批次上传
 */
func (client *Client) UploadShippingEblbatchEx(request *UploadShippingEblbatchRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *UploadShippingEblbatchResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &UploadShippingEblbatchResponse{}
	_body, _err := client.DoRequest(tea.String("1.0"), tea.String("digital.logistic.shipping.eblbatch.upload"), tea.String("HTTPS"), tea.String("POST"), tea.String("/gateway.do"), tea.ToMap(request), headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

/**
 * Description: 感知收货人提单状态的变更
 * Summary: 电子提单批次状态变更
 */
func (client *Client) UpdateShippingEblbatchstatus(request *UpdateShippingEblbatchstatusRequest) (_result *UpdateShippingEblbatchstatusResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpdateShippingEblbatchstatusResponse{}
	_body, _err := client.UpdateShippingEblbatchstatusEx(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

/**
 * Description: 感知收货人提单状态的变更
 * Summary: 电子提单批次状态变更
 */
func (client *Client) UpdateShippingEblbatchstatusEx(request *UpdateShippingEblbatchstatusRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *UpdateShippingEblbatchstatusResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	_result = &UpdateShippingEblbatchstatusResponse{}
	_body, _err := client.DoRequest(tea.String("1.0"), tea.String("digital.logistic.shipping.eblbatchstatus.update"), tea.String("HTTPS"), tea.String("POST"), tea.String("/gateway.do"), tea.ToMap(request), headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}
