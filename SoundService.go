package ozsndqueue

type SoundService interface {
	Play(fileUri string) error
}
