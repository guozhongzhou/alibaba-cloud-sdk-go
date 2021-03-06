package iot

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

// QueryDeviceByDriver invokes the iot.QueryDeviceByDriver API synchronously
// api document: https://help.aliyun.com/api/iot/querydevicebydriver.html
func (client *Client) QueryDeviceByDriver(request *QueryDeviceByDriverRequest) (response *QueryDeviceByDriverResponse, err error) {
	response = CreateQueryDeviceByDriverResponse()
	err = client.DoAction(request, response)
	return
}

// QueryDeviceByDriverWithChan invokes the iot.QueryDeviceByDriver API asynchronously
// api document: https://help.aliyun.com/api/iot/querydevicebydriver.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) QueryDeviceByDriverWithChan(request *QueryDeviceByDriverRequest) (<-chan *QueryDeviceByDriverResponse, <-chan error) {
	responseChan := make(chan *QueryDeviceByDriverResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.QueryDeviceByDriver(request)
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

// QueryDeviceByDriverWithCallback invokes the iot.QueryDeviceByDriver API asynchronously
// api document: https://help.aliyun.com/api/iot/querydevicebydriver.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) QueryDeviceByDriverWithCallback(request *QueryDeviceByDriverRequest, callback func(response *QueryDeviceByDriverResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *QueryDeviceByDriverResponse
		var err error
		defer close(result)
		response, err = client.QueryDeviceByDriver(request)
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

// QueryDeviceByDriverRequest is the request struct for api QueryDeviceByDriver
type QueryDeviceByDriverRequest struct {
	*requests.RpcRequest
	CurrentPage   requests.Integer `position:"Query" name:"CurrentPage"`
	InstanceId    string           `position:"Query" name:"InstanceId"`
	DriverId      string           `position:"Query" name:"DriverId"`
	IotInstanceId string           `position:"Query" name:"IotInstanceId"`
	PageSize      requests.Integer `position:"Query" name:"PageSize"`
}

// QueryDeviceByDriverResponse is the response struct for api QueryDeviceByDriver
type QueryDeviceByDriverResponse struct {
	*responses.BaseResponse
	RequestId    string                    `json:"RequestId" xml:"RequestId"`
	Success      bool                      `json:"Success" xml:"Success"`
	Code         string                    `json:"Code" xml:"Code"`
	ErrorMessage string                    `json:"ErrorMessage" xml:"ErrorMessage"`
	Data         DataInQueryDeviceByDriver `json:"Data" xml:"Data"`
}

// CreateQueryDeviceByDriverRequest creates a request to invoke QueryDeviceByDriver API
func CreateQueryDeviceByDriverRequest() (request *QueryDeviceByDriverRequest) {
	request = &QueryDeviceByDriverRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Iot", "2018-01-20", "QueryDeviceByDriver", "Iot", "openAPI")
	return
}

// CreateQueryDeviceByDriverResponse creates a response to parse from QueryDeviceByDriver response
func CreateQueryDeviceByDriverResponse() (response *QueryDeviceByDriverResponse) {
	response = &QueryDeviceByDriverResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
