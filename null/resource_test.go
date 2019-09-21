package null

import (
	"reflect"
	"testing"
)

func TestResource(t *testing.T) {
	wd := testHelper.RequireNewWorkingDir(t)
	defer wd.Close()

	wd.RequireSetConfig(t, `
resource "null_resource" "test" {
  triggers = {
    a = 1
  }
}
`)

	wd.RequireInit(t)
	defer wd.RequireDestroy(t)
	wd.RequireApply(t)

	state := wd.RequireState(t)
	if got, want := len(state.Values.RootModule.Resources), 1; got != want {
		t.Fatalf("wrong number of resource instance objects in state %d; want %d", got, want)
	}
	instanceState := state.Values.RootModule.Resources[0]
	if got, want := instanceState.Type, "null_resource"; got != want {
		t.Errorf("wrong resource type in state\ngot:  %s\nwant: %s", got, want)
	}
	if got, want := instanceState.AttributeValues["triggers"], map[string]interface{}{"a": "1"}; !reflect.DeepEqual(got, want) {
		t.Errorf("wrong 'triggers' value\ngot:  %#v (%T)\nwant: %#v (%T)", got, got, want, want)
	}

	// Now we'll plan without changing anything. That should produce an empty plan.
	wd.RequireCreatePlan(t)
	plan := wd.RequireSavedPlan(t)
	if got, want := len(plan.ResourceChanges), 1; got != want {
		t.Fatalf("wrong number of resource changes in plan %d; want %d", got, want)
	}
	instanceChange := plan.ResourceChanges[0]
	if got, want := instanceChange.Type, "null_resource"; got != want {
		t.Errorf("wrong resource type in plan\ngot:  %s\nwant: %s", got, want)
	}
	if !instanceChange.Change.Actions.NoOp() {
		t.Errorf("wrong action in plan\ngot:  %#v\nwant no-op", instanceChange.Change.Actions)
	}

	// No we'll change the triggers, which should cause null_resource to show
	// as needing replacement in the plan.
	wd.RequireSetConfig(t, `
resource "null_resource" "test" {
  triggers = {
    a = 2
  }
}
`)
	wd.RequireCreatePlan(t)
	plan = wd.RequireSavedPlan(t)
	if got, want := len(plan.ResourceChanges), 1; got != want {
		t.Fatalf("wrong number of resource changes in plan %d; want %d", got, want)
	}
	instanceChange = plan.ResourceChanges[0]
	if got, want := instanceChange.Type, "null_resource"; got != want {
		t.Errorf("wrong resource type in plan\ngot:  %s\nwant: %s", got, want)
	}
	if !instanceChange.Change.Actions.DestroyBeforeCreate() {
		t.Errorf("wrong action in plan\ngot:  %#v\nwant destroy then create (replace)", instanceChange.Change.Actions)
	}

	// For good measure, we'll apply that saved plan to make sure the destroy
	// works (should be a no-op, but successful).
	wd.RequireApply(t)
}
