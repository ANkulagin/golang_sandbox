package commentary

/*
Вывод значения переменной a <- go doc returnNum или из другой директории go doc commentary.returnNum
*/
func returnNum(a int) any {
	//Возвращает значение переменной a
	return a
}

// Вывод   a <- go doc ReturnString или из другой директории go doc commentary.ReturnString
func ReturnString() any {
	return 'a'
}
