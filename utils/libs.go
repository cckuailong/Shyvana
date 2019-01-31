package utils

func StrIsIn(item string, list []string)bool{
	for _, c := range(list){
		if c == item{
			return true
		}
	}
	return false
}
