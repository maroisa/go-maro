package server

type IndexData struct {
	SourceMessage string
	AliasMessage  string
}

func (i *IndexData) validateSource(source string) {
	if source == "" {
		i.SourceMessage = "source cannot be empty!"
		return
	}
}

func (i *IndexData) validateAlias(alias string) {
	if alias == "" {
		i.AliasMessage = "alias cannot be empty!"
		return
	}
}

func (i IndexData) isValid() bool {
	if i.SourceMessage != "" || i.AliasMessage != "" {
		return false
	}
	return true
}
