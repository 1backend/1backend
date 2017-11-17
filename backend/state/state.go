package state

import (
	"errors"
	"time"

	"github.com/go-redis/redis"
)

// State stores anythin that you would store in memory
type State struct {
	redisClient *redis.Client
}

func NewState(redisClient *redis.Client) *State {
	return &State{
		redisClient: redisClient,
	}
}

func (s *State) Port(author, projectName string) (int, error) {
	port, err := s.redisClient.Get("port:" + author + "_" + projectName).Int64()
	return int(port), err
}

func (s *State) SetPort(author, projectName string, port int) error {
	return s.redisClient.Set(author+"_"+projectName, port, 0).Err()
}

// Use this to remove the service after stopping the container
func (s *State) MarkAsDown(author, projectName string) error {
	return s.redisClient.Set("isup:"+author+"_"+projectName, 0, 0).Err()
}

// Use this to remove the service after stopping the container
func (s *State) MarkAsUp(author, projectName string) error {
	return s.redisClient.Set("isup:"+author+"_"+projectName, 1, 0).Err()
}

func (s *State) IsUp(author, projectName string) (bool, error) {
	val, err := s.redisClient.Get("isup:" + author + "_" + projectName).Int64()
	if err != nil {
		return false, err
	}
	return val == 1, nil
}

// Use this to remove the service after stopping the container
func (s *State) MarkAsUnderStartup(author, projectName string) error {
	return s.redisClient.Set("starting:"+author+"_"+projectName, 1, 0).Err()
}

// Use this to remove the service after stopping the container
func (s *State) MarkAsNotUnderStartup(author, projectName string) error {
	return s.redisClient.Set("starting:"+author+"_"+projectName, 0, 0).Err()
}

func (s *State) IsUnderStartup(author, projectName string) (bool, error) {
	val, err := s.redisClient.Get("starting:" + author + "_" + projectName).Int64()
	if err != nil {
		return false, err
	}
	return val == 1, nil
}

// Returns wether you can use the service or has to start it
func (s *State) LastCallIn(author, projectName string, d time.Duration) (bool, error) {
	val, err := s.redisClient.Get("lastcall:" + author + "_" + projectName).Int64()
	if err != nil {
		return false, err
	}
	return time.Now().Sub(time.Unix(val, 0)) < d, nil
}

// Call this when proxying a request
func (s *State) SetLastCall(author, projectName string) error {
	return s.redisClient.Set("lastcall:"+author+"_"+projectName, time.Now().Unix(), 0).Err()
}

func (s *State) Decrement(tokenId string) error {
	val, err := s.redisClient.Get("quota:" + tokenId).Int64()
	if err != nil {
		return err
	}
	if val <= 0 {
		return errors.New("Quote exceeded")
	}
	return s.redisClient.Decr("quota:" + tokenId).Err()
}

func (s *State) DecrementBy(tokenId string, amt int64) error {
	val, err := s.redisClient.Get("quota:" + tokenId).Int64()
	if err != nil {
		return err
	}
	if val <= amt {
		return errors.New("Quota would be less than 0")
	}
	return s.redisClient.DecrBy("quota:"+tokenId, amt).Err()
}

func (s *State) IncrementBy(tokenId string, amt int64) error {
	return s.redisClient.IncrBy("quota:"+tokenId, amt).Err()
}

// Use at initialization
func (s *State) SetQuota(tokenId string, quota int64) error {
	return s.redisClient.Set("quota:"+tokenId, quota, 0).Err()
}

// Use at initialization
func (s *State) GetQuota(tokenId string, quota int64) error {
	return s.redisClient.Set("quota:"+tokenId, quota, 0).Err()
}
