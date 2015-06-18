package strike

type Torrent struct {
	Torrent_Hash string
	Imdbid string
	Torrent_Title string
	Torrent_Category string
	Sub_Category string
	Seeds uint64
	Leeches uint64
	FileCount uint64
	Size uint64
	DownloadCount uint64
	UploadDate string
	UploaderUsername string
	Page string
	RssFeed string
	MagnetUrl string
}

type SearchResults struct {
	Results uint64
	StatusCode uint16
	ResponseTime float32
	Torrents []Torrent
}