package main

func GetVaryingLenghtSimpleHostnames(num int) []string {
	baseName := "agentname567_that_is_long_and6578uhiii"

	hostnames := []string{}
	for i := 0; i < num; i++ {
		hostnames = append(hostnames, baseName[:(i+1)])
	}

	return hostnames
}

func GetVaryingLenghtHostnames(num int) []string {
	baseName := GetRandomString(num)

	hostnames := []string{}
	for i := 0; i < num; i++ {
		hostnames = append(hostnames, baseName[:(i+1)])
	}

	return hostnames
}
