![](https://github.com/kianashnad/ServerDoctor/banner.jpg)

# Server Doctor
it's just a simple go program to check the URLs you provide. 

## Usage
1. Clone the repo.
2. Create a telegram bot using the `@BotFather` Bot and get the bot token.
3. Get your account UUID (you can use a bot like `@userinfobot`).
4. prepare the .env file and place the telegram bot token, telegram UUID, and the URLS.
5. make sure that you seprated the URLs by a comma (`,`) and there's no space. They should also contain `https://`.
6. run the program!
---
## Important points
- You can deploy it via docker and docker and docker compose as well.
- In the `docker-compose.yml` file, the resource usage is limited to these:
   ```
  cpus: '0.060'
  memory: 50M
  ```
  change them if you desire to.

- Feel free to improve the project with a PR
- You're welcomed to email me if I've had a mistake.
---
## License
```
Copyright 2022 Kian Ashnad

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
```
Plus, please don't use it in any way that harms humans, animals, nature. Use it in the way of goodness and love❤. We need more of that. ️