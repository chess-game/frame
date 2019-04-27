package cljson

import (
	"fmt"
	"github.com/CodisLabs/codis/pkg/utils/log"
	"testing"
)

type Object struct {
	Hello string
	Nima  []string
}

func TestJsonStream_GetArray(t *testing.T) {
	var jsonMap = make(JsonMap)
	fmt.Println("test:", jsonMap)
}

func TestJsonStream_GetMap(t *testing.T) {

	js := New([]byte(`{"markets": {"GLss": {"xx":111, "yy":{"zz":"zzzzzz"}}, "Hcw": [{"aa":"cc"}, {"bb":"dd"}], "1X2": true, "HGL": true, "HHc": true, "H1X2": true}}`))

	markets := js.GetMap("markets") //.ToTree()

	if markets == nil {
		fmt.Printf(">>> %v\n", markets)
		return
	}

	fmt.Printf(">>> %v\n", markets)
	fmt.Printf("GET:%v\n", markets.ToTree())
}

func TestJsonStream_GetStr(t *testing.T) {

	js := New([]byte(`{"HC": {"H": null, "V": null}, "HHC": {"H": null, "V": null}}`))

	fmt.Printf("json data: %v\n", js.ToMap().ToTree())
}

func TestCreateBy(t *testing.T) {
	fmt.Printf("Find: |%v|", FindKeyPos([]byte(`{"data":{"msg":"oka阿四大大哭sdas"dasd"}}`), []byte("data")))
}

func TestJsonStream_ToArray(t *testing.T) {
	//js := New([]byte(`{"msg":0, "data": "["asdasd001", "asdasd002"]"}`))
	//{"1":{"HandCards":[19,30,42]}}
	js := New([]byte(`{"FlopCards":{"2":4,"52":1,"0":4,"6":4,"7":4,"9":4,"12":4,"1":4,"5":4,"10":4,"11":4,"53":1,"3":4,"4":4,"8":4},"LastCard":[],"GrabTurnUid":0,"LastCardType":14,"GameTimes":0,"BroadCards":[30,41,40],"RoundId":"yezhurdskey_806001_8001_8e645b29-3154-4b4a-85d4-8d7349a696e3","0":{"PlayType":0,"HeadId":2,"HandCards":[30,40,41],"Income":0,"Status":4,"Doubletimes":0,"NickName":"zongdai1","Multiple":16,"BoomTimes":0,"SpringTimes":0,"Uid":2,"Balance":55186,"Integral":0,"MaxTimes":0,"WinTimes":0,"LastTime":1542866401},"1":{"Uid":4,"HeadId":4,"Integral":0,"HandCards":[26,21,1,35,10,32,53,50,33,45,42,13,24,17,2,29,4],"Status":4,"NickName":"zongdai3","LastTime":1542866401,"BoomTimes":0,"SpringTimes":0,"Income":0,"Doubletimes":0,"Multiple":16,"WinTimes":0,"PlayType":0,"Balance":13097,"MaxTimes":0},"2":{"BoomTimes":0,"WinTimes":0,"LastTime":1542866401,"Doubletimes":0,"HandCards":[18,39,8,49,12,6,38,11,9,43,47,14,0,23,44,22,48],"Status":4,"NickName":"zongdai2","Balance":48429,"Integral":0,"SpringTimes":0,"Uid":3,"HeadId":3,"Multiple":16,"MaxTimes":0,"Income":0,"PlayType":0},"GameAc":5,"LastTime":1542866437,"CurUid":0,"Integral":0,"PassTimes":0,"Nums":3,"RoomType":8001,"AcTimes":0}`))

	//arr := js.GetArray("1","HandCards").ToCustom()
	if js == nil {
		fmt.Printf("nil\n")
		return
	}
	arr := js.GetArray("0", "HandCards")
	//arr := js.GetKey("0").GetKey("HandCards").ToArray().ToCustom()
	//arr := js.GetMap("0").ToCustom()["HandCards"]
	fmt.Printf("arr[%v]\n", arr.ToCustom())
	//fmt.Printf("Arr: %v\n", arr)
}

// 测试删除某个key
func TestGetJsonValueEx(t *testing.T) {

	js := New([]byte(`{"msg":["haha","isMessageBox"], "param":"ok"}`))
	js.GetKey("param").SetStr("okokok")

	fmt.Printf("BYTE: %v\n", js.ToStr())
}

