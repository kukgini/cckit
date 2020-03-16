package testdata

import (
	"github.com/kukgini/cckit2/extensions/debug"
	"github.com/kukgini/cckit2/extensions/owner"
	"github.com/kukgini/cckit2/router"
	"github.com/kukgini/cckit2/router/param"
	"github.com/kukgini/cckit2/router/param/defparam"
	"github.com/kukgini/cckit2/state"
	m "github.com/kukgini/cckit2/state/mapping"
	"github.com/kukgini/cckit2/state/mapping/testdata/schema"
)

func NewSliceIdCC() *router.Chaincode {
	r := router.New(`complexId`)
	debug.AddHandlers(r, `debug`, owner.Only)

	// Mappings for chaincode state
	r.Use(m.MapStates(m.StateMappings{}.
		//key will be <`EntityWithSliceId`, {Id[0]}, {Id[1]},... {Id[len(Id)-1]} >
		Add(&schema.EntityWithSliceId{}, m.PKeyId())))

	r.Init(owner.InvokeSetFromCreator)

	r.Group(`entity`).
		Invoke(`List`, func(c router.Context) (interface{}, error) {
			return c.State().List(&schema.EntityWithSliceId{})
		}).
		Invoke(`Get`, func(c router.Context) (interface{}, error) {
			return c.State().Get(&schema.EntityWithSliceId{Id: state.StringsIdFromStr(c.ParamString(`Id`))})
		}, param.String(`Id`)).
		Invoke(`Insert`, func(c router.Context) (interface{}, error) {
			return nil, c.State().Insert(c.Param())
		}, defparam.Proto(&schema.EntityWithSliceId{}))

	return router.NewChaincode(r)
}
