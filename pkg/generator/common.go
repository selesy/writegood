package generator

import (
	"errors"
	"io/ioutil"
	"net/http"
	"regexp"
	"strings"

	log "github.com/sirupsen/logrus"
)

func LineSlice1(lines []string, start, end string, startInclusive, endInclusive bool) ([]string, error) {
	startExp, err := regexp.Compile(start)
	if err != nil {
		return nil, err
	}

	endExp, err := regexp.Compile(end)
	if err != nil {
		return nil, err
	}

	type State int

	const (
		awaitingStart State = iota
		awaitingStartInclusive
		awaitingEnd
		awaitingEndInclusive
	)

	state := awaitingStart
	if startInclusive {
		state = awaitingEndInclusive
	}

	endState := awaitingEnd
	if endInclusive {
		endState = awaitingEndInclusive
	}

	matchExp := startExp
	out := []string{}

ToEnd:
	for _, line := range lines {
		matches := matchExp.Match([]byte(line))

		switch state {
		case awaitingStartInclusive:
			out = append(out, line)
			continue
		case awaitingStart:
			if matches {
				state = endState
				matchExp = endExp
			}
		case awaitingEndInclusive:
			out = append(out, line)
			continue
			// if matches {
			// 	break
			// }
		case awaitingEnd:
			if matches {
				break ToEnd
			}

			out = append(out, line)
		}
	}

	return out, nil
}

func LineSlice(lines []string, start, end string, startInclusive, endInclusive bool) ([]string, error) {
	startExp, err := regexp.Compile(start)
	if err != nil {
		return nil, err
	}

	endExp, err := regexp.Compile(end)
	if err != nil {
		return nil, err
	}

	out, collecting := []string{}, false
	for _, line := range lines {
		if !collecting && startExp.Match([]byte(line)) {
			collecting = true
		}
		if startExp.Match([]byte(line)) && !startInclusive {
			continue
		}
		if endExp.Match([]byte(line)) && !endInclusive {
			break
		}
		if collecting && endExp.Match([]byte(line)) {
			collecting = false
		}
		if collecting {
			out = append(out, line)
		}
	}
	return out, nil
}

func LineSliceExclusive(lines []string, start, end string) ([]string, error) {
	return LineSlice1(lines, start, end, false, false)
}

func RemoteFileLines(url string) ([]string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	log.Debug("Response status: ", resp)

	if resp.StatusCode != 200 {
		return nil, errors.New(resp.Status)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	log.Debug("Response body: ", string(body))

	return strings.Split(string(body), "\n"), nil
}

func TrimList(list []string) {
	for idx, item := range list {
		list[idx] = strings.Trim(item, "` '\t\n\"")
		log.Info("Item: ", list[idx])
	}
}