// 测试修改某个key
func TestJsonStream_SetObject(t *testing.T) {
	js := New([]byte("[{\"history_id\":0,\"price\":100,\"sport_id\":1,\"schedule_id\":2,\"bet_data\":{\"1X2\":{\"k\":\"H\",\"p\":\"2.5\"}}}]"))

	fmt.Printf(">> %v", js.ToArray().GetOffset(0).ToStr())
}

// 测试删除指定索引
func TestJsonStream_DelOffset(t *testing.T) {

	js := CreateBy(M{
		"cards": M{
			"1": 4,
			"2": 4,
		},
		"param": M{
			"game_type":    0,
			"play_methods": "HC",

			"is_better": 0,
			"data":      "[{\"history_id\":0,\"price\":100,\"sport_id\":1,\"schedule_id\": 1,\"bet_data\":{\"HC\":{\"H\":{\"k\":\"-1.5/2\",\"p\": \"0.830\", \"data\":\"{\"history_id\":0}\"}}}}]"},
		"game_type": A{
			"aaa", "bbb", "ccc",
		}},
	)

	fmt.Printf("js[%s]\n", js.ToStr())
	//param := js.GetKey("game_type").ToArray()
	js.GetKey("cards", "1").SetObject(3)
	param := js.GetInt32("cards", "1")
	js.ResetOffset()
	fmt.Printf("param[%v]\n", param)
	fmt.Printf("js[%v]\n", js.ToStr())
	//param2 := js.GetInt32("cards", "1")

	json := CreateBy([]byte(js.ToStr()))
	fmt.Printf("param2[%v]\n", json.GetInt32("cards", "1"))

	//for _, val := range param.ToCustom() {
	//	fmt.Printf(" >> ITEM: %v\n", val)
	//}

	//js.ResetOffset()
	//fmt.Printf("%v\n", js.GetKey("param", "game_type").ToStr() )

	return
}

// 测试删除指定索引
func TestJsonStream_DelOffset2(t *testing.T) {

	js := CreateBy(
		M{
			"param":  M{},
			"config": M{},
			"server": M{
				"Host":      "172.17.0.1",
				"remote_ip": "127.0.0.1",
			},
			"cache": M{
				"bank_type": A{
					M{"id": 1, "bank_name": "工商银行"},
					M{"id": 2, "bank_name": "建设银行"},
					M{"id": 3, "bank_name": "招商银行"},
					M{"id": 4, "bank_name": "农业银行"},
				},
			},
		},
	)

	fmt.Printf("js[%v]", js.ToStr())
	param := js.GetKey("cache", "bank_type").ToArray()

	for _, val := range param.ToCustom() {
		fmt.Printf(" >> ITEM: %v\n", val)
	}

	js.ResetOffset()

	fmt.Printf("HOST:%v\n", js.GetKey("server", "Host").ToStr())
	return
}

func TestJsonMap_ToTree(t *testing.T) {

	js := CreateBy(M{
		"a1": M{
			"b1": M{
				"c1": A{
					"111", "112", "113",
				},
			},
		},
		"a2": M{
			"b2": M{
				"c2": A{
					"111", "112", "113",
				},
			},
		},
	})

	fmt.Printf("TREE: %v", js.ToMap().ToTree())

}

func TestJsonStream_GetMap2(t *testing.T) {
	jsonStr := `{"GL":{"OV":[{"k":"2","p":"0.700","t":"e"}],"UN":[{"k":"2","p":"1.120","t":"e"}]},"HC":{"H":[{"k":"-0.5","p":"0.770","t":"e"}],"V":[{"k":"0.5","p":"1.070","t":"e"}]},"1X2":[{"k":"H","p":"1.77","t":"e"},{"k":"V","p":"4.20","t":"e"},{"k":"X","p":"3.35","t":"e"}],"APP":[{"GL":{"OV":{"k":"2","p":"0.700","t":"e"},"UN":{"k":"2","p":"1.120","t":"e"}},"HC":{"H":{"k":"-0.5","p":"0.770","t":"e"},"V":{"k":"0.5","p":"1.070","t":"e"}},"1X2":[{"k":"H","p":"1.77","t":"e"},{"k":"V","p":"4.20","t":"e"},{"k":"X","p":"3.35","t":"e"}],"HGL":{"OV":{"k1":"0.5/1","p":"0.700","t":"e"},"UN":{"k":"0.5/1","p":"1.120","t":"e"}},"HHC":{"H":{"k":"-0/0.5","p":"1.020","t":"e"},"V":{"k":"0/0.5","p":"0.820","t":"e"}},"H1X2":[{"k":"H","p":"2.42","t":"e"},{"k":"V","p":"4.70","t":"e"},{"k":"X","p":"2.02","t":"e"}],"TGOE":[{"k":"Odd","p":"1.96","t":"e"},{"k":"Even","p":"1.91","t":"e"}],"HTGOE":[{"k":"Odd","p":"/","t":"e"},{"k":"Even","p":"/","t":"e"}]}],"HGL":{"OV":[{"k2":"0.5/1","p":"0.700","t":"e"}],"UN":[{"k":"0.5/1","p":"1.120","t":"e"}]},"HHC":{"H":[{"k":"-0/0.5","p":"1.020","t":"e"}],"V":[{"k":"0/0.5","p":"0.820","t":"e"}]},"H1X2":[{"k":"H","p":"2.42","t":"e"},{"k":"V","p":"4.70","t":"e"},{"k":"X","p":"2.02","t":"e"}],"TGOE":[{"k":"Odd","p":"1.96","t":"e"},{"k":"Even","p":"1.91","t":"e"}]}`

	js := New([]byte(jsonStr))

	fmt.Printf("TARGET: %v\n", js.GetMap("HGL").Find("OV").ToStr())

}

