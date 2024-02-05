package pinecone_types

type Serverless struct {
	Cloud  string `json:"cloud"`
	Region string `json:"region"`
}

type Spec struct {
	Serverless Serverless `json:"serverless"`
}

type CreateIndexRequest struct {
	Name      string `json:"name"`
	Dimension int32  `json:"dimension"`
	Metric    string `json:"metric"`
	Spec      Spec   `json:"spec"`
}

type Vector struct {
	ID     string    `json:"id"`
	Values []float32 `json:"values"`
}

type UpsertRequest struct {
	Namespace string   `json:"namespace"`
	Vectors   []Vector `json:"vectors"`
}
