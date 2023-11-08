package misc

import (
	"os"
	"strings"

	"github.com/charmbracelet/lipgloss"
)

func Summary(err []string) string {
	styleFILE := lipgloss.NewStyle().Underline(false).Foreground(lipgloss.Color("#38ef7d")).PaddingTop(5)
	fileList := lipgloss.NewStyle().Foreground(lipgloss.Color("#A6CF98")).PaddingTop(1)
	styleErr := lipgloss.NewStyle().Underline(false).Foreground(lipgloss.Color("#11998e")).PaddingTop(5)
	errList := lipgloss.NewStyle().Foreground(lipgloss.Color("#C70039")).PaddingTop(1)
	res := lipgloss.JoinVertical(
		lipgloss.Left,
		styleFILE.Render("Files Summary:\n"),
		fileList.Render(FileCheck()),
		styleErr.Render("Error Report:"),
		errList.Render(joinStrings(err)),
		"Thank you for using !!! "+CONGRAT,
		"---press <ESC> or Enter to exit---",
	)
	return res
}
func FileCheck() string {
	res := ""
	pwd, err := os.Getwd()
	if err != nil {
		mf := pwd + "./Makefile"
		conf := pwd + "./.vscode/c_cpp_properties.json"
		if fileExists(mf) {
			res += mf + "EXISTS ✔\n"
		}
		if fileExists(conf) {
			res += conf + "EXISTS ✔\n"
		}
	}
	return res
}

func joinStrings(rp []string) string {
	for i := 0; i < len(rp); i++ {
		rp[i] = "👇\t" + rp[i]
	}
	return strings.Join(rp, "\n")
}
func fileExists(filename string) bool {
	_, err := os.Stat(filename)
	if err == nil {
		return true // File exists
	}
	if os.IsNotExist(err) {
		return false // File does not exist
	}
	return false // Error occurred while checking file existence
}