func TestJsonStream_GetArray2(t *testing.T) {

	jsonStr := `{"quick_data":[],"payment":[{"sort":2,"icon":"quick_bank.png","utype":"1,2,3,4,9,7","id":12,"name":"????","type":0,"code":"bank_wap","status":0},{"name":"????","type":1,"code":"bank_app","status":0,"sort":3,"icon":"quick_bank.png","utype":"1,2,3,4,9,7","id":13},{"sort":12,"icon":"alipay.png","utype":"1,2,3,4,9,7,5,6,8","id":4,"name":"?????","type":0,"code":"alipay_scan","status":0},{"id":5,"name":"?????","type":1,"code":"alipay_wap","status":0,"sort":4,"icon":"alipay.png","utype":"1,2,3,4,9,7,5,6,8"},{"id":11,"name":"QQ??","type":1,"code":"qq_wap","status":0,"sort":6,"icon":"qq.png","utype":"1,2,3,4,5,9,7,6,8"},{"icon":"weixin.png","utype":"1,2,3,4,9,7,5,6,8","id":3,"name":"????","type":1,"code":"weixin_wap","status":0,"sort":5},{"sort":9,"icon":"weixin.png","utype":"1,2,3,4,9,7,5,6,8","id":2,"name":"????","type":0,"code":"weixin_scan","status":0},{"type":2,"code":"man_bank","status":0,"sort":10,"icon":"man_bank.png","utype":"1,2,3,4,5,9,7,6,8","id":15,"name":"????"},{"icon":"qq.png","utype":"1,2,3,4,5,9,7,6,8","id":7,"name":"QQ??","type":0,"code":"qq_scan","status":0,"sort":11}],"thrid_bank":null,"pay_data":[{"app_tip":"??????,?????......","device":0,"sort":10,"account":"","id":35,"remark":"","min_price":0,"qrcode":"0","show":0,"group":2,"max_price":0,"pc_tip":"??????,?????......","man":0,"nickname":"","type":2,"title":"??","utype":"1,2,3,4,9,7","icon":"pay_icon_weixin.png","name":"??????"},{"sort":11,"account":"","show":0,"min_price":0,"max_price":0,"utype":"1,2,3,4,9,7","icon":"pay_icon_weixin.png","type":3,"title":"??","qrcode":"1","pc_tip":"??????,?????......","device":1,"nickname":"","id":35,"app_tip":"??????,?????......","man":0,"group":3,"name":"????wap","remark":""},{"name":"???????","nickname":"","account":"","group":4,"device":0,"sort":7,"min_price":0,"app_tip":"??????,?????......","pc_tip":"??????,?????......","man":0,"show":0,"remark":"","max_price":0,"utype":"1,2,3,4,9,7","qrcode":"0","icon":"pay_icon_alipay.png","id":35,"type":4,"title":"???"},{"sort":17,"man":0,"show":0,"min_price":0,"max_price":0,"device":1,"icon":"pay_icon_alipay.png","type":5,"title":"???","app_tip":"??????,?????......","pc_tip":"??????,?????......","nickname":"","account":"","name":"?????wap","remark":"","qrcode":"1","utype":"1,2,3,4,9,7","id":35,"group":5},{"group":12,"type":12,"remark":"","qrcode":"0","man":0,"nickname":"","icon":"pay_icon_bankq.png","sort":0,"name":"????YQ(30~5000)","title":"????","max_price":0,"utype":"1,2,3,4,9,7","app_tip":"??????,?????......","pc_tip":"??????,?????......","device":0,"id":35,"min_price":0,"account":"","show":0},{"type":13,"remark":"","utype":"1,2,3,4,9,7","app_tip":"??????,?????......","id":35,"qrcode":"1","device":1,"man":0,"nickname":"","show":0,"name":"????YQ(30~10000)","max_price":0,"sort":6,"account":"","icon":"pay_icon_bankq.png","min_price":0,"pc_tip":"??????,?????......","group":13,"title":"????"},{"id":87,"remark":"","account":"","name":"??","man":0,"show":0,"group":2,"type":2,"title":"??","max_price":0,"utype":"1,2,3,4,5,9,7,6,8","icon":"pay_icon_weixin.png","min_price":0,"qrcode":"0","app_tip":"","pc_tip":"","device":0,"sort":4,"nickname":""},{"utype":"1,2,3,4,5,9,7,6,8","type":3,"app_tip":"","pc_tip":"","nickname":"","account":"","show":0,"name":"??","max_price":0,"remark":"","title":"??","min_price":0,"qrcode":"1","sort":4,"icon":"pay_icon_weixin.png","group":3,"id":87,"man":0,"device":1},{"id":87,"title":"???","min_price":0,"max_price":0,"man":0,"nickname":"","name":"??","remark":"","qrcode":"0","device":0,"sort":3,"type":4,"app_tip":"","group":4,"utype":"1,2,3,4,5,9,7,6,8","pc_tip":"","account":"","icon":"pay_icon_alipay.png","show":0},{"pc_tip":"","nickname":"","name":"??","remark":"","max_price":0,"app_tip":"","utype":"1,2,3,4,5,9,7,6,8","account":"","show":0,"id":87,"title":"???","qrcode":"1","man":0,"icon":"pay_icon_alipay.png","sort":7,"group":5,"type":5,"min_price":0,"device":1},{"type":7,"utype":"1,2,3,4,5,9,7,6,8","device":0,"man":0,"nickname":"","icon":"pay_icon_qq.png","title":"QQ","min_price":0,"max_price":0,"pc_tip":"","sort":0,"account":"","qrcode":"0","app_tip":"","name":"??","id":87,"remark":"","show":0,"group":7},{"account":"","show":0,"id":87,"max_price":0,"pc_tip":"","nickname":"","type":11,"remark":"","device":1,"name":"??","utype":"1,2,3,4,5,9,7,6,8","man":0,"group":11,"app_tip":"","sort":4,"icon":"pay_icon_qq.png","title":"QQ","min_price":0,"qrcode":"1"},{"name":"??????????","utype":"8,7,9,5,4,3,2,1","app_tip":"???????,????????1%","man":1,"group":0,"pc_tip":"???????,????????1%","account":"??KTV","remark":"","title":"??","min_price":0,"qrcode":"http://upload.bxvip588.com/sg04/Uploads/qrcode/06f528bac1bb3c5cc0edbc217e25903c.jpg","id":16,"max_price":0,"icon":"pay_icon_weixin2.png","show":1,"type":2,"device":2,"sort":11,"nickname":"??KTV"}],"man_bank":[{"bank_type":"??????","user_type":"1,2,3,4,5,9,7,6,8","pc_tip":"???????,???,??3????","price":0,"id":4,"real_name":"???","bank_card":"6217000010117415946","bank_type_id":2,"app_tip":"???????,???,??3????","bank_home":"?????"},{"id":7,"bank_type":"????????","bank_home":"????????????","real_name":"???","bank_card":"6210676862079669656","bank_type_id":28,"user_type":"8,6,7,9,5,4,3,2,1","pc_tip":"????????3????,????????","app_tip":"????????3????,????????","price":0}]}`

	js := New([]byte(jsonStr))

	if js == nil {
		return
	}

	arr := js.GetArray("payment")
	if arr == nil {
		return
	}

	for _, val := range arr.ToCustom() {
		fmt.Printf(">> %v\n", val)
	}
}

