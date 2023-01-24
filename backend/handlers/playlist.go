package handlers

import (
	"github.com/McCune1224/Echo/models"
	"github.com/McCune1224/Echo/repository"
	"github.com/gofiber/fiber/v2"
)

func GetAllPlaylists(c *fiber.Ctx) error {
	userID, err := GetUserIDFromJWT(c)
	if err != nil {
		return c.Status(401).JSON(fiber.Map{
			"message": "Unauthorized",
		})
	}

	// Get all playlists from database where user_id = userID
	userPlaylists := []models.Playlist{}
	repository.DBConnection.Where("user_id = ?", userID).Find(&userPlaylists)

	playlistResponse := []models.PlaylistResponse{}
	for _, currPlaylist := range userPlaylists {
		newPlaylistResponse := models.NewPlaylistResponse(&currPlaylist)
		playlistResponse = append(playlistResponse, *newPlaylistResponse)
	}
	return c.JSON(playlistResponse)
}

func GetPlaylist(c *fiber.Ctx) error {
	userID, err := GetUserIDFromJWT(c)
	if err != nil {
		return c.Status(401).JSON(fiber.Map{
			"message": "Unauthorized",
		})
	}

	// Get playlist id from url
	playlistID := c.Params("id")
	userPlaylist := models.Playlist{}

	// Get playlist from database where user_id = userID and id = playlistID
	repository.DBConnection.Where("user_id = ? AND id = ?", userID, playlistID).First(&userPlaylist)

	playlistResponse := models.NewPlaylistResponse(&userPlaylist)

	return c.JSON(playlistResponse)
}

func CreatePlaylist(c *fiber.Ctx) error {
	userID, err := GetUserIDFromJWT(c)
	if err != nil {
		return c.Status(401).JSON(fiber.Map{
			"message": "Unauthorized",
		})
	}

	// Get playlist data from request
	var playlistData models.Playlist
	err = c.BodyParser(&playlistData)
	if err != nil {
		return err
	}

	// Check if playlist name is empty
	if playlistData.Name == "" {
		return c.Status(400).JSON(fiber.Map{
			"message": "Playlist name is required",
		})
	}
	// Check if playlist already exists
	var duplicatePlaylist models.Playlist
	repository.DBConnection.Where("user_id = ? AND name = ?", userID, playlistData.Name).First(&duplicatePlaylist)
	if duplicatePlaylist.ID != 0 {
		return c.Status(400).JSON(fiber.Map{
			"message": "Playlist already exists",
		})
	}

	// Create playlist in database
	playlist := models.Playlist{
		Name:        playlistData.Name,
		Description: playlistData.Description,
		UserID:      userID,
	}
	repository.DBConnection.Create(&playlist)

	// Return the playlist as JSON with a status code of 201 (Created)
	playlistResponse := models.NewPlaylistResponse(&playlist)
	return c.Status(201).JSON(playlistResponse)
}

func UpdatePlaylistTracks(c *fiber.Ctx) error {
	userID, err := GetUserIDFromJWT(c)
	if err != nil {
		return c.Status(401).JSON(fiber.Map{
			"message": "Unauthorized",
		})
	}

	// Get playlist id from request
	playlistID := c.Params("id")
	if playlistID == "" {
		return c.Status(400).JSON(fiber.Map{
			"message": "Playlist id is required",
		})
	}

	// Get playlist from database
	var playlist models.Playlist
	// query the database for the playlist with the given id and user_id
	repository.DBConnection.Where("id = ? AND user_id = ?", playlistID, userID).First(&playlist)
	// Check if playlist exists
	if playlist.ID == 0 {
		return c.Status(404).JSON(fiber.Map{
			"message": "Playlist does not exist",
		})
	}

	// Get tracks from request
	tracks := []models.TrackResponse{}
	err = c.BodyParser(&tracks)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"message": "Invalid track data",
			"error":   err.Error(),
		})
	}

	// Check if tracks are empty
	if len(tracks) == 0 {
		return c.Status(400).JSON(fiber.Map{
			"message": "Tracks are required",
		})
	}

	// add tracks to playlist
	for _, track := range tracks {
		dbTrackEntry := models.Track{
			Name:       track.Name,
			Artist:     track.Artist,
			Album:      track.Album,
			PlaylistID: playlist.ID,
		}
		playlist.Tracks = append(playlist.Tracks, dbTrackEntry)
	}

	// Save playlist to database
	repository.DBConnection.Save(&playlist)

	// Return the playlist as JSON with a status code of 200 (OK)
	playlistResponse := models.NewPlaylistResponse(&playlist)
	return c.Status(200).JSON(playlistResponse)
}

func UpdatePlaylistDetails(c *fiber.Ctx) error {
	userID, err := GetUserIDFromJWT(c)
	if err != nil {
		return c.Status(401).JSON(fiber.Map{
			"message": "Unauthorized",
		})
	}

	// Get playlist id from request
	playlistID := c.Params("id")
	// Get playlist from database
	var playlist models.Playlist
	// query the database for the playlist with the given id and user_id
	repository.DBConnection.Where("id = ? AND user_id = ?", playlistID, userID).First(&playlist)
	// Check if playlist exists
	if playlist.ID == 0 {
		return c.Status(404).JSON(fiber.Map{
			"message": "Playlist does not exist",
		})
	}

	newPlaylistData := models.Playlist{}
	err = c.BodyParser(&newPlaylistData)
	if err != nil {
		return c.Status(400).JSON(
			fiber.Map{
				"message": "Invalid playlist data",
				"error":   err.Error(),
			},
		)
	}

	// Update playlist data
	models.OverwriteModel(&playlist, &newPlaylistData)

	// Save playlist to database, do not create a new playlist if it does not exist
	info := repository.DBConnection.Save(&playlist)
	if info.Error != nil {
		return c.Status(400).JSON(
			fiber.Map{
				"message": "Invalid playlist data",
				"error":   info.Error.Error(),
			},
		)
	}
	return c.JSON(fiber.Map{
		"message": "Playlist updated",
	})
}

func DeletePlaylist(c *fiber.Ctx) error {
	userID, err := GetUserIDFromJWT(c)
	if err != nil {
		return c.Status(401).JSON(fiber.Map{
			"message": "Unauthorized",
		})
	}

	// Get playlist id from request
	playlistID := c.Params("id")

	// Get playlist from database
	var playlist models.Playlist
	// query the database for the playlist with the given id and user_id
	repository.DBConnection.Where("id = ? AND user_id = ?", playlistID, userID).First(&playlist)

	// Check if playlist exists
	if playlist.ID == 0 {
		return c.Status(404).JSON(fiber.Map{
			"message": "Playlist does not exist",
		})
	}

	// Delete playlist from database
	repository.DBConnection.Delete(&playlist)

	return c.JSON(fiber.Map{
		"message": "Playlist deleted",
	})
}

func DeletePlaylistTracks(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"TODO": "Delete playlist tracks",
	})
}
