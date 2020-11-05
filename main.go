package main

import (
	"fmt"
	"strings"
	"sync"
	"time"
)

var (
	start time.Time
	wg    sync.WaitGroup
)

func main() {
	input := `"Abstract Violence"
	"Alcohol and Tobacco Reference"
	"Alcohol Reference"
	"Alcohol Use"
	"Blood"
	"Blood and Gore"
	"Cartoon Violence"
	"Comic Mischief"
	"Contents for Different Age Groups"
	"Criminal Technique Instructions"
	"Crude Humor"
	"Discrimination"
	"Diverse Content: Discretion Advised"
	"Drug and Alcohol Reference"
	"Drug Reference"
	"Drugs"
	"Drug Use"
	"Explicit Language"
	"Explicit Violence"
	"Extreme Violence"
	"Fantasy Violence"
	"Fear"
	"Gambling"
	"General"
	"Horror"
	"Horror Themes"
	"Illegal Drugs"
	"Impacting Content"
	"Implied Violence"
	"Inappropriate Language"
	"Intense Violence"
	"Language"
	"Legal Drugs"
	"Mature Themes"
	"Mild Blood"
	"Mild Fantasy Violence"
	"Mild Language"
	"Mild Sexual Themes"
	"Mild Suggestive Themes"
	"Mild Swearing"
	"Mild Themes"
	"Mild Violence"
	"Moderate Violence"
	"Nudity"
	"Nudity/Eroticism"
	"Occasional Swearing"
	"Online Interactivity"
	"Parental Guidance Recommended"
	"Partial Nudity"
	"Rare Scary Situations"
	"Rated for 12+"
	"Rated for 3+"
	"Real Gambling/Paid Contests"
	"Referenced Sexual Violence"
	"Scary Content"
	"Scary Scenes"
	"Sex"
	"Sex/Eroticism"
	"Sex Scenes"
	"Sexual Content"
	"Sexual Innuendo"
	"Sexuality"
	"Sexual References"
	"Sexual Themes"
	"Sexual Violence"
	"Shock and/or Horror Content"
	"Shop/Streaming Service - Contents May Vary"
	"Simulated Gambling"
	"Strong Language"
	"Strong Nudity"
	"Strong Sexual Content"
	"Strong Violence"
	"Suggestive Themes"
	"Tobacco Reference"
	"Tobacco Use"
	"Use of Alcohol"
	"Use of Alcohol and Tobacco"
	"Use of Alcohol/Tobacco"
	"Use of Drugs"
	"Use of Drugs and Alcohol"
	"Use of Tobacco"
	"Very Mild Themes"
	"Very Mild Violence"
	"Violence"
	"Violent References"
	"Warning – content has not yet been rated. Unrated apps may potentially contain content appropriate for mature audiences only."
	"War Themes"`

	splitted := strings.Split(input, `"`)
	var results []string
	for _, s := range splitted {
		ts := strings.TrimSpace(s)
		if ts != "" {
			results = append(results, output(ts))
		}
	}
	fmt.Println(strings.Join(results[:], "\n"))
}

func clean(input string) string {
	input = strings.ToUpper(input)
	replacer := strings.NewReplacer(".", "", "-", "_", " ", "_", "\\", "", "/", "", "'", "")
	return replacer.Replace(input)
}

func output(input string) string {
	return fmt.Sprintf(`,[
		'title' => '%s',
		'id' => '%s',
		'tags' => [],
		'description' => ''
		]`, input, clean(input))
}
