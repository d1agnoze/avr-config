package misc

import (
	"os"
	"strings"

	"github.com/charmbracelet/lipgloss"
)

func Summary(err []string) string {
	styleFILE := lipgloss.NewStyle().Underline(false).Foreground(lipgloss.Color("#38ef7d"))
	fileList := lipgloss.NewStyle().Foreground(lipgloss.Color("#A6CF98")).PaddingTop(1)
	styleErr := lipgloss.NewStyle().Underline(false).Foreground(lipgloss.Color("#11998e")).PaddingTop(1)
	errList := lipgloss.NewStyle().Foreground(lipgloss.Color("#C70039")).PaddingTop(0)
	res := lipgloss.JoinVertical(
		lipgloss.Left,
		styleFILE.Render("Files Summary:\n"),
		fileList.Render(FileCheck()),
		styleErr.Render("Error Report:"),
		errList.Render(joinStrings(err)),
		lipgloss.NewStyle().Foreground(lipgloss.Color("#D2DE32")).Render("Thank you for using !!! "+CONGRAT),
		lipgloss.NewStyle().PaddingTop(3).Render("---press <ESC> or Enter to exit---"),
	)
	return res
}
func FileCheck() string {
	res := ""
	pwd, err := os.Getwd()
	var conf, mf string
	if err == nil {
		mf = pwd + "/Makefile"
		conf = pwd + "/.vscode/c_cpp_properties.json"
		if FileExists(mf) {
			res += mf + "  <EXISTS ✅>\n"
		}
		if FileExists(conf) {
			res += conf + "  <EXISTS ✅>"
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
func FileExists(filename string) bool {
	_, err := os.Stat(filename)
	if err == nil {
		return true // File exists
	}
	if os.IsNotExist(err) {
		return false // File does not exist
	}
	return false // Error occurred while checking file existence
}
