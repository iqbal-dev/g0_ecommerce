package user

import (
	domain "ecommerce/internal/domain/user"
	"errors"
	"reflect"
	"testing"
)

type fakeUserRepo struct {
	createArg      domain.User
	updateArgID    int
	updateArg      domain.User
	deleteArgID    int
	findOneArgID   int
	findByEmail    string
	findByPassword string

	createRes      *domain.User
	updateRes      *domain.User
	deleteRes      bool
	findAllRes     []*domain.User
	findOneRes     *domain.User
	findByEmailRes *domain.User

	err error
}

func (f *fakeUserRepo) Create(u domain.User) (*domain.User, error) {
	f.createArg = u
	return f.createRes, f.err
}

func (f *fakeUserRepo) Update(id int, u domain.User) (*domain.User, error) {
	f.updateArgID = id
	f.updateArg = u
	return f.updateRes, f.err
}

func (f *fakeUserRepo) Delete(id int) (bool, error) {
	f.deleteArgID = id
	return f.deleteRes, f.err
}

func (f *fakeUserRepo) FindAll() ([]*domain.User, error) {
	return f.findAllRes, f.err
}

func (f *fakeUserRepo) FindOne(id int) (*domain.User, error) {
	f.findOneArgID = id
	return f.findOneRes, f.err
}

func (f *fakeUserRepo) FindByEmailAndPassword(email string, password string) (*domain.User, error) {
	f.findByEmail = email
	f.findByPassword = password
	return f.findByEmailRes, f.err
}

func TestUserServiceDelegates(t *testing.T) {
	repo := &fakeUserRepo{
		createRes:      &domain.User{Id: 1, Name: "created"},
		updateRes:      &domain.User{Id: 2, Name: "updated"},
		deleteRes:      true,
		findAllRes:     []*domain.User{{Id: 3}},
		findOneRes:     &domain.User{Id: 4},
		findByEmailRes: &domain.User{Id: 5},
	}
	svc := NewService(repo)

	if _, err := svc.Create(domain.User{Name: "n"}); err != nil {
		t.Fatalf("Create error: %v", err)
	}
	if repo.createArg.Name != "n" {
		t.Fatalf("Create not delegated")
	}

	if _, err := svc.Update(9, domain.User{Name: "u"}); err != nil {
		t.Fatalf("Update error: %v", err)
	}
	if repo.updateArgID != 9 || repo.updateArg.Name != "u" {
		t.Fatalf("Update args wrong")
	}

	if _, err := svc.Delete(7); err != nil {
		t.Fatalf("Delete error: %v", err)
	}
	if repo.deleteArgID != 7 {
		t.Fatalf("Delete arg wrong")
	}

	if _, err := svc.FindAll(); err != nil {
		t.Fatalf("FindAll error: %v", err)
	}

	if _, err := svc.FindOne(11); err != nil {
		t.Fatalf("FindOne error: %v", err)
	}
	if repo.findOneArgID != 11 {
		t.Fatalf("FindOne arg wrong")
	}

	if _, err := svc.FindByEmailAndPassword("a@b.c", "pw"); err != nil {
		t.Fatalf("FindByEmailAndPassword error: %v", err)
	}
	if repo.findByEmail != "a@b.c" || repo.findByPassword != "pw" {
		t.Fatalf("FindByEmailAndPassword args wrong")
	}
}

func TestUserServicePropagatesErrors(t *testing.T) {
	expectedErr := errors.New("boom")
	repo := &fakeUserRepo{err: expectedErr}
	svc := NewService(repo)

	if _, err := svc.Create(domain.User{}); !errors.Is(err, expectedErr) {
		t.Fatalf("Create error not propagated: %v", err)
	}
	if _, err := svc.Update(0, domain.User{}); !errors.Is(err, expectedErr) {
		t.Fatalf("Update error not propagated: %v", err)
	}
	if _, err := svc.Delete(0); !errors.Is(err, expectedErr) {
		t.Fatalf("Delete error not propagated: %v", err)
	}
	if _, err := svc.FindAll(); !errors.Is(err, expectedErr) {
		t.Fatalf("FindAll error not propagated: %v", err)
	}
	if _, err := svc.FindOne(0); !errors.Is(err, expectedErr) {
		t.Fatalf("FindOne error not propagated: %v", err)
	}
	if _, err := svc.FindByEmailAndPassword("", ""); !errors.Is(err, expectedErr) {
		t.Fatalf("FindByEmailAndPassword error not propagated: %v", err)
	}
}

func TestUserServiceReturnsRepoValues(t *testing.T) {
	repo := &fakeUserRepo{
		createRes:      &domain.User{Id: 1},
		updateRes:      &domain.User{Id: 2},
		deleteRes:      true,
		findAllRes:     []*domain.User{{Id: 3}},
		findOneRes:     &domain.User{Id: 4},
		findByEmailRes: &domain.User{Id: 5},
	}
	svc := NewService(repo)

	if got, _ := svc.Create(domain.User{}); !reflect.DeepEqual(got, repo.createRes) {
		t.Fatalf("Create mismatch")
	}
	if got, _ := svc.Update(0, domain.User{}); !reflect.DeepEqual(got, repo.updateRes) {
		t.Fatalf("Update mismatch")
	}
	if got, _ := svc.Delete(0); got != repo.deleteRes {
		t.Fatalf("Delete mismatch")
	}
	if got, _ := svc.FindAll(); !reflect.DeepEqual(got, repo.findAllRes) {
		t.Fatalf("FindAll mismatch")
	}
	if got, _ := svc.FindOne(0); !reflect.DeepEqual(got, repo.findOneRes) {
		t.Fatalf("FindOne mismatch")
	}
	if got, _ := svc.FindByEmailAndPassword("", ""); !reflect.DeepEqual(got, repo.findByEmailRes) {
		t.Fatalf("FindByEmailAndPassword mismatch")
	}
}
