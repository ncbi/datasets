package util

import (
	"regexp"
	"strings"

	"github.com/spf13/cobra"
)

type UsageSection struct {
	Section     string
	SectionText string
	Commands    []string
}

type UsageSections []*UsageSection
type UsageCmdMap map[string]*UsageSections

type FlagAdjustment struct {
	RegexpMatch string
	Content     func(reportName, flagContent string) []string
}

type UsageConfig struct {
	Sections    UsageCmdMap
	AdjustFlags []*FlagAdjustment
}

var (
	usageConfig = UsageConfig{Sections: UsageCmdMap{}}
)

func getDefaultSections(cmd *cobra.Command) *UsageSections {
	defaultSection := &UsageSection{
		SectionText: "Available Commands",
	}
	for _, c := range cmd.Commands() {
		defaultSection.Commands = append(defaultSection.Commands, c.Name())
	}
	return &UsageSections{defaultSection}
}

func AddUsageSections(group string, cmdmap *UsageSections) {
	usageConfig.Sections[group] = cmdmap
}

func AddFlagAdjustment(adjust *FlagAdjustment) {
	usageConfig.AdjustFlags = append(usageConfig.AdjustFlags, adjust)
}

func getUsageSections(cmd *cobra.Command) *UsageSections {
	if s, ok := usageConfig.Sections[cmd.Name()]; ok {
		return s
	} else {
		return getDefaultSections(cmd)
	}
}

func getUsageCommands(commands []*cobra.Command, section *UsageSection) (retval []*cobra.Command) {
	for _, sectionCommand := range section.Commands {
		for _, c := range commands {
			if sectionCommand == c.Name() {
				retval = append(retval, c)
			}
		}
	}
	return
}

func expandFields(cmd *cobra.Command, flags string) (retval string) {
	if len(usageConfig.AdjustFlags) == 0 {
		retval = flags
		return
	}
	for _, aflag := range usageConfig.AdjustFlags {
		re := regexp.MustCompile(aflag.RegexpMatch)
		lines := strings.Split(flags, "\n")
		retval = ""
		for _, line := range lines {
			retval += line + "\n"
			if locs := re.FindStringIndex(line); len(locs) == 2 {
				prefix := strings.Repeat(" ", locs[1]+4) + "- "
				content := aflag.Content(cmd.Use, flags)
				for _, cline := range content {
					retval += prefix + cline + "\n"
				}
			}
		}
	}
	return
}

func getUsageTemplate() string {
	return `
{{- "Usage" }}
{{- if .Runnable}}
  {{- "\n  "}}{{.UseLine}}
{{- end}}
{{- if .HasAvailableSubCommands}}
  {{- "\n  "}}{{.CommandPath}} [command]
{{- end}}
{{- if gt (len .Aliases) 0}}
{{- "\n\n"}}Aliases
  {{- "\n  "}}{{.NameAndAliases}}
{{- end}}
{{- if .HasExample -}}
  {{- "\n\n"}}Examples
{{.Example}}
{{- end}}
{{- $cmd := .}}
{{- if .HasAvailableSubCommands}}
  {{- range (usageSections $cmd) -}}
    {{- $section := .}}
    {{- "\n\n"}}{{$section.SectionText}}
    {{- range (usageCommands $cmd.Commands $section) -}}
      {{- if (or .IsAvailableCommand (eq .Name "help")) -}}
	    {{- "\n  " }}{{rpad .Name .NamePadding}} {{.Short}}
      {{- end}}
    {{- end}}
  {{- end}}
{{- end}}
{{- if .HasAvailableLocalFlags}}
{{- "\n\n"}}Flags
  {{- "\n"}}{{expandFields $cmd .LocalFlags.FlagUsages }}
{{- end}}
{{- if .HasAvailableInheritedFlags}}
  {{- "\n\n"}}Global Flags
  {{- "\n"}}{{.InheritedFlags.FlagUsages}}
{{- end}}

{{- if .HasHelpSubCommands}}

Additional help topics:
  {{- range .Commands -}}
    {{- if .IsAdditionalHelpTopicCommand -}}
	  {{- "\n  "}}{{rpad .CommandPath .CommandPathPadding}} {{.Short}}
	{{- end}}
  {{- end}}
{{- end}}

{{- if .HasAvailableSubCommands}}
Use {{.CommandPath}} <command> --help for detailed help about a command.
{{- end}}
`
}

func InitRootCommand(cmd *cobra.Command) {
	cmd.SetUsageTemplate(getUsageTemplate())
}

func init() {
	cobra.AddTemplateFunc("usageSections", getUsageSections)
	cobra.AddTemplateFunc("usageCommands", getUsageCommands)
	cobra.AddTemplateFunc("expandFields", expandFields)
}
