package database

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockConnection struct {
	mock.Mock
}

// func (m *MockConnection) Connect(config *gorm.Config) (*gorm.DB, error) {
// 	args := m.Called(config)
// 	return args.Get(0).(*gorm.DB), args.Error(1)
// }

// func TestInit(t *testing.T) {
// 	mockConnection := new(MockConnection)
// 	mockConnection.On("Connect", mock.Anything).Return(&gorm.DB{}, nil)

// 	oldConnectionFactory := connectionFactory
// 	defer func() { connectionFactory = oldConnectionFactory }()
// 	connectionFactory = func(factory string) (connection, error) {
// 		return mockConnection, nil
// 	}

// 	Init()

// 	mockConnection.AssertExpectations(t)
// }

func TestGetConnection(t *testing.T) {
	db := GetConnection()
	assert.NotNil(t, db)
}

// func TestMakeMigration(t *testing.T) {
// 	db := &gorm.DB{}
// 	db.On("AutoMigrate", mock.Anything).Return(nil)

// 	MakeMigration(db)

// 	db.AssertExpectations(t)
// }

func TestConnectionFactory(t *testing.T) {
	_, err := connectionFactory("invalid")
	assert.Error(t, err)

	_, err = connectionFactory("postgres")
	assert.NoError(t, err)
}
