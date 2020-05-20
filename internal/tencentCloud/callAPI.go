package tencentcloud

import (
	"fmt"
	"os"

	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/regions"

	// cvm "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/cvm/v20170312"
	billing "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/billing/v20180709"
)

func callAPI(GETobject string) float64 {

	// 必要步骤：
	// 实例化一个认证对象，入参需要传入腾讯云账户密钥对secretId，secretKey。
	// 这里采用的是从环境变量读取的方式，需要在环境变量中先设置这两个值。
	// 你也可以直接在代码中写死密钥对，但是小心不要将代码复制、上传或者分享给他人，
	// 以免泄露密钥对危及你的财产安全。
	credential := common.NewCredential(
		os.Getenv("TENCENTCLOUD_SECRET_ID"),
		os.Getenv("TENCENTCLOUD_SECRET_KEY"),
	)

	// 非必要步骤
	// 实例化一个客户端配置对象，可以指定超时时间等配置
	cpf := profile.NewClientProfile()
	// SDK默认使用POST方法。
	// 如果你一定要使用GET方法，可以在这里设置。GET方法无法处理一些较大的请求。
	// 如非必要请不要修改默认设置。
	//cpf.HttpProfile.ReqMethod = "GET"
	// SDK有默认的超时时间，如非必要请不要修改默认设置。
	// 如有需要请在代码中查阅以获取最新的默认值。
	//cpf.HttpProfile.ReqTimeout = 10
	// SDK会自动指定域名。通常是不需要特地指定域名的，但是如果你访问的是金融区的服务，
	// 则必须手动指定域名，例如云服务器的上海金融区域名： cvm.ap-shanghai-fsi.tencentcloudapi.com
	//cpf.HttpProfile.Endpoint = "cvm.tencentcloudapi.com"
	// SDK默认用HmacSHA256进行签名，它更安全但是会轻微降低性能。
	// 如非必要请不要修改默认设置。
	//cpf.SignMethod = "HmacSHA1"
	// SDK 默认用 zh-CN 调用返回中文。此外还可以设置 en-US 返回全英文。
	// 但大部分产品或接口并不支持全英文的返回。
	// 如非必要请不要修改默认设置。
	//cpf.Language = "en-US"

	// 实例化要请求产品(以cvm为例)的client对象
	// 第二个参数是地域信息，可以直接填写字符串ap-guangzhou，或者引用预设的常量
	// client, _ := cvm.NewClient(credential, regions.Guangzhou, cpf)
	//billingclient, _ := billing.NewClient(credential, regions.Guangzhou, cpf)
	// 实例化一个请求对象，根据调用的接口和实际情况，可以进一步设置请求参数
	// 你可以直接查询SDK源码确定DescribeInstancesRequest有哪些属性可以设置，
	// 属性可能是基本类型，也可能引用了另一个数据结构。
	// 推荐使用IDE进行开发，可以方便的跳转查阅各个接口和数据结构的文档说明。
	//request := cvm.NewDescribeInstancesRequest()

	switch GETobject {
	default: 
		return 0.0
	case "AccountBalance":
		billingclient, _ := billing.NewClient(credential, regions.Guangzhou, cpf)
		// 实例化一个请求对象，根据调用的接口和实际情况，可以进一步设置请求参数
		// 你可以直接查询SDK源码确定DescribeInstancesRequest有哪些属性可以设置，
		// 属性可能是基本类型，也可能引用了另一个数据结构。
		// 推荐使用IDE进行开发，可以方便的跳转查阅各个接口和数据结构的文档说明。
		//request := cvm.NewDescribeInstancesRequest()
		balanceReqaest := billing.NewDescribeAccountBalanceRequest()

		balanceResponse, err := billingclient.DescribeAccountBalance(balanceReqaest)

		if _, ok := err.(*errors.TencentCloudSDKError); ok {
			fmt.Printf("An API error has returned: %s", err)
			return 0.0
		}

		if err != nil {
			panic(err)
		}

		fmt.Printf("%s", balanceResponse.ToJsonString())
		fmt.Println("--------Balance----------")
		var balanceResulta float64
		balanceResulta = float64(*(balanceResponse.Response.Balance)/100)
		fmt.Println(balanceResulta)
		return balanceResulta
	}
}
