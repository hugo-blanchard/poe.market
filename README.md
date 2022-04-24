# POE.MARKET

poe.market is a fan-made app that is not endorsed by Grinding Gear Games

It's purpose is to build a database of items that may have been sold by reading the stash stream from `pathofexile.com/api/stashes`.

Disclaimer: The app is a bit messy at the moment as it's the first app I made in Go and was meant as an exercise more than anything else. If I find that some peoples are interested in the app or if I just feel like it I might work on it further, you can look at the Roadmap at the bottom if you want to know more about where the app is going.

## How to run the app locally

1. Clone this repository.
2. Make sure you have Go installed on your machine.
3. Make sure you have a postgres database available, or if you have Docker installed you can edit and use the Makefile in this repo.
4. Edit the poe.market.config.json, add you postgres connection under `"postgres"` and add items that you want to follow under `"items"`.
5. Run the migrate command in the Makefile or create the tables in the way you prefer.
6. You can now launch the app with `go run .` or compile it to an executable, everything should be working.

## What does the app do

The app reads from the `pathofexile.com/api/stashes` starting at the `next_change_id` that poe.ninja is at.

Everytime it gets a stash from the api, it looks if there already are items in the `items` table under that `stash_id`.

Items that came from the api but are not in the `items` table are added to the table, and items that are in the `items` table but are not comming from the api anymore are assumed to be sold and are moved to the `solditems` table.

Every unique item types from the `solditems` table can be queried with a `GET http://{url_of_the_app}/solditems` and every registered price of a specific item can be queried with a `GET http://{url_of_the_app}/solditems/{item_name}`.

I'm aware that this can be easily exploited by adding and removing an item from a stash periodically, but a ban system scan be implemented using the fact that `item_id`s are consistant through time even after being moved out of a stash.

## Roadmap

1. Clean up the codebase before expanding on anything.
2. Add a ban system to keep `sold_item` data as reliable as possible.
3. Add support for different types of valuables, at the moment only unique items are supported and with no regards to their ilvl, their rolls and them being identified or not.
4. Build a frontend app to serve the data.
5. Optimization for lower cost hosting.