package thread

type Thread struct {
  id string
  title string
  content string
}

func NewThread(id, title, content string) *Thread {
  return &Thread{id, title, content}
}

func (t *Thread) Id() string {
  return t.id
}

func (t *Thread) Title() string {
  return t.title
}

func (t *Thread) Content() string {
  return t.content
}


