package product

import (
	domain "ecommerce/internal/domain/product"
	"errors"
	"reflect"
	"testing"
)

type fakeProductRepo struct {
	createArg        domain.Product
	updateArgID      int
	updateArgProduct domain.Product
	deleteArgID      int
	findAllPage      int64
	findAllLimit     int64
	findOneArgID     int

	createRes    *domain.Product
	updateRes    *domain.Product
	deleteRes    bool
	findAllRes   []*domain.Product
	findAllTotal int64
	findOneRes   *domain.Product

	err error
}

func (f *fakeProductRepo) Create(p domain.Product) (*domain.Product, error) {
	f.createArg = p
	return f.createRes, f.err
}

func (f *fakeProductRepo) Update(id int, p domain.Product) (*domain.Product, error) {
	f.updateArgID = id
	f.updateArgProduct = p
	return f.updateRes, f.err
}

func (f *fakeProductRepo) Delete(id int) (bool, error) {
	f.deleteArgID = id
	return f.deleteRes, f.err
}

func (f *fakeProductRepo) FindAll(page, limit int64) ([]*domain.Product, int64, error) {
	f.findAllPage = page
	f.findAllLimit = limit
	return f.findAllRes, f.findAllTotal, f.err
}

func (f *fakeProductRepo) FindOne(id int) (*domain.Product, error) {
	f.findOneArgID = id
	return f.findOneRes, f.err
}

func TestServiceDelegatesToRepository(t *testing.T) {
	repo := &fakeProductRepo{
		createRes:    &domain.Product{Id: 1, Name: "created"},
		updateRes:    &domain.Product{Id: 2, Name: "updated"},
		deleteRes:    true,
		findAllRes:   []*domain.Product{{Id: 3}},
		findAllTotal: 1,
		findOneRes:   &domain.Product{Id: 4, Name: "found"},
	}
	svc := NewService(repo)

	if _, err := svc.Create(domain.Product{Name: "n"}); err != nil {
		t.Fatalf("Create returned error: %v", err)
	}
	if repo.createArg.Name != "n" {
		t.Fatalf("Create not delegated: got %s", repo.createArg.Name)
	}

	if _, err := svc.Update(10, domain.Product{Name: "u"}); err != nil {
		t.Fatalf("Update returned error: %v", err)
	}
	if repo.updateArgID != 10 || repo.updateArgProduct.Name != "u" {
		t.Fatalf("Update not delegated correctly: id %d name %s", repo.updateArgID, repo.updateArgProduct.Name)
	}

	if _, err := svc.Delete(5); err != nil {
		t.Fatalf("Delete returned error: %v", err)
	}
	if repo.deleteArgID != 5 {
		t.Fatalf("Delete not delegated: got %d", repo.deleteArgID)
	}

	products, total, err := svc.FindAll(2, 20)
	if err != nil {
		t.Fatalf("FindAll returned error: %v", err)
	}
	if repo.findAllPage != 2 || repo.findAllLimit != 20 {
		t.Fatalf("FindAll args wrong: page %d limit %d", repo.findAllPage, repo.findAllLimit)
	}
	if total != 1 || len(products) != 1 || products[0].Id != 3 {
		t.Fatalf("FindAll results unexpected: total %d products %+v", total, products)
	}

	p, err := svc.FindOne(7)
	if err != nil {
		t.Fatalf("FindOne returned error: %v", err)
	}
	if repo.findOneArgID != 7 || p.Id != 4 {
		t.Fatalf("FindOne mismatch: id %d product %+v", repo.findOneArgID, p)
	}
}

func TestServicePropagatesErrors(t *testing.T) {
	expectedErr := errors.New("boom")
	repo := &fakeProductRepo{err: expectedErr}
	svc := NewService(repo)

	if _, err := svc.Create(domain.Product{}); !errors.Is(err, expectedErr) {
		t.Fatalf("Create error not propagated: %v", err)
	}
	if _, err := svc.Update(1, domain.Product{}); !errors.Is(err, expectedErr) {
		t.Fatalf("Update error not propagated: %v", err)
	}
	if _, err := svc.Delete(1); !errors.Is(err, expectedErr) {
		t.Fatalf("Delete error not propagated: %v", err)
	}
	if _, _, err := svc.FindAll(1, 1); !errors.Is(err, expectedErr) {
		t.Fatalf("FindAll error not propagated: %v", err)
	}
	if _, err := svc.FindOne(1); !errors.Is(err, expectedErr) {
		t.Fatalf("FindOne error not propagated: %v", err)
	}
}

func TestServiceUsesRepositoryReturnValues(t *testing.T) {
	repo := &fakeProductRepo{
		createRes:    &domain.Product{Id: 1},
		updateRes:    &domain.Product{Id: 2},
		deleteRes:    true,
		findAllRes:   []*domain.Product{{Id: 3}},
		findAllTotal: 99,
		findOneRes:   &domain.Product{Id: 4},
	}
	svc := NewService(repo)

	if got, _ := svc.Create(domain.Product{}); !reflect.DeepEqual(got, repo.createRes) {
		t.Fatalf("Create result mismatch")
	}
	if got, _ := svc.Update(0, domain.Product{}); !reflect.DeepEqual(got, repo.updateRes) {
		t.Fatalf("Update result mismatch")
	}
	if got, _ := svc.Delete(0); got != repo.deleteRes {
		t.Fatalf("Delete result mismatch")
	}
	if res, total, _ := svc.FindAll(1, 1); total != repo.findAllTotal || !reflect.DeepEqual(res, repo.findAllRes) {
		t.Fatalf("FindAll result mismatch")
	}
	if got, _ := svc.FindOne(0); !reflect.DeepEqual(got, repo.findOneRes) {
		t.Fatalf("FindOne result mismatch")
	}
}
