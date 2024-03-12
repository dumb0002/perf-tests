/*
Copyright 2020 The Kubernetes Authors.

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

package provider

import (
	clientset "k8s.io/client-go/kubernetes"
	prom "k8s.io/perf-tests/clusterloader2/pkg/prometheus/clients"
)

type KSProvider struct {
	features Features
}

func NewKSProvider(_ map[string]string) Provider {
	return &KSProvider{
		features: Features{
			SupportProbe:                        false,
			SupportSSHToMaster:                  false,
			SupportImagePreload:                 false,
			SupportEnablePrometheusServer:       false,
			SupportGrabMetricsFromKubelets:      false,
			SupportAccessAPIServerPprofEndpoint: false,
			SupportResourceUsageMetering:        false,
			ShouldScrapeKubeProxy:               false,
		},
	}
}

func (p *KSProvider) Name() string {
	return KSName
}

func (p *KSProvider) Features() *Features {
	return &p.features
}

func (p *KSProvider) GetComponentProtocolAndPort(componentName string) (string, int, error) {
	return getComponentProtocolAndPort(componentName)
}

func (p *KSProvider) GetConfig() Config {
	return Config{}
}

func (p *KSProvider) RunSSHCommand(cmd, host string) (string, string, int, error) {
	return "", "", 0, nil
}

func (p *KSProvider) Metadata(client clientset.Interface) (map[string]string, error) {
	return nil, nil
}

func (p *KSProvider) GetManagedPrometheusClient() (prom.Client, error) {
	return nil, ErrNoManagedPrometheus
}
