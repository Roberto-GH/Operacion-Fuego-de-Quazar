package services

import (
	"log"
	"strings"
)

func GetLocation(distances ...float32) (x, y float32) {

	if len(distances) < 3 {
		x = 0
		y = 0
	} else {
		kenobiDistance := distances[0]
		skywalkerDistance := distances[1]
		satoDistance := distances[2]

		kenobi := []float32{-500, -200}
		skywalker := []float32{100, -100}
		sato := []float32{500, 100}

		distanceA := 2*skywalker[0] - 2*kenobi[0]
		distanceB := 2*skywalker[1] - 2*kenobi[1]
		distanceC := kenobiDistance*kenobiDistance - skywalkerDistance*skywalkerDistance - kenobi[0]*kenobi[0] +
			skywalker[0]*skywalker[0] - kenobi[1]*kenobi[1] + skywalker[1]*skywalker[1]
		distanceD := 2*sato[0] - 2*skywalker[0]
		distanceE := 2*sato[1] - 2*skywalker[1]
		distanceF := skywalkerDistance*skywalkerDistance - satoDistance*satoDistance - skywalker[0]*skywalker[0] +
			sato[0]*sato[0] - skywalker[1]*skywalker[1] + sato[1]*sato[1]

		x = (distanceC*distanceE - distanceF*distanceB) / (distanceE*distanceA - distanceB*distanceD)
		y = (distanceC*distanceD - distanceA*distanceF) / (distanceB*distanceD - distanceA*distanceE)
	}
	log.Printf("Las coordenadas: %v %v", x, y)
	return x, y
}

//FIRMA GET MESSAGE
func GetMessage(messages ...[]string) (msg string) {
	messageDecode := []string{}
	for i := 0; i < len(messages[0]) && i < len(messages[1]) && i < len(messages[2]); i++ {
		messageDecode = append(messageDecode, messages[0][i], messages[1][i], messages[2][i])
	}
	resultado := removeDuplicateElement(messageDecode)
	for i := 0; i < len(resultado); i++ {
		msg += resultado[i] + " "
	}
	msg = strings.Trim(msg, " ")
	log.Println("El mensaje: " + msg)
	return msg
}

//REMOVER DUPLICADOS
func removeDuplicateElement(addrs []string) []string {
	result := make([]string, 0, len(addrs))
	temp := map[string]struct{}{}
	for _, item := range addrs {
		if _, ok := temp[item]; !ok {
			temp[item] = struct{}{}
			result = append(result, item)
		}
	}
	return result
}
