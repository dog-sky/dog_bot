package mocks

// Code generated by http://github.com/gojuno/minimock (dev). DO NOT EDIT.

//go:generate minimock -i github.com/dog-sky/dog_bot/internal/service/db.DB -o ./db_mock_test.go -n DBMock

import (
	"sync"
	mm_atomic "sync/atomic"
	mm_time "time"

	"github.com/gojuno/minimock/v3"
)

// DBMock implements DB
type DBMock struct {
	t minimock.Tester

	funcSetStatus          func(status string) (err error)
	inspectFuncSetStatus   func(status string)
	afterSetStatusCounter  uint64
	beforeSetStatusCounter uint64
	SetStatusMock          mDBMockSetStatus

	funcShutDown          func()
	inspectFuncShutDown   func()
	afterShutDownCounter  uint64
	beforeShutDownCounter uint64
	ShutDownMock          mDBMockShutDown
}

// NewDBMock returns a mock for DB
func NewDBMock(t minimock.Tester) *DBMock {
	m := &DBMock{t: t}
	if controller, ok := t.(minimock.MockController); ok {
		controller.RegisterMocker(m)
	}

	m.SetStatusMock = mDBMockSetStatus{mock: m}
	m.SetStatusMock.callArgs = []*DBMockSetStatusParams{}

	m.ShutDownMock = mDBMockShutDown{mock: m}

	return m
}

type mDBMockSetStatus struct {
	mock               *DBMock
	defaultExpectation *DBMockSetStatusExpectation
	expectations       []*DBMockSetStatusExpectation

	callArgs []*DBMockSetStatusParams
	mutex    sync.RWMutex
}

// DBMockSetStatusExpectation specifies expectation struct of the DB.SetStatus
type DBMockSetStatusExpectation struct {
	mock    *DBMock
	params  *DBMockSetStatusParams
	results *DBMockSetStatusResults
	Counter uint64
}

// DBMockSetStatusParams contains parameters of the DB.SetStatus
type DBMockSetStatusParams struct {
	status string
}

// DBMockSetStatusResults contains results of the DB.SetStatus
type DBMockSetStatusResults struct {
	err error
}

// Expect sets up expected params for DB.SetStatus
func (mmSetStatus *mDBMockSetStatus) Expect(status string) *mDBMockSetStatus {
	if mmSetStatus.mock.funcSetStatus != nil {
		mmSetStatus.mock.t.Fatalf("DBMock.SetStatus mock is already set by Set")
	}

	if mmSetStatus.defaultExpectation == nil {
		mmSetStatus.defaultExpectation = &DBMockSetStatusExpectation{}
	}

	mmSetStatus.defaultExpectation.params = &DBMockSetStatusParams{status}
	for _, e := range mmSetStatus.expectations {
		if minimock.Equal(e.params, mmSetStatus.defaultExpectation.params) {
			mmSetStatus.mock.t.Fatalf("Expectation set by When has same params: %#v", *mmSetStatus.defaultExpectation.params)
		}
	}

	return mmSetStatus
}

// Inspect accepts an inspector function that has same arguments as the DB.SetStatus
func (mmSetStatus *mDBMockSetStatus) Inspect(f func(status string)) *mDBMockSetStatus {
	if mmSetStatus.mock.inspectFuncSetStatus != nil {
		mmSetStatus.mock.t.Fatalf("Inspect function is already set for DBMock.SetStatus")
	}

	mmSetStatus.mock.inspectFuncSetStatus = f

	return mmSetStatus
}

// Return sets up results that will be returned by DB.SetStatus
func (mmSetStatus *mDBMockSetStatus) Return(err error) *DBMock {
	if mmSetStatus.mock.funcSetStatus != nil {
		mmSetStatus.mock.t.Fatalf("DBMock.SetStatus mock is already set by Set")
	}

	if mmSetStatus.defaultExpectation == nil {
		mmSetStatus.defaultExpectation = &DBMockSetStatusExpectation{mock: mmSetStatus.mock}
	}
	mmSetStatus.defaultExpectation.results = &DBMockSetStatusResults{err}
	return mmSetStatus.mock
}

