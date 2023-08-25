Whisper
=======

Author: **[Afaan Bilal](https://afaan.dev)**

## Introduction
**Whisper** is a micro-blogging platform written in Go, React Native and TypeScript. This is the repo for the backend API server written in Go. See [Whisper App](https://github.com/AfaanBilal/whisper-app) for the React Native app.

---

## Screenshots

|              ![SignIn](https://raw.githubusercontent.com/AfaanBilal/whisper-app/master/screenshots/SignIn.webp)              |             ![SignUp](https://raw.githubusercontent.com/AfaanBilal/whisper-app/master/screenshots/SignUp.webp)             |        ![Reset Password](https://raw.githubusercontent.com/AfaanBilal/whisper-app/master/screenshots/ResetPassword.webp)         |
| :--------------------------------------------------------------------------------------------------------------------------: | :------------------------------------------------------------------------------------------------------------------------: | :------------------------------------------------------------------------------------------------------------------------------: |
|             ![Welcome](https://raw.githubusercontent.com/AfaanBilal/whisper-app/master/screenshots/Welcome.webp)             |               ![Home](https://raw.githubusercontent.com/AfaanBilal/whisper-app/master/screenshots/Home.webp)               |               ![Compose](https://raw.githubusercontent.com/AfaanBilal/whisper-app/master/screenshots/Compose.webp)               |
|             ![Explore](https://raw.githubusercontent.com/AfaanBilal/whisper-app/master/screenshots/Explore.webp)             |     ![Explore Search](https://raw.githubusercontent.com/AfaanBilal/whisper-app/master/screenshots/Explore-Search.webp)     |          ![User Profile](https://raw.githubusercontent.com/AfaanBilal/whisper-app/master/screenshots/UserProfile.webp)           |
| ![Notifications Empty](https://raw.githubusercontent.com/AfaanBilal/whisper-app/master/screenshots/Notifications-Empty.webp) | ![Notifications List](https://raw.githubusercontent.com/AfaanBilal/whisper-app/master/screenshots/Notifications-List.webp) | ![Notifications Request](https://raw.githubusercontent.com/AfaanBilal/whisper-app/master/screenshots/Notifications-Request.webp) |
|             ![Profile](https://raw.githubusercontent.com/AfaanBilal/whisper-app/master/screenshots/Profile.webp)             |       ![Profile Menu](https://raw.githubusercontent.com/AfaanBilal/whisper-app/master/screenshots/Profile-Menu.webp)       |          ![Edit Profile](https://raw.githubusercontent.com/AfaanBilal/whisper-app/master/screenshots/EditProfile.webp)           |

---

## Run with Docker
```bash
docker run --env DB_DSN="root:@tcp(host.docker.internal:3306)/whisper?charset=utf8mb4&parseTime=True&loc=Local" --env PORT=8080 --env VERSION=0.1.0 -p 8080:8080 afaanbilal/whisper
```

This will start the server on port 8080.

## API
See [routes.go](./routes/routes.go) for routing.

## Contributing
All contributions are welcome. Please create an issue first for any feature request
or bug. Then fork the repository, create a branch and make any changes to fix the bug
or add the feature and create a pull request. That's it!
Thanks!

---

## License
**Whisper** is released under the MIT License.
Check out the full license [here](LICENSE).
