package models

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestInsertClass(t *testing.T) {
	db := initTestDB()
	createTestDbIfNotExists(db)
	clearDb(db)
	m := Class{}
	err := m.Insert(db)
	assert.NotNil(t, err)
	fmt.Printf("%+v", err)
}
