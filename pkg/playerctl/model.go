package playerctl

const (
	STATUS_PLAYING    = "Playing"
	STATUS_PAUSED     = "Paused"
	STATUS_NO_PLAYERS = "No players found"

	COMMAND  = "playerctl"
	STATUS   = "status"
	METADATA = "metadata"

	LIST_PLAYERS  = "-l"
	PLAYER_OPTION = "-p"

	SPOTIFY = "spotify"
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