func TestJsonArray_GetOffset(t *testing.T) {

	jsonStr := []byte(`[{"aa":"cc"}, {"bb":"dd"}]`)

	jsonArr := New(jsonStr)

	if jsonArr == nil {
		fmt.Printf(">> 失败!!\n")
	} else {
		fmt.Printf(">> 成功!!\n")
	}

	for _, val := range jsonArr.ToArray().ToCustom() {
		fmt.Printf(">> %v\n", val)
	}

}

func TestJsonMap_Find(t *testing.T) {

	jsonStr := []byte(`{"hwh"   :"a  sda   sd", "ww":  "2223", "dd":["11aasdasd11","2222"]}`)

	jsonMap := New(jsonStr)

	if jsonMap == nil {
		fmt.Printf(">> 失败!!\n")
		return
	}

	fmt.Printf(">> 成功!!!\n")

	for key, val := range jsonMap.ToMap().ToCustom() {
		fmt.Printf(">> %v => %v\n", key, val)
	}
}

func TestJsonMap_Each(t *testing.T) {
	jsonStr := []byte(`{"CHP": [{"k": "乌拉圭/俄罗斯", "p": "1.75"}, {"k": "乌拉圭/埃及", "p": "4"}, {"k": "俄罗斯/埃及", "p": "5"}, {"k": "乌拉圭/沙特阿拉伯", "p": "17"}, {"k": "俄罗斯/沙特阿拉伯", "p": "21"}, {"k": "埃及/沙特阿拉伯", "p": "41"}]}`)

	json := New(jsonStr)

	if json == nil {
		fmt.Printf(">> 失败!!")
		return
	}

	fmt.Printf(">> 成功!!!\n")

	jsonMap := json.ToMap()
	if jsonMap == nil {
		fmt.Printf(">> 转换jsonMap失败!\n")
		return
	}

	fmt.Printf(">> 转换JsonMap成功!!! \n")
	for key, val := range jsonMap.ToCustom() {
		fmt.Printf(">> %v => %v\n", key, val)
	}

}

