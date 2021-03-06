// +build !ignore_autogenerated

/*
Copyright 2018 The Kubernetes Authors.

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

// This file was autogenerated by conversion-gen. Do not edit it manually!

package v1

import (
	sho "github.com/sanjid133/ksd/apis/sho"
	conversion "k8s.io/apimachinery/pkg/conversion"
	runtime "k8s.io/apimachinery/pkg/runtime"
	unsafe "unsafe"
)

func init() {
	localSchemeBuilder.Register(RegisterConversions)
}

// RegisterConversions adds conversion functions to the given scheme.
// Public to allow building arbitrary schemes.
func RegisterConversions(scheme *runtime.Scheme) error {
	return scheme.AddGeneratedConversionFuncs(
		Convert_v1_InfoSpec_To_sho_InfoSpec,
		Convert_sho_InfoSpec_To_v1_InfoSpec,
		Convert_v1_Ksd_To_sho_Ksd,
		Convert_sho_Ksd_To_v1_Ksd,
		Convert_v1_KsdList_To_sho_KsdList,
		Convert_sho_KsdList_To_v1_KsdList,
		Convert_v1_KsdSpec_To_sho_KsdSpec,
		Convert_sho_KsdSpec_To_v1_KsdSpec,
	)
}

func autoConvert_v1_InfoSpec_To_sho_InfoSpec(in *InfoSpec, out *sho.InfoSpec, s conversion.Scope) error {
	out.SecretName = in.SecretName
	out.Data = *(*map[string]string)(unsafe.Pointer(&in.Data))
	return nil
}

// Convert_v1_InfoSpec_To_sho_InfoSpec is an autogenerated conversion function.
func Convert_v1_InfoSpec_To_sho_InfoSpec(in *InfoSpec, out *sho.InfoSpec, s conversion.Scope) error {
	return autoConvert_v1_InfoSpec_To_sho_InfoSpec(in, out, s)
}

func autoConvert_sho_InfoSpec_To_v1_InfoSpec(in *sho.InfoSpec, out *InfoSpec, s conversion.Scope) error {
	out.SecretName = in.SecretName
	out.Data = *(*map[string]string)(unsafe.Pointer(&in.Data))
	return nil
}

// Convert_sho_InfoSpec_To_v1_InfoSpec is an autogenerated conversion function.
func Convert_sho_InfoSpec_To_v1_InfoSpec(in *sho.InfoSpec, out *InfoSpec, s conversion.Scope) error {
	return autoConvert_sho_InfoSpec_To_v1_InfoSpec(in, out, s)
}

func autoConvert_v1_Ksd_To_sho_Ksd(in *Ksd, out *sho.Ksd, s conversion.Scope) error {
	out.ObjectMeta = in.ObjectMeta
	if err := Convert_v1_KsdSpec_To_sho_KsdSpec(&in.Spec, &out.Spec, s); err != nil {
		return err
	}
	return nil
}

// Convert_v1_Ksd_To_sho_Ksd is an autogenerated conversion function.
func Convert_v1_Ksd_To_sho_Ksd(in *Ksd, out *sho.Ksd, s conversion.Scope) error {
	return autoConvert_v1_Ksd_To_sho_Ksd(in, out, s)
}

func autoConvert_sho_Ksd_To_v1_Ksd(in *sho.Ksd, out *Ksd, s conversion.Scope) error {
	out.ObjectMeta = in.ObjectMeta
	if err := Convert_sho_KsdSpec_To_v1_KsdSpec(&in.Spec, &out.Spec, s); err != nil {
		return err
	}
	return nil
}

// Convert_sho_Ksd_To_v1_Ksd is an autogenerated conversion function.
func Convert_sho_Ksd_To_v1_Ksd(in *sho.Ksd, out *Ksd, s conversion.Scope) error {
	return autoConvert_sho_Ksd_To_v1_Ksd(in, out, s)
}

func autoConvert_v1_KsdList_To_sho_KsdList(in *KsdList, out *sho.KsdList, s conversion.Scope) error {
	out.ListMeta = in.ListMeta
	out.Items = *(*[]sho.Ksd)(unsafe.Pointer(&in.Items))
	return nil
}

// Convert_v1_KsdList_To_sho_KsdList is an autogenerated conversion function.
func Convert_v1_KsdList_To_sho_KsdList(in *KsdList, out *sho.KsdList, s conversion.Scope) error {
	return autoConvert_v1_KsdList_To_sho_KsdList(in, out, s)
}

func autoConvert_sho_KsdList_To_v1_KsdList(in *sho.KsdList, out *KsdList, s conversion.Scope) error {
	out.ListMeta = in.ListMeta
	out.Items = *(*[]Ksd)(unsafe.Pointer(&in.Items))
	return nil
}

// Convert_sho_KsdList_To_v1_KsdList is an autogenerated conversion function.
func Convert_sho_KsdList_To_v1_KsdList(in *sho.KsdList, out *KsdList, s conversion.Scope) error {
	return autoConvert_sho_KsdList_To_v1_KsdList(in, out, s)
}

func autoConvert_v1_KsdSpec_To_sho_KsdSpec(in *KsdSpec, out *sho.KsdSpec, s conversion.Scope) error {
	out.SecretType = in.SecretType
	out.Info = *(*[]sho.InfoSpec)(unsafe.Pointer(&in.Info))
	return nil
}

// Convert_v1_KsdSpec_To_sho_KsdSpec is an autogenerated conversion function.
func Convert_v1_KsdSpec_To_sho_KsdSpec(in *KsdSpec, out *sho.KsdSpec, s conversion.Scope) error {
	return autoConvert_v1_KsdSpec_To_sho_KsdSpec(in, out, s)
}

func autoConvert_sho_KsdSpec_To_v1_KsdSpec(in *sho.KsdSpec, out *KsdSpec, s conversion.Scope) error {
	out.SecretType = in.SecretType
	out.Info = *(*[]InfoSpec)(unsafe.Pointer(&in.Info))
	return nil
}

// Convert_sho_KsdSpec_To_v1_KsdSpec is an autogenerated conversion function.
func Convert_sho_KsdSpec_To_v1_KsdSpec(in *sho.KsdSpec, out *KsdSpec, s conversion.Scope) error {
	return autoConvert_sho_KsdSpec_To_v1_KsdSpec(in, out, s)
}