//Set uses given function f to mock the DB.SetStatus method
func (mmSetStatus *mDBMockSetStatus) Set(f func(status string) (err error)) *DBMock {
	if mmSetStatus.defaultExpectation != nil {
		mmSetStatus.mock.t.Fatalf("Default expectation is already set for the DB.SetStatus method")
	}

	if len(mmSetStatus.expectations) > 0 {
		mmSetStatus.mock.t.Fatalf("Some expectations are already set for the DB.SetStatus method")
	}

	mmSetStatus.mock.funcSetStatus = f
	return mmSetStatus.mock
}

// When sets expectation for the DB.SetStatus which will trigger the result defined by the following
// Then helper
func (mmSetStatus *mDBMockSetStatus) When(status string) *DBMockSetStatusExpectation {
	if mmSetStatus.mock.funcSetStatus != nil {
		mmSetStatus.mock.t.Fatalf("DBMock.SetStatus mock is already set by Set")
	}

	expectation := &DBMockSetStatusExpectation{
		mock:   mmSetStatus.mock,
		params: &DBMockSetStatusParams{status},
	}
	mmSetStatus.expectations = append(mmSetStatus.expectations, expectation)
	return expectation
}

// Then sets up DB.SetStatus return parameters for the expectation previously defined by the When method
func (e *DBMockSetStatusExpectation) Then(err error) *DBMock {
	e.results = &DBMockSetStatusResults{err}
	return e.mock
}

// SetStatus implements DB
func (mmSetStatus *DBMock) SetStatus(status string) (err error) {
	mm_atomic.AddUint64(&mmSetStatus.beforeSetStatusCounter, 1)
	defer mm_atomic.AddUint64(&mmSetStatus.afterSetStatusCounter, 1)

	if mmSetStatus.inspectFuncSetStatus != nil {
		mmSetStatus.inspectFuncSetStatus(status)
	}

	mm_params := &DBMockSetStatusParams{status}

	// Record call args
	mmSetStatus.SetStatusMock.mutex.Lock()
	mmSetStatus.SetStatusMock.callArgs = append(mmSetStatus.SetStatusMock.callArgs, mm_params)
	mmSetStatus.SetStatusMock.mutex.Unlock()

	for _, e := range mmSetStatus.SetStatusMock.expectations {
		if minimock.Equal(e.params, mm_params) {
			mm_atomic.AddUint64(&e.Counter, 1)
			return e.results.err
		}
	}

	if mmSetStatus.SetStatusMock.defaultExpectation != nil {
		mm_atomic.AddUint64(&mmSetStatus.SetStatusMock.defaultExpectation.Counter, 1)
		mm_want := mmSetStatus.SetStatusMock.defaultExpectation.params
		mm_got := DBMockSetStatusParams{status}
		if mm_want != nil && !minimock.Equal(*mm_want, mm_got) {
			mmSetStatus.t.Errorf("DBMock.SetStatus got unexpected parameters, want: %#v, got: %#v%s\n", *mm_want, mm_got, minimock.Diff(*mm_want, mm_got))
		}

		mm_results := mmSetStatus.SetStatusMock.defaultExpectation.results
		if mm_results == nil {
			mmSetStatus.t.Fatal("No results are set for the DBMock.SetStatus")
		}
		return (*mm_results).err
	}
	if mmSetStatus.funcSetStatus != nil {
		return mmSetStatus.funcSetStatus(status)
	}
	mmSetStatus.t.Fatalf("Unexpected call to DBMock.SetStatus. %v", status)
	return
}

// SetStatusAfterCounter returns a count of finished DBMock.SetStatus invocations
func (mmSetStatus *DBMock) SetStatusAfterCounter() uint64 {
	return mm_atomic.LoadUint64(&mmSetStatus.afterSetStatusCounter)
}

