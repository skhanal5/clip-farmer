# clip-farmer
A cli tool that simplifies the process of selecting, editing, and producing short-form content from existing media sources

## Disclaimer
This project is intended for educational purposes only. The author(s) of this project are not liable for any misuse or damage that may arise from the use of this project. Users of this project are responsible for ensuring that their use complies with all applicable laws, terms of service, and policies of third-party services. 

Please use this project responsibly and ethically.  

## Usage

### Build
Run `go build -o clip-farmer.exe` to build the executable in the project root if it already doesn't exist.

### Config

You will need secret values from Twitch and TikTok to use this application, refer to [local development](#local-development)

#### TikTok Config
To set the TikTok environment variables in the app config use the following command:
```
.\clip-farmer.exe config tiktok --client-key [client-key] --client-secret [client-secret]
```

#### Twitch Config
To set the Twitch environment variables in the app config use the following command:
```
.\clip-farmer.exe config twitch --client-id [client-id] --client-oauth [client-oauth]
```

**Note**: A side effect of running either command is that it will generate a `config.yaml` file in the project root directory.The values in this file are loaded into memory during the initialization of the "root" command. Producing this file makes it so that you do not have to set these values each time you want to use the CLI.

### Fetch

There are two subcommands under the fetch command, `oauth` and `clips`

#### OAuth
You can fetch the TikTok OAuth token for the account you'd like to post on
with the following command:
```
.\clip-farmer.exe fetch oauth tiktok
```
This will produce a link where you can authorize the application to access your account with the necessary
scopes. Once the authorization flow is complete, it will fetch the OAuth token and set it as an environment variable
for later use. One side effect is that this will also produce a `tiktok_oauth_resp.json` file which will allow us
to re-use existing tokens and refresh tokens on expiration.

#### Clips
You can fetch clips from Twitch from a given user with the following command:
```
.\clip-farmer.exe fetch clips twitch --user [twitch-username]
```
This will download clips of that user under the `clips/[username]/` directory in the project root.

Alternatively, you can pass in query parameters to filter clips by **time period** and by a **sort** filter. For example:
```
.\clip-farmer.exe fetch clips twitch -u [username] -p [period-of-time] -s [sort-filter]
```

### Edit
**Prerequisite:** To edit clips, you must have ffmpeg installed on your local environment. I am not distributing this.

You can edit downloaded clips by using the edit command and the type of edit you would like. To edit in bulk:
```
.\clip-farmer.exe edit --directory [directory] --output [output-directory] --blurred
```

Alternatively, to edit an individual file
```
.\clip-farmer.exe edit --file [file] --output [output-directory] --blurred
```

This library supports one edit option which is `blurred`. This will take a clip and produce a 1080x1920 video where the clip is centered in the middle of the screen and overlayed ontop of a blurred background. The blurred background is the same clip but stretched out to fit the resolution. 

### Post

**Note**: for this command to work as intended, you must invoke the fetch oauth command as a prerequisite at least
once. This is because posting requires a valid OAuth token.

You can post a video onto TikTok with the following command:
```
.\clip-farmer.exe post tiktok --file [path-to-video-file]       
```

Alternatively, you can post all videos under a directory with the following command:
```
.\clip-farmer.exe post tiktok --directory [path-to-video-file] 
```

### Clean
This command will clean up any directories containing clips on your local filesystem.
```
 .\clip-farmer.exe clean --directory [directory] 
```
Alternatively, you can add a filter to delete all videos that are less than or equal to a `duration`:
```
 .\clip-farmer.exe clean --directory [directory] --duration [seconds]
```
**Note:** To clean clips with filter, you must have ffmpeg/ffprobe installed on your local environment.


### Help
You can get help when interacting with the cli with either the `--help` or `-h` flags at any command-level.

For example, if you are having setting your TikTok configuration, you can run:
```
    .\clip-farmer.exe config tiktok -h
```

You will get an output like so:
```
Configure TikTok environment variables

Usage:
  clip-farmer config tiktok [flags]

Flags:
  -k, --client-key string      Set the client-key of the TikTok app that we want to connect to.
  -s, --client-secret string   Set the client-secret of the TikTok app that we want to connect to.
  -h, --help                   help for tiktok
```

Or, if you are having trouble posting specifically to TikTok
```
.\clip-farmer.exe post tiktok -h  
```
You will get an output like so:
```
Post short-form content onto TikTok

Usage:
  clip-farmer post tiktok [flags]

Flags:
  -d, --directory string   Path to the directory with all the videos that we want to post onto TikTok.
  -f, --file string        Path to the file containing the video that we want to post onto TikTok.
  -h, --help               help for tiktok
```

### Local Development

This application makes downstream API calls to TikTok and Twitch's API's. Thus, you will
need the appropriate credentials from both services.

#### Twitch Credentials
You will need the `TWITCH_CLIENT_ID` and the `TWITCH_CLIENT_OAUTH` values from your Twitch account.
This can be retrieved from your browser's console after authenticating into Twitch.

#### TikTok Credentials
Next, you will need the `TIKTOK_CLIENT_KEY` and `TIKTOK_CLIENT_SECRET` values from
your registered application on the [TikTok Developer Dashboard](https://developers.tiktok.com/apps). When registering, your
application it is recommended that you make a Sandbox account. In your register application, allocate a `Target User` so that you can post content
on that account's behalf. We make use of the `TIKTOK_CLIENT_KEY` and `TIKTOK_CLIENT_SECRET` to fetch an OAuth
token on behalf of that user.
