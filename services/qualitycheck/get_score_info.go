package qualitycheck

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

// GetScoreInfo invokes the qualitycheck.GetScoreInfo API synchronously
// api document: https://help.aliyun.com/api/qualitycheck/getscoreinfo.html
func (client *Client) GetScoreInfo(request *GetScoreInfoRequest) (response *GetScoreInfoResponse, err error) {
	response = CreateGetScoreInfoResponse()
	err = client.DoAction(request, response)
	return
}

// GetScoreInfoWithChan invokes the qualitycheck.GetScoreInfo API asynchronously
// api document: https://help.aliyun.com/api/qualitycheck/getscoreinfo.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) GetScoreInfoWithChan(request *GetScoreInfoRequest) (<-chan *GetScoreInfoResponse, <-chan error) {
	responseChan := make(chan *GetScoreInfoResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.GetScoreInfo(request)
		if err != nil {
			errChan <- err
		} else {
			responseChan <- response
		}
	})
	if err != nil {
		errChan <- err
		close(responseChan)
		close(errChan)
	}
	return responseChan, errChan
}

// GetScoreInfoWithCallback invokes the qualitycheck.GetScoreInfo API asynchronously
// api document: https://help.aliyun.com/api/qualitycheck/getscoreinfo.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) GetScoreInfoWithCallback(request *GetScoreInfoRequest, callback func(response *GetScoreInfoResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *GetScoreInfoResponse
		var err error
		defer close(result)
		response, err = client.GetScoreInfo(request)
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

// GetScoreInfoRequest is the request struct for api GetScoreInfo
type GetScoreInfoRequest struct {
	*requests.RpcRequest
	ResourceOwnerId requests.Integer `position:"Query" name:"ResourceOwnerId"`
	JsonStr         string           `position:"Query" name:"JsonStr"`
}

// GetScoreInfoResponse is the response struct for api GetScoreInfo
type GetScoreInfoResponse struct {
	*responses.BaseResponse
	RequestId string             `json:"RequestId" xml:"RequestId"`
	Success   bool               `json:"Success" xml:"Success"`
	Code      string             `json:"Code" xml:"Code"`
	Message   string             `json:"Message" xml:"Message"`
	Data      DataInGetScoreInfo `json:"Data" xml:"Data"`
}

// CreateGetScoreInfoRequest creates a request to invoke GetScoreInfo API
func CreateGetScoreInfoRequest() (request *GetScoreInfoRequest) {
	request = &GetScoreInfoRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Qualitycheck", "2019-01-15", "GetScoreInfo", "", "")
	return
}

// CreateGetScoreInfoResponse creates a response to parse from GetScoreInfo response
func CreateGetScoreInfoResponse() (response *GetScoreInfoResponse) {
	response = &GetScoreInfoResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
