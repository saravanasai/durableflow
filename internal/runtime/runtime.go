package runtime

import (
	"fmt"
	"time"
)

type Context struct {
	runtime *Runtime
}

func (c *Context) Activity(name string, args ...any) (any, error) {
	activity, ok := c.runtime.activities.Get(name)
	if !ok {
		return nil, fmt.Errorf("activity not found: %s", name)
	}
	return activity.Execute(args...)
}

func (c *Context) Sleep(duration time.Duration) error {
	time.Sleep(duration)
	return nil
}

type Runtime struct {
	workflows  *WorkflowRegistry
	activities *ActivityRegistry
}

func NewRuntime() *Runtime {
	return &Runtime{
		workflows:  &WorkflowRegistry{workflows: make(map[string]Workflow)},
		activities: &ActivityRegistry{activities: make(map[string]Activity)},
	}
}

func (r *Runtime) RegisterWorkflow(name string, workflow Workflow) {
	r.workflows.Register(name, workflow)
}

func (r *Runtime) RegisterActivity(name string, activity Activity) {
	r.activities.Register(name, activity)
}

func (r *Runtime) Start(workflow string, input any) error {
	flow, ok := r.workflows.Get(workflow)
	if !ok {
		return fmt.Errorf("workflow not found: %s", workflow)
	}

	return flow.Execute(&Context{runtime: r}, input)
}
