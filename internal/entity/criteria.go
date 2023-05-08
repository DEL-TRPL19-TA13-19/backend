package entity

type CriteriaData struct {
	PairwiseFromJson        [][]float64 `json:"pairwise"`
	PairwiseAfterCalculated [][]float64 `json:"pairwise_after_calculated"`
	Criteria                []float64   `json:"criteria"`
}

type Matrix [][]float64
