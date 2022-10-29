package main

func countMatches(items [][]string, ruleKey string, ruleValue string) int {
	typeMap := make(map[string]int)
	colorMap := make(map[string]int)
	nameMap := make(map[string]int)
	for _, item := range items {
		_type := item[0]
		color := item[1]
		name := item[2]
		typeCount, ok := typeMap[_type]
		if !ok {
			typeCount = 0
		}
		colorCount, ok := colorMap[color]
		if !ok {
			colorCount = 0
		}
		nameCount, ok := nameMap[name]
		if !ok {
			nameCount = 0
		}
		typeMap[_type] = typeCount + 1
		colorMap[color] = colorCount + 1
		nameMap[name] = nameCount + 1
	}
	switch ruleKey {
	case "type":
		return typeMap[ruleValue]
	case "color":
		return colorMap[ruleValue]
	case "name":
		return nameMap[ruleValue]
	}
	return 0
}
