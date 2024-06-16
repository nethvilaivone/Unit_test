package services

// import "golang.org/x/text/cases"
func CheckGrade(val int ) string{

switch {
case val >= 100:
	return "A"
case val >= 80:
	return "B"
case val >=70:
	return "C"
case val >= 60:
 return "D"
 default :
 return "F"
}}