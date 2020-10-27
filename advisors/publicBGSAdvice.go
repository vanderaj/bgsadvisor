package advisors

import (
	"github.com/bwmarrin/discordgo"
)

// PublicBGSAdvice gives positive urgent actions for controlled stations only
// This option is safe for public BGS to uninitiated players
// - Will not suggest raiding
// - Will not suggest negative actions
// - Does not provide intelligence support for undermining actions
// - Will not suggest taking over a station
func PublicBGSAdvice(s *discordgo.Session, m *discordgo.MessageCreate) string {

	// obtain tick date

	tick := apihelper.getTick()

	// obtain faction data for OpIda

	// obtain system data for each system we're present in

	// Create a list of stale systems that need to be scanned

	// iterate over all stations for urgent actions, e.g. famine, drought, low INF, etc

	// iterate over controlled stations and create a list of positive actions for public consumption

	return "```diff\n+ Public BGS Advice\nWell hello there```"
}
