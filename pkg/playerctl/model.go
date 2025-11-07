package playerctl

const (
	STATUS_PLAYING    = "Playing"
	STATUS_PAUSED     = "Paused"
	STATUS_NO_PLAYERS = "No players found"

	COMMAND  = "playerctl"
	STATUS   = "status"
	METADATA = "metadata"
)

type Metadata struct {
	TrackID     string
	Length      int
	ArtURL      string
	Album       string
	AlbumArtist string
	Artist      string
	AutoRating  float64
	DiscNumber  int
	Title       string
	TrackNumber int
	URL         string
}
