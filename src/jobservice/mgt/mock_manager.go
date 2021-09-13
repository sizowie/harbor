// Code generated by mockery v2.1.0. DO NOT EDIT.

package mgt

import (
	job "github.com/goharbor/harbor/src/jobservice/job"
	mock "github.com/stretchr/testify/mock"

	query "github.com/goharbor/harbor/src/jobservice/common/query"
)

// MockManager is an autogenerated mock type for the Manager type
type MockManager struct {
	mock.Mock
}

// GetJob provides a mock function with given fields: jobID
func (_m *MockManager) GetJob(jobID string) (*job.Stats, error) {
	ret := _m.Called(jobID)

	var r0 *job.Stats
	if rf, ok := ret.Get(0).(func(string) *job.Stats); ok {
		r0 = rf(jobID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*job.Stats)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(jobID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetJobs provides a mock function with given fields: q
func (_m *MockManager) GetJobs(q *query.Parameter) ([]*job.Stats, int64, error) {
	ret := _m.Called(q)

	var r0 []*job.Stats
	if rf, ok := ret.Get(0).(func(*query.Parameter) []*job.Stats); ok {
		r0 = rf(q)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*job.Stats)
		}
	}

	var r1 int64
	if rf, ok := ret.Get(1).(func(*query.Parameter) int64); ok {
		r1 = rf(q)
	} else {
		r1 = ret.Get(1).(int64)
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(*query.Parameter) error); ok {
		r2 = rf(q)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// GetPeriodicExecution provides a mock function with given fields: pID, q
func (_m *MockManager) GetPeriodicExecution(pID string, q *query.Parameter) ([]*job.Stats, int64, error) {
	ret := _m.Called(pID, q)

	var r0 []*job.Stats
	if rf, ok := ret.Get(0).(func(string, *query.Parameter) []*job.Stats); ok {
		r0 = rf(pID, q)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*job.Stats)
		}
	}

	var r1 int64
	if rf, ok := ret.Get(1).(func(string, *query.Parameter) int64); ok {
		r1 = rf(pID, q)
	} else {
		r1 = ret.Get(1).(int64)
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(string, *query.Parameter) error); ok {
		r2 = rf(pID, q)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// GetScheduledJobs provides a mock function with given fields: q
func (_m *MockManager) GetScheduledJobs(q *query.Parameter) ([]*job.Stats, int64, error) {
	ret := _m.Called(q)

	var r0 []*job.Stats
	if rf, ok := ret.Get(0).(func(*query.Parameter) []*job.Stats); ok {
		r0 = rf(q)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*job.Stats)
		}
	}

	var r1 int64
	if rf, ok := ret.Get(1).(func(*query.Parameter) int64); ok {
		r1 = rf(q)
	} else {
		r1 = ret.Get(1).(int64)
	}

	var r2 error
	if rf, ok := ret.Get(2).(func(*query.Parameter) error); ok {
		r2 = rf(q)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// SaveJob provides a mock function with given fields: _a0
func (_m *MockManager) SaveJob(_a0 *job.Stats) error {
	ret := _m.Called(_a0)

	var r0 error
	if rf, ok := ret.Get(0).(func(*job.Stats) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
