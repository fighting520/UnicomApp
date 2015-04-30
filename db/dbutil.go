package db

import (
	"gopkg.in/mgo.v2"
	//"gopkg.in/mgo.v2/bson"
)

var DB_NAME = "payment"

type Goods struct {
	PayDetailId string // 必选	 支付订单号
	CpTradeId   string // 必选	 开发者平台交易流水号
	ProductName string //必选	 支付商品名称
	ProductType string //必选	 商品类型 1:包月 2:点播
	Price       string //必选	 单价
	Amount      string //必选	 数量
	Money       string //必选	 支付总金额
	PayStatus   string //必选	 支付状态 成功:200
	PayType     string //必选	 支付类型1:vac 2:短信
	TimeStamp   string // 必选	 支付时间
	Sign        string // 必选	 开发者验证平台信息的签名，格式是BASE64加密的数据，按照下列顺序拼装： (payDetailId=payDetailId值& cpTradeId= cpTradeId值& productName= productName值 & productType= productType值&price=price值&amount=amount值&money=money值 &payStatus=payStatus值&payType=payType值×tamp=timestamp值)使用服务器RSA公钥进行验签

}

func Save(beam Goods) (e error) {

	session, err := mgo.Dial("localhost")
	if err != nil {
		panic(err)
	}
	defer session.Close()

	// Optional. Switch the session to a monotonic behavior.
	session.SetMode(mgo.Monotonic, true)

	c := session.DB(DB_NAME).C("goods")
	err = c.Insert(&beam)
	if err != nil {
		return err
	}

	return nil
}
