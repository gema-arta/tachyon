package tachyon

import (
	"testing"
)

func TestSimplePlaybook(t *testing.T) {
	env := &Environment{Vars: NewNestedScope(nil)}
	p, err := LoadPlaybook("test/playbook1.yml", env.Vars, env)

	if err != nil {
		panic(err)
	}

	if len(p.Plays) != 1 {
		t.Fatalf("Didn't load 1 playbook, loaded: %d", len(p.Plays))
	}

	x := p.Plays[0]

	if x.Hosts != "all" {
		t.Errorf("Hosts not all: was %s", x.Hosts)
	}

	vars := x.Vars

	a, ok := vars.Get("answer")

	if !ok {
		t.Fatalf("No var 'answer'")
	}

	if a != "Wuh, I think so" {
		t.Errorf("Unable to decode string var: %#v", a)
	}

	a, ok = vars.Get("port")

	if !ok {
		t.Fatalf("No var 'port'")
	}

	if a != 5150 {
		t.Errorf("Unable to decode numeric var: %#v", a)
	}

	if len(x.VarsFiles) != 2 {
		t.Fatalf("Unable to decode varsfiles, got %d", len(x.VarsFiles))
	}

	f := x.VarsFiles[0]

	if f != "common_vars.yml" {
		t.Errorf("Unable to decode literal vars_files")
	}

	f2 := x.VarsFiles[1].([]interface{})

	if f2[1].(string) != "default_os.yml" {
		t.Errorf("Unable to decode list vars_files")
	}

	tasks := x.Tasks

	if len(tasks) < 5 {
		t.Errorf("Failed to decode the proper number of tasks: %d", len(tasks))
	}

	if tasks[3].Args() != "echo {{port}}" {
		t.Errorf("Failed to decode templating in action: %#v", tasks[3].Args())
	}
}
