package strike

/*
FileInfo contains information on the torrent's files
*/
type FileInfo struct {
	FileNames			[]string	`json:"file_names"`
	FileLengths			[]uint64	`json:"file_lengths"`
}

/*
Torrent contains information about the torrent
*/
type Torrent struct {
	Hash				string		`json:"torrent_hash"`
	ImdbID				string		`json:"imdbid"`
	Title				string		`json:"torrent_title"`
	Category			string		`json:"torrent_category"`
	Subcategory			string		`json:"sub_category"`
	Seeds				uint64		`json:"seeds"`
	Leeches				uint64		`json:"leeches"`
	FileCount			uint64		`json:"file_count"`
	Size				float64		`json:"size"`
	DownloadCount		uint64		`json:"download_count"`
	UploadDate			string		`json:"upload_date"`
	UploaderUsername	string		`json:"uploader_username"`
	FileInfo			FileInfo	`json:"file_info"`
	Page				string		`json:"page"`
	RssFeed				string		`json:"rss_feed"`
	MagnetUri			string		`json:"magnet_uri"`
}

/*
getDownload returns a TorrentDownload
*/
func (t Torrent) getDownload(string) (TorrentDownload, error) {
	return Download(t.Hash)
}

/*
getInfo returns an Results object containing torrent info
*/
func (t Torrent) getInfo(string) (Results, error) {
	return Info(t.Hash)
}

/*
getIMDb returns an IMDb movie information object.
*/
func (t Torrent) getIMDb(string) (IMDb, error) {
	return Imdb(t.ImdbID)
}

/*
getDescription returns a TorrentDescription containing a base64 encoded HTML description, 
or a TorrentDescription containing a plain text description if true is passed as the 
first parameter. 
*/
func (t Torrent) getDescription(params ...bool) (result TorrentDescription, err error) {
	decode := false
	if (len(params) > 0) {
		decode = params[0]
	}
	return Description(t.ImdbID, decode)
}

/*
Results contains information about an API response and holds an array of Torrent objects
*/
type Results struct {
	Results				uint64		`json:"results"`
	StatusCode			uint16		`json:"statuscode"`
	ResponseTime		float32		`json:"responsetime"`
	Torrents			[]Torrent	`json:"torrents"`
}

/*
TorrentCount contains a Count from a count API response
*/
type TorrentCount struct {
	StatusCode			uint16		`json:"statuscode"`
	Count				uint64		`json:"message"`
}

/*
TorrentDescription contains a Description from a description API response
*/
type TorrentDescription struct {
	StatusCode			uint16		`json:"statuscode"`
	Description			string		`json:"message"`
}

/*
TorrentDownload contains a Uri from a download API response
*/
type TorrentDownload struct {
	StatusCode			uint16		`json:"statuscode"`
	Uri					string		`json:"message"`
}

/*
IMDb contains movie information from an imdb API response
*/
type IMDb struct {
	ID					uint64		`json:"id"`
	ImdbID				string		`json:"imdbID"`
	Title				string		`json:"Title"`
	Year				uint16		`json:"Year"`
	Rating				float32		`json:"Rating"`
	Runtime				string		`json:"Runtime"`
	Genre				string		`json:"Genre"`
	Released			string		`json:"Released"`
	Director			string		`json:"Director"`
	Writer				string		`json:"Writer"`
	Cast				string		`json:"Cast"`
	Metacritic			string		`json:"Metacritic"`
	ImdbRating			string		`json:"imdbRating"`
	ImdbVotes			string		`json:"imdbVotes"`
	Poster				string		`json:"Poster"`
	Plot				string		`json:"Plot"`
	FullPlot			string		`json:"FullPlot"`
	Language			string		`json:"Language"`
	Country				string		`json:"Country"`
	Awards				string		`json:"Awards"`
	LastUpdated			string		`json:"lastUpdated"`
	Type				string		`json:"Type"`	
}