<div align="center">
<br />
<p align="center">
<a href="https://wapijs.co"><img src="https://media.discordapp.net/attachments/907937769014325288/1248544029705240629/image.png?ex=66640cbd&is=6662bb3d&hm=9339e051f865880d2d8bfa7d04288f0eaeab042614aca4dd127e210377aeecb9&=&format=webp&quality=lossless&width=2261&height=1034" alt="@wapijs/Wapi.go"  height="200" width="360" /></a>
</p>
<br />
</div>

## 📌 Status

Beta Version - This SDK is not stable right now. It is currently in beta version. Report issues [here](https://github.com/sarthakjdev/wapi.go/issues).

## 📖 About

Wapi.go is a JavaScript module, written in TypeScript, designed to interact with the WhatsApp cloud API in a user-friendly manner.

## ✨ Features

- Single Client Model
- Send Messages with the least configuration
- Event Listener for Notifications (support both User and System Notifications)
- Upload Media to WhatsApp servers
- Reply and React to incoming messages.

## 💻 Installation

This assumes you already have a working Go environment, if not please see
[this page](https://golang.org/doc/install) first.

`go get` _will always pull the latest tagged release from the master branch._

```sh
go get github.com/sarthakjdev/wapi.go
```

> Note: This SDK is not affiliated with the official WhatsApp Cloud API or does not act as any official solution provided the the Meta Inclusive Private Limited, this is just a open source SDK built for developers to support them in building whatsapp cloud api based chat bots easily.

## 🚀 Usage

You can check out the example WhatsApp bot here. [Example Chatbot](./example-chat-bot/)

# Example Usage

Import the package into your project.
This repository has three packages exported:

- github.com/sarthakjdev/wapi.go/components
- github.com/sarthakjdev/wapi.go/wapi/wapi
- github.com/sarthakjdev/wapi.go/wapi/events

```go
import "github.com/sarthakjdev/wapi.go/wapi/wapi"
```

Construct a new Wapi Client to access the managers in order to send messages and listen to incoming notifications.

```go
whatsappClient, err := wapi.New(wapi.ClientConfig{
		PhoneNumberId:     "",
		ApiAccessToken: "",
		BusinessAccountId: "",
		WebhookPath:       "/webhook",
		WebhookSecret:     "",
		WebhookServerPort: 8080,
	})
```

## 🔗 References

- **Message Structures**: Refer to the WhatsApp Docs [here](https://developers.facebook.com/docs/whatsapp/cloud-api/reference/messages).

- **Notification Payloads**: Details can be found [here](https://developers.facebook.com/docs/whatsapp/cloud-api/webhooks/components).

<!-- ## 🔗 Other Links

- [Website](https://wapijs.co)
- [Documentation](https://wapijs.co/docs) -->

## 🤝 Contribution Guidelines

1. Fork the Project
2. Create your Feature Branch (`git checkout -b feature/AmazingFeature`)
3. Commit changes (`git commit -m 'Add some AmazingFeature'`)
4. Push to the Branch (`git push origin feature/AmazingFeature`)
5. Open a Pull Request

For detailed guidelines, check [Contributing.md](./CONTRIBUTING.md).

## 📜 License

Distributed under the Apache 2.0 License. View [LICENSE](./LICENSE).

## 📞 Contact

- [Sarthak Jain](https://sarthakjdev.com)
- Email: sarthak@softlancer.co
- [Twitter](https://twitter.com/sarthakjdev) | [LinkedIn](https://www.linkedin.com/in/sarthakjdev)

Note: This SDK is part of an open-source product-building initiative by [Softlancer](https://github.com/softlancerhq), and this repository will soon be moved under the same organization.
