package cdn

//Licensed under the Apache License, Version 2.0 (the "License");
//you may not use this file except in compliance with the License.
//You may obtain a copy of the License at
//
//http://www.apache.org/licenses/LICENSE-2.0
//
//Unless required by applicable law or agreed to in writing, software
//distributed under the License is distributed on an "AS IS" BASIS,
//WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//See the License for the specific language governing permissions and
//limitations under the License.
//
// Code generated by Alibaba Cloud SDK Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

func (client *Client) DescribeLiveStreamFrameLossRatio(request *DescribeLiveStreamFrameLossRatioRequest) (response *DescribeLiveStreamFrameLossRatioResponse, err error) {
	response = CreateDescribeLiveStreamFrameLossRatioResponse()
	err = client.DoAction(request, response)
	return
}

func (client *Client) DescribeLiveStreamFrameLossRatioWithChan(request *DescribeLiveStreamFrameLossRatioRequest) (<-chan *DescribeLiveStreamFrameLossRatioResponse, <-chan error) {
	responseChan := make(chan *DescribeLiveStreamFrameLossRatioResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeLiveStreamFrameLossRatio(request)
		responseChan <- response
		errChan <- err
	})
	if err != nil {
		errChan <- err
		close(responseChan)
		close(errChan)
	}
	return responseChan, errChan
}

func (client *Client) DescribeLiveStreamFrameLossRatioWithCallback(request *DescribeLiveStreamFrameLossRatioRequest, callback func(response *DescribeLiveStreamFrameLossRatioResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeLiveStreamFrameLossRatioResponse
		var err error
		defer close(result)
		response, err = client.DescribeLiveStreamFrameLossRatio(request)
		callback(response, err)
		result <- 1
	})
	if err != nil {
		defer close(result)
		callback(nil, err)
		result <- 0
	}
	return result
}

type DescribeLiveStreamFrameLossRatioRequest struct {
	*requests.RpcRequest
	AppName       string `position:"Query" name:"AppName"`
	SecurityToken string `position:"Query" name:"SecurityToken"`
	DomainName    string `position:"Query" name:"DomainName"`
	Action        string `position:"Query" name:"Action"`
	EndTime       string `position:"Query" name:"EndTime"`
	StartTime     string `position:"Query" name:"StartTime"`
	OwnerId       string `position:"Query" name:"OwnerId"`
	StreamName    string `position:"Query" name:"StreamName"`
	AccessKeyId   string `position:"Query" name:"AccessKeyId"`
}

type DescribeLiveStreamFrameLossRatioResponse struct {
	*responses.BaseResponse
	RequestId           string `json:"RequestId"`
	FrameLossRatioInfos []struct {
		StreamUrl      string  `json:"StreamUrl"`
		FrameLossRatio float64 `json:"FrameLossRatio"`
		Time           string  `json:"Time"`
	} `json:"FrameLossRatioInfos"`
}

func CreateDescribeLiveStreamFrameLossRatioRequest() (request *DescribeLiveStreamFrameLossRatioRequest) {
	request = &DescribeLiveStreamFrameLossRatioRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Cdn", "2014-11-11", "DescribeLiveStreamFrameLossRatio", "", "")
	return
}

func CreateDescribeLiveStreamFrameLossRatioResponse() (response *DescribeLiveStreamFrameLossRatioResponse) {
	response = &DescribeLiveStreamFrameLossRatioResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
