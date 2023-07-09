Whisper
=======

Author: **[Afaan Bilal](https://afaan.dev)**

## Introduction
**Whisper** is a micro-blogging platform written in Go, React Native and TypeScript. This is the repo for the backend API server written in Go. See [Whisper App](https://github.com/AfaanBilal/whisper-app) for the React Native app.

---

## Screenshots

|        ![SignIn](/screenshots/1-SignIn.png)        |      ![SignUp](/screenshots/2-SignUp.png)       |
| :------------------------------------------------: | :---------------------------------------------: |
| ![ResetPassword](/screenshots/3-ResetPassword.png) |        ![Home](/screenshots/4-Home.png)         |
|       ![Compose](/screenshots/5-Compose.png)       |     ![Explore](/screenshots/6-Explore.png)      |
| ![Notifications](/screenshots/7-Notifications.png) |     ![Profile](/screenshots/8-Profile.png)      |
|          ![Menu](/screenshots/9-Menu.png)          | ![EditProfile](/screenshots/10-EditProfile.png) |

---

## Run with Docker
```bash
docker run --env DB_DSN="root:@tcp(host.docker.internal:3306)/whisper?charset=utf8mb4&parseTime=True&loc=Local" --env PORT=8080 --env VERSION=0.1.0 -p 8080:8080 afaanbilal/whisper
```

This will start the server on port 8080.

## API
See [main.go](./main.go) for routing.

## Contributing
All contributions are welcome. Please create an issue first for any feature request
or bug. Then fork the repository, create a branch and make any changes to fix the bug
or add the feature and create a pull request. That's it!
Thanks!

---

## License
**Whisper** is released under the MIT License.
Check out the full license [here](LICENSE).
