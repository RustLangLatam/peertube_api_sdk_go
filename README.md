# **Unofficial Go SDK for PeerTube API**

This is an **unofficial Go SDK** for interacting with the **PeerTube API**, a decentralized video hosting platform. It provides a simple way to interact with PeerTube's **RESTful API**, enabling authentication, video retrieval, and more.

📖 **Official PeerTube API Docs:** [REST API Quick Start](https://docs.joinpeertube.org/api/rest-getting-started)

---

## **🔹 Features**
- ✅ **OAuth2 Authentication** (password grant)
- ✅ Fetch **video details**
- ✅ Interact with **PeerTube users, channels, and videos**
- ✅ Implements **REST API** using **Go**

---

## **📌 Installation**
To use this SDK, install the required dependencies:

```sh
go get github.com/stretchr/testify/assert
go get golang.org/x/oauth2
go get golang.org/x/net/context
go get github.com/RustLangLatam/peertube_api_sdk
```

Import the package into your Go project:

```go
import peertube_api_sdk "github.com/RustLangLatam/peertube_api_sdk"
```

---

## **🔐 Authentication**
PeerTube uses **OAuth2 (password grant)** for authentication.  
Use the following method to obtain an **access token**:

```go
package main

import (
	"context"
	"fmt"
	"log"

	openapiclient "github.com/RustLangLatam/peertube_api_sdk"
)

func main() {
	apiClient := openapiclient.NewAPIClient(
		openapiclient.NewConfigurationFromBaseURL("https://your-peertube-instance.com"),
	)

	ctx := context.Background()
	oauthClient, _, err := apiClient.SessionAPI.GetOAuthClient(ctx).Execute()
	if err != nil {
		log.Fatal(err)
	}

	token, _, err := apiClient.SessionAPI.GetOAuthToken(ctx).
		ClientId(oauthClient.GetClientId()).
		ClientSecret(oauthClient.GetClientSecret()).
		Password("your-password").
		Username("your-username").
		GrantType("password").
		Execute()

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Access Token:", token.GetAccessToken())
}
```

More details can be found in **[docs/SessionAPI.md](docs/SessionAPI.md)**.

---

## **📺 Fetch Video Details**
To retrieve **video details** by **video ID**, use:

```go
package main

import (
	"context"
	"fmt"
	"log"

	openapiclient "github.com/RustLangLatam/peertube_api_sdk"
)

func main() {
	apiClient := openapiclient.NewAPIClient(
		openapiclient.NewConfigurationFromBaseURL("https://your-peertube-instance.com"),
	)

	videoID := openapiclient.Int32AsApiV1VideosOwnershipIdAcceptPostIdParameter(
		openapiclient.PtrInt32(29),
	)

	video, _, err := apiClient.VideoAPI.GetVideo(context.Background(), videoID).Execute()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Video Title:", video.GetName())
}
```

More examples can be found in the **[examples](examples/)** folder.

---

## **🌎 Proxy Support**
To use a proxy, set the `HTTP_PROXY` environment variable:

```go
os.Setenv("HTTP_PROXY", "http://proxy_name:proxy_port")
```

---

## **📖 API Documentation**
This SDK provides wrappers for **various PeerTube API endpoints**, including:
- **Videos** (`GetVideo`, `UploadVideo`, `DeleteVideo`)
- **Users** (`GetUser`, `RegisterUser`)
- **Playlists** (`GetPlaylists`, `AddPlaylist`)
- **Subscriptions** (`GetSubscriptions`, `SubscribeChannel`)

Detailed API documentation is available in:  
📖 **[docs/api_reference.md](docs/api_reference.md)**

For a full API reference, check the official **PeerTube API docs**:  
🔗 **[API Reference](https://docs.joinpeertube.org/api/rest-reference.html)**

---

## **🛠️ Contributing**
This is an **unofficial** SDK. Contributions and improvements are welcome!  
Feel free to **submit issues** or **pull requests**.

---

## **📜 License**
This project is **open-source** and distributed under the **MIT License**.

🚀 **Happy coding with PeerTube in Go!** 🎥
