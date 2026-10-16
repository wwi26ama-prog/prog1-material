package listen

import "fmt"

func ListenDemo() {
	liste1 := []string{"BMW", "Audi", "Skoda"}

	fmt.Println(liste1)
	fmt.Println(liste1[0])
	fmt.Println(liste1[1])
	fmt.Println(liste1[2])

	liste1[1] = "VW"
	fmt.Println(liste1)
	fmt.Println(liste1[1])

	liste1 = append(liste1, "Ferrari")
	fmt.Println(liste1[3])
}
