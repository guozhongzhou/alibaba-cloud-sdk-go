package drds

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

// RemoveDrdsInstance invokes the drds.RemoveDrdsInstance API synchronously
// api document: https://help.aliyun.com/api/drds/removedrdsinstance.html
func (client *Client) RemoveDrdsInstance(request *RemoveDrdsInstanceRequest) (response *RemoveDrdsInstanceResponse, err error) {
	response = CreateRemoveDrdsInstanceResponse()
	err = client.DoAction(request, response)
	return
}

// RemoveDrdsInstanceWithChan invokes the drds.RemoveDrdsInstance API asynchronously
// api document: https://help.aliyun.com/api/drds/removedrdsinstance.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) RemoveDrdsInstanceWithChan(request *RemoveDrdsInstanceRequest) (<-chan *RemoveDrdsInstanceResponse, <-chan error) {
	responseChan := make(chan *RemoveDrdsInstanceResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.RemoveDrdsInstance(request)
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

// RemoveDrdsInstanceWithCallback invokes the drds.RemoveDrdsInstance API asynchronously
// api document: https://help.aliyun.com/api/drds/removedrdsinstance.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) RemoveDrdsInstanceWithCallback(request *RemoveDrdsInstanceRequest, callback func(response *RemoveDrdsInstanceResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *RemoveDrdsInstanceResponse
		var err error
		defer close(result)
		response, err = client.RemoveDrdsInstance(request)
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

// RemoveDrdsInstanceRequest is the request struct for api RemoveDrdsInstance
type RemoveDrdsInstanceRequest struct {
	*requests.RpcRequest
	DrdsInstanceId string `position:"Query" name:"DrdsInstanceId"`
}

// RemoveDrdsInstanceResponse is the response struct for api RemoveDrdsInstance
type RemoveDrdsInstanceResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Success   bool   `json:"Success" xml:"Success"`
}

// CreateRemoveDrdsInstanceRequest creates a request to invoke RemoveDrdsInstance API
func CreateRemoveDrdsInstanceRequest() (request *RemoveDrdsInstanceRequest) {
	request = &RemoveDrdsInstanceRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Drds", "2017-10-16", "RemoveDrdsInstance", "drds", "openAPI")
	return
}

// CreateRemoveDrdsInstanceResponse creates a response to parse from RemoveDrdsInstance response
func CreateRemoveDrdsInstanceResponse() (response *RemoveDrdsInstanceResponse) {
	response = &RemoveDrdsInstanceResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
