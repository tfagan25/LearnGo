package main

import (
	"fmt"
)

type Person struct {
	name string
	age int
	job string
}

type PersonBuilder struct {
	person *Person
}

func (p *PersonBuilder) setName(name string) *PersonBuilder {
	p.person.name = name
	return p
}

func (p *PersonBuilder) setAge(age int) *PersonBuilder {
	p.person.age = age
	return p
}

func (p *PersonBuilder) setJob(job string) *PersonBuilder {
	p.person.job = job
	return p
}

func main () {
	builder := &PersonBuilder{person: &Person{}}

	builder.
	setAge(25).
	setName("John")//.
	//setJob("MANGO") => In this job market??

	fmt.Println(builder.person)
}