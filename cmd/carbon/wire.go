//go:build wireinject
// +build wireinject

package carbon

import "github.com/google/wire"

func wireF() (func(), error) {
	panic(wire.Bind(nil, nil))
}
