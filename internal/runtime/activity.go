package runtime

import "fmt"

type Activity interface {
	Execute(args ...any) (any, error)
}

type DebitUser struct{}

func (d *DebitUser) Execute(args ...any) (any, error) {
	// Implement the logic to debit the user
	fmt.Println("Debiting user with args:", args)
	return nil, nil
}

type ActivityRegistry struct {
	activities map[string]Activity
}

func (r *ActivityRegistry) Register(name string, activity Activity) {
	r.activities[name] = activity
}

func (r *ActivityRegistry) Get(name string) (Activity, bool) {
	activity := r.activities[name]
	if activity != nil {
		return activity, true
	}
	return nil, false
}
