package route

import "tardis_web2/context"

type Routable interface {
	Route(method, pattern string, handlerFunc func(c context.AbstractContext))
}
