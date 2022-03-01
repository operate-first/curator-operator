//go:build !ignore_autogenerated
// +build !ignore_autogenerated

/*
Copyright 2022.

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
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CuratorConfig) DeepCopyInto(out *CuratorConfig) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	out.Status = in.Status
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CuratorConfig.
func (in *CuratorConfig) DeepCopy() *CuratorConfig {
	if in == nil {
		return nil
	}
	out := new(CuratorConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *CuratorConfig) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CuratorConfigList) DeepCopyInto(out *CuratorConfigList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]CuratorConfig, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CuratorConfigList.
func (in *CuratorConfigList) DeepCopy() *CuratorConfigList {
	if in == nil {
		return nil
	}
	out := new(CuratorConfigList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *CuratorConfigList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CuratorConfigSpec) DeepCopyInto(out *CuratorConfigSpec) {
	*out = *in
	out.Storage = in.Storage
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CuratorConfigSpec.
func (in *CuratorConfigSpec) DeepCopy() *CuratorConfigSpec {
	if in == nil {
		return nil
	}
	out := new(CuratorConfigSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CuratorConfigStatus) DeepCopyInto(out *CuratorConfigStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CuratorConfigStatus.
func (in *CuratorConfigStatus) DeepCopy() *CuratorConfigStatus {
	if in == nil {
		return nil
	}
	out := new(CuratorConfigStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Report) DeepCopyInto(out *Report) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Report.
func (in *Report) DeepCopy() *Report {
	if in == nil {
		return nil
	}
	out := new(Report)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Report) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ReportList) DeepCopyInto(out *ReportList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Report, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ReportList.
func (in *ReportList) DeepCopy() *ReportList {
	if in == nil {
		return nil
	}
	out := new(ReportList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ReportList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ReportSchedule) DeepCopyInto(out *ReportSchedule) {
	*out = *in
	if in.Cron != nil {
		in, out := &in.Cron, &out.Cron
		*out = new(ReportScheduleCron)
		**out = **in
	}
	if in.Hourly != nil {
		in, out := &in.Hourly, &out.Hourly
		*out = new(ReportScheduleHourly)
		**out = **in
	}
	if in.Daily != nil {
		in, out := &in.Daily, &out.Daily
		*out = new(ReportScheduleDaily)
		**out = **in
	}
	if in.Weekly != nil {
		in, out := &in.Weekly, &out.Weekly
		*out = new(ReportScheduleWeekly)
		(*in).DeepCopyInto(*out)
	}
	if in.Monthly != nil {
		in, out := &in.Monthly, &out.Monthly
		*out = new(ReportScheduleMonthly)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ReportSchedule.
func (in *ReportSchedule) DeepCopy() *ReportSchedule {
	if in == nil {
		return nil
	}
	out := new(ReportSchedule)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ReportScheduleCron) DeepCopyInto(out *ReportScheduleCron) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ReportScheduleCron.
func (in *ReportScheduleCron) DeepCopy() *ReportScheduleCron {
	if in == nil {
		return nil
	}
	out := new(ReportScheduleCron)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ReportScheduleDaily) DeepCopyInto(out *ReportScheduleDaily) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ReportScheduleDaily.
func (in *ReportScheduleDaily) DeepCopy() *ReportScheduleDaily {
	if in == nil {
		return nil
	}
	out := new(ReportScheduleDaily)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ReportScheduleHourly) DeepCopyInto(out *ReportScheduleHourly) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ReportScheduleHourly.
func (in *ReportScheduleHourly) DeepCopy() *ReportScheduleHourly {
	if in == nil {
		return nil
	}
	out := new(ReportScheduleHourly)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ReportScheduleMonthly) DeepCopyInto(out *ReportScheduleMonthly) {
	*out = *in
	if in.DayOfMonth != nil {
		in, out := &in.DayOfMonth, &out.DayOfMonth
		*out = new(int64)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ReportScheduleMonthly.
func (in *ReportScheduleMonthly) DeepCopy() *ReportScheduleMonthly {
	if in == nil {
		return nil
	}
	out := new(ReportScheduleMonthly)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ReportScheduleWeekly) DeepCopyInto(out *ReportScheduleWeekly) {
	*out = *in
	if in.DayOfWeek != nil {
		in, out := &in.DayOfWeek, &out.DayOfWeek
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ReportScheduleWeekly.
func (in *ReportScheduleWeekly) DeepCopy() *ReportScheduleWeekly {
	if in == nil {
		return nil
	}
	out := new(ReportScheduleWeekly)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ReportSpec) DeepCopyInto(out *ReportSpec) {
	*out = *in
	if in.Schedule != nil {
		in, out := &in.Schedule, &out.Schedule
		*out = new(ReportSchedule)
		(*in).DeepCopyInto(*out)
	}
	if in.ReportingEnd != nil {
		in, out := &in.ReportingEnd, &out.ReportingEnd
		*out = (*in).DeepCopy()
	}
	if in.ReportingStart != nil {
		in, out := &in.ReportingStart, &out.ReportingStart
		*out = (*in).DeepCopy()
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ReportSpec.
func (in *ReportSpec) DeepCopy() *ReportSpec {
	if in == nil {
		return nil
	}
	out := new(ReportSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ReportStatus) DeepCopyInto(out *ReportStatus) {
	*out = *in
	if in.LastReportTime != nil {
		in, out := &in.LastReportTime, &out.LastReportTime
		*out = (*in).DeepCopy()
	}
	if in.NextReportTime != nil {
		in, out := &in.NextReportTime, &out.NextReportTime
		*out = (*in).DeepCopy()
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ReportStatus.
func (in *ReportStatus) DeepCopy() *ReportStatus {
	if in == nil {
		return nil
	}
	out := new(ReportStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Storage) DeepCopyInto(out *Storage) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Storage.
func (in *Storage) DeepCopy() *Storage {
	if in == nil {
		return nil
	}
	out := new(Storage)
	in.DeepCopyInto(out)
	return out
}
