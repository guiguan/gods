package unionfind

import (
	"encoding/json"

	"github.com/emirpasic/gods/containers"
)

func assertSerializationImplementation() {
	var _ containers.JSONSerializer = (*Graph)(nil)
	var _ containers.JSONDeserializer = (*Graph)(nil)
}

type graph struct {
	Count   int   `json:"count"`
	Parents []int `json:"parents"`
	Sizes   []int `json:"sizes"`
}

// ToJSON outputs the JSON representation of the graph
func (g *Graph) ToJSON() ([]byte, error) {
	gr := &graph{}
	gr.Count = g.count
	gr.Parents = g.parents
	gr.Sizes = g.sizes

	return json.Marshal(gr)
}

// FromJSON populates the graph from the input JSON representation.
func (g *Graph) FromJSON(data []byte) error {
	gr := &graph{}

	err := json.Unmarshal(data, gr)
	if err != nil {
		return err
	}

	g.count = gr.Count
	g.parents = gr.Parents
	g.sizes = gr.Sizes

	return nil
}
