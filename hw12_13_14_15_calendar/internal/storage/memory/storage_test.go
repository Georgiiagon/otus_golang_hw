package memorystorage

import (
	"testing"
	"time"

	storagemodels "github.com/Georgiiagon/otus_golang_hw/hw12_13_14_15_calendar/internal/storage/models"
	"github.com/stretchr/testify/require"
)

func TestStorage(t *testing.T) {
	s := New()
	emptyStorage := make([]*storagemodels.Event, 0)
	require.Equal(t, emptyStorage, s.GetEvents())
}

func TestCreateEvent(t *testing.T) {
	s := New()
	e := storagemodels.Event{
		Title:       "Test event",
		Duration:    time.Minute * 30,
		Description: "Event desc",
		NotifyAt:    time.Now().AddDate(0, 0, 1),
	}
	event, err := s.CreateEvent(e)
	require.NoError(t, err)
	require.Equal(t, s.GetEvents()[0], event)
}

func TestDeleteEvent(t *testing.T) {
	s := New()
	e := storagemodels.Event{
		Title:       "Test event",
		Duration:    time.Minute * 30,
		Description: "Event desc",
		NotifyAt:    time.Now().AddDate(0, 0, 1),
	}
	event, err := s.CreateEvent(e)
	require.NoError(t, err)
	require.Equal(t, s.GetEvents()[0], event)
	err = s.DeleteEvent(event.ID)
	require.NoError(t, err)
	require.Equal(t, 0, len(s.GetEvents()))
}

func TestUpdateEvent(t *testing.T) {
	s := New()
	e := storagemodels.Event{
		Title:       "Test event",
		Duration:    time.Minute * 30,
		Description: "Event desc",
		NotifyAt:    time.Now().AddDate(0, 0, 1),
	}
	event, err := s.CreateEvent(e)
	require.NoError(t, err)
	require.Equal(t, s.GetEvents()[0], event)
	require.Equal(t, s.GetEvents()[0].Title, event.Title)
	event.Title = "new title"
	_, err = s.UpdateEvent(*event)
	require.NoError(t, err)
	require.Equal(t, "new title", s.GetEvents()[0].Title)
}

func TestUpdateError(t *testing.T) {
	s := New()
	e := storagemodels.Event{
		Title:       "Test event",
		Duration:    time.Minute * 30,
		Description: "Event desc",
		NotifyAt:    time.Now().AddDate(0, 0, 1),
	}
	_, err := s.UpdateEvent(e)
	require.ErrorIs(t, err, ErrEventNotFound)
}
