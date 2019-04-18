package util


func CheckErr(err error){
	if err != nil {
		panic(err)
	}
}

func ContainsZero(a[] int64) bool {
	for i:=range a{
		if(a[i]!=0){
			continue
		}else {
			return true //含有0
		}
	}
	return false	//不含有0
}