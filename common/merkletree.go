package common

import (
	"crypto/sha256"
)

type MerkleTree struct{
	RootNode *MerkleNode
	//tree [][]MerkleNode
}

type MerkleNode struct{
	Left *MerkleNode
	Right *MerkleNode
	Data []byte
}
/*
머클 트리의 노드를 제작하는 과정 left,right가 없을 경우에 리프노드로 자신의 데이터만 해싱한다.
그 외의 경우는 left,right의 데이터값을 append한 뒤 해싱한다.
새로 만든 노드의 데이터를 입력한 뒤 마지막에는 left, right 값을 입력해준다.
이후 새로 만든 노드를 반환한다. left,right 노드를 이용해 새로운 노드를 만들어내는 함수
 */
func NewMerkleNode(left *MerkleNode,right *MerkleNode,data []byte)*MerkleNode{
	node :=MerkleNode{}
	if left==nil&&right==nil{
		hash:=sha256.Sum256(data)
		node.Data=hash[:]
	}else{
		childData:=append(left.Data,right.Data...)
		hash:=sha256.Sum256(childData)
		node.Data=hash[:]
		//fmt.Println(hex.EncodeToString(node.Data))
		//sha256.sum256하면 그냥 sha256한것과 다른가? 단순히 좌,우 데이터 붙여서 sha256온라인에서 구한 sha256 값과 다른값이 나온다
	}
	node.Left=left
	node.Right =right
	return &node
}

func NewMerkleTree(data [][]byte)*MerkleTree {
	var nodes []MerkleNode
	if len(data)%2 == 1 {
		data = append(data, data[len(data)-1])
	}
	for _, nodeData := range data {
		node := NewMerkleNode(nil, nil, nodeData)
		nodes = append(nodes, *node)
	}
	for j := 0; j < len(data)/2; j++ {
		var parentNode []MerkleNode
		for i := 0; i < len(nodes); i += 2 {
			newNode := NewMerkleNode(&nodes[i], &nodes[i+1], nil)
			parentNode = append(parentNode, *newNode)
		}
		nodes = parentNode
		if len(nodes) < 2 {
			break
		}
	}
	Tree := MerkleTree{&nodes[0]}
	return &Tree
}

//2차원 배열써서 트리 층별 요소 저장하려했는데 2차원 배열 사용하는데 index out of range 오류발생 2차원 배열 사용법
//func NeswMerkleTree(data [][]byte)*MerkleTree{
//	var nodes [][]MerkleNode
//	if len(data)%2==1 {
//		data = append(data, data[len(data)-1])
//	}
//	for _,nodeData :=range data {
//		node := NewMerkleNode(nil, nil, nodeData)
//		nodes[0] = append(nodes[0], *node)
//		fmt.Println(nodes)
//	}
//	step:=0
//	for j:=0;j<len(data);j++ {
//		var parentNode []MerkleNode
//		for i := 0; i < len(nodes); i += 2 {
//			node := NewMerkleNode(&nodes[step][j], &nodes[step][j+1], nil)
//			parentNode = append(parentNode, *node)
//		}
//		step++
//		nodes[step] = parentNode
//		if len(nodes[step]) < 2 {
//			break
//		}
//	}
//	Tree := MerkleTree{&nodes[step][0],nodes}
//	return &Tree
//}