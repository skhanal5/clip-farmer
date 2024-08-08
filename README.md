## clip-farmer
Automating the process of selecting, editing, and producing short-form content from existing media sources

### Disclaimer
This project is intended for educational purposes only. The author(s) of this project are not liable for any misuse or damage that may arise from the use of this project. Users of this project are responsible for ensuring that their use complies with all applicable laws, terms of service, and policies of third-party services.

Please use this project responsibly and ethically.

### Local Development

This makes downstream API to TikTok and Twitch's API's. This means
to run the application you will need to register an application with
both Twitch and TikTok and pass in their secret values. 

#### Twitch Credentials
You will need the `TWITCH_CLIENT_ID` and the `TWITCH_CLIENT_OAUTH` values from your Twitch account.
This can be retrieved from your browser's console after authenticating into Twitch.

#### TikTok Credentials
Next, you will need the `TIKTOK_CLIENT_KEY` and `TIKTOK_CLIENT_SECRET` values from
your registered application on the [TikTok Developer Dashboard](https://developers.tiktok.com/apps). When registering, your
application it is recommended that you make a Sandbox account.

In addition to the above, the application requires you to register an `Target User` so that you can post content
on that account's behalf. We make use of the `TIKTOK_CLIENT_KEY` and `TIKTOK_CLIENT_SECRET` to fetch an OAuth
token on behalf of that user.

#### Environment Variables

Altogether, you will need to pass in the following secret values in the config.yaml file:

```
secrets:
  twitch:
    client-id: 
    client-oauth:
  tiktok:
    client-key: 
    client-secret:
query:
  twitch-creator: stableronaldo
```

#### Running the Application
From there, you can run the `main.go` file using the command `go run ./cmd` to start the application