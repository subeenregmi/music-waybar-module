package playerctl

import (
	"fmt"
	"os/exec"
	"regexp"
	"strconv"
	"strings"
)

func metadataRegex(key string) *regexp.Regexp {
	return regexp.MustCompile(fmt.Sprintf(`spotify (?:xesam|mpris):%s\s*(.*)`, key))
}

func Players() ([]string, error) {
	output, err := exec.Command(COMMAND, "-l").Output()
	if err != nil {
		return nil, err
	}

	players := string(output)
	return strings.Split(players, " "), nil
}

func Status() (string, error) {
	output, err := exec.Command(COMMAND, STATUS).Output()
	if err != nil {
		return "", err
	}

	return string(output), nil
}

func GetMetadata() (Metadata, error) {
	output, err := exec.Command(COMMAND, METADATA).Output()
	if err != nil {
		return Metadata{}, err
	}

	return parseMetadata(string(output))
}

func parseMetadata(raw string) (Metadata, error) {
	m := Metadata{}

	m.getTrackID(raw)
	if err := m.getLength(raw); err != nil {
		return Metadata{}, err
	}
	m.getArtURL(raw)
	m.getAlbum(raw)
	m.getAlbumArtist(raw)
	m.getArtist(raw)
	if err := m.getAutoRating(raw); err != nil {
		return Metadata{}, err
	}
	if err := m.getDiscNumber(raw); err != nil {
		return Metadata{}, err
	}
	m.getTitle(raw)
	if err := m.getTrackNumber(raw); err != nil {
		return Metadata{}, err
	}
	m.getURL(raw)

	return m, nil
}

func (m *Metadata) getTrackID(raw string) {
	regexp := metadataRegex("trackid")
	matches := regexp.FindStringSubmatch(raw)

	m.TrackID = matches[len(matches)-1]
}

func (m *Metadata) getLength(raw string) error {
	regexp := metadataRegex("length")
	matches := regexp.FindStringSubmatch(raw)

	match := matches[len(matches)-1]

	length, err := strconv.Atoi(match)
	if err != nil {
		return err
	}

	m.Length = length
	return nil
}

func (m *Metadata) getArtURL(raw string) {
	regexp := metadataRegex("artUrl")
	matches := regexp.FindStringSubmatch(raw)

	m.ArtURL = matches[len(matches)-1]
}

func (m *Metadata) getAlbum(raw string) {
	regexp := metadataRegex("album")
	matches := regexp.FindStringSubmatch(raw)

	m.Album = matches[len(matches)-1]
}

func (m *Metadata) getAlbumArtist(raw string) {
	regexp := metadataRegex("albumArtist")
	matches := regexp.FindStringSubmatch(raw)

	m.AlbumArtist = matches[len(matches)-1]
}

func (m *Metadata) getArtist(raw string) {
	regexp := metadataRegex("artist")
	matches := regexp.FindStringSubmatch(raw)

	m.Artist = matches[len(matches)-1]
}

func (m *Metadata) getAutoRating(raw string) error {
	regexp := metadataRegex("autoRating")
	matches := regexp.FindStringSubmatch(raw)

	match := matches[len(matches)-1]

	rating, err := strconv.ParseFloat(match, 64)
	if err != nil {
		return err
	}

	m.AutoRating = rating
	return nil
}

func (m *Metadata) getDiscNumber(raw string) error {
	regexp := metadataRegex("discNumber")
	matches := regexp.FindStringSubmatch(raw)

	match := matches[len(matches)-1]

	discNumber, err := strconv.Atoi(match)
	if err != nil {
		return err
	}

	m.DiscNumber = discNumber
	return nil
}

func (m *Metadata) getTitle(raw string) {
	regexp := metadataRegex("title")
	matches := regexp.FindStringSubmatch(raw)

	m.Title = matches[len(matches)-1]
}

func (m *Metadata) getTrackNumber(raw string) error {
	regexp := metadataRegex("trackNumber")
	matches := regexp.FindStringSubmatch(raw)

	match := matches[len(matches)-1]

	trackNumber, err := strconv.Atoi(match)
	if err != nil {
		return err
	}

	m.TrackNumber = trackNumber
	return nil
}

func (m *Metadata) getURL(raw string) {
	regexp := metadataRegex("url")
	matches := regexp.FindStringSubmatch(raw)

	m.URL = matches[len(matches)-1]
}
