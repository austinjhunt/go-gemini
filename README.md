# Go Gemini

Golang implementation of a cryptocurrency trading bot specifically for the Gemini exchange.

## Usage

1. Navigate to [https://exchange.gemini.com/settings/api](https://exchange.gemini.com/settings/api) and create an API key for yourself. You can name it anything you want; give it the primary scope. Choosing the master scope will not work with this project (for now).

2. Copy [./.env.sample](./.env.sample) to your own `.env` file. Copy and paste the API key shown in the Gemini GUI as the value of `GEMINI_EXCHANGE_API_KEY` and do the same for the API Secret (`GEMINI_EXCHANGE_API_SECRET`).

3. Check the acknowledgement box and close when step 2 is complete.

4. A note about Gemini Exchange API authentication: **Authenticated APIs do not submit their payload as POSTed data, but instead put it in the X-GEMINI-PAYLOAD header**
