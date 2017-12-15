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

func (client *Client) AddLiveAppSnapshotConfig(request *AddLiveAppSnapshotConfigRequest) (response *AddLiveAppSnapshotConfigResponse, err error) {
	response = CreateAddLiveAppSnapshotConfigResponse()
	err = client.DoAction(request, response)
	return
}

func (client *Client) AddLiveAppSnapshotConfigWithChan(request *AddLiveAppSnapshotConfigRequest) (<-chan *AddLiveAppSnapshotConfigResponse, <-chan error) {
	responseChan := make(chan *AddLiveAppSnapshotConfigResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.AddLiveAppSnapshotConfig(request)
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

func (client *Client) AddLiveAppSnapshotConfigWithCallback(request *AddLiveAppSnapshotConfigRequest, callback func(response *AddLiveAppSnapshotConfigResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *AddLiveAppSnapshotConfigResponse
		var err error
		defer close(result)
		response, err = client.AddLiveAppSnapshotConfig(request)
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

type AddLiveAppSnapshotConfigRequest struct {
	*requests.RpcRequest
	TimeInterval       string `position:"Query" name:"TimeInterval"`
	OssBucket          string `position:"Query" name:"OssBucket"`
	AppName            string `position:"Query" name:"AppName"`
	SecurityToken      string `position:"Query" name:"SecurityToken"`
	DomainName         string `position:"Query" name:"DomainName"`
	OssEndpoint        string `position:"Query" name:"OssEndpoint"`
	Action             string `position:"Query" name:"Action"`
	SequenceOssObject  string `position:"Query" name:"SequenceOssObject"`
	OverwriteOssObject string `position:"Query" name:"OverwriteOssObject"`
	OwnerId            string `position:"Query" name:"OwnerId"`
	AccessKeyId        string `position:"Query" name:"AccessKeyId"`
}

type AddLiveAppSnapshotConfigResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId"`
}

func CreateAddLiveAppSnapshotConfigRequest() (request *AddLiveAppSnapshotConfigRequest) {
	request = &AddLiveAppSnapshotConfigRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Cdn", "2014-11-11", "AddLiveAppSnapshotConfig", "", "")
	return
}

func CreateAddLiveAppSnapshotConfigResponse() (response *AddLiveAppSnapshotConfigResponse) {
	response = &AddLiveAppSnapshotConfigResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