func TestJsonStream_GetBool2(t *testing.T) {
	json := New([]byte(`{"GL":{"OV":[{"k": "8.5", "p": "0.950"}],"UN":[{"k": "8.5", "p": "0.850"}]}, "APP":[{"GL": {"OV": {"k": "8.5", "p": "0.950"},"UN":{"k": "8.5", "p": "0.850"}}, "HC": {"H": {"k": "-", "p": ""}, "V": {"k": "", "p": ""}},"1X2": [{"k": "H", "p": "/"}, {"k": "V", "p": "/"}, {"k": "X", "p": "/"}], "HGL": {"OV": {"k": "3.5", "p": "1.050"}, "UN": {"k": "3.5", "p": "0.750"}}, "HHC": {"H": {"k": "-", "p": ""}, "V": {"k": "", "p": ""}}, "H1X2": [{"k": "H", "p": "/"}, {"k": "V", "p": "/"}, {"k": "X", "p": "/"}], "TGOE": [{"k": "Odd", "p": "/"}, {"k": "Even", "p": "/"}], "HTGOE": [{"k": "Odd", "p": "/"}, {"k": "Even", "p": "/"}]}], "HGL": {"OV": [{"k": "3.5", "p": "1.050"}], "UN": [{"k": "3.5", "p": "0.750"}]}}`))

	if json == nil {
		fmt.Printf(">> 失败!!")
		return
	}

	fmt.Printf(">> 成功!!\n")
	fmt.Printf("The json:%v\n", json.GetArray("APP"))
}

func TestJsonMap_GetFloat32(t *testing.T) {

	json := New([]byte(`{"bet":"1.2"}`))

	if json == nil {
		fmt.Printf(" >> 失败!!")
		return
	}

	fmt.Printf(">> bet: %v\n", json.ToMap().GetFloat32("bet", 0))

}

func TestJsonArray_GetFloat32(t *testing.T) {

	betJson := New([]byte(`{"GL":{"OV":[{"k": "8.5", "p": "true"}],"UN":[{"k": "8.5", "p": "0.850"}]}, "APP":[{"GL": {"OV": {"k": "8.5", "p": "true"},"UN":{"k": "8.5", "p": "0.850"}}, "HC": {"H": {"k": "-", "p": ""}, "V": {"k": "", "p": ""}},"1X2": [{"k": "H", "p": "/"}, {"k": "V", "p": "/"}, {"k": "X", "p": "/"}], "HGL": {"OV": {"k": "3.5", "p": "1.050"}, "UN": {"k": "3.5", "p": "0.750"}}, "HHC": {"H": {"k": "-", "p": ""}, "V": {"k": "", "p": ""}}, "H1X2": [{"k": "H", "p": "/"}, {"k": "V", "p": "/"}, {"k": "X", "p": "/"}], "TGOE": [{"k": "Odd", "p": "/"}, {"k": "Even", "p": "/"}], "HTGOE": [{"k": "Odd", "p": "/"}, {"k": "Even", "p": "/"}]}], "HGL": {"OV": [{"k": "3.5", "p": "1.050"}], "UN": [{"k": "3.5", "p": "0.750"}]}}`))

	if betJson == nil {
		fmt.Printf("it is nil\n")
		return
	}

	fmt.Printf("The json:%v\n", betJson.GetMap("GL"))
	oddsData := betJson.GetMap("GL")

	oddsData.Find("OV").ToArray().Each(func(key int, value *JsonStream) bool {
		fmt.Printf("value: %v\n", value.GetFloat32("k"))
		fmt.Printf("value: %v\n", value.GetBool("p"))
		return true
	})

}

