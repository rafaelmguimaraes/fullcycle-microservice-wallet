package database

import (
	"database/sql"
	"testing"

	"github.com/rafaelmguimaraes/fullcycle-microservice-wallet/internal/entity"
	"github.com/stretchr/testify/suite"
)

type AccountDBTestSuite struct {
	suite.Suite
	db        *sql.DB
	accountDB *AccountDB
	client    *entity.Client
}

func (s *AccountDBTestSuite) SetupSuite() {
	db, err := sql.Open("sqlite3", ":memory:")
	if err != nil {
		s.T().Fatal(err)
	}
	s.db = db
	db.Exec("CREATE TABLE clients (id varchar(255), name varchar(255), email varchar(255), created_at date)")
	db.Exec("CREATE TABLE accounts (id varchar(255), client_id varchar(255), balance float, created_at date)")
	s.accountDB = NewAccountDB(db)
	s.client, _ = entity.NewClient("John", "j@j.com")

}

func (s *AccountDBTestSuite) TearDownSuite() {
	defer s.db.Close()
	s.db.Exec("DROP TABLE accounts")
	s.db.Exec("DROP TABLE clients")
}

func TestAccountDBTestSuite(t *testing.T) {
	suite.Run(t, new(AccountDBTestSuite))
}

func (s *AccountDBTestSuite) TestAccountDB_Save() {
	account, _ := entity.NewAccount(s.client)
	err := s.accountDB.Save(account)
	s.Nil(err)
}

func (s *AccountDBTestSuite) TestAccountDB_Get() {
	s.db.Exec("INSERT INTO clients (id, name, email, created_at) VALUES (?, ?, ?, ?)", s.client.ID, s.client.Name, s.client.Email, s.client.CreatedAt)
	account, _ := entity.NewAccount(s.client)
	err := s.accountDB.Save(account)
	s.Nil(err)

	accountFromDB, err := s.accountDB.Get(account.ID)
	s.Nil(err)
	s.Equal(account.ID, accountFromDB.ID)
	s.Equal(account.Client.ID, accountFromDB.Client.ID)
	s.Equal(account.Balance, accountFromDB.Balance)
	s.Equal(account.Client.Name, accountFromDB.Client.Name)
	s.Equal(account.Client.Email, accountFromDB.Client.Email)
}
