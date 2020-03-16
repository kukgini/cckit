package testdata

import (
	"github.com/kukgini/cckit2/extensions/debug"
	"github.com/kukgini/cckit2/extensions/owner"
	"github.com/kukgini/cckit2/router"
	"github.com/kukgini/cckit2/router/param/defparam"
	m "github.com/kukgini/cckit2/state/mapping"
	"github.com/kukgini/cckit2/state/mapping/testdata/schema"
)

func NewComplexIdCC() *router.Chaincode {
	r := router.New(`complexId`)
	debug.AddHandlers(r, `debug`, owner.Only)

	// Mappings for chaincode state
	r.Use(m.MapStates(m.StateMappings{}.
		//key will be <`EntityWithComplexId`, {Id.IdPart1}, {Id.IdPart2} >
		Add(&schema.EntityWithComplexId{}, m.PKeyComplexId(&schema.EntityComplexId{}))))

	r.Init(owner.InvokeSetFromCreator)

	r.Group(`entity`).
		Invoke(`List`, entityList).
		Invoke(`Get`, entityGet, defparam.Proto(&schema.EntityComplexId{})).
		Invoke(`Insert`, entityInsert, defparam.Proto(&schema.EntityWithComplexId{}))

	return router.NewChaincode(r)
}

func entityList(c router.Context) (interface{}, error) {
	return c.State().List(&schema.EntityWithComplexId{})
}

func entityInsert(c router.Context) (interface{}, error) {
	var (
		entity = c.Param()
	)
	return entity, c.State().Insert(entity)
}

func entityGet(c router.Context) (interface{}, error) {
	var (
		id = c.Param()
	)
	return c.State().Get(id)
}
