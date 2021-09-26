package gitee

import (
	"fmt"
	"net/http"
	"time"
)

type RepositoryFilesService struct {
	client *Client
}

type File struct {
	Type        string `json:"type"`
	Encoding    string `json:"encoding"`
	Size        int    `json:"size"`
	Name        string `json:"name"`
	Path        string `json:"path"`
	Content     string `json:"content"`
	SHA         string `json:"sha"`
	URL         string `json:"url"`
	HTMLURL     string `json:"html_url"`
	DownloadURL string `json:"download_url"`
}

func (r File) String() string {
	return Stringify(r)
}

type GetFileOptions struct {
	Ref *string `url:"ref,omitempty""`
}

func (s *RepositoryFilesService) GetFile(owner, repo, path string, opts *GetFileOptions) (*File, *Response, error) {

	u := fmt.Sprintf("repos/%s/%s/contents/%s", pathEscape(owner), pathEscape(repo), pathEscape(path))

	req, err := s.client.NewRequest(http.MethodGet, u, opts)
	if err != nil {
		return nil, nil, err
	}

	f := new(File)
	resp, err := s.client.Do(req, f)
	if err != nil {
		return nil, resp, err
	}

	return f, resp, err
}

type CreateFileOptions struct {
	Content        string `json:"content"`
	Message        string `json:"message"`
	Branch         string `json:"branch"`
	CommitterName  string `json:"committer[name]"`
	CommitterEmail string `json:"commtter[email]"`
	AuthorName     string `json:"author[name]"`
	AuthorEmail    string `json:"author[email]"`
}

type FileOperationResponse struct {
	Content *struct {
		Name        string `json:"name"`
		Path        string `json:"path"`
		Size        int    `json:"size"`
		SHA         string `json:"sha"`
		Type        string `json:"file"`
		URL         string `json:"url"`
		HTMLURL     string `json:"html_url"`
		DownloadURL string `json:"download_url"`
	} `json:"content"`
	Commit *struct {
		SHA    string `json:"sha"`
		Author *struct {
			Name  string    `json:"name"`
			Date  *time.Time `json:"date"`
			Email string    `json:"email"`
		} `json:"author"`
		Committer *struct {
			Name  string    `json:"name"`
			Date  *time.Time `json:"date"`
			Email string    `json:"email"`
		} `json:"committer"`
		Message string `json:"message"`
		Tree    *struct {
			SHA string `json:"sha"`
			URL string `json:"url"`
		} `json:"tree"`
		Parents []*struct {
			SHA string `json:"sha"`
			URL string `json:"url"`
		} `json:"parents"`
	} `json:"commit"`
}

// CreateFile return 400 {message: A file with this name already exists}
func (s *RepositoryFilesService) CreateFile(owner, repo, path string, opts *CreateFileOptions) (*FileOperationResponse, *Response, error) {

	u := fmt.Sprintf("repos/%s/%s/contents/%s", pathEscape(owner), pathEscape(repo), pathEscape(path))

	req, err := s.client.NewRequest(http.MethodPost, u, opts)
	if err != nil {
		return nil, nil, err
	}

	f := new(FileOperationResponse)
	resp, err := s.client.Do(req, f)
	if err != nil {
		return nil, resp, err
	}

	return f, resp, err
}

type UpdateFileOptions struct {
	Content        string `json:"content"`
	SHA            string `json:"sha"`
	Message        string `json:"message"`
	Branch         string `json:"branch"`
	CommitterName  string `json:"committer[name]"`
	CommitterEmail string `json:"commtter[email]"`
	AuthorName     string `json:"author[name]"`
	AuthorEmail    string `json:"author[email]"`
}

// UpdateFile if path does not exist, 404 {message: File}
func (s *RepositoryFilesService) UpdateFile(owner, repo, path string, opts *UpdateFileOptions) (*FileOperationResponse, *Response, error) {

	u := fmt.Sprintf("repos/%s/%s/contents/%s", pathEscape(owner), pathEscape(repo), pathEscape(path))

	req, err := s.client.NewRequest(http.MethodPut, u, opts)
	if err != nil {
		return nil, nil, err
	}

	f := new(FileOperationResponse)
	resp, err := s.client.Do(req, f)
	if err != nil {
		return nil, resp, err
	}

	return f, resp, err
}

type DeleteFileOptions struct {
	SHA            string `json:"sha"`
	Message        string `json:"message"`
	Branch         string `json:"branch"`
	CommitterName  string `json:"committer[name]"`
	CommitterEmail string `json:"commtter[email]"`
	AuthorName     string `json:"author[name]"`
	AuthorEmail    string `json:"author[email]"`
}

// DeleteFile if path does not exist, 404 {message: File}
func (s *RepositoryFilesService) DeleteFile(owner, repo, path string, opts *DeleteFileOptions) (*FileOperationResponse, *Response, error) {

	u := fmt.Sprintf("repos/%s/%s/contents/%s", pathEscape(owner), pathEscape(repo), pathEscape(path))

	req, err := s.client.NewRequest(http.MethodDelete, u, opts)
	if err != nil {
		return nil, nil, err
	}

	f := new(FileOperationResponse)
	resp, err := s.client.Do(req, f)
	if err != nil {
		return nil, resp, err
	}

	return f, resp, err
}
