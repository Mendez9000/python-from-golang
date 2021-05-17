package main

import (
	"fmt"
	"github.com/sbinet/go-python"
	"log"
)

func init() {
	err := python.Initialize()
	if err != nil {
		panic(err.Error())
	}
}

// Posiblemente sea necesario establecer PYTHONPATH:
// un a forma es ejecutarlo desde dentro de la carpeta que contiene este archivo:
// export PYTHONPATH=./

func main() {
	fmt.Printf("importing ...\n")
	m := python.PyImport_ImportModule("go-python")
	if m == nil {
		log.Fatalf("could not import 'modoe'. Use command export PYTHONPATH=./\n")
	}

	foo := m.GetAttrString("foo")
	if foo == nil {
		log.Fatalf("could not getattr(kwargs, 'foo')\n")
	}

	// keyword arguments
	kw := python.PyDict_New()
	err := python.PyDict_SetItem(
		kw,
		python.PyString_FromString("Fernando Perez"),
		python.PyInt_FromLong(327524),
	)
	if err != nil {
		log.Fatalf("error: %v\n", err)
	}

	args := python.PyList_New(0)
	out := foo.Call(args, kw)
	if out == nil {
		log.Fatalf("%s\n", out)
	}

	str := python.PyString_AsString(out)
	fmt.Printf("%s\n", str)
}
