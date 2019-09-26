package main

import (
	"fmt"
	mapset "github.com/deckarep/golang-set"
)

func main() {

	/*
		groups := []interface{}{"SGT", "CBS", "RST", "Game", "Web", "SCAS", "SBS", "SBOS", "ADS", "SPRS", "BDP", "Pay",
			"UGS", "SLS", "WeShow", "Co", "DA", "SI", "UO", "CO"}
	*/
	tags := []interface{}{"env", "group", "project"}

	envs := mapset.NewSet()
	envs.Add("prod")
	envs.Add("test")
	envs.Add("pre")

	env := "alpha"
	if envs.Contains(env) {
		fmt.Println("包含")
	} else {
		fmt.Println("不包含")
	}

	currentTags := mapset.NewSetFromSlice([]interface{}{"env", "group", "project", "panda"})
	currentEnvs := mapset.NewSetFromSlice([]interface{}{"prod", "test", "alpha"})
	// envs 对 currentEnvs 的差集
	dset := envs.Difference(currentEnvs)

	is := currentTags.IsSuperset(mapset.NewSetFromSlice(tags))
	isSubset := currentEnvs.IsSubset(envs)

	fmt.Println(isSubset)

	fmt.Println(dset.ToSlice())

	fmt.Println(envs.Intersect(currentEnvs))

	fmt.Println(envs.Cardinality())

	fmt.Println(envs.Union(currentEnvs))

	envs.Remove("pres")
	fmt.Println(envs)

	fmt.Println(is)
}
