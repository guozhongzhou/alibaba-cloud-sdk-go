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

func (client *Client) DescribeLiveStreamLimitInfo(request *DescribeLiveStreamLimitInfoRequest) (response *DescribeLiveStreamLimitInfoResponse, err error) {
	response = CreateDescribeLiveStreamLimitInfoResponse()
	err = client.DoAction(request, response)
	return
}

func (client *Client) DescribeLiveStreamLimitInfoWithChan(request *DescribeLiveStreamLimitInfoRequest) (<-chan *DescribeLiveStreamLimitInfoResponse, <-chan error) {
	responseChan := make(chan *DescribeLiveStreamLimitInfoResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeLiveStreamLimitInfo(request)
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

func (client *Client) DescribeLiveStreamLimitInfoWithCallback(request *DescribeLiveStreamLimitInfoRequest, callback func(response *DescribeLiveStreamLimitInfoResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeLiveStreamLimitInfoResponse
		var err error
		defer close(result)
		response, err = client.DescribeLiveStreamLimitInfo(request)
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

type DescribeLiveStreamLimitInfoRequest struct {
	*requests.RpcRequest
	SecurityToken string `position:"Query" name:"SecurityToken"`
	Action        string `position:"Query" name:"Action"`
	OwnerId       string `position:"Query" name:"OwnerId"`
	AccessKeyId   string `position:"Query" name:"AccessKeyId"`
	LimitDomain   string `position:"Query" name:"LimitDomain"`
}

type DescribeLiveStreamLimitInfoResponse struct {
	*responses.BaseResponse
	RequestId      string `json:"RequestId"`
	UserLimitLists []struct {
		LimitDomain       string `json:"LimitDomain"`
		LimitNum          string `json:"LimitNum"`
		LimitTranscodeNum string `json:"LimitTranscodeNum"`
	} `json:"UserLimitLists"`
}

func CreateDescribeLiveStreamLimitInfoRequest() (request *DescribeLiveStreamLimitInfoRequest) {
	request = &DescribeLiveStreamLimitInfoRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Cdn", "2014-11-11", "DescribeLiveStreamLimitInfo", "", "")
	return
}

func CreateDescribeLiveStreamLimitInfoResponse() (response *DescribeLiveStreamLimitInfoResponse) {
	response = &DescribeLiveStreamLimitInfoResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
