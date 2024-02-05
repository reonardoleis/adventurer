package vector_search

var _ VectorSearch = (*Pinecone)(nil)

type VectorSearch interface {
	Upsert(namespace string, vector []float32, id string) error
	Search(vector []float32, k int) ([]string, error)
}
