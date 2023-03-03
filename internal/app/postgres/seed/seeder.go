package seed

import (
	"gorm.io/gorm"
	"log"
	"reflect"
)

type Seed struct {
	db *gorm.DB
}

func Execute(db *gorm.DB, seedMethodNames ...string) {
	s := Seed{db}

	seedType := reflect.TypeOf(s)

	// Execute all seeders if no method name is given
	if len(seedMethodNames) == 0 {
		log.Println("Running all seeder...")
		// We are looping over the method on a Seed struct
		for i := 0; i < seedType.NumMethod(); i++ {
			// Get the method in the current iteration
			method := seedType.Method(i)
			// Execute seeder
			seed(s, method.Name)
		}
	}

	// Execute only the given method names
	for _, item := range seedMethodNames {
		seed(s, item)
	}
}

func seed(s Seed, seedMethodName string) {
	m := reflect.ValueOf(s).MethodByName(seedMethodName)

	if !m.IsValid() {
		log.Fatal("No method called ", seedMethodName)
	}
	// Execute the method
	log.Println("Seeding", seedMethodName, "...")
	m.Call(nil)
	log.Println("Seed", seedMethodName, "succeeded")
}
