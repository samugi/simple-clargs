package clargs

import (
	"fmt"
	"os"
	"strings"
	s "strings"
)

var usage string
var options []*Option
var args []string

func Init(u string, opts []*Option, a []string) {
	usage = u
	options = opts
	args = a
	usage += getUsageOptions()
}

func CheckArgs() {
	if len(args) == 0 {
		PrintUsage()
		os.Exit(1)
		return
	}
	for i, o := range options {
		parameterValue := getParameterValue(args, o.ShortOption)
		if parameterValue == "" {
			parameterValue = getParameterValue(args, o.LongOption)
		}
		if o.HasArgs && parameterValue == "" {
			PrintUsage()
			os.Exit(1)
			return
		}
		if o.Required && !findOptionInArgs(o, args) {
			PrintUsage()
			os.Exit(1)
			return
		}
		options[i].SetValue(parameterValue)
		options[i].SetValueB(findOptionInArgs(o, args))
	}
}

func PrintUsage() {
	fmt.Println(usage)
}

func getParameterValue(slice []string, parameter string) string {
	for index, par := range slice {
		if s.HasPrefix(par, parameter) {
			if len(s.Split(par, "=")) > 1 {
				return s.Split(par, "=")[1]
			}
			if index+1 < len(slice) {
				return slice[index+1]
			}
		}
	}
	return ""
}

func findOptionInArgs(opt *Option, parameters []string) bool {
	for _, par := range parameters {
		par = s.Split(par, "=")[0]

		var pars []string
		if len(par) > 2 && par[0] == '-' && par[1] != '-' { //this means it's a multiparam of the kind -abc
			pars = strings.Split(par[1:len(par)], "")
			pars = addHyphen(pars)
		} else {
			pars = append(pars, par)
		}

		for _, p := range pars {
			if opt.ShortOption == p || opt.LongOption == p {
				return true
			}
		}
	}
	return false
}

func addHyphen(parameters []string) []string {
	for i, str := range parameters {
		parameters[i] = "-" + str
	}
	return parameters
}

func getUsageOptions() string {
	retStr := "\n\nOptions\n"
	for _, opt := range options {
		optStr := opt.ShortOption + ", " + opt.LongOption
		retStr += optStr
		for i := 0; i < 30-len(optStr); i++ {
			retStr += " "
		}
		retStr += opt.Description
		retStr += "\n"
	}
	return retStr
}
