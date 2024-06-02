package main

import "time"

// Object Methods
type DataContainer struct {
	ID         string    `json:"id"`
	Index      int       `json:"index"`
	SomeBool   bool      `json:"someBool"`
	SomeFloat  float64   `json:"someFloat"`
	SomeInt    int       `json:"someInt"`
	FirstName  string    `json:"firstName"`
	LastName   string    `json:"lastName"`
	SomeString string    `json:"someString"`
	SomeDate   time.Time `json:"someDate"`
}

var data = DataContainer{
	ID:         "0123456789abcdefghijklmnopqrstuvwxyz",
	Index:      0,
	SomeBool:   false,
	SomeFloat:  2238.57,
	SomeInt:    3,
	FirstName:  "FName",
	LastName:   "LName",
	SomeString: "Pariatur sunt eu tempor mollit sint sint sint non. Duis duis anim minim ad velit. Magna nostrud id amet consequat. Irure laboris tempor culpa mollit est ut dolor cupidatat elit ipsum proident ad pariatur tempor.\r\nAliquip voluptate officia deserunt ex tempor. Aliquip non non officia commodo consequat non sit qui minim sunt reprehenderit laborum esse. Ut id elit exercitation ad reprehenderit qui. Anim et reprehenderit culpa ipsum ad nulla. Dolor eiusmod veniam velit consectetur commodo. Aliqua nulla ullamco laborum exercitation aute reprehenderit do sunt et labore incididunt proident sunt.\r\n",
	SomeDate:   time.Date(2024, 05, 24, 8, 58, 50, 0, time.UTC),
}

func returnData() DataContainer {
	return data
}

func returnDataPtr() *DataContainer {
	return &data
}
