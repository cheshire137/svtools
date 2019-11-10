package models

import (
	"encoding/xml"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestProfessions(t *testing.T) {
	xmlProfessions := `<professions><int>13</int><int>1</int><int>6</int><int>24</int><int>4</int><int>16</int><int>18</int><int>21</int><int>9</int><int>26</int></professions>`
	var professions Professions

	err := xml.Unmarshal([]byte(xmlProfessions), &professions)

	require.NoError(t, err)
	require.Equal(t, []int{13, 1, 6, 24, 4, 16, 18, 21, 9, 26}, professions.Values)
}
