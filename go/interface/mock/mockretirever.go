package mock

type Retriever struct {
	Contents string
}

// 實現者只要實現interface裡的方法(Get)就行，不需要在乎是哪個interface)
func (r Retriever) Get(url string) string {
	return r.Contents
}

