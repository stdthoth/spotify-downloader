package playlist

import (
	"context"
	"net/http"
	"time"

	"github.com/zmb3/spotify/v2"
)

func GetUserPlaylists() ([]spotify.SimplePlaylist, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*120)
	defer cancel()

	client := spotify.New(http.DefaultClient, spotify.WithRetry(true))
	user, err := client.CurrentUser(ctx)
	if err != nil {
		return nil, err
	}
	playlists, err := client.GetPlaylistsForUser(ctx, user.ID)
	if err != nil {
		return nil, err
	}

	return playlists.Playlists, nil

}
