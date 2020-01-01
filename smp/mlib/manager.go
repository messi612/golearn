package mlib

import "errors"

type MusicEntry struct {
	Id     string
	Name   string
	Artist string
	Source string
	Type   string
}

type MusicManager struct {
	musics []MusicEntry
}

func NewMusicManager() *MusicManager {
	return &MusicManager{make([]MusicEntry, 0)}
}

func (m *MusicManager) Len() int {
	return len(m.musics)
}

func (m *MusicManager) Get(index int) (music *MusicEntry, err error) {
	if index < 0 || index >= len(m.musics) {
		return nil, errors.New("Index out of range")
	}
	return &m.musics[index], nil
}

func (m *MusicManager) Find(name string) (*MusicEntry, int) {
	if len(m.musics) == 0 {
		return nil, -1
	}

	for idx, mu := range m.musics {
		if mu.Name == name {
			return &mu, idx
		}
	}
	return nil, -1
}

func (m *MusicManager) Add(music *MusicEntry) {
	m.musics = append(m.musics, *music)
}

func (m *MusicManager) Remove(index int) *MusicEntry {
	if index < 0 || index >= len(m.musics) {
		return nil
	}
	removeMusic := &m.musics[index]
	m.musics = append(m.musics[:index], m.musics[index+1:]...)
	return removeMusic
}

func (m *MusicManager) RemoveByName(source string) *MusicEntry {
	removeMusic, index := m.Find(source)
	m.musics = append(m.musics[:index], m.musics[index+1:]...)
	return removeMusic
}