// SetStatusBeforeCounter returns a count of DBMock.SetStatus invocations
func (mmSetStatus *DBMock) SetStatusBeforeCounter() uint64 {
	return mm_atomic.LoadUint64(&mmSetStatus.beforeSetStatusCounter)
}

// Calls returns a list of arguments used in each call to DBMock.SetStatus.
// The list is in the same order as the calls were made (i.e. recent calls have a higher index)
func (mmSetStatus *mDBMockSetStatus) Calls() []*DBMockSetStatusParams {
	mmSetStatus.mutex.RLock()

	argCopy := make([]*DBMockSetStatusParams, len(mmSetStatus.callArgs))
	copy(argCopy, mmSetStatus.callArgs)

	mmSetStatus.mutex.RUnlock()

	return argCopy
}

// MinimockSetStatusDone returns true if the count of the SetStatus invocations corresponds
// the number of defined expectations
func (m *DBMock) MinimockSetStatusDone() bool {
	for _, e := range m.SetStatusMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			return false
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.SetStatusMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterSetStatusCounter) < 1 {
		return false
	}
	// if func was set then invocations count should be greater than zero
	if m.funcSetStatus != nil && mm_atomic.LoadUint64(&m.afterSetStatusCounter) < 1 {
		return false
	}
	return true
}

// MinimockSetStatusInspect logs each unmet expectation
func (m *DBMock) MinimockSetStatusInspect() {
	for _, e := range m.SetStatusMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			m.t.Errorf("Expected call to DBMock.SetStatus with params: %#v", *e.params)
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.SetStatusMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterSetStatusCounter) < 1 {
		if m.SetStatusMock.defaultExpectation.params == nil {
			m.t.Error("Expected call to DBMock.SetStatus")
		} else {
			m.t.Errorf("Expected call to DBMock.SetStatus with params: %#v", *m.SetStatusMock.defaultExpectation.params)
		}
	}
	// if func was set then invocations count should be greater than zero
	if m.funcSetStatus != nil && mm_atomic.LoadUint64(&m.afterSetStatusCounter) < 1 {
		m.t.Error("Expected call to DBMock.SetStatus")
	}
}

type mDBMockShutDown struct {
	mock               *DBMock
	defaultExpectation *DBMockShutDownExpectation
	expectations       []*DBMockShutDownExpectation
}

// DBMockShutDownExpectation specifies expectation struct of the DB.ShutDown
type DBMockShutDownExpectation struct {
	mock *DBMock

	Counter uint64
}

// Expect sets up expected params for DB.ShutDown
func (mmShutDown *mDBMockShutDown) Expect() *mDBMockShutDown {
	if mmShutDown.mock.funcShutDown != nil {
		mmShutDown.mock.t.Fatalf("DBMock.ShutDown mock is already set by Set")
	}

	if mmShutDown.defaultExpectation == nil {
		mmShutDown.defaultExpectation = &DBMockShutDownExpectation{}
	}

	return mmShutDown
}

// Inspect accepts an inspector function that has same arguments as the DB.ShutDown
func (mmShutDown *mDBMockShutDown) Inspect(f func()) *mDBMockShutDown {
	if mmShutDown.mock.inspectFuncShutDown != nil {
		mmShutDown.mock.t.Fatalf("Inspect function is already set for DBMock.ShutDown")
	}

	mmShutDown.mock.inspectFuncShutDown = f

	return mmShutDown
}

// Return sets up results that will be returned by DB.ShutDown
func (mmShutDown *mDBMockShutDown) Return() *DBMock {
	if mmShutDown.mock.funcShutDown != nil {
		mmShutDown.mock.t.Fatalf("DBMock.ShutDown mock is already set by Set")
	}

	if mmShutDown.defaultExpectation == nil {
		mmShutDown.defaultExpectation = &DBMockShutDownExpectation{mock: mmShutDown.mock}
	}

	return mmShutDown.mock
}

