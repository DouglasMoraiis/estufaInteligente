package main

import (
	"../camada"
	"fmt"
	"github.com/google/gopacket"
	"io/ioutil"
	"net"
	"os"
)

func checkError(err error, msg string){
	if err != nil {
		fmt.Fprintf(os.Stderr, "Erro em " + msg, err.Error())
		os.Exit(1)
	}
}

func main() {
	tcpAddr, err := net.ResolveTCPAddr("tcp", ":3200")
	checkError(err, "ResolveTCPAddr")

	listener, err := net.ListenTCP("tcp", tcpAddr)
	checkError(err, "ListenTCP")

	for {
		conn, err := listener.Accept()

		if err != nil {
			return
		}

		//Fazer um if pra saber o tipo das conexão
		//handleClient
		result, err := ioutil.ReadAll(conn)
		checkError(err, "ReadAll")

		packet := gopacket.NewPacket(
			result,
			camada.ParametersLayerType,
			gopacket.Default,
		)

		decodePacket := packet.Layer(camada.ParametersLayerType)

		if decodePacket != nil {
			fmt.Println("Pacote decodificado com sucesso!")
			content, _ := decodePacket.(*camada.ParametersLayer)

			fmt.Println("TemperaturaMin:", int32(content.TemperaturaMin))
			fmt.Println("TemperaturaMax:", int32(content.TemperaturaMax))
			fmt.Println("UmidadeMin:", content.UmidadeMin)
			fmt.Println("NivelCO2Min:", content.NivelCO2Min)
		}
	}
}