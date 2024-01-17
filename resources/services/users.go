package services

import (
	"context"
	"encoding/json"
	"fmt"
	"net/url"
	"time"

	"github.com/cloudquery/plugin-sdk/v4/schema"
	"github.com/cloudquery/plugin-sdk/v4/transformers"
	"github.com/razin99/cq-source-knowbe4/client"
)

func UsersTable() *schema.Table {
	return &schema.Table{
		Name:      "knowbe4_users",
		Resolver:  fetchUsers,
		Transform: transformers.TransformWithStruct(&User{}, transformers.WithPrimaryKeys("Id")),
	}
}

type User struct {
	Id                   uint64    `json:"id"`
	FirstName            string    `json:"first_name"`
	LastName             string    `json:"last_name"`
	Email                string    `json:"email"`
	PhishPronePercentage float64   `json:"phish_prone_percentage"`
	Groups               []uint64  `json:"groups"`
	CurrentRiskScore     float64   `json:"current_risk_score"`
	JoinedOn             time.Time `json:"joined_on"`
	LastSignIn           time.Time `json:"last_sign_in"`
}

func fetchUsers(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)

	var page uint64 = 1
	path := "/v1/users"
	query := url.Values{}
	query.Set("status", "active")

	for {
		query.Set("page", fmt.Sprint(page))
		response, err := cl.KnowBe4.NewRequest(ctx, path, query)
		if err != nil {
			return err
		}

		var users []User
		err = func() error {
			defer response.Body.Close()
			return json.NewDecoder(response.Body).Decode(&users)
		}()
		if err != nil {
			return err
		}

		if len(users) == 0 {
			break
		}

		for _, enrollment := range users {
			res <- enrollment
		}

		page++
	}

	return nil
}
