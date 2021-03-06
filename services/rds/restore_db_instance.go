package rds

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

// RestoreDBInstance invokes the rds.RestoreDBInstance API synchronously
// api document: https://help.aliyun.com/api/rds/restoredbinstance.html
func (client *Client) RestoreDBInstance(request *RestoreDBInstanceRequest) (response *RestoreDBInstanceResponse, err error) {
	response = CreateRestoreDBInstanceResponse()
	err = client.DoAction(request, response)
	return
}

// RestoreDBInstanceWithChan invokes the rds.RestoreDBInstance API asynchronously
// api document: https://help.aliyun.com/api/rds/restoredbinstance.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) RestoreDBInstanceWithChan(request *RestoreDBInstanceRequest) (<-chan *RestoreDBInstanceResponse, <-chan error) {
	responseChan := make(chan *RestoreDBInstanceResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.RestoreDBInstance(request)
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

// RestoreDBInstanceWithCallback invokes the rds.RestoreDBInstance API asynchronously
// api document: https://help.aliyun.com/api/rds/restoredbinstance.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) RestoreDBInstanceWithCallback(request *RestoreDBInstanceRequest, callback func(response *RestoreDBInstanceResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *RestoreDBInstanceResponse
		var err error
		defer close(result)
		response, err = client.RestoreDBInstance(request)
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

// RestoreDBInstanceRequest is the request struct for api RestoreDBInstance
type RestoreDBInstanceRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	ClientToken          string           `position:"Query" name:"ClientToken"`
	DBInstanceId         string           `position:"Query" name:"DBInstanceId"`
	RestoreTime          string           `position:"Query" name:"RestoreTime"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	BackupId             string           `position:"Query" name:"BackupId"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
}

// RestoreDBInstanceResponse is the response struct for api RestoreDBInstance
type RestoreDBInstanceResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateRestoreDBInstanceRequest creates a request to invoke RestoreDBInstance API
func CreateRestoreDBInstanceRequest() (request *RestoreDBInstanceRequest) {
	request = &RestoreDBInstanceRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Rds", "2014-08-15", "RestoreDBInstance", "rds", "openAPI")
	return
}

// CreateRestoreDBInstanceResponse creates a response to parse from RestoreDBInstance response
func CreateRestoreDBInstanceResponse() (response *RestoreDBInstanceResponse) {
	response = &RestoreDBInstanceResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
