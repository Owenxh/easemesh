/*
Copyright (c) 2021, MegaEase
All rights reserved.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

	http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
// code generated by github.com/megaease/easemeshctl/cmd/generator, DO NOT EDIT.
package meshclient

import (
	"context"
	"encoding/json"
	"fmt"
	v1alpha1 "github.com/megaease/easemesh-api/v1alpha1"
	resource "github.com/megaease/easemeshctl/cmd/client/resource"
	client "github.com/megaease/easemeshctl/cmd/common/client"
	errors "github.com/pkg/errors"
	"net/http"
)

type serviceCanaryGetter struct {
	client *meshClient
}
type serviceCanaryInterface struct {
	client *meshClient
}

func (s *serviceCanaryGetter) ServiceCanary() ServiceCanaryInterface {
	return &serviceCanaryInterface{client: s.client}
}
func (s *serviceCanaryInterface) Get(args0 context.Context, args1 string) (*resource.ServiceCanary, error) {
	url := fmt.Sprintf("http://"+s.client.server+apiURL+"/mesh/"+"servicecanaries/%s", args1)
	r0, err := client.NewHTTPJSON().GetByContext(args0, url, nil, nil).HandleResponse(func(buff []byte, statusCode int) (interface{}, error) {
		if statusCode == http.StatusNotFound {
			return nil, errors.Wrapf(NotFoundError, "get ServiceCanary %s", args1)
		}
		if statusCode >= 300 {
			return nil, errors.Errorf("call %s failed, return status code %d text %+v", url, statusCode, string(buff))
		}
		ServiceCanary := &v1alpha1.ServiceCanary{}
		err := json.Unmarshal(buff, ServiceCanary)
		if err != nil {
			return nil, errors.Wrapf(err, "unmarshal data to v1alpha1.ServiceCanary")
		}
		return resource.ToServiceCanary(ServiceCanary), nil
	})
	if err != nil {
		return nil, err
	}
	return r0.(*resource.ServiceCanary), nil
}
func (s *serviceCanaryInterface) Patch(args0 context.Context, args1 *resource.ServiceCanary) error {
	url := fmt.Sprintf("http://"+s.client.server+apiURL+"/mesh/"+"servicecanaries/%s", args1.Name())
	object := args1.ToV1Alpha1()
	_, err := client.NewHTTPJSON().PutByContext(args0, url, object, nil).HandleResponse(func(b []byte, statusCode int) (interface{}, error) {
		if statusCode == http.StatusNotFound {
			return nil, errors.Wrapf(NotFoundError, "patch ServiceCanary %s", args1.Name())
		}
		if statusCode < 300 && statusCode >= 200 {
			return nil, nil
		}
		return nil, errors.Errorf("call PUT %s failed, return statuscode %d text %+v", url, statusCode, string(b))
	})
	return err
}
func (s *serviceCanaryInterface) Create(args0 context.Context, args1 *resource.ServiceCanary) error {
	url := "http://" + s.client.server + apiURL + "/mesh/servicecanaries"
	object := args1.ToV1Alpha1()
	_, err := client.NewHTTPJSON().PostByContext(args0, url, object, nil).HandleResponse(func(b []byte, statusCode int) (interface{}, error) {
		if statusCode == http.StatusConflict {
			return nil, errors.Wrapf(ConflictError, "create ServiceCanary %s", args1.Name())
		}
		if statusCode < 300 && statusCode >= 200 {
			return nil, nil
		}
		return nil, errors.Errorf("call Post %s failed, return statuscode %d text %+v", url, statusCode, string(b))
	})
	return err
}
func (s *serviceCanaryInterface) Delete(args0 context.Context, args1 string) error {
	url := fmt.Sprintf("http://"+s.client.server+apiURL+"/mesh/"+"servicecanaries/%s", args1)
	_, err := client.NewHTTPJSON().DeleteByContext(args0, url, nil, nil).HandleResponse(func(b []byte, statusCode int) (interface{}, error) {
		if statusCode == http.StatusNotFound {
			return nil, errors.Wrapf(NotFoundError, "Delete ServiceCanary %s", args1)
		}
		if statusCode < 300 && statusCode >= 200 {
			return nil, nil
		}
		return nil, errors.Errorf("call Delete %s failed, return statuscode %d text %+v", url, statusCode, string(b))
	})
	return err
}
func (s *serviceCanaryInterface) List(args0 context.Context) ([]*resource.ServiceCanary, error) {
	url := "http://" + s.client.server + apiURL + "/mesh/servicecanaries"
	result, err := client.NewHTTPJSON().GetByContext(args0, url, nil, nil).HandleResponse(func(b []byte, statusCode int) (interface{}, error) {
		if statusCode == http.StatusNotFound {
			return nil, errors.Wrapf(NotFoundError, "list service")
		}
		if statusCode >= 300 && statusCode < 200 {
			return nil, errors.Errorf("call GET %s failed, return statuscode %d text %+v", url, statusCode, b)
		}
		serviceCanary := []v1alpha1.ServiceCanary{}
		err := json.Unmarshal(b, &serviceCanary)
		if err != nil {
			return nil, errors.Wrapf(err, "unmarshal data to v1alpha1.")
		}
		results := []*resource.ServiceCanary{}
		for _, item := range serviceCanary {
			copy := item
			results = append(results, resource.ToServiceCanary(&copy))
		}
		return results, nil
	})
	if err != nil {
		return nil, err
	}
	return result.([]*resource.ServiceCanary), nil
}