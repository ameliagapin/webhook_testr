# webhook_testr

Simple HTTP server for testing webhook development

## Usage

1. Install [ngrok](https://ngrok.com/download)
2. Start ngrok: `ngrok http 8080`
3. Take the `HTTPS` link from ngrok and add `/webhook` to the end for your webhook URL
4. Run `webhook_testr`
5. Set up your webhook with your endpoint
6. Trigger your webhook
7. Check the logs of both `ngrok` and `webhook_testr`
