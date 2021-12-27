package tournament

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"sort"
	"strings"
)

type Team struct {
	Name   string
	Played int
	Won    int
	Lost   int
	Drawn  int
	Points int
}

var teamMap map[string]*Team

func UpdateTeamScores(strA, strB, result string) error {
	errInvalidMatchResult := errors.New("invalid match result")
	teamA := teamMap[strA]
	if teamA == nil {
		teamA = &Team{Name: strA}
		teamMap[strA] = teamA
	}
	teamB := teamMap[strB]
	if teamB == nil {
		teamB = &Team{Name: strB}
		teamMap[strB] = teamB
	}
	teamA.Played++
	teamB.Played++
	switch result {
	case "win":
		teamA.Won++
		teamB.Lost++
		teamA.Points += 3
	case "loss":
		teamA.Lost++
		teamB.Won++
		teamB.Points += 3
	case "draw":
		teamA.Drawn++
		teamB.Drawn++
		teamA.Points++
		teamB.Points++
	default:
		return errInvalidMatchResult
	}
	return nil
}

func GetSortedResults() []*Team {
	sortedTeams := make([]*Team, 0)
	for _, team := range teamMap {
		sortedTeams = append(sortedTeams, team)
	}
	sort.Slice(sortedTeams, func(i, j int) bool {
		if sortedTeams[i].Points != sortedTeams[j].Points {
			return sortedTeams[i].Points > sortedTeams[j].Points
		} else {
			return sortedTeams[i].Name < sortedTeams[j].Name
		}
	})
	return sortedTeams
}

func WriteResults(writer io.Writer, results []*Team) error {
	format := "%-30s |%3d |%3d |%3d |%3d |%3d\n"
	headerLine := "Team                           | MP |  W |  D |  L |  P\n"
	_, err := io.WriteString(writer, headerLine)
	if err != nil {
		return err
	}
	for _, team := range results {
		teamData := fmt.Sprintf(format, team.Name, team.Played, team.Won, team.Drawn, team.Lost, team.Points)
		_, err = io.WriteString(writer, teamData)
		if err != nil {
			return err
		}
	}
	return nil
}

func Tally(reader io.Reader, writer io.Writer) error {
	teamMap = make(map[string]*Team)
	scanner := bufio.NewScanner(reader)
	errIncorrectDataFormat := errors.New("the file data is not in the expected format")
	for scanner.Scan() {
		matchDetails := strings.TrimSpace(scanner.Text())
		if matchDetails == "" || strings.Index(matchDetails, "#") == 0 {
			continue
		}
		parts := strings.Split(matchDetails, ";")
		if len(parts) != 3 {
			return errIncorrectDataFormat
		}
		err := UpdateTeamScores(parts[0], parts[1], parts[2])
		if err != nil {
			return err
		}
	}
	results := GetSortedResults()
	err := WriteResults(writer, results)
	if err != nil {
		return err
	}
	return nil
}
