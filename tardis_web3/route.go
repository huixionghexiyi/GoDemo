package route

import "tardis_web3/context"

type Routable interface {
	Route(method, pattern string, handlerFunc func(c context.AbstractContext))
}
