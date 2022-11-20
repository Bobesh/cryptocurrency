## API SERVER

- contains two endpoints

1. `/cryptocurrency` with query parameter `currency` that returns currency rates for some cryptocurrencies
2. `/find_hash` with query parameter `hash` hat finds if hash is within files that are randomly generated on app startup

## HOW TO RUN

- you should run app by typing `docker-compose -f docker-compose.yaml up server`