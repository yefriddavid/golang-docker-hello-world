package main

import (
  "fmt"
  "bytes"
  "net/http"
  //"os"
  "io/ioutil"
  //"encoding/json"
)


func main(){

  //var jsonStr = []byte(`{"title":"Buy cheese and bread for breakfast."}`)
  //req, err := http.NewRequest("POST","http://181.143.147.114:96/WebServices/WebServices/MinaServicios/getPlanillaEspecial.asmx", bytes.NewBuffer(jsonStr))
  var jsonStr = bytes.NewBufferString(`<x:Envelope xmlns:x="http://schemas.xmlsoap.org/soap/envelope/" xmlns:sof="Softtools.Gespasajeros.Servicios"> <x:Header/> <x:Body> <sof:Retorna_Planilla_Especial> <sof:FiltrosValidarUsuario> <sof:intCodigoEmpresaUsuario>1</sof:intCodigoEmpresaUsuario> <sof:strIdentificadorUsuario>WEBSERVICE</sof:strIdentificadorUsuario> <sof:strClave>mina2017*</sof:strClave> </sof:FiltrosValidarUsuario> <sof:FiltroPlanillas> <sof:dteFechaInicial>2018-04-01T00:00:00</sof:dteFechaInicial> <sof:dteFechaFinal>2018-04-06T00:00:00</sof:dteFechaFinal> <!--<sof:lonNumeroPlanillaInicial>0</sof:lonNumeroPlanillaInicial> <sof:lonNumeroPlanillaFinal>10000000000</sof:lonNumeroPlanillaFinal>--> <sof:intPlanillaAnulada>0</sof:intPlanillaAnulada> <!--<sof:intPlanillaCumplida>0</sof:intPlanillaCumplida>--> </sof:FiltroPlanillas> </sof:Retorna_Planilla_Especial> </x:Body> </x:Envelope>`)
  req, err := http.NewRequest("POST","http://181.143.147.115:96/WebServices/WebServices/MinaServicios/getPlanillaEspecial.asmx", jsonStr)

    //req.Header.Set("X-Custom-Header", "myvalue")
    req.Header.Set("Content-Type", "text/xml")

    client := &http.Client{}
    resp, err := client.Do(req)
    if err != nil {
        panic(err)
    }
    defer resp.Body.Close()

    fmt.Println("response Status:", resp.Status)
    fmt.Println("response Headers:", resp.Header)
    body, _ := ioutil.ReadAll(resp.Body)
    fmt.Println("response Body:", string(body))

}
