package solr

type SolrResponse struct {
	Status int `json:"status"`
	QTime  int `json:"qtime"`
	Params struct {
		Query  string `json:"q"`
		Indent string `json:"indent"`
		Wt     string `json:"wt"`
	} `json:"params"`
	Response       Response `json:"response"`
	NextCursorMark string   `json:"nextCursorMark"`
	Adds           Adds     `json:"adds"`
}

type Response struct {
	NumFound uint32                   `json:"numFound"`
	Start    int                      `json:"start"`
	Docs     []map[string]interface{} `json:"docs"`
}

func GetDocIdFromDoc(m map[string]interface{}) string {
	return m["id"].(string)
}

func GetVersionFromDoc(m map[string]interface{}) int {
	return m["_version_"].(int)
}

type Adds map[string]int

type updateResponse struct {
	Response struct {
		Status int `json:"status"`
		QTime  int `json:"QTime"`
		RF     int `json:"rf"`
		MinRF  int `json:"min_rf"`
	} `json:"responseHeader"`
	Error struct {
		Metadata []string `json:"metadata"`
		Msg      string   `json:"msg"`
		Code     int      `json:"code"`
	}
}

type deleteRequest struct {
	Delete []string `json:"delete"`
}
