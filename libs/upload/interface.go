package upload

type Uploader interface {
	Upload()(path string, err error)
}
