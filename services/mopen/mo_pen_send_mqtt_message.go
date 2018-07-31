package mopen

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

// MoPenSendMqttMessage invokes the mopen.MoPenSendMqttMessage API synchronously
// api document: https://help.aliyun.com/api/mopen/mopensendmqttmessage.html
func (client *Client) MoPenSendMqttMessage(request *MoPenSendMqttMessageRequest) (response *MoPenSendMqttMessageResponse, err error) {
	response = CreateMoPenSendMqttMessageResponse()
	err = client.DoAction(request, response)
	return
}

// MoPenSendMqttMessageWithChan invokes the mopen.MoPenSendMqttMessage API asynchronously
// api document: https://help.aliyun.com/api/mopen/mopensendmqttmessage.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) MoPenSendMqttMessageWithChan(request *MoPenSendMqttMessageRequest) (<-chan *MoPenSendMqttMessageResponse, <-chan error) {
	responseChan := make(chan *MoPenSendMqttMessageResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.MoPenSendMqttMessage(request)
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

// MoPenSendMqttMessageWithCallback invokes the mopen.MoPenSendMqttMessage API asynchronously
// api document: https://help.aliyun.com/api/mopen/mopensendmqttmessage.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) MoPenSendMqttMessageWithCallback(request *MoPenSendMqttMessageRequest, callback func(response *MoPenSendMqttMessageResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *MoPenSendMqttMessageResponse
		var err error
		defer close(result)
		response, err = client.MoPenSendMqttMessage(request)
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

// MoPenSendMqttMessageRequest is the request struct for api MoPenSendMqttMessage
type MoPenSendMqttMessageRequest struct {
	*requests.RpcRequest
	Payload    string `position:"Body" name:"Payload"`
	DeviceName string `position:"Body" name:"DeviceName"`
}

// MoPenSendMqttMessageResponse is the response struct for api MoPenSendMqttMessage
type MoPenSendMqttMessageResponse struct {
	*responses.BaseResponse
	RequestId   string `json:"RequestId" xml:"RequestId"`
	Code        string `json:"Code" xml:"Code"`
	Message     string `json:"Message" xml:"Message"`
	Success     bool   `json:"Success" xml:"Success"`
	Description string `json:"Description" xml:"Description"`
}

// CreateMoPenSendMqttMessageRequest creates a request to invoke MoPenSendMqttMessage API
func CreateMoPenSendMqttMessageRequest() (request *MoPenSendMqttMessageRequest) {
	request = &MoPenSendMqttMessageRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("MoPen", "2018-02-11", "MoPenSendMqttMessage", "", "")
	return
}

// CreateMoPenSendMqttMessageResponse creates a response to parse from MoPenSendMqttMessage response
func CreateMoPenSendMqttMessageResponse() (response *MoPenSendMqttMessageResponse) {
	response = &MoPenSendMqttMessageResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
