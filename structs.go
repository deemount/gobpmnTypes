package gobpmn_types

// CoreID ...
type CoreID struct {
	ID string `xml:"id,attr,omitempty" json:"id"`
}

// CoreInnerID ...
type CoreInnerID struct {
	ID string `xml:",innerxml" json:"id"`
}

// BaseAttributes ...
type BaseAttributes struct {
	ID   string `xml:"id,attr,omitempty" json:"id"`
	Name string `xml:"name,attr,omitempty" json:"name,omitempty"`
}
