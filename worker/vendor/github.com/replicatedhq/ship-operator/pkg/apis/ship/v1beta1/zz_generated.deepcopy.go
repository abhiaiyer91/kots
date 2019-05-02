// +build !ignore_autogenerated

// Code generated by main. DO NOT EDIT.

package v1beta1

import (
	v1 "k8s.io/api/core/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ActionSpec) DeepCopyInto(out *ActionSpec) {
	*out = *in
	if in.PullRequest != nil {
		in, out := &in.PullRequest, &out.PullRequest
		*out = new(PullRequestActionSpec)
		(*in).DeepCopyInto(*out)
	}
	if in.Webhook != nil {
		in, out := &in.Webhook, &out.Webhook
		*out = new(WebhookActionSpec)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ActionSpec.
func (in *ActionSpec) DeepCopy() *ActionSpec {
	if in == nil {
		return nil
	}
	out := new(ActionSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GitHubRef) DeepCopyInto(out *GitHubRef) {
	*out = *in
	out.Token = in.Token
	out.Key = in.Key
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GitHubRef.
func (in *GitHubRef) DeepCopy() *GitHubRef {
	if in == nil {
		return nil
	}
	out := new(GitHubRef)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ImageSpec) DeepCopyInto(out *ImageSpec) {
	*out = *in
	return
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
func (in *KeySpec) DeepCopyInto(out *KeySpec) {
	*out = *in
	out.ValueFrom = in.ValueFrom
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KeySpec.
func (in *KeySpec) DeepCopy() *KeySpec {
	if in == nil {
		return nil
	}
	out := new(KeySpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PullRequestActionSpec) DeepCopyInto(out *PullRequestActionSpec) {
	*out = *in
	if in.GitHub != nil {
		in, out := &in.GitHub, &out.GitHub
		*out = new(GitHubRef)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PullRequestActionSpec.
func (in *PullRequestActionSpec) DeepCopy() *PullRequestActionSpec {
	if in == nil {
		return nil
	}
	out := new(PullRequestActionSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SecretKeyRef) DeepCopyInto(out *SecretKeyRef) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SecretKeyRef.
func (in *SecretKeyRef) DeepCopy() *SecretKeyRef {
	if in == nil {
		return nil
	}
	out := new(SecretKeyRef)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ShipWatch) DeepCopyInto(out *ShipWatch) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ShipWatch.
func (in *ShipWatch) DeepCopy() *ShipWatch {
	if in == nil {
		return nil
	}
	out := new(ShipWatch)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ShipWatch) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ShipWatchList) DeepCopyInto(out *ShipWatchList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ShipWatch, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ShipWatchList.
func (in *ShipWatchList) DeepCopy() *ShipWatchList {
	if in == nil {
		return nil
	}
	out := new(ShipWatchList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ShipWatchList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ShipWatchSpec) DeepCopyInto(out *ShipWatchSpec) {
	*out = *in
	out.State = in.State
	if in.Images != nil {
		in, out := &in.Images, &out.Images
		*out = make([]ImageSpec, len(*in))
		copy(*out, *in)
	}
	if in.Environment != nil {
		in, out := &in.Environment, &out.Environment
		*out = make([]v1.EnvVar, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Actions != nil {
		in, out := &in.Actions, &out.Actions
		*out = make([]ActionSpec, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ShipWatchSpec.
func (in *ShipWatchSpec) DeepCopy() *ShipWatchSpec {
	if in == nil {
		return nil
	}
	out := new(ShipWatchSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ShipWatchStatus) DeepCopyInto(out *ShipWatchStatus) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ShipWatchStatus.
func (in *ShipWatchStatus) DeepCopy() *ShipWatchStatus {
	if in == nil {
		return nil
	}
	out := new(ShipWatchStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ShipWatchValueFromSpec) DeepCopyInto(out *ShipWatchValueFromSpec) {
	*out = *in
	out.SecretKeyRef = in.SecretKeyRef
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ShipWatchValueFromSpec.
func (in *ShipWatchValueFromSpec) DeepCopy() *ShipWatchValueFromSpec {
	if in == nil {
		return nil
	}
	out := new(ShipWatchValueFromSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StateSpec) DeepCopyInto(out *StateSpec) {
	*out = *in
	out.ValueFrom = in.ValueFrom
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StateSpec.
func (in *StateSpec) DeepCopy() *StateSpec {
	if in == nil {
		return nil
	}
	out := new(StateSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TokenSpec) DeepCopyInto(out *TokenSpec) {
	*out = *in
	out.ValueFrom = in.ValueFrom
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TokenSpec.
func (in *TokenSpec) DeepCopy() *TokenSpec {
	if in == nil {
		return nil
	}
	out := new(TokenSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WebhookActionSpec) DeepCopyInto(out *WebhookActionSpec) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WebhookActionSpec.
func (in *WebhookActionSpec) DeepCopy() *WebhookActionSpec {
	if in == nil {
		return nil
	}
	out := new(WebhookActionSpec)
	in.DeepCopyInto(out)
	return out
}
