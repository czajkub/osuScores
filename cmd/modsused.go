package cmd

import (
	"fmt"
	"sort"
)

func Usedmods(scores []Score) {
	modmap := make(map[string]int)
	modname := make(map[string]string)
	modname = setdict(modname)
	for _, score := range scores {
		mods := score.Mods
		for _, mod := range mods {
			modmap[mod.Modacronym] += 1
		}
	}
	keys := make([]string, 0, len(modmap))
	for key := range modmap {
		keys = append(keys, key)
	}
	sort.Slice(keys, func(i, j int) bool {
		return modmap[keys[i]] > modmap[keys[j]]
	})
	fmt.Printf("Most commonly used mods:\n")
	for _, key := range keys {
		fullmod, ok := modname[key]
		if ok {
			fmt.Printf("%s: %d\n", fullmod, modmap[key])
		} else {
			fmt.Printf("%s: %d\n", key, modmap[key])
		}

	}
}

func setdict(modname map[string]string) map[string]string {
	modname["CL"] = "Nomod"
	modname["HD"] = "Hidden"
	modname["DT"] = "Double Time"
	modname["HR"] = "Hard Rock"
	modname["NF"] = "NoFail"
	modname["NC"] = "Nightcore"
	modname["EZ"] = "Easy"
	modname["FL"] = "Flashlight"
	modname["TD"] = "Touchscreen device"
	modname["RX"] = "Relax"
	modname["HT"] = "Half Time"
	modname["SD"] = "Sudden Death"
	modname["PF"] = "Perfect"
	modname["AP"] = "Autopilot"
	modname["SO"] = "Spinout"
	return modname
}
