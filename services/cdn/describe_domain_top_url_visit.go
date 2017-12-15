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

func (client *Client) DescribeDomainTopUrlVisit(request *DescribeDomainTopUrlVisitRequest) (response *DescribeDomainTopUrlVisitResponse, err error) {
	response = CreateDescribeDomainTopUrlVisitResponse()
	err = client.DoAction(request, response)
	return
}

func (client *Client) DescribeDomainTopUrlVisitWithChan(request *DescribeDomainTopUrlVisitRequest) (<-chan *DescribeDomainTopUrlVisitResponse, <-chan error) {
	responseChan := make(chan *DescribeDomainTopUrlVisitResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeDomainTopUrlVisit(request)
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

func (client *Client) DescribeDomainTopUrlVisitWithCallback(request *DescribeDomainTopUrlVisitRequest, callback func(response *DescribeDomainTopUrlVisitResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeDomainTopUrlVisitResponse
		var err error
		defer close(result)
		response, err = client.DescribeDomainTopUrlVisit(request)
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

type DescribeDomainTopUrlVisitRequest struct {
	*requests.RpcRequest
	SecurityToken string `position:"Query" name:"SecurityToken"`
	DomainName    string `position:"Query" name:"DomainName"`
	Action        string `position:"Query" name:"Action"`
	SortBy        string `position:"Query" name:"SortBy"`
	StartTime     string `position:"Query" name:"StartTime"`
	OwnerId       string `position:"Query" name:"OwnerId"`
	AccessKeyId   string `position:"Query" name:"AccessKeyId"`
}

type DescribeDomainTopUrlVisitResponse struct {
	*responses.BaseResponse
	RequestId  string `json:"RequestId"`
	DomainName string `json:"DomainName"`
	StartTime  string `json:"StartTime"`
	AllUrlList []struct {
		UrlDetail       string  `json:"UrlDetail"`
		VisitData       string  `json:"VisitData"`
		VisitProportion float64 `json:"VisitProportion"`
		Flow            string  `json:"Flow"`
		FlowProportion  float64 `json:"FlowProportion"`
	} `json:"AllUrlList"`
	Url200List []struct {
		UrlDetail       string  `json:"UrlDetail"`
		VisitData       string  `json:"VisitData"`
		VisitProportion float64 `json:"VisitProportion"`
		Flow            string  `json:"Flow"`
		FlowProportion  float64 `json:"FlowProportion"`
	} `json:"Url200List"`
	Url300List []struct {
		UrlDetail       string  `json:"UrlDetail"`
		VisitData       string  `json:"VisitData"`
		VisitProportion float64 `json:"VisitProportion"`
		Flow            string  `json:"Flow"`
		FlowProportion  float64 `json:"FlowProportion"`
	} `json:"Url300List"`
	Url400List []struct {
		UrlDetail       string  `json:"UrlDetail"`
		VisitData       string  `json:"VisitData"`
		VisitProportion float64 `json:"VisitProportion"`
		Flow            string  `json:"Flow"`
		FlowProportion  float64 `json:"FlowProportion"`
	} `json:"Url400List"`
	Url500List []struct {
		UrlDetail       string  `json:"UrlDetail"`
		VisitData       string  `json:"VisitData"`
		VisitProportion float64 `json:"VisitProportion"`
		Flow            string  `json:"Flow"`
		FlowProportion  float64 `json:"FlowProportion"`
	} `json:"Url500List"`
}

func CreateDescribeDomainTopUrlVisitRequest() (request *DescribeDomainTopUrlVisitRequest) {
	request = &DescribeDomainTopUrlVisitRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Cdn", "2014-11-11", "DescribeDomainTopUrlVisit", "", "")
	return
}

func CreateDescribeDomainTopUrlVisitResponse() (response *DescribeDomainTopUrlVisitResponse) {
	response = &DescribeDomainTopUrlVisitResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
