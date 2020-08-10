package utils

type HttpClient interface {
	Get(url string, mp map[string]string, args ...interface{}) error
	Delete(url string, mp map[string]string, args ...interface{}) error
	Post(url string, mp map[string]string, args ...interface{}) error
	//Download(url string)(http.Header, io.ReadCloser, error)
	//Upload(url string, header *multipart.FileHeader, fields map[string]string,respStruct interface{})error
}

type DefautHttpClient struct {
}

func (d *DefautHttpClient) Get(url string, mp map[string]string, args ...interface{}) error {
	return http_request(HTTP_GET_METHOD, url, mp, args)
}

func (d *DefautHttpClient) Delete(url string, mp map[string]string, args ...interface{}) error {
	return http_request(HTTP_DELETE_METHOD, url, mp, args)
}

func (d *DefautHttpClient) Post(url string, mp map[string]string, args ...interface{}) error {
	return http_request(HTTP_POST_METHOD, url, mp, args)
}
