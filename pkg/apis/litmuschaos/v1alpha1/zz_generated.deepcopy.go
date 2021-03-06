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

// Code generated by deepcopy-gen. DO NOT EDIT.

package v1alpha1

import (
	v1 "k8s.io/api/core/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ChaosSchedule) DeepCopyInto(out *ChaosSchedule) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ChaosSchedule.
func (in *ChaosSchedule) DeepCopy() *ChaosSchedule {
	if in == nil {
		return nil
	}
	out := new(ChaosSchedule)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ChaosSchedule) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ChaosScheduleList) DeepCopyInto(out *ChaosScheduleList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ChaosSchedule, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ChaosScheduleList.
func (in *ChaosScheduleList) DeepCopy() *ChaosScheduleList {
	if in == nil {
		return nil
	}
	out := new(ChaosScheduleList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ChaosScheduleList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ChaosScheduleSpec) DeepCopyInto(out *ChaosScheduleSpec) {
	*out = *in
	in.Schedule.DeepCopyInto(&out.Schedule)
	in.EngineTemplateSpec.DeepCopyInto(&out.EngineTemplateSpec)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ChaosScheduleSpec.
func (in *ChaosScheduleSpec) DeepCopy() *ChaosScheduleSpec {
	if in == nil {
		return nil
	}
	out := new(ChaosScheduleSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ChaosScheduleStatus) DeepCopyInto(out *ChaosScheduleStatus) {
	*out = *in
	in.Schedule.DeepCopyInto(&out.Schedule)
	if in.LastScheduleTime != nil {
		in, out := &in.LastScheduleTime, &out.LastScheduleTime
		*out = (*in).DeepCopy()
	}
	if in.Active != nil {
		in, out := &in.Active, &out.Active
		*out = make([]v1.ObjectReference, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ChaosScheduleStatus.
func (in *ChaosScheduleStatus) DeepCopy() *ChaosScheduleStatus {
	if in == nil {
		return nil
	}
	out := new(ChaosScheduleStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Schedule) DeepCopyInto(out *Schedule) {
	*out = *in
	if in.Once != nil {
		in, out := &in.Once, &out.Once
		*out = new(ScheduleOnce)
		(*in).DeepCopyInto(*out)
	}
	if in.Repeat != nil {
		in, out := &in.Repeat, &out.Repeat
		*out = new(ScheduleRepeat)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Schedule.
func (in *Schedule) DeepCopy() *Schedule {
	if in == nil {
		return nil
	}
	out := new(Schedule)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ScheduleOnce) DeepCopyInto(out *ScheduleOnce) {
	*out = *in
	in.ExecutionTime.DeepCopyInto(&out.ExecutionTime)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ScheduleOnce.
func (in *ScheduleOnce) DeepCopy() *ScheduleOnce {
	if in == nil {
		return nil
	}
	out := new(ScheduleOnce)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ScheduleRepeat) DeepCopyInto(out *ScheduleRepeat) {
	*out = *in
	if in.TimeRange != nil {
		in, out := &in.TimeRange, &out.TimeRange
		*out = new(TimeRange)
		(*in).DeepCopyInto(*out)
	}
	out.Properties = in.Properties
	if in.WorkHours != nil {
		in, out := &in.WorkHours, &out.WorkHours
		*out = new(WorkHours)
		**out = **in
	}
	if in.WorkDays != nil {
		in, out := &in.WorkDays, &out.WorkDays
		*out = new(WorkDays)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ScheduleRepeat.
func (in *ScheduleRepeat) DeepCopy() *ScheduleRepeat {
	if in == nil {
		return nil
	}
	out := new(ScheduleRepeat)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ScheduleRepeatProperties) DeepCopyInto(out *ScheduleRepeatProperties) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ScheduleRepeatProperties.
func (in *ScheduleRepeatProperties) DeepCopy() *ScheduleRepeatProperties {
	if in == nil {
		return nil
	}
	out := new(ScheduleRepeatProperties)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ScheduleStatus) DeepCopyInto(out *ScheduleStatus) {
	*out = *in
	if in.StartTime != nil {
		in, out := &in.StartTime, &out.StartTime
		*out = (*in).DeepCopy()
	}
	if in.EndTime != nil {
		in, out := &in.EndTime, &out.EndTime
		*out = (*in).DeepCopy()
	}
	if in.ExpectedNextRunTime != nil {
		in, out := &in.ExpectedNextRunTime, &out.ExpectedNextRunTime
		*out = (*in).DeepCopy()
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ScheduleStatus.
func (in *ScheduleStatus) DeepCopy() *ScheduleStatus {
	if in == nil {
		return nil
	}
	out := new(ScheduleStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TimeRange) DeepCopyInto(out *TimeRange) {
	*out = *in
	if in.StartTime != nil {
		in, out := &in.StartTime, &out.StartTime
		*out = (*in).DeepCopy()
	}
	if in.EndTime != nil {
		in, out := &in.EndTime, &out.EndTime
		*out = (*in).DeepCopy()
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TimeRange.
func (in *TimeRange) DeepCopy() *TimeRange {
	if in == nil {
		return nil
	}
	out := new(TimeRange)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WorkDays) DeepCopyInto(out *WorkDays) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WorkDays.
func (in *WorkDays) DeepCopy() *WorkDays {
	if in == nil {
		return nil
	}
	out := new(WorkDays)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WorkHours) DeepCopyInto(out *WorkHours) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WorkHours.
func (in *WorkHours) DeepCopy() *WorkHours {
	if in == nil {
		return nil
	}
	out := new(WorkHours)
	in.DeepCopyInto(out)
	return out
}
