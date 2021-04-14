package app

import (
	"github.com/stretchr/testify/mock"

	"github.com/ZupIT/horusec-devkit/pkg/enums/auth"
)

type Mock struct {
	mock.Mock
}

func (m *Mock) IsBrokerDisabled() bool {
	args := m.MethodCalled("IsBrokerDisabled")
	return args.Get(0).(bool)
}

func (m *Mock) IsApplicationAdmEnabled() bool {
	args := m.MethodCalled("IsApplicationAdmEnabled")
	return args.Get(0).(bool)
}

func (m *Mock) GetAuthenticationType() auth.AuthenticationType {
	args := m.MethodCalled("GetAuthenticationType")
	return args.Get(0).(auth.AuthenticationType)
}
