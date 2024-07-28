## clip-farmer
Automating the process of selecting, editing, and producing short-form content from existing media sources

### Local Development

This makes downstream API to TikTok and Twitch's API's. This means
to run the application you will need to register an application with
both Twitch and TikTok and pass in their secret values. 

#### Twitch Credentials
You will need the `TWITCH_CLIENT_ID` and the `TWITCH_CLIENT_SECRET` values
from your register application on the [Twitch Developer Dashboard](https://dev.twitch.tv/console). Using these values, the application
will handle fetching OAuth tokens and sending authorized requests on your behalf.

This application as of now only uses `GET Users` and `GET Clips` from Twitch's API. 

#### TikTok Credentials
Next, you will need the `TIKTOK_CLIENT_KEY` and `TIKTOK_CLIENT_SECRET` values from
your registered application on the [TikTok Developer Dashboard](https://developers.tiktok.com/apps). When registering, your
application it is recommended that you make a Sandbox account.

In addition to the above, the application requires you to register an `Target User` so that you can post content
on that account's behalf.

#### Environment Variables

Altogether, you will need to pass in the following secret values:

```
// found on twitch developer dashboard
TWITCH_CLIENT_ID = ""
TWITCH_CLIENT_SECRET = ""

// found on tiktok developer dashboard
TIKTOK_CLIENT_KEY = ""
TIKTOK_CLIENT_SECRET = ""
```

#### Running the Application
From there, you can run the `main.go` file using the command `go run ./cmd` to start the application