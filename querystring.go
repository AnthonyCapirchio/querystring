package querystring

import (
	"regexp"
	"strings"
)

type QueryString struct {
	Data map[string]string
	Path string
}

const QueryStringDelimiter = "?"
const ParametersSeparator = "&"
const ParametersSeparatorPattern = ";|&"
const KeyValueSeparator = "="

func MapToQueryString(Data map[string]string) string {
	lines := []string{}
	for index, row := range Data {
		lines = append(lines, index+KeyValueSeparator+row)
	}

	return strings.Join(lines, ParametersSeparator)
}

func (q *QueryString) BasePath(Path string) *QueryString {
	splitedPath := strings.Split(Path, QueryStringDelimiter)

	if len(splitedPath) == 2 && splitedPath[1] != "" {
		q.HydrateFromQueryString(splitedPath[1])
	}

	q.Path = splitedPath[0]

	return q
}

func (q *QueryString) HydrateFromQueryString(arguments string) *QueryString {

	a := regexp.MustCompile(ParametersSeparatorPattern)

	argumentPairs := a.Split(arguments, -1)

	for _, pair := range argumentPairs {
		splitedPair := strings.Split(pair, KeyValueSeparator)
		if splitedPair[0] == "" {
			continue
		}
		q.SetParameter(splitedPair[0], splitedPair[1])
	}

	return q
}

func (q *QueryString) Get(key string) string {
	return q.Data[key]
}

func (q *QueryString) SetParameter(key, value string) *QueryString {
	q.Data[key] = value
	return q
}

func (q *QueryString) SetParameters(Data map[string]string) *QueryString {
	q.Data = Data
	return q
}

func (q *QueryString) Build() string {
	return q.Path + MapToQueryString(q.Data)
}

func CreateInstance(basePath string) *QueryString {
	q := QueryString{
		Data: map[string]string{},
	}
	q.BasePath(basePath)

	return &q
}
