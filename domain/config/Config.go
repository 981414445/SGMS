package config

// type configKV struct{ Key, Value string }

// func HttpUrl() string {
// 	port := ""
// 	if 80 != ProxyPort {
// 		port = strconv.Itoa(ProxyPort)
// 	}
// 	return "http://" + AppDoamin + port
// }

// func init() {
// 	Load()
// }

// func Load() {
// 	mysql := db.InitMysql()
// 	defer mysql.Db.Close()
// 	var kvs []configKV
// 	_, err := mysql.Select(&kvs, "select `Key`,Value from Config")
// 	if nil != err {
// 		panic(err)
// 	}
// 	//	var err error
// 	for _, item := range kvs {
// 		switch item.Key {
// 		case "Version":
// 			Version = item.Value
// 		case "UserIcon1":
// 			UserIcon1 = item.Value
// 		case "UserIcon2":
// 			UserIcon2 = item.Value
// 		case "UserIcon3":
// 			UserIcon3 = item.Value
// 		case "AppDoamin":
// 			AppDoamin = item.Value
// 		case "AppPort":
// 			AppPort, err = strconv.Atoi(item.Value)
// 		case "IsDevMode":
// 			IsDevMode, _ = strconv.ParseBool(item.Value)
// 		case "WeixinServiceAppSecret":
// 			WeixinServiceAppSecret = item.Value
// 		case "WeixinServiceAppId":
// 			WeixinServiceAppId = item.Value
// 		case "WeixinSigninAppId":
// 			WeixinSigninAppId = item.Value
// 		case "WeixinSigninAppSecret":
// 			WeixinSigninAppSecret = item.Value
// 		case "GameTimer":
// 			GameTimer, err = strconv.Atoi(item.Value)
// 		case "EmailSupportSmtp":
// 			EmailSupportSmtp = item.Value
// 		case "EmailSupportSsl":
// 			EmailSupportSsl = item.Value
// 		case "EmailSupportPort":
// 			EmailSupportPort, err = strconv.Atoi(item.Value)
// 		case "EmailSupportUsername":
// 			EmailSupportUsername = item.Value
// 		case "EmailSupportPassword":
// 			EmailSupportPassword = item.Value
// 		case "PolyvUserId":
// 			PolyvUserId = item.Value
// 		case "PolyvWriteToken":
// 			PolyvWriteToken = item.Value
// 		case "PolyvReadToken":
// 			PolyvReadToken = item.Value
// 		case "PolyvSecretkey":
// 			PolyvSecretkey = item.Value
// 		case "LuosimaoKey":
// 			LuosimaoKey = item.Value
// 		case "Redis":
// 			Redis = item.Value
// 		case "VerifyCodeExpire":
// 			VerifyCodeExpire, err = strconv.ParseInt(item.Value, 10, 64)
// 		case "UploadRootDir":
// 			UploadRootDir = item.Value
// 		case "UploadDirName":
// 			UploadDirName = item.Value
// 		case "WeixinPayAppId":
// 			WeixinPayAppId = item.Value
// 		case "WeixinPayAppKey":
// 			WeixinPayAppKey = item.Value
// 		case "WeixinPayPayKey":
// 			WeixinPayPayKey = item.Value
// 		case "WeixinPayMchId":
// 			WeixinPayMchId = item.Value
// 		case "WeixinPayNotifyUrl":
// 			WeixinPayNotifyUrl = item.Value
// 		case "ProxyPort":
// 			ProxyPort, err = strconv.Atoi(item.Value)
// 		case "AlipayPrivateKey":
// 			AlipayPrivateKey = item.Value
// 		case "AlipayPublicKey":
// 			AlipayPublicKey = item.Value
// 		case "AlipayAppId":
// 			AlipayAppId = item.Value
// 		case "AlipayPartner":
// 			AlipayPartner = item.Value
// 		case "AlipayNotifyUrl":
// 			AlipayNotifyUrl = item.Value
// 		case "WeixinServiceToken":
// 			WeixinServiceToken = item.Value
// 		case "WsPort":
// 			WsPort, err = strconv.Atoi(item.Value)
// 		case "VideoHost":
// 			VideoHost = item.Value
// 		case "AutoLoginAESKey":
// 			AutoLoginAESKey = item.Value
// 		case "AutoLoginMaxAge":
// 			AutoLoginMaxAge, err = strconv.Atoi(item.Value)
// 		case "ActivityGroupSize":
// 			ActivityGroupSize, err = strconv.Atoi(item.Value)
// 		case "CapacityTestTime":
// 			CapacityTestTime, err = strconv.Atoi(item.Value)
// 		case "CapacityTestQuizCount":
// 			CapacityTestQuizCount, err = strconv.Atoi(item.Value)
// 		case "CapacityTestInteval":
// 			CapacityTestInteval, err = strconv.Atoi(item.Value)
// 		case "PolyvAppId":
// 			PolyvAppId = item.Value
// 		case "PolyvAppSecret":
// 			PolyvAppSecret = item.Value
// 		case "EvaluationQuizCount":
// 			EvaluationQuizCount, err = strconv.Atoi(item.Value)
// 		case "EvaluationInterval":
// 			EvaluationInterval, err = strconv.ParseFloat(item.Value, 64)
// 		case "EvaluationFalseCo":
// 			EvaluationFalseCo, err = strconv.ParseFloat(item.Value, 64)
// 		}
// 	}
// 	if nil != err {
// 		panic(err)
// 	}
// }
