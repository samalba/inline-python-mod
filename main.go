package main

import (
	"context"
	"slices"
)

type InlinePython struct {
	Ctr *Container
}

func (m *InlinePython) initBaseContainer() {
	if m.Ctr == nil {
		m.Ctr = dag.Container().From("python:3-alpine")
	}
}

func (m *InlinePython) WithPackage(ctx context.Context, name string) (*InlinePython, error) {
	m.initBaseContainer()
	m.Ctr = m.Ctr.WithExec([]string{"pip", "install", name})
	return m, nil
}

func (m *InlinePython) WithPackages(ctx context.Context, packages []string) (*InlinePython, error) {
	// sort the requirements to optimize caching
	slices.Sort(packages)

	for _, name := range packages {
		m.WithPackage(ctx, name)
	}

	return m, nil
}

func (m *InlinePython) Code(ctx context.Context, code string) (*Container, error) {
	m.initBaseContainer()
	return m.Ctr.WithExec([]string{"python", "-c", code}).Sync(ctx)
}
