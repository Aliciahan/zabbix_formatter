package zabbix_formatter

func AddUnique(listStr *[]string, newStr string) {
	addOrNot := true

	if len(*listStr) == 0 {
		addOrNot = true
	} else {
		for _,v := range *listStr {
			if v == newStr {
				addOrNot = false
			}
		}
	}
	if addOrNot == true {
		*listStr = append(*listStr,newStr)
	}
}

func FormatterDiscover(detectName string, dataSlice []string) string {
	result := `{"data": [`
	for i:=0; i< len(dataSlice)-1; i++ {
		result = result + `{"{#`+ detectName + `}": ` +`"`+ dataSlice[i] +`"` +    `},`
	}
	result = result + `{"{#`+ detectName + `}": ` +`"`+ dataSlice[len(dataSlice)-1] +`"` +`}`
	return result + `]}`
}
