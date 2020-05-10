package	main

import(
	"fmt"
	"math/rand"
	"time"
)

type neuron struct{
	neuronIndex 	int
	valueInput		float64
	weight 			float64
	valueOutput 	float64
}

type layer struct{
	layerIndex 		int
	qtdNeurons 		int
	bias			float64
	weightBias		float64
	neuron 			[]neuron
}

type network struct{
	qtdLayers		int
	qtdNeurons 		int
	layer 			[]layer
}

func (l *layer) addBias(){
	for i := range l.neuron {
		l.neuron[i].valueOutput = l.bias*l.weightBias + l.neuron[i].valueInput*l.neuron[i].weight
	}
}

func (l *layer) print(){

	for i := 0; i < l.qtdNeurons; i++ {

		fmt.Println(l.neuron[i].valueInput)

	}
}

func (n *network) makeNetwork (neuronByLayer ...int){

	n.layer = make([]layer, len(neuronByLayer))
	n.qtdLayers = len(neuronByLayer)

	for i, v := range neuronByLayer {

			n.layer[i].neuron = make([]neuron, v)
			n.qtdNeurons += v
			n.layer[i].qtdNeurons = v
	}
}

func (n *network) generateWeight(){

	rand.Seed(time.Now().UnixNano())

	for i := range n.layer {

		n.layer[i].weightBias = rand.Float64()

		for j := range n.layer[i].neuron {

			n.layer[i].neuron[j].weight = rand.Float64()

		}
	}
}

func (n *network) forwardPropagation(){

	for i := range n.layer {
		n.layer[i].addBias()

		for j := range n.layer[i].neuron {

			if i + 1 < len(n.layer){
				for k := range n.layer[i+1].neuron{
					n.layer[i+1].neuron[k].valueInput += n.layer[i].neuron[j].valueOutput
				}
			}else{
				break
			}

		}
	}

	
}

func (n *network) print() {

	for j := range n.layer {

		fmt.Printf("Layer[%d]\n", j)

		for i := range n.layer[j].neuron {

			fmt.Printf("Neuronio[%d] - Weight: %f\n", i, n.layer[j].neuron[i].weight)
		}

		fmt.Println()
	}

}

func main(){

	var n network

	n.makeNetwork( 2, 3, 2 )
	n.generateWeight()

	fmt.Println("Qtd de Layers: ", n.qtdLayers)
	fmt.Println("Qtd de Neurinios: ", n.qtdNeurons)
	fmt.Println()
	
	n.layer[0].neuron[0].valueInput = 6.77
	n.layer[0].neuron[1].valueInput = 3.62
	

	n.layer[0].bias = 1.0
	n.layer[1].bias = 1.0

	n.forwardPropagation()

	for i := range n.layer {
		fmt.Println("Layer -",i)
		for j := range n.layer[i].neuron {
			fmt.Printf("Neuronio[%d]:\n", j )
			fmt.Printf("Entrada:%f\n", n.layer[i].neuron[j].valueInput )
			fmt.Printf("Saída:%f\n\n", n.layer[i].neuron[j].valueOutput )
		}
	}

}