func TestJsonStream_GetBool(t *testing.T) {

	json := New([]byte(`{  "key"  :   "  value" }`))

	if json == nil {
		t.Fatal("it is nil\n")
	}

	json.GetKey("key").SetObject(M{"A": "B"})
	json.GetKey("key").SetValues(123)
	fmt.Printf("key[%v]", json.GetInt32("key"))

	fmt.Printf(">> %v\n", json.ToStr())
}

func TestJsonArray_Each(t *testing.T) {

	json := New([]byte(`["128#1.0#1"]`))

	if json == nil {
		t.Fatal("it is nil\n")
	}

	fmt.Printf(">> %v\n", json.ToArray().ToCustom())

}

func TestJsonArray_IsEmpty(t *testing.T) {

	json := New([]byte(`{"key1": { "key2":"2018-01-02 23:32:22"}}`))

	if json == nil {
		t.Fatal("it is nil\n")
	}

	if !json.GetMap("key1").IsEmpty() {
		fmt.Printf("key1 is Exists!!\n")
	} else {
		fmt.Printf("key1 is Not Exists!!\n")
	}

	if !json.GetMap("key3").IsEmpty() {
		fmt.Printf("key3 is Exists!!\n")
	} else {
		fmt.Printf("key3 is Not Exists!!\n")
	}

	fmt.Printf(" >> 日期: %v\n", json.GetMap("key1").GetStr("key2", ""))

}

func TestNew(t *testing.T) {

	js := New([]byte(`{"version":"v1.0.2 Beta","facility":"gocpclient","short_message":"验证服务器回传","full_message":"[2018-05-20 18:10:24  cpsystem-client.go:126\u003e\u003eERR]\u003e\u003e 错误的JSON: ({\")","level":"2"}`))

	if js == nil {
		fmt.Printf(">> 失败!! \n")
	} else {
		fmt.Printf(">> 成功!! \n")
	}

}

func GetNewFlopCards(FlopCards *JsonStream, tag string, usercards []int) {

	if jsmap := FlopCards.GetMap(tag); jsmap != nil {
		for _, v := range usercards {
			//获取卡牌逻辑值
			if val := jsmap.GetUint32(fmt.Sprintf("%d", v), 100); val != 100 {
				log.Infof("val[%v]logic[%v]flodcard[%v]", val, v, FlopCards.ToStr())
				FlopCards.GetKey(tag, fmt.Sprintf("%d", v)).SetObject(val - 1)
			}
		}
	}

}

func TestAddKey(t *testing.T) {

	//data := "{\"key1\":{\"key2\":\"2018-01-02 23:32:22\"}}"
	//js := New([]byte(data))
	//fmt.Printf("test[%v]key1[%v]\n",js.GetInt32("test"),js.GetStr("key1","key2") )
	//fmt.Printf("js[%v]\n",js.ToStr())
	//js.Add("test", 12314)
	//fmt.Printf("js[%v]\n",js.ToStr())
	//fmt.Printf("test[%v]key1[%v]\n",js.GetInt32("test"),js.GetStr("key1","key2") )

	var usercards []int
	usercards = append(usercards, 6)
	usercards = append(usercards, 9)
	js := New([]byte(`{"FlopCards": {"6": 4,"9": 4,"0": 4}, "param":"ok"}`))
	fmt.Printf("1111111[%v]1111111\n", js.ToStr())
	fmt.Printf("map[%v]\n", js.GetMap("FlopCards").ToCustom())
	GetNewFlopCards(js, "FlopCards", usercards)
	fmt.Printf("------js[%v]------\n", js.ToStr())
	fmt.Printf("map[%v]\n", js.GetMap("FlopCards").ToCustom())
	fmt.Printf("mapstr[%v]\n", js.GetStr("FlopCards"))

}
