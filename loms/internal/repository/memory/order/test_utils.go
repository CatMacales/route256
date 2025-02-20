//go:build testmode

package order_repository

import "context"

type TestRepository interface {
	Repository
	Clear(context.Context)
}

func (r *Repository) Clear(_ context.Context) {
	clear(r.storage)
}
