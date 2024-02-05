package vector_search

import (
	"github.com/reonardoleis/adventurer/internal/http"
	pt "github.com/reonardoleis/adventurer/internal/vector_search/pinecone_types"
)

const (
	URL_PREFIX = "https://api.pinecone.io"
)

type Pinecone struct {
	apiKey    string
	indexHost string
}

func (p Pinecone) headers() map[string]string {
	return map[string]string{
		"Api-Key": p.apiKey,
	}
}

func NewPinecone(apiKey, indexHost string) (*Pinecone, error) {
	return &Pinecone{apiKey: apiKey, indexHost: indexHost}, nil
}

func (p *Pinecone) CreateIndex(name string, dimension EmbeddingDimension) error {
	req := pt.CreateIndexRequest{
		Name:      name,
		Dimension: int32(dimension),
		Metric:    "cosine",
		Spec: pt.Spec{
			Serverless: pt.Serverless{
				Cloud:  "aws",
				Region: "us-west-2",
			},
		},
	}

	response := make(map[string]interface{})
	err := http.PostAndDecode(
		URL_PREFIX+"/indexes",
		p.headers(),
		req,
		&response,
	)
	if err != nil {
		return err
	}

	return nil
}

func (p *Pinecone) Upsert(namespace string, vector []float32, id string) error {
	req := pt.UpsertRequest{
		Namespace: namespace,
		Vectors: []pt.Vector{
			{
				ID:     id,
				Values: vector,
			},
		},
	}

	response := make(map[string]interface{})
	err := http.PostAndDecode(
		p.indexHost+"/vectors/upsert",
		p.headers(),
		req,
		&response,
	)
	if err != nil {
		return err
	}

	return nil
}

func (p *Pinecone) Search(vector []float32, k int) ([]string, error) {
	return nil, nil
}
