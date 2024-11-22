package stub

import (
	"errors"
	"fmt"
	"testing"

	"github.com/google/go-cmp/cmp"
)

type GetPetNamesStub struct {
	Entities
}

func (ps GetPetNamesStub) GetPets(userID string) ([]Pet, error) {
	switch userID {
	case "1":
		return []Pet{{Name: "Bubbles"}}, nil
	case "2":
		return []Pet{{Name: "Stampy"}, {Name: "Snowball II"}}, nil
	default:
		return nil, fmt.Errorf("invalid id: %s", userID)
	}
}

func TestGetPetName(t *testing.T) {
	data := []struct {
		userID string
		err    error
		result []string
	}{
		{"1", nil, []string{"Bubbles"}},
		{"2", nil, []string{"Stampy", "Snowball II"}},
		{"3", fmt.Errorf("invalid id: 3"), nil},
	}

	l := Logic{Entities: GetPetNamesStub{}}

	for _, d := range data {
		t.Run(d.userID, func(t *testing.T) {
			result, err := l.GetPetNames(d.userID)

			fmt.Println(result)
			if err != nil && err.Error() != d.err.Error() {
				t.Error(err)
			}

			if diff := cmp.Diff(d.result, result); diff != "" {
				t.Error(diff)
			}

		})

	}
}

type EntitiesStub struct {
	getUser     func(id string) (User, error)
	getPets     func(userID string) ([]Pet, error)
	getChildren func(userID string) ([]Person, error)
	getFriends  func(userID string) ([]Person, error)
	saveUser    func(user User) error
}

func (es EntitiesStub) GetPets(userID string) ([]Pet, error) {
	return es.getPets(userID)
}

func (es EntitiesStub) GetUser(userId string) (User, error) {
	return es.getUser(userId)
}

func (es EntitiesStub) GetChildren(userID string) ([]Person, error) {
	return es.getChildren(userID)
}

func (es EntitiesStub) GetFriends(userID string) ([]Person, error) {
	return es.getFriends(userID)
}

func (es EntitiesStub) SaveUser(user User) error {
	return es.saveUser(user)
}

func TestGetPetNamesComprehensive(t *testing.T) {
	data := []struct {
		name     string
		getPets  func(userID string) ([]Pet, error)
		userID   string
		petNames []string
		errMsg   string
	}{
		{"case1", func(userID string) ([]Pet, error) {
			return []Pet{{Name: "Bubbles"}}, nil
		}, "1", []string{"Bubbles"}, ""},
		{"case2", func(userID string) ([]Pet, error) {
			return nil, errors.New("invalid id: 3")
		}, "3", nil, "invalid id: 3"},
	}

	l := Logic{}
	for _, d := range data {
		t.Run(d.name, func(t *testing.T) {
			l.Entities = EntitiesStub{getPets: d.getPets}

			petNames, err := l.GetPetNames(d.userID)

			if err != nil && err.Error() != d.errMsg {
				t.Error(err)
			}

			if diff := cmp.Diff(d.petNames, petNames); diff != "" {
				t.Error(diff)
			}
		})
	}
}
