package entity

type OrderQueue struct {
	Orders []*Order
}

func (oq OrderQueue) Less(i, j int) bool {
	return oq.Orders[i].Price
}

//Less que compara 2 valores j e i
//Swap que inverte os valores valor i passa para j e o inverso
//Len para ver o tamanho dos dados
//Push adiciona nos valores
//Pop remove os valores
