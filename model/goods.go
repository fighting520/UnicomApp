package model

type Goods struct {
	PayDetailId String // 必选	 支付订单号
	CpTradeId   String // 必选	 开发者平台交易流水号
	ProductName String //必选	 支付商品名称
	ProductType String //必选	 商品类型 1:包月 2:点播
	Price       String //必选	 单价
	Amount      String //必选	 数量
	Money       String //必选	 支付总金额
	PayStatus   String //必选	 支付状态 成功:200
	PayType     String //必选	 支付类型1:vac 2:短信
	TimeStamp   String // 必选	 支付时间
	Sign        String // 必选	 开发者验证平台信息的签名，格式是BASE64加密的数据，按照下列顺序拼装： (payDetailId=payDetailId值& cpTradeId= cpTradeId值& productName= productName值 & productType= productType值&price=price值&amount=amount值&money=money值 &payStatus=payStatus值&payType=payType值×tamp=timestamp值)使用服务器RSA公钥进行验签

}
