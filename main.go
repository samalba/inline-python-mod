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

func (m *InlinePython) WithPackage(ctx context.Context, name string) *InlinePython {
	m.initBaseContainer()
	m.Ctr = m.Ctr.WithExec([]string{"pip", "install", name})
	return m
}

func (m *InlinePython) WithPackages(ctx context.Context, packages []string) *InlinePython {
	// sort the requirements to optimize caching
	slices.Sort(packages)

	for _, name := range packages {
		m.WithPackage(ctx, name)
	}

	return m
}

func (m *InlinePython) CacheKey(ctx context.Context, key string) *InlinePython {
	m.initBaseContainer()
	m.Ctr = m.Ctr.WithEnvVariable("_CACHE_KEY", key)
	return m
}

func (m *InlinePython) Code(ctx context.Context, code string) (*Container, error) {
	m.initBaseContainer()
	return m.Ctr.WithExec([]string{"python", "-c", code}).Sync(ctx)
}
