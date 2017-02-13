// This file was generated by github.com/nelsam/hel.  Do not
// edit this code by hand unless you *really* know what you're
// doing.  Expect any changes made manually to be overwritten
// the next time hel regenerates this file.

package logcounter_test

type mockUAA struct {
	GetAuthTokenCalled chan bool
	GetAuthTokenOutput struct {
		Ret0 chan string
		Ret1 chan error
	}
}

func newMockUAA() *mockUAA {
	m := &mockUAA{}
	m.GetAuthTokenCalled = make(chan bool, 100)
	m.GetAuthTokenOutput.Ret0 = make(chan string, 100)
	m.GetAuthTokenOutput.Ret1 = make(chan error, 100)
	return m
}
func (m *mockUAA) GetAuthToken() (string, error) {
	m.GetAuthTokenCalled <- true
	return <-m.GetAuthTokenOutput.Ret0, <-m.GetAuthTokenOutput.Ret1
}
