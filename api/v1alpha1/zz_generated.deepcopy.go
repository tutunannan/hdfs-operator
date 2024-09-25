//go:build !ignore_autogenerated

/*
Copyright 2024.

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

// Code generated by controller-gen. DO NOT EDIT.

package v1alpha1

import (
	commonsv1alpha1 "github.com/zncdatadev/operator-go/pkg/apis/commons/v1alpha1"
	"k8s.io/api/core/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AuthenticationSpec) DeepCopyInto(out *AuthenticationSpec) {
	*out = *in
	if in.Oidc != nil {
		in, out := &in.Oidc, &out.Oidc
		*out = new(OidcSpec)
		(*in).DeepCopyInto(*out)
	}
	if in.Tls != nil {
		in, out := &in.Tls, &out.Tls
		*out = new(TlsSpec)
		**out = **in
	}
	if in.Kerberos != nil {
		in, out := &in.Kerberos, &out.Kerberos
		*out = new(KerberosSpec)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AuthenticationSpec.
func (in *AuthenticationSpec) DeepCopy() *AuthenticationSpec {
	if in == nil {
		return nil
	}
	out := new(AuthenticationSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterConfigSpec) DeepCopyInto(out *ClusterConfigSpec) {
	*out = *in
	if in.Service != nil {
		in, out := &in.Service, &out.Service
		*out = new(ServiceSpec)
		(*in).DeepCopyInto(*out)
	}
	if in.Authentication != nil {
		in, out := &in.Authentication, &out.Authentication
		*out = new(AuthenticationSpec)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterConfigSpec.
func (in *ClusterConfigSpec) DeepCopy() *ClusterConfigSpec {
	if in == nil {
		return nil
	}
	out := new(ClusterConfigSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ConfigOverridesSpec) DeepCopyInto(out *ConfigOverridesSpec) {
	*out = *in
	if in.CoreSite != nil {
		in, out := &in.CoreSite, &out.CoreSite
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.HdfsSite != nil {
		in, out := &in.HdfsSite, &out.HdfsSite
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.Log4j != nil {
		in, out := &in.Log4j, &out.Log4j
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.Security != nil {
		in, out := &in.Security, &out.Security
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.HadoopPolicy != nil {
		in, out := &in.HadoopPolicy, &out.HadoopPolicy
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.SslServer != nil {
		in, out := &in.SslServer, &out.SslServer
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.SslClient != nil {
		in, out := &in.SslClient, &out.SslClient
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ConfigOverridesSpec.
func (in *ConfigOverridesSpec) DeepCopy() *ConfigOverridesSpec {
	if in == nil {
		return nil
	}
	out := new(ConfigOverridesSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DataNodeConfigSpec) DeepCopyInto(out *DataNodeConfigSpec) {
	*out = *in
	if in.Resources != nil {
		in, out := &in.Resources, &out.Resources
		*out = new(commonsv1alpha1.ResourcesSpec)
		(*in).DeepCopyInto(*out)
	}
	if in.SecurityContext != nil {
		in, out := &in.SecurityContext, &out.SecurityContext
		*out = new(v1.PodSecurityContext)
		(*in).DeepCopyInto(*out)
	}
	if in.Affinity != nil {
		in, out := &in.Affinity, &out.Affinity
		*out = new(v1.Affinity)
		(*in).DeepCopyInto(*out)
	}
	if in.NodeSelector != nil {
		in, out := &in.NodeSelector, &out.NodeSelector
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.Tolerations != nil {
		in, out := &in.Tolerations, &out.Tolerations
		*out = make([]v1.Toleration, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.PodDisruptionBudget != nil {
		in, out := &in.PodDisruptionBudget, &out.PodDisruptionBudget
		*out = new(PodDisruptionBudgetSpec)
		**out = **in
	}
	if in.ExtraEnv != nil {
		in, out := &in.ExtraEnv, &out.ExtraEnv
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.ExtraSecret != nil {
		in, out := &in.ExtraSecret, &out.ExtraSecret
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.Logging != nil {
		in, out := &in.Logging, &out.Logging
		*out = new(DataNodeContainerLoggingSpec)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DataNodeConfigSpec.
func (in *DataNodeConfigSpec) DeepCopy() *DataNodeConfigSpec {
	if in == nil {
		return nil
	}
	out := new(DataNodeConfigSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DataNodeContainerLoggingSpec) DeepCopyInto(out *DataNodeContainerLoggingSpec) {
	*out = *in
	if in.DataNode != nil {
		in, out := &in.DataNode, &out.DataNode
		*out = new(LoggingConfigSpec)
		(*in).DeepCopyInto(*out)
	}
	if in.WaitNameNode != nil {
		in, out := &in.WaitNameNode, &out.WaitNameNode
		*out = new(LoggingConfigSpec)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DataNodeContainerLoggingSpec.
func (in *DataNodeContainerLoggingSpec) DeepCopy() *DataNodeContainerLoggingSpec {
	if in == nil {
		return nil
	}
	out := new(DataNodeContainerLoggingSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DataNodeRoleGroupSpec) DeepCopyInto(out *DataNodeRoleGroupSpec) {
	*out = *in
	if in.Config != nil {
		in, out := &in.Config, &out.Config
		*out = new(DataNodeConfigSpec)
		(*in).DeepCopyInto(*out)
	}
	if in.CommandArgsOverrides != nil {
		in, out := &in.CommandArgsOverrides, &out.CommandArgsOverrides
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.ConfigOverrides != nil {
		in, out := &in.ConfigOverrides, &out.ConfigOverrides
		*out = new(ConfigOverridesSpec)
		(*in).DeepCopyInto(*out)
	}
	if in.EnvOverrides != nil {
		in, out := &in.EnvOverrides, &out.EnvOverrides
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DataNodeRoleGroupSpec.
func (in *DataNodeRoleGroupSpec) DeepCopy() *DataNodeRoleGroupSpec {
	if in == nil {
		return nil
	}
	out := new(DataNodeRoleGroupSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DataNodeSpec) DeepCopyInto(out *DataNodeSpec) {
	*out = *in
	if in.Config != nil {
		in, out := &in.Config, &out.Config
		*out = new(DataNodeConfigSpec)
		(*in).DeepCopyInto(*out)
	}
	if in.RoleGroups != nil {
		in, out := &in.RoleGroups, &out.RoleGroups
		*out = make(map[string]*DataNodeRoleGroupSpec, len(*in))
		for key, val := range *in {
			var outVal *DataNodeRoleGroupSpec
			if val == nil {
				(*out)[key] = nil
			} else {
				inVal := (*in)[key]
				in, out := &inVal, &outVal
				*out = new(DataNodeRoleGroupSpec)
				(*in).DeepCopyInto(*out)
			}
			(*out)[key] = outVal
		}
	}
	if in.PodDisruptionBudget != nil {
		in, out := &in.PodDisruptionBudget, &out.PodDisruptionBudget
		*out = new(PodDisruptionBudgetSpec)
		**out = **in
	}
	if in.CommandArgsOverrides != nil {
		in, out := &in.CommandArgsOverrides, &out.CommandArgsOverrides
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.ConfigOverrides != nil {
		in, out := &in.ConfigOverrides, &out.ConfigOverrides
		*out = new(ConfigOverridesSpec)
		(*in).DeepCopyInto(*out)
	}
	if in.EnvOverrides != nil {
		in, out := &in.EnvOverrides, &out.EnvOverrides
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DataNodeSpec.
func (in *DataNodeSpec) DeepCopy() *DataNodeSpec {
	if in == nil {
		return nil
	}
	out := new(DataNodeSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HdfsCluster) DeepCopyInto(out *HdfsCluster) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HdfsCluster.
func (in *HdfsCluster) DeepCopy() *HdfsCluster {
	if in == nil {
		return nil
	}
	out := new(HdfsCluster)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *HdfsCluster) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HdfsClusterList) DeepCopyInto(out *HdfsClusterList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]HdfsCluster, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HdfsClusterList.
func (in *HdfsClusterList) DeepCopy() *HdfsClusterList {
	if in == nil {
		return nil
	}
	out := new(HdfsClusterList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *HdfsClusterList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HdfsClusterSpec) DeepCopyInto(out *HdfsClusterSpec) {
	*out = *in
	if in.Image != nil {
		in, out := &in.Image, &out.Image
		*out = new(ImageSpec)
		**out = **in
	}
	if in.ClusterConfigSpec != nil {
		in, out := &in.ClusterConfigSpec, &out.ClusterConfigSpec
		*out = new(ClusterConfigSpec)
		(*in).DeepCopyInto(*out)
	}
	if in.NameNode != nil {
		in, out := &in.NameNode, &out.NameNode
		*out = new(NameNodeSpec)
		(*in).DeepCopyInto(*out)
	}
	if in.DataNode != nil {
		in, out := &in.DataNode, &out.DataNode
		*out = new(DataNodeSpec)
		(*in).DeepCopyInto(*out)
	}
	if in.JournalNode != nil {
		in, out := &in.JournalNode, &out.JournalNode
		*out = new(JournalNodeSpec)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HdfsClusterSpec.
func (in *HdfsClusterSpec) DeepCopy() *HdfsClusterSpec {
	if in == nil {
		return nil
	}
	out := new(HdfsClusterSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ImageSpec) DeepCopyInto(out *ImageSpec) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ImageSpec.
func (in *ImageSpec) DeepCopy() *ImageSpec {
	if in == nil {
		return nil
	}
	out := new(ImageSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *JournalNodeConfigSpec) DeepCopyInto(out *JournalNodeConfigSpec) {
	*out = *in
	if in.Resources != nil {
		in, out := &in.Resources, &out.Resources
		*out = new(commonsv1alpha1.ResourcesSpec)
		(*in).DeepCopyInto(*out)
	}
	if in.SecurityContext != nil {
		in, out := &in.SecurityContext, &out.SecurityContext
		*out = new(v1.PodSecurityContext)
		(*in).DeepCopyInto(*out)
	}
	if in.Affinity != nil {
		in, out := &in.Affinity, &out.Affinity
		*out = new(v1.Affinity)
		(*in).DeepCopyInto(*out)
	}
	if in.NodeSelector != nil {
		in, out := &in.NodeSelector, &out.NodeSelector
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.Tolerations != nil {
		in, out := &in.Tolerations, &out.Tolerations
		*out = make([]v1.Toleration, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.PodDisruptionBudget != nil {
		in, out := &in.PodDisruptionBudget, &out.PodDisruptionBudget
		*out = new(PodDisruptionBudgetSpec)
		**out = **in
	}
	if in.ExtraEnv != nil {
		in, out := &in.ExtraEnv, &out.ExtraEnv
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.ExtraSecret != nil {
		in, out := &in.ExtraSecret, &out.ExtraSecret
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.Logging != nil {
		in, out := &in.Logging, &out.Logging
		*out = new(JournalNodeContainerLoggingSpec)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new JournalNodeConfigSpec.
func (in *JournalNodeConfigSpec) DeepCopy() *JournalNodeConfigSpec {
	if in == nil {
		return nil
	}
	out := new(JournalNodeConfigSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *JournalNodeContainerLoggingSpec) DeepCopyInto(out *JournalNodeContainerLoggingSpec) {
	*out = *in
	if in.JournalNode != nil {
		in, out := &in.JournalNode, &out.JournalNode
		*out = new(LoggingConfigSpec)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new JournalNodeContainerLoggingSpec.
func (in *JournalNodeContainerLoggingSpec) DeepCopy() *JournalNodeContainerLoggingSpec {
	if in == nil {
		return nil
	}
	out := new(JournalNodeContainerLoggingSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *JournalNodeRoleGroupSpec) DeepCopyInto(out *JournalNodeRoleGroupSpec) {
	*out = *in
	if in.Config != nil {
		in, out := &in.Config, &out.Config
		*out = new(JournalNodeConfigSpec)
		(*in).DeepCopyInto(*out)
	}
	if in.CommandArgsOverrides != nil {
		in, out := &in.CommandArgsOverrides, &out.CommandArgsOverrides
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.ConfigOverrides != nil {
		in, out := &in.ConfigOverrides, &out.ConfigOverrides
		*out = new(ConfigOverridesSpec)
		(*in).DeepCopyInto(*out)
	}
	if in.EnvOverrides != nil {
		in, out := &in.EnvOverrides, &out.EnvOverrides
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new JournalNodeRoleGroupSpec.
func (in *JournalNodeRoleGroupSpec) DeepCopy() *JournalNodeRoleGroupSpec {
	if in == nil {
		return nil
	}
	out := new(JournalNodeRoleGroupSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *JournalNodeSpec) DeepCopyInto(out *JournalNodeSpec) {
	*out = *in
	if in.Config != nil {
		in, out := &in.Config, &out.Config
		*out = new(JournalNodeConfigSpec)
		(*in).DeepCopyInto(*out)
	}
	if in.RoleGroups != nil {
		in, out := &in.RoleGroups, &out.RoleGroups
		*out = make(map[string]*JournalNodeRoleGroupSpec, len(*in))
		for key, val := range *in {
			var outVal *JournalNodeRoleGroupSpec
			if val == nil {
				(*out)[key] = nil
			} else {
				inVal := (*in)[key]
				in, out := &inVal, &outVal
				*out = new(JournalNodeRoleGroupSpec)
				(*in).DeepCopyInto(*out)
			}
			(*out)[key] = outVal
		}
	}
	if in.PodDisruptionBudget != nil {
		in, out := &in.PodDisruptionBudget, &out.PodDisruptionBudget
		*out = new(PodDisruptionBudgetSpec)
		**out = **in
	}
	if in.CommandArgsOverrides != nil {
		in, out := &in.CommandArgsOverrides, &out.CommandArgsOverrides
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.ConfigOverrides != nil {
		in, out := &in.ConfigOverrides, &out.ConfigOverrides
		*out = new(ConfigOverridesSpec)
		(*in).DeepCopyInto(*out)
	}
	if in.EnvOverrides != nil {
		in, out := &in.EnvOverrides, &out.EnvOverrides
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new JournalNodeSpec.
func (in *JournalNodeSpec) DeepCopy() *JournalNodeSpec {
	if in == nil {
		return nil
	}
	out := new(JournalNodeSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KerberosSpec) DeepCopyInto(out *KerberosSpec) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KerberosSpec.
func (in *KerberosSpec) DeepCopy() *KerberosSpec {
	if in == nil {
		return nil
	}
	out := new(KerberosSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LogLevelSpec) DeepCopyInto(out *LogLevelSpec) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LogLevelSpec.
func (in *LogLevelSpec) DeepCopy() *LogLevelSpec {
	if in == nil {
		return nil
	}
	out := new(LogLevelSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LoggingConfigSpec) DeepCopyInto(out *LoggingConfigSpec) {
	*out = *in
	if in.Loggers != nil {
		in, out := &in.Loggers, &out.Loggers
		*out = make(map[string]*LogLevelSpec, len(*in))
		for key, val := range *in {
			var outVal *LogLevelSpec
			if val == nil {
				(*out)[key] = nil
			} else {
				inVal := (*in)[key]
				in, out := &inVal, &outVal
				*out = new(LogLevelSpec)
				**out = **in
			}
			(*out)[key] = outVal
		}
	}
	if in.Console != nil {
		in, out := &in.Console, &out.Console
		*out = new(LogLevelSpec)
		**out = **in
	}
	if in.File != nil {
		in, out := &in.File, &out.File
		*out = new(LogLevelSpec)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LoggingConfigSpec.
func (in *LoggingConfigSpec) DeepCopy() *LoggingConfigSpec {
	if in == nil {
		return nil
	}
	out := new(LoggingConfigSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NameNodeConfigSpec) DeepCopyInto(out *NameNodeConfigSpec) {
	*out = *in
	if in.Resources != nil {
		in, out := &in.Resources, &out.Resources
		*out = new(commonsv1alpha1.ResourcesSpec)
		(*in).DeepCopyInto(*out)
	}
	if in.SecurityContext != nil {
		in, out := &in.SecurityContext, &out.SecurityContext
		*out = new(v1.PodSecurityContext)
		(*in).DeepCopyInto(*out)
	}
	if in.Affinity != nil {
		in, out := &in.Affinity, &out.Affinity
		*out = new(v1.Affinity)
		(*in).DeepCopyInto(*out)
	}
	if in.NodeSelector != nil {
		in, out := &in.NodeSelector, &out.NodeSelector
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.Tolerations != nil {
		in, out := &in.Tolerations, &out.Tolerations
		*out = make([]v1.Toleration, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.PodDisruptionBudget != nil {
		in, out := &in.PodDisruptionBudget, &out.PodDisruptionBudget
		*out = new(PodDisruptionBudgetSpec)
		**out = **in
	}
	if in.ExtraEnv != nil {
		in, out := &in.ExtraEnv, &out.ExtraEnv
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.ExtraSecret != nil {
		in, out := &in.ExtraSecret, &out.ExtraSecret
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.Logging != nil {
		in, out := &in.Logging, &out.Logging
		*out = new(NameNodeContainerLoggingSpec)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NameNodeConfigSpec.
func (in *NameNodeConfigSpec) DeepCopy() *NameNodeConfigSpec {
	if in == nil {
		return nil
	}
	out := new(NameNodeConfigSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NameNodeContainerLoggingSpec) DeepCopyInto(out *NameNodeContainerLoggingSpec) {
	*out = *in
	if in.NameNode != nil {
		in, out := &in.NameNode, &out.NameNode
		*out = new(LoggingConfigSpec)
		(*in).DeepCopyInto(*out)
	}
	if in.Zkfc != nil {
		in, out := &in.Zkfc, &out.Zkfc
		*out = new(LoggingConfigSpec)
		(*in).DeepCopyInto(*out)
	}
	if in.FormatZookeeper != nil {
		in, out := &in.FormatZookeeper, &out.FormatZookeeper
		*out = new(LoggingConfigSpec)
		(*in).DeepCopyInto(*out)
	}
	if in.FormatNameNode != nil {
		in, out := &in.FormatNameNode, &out.FormatNameNode
		*out = new(LoggingConfigSpec)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NameNodeContainerLoggingSpec.
func (in *NameNodeContainerLoggingSpec) DeepCopy() *NameNodeContainerLoggingSpec {
	if in == nil {
		return nil
	}
	out := new(NameNodeContainerLoggingSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NameNodeRoleGroupSpec) DeepCopyInto(out *NameNodeRoleGroupSpec) {
	*out = *in
	if in.Config != nil {
		in, out := &in.Config, &out.Config
		*out = new(NameNodeConfigSpec)
		(*in).DeepCopyInto(*out)
	}
	if in.CommandArgsOverrides != nil {
		in, out := &in.CommandArgsOverrides, &out.CommandArgsOverrides
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.ConfigOverrides != nil {
		in, out := &in.ConfigOverrides, &out.ConfigOverrides
		*out = new(ConfigOverridesSpec)
		(*in).DeepCopyInto(*out)
	}
	if in.EnvOverrides != nil {
		in, out := &in.EnvOverrides, &out.EnvOverrides
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NameNodeRoleGroupSpec.
func (in *NameNodeRoleGroupSpec) DeepCopy() *NameNodeRoleGroupSpec {
	if in == nil {
		return nil
	}
	out := new(NameNodeRoleGroupSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NameNodeSpec) DeepCopyInto(out *NameNodeSpec) {
	*out = *in
	if in.Config != nil {
		in, out := &in.Config, &out.Config
		*out = new(NameNodeConfigSpec)
		(*in).DeepCopyInto(*out)
	}
	if in.RoleGroups != nil {
		in, out := &in.RoleGroups, &out.RoleGroups
		*out = make(map[string]*NameNodeRoleGroupSpec, len(*in))
		for key, val := range *in {
			var outVal *NameNodeRoleGroupSpec
			if val == nil {
				(*out)[key] = nil
			} else {
				inVal := (*in)[key]
				in, out := &inVal, &outVal
				*out = new(NameNodeRoleGroupSpec)
				(*in).DeepCopyInto(*out)
			}
			(*out)[key] = outVal
		}
	}
	if in.PodDisruptionBudget != nil {
		in, out := &in.PodDisruptionBudget, &out.PodDisruptionBudget
		*out = new(PodDisruptionBudgetSpec)
		**out = **in
	}
	if in.CommandArgsOverrides != nil {
		in, out := &in.CommandArgsOverrides, &out.CommandArgsOverrides
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.ConfigOverrides != nil {
		in, out := &in.ConfigOverrides, &out.ConfigOverrides
		*out = new(ConfigOverridesSpec)
		(*in).DeepCopyInto(*out)
	}
	if in.EnvOverrides != nil {
		in, out := &in.EnvOverrides, &out.EnvOverrides
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NameNodeSpec.
func (in *NameNodeSpec) DeepCopy() *NameNodeSpec {
	if in == nil {
		return nil
	}
	out := new(NameNodeSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OidcSpec) DeepCopyInto(out *OidcSpec) {
	*out = *in
	if in.ExtraScopes != nil {
		in, out := &in.ExtraScopes, &out.ExtraScopes
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OidcSpec.
func (in *OidcSpec) DeepCopy() *OidcSpec {
	if in == nil {
		return nil
	}
	out := new(OidcSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PodDisruptionBudgetSpec) DeepCopyInto(out *PodDisruptionBudgetSpec) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PodDisruptionBudgetSpec.
func (in *PodDisruptionBudgetSpec) DeepCopy() *PodDisruptionBudgetSpec {
	if in == nil {
		return nil
	}
	out := new(PodDisruptionBudgetSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ServiceSpec) DeepCopyInto(out *ServiceSpec) {
	*out = *in
	if in.Annotations != nil {
		in, out := &in.Annotations, &out.Annotations
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServiceSpec.
func (in *ServiceSpec) DeepCopy() *ServiceSpec {
	if in == nil {
		return nil
	}
	out := new(ServiceSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TlsSpec) DeepCopyInto(out *TlsSpec) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TlsSpec.
func (in *TlsSpec) DeepCopy() *TlsSpec {
	if in == nil {
		return nil
	}
	out := new(TlsSpec)
	in.DeepCopyInto(out)
	return out
}
