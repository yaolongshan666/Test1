package main

func main() {
	var err error
	if err == nil{
		println("kong")
	}
	println(my(err))
}

func my(err error) bool {
	return err == nil
}
