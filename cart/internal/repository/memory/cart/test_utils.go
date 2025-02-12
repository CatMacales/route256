//go:build testmode

package cart_repository

import "context"

type TestRepository interface {
	Repository
	Clear(context.Context)
}

func (r *Repository) Clear(_ context.Context) {
	clear(r.storage)
}
