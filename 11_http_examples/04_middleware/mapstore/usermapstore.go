package mapstore

import (
	"fmt"
	"go-examples/11_http_examples/04_middleware/domain"
)

type MapStore struct {
	store map[string]domain.User
}

func NewUserMapStore() *MapStore {
	return &MapStore{
		store: make(map[string]domain.User),
	}
}

func (m MapStore) isRecordExist(id string) bool {
	_, ok := m.store[id]
	return ok
}

func (m MapStore) Create(doc domain.User) error {
	if m.isRecordExist(doc.ID) {
		return fmt.Errorf("user already exists")
	}
	m.store[doc.ID] = doc
	return nil
}

func (m MapStore) Update(id string, doc domain.User) error {
	if m.isRecordExist(id) {
		m.store[id] = doc
		return nil
	}
	return fmt.Errorf("user does not exist for ID: %s", id)
}

func (m MapStore) GetRecordByID(id string) (domain.User, error) {
	if m.isRecordExist(id) {
		doc := m.store[id]
		return doc, nil
	}
	return domain.User{}, nil
}

func (m MapStore) GetAll() ([]domain.User, error) {
	var docs []domain.User
	for k := range m.store {
		docs = append(docs, m.store[k])
	}
	return docs, nil
}

func (m MapStore) Delete(id string) error {
	if m.isRecordExist(id) {
		delete(m.store, id)
		return nil
	}
	return fmt.Errorf("user does not exist for ID: %s", id)
}
