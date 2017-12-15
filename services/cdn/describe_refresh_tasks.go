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

func (client *Client) DescribeRefreshTasks(request *DescribeRefreshTasksRequest) (response *DescribeRefreshTasksResponse, err error) {
	response = CreateDescribeRefreshTasksResponse()
	err = client.DoAction(request, response)
	return
}

func (client *Client) DescribeRefreshTasksWithChan(request *DescribeRefreshTasksRequest) (<-chan *DescribeRefreshTasksResponse, <-chan error) {
	responseChan := make(chan *DescribeRefreshTasksResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeRefreshTasks(request)
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

func (client *Client) DescribeRefreshTasksWithCallback(request *DescribeRefreshTasksRequest, callback func(response *DescribeRefreshTasksResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeRefreshTasksResponse
		var err error
		defer close(result)
		response, err = client.DescribeRefreshTasks(request)
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

type DescribeRefreshTasksRequest struct {
	*requests.RpcRequest
	ObjectPath      string `position:"Query" name:"ObjectPath"`
	DomainName      string `position:"Query" name:"DomainName"`
	EndTime         string `position:"Query" name:"EndTime"`
	StartTime       string `position:"Query" name:"StartTime"`
	OwnerId         string `position:"Query" name:"OwnerId"`
	PageNumber      string `position:"Query" name:"PageNumber"`
	AccessKeyId     string `position:"Query" name:"AccessKeyId"`
	ResourceGroupId string `position:"Query" name:"ResourceGroupId"`
	SecurityToken   string `position:"Query" name:"SecurityToken"`
	PageSize        string `position:"Query" name:"PageSize"`
	Action          string `position:"Query" name:"Action"`
	ObjectType      string `position:"Query" name:"ObjectType"`
	TaskId          string `position:"Query" name:"TaskId"`
	Status          string `position:"Query" name:"Status"`
}

type DescribeRefreshTasksResponse struct {
	*responses.BaseResponse
	RequestId  string `json:"RequestId"`
	PageNumber int64  `json:"PageNumber"`
	PageSize   int64  `json:"PageSize"`
	TotalCount int64  `json:"TotalCount"`
	Tasks      []struct {
		TaskId       string `json:"TaskId"`
		ObjectPath   string `json:"ObjectPath"`
		Process      string `json:"Process"`
		Status       string `json:"Status"`
		CreationTime string `json:"CreationTime"`
		Description  string `json:"Description"`
		ObjectType   string `json:"ObjectType"`
	} `json:"Tasks"`
}

func CreateDescribeRefreshTasksRequest() (request *DescribeRefreshTasksRequest) {
	request = &DescribeRefreshTasksRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Cdn", "2014-11-11", "DescribeRefreshTasks", "", "")
	return
}

func CreateDescribeRefreshTasksResponse() (response *DescribeRefreshTasksResponse) {
	response = &DescribeRefreshTasksResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
