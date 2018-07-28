package main

import (
	"encoding/xml"
	"fmt"
	"encoding/json"
	//"reflect"
)

var payload = `
<soapenv:Envelope xmlns:soapenv="http://schemas.xmlsoap.org/soap/envelope/" xmlns:ns="http://xmlns.fmr.com/systems/dev/xtrac/2004/06/" xmlns:ser="http://xmlns.fmr.com/common/headers/2005/12/ServiceProcessingDirectives" xmlns:ser1="http://xmlns.fmr.com/common/headers/2005/12/ServiceCallContext" xmlns:typ="http://xmlns.fmr.com/systems/dev/xtrac/2004/06/types">
<soapenv:Body>
<Retorna_Planilla_EspecialResponse>
<Retorna_Planilla_EspecialResult>
	<lstPlanillasEspeciales>
                        <lonNumeroPlanilla>1831595</lonNumeroPlanilla>
                        <lonNumeroSolicitudEspecial>1585012</lonNumeroSolicitudEspecial>	
	</lstPlanillasEspeciales>
	<lstPlanillasEspeciales>
                        <lonNumeroPlanilla>1831595</lonNumeroPlanilla>
                        <lonNumeroSolicitudEspecial>1585012</lonNumeroSolicitudEspecial>	
	</lstPlanillasEspeciales>	
</Retorna_Planilla_EspecialResult>
</Retorna_Planilla_EspecialResponse>
</soapenv:Body>
</soapenv:Envelope>
`

type createEnvelope struct {
	CreateBody createBody `xml:"Body"`
}

type createBody struct {
	Create create `xml:"Retorna_Planilla_EspecialResponse"`
}

type create struct {
	Result Planillas `xml:"Retorna_Planilla_EspecialResult"`
}

type Planillas struct {
	Items []Planilla `xml:"lstPlanillasEspeciales"`
}
type Planilla struct {
	LonNumeroPlanilla int `xml:"lonNumeroPlanilla"`
	LonNumeroSolicitudEspecial int `xml:"lonNumeroSolicitudEspecial"`
}

func main() {
	var createEnv createEnvelope

	err := xml.Unmarshal([]byte(payload), &createEnv)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
    	jsonData, _ := json.Marshal(createEnv.CreateBody.Create.Result)
    	fmt.Println(string(jsonData))
	
	planillas := createEnv.CreateBody.Create.Result.Items
	fmt.Println(planillas[1])
	//fmt.Println(reflect.TypeOf(planillas))
	
//	fmt.Println(reflect.TypeOf(createEnv.CreateBody.Create))

	/*for i := range createEnv.CreateBody.Create.Planillas {
		fmt.Println("%s", i.LonNumeroSolicitudEspecial)
	

	}*/
//	fmt.Printf("%v\n", createEnv.CreateBody.Create.Planillas)
}




