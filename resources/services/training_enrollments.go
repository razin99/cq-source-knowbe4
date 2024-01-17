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

func TrainingEnrollmentsTable() *schema.Table {
	return &schema.Table{
		Name:      "knowbe4_training_enrollments",
		Resolver:  fetchTrainingEnrollments,
		Transform: transformers.TransformWithStruct(&TrainingEnrollment{}, transformers.WithPrimaryKeys("EnrollmentID")),
	}
}

type TrainingEnrollment struct {
	EnrollmentID uint64 `json:"enrollment_id"`
	ContentType  string `json:"content_type"`
	ModuleName   string `json:"module_name"`
	User         struct {
		ID        uint64 `json:"id"`
		FirstName string `json:"first_name"`
		LastName  string `json:"last_name"`
		Email     string `json:"email"`
	} `json:"user"`
	CampaignName       string    `json:"campaign_name"`
	EnrollmentDate     time.Time `json:"enrollment_date"`
	StartDate          time.Time `json:"start_date"`
	CompletionDate     time.Time `json:"completion_date"`
	Status             string    `json:"status"`
	TimeSpent          uint64    `json:"time_spent"`
	PolicyAcknowledged bool      `json:"policy_acknowledged"`
}

func fetchTrainingEnrollments(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- any) error {
	cl := meta.(*client.Client)

	var page uint64 = 1
	path := "/v1/training/enrollments"
	query := url.Values{}

	for {
		query.Set("page", fmt.Sprint(page))
		response, err := cl.KnowBe4.NewRequest(ctx, path, query)
		if err != nil {
			return err
		}

		var enrollmentsData []TrainingEnrollment
		err = func() error {
			defer response.Body.Close()
			return json.NewDecoder(response.Body).Decode(&enrollmentsData)
		}()
		if err != nil {
			return err
		}

		if len(enrollmentsData) == 0 {
			break
		}

		for _, enrollment := range enrollmentsData {
			res <- enrollment
		}

		page++
	}

	return nil
}
