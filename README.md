# Elite Dangerous BGS Advisor Discord Bot

BGS Advisor is a Discord Bot that is useful for factions who want to have the general public, trial, or early stage BGS operatives assist them, but don't want to expose their long term planning strategy, upset player faction allies, or do negative things to their influence. The sort of thing anyone can work out by viewing their faction's page on EliteBGS, basically.

Future versions of this bot will encompass more functionality to allow BGS planners to create better "orders of the day" to maintain a desired state, but these are currently unimplemented.

## What is BGS Advisor

BGS Advisor's role is to create equilibrium by providing advice that promotes positive actions for the faction, whilst not upsetting player faction allies or rivals in non-controlled systems. It does this by prioritizing common sense "orders" that will be the most beneficial to harnessing unskilled, trial or early career BGS operatives, or otherwise undirected BGS activity, and take the load off BGS Planners so they can concentrate on long term plans for galaxy domination.

BGS Advisor pulls in data from EliteBGS once a day, and analyzes it to provide easy to follow recommendations on:

* Urgent actions to get out of negative or harmful states in both controlled and non-controlled systems
* Priority actions to keep controlled factions system influence over 60%
* Aims to maintain perma-boom, rather than expansion
* Aims to maintain between 10% and 25% system influence in non-controlled systems, and capped at 20% behind the controlling faction to prevent ally / rival misunderstandings (i.e. if a controlling faction is at 35%, BGS Advisor will only offer positive action advice to get above 10% and stay below 15%)
* Aims to stay out of the last faction spot in non-controlled systems to avoid retreat states. If the last faction spot INF is say 26% and the controlling faction is 45%, BGS Advisor will note this to the BGS Planner, as other actions will need to be planned, including diplomacy and so on that are really far outside this bot's tactical advice capabilities.
* Alerts to election or war states, but doesn't suggest winning them unless given permission to do so.
* For controlled boom systems, suggests in demand, high profit commodities for trade loops to maximize trading profits

## How is BGS Advisor different to BGSBot

EliteBGS has an official BGSBot that provides detailed raw information about systems. There is no analysis, so BGS planners need to interpret BGSBot's output, and occasionally go check out a faction's security or economic state directly to create more detailed plans.

## What this bot is *NOT*

BGS Advisor is not a BGS intelligence agency. Yet. A future version of this bot might add features to work out fast movers, systems likely to get into trouble soon, expansion spheres, or where enemy action has likely occurred, but these features are primarily focused on BGS planners, rather than entry level BGS operatives.

More skilled advice commonly used by BGS planners to expand to specific systems, fight wars in CZs, killing other factions, smuggle black market items, retreat actions, and so on are out of the scope of this bot at this time. A future version might add these features to specific guild roles so that more skilled BGS operatives can take the full set of relevant actions if those actions are desirable. Having unskilled BGS operatives blithely fight in CZs that they shouldn't be in or trample over allies or rivals INF, could easily spark a BGS war so this is unlikely to ever be offered as a public option.

## Instructions

BGS Advisor is in its earliest stages of development, so expect change.

### Set up a BGS Advisor on your local system

Each faction should run their own BGS Bot.

* Clone the repo on a VPS or other system that is highly available.
* Install the [go tool chain](https://golang.org/doc/install) as you're going to be compiling it
* Build BGS Advisor

```go
go build main.go
```

It should compile cleanly as it's very self-contained and only has one dependency.

This is development code. One day, you might be able to use Docker or a pre-packaged cloud option if you don't feel up to compiling and running your own BGS Bot.

### Invite BGS Advisor to your server

Development version - follow the Discord advice on inviting a random bot to your server.

Production version - TBA.

### Configure your faction

At this time, BGS Advisor is stateless, so BGS Advisor administrators are Discord guild (server) administrators with the manage server permission. There are only a few administration commands, such as setting your faction, and they will need to be re-entered if BGS Advisor is restarted.

Once BGS Advisor is established in a channel, guild administrators can set the faction:

```html
!bgsadvisor faction set <factionid>
```

### Updating stale systems after scanning, and caching behavior

BGS only changes once per day after a tick, so constantly asking about new results is harmful to EliteBGS performance. Advice for stale systems may change after being scanned using ED Market Connector and data uploaded to EDDN, so you can flush the BGS Advisor cache, but in general, BGS data is only requested once per day.

BGS Advisor will give advice on stale systems where the faction is present when you ask it for advice. You can also get the list of systems to be scanned by:

```html
!bgsadvisor cache stale
```

Say you've now scanned a stale system, you can force a reload of BGS Advisor cached data by:

```html
!bgsadvisor cache flush
```

#### Cache Architecture

BGS Advisor caches EliteBGS API results upon first use, keeping the results until the tick is likely to have changed again, the bot is restarted, or the cache is flushed.

Once a tick is cached, around the next tick time, BGS Advisor will test the tick every five minutes until the tick has happened. Once that happens, the cache is flushed, and the new tick time cached. The faction and system cache will be repopulated upon first use.

#### Note on cache performance

In memory caching is extremely fast, so outside of Discord API limits, there should be no performance issues once data has been obtained. As no persistent state is maintained, the first few queries will take slightly longer or may even timeout while data is cached from EliteBGS API. This is normal. If you get a timeout, just re-issue the command again after a few minutes in case of Discord or EliteBGS API limits.
