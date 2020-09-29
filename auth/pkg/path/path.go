package path

type Path struct {
	ID          string   `json:"id" bson:"_id,omitempty"`
	Host        string   `json:"host" bson:"host"`
	Path        string   `json:"path" bson:"path"`
	Method      string   `json:"method" bson:"method"`
	Permissions []string `json:"permissions" bson:"permissions"`
}
