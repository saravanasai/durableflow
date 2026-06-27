package runtime

type Workflow interface {
	Execute(ctx *Context, input any) error
}

type PurchaseWorkflow struct{}

func (p *PurchaseWorkflow) Execute(ctx *Context, input any) error {

	ctx.Activity("debit", input)

	return nil
}

type WorkflowRegistry struct {
	workflows map[string]Workflow
}

func (r *WorkflowRegistry) Register(name string, workflow Workflow) {
	r.workflows[name] = workflow
}

func (r *WorkflowRegistry) Get(name string) (Workflow, bool) {
	flow := r.workflows[name]
	if flow != nil {
		return flow, true
	}

	return nil, false
}
