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

func (client *Client) ClearUserBlackList(request *ClearUserBlackListRequest) (response *ClearUserBlackListResponse, err error) {
	response = CreateClearUserBlackListResponse()
	err = client.DoAction(request, response)
	return
}

func (client *Client) ClearUserBlackListWithChan(request *ClearUserBlackListRequest) (<-chan *ClearUserBlackListResponse, <-chan error) {
	responseChan := make(chan *ClearUserBlackListResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ClearUserBlackList(request)
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

func (client *Client) ClearUserBlackListWithCallback(request *ClearUserBlackListRequest, callback func(response *ClearUserBlackListResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ClearUserBlackListResponse
		var err error
		defer close(result)
		response, err = client.ClearUserBlackList(request)
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

type ClearUserBlackListRequest struct {
	*requests.RpcRequest
	SecurityToken string `position:"Query" name:"SecurityToken"`
	OwnerAccount  string `position:"Query" name:"OwnerAccount"`
	Action        string `position:"Query" name:"Action"`
	OwnerId       string `position:"Query" name:"OwnerId"`
	AccessKeyId   string `position:"Query" name:"AccessKeyId"`
}

type ClearUserBlackListResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId"`
}

func CreateClearUserBlackListRequest() (request *ClearUserBlackListRequest) {
	request = &ClearUserBlackListRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Cdn", "2014-11-11", "ClearUserBlackList", "", "")
	return
}

func CreateClearUserBlackListResponse() (response *ClearUserBlackListResponse) {
	response = &ClearUserBlackListResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
