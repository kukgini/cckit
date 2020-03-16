package defparam

import (
	"github.com/kukgini/cckit2/router"
	"github.com/kukgini/cckit2/router/param"
)

func Proto(target interface{}, argPoss ...int) router.MiddlewareFunc {
	return param.Proto(router.DefaultParam, target, argPoss...)
}