//Set uses given function f to mock the DB.ShutDown method
func (mmShutDown *mDBMockShutDown) Set(f func()) *DBMock {
	if mmShutDown.defaultExpectation != nil {
		mmShutDown.mock.t.Fatalf("Default expectation is already set for the DB.ShutDown method")
	}

	if len(mmShutDown.expectations) > 0 {
		mmShutDown.mock.t.Fatalf("Some expectations are already set for the DB.ShutDown method")
	}

	mmShutDown.mock.funcShutDown = f
	return mmShutDown.mock
}

// ShutDown implements DB
func (mmShutDown *DBMock) ShutDown() {
	mm_atomic.AddUint64(&mmShutDown.beforeShutDownCounter, 1)
	defer mm_atomic.AddUint64(&mmShutDown.afterShutDownCounter, 1)

	if mmShutDown.inspectFuncShutDown != nil {
		mmShutDown.inspectFuncShutDown()
	}

	if mmShutDown.ShutDownMock.defaultExpectation != nil {
		mm_atomic.AddUint64(&mmShutDown.ShutDownMock.defaultExpectation.Counter, 1)

		return

	}
	if mmShutDown.funcShutDown != nil {
		mmShutDown.funcShutDown()
		return
	}
	mmShutDown.t.Fatalf("Unexpected call to DBMock.ShutDown.")

}

// ShutDownAfterCounter returns a count of finished DBMock.ShutDown invocations
func (mmShutDown *DBMock) ShutDownAfterCounter() uint64 {
	return mm_atomic.LoadUint64(&mmShutDown.afterShutDownCounter)
}

// ShutDownBeforeCounter returns a count of DBMock.ShutDown invocations
func (mmShutDown *DBMock) ShutDownBeforeCounter() uint64 {
	return mm_atomic.LoadUint64(&mmShutDown.beforeShutDownCounter)
}

// MinimockShutDownDone returns true if the count of the ShutDown invocations corresponds
// the number of defined expectations
func (m *DBMock) MinimockShutDownDone() bool {
	for _, e := range m.ShutDownMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			return false
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.ShutDownMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterShutDownCounter) < 1 {
		return false
	}
	// if func was set then invocations count should be greater than zero
	if m.funcShutDown != nil && mm_atomic.LoadUint64(&m.afterShutDownCounter) < 1 {
		return false
	}
	return true
}

// MinimockShutDownInspect logs each unmet expectation
func (m *DBMock) MinimockShutDownInspect() {
	for _, e := range m.ShutDownMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			m.t.Error("Expected call to DBMock.ShutDown")
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.ShutDownMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterShutDownCounter) < 1 {
		m.t.Error("Expected call to DBMock.ShutDown")
	}
	// if func was set then invocations count should be greater than zero
	if m.funcShutDown != nil && mm_atomic.LoadUint64(&m.afterShutDownCounter) < 1 {
		m.t.Error("Expected call to DBMock.ShutDown")
	}
}

// MinimockFinish checks that all mocked methods have been called the expected number of times
func (m *DBMock) MinimockFinish() {
	if !m.minimockDone() {
		m.MinimockSetStatusInspect()

		m.MinimockShutDownInspect()
		m.t.FailNow()
	}
}

// MinimockWait waits for all mocked methods to be called the expected number of times
func (m *DBMock) MinimockWait(timeout mm_time.Duration) {
	timeoutCh := mm_time.After(timeout)
	for {
		if m.minimockDone() {
			return
		}
		select {
		case <-timeoutCh:
			m.MinimockFinish()
			return
		case <-mm_time.After(10 * mm_time.Millisecond):
		}
	}
}

func (m *DBMock) minimockDone() bool {
	done := true
	return done &&
		m.MinimockSetStatusDone() &&
		m.MinimockShutDownDone()
}
