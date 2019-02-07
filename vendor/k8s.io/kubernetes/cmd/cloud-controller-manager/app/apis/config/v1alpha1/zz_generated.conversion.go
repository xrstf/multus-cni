// +build !ignore_autogenerated

/*
Copyright The Kubernetes Authors.

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

// Code generated by conversion-gen. DO NOT EDIT.

package v1alpha1

import (
	conversion "k8s.io/apimachinery/pkg/conversion"
	runtime "k8s.io/apimachinery/pkg/runtime"
	config "k8s.io/kubernetes/cmd/cloud-controller-manager/app/apis/config"
	configv1alpha1 "k8s.io/kubernetes/pkg/controller/apis/config/v1alpha1"
)

func init() {
	localSchemeBuilder.Register(RegisterConversions)
}

// RegisterConversions adds conversion functions to the given scheme.
// Public to allow building arbitrary schemes.
func RegisterConversions(s *runtime.Scheme) error {
	if err := s.AddGeneratedConversionFunc((*CloudControllerManagerConfiguration)(nil), (*config.CloudControllerManagerConfiguration)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha1_CloudControllerManagerConfiguration_To_config_CloudControllerManagerConfiguration(a.(*CloudControllerManagerConfiguration), b.(*config.CloudControllerManagerConfiguration), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*config.CloudControllerManagerConfiguration)(nil), (*CloudControllerManagerConfiguration)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_config_CloudControllerManagerConfiguration_To_v1alpha1_CloudControllerManagerConfiguration(a.(*config.CloudControllerManagerConfiguration), b.(*CloudControllerManagerConfiguration), scope)
	}); err != nil {
		return err
	}
	return nil
}

func autoConvert_v1alpha1_CloudControllerManagerConfiguration_To_config_CloudControllerManagerConfiguration(in *CloudControllerManagerConfiguration, out *config.CloudControllerManagerConfiguration, s conversion.Scope) error {
	if err := configv1alpha1.Convert_v1alpha1_GenericControllerManagerConfiguration_To_config_GenericControllerManagerConfiguration(&in.Generic, &out.Generic, s); err != nil {
		return err
	}
	if err := configv1alpha1.Convert_v1alpha1_KubeCloudSharedConfiguration_To_config_KubeCloudSharedConfiguration(&in.KubeCloudShared, &out.KubeCloudShared, s); err != nil {
		return err
	}
	if err := configv1alpha1.Convert_v1alpha1_ServiceControllerConfiguration_To_config_ServiceControllerConfiguration(&in.ServiceController, &out.ServiceController, s); err != nil {
		return err
	}
	out.NodeStatusUpdateFrequency = in.NodeStatusUpdateFrequency
	return nil
}

// Convert_v1alpha1_CloudControllerManagerConfiguration_To_config_CloudControllerManagerConfiguration is an autogenerated conversion function.
func Convert_v1alpha1_CloudControllerManagerConfiguration_To_config_CloudControllerManagerConfiguration(in *CloudControllerManagerConfiguration, out *config.CloudControllerManagerConfiguration, s conversion.Scope) error {
	return autoConvert_v1alpha1_CloudControllerManagerConfiguration_To_config_CloudControllerManagerConfiguration(in, out, s)
}

func autoConvert_config_CloudControllerManagerConfiguration_To_v1alpha1_CloudControllerManagerConfiguration(in *config.CloudControllerManagerConfiguration, out *CloudControllerManagerConfiguration, s conversion.Scope) error {
	if err := configv1alpha1.Convert_config_GenericControllerManagerConfiguration_To_v1alpha1_GenericControllerManagerConfiguration(&in.Generic, &out.Generic, s); err != nil {
		return err
	}
	if err := configv1alpha1.Convert_config_KubeCloudSharedConfiguration_To_v1alpha1_KubeCloudSharedConfiguration(&in.KubeCloudShared, &out.KubeCloudShared, s); err != nil {
		return err
	}
	if err := configv1alpha1.Convert_config_ServiceControllerConfiguration_To_v1alpha1_ServiceControllerConfiguration(&in.ServiceController, &out.ServiceController, s); err != nil {
		return err
	}
	out.NodeStatusUpdateFrequency = in.NodeStatusUpdateFrequency
	return nil
}

// Convert_config_CloudControllerManagerConfiguration_To_v1alpha1_CloudControllerManagerConfiguration is an autogenerated conversion function.
func Convert_config_CloudControllerManagerConfiguration_To_v1alpha1_CloudControllerManagerConfiguration(in *config.CloudControllerManagerConfiguration, out *CloudControllerManagerConfiguration, s conversion.Scope) error {
	return autoConvert_config_CloudControllerManagerConfiguration_To_v1alpha1_CloudControllerManagerConfiguration(in, out, s)
}