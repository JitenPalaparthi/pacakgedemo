package types

import "fmt"

// int is a buitin type or defined type
// methods can be implemented not only on structs but also on any userdefiend type

//ToString medhod on int
type MyInt int

// assignment to do on wekend
// type Umap map[interface{}]interface{}
// func(um Umap) Insert(key,value)
// func(um Umap) Update(key,value)(error) // IF the key does not exist or map is nil return error
// func (um Umap) GetAll()(string)
// func (um Umapo)Delete(key) // if the key does not exist or map is nil then return error
// func (um Umap)GetKeys()[]interface{}
// func (um Umap)GetValues()[]interface{}

func (mi MyInt) ToString() string {
	return fmt.Sprint(mi)
}
