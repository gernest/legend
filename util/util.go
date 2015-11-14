package util

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"strconv"
	"strings"

	"github.com/pborman/uuid"

	"github.com/gernest/legend/core"
)

type JSONHelper struct {
	data map[string]interface{}
}

func NewJSONHelper() *JSONHelper {
	return &JSONHelper{data: make(map[string]interface{})}
}

func (j *JSONHelper) Decode(input io.Reader) error {
	return json.NewDecoder(input).Decode(&j.data)
}

func (j *JSONHelper) HasKey(key string) bool {
	_, ok := j.data[key]
	return ok
}

func (j *JSONHelper) Get(key string) interface{} {
	return j.data[key]
}

func (j *JSONHelper) GetString(key string) string {
	return fmt.Sprint(j.Get(key))
}

func (j *JSONHelper) Set(key string, val interface{}) {
	j.data[key] = val
}

func (j *JSONHelper) Encode() ([]byte, error) {
	rst, err := json.Marshal(j.data)
	if err != nil {
		return nil, err
	}
	return rst, nil
}

func (j *JSONHelper) HasError() bool {
	return j.Get("error") != nil
}

type serviceInfo struct {
	name   string
	status int
	id     string
}

func (s *serviceInfo) ID() string {
	return s.id
}

func (s *serviceInfo) Name() string {
	return s.name
}

func (s *serviceInfo) Status() int {
	return s.status
}

func ParseServiceID(serviceID string) (core.ServiceInfo, error) {
	return parseServiceID(serviceID)
}

func parseServiceID(serviceID string) (core.ServiceInfo, error) {
	sections := strings.SplitN(serviceID, ".", 3)
	switch len(sections) {
	case 1:
		return &serviceInfo{name: sections[0]}, nil
	case 2:
		name := sections[0]
		status, err := strconv.Atoi(sections[1])
		if err != nil {
			return nil, err
		}

		return &serviceInfo{
			name:   name,
			status: status,
		}, nil
	case 3:
		name := sections[0]
		status, err := strconv.Atoi(sections[1])
		if err != nil {
			return nil, err
		}

		return &serviceInfo{
			name:   name,
			status: status,
			id:     sections[2],
		}, nil

	}
	return nil, errors.New("soome fish parsing service identifier")
}

func GenUUID() string {
	return uuid.NewRandom().String()
}

func GetServiceName(info core.ServiceInfo) string {
	return info.Name() + "." + fmt.Sprint(info.Status()) + "." + info.ID()
}
