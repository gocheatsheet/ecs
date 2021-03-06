package ecs_test

import (
	"github.com/andygeiss/assert"
	"github.com/andygeiss/ecs"
	"strconv"
	"testing"
)

func TestEntityManager_Entities_Should_Have_No_Entity_At_Start(t *testing.T) {
	m := ecs.NewEntityManager()
	assert.That("manager should have no entity at start", t, len(m.Entities()), 0)
}

type mockComponent struct {
	name string
}

func (c *mockComponent) Name() string { return c.name }

func TestEntityManager_Entities_Should_Have_One_Entity_After_Adding_One_Entity(t *testing.T) {
	m := ecs.NewEntityManager()
	m.Add(&ecs.Entity{})
	assert.That("manager should have one entity", t, len(m.Entities()), 1)
}

func TestEntityManager_Entities_Should_Have_Two_Entities_After_Adding_Two_Entities(t *testing.T) {
	m := ecs.NewEntityManager()
	m.Add(&ecs.Entity{Id: "1"})
	m.Add(&ecs.Entity{Id: "2"})
	assert.That("manager should have two entities", t, len(m.Entities()), 2)
}

func TestEntityManager_Entities_Should_Have_One_Entity_After_Removing_One_Of_Two_Entities(t *testing.T) {
	m := ecs.NewEntityManager()
	e1 := &ecs.Entity{Id: "e1"}
	e2 := &ecs.Entity{Id: "e2"}
	m.Add(e1)
	m.Add(e2)
	m.Remove(e2)
	assert.That("manager should have one entity after removing one out of two", t, len(m.Entities()), 1)
	assert.That("remaining entity should have id e1", t, m.Entities()[0].Id, "e1")
}

func TestEntityManager_FilterBy_Should_Return_One_Entity_Out_Of_One(t *testing.T) {
	em := ecs.NewEntityManager()
	e := &ecs.Entity{Id: "e1", Components: []ecs.Component{
		&mockComponent{name: "position"},
	}}
	em.Add(e)
	entities := em.FilterBy("position")
	assert.That("filter should return one entity", t, len(entities), 1)
}

func TestEntityManager_FilterBy_Should_Return_One_Entity_Out_Of_Two(t *testing.T) {
	em := ecs.NewEntityManager()
	e1 := &ecs.Entity{Id: "e1", Components: []ecs.Component{
		&mockComponent{name: "position"},
	}}
	e2 := &ecs.Entity{Id: "e2", Components: []ecs.Component{
		&mockComponent{name: "velocity"},
	}}
	em.Add(e1, e2)
	entities := em.FilterBy("position")
	assert.That("filter should return one entity", t, len(entities), 1)
	assert.That("entity should be e1", t, entities[0], e1)
}

func BenchmarkEntityManager_Get_With_1_Entity_Id_Found(b *testing.B) {
	m := ecs.NewEntityManager()
	m.Add(&ecs.Entity{Id: "foo"})
	for i := 0; i < b.N; i++ {
		m.Get("foo")
	}
}

func BenchmarkEntityManager_Get_With_1000_Entities_Id_Not_Found(b *testing.B) {
	m := ecs.NewEntityManager()
	for i := 0; i < 1000; i++ {
		m.Add(&ecs.Entity{Id: strconv.Itoa(i)})
	}
	for i := 0; i < b.N; i++ {
		m.Get("1000")
	}
}
